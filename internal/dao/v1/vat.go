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

// VatMetaData contains all meta data concerning the Vat contract.
var VatMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Line\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"behalf\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"can\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"debt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"what\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"name\":\"file\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"flux\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"rate\",\"type\":\"int256\"}],\"name\":\"fold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"fork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"frob\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"i\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"w\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"dink\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"dart\",\"type\":\"int256\"}],\"name\":\"grab\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"heal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"hope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ilks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"Art\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"line\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dust\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"live\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"nope\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"regard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sikka\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"wad\",\"type\":\"int256\"}],\"name\":\"slip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"u\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rad\",\"type\":\"uint256\"}],\"name\":\"suck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"urns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ink\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"art\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// VatABI is the input ABI used to generate the binding from.
// Deprecated: Use VatMetaData.ABI instead.
var VatABI = VatMetaData.ABI

// Vat is an auto generated Go binding around an Ethereum contract.
type Vat struct {
	VatCaller     // Read-only binding to the contract
	VatTransactor // Write-only binding to the contract
	VatFilterer   // Log filterer for contract events
}

// VatCaller is an auto generated read-only Go binding around an Ethereum contract.
type VatCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VatTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VatFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VatSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VatSession struct {
	Contract     *Vat              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VatCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VatCallerSession struct {
	Contract *VatCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VatTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VatTransactorSession struct {
	Contract     *VatTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VatRaw is an auto generated low-level Go binding around an Ethereum contract.
type VatRaw struct {
	Contract *Vat // Generic contract binding to access the raw methods on
}

// VatCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VatCallerRaw struct {
	Contract *VatCaller // Generic read-only contract binding to access the raw methods on
}

// VatTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VatTransactorRaw struct {
	Contract *VatTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVat creates a new instance of Vat, bound to a specific deployed contract.
func NewVat(address common.Address, backend bind.ContractBackend) (*Vat, error) {
	contract, err := bindVat(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vat{VatCaller: VatCaller{contract: contract}, VatTransactor: VatTransactor{contract: contract}, VatFilterer: VatFilterer{contract: contract}}, nil
}

// NewVatCaller creates a new read-only instance of Vat, bound to a specific deployed contract.
func NewVatCaller(address common.Address, caller bind.ContractCaller) (*VatCaller, error) {
	contract, err := bindVat(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VatCaller{contract: contract}, nil
}

// NewVatTransactor creates a new write-only instance of Vat, bound to a specific deployed contract.
func NewVatTransactor(address common.Address, transactor bind.ContractTransactor) (*VatTransactor, error) {
	contract, err := bindVat(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VatTransactor{contract: contract}, nil
}

// NewVatFilterer creates a new log filterer instance of Vat, bound to a specific deployed contract.
func NewVatFilterer(address common.Address, filterer bind.ContractFilterer) (*VatFilterer, error) {
	contract, err := bindVat(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VatFilterer{contract: contract}, nil
}

// bindVat binds a generic wrapper to an already deployed contract.
func bindVat(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VatABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vat *VatRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vat.Contract.VatCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vat *VatRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.Contract.VatTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vat *VatRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vat.Contract.VatTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vat *VatCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vat.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vat *VatTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vat *VatTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vat.Contract.contract.Transact(opts, method, params...)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vat *VatCaller) Line(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "Line")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vat *VatSession) Line() (*big.Int, error) {
	return _Vat.Contract.Line(&_Vat.CallOpts)
}

// Line is a free data retrieval call binding the contract method 0xbabe8a3f.
//
// Solidity: function Line() view returns(uint256)
func (_Vat *VatCallerSession) Line() (*big.Int, error) {
	return _Vat.Contract.Line(&_Vat.CallOpts)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vat *VatCaller) Can(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "can", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vat *VatSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Can(&_Vat.CallOpts, arg0, arg1)
}

// Can is a free data retrieval call binding the contract method 0x4538c4eb.
//
// Solidity: function can(address , address ) view returns(uint256)
func (_Vat *VatCallerSession) Can(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Can(&_Vat.CallOpts, arg0, arg1)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vat *VatCaller) Debt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "debt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vat *VatSession) Debt() (*big.Int, error) {
	return _Vat.Contract.Debt(&_Vat.CallOpts)
}

// Debt is a free data retrieval call binding the contract method 0x0dca59c1.
//
// Solidity: function debt() view returns(uint256)
func (_Vat *VatCallerSession) Debt() (*big.Int, error) {
	return _Vat.Contract.Debt(&_Vat.CallOpts)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vat *VatCaller) Gem(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "gem", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vat *VatSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Gem(&_Vat.CallOpts, arg0, arg1)
}

// Gem is a free data retrieval call binding the contract method 0x214414d5.
//
// Solidity: function gem(bytes32 , address ) view returns(uint256)
func (_Vat *VatCallerSession) Gem(arg0 [32]byte, arg1 common.Address) (*big.Int, error) {
	return _Vat.Contract.Gem(&_Vat.CallOpts, arg0, arg1)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatCaller) Ilks(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "ilks", arg0)

	outstruct := new(struct {
		Art  *big.Int
		Rate *big.Int
		Spot *big.Int
		Line *big.Int
		Dust *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Art = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Rate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Spot = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Line = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Dust = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _Vat.Contract.Ilks(&_Vat.CallOpts, arg0)
}

// Ilks is a free data retrieval call binding the contract method 0xd9638d36.
//
// Solidity: function ilks(bytes32 ) view returns(uint256 Art, uint256 rate, uint256 spot, uint256 line, uint256 dust)
func (_Vat *VatCallerSession) Ilks(arg0 [32]byte) (struct {
	Art  *big.Int
	Rate *big.Int
	Spot *big.Int
	Line *big.Int
	Dust *big.Int
}, error) {
	return _Vat.Contract.Ilks(&_Vat.CallOpts, arg0)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vat *VatCaller) Live(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "live")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vat *VatSession) Live() (*big.Int, error) {
	return _Vat.Contract.Live(&_Vat.CallOpts)
}

// Live is a free data retrieval call binding the contract method 0x957aa58c.
//
// Solidity: function live() view returns(uint256)
func (_Vat *VatCallerSession) Live() (*big.Int, error) {
	return _Vat.Contract.Live(&_Vat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vat *VatCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vat *VatSession) Owner() (common.Address, error) {
	return _Vat.Contract.Owner(&_Vat.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vat *VatCallerSession) Owner() (common.Address, error) {
	return _Vat.Contract.Owner(&_Vat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vat *VatCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vat *VatSession) ProxiableUUID() ([32]byte, error) {
	return _Vat.Contract.ProxiableUUID(&_Vat.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Vat *VatCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Vat.Contract.ProxiableUUID(&_Vat.CallOpts)
}

// Sikka is a free data retrieval call binding the contract method 0xc34501fb.
//
// Solidity: function sikka(address ) view returns(uint256)
func (_Vat *VatCaller) Sikka(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "sikka", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sikka is a free data retrieval call binding the contract method 0xc34501fb.
//
// Solidity: function sikka(address ) view returns(uint256)
func (_Vat *VatSession) Sikka(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sikka(&_Vat.CallOpts, arg0)
}

// Sikka is a free data retrieval call binding the contract method 0xc34501fb.
//
// Solidity: function sikka(address ) view returns(uint256)
func (_Vat *VatCallerSession) Sikka(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sikka(&_Vat.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vat *VatCaller) Sin(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "sin", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vat *VatSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sin(&_Vat.CallOpts, arg0)
}

// Sin is a free data retrieval call binding the contract method 0xf059212a.
//
// Solidity: function sin(address ) view returns(uint256)
func (_Vat *VatCallerSession) Sin(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Sin(&_Vat.CallOpts, arg0)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vat *VatCaller) Urns(opts *bind.CallOpts, arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "urns", arg0, arg1)

	outstruct := new(struct {
		Ink *big.Int
		Art *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ink = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Art = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vat *VatSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _Vat.Contract.Urns(&_Vat.CallOpts, arg0, arg1)
}

// Urns is a free data retrieval call binding the contract method 0x2424be5c.
//
// Solidity: function urns(bytes32 , address ) view returns(uint256 ink, uint256 art)
func (_Vat *VatCallerSession) Urns(arg0 [32]byte, arg1 common.Address) (struct {
	Ink *big.Int
	Art *big.Int
}, error) {
	return _Vat.Contract.Urns(&_Vat.CallOpts, arg0, arg1)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vat *VatCaller) Vice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "vice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vat *VatSession) Vice() (*big.Int, error) {
	return _Vat.Contract.Vice(&_Vat.CallOpts)
}

// Vice is a free data retrieval call binding the contract method 0x2d61a355.
//
// Solidity: function vice() view returns(uint256)
func (_Vat *VatCallerSession) Vice() (*big.Int, error) {
	return _Vat.Contract.Vice(&_Vat.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vat *VatCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vat.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vat *VatSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Wards(&_Vat.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Vat *VatCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Vat.Contract.Wards(&_Vat.CallOpts, arg0)
}

// Behalf is a paid mutator transaction binding the contract method 0x5216788d.
//
// Solidity: function behalf(address bit, address usr) returns()
func (_Vat *VatTransactor) Behalf(opts *bind.TransactOpts, bit common.Address, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "behalf", bit, usr)
}

// Behalf is a paid mutator transaction binding the contract method 0x5216788d.
//
// Solidity: function behalf(address bit, address usr) returns()
func (_Vat *VatSession) Behalf(bit common.Address, usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Behalf(&_Vat.TransactOpts, bit, usr)
}

// Behalf is a paid mutator transaction binding the contract method 0x5216788d.
//
// Solidity: function behalf(address bit, address usr) returns()
func (_Vat *VatTransactorSession) Behalf(bit common.Address, usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Behalf(&_Vat.TransactOpts, bit, usr)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatTransactor) Cage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "cage")
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatSession) Cage() (*types.Transaction, error) {
	return _Vat.Contract.Cage(&_Vat.TransactOpts)
}

// Cage is a paid mutator transaction binding the contract method 0x69245009.
//
// Solidity: function cage() returns()
func (_Vat *VatTransactorSession) Cage() (*types.Transaction, error) {
	return _Vat.Contract.Cage(&_Vat.TransactOpts)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Deny(&_Vat.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Vat *VatTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Deny(&_Vat.TransactOpts, usr)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatTransactor) File(opts *bind.TransactOpts, ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "file", ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File(&_Vat.TransactOpts, ilk, what, data)
}

// File is a paid mutator transaction binding the contract method 0x1a0b287e.
//
// Solidity: function file(bytes32 ilk, bytes32 what, uint256 data) returns()
func (_Vat *VatTransactorSession) File(ilk [32]byte, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File(&_Vat.TransactOpts, ilk, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatTransactor) File0(opts *bind.TransactOpts, what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "file0", what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File0(&_Vat.TransactOpts, what, data)
}

// File0 is a paid mutator transaction binding the contract method 0x29ae8114.
//
// Solidity: function file(bytes32 what, uint256 data) returns()
func (_Vat *VatTransactorSession) File0(what [32]byte, data *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.File0(&_Vat.TransactOpts, what, data)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatTransactor) Flux(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "flux", ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Flux(&_Vat.TransactOpts, ilk, src, dst, wad)
}

// Flux is a paid mutator transaction binding the contract method 0x6111be2e.
//
// Solidity: function flux(bytes32 ilk, address src, address dst, uint256 wad) returns()
func (_Vat *VatTransactorSession) Flux(ilk [32]byte, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Flux(&_Vat.TransactOpts, ilk, src, dst, wad)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatTransactor) Fold(opts *bind.TransactOpts, i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "fold", i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fold(&_Vat.TransactOpts, i, u, rate)
}

// Fold is a paid mutator transaction binding the contract method 0xb65337df.
//
// Solidity: function fold(bytes32 i, address u, int256 rate) returns()
func (_Vat *VatTransactorSession) Fold(i [32]byte, u common.Address, rate *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fold(&_Vat.TransactOpts, i, u, rate)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Fork(opts *bind.TransactOpts, ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "fork", ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fork(&_Vat.TransactOpts, ilk, src, dst, dink, dart)
}

// Fork is a paid mutator transaction binding the contract method 0x870c616d.
//
// Solidity: function fork(bytes32 ilk, address src, address dst, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Fork(ilk [32]byte, src common.Address, dst common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Fork(&_Vat.TransactOpts, ilk, src, dst, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Frob(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "frob", i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Frob(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Frob is a paid mutator transaction binding the contract method 0x76088703.
//
// Solidity: function frob(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Frob(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Frob(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactor) Grab(opts *bind.TransactOpts, i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "grab", i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Grab(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Grab is a paid mutator transaction binding the contract method 0x7bab3f40.
//
// Solidity: function grab(bytes32 i, address u, address v, address w, int256 dink, int256 dart) returns()
func (_Vat *VatTransactorSession) Grab(i [32]byte, u common.Address, v common.Address, w common.Address, dink *big.Int, dart *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Grab(&_Vat.TransactOpts, i, u, v, w, dink, dart)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatTransactor) Heal(opts *bind.TransactOpts, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "heal", rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Heal(&_Vat.TransactOpts, rad)
}

// Heal is a paid mutator transaction binding the contract method 0xf37ac61c.
//
// Solidity: function heal(uint256 rad) returns()
func (_Vat *VatTransactorSession) Heal(rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Heal(&_Vat.TransactOpts, rad)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatTransactor) Hope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "hope", usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Hope(&_Vat.TransactOpts, usr)
}

// Hope is a paid mutator transaction binding the contract method 0xa3b22fc4.
//
// Solidity: function hope(address usr) returns()
func (_Vat *VatTransactorSession) Hope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Hope(&_Vat.TransactOpts, usr)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatTransactor) Init(opts *bind.TransactOpts, ilk [32]byte) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "init", ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Vat.Contract.Init(&_Vat.TransactOpts, ilk)
}

// Init is a paid mutator transaction binding the contract method 0x3b663195.
//
// Solidity: function init(bytes32 ilk) returns()
func (_Vat *VatTransactorSession) Init(ilk [32]byte) (*types.Transaction, error) {
	return _Vat.Contract.Init(&_Vat.TransactOpts, ilk)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vat *VatTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vat *VatSession) Initialize() (*types.Transaction, error) {
	return _Vat.Contract.Initialize(&_Vat.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vat *VatTransactorSession) Initialize() (*types.Transaction, error) {
	return _Vat.Contract.Initialize(&_Vat.TransactOpts)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatTransactor) Move(opts *bind.TransactOpts, src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "move", src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Move(&_Vat.TransactOpts, src, dst, rad)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address src, address dst, uint256 rad) returns()
func (_Vat *VatTransactorSession) Move(src common.Address, dst common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Move(&_Vat.TransactOpts, src, dst, rad)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatTransactor) Nope(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "nope", usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Nope(&_Vat.TransactOpts, usr)
}

// Nope is a paid mutator transaction binding the contract method 0xdc4d20fa.
//
// Solidity: function nope(address usr) returns()
func (_Vat *VatTransactorSession) Nope(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Nope(&_Vat.TransactOpts, usr)
}

// Regard is a paid mutator transaction binding the contract method 0xd7a9b792.
//
// Solidity: function regard(address bit, address usr) returns()
func (_Vat *VatTransactor) Regard(opts *bind.TransactOpts, bit common.Address, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "regard", bit, usr)
}

// Regard is a paid mutator transaction binding the contract method 0xd7a9b792.
//
// Solidity: function regard(address bit, address usr) returns()
func (_Vat *VatSession) Regard(bit common.Address, usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Regard(&_Vat.TransactOpts, bit, usr)
}

// Regard is a paid mutator transaction binding the contract method 0xd7a9b792.
//
// Solidity: function regard(address bit, address usr) returns()
func (_Vat *VatTransactorSession) Regard(bit common.Address, usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Regard(&_Vat.TransactOpts, bit, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Rely(&_Vat.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Vat *VatTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Vat.Contract.Rely(&_Vat.TransactOpts, usr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vat *VatTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vat *VatSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vat.Contract.RenounceOwnership(&_Vat.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vat *VatTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vat.Contract.RenounceOwnership(&_Vat.TransactOpts)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatTransactor) Slip(opts *bind.TransactOpts, ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "slip", ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Slip(&_Vat.TransactOpts, ilk, usr, wad)
}

// Slip is a paid mutator transaction binding the contract method 0x7cdd3fde.
//
// Solidity: function slip(bytes32 ilk, address usr, int256 wad) returns()
func (_Vat *VatTransactorSession) Slip(ilk [32]byte, usr common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Slip(&_Vat.TransactOpts, ilk, usr, wad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatTransactor) Suck(opts *bind.TransactOpts, u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "suck", u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Suck(&_Vat.TransactOpts, u, v, rad)
}

// Suck is a paid mutator transaction binding the contract method 0xf24e23eb.
//
// Solidity: function suck(address u, address v, uint256 rad) returns()
func (_Vat *VatTransactorSession) Suck(u common.Address, v common.Address, rad *big.Int) (*types.Transaction, error) {
	return _Vat.Contract.Suck(&_Vat.TransactOpts, u, v, rad)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vat *VatTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vat *VatSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vat.Contract.TransferOwnership(&_Vat.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vat *VatTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vat.Contract.TransferOwnership(&_Vat.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vat *VatTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vat *VatSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vat.Contract.UpgradeTo(&_Vat.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Vat *VatTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Vat.Contract.UpgradeTo(&_Vat.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vat *VatTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vat.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vat *VatSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vat.Contract.UpgradeToAndCall(&_Vat.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Vat *VatTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Vat.Contract.UpgradeToAndCall(&_Vat.TransactOpts, newImplementation, data)
}

// VatAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Vat contract.
type VatAdminChangedIterator struct {
	Event *VatAdminChanged // Event containing the contract specifics and raw log

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
func (it *VatAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatAdminChanged)
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
		it.Event = new(VatAdminChanged)
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
func (it *VatAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatAdminChanged represents a AdminChanged event raised by the Vat contract.
type VatAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vat *VatFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*VatAdminChangedIterator, error) {

	logs, sub, err := _Vat.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &VatAdminChangedIterator{contract: _Vat.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vat *VatFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *VatAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Vat.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatAdminChanged)
				if err := _Vat.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Vat *VatFilterer) ParseAdminChanged(log types.Log) (*VatAdminChanged, error) {
	event := new(VatAdminChanged)
	if err := _Vat.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Vat contract.
type VatBeaconUpgradedIterator struct {
	Event *VatBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *VatBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatBeaconUpgraded)
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
		it.Event = new(VatBeaconUpgraded)
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
func (it *VatBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatBeaconUpgraded represents a BeaconUpgraded event raised by the Vat contract.
type VatBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vat *VatFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*VatBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &VatBeaconUpgradedIterator{contract: _Vat.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vat *VatFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *VatBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatBeaconUpgraded)
				if err := _Vat.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Vat *VatFilterer) ParseBeaconUpgraded(log types.Log) (*VatBeaconUpgraded, error) {
	event := new(VatBeaconUpgraded)
	if err := _Vat.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vat contract.
type VatOwnershipTransferredIterator struct {
	Event *VatOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VatOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatOwnershipTransferred)
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
		it.Event = new(VatOwnershipTransferred)
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
func (it *VatOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatOwnershipTransferred represents a OwnershipTransferred event raised by the Vat contract.
type VatOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vat *VatFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VatOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VatOwnershipTransferredIterator{contract: _Vat.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vat *VatFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VatOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatOwnershipTransferred)
				if err := _Vat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Vat *VatFilterer) ParseOwnershipTransferred(log types.Log) (*VatOwnershipTransferred, error) {
	event := new(VatOwnershipTransferred)
	if err := _Vat.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VatUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Vat contract.
type VatUpgradedIterator struct {
	Event *VatUpgraded // Event containing the contract specifics and raw log

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
func (it *VatUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VatUpgraded)
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
		it.Event = new(VatUpgraded)
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
func (it *VatUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VatUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VatUpgraded represents a Upgraded event raised by the Vat contract.
type VatUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vat *VatFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*VatUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vat.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &VatUpgradedIterator{contract: _Vat.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vat *VatFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *VatUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Vat.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VatUpgraded)
				if err := _Vat.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Vat *VatFilterer) ParseUpgraded(log types.Log) (*VatUpgraded, error) {
	event := new(VatUpgraded)
	if err := _Vat.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
