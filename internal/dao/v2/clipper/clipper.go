// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clipper

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

// Sale is an auto generated low-level Go binding around an user-defined struct.
type Sale struct {
	Pos *big.Int
	Tab *big.Int
	Lot *big.Int
	Usr common.Address
	Tic *big.Int
	Top *big.Int
}

// ClipperMetaData contains all meta data concerning the Clipper contract.
var ClipperMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"File\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"top\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kpr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coin\",\"type\":\"uint256\"}],\"name\":\"Kick\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"top\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"kpr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"coin\",\"type\":\"uint256\"}],\"name\":\"Redo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Rely\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"owe\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"Take\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Yank\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"active\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calc\",\"outputs\":[{\"internalType\":\"contractAbacusLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chip\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cusp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dog\",\"outputs\":[{\"internalType\":\"contractDogLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"needsRedo\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ilk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spotter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dog_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk_\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kpr\",\"type\":\"address\"}],\"name\":\"kick\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kicks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"kpr\",\"type\":\"address\"}],\"name\":\"redo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"sales\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pos\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lot\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"tic\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"top\",\"type\":\"uint256\"}],\"internalType\":\"structSale\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spotter\",\"outputs\":[{\"internalType\":\"contractSpotterLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tail\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"take\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tip\",\"outputs\":[{\"internalType\":\"uint192\",\"name\":\"\",\"type\":\"uint192\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upchost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vow\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"yank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ClipperABI is the input ABI used to generate the binding from.
// Deprecated: Use ClipperMetaData.ABI instead.
var ClipperABI = ClipperMetaData.ABI

// Clipper is an auto generated Go binding around an Ethereum contract.
type Clipper struct {
	ClipperCaller     // Read-only binding to the contract
	ClipperTransactor // Write-only binding to the contract
	ClipperFilterer   // Log filterer for contract events
}

// ClipperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClipperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClipperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClipperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClipperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClipperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClipperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClipperSession struct {
	Contract     *Clipper          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClipperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClipperCallerSession struct {
	Contract *ClipperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ClipperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClipperTransactorSession struct {
	Contract     *ClipperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ClipperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClipperRaw struct {
	Contract *Clipper // Generic contract binding to access the raw methods on
}

// ClipperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClipperCallerRaw struct {
	Contract *ClipperCaller // Generic read-only contract binding to access the raw methods on
}

// ClipperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClipperTransactorRaw struct {
	Contract *ClipperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClipper creates a new instance of Clipper, bound to a specific deployed contract.
func NewClipper(address common.Address, backend bind.ContractBackend) (*Clipper, error) {
	contract, err := bindClipper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clipper{ClipperCaller: ClipperCaller{contract: contract}, ClipperTransactor: ClipperTransactor{contract: contract}, ClipperFilterer: ClipperFilterer{contract: contract}}, nil
}

// NewClipperCaller creates a new read-only instance of Clipper, bound to a specific deployed contract.
func NewClipperCaller(address common.Address, caller bind.ContractCaller) (*ClipperCaller, error) {
	contract, err := bindClipper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClipperCaller{contract: contract}, nil
}

// NewClipperTransactor creates a new write-only instance of Clipper, bound to a specific deployed contract.
func NewClipperTransactor(address common.Address, transactor bind.ContractTransactor) (*ClipperTransactor, error) {
	contract, err := bindClipper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClipperTransactor{contract: contract}, nil
}

// NewClipperFilterer creates a new log filterer instance of Clipper, bound to a specific deployed contract.
func NewClipperFilterer(address common.Address, filterer bind.ContractFilterer) (*ClipperFilterer, error) {
	contract, err := bindClipper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClipperFilterer{contract: contract}, nil
}

// bindClipper binds a generic wrapper to an already deployed contract.
func bindClipper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClipperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clipper *ClipperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clipper.Contract.ClipperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clipper *ClipperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clipper.Contract.ClipperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clipper *ClipperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clipper.Contract.ClipperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clipper *ClipperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clipper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clipper *ClipperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clipper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clipper *ClipperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clipper.Contract.contract.Transact(opts, method, params...)
}

// Active is a free data retrieval call binding the contract method 0x8033d581.
//
// Solidity: function active(uint256 ) view returns(uint256)
func (_Clipper *ClipperCaller) Active(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "active", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Active is a free data retrieval call binding the contract method 0x8033d581.
//
// Solidity: function active(uint256 ) view returns(uint256)
func (_Clipper *ClipperSession) Active(arg0 *big.Int) (*big.Int, error) {
	return _Clipper.Contract.Active(&_Clipper.CallOpts, arg0)
}

// Active is a free data retrieval call binding the contract method 0x8033d581.
//
// Solidity: function active(uint256 ) view returns(uint256)
func (_Clipper *ClipperCallerSession) Active(arg0 *big.Int) (*big.Int, error) {
	return _Clipper.Contract.Active(&_Clipper.CallOpts, arg0)
}

// Buf is a free data retrieval call binding the contract method 0x15232515.
//
// Solidity: function buf() view returns(uint256)
func (_Clipper *ClipperCaller) Buf(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "buf")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buf is a free data retrieval call binding the contract method 0x15232515.
//
// Solidity: function buf() view returns(uint256)
func (_Clipper *ClipperSession) Buf() (*big.Int, error) {
	return _Clipper.Contract.Buf(&_Clipper.CallOpts)
}

// Buf is a free data retrieval call binding the contract method 0x15232515.
//
// Solidity: function buf() view returns(uint256)
func (_Clipper *ClipperCallerSession) Buf() (*big.Int, error) {
	return _Clipper.Contract.Buf(&_Clipper.CallOpts)
}

// Calc is a free data retrieval call binding the contract method 0x96f1b6be.
//
// Solidity: function calc() view returns(address)
func (_Clipper *ClipperCaller) Calc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "calc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Calc is a free data retrieval call binding the contract method 0x96f1b6be.
//
// Solidity: function calc() view returns(address)
func (_Clipper *ClipperSession) Calc() (common.Address, error) {
	return _Clipper.Contract.Calc(&_Clipper.CallOpts)
}

// Calc is a free data retrieval call binding the contract method 0x96f1b6be.
//
// Solidity: function calc() view returns(address)
func (_Clipper *ClipperCallerSession) Calc() (common.Address, error) {
	return _Clipper.Contract.Calc(&_Clipper.CallOpts)
}

// Chip is a free data retrieval call binding the contract method 0xb61500e4.
//
// Solidity: function chip() view returns(uint64)
func (_Clipper *ClipperCaller) Chip(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "chip")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Chip is a free data retrieval call binding the contract method 0xb61500e4.
//
// Solidity: function chip() view returns(uint64)
func (_Clipper *ClipperSession) Chip() (uint64, error) {
	return _Clipper.Contract.Chip(&_Clipper.CallOpts)
}

// Chip is a free data retrieval call binding the contract method 0xb61500e4.
//
// Solidity: function chip() view returns(uint64)
func (_Clipper *ClipperCallerSession) Chip() (uint64, error) {
	return _Clipper.Contract.Chip(&_Clipper.CallOpts)
}

// Chost is a free data retrieval call binding the contract method 0xba2cdc75.
//
// Solidity: function chost() view returns(uint256)
func (_Clipper *ClipperCaller) Chost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "chost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Chost is a free data retrieval call binding the contract method 0xba2cdc75.
//
// Solidity: function chost() view returns(uint256)
func (_Clipper *ClipperSession) Chost() (*big.Int, error) {
	return _Clipper.Contract.Chost(&_Clipper.CallOpts)
}

// Chost is a free data retrieval call binding the contract method 0xba2cdc75.
//
// Solidity: function chost() view returns(uint256)
func (_Clipper *ClipperCallerSession) Chost() (*big.Int, error) {
	return _Clipper.Contract.Chost(&_Clipper.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Clipper *ClipperCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Clipper *ClipperSession) Count() (*big.Int, error) {
	return _Clipper.Contract.Count(&_Clipper.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_Clipper *ClipperCallerSession) Count() (*big.Int, error) {
	return _Clipper.Contract.Count(&_Clipper.CallOpts)
}

// Cusp is a free data retrieval call binding the contract method 0x49ed5931.
//
// Solidity: function cusp() view returns(uint256)
func (_Clipper *ClipperCaller) Cusp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "cusp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cusp is a free data retrieval call binding the contract method 0x49ed5931.
//
// Solidity: function cusp() view returns(uint256)
func (_Clipper *ClipperSession) Cusp() (*big.Int, error) {
	return _Clipper.Contract.Cusp(&_Clipper.CallOpts)
}

// Cusp is a free data retrieval call binding the contract method 0x49ed5931.
//
// Solidity: function cusp() view returns(uint256)
func (_Clipper *ClipperCallerSession) Cusp() (*big.Int, error) {
	return _Clipper.Contract.Cusp(&_Clipper.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Clipper *ClipperCaller) Dog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "dog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Clipper *ClipperSession) Dog() (common.Address, error) {
	return _Clipper.Contract.Dog(&_Clipper.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Clipper *ClipperCallerSession) Dog() (common.Address, error) {
	return _Clipper.Contract.Dog(&_Clipper.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 id) view returns(bool needsRedo, uint256 price, uint256 lot, uint256 tab)
func (_Clipper *ClipperCaller) GetStatus(opts *bind.CallOpts, id *big.Int) (struct {
	NeedsRedo bool
	Price     *big.Int
	Lot       *big.Int
	Tab       *big.Int
}, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "getStatus", id)

	outstruct := new(struct {
		NeedsRedo bool
		Price     *big.Int
		Lot       *big.Int
		Tab       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NeedsRedo = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Lot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Tab = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 id) view returns(bool needsRedo, uint256 price, uint256 lot, uint256 tab)
func (_Clipper *ClipperSession) GetStatus(id *big.Int) (struct {
	NeedsRedo bool
	Price     *big.Int
	Lot       *big.Int
	Tab       *big.Int
}, error) {
	return _Clipper.Contract.GetStatus(&_Clipper.CallOpts, id)
}

// GetStatus is a free data retrieval call binding the contract method 0x5c622a0e.
//
// Solidity: function getStatus(uint256 id) view returns(bool needsRedo, uint256 price, uint256 lot, uint256 tab)
func (_Clipper *ClipperCallerSession) GetStatus(id *big.Int) (struct {
	NeedsRedo bool
	Price     *big.Int
	Lot       *big.Int
	Tab       *big.Int
}, error) {
	return _Clipper.Contract.GetStatus(&_Clipper.CallOpts, id)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_Clipper *ClipperCaller) Ilk(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "ilk")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_Clipper *ClipperSession) Ilk() ([32]byte, error) {
	return _Clipper.Contract.Ilk(&_Clipper.CallOpts)
}

// Ilk is a free data retrieval call binding the contract method 0xc5ce281e.
//
// Solidity: function ilk() view returns(bytes32)
func (_Clipper *ClipperCallerSession) Ilk() ([32]byte, error) {
	return _Clipper.Contract.Ilk(&_Clipper.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_Clipper *ClipperCaller) Kicks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "kicks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_Clipper *ClipperSession) Kicks() (*big.Int, error) {
	return _Clipper.Contract.Kicks(&_Clipper.CallOpts)
}

// Kicks is a free data retrieval call binding the contract method 0xcfdd3302.
//
// Solidity: function kicks() view returns(uint256)
func (_Clipper *ClipperCallerSession) Kicks() (*big.Int, error) {
	return _Clipper.Contract.Kicks(&_Clipper.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(uint256[])
func (_Clipper *ClipperCaller) List(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(uint256[])
func (_Clipper *ClipperSession) List() ([]*big.Int, error) {
	return _Clipper.Contract.List(&_Clipper.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns(uint256[])
func (_Clipper *ClipperCallerSession) List() ([]*big.Int, error) {
	return _Clipper.Contract.List(&_Clipper.CallOpts)
}

// Sales is a free data retrieval call binding the contract method 0xb5f522f7.
//
// Solidity: function sales(uint256 id) view returns((uint256,uint256,uint256,address,uint96,uint256))
func (_Clipper *ClipperCaller) Sales(opts *bind.CallOpts, id *big.Int) (Sale, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "sales", id)

	if err != nil {
		return *new(Sale), err
	}

	out0 := *abi.ConvertType(out[0], new(Sale)).(*Sale)

	return out0, err

}

// Sales is a free data retrieval call binding the contract method 0xb5f522f7.
//
// Solidity: function sales(uint256 id) view returns((uint256,uint256,uint256,address,uint96,uint256))
func (_Clipper *ClipperSession) Sales(id *big.Int) (Sale, error) {
	return _Clipper.Contract.Sales(&_Clipper.CallOpts, id)
}

// Sales is a free data retrieval call binding the contract method 0xb5f522f7.
//
// Solidity: function sales(uint256 id) view returns((uint256,uint256,uint256,address,uint96,uint256))
func (_Clipper *ClipperCallerSession) Sales(id *big.Int) (Sale, error) {
	return _Clipper.Contract.Sales(&_Clipper.CallOpts, id)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Clipper *ClipperCaller) Spotter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "spotter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Clipper *ClipperSession) Spotter() (common.Address, error) {
	return _Clipper.Contract.Spotter(&_Clipper.CallOpts)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Clipper *ClipperCallerSession) Spotter() (common.Address, error) {
	return _Clipper.Contract.Spotter(&_Clipper.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_Clipper *ClipperCaller) Stopped(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "stopped")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_Clipper *ClipperSession) Stopped() (*big.Int, error) {
	return _Clipper.Contract.Stopped(&_Clipper.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() view returns(uint256)
func (_Clipper *ClipperCallerSession) Stopped() (*big.Int, error) {
	return _Clipper.Contract.Stopped(&_Clipper.CallOpts)
}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(uint256)
func (_Clipper *ClipperCaller) Tail(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "tail")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(uint256)
func (_Clipper *ClipperSession) Tail() (*big.Int, error) {
	return _Clipper.Contract.Tail(&_Clipper.CallOpts)
}

// Tail is a free data retrieval call binding the contract method 0x13d8c840.
//
// Solidity: function tail() view returns(uint256)
func (_Clipper *ClipperCallerSession) Tail() (*big.Int, error) {
	return _Clipper.Contract.Tail(&_Clipper.CallOpts)
}

// Tip is a free data retrieval call binding the contract method 0x2755cd2d.
//
// Solidity: function tip() view returns(uint192)
func (_Clipper *ClipperCaller) Tip(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "tip")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tip is a free data retrieval call binding the contract method 0x2755cd2d.
//
// Solidity: function tip() view returns(uint192)
func (_Clipper *ClipperSession) Tip() (*big.Int, error) {
	return _Clipper.Contract.Tip(&_Clipper.CallOpts)
}

// Tip is a free data retrieval call binding the contract method 0x2755cd2d.
//
// Solidity: function tip() view returns(uint192)
func (_Clipper *ClipperCallerSession) Tip() (*big.Int, error) {
	return _Clipper.Contract.Tip(&_Clipper.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Clipper *ClipperCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Clipper *ClipperSession) Vat() (common.Address, error) {
	return _Clipper.Contract.Vat(&_Clipper.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Clipper *ClipperCallerSession) Vat() (common.Address, error) {
	return _Clipper.Contract.Vat(&_Clipper.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Clipper *ClipperCaller) Vow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "vow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Clipper *ClipperSession) Vow() (common.Address, error) {
	return _Clipper.Contract.Vow(&_Clipper.CallOpts)
}

// Vow is a free data retrieval call binding the contract method 0x626cb3c5.
//
// Solidity: function vow() view returns(address)
func (_Clipper *ClipperCallerSession) Vow() (common.Address, error) {
	return _Clipper.Contract.Vow(&_Clipper.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Clipper *ClipperCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Clipper.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Clipper *ClipperSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Clipper.Contract.Wards(&_Clipper.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Clipper *ClipperCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Clipper.Contract.Wards(&_Clipper.CallOpts, arg0)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Clipper *ClipperTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Clipper *ClipperSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.Deny(&_Clipper.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Clipper *ClipperTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.Deny(&_Clipper.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Clipper *ClipperTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Clipper *ClipperSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Clipper.Contract.File(&_Clipper.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Clipper *ClipperTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Clipper.Contract.File(&_Clipper.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Clipper *ClipperTransactor) File0(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Clipper *ClipperSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.File0(&_Clipper.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Clipper *ClipperTransactorSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.File0(&_Clipper.TransactOpts, what, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x63acc14d.
//
// Solidity: function initialize(address vat_, address spotter_, address dog_, bytes32 ilk_) returns()
func (_Clipper *ClipperTransactor) Initialize(opts *bind.TransactOpts, vat_ common.Address, spotter_ common.Address, dog_ common.Address, ilk_ [32]byte) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "initialize", vat_, spotter_, dog_, ilk_)
}

// Initialize is a paid mutator transaction binding the contract method 0x63acc14d.
//
// Solidity: function initialize(address vat_, address spotter_, address dog_, bytes32 ilk_) returns()
func (_Clipper *ClipperSession) Initialize(vat_ common.Address, spotter_ common.Address, dog_ common.Address, ilk_ [32]byte) (*types.Transaction, error) {
	return _Clipper.Contract.Initialize(&_Clipper.TransactOpts, vat_, spotter_, dog_, ilk_)
}

// Initialize is a paid mutator transaction binding the contract method 0x63acc14d.
//
// Solidity: function initialize(address vat_, address spotter_, address dog_, bytes32 ilk_) returns()
func (_Clipper *ClipperTransactorSession) Initialize(vat_ common.Address, spotter_ common.Address, dog_ common.Address, ilk_ [32]byte) (*types.Transaction, error) {
	return _Clipper.Contract.Initialize(&_Clipper.TransactOpts, vat_, spotter_, dog_, ilk_)
}

// Kick is a paid mutator transaction binding the contract method 0x898eb267.
//
// Solidity: function kick(uint256 tab, uint256 lot, address usr, address kpr) returns(uint256 id)
func (_Clipper *ClipperTransactor) Kick(opts *bind.TransactOpts, tab *big.Int, lot *big.Int, usr common.Address, kpr common.Address) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "kick", tab, lot, usr, kpr)
}

// Kick is a paid mutator transaction binding the contract method 0x898eb267.
//
// Solidity: function kick(uint256 tab, uint256 lot, address usr, address kpr) returns(uint256 id)
func (_Clipper *ClipperSession) Kick(tab *big.Int, lot *big.Int, usr common.Address, kpr common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.Kick(&_Clipper.TransactOpts, tab, lot, usr, kpr)
}

// Kick is a paid mutator transaction binding the contract method 0x898eb267.
//
// Solidity: function kick(uint256 tab, uint256 lot, address usr, address kpr) returns(uint256 id)
func (_Clipper *ClipperTransactorSession) Kick(tab *big.Int, lot *big.Int, usr common.Address, kpr common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.Kick(&_Clipper.TransactOpts, tab, lot, usr, kpr)
}

// Redo is a paid mutator transaction binding the contract method 0xd843416d.
//
// Solidity: function redo(uint256 id, address kpr) returns()
func (_Clipper *ClipperTransactor) Redo(opts *bind.TransactOpts, id *big.Int, kpr common.Address) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "redo", id, kpr)
}

// Redo is a paid mutator transaction binding the contract method 0xd843416d.
//
// Solidity: function redo(uint256 id, address kpr) returns()
func (_Clipper *ClipperSession) Redo(id *big.Int, kpr common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.Redo(&_Clipper.TransactOpts, id, kpr)
}

// Redo is a paid mutator transaction binding the contract method 0xd843416d.
//
// Solidity: function redo(uint256 id, address kpr) returns()
func (_Clipper *ClipperTransactorSession) Redo(id *big.Int, kpr common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.Redo(&_Clipper.TransactOpts, id, kpr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Clipper *ClipperTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Clipper *ClipperSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.Rely(&_Clipper.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Clipper *ClipperTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Clipper.Contract.Rely(&_Clipper.TransactOpts, usr)
}

// Take is a paid mutator transaction binding the contract method 0x81a794cb.
//
// Solidity: function take(uint256 id, uint256 amt, uint256 max, address who, bytes data) returns()
func (_Clipper *ClipperTransactor) Take(opts *bind.TransactOpts, id *big.Int, amt *big.Int, max *big.Int, who common.Address, data []byte) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "take", id, amt, max, who, data)
}

// Take is a paid mutator transaction binding the contract method 0x81a794cb.
//
// Solidity: function take(uint256 id, uint256 amt, uint256 max, address who, bytes data) returns()
func (_Clipper *ClipperSession) Take(id *big.Int, amt *big.Int, max *big.Int, who common.Address, data []byte) (*types.Transaction, error) {
	return _Clipper.Contract.Take(&_Clipper.TransactOpts, id, amt, max, who, data)
}

// Take is a paid mutator transaction binding the contract method 0x81a794cb.
//
// Solidity: function take(uint256 id, uint256 amt, uint256 max, address who, bytes data) returns()
func (_Clipper *ClipperTransactorSession) Take(id *big.Int, amt *big.Int, max *big.Int, who common.Address, data []byte) (*types.Transaction, error) {
	return _Clipper.Contract.Take(&_Clipper.TransactOpts, id, amt, max, who, data)
}

// Upchost is a paid mutator transaction binding the contract method 0x0cbb5862.
//
// Solidity: function upchost() returns()
func (_Clipper *ClipperTransactor) Upchost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "upchost")
}

// Upchost is a paid mutator transaction binding the contract method 0x0cbb5862.
//
// Solidity: function upchost() returns()
func (_Clipper *ClipperSession) Upchost() (*types.Transaction, error) {
	return _Clipper.Contract.Upchost(&_Clipper.TransactOpts)
}

// Upchost is a paid mutator transaction binding the contract method 0x0cbb5862.
//
// Solidity: function upchost() returns()
func (_Clipper *ClipperTransactorSession) Upchost() (*types.Transaction, error) {
	return _Clipper.Contract.Upchost(&_Clipper.TransactOpts)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Clipper *ClipperTransactor) Yank(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Clipper.contract.Transact(opts, "yank", id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Clipper *ClipperSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _Clipper.Contract.Yank(&_Clipper.TransactOpts, id)
}

// Yank is a paid mutator transaction binding the contract method 0x26e027f1.
//
// Solidity: function yank(uint256 id) returns()
func (_Clipper *ClipperTransactorSession) Yank(id *big.Int) (*types.Transaction, error) {
	return _Clipper.Contract.Yank(&_Clipper.TransactOpts, id)
}

// ClipperDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Clipper contract.
type ClipperDenyIterator struct {
	Event *ClipperDeny // Event containing the contract specifics and raw log

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
func (it *ClipperDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClipperDeny)
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
		it.Event = new(ClipperDeny)
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
func (it *ClipperDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClipperDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClipperDeny represents a Deny event raised by the Clipper contract.
type ClipperDeny struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Clipper *ClipperFilterer) FilterDeny(opts *bind.FilterOpts, usr []common.Address) (*ClipperDenyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Clipper.contract.FilterLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return &ClipperDenyIterator{contract: _Clipper.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Clipper *ClipperFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ClipperDeny, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Clipper.contract.WatchLogs(opts, "Deny", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClipperDeny)
				if err := _Clipper.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0x184450df2e323acec0ed3b5c7531b81f9b4cdef7914dfd4c0a4317416bb5251b.
//
// Solidity: event Deny(address indexed usr)
func (_Clipper *ClipperFilterer) ParseDeny(log types.Log) (*ClipperDeny, error) {
	event := new(ClipperDeny)
	if err := _Clipper.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClipperFileIterator is returned from FilterFile and is used to iterate over the raw logs and unpacked data for File events raised by the Clipper contract.
type ClipperFileIterator struct {
	Event *ClipperFile // Event containing the contract specifics and raw log

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
func (it *ClipperFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClipperFile)
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
		it.Event = new(ClipperFile)
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
func (it *ClipperFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClipperFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClipperFile represents a File event raised by the Clipper contract.
type ClipperFile struct {
	What [32]byte
	Data *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile is a free log retrieval operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Clipper *ClipperFilterer) FilterFile(opts *bind.FilterOpts, what [][32]byte) (*ClipperFileIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Clipper.contract.FilterLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return &ClipperFileIterator{contract: _Clipper.contract, event: "File", logs: logs, sub: sub}, nil
}

// WatchFile is a free log subscription operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Clipper *ClipperFilterer) WatchFile(opts *bind.WatchOpts, sink chan<- *ClipperFile, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Clipper.contract.WatchLogs(opts, "File", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClipperFile)
				if err := _Clipper.contract.UnpackLog(event, "File", log); err != nil {
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

// ParseFile is a log parse operation binding the contract event 0xe986e40cc8c151830d4f61050f4fb2e4add8567caad2d5f5496f9158e91fe4c7.
//
// Solidity: event File(bytes32 indexed what, uint256 data)
func (_Clipper *ClipperFilterer) ParseFile(log types.Log) (*ClipperFile, error) {
	event := new(ClipperFile)
	if err := _Clipper.contract.UnpackLog(event, "File", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClipperFile0Iterator is returned from FilterFile0 and is used to iterate over the raw logs and unpacked data for File0 events raised by the Clipper contract.
type ClipperFile0Iterator struct {
	Event *ClipperFile0 // Event containing the contract specifics and raw log

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
func (it *ClipperFile0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClipperFile0)
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
		it.Event = new(ClipperFile0)
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
func (it *ClipperFile0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClipperFile0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClipperFile0 represents a File0 event raised by the Clipper contract.
type ClipperFile0 struct {
	What [32]byte
	Data common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFile0 is a free log retrieval operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Clipper *ClipperFilterer) FilterFile0(opts *bind.FilterOpts, what [][32]byte) (*ClipperFile0Iterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Clipper.contract.FilterLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return &ClipperFile0Iterator{contract: _Clipper.contract, event: "File0", logs: logs, sub: sub}, nil
}

// WatchFile0 is a free log subscription operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Clipper *ClipperFilterer) WatchFile0(opts *bind.WatchOpts, sink chan<- *ClipperFile0, what [][32]byte) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}

	logs, sub, err := _Clipper.contract.WatchLogs(opts, "File0", whatRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClipperFile0)
				if err := _Clipper.contract.UnpackLog(event, "File0", log); err != nil {
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

// ParseFile0 is a log parse operation binding the contract event 0x8fef588b5fc1afbf5b2f06c1a435d513f208da2e6704c3d8f0e0ec91167066ba.
//
// Solidity: event File(bytes32 indexed what, address data)
func (_Clipper *ClipperFilterer) ParseFile0(log types.Log) (*ClipperFile0, error) {
	event := new(ClipperFile0)
	if err := _Clipper.contract.UnpackLog(event, "File0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClipperKickIterator is returned from FilterKick and is used to iterate over the raw logs and unpacked data for Kick events raised by the Clipper contract.
type ClipperKickIterator struct {
	Event *ClipperKick // Event containing the contract specifics and raw log

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
func (it *ClipperKickIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClipperKick)
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
		it.Event = new(ClipperKick)
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
func (it *ClipperKickIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClipperKickIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClipperKick represents a Kick event raised by the Clipper contract.
type ClipperKick struct {
	Id   *big.Int
	Top  *big.Int
	Tab  *big.Int
	Lot  *big.Int
	Usr  common.Address
	Kpr  common.Address
	Coin *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterKick is a free log retrieval operation binding the contract event 0x7c5bfdc0a5e8192f6cd4972f382cec69116862fb62e6abff8003874c58e064b8.
//
// Solidity: event Kick(uint256 indexed id, uint256 top, uint256 tab, uint256 lot, address indexed usr, address indexed kpr, uint256 coin)
func (_Clipper *ClipperFilterer) FilterKick(opts *bind.FilterOpts, id []*big.Int, usr []common.Address, kpr []common.Address) (*ClipperKickIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var kprRule []interface{}
	for _, kprItem := range kpr {
		kprRule = append(kprRule, kprItem)
	}

	logs, sub, err := _Clipper.contract.FilterLogs(opts, "Kick", idRule, usrRule, kprRule)
	if err != nil {
		return nil, err
	}
	return &ClipperKickIterator{contract: _Clipper.contract, event: "Kick", logs: logs, sub: sub}, nil
}

// WatchKick is a free log subscription operation binding the contract event 0x7c5bfdc0a5e8192f6cd4972f382cec69116862fb62e6abff8003874c58e064b8.
//
// Solidity: event Kick(uint256 indexed id, uint256 top, uint256 tab, uint256 lot, address indexed usr, address indexed kpr, uint256 coin)
func (_Clipper *ClipperFilterer) WatchKick(opts *bind.WatchOpts, sink chan<- *ClipperKick, id []*big.Int, usr []common.Address, kpr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var kprRule []interface{}
	for _, kprItem := range kpr {
		kprRule = append(kprRule, kprItem)
	}

	logs, sub, err := _Clipper.contract.WatchLogs(opts, "Kick", idRule, usrRule, kprRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClipperKick)
				if err := _Clipper.contract.UnpackLog(event, "Kick", log); err != nil {
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

// ParseKick is a log parse operation binding the contract event 0x7c5bfdc0a5e8192f6cd4972f382cec69116862fb62e6abff8003874c58e064b8.
//
// Solidity: event Kick(uint256 indexed id, uint256 top, uint256 tab, uint256 lot, address indexed usr, address indexed kpr, uint256 coin)
func (_Clipper *ClipperFilterer) ParseKick(log types.Log) (*ClipperKick, error) {
	event := new(ClipperKick)
	if err := _Clipper.contract.UnpackLog(event, "Kick", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClipperRedoIterator is returned from FilterRedo and is used to iterate over the raw logs and unpacked data for Redo events raised by the Clipper contract.
type ClipperRedoIterator struct {
	Event *ClipperRedo // Event containing the contract specifics and raw log

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
func (it *ClipperRedoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClipperRedo)
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
		it.Event = new(ClipperRedo)
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
func (it *ClipperRedoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClipperRedoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClipperRedo represents a Redo event raised by the Clipper contract.
type ClipperRedo struct {
	Id   *big.Int
	Top  *big.Int
	Tab  *big.Int
	Lot  *big.Int
	Usr  common.Address
	Kpr  common.Address
	Coin *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRedo is a free log retrieval operation binding the contract event 0x275de7ecdd375b5e8049319f8b350686131c219dd4dc450a08e9cf83b03c865f.
//
// Solidity: event Redo(uint256 indexed id, uint256 top, uint256 tab, uint256 lot, address indexed usr, address indexed kpr, uint256 coin)
func (_Clipper *ClipperFilterer) FilterRedo(opts *bind.FilterOpts, id []*big.Int, usr []common.Address, kpr []common.Address) (*ClipperRedoIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var kprRule []interface{}
	for _, kprItem := range kpr {
		kprRule = append(kprRule, kprItem)
	}

	logs, sub, err := _Clipper.contract.FilterLogs(opts, "Redo", idRule, usrRule, kprRule)
	if err != nil {
		return nil, err
	}
	return &ClipperRedoIterator{contract: _Clipper.contract, event: "Redo", logs: logs, sub: sub}, nil
}

// WatchRedo is a free log subscription operation binding the contract event 0x275de7ecdd375b5e8049319f8b350686131c219dd4dc450a08e9cf83b03c865f.
//
// Solidity: event Redo(uint256 indexed id, uint256 top, uint256 tab, uint256 lot, address indexed usr, address indexed kpr, uint256 coin)
func (_Clipper *ClipperFilterer) WatchRedo(opts *bind.WatchOpts, sink chan<- *ClipperRedo, id []*big.Int, usr []common.Address, kpr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}
	var kprRule []interface{}
	for _, kprItem := range kpr {
		kprRule = append(kprRule, kprItem)
	}

	logs, sub, err := _Clipper.contract.WatchLogs(opts, "Redo", idRule, usrRule, kprRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClipperRedo)
				if err := _Clipper.contract.UnpackLog(event, "Redo", log); err != nil {
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

// ParseRedo is a log parse operation binding the contract event 0x275de7ecdd375b5e8049319f8b350686131c219dd4dc450a08e9cf83b03c865f.
//
// Solidity: event Redo(uint256 indexed id, uint256 top, uint256 tab, uint256 lot, address indexed usr, address indexed kpr, uint256 coin)
func (_Clipper *ClipperFilterer) ParseRedo(log types.Log) (*ClipperRedo, error) {
	event := new(ClipperRedo)
	if err := _Clipper.contract.UnpackLog(event, "Redo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClipperRelyIterator is returned from FilterRely and is used to iterate over the raw logs and unpacked data for Rely events raised by the Clipper contract.
type ClipperRelyIterator struct {
	Event *ClipperRely // Event containing the contract specifics and raw log

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
func (it *ClipperRelyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClipperRely)
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
		it.Event = new(ClipperRely)
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
func (it *ClipperRelyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClipperRelyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClipperRely represents a Rely event raised by the Clipper contract.
type ClipperRely struct {
	Usr common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRely is a free log retrieval operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Clipper *ClipperFilterer) FilterRely(opts *bind.FilterOpts, usr []common.Address) (*ClipperRelyIterator, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Clipper.contract.FilterLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return &ClipperRelyIterator{contract: _Clipper.contract, event: "Rely", logs: logs, sub: sub}, nil
}

// WatchRely is a free log subscription operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Clipper *ClipperFilterer) WatchRely(opts *bind.WatchOpts, sink chan<- *ClipperRely, usr []common.Address) (event.Subscription, error) {

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Clipper.contract.WatchLogs(opts, "Rely", usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClipperRely)
				if err := _Clipper.contract.UnpackLog(event, "Rely", log); err != nil {
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

// ParseRely is a log parse operation binding the contract event 0xdd0e34038ac38b2a1ce960229778ac48a8719bc900b6c4f8d0475c6e8b385a60.
//
// Solidity: event Rely(address indexed usr)
func (_Clipper *ClipperFilterer) ParseRely(log types.Log) (*ClipperRely, error) {
	event := new(ClipperRely)
	if err := _Clipper.contract.UnpackLog(event, "Rely", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClipperTakeIterator is returned from FilterTake and is used to iterate over the raw logs and unpacked data for Take events raised by the Clipper contract.
type ClipperTakeIterator struct {
	Event *ClipperTake // Event containing the contract specifics and raw log

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
func (it *ClipperTakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClipperTake)
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
		it.Event = new(ClipperTake)
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
func (it *ClipperTakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClipperTakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClipperTake represents a Take event raised by the Clipper contract.
type ClipperTake struct {
	Id    *big.Int
	Max   *big.Int
	Price *big.Int
	Owe   *big.Int
	Tab   *big.Int
	Lot   *big.Int
	Usr   common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTake is a free log retrieval operation binding the contract event 0x05e309fd6ce72f2ab888a20056bb4210df08daed86f21f95053deb19964d86b1.
//
// Solidity: event Take(uint256 indexed id, uint256 max, uint256 price, uint256 owe, uint256 tab, uint256 lot, address indexed usr)
func (_Clipper *ClipperFilterer) FilterTake(opts *bind.FilterOpts, id []*big.Int, usr []common.Address) (*ClipperTakeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Clipper.contract.FilterLogs(opts, "Take", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return &ClipperTakeIterator{contract: _Clipper.contract, event: "Take", logs: logs, sub: sub}, nil
}

// WatchTake is a free log subscription operation binding the contract event 0x05e309fd6ce72f2ab888a20056bb4210df08daed86f21f95053deb19964d86b1.
//
// Solidity: event Take(uint256 indexed id, uint256 max, uint256 price, uint256 owe, uint256 tab, uint256 lot, address indexed usr)
func (_Clipper *ClipperFilterer) WatchTake(opts *bind.WatchOpts, sink chan<- *ClipperTake, id []*big.Int, usr []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	var usrRule []interface{}
	for _, usrItem := range usr {
		usrRule = append(usrRule, usrItem)
	}

	logs, sub, err := _Clipper.contract.WatchLogs(opts, "Take", idRule, usrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClipperTake)
				if err := _Clipper.contract.UnpackLog(event, "Take", log); err != nil {
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

// ParseTake is a log parse operation binding the contract event 0x05e309fd6ce72f2ab888a20056bb4210df08daed86f21f95053deb19964d86b1.
//
// Solidity: event Take(uint256 indexed id, uint256 max, uint256 price, uint256 owe, uint256 tab, uint256 lot, address indexed usr)
func (_Clipper *ClipperFilterer) ParseTake(log types.Log) (*ClipperTake, error) {
	event := new(ClipperTake)
	if err := _Clipper.contract.UnpackLog(event, "Take", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClipperYankIterator is returned from FilterYank and is used to iterate over the raw logs and unpacked data for Yank events raised by the Clipper contract.
type ClipperYankIterator struct {
	Event *ClipperYank // Event containing the contract specifics and raw log

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
func (it *ClipperYankIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClipperYank)
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
		it.Event = new(ClipperYank)
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
func (it *ClipperYankIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClipperYankIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClipperYank represents a Yank event raised by the Clipper contract.
type ClipperYank struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterYank is a free log retrieval operation binding the contract event 0x2c5d2826eb5903b8fc201cf48094b858f42f61c7eaac9aaf43ebed490138144e.
//
// Solidity: event Yank(uint256 id)
func (_Clipper *ClipperFilterer) FilterYank(opts *bind.FilterOpts) (*ClipperYankIterator, error) {

	logs, sub, err := _Clipper.contract.FilterLogs(opts, "Yank")
	if err != nil {
		return nil, err
	}
	return &ClipperYankIterator{contract: _Clipper.contract, event: "Yank", logs: logs, sub: sub}, nil
}

// WatchYank is a free log subscription operation binding the contract event 0x2c5d2826eb5903b8fc201cf48094b858f42f61c7eaac9aaf43ebed490138144e.
//
// Solidity: event Yank(uint256 id)
func (_Clipper *ClipperFilterer) WatchYank(opts *bind.WatchOpts, sink chan<- *ClipperYank) (event.Subscription, error) {

	logs, sub, err := _Clipper.contract.WatchLogs(opts, "Yank")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClipperYank)
				if err := _Clipper.contract.UnpackLog(event, "Yank", log); err != nil {
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

// ParseYank is a log parse operation binding the contract event 0x2c5d2826eb5903b8fc201cf48094b858f42f61c7eaac9aaf43ebed490138144e.
//
// Solidity: event Yank(uint256 id)
func (_Clipper *ClipperFilterer) ParseYank(log types.Log) (*ClipperYank, error) {
	event := new(ClipperYank)
	if err := _Clipper.contract.UnpackLog(event, "Yank", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
