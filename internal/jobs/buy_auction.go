package jobs

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	daov1 "github.com/helio-money/auctionbot/internal/dao/v1"
	daov2 "github.com/helio-money/auctionbot/internal/dao/v2"
	"github.com/helio-money/auctionbot/internal/wallet"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"math/big"
)

func NewBuyAuctionJob(
	ctx context.Context,
	log *logrus.Logger,
	wall wallet.Walleter,
	ethCli *ethclient.Client,
	interactAddr common.Address,
	collateralAddr common.Address,
	hayAddr common.Address,
	withWait bool,
) cron.Job {
	job := &buyAuctionJob{
		ctx:            ctx,
		ethCli:         ethCli,
		interactAddr:   interactAddr,
		collateralAddr: collateralAddr,
		wallet:         wall,
		log: log.WithFields(logrus.Fields{
			"job":      "buy_auction",
			"operator": wall.Address(),
		}),
		hayAddr:  hayAddr,
		withWait: withWait,
	}

	return job
}

var _ cron.Job = (*buyAuctionJob)(nil)

type buyAuctionJob struct {
	ctx context.Context

	wallet         wallet.Walleter
	ethCli         *ethclient.Client
	log            *logrus.Entry
	collateralAddr common.Address
	hayAddr        common.Address
	interactAddr   common.Address
	spotAddr       common.Address
	collateralIlk  [32]byte

	inter   *daov2.Interaction
	clipper *daov2.Clipper
	oracle  *daov1.MockOracle
	hay     *daov2.Token

	withWait bool
}

func (j *buyAuctionJob) init() {
	if j == nil {
		panic("buy auction job is null")
	}
	var err error

	j.inter, err = daov2.NewInteraction(j.interactAddr, j.ethCli)
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

	j.clipper, err = daov2.NewClipper(collatDetails.Clip, j.ethCli)
	if err != nil {
		panic(err)
	}

	j.hay, err = daov2.NewToken(j.hayAddr, j.ethCli)
	if err != nil {
		panic(err)
	}

	allowance, err := j.hay.Allowance(&bind.CallOpts{}, j.wallet.Address(), j.interactAddr)
	if err != nil {
		panic(err)
	}

	if allowance.Cmp(abi.MaxUint256) != 0 {
		j.log.Debug("approving hay to interactor")
		if err := j.approveTokens(); err != nil {
			panic(err)
		}
		j.log.Debug("hay approved")
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

func (j *buyAuctionJob) Run() {
	j.log.Debug("start")

	j.init()

	auctionIds, err := j.clipper.List(&bind.CallOpts{})
	if err != nil {
		j.log.WithError(err).Error("failed to list auction ids from clipper")
	}

	for _, auctionID := range auctionIds {
		j.processAuction(auctionID)
	}
}

func (j *buyAuctionJob) processAuction(auctionID *big.Int) {
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
		auctionPrice.Cmp(calcPercent(actualPrice, big.NewInt(95))) >= 0 {
		log.Debug("auctions is skipped")
		return
	}

	balance, err := j.hay.BalanceOf(&bind.CallOpts{}, j.wallet.Address())
	if err != nil {
		log.WithError(err).Error("failed to get balance from hay")
		return
	}

	collatAmount := big.NewInt(0).Div(big.NewInt(0).Mul(balance, status.Price), WAD)
	if collatAmount.Cmp(status.Lot) > 0 {
		collatAmount = status.Lot
	}

	j.buyAuction(log, auctionID, collatAmount, status.Price)
}

func (j *buyAuctionJob) buyAuction(log *logrus.Entry, auctionID *big.Int, collatAmount, maxPrice *big.Int) {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		log.WithError(err).Error("failed to get tx opts")
		return
	}

	tx, err := j.inter.BuyFromAuction(
		opts,
		j.collateralAddr,
		auctionID,
		collatAmount,
		maxPrice,
		opts.From,
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
			log.Debug("auction bought")
		}
	}

	return
}

func (j *buyAuctionJob) approveTokens() error {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx opts")
	}

	tx, err := j.hay.Approve(
		opts,
		j.interactAddr,
		abi.MaxUint256,
	)
	if err != nil {
		return errors.Wrap(err, "failed to send tx")
	}

	if j.withWait {
		receipt, err := bind.WaitMined(j.ctx, j.ethCli, tx)
		if err != nil {
			return errors.Wrapf(err, "failed to wait for tx %s mint", tx.Hash())
		}

		if receipt.Status == types.ReceiptStatusFailed {
			_, err := getRevertReason(j.ctx, j.ethCli, tx, receipt)
			return errors.Wrapf(err, "failed to wait for tx %s mint", tx.Hash())
		}
	}

	return nil
}
