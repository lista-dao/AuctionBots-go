// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iauctionproxy

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

// IauctionproxyMetaData contains all meta data concerning the Iauctionproxy contract.
var IauctionproxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"}],\"name\":\"buyFromAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IauctionproxyABI is the input ABI used to generate the binding from.
// Deprecated: Use IauctionproxyMetaData.ABI instead.
var IauctionproxyABI = IauctionproxyMetaData.ABI

// Iauctionproxy is an auto generated Go binding around an Ethereum contract.
type Iauctionproxy struct {
	IauctionproxyCaller     // Read-only binding to the contract
	IauctionproxyTransactor // Write-only binding to the contract
	IauctionproxyFilterer   // Log filterer for contract events
}

// IauctionproxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IauctionproxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IauctionproxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IauctionproxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IauctionproxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IauctionproxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IauctionproxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IauctionproxySession struct {
	Contract     *Iauctionproxy    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IauctionproxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IauctionproxyCallerSession struct {
	Contract *IauctionproxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IauctionproxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IauctionproxyTransactorSession struct {
	Contract     *IauctionproxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IauctionproxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IauctionproxyRaw struct {
	Contract *Iauctionproxy // Generic contract binding to access the raw methods on
}

// IauctionproxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IauctionproxyCallerRaw struct {
	Contract *IauctionproxyCaller // Generic read-only contract binding to access the raw methods on
}

// IauctionproxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IauctionproxyTransactorRaw struct {
	Contract *IauctionproxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIauctionproxy creates a new instance of Iauctionproxy, bound to a specific deployed contract.
func NewIauctionproxy(address common.Address, backend bind.ContractBackend) (*Iauctionproxy, error) {
	contract, err := bindIauctionproxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iauctionproxy{IauctionproxyCaller: IauctionproxyCaller{contract: contract}, IauctionproxyTransactor: IauctionproxyTransactor{contract: contract}, IauctionproxyFilterer: IauctionproxyFilterer{contract: contract}}, nil
}

// NewIauctionproxyCaller creates a new read-only instance of Iauctionproxy, bound to a specific deployed contract.
func NewIauctionproxyCaller(address common.Address, caller bind.ContractCaller) (*IauctionproxyCaller, error) {
	contract, err := bindIauctionproxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IauctionproxyCaller{contract: contract}, nil
}

// NewIauctionproxyTransactor creates a new write-only instance of Iauctionproxy, bound to a specific deployed contract.
func NewIauctionproxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IauctionproxyTransactor, error) {
	contract, err := bindIauctionproxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IauctionproxyTransactor{contract: contract}, nil
}

// NewIauctionproxyFilterer creates a new log filterer instance of Iauctionproxy, bound to a specific deployed contract.
func NewIauctionproxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IauctionproxyFilterer, error) {
	contract, err := bindIauctionproxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IauctionproxyFilterer{contract: contract}, nil
}

// bindIauctionproxy binds a generic wrapper to an already deployed contract.
func bindIauctionproxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IauctionproxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iauctionproxy *IauctionproxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iauctionproxy.Contract.IauctionproxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iauctionproxy *IauctionproxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iauctionproxy.Contract.IauctionproxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iauctionproxy *IauctionproxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iauctionproxy.Contract.IauctionproxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iauctionproxy *IauctionproxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iauctionproxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iauctionproxy *IauctionproxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iauctionproxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iauctionproxy *IauctionproxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iauctionproxy.Contract.contract.Transact(opts, method, params...)
}

// BuyFromAuction is a paid mutator transaction binding the contract method 0x71e880ea.
//
// Solidity: function buyFromAuction(address user, uint256 auctionId, uint256 collateralAmount, uint256 maxPrice, address receiverAddress) returns()
func (_Iauctionproxy *IauctionproxyTransactor) BuyFromAuction(opts *bind.TransactOpts, user common.Address, auctionId *big.Int, collateralAmount *big.Int, maxPrice *big.Int, receiverAddress common.Address) (*types.Transaction, error) {
	return _Iauctionproxy.contract.Transact(opts, "buyFromAuction", user, auctionId, collateralAmount, maxPrice, receiverAddress)
}

// BuyFromAuction is a paid mutator transaction binding the contract method 0x71e880ea.
//
// Solidity: function buyFromAuction(address user, uint256 auctionId, uint256 collateralAmount, uint256 maxPrice, address receiverAddress) returns()
func (_Iauctionproxy *IauctionproxySession) BuyFromAuction(user common.Address, auctionId *big.Int, collateralAmount *big.Int, maxPrice *big.Int, receiverAddress common.Address) (*types.Transaction, error) {
	return _Iauctionproxy.Contract.BuyFromAuction(&_Iauctionproxy.TransactOpts, user, auctionId, collateralAmount, maxPrice, receiverAddress)
}

// BuyFromAuction is a paid mutator transaction binding the contract method 0x71e880ea.
//
// Solidity: function buyFromAuction(address user, uint256 auctionId, uint256 collateralAmount, uint256 maxPrice, address receiverAddress) returns()
func (_Iauctionproxy *IauctionproxyTransactorSession) BuyFromAuction(user common.Address, auctionId *big.Int, collateralAmount *big.Int, maxPrice *big.Int, receiverAddress common.Address) (*types.Transaction, error) {
	return _Iauctionproxy.Contract.BuyFromAuction(&_Iauctionproxy.TransactOpts, user, auctionId, collateralAmount, maxPrice, receiverAddress)
}
