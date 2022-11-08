package config

import (
	analyticsv1 "github.com/helio-money/auctionbot/internal/analytics/v1"
	"github.com/sirupsen/logrus"
)

type configer interface {
	populate()
	validate() error
}

type Config interface {
	Log() *logrus.Logger
	AnalyticsCli() *analyticsv1.Client
	contracter
	ethereumer
}

type config struct {
	analyticsClient
	logger
	ethereum
	contracts
}

func NewConfig() (Config, error) {
	cfg := config{}

	cfg.populate()

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
