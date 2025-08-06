package v1

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type LiquidationUsersResp struct {
	Users []User `json:"users"`
}

type CommonDataResp struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Users []User `json:"users"`
	} `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type User struct {
	UserAddress common.Address  `json:"userAddress"`
	TokenName   string          `json:"tokenName"`
	Collateral  decimal.Decimal `json:"collateral"`
	//LiquidationCost decimal.Decimal `json:"liquidationCost"`
	//RangeFromLiquidation decimal.Decimal `json:"rangeFromLiquidation"`
	LiquidationPrice decimal.Decimal `json:"liquidationPrice"`
	ClipperAddress   common.Address  `json:"clipperAddress"`
}
