package main

import (
	"bytes"
	"context"
	"flag"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lista-dao/AuctionBots-go/internal/jobs"
	"github.com/lista-dao/AuctionBots-go/pkg/config"
	"github.com/sirupsen/logrus"
	"math/big"
)

var configFile = flag.String("config", "./config/config.yaml", "config file path")

func main() {
	flag.Parse()

	logrus.SetReportCaller(true)

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		logrus.Errorf("config.LoadConfig err: %v", err)
	}

	logrus.Infof("log.level: %+v", cfg.Log.Level)

	Run(cfg)

	block := make(chan struct{})
	<-block
}

const (
	commandBuyFlashAuction = "buy_flash_auction"
	commandBuyAction       = "buy_auction"
	commandResetAction     = "reset_auction"
	commandStartAction     = "start_auction"
)

func Run(cfg *config.Config) {
	args := cfg.Commands
	if len(args) == 0 {
		logrus.Errorf("please set bot mode %v", []string{
			commandBuyAction,
			commandResetAction,
			commandStartAction,
			commandBuyFlashAuction,
		})
		return
	}

	resource, err := config.LoadEnvironmentResource(cfg)
	if err != nil {
		logrus.Fatalf("config.LoadEnvironmentResource err: %v", err)
	}

	interaction := common.HexToAddress(cfg.Contract.Interaction)
	hay := common.HexToAddress(cfg.Contract.Hay)
	flushBuy := common.HexToAddress(cfg.Contract.FlushBuy)
	collaterals := make([]common.Address, 0, len(cfg.Contract.Collaterals))
	for _, c := range cfg.Contract.Collaterals {
		collaterals = append(collaterals, common.HexToAddress(c))
	}

	jj := make([]jobs.Job, len(args)*len(collaterals))
	i := 0
	for _, arg := range args {
		for _, collateral := range collaterals {
			switch arg {
			case commandResetAction:
				jj[i] = jobs.NewResetAuctionJob(
					context.Background(),
					resource.Log,
					resource.Wallet,
					resource.HttpNodeClient,
					interaction,
					collateral,
					true,
				)
			case commandBuyAction:
				jj[i] = jobs.NewBuyAuctionJob(
					context.Background(),
					resource.Log,
					resource.Wallet,
					resource.HttpNodeClient,
					interaction,
					collateral,
					hay,
					big.NewInt(cfg.Settings.MaxPricePercentage),
					true,
				)
			case commandStartAction:
				jj[i] = jobs.NewStartAuctionJob(
					context.Background(),
					resource.Log,
					resource.Wallet,
					resource.AnalyticsClient,
					resource.HttpNodeClient,
					interaction,
					collateral,
					true,
				)
			case commandBuyFlashAuction:
				// is address zero checking
				if bytes.Compare(flushBuy.Bytes(), common.Address{}.Bytes()) == 0 {
					logrus.Errorf("FLASHBUY contract must be set for %s mode", commandBuyFlashAuction)
					return
				}

				jj[i] = jobs.NewBuyFlashAuctionJob(
					context.Background(),
					resource.Log,
					resource.Wallet,
					resource.HttpNodeClient,
					interaction,
					collateral,
					hay,
					flushBuy,
					big.NewInt(cfg.Settings.MaxPricePercentage),
					true,
				)
			default:
				logrus.Errorf("such command %s is not exists", arg)
				return
			}
			i++
		}

	}

	ctx := context.Background()

	for _, j := range jj {
		j.Run(ctx)
	}

	logrus.Infof("start bot success!")
}
