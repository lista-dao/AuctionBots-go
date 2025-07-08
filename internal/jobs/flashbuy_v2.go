package jobs

import (
	"context"
	"encoding/json"
	"github.com/1inch/1inch-sdk-go/constants"
	"github.com/1inch/1inch-sdk-go/sdk-clients/aggregation"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	daov1 "github.com/lista-dao/AuctionBots-go/internal/dao/v1"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/clipper"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/flash/flashbuy"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/flash/interfaces/ierc3156flashlender"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/flash/liquidator"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/stakemanager"
	"github.com/lista-dao/AuctionBots-go/pkg/config"
	"strconv"
	"strings"
	"time"

	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/interaction"
	"github.com/lista-dao/AuctionBots-go/internal/wallet"
	"github.com/sirupsen/logrus"
	"math/big"
)

func NewBuyFlashAuctionV2Job(
	ctx context.Context,
	log *logrus.Logger,
	wall wallet.Walleter,
	ethCli *ethclient.Client,
	interactAddr common.Address,
	collateralAddr common.Address,
	hayAddr common.Address,
	liquidatorAddr common.Address,
	maxPricePerc *big.Int,
	withWait bool,
	cfg *config.Config,
) Job {
	return &buyFlashAuctionV2Job{
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
		hayAddr:        hayAddr,
		withWait:       withWait,
		liquidatorAddr: liquidatorAddr,
		//maxPricePerc: maxPricePerc,
		maxPricePerc: big.NewInt(95),
		cfg:          cfg,
	}
}

var _ Job = (*buyFlashAuctionV2Job)(nil)

type buyFlashAuctionV2Job struct {
	ctx context.Context

	// job will buy action when
	// auction_price < actual_price * percentage (1-100%)
	maxPricePerc *big.Int

	wallet           wallet.Walleter
	ethCli           *ethclient.Client
	log              *logrus.Entry
	collateralAddr   common.Address
	hayAddr          common.Address
	interactAddr     common.Address
	spotAddr         common.Address
	liquidatorAddr   common.Address
	stakeManagerAddr common.Address
	collateralIlk    [32]byte

	flash           *flashbuy.Flashbuy
	liquidator      *liquidator.Liquidator
	stakeManager    *stakemanager.Stakemanager
	flashAbi        *abi.ABI
	liquidatorAbi   *abi.ABI
	stakeManagerAbi *abi.ABI
	lender          *ierc3156flashlender.Ierc3156flashlender

	inter   *interaction.Interaction
	clipper *clipper.Clipper
	oracle  *daov1.MockOracle

	withWait bool
	cfg      *config.Config
}

func (j *buyFlashAuctionV2Job) init() {
	if j == nil {
		panic("buy auction job is null")
	}
	var err error
	//j.flash, _ = flashbuy.NewFlashbuy(j.flashBuyAddr, j.ethCli)
	//j.flashAbi, err = flashbuy.FlashbuyMetaData.GetAbi()
	//if err != nil {
	//	panic(err)
	//}
	j.stakeManagerAddr = common.HexToAddress(j.cfg.Contract.StakeManager)
	j.liquidator, _ = liquidator.NewLiquidator(j.liquidatorAddr, j.ethCli)
	j.liquidatorAbi, err = liquidator.LiquidatorMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	j.stakeManager, _ = stakemanager.NewStakemanager(j.stakeManagerAddr, j.ethCli)
	j.stakeManagerAbi, err = stakemanager.StakemanagerMetaData.GetAbi()
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

func (j *buyFlashAuctionV2Job) Run(ctx context.Context) {
	j.init()
	ticker := time.NewTicker(time.Minute)
	go func() {
		j.log.Debug("start")

		for {
			select {
			case <-ticker.C:
				Monitor.Beat("buyFlashAuctionJob")
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

func (j *buyFlashAuctionV2Job) processAuction(auctionID *big.Int) {
	log := j.log.WithFields(logrus.Fields{
		"auction_id": auctionID.String(),
	})

	status, err := j.clipper.GetStatus(&bind.CallOpts{}, auctionID)
	if err != nil {
		log.WithError(err).Error("failed to get status from clipper")
		return
	}

	//chost, err := j.clipper.Chost(&bind.CallOpts{})
	//if err != nil {
	//	log.WithError(err).Error("failed to get chost from clipper")
	//	return
	//}

	// Normalize values (price, tab, chost) from WAD/BLN scaling
	auctionPrice := new(big.Int).Div(status.Price, BLN) // uint256 price
	tab := status.Tab

	// collateralAmt = (tab - chost) / auctionPrice
	if auctionPrice.Cmp(big.NewInt(0)) == 0 {
		log.Error("auction price is zero, cannot calculate collateral amount")
		return
	}
	inchCollatAmount := new(big.Int).Div(status.Tab, status.Price)
	//stakeManager.convertBnbToSnBnb(109249908013584147)
	bnbAddr := strings.ToLower(j.cfg.Contract.Collaterals[0])
	if bnbAddr == strings.ToLower(j.collateralAddr.Hex()) { //bnb
		converted, err := j.stakeManager.ConvertBnbToSnBnb(&bind.CallOpts{}, inchCollatAmount)
		log.Infof("stakeManager orgin: %v, converted: %v", inchCollatAmount, converted)
		if err != nil {
			log.WithError(err).Error("failed to get status from stakeManager")
			return
		}
		inchCollatAmount = converted
	}
	log.WithFields(logrus.Fields{
		"auctionPrice": auctionPrice.String(),
		"statusPrice":  status.Price.String(),
		"tab":          tab.String(),
		//"chost":            chost.String(),
		"inchCollatAmount": inchCollatAmount.String(),
	}).Info("calculated collateral amount")

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

	collatHay := big.NewInt(0).Div(big.NewInt(0).Mul(status.Price, status.Lot), RAY)
	if collatHay.Cmp(AUTIION_MAX_CAP) >= 0 {
		log.Infof("buy_auction : ignoring users with excessive max auction amounts...")
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
		hayMax = big.NewInt(0).Add(big.NewInt(0), AUCTION_CAP)
	}
	//j.flashBuyAuction(log, auctionID, collatAmount, hayMax, status.Price)
	j.flashLiquidate(log, auctionID, collatAmount, hayMax, status.Price, inchCollatAmount)
	log.Debug("flash auction bought")
}

func (j *buyFlashAuctionV2Job) flashLiquidate(log *logrus.Entry, auctionID *big.Int, collateralAmt *big.Int, borrowAmount, maxPrice *big.Int, inchCollAmt *big.Int) {
	opts, err := j.wallet.Opts(j.ctx)
	if err != nil {
		log.WithError(err).Error("failed to get tx opts")
		return
	}
	fb, ok := j.cfg.FlushBuy.Paths[strings.ToLower(j.collateralAddr.Hex())]
	if !ok {
		fb = config.FlushBuyConfig{
			Received: j.collateralAddr.Hex(),
			Scale:    "18",
		}
	}
	scale := fb.Scale
	if scale == "" || scale == "0" {
		scale = "18" // default to 18 decimals
	}

	scaledInchAmt := new(big.Int).Set(inchCollAmt)
	if scale != "18" {
		if tokenDecimals, err := strconv.ParseInt(scale, 10, 64); err == nil && tokenDecimals < 18 {
			divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18-tokenDecimals), nil)
			scaledInchAmt = new(big.Int).Div(inchCollAmt, divisor)
		}
	}
	collateralReal := common.HexToAddress(fb.Received)
	rpcUrl := j.cfg.RpcNode.Http
	swapConf, err := aggregation.NewConfiguration(aggregation.ConfigurationParams{
		NodeUrl:    rpcUrl,
		PrivateKey: j.cfg.Wallet.PrivateKey,
		ChainId:    constants.BscChainId,
		ApiUrl:     "https://api.1inch.dev",
		ApiKey:     j.cfg.FlushBuy.OneInchKey,
	})
	if err != nil {
		log.Fatalf("Failed to create 1inch configuration: %v\n", err)
	}
	client, err := aggregation.NewClient(swapConf)
	if err != nil {
		log.Fatalf("Failed to create 1inch client: %v\n", err)
	}
	ctx := context.Background()

	swapData, err := client.GetSwap(ctx, aggregation.GetSwapParams{
		Src:              collateralReal.Hex(),
		Dst:              j.hayAddr.Hex(),
		Amount:           scaledInchAmt.String(),
		From:             j.cfg.Contract.Liquidator,
		Slippage:         float32(j.cfg.FlushBuy.OneInchSlip),
		DisableEstimate:  true,
		AllowPartialFill: false, // Set to true to allow partial filling of the swap order
	})
	if err != nil {
		log.WithError(err).Error("Failed to get swap data")
		return
	}

	output, err := json.MarshalIndent(swapData, "", "  ")
	if err != nil {
		log.WithError(err).Error("Failed to marshal swap data")
		return
	}
	log.Infof("1inch swap resp: %s\n", string(output))

	// slip := big.NewInt(j.cfg.FlushBuy.Slip)
	// path := utils.PathToBytes(fb.Tokens, fb.Fees)
	onePercent := new(big.Int).Div(new(big.Int).Mul(borrowAmount, big.NewInt(1)), big.NewInt(100))
	borrowWithBuffer := new(big.Int).Add(borrowAmount, onePercent)
	swapCallData, err := hexutil.Decode(swapData.SwapResponse.Tx.Data)
	if err != nil {
		log.WithError(err).Error("invalid hex calldata")
		return
	}
	swapTarget := common.HexToAddress(swapData.SwapResponse.Tx.To)
	input, err := j.liquidatorAbi.Pack(
		"flashLiquidate",
		auctionID,
		borrowWithBuffer,
		j.collateralAddr,
		collateralAmt,
		maxPrice,
		collateralReal,
		swapTarget,
		swapCallData,
	)
	if err != nil {
		log.WithError(err).Error("failed to pack flashLiquidate input")
		return
	}
	log.WithFields(logrus.Fields{
		"auctionID":      auctionID,
		"collateral":     j.collateralAddr.Hex(),
		"inchCollAmt":    inchCollAmt.String(),
		"maxPrice":       maxPrice.String(),
		"collateralReal": collateralReal.Hex(),
		"swapTarget":     swapTarget.Hex(),
		"scaleInchAmt":   scaledInchAmt.String(),
	}).Debug("Prepared flashLiquidate call")

	_, err = j.ethCli.EstimateGas(ctx, ethereum.CallMsg{
		From:     j.wallet.Address(),
		To:       &j.liquidatorAddr,
		Gas:      opts.GasLimit,
		GasPrice: opts.GasPrice,
		Value:    opts.Value,
		Data:     input,
	})
	if err != nil {
		log.WithError(err).Error("failed to estimate gas")
		return
	}

	tx, err := j.liquidator.FlashLiquidate(
		opts,
		auctionID,
		borrowWithBuffer,
		j.collateralAddr,
		collateralAmt,
		maxPrice,
		collateralReal,
		swapTarget,
		swapCallData,
	)
	if err != nil {
		log.WithError(err).Error("failed to send flashLiquidate tx")
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
}
