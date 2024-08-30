package utils

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
)

func TestUint24Bytes(t *testing.T) {
	bs := Uint24Bytes(500)
	h := hexutil.Encode(bs)
	t.Log("hex", h)
}

func TestBytesToUint24(t *testing.T) {
	bs, err := hexutil.Decode("0x0001f4")
	if err != nil {
		t.Errorf("hexutil.Decode err: %v", err)
		return
	}

	i := BytesToUint24(bs)
	t.Log("value:", i)
}

func TestBytesToPath(t *testing.T) {
	bs, err := hexutil.Decode("0xb0b84d294e0c75a6abe60171b70edeb2efd14a1b0001f4bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c0009c40782b6d8c4551b9760e74c0545a9bcd90bdc41e5")
	if err != nil {
		t.Errorf("hexutil.Decode err: %v", err)
		return
	}

	tokens, fees := BytesToPath(bs)
	t.Log("tokens", tokens)
	t.Log("fees", fees)
}

func TestPathToBytes(t *testing.T) {
	tokens := []string{
		"0xfceb31a79f71ac9cbdcf853519c1b12d379edc46",
		"0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
		"0x55d398326f99059ff775485246999027b3197955",
	}
	fees := []uint64{
		10000,
		500,
	}
	bs := PathToBytes(tokens, fees)
	h := hexutil.Encode(bs)
	t.Log("hex", h)
}
