package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func PathToBytes(tokens []string, fee []uint64) []byte {
	path := make([]byte, 0, 1000)

	for i := range tokens {
		token := common.HexToAddress(tokens[i])
		path = append(path, token.Bytes()...)
		if i <= len(fee)-1 {
			path = append(path, Uint24Bytes(fee[i])...)
		}
	}
	return path
}

func Uint24Bytes(i uint64) []byte {
	b := make([]byte, 3)
	b[2] = byte(i)
	b[1] = byte(i >> 8)
	b[0] = byte(i >> 16)
	return b
}

func BytesToUint24(b []byte) uint64 {
	return uint64(b[2]) | uint64(b[1])<<8 | uint64(b[0])<<16
}

func BytesToPath(b []byte) ([]string, []uint64) {
	var (
		tokens []string
		fees   []uint64
	)

	for i := 0; i < len(b); i += 23 {
		tokens = append(tokens, hexutil.Encode(b[i:i+20]))
		if len(b)-i > 20 {
			fees = append(fees, BytesToUint24(b[i+20:i+23]))
		}

	}
	return tokens, fees
}
