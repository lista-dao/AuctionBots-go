// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ierc3156flashborrower

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

// Ierc3156flashborrowerMetaData contains all meta data concerning the Ierc3156flashborrower contract.
var Ierc3156flashborrowerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onFlashLoan\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Ierc3156flashborrowerABI is the input ABI used to generate the binding from.
// Deprecated: Use Ierc3156flashborrowerMetaData.ABI instead.
var Ierc3156flashborrowerABI = Ierc3156flashborrowerMetaData.ABI

// Ierc3156flashborrower is an auto generated Go binding around an Ethereum contract.
type Ierc3156flashborrower struct {
	Ierc3156flashborrowerCaller     // Read-only binding to the contract
	Ierc3156flashborrowerTransactor // Write-only binding to the contract
	Ierc3156flashborrowerFilterer   // Log filterer for contract events
}

// Ierc3156flashborrowerCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ierc3156flashborrowerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc3156flashborrowerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ierc3156flashborrowerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc3156flashborrowerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ierc3156flashborrowerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ierc3156flashborrowerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ierc3156flashborrowerSession struct {
	Contract     *Ierc3156flashborrower // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Ierc3156flashborrowerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ierc3156flashborrowerCallerSession struct {
	Contract *Ierc3156flashborrowerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// Ierc3156flashborrowerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ierc3156flashborrowerTransactorSession struct {
	Contract     *Ierc3156flashborrowerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// Ierc3156flashborrowerRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ierc3156flashborrowerRaw struct {
	Contract *Ierc3156flashborrower // Generic contract binding to access the raw methods on
}

// Ierc3156flashborrowerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ierc3156flashborrowerCallerRaw struct {
	Contract *Ierc3156flashborrowerCaller // Generic read-only contract binding to access the raw methods on
}

// Ierc3156flashborrowerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ierc3156flashborrowerTransactorRaw struct {
	Contract *Ierc3156flashborrowerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIerc3156flashborrower creates a new instance of Ierc3156flashborrower, bound to a specific deployed contract.
func NewIerc3156flashborrower(address common.Address, backend bind.ContractBackend) (*Ierc3156flashborrower, error) {
	contract, err := bindIerc3156flashborrower(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ierc3156flashborrower{Ierc3156flashborrowerCaller: Ierc3156flashborrowerCaller{contract: contract}, Ierc3156flashborrowerTransactor: Ierc3156flashborrowerTransactor{contract: contract}, Ierc3156flashborrowerFilterer: Ierc3156flashborrowerFilterer{contract: contract}}, nil
}

// NewIerc3156flashborrowerCaller creates a new read-only instance of Ierc3156flashborrower, bound to a specific deployed contract.
func NewIerc3156flashborrowerCaller(address common.Address, caller bind.ContractCaller) (*Ierc3156flashborrowerCaller, error) {
	contract, err := bindIerc3156flashborrower(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc3156flashborrowerCaller{contract: contract}, nil
}

// NewIerc3156flashborrowerTransactor creates a new write-only instance of Ierc3156flashborrower, bound to a specific deployed contract.
func NewIerc3156flashborrowerTransactor(address common.Address, transactor bind.ContractTransactor) (*Ierc3156flashborrowerTransactor, error) {
	contract, err := bindIerc3156flashborrower(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ierc3156flashborrowerTransactor{contract: contract}, nil
}

// NewIerc3156flashborrowerFilterer creates a new log filterer instance of Ierc3156flashborrower, bound to a specific deployed contract.
func NewIerc3156flashborrowerFilterer(address common.Address, filterer bind.ContractFilterer) (*Ierc3156flashborrowerFilterer, error) {
	contract, err := bindIerc3156flashborrower(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ierc3156flashborrowerFilterer{contract: contract}, nil
}

// bindIerc3156flashborrower binds a generic wrapper to an already deployed contract.
func bindIerc3156flashborrower(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ierc3156flashborrowerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc3156flashborrower *Ierc3156flashborrowerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc3156flashborrower.Contract.Ierc3156flashborrowerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc3156flashborrower *Ierc3156flashborrowerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc3156flashborrower.Contract.Ierc3156flashborrowerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc3156flashborrower *Ierc3156flashborrowerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc3156flashborrower.Contract.Ierc3156flashborrowerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ierc3156flashborrower *Ierc3156flashborrowerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ierc3156flashborrower.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ierc3156flashborrower *Ierc3156flashborrowerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ierc3156flashborrower.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ierc3156flashborrower *Ierc3156flashborrowerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ierc3156flashborrower.Contract.contract.Transact(opts, method, params...)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Ierc3156flashborrower *Ierc3156flashborrowerTransactor) OnFlashLoan(opts *bind.TransactOpts, initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Ierc3156flashborrower.contract.Transact(opts, "onFlashLoan", initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Ierc3156flashborrower *Ierc3156flashborrowerSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Ierc3156flashborrower.Contract.OnFlashLoan(&_Ierc3156flashborrower.TransactOpts, initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Ierc3156flashborrower *Ierc3156flashborrowerTransactorSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Ierc3156flashborrower.Contract.OnFlashLoan(&_Ierc3156flashborrower.TransactOpts, initiator, token, amount, fee, data)
}
