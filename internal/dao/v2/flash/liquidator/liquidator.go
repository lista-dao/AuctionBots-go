// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package liquidator

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

// LiquidatorMetaData contains all meta data concerning the Liquidator contract.
var LiquidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"added\",\"type\":\"bool\"}],\"name\":\"PairWhitelistChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"}],\"name\":\"SellToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"added\",\"type\":\"bool\"}],\"name\":\"TokenWhitelistChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAm\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralAm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralReal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapData\",\"type\":\"bytes\"}],\"name\":\"flashLiquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAm\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralAm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralReal\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapData\",\"type\":\"bytes\"}],\"name\":\"flashLiquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"},{\"internalType\":\"contractIERC3156FlashLender\",\"name\":\"_lender\",\"type\":\"address\"},{\"internalType\":\"contractIInteraction\",\"name\":\"_interaction\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lisUSD\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interaction\",\"outputs\":[{\"internalType\":\"contractIInteraction\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lender\",\"outputs\":[{\"internalType\":\"contractIERC3156FlashLender\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"auctionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralAm\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lisUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onFlashLoan\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pairWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"swapData\",\"type\":\"bytes\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"swapData\",\"type\":\"bytes\"}],\"name\":\"sellToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setPairWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setTokenWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LiquidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidatorMetaData.ABI instead.
var LiquidatorABI = LiquidatorMetaData.ABI

// Liquidator is an auto generated Go binding around an Ethereum contract.
type Liquidator struct {
	LiquidatorCaller     // Read-only binding to the contract
	LiquidatorTransactor // Write-only binding to the contract
	LiquidatorFilterer   // Log filterer for contract events
}

// LiquidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidatorSession struct {
	Contract     *Liquidator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidatorCallerSession struct {
	Contract *LiquidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LiquidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidatorTransactorSession struct {
	Contract     *LiquidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LiquidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidatorRaw struct {
	Contract *Liquidator // Generic contract binding to access the raw methods on
}

// LiquidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidatorCallerRaw struct {
	Contract *LiquidatorCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidatorTransactorRaw struct {
	Contract *LiquidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidator creates a new instance of Liquidator, bound to a specific deployed contract.
func NewLiquidator(address common.Address, backend bind.ContractBackend) (*Liquidator, error) {
	contract, err := bindLiquidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Liquidator{LiquidatorCaller: LiquidatorCaller{contract: contract}, LiquidatorTransactor: LiquidatorTransactor{contract: contract}, LiquidatorFilterer: LiquidatorFilterer{contract: contract}}, nil
}

// NewLiquidatorCaller creates a new read-only instance of Liquidator, bound to a specific deployed contract.
func NewLiquidatorCaller(address common.Address, caller bind.ContractCaller) (*LiquidatorCaller, error) {
	contract, err := bindLiquidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidatorCaller{contract: contract}, nil
}

// NewLiquidatorTransactor creates a new write-only instance of Liquidator, bound to a specific deployed contract.
func NewLiquidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidatorTransactor, error) {
	contract, err := bindLiquidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidatorTransactor{contract: contract}, nil
}

// NewLiquidatorFilterer creates a new log filterer instance of Liquidator, bound to a specific deployed contract.
func NewLiquidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidatorFilterer, error) {
	contract, err := bindLiquidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidatorFilterer{contract: contract}, nil
}

// bindLiquidator binds a generic wrapper to an already deployed contract.
func bindLiquidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidator *LiquidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidator.Contract.LiquidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidator *LiquidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidator.Contract.LiquidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidator *LiquidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidator.Contract.LiquidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Liquidator *LiquidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Liquidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Liquidator *LiquidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Liquidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Liquidator *LiquidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Liquidator.Contract.contract.Transact(opts, method, params...)
}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Liquidator *LiquidatorCaller) BOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "BOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Liquidator *LiquidatorSession) BOT() ([32]byte, error) {
	return _Liquidator.Contract.BOT(&_Liquidator.CallOpts)
}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Liquidator *LiquidatorCallerSession) BOT() ([32]byte, error) {
	return _Liquidator.Contract.BOT(&_Liquidator.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Liquidator *LiquidatorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Liquidator *LiquidatorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Liquidator.Contract.DEFAULTADMINROLE(&_Liquidator.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Liquidator *LiquidatorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Liquidator.Contract.DEFAULTADMINROLE(&_Liquidator.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Liquidator *LiquidatorCaller) MANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Liquidator *LiquidatorSession) MANAGER() ([32]byte, error) {
	return _Liquidator.Contract.MANAGER(&_Liquidator.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Liquidator *LiquidatorCallerSession) MANAGER() ([32]byte, error) {
	return _Liquidator.Contract.MANAGER(&_Liquidator.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Liquidator *LiquidatorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Liquidator *LiquidatorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Liquidator.Contract.GetRoleAdmin(&_Liquidator.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Liquidator *LiquidatorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Liquidator.Contract.GetRoleAdmin(&_Liquidator.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Liquidator *LiquidatorCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Liquidator *LiquidatorSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Liquidator.Contract.GetRoleMember(&_Liquidator.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Liquidator *LiquidatorCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Liquidator.Contract.GetRoleMember(&_Liquidator.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Liquidator *LiquidatorCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Liquidator *LiquidatorSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Liquidator.Contract.GetRoleMemberCount(&_Liquidator.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Liquidator *LiquidatorCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Liquidator.Contract.GetRoleMemberCount(&_Liquidator.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Liquidator *LiquidatorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Liquidator *LiquidatorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Liquidator.Contract.HasRole(&_Liquidator.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Liquidator *LiquidatorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Liquidator.Contract.HasRole(&_Liquidator.CallOpts, role, account)
}

// Interaction is a free data retrieval call binding the contract method 0x6c697331.
//
// Solidity: function interaction() view returns(address)
func (_Liquidator *LiquidatorCaller) Interaction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "interaction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Interaction is a free data retrieval call binding the contract method 0x6c697331.
//
// Solidity: function interaction() view returns(address)
func (_Liquidator *LiquidatorSession) Interaction() (common.Address, error) {
	return _Liquidator.Contract.Interaction(&_Liquidator.CallOpts)
}

// Interaction is a free data retrieval call binding the contract method 0x6c697331.
//
// Solidity: function interaction() view returns(address)
func (_Liquidator *LiquidatorCallerSession) Interaction() (common.Address, error) {
	return _Liquidator.Contract.Interaction(&_Liquidator.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Liquidator *LiquidatorCaller) Lender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "lender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Liquidator *LiquidatorSession) Lender() (common.Address, error) {
	return _Liquidator.Contract.Lender(&_Liquidator.CallOpts)
}

// Lender is a free data retrieval call binding the contract method 0xbcead63e.
//
// Solidity: function lender() view returns(address)
func (_Liquidator *LiquidatorCallerSession) Lender() (common.Address, error) {
	return _Liquidator.Contract.Lender(&_Liquidator.CallOpts)
}

// LisUSD is a free data retrieval call binding the contract method 0xad15502b.
//
// Solidity: function lisUSD() view returns(address)
func (_Liquidator *LiquidatorCaller) LisUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "lisUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LisUSD is a free data retrieval call binding the contract method 0xad15502b.
//
// Solidity: function lisUSD() view returns(address)
func (_Liquidator *LiquidatorSession) LisUSD() (common.Address, error) {
	return _Liquidator.Contract.LisUSD(&_Liquidator.CallOpts)
}

// LisUSD is a free data retrieval call binding the contract method 0xad15502b.
//
// Solidity: function lisUSD() view returns(address)
func (_Liquidator *LiquidatorCallerSession) LisUSD() (common.Address, error) {
	return _Liquidator.Contract.LisUSD(&_Liquidator.CallOpts)
}

// PairWhitelist is a free data retrieval call binding the contract method 0xd06ba892.
//
// Solidity: function pairWhitelist(address ) view returns(bool)
func (_Liquidator *LiquidatorCaller) PairWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "pairWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PairWhitelist is a free data retrieval call binding the contract method 0xd06ba892.
//
// Solidity: function pairWhitelist(address ) view returns(bool)
func (_Liquidator *LiquidatorSession) PairWhitelist(arg0 common.Address) (bool, error) {
	return _Liquidator.Contract.PairWhitelist(&_Liquidator.CallOpts, arg0)
}

// PairWhitelist is a free data retrieval call binding the contract method 0xd06ba892.
//
// Solidity: function pairWhitelist(address ) view returns(bool)
func (_Liquidator *LiquidatorCallerSession) PairWhitelist(arg0 common.Address) (bool, error) {
	return _Liquidator.Contract.PairWhitelist(&_Liquidator.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Liquidator *LiquidatorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Liquidator *LiquidatorSession) ProxiableUUID() ([32]byte, error) {
	return _Liquidator.Contract.ProxiableUUID(&_Liquidator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Liquidator *LiquidatorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Liquidator.Contract.ProxiableUUID(&_Liquidator.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Liquidator *LiquidatorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Liquidator *LiquidatorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Liquidator.Contract.SupportsInterface(&_Liquidator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Liquidator *LiquidatorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Liquidator.Contract.SupportsInterface(&_Liquidator.CallOpts, interfaceId)
}

// TokenWhitelist is a free data retrieval call binding the contract method 0x753d7563.
//
// Solidity: function tokenWhitelist(address ) view returns(bool)
func (_Liquidator *LiquidatorCaller) TokenWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Liquidator.contract.Call(opts, &out, "tokenWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenWhitelist is a free data retrieval call binding the contract method 0x753d7563.
//
// Solidity: function tokenWhitelist(address ) view returns(bool)
func (_Liquidator *LiquidatorSession) TokenWhitelist(arg0 common.Address) (bool, error) {
	return _Liquidator.Contract.TokenWhitelist(&_Liquidator.CallOpts, arg0)
}

// TokenWhitelist is a free data retrieval call binding the contract method 0x753d7563.
//
// Solidity: function tokenWhitelist(address ) view returns(bool)
func (_Liquidator *LiquidatorCallerSession) TokenWhitelist(arg0 common.Address) (bool, error) {
	return _Liquidator.Contract.TokenWhitelist(&_Liquidator.CallOpts, arg0)
}

// FlashLiquidate is a paid mutator transaction binding the contract method 0x7209d2c7.
//
// Solidity: function flashLiquidate(uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, address collateralReal, address pair, bytes swapData) returns()
func (_Liquidator *LiquidatorTransactor) FlashLiquidate(opts *bind.TransactOpts, auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, collateralReal common.Address, pair common.Address, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "flashLiquidate", auctionId, borrowAm, collateral, collateralAm, maxPrice, collateralReal, pair, swapData)
}

// FlashLiquidate is a paid mutator transaction binding the contract method 0x7209d2c7.
//
// Solidity: function flashLiquidate(uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, address collateralReal, address pair, bytes swapData) returns()
func (_Liquidator *LiquidatorSession) FlashLiquidate(auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, collateralReal common.Address, pair common.Address, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.FlashLiquidate(&_Liquidator.TransactOpts, auctionId, borrowAm, collateral, collateralAm, maxPrice, collateralReal, pair, swapData)
}

// FlashLiquidate is a paid mutator transaction binding the contract method 0x7209d2c7.
//
// Solidity: function flashLiquidate(uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, address collateralReal, address pair, bytes swapData) returns()
func (_Liquidator *LiquidatorTransactorSession) FlashLiquidate(auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, collateralReal common.Address, pair common.Address, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.FlashLiquidate(&_Liquidator.TransactOpts, auctionId, borrowAm, collateral, collateralAm, maxPrice, collateralReal, pair, swapData)
}

// FlashLiquidate0 is a paid mutator transaction binding the contract method 0xda92259a.
//
// Solidity: function flashLiquidate(uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, address collateralReal, address pair, address spender, bytes swapData) returns()
func (_Liquidator *LiquidatorTransactor) FlashLiquidate0(opts *bind.TransactOpts, auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, collateralReal common.Address, pair common.Address, spender common.Address, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "flashLiquidate0", auctionId, borrowAm, collateral, collateralAm, maxPrice, collateralReal, pair, spender, swapData)
}

// FlashLiquidate0 is a paid mutator transaction binding the contract method 0xda92259a.
//
// Solidity: function flashLiquidate(uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, address collateralReal, address pair, address spender, bytes swapData) returns()
func (_Liquidator *LiquidatorSession) FlashLiquidate0(auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, collateralReal common.Address, pair common.Address, spender common.Address, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.FlashLiquidate0(&_Liquidator.TransactOpts, auctionId, borrowAm, collateral, collateralAm, maxPrice, collateralReal, pair, spender, swapData)
}

// FlashLiquidate0 is a paid mutator transaction binding the contract method 0xda92259a.
//
// Solidity: function flashLiquidate(uint256 auctionId, uint256 borrowAm, address collateral, uint256 collateralAm, uint256 maxPrice, address collateralReal, address pair, address spender, bytes swapData) returns()
func (_Liquidator *LiquidatorTransactorSession) FlashLiquidate0(auctionId *big.Int, borrowAm *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int, collateralReal common.Address, pair common.Address, spender common.Address, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.FlashLiquidate0(&_Liquidator.TransactOpts, auctionId, borrowAm, collateral, collateralAm, maxPrice, collateralReal, pair, spender, swapData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.GrantRole(&_Liquidator.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.GrantRole(&_Liquidator.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address admin, address manager, address bot, address _lender, address _interaction, address _lisUSD) returns()
func (_Liquidator *LiquidatorTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, manager common.Address, bot common.Address, _lender common.Address, _interaction common.Address, _lisUSD common.Address) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "initialize", admin, manager, bot, _lender, _interaction, _lisUSD)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address admin, address manager, address bot, address _lender, address _interaction, address _lisUSD) returns()
func (_Liquidator *LiquidatorSession) Initialize(admin common.Address, manager common.Address, bot common.Address, _lender common.Address, _interaction common.Address, _lisUSD common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.Initialize(&_Liquidator.TransactOpts, admin, manager, bot, _lender, _interaction, _lisUSD)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address admin, address manager, address bot, address _lender, address _interaction, address _lisUSD) returns()
func (_Liquidator *LiquidatorTransactorSession) Initialize(admin common.Address, manager common.Address, bot common.Address, _lender common.Address, _interaction common.Address, _lisUSD common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.Initialize(&_Liquidator.TransactOpts, admin, manager, bot, _lender, _interaction, _lisUSD)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0fb548e9.
//
// Solidity: function liquidate(uint256 auctionId, address collateral, uint256 collateralAm, uint256 maxPrice) returns()
func (_Liquidator *LiquidatorTransactor) Liquidate(opts *bind.TransactOpts, auctionId *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "liquidate", auctionId, collateral, collateralAm, maxPrice)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0fb548e9.
//
// Solidity: function liquidate(uint256 auctionId, address collateral, uint256 collateralAm, uint256 maxPrice) returns()
func (_Liquidator *LiquidatorSession) Liquidate(auctionId *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Liquidator.Contract.Liquidate(&_Liquidator.TransactOpts, auctionId, collateral, collateralAm, maxPrice)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0fb548e9.
//
// Solidity: function liquidate(uint256 auctionId, address collateral, uint256 collateralAm, uint256 maxPrice) returns()
func (_Liquidator *LiquidatorTransactorSession) Liquidate(auctionId *big.Int, collateral common.Address, collateralAm *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Liquidator.Contract.Liquidate(&_Liquidator.TransactOpts, auctionId, collateral, collateralAm, maxPrice)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Liquidator *LiquidatorTransactor) OnFlashLoan(opts *bind.TransactOpts, initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "onFlashLoan", initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Liquidator *LiquidatorSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.OnFlashLoan(&_Liquidator.TransactOpts, initiator, token, amount, fee, data)
}

// OnFlashLoan is a paid mutator transaction binding the contract method 0x23e30c8b.
//
// Solidity: function onFlashLoan(address initiator, address token, uint256 amount, uint256 fee, bytes data) returns(bytes32)
func (_Liquidator *LiquidatorTransactorSession) OnFlashLoan(initiator common.Address, token common.Address, amount *big.Int, fee *big.Int, data []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.OnFlashLoan(&_Liquidator.TransactOpts, initiator, token, amount, fee, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.RenounceRole(&_Liquidator.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.RenounceRole(&_Liquidator.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.RevokeRole(&_Liquidator.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Liquidator *LiquidatorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.RevokeRole(&_Liquidator.TransactOpts, role, account)
}

// SellToken is a paid mutator transaction binding the contract method 0x873bf8a7.
//
// Solidity: function sellToken(address pair, address tokenIn, uint256 amountIn, uint256 amountOutMin, bytes swapData) returns()
func (_Liquidator *LiquidatorTransactor) SellToken(opts *bind.TransactOpts, pair common.Address, tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "sellToken", pair, tokenIn, amountIn, amountOutMin, swapData)
}

// SellToken is a paid mutator transaction binding the contract method 0x873bf8a7.
//
// Solidity: function sellToken(address pair, address tokenIn, uint256 amountIn, uint256 amountOutMin, bytes swapData) returns()
func (_Liquidator *LiquidatorSession) SellToken(pair common.Address, tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.SellToken(&_Liquidator.TransactOpts, pair, tokenIn, amountIn, amountOutMin, swapData)
}

// SellToken is a paid mutator transaction binding the contract method 0x873bf8a7.
//
// Solidity: function sellToken(address pair, address tokenIn, uint256 amountIn, uint256 amountOutMin, bytes swapData) returns()
func (_Liquidator *LiquidatorTransactorSession) SellToken(pair common.Address, tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.SellToken(&_Liquidator.TransactOpts, pair, tokenIn, amountIn, amountOutMin, swapData)
}

// SellToken0 is a paid mutator transaction binding the contract method 0xeb2f51a6.
//
// Solidity: function sellToken(address pair, address spender, address tokenIn, uint256 amountIn, uint256 amountOutMin, bytes swapData) returns()
func (_Liquidator *LiquidatorTransactor) SellToken0(opts *bind.TransactOpts, pair common.Address, spender common.Address, tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "sellToken0", pair, spender, tokenIn, amountIn, amountOutMin, swapData)
}

// SellToken0 is a paid mutator transaction binding the contract method 0xeb2f51a6.
//
// Solidity: function sellToken(address pair, address spender, address tokenIn, uint256 amountIn, uint256 amountOutMin, bytes swapData) returns()
func (_Liquidator *LiquidatorSession) SellToken0(pair common.Address, spender common.Address, tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.SellToken0(&_Liquidator.TransactOpts, pair, spender, tokenIn, amountIn, amountOutMin, swapData)
}

// SellToken0 is a paid mutator transaction binding the contract method 0xeb2f51a6.
//
// Solidity: function sellToken(address pair, address spender, address tokenIn, uint256 amountIn, uint256 amountOutMin, bytes swapData) returns()
func (_Liquidator *LiquidatorTransactorSession) SellToken0(pair common.Address, spender common.Address, tokenIn common.Address, amountIn *big.Int, amountOutMin *big.Int, swapData []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.SellToken0(&_Liquidator.TransactOpts, pair, spender, tokenIn, amountIn, amountOutMin, swapData)
}

// SetPairWhitelist is a paid mutator transaction binding the contract method 0xbf75ca41.
//
// Solidity: function setPairWhitelist(address pair, bool status) returns()
func (_Liquidator *LiquidatorTransactor) SetPairWhitelist(opts *bind.TransactOpts, pair common.Address, status bool) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "setPairWhitelist", pair, status)
}

// SetPairWhitelist is a paid mutator transaction binding the contract method 0xbf75ca41.
//
// Solidity: function setPairWhitelist(address pair, bool status) returns()
func (_Liquidator *LiquidatorSession) SetPairWhitelist(pair common.Address, status bool) (*types.Transaction, error) {
	return _Liquidator.Contract.SetPairWhitelist(&_Liquidator.TransactOpts, pair, status)
}

// SetPairWhitelist is a paid mutator transaction binding the contract method 0xbf75ca41.
//
// Solidity: function setPairWhitelist(address pair, bool status) returns()
func (_Liquidator *LiquidatorTransactorSession) SetPairWhitelist(pair common.Address, status bool) (*types.Transaction, error) {
	return _Liquidator.Contract.SetPairWhitelist(&_Liquidator.TransactOpts, pair, status)
}

// SetTokenWhitelist is a paid mutator transaction binding the contract method 0xc9bcc97e.
//
// Solidity: function setTokenWhitelist(address token, bool status) returns()
func (_Liquidator *LiquidatorTransactor) SetTokenWhitelist(opts *bind.TransactOpts, token common.Address, status bool) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "setTokenWhitelist", token, status)
}

// SetTokenWhitelist is a paid mutator transaction binding the contract method 0xc9bcc97e.
//
// Solidity: function setTokenWhitelist(address token, bool status) returns()
func (_Liquidator *LiquidatorSession) SetTokenWhitelist(token common.Address, status bool) (*types.Transaction, error) {
	return _Liquidator.Contract.SetTokenWhitelist(&_Liquidator.TransactOpts, token, status)
}

// SetTokenWhitelist is a paid mutator transaction binding the contract method 0xc9bcc97e.
//
// Solidity: function setTokenWhitelist(address token, bool status) returns()
func (_Liquidator *LiquidatorTransactorSession) SetTokenWhitelist(token common.Address, status bool) (*types.Transaction, error) {
	return _Liquidator.Contract.SetTokenWhitelist(&_Liquidator.TransactOpts, token, status)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Liquidator *LiquidatorTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Liquidator *LiquidatorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.UpgradeTo(&_Liquidator.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Liquidator *LiquidatorTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Liquidator.Contract.UpgradeTo(&_Liquidator.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Liquidator *LiquidatorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Liquidator *LiquidatorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.UpgradeToAndCall(&_Liquidator.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Liquidator *LiquidatorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Liquidator.Contract.UpgradeToAndCall(&_Liquidator.TransactOpts, newImplementation, data)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_Liquidator *LiquidatorTransactor) WithdrawERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidator.contract.Transact(opts, "withdrawERC20", token, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_Liquidator *LiquidatorSession) WithdrawERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidator.Contract.WithdrawERC20(&_Liquidator.TransactOpts, token, amount)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xa1db9782.
//
// Solidity: function withdrawERC20(address token, uint256 amount) returns()
func (_Liquidator *LiquidatorTransactorSession) WithdrawERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Liquidator.Contract.WithdrawERC20(&_Liquidator.TransactOpts, token, amount)
}

// LiquidatorAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Liquidator contract.
type LiquidatorAdminChangedIterator struct {
	Event *LiquidatorAdminChanged // Event containing the contract specifics and raw log

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
func (it *LiquidatorAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorAdminChanged)
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
		it.Event = new(LiquidatorAdminChanged)
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
func (it *LiquidatorAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorAdminChanged represents a AdminChanged event raised by the Liquidator contract.
type LiquidatorAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Liquidator *LiquidatorFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*LiquidatorAdminChangedIterator, error) {

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &LiquidatorAdminChangedIterator{contract: _Liquidator.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Liquidator *LiquidatorFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *LiquidatorAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorAdminChanged)
				if err := _Liquidator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Liquidator *LiquidatorFilterer) ParseAdminChanged(log types.Log) (*LiquidatorAdminChanged, error) {
	event := new(LiquidatorAdminChanged)
	if err := _Liquidator.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Liquidator contract.
type LiquidatorBeaconUpgradedIterator struct {
	Event *LiquidatorBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *LiquidatorBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorBeaconUpgraded)
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
		it.Event = new(LiquidatorBeaconUpgraded)
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
func (it *LiquidatorBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorBeaconUpgraded represents a BeaconUpgraded event raised by the Liquidator contract.
type LiquidatorBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Liquidator *LiquidatorFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*LiquidatorBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &LiquidatorBeaconUpgradedIterator{contract: _Liquidator.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Liquidator *LiquidatorFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *LiquidatorBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorBeaconUpgraded)
				if err := _Liquidator.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Liquidator *LiquidatorFilterer) ParseBeaconUpgraded(log types.Log) (*LiquidatorBeaconUpgraded, error) {
	event := new(LiquidatorBeaconUpgraded)
	if err := _Liquidator.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Liquidator contract.
type LiquidatorInitializedIterator struct {
	Event *LiquidatorInitialized // Event containing the contract specifics and raw log

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
func (it *LiquidatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorInitialized)
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
		it.Event = new(LiquidatorInitialized)
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
func (it *LiquidatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorInitialized represents a Initialized event raised by the Liquidator contract.
type LiquidatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Liquidator *LiquidatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*LiquidatorInitializedIterator, error) {

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LiquidatorInitializedIterator{contract: _Liquidator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Liquidator *LiquidatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LiquidatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorInitialized)
				if err := _Liquidator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Liquidator *LiquidatorFilterer) ParseInitialized(log types.Log) (*LiquidatorInitialized, error) {
	event := new(LiquidatorInitialized)
	if err := _Liquidator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorPairWhitelistChangedIterator is returned from FilterPairWhitelistChanged and is used to iterate over the raw logs and unpacked data for PairWhitelistChanged events raised by the Liquidator contract.
type LiquidatorPairWhitelistChangedIterator struct {
	Event *LiquidatorPairWhitelistChanged // Event containing the contract specifics and raw log

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
func (it *LiquidatorPairWhitelistChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorPairWhitelistChanged)
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
		it.Event = new(LiquidatorPairWhitelistChanged)
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
func (it *LiquidatorPairWhitelistChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorPairWhitelistChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorPairWhitelistChanged represents a PairWhitelistChanged event raised by the Liquidator contract.
type LiquidatorPairWhitelistChanged struct {
	Pair  common.Address
	Added bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPairWhitelistChanged is a free log retrieval operation binding the contract event 0xf84396a18addacf98f41e3fb196a7ddb639ee0c11d3d790510e2eac3e97549a3.
//
// Solidity: event PairWhitelistChanged(address pair, bool added)
func (_Liquidator *LiquidatorFilterer) FilterPairWhitelistChanged(opts *bind.FilterOpts) (*LiquidatorPairWhitelistChangedIterator, error) {

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "PairWhitelistChanged")
	if err != nil {
		return nil, err
	}
	return &LiquidatorPairWhitelistChangedIterator{contract: _Liquidator.contract, event: "PairWhitelistChanged", logs: logs, sub: sub}, nil
}

// WatchPairWhitelistChanged is a free log subscription operation binding the contract event 0xf84396a18addacf98f41e3fb196a7ddb639ee0c11d3d790510e2eac3e97549a3.
//
// Solidity: event PairWhitelistChanged(address pair, bool added)
func (_Liquidator *LiquidatorFilterer) WatchPairWhitelistChanged(opts *bind.WatchOpts, sink chan<- *LiquidatorPairWhitelistChanged) (event.Subscription, error) {

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "PairWhitelistChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorPairWhitelistChanged)
				if err := _Liquidator.contract.UnpackLog(event, "PairWhitelistChanged", log); err != nil {
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

// ParsePairWhitelistChanged is a log parse operation binding the contract event 0xf84396a18addacf98f41e3fb196a7ddb639ee0c11d3d790510e2eac3e97549a3.
//
// Solidity: event PairWhitelistChanged(address pair, bool added)
func (_Liquidator *LiquidatorFilterer) ParsePairWhitelistChanged(log types.Log) (*LiquidatorPairWhitelistChanged, error) {
	event := new(LiquidatorPairWhitelistChanged)
	if err := _Liquidator.contract.UnpackLog(event, "PairWhitelistChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Liquidator contract.
type LiquidatorRoleAdminChangedIterator struct {
	Event *LiquidatorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LiquidatorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorRoleAdminChanged)
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
		it.Event = new(LiquidatorRoleAdminChanged)
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
func (it *LiquidatorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorRoleAdminChanged represents a RoleAdminChanged event raised by the Liquidator contract.
type LiquidatorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Liquidator *LiquidatorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LiquidatorRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LiquidatorRoleAdminChangedIterator{contract: _Liquidator.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Liquidator *LiquidatorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LiquidatorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorRoleAdminChanged)
				if err := _Liquidator.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Liquidator *LiquidatorFilterer) ParseRoleAdminChanged(log types.Log) (*LiquidatorRoleAdminChanged, error) {
	event := new(LiquidatorRoleAdminChanged)
	if err := _Liquidator.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Liquidator contract.
type LiquidatorRoleGrantedIterator struct {
	Event *LiquidatorRoleGranted // Event containing the contract specifics and raw log

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
func (it *LiquidatorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorRoleGranted)
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
		it.Event = new(LiquidatorRoleGranted)
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
func (it *LiquidatorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorRoleGranted represents a RoleGranted event raised by the Liquidator contract.
type LiquidatorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Liquidator *LiquidatorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LiquidatorRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidatorRoleGrantedIterator{contract: _Liquidator.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Liquidator *LiquidatorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LiquidatorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorRoleGranted)
				if err := _Liquidator.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Liquidator *LiquidatorFilterer) ParseRoleGranted(log types.Log) (*LiquidatorRoleGranted, error) {
	event := new(LiquidatorRoleGranted)
	if err := _Liquidator.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Liquidator contract.
type LiquidatorRoleRevokedIterator struct {
	Event *LiquidatorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LiquidatorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorRoleRevoked)
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
		it.Event = new(LiquidatorRoleRevoked)
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
func (it *LiquidatorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorRoleRevoked represents a RoleRevoked event raised by the Liquidator contract.
type LiquidatorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Liquidator *LiquidatorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LiquidatorRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LiquidatorRoleRevokedIterator{contract: _Liquidator.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Liquidator *LiquidatorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LiquidatorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorRoleRevoked)
				if err := _Liquidator.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Liquidator *LiquidatorFilterer) ParseRoleRevoked(log types.Log) (*LiquidatorRoleRevoked, error) {
	event := new(LiquidatorRoleRevoked)
	if err := _Liquidator.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorSellTokenIterator is returned from FilterSellToken and is used to iterate over the raw logs and unpacked data for SellToken events raised by the Liquidator contract.
type LiquidatorSellTokenIterator struct {
	Event *LiquidatorSellToken // Event containing the contract specifics and raw log

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
func (it *LiquidatorSellTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorSellToken)
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
		it.Event = new(LiquidatorSellToken)
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
func (it *LiquidatorSellTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorSellTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorSellToken represents a SellToken event raised by the Liquidator contract.
type LiquidatorSellToken struct {
	Pair         common.Address
	TokenIn      common.Address
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSellToken is a free log retrieval operation binding the contract event 0xbbcd6525e9abb18370c96747b5bdabcbfb595afb38cb853a1ed34b210ad437c6.
//
// Solidity: event SellToken(address pair, address tokenIn, uint256 amountIn, uint256 amountOutMin)
func (_Liquidator *LiquidatorFilterer) FilterSellToken(opts *bind.FilterOpts) (*LiquidatorSellTokenIterator, error) {

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "SellToken")
	if err != nil {
		return nil, err
	}
	return &LiquidatorSellTokenIterator{contract: _Liquidator.contract, event: "SellToken", logs: logs, sub: sub}, nil
}

// WatchSellToken is a free log subscription operation binding the contract event 0xbbcd6525e9abb18370c96747b5bdabcbfb595afb38cb853a1ed34b210ad437c6.
//
// Solidity: event SellToken(address pair, address tokenIn, uint256 amountIn, uint256 amountOutMin)
func (_Liquidator *LiquidatorFilterer) WatchSellToken(opts *bind.WatchOpts, sink chan<- *LiquidatorSellToken) (event.Subscription, error) {

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "SellToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorSellToken)
				if err := _Liquidator.contract.UnpackLog(event, "SellToken", log); err != nil {
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

// ParseSellToken is a log parse operation binding the contract event 0xbbcd6525e9abb18370c96747b5bdabcbfb595afb38cb853a1ed34b210ad437c6.
//
// Solidity: event SellToken(address pair, address tokenIn, uint256 amountIn, uint256 amountOutMin)
func (_Liquidator *LiquidatorFilterer) ParseSellToken(log types.Log) (*LiquidatorSellToken, error) {
	event := new(LiquidatorSellToken)
	if err := _Liquidator.contract.UnpackLog(event, "SellToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorTokenWhitelistChangedIterator is returned from FilterTokenWhitelistChanged and is used to iterate over the raw logs and unpacked data for TokenWhitelistChanged events raised by the Liquidator contract.
type LiquidatorTokenWhitelistChangedIterator struct {
	Event *LiquidatorTokenWhitelistChanged // Event containing the contract specifics and raw log

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
func (it *LiquidatorTokenWhitelistChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorTokenWhitelistChanged)
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
		it.Event = new(LiquidatorTokenWhitelistChanged)
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
func (it *LiquidatorTokenWhitelistChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorTokenWhitelistChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorTokenWhitelistChanged represents a TokenWhitelistChanged event raised by the Liquidator contract.
type LiquidatorTokenWhitelistChanged struct {
	Token common.Address
	Added bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenWhitelistChanged is a free log retrieval operation binding the contract event 0x6b2d0400db56a6ee7854a654e69e10f7fa8e3e8a75361c0134821522358f22a2.
//
// Solidity: event TokenWhitelistChanged(address indexed token, bool added)
func (_Liquidator *LiquidatorFilterer) FilterTokenWhitelistChanged(opts *bind.FilterOpts, token []common.Address) (*LiquidatorTokenWhitelistChangedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "TokenWhitelistChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return &LiquidatorTokenWhitelistChangedIterator{contract: _Liquidator.contract, event: "TokenWhitelistChanged", logs: logs, sub: sub}, nil
}

// WatchTokenWhitelistChanged is a free log subscription operation binding the contract event 0x6b2d0400db56a6ee7854a654e69e10f7fa8e3e8a75361c0134821522358f22a2.
//
// Solidity: event TokenWhitelistChanged(address indexed token, bool added)
func (_Liquidator *LiquidatorFilterer) WatchTokenWhitelistChanged(opts *bind.WatchOpts, sink chan<- *LiquidatorTokenWhitelistChanged, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "TokenWhitelistChanged", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorTokenWhitelistChanged)
				if err := _Liquidator.contract.UnpackLog(event, "TokenWhitelistChanged", log); err != nil {
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

// ParseTokenWhitelistChanged is a log parse operation binding the contract event 0x6b2d0400db56a6ee7854a654e69e10f7fa8e3e8a75361c0134821522358f22a2.
//
// Solidity: event TokenWhitelistChanged(address indexed token, bool added)
func (_Liquidator *LiquidatorFilterer) ParseTokenWhitelistChanged(log types.Log) (*LiquidatorTokenWhitelistChanged, error) {
	event := new(LiquidatorTokenWhitelistChanged)
	if err := _Liquidator.contract.UnpackLog(event, "TokenWhitelistChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidatorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Liquidator contract.
type LiquidatorUpgradedIterator struct {
	Event *LiquidatorUpgraded // Event containing the contract specifics and raw log

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
func (it *LiquidatorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidatorUpgraded)
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
		it.Event = new(LiquidatorUpgraded)
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
func (it *LiquidatorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidatorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidatorUpgraded represents a Upgraded event raised by the Liquidator contract.
type LiquidatorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Liquidator *LiquidatorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*LiquidatorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Liquidator.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &LiquidatorUpgradedIterator{contract: _Liquidator.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Liquidator *LiquidatorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *LiquidatorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Liquidator.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidatorUpgraded)
				if err := _Liquidator.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Liquidator *LiquidatorFilterer) ParseUpgraded(log types.Log) (*LiquidatorUpgraded, error) {
	event := new(LiquidatorUpgraded)
	if err := _Liquidator.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
