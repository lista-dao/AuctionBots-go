package jobs

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

var WAD = big.NewInt(1).Exp(big.NewInt(10), big.NewInt(18), nil)
var BLN = big.NewInt(1).Exp(big.NewInt(10), big.NewInt(9), nil)
var RAY = big.NewInt(1).Exp(big.NewInt(10), big.NewInt(27), nil)
var AUCTION_CAP = big.NewInt(0).Mul(big.NewInt(10000), WAD)

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
