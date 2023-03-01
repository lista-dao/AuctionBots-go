// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flashbuy

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FlashbuyMetaData contains all meta data concerning the Flashbuy contract.
var FlashbuyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC3156FlashLender\",\"name\":\"lender_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"contractIAuctionProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAm\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralAm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"flashBuyAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"contractIERC3156FlashLender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onFlashLoan\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FlashbuyABI is the input ABI used to generate the binding from.
// Deprecated: Use FlashbuyMetaData.ABI instead.
var FlashbuyABI = FlashbuyMetaData.ABI

// Flashbuy is an auto generated Go binding around an Ethereum contract.
type Flashbuy struct {
	FlashbuyCaller     // Read-only binding to the contract
	FlashbuyTransactor // Write-only binding to the contract
	FlashbuyFilterer   // Log filterer for contract events
}

// FlashbuyCaller is an auto generated read-only Go binding around an Ethereum contract.
type FlashbuyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbuyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FlashbuyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbuyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FlashbuyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FlashbuySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FlashbuySession struct {
	Contract     *Flashbuy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FlashbuyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FlashbuyCallerSession struct {
	Contract *FlashbuyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FlashbuyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FlashbuyTransactorSession struct {
	Contract     *FlashbuyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FlashbuyRaw is an auto generated low-level Go binding around an Ethereum contract.
type FlashbuyRaw struct {
	Contract *Flashbuy // Generic contract binding to access the raw methods on
}

// FlashbuyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FlashbuyCallerRaw struct {
	Contract *FlashbuyCaller // Generic read-only contract binding to access the raw methods on
}

// FlashbuyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FlashbuyTransactorRaw struct {
	Contract *FlashbuyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFlashbuy creates a new instance of Flashbuy, bound to a specific deployed contract.
func NewFlashbuy(address common.Address, backend bind.ContractBackend) (*Flashbuy, error) {
	contract, err := bindFlashbuy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Flashbuy{FlashbuyCaller: FlashbuyCaller{contract: contract}, FlashbuyTransactor: FlashbuyTransactor{contract: contract}, FlashbuyFilterer: FlashbuyFilterer{contract: contract}}, nil
}

// NewFlashbuyCaller creates a new read-only instance of Flashbuy, bound to a specific deployed contract.
func NewFlashbuyCaller(address common.Address, caller bind.ContractCaller) (*FlashbuyCaller, error) {
	contract, err := bindFlashbuy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FlashbuyCaller{contract: contract}, nil
}

// NewFlashbuyTransactor creates a new write-only instance of Flashbuy, bound to a specific deployed contract.
func NewFlashbuyTransactor(address common.Address, transactor bind.ContractTransactor) (*FlashbuyTransactor, error) {
	contract, err := bindFlashbuy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FlashbuyTransactor{contract: contract}, nil
}

// NewFlashbuyFilterer creates a new log filterer instance of Flashbuy, bound to a specific deployed contract.
func NewFlashbuyFilterer(address common.Address, filterer bind.ContractFilterer) (*FlashbuyFilterer, error) {
	contract, err := bindFlashbuy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FlashbuyFilterer{contract: contract}, nil
}

// bindFlashbuy binds a generic wrapper to an already deployed contract.
func bindFlashbuy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FlashbuyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flashbuy *FlashbuyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flashbuy.Contract.FlashbuyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flashbuy *FlashbuyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flashbuy.Contract.FlashbuyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flashbuy *FlashbuyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flashbuy.Contract.FlashbuyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Flashbuy *FlashbuyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Flashbuy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Flashbuy *FlashbuyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flashbuy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Flashbuy *FlashbuyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Flashbuy.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_Flashbuy *FlashbuyCaller) Auction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flashbuy.contract.Call(opts, &out, "auction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_Flashbuy *FlashbuySession) Auction() (common.Address, error) {
	return _Flashbuy.Contract.Auction(&_Flashbuy.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(address)
func (_Flashbuy *FlashbuyCallerSession) Auction() (common.Address, error) {
	return _Flashbuy.Contract.Auction(&_Flashbuy.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Flashbuy *FlashbuyCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flashbuy.contract.Call(opts, &out, "lender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Flashbuy *FlashbuySession) Lender() (common.Address, error) {
	return _Flashbuy.Contract.Lender(&_Flashbuy.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Flashbuy *FlashbuyCallerSession) Lender() (common.Address, error) {
	return _Flashbuy.Contract.Lender(&_Flashbuy.CallOpts)
}

// FlashBuyAuction is a paid mutator transaction binding the contract method 0xb131e9d6.
//
// Solidity: function flashBuyAuction(address token, uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice) returns()
func (_Flashbuy *FlashbuyTransactor) FlashBuyAuction(opts *bind.TransactOpts, token common.Address, auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Flashbuy.contract.Transact(opts, "flashBuyAuction", token, auctionId, borrowAm, collateral, collateralAm, maxPrice)
}

// FlashBuyAuction is a paid mutator transaction binding the contract method 0xb131e9d6.
//
// Solidity: function flashBuyAuction(address token, uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice) returns()
func (_Flashbuy *FlashbuySession) FlashBuyAuction(token common.Address, auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Flashbuy.Contract.FlashBuyAuction(&_Flashbuy.TransactOpts, token, auctionId, borrowAm, collateral, collateralAm, maxPrice)
}

// FlashBuyAuction is a paid mutator transaction binding the contract method 0xb131e9d6.
//
// Solidity: function flashBuyAuction(address token, uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice) returns()
func (_Flashbuy *FlashbuyTransactorSession) FlashBuyAuction(token common.Address, auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Flashbuy.Contract.FlashBuyAuction(&_Flashbuy.TransactOpts, token, auctionId, borrowAm, collateral, collateralAm, maxPrice)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Flashbuy *FlashbuyTransactor) OnFlashLoan(opts *bind.TransactOpts, initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Flashbuy.contract.Transact(opts, "onFlashLoan", initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Flashbuy *FlashbuySession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Flashbuy.Contract.OnFlashLoan(&_Flashbuy.TransactOpts, initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Flashbuy *FlashbuyTransactorSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Flashbuy.Contract.OnFlashLoan(&_Flashbuy.TransactOpts, initiator, token, amount, fee, data)
}
