package config

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
	analyticsv3 "github.com/lista-dao/AuctionBots-go/internal/analytics/v3"
	"net/http"
	"net/url"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	analyticsv1 "github.com/lista-dao/AuctionBots-go/internal/analytics/v1"
	"github.com/lista-dao/AuctionBots-go/internal/wallet"
	"github.com/sirupsen/logrus"
)

type Resource struct {
	Log               *logrus.Logger
	HttpNodeClient    *ethclient.Client
	Wallet            wallet.Walleter
	AnalyticsClient   *analyticsv1.Client
	AnalyticsClientV3 *analyticsv3.Client
}

func LoadEnvironmentResource(config *Config) (*Resource, error) {
	var (
		resource Resource
		err      error
	)
	resource.Log, err = initLogger(config.Log.Level)
	if err != nil {
		return nil, fmt.Errorf("initLogger err: %w", err)
	}

	resource.Log.SetReportCaller(config.Log.Caller)

	analyticsUrl, err := url.Parse(config.Analytics.Url)
	analyticsUrlV3, err := url.Parse(config.Analytics.ListaApiUrl)
	if err != nil {
		return nil, fmt.Errorf("url.Parse err: %w", err)
	}
	resource.AnalyticsClient = analyticsv1.NewClient(analyticsUrl)
	resource.AnalyticsClientV3 = analyticsv3.NewClient(analyticsUrlV3, resource.Log)

	resource.HttpNodeClient, err = initHttpNodeClient(config.RpcNode.Http)
	if err != nil {
		return nil, fmt.Errorf("initHttpNodeClient err: %w", err)

	}

	resource.Wallet, err = initWallet(config.Wallet.PrivateKey, resource.HttpNodeClient)
	if err != nil {
		return nil, fmt.Errorf("initWallet err: %w", err)
	}

	return &resource, nil
}

func initLogger(level string) (*logrus.Logger, error) {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, fmt.Errorf("logrus.ParseLevel err: %w", err)
	}

	log := logrus.StandardLogger()
	log.Level = l
	return log, nil
}

func initWallet(privateKey string, httpClient *ethclient.Client) (wallet.Walleter, error) {
	var err error
	prvk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("crypto.HexToECDSA err: %w", err)
	}

	w := wallet.NewWallet(httpClient, prvk)

	return w, nil
}

func initHttpNodeClient(httpURL string) (*ethclient.Client, error) {
	httpCli, err := ethclient.Dial(httpURL)
	if err != nil {
		return nil, fmt.Errorf("ethclient.Dial err: %w", err)
	}

	return httpCli, nil
}

func initWsNodeClient(wsURL string) (*ethclient.Client, error) {
	ctx := context.Background()

	dialer := websocket.Dialer{
		Proxy:           http.ProxyFromEnvironment,
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		WriteBufferPool: new(sync.Pool),
	}

	conn, err := rpc.DialWebsocketWithDialer(ctx, wsURL, "", dialer)
	if err != nil {
		return nil, fmt.Errorf("rpc.DialWebsocketWithDialer err: %w", err)
	}

	wsCli := ethclient.NewClient(conn)

	if err != nil {
		return nil, fmt.Errorf("ethclient.Dial err: %w", err)
	}

	return wsCli, nil
}
