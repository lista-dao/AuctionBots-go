package config

import (
	"github.com/helio-money/auctionbot/internal/wallet"
	"github.com/namsral/flag"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"net/url"
)

type ethereumer interface {
	WS() *ethclient.Client
	HTTP() *ethclient.Client
	Wallet() wallet.Walleter
}

var _ configer = (*ethereum)(nil)

type ethereum struct {
	ethereumParams
	wsCli   *ethclient.Client
	httpCli *ethclient.Client
	wallet  wallet.Walleter
}

type ethereumParams struct {
	WSURL      string
	HTTPURL    string
	PrivateKey string
}

func (d *ethereum) populate() {
	flag.StringVar(&d.HTTPURL, "HTTP_RPC_URL", "", "ethereum node rpc wss url")
	flag.StringVar(&d.WSURL, "WS_RPC_URL", "", "ethereum node rpc wss url")
	flag.StringVar(&d.ethereumParams.PrivateKey, "PRIVATE_KEY", "", "signer private key")
	flag.Parse()
}

func (d *ethereum) validate() error {
	return errors.Wrap(d.check(), "failed to validate ethereum")
}

func (d *ethereum) check() error {
	if u, err := url.Parse(d.HTTPURL); err != nil {
		return err
	} else {
		if !u.IsAbs() {
			return errors.New("HTTP_RPC_URL scheme (http|https) is missing")
		}
	}
	if u, err := url.Parse(d.WSURL); err != nil {
		return err
	} else {
		if !u.IsAbs() {
			return errors.New("WS_RPC_URL scheme (ws|wss) is missing")
		}
	}
	if d.ethereumParams.PrivateKey == "" {
		return errors.New("private_key is missing")
	}

	return nil
}

func (d *ethereum) HTTP() *ethclient.Client {
	if d.httpCli == nil {
		var err error
		if d.httpCli, err = ethclient.Dial(d.HTTPURL); err != nil {
			panic(err)
		}
	}

	return d.httpCli
}

func (d *ethereum) WS() *ethclient.Client {
	if d.wsCli == nil {
		var err error
		if d.wsCli, err = ethclient.Dial(d.WSURL); err != nil {
			panic(err)
		}
	}

	return d.wsCli
}

func (d *ethereum) Wallet() wallet.Walleter {
	if d.wallet == nil {
		var err error
		prvk, err := crypto.HexToECDSA(d.ethereumParams.PrivateKey)
		if err != nil {
			panic(err)
		}

		d.wallet = wallet.NewWallet(d.HTTP(), prvk)
	}

	return d.wallet
}
