package jobs

import (
	"context"
	"github.com/dgraph-io/ristretto"
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
	"math/big"
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
	cache          *ristretto.Cache

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
	j.cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
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
					continue
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
	logrus.Infof("start auction for user: %s collateral: %s", user.UserAddress.String(), j.collateralAddr.String())

	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get tx opts")
	}

	_, ok := j.cache.Get("poke")
	if !ok {
		// didn't poke for one minute
		tx, err := j.inter.Poke(opts, j.collateralAddr)
		if err != nil {
			return errors.Wrap(err, "j.inter.Poke")
		}

		j.cache.SetWithTTL("poke", true, 1, 3*time.Minute)

		logrus.Infof("poke tx %s", tx.Hash().String())
		opts.Nonce = opts.Nonce.Add(opts.Nonce, big.NewInt(1))
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

	logrus.Infof("start auction tx %s", tx.Hash().String())
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
