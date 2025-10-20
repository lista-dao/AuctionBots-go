package jobs

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sirupsen/logrus"
)

var WAD = big.NewInt(1).Exp(big.NewInt(10), big.NewInt(18), nil)
var BLN = big.NewInt(1).Exp(big.NewInt(10), big.NewInt(9), nil)
var RAY = big.NewInt(1).Exp(big.NewInt(10), big.NewInt(27), nil)
var AUCTION_CAP = big.NewInt(0).Mul(big.NewInt(5000), WAD)
var AUTIION_MAX_CAP = big.NewInt(0).Mul(big.NewInt(500000), WAD)

func getRevertReason(ctx context.Context, caller ethereum.ContractCaller, tx *types.Transaction, receipt *types.Receipt) (*string, error) {
	from, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx)
	if err != nil {
		return nil, err
	}

	resBytes, errCall := caller.CallContract(ctx, ethereum.CallMsg{
		From:     from,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Data:     tx.Data(),
	}, receipt.BlockNumber)

	return bytesToStr(resBytes), errCall
}

func bytesToStr(bytes []byte) *string {
	if bytes == nil {
		return nil
	}

	result := string(bytes)
	return &result
}

func calcPercent(x, p *big.Int) *big.Int {
	if p.Cmp(big.NewInt(0)) < 0 || p.Cmp(big.NewInt(100)) > 0 {
		panic("bad x")
	}
	return big.NewInt(0).Div(big.NewInt(0).Mul(x, p), big.NewInt(100))
}

var (
	Monitor = NewHeartbeatMonitor(1*time.Minute, 2*time.Minute, func(id string) {
		fmt.Printf("⚠️ Timeout: %s\n", id)
	})
)

type HeartbeatMonitor struct {
	mu        sync.Mutex
	interval  time.Duration
	timeout   time.Duration
	lastBeats map[string]time.Time
	onTimeout func(id string)
	stopCh    chan struct{}
	stopped   bool
}

func NewHeartbeatMonitor(interval, timeout time.Duration, onTimeout func(id string)) *HeartbeatMonitor {
	h := &HeartbeatMonitor{
		interval:  interval,
		timeout:   timeout,
		lastBeats: make(map[string]time.Time),
		onTimeout: onTimeout,
		stopCh:    make(chan struct{}),
	}
	go h.watch()
	return h
}

func (h *HeartbeatMonitor) Beat(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	logrus.Debugf("[Heartbeat] %s", id)
	h.lastBeats[id] = time.Now()
}

func (h *HeartbeatMonitor) watch() {
	ticker := time.NewTicker(h.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.checkTimeouts()
		case <-h.stopCh:
			return
		}
	}
}

func (h *HeartbeatMonitor) checkTimeouts() {
	h.mu.Lock()
	defer h.mu.Unlock()
	now := time.Now()
	for id, t := range h.lastBeats {
		if now.Sub(t) > h.timeout {
			log.Printf("[Heartbeat] Timeout error: %s", id)
			//delete(h.lastBeats, id)
			if h.onTimeout != nil {
				go h.onTimeout(id)
			}
		}
	}
}

func (h *HeartbeatMonitor) CheckStuckJobs() string {
	h.mu.Lock()
	defer h.mu.Unlock()
	now := time.Now()
	for id, t := range h.lastBeats {
		//log.Printf(fmt.Sprintf("[hb list] %s", id))
		if now.Sub(t) > 5*time.Minute {
			return id
		}
	}
	return ""
}

func (h *HeartbeatMonitor) Stop() {
	h.mu.Lock()
	defer h.mu.Unlock()
	if !h.stopped {
		h.stopped = true
		close(h.stopCh)
	}
}
