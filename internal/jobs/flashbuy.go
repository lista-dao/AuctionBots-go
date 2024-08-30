package jobs

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	daov1 "github.com/lista-dao/AuctionBots-go/internal/dao/v1"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/clipper"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/flash/flashbuy"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/flash/interfaces/ierc3156flashlender"
	"github.com/lista-dao/AuctionBots-go/pkg/config"
	"github.com/lista-dao/AuctionBots-go/pkg/utils"
	"strings"
	"time"

	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/interaction"
	"github.com/lista-dao/AuctionBots-go/internal/wallet"
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
	flashBuyAddr common.Address,
	maxPricePerc *big.Int,
	withWait bool,
	cfg *config.Config,
) Job {
	return &buyFlashAuctionJob{
		ctx:            ctx,
		ethCli:         ethCli,
		interactAddr:   interactAddr,
		collateralAddr: collateralAddr,
		wallet:         wall,
		log: log.WithFields(logrus.Fields{
			"job":        "flash_buy_auction",
			"collateral": collateralAddr.String(),
			"operator":   wall.Address(),
		}),
		hayAddr:      hayAddr,
		withWait:     withWait,
		flashBuyAddr: flashBuyAddr,
		//maxPricePerc: maxPricePerc,
		maxPricePerc: big.NewInt(88),
		cfg:          cfg,
	}
}

var _ Job = (*buyFlashAuctionJob)(nil)

type buyFlashAuctionJob struct {
	ctx context.Context

	// job will buy action when
	// auction_price < actual_price * percentage (1-100%)
	maxPricePerc *big.Int

	wallet         wallet.Walleter
	ethCli         *ethclient.Client
	log            *logrus.Entry
	collateralAddr common.Address
	hayAddr        common.Address
	interactAddr   common.Address
	spotAddr       common.Address
	flashBuyAddr   common.Address
	collateralIlk  [32]byte

	flash    *flashbuy.Flashbuy
	flashAbi *abi.ABI
	lender   *ierc3156flashlender.Ierc3156flashlender

	inter   *interaction.Interaction
	clipper *clipper.Clipper
	oracle  *daov1.MockOracle

	withWait bool
	cfg      *config.Config
}

func (j *buyFlashAuctionJob) init() {
	if j == nil {
		panic("buy auction job is null")
	}
	var err error
	j.flash, _ = flashbuy.NewFlashbuy(j.flashBuyAddr, j.ethCli)
	j.flashAbi, err = flashbuy.FlashbuyMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
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
	j.log.Debugf("config addr:%v value:%v", j.cfg.Contract.FlushBuy, j.cfg.FlushBuy)
}

func (j *buyFlashAuctionJob) Run(ctx context.Context) {
	j.init()
	ticker := time.NewTicker(time.Minute)
	go func() {
		j.log.Debug("start")

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
		status.Lot.Cmp(big.NewInt(0)) < 0 ||
		auctionPrice.Cmp(calcPercent(actualPrice, j.maxPricePerc)) >= 0 {
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
	if hayMax.Cmp(AUCTION_CAP) >= 0 {
		collatAmount = big.NewInt(0).Div(big.NewInt(0).Mul(AUCTION_CAP, RAY), status.Price)
		collatAmount.Sub(collatAmount, big.NewInt(100))
		log.Infof("FlashBuyAuction : split auction cap:%v price:%v amt:%v", AUCTION_CAP, auctionPrice, collatAmount)
	}
	j.flashBuyAuction(log, auctionID, collatAmount, hayMax, status.Price)
	log.Debug("flash auction bought")
}

func (j *buyFlashAuctionJob) flashBuyAuction(log *logrus.Entry, auctionID *big.Int, collatAmount, borrowAmount, maxPrice *big.Int) {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		log.WithError(err).Error("failed to get tx opts")
		return
	}

	slip := big.NewInt(j.cfg.FlushBuy.Slip)
	fb := j.cfg.FlushBuy.Paths[strings.ToLower(j.collateralAddr.Hex())]
	collateralReal := common.HexToAddress(fb.Received)
	path := utils.PathToBytes(fb.Tokens, fb.Fees)

	ctx := context.Background()
	input, err := j.flashAbi.Pack(
		"flashBuyAuction",
		j.hayAddr,
		auctionID,
		borrowAmount.Add(borrowAmount, big.NewInt(100)),
		j.collateralAddr,
		collatAmount,
		maxPrice,
		slip,
		collateralReal,
		path,
	)
	log.Debugf("FlashBuyAuction params: %v %v %v %v %v %v %v %v %v %v %v", j.hayAddr, auctionID, borrowAmount, j.collateralAddr, collatAmount, maxPrice, slip, collateralReal, hex.EncodeToString(path), fb.Tokens, fb.Fees)

	if err != nil {
		log.WithError(err).Error("j.interAbi.Pack")
		return
	}
	_, err = j.ethCli.EstimateGas(ctx, ethereum.CallMsg{
		From:     j.wallet.Address(),
		To:       &j.flashBuyAddr,
		Gas:      opts.GasLimit,
		GasPrice: opts.GasPrice,
		Value:    opts.Value,
		Data:     input,
	})
	if err != nil {
		log.WithError(err).Error("j.ethCli.EstimateGas")
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
		slip,
		collateralReal,
		path,
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
