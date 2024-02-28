package jobs

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	analyticsv1 "github.com/lista-dao/AuctionBots-go/internal/analytics/v1"
	daov2 "github.com/lista-dao/AuctionBots-go/internal/dao/v2/interaction"
	"github.com/lista-dao/AuctionBots-go/internal/wallet"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
		log: log.WithFields(
			logrus.Fields{
				"job":        "start_auction",
				"collateral": collateralAddr.String(),
				"operator":   wall.Address(),
			},
		),
		withWait: withWait,
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
	inter          *daov2.Interaction
	interAbi       *abi.ABI
	collateralAddr common.Address

	withWait bool
}

func (j *startAuctionJob) Run(ctx context.Context) {
	var err error
	j.inter, err = daov2.NewInteraction(j.interactAddr, j.ethCli)
	if err != nil {
		panic(err)
	}
	j.interAbi, err = daov2.InteractionMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	ticker := time.NewTicker(time.Minute)
	go func() {
		j.log.Debug("start")

		for {
			select {
			case <-ticker.C:
				users, err := j.analyticsCli.GetRedUsers(j.ctx)
				if err != nil {
					j.log.WithError(err).Error("failed to get red users")
					return
				}

				if len(users) == 0 {
					j.log.Debug("nothing to start")
					continue
				}

				for ind := range users {
					log := j.log.WithFields(logrus.Fields{
						"user_address": users[ind].UserAddress.String(),
					})

					//TODO: add liquidation neediness checking
					//TODO: add profitability checking

					log.Debug("starting auction...")
					if err := j.startAuction(users[ind]); err != nil {
						log.WithError(err).Error("failed to start auction")
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (j *startAuctionJob) startAuction(user analyticsv1.User) error {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx opts")
	}

	ctx := context.Background()
	input, err := j.interAbi.Pack(
		"startAuction",
		j.collateralAddr,
		user.UserAddress,
		opts.From,
	)
	if err != nil {
		return errors.Wrap(err, "j.interAbi.Pack")
	}
	_, err = j.ethCli.EstimateGas(ctx, ethereum.CallMsg{
		From:     j.wallet.Address(),
		To:       &j.interactAddr,
		Gas:      opts.GasLimit,
		GasPrice: opts.GasPrice,
		Value:    opts.Value,
		Data:     input,
	})
	if err != nil {
		return errors.Wrap(err, "j.ethCli.EstimateGas")
	}

	tx, err := j.inter.StartAuction(
		opts,
		j.collateralAddr,
		user.UserAddress,
		opts.From,
	)
	if err != nil {
		return errors.Wrap(err, "j.inter.StartAuction")
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
