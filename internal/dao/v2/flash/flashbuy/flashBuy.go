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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SLIPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dex\",\"outputs\":[{\"internalType\":\"contractIDEX\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAm\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralAm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slip\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralReal\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"}],\"name\":\"flashBuyAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashLender\",\"name\":\"lender_\",\"type\":\"address\"},{\"internalType\":\"contractIInteraction\",\"name\":\"interaction_\",\"type\":\"address\"},{\"internalType\":\"contractIDEX\",\"name\":\"dex_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interaction\",\"outputs\":[{\"internalType\":\"contractIInteraction\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"contractIERC3156FlashLender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onFlashLoan\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// MAXSLIPE is a free data retrieval call binding the contract method 0xc966d05c.
//
// Solidity: function MAX_SLIPE() view returns(uint256)
func (_Flashbuy *FlashbuyCaller) MAXSLIPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Flashbuy.contract.Call(opts, &out, "MAX_SLIPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSLIPE is a free data retrieval call binding the contract method 0xc966d05c.
//
// Solidity: function MAX_SLIPE() view returns(uint256)
func (_Flashbuy *FlashbuySession) MAXSLIPE() (*big.Int, error) {
	return _Flashbuy.Contract.MAXSLIPE(&_Flashbuy.CallOpts)
}

// MAXSLIPE is a free data retrieval call binding the contract method 0xc966d05c.
//
// Solidity: function MAX_SLIPE() view returns(uint256)
func (_Flashbuy *FlashbuyCallerSession) MAXSLIPE() (*big.Int, error) {
	return _Flashbuy.Contract.MAXSLIPE(&_Flashbuy.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_Flashbuy *FlashbuyCaller) Dex(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flashbuy.contract.Call(opts, &out, "dex")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_Flashbuy *FlashbuySession) Dex() (common.Address, error) {
	return _Flashbuy.Contract.Dex(&_Flashbuy.CallOpts)
}

// Dex is a free data retrieval call binding the contract method 0x692058c2.
//
// Solidity: function dex() view returns(address)
func (_Flashbuy *FlashbuyCallerSession) Dex() (common.Address, error) {
	return _Flashbuy.Contract.Dex(&_Flashbuy.CallOpts)
}

// Interaction is a free data retrieval call binding the contract method 0x6c697331.
//
// Solidity: function interaction() view returns(address)
func (_Flashbuy *FlashbuyCaller) Interaction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flashbuy.contract.Call(opts, &out, "interaction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Interaction is a free data retrieval call binding the contract method 0x6c697331.
//
// Solidity: function interaction() view returns(address)
func (_Flashbuy *FlashbuySession) Interaction() (common.Address, error) {
	return _Flashbuy.Contract.Interaction(&_Flashbuy.CallOpts)
}

// Interaction is a free data retrieval call binding the contract method 0x6c697331.
//
// Solidity: function interaction() view returns(address)
func (_Flashbuy *FlashbuyCallerSession) Interaction() (common.Address, error) {
	return _Flashbuy.Contract.Interaction(&_Flashbuy.CallOpts)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Flashbuy *FlashbuyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Flashbuy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Flashbuy *FlashbuySession) Owner() (common.Address, error) {
	return _Flashbuy.Contract.Owner(&_Flashbuy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Flashbuy *FlashbuyCallerSession) Owner() (common.Address, error) {
	return _Flashbuy.Contract.Owner(&_Flashbuy.CallOpts)
}

// FlashBuyAuction is a paid mutator transaction binding the contract method 0x3bfdba5e.
//
// Solidity: function flashBuyAuction(address token, uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, uint256 slip, address collateralReal, bytes path) returns()
func (_Flashbuy *FlashbuyTransactor) FlashBuyAuction(opts *bind.TransactOpts, token common.Address, auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, slip *big.Int, collateralReal common.Address, path []byte) (*types.Transaction, error) {
	return _Flashbuy.contract.Transact(opts, "flashBuyAuction", token, auctionId, borrowAm, collateral, collateralAm, maxPrice, slip, collateralReal, path)
}

// FlashBuyAuction is a paid mutator transaction binding the contract method 0x3bfdba5e.
//
// Solidity: function flashBuyAuction(address token, uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, uint256 slip, address collateralReal, bytes path) returns()
func (_Flashbuy *FlashbuySession) FlashBuyAuction(token common.Address, auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, slip *big.Int, collateralReal common.Address, path []byte) (*types.Transaction, error) {
	return _Flashbuy.Contract.FlashBuyAuction(&_Flashbuy.TransactOpts, token, auctionId, borrowAm, collateral, collateralAm, maxPrice, slip, collateralReal, path)
}

// FlashBuyAuction is a paid mutator transaction binding the contract method 0x3bfdba5e.
//
// Solidity: function flashBuyAuction(address token, uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, uint256 slip, address collateralReal, bytes path) returns()
func (_Flashbuy *FlashbuyTransactorSession) FlashBuyAuction(token common.Address, auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, slip *big.Int, collateralReal common.Address, path []byte) (*types.Transaction, error) {
	return _Flashbuy.Contract.FlashBuyAuction(&_Flashbuy.TransactOpts, token, auctionId, borrowAm, collateral, collateralAm, maxPrice, slip, collateralReal, path)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address lender_, address interaction_, address dex_) returns()
func (_Flashbuy *FlashbuyTransactor) Initialize(opts *bind.TransactOpts, lender_ common.Address, interaction_ common.Address, dex_ common.Address) (*types.Transaction, error) {
	return _Flashbuy.contract.Transact(opts, "initialize", lender_, interaction_, dex_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address lender_, address interaction_, address dex_) returns()
func (_Flashbuy *FlashbuySession) Initialize(lender_ common.Address, interaction_ common.Address, dex_ common.Address) (*types.Transaction, error) {
	return _Flashbuy.Contract.Initialize(&_Flashbuy.TransactOpts, lender_, interaction_, dex_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address lender_, address interaction_, address dex_) returns()
func (_Flashbuy *FlashbuyTransactorSession) Initialize(lender_ common.Address, interaction_ common.Address, dex_ common.Address) (*types.Transaction, error) {
	return _Flashbuy.Contract.Initialize(&_Flashbuy.TransactOpts, lender_, interaction_, dex_)
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

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Flashbuy *FlashbuyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Flashbuy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Flashbuy *FlashbuySession) RenounceOwnership() (*types.Transaction, error) {
	return _Flashbuy.Contract.RenounceOwnership(&_Flashbuy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Flashbuy *FlashbuyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Flashbuy.Contract.RenounceOwnership(&_Flashbuy.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address token) returns()
func (_Flashbuy *FlashbuyTransactor) Transfer(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Flashbuy.contract.Transact(opts, "transfer", token)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address token) returns()
func (_Flashbuy *FlashbuySession) Transfer(token common.Address) (*types.Transaction, error) {
	return _Flashbuy.Contract.Transfer(&_Flashbuy.TransactOpts, token)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address token) returns()
func (_Flashbuy *FlashbuyTransactorSession) Transfer(token common.Address) (*types.Transaction, error) {
	return _Flashbuy.Contract.Transfer(&_Flashbuy.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Flashbuy *FlashbuyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Flashbuy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Flashbuy *FlashbuySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Flashbuy.Contract.TransferOwnership(&_Flashbuy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Flashbuy *FlashbuyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Flashbuy.Contract.TransferOwnership(&_Flashbuy.TransactOpts, newOwner)
}

// FlashbuyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Flashbuy contract.
type FlashbuyInitializedIterator struct {
	Event *FlashbuyInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlashbuyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashbuyInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlashbuyInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlashbuyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashbuyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashbuyInitialized represents a Initialized event raised by the Flashbuy contract.
type FlashbuyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Flashbuy *FlashbuyFilterer) FilterInitialized(opts *bind.FilterOpts) (*FlashbuyInitializedIterator, error) {

	logs, sub, err := _Flashbuy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FlashbuyInitializedIterator{contract: _Flashbuy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Flashbuy *FlashbuyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FlashbuyInitialized) (event.Subscription, error) {

	logs, sub, err := _Flashbuy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashbuyInitialized)
				if err := _Flashbuy.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Flashbuy *FlashbuyFilterer) ParseInitialized(log types.Log) (*FlashbuyInitialized, error) {
	event := new(FlashbuyInitialized)
	if err := _Flashbuy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FlashbuyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Flashbuy contract.
type FlashbuyOwnershipTransferredIterator struct {
	Event *FlashbuyOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FlashbuyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FlashbuyOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FlashbuyOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FlashbuyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FlashbuyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FlashbuyOwnershipTransferred represents a OwnershipTransferred event raised by the Flashbuy contract.
type FlashbuyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Flashbuy *FlashbuyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FlashbuyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Flashbuy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FlashbuyOwnershipTransferredIterator{contract: _Flashbuy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Flashbuy *FlashbuyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FlashbuyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Flashbuy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FlashbuyOwnershipTransferred)
				if err := _Flashbuy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Flashbuy *FlashbuyFilterer) ParseOwnershipTransferred(log types.Log) (*FlashbuyOwnershipTransferred, error) {
	event := new(FlashbuyOwnershipTransferred)
	if err := _Flashbuy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
