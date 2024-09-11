package main

import (
	"context"
	"flag"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lista-dao/AuctionBots-go/internal/dao/v2/flash/flashbuy"
	"github.com/lista-dao/AuctionBots-go/pkg/config"
	"github.com/lista-dao/AuctionBots-go/pkg/utils"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
	"path/filepath"
)

var configFile = flag.String("config", "./config/config.yaml", "config file path")

func main() {
	cfg, err := loadConfig()
	if err != nil {
		logrus.Fatalf("loadConfig err: %v", err)
	}

	env, err := config.LoadEnvironmentResource(cfg)
	if err != nil {
		logrus.Fatalf("config.LoadEnvironmentResource err: %v", err)
	}

	buyAuction(env)
}

func buyAuction(resource *config.Resource) {
	ctx := context.Background()

	opts, err := resource.Wallet.Opts(ctx)
	if err != nil {
		logrus.Errorf("resource.Wallet.Opts err: %v", err)
		return
	}

	flashBuyAddress := common.HexToAddress("0x5C25A9FC1CFfda5D7E871C73929Dfca85ef6c92d")

	flashBuyContract, err := flashbuy.NewFlashbuy(flashBuyAddress, resource.HttpNodeClient)
	if err != nil {
		logrus.Errorf("flashBuy.NewFlashbuy err: %v", err)
		return
	}

	// lisUSD
	token := common.HexToAddress("0x0782b6d8c4551B9760e74c0545a9bCD90bdc41E5")
	auctionId := big.NewInt(591)
	borrowAmount, _ := big.NewInt(0).SetString("10000000000000000000000", 10)
	collateral := common.HexToAddress("0x563282106A5B0538f8673c787B3A16D3Cc1DbF1a")
	collateralAmount, _ := big.NewInt(0).SetString("298000000000000000", 10)
	maxPrice, _ := big.NewInt(0).SetString("384205199372210944444444444343", 10)
	slip := big.NewInt(10)
	collateralReal := common.HexToAddress("0xB0b84D294e0C75A6abe60171b70edEb2EFd14A1B")

	// slisBNB - BNB - lisUSD
	// resource
	tokens := []string{
		"0xB0b84D294e0C75A6abe60171b70edEb2EFd14A1B", // slisBNB
		"0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", // BNB
		"0x0782b6d8c4551B9760e74c0545a9bCD90bdc41E5", // lisUSD
	}
	// fee 500 2500
	fees := []uint64{
		500,
		2500,
	}

	path := utils.PathToBytes(tokens, fees)

	logrus.Infof("path: %v", hexutil.Encode(path))
	// collateral token
	tx, err := flashBuyContract.FlashBuyAuction(
		opts,
		token,
		auctionId,
		borrowAmount,
		collateral,
		collateralAmount,
		maxPrice,
		slip,
		collateralReal,
		path,
	)
	if err != nil {
		logrus.Errorf("flashBuyContract.FlashBuyAuction err: %v", err)
		return
	}
	logrus.Infof("flashBuy tx: %v", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, resource.HttpNodeClient, tx)
	if err != nil {
		logrus.Errorf("bind.WaitMined err: %v", err)
		return
	}

	if receipt.Status == types.ReceiptStatusFailed {
		if _, err := getRevertReason(ctx, resource.HttpNodeClient, tx, receipt); err != nil {
			logrus.WithError(err).Error("tx failed")
			return
		}
	}

}

func loadConfig() (*config.Config, error) {
	exe, err := os.Executable()
	if err != nil {
		logrus.Errorf("os.Executable err: %v", err)
		return nil, err
	}

	pwd, err := os.Getwd()
	if err != nil {
		logrus.Errorf("os.Getwd err: %v", err)
		return nil, err
	}

	dir := filepath.Dir(exe)

	logrus.Infof("current dir: %s", dir)
	logrus.Infof("current work dir: %s", pwd)

	cfg, err := config.LoadConfig(*configFile, dir)
	if err != nil {
		logrus.Errorf("config.LoadConfig err: %v", err)
		cfg, err = config.LoadConfig(*configFile, pwd)
	}

	return cfg, nil
}

func getRevertReason(ctx context.Context, caller ethereum.ContractCaller, tx *types.Transaction, receipt *types.Receipt) (*string, error) {
	from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return nil, err
	}

	resBytes, errCall := caller.CallContract(ctx, ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Data:     tx.Data(),
	}, receipt.BlockNumber)

	return bytesToStr(resBytes), errCall
}

func bytesToStr(bytes []byte) *string {
	if bytes == nil {
		return nil
	}

	result := string(bytes)
	return &result
}
