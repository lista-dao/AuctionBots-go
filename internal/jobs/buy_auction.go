package jobs

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	daov1 "github.com/lista-dao/AuctionBots-go/internal/dao/v1"
	daov2 "github.com/lista-dao/AuctionBots-go/internal/dao/v2"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/clipper"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/interaction"
	"github.com/lista-dao/AuctionBots-go/internal/wallet"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
	"time"
)

func NewBuyAuctionJob(
	ctx context.Context,
	log *logrus.Logger,
	wall wallet.Walleter,
	ethCli *ethclient.Client,
	interactAddr common.Address,
	collateralAddr common.Address,
	hayAddr common.Address,
	maxPricePerc *big.Int,
	withWait bool,
) Job {
	job := &buyAuctionJob{
		ctx:            ctx,
		ethCli:         ethCli,
		interactAddr:   interactAddr,
		collateralAddr: collateralAddr,
		wallet:         wall,
		log: log.WithFields(logrus.Fields{
			"job":        "buy_auction",
			"collateral": collateralAddr.String(),
			"operator":   wall.Address(),
		}),
		hayAddr:      hayAddr,
		withWait:     withWait,
		maxPricePerc: maxPricePerc,
	}

	return job
}

var _ Job = (*buyAuctionJob)(nil)

type buyAuctionJob struct {
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
	interAbi       *abi.ABI
	spotAddr       common.Address
	collateralIlk  [32]byte

	inter   *interaction.Interaction
	clipper *clipper.Clipper
	oracle  *daov1.MockOracle
	hay     *daov2.Token

	withWait bool
}

func (j *buyAuctionJob) init() {
	if j == nil {
		panic("buy auction job is null")
	}
	var err error

	j.inter, err = interaction.NewInteraction(j.interactAddr, j.ethCli)
	if err != nil {
		panic(err)
	}

	j.interAbi, err = interaction.InteractionMetaData.GetAbi()
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

func (j *buyAuctionJob) Run(ctx context.Context) {
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
		auctionPrice.Cmp(calcPercent(actualPrice, j.maxPricePerc)) >= 0 {
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

	hayMax := big.NewInt(0).Div(big.NewInt(0).Mul(status.Price, collatAmount), RAY)
	log = log.WithFields(logrus.Fields{
		"collat_to_buy": collatAmount.String(),
		"hay_max":       hayMax.String(),
	})

	log.Debug("buying auction...")
	j.buyAuction(log, auctionID, collatAmount, status.Price)
	log.Debug("auction bought")
}

func (j *buyAuctionJob) buyAuction(log *logrus.Entry, auctionID *big.Int, collatAmount, maxPrice *big.Int) {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		log.WithError(err).Error("failed to get tx opts")
		return
	}

	ctx := context.Background()
	input, err := j.interAbi.Pack(
		"buyFromAuction",
		j.collateralAddr,
		auctionID,
		collatAmount,
		maxPrice,
		opts.From,
	)
	if err != nil {
		log.WithError(err).Error("j.interAbi.Pack")
		return
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
		log.WithError(err).Error("j.ethCli.EstimateGas")
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
		return errors.Wrap(err, " j.hay.Approv")
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
