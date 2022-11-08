// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v1

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
)

// VowMetaData contains all meta data concerning the Vow contract.
var VowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vat_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flapper_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flopper_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"multisig_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Ash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Sin2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tab\",\"type\":\"uint256\"}],\"name\":\"fess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"data\",\"type\":\"address\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flapper\",\"outputs\":[{\"internalType\":\"contractFlapLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"era\",\"type\":\"uint256\"}],\"name\":\"flog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flopper\",\"outputs\":[{\"internalType\":\"contractFlopLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"heal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"kiss\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lever\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multisig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sump\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vat\",\"outputs\":[{\"internalType\":\"contractVatLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wait\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VowABI is the input ABI used to generate the binding from.
// Deprecated: Use VowMetaData.ABI instead.
var VowABI = VowMetaData.ABI

// Vow is an auto generated Go binding around an Ethereum contract.
type Vow struct {
	VowCaller     // Read-only binding to the contract
	VowTransactor // Write-only binding to the contract
	VowFilterer   // Log filterer for contract events
}

// VowCaller is an auto generated read-only Go binding around an Ethereum contract.
type VowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VowSession struct {
	Contract     *Vow              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VowCallerSession struct {
	Contract *VowCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VowTransactorSession struct {
	Contract     *VowTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VowRaw is an auto generated low-level Go binding around an Ethereum contract.
type VowRaw struct {
	Contract *Vow // Generic contract binding to access the raw methods on
}

// VowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VowCallerRaw struct {
	Contract *VowCaller // Generic read-only contract binding to access the raw methods on
}

// VowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VowTransactorRaw struct {
	Contract *VowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVow creates a new instance of Vow, bound to a specific deployed contract.
func NewVow(address common.Address, backend bind.ContractBackend) (*Vow, error) {
	contract, err := bindVow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vow{VowCaller: VowCaller{contract: contract}, VowTransactor: VowTransactor{contract: contract}, VowFilterer: VowFilterer{contract: contract}}, nil
}

// NewVowCaller creates a new read-only instance of Vow, bound to a specific deployed contract.
func NewVowCaller(address common.Address, caller bind.ContractCaller) (*VowCaller, error) {
	contract, err := bindVow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VowCaller{contract: contract}, nil
}

// NewVowTransactor creates a new write-only instance of Vow, bound to a specific deployed contract.
func NewVowTransactor(address common.Address, transactor bind.ContractTransactor) (*VowTransactor, error) {
	contract, err := bindVow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VowTransactor{contract: contract}, nil
}

// NewVowFilterer creates a new log filterer instance of Vow, bound to a specific deployed contract.
func NewVowFilterer(address common.Address, filterer bind.ContractFilterer) (*VowFilterer, error) {
	contract, err := bindVow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VowFilterer{contract: contract}, nil
}

// bindVow binds a generic wrapper to an already deployed contract.
func bindVow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VowABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vow *VowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vow.Contract.VowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vow *VowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.Contract.VowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vow *VowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vow.Contract.VowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vow *VowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vow *VowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vow *VowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vow.Contract.contract.Transact(opts, method, params...)
}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_Vow *VowCaller) Ash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "Ash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_Vow *VowSession) Ash() (*big.Int, error) {
	return _Vow.Contract.Ash(&_Vow.CallOpts)
}

// Ash is a free data retrieval call binding the contract method 0x2a1d2b3c.
//
// Solidity: function Ash() view returns(uint256)
func (_Vow *VowCallerSession) Ash() (*big.Int, error) {
	return _Vow.Contract.Ash(&_Vow.CallOpts)
}

// Sin2 is a free data retrieval call binding the contract method 0xa7206ec0.
//
// Solidity: function Sin2() view returns(uint256)
func (_Vow *VowCaller) Sin2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "Sin2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin2 is a free data retrieval call binding the contract method 0xa7206ec0.
//
// Solidity: function Sin2() view returns(uint256)
func (_Vow *VowSession) Sin2() (*big.Int, error) {
	return _Vow.Contract.Sin2(&_Vow.CallOpts)
}

// Sin2 is a free data retrieval call binding the contract method 0xa7206ec0.
//
// Solidity: function Sin2() view returns(uint256)
func (_Vow *VowCallerSession) Sin2() (*big.Int, error) {
	return _Vow.Contract.Sin2(&_Vow.CallOpts)
}

// Bump is a free data retrieval call binding the contract method 0x68110b2f.
//
// Solidity: function bump() view returns(uint256)
func (_Vow *VowCaller) Bump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "bump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bump is a free data retrieval call binding the contract method 0x68110b2f.
//
// Solidity: function bump() view returns(uint256)
func (_Vow *VowSession) Bump() (*big.Int, error) {
	return _Vow.Contract.Bump(&_Vow.CallOpts)
}

// Bump is a free data retrieval call binding the contract method 0x68110b2f.
//
// Solidity: function bump() view returns(uint256)
func (_Vow *VowCallerSession) Bump() (*big.Int, error) {
	return _Vow.Contract.Bump(&_Vow.CallOpts)
}

// Dump is a free data retrieval call binding the contract method 0xe4330545.
//
// Solidity: function dump() view returns(uint256)
func (_Vow *VowCaller) Dump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "dump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dump is a free data retrieval call binding the contract method 0xe4330545.
//
// Solidity: function dump() view returns(uint256)
func (_Vow *VowSession) Dump() (*big.Int, error) {
	return _Vow.Contract.Dump(&_Vow.CallOpts)
}

// Dump is a free data retrieval call binding the contract method 0xe4330545.
//
// Solidity: function dump() view returns(uint256)
func (_Vow *VowCallerSession) Dump() (*big.Int, error) {
	return _Vow.Contract.Dump(&_Vow.CallOpts)
}

// Flapper is a free data retrieval call binding the contract method 0x5ca0d723.
//
// Solidity: function flapper() view returns(address)
func (_Vow *VowCaller) Flapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "flapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flapper is a free data retrieval call binding the contract method 0x5ca0d723.
//
// Solidity: function flapper() view returns(address)
func (_Vow *VowSession) Flapper() (common.Address, error) {
	return _Vow.Contract.Flapper(&_Vow.CallOpts)
}

// Flapper is a free data retrieval call binding the contract method 0x5ca0d723.
//
// Solidity: function flapper() view returns(address)
func (_Vow *VowCallerSession) Flapper() (common.Address, error) {
	return _Vow.Contract.Flapper(&_Vow.CallOpts)
}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_Vow *VowCaller) Flopper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "flopper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_Vow *VowSession) Flopper() (common.Address, error) {
	return _Vow.Contract.Flopper(&_Vow.CallOpts)
}

// Flopper is a free data retrieval call binding the contract method 0x4081d73a.
//
// Solidity: function flopper() view returns(address)
func (_Vow *VowCallerSession) Flopper() (common.Address, error) {
	return _Vow.Contract.Flopper(&_Vow.CallOpts)
}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_Vow *VowCaller) Hump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "hump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_Vow *VowSession) Hump() (*big.Int, error) {
	return _Vow.Contract.Hump(&_Vow.CallOpts)
}

// Hump is a free data retrieval call binding the contract method 0x1b8e8cfa.
//
// Solidity: function hump() view returns(uint256)
func (_Vow *VowCallerSession) Hump() (*big.Int, error) {
	return _Vow.Contract.Hump(&_Vow.CallOpts)
}

// Lever is a free data retrieval call binding the contract method 0x36cd6052.
//
// Solidity: function lever() view returns(uint256)
func (_Vow *VowCaller) Lever(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "lever")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Lever is a free data retrieval call binding the contract method 0x36cd6052.
//
// Solidity: function lever() view returns(uint256)
func (_Vow *VowSession) Lever() (*big.Int, error) {
	return _Vow.Contract.Lever(&_Vow.CallOpts)
}

// Lever is a free data retrieval call binding the contract method 0x36cd6052.
//
// Solidity: function lever() view returns(uint256)
func (_Vow *VowCallerSession) Lever() (*big.Int, error) {
	return _Vow.Contract.Lever(&_Vow.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vow *VowCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vow *VowSession) Live() (*big.Int, error) {
	return _Vow.Contract.Live(&_Vow.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vow *VowCallerSession) Live() (*big.Int, error) {
	return _Vow.Contract.Live(&_Vow.CallOpts)
}

// Multisig is a free data retrieval call binding the contract method 0x4783c35b.
//
// Solidity: function multisig() view returns(address)
func (_Vow *VowCaller) Multisig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "multisig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Multisig is a free data retrieval call binding the contract method 0x4783c35b.
//
// Solidity: function multisig() view returns(address)
func (_Vow *VowSession) Multisig() (common.Address, error) {
	return _Vow.Contract.Multisig(&_Vow.CallOpts)
}

// Multisig is a free data retrieval call binding the contract method 0x4783c35b.
//
// Solidity: function multisig() view returns(address)
func (_Vow *VowCallerSession) Multisig() (common.Address, error) {
	return _Vow.Contract.Multisig(&_Vow.CallOpts)
}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_Vow *VowCaller) Sin(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "sin", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_Vow *VowSession) Sin(arg0 *big.Int) (*big.Int, error) {
	return _Vow.Contract.Sin(&_Vow.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xcb5cc109.
//
// Solidity: function sin(uint256 ) view returns(uint256)
func (_Vow *VowCallerSession) Sin(arg0 *big.Int) (*big.Int, error) {
	return _Vow.Contract.Sin(&_Vow.CallOpts, arg0)
}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_Vow *VowCaller) Sump(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "sump")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_Vow *VowSession) Sump() (*big.Int, error) {
	return _Vow.Contract.Sump(&_Vow.CallOpts)
}

// Sump is a free data retrieval call binding the contract method 0xc349d362.
//
// Solidity: function sump() view returns(uint256)
func (_Vow *VowCallerSession) Sump() (*big.Int, error) {
	return _Vow.Contract.Sump(&_Vow.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Vow *VowCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Vow *VowSession) Vat() (common.Address, error) {
	return _Vow.Contract.Vat(&_Vow.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Vow *VowCallerSession) Vat() (common.Address, error) {
	return _Vow.Contract.Vat(&_Vow.CallOpts)
}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_Vow *VowCaller) Wait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "wait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_Vow *VowSession) Wait() (*big.Int, error) {
	return _Vow.Contract.Wait(&_Vow.CallOpts)
}

// Wait is a free data retrieval call binding the contract method 0x64bd7013.
//
// Solidity: function wait() view returns(uint256)
func (_Vow *VowCallerSession) Wait() (*big.Int, error) {
	return _Vow.Contract.Wait(&_Vow.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vow *VowCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vow.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vow *VowSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vow.Contract.Wards(&_Vow.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vow *VowCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vow.Contract.Wards(&_Vow.CallOpts, arg0)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vow *VowTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vow *VowSession) Cage() (*types.Transaction, error) {
	return _Vow.Contract.Cage(&_Vow.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vow *VowTransactorSession) Cage() (*types.Transaction, error) {
	return _Vow.Contract.Cage(&_Vow.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vow *VowTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vow *VowSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vow.Contract.Deny(&_Vow.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vow *VowTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vow.Contract.Deny(&_Vow.TransactOpts, usr)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_Vow *VowTransactor) Fess(opts *bind.TransactOpts, tab *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "fess", tab)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_Vow *VowSession) Fess(tab *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Fess(&_Vow.TransactOpts, tab)
}

// Fess is a paid mutator transaction binding the contract method 0x697efb78.
//
// Solidity: function fess(uint256 tab) returns()
func (_Vow *VowTransactorSession) Fess(tab *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Fess(&_Vow.TransactOpts, tab)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vow *VowTransactor) File(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "file", what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vow *VowSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.File(&_Vow.TransactOpts, what, data)
}

// File is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vow *VowTransactorSession) File(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.File(&_Vow.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Vow *VowTransactor) File0(opts *bind.TransactOpts, what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Vow *VowSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Vow.Contract.File0(&_Vow.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0xd4e8be83.
//
// Solidity: function file(bytes32 what, address data) returns()
func (_Vow *VowTransactorSession) File0(what [32]byte, data common.Address) (*types.Transaction, error) {
	return _Vow.Contract.File0(&_Vow.TransactOpts, what, data)
}

// Flap is a paid mutator transaction binding the contract method 0x0e01198b.
//
// Solidity: function flap() returns(uint256 id)
func (_Vow *VowTransactor) Flap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "flap")
}

// Flap is a paid mutator transaction binding the contract method 0x0e01198b.
//
// Solidity: function flap() returns(uint256 id)
func (_Vow *VowSession) Flap() (*types.Transaction, error) {
	return _Vow.Contract.Flap(&_Vow.TransactOpts)
}

// Flap is a paid mutator transaction binding the contract method 0x0e01198b.
//
// Solidity: function flap() returns(uint256 id)
func (_Vow *VowTransactorSession) Flap() (*types.Transaction, error) {
	return _Vow.Contract.Flap(&_Vow.TransactOpts)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_Vow *VowTransactor) Flog(opts *bind.TransactOpts, era *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "flog", era)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_Vow *VowSession) Flog(era *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Flog(&_Vow.TransactOpts, era)
}

// Flog is a paid mutator transaction binding the contract method 0xd7ee674b.
//
// Solidity: function flog(uint256 era) returns()
func (_Vow *VowTransactorSession) Flog(era *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Flog(&_Vow.TransactOpts, era)
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns(uint256 id)
func (_Vow *VowTransactor) Flop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "flop")
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns(uint256 id)
func (_Vow *VowSession) Flop() (*types.Transaction, error) {
	return _Vow.Contract.Flop(&_Vow.TransactOpts)
}

// Flop is a paid mutator transaction binding the contract method 0xbbbb0d7b.
//
// Solidity: function flop() returns(uint256 id)
func (_Vow *VowTransactorSession) Flop() (*types.Transaction, error) {
	return _Vow.Contract.Flop(&_Vow.TransactOpts)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vow *VowTransactor) Heal(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "heal", rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vow *VowSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Heal(&_Vow.TransactOpts, rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vow *VowTransactorSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Heal(&_Vow.TransactOpts, rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_Vow *VowTransactor) Kiss(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "kiss", rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_Vow *VowSession) Kiss(rad *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Kiss(&_Vow.TransactOpts, rad)
}

// Kiss is a paid mutator transaction binding the contract method 0x2506855a.
//
// Solidity: function kiss(uint256 rad) returns()
func (_Vow *VowTransactorSession) Kiss(rad *big.Int) (*types.Transaction, error) {
	return _Vow.Contract.Kiss(&_Vow.TransactOpts, rad)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vow *VowTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vow.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vow *VowSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vow.Contract.Rely(&_Vow.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vow *VowTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vow.Contract.Rely(&_Vow.TransactOpts, usr)
}
