// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc3156flashlender

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

// Ierc3156flashlenderMetaData contains all meta data concerning the Ierc3156flashlender contract.
var Ierc3156flashlenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Ierc3156flashlenderABI is the input ABI used to generate the binding from.
// Deprecated: Use Ierc3156flashlenderMetaData.ABI instead.
var Ierc3156flashlenderABI = Ierc3156flashlenderMetaData.ABI

// Ierc3156flashlender is an auto generated Go binding around an Ethereum contract.
type Ierc3156flashlender struct {
	Ierc3156flashlenderCaller     // Read-only binding to the contract
	Ierc3156flashlenderTransactor // Write-only binding to the contract
	Ierc3156flashlenderFilterer   // Log filterer for contract events
}

// Ierc3156flashlenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ierc3156flashlenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc3156flashlenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ierc3156flashlenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc3156flashlenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ierc3156flashlenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc3156flashlenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ierc3156flashlenderSession struct {
	Contract     *Ierc3156flashlender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Ierc3156flashlenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ierc3156flashlenderCallerSession struct {
	Contract *Ierc3156flashlenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// Ierc3156flashlenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ierc3156flashlenderTransactorSession struct {
	Contract     *Ierc3156flashlenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// Ierc3156flashlenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ierc3156flashlenderRaw struct {
	Contract *Ierc3156flashlender // Generic contract binding to access the raw methods on
}

// Ierc3156flashlenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ierc3156flashlenderCallerRaw struct {
	Contract *Ierc3156flashlenderCaller // Generic read-only contract binding to access the raw methods on
}

// Ierc3156flashlenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ierc3156flashlenderTransactorRaw struct {
	Contract *Ierc3156flashlenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIerc3156flashlender creates a new instance of Ierc3156flashlender, bound to a specific deployed contract.
func NewIerc3156flashlender(address common.Address, backend bind.ContractBackend) (*Ierc3156flashlender, error) {
	contract, err := bindIerc3156flashlender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ierc3156flashlender{Ierc3156flashlenderCaller: Ierc3156flashlenderCaller{contract: contract}, Ierc3156flashlenderTransactor: Ierc3156flashlenderTransactor{contract: contract}, Ierc3156flashlenderFilterer: Ierc3156flashlenderFilterer{contract: contract}}, nil
}

// NewIerc3156flashlenderCaller creates a new read-only instance of Ierc3156flashlender, bound to a specific deployed contract.
func NewIerc3156flashlenderCaller(address common.Address, caller bind.ContractCaller) (*Ierc3156flashlenderCaller, error) {
	contract, err := bindIerc3156flashlender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc3156flashlenderCaller{contract: contract}, nil
}

// NewIerc3156flashlenderTransactor creates a new write-only instance of Ierc3156flashlender, bound to a specific deployed contract.
func NewIerc3156flashlenderTransactor(address common.Address, transactor bind.ContractTransactor) (*Ierc3156flashlenderTransactor, error) {
	contract, err := bindIerc3156flashlender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc3156flashlenderTransactor{contract: contract}, nil
}

// NewIerc3156flashlenderFilterer creates a new log filterer instance of Ierc3156flashlender, bound to a specific deployed contract.
func NewIerc3156flashlenderFilterer(address common.Address, filterer bind.ContractFilterer) (*Ierc3156flashlenderFilterer, error) {
	contract, err := bindIerc3156flashlender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ierc3156flashlenderFilterer{contract: contract}, nil
}

// bindIerc3156flashlender binds a generic wrapper to an already deployed contract.
func bindIerc3156flashlender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ierc3156flashlenderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc3156flashlender *Ierc3156flashlenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc3156flashlender.Contract.Ierc3156flashlenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc3156flashlender *Ierc3156flashlenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc3156flashlender.Contract.Ierc3156flashlenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc3156flashlender *Ierc3156flashlenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc3156flashlender.Contract.Ierc3156flashlenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc3156flashlender *Ierc3156flashlenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc3156flashlender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc3156flashlender *Ierc3156flashlenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc3156flashlender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc3156flashlender *Ierc3156flashlenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc3156flashlender.Contract.contract.Transact(opts, method, params...)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_Ierc3156flashlender *Ierc3156flashlenderCaller) FlashFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Ierc3156flashlender.contract.Call(opts, &out, "flashFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_Ierc3156flashlender *Ierc3156flashlenderSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Ierc3156flashlender.Contract.FlashFee(&_Ierc3156flashlender.CallOpts, token, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address token, uint256 amount) view returns(uint256)
func (_Ierc3156flashlender *Ierc3156flashlenderCallerSession) FlashFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Ierc3156flashlender.Contract.FlashFee(&_Ierc3156flashlender.CallOpts, token, amount)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Ierc3156flashlender *Ierc3156flashlenderCaller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ierc3156flashlender.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Ierc3156flashlender *Ierc3156flashlenderSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _Ierc3156flashlender.Contract.MaxFlashLoan(&_Ierc3156flashlender.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Ierc3156flashlender *Ierc3156flashlenderCallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _Ierc3156flashlender.Contract.MaxFlashLoan(&_Ierc3156flashlender.CallOpts, token)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_Ierc3156flashlender *Ierc3156flashlenderTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Ierc3156flashlender.contract.Transact(opts, "flashLoan", receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_Ierc3156flashlender *Ierc3156flashlenderSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Ierc3156flashlender.Contract.FlashLoan(&_Ierc3156flashlender.TransactOpts, receiver, token, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes data) returns(bool)
func (_Ierc3156flashlender *Ierc3156flashlenderTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Ierc3156flashlender.Contract.FlashLoan(&_Ierc3156flashlender.TransactOpts, receiver, token, amount, data)
}
