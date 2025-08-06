// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package provider

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

// ProviderMetaData contains all meta data concerning the Provider contract.
var ProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cdp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nonFungiblePositionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_masterChefV3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lpUsd\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpValue\",\"type\":\"uint256\"}],\"name\":\"DepositLp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userLpTotalValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userLockedCollateral\",\"type\":\"uint256\"}],\"name\":\"Liquidatable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLeftover\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token0Left\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"token1Left\",\"type\":\"uint256\"}],\"name\":\"Liquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidatableAmount\",\"type\":\"uint256\"}],\"name\":\"LiquidationBegan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldExchangeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newExchangeRate\",\"type\":\"uint256\"}],\"name\":\"LpDiscountRateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaxLpValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxLpValue\",\"type\":\"uint256\"}],\"name\":\"MaxLpPerUserSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMinLpValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinLpValue\",\"type\":\"uint256\"}],\"name\":\"MinLpValueSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userTotalLpValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userCdpPositionValue\",\"type\":\"uint256\"}],\"name\":\"UserCdpPositionSynced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpValue\",\"type\":\"uint256\"}],\"name\":\"WithdrawLp\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESILIENT_ORACLE_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"batchSyncUserLpValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cdp\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"claimableStakingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"}],\"name\":\"daoBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getLatestUserTotalLpValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"userLpTotalValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLpValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"appraisedValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserLps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bot\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pauser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pancakeStakingHub\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pancakeStakingVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resilientOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxLpPerUser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minLpValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lpDiscountRate\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isLeftOver\",\"type\":\"bool\"}],\"name\":\"liquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpDiscountRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpUsd\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lpValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"masterChefV3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLpPerUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minLpValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonFungiblePositionManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeLpStakingVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pancakeStakingHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"provide\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resilientOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lpDiscountRate\",\"type\":\"uint256\"}],\"name\":\"setLpDiscountRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLpPerUser\",\"type\":\"uint256\"}],\"name\":\"setMaxLpPerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minLpValue\",\"type\":\"uint256\"}],\"name\":\"setMinLpValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"syncUserLpValues\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLiquidations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ongoing\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"token0Left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Left\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userLps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userTotalLpValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"vaultClaimStakingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalReward\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use ProviderMetaData.ABI instead.
var ProviderABI = ProviderMetaData.ABI

// Provider is an auto generated Go binding around an Ethereum contract.
type Provider struct {
	ProviderCaller     // Read-only binding to the contract
	ProviderTransactor // Write-only binding to the contract
	ProviderFilterer   // Log filterer for contract events
}

// ProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProviderSession struct {
	Contract     *Provider         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProviderCallerSession struct {
	Contract *ProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProviderTransactorSession struct {
	Contract     *ProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProviderRaw struct {
	Contract *Provider // Generic contract binding to access the raw methods on
}

// ProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProviderCallerRaw struct {
	Contract *ProviderCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProviderTransactorRaw struct {
	Contract *ProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProvider creates a new instance of Provider, bound to a specific deployed contract.
func NewProvider(address common.Address, backend bind.ContractBackend) (*Provider, error) {
	contract, err := bindProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Provider{ProviderCaller: ProviderCaller{contract: contract}, ProviderTransactor: ProviderTransactor{contract: contract}, ProviderFilterer: ProviderFilterer{contract: contract}}, nil
}

// NewProviderCaller creates a new read-only instance of Provider, bound to a specific deployed contract.
func NewProviderCaller(address common.Address, caller bind.ContractCaller) (*ProviderCaller, error) {
	contract, err := bindProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderCaller{contract: contract}, nil
}

// NewProviderTransactor creates a new write-only instance of Provider, bound to a specific deployed contract.
func NewProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderTransactor, error) {
	contract, err := bindProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderTransactor{contract: contract}, nil
}

// NewProviderFilterer creates a new log filterer instance of Provider, bound to a specific deployed contract.
func NewProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderFilterer, error) {
	contract, err := bindProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderFilterer{contract: contract}, nil
}

// bindProvider binds a generic wrapper to an already deployed contract.
func bindProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provider *ProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provider.Contract.ProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provider *ProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.Contract.ProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provider *ProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provider.Contract.ProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provider *ProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provider *ProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provider *ProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provider.Contract.contract.Transact(opts, method, params...)
}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Provider *ProviderCaller) BOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "BOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Provider *ProviderSession) BOT() ([32]byte, error) {
	return _Provider.Contract.BOT(&_Provider.CallOpts)
}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Provider *ProviderCallerSession) BOT() ([32]byte, error) {
	return _Provider.Contract.BOT(&_Provider.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Provider *ProviderCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Provider *ProviderSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Provider.Contract.DEFAULTADMINROLE(&_Provider.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Provider *ProviderCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Provider.Contract.DEFAULTADMINROLE(&_Provider.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_Provider *ProviderCaller) DENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_Provider *ProviderSession) DENOMINATOR() (*big.Int, error) {
	return _Provider.Contract.DENOMINATOR(&_Provider.CallOpts)
}

// DENOMINATOR is a free data retrieval call binding the contract method 0x918f8674.
//
// Solidity: function DENOMINATOR() view returns(uint256)
func (_Provider *ProviderCallerSession) DENOMINATOR() (*big.Int, error) {
	return _Provider.Contract.DENOMINATOR(&_Provider.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Provider *ProviderCaller) MANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Provider *ProviderSession) MANAGER() ([32]byte, error) {
	return _Provider.Contract.MANAGER(&_Provider.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Provider *ProviderCallerSession) MANAGER() ([32]byte, error) {
	return _Provider.Contract.MANAGER(&_Provider.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_Provider *ProviderCaller) PAUSER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "PAUSER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_Provider *ProviderSession) PAUSER() ([32]byte, error) {
	return _Provider.Contract.PAUSER(&_Provider.CallOpts)
}

// PAUSER is a free data retrieval call binding the contract method 0xd9dc8694.
//
// Solidity: function PAUSER() view returns(bytes32)
func (_Provider *ProviderCallerSession) PAUSER() ([32]byte, error) {
	return _Provider.Contract.PAUSER(&_Provider.CallOpts)
}

// RESILIENTORACLEDECIMALS is a free data retrieval call binding the contract method 0xb532c078.
//
// Solidity: function RESILIENT_ORACLE_DECIMALS() view returns(uint256)
func (_Provider *ProviderCaller) RESILIENTORACLEDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "RESILIENT_ORACLE_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RESILIENTORACLEDECIMALS is a free data retrieval call binding the contract method 0xb532c078.
//
// Solidity: function RESILIENT_ORACLE_DECIMALS() view returns(uint256)
func (_Provider *ProviderSession) RESILIENTORACLEDECIMALS() (*big.Int, error) {
	return _Provider.Contract.RESILIENTORACLEDECIMALS(&_Provider.CallOpts)
}

// RESILIENTORACLEDECIMALS is a free data retrieval call binding the contract method 0xb532c078.
//
// Solidity: function RESILIENT_ORACLE_DECIMALS() view returns(uint256)
func (_Provider *ProviderCallerSession) RESILIENTORACLEDECIMALS() (*big.Int, error) {
	return _Provider.Contract.RESILIENTORACLEDECIMALS(&_Provider.CallOpts)
}

// Cdp is a free data retrieval call binding the contract method 0x98bc8692.
//
// Solidity: function cdp() view returns(address)
func (_Provider *ProviderCaller) Cdp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "cdp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Cdp is a free data retrieval call binding the contract method 0x98bc8692.
//
// Solidity: function cdp() view returns(address)
func (_Provider *ProviderSession) Cdp() (common.Address, error) {
	return _Provider.Contract.Cdp(&_Provider.CallOpts)
}

// Cdp is a free data retrieval call binding the contract method 0x98bc8692.
//
// Solidity: function cdp() view returns(address)
func (_Provider *ProviderCallerSession) Cdp() (common.Address, error) {
	return _Provider.Contract.Cdp(&_Provider.CallOpts)
}

// ClaimableStakingRewards is a free data retrieval call binding the contract method 0xd8d97417.
//
// Solidity: function claimableStakingRewards(address account) view returns(uint256)
func (_Provider *ProviderCaller) ClaimableStakingRewards(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "claimableStakingRewards", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableStakingRewards is a free data retrieval call binding the contract method 0xd8d97417.
//
// Solidity: function claimableStakingRewards(address account) view returns(uint256)
func (_Provider *ProviderSession) ClaimableStakingRewards(account common.Address) (*big.Int, error) {
	return _Provider.Contract.ClaimableStakingRewards(&_Provider.CallOpts, account)
}

// ClaimableStakingRewards is a free data retrieval call binding the contract method 0xd8d97417.
//
// Solidity: function claimableStakingRewards(address account) view returns(uint256)
func (_Provider *ProviderCallerSession) ClaimableStakingRewards(account common.Address) (*big.Int, error) {
	return _Provider.Contract.ClaimableStakingRewards(&_Provider.CallOpts, account)
}

// GetAmounts is a free data retrieval call binding the contract method 0x59f7cf50.
//
// Solidity: function getAmounts(uint256 tokenId) view returns(uint256 amount0, uint256 amount1)
func (_Provider *ProviderCaller) GetAmounts(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getAmounts", tokenId)

	outstruct := new(struct {
		Amount0 *big.Int
		Amount1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmounts is a free data retrieval call binding the contract method 0x59f7cf50.
//
// Solidity: function getAmounts(uint256 tokenId) view returns(uint256 amount0, uint256 amount1)
func (_Provider *ProviderSession) GetAmounts(tokenId *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _Provider.Contract.GetAmounts(&_Provider.CallOpts, tokenId)
}

// GetAmounts is a free data retrieval call binding the contract method 0x59f7cf50.
//
// Solidity: function getAmounts(uint256 tokenId) view returns(uint256 amount0, uint256 amount1)
func (_Provider *ProviderCallerSession) GetAmounts(tokenId *big.Int) (struct {
	Amount0 *big.Int
	Amount1 *big.Int
}, error) {
	return _Provider.Contract.GetAmounts(&_Provider.CallOpts, tokenId)
}

// GetLatestUserTotalLpValue is a free data retrieval call binding the contract method 0x55f7e0bc.
//
// Solidity: function getLatestUserTotalLpValue(address user) view returns(uint256 userLpTotalValue)
func (_Provider *ProviderCaller) GetLatestUserTotalLpValue(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getLatestUserTotalLpValue", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestUserTotalLpValue is a free data retrieval call binding the contract method 0x55f7e0bc.
//
// Solidity: function getLatestUserTotalLpValue(address user) view returns(uint256 userLpTotalValue)
func (_Provider *ProviderSession) GetLatestUserTotalLpValue(user common.Address) (*big.Int, error) {
	return _Provider.Contract.GetLatestUserTotalLpValue(&_Provider.CallOpts, user)
}

// GetLatestUserTotalLpValue is a free data retrieval call binding the contract method 0x55f7e0bc.
//
// Solidity: function getLatestUserTotalLpValue(address user) view returns(uint256 userLpTotalValue)
func (_Provider *ProviderCallerSession) GetLatestUserTotalLpValue(user common.Address) (*big.Int, error) {
	return _Provider.Contract.GetLatestUserTotalLpValue(&_Provider.CallOpts, user)
}

// GetLpValue is a free data retrieval call binding the contract method 0x058062b3.
//
// Solidity: function getLpValue(uint256 tokenId) view returns(uint256 appraisedValue)
func (_Provider *ProviderCaller) GetLpValue(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getLpValue", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLpValue is a free data retrieval call binding the contract method 0x058062b3.
//
// Solidity: function getLpValue(uint256 tokenId) view returns(uint256 appraisedValue)
func (_Provider *ProviderSession) GetLpValue(tokenId *big.Int) (*big.Int, error) {
	return _Provider.Contract.GetLpValue(&_Provider.CallOpts, tokenId)
}

// GetLpValue is a free data retrieval call binding the contract method 0x058062b3.
//
// Solidity: function getLpValue(uint256 tokenId) view returns(uint256 appraisedValue)
func (_Provider *ProviderCallerSession) GetLpValue(tokenId *big.Int) (*big.Int, error) {
	return _Provider.Contract.GetLpValue(&_Provider.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Provider *ProviderCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Provider *ProviderSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Provider.Contract.GetRoleAdmin(&_Provider.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Provider *ProviderCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Provider.Contract.GetRoleAdmin(&_Provider.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Provider *ProviderCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Provider *ProviderSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Provider.Contract.GetRoleMember(&_Provider.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Provider *ProviderCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Provider.Contract.GetRoleMember(&_Provider.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Provider *ProviderCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Provider *ProviderSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Provider.Contract.GetRoleMemberCount(&_Provider.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Provider *ProviderCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Provider.Contract.GetRoleMemberCount(&_Provider.CallOpts, role)
}

// GetUserLps is a free data retrieval call binding the contract method 0x7c6cf832.
//
// Solidity: function getUserLps(address user) view returns(uint256[])
func (_Provider *ProviderCaller) GetUserLps(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "getUserLps", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetUserLps is a free data retrieval call binding the contract method 0x7c6cf832.
//
// Solidity: function getUserLps(address user) view returns(uint256[])
func (_Provider *ProviderSession) GetUserLps(user common.Address) ([]*big.Int, error) {
	return _Provider.Contract.GetUserLps(&_Provider.CallOpts, user)
}

// GetUserLps is a free data retrieval call binding the contract method 0x7c6cf832.
//
// Solidity: function getUserLps(address user) view returns(uint256[])
func (_Provider *ProviderCallerSession) GetUserLps(user common.Address) ([]*big.Int, error) {
	return _Provider.Contract.GetUserLps(&_Provider.CallOpts, user)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Provider *ProviderCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Provider *ProviderSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Provider.Contract.HasRole(&_Provider.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Provider *ProviderCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Provider.Contract.HasRole(&_Provider.CallOpts, role, account)
}

// LpDiscountRate is a free data retrieval call binding the contract method 0x4c844bd9.
//
// Solidity: function lpDiscountRate() view returns(uint256)
func (_Provider *ProviderCaller) LpDiscountRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "lpDiscountRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpDiscountRate is a free data retrieval call binding the contract method 0x4c844bd9.
//
// Solidity: function lpDiscountRate() view returns(uint256)
func (_Provider *ProviderSession) LpDiscountRate() (*big.Int, error) {
	return _Provider.Contract.LpDiscountRate(&_Provider.CallOpts)
}

// LpDiscountRate is a free data retrieval call binding the contract method 0x4c844bd9.
//
// Solidity: function lpDiscountRate() view returns(uint256)
func (_Provider *ProviderCallerSession) LpDiscountRate() (*big.Int, error) {
	return _Provider.Contract.LpDiscountRate(&_Provider.CallOpts)
}

// LpOwners is a free data retrieval call binding the contract method 0x8f718f3f.
//
// Solidity: function lpOwners(uint256 ) view returns(address)
func (_Provider *ProviderCaller) LpOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "lpOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpOwners is a free data retrieval call binding the contract method 0x8f718f3f.
//
// Solidity: function lpOwners(uint256 ) view returns(address)
func (_Provider *ProviderSession) LpOwners(arg0 *big.Int) (common.Address, error) {
	return _Provider.Contract.LpOwners(&_Provider.CallOpts, arg0)
}

// LpOwners is a free data retrieval call binding the contract method 0x8f718f3f.
//
// Solidity: function lpOwners(uint256 ) view returns(address)
func (_Provider *ProviderCallerSession) LpOwners(arg0 *big.Int) (common.Address, error) {
	return _Provider.Contract.LpOwners(&_Provider.CallOpts, arg0)
}

// LpUsd is a free data retrieval call binding the contract method 0x42f47187.
//
// Solidity: function lpUsd() view returns(address)
func (_Provider *ProviderCaller) LpUsd(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "lpUsd")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpUsd is a free data retrieval call binding the contract method 0x42f47187.
//
// Solidity: function lpUsd() view returns(address)
func (_Provider *ProviderSession) LpUsd() (common.Address, error) {
	return _Provider.Contract.LpUsd(&_Provider.CallOpts)
}

// LpUsd is a free data retrieval call binding the contract method 0x42f47187.
//
// Solidity: function lpUsd() view returns(address)
func (_Provider *ProviderCallerSession) LpUsd() (common.Address, error) {
	return _Provider.Contract.LpUsd(&_Provider.CallOpts)
}

// LpValues is a free data retrieval call binding the contract method 0x015f3073.
//
// Solidity: function lpValues(uint256 ) view returns(uint256)
func (_Provider *ProviderCaller) LpValues(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "lpValues", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpValues is a free data retrieval call binding the contract method 0x015f3073.
//
// Solidity: function lpValues(uint256 ) view returns(uint256)
func (_Provider *ProviderSession) LpValues(arg0 *big.Int) (*big.Int, error) {
	return _Provider.Contract.LpValues(&_Provider.CallOpts, arg0)
}

// LpValues is a free data retrieval call binding the contract method 0x015f3073.
//
// Solidity: function lpValues(uint256 ) view returns(uint256)
func (_Provider *ProviderCallerSession) LpValues(arg0 *big.Int) (*big.Int, error) {
	return _Provider.Contract.LpValues(&_Provider.CallOpts, arg0)
}

// MasterChefV3 is a free data retrieval call binding the contract method 0xcb65b5b5.
//
// Solidity: function masterChefV3() view returns(address)
func (_Provider *ProviderCaller) MasterChefV3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "masterChefV3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MasterChefV3 is a free data retrieval call binding the contract method 0xcb65b5b5.
//
// Solidity: function masterChefV3() view returns(address)
func (_Provider *ProviderSession) MasterChefV3() (common.Address, error) {
	return _Provider.Contract.MasterChefV3(&_Provider.CallOpts)
}

// MasterChefV3 is a free data retrieval call binding the contract method 0xcb65b5b5.
//
// Solidity: function masterChefV3() view returns(address)
func (_Provider *ProviderCallerSession) MasterChefV3() (common.Address, error) {
	return _Provider.Contract.MasterChefV3(&_Provider.CallOpts)
}

// MaxLpPerUser is a free data retrieval call binding the contract method 0xf1e0fde9.
//
// Solidity: function maxLpPerUser() view returns(uint256)
func (_Provider *ProviderCaller) MaxLpPerUser(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "maxLpPerUser")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLpPerUser is a free data retrieval call binding the contract method 0xf1e0fde9.
//
// Solidity: function maxLpPerUser() view returns(uint256)
func (_Provider *ProviderSession) MaxLpPerUser() (*big.Int, error) {
	return _Provider.Contract.MaxLpPerUser(&_Provider.CallOpts)
}

// MaxLpPerUser is a free data retrieval call binding the contract method 0xf1e0fde9.
//
// Solidity: function maxLpPerUser() view returns(uint256)
func (_Provider *ProviderCallerSession) MaxLpPerUser() (*big.Int, error) {
	return _Provider.Contract.MaxLpPerUser(&_Provider.CallOpts)
}

// MinLpValue is a free data retrieval call binding the contract method 0x9a252197.
//
// Solidity: function minLpValue() view returns(uint256)
func (_Provider *ProviderCaller) MinLpValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "minLpValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinLpValue is a free data retrieval call binding the contract method 0x9a252197.
//
// Solidity: function minLpValue() view returns(uint256)
func (_Provider *ProviderSession) MinLpValue() (*big.Int, error) {
	return _Provider.Contract.MinLpValue(&_Provider.CallOpts)
}

// MinLpValue is a free data retrieval call binding the contract method 0x9a252197.
//
// Solidity: function minLpValue() view returns(uint256)
func (_Provider *ProviderCallerSession) MinLpValue() (*big.Int, error) {
	return _Provider.Contract.MinLpValue(&_Provider.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Provider *ProviderCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Provider *ProviderSession) Name() (string, error) {
	return _Provider.Contract.Name(&_Provider.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Provider *ProviderCallerSession) Name() (string, error) {
	return _Provider.Contract.Name(&_Provider.CallOpts)
}

// NonFungiblePositionManager is a free data retrieval call binding the contract method 0x26dd2c2a.
//
// Solidity: function nonFungiblePositionManager() view returns(address)
func (_Provider *ProviderCaller) NonFungiblePositionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "nonFungiblePositionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonFungiblePositionManager is a free data retrieval call binding the contract method 0x26dd2c2a.
//
// Solidity: function nonFungiblePositionManager() view returns(address)
func (_Provider *ProviderSession) NonFungiblePositionManager() (common.Address, error) {
	return _Provider.Contract.NonFungiblePositionManager(&_Provider.CallOpts)
}

// NonFungiblePositionManager is a free data retrieval call binding the contract method 0x26dd2c2a.
//
// Solidity: function nonFungiblePositionManager() view returns(address)
func (_Provider *ProviderCallerSession) NonFungiblePositionManager() (common.Address, error) {
	return _Provider.Contract.NonFungiblePositionManager(&_Provider.CallOpts)
}

// PancakeLpStakingVault is a free data retrieval call binding the contract method 0xa1743125.
//
// Solidity: function pancakeLpStakingVault() view returns(address)
func (_Provider *ProviderCaller) PancakeLpStakingVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "pancakeLpStakingVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeLpStakingVault is a free data retrieval call binding the contract method 0xa1743125.
//
// Solidity: function pancakeLpStakingVault() view returns(address)
func (_Provider *ProviderSession) PancakeLpStakingVault() (common.Address, error) {
	return _Provider.Contract.PancakeLpStakingVault(&_Provider.CallOpts)
}

// PancakeLpStakingVault is a free data retrieval call binding the contract method 0xa1743125.
//
// Solidity: function pancakeLpStakingVault() view returns(address)
func (_Provider *ProviderCallerSession) PancakeLpStakingVault() (common.Address, error) {
	return _Provider.Contract.PancakeLpStakingVault(&_Provider.CallOpts)
}

// PancakeStakingHub is a free data retrieval call binding the contract method 0xc5223561.
//
// Solidity: function pancakeStakingHub() view returns(address)
func (_Provider *ProviderCaller) PancakeStakingHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "pancakeStakingHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PancakeStakingHub is a free data retrieval call binding the contract method 0xc5223561.
//
// Solidity: function pancakeStakingHub() view returns(address)
func (_Provider *ProviderSession) PancakeStakingHub() (common.Address, error) {
	return _Provider.Contract.PancakeStakingHub(&_Provider.CallOpts)
}

// PancakeStakingHub is a free data retrieval call binding the contract method 0xc5223561.
//
// Solidity: function pancakeStakingHub() view returns(address)
func (_Provider *ProviderCallerSession) PancakeStakingHub() (common.Address, error) {
	return _Provider.Contract.PancakeStakingHub(&_Provider.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Provider *ProviderCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Provider *ProviderSession) Paused() (bool, error) {
	return _Provider.Contract.Paused(&_Provider.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Provider *ProviderCallerSession) Paused() (bool, error) {
	return _Provider.Contract.Paused(&_Provider.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() pure returns(bytes32, bool)
func (_Provider *ProviderCaller) Peek(opts *bind.CallOpts) ([32]byte, bool, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "peek")

	if err != nil {
		return *new([32]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() pure returns(bytes32, bool)
func (_Provider *ProviderSession) Peek() ([32]byte, bool, error) {
	return _Provider.Contract.Peek(&_Provider.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() pure returns(bytes32, bool)
func (_Provider *ProviderCallerSession) Peek() ([32]byte, bool, error) {
	return _Provider.Contract.Peek(&_Provider.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Provider *ProviderCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Provider *ProviderSession) ProxiableUUID() ([32]byte, error) {
	return _Provider.Contract.ProxiableUUID(&_Provider.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Provider *ProviderCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Provider.Contract.ProxiableUUID(&_Provider.CallOpts)
}

// ResilientOracle is a free data retrieval call binding the contract method 0x60de3604.
//
// Solidity: function resilientOracle() view returns(address)
func (_Provider *ProviderCaller) ResilientOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "resilientOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResilientOracle is a free data retrieval call binding the contract method 0x60de3604.
//
// Solidity: function resilientOracle() view returns(address)
func (_Provider *ProviderSession) ResilientOracle() (common.Address, error) {
	return _Provider.Contract.ResilientOracle(&_Provider.CallOpts)
}

// ResilientOracle is a free data retrieval call binding the contract method 0x60de3604.
//
// Solidity: function resilientOracle() view returns(address)
func (_Provider *ProviderCallerSession) ResilientOracle() (common.Address, error) {
	return _Provider.Contract.ResilientOracle(&_Provider.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Provider *ProviderCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Provider *ProviderSession) RewardToken() (common.Address, error) {
	return _Provider.Contract.RewardToken(&_Provider.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Provider *ProviderCallerSession) RewardToken() (common.Address, error) {
	return _Provider.Contract.RewardToken(&_Provider.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Provider *ProviderCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Provider *ProviderSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Provider.Contract.SupportsInterface(&_Provider.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Provider *ProviderCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Provider.Contract.SupportsInterface(&_Provider.CallOpts, interfaceId)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Provider *ProviderCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Provider *ProviderSession) Token0() (common.Address, error) {
	return _Provider.Contract.Token0(&_Provider.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Provider *ProviderCallerSession) Token0() (common.Address, error) {
	return _Provider.Contract.Token0(&_Provider.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Provider *ProviderCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Provider *ProviderSession) Token1() (common.Address, error) {
	return _Provider.Contract.Token1(&_Provider.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Provider *ProviderCallerSession) Token1() (common.Address, error) {
	return _Provider.Contract.Token1(&_Provider.CallOpts)
}

// UserLiquidations is a free data retrieval call binding the contract method 0x0fc7b53f.
//
// Solidity: function userLiquidations(address ) view returns(bool ongoing, uint256 token0Left, uint256 token1Left)
func (_Provider *ProviderCaller) UserLiquidations(opts *bind.CallOpts, arg0 common.Address) (struct {
	Ongoing    bool
	Token0Left *big.Int
	Token1Left *big.Int
}, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "userLiquidations", arg0)

	outstruct := new(struct {
		Ongoing    bool
		Token0Left *big.Int
		Token1Left *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ongoing = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Token0Left = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Token1Left = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserLiquidations is a free data retrieval call binding the contract method 0x0fc7b53f.
//
// Solidity: function userLiquidations(address ) view returns(bool ongoing, uint256 token0Left, uint256 token1Left)
func (_Provider *ProviderSession) UserLiquidations(arg0 common.Address) (struct {
	Ongoing    bool
	Token0Left *big.Int
	Token1Left *big.Int
}, error) {
	return _Provider.Contract.UserLiquidations(&_Provider.CallOpts, arg0)
}

// UserLiquidations is a free data retrieval call binding the contract method 0x0fc7b53f.
//
// Solidity: function userLiquidations(address ) view returns(bool ongoing, uint256 token0Left, uint256 token1Left)
func (_Provider *ProviderCallerSession) UserLiquidations(arg0 common.Address) (struct {
	Ongoing    bool
	Token0Left *big.Int
	Token1Left *big.Int
}, error) {
	return _Provider.Contract.UserLiquidations(&_Provider.CallOpts, arg0)
}

// UserLps is a free data retrieval call binding the contract method 0xfb45b39b.
//
// Solidity: function userLps(address , uint256 ) view returns(uint256)
func (_Provider *ProviderCaller) UserLps(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "userLps", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLps is a free data retrieval call binding the contract method 0xfb45b39b.
//
// Solidity: function userLps(address , uint256 ) view returns(uint256)
func (_Provider *ProviderSession) UserLps(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Provider.Contract.UserLps(&_Provider.CallOpts, arg0, arg1)
}

// UserLps is a free data retrieval call binding the contract method 0xfb45b39b.
//
// Solidity: function userLps(address , uint256 ) view returns(uint256)
func (_Provider *ProviderCallerSession) UserLps(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Provider.Contract.UserLps(&_Provider.CallOpts, arg0, arg1)
}

// UserTotalLpValue is a free data retrieval call binding the contract method 0x935b9463.
//
// Solidity: function userTotalLpValue(address ) view returns(uint256)
func (_Provider *ProviderCaller) UserTotalLpValue(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Provider.contract.Call(opts, &out, "userTotalLpValue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTotalLpValue is a free data retrieval call binding the contract method 0x935b9463.
//
// Solidity: function userTotalLpValue(address ) view returns(uint256)
func (_Provider *ProviderSession) UserTotalLpValue(arg0 common.Address) (*big.Int, error) {
	return _Provider.Contract.UserTotalLpValue(&_Provider.CallOpts, arg0)
}

// UserTotalLpValue is a free data retrieval call binding the contract method 0x935b9463.
//
// Solidity: function userTotalLpValue(address ) view returns(uint256)
func (_Provider *ProviderCallerSession) UserTotalLpValue(arg0 common.Address) (*big.Int, error) {
	return _Provider.Contract.UserTotalLpValue(&_Provider.CallOpts, arg0)
}

// BatchSyncUserLpValues is a paid mutator transaction binding the contract method 0x2181728d.
//
// Solidity: function batchSyncUserLpValues(address[] users) returns()
func (_Provider *ProviderTransactor) BatchSyncUserLpValues(opts *bind.TransactOpts, users []common.Address) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "batchSyncUserLpValues", users)
}

// BatchSyncUserLpValues is a paid mutator transaction binding the contract method 0x2181728d.
//
// Solidity: function batchSyncUserLpValues(address[] users) returns()
func (_Provider *ProviderSession) BatchSyncUserLpValues(users []common.Address) (*types.Transaction, error) {
	return _Provider.Contract.BatchSyncUserLpValues(&_Provider.TransactOpts, users)
}

// BatchSyncUserLpValues is a paid mutator transaction binding the contract method 0x2181728d.
//
// Solidity: function batchSyncUserLpValues(address[] users) returns()
func (_Provider *ProviderTransactorSession) BatchSyncUserLpValues(users []common.Address) (*types.Transaction, error) {
	return _Provider.Contract.BatchSyncUserLpValues(&_Provider.TransactOpts, users)
}

// DaoBurn is a paid mutator transaction binding the contract method 0xdf3d23ae.
//
// Solidity: function daoBurn(address user, uint256 lpAmount) returns()
func (_Provider *ProviderTransactor) DaoBurn(opts *bind.TransactOpts, user common.Address, lpAmount *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "daoBurn", user, lpAmount)
}

// DaoBurn is a paid mutator transaction binding the contract method 0xdf3d23ae.
//
// Solidity: function daoBurn(address user, uint256 lpAmount) returns()
func (_Provider *ProviderSession) DaoBurn(user common.Address, lpAmount *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.DaoBurn(&_Provider.TransactOpts, user, lpAmount)
}

// DaoBurn is a paid mutator transaction binding the contract method 0xdf3d23ae.
//
// Solidity: function daoBurn(address user, uint256 lpAmount) returns()
func (_Provider *ProviderTransactorSession) DaoBurn(user common.Address, lpAmount *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.DaoBurn(&_Provider.TransactOpts, user, lpAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Provider *ProviderTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Provider *ProviderSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.Contract.GrantRole(&_Provider.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Provider *ProviderTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.Contract.GrantRole(&_Provider.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x0974f94c.
//
// Solidity: function initialize(address _admin, address _manager, address _bot, address _pauser, address _pancakeStakingHub, address _pancakeStakingVault, address _resilientOracle, uint256 _maxLpPerUser, uint256 _minLpValue, uint256 _lpDiscountRate) returns()
func (_Provider *ProviderTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _manager common.Address, _bot common.Address, _pauser common.Address, _pancakeStakingHub common.Address, _pancakeStakingVault common.Address, _resilientOracle common.Address, _maxLpPerUser *big.Int, _minLpValue *big.Int, _lpDiscountRate *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "initialize", _admin, _manager, _bot, _pauser, _pancakeStakingHub, _pancakeStakingVault, _resilientOracle, _maxLpPerUser, _minLpValue, _lpDiscountRate)
}

// Initialize is a paid mutator transaction binding the contract method 0x0974f94c.
//
// Solidity: function initialize(address _admin, address _manager, address _bot, address _pauser, address _pancakeStakingHub, address _pancakeStakingVault, address _resilientOracle, uint256 _maxLpPerUser, uint256 _minLpValue, uint256 _lpDiscountRate) returns()
func (_Provider *ProviderSession) Initialize(_admin common.Address, _manager common.Address, _bot common.Address, _pauser common.Address, _pancakeStakingHub common.Address, _pancakeStakingVault common.Address, _resilientOracle common.Address, _maxLpPerUser *big.Int, _minLpValue *big.Int, _lpDiscountRate *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.Initialize(&_Provider.TransactOpts, _admin, _manager, _bot, _pauser, _pancakeStakingHub, _pancakeStakingVault, _resilientOracle, _maxLpPerUser, _minLpValue, _lpDiscountRate)
}

// Initialize is a paid mutator transaction binding the contract method 0x0974f94c.
//
// Solidity: function initialize(address _admin, address _manager, address _bot, address _pauser, address _pancakeStakingHub, address _pancakeStakingVault, address _resilientOracle, uint256 _maxLpPerUser, uint256 _minLpValue, uint256 _lpDiscountRate) returns()
func (_Provider *ProviderTransactorSession) Initialize(_admin common.Address, _manager common.Address, _bot common.Address, _pauser common.Address, _pancakeStakingHub common.Address, _pancakeStakingVault common.Address, _resilientOracle common.Address, _maxLpPerUser *big.Int, _minLpValue *big.Int, _lpDiscountRate *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.Initialize(&_Provider.TransactOpts, _admin, _manager, _bot, _pauser, _pancakeStakingHub, _pancakeStakingVault, _resilientOracle, _maxLpPerUser, _minLpValue, _lpDiscountRate)
}

// Liquidation is a paid mutator transaction binding the contract method 0x01852496.
//
// Solidity: function liquidation(address owner, address recipient, uint256 amount, bytes data, bool isLeftOver) returns()
func (_Provider *ProviderTransactor) Liquidation(opts *bind.TransactOpts, owner common.Address, recipient common.Address, amount *big.Int, data []byte, isLeftOver bool) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "liquidation", owner, recipient, amount, data, isLeftOver)
}

// Liquidation is a paid mutator transaction binding the contract method 0x01852496.
//
// Solidity: function liquidation(address owner, address recipient, uint256 amount, bytes data, bool isLeftOver) returns()
func (_Provider *ProviderSession) Liquidation(owner common.Address, recipient common.Address, amount *big.Int, data []byte, isLeftOver bool) (*types.Transaction, error) {
	return _Provider.Contract.Liquidation(&_Provider.TransactOpts, owner, recipient, amount, data, isLeftOver)
}

// Liquidation is a paid mutator transaction binding the contract method 0x01852496.
//
// Solidity: function liquidation(address owner, address recipient, uint256 amount, bytes data, bool isLeftOver) returns()
func (_Provider *ProviderTransactorSession) Liquidation(owner common.Address, recipient common.Address, amount *big.Int, data []byte, isLeftOver bool) (*types.Transaction, error) {
	return _Provider.Contract.Liquidation(&_Provider.TransactOpts, owner, recipient, amount, data, isLeftOver)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes ) returns(bytes4)
func (_Provider *ProviderTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, from common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "onERC721Received", arg0, from, tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes ) returns(bytes4)
func (_Provider *ProviderSession) OnERC721Received(arg0 common.Address, from common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Provider.Contract.OnERC721Received(&_Provider.TransactOpts, arg0, from, tokenId, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address from, uint256 tokenId, bytes ) returns(bytes4)
func (_Provider *ProviderTransactorSession) OnERC721Received(arg0 common.Address, from common.Address, tokenId *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Provider.Contract.OnERC721Received(&_Provider.TransactOpts, arg0, from, tokenId, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Provider *ProviderTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Provider *ProviderSession) Pause() (*types.Transaction, error) {
	return _Provider.Contract.Pause(&_Provider.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Provider *ProviderTransactorSession) Pause() (*types.Transaction, error) {
	return _Provider.Contract.Pause(&_Provider.TransactOpts)
}

// Provide is a paid mutator transaction binding the contract method 0x2e2ebe06.
//
// Solidity: function provide(uint256 tokenId) returns()
func (_Provider *ProviderTransactor) Provide(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "provide", tokenId)
}

// Provide is a paid mutator transaction binding the contract method 0x2e2ebe06.
//
// Solidity: function provide(uint256 tokenId) returns()
func (_Provider *ProviderSession) Provide(tokenId *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.Provide(&_Provider.TransactOpts, tokenId)
}

// Provide is a paid mutator transaction binding the contract method 0x2e2ebe06.
//
// Solidity: function provide(uint256 tokenId) returns()
func (_Provider *ProviderTransactorSession) Provide(tokenId *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.Provide(&_Provider.TransactOpts, tokenId)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(uint256 tokenId) returns()
func (_Provider *ProviderTransactor) Release(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "release", tokenId)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(uint256 tokenId) returns()
func (_Provider *ProviderSession) Release(tokenId *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.Release(&_Provider.TransactOpts, tokenId)
}

// Release is a paid mutator transaction binding the contract method 0x37bdc99b.
//
// Solidity: function release(uint256 tokenId) returns()
func (_Provider *ProviderTransactorSession) Release(tokenId *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.Release(&_Provider.TransactOpts, tokenId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Provider *ProviderTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Provider *ProviderSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.Contract.RenounceRole(&_Provider.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Provider *ProviderTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.Contract.RenounceRole(&_Provider.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Provider *ProviderTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Provider *ProviderSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.Contract.RevokeRole(&_Provider.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Provider *ProviderTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Provider.Contract.RevokeRole(&_Provider.TransactOpts, role, account)
}

// SetLpDiscountRate is a paid mutator transaction binding the contract method 0x9bec8603.
//
// Solidity: function setLpDiscountRate(uint256 _lpDiscountRate) returns()
func (_Provider *ProviderTransactor) SetLpDiscountRate(opts *bind.TransactOpts, _lpDiscountRate *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "setLpDiscountRate", _lpDiscountRate)
}

// SetLpDiscountRate is a paid mutator transaction binding the contract method 0x9bec8603.
//
// Solidity: function setLpDiscountRate(uint256 _lpDiscountRate) returns()
func (_Provider *ProviderSession) SetLpDiscountRate(_lpDiscountRate *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.SetLpDiscountRate(&_Provider.TransactOpts, _lpDiscountRate)
}

// SetLpDiscountRate is a paid mutator transaction binding the contract method 0x9bec8603.
//
// Solidity: function setLpDiscountRate(uint256 _lpDiscountRate) returns()
func (_Provider *ProviderTransactorSession) SetLpDiscountRate(_lpDiscountRate *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.SetLpDiscountRate(&_Provider.TransactOpts, _lpDiscountRate)
}

// SetMaxLpPerUser is a paid mutator transaction binding the contract method 0x61e289bc.
//
// Solidity: function setMaxLpPerUser(uint256 _maxLpPerUser) returns()
func (_Provider *ProviderTransactor) SetMaxLpPerUser(opts *bind.TransactOpts, _maxLpPerUser *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "setMaxLpPerUser", _maxLpPerUser)
}

// SetMaxLpPerUser is a paid mutator transaction binding the contract method 0x61e289bc.
//
// Solidity: function setMaxLpPerUser(uint256 _maxLpPerUser) returns()
func (_Provider *ProviderSession) SetMaxLpPerUser(_maxLpPerUser *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.SetMaxLpPerUser(&_Provider.TransactOpts, _maxLpPerUser)
}

// SetMaxLpPerUser is a paid mutator transaction binding the contract method 0x61e289bc.
//
// Solidity: function setMaxLpPerUser(uint256 _maxLpPerUser) returns()
func (_Provider *ProviderTransactorSession) SetMaxLpPerUser(_maxLpPerUser *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.SetMaxLpPerUser(&_Provider.TransactOpts, _maxLpPerUser)
}

// SetMinLpValue is a paid mutator transaction binding the contract method 0x2f57c52a.
//
// Solidity: function setMinLpValue(uint256 _minLpValue) returns()
func (_Provider *ProviderTransactor) SetMinLpValue(opts *bind.TransactOpts, _minLpValue *big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "setMinLpValue", _minLpValue)
}

// SetMinLpValue is a paid mutator transaction binding the contract method 0x2f57c52a.
//
// Solidity: function setMinLpValue(uint256 _minLpValue) returns()
func (_Provider *ProviderSession) SetMinLpValue(_minLpValue *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.SetMinLpValue(&_Provider.TransactOpts, _minLpValue)
}

// SetMinLpValue is a paid mutator transaction binding the contract method 0x2f57c52a.
//
// Solidity: function setMinLpValue(uint256 _minLpValue) returns()
func (_Provider *ProviderTransactorSession) SetMinLpValue(_minLpValue *big.Int) (*types.Transaction, error) {
	return _Provider.Contract.SetMinLpValue(&_Provider.TransactOpts, _minLpValue)
}

// SyncUserLpValues is a paid mutator transaction binding the contract method 0xdd66f26d.
//
// Solidity: function syncUserLpValues(address user) returns()
func (_Provider *ProviderTransactor) SyncUserLpValues(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "syncUserLpValues", user)
}

// SyncUserLpValues is a paid mutator transaction binding the contract method 0xdd66f26d.
//
// Solidity: function syncUserLpValues(address user) returns()
func (_Provider *ProviderSession) SyncUserLpValues(user common.Address) (*types.Transaction, error) {
	return _Provider.Contract.SyncUserLpValues(&_Provider.TransactOpts, user)
}

// SyncUserLpValues is a paid mutator transaction binding the contract method 0xdd66f26d.
//
// Solidity: function syncUserLpValues(address user) returns()
func (_Provider *ProviderTransactorSession) SyncUserLpValues(user common.Address) (*types.Transaction, error) {
	return _Provider.Contract.SyncUserLpValues(&_Provider.TransactOpts, user)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Provider *ProviderTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Provider *ProviderSession) Unpause() (*types.Transaction, error) {
	return _Provider.Contract.Unpause(&_Provider.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Provider *ProviderTransactorSession) Unpause() (*types.Transaction, error) {
	return _Provider.Contract.Unpause(&_Provider.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Provider *ProviderTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Provider *ProviderSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Provider.Contract.UpgradeTo(&_Provider.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Provider *ProviderTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Provider.Contract.UpgradeTo(&_Provider.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Provider *ProviderTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Provider *ProviderSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Provider.Contract.UpgradeToAndCall(&_Provider.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Provider *ProviderTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Provider.Contract.UpgradeToAndCall(&_Provider.TransactOpts, newImplementation, data)
}

// VaultClaimStakingReward is a paid mutator transaction binding the contract method 0xb87598bc.
//
// Solidity: function vaultClaimStakingReward(address account, uint256[] tokenIds) returns(uint256 totalReward)
func (_Provider *ProviderTransactor) VaultClaimStakingReward(opts *bind.TransactOpts, account common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Provider.contract.Transact(opts, "vaultClaimStakingReward", account, tokenIds)
}

// VaultClaimStakingReward is a paid mutator transaction binding the contract method 0xb87598bc.
//
// Solidity: function vaultClaimStakingReward(address account, uint256[] tokenIds) returns(uint256 totalReward)
func (_Provider *ProviderSession) VaultClaimStakingReward(account common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Provider.Contract.VaultClaimStakingReward(&_Provider.TransactOpts, account, tokenIds)
}

// VaultClaimStakingReward is a paid mutator transaction binding the contract method 0xb87598bc.
//
// Solidity: function vaultClaimStakingReward(address account, uint256[] tokenIds) returns(uint256 totalReward)
func (_Provider *ProviderTransactorSession) VaultClaimStakingReward(account common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Provider.Contract.VaultClaimStakingReward(&_Provider.TransactOpts, account, tokenIds)
}

// ProviderAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Provider contract.
type ProviderAdminChangedIterator struct {
	Event *ProviderAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProviderAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderAdminChanged)
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
		it.Event = new(ProviderAdminChanged)
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
func (it *ProviderAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderAdminChanged represents a AdminChanged event raised by the Provider contract.
type ProviderAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Provider *ProviderFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ProviderAdminChangedIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ProviderAdminChangedIterator{contract: _Provider.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Provider *ProviderFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ProviderAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderAdminChanged)
				if err := _Provider.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Provider *ProviderFilterer) ParseAdminChanged(log types.Log) (*ProviderAdminChanged, error) {
	event := new(ProviderAdminChanged)
	if err := _Provider.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Provider contract.
type ProviderBeaconUpgradedIterator struct {
	Event *ProviderBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ProviderBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderBeaconUpgraded)
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
		it.Event = new(ProviderBeaconUpgraded)
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
func (it *ProviderBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderBeaconUpgraded represents a BeaconUpgraded event raised by the Provider contract.
type ProviderBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Provider *ProviderFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ProviderBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Provider.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ProviderBeaconUpgradedIterator{contract: _Provider.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Provider *ProviderFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ProviderBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Provider.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderBeaconUpgraded)
				if err := _Provider.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Provider *ProviderFilterer) ParseBeaconUpgraded(log types.Log) (*ProviderBeaconUpgraded, error) {
	event := new(ProviderBeaconUpgraded)
	if err := _Provider.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderDepositLpIterator is returned from FilterDepositLp and is used to iterate over the raw logs and unpacked data for DepositLp events raised by the Provider contract.
type ProviderDepositLpIterator struct {
	Event *ProviderDepositLp // Event containing the contract specifics and raw log

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
func (it *ProviderDepositLpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderDepositLp)
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
		it.Event = new(ProviderDepositLp)
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
func (it *ProviderDepositLpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderDepositLpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderDepositLp represents a DepositLp event raised by the Provider contract.
type ProviderDepositLp struct {
	User    common.Address
	TokenId *big.Int
	LpValue *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositLp is a free log retrieval operation binding the contract event 0xa83bffa82b1abbe9b0782ac5e7b9bb97c1b83698b5c874955af5541d9b9a6791.
//
// Solidity: event DepositLp(address user, uint256 tokenId, uint256 lpValue)
func (_Provider *ProviderFilterer) FilterDepositLp(opts *bind.FilterOpts) (*ProviderDepositLpIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "DepositLp")
	if err != nil {
		return nil, err
	}
	return &ProviderDepositLpIterator{contract: _Provider.contract, event: "DepositLp", logs: logs, sub: sub}, nil
}

// WatchDepositLp is a free log subscription operation binding the contract event 0xa83bffa82b1abbe9b0782ac5e7b9bb97c1b83698b5c874955af5541d9b9a6791.
//
// Solidity: event DepositLp(address user, uint256 tokenId, uint256 lpValue)
func (_Provider *ProviderFilterer) WatchDepositLp(opts *bind.WatchOpts, sink chan<- *ProviderDepositLp) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "DepositLp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderDepositLp)
				if err := _Provider.contract.UnpackLog(event, "DepositLp", log); err != nil {
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

// ParseDepositLp is a log parse operation binding the contract event 0xa83bffa82b1abbe9b0782ac5e7b9bb97c1b83698b5c874955af5541d9b9a6791.
//
// Solidity: event DepositLp(address user, uint256 tokenId, uint256 lpValue)
func (_Provider *ProviderFilterer) ParseDepositLp(log types.Log) (*ProviderDepositLp, error) {
	event := new(ProviderDepositLp)
	if err := _Provider.contract.UnpackLog(event, "DepositLp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Provider contract.
type ProviderInitializedIterator struct {
	Event *ProviderInitialized // Event containing the contract specifics and raw log

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
func (it *ProviderInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderInitialized)
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
		it.Event = new(ProviderInitialized)
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
func (it *ProviderInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderInitialized represents a Initialized event raised by the Provider contract.
type ProviderInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Provider *ProviderFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProviderInitializedIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProviderInitializedIterator{contract: _Provider.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Provider *ProviderFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProviderInitialized) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderInitialized)
				if err := _Provider.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Provider *ProviderFilterer) ParseInitialized(log types.Log) (*ProviderInitialized, error) {
	event := new(ProviderInitialized)
	if err := _Provider.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderLiquidatableIterator is returned from FilterLiquidatable and is used to iterate over the raw logs and unpacked data for Liquidatable events raised by the Provider contract.
type ProviderLiquidatableIterator struct {
	Event *ProviderLiquidatable // Event containing the contract specifics and raw log

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
func (it *ProviderLiquidatableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderLiquidatable)
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
		it.Event = new(ProviderLiquidatable)
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
func (it *ProviderLiquidatableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderLiquidatableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderLiquidatable represents a Liquidatable event raised by the Provider contract.
type ProviderLiquidatable struct {
	User                 common.Address
	UserLpTotalValue     *big.Int
	UserLockedCollateral *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterLiquidatable is a free log retrieval operation binding the contract event 0x6ca8decc94bd7d3371cd4e0e10fc32b3500946b13910df760b556ffd1fc4ad2e.
//
// Solidity: event Liquidatable(address user, uint256 userLpTotalValue, uint256 userLockedCollateral)
func (_Provider *ProviderFilterer) FilterLiquidatable(opts *bind.FilterOpts) (*ProviderLiquidatableIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "Liquidatable")
	if err != nil {
		return nil, err
	}
	return &ProviderLiquidatableIterator{contract: _Provider.contract, event: "Liquidatable", logs: logs, sub: sub}, nil
}

// WatchLiquidatable is a free log subscription operation binding the contract event 0x6ca8decc94bd7d3371cd4e0e10fc32b3500946b13910df760b556ffd1fc4ad2e.
//
// Solidity: event Liquidatable(address user, uint256 userLpTotalValue, uint256 userLockedCollateral)
func (_Provider *ProviderFilterer) WatchLiquidatable(opts *bind.WatchOpts, sink chan<- *ProviderLiquidatable) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "Liquidatable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderLiquidatable)
				if err := _Provider.contract.UnpackLog(event, "Liquidatable", log); err != nil {
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

// ParseLiquidatable is a log parse operation binding the contract event 0x6ca8decc94bd7d3371cd4e0e10fc32b3500946b13910df760b556ffd1fc4ad2e.
//
// Solidity: event Liquidatable(address user, uint256 userLpTotalValue, uint256 userLockedCollateral)
func (_Provider *ProviderFilterer) ParseLiquidatable(log types.Log) (*ProviderLiquidatable, error) {
	event := new(ProviderLiquidatable)
	if err := _Provider.contract.UnpackLog(event, "Liquidatable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderLiquidatedIterator is returned from FilterLiquidated and is used to iterate over the raw logs and unpacked data for Liquidated events raised by the Provider contract.
type ProviderLiquidatedIterator struct {
	Event *ProviderLiquidated // Event containing the contract specifics and raw log

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
func (it *ProviderLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderLiquidated)
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
		it.Event = new(ProviderLiquidated)
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
func (it *ProviderLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderLiquidated represents a Liquidated event raised by the Provider contract.
type ProviderLiquidated struct {
	User       common.Address
	Receipient common.Address
	Amount     *big.Int
	IsLeftover bool
	Token0Left *big.Int
	Token1Left *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidated is a free log retrieval operation binding the contract event 0xb315944923988cd53fe5caa5e109f7274a7e67bce0040d4761f85d2fda97ac9e.
//
// Solidity: event Liquidated(address user, address receipient, uint256 amount, bool isLeftover, uint256 token0Left, uint256 token1Left)
func (_Provider *ProviderFilterer) FilterLiquidated(opts *bind.FilterOpts) (*ProviderLiquidatedIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "Liquidated")
	if err != nil {
		return nil, err
	}
	return &ProviderLiquidatedIterator{contract: _Provider.contract, event: "Liquidated", logs: logs, sub: sub}, nil
}

// WatchLiquidated is a free log subscription operation binding the contract event 0xb315944923988cd53fe5caa5e109f7274a7e67bce0040d4761f85d2fda97ac9e.
//
// Solidity: event Liquidated(address user, address receipient, uint256 amount, bool isLeftover, uint256 token0Left, uint256 token1Left)
func (_Provider *ProviderFilterer) WatchLiquidated(opts *bind.WatchOpts, sink chan<- *ProviderLiquidated) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "Liquidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderLiquidated)
				if err := _Provider.contract.UnpackLog(event, "Liquidated", log); err != nil {
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

// ParseLiquidated is a log parse operation binding the contract event 0xb315944923988cd53fe5caa5e109f7274a7e67bce0040d4761f85d2fda97ac9e.
//
// Solidity: event Liquidated(address user, address receipient, uint256 amount, bool isLeftover, uint256 token0Left, uint256 token1Left)
func (_Provider *ProviderFilterer) ParseLiquidated(log types.Log) (*ProviderLiquidated, error) {
	event := new(ProviderLiquidated)
	if err := _Provider.contract.UnpackLog(event, "Liquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderLiquidationBeganIterator is returned from FilterLiquidationBegan and is used to iterate over the raw logs and unpacked data for LiquidationBegan events raised by the Provider contract.
type ProviderLiquidationBeganIterator struct {
	Event *ProviderLiquidationBegan // Event containing the contract specifics and raw log

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
func (it *ProviderLiquidationBeganIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderLiquidationBegan)
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
		it.Event = new(ProviderLiquidationBegan)
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
func (it *ProviderLiquidationBeganIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderLiquidationBeganIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderLiquidationBegan represents a LiquidationBegan event raised by the Provider contract.
type ProviderLiquidationBegan struct {
	User               common.Address
	LiquidatableAmount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterLiquidationBegan is a free log retrieval operation binding the contract event 0x1f728ee44d95f87cfdfb2a43333ac3e0b804dcdddeae7fa4983950650612e83c.
//
// Solidity: event LiquidationBegan(address user, uint256 liquidatableAmount)
func (_Provider *ProviderFilterer) FilterLiquidationBegan(opts *bind.FilterOpts) (*ProviderLiquidationBeganIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "LiquidationBegan")
	if err != nil {
		return nil, err
	}
	return &ProviderLiquidationBeganIterator{contract: _Provider.contract, event: "LiquidationBegan", logs: logs, sub: sub}, nil
}

// WatchLiquidationBegan is a free log subscription operation binding the contract event 0x1f728ee44d95f87cfdfb2a43333ac3e0b804dcdddeae7fa4983950650612e83c.
//
// Solidity: event LiquidationBegan(address user, uint256 liquidatableAmount)
func (_Provider *ProviderFilterer) WatchLiquidationBegan(opts *bind.WatchOpts, sink chan<- *ProviderLiquidationBegan) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "LiquidationBegan")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderLiquidationBegan)
				if err := _Provider.contract.UnpackLog(event, "LiquidationBegan", log); err != nil {
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

// ParseLiquidationBegan is a log parse operation binding the contract event 0x1f728ee44d95f87cfdfb2a43333ac3e0b804dcdddeae7fa4983950650612e83c.
//
// Solidity: event LiquidationBegan(address user, uint256 liquidatableAmount)
func (_Provider *ProviderFilterer) ParseLiquidationBegan(log types.Log) (*ProviderLiquidationBegan, error) {
	event := new(ProviderLiquidationBegan)
	if err := _Provider.contract.UnpackLog(event, "LiquidationBegan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderLpDiscountRateSetIterator is returned from FilterLpDiscountRateSet and is used to iterate over the raw logs and unpacked data for LpDiscountRateSet events raised by the Provider contract.
type ProviderLpDiscountRateSetIterator struct {
	Event *ProviderLpDiscountRateSet // Event containing the contract specifics and raw log

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
func (it *ProviderLpDiscountRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderLpDiscountRateSet)
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
		it.Event = new(ProviderLpDiscountRateSet)
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
func (it *ProviderLpDiscountRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderLpDiscountRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderLpDiscountRateSet represents a LpDiscountRateSet event raised by the Provider contract.
type ProviderLpDiscountRateSet struct {
	OldExchangeRate *big.Int
	NewExchangeRate *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLpDiscountRateSet is a free log retrieval operation binding the contract event 0xf367959bbb2cfc7bb7229fad804e4c8daebca65a69e365cc87451628b55842d7.
//
// Solidity: event LpDiscountRateSet(uint256 oldExchangeRate, uint256 newExchangeRate)
func (_Provider *ProviderFilterer) FilterLpDiscountRateSet(opts *bind.FilterOpts) (*ProviderLpDiscountRateSetIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "LpDiscountRateSet")
	if err != nil {
		return nil, err
	}
	return &ProviderLpDiscountRateSetIterator{contract: _Provider.contract, event: "LpDiscountRateSet", logs: logs, sub: sub}, nil
}

// WatchLpDiscountRateSet is a free log subscription operation binding the contract event 0xf367959bbb2cfc7bb7229fad804e4c8daebca65a69e365cc87451628b55842d7.
//
// Solidity: event LpDiscountRateSet(uint256 oldExchangeRate, uint256 newExchangeRate)
func (_Provider *ProviderFilterer) WatchLpDiscountRateSet(opts *bind.WatchOpts, sink chan<- *ProviderLpDiscountRateSet) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "LpDiscountRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderLpDiscountRateSet)
				if err := _Provider.contract.UnpackLog(event, "LpDiscountRateSet", log); err != nil {
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

// ParseLpDiscountRateSet is a log parse operation binding the contract event 0xf367959bbb2cfc7bb7229fad804e4c8daebca65a69e365cc87451628b55842d7.
//
// Solidity: event LpDiscountRateSet(uint256 oldExchangeRate, uint256 newExchangeRate)
func (_Provider *ProviderFilterer) ParseLpDiscountRateSet(log types.Log) (*ProviderLpDiscountRateSet, error) {
	event := new(ProviderLpDiscountRateSet)
	if err := _Provider.contract.UnpackLog(event, "LpDiscountRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderMaxLpPerUserSetIterator is returned from FilterMaxLpPerUserSet and is used to iterate over the raw logs and unpacked data for MaxLpPerUserSet events raised by the Provider contract.
type ProviderMaxLpPerUserSetIterator struct {
	Event *ProviderMaxLpPerUserSet // Event containing the contract specifics and raw log

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
func (it *ProviderMaxLpPerUserSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderMaxLpPerUserSet)
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
		it.Event = new(ProviderMaxLpPerUserSet)
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
func (it *ProviderMaxLpPerUserSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderMaxLpPerUserSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderMaxLpPerUserSet represents a MaxLpPerUserSet event raised by the Provider contract.
type ProviderMaxLpPerUserSet struct {
	OldMaxLpValue *big.Int
	NewMaxLpValue *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaxLpPerUserSet is a free log retrieval operation binding the contract event 0xc3ef270bd7176fb3c3c896c693244536000240e9feea94fbf207bb381a9acb5e.
//
// Solidity: event MaxLpPerUserSet(uint256 oldMaxLpValue, uint256 newMaxLpValue)
func (_Provider *ProviderFilterer) FilterMaxLpPerUserSet(opts *bind.FilterOpts) (*ProviderMaxLpPerUserSetIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "MaxLpPerUserSet")
	if err != nil {
		return nil, err
	}
	return &ProviderMaxLpPerUserSetIterator{contract: _Provider.contract, event: "MaxLpPerUserSet", logs: logs, sub: sub}, nil
}

// WatchMaxLpPerUserSet is a free log subscription operation binding the contract event 0xc3ef270bd7176fb3c3c896c693244536000240e9feea94fbf207bb381a9acb5e.
//
// Solidity: event MaxLpPerUserSet(uint256 oldMaxLpValue, uint256 newMaxLpValue)
func (_Provider *ProviderFilterer) WatchMaxLpPerUserSet(opts *bind.WatchOpts, sink chan<- *ProviderMaxLpPerUserSet) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "MaxLpPerUserSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderMaxLpPerUserSet)
				if err := _Provider.contract.UnpackLog(event, "MaxLpPerUserSet", log); err != nil {
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

// ParseMaxLpPerUserSet is a log parse operation binding the contract event 0xc3ef270bd7176fb3c3c896c693244536000240e9feea94fbf207bb381a9acb5e.
//
// Solidity: event MaxLpPerUserSet(uint256 oldMaxLpValue, uint256 newMaxLpValue)
func (_Provider *ProviderFilterer) ParseMaxLpPerUserSet(log types.Log) (*ProviderMaxLpPerUserSet, error) {
	event := new(ProviderMaxLpPerUserSet)
	if err := _Provider.contract.UnpackLog(event, "MaxLpPerUserSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderMinLpValueSetIterator is returned from FilterMinLpValueSet and is used to iterate over the raw logs and unpacked data for MinLpValueSet events raised by the Provider contract.
type ProviderMinLpValueSetIterator struct {
	Event *ProviderMinLpValueSet // Event containing the contract specifics and raw log

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
func (it *ProviderMinLpValueSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderMinLpValueSet)
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
		it.Event = new(ProviderMinLpValueSet)
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
func (it *ProviderMinLpValueSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderMinLpValueSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderMinLpValueSet represents a MinLpValueSet event raised by the Provider contract.
type ProviderMinLpValueSet struct {
	OldMinLpValue *big.Int
	NewMinLpValue *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMinLpValueSet is a free log retrieval operation binding the contract event 0xd6d46991a0df72028a56ab4202e9cbf8ad74f4919cb27183cea44107c4a8a068.
//
// Solidity: event MinLpValueSet(uint256 oldMinLpValue, uint256 newMinLpValue)
func (_Provider *ProviderFilterer) FilterMinLpValueSet(opts *bind.FilterOpts) (*ProviderMinLpValueSetIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "MinLpValueSet")
	if err != nil {
		return nil, err
	}
	return &ProviderMinLpValueSetIterator{contract: _Provider.contract, event: "MinLpValueSet", logs: logs, sub: sub}, nil
}

// WatchMinLpValueSet is a free log subscription operation binding the contract event 0xd6d46991a0df72028a56ab4202e9cbf8ad74f4919cb27183cea44107c4a8a068.
//
// Solidity: event MinLpValueSet(uint256 oldMinLpValue, uint256 newMinLpValue)
func (_Provider *ProviderFilterer) WatchMinLpValueSet(opts *bind.WatchOpts, sink chan<- *ProviderMinLpValueSet) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "MinLpValueSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderMinLpValueSet)
				if err := _Provider.contract.UnpackLog(event, "MinLpValueSet", log); err != nil {
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

// ParseMinLpValueSet is a log parse operation binding the contract event 0xd6d46991a0df72028a56ab4202e9cbf8ad74f4919cb27183cea44107c4a8a068.
//
// Solidity: event MinLpValueSet(uint256 oldMinLpValue, uint256 newMinLpValue)
func (_Provider *ProviderFilterer) ParseMinLpValueSet(log types.Log) (*ProviderMinLpValueSet, error) {
	event := new(ProviderMinLpValueSet)
	if err := _Provider.contract.UnpackLog(event, "MinLpValueSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Provider contract.
type ProviderPausedIterator struct {
	Event *ProviderPaused // Event containing the contract specifics and raw log

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
func (it *ProviderPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderPaused)
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
		it.Event = new(ProviderPaused)
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
func (it *ProviderPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderPaused represents a Paused event raised by the Provider contract.
type ProviderPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Provider *ProviderFilterer) FilterPaused(opts *bind.FilterOpts) (*ProviderPausedIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ProviderPausedIterator{contract: _Provider.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Provider *ProviderFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ProviderPaused) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderPaused)
				if err := _Provider.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Provider *ProviderFilterer) ParsePaused(log types.Log) (*ProviderPaused, error) {
	event := new(ProviderPaused)
	if err := _Provider.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Provider contract.
type ProviderRoleAdminChangedIterator struct {
	Event *ProviderRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProviderRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRoleAdminChanged)
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
		it.Event = new(ProviderRoleAdminChanged)
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
func (it *ProviderRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRoleAdminChanged represents a RoleAdminChanged event raised by the Provider contract.
type ProviderRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Provider *ProviderFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ProviderRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Provider.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRoleAdminChangedIterator{contract: _Provider.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Provider *ProviderFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ProviderRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Provider.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRoleAdminChanged)
				if err := _Provider.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Provider *ProviderFilterer) ParseRoleAdminChanged(log types.Log) (*ProviderRoleAdminChanged, error) {
	event := new(ProviderRoleAdminChanged)
	if err := _Provider.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Provider contract.
type ProviderRoleGrantedIterator struct {
	Event *ProviderRoleGranted // Event containing the contract specifics and raw log

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
func (it *ProviderRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRoleGranted)
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
		it.Event = new(ProviderRoleGranted)
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
func (it *ProviderRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRoleGranted represents a RoleGranted event raised by the Provider contract.
type ProviderRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Provider *ProviderFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProviderRoleGrantedIterator, error) {

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

	logs, sub, err := _Provider.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRoleGrantedIterator{contract: _Provider.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Provider *ProviderFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ProviderRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Provider.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRoleGranted)
				if err := _Provider.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Provider *ProviderFilterer) ParseRoleGranted(log types.Log) (*ProviderRoleGranted, error) {
	event := new(ProviderRoleGranted)
	if err := _Provider.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Provider contract.
type ProviderRoleRevokedIterator struct {
	Event *ProviderRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ProviderRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRoleRevoked)
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
		it.Event = new(ProviderRoleRevoked)
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
func (it *ProviderRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRoleRevoked represents a RoleRevoked event raised by the Provider contract.
type ProviderRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Provider *ProviderFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProviderRoleRevokedIterator, error) {

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

	logs, sub, err := _Provider.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRoleRevokedIterator{contract: _Provider.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Provider *ProviderFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ProviderRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Provider.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRoleRevoked)
				if err := _Provider.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Provider *ProviderFilterer) ParseRoleRevoked(log types.Log) (*ProviderRoleRevoked, error) {
	event := new(ProviderRoleRevoked)
	if err := _Provider.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Provider contract.
type ProviderUnpausedIterator struct {
	Event *ProviderUnpaused // Event containing the contract specifics and raw log

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
func (it *ProviderUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderUnpaused)
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
		it.Event = new(ProviderUnpaused)
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
func (it *ProviderUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderUnpaused represents a Unpaused event raised by the Provider contract.
type ProviderUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Provider *ProviderFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ProviderUnpausedIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ProviderUnpausedIterator{contract: _Provider.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Provider *ProviderFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ProviderUnpaused) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderUnpaused)
				if err := _Provider.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Provider *ProviderFilterer) ParseUnpaused(log types.Log) (*ProviderUnpaused, error) {
	event := new(ProviderUnpaused)
	if err := _Provider.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Provider contract.
type ProviderUpgradedIterator struct {
	Event *ProviderUpgraded // Event containing the contract specifics and raw log

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
func (it *ProviderUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderUpgraded)
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
		it.Event = new(ProviderUpgraded)
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
func (it *ProviderUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderUpgraded represents a Upgraded event raised by the Provider contract.
type ProviderUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Provider *ProviderFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ProviderUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Provider.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ProviderUpgradedIterator{contract: _Provider.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Provider *ProviderFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ProviderUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Provider.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderUpgraded)
				if err := _Provider.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Provider *ProviderFilterer) ParseUpgraded(log types.Log) (*ProviderUpgraded, error) {
	event := new(ProviderUpgraded)
	if err := _Provider.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderUserCdpPositionSyncedIterator is returned from FilterUserCdpPositionSynced and is used to iterate over the raw logs and unpacked data for UserCdpPositionSynced events raised by the Provider contract.
type ProviderUserCdpPositionSyncedIterator struct {
	Event *ProviderUserCdpPositionSynced // Event containing the contract specifics and raw log

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
func (it *ProviderUserCdpPositionSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderUserCdpPositionSynced)
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
		it.Event = new(ProviderUserCdpPositionSynced)
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
func (it *ProviderUserCdpPositionSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderUserCdpPositionSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderUserCdpPositionSynced represents a UserCdpPositionSynced event raised by the Provider contract.
type ProviderUserCdpPositionSynced struct {
	User                 common.Address
	UserTotalLpValue     *big.Int
	UserCdpPositionValue *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUserCdpPositionSynced is a free log retrieval operation binding the contract event 0xc01a01a0ac3e54f7056a33a026f41039e5ad87f7074f52b37252e253f6d91efb.
//
// Solidity: event UserCdpPositionSynced(address user, uint256 userTotalLpValue, uint256 userCdpPositionValue)
func (_Provider *ProviderFilterer) FilterUserCdpPositionSynced(opts *bind.FilterOpts) (*ProviderUserCdpPositionSyncedIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "UserCdpPositionSynced")
	if err != nil {
		return nil, err
	}
	return &ProviderUserCdpPositionSyncedIterator{contract: _Provider.contract, event: "UserCdpPositionSynced", logs: logs, sub: sub}, nil
}

// WatchUserCdpPositionSynced is a free log subscription operation binding the contract event 0xc01a01a0ac3e54f7056a33a026f41039e5ad87f7074f52b37252e253f6d91efb.
//
// Solidity: event UserCdpPositionSynced(address user, uint256 userTotalLpValue, uint256 userCdpPositionValue)
func (_Provider *ProviderFilterer) WatchUserCdpPositionSynced(opts *bind.WatchOpts, sink chan<- *ProviderUserCdpPositionSynced) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "UserCdpPositionSynced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderUserCdpPositionSynced)
				if err := _Provider.contract.UnpackLog(event, "UserCdpPositionSynced", log); err != nil {
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

// ParseUserCdpPositionSynced is a log parse operation binding the contract event 0xc01a01a0ac3e54f7056a33a026f41039e5ad87f7074f52b37252e253f6d91efb.
//
// Solidity: event UserCdpPositionSynced(address user, uint256 userTotalLpValue, uint256 userCdpPositionValue)
func (_Provider *ProviderFilterer) ParseUserCdpPositionSynced(log types.Log) (*ProviderUserCdpPositionSynced, error) {
	event := new(ProviderUserCdpPositionSynced)
	if err := _Provider.contract.UnpackLog(event, "UserCdpPositionSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderWithdrawLpIterator is returned from FilterWithdrawLp and is used to iterate over the raw logs and unpacked data for WithdrawLp events raised by the Provider contract.
type ProviderWithdrawLpIterator struct {
	Event *ProviderWithdrawLp // Event containing the contract specifics and raw log

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
func (it *ProviderWithdrawLpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderWithdrawLp)
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
		it.Event = new(ProviderWithdrawLp)
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
func (it *ProviderWithdrawLpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderWithdrawLpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderWithdrawLp represents a WithdrawLp event raised by the Provider contract.
type ProviderWithdrawLp struct {
	User    common.Address
	TokenId *big.Int
	LpValue *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawLp is a free log retrieval operation binding the contract event 0xa36baeea9918f1df08335adc43f41a3767beff15a7fb1ed588e9074c8574daa9.
//
// Solidity: event WithdrawLp(address user, uint256 tokenId, uint256 lpValue)
func (_Provider *ProviderFilterer) FilterWithdrawLp(opts *bind.FilterOpts) (*ProviderWithdrawLpIterator, error) {

	logs, sub, err := _Provider.contract.FilterLogs(opts, "WithdrawLp")
	if err != nil {
		return nil, err
	}
	return &ProviderWithdrawLpIterator{contract: _Provider.contract, event: "WithdrawLp", logs: logs, sub: sub}, nil
}

// WatchWithdrawLp is a free log subscription operation binding the contract event 0xa36baeea9918f1df08335adc43f41a3767beff15a7fb1ed588e9074c8574daa9.
//
// Solidity: event WithdrawLp(address user, uint256 tokenId, uint256 lpValue)
func (_Provider *ProviderFilterer) WatchWithdrawLp(opts *bind.WatchOpts, sink chan<- *ProviderWithdrawLp) (event.Subscription, error) {

	logs, sub, err := _Provider.contract.WatchLogs(opts, "WithdrawLp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderWithdrawLp)
				if err := _Provider.contract.UnpackLog(event, "WithdrawLp", log); err != nil {
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

// ParseWithdrawLp is a log parse operation binding the contract event 0xa36baeea9918f1df08335adc43f41a3767beff15a7fb1ed588e9074c8574daa9.
//
// Solidity: event WithdrawLp(address user, uint256 tokenId, uint256 lpValue)
func (_Provider *ProviderFilterer) ParseWithdrawLp(log types.Log) (*ProviderWithdrawLp, error) {
	event := new(ProviderWithdrawLp)
	if err := _Provider.contract.UnpackLog(event, "WithdrawLp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
