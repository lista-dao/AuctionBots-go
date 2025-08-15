package jobs

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	analyticsv3 "github.com/lista-dao/AuctionBots-go/internal/analytics/v3"
	daov1 "github.com/lista-dao/AuctionBots-go/internal/dao/v1"
	daov2 "github.com/lista-dao/AuctionBots-go/internal/dao/v2"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/clipper"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v3/interaction"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v3/oracle"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v3/provider"
	"github.com/lista-dao/AuctionBots-go/internal/wallet"
	"github.com/lista-dao/AuctionBots-go/pkg/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"strings"
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
	analyticsCli *analyticsv3.Client,
	multiOracleAddr common.Address,
	cfg *config.Config,
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
		hayAddr:  hayAddr,
		withWait: withWait,
		//maxPricePerc: maxPricePerc,
		maxPricePerc:    big.NewInt(90),
		analyticsCli:    analyticsCli,
		multiOracleAddr: multiOracleAddr,
		cfg:             cfg,
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

	inter       *interaction.Interaction
	pcsProvider *provider.Provider
	multiOracle *oracle.Oracle
	clipper     *clipper.Clipper
	oracle      *daov1.MockOracle
	hay         *daov2.Token

	withWait        bool
	lp              bool
	analyticsCli    *analyticsv3.Client
	pcsProviderAbi  *abi.ABI
	multiOracleAddr common.Address
	multiOracleAbi  *abi.ABI
	cfg             *config.Config
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
	j.multiOracle, err = oracle.NewOracle(j.multiOracleAddr, j.ethCli)
	if err != nil {
		panic(err)
	}

	j.interAbi, err = interaction.InteractionMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	j.pcsProviderAbi, err = provider.ProviderMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	j.multiOracleAbi, err = oracle.OracleMetaData.GetAbi()
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
				Monitor.Beat("buyAuctionJob")
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
	if status.Price == nil || len(status.Price.Bits()) == 0 {
		log.WithError(err).Error("failed to get status.Price")
		return
	}

	collatHay := big.NewInt(0).Div(big.NewInt(0).Mul(status.Price, status.Lot), RAY)
	if collatHay.Cmp(AUTIION_MAX_CAP) >= 0 {
		log.Infof("buy_auction : ignoring users with excessive max auction amounts...")
		return
	}

	// balance * RAY / price
	collatAmount := big.NewInt(0).Div(big.NewInt(0).Mul(balance, RAY), status.Price)
	if collatAmount.Cmp(status.Lot) > 0 {
		collatAmount = status.Lot
	}

	hayMax := big.NewInt(0).Div(big.NewInt(0).Mul(status.Price, collatAmount), RAY)
	log = log.WithFields(logrus.Fields{
		"collat_to_buy": collatAmount.String(),
		"hay_max":       hayMax.String(),
	})

	if hayMax.Cmp(AUCTION_CAP) >= 0 {
		log.Infof("buy_auction : ignoring users with excessive auction amounts...")
		collatAmount = big.NewInt(0).Div(big.NewInt(0).Mul(AUCTION_CAP, RAY), status.Price)
		log.Infof("buy_auction : split auction %v", collatAmount)
	}

	log.Debug("buying auction...")
	j.buyAuction(log, auctionID, collatAmount, status.Price)
	log.Debug("auction bought")
}
func parseType(typeStr string) abi.Type {
	t, err := abi.NewType(typeStr, "", nil)
	if err != nil {
		log.Fatalf("Failed to parse type '%s': %v", typeStr, err)
	}
	return t
}

var multicall3Address = common.HexToAddress("0x1Ee38d535d541c55C9dae27B12edf090C608E6Fb")

const multicall3ABIJson = `[
  {"inputs":[{"components":[{"internalType":"address","name":"target","type":"address"},{"internalType":"bool","name":"allowFailure","type":"bool"},{"internalType":"bytes","name":"callData","type":"bytes"}],"internalType":"struct Multicall3.Call3[]","name":"calls","type":"tuple[]"}],"name":"aggregate3","outputs":[{"components":[{"internalType":"bool","name":"success","type":"bool"},{"internalType":"bytes","name":"returnData","type":"bytes"}],"internalType":"struct Multicall3.Result[]","name":"returnData","type":"tuple[]"}],"stateMutability":"payable","type":"function"}
]`

type mcCall struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

type mcResult struct {
	Success    bool
	ReturnData []byte
}

func doMulticall3(ctx context.Context, client bind.ContractBackend, calls []mcCall) ([]mcResult, error) {
	mcABI, err := abi.JSON(strings.NewReader(multicall3ABIJson))
	if err != nil {
		return nil, err
	}
	// pack calls
	type call3 struct {
		Target       common.Address
		AllowFailure bool
		CallData     []byte
	}
	in := make([]call3, 0, len(calls))
	for _, c := range calls {
		in = append(in, call3{Target: c.Target, AllowFailure: c.AllowFailure, CallData: c.CallData})
	}
	input, err := mcABI.Pack("aggregate3", in)
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &multicall3Address,
		Data: input,
	}
	raw, err := client.(bind.ContractCaller).CallContract(ctx, msg, nil)
	if err != nil {
		return nil, err
	}

	// decode results
	var outs []struct {
		Success    bool
		ReturnData []byte
	}
	if err := mcABI.UnpackIntoInterface(&outs, "aggregate3", raw); err != nil {
		return nil, err
	}
	res := make([]mcResult, len(outs))
	for i, o := range outs {
		res[i] = mcResult{Success: o.Success, ReturnData: o.ReturnData}
	}
	return res, nil
}

func (j *buyAuctionJob) buyAuction(log *logrus.Entry, auctionID *big.Int, collatAmount, maxPrice *big.Int) {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		log.WithError(err).Error("failed to get tx opts")
		return
	}

	ctx := context.Background()
	user, err := j.analyticsCli.GetUserByAuctionId(ctx, auctionID.String(), j.collateralAddr.Hex())
	if err != nil {
		j.log.WithError(err).Error("failed to GetUserByAuctionId")
		return
	}
	realLpProviderAddr := ""
	if v, ok := j.cfg.Contract.LPProviders[user.ClipperAddress.Hex()]; ok {
		realLpProviderAddr = v
	}
	data := []byte{}
	if realLpProviderAddr != "" {
		j.pcsProvider, err = provider.NewProvider(common.HexToAddress(realLpProviderAddr), j.ethCli)
		if err != nil {
			j.log.WithError(err).Error("failed to NewProvider", realLpProviderAddr)
			return
		}
		args := abi.Arguments{
			{Type: parseType("uint256")},
			{Type: parseType("uint256")},
			{Type: parseType("uint256")},
		}
		callGetUserLps, _ := j.pcsProviderAbi.Pack("getUserLps", user.UserAddress)
		callToken0, _ := j.pcsProviderAbi.Pack("token0")
		callToken1, _ := j.pcsProviderAbi.Pack("token1")
		callUserLiq, _ := j.pcsProviderAbi.Pack("userLiquidations", user.UserAddress)
		firstCalls := []mcCall{
			{Target: common.HexToAddress(realLpProviderAddr), AllowFailure: false, CallData: callGetUserLps},
			{Target: common.HexToAddress(realLpProviderAddr), AllowFailure: false, CallData: callToken0},
			{Target: common.HexToAddress(realLpProviderAddr), AllowFailure: false, CallData: callToken1},
			{Target: common.HexToAddress(realLpProviderAddr), AllowFailure: true, CallData: callUserLiq},
		}
		firstRes, err := doMulticall3(ctx, j.ethCli, firstCalls)
		if err != nil {
			j.log.WithError(err).Error("multicall phase1 failed")
			return
		}

		var lps []*big.Int
		if err := j.pcsProviderAbi.UnpackIntoInterface(&lps, "getUserLps", firstRes[0].ReturnData); err != nil {
			j.log.WithError(err).Error("decode getUserLps failed")
			return
		}

		// decode token addresses
		var token0Addr, token1Addr common.Address
		if err := j.pcsProviderAbi.UnpackIntoInterface(&token0Addr, "token0", firstRes[1].ReturnData); err != nil {
			j.log.WithError(err).Error("decode token0 failed")
			return
		}
		if err := j.pcsProviderAbi.UnpackIntoInterface(&token1Addr, "token1", firstRes[2].ReturnData); err != nil {
			j.log.WithError(err).Error("decode token1 failed")
			return
		}

		// decode userLiquidations
		var userLiq struct {
			Ongoing    bool
			Token0Left *big.Int
			Token1Left *big.Int
		}
		if len(firstRes[3].ReturnData) > 0 {
			if err := j.pcsProviderAbi.UnpackIntoInterface(&userLiq, "userLiquidations", firstRes[3].ReturnData); err != nil {
				j.log.WithError(err).Error("decode userLiquidations failed")
				return
			}
		}

		if userLiq.Ongoing {
			j.log.Warn("user in liquidation, skip auction")
			return
		}

		secondCalls := make([]mcCall, 0, len(lps)*2)
		for _, lp := range lps {
			cd, _ := j.pcsProviderAbi.Pack("getLpValue", lp)
			cdAmounts, _ := j.pcsProviderAbi.Pack("getAmounts", lp)
			secondCalls = append(secondCalls, mcCall{Target: common.HexToAddress(realLpProviderAddr), AllowFailure: true, CallData: cd})
			secondCalls = append(secondCalls, mcCall{Target: common.HexToAddress(realLpProviderAddr), AllowFailure: true, CallData: cdAmounts})
		}
		secondRes, err := doMulticall3(ctx, j.ethCli, secondCalls)
		if err != nil {
			j.log.WithError(err).Error("multicall phase2 failed")
			return
		}
		var minLp *big.Int
		var minLpVal *big.Int
		var minLpAmounts struct {
			Amount0 *big.Int
			Amount1 *big.Int
		}
		for i := 0; i < len(lps); i++ {
			valIdx := i * 2
			amtIdx := i*2 + 1
			if !secondRes[valIdx].Success || len(secondRes[valIdx].ReturnData) == 0 {
				continue
			}
			var v *big.Int
			if err := j.pcsProviderAbi.UnpackIntoInterface(&v, "getLpValue", secondRes[valIdx].ReturnData); err != nil {
				continue
			}
			if minLpVal == nil || v.Cmp(minLpVal) < 0 {
				minLpVal = v
				minLp = lps[i]
				var amt struct {
					Amount0 *big.Int
					Amount1 *big.Int
				}
				if len(secondRes[amtIdx].ReturnData) > 0 {
					if err := j.pcsProviderAbi.UnpackIntoInterface(&amt, "getAmounts", secondRes[amtIdx].ReturnData); err == nil {
						minLpAmounts = amt
					}
				}
			}
		}
		if minLp == nil {
			j.log.Error("no LP found for user")
			return
		}
		// --- compare LP with collateral value ---
		collateralValue := new(big.Int).Mul(collatAmount, maxPrice) // 1e18 * 1e27 = 1e45
		scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(27), nil)
		collateralValue.Div(collateralValue, scale)
		if collateralValue.Cmp(minLpVal) > 0 {
			// User has more LP than needed — check token liquidation values
			//tokens, err := j.pcsProvider.UserLiquidations(&bind.CallOpts{}, user.UserAddress)
			//if err != nil {
			//	j.log.WithError(err).Error("failed to GetUserLiquidations")
			//	return
			//}

			//t0Addr, _ := j.pcsProvider.Token0(&bind.CallOpts{})
			//t1Addr, _ := j.pcsProvider.Token1(&bind.CallOpts{})

			// prices in 1e8
			t0Price, err := j.multiOracle.Peek(&bind.CallOpts{}, token0Addr)
			if err != nil {
				j.log.WithError(err).Error("failed to get t0 price")
				return
			}
			t1Price, err := j.multiOracle.Peek(&bind.CallOpts{}, token1Addr)
			if err != nil {
				j.log.WithError(err).Error("failed to get t1 price")
				return
			}

			t0Amt := userLiq.Token0Left
			t1Amt := userLiq.Token1Left

			// convert to value in 1e8 scale
			scale18 := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
			scale10 := new(big.Int).Exp(big.NewInt(10), big.NewInt(10), nil)
			t0Val := new(big.Int).Div(new(big.Int).Mul(t0Amt, t0Price), scale18)
			t1Val := new(big.Int).Div(new(big.Int).Mul(t1Amt, t1Price), scale18)
			totalVal := new(big.Int).Add(t0Val, t1Val)

			collateralValue1e8 := new(big.Int).Div(collateralValue, scale10)
			j.log.WithFields(logrus.Fields{
				"t0_amount":        t0Amt.String(),
				"t1_amount":        t1Amt.String(),
				"t0_price_1e8":     t0Price.String(),
				"t1_price_1e8":     t1Price.String(),
				"t0_value_1e8":     t0Val.String(),
				"t1_value_1e8":     t1Val.String(),
				"total_value_1e8":  totalVal.String(),
				"collat_value_1e8": collateralValue.String(),
			}).Debug("liquidation value check")

			if totalVal.Cmp(collateralValue1e8) < 0 {
				encodedData, err := args.Pack(t0Amt, t1Amt, minLp)
				if err != nil {
					j.log.WithError(err).Error("failed to pack LP liquidation data")
					return
				}
				data = encodedData
			}
		} else {
			// LP is smaller than needed collateral — use LP directly
			encodedData, err := args.Pack(minLpAmounts.Amount0, minLpAmounts.Amount1, minLp)
			if err != nil {
				j.log.WithError(err).Error("failed to pack LP data")
				return
			}
			data = encodedData
		}
	}

	input, err := j.interAbi.Pack(
		"buyFromAuction",
		j.collateralAddr,
		auctionID,
		collatAmount,
		maxPrice,
		opts.From,
		data,
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
		data,
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
