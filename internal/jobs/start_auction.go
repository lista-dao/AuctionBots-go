package jobs

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	analyticsv1 "github.com/helio-money/auctionbot/internal/analytics/v1"
	daov2 "github.com/helio-money/auctionbot/internal/dao/v2/interaction"
	"github.com/helio-money/auctionbot/internal/wallet"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func NewStartAuctionJob(
	ctx context.Context,
	log *logrus.Logger,
	wall wallet.Walleter,
	analyticsCli *analyticsv1.Client,
	ethCli *ethclient.Client,
	interactAddr common.Address,
	collateralAddr common.Address,
	withWait bool,
) Job {
	return &startAuctionJob{
		ctx:            ctx,
		ethCli:         ethCli,
		analyticsCli:   analyticsCli,
		interactAddr:   interactAddr,
		collateralAddr: collateralAddr,
		wallet:         wall,
		log:            log.WithField("job", "start_auction"),
		withWait:       withWait,
	}
}

var _ Job = (*startAuctionJob)(nil)

type startAuctionJob struct {
	ctx context.Context

	wallet         wallet.Walleter
	analyticsCli   *analyticsv1.Client
	ethCli         *ethclient.Client
	log            *logrus.Entry
	interactAddr   common.Address
	collateralAddr common.Address

	withWait bool
}

func (j *startAuctionJob) Run(ctx context.Context, wg *sync.WaitGroup) {
	j.log.Debug("start")
	inter, err := daov2.NewInteraction(j.interactAddr, j.ethCli)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Minute)
	go func() {
		j.log.Debug("start")

		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				users, err := j.analyticsCli.GetRedUsers(j.ctx)
				if err != nil {
					j.log.WithError(err).Error("failed to get red users")
					return
				}

				for ind := range users {
					log := j.log.WithFields(logrus.Fields{
						"user_address": users[ind].UserAddress.String(),
					})

					//TODO: add liquidation neediness checking
					//TODO: add profitability checking

					log.Debug("starting auction...")
					if err := j.startAuction(inter, users[ind]); err != nil {
						log.WithError(err).Error("failed to start auction")
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (j *startAuctionJob) startAuction(inter *daov2.Interaction, user analyticsv1.User) error {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx opts")
	}

	tx, err := inter.StartAuction(
		opts,
		j.collateralAddr,
		user.UserAddress,
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
