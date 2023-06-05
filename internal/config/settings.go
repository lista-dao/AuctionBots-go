package config

import (
	"github.com/namsral/flag"
	"github.com/pkg/errors"
	"math/big"
	"sync"
)

var _ configer = (*settings)(nil)

type settings struct {
	settingsParams
	maxPricePerc *big.Int
	o            sync.Once
}

type settingsParams struct {
	maxPricePerc int64
}

func (conf *settings) validate() error {
	return errors.Wrap(conf.check(), "failed to validate analytics client")
}

func (conf *settings) populate() {
	conf.o.Do(func() {
		flag.Int64Var(&conf.settingsParams.maxPricePerc, "MAX_PRICE_PERCENTAGE", 80, "maximum actual_price/auction_price proportion")
		flag.Parse()
	})
}

func (cli *settings) check() error {
	if cli.settingsParams.maxPricePerc > 0 && cli.settingsParams.maxPricePerc < 100 {
		return errors.New("MAX_PRICE_PERCENTAGE must be < 100 and > 0")
	}
	cli.maxPricePerc = big.NewInt(cli.settingsParams.maxPricePerc)
	return nil
}

func (conf *settings) MaxPricePercentage() *big.Int {
	return conf.maxPricePerc
}
