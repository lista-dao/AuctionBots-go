package cli

import (
	"context"
	"github.com/helio-money/auctionbot/internal/config"
	"github.com/helio-money/auctionbot/internal/jobs"
	"github.com/pkg/errors"
)

func Run() bool {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(errors.Wrap(err, "failed to create config"))
	}

	//jobs.NewBuyAuctionJob(
	//	context.Background(),
	//	cfg.Log(),
	//	cfg.Wallet(),
	//	cfg.HTTP(),
	//	cfg.InteractionContract(),
	//	cfg.Token0Contract(),
	//	cfg.HayContract(),
	//	true,
	//).Run()

	jobs.NewBuyFlashAuctionJob(
		context.Background(),
		cfg.Log(),
		cfg.Wallet(),
		cfg.HTTP(),
		cfg.InteractionContract(),
		cfg.Token0Contract(),
		cfg.HayContract(),
		true,
	).Run()

	//jobs.NewStartAuctionJob(
	//	context.Background(),
	//	cfg.Log(),
	//	cfg.Wallet(),
	//	cfg.AnalyticsCli(),
	//	cfg.HTTP(),
	//	cfg.InteractionContract(),
	//	cfg.Token0Contract(),
	//	true,
	//).Run()
	//
	//jobs.NewResetAuctionJob(
	//	context.Background(),
	//	cfg.Log(),
	//	cfg.Wallet(),
	//	cfg.HTTP(),
	//	cfg.InteractionContract(),
	//	cfg.Token0Contract(),
	//	true,
	).Run()

	return true
}
