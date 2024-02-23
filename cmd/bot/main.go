package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lista-dao/AuctionBots-go/internal/jobs"
	"github.com/lista-dao/AuctionBots-go/pkg/config"
	"github.com/sirupsen/logrus"
	"math/big"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var configFile = flag.String("config", "./config/config.yaml", "config file path")

func main() {
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		logrus.Errorf("config.LoadConfig err: %v", err)
	}

	logrus.Infof("log.level: %+v", cfg.Log.Level)
	logrus.Infof("rpcNode.Http: %+v", cfg.RpcNode.Http)
	logrus.Infof("rpcNode.Ws: %+v", cfg.RpcNode.Ws)

	if !Run(cfg) {
		os.Exit(1)
	}
}

const (
	commandBuyFlashAuction = "buy_flash_auction"
	commandBuyAction       = "buy_auction"
	commandResetAction     = "reset_auction"
	commandStartAction     = "start_auction"
)

func Run(cfg *config.Config) bool {
	args := cfg.Commands
	if len(args) == 0 {
		panic(fmt.Sprintf("please set bot mode %v", []string{
			commandBuyAction,
			commandResetAction,
			commandStartAction,
			commandBuyFlashAuction,
		}))
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
	for ind, arg := range args {
		for _, collateral := range collaterals {
			switch arg {
			case commandResetAction:
				jj[ind] = jobs.NewResetAuctionJob(
					context.Background(),
					resource.Log,
					resource.Wallet,
					resource.HttpNodeClient,
					interaction,
					collateral,
					true,
				)
			case commandBuyAction:
				jj[ind] = jobs.NewBuyAuctionJob(
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
				jj[ind] = jobs.NewStartAuctionJob(
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
					panic(fmt.Sprintf("FLASHBUY contract must be set for %s mode", commandBuyFlashAuction))
				}

				jj[ind] = jobs.NewBuyFlashAuctionJob(
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
				panic(fmt.Sprintf("such command %s is not exists", arg))
			}
		}

	}

	ctx, cancel := ctxWithSig()
	defer cancel()
	defer func() {
		if err := recover(); err != nil {
			resource.Log.Error(err)
			cancel()
		}
	}()

	wg := &sync.WaitGroup{}
	for _, j := range jj {
		wg.Add(1)
		j.Run(ctx, wg)
	}

	wg.Wait()
	return true
}

func ctxWithSig() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	go func() {
		select {
		case <-ch:
			cancel()
		}
	}()

	return ctx, cancel
}
