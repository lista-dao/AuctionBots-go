package jobs

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	daov2 "github.com/helio-money/auctionbot/internal/dao/v2"
	"github.com/helio-money/auctionbot/internal/wallet"
	"github.com/pkg/errors"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"math/big"
)

func NewResetAuctionJob(
	ctx context.Context,
	log *logrus.Logger,
	wall wallet.Walleter,
	ethCli *ethclient.Client,
	interactAddr common.Address,
	collateralAddr common.Address,
	withWait bool,
) cron.Job {
	job := &resetJob{
		ctx:            ctx,
		ethCli:         ethCli,
		interactAddr:   interactAddr,
		collateralAddr: collateralAddr,
		wallet:         wall,
		log: log.WithFields(logrus.Fields{
			"job":      "buy_auction",
			"operator": wall.Address(),
		}),
		withWait: withWait,
	}

	return job
}

var _ cron.Job = (*resetJob)(nil)

type resetJob struct {
	ctx context.Context

	wallet         wallet.Walleter
	ethCli         *ethclient.Client
	log            *logrus.Entry
	collateralAddr common.Address
	interactAddr   common.Address

	inter   *daov2.Interaction
	clipper *daov2.Clipper

	withWait bool
}

func (j *resetJob) init() {
	if j == nil {
		panic("buy auction job is null")
	}
	var err error

	j.inter, err = daov2.NewInteraction(j.interactAddr, j.ethCli)
	if err != nil {
		panic(err)
	}

	// get collateral details to retrieve clip and ilk
	collatDetails, err := j.inter.Collaterals(&bind.CallOpts{}, j.collateralAddr)
	if err != nil {
		panic(err)
	}

	j.clipper, err = daov2.NewClipper(collatDetails.Clip, j.ethCli)
	if err != nil {
		panic(err)
	}
}

func (j *resetJob) Run() {
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

func (j *resetJob) processAuction(auctionID *big.Int) {
	log := j.log.WithFields(logrus.Fields{
		"auction_id": auctionID.String(),
	})

	status, err := j.clipper.GetStatus(&bind.CallOpts{}, auctionID)
	if err != nil {
		log.WithError(err).Error("failed to get status from clipper")
		return
	}

	log = log.WithFields(logrus.Fields{
		"needs_redo": status.NeedsRedo,
	})

	if status.NeedsRedo {
		log.Debug("auction resetting...")
		if err := j.redoAuction(auctionID); err != nil {
			log.WithError(err).Error("failed to redo auction")
		}
		log.Debug("auction reset")
	}
}

func (j *resetJob) redoAuction(auctionID *big.Int) error {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx opts")
	}

	tx, err := j.inter.ResetAuction(
		opts,
		j.collateralAddr,
		auctionID,
		opts.From,
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
