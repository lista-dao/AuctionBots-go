package jobs

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	daov1 "github.com/helio-money/auctionbot/internal/dao/v1"
	"github.com/helio-money/auctionbot/internal/dao/v2/clipper"
	"github.com/helio-money/auctionbot/internal/dao/v2/flash/flashbuy"
	"github.com/helio-money/auctionbot/internal/dao/v2/flash/interfaces/ierc3156flashlender"
	"sync"
	"time"

	"github.com/helio-money/auctionbot/internal/dao/v2/interaction"
	"github.com/helio-money/auctionbot/internal/wallet"
	"github.com/sirupsen/logrus"
	"math/big"
)

func NewBuyFlashAuctionJob(
	ctx context.Context,
	log *logrus.Logger,
	wall wallet.Walleter,
	ethCli *ethclient.Client,
	interactAddr common.Address,
	collateralAddr common.Address,
	hayAddr common.Address,
	flashLoanAddr common.Address,
	withWait bool,
) Job {
	return &buyFlashAuctionJob{
		ctx:            ctx,
		ethCli:         ethCli,
		interactAddr:   interactAddr,
		collateralAddr: collateralAddr,
		wallet:         wall,
		log: log.WithFields(logrus.Fields{
			"job":      "flash_buy_auction",
			"operator": wall.Address(),
		}),
		hayAddr:       hayAddr,
		withWait:      withWait,
		flashLoanAddr: flashLoanAddr,
	}
}

var _ Job = (*buyFlashAuctionJob)(nil)

type buyFlashAuctionJob struct {
	ctx context.Context

	wallet         wallet.Walleter
	ethCli         *ethclient.Client
	log            *logrus.Entry
	collateralAddr common.Address
	hayAddr        common.Address
	interactAddr   common.Address
	spotAddr       common.Address
	flashLoanAddr  common.Address
	collateralIlk  [32]byte

	flash  *flashbuy.Flashbuy
	lender *ierc3156flashlender.Ierc3156flashlender

	inter   *interaction.Interaction
	clipper *clipper.Clipper
	oracle  *daov1.MockOracle

	withWait bool
}

func (j *buyFlashAuctionJob) init() {
	if j == nil {
		panic("buy auction job is null")
	}
	var err error
	j.flash, _ = flashbuy.NewFlashbuy(j.flashLoanAddr, j.ethCli)
	j.inter, err = interaction.NewInteraction(j.interactAddr, j.ethCli)
	if err != nil {
		panic(err)
	}

	spotAddr, err := j.inter.Spotter(&bind.CallOpts{})
	if err != nil {
		panic(err)
	}

	// get collateral details to retrieve clip and ilk
	collatDetails, err := j.inter.Collaterals(&bind.CallOpts{}, j.collateralAddr)
	if err != nil {
		panic(err)
	}
	j.collateralIlk = collatDetails.Ilk

	spot, err := daov1.NewSpot(spotAddr, j.ethCli)
	if err != nil {
		panic(err)
	}

	j.clipper, err = clipper.NewClipper(collatDetails.Clip, j.ethCli)
	if err != nil {
		panic(err)
	}

	spotIlk, err := spot.Ilks(&bind.CallOpts{}, collatDetails.Ilk)
	if err != nil {
		panic(err)
	}

	j.oracle, err = daov1.NewMockOracle(spotIlk.Pip, j.ethCli)
	if err != nil {
		panic(err)
	}
}

func (j *buyFlashAuctionJob) Run(ctx context.Context, wg *sync.WaitGroup) {
	j.init()
	ticker := time.NewTicker(time.Minute)
	go func() {
		j.log.Debug("start")

		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				auctionIds, err := j.clipper.List(&bind.CallOpts{})
				if err != nil {
					j.log.WithError(err).Error("failed to list auction ids from clipper")
					continue
				}

				if len(auctionIds) == 0 {
					j.log.Debug("nothing to buy")
					continue
				}

				for _, auctionID := range auctionIds {
					j.processAuction(auctionID)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (j *buyFlashAuctionJob) processAuction(auctionID *big.Int) {
	log := j.log.WithFields(logrus.Fields{
		"auction_id": auctionID.String(),
	})

	status, err := j.clipper.GetStatus(&bind.CallOpts{}, auctionID)
	if err != nil {
		log.WithError(err).Error("failed to get status from clipper")
		return
	}

	auctionPrice := big.NewInt(1).Div(status.Price, BLN)

	log = log.WithFields(logrus.Fields{
		"auction_price": auctionPrice,
		"needs_redo":    status.NeedsRedo,
		"lot":           status.Lot.String(),
	})

	// get collateral price
	bytesPrice, _, err := j.oracle.Peek(&bind.CallOpts{})
	if err != nil {
		log.WithError(err).Error("failed to pick price from oracle")
		return
	}
	actualPrice := big.NewInt(0).SetBytes(bytesPrice[:])

	log = log.WithField("actual_price", actualPrice.String())

	// check if auction lot is 0 or auction needs redo or
	// auction collateral price is higher than actual
	if status.NeedsRedo ||
		status.Lot.Cmp(big.NewInt(0)) < 0 {
		//auctionPrice.Cmp(calcPercent(actualPrice, big.NewInt(95))) >= 0
		log.Debug("auctions is skipped")
		return
	}

	collatAmount := status.Lot

	hayMax := big.NewInt(0).Div(big.NewInt(0).Mul(status.Price, status.Lot), RAY)
	log = log.WithFields(logrus.Fields{
		"collat_to_buy": collatAmount.String(),
		"hay_max":       hayMax.String(),
	})

	log.Debug("flash buying auction...")
	j.flashBuyAuction(log, auctionID, collatAmount, hayMax, status.Price)
	log.Debug("flash auction bought")
}

func (j *buyFlashAuctionJob) flashBuyAuction(log *logrus.Entry, auctionID *big.Int, collatAmount, borrowAmount, maxPrice *big.Int) {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		log.WithError(err).Error("failed to get tx opts")
		return
	}

	tx, err := j.flash.FlashBuyAuction(
		opts,
		j.hayAddr,
		auctionID,
		borrowAmount.Add(borrowAmount, big.NewInt(100)),
		j.collateralAddr,
		collatAmount,
		maxPrice,
	)
	if err != nil {
		log.WithError(err).Error("failed to send buy tx")
		return
	}
	log = log.WithField("tx_hash", tx.Hash().String())

	if j.withWait {
		receipt, err := bind.WaitMined(j.ctx, j.ethCli, tx)
		if err != nil {
			log.WithError(err).Error("failed to wait for mint")
			return
		}

		if receipt.Status == types.ReceiptStatusFailed {
			if _, err := getRevertReason(j.ctx, j.ethCli, tx, receipt); err != nil {
				log.WithError(err).Error("tx failed")
				return
			}
		}
	}

	return
}
