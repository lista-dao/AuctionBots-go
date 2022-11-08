package config

import (
	analyticsv1 "github.com/helio-money/auctionbot/internal/analytics/v1"
	"github.com/namsral/flag"
	"github.com/pkg/errors"
	"net/url"
	"sync"
)

var _ configer = (*analyticsClient)(nil)

type analyticsClient struct {
	analyticsClientParams
	cli *analyticsv1.Client
	url *url.URL
	o   sync.Once
}

type analyticsClientParams struct {
	url string
}

func (cli *analyticsClient) validate() error {
	return errors.Wrap(cli.check(), "failed to validate analytics client")
}

func (contr *analyticsClient) populate() {
	flag.StringVar(&contr.analyticsClientParams.url, "ANALYTICS_URL", "", "url to the analytics api")
	flag.Parse()
}

func (cli *analyticsClient) check() error {
	if cli.analyticsClientParams.url == "" {
		return errors.New("url is empty")
	}
	var err error
	if cli.url, err = url.Parse(cli.analyticsClientParams.url); err != nil {
		return errors.New("failed to parse url")
	}

	return nil
}

func (cli *analyticsClient) AnalyticsCli() *analyticsv1.Client {
	cli.o.Do(func() {
		if cli.cli == nil {
			cli.cli = analyticsv1.NewClient(cli.url)
		}
	})
	return cli.cli
}
