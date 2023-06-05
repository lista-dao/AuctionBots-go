package cli

import (
	"context"
	"fmt"
	"github.com/helio-money/auctionbot/internal/config"
	"github.com/helio-money/auctionbot/internal/jobs"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	commandBuyFlashAuction = "buy_flash_auction"
	commandBuyAction       = "buy_auction"
	commandResetAction     = "reset_auction"
	commandStartAction     = "start_auction"
)

func Run(args []string) bool {
	if len(args) == 0 {
		panic(fmt.Sprintf("please set bot mode %v", []string{
			commandBuyAction,
			commandResetAction,
			commandStartAction,
			commandBuyFlashAuction,
		}))
	}

	cfg, err := config.NewConfig()
	if err != nil {
		panic(errors.Wrap(err, "failed to create config"))
	}

	jj := make([]jobs.Job, len(args))
	for ind, arg := range args {
		switch arg {
		case commandResetAction:
			jj[ind] = jobs.NewResetAuctionJob(
				context.Background(),
				cfg.Log(),
				cfg.Wallet(),
				cfg.HTTP(),
				cfg.InteractionContract(),
				cfg.Token0Contract(),
				true,
			)
		case commandBuyAction:
			jj[ind] = jobs.NewBuyAuctionJob(
				context.Background(),
				cfg.Log(),
				cfg.Wallet(),
				cfg.HTTP(),
				cfg.InteractionContract(),
				cfg.Token0Contract(),
				cfg.HayContract(),
				true,
			)
		case commandStartAction:
			jj[ind] = jobs.NewStartAuctionJob(
				context.Background(),
				cfg.Log(),
				cfg.Wallet(),
				cfg.AnalyticsCli(),
				cfg.HTTP(),
				cfg.InteractionContract(),
				cfg.Token0Contract(),
				true,
			)
		case commandBuyFlashAuction:
			if cfg.FlashBuyContract() == nil {
				panic(fmt.Sprintf("FLASHBUY contract must be set for %s mode", commandBuyFlashAuction))
			}

			jj[ind] = jobs.NewBuyFlashAuctionJob(
				context.Background(),
				cfg.Log(),
				cfg.Wallet(),
				cfg.HTTP(),
				cfg.InteractionContract(),
				cfg.Token0Contract(),
				cfg.HayContract(),
				*cfg.FlashBuyContract(),
				true,
			)
		default:
			panic(fmt.Sprintf("such command %s is not exists", arg))
		}
	}

	ctx, cancel := ctxWithSig()
	defer cancel()
	defer func() {
		if err := recover(); err != nil {
			cfg.Log().Error(err)
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
