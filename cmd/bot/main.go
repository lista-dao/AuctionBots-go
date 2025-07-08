package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	gosundheit "github.com/AppsFlyer/go-sundheit"
	"github.com/AppsFlyer/go-sundheit/checks"
	healthhttp "github.com/AppsFlyer/go-sundheit/http"
	"github.com/ethereum/go-ethereum/common"
	"github.com/lista-dao/AuctionBots-go/internal/jobs"
	"github.com/lista-dao/AuctionBots-go/pkg/config"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var configFile = flag.String("config", "./config/config.yaml", "config file path")
var LIQUIDATION_API_URL = "https://api.lista.org/api/v2/liquidations/red?start=0&count=1"

func main() {
	flag.Parse()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	exe, err := os.Executable()
	if err != nil {
		logrus.Errorf("os.Executable err: %v", err)
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		logrus.Errorf("os.Getwd err: %v", err)
		return
	}

	dir := filepath.Dir(exe)

	logrus.Infof("current dir: %s", dir)
	logrus.Infof("current work dir: %s", pwd)

	cfg, err := config.LoadConfig(*configFile, dir)
	if err != nil {
		logrus.Errorf("config.LoadConfig err: %v", err)
		cfg, err = config.LoadConfig(*configFile, pwd)
	}

	if err != nil {
		logrus.Errorf("config.LoadConfig err: %v", err)
		return
	}
	logrus.SetReportCaller(cfg.Log.Caller)

	logrus.Infof("log.level: %+v", cfg.Log.Level)

	if cfg.Wallet.PrivateKey == "" {
		logrus.Errorf("privateKey must be set, please set it in config.txt")
		return
	}

	if cfg.FlushBuy.OneInchKey == "" {
		logrus.Errorf("OneInchKey must be set, please set it in config.txt")
		return
	}

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
	defer func() {
		if r := recover(); r != nil {
			logrus.Infof("Recovering from panic : %v \n", r)
		}
	}()
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
	//flushBuy := common.HexToAddress(cfg.Contract.FlushBuy)
	liquidator := common.HexToAddress(cfg.Contract.Liquidator)
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
					resource.AnalyticsClientV3,
					resource.HttpNodeClient,
					interaction,
					collateral,
					true,
				)
			case commandBuyFlashAuction:
				// is address zero checking
				if bytes.Compare(liquidator.Bytes(), common.Address{}.Bytes()) == 0 {
					logrus.Errorf("FLASHBUY contract must be set for %s mode", commandBuyFlashAuction)
					return
				}

				jj[i] = jobs.NewBuyFlashAuctionV2Job(
					context.Background(),
					resource.Log,
					resource.Wallet,
					resource.HttpNodeClient,
					interaction,
					collateral,
					hay,
					liquidator,
					big.NewInt(cfg.Settings.MaxPricePercentage),
					true,
					cfg,
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
	HealthCheck()
}

func HealthCheck() {
	// create a new health instance
	h := gosundheit.New()

	// define an HTTP dependency check
	httpCheckConf := checks.HTTPCheckConfig{
		CheckName: "httpbin.liquidation.check",
		Timeout:   3 * time.Second,
		URL:       LIQUIDATION_API_URL,
	}
	httpCheck, err := checks.NewHTTPCheck(httpCheckConf)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = h.RegisterCheck(
		httpCheck,
		gosundheit.InitialDelay(time.Second),
		gosundheit.ExecutionPeriod(10*time.Second),
	)

	if err != nil {
		fmt.Println("Failed to register check: ", err)
		return
	}

	err = h.RegisterCheck(
		&checks.CustomCheck{
			CheckName: "jobs.heartbeat.check",
			CheckFunc: jobHealthCheck,
		},
		gosundheit.InitialDelay(5*time.Second),
		gosundheit.ExecutionPeriod(2*time.Minute),
		gosundheit.ExecutionTimeout(15*time.Second),
	)

	if err != nil {
		fmt.Println("Failed to register check: ", err)
		return
	}
	// register a health endpoint
	http.Handle("/admin/health", healthhttp.HandleHealthJSON(h))

	// serve HTTP
	log.Fatal(http.ListenAndServe(":6565", nil))
}

func jobHealthCheck(ctx context.Context) (details interface{}, err error) {
	errId := jobs.Monitor.CheckStuckJobs()
	if errId != "" {
		return nil, fmt.Errorf("stuck jobs: %s", errId)
	}
	return nil, nil
}
