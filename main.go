package main

import (
	"github.com/helio-money/auctionbot/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args[1:]) {
		os.Exit(1)
	}
}
