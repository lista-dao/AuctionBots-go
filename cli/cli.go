package cli

import (
	"context"
	"github.com/helio-money/auctionbot/internal/config"
	"github.com/helio-money/auctionbot/internal/jobs"
	"github.com/pkg/errors"
)

func Run(args []string) bool {

	cfg, err := config.NewConfig()
	if err != nil {
		panic(errors.Wrap(err, "failed to create config"))
	}

	//app := cli.NewApp()
	//ctx := context.Background()
	//
	//app.Commands = cli.Commands{
	//	{
	//		Name: "run",
	//		Action: func(c *cli.Context) error {
	//			return cron.New().AddJob("1 * * * *")
	//		},
	//	},
	//}
	//
	//if err := app.RunContext(ctx, args); err != nil {
	//	logrus.WithError(err).Error("app finished")
	//	return false
	//}
	//
	//if err := cr.AddJob("1 * * * *", jobs.NewBuyAuctionJob(
	//	context.Background(),
	//	cfg.Log(),
	//	cfg.Wallet(),
	//	cfg.HTTP(),
	//	cfg.InteractionContract(),
	//	cfg.SpotContract(),
	//	true,
	//)); err != nil {
	//	panic(err)
	//}

	jobs.NewBuyAuctionJob(
		context.Background(),
		cfg.Log(),
		cfg.Wallet(),
		cfg.HTTP(),
		cfg.InteractionContract(),
		cfg.Token0Contract(),
		cfg.HayContract(),
		true,
	).Run()

	//jobs.NewResetAuctionJob(
	//	context.Background(),
	//	cfg.Log(),
	//	cfg.Wallet(),
	//	cfg.HTTP(),
	//	cfg.InteractionContract(),
	//	cfg.Token0Contract(),
	//	true,
	//).Run()

	//cr.Run()
	return true
}
