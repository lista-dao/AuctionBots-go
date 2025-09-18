// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package interaction

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

// InteractionMetaData contains all meta data concerning the Interaction contract.
var InteractionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addToAuctionWhitelist\",\"inputs\":[{\"name\":\"usrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToBlacklist\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToWhitelist\",\"inputs\":[{\"name\":\"usrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"auctionWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctionWhitelistMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"availableToBorrow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"borrow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hayAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"borrowApr\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"borrowLisUSDListaDistributor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBorrowLisUSDListaDistributor\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"borrowed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buyFromAuction\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiverAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collateralPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralRate\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collateralTVL\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collaterals\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"gem\",\"type\":\"address\",\"internalType\":\"contractGemJoinLike\"},{\"name\":\"ilk\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"live\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"clip\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentLiquidationPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deny\",\"inputs\":[{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dink\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositTVL\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositsDeprecated\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableAuctionWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dog\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"drip\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dutyCalculator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDynamicDutyCalculator\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enableAuctionWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableWhitelist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"free\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllActiveAuctionsForToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"sales\",\"type\":\"tuple[]\",\"internalType\":\"structSale[]\",\"components\":[{\"name\":\"pos\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tab\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"lot\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tic\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"top\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAuctionStatus\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20Upgradeable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hayJoin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractHayJoinLike\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hayPrice\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"helioProviders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"helioRewards\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRewards\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"vat_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spot_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hay_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hayJoin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"jug_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dog_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"jug\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractJugLike\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"locked\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payback\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hayAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paybackFor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hayAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"borrower\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"poke\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"providerCompatibilityMode\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rely\",\"inputs\":[{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeCollateralType\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFromAuctionWhitelist\",\"inputs\":[{\"name\":\"usrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFromBlacklist\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFromWhitelist\",\"inputs\":[{\"name\":\"usrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetAuction\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCollateralDuty\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"duty\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCollateralType\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gemJoin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"ilk\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"clip\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mat\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHelioProvider\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"helioProvider\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"compatibilityMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setListaDistributor\",\"inputs\":[{\"name\":\"distributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelistOperator\",\"inputs\":[{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"spotter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSpotLike\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"startAuction\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"keeper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"syncSnapshot\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokensBlacklist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vat\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractVatLike\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"wards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistOperator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"willBorrow\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dink\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AuctionFinished\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"keeper\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AuctionStarted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Borrow\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"liquidationPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralDisabled\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ilk\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralEnabled\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"ilk\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Liquidation\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"leftover\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Payback\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"collateral\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"debt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"liquidationPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// InteractionABI is the input ABI used to generate the binding from.
// Deprecated: Use InteractionMetaData.ABI instead.
var InteractionABI = InteractionMetaData.ABI

// Interaction is an auto generated Go binding around an Ethereum contract.
type Interaction struct {
	InteractionCaller     // Read-only binding to the contract
	InteractionTransactor // Write-only binding to the contract
	InteractionFilterer   // Log filterer for contract events
}

// InteractionCaller is an auto generated read-only Go binding around an Ethereum contract.
type InteractionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InteractionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InteractionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InteractionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InteractionSession struct {
	Contract     *Interaction      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InteractionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InteractionCallerSession struct {
	Contract *InteractionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// InteractionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InteractionTransactorSession struct {
	Contract     *InteractionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// InteractionRaw is an auto generated low-level Go binding around an Ethereum contract.
type InteractionRaw struct {
	Contract *Interaction // Generic contract binding to access the raw methods on
}

// InteractionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InteractionCallerRaw struct {
	Contract *InteractionCaller // Generic read-only contract binding to access the raw methods on
}

// InteractionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InteractionTransactorRaw struct {
	Contract *InteractionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInteraction creates a new instance of Interaction, bound to a specific deployed contract.
func NewInteraction(address common.Address, backend bind.ContractBackend) (*Interaction, error) {
	contract, err := bindInteraction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Interaction{InteractionCaller: InteractionCaller{contract: contract}, InteractionTransactor: InteractionTransactor{contract: contract}, InteractionFilterer: InteractionFilterer{contract: contract}}, nil
}

// NewInteractionCaller creates a new read-only instance of Interaction, bound to a specific deployed contract.
func NewInteractionCaller(address common.Address, caller bind.ContractCaller) (*InteractionCaller, error) {
	contract, err := bindInteraction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InteractionCaller{contract: contract}, nil
}

// NewInteractionTransactor creates a new write-only instance of Interaction, bound to a specific deployed contract.
func NewInteractionTransactor(address common.Address, transactor bind.ContractTransactor) (*InteractionTransactor, error) {
	contract, err := bindInteraction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InteractionTransactor{contract: contract}, nil
}

// NewInteractionFilterer creates a new log filterer instance of Interaction, bound to a specific deployed contract.
func NewInteractionFilterer(address common.Address, filterer bind.ContractFilterer) (*InteractionFilterer, error) {
	contract, err := bindInteraction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InteractionFilterer{contract: contract}, nil
}

// bindInteraction binds a generic wrapper to an already deployed contract.
func bindInteraction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InteractionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interaction *InteractionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interaction.Contract.InteractionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interaction *InteractionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interaction.Contract.InteractionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interaction *InteractionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interaction.Contract.InteractionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Interaction *InteractionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Interaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Interaction *InteractionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Interaction *InteractionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Interaction.Contract.contract.Transact(opts, method, params...)
}

// AuctionWhitelist is a free data retrieval call binding the contract method 0x9e068764.
//
// Solidity: function auctionWhitelist(address ) view returns(uint256)
func (_Interaction *InteractionCaller) AuctionWhitelist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "auctionWhitelist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionWhitelist is a free data retrieval call binding the contract method 0x9e068764.
//
// Solidity: function auctionWhitelist(address ) view returns(uint256)
func (_Interaction *InteractionSession) AuctionWhitelist(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.AuctionWhitelist(&_Interaction.CallOpts, arg0)
}

// AuctionWhitelist is a free data retrieval call binding the contract method 0x9e068764.
//
// Solidity: function auctionWhitelist(address ) view returns(uint256)
func (_Interaction *InteractionCallerSession) AuctionWhitelist(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.AuctionWhitelist(&_Interaction.CallOpts, arg0)
}

// AuctionWhitelistMode is a free data retrieval call binding the contract method 0x1e9990ff.
//
// Solidity: function auctionWhitelistMode() view returns(uint256)
func (_Interaction *InteractionCaller) AuctionWhitelistMode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "auctionWhitelistMode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionWhitelistMode is a free data retrieval call binding the contract method 0x1e9990ff.
//
// Solidity: function auctionWhitelistMode() view returns(uint256)
func (_Interaction *InteractionSession) AuctionWhitelistMode() (*big.Int, error) {
	return _Interaction.Contract.AuctionWhitelistMode(&_Interaction.CallOpts)
}

// AuctionWhitelistMode is a free data retrieval call binding the contract method 0x1e9990ff.
//
// Solidity: function auctionWhitelistMode() view returns(uint256)
func (_Interaction *InteractionCallerSession) AuctionWhitelistMode() (*big.Int, error) {
	return _Interaction.Contract.AuctionWhitelistMode(&_Interaction.CallOpts)
}

// AvailableToBorrow is a free data retrieval call binding the contract method 0xdc7e91dd.
//
// Solidity: function availableToBorrow(address token, address usr) view returns(int256 amount)
func (_Interaction *InteractionCaller) AvailableToBorrow(opts *bind.CallOpts, token common.Address, usr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "availableToBorrow", token, usr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableToBorrow is a free data retrieval call binding the contract method 0xdc7e91dd.
//
// Solidity: function availableToBorrow(address token, address usr) view returns(int256 amount)
func (_Interaction *InteractionSession) AvailableToBorrow(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.AvailableToBorrow(&_Interaction.CallOpts, token, usr)
}

// AvailableToBorrow is a free data retrieval call binding the contract method 0xdc7e91dd.
//
// Solidity: function availableToBorrow(address token, address usr) view returns(int256 amount)
func (_Interaction *InteractionCallerSession) AvailableToBorrow(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.AvailableToBorrow(&_Interaction.CallOpts, token, usr)
}

// BorrowApr is a free data retrieval call binding the contract method 0x9fecc37d.
//
// Solidity: function borrowApr(address token) view returns(uint256)
func (_Interaction *InteractionCaller) BorrowApr(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "borrowApr", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowApr is a free data retrieval call binding the contract method 0x9fecc37d.
//
// Solidity: function borrowApr(address token) view returns(uint256)
func (_Interaction *InteractionSession) BorrowApr(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.BorrowApr(&_Interaction.CallOpts, token)
}

// BorrowApr is a free data retrieval call binding the contract method 0x9fecc37d.
//
// Solidity: function borrowApr(address token) view returns(uint256)
func (_Interaction *InteractionCallerSession) BorrowApr(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.BorrowApr(&_Interaction.CallOpts, token)
}

// BorrowLisUSDListaDistributor is a free data retrieval call binding the contract method 0xe68d57cf.
//
// Solidity: function borrowLisUSDListaDistributor() view returns(address)
func (_Interaction *InteractionCaller) BorrowLisUSDListaDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "borrowLisUSDListaDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowLisUSDListaDistributor is a free data retrieval call binding the contract method 0xe68d57cf.
//
// Solidity: function borrowLisUSDListaDistributor() view returns(address)
func (_Interaction *InteractionSession) BorrowLisUSDListaDistributor() (common.Address, error) {
	return _Interaction.Contract.BorrowLisUSDListaDistributor(&_Interaction.CallOpts)
}

// BorrowLisUSDListaDistributor is a free data retrieval call binding the contract method 0xe68d57cf.
//
// Solidity: function borrowLisUSDListaDistributor() view returns(address)
func (_Interaction *InteractionCallerSession) BorrowLisUSDListaDistributor() (common.Address, error) {
	return _Interaction.Contract.BorrowLisUSDListaDistributor(&_Interaction.CallOpts)
}

// Borrowed is a free data retrieval call binding the contract method 0xb0a02abe.
//
// Solidity: function borrowed(address token, address usr) view returns(uint256)
func (_Interaction *InteractionCaller) Borrowed(opts *bind.CallOpts, token common.Address, usr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "borrowed", token, usr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Borrowed is a free data retrieval call binding the contract method 0xb0a02abe.
//
// Solidity: function borrowed(address token, address usr) view returns(uint256)
func (_Interaction *InteractionSession) Borrowed(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.Borrowed(&_Interaction.CallOpts, token, usr)
}

// Borrowed is a free data retrieval call binding the contract method 0xb0a02abe.
//
// Solidity: function borrowed(address token, address usr) view returns(uint256)
func (_Interaction *InteractionCallerSession) Borrowed(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.Borrowed(&_Interaction.CallOpts, token, usr)
}

// CollateralPrice is a free data retrieval call binding the contract method 0x08b1cb23.
//
// Solidity: function collateralPrice(address token) view returns(uint256)
func (_Interaction *InteractionCaller) CollateralPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "collateralPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralPrice is a free data retrieval call binding the contract method 0x08b1cb23.
//
// Solidity: function collateralPrice(address token) view returns(uint256)
func (_Interaction *InteractionSession) CollateralPrice(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.CollateralPrice(&_Interaction.CallOpts, token)
}

// CollateralPrice is a free data retrieval call binding the contract method 0x08b1cb23.
//
// Solidity: function collateralPrice(address token) view returns(uint256)
func (_Interaction *InteractionCallerSession) CollateralPrice(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.CollateralPrice(&_Interaction.CallOpts, token)
}

// CollateralRate is a free data retrieval call binding the contract method 0x72c8fef0.
//
// Solidity: function collateralRate(address token) view returns(uint256)
func (_Interaction *InteractionCaller) CollateralRate(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "collateralRate", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralRate is a free data retrieval call binding the contract method 0x72c8fef0.
//
// Solidity: function collateralRate(address token) view returns(uint256)
func (_Interaction *InteractionSession) CollateralRate(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.CollateralRate(&_Interaction.CallOpts, token)
}

// CollateralRate is a free data retrieval call binding the contract method 0x72c8fef0.
//
// Solidity: function collateralRate(address token) view returns(uint256)
func (_Interaction *InteractionCallerSession) CollateralRate(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.CollateralRate(&_Interaction.CallOpts, token)
}

// CollateralTVL is a free data retrieval call binding the contract method 0x4eb3048a.
//
// Solidity: function collateralTVL(address token) view returns(uint256)
func (_Interaction *InteractionCaller) CollateralTVL(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "collateralTVL", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollateralTVL is a free data retrieval call binding the contract method 0x4eb3048a.
//
// Solidity: function collateralTVL(address token) view returns(uint256)
func (_Interaction *InteractionSession) CollateralTVL(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.CollateralTVL(&_Interaction.CallOpts, token)
}

// CollateralTVL is a free data retrieval call binding the contract method 0x4eb3048a.
//
// Solidity: function collateralTVL(address token) view returns(uint256)
func (_Interaction *InteractionCallerSession) CollateralTVL(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.CollateralTVL(&_Interaction.CallOpts, token)
}

// Collaterals is a free data retrieval call binding the contract method 0xeeb97d3b.
//
// Solidity: function collaterals(address ) view returns(address gem, bytes32 ilk, uint32 live, address clip)
func (_Interaction *InteractionCaller) Collaterals(opts *bind.CallOpts, arg0 common.Address) (struct {
	Gem  common.Address
	Ilk  [32]byte
	Live uint32
	Clip common.Address
}, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "collaterals", arg0)

	outstruct := new(struct {
		Gem  common.Address
		Ilk  [32]byte
		Live uint32
		Clip common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Gem = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Ilk = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Live = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Clip = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Collaterals is a free data retrieval call binding the contract method 0xeeb97d3b.
//
// Solidity: function collaterals(address ) view returns(address gem, bytes32 ilk, uint32 live, address clip)
func (_Interaction *InteractionSession) Collaterals(arg0 common.Address) (struct {
	Gem  common.Address
	Ilk  [32]byte
	Live uint32
	Clip common.Address
}, error) {
	return _Interaction.Contract.Collaterals(&_Interaction.CallOpts, arg0)
}

// Collaterals is a free data retrieval call binding the contract method 0xeeb97d3b.
//
// Solidity: function collaterals(address ) view returns(address gem, bytes32 ilk, uint32 live, address clip)
func (_Interaction *InteractionCallerSession) Collaterals(arg0 common.Address) (struct {
	Gem  common.Address
	Ilk  [32]byte
	Live uint32
	Clip common.Address
}, error) {
	return _Interaction.Contract.Collaterals(&_Interaction.CallOpts, arg0)
}

// CurrentLiquidationPrice is a free data retrieval call binding the contract method 0xfc085c11.
//
// Solidity: function currentLiquidationPrice(address token, address usr) view returns(uint256)
func (_Interaction *InteractionCaller) CurrentLiquidationPrice(opts *bind.CallOpts, token common.Address, usr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "currentLiquidationPrice", token, usr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentLiquidationPrice is a free data retrieval call binding the contract method 0xfc085c11.
//
// Solidity: function currentLiquidationPrice(address token, address usr) view returns(uint256)
func (_Interaction *InteractionSession) CurrentLiquidationPrice(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.CurrentLiquidationPrice(&_Interaction.CallOpts, token, usr)
}

// CurrentLiquidationPrice is a free data retrieval call binding the contract method 0xfc085c11.
//
// Solidity: function currentLiquidationPrice(address token, address usr) view returns(uint256)
func (_Interaction *InteractionCallerSession) CurrentLiquidationPrice(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.CurrentLiquidationPrice(&_Interaction.CallOpts, token, usr)
}

// DepositTVL is a free data retrieval call binding the contract method 0x89a283b9.
//
// Solidity: function depositTVL(address token) view returns(uint256)
func (_Interaction *InteractionCaller) DepositTVL(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "depositTVL", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositTVL is a free data retrieval call binding the contract method 0x89a283b9.
//
// Solidity: function depositTVL(address token) view returns(uint256)
func (_Interaction *InteractionSession) DepositTVL(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.DepositTVL(&_Interaction.CallOpts, token)
}

// DepositTVL is a free data retrieval call binding the contract method 0x89a283b9.
//
// Solidity: function depositTVL(address token) view returns(uint256)
func (_Interaction *InteractionCallerSession) DepositTVL(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.DepositTVL(&_Interaction.CallOpts, token)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address token) view returns(uint256)
func (_Interaction *InteractionCaller) Deposits(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "deposits", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address token) view returns(uint256)
func (_Interaction *InteractionSession) Deposits(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.Deposits(&_Interaction.CallOpts, token)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address token) view returns(uint256)
func (_Interaction *InteractionCallerSession) Deposits(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.Deposits(&_Interaction.CallOpts, token)
}

// DepositsDeprecated is a free data retrieval call binding the contract method 0xe57fc302.
//
// Solidity: function depositsDeprecated(address ) view returns(uint256)
func (_Interaction *InteractionCaller) DepositsDeprecated(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "depositsDeprecated", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsDeprecated is a free data retrieval call binding the contract method 0xe57fc302.
//
// Solidity: function depositsDeprecated(address ) view returns(uint256)
func (_Interaction *InteractionSession) DepositsDeprecated(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.DepositsDeprecated(&_Interaction.CallOpts, arg0)
}

// DepositsDeprecated is a free data retrieval call binding the contract method 0xe57fc302.
//
// Solidity: function depositsDeprecated(address ) view returns(uint256)
func (_Interaction *InteractionCallerSession) DepositsDeprecated(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.DepositsDeprecated(&_Interaction.CallOpts, arg0)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Interaction *InteractionCaller) Dog(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "dog")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Interaction *InteractionSession) Dog() (common.Address, error) {
	return _Interaction.Contract.Dog(&_Interaction.CallOpts)
}

// Dog is a free data retrieval call binding the contract method 0xc3b3ad7f.
//
// Solidity: function dog() view returns(address)
func (_Interaction *InteractionCallerSession) Dog() (common.Address, error) {
	return _Interaction.Contract.Dog(&_Interaction.CallOpts)
}

// DutyCalculator is a free data retrieval call binding the contract method 0xebb95c5d.
//
// Solidity: function dutyCalculator() view returns(address)
func (_Interaction *InteractionCaller) DutyCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "dutyCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DutyCalculator is a free data retrieval call binding the contract method 0xebb95c5d.
//
// Solidity: function dutyCalculator() view returns(address)
func (_Interaction *InteractionSession) DutyCalculator() (common.Address, error) {
	return _Interaction.Contract.DutyCalculator(&_Interaction.CallOpts)
}

// DutyCalculator is a free data retrieval call binding the contract method 0xebb95c5d.
//
// Solidity: function dutyCalculator() view returns(address)
func (_Interaction *InteractionCallerSession) DutyCalculator() (common.Address, error) {
	return _Interaction.Contract.DutyCalculator(&_Interaction.CallOpts)
}

// Free is a free data retrieval call binding the contract method 0xc5cafb88.
//
// Solidity: function free(address token, address usr) view returns(uint256)
func (_Interaction *InteractionCaller) Free(opts *bind.CallOpts, token common.Address, usr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "free", token, usr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Free is a free data retrieval call binding the contract method 0xc5cafb88.
//
// Solidity: function free(address token, address usr) view returns(uint256)
func (_Interaction *InteractionSession) Free(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.Free(&_Interaction.CallOpts, token, usr)
}

// Free is a free data retrieval call binding the contract method 0xc5cafb88.
//
// Solidity: function free(address token, address usr) view returns(uint256)
func (_Interaction *InteractionCallerSession) Free(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.Free(&_Interaction.CallOpts, token, usr)
}

// GetAllActiveAuctionsForToken is a free data retrieval call binding the contract method 0x5c49b500.
//
// Solidity: function getAllActiveAuctionsForToken(address token) view returns((uint256,uint256,uint256,address,uint96,uint256)[] sales)
func (_Interaction *InteractionCaller) GetAllActiveAuctionsForToken(opts *bind.CallOpts, token common.Address) ([]Sale, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "getAllActiveAuctionsForToken", token)

	if err != nil {
		return *new([]Sale), err
	}

	out0 := *abi.ConvertType(out[0], new([]Sale)).(*[]Sale)

	return out0, err

}

// GetAllActiveAuctionsForToken is a free data retrieval call binding the contract method 0x5c49b500.
//
// Solidity: function getAllActiveAuctionsForToken(address token) view returns((uint256,uint256,uint256,address,uint96,uint256)[] sales)
func (_Interaction *InteractionSession) GetAllActiveAuctionsForToken(token common.Address) ([]Sale, error) {
	return _Interaction.Contract.GetAllActiveAuctionsForToken(&_Interaction.CallOpts, token)
}

// GetAllActiveAuctionsForToken is a free data retrieval call binding the contract method 0x5c49b500.
//
// Solidity: function getAllActiveAuctionsForToken(address token) view returns((uint256,uint256,uint256,address,uint96,uint256)[] sales)
func (_Interaction *InteractionCallerSession) GetAllActiveAuctionsForToken(token common.Address) ([]Sale, error) {
	return _Interaction.Contract.GetAllActiveAuctionsForToken(&_Interaction.CallOpts, token)
}

// GetAuctionStatus is a free data retrieval call binding the contract method 0xf8a206fa.
//
// Solidity: function getAuctionStatus(address token, uint256 auctionId) view returns(bool, uint256, uint256, uint256)
func (_Interaction *InteractionCaller) GetAuctionStatus(opts *bind.CallOpts, token common.Address, auctionId *big.Int) (bool, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "getAuctionStatus", token, auctionId)

	if err != nil {
		return *new(bool), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAuctionStatus is a free data retrieval call binding the contract method 0xf8a206fa.
//
// Solidity: function getAuctionStatus(address token, uint256 auctionId) view returns(bool, uint256, uint256, uint256)
func (_Interaction *InteractionSession) GetAuctionStatus(token common.Address, auctionId *big.Int) (bool, *big.Int, *big.Int, *big.Int, error) {
	return _Interaction.Contract.GetAuctionStatus(&_Interaction.CallOpts, token, auctionId)
}

// GetAuctionStatus is a free data retrieval call binding the contract method 0xf8a206fa.
//
// Solidity: function getAuctionStatus(address token, uint256 auctionId) view returns(bool, uint256, uint256, uint256)
func (_Interaction *InteractionCallerSession) GetAuctionStatus(token common.Address, auctionId *big.Int) (bool, *big.Int, *big.Int, *big.Int, error) {
	return _Interaction.Contract.GetAuctionStatus(&_Interaction.CallOpts, token, auctionId)
}

// Hay is a free data retrieval call binding the contract method 0xb02d1e9e.
//
// Solidity: function hay() view returns(address)
func (_Interaction *InteractionCaller) Hay(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "hay")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Hay is a free data retrieval call binding the contract method 0xb02d1e9e.
//
// Solidity: function hay() view returns(address)
func (_Interaction *InteractionSession) Hay() (common.Address, error) {
	return _Interaction.Contract.Hay(&_Interaction.CallOpts)
}

// Hay is a free data retrieval call binding the contract method 0xb02d1e9e.
//
// Solidity: function hay() view returns(address)
func (_Interaction *InteractionCallerSession) Hay() (common.Address, error) {
	return _Interaction.Contract.Hay(&_Interaction.CallOpts)
}

// HayJoin is a free data retrieval call binding the contract method 0x4d608102.
//
// Solidity: function hayJoin() view returns(address)
func (_Interaction *InteractionCaller) HayJoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "hayJoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HayJoin is a free data retrieval call binding the contract method 0x4d608102.
//
// Solidity: function hayJoin() view returns(address)
func (_Interaction *InteractionSession) HayJoin() (common.Address, error) {
	return _Interaction.Contract.HayJoin(&_Interaction.CallOpts)
}

// HayJoin is a free data retrieval call binding the contract method 0x4d608102.
//
// Solidity: function hayJoin() view returns(address)
func (_Interaction *InteractionCallerSession) HayJoin() (common.Address, error) {
	return _Interaction.Contract.HayJoin(&_Interaction.CallOpts)
}

// HayPrice is a free data retrieval call binding the contract method 0xdddca8af.
//
// Solidity: function hayPrice(address token) view returns(uint256)
func (_Interaction *InteractionCaller) HayPrice(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "hayPrice", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HayPrice is a free data retrieval call binding the contract method 0xdddca8af.
//
// Solidity: function hayPrice(address token) view returns(uint256)
func (_Interaction *InteractionSession) HayPrice(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.HayPrice(&_Interaction.CallOpts, token)
}

// HayPrice is a free data retrieval call binding the contract method 0xdddca8af.
//
// Solidity: function hayPrice(address token) view returns(uint256)
func (_Interaction *InteractionCallerSession) HayPrice(token common.Address) (*big.Int, error) {
	return _Interaction.Contract.HayPrice(&_Interaction.CallOpts, token)
}

// HelioProviders is a free data retrieval call binding the contract method 0x1ef0bb94.
//
// Solidity: function helioProviders(address ) view returns(address)
func (_Interaction *InteractionCaller) HelioProviders(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "helioProviders", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HelioProviders is a free data retrieval call binding the contract method 0x1ef0bb94.
//
// Solidity: function helioProviders(address ) view returns(address)
func (_Interaction *InteractionSession) HelioProviders(arg0 common.Address) (common.Address, error) {
	return _Interaction.Contract.HelioProviders(&_Interaction.CallOpts, arg0)
}

// HelioProviders is a free data retrieval call binding the contract method 0x1ef0bb94.
//
// Solidity: function helioProviders(address ) view returns(address)
func (_Interaction *InteractionCallerSession) HelioProviders(arg0 common.Address) (common.Address, error) {
	return _Interaction.Contract.HelioProviders(&_Interaction.CallOpts, arg0)
}

// HelioRewards is a free data retrieval call binding the contract method 0xff4e65f3.
//
// Solidity: function helioRewards() view returns(address)
func (_Interaction *InteractionCaller) HelioRewards(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "helioRewards")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HelioRewards is a free data retrieval call binding the contract method 0xff4e65f3.
//
// Solidity: function helioRewards() view returns(address)
func (_Interaction *InteractionSession) HelioRewards() (common.Address, error) {
	return _Interaction.Contract.HelioRewards(&_Interaction.CallOpts)
}

// HelioRewards is a free data retrieval call binding the contract method 0xff4e65f3.
//
// Solidity: function helioRewards() view returns(address)
func (_Interaction *InteractionCallerSession) HelioRewards() (common.Address, error) {
	return _Interaction.Contract.HelioRewards(&_Interaction.CallOpts)
}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_Interaction *InteractionCaller) Jug(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "jug")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_Interaction *InteractionSession) Jug() (common.Address, error) {
	return _Interaction.Contract.Jug(&_Interaction.CallOpts)
}

// Jug is a free data retrieval call binding the contract method 0x84718d89.
//
// Solidity: function jug() view returns(address)
func (_Interaction *InteractionCallerSession) Jug() (common.Address, error) {
	return _Interaction.Contract.Jug(&_Interaction.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address token, address usr) view returns(uint256)
func (_Interaction *InteractionCaller) Locked(opts *bind.CallOpts, token common.Address, usr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "locked", token, usr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address token, address usr) view returns(uint256)
func (_Interaction *InteractionSession) Locked(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.Locked(&_Interaction.CallOpts, token, usr)
}

// Locked is a free data retrieval call binding the contract method 0xdb20266f.
//
// Solidity: function locked(address token, address usr) view returns(uint256)
func (_Interaction *InteractionCallerSession) Locked(token common.Address, usr common.Address) (*big.Int, error) {
	return _Interaction.Contract.Locked(&_Interaction.CallOpts, token, usr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Interaction *InteractionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Interaction *InteractionSession) Owner() (common.Address, error) {
	return _Interaction.Contract.Owner(&_Interaction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Interaction *InteractionCallerSession) Owner() (common.Address, error) {
	return _Interaction.Contract.Owner(&_Interaction.CallOpts)
}

// ProviderCompatibilityMode is a free data retrieval call binding the contract method 0xe107f49b.
//
// Solidity: function providerCompatibilityMode(address ) view returns(bool)
func (_Interaction *InteractionCaller) ProviderCompatibilityMode(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "providerCompatibilityMode", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProviderCompatibilityMode is a free data retrieval call binding the contract method 0xe107f49b.
//
// Solidity: function providerCompatibilityMode(address ) view returns(bool)
func (_Interaction *InteractionSession) ProviderCompatibilityMode(arg0 common.Address) (bool, error) {
	return _Interaction.Contract.ProviderCompatibilityMode(&_Interaction.CallOpts, arg0)
}

// ProviderCompatibilityMode is a free data retrieval call binding the contract method 0xe107f49b.
//
// Solidity: function providerCompatibilityMode(address ) view returns(bool)
func (_Interaction *InteractionCallerSession) ProviderCompatibilityMode(arg0 common.Address) (bool, error) {
	return _Interaction.Contract.ProviderCompatibilityMode(&_Interaction.CallOpts, arg0)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Interaction *InteractionCaller) Spotter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "spotter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Interaction *InteractionSession) Spotter() (common.Address, error) {
	return _Interaction.Contract.Spotter(&_Interaction.CallOpts)
}

// Spotter is a free data retrieval call binding the contract method 0x2e77468d.
//
// Solidity: function spotter() view returns(address)
func (_Interaction *InteractionCallerSession) Spotter() (common.Address, error) {
	return _Interaction.Contract.Spotter(&_Interaction.CallOpts)
}

// TokensBlacklist is a free data retrieval call binding the contract method 0x293e10ad.
//
// Solidity: function tokensBlacklist(address ) view returns(uint256)
func (_Interaction *InteractionCaller) TokensBlacklist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "tokensBlacklist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensBlacklist is a free data retrieval call binding the contract method 0x293e10ad.
//
// Solidity: function tokensBlacklist(address ) view returns(uint256)
func (_Interaction *InteractionSession) TokensBlacklist(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.TokensBlacklist(&_Interaction.CallOpts, arg0)
}

// TokensBlacklist is a free data retrieval call binding the contract method 0x293e10ad.
//
// Solidity: function tokensBlacklist(address ) view returns(uint256)
func (_Interaction *InteractionCallerSession) TokensBlacklist(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.TokensBlacklist(&_Interaction.CallOpts, arg0)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Interaction *InteractionCaller) Vat(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "vat")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Interaction *InteractionSession) Vat() (common.Address, error) {
	return _Interaction.Contract.Vat(&_Interaction.CallOpts)
}

// Vat is a free data retrieval call binding the contract method 0x36569e77.
//
// Solidity: function vat() view returns(address)
func (_Interaction *InteractionCallerSession) Vat() (common.Address, error) {
	return _Interaction.Contract.Vat(&_Interaction.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Interaction *InteractionCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "wards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Interaction *InteractionSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.Wards(&_Interaction.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_Interaction *InteractionCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.Wards(&_Interaction.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(uint256)
func (_Interaction *InteractionCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(uint256)
func (_Interaction *InteractionSession) Whitelist(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.Whitelist(&_Interaction.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(uint256)
func (_Interaction *InteractionCallerSession) Whitelist(arg0 common.Address) (*big.Int, error) {
	return _Interaction.Contract.Whitelist(&_Interaction.CallOpts, arg0)
}

// WhitelistMode is a free data retrieval call binding the contract method 0x70c757ec.
//
// Solidity: function whitelistMode() view returns(uint256)
func (_Interaction *InteractionCaller) WhitelistMode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "whitelistMode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistMode is a free data retrieval call binding the contract method 0x70c757ec.
//
// Solidity: function whitelistMode() view returns(uint256)
func (_Interaction *InteractionSession) WhitelistMode() (*big.Int, error) {
	return _Interaction.Contract.WhitelistMode(&_Interaction.CallOpts)
}

// WhitelistMode is a free data retrieval call binding the contract method 0x70c757ec.
//
// Solidity: function whitelistMode() view returns(uint256)
func (_Interaction *InteractionCallerSession) WhitelistMode() (*big.Int, error) {
	return _Interaction.Contract.WhitelistMode(&_Interaction.CallOpts)
}

// WhitelistOperator is a free data retrieval call binding the contract method 0x44c43782.
//
// Solidity: function whitelistOperator() view returns(address)
func (_Interaction *InteractionCaller) WhitelistOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "whitelistOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistOperator is a free data retrieval call binding the contract method 0x44c43782.
//
// Solidity: function whitelistOperator() view returns(address)
func (_Interaction *InteractionSession) WhitelistOperator() (common.Address, error) {
	return _Interaction.Contract.WhitelistOperator(&_Interaction.CallOpts)
}

// WhitelistOperator is a free data retrieval call binding the contract method 0x44c43782.
//
// Solidity: function whitelistOperator() view returns(address)
func (_Interaction *InteractionCallerSession) WhitelistOperator() (common.Address, error) {
	return _Interaction.Contract.WhitelistOperator(&_Interaction.CallOpts)
}

// WillBorrow is a free data retrieval call binding the contract method 0xaa592eee.
//
// Solidity: function willBorrow(address token, address usr, int256 amount) view returns(int256)
func (_Interaction *InteractionCaller) WillBorrow(opts *bind.CallOpts, token common.Address, usr common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Interaction.contract.Call(opts, &out, "willBorrow", token, usr, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WillBorrow is a free data retrieval call binding the contract method 0xaa592eee.
//
// Solidity: function willBorrow(address token, address usr, int256 amount) view returns(int256)
func (_Interaction *InteractionSession) WillBorrow(token common.Address, usr common.Address, amount *big.Int) (*big.Int, error) {
	return _Interaction.Contract.WillBorrow(&_Interaction.CallOpts, token, usr, amount)
}

// WillBorrow is a free data retrieval call binding the contract method 0xaa592eee.
//
// Solidity: function willBorrow(address token, address usr, int256 amount) view returns(int256)
func (_Interaction *InteractionCallerSession) WillBorrow(token common.Address, usr common.Address, amount *big.Int) (*big.Int, error) {
	return _Interaction.Contract.WillBorrow(&_Interaction.CallOpts, token, usr, amount)
}

// AddToAuctionWhitelist is a paid mutator transaction binding the contract method 0xf746d8f4.
//
// Solidity: function addToAuctionWhitelist(address[] usrs) returns()
func (_Interaction *InteractionTransactor) AddToAuctionWhitelist(opts *bind.TransactOpts, usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "addToAuctionWhitelist", usrs)
}

// AddToAuctionWhitelist is a paid mutator transaction binding the contract method 0xf746d8f4.
//
// Solidity: function addToAuctionWhitelist(address[] usrs) returns()
func (_Interaction *InteractionSession) AddToAuctionWhitelist(usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.AddToAuctionWhitelist(&_Interaction.TransactOpts, usrs)
}

// AddToAuctionWhitelist is a paid mutator transaction binding the contract method 0xf746d8f4.
//
// Solidity: function addToAuctionWhitelist(address[] usrs) returns()
func (_Interaction *InteractionTransactorSession) AddToAuctionWhitelist(usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.AddToAuctionWhitelist(&_Interaction.TransactOpts, usrs)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] tokens) returns()
func (_Interaction *InteractionTransactor) AddToBlacklist(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "addToBlacklist", tokens)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] tokens) returns()
func (_Interaction *InteractionSession) AddToBlacklist(tokens []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.AddToBlacklist(&_Interaction.TransactOpts, tokens)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x935eb35f.
//
// Solidity: function addToBlacklist(address[] tokens) returns()
func (_Interaction *InteractionTransactorSession) AddToBlacklist(tokens []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.AddToBlacklist(&_Interaction.TransactOpts, tokens)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] usrs) returns()
func (_Interaction *InteractionTransactor) AddToWhitelist(opts *bind.TransactOpts, usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "addToWhitelist", usrs)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] usrs) returns()
func (_Interaction *InteractionSession) AddToWhitelist(usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.AddToWhitelist(&_Interaction.TransactOpts, usrs)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0x7f649783.
//
// Solidity: function addToWhitelist(address[] usrs) returns()
func (_Interaction *InteractionTransactorSession) AddToWhitelist(usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.AddToWhitelist(&_Interaction.TransactOpts, usrs)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address token, uint256 hayAmount) returns(uint256)
func (_Interaction *InteractionTransactor) Borrow(opts *bind.TransactOpts, token common.Address, hayAmount *big.Int) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "borrow", token, hayAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address token, uint256 hayAmount) returns(uint256)
func (_Interaction *InteractionSession) Borrow(token common.Address, hayAmount *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.Borrow(&_Interaction.TransactOpts, token, hayAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b8a3529.
//
// Solidity: function borrow(address token, uint256 hayAmount) returns(uint256)
func (_Interaction *InteractionTransactorSession) Borrow(token common.Address, hayAmount *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.Borrow(&_Interaction.TransactOpts, token, hayAmount)
}

// BuyFromAuction is a paid mutator transaction binding the contract method 0xe5ff2f8e.
//
// Solidity: function buyFromAuction(address token, uint256 auctionId, uint256 collateralAmount, uint256 maxPrice, address receiverAddress, bytes data) returns()
func (_Interaction *InteractionTransactor) BuyFromAuction(opts *bind.TransactOpts, token common.Address, auctionId *big.Int, collateralAmount *big.Int, maxPrice *big.Int, receiverAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "buyFromAuction", token, auctionId, collateralAmount, maxPrice, receiverAddress, data)
}

// BuyFromAuction is a paid mutator transaction binding the contract method 0xe5ff2f8e.
//
// Solidity: function buyFromAuction(address token, uint256 auctionId, uint256 collateralAmount, uint256 maxPrice, address receiverAddress, bytes data) returns()
func (_Interaction *InteractionSession) BuyFromAuction(token common.Address, auctionId *big.Int, collateralAmount *big.Int, maxPrice *big.Int, receiverAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Interaction.Contract.BuyFromAuction(&_Interaction.TransactOpts, token, auctionId, collateralAmount, maxPrice, receiverAddress, data)
}

// BuyFromAuction is a paid mutator transaction binding the contract method 0xe5ff2f8e.
//
// Solidity: function buyFromAuction(address token, uint256 auctionId, uint256 collateralAmount, uint256 maxPrice, address receiverAddress, bytes data) returns()
func (_Interaction *InteractionTransactorSession) BuyFromAuction(token common.Address, auctionId *big.Int, collateralAmount *big.Int, maxPrice *big.Int, receiverAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Interaction.Contract.BuyFromAuction(&_Interaction.TransactOpts, token, auctionId, collateralAmount, maxPrice, receiverAddress, data)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Interaction *InteractionTransactor) Deny(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "deny", usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Interaction *InteractionSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Deny(&_Interaction.TransactOpts, usr)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address usr) returns()
func (_Interaction *InteractionTransactorSession) Deny(usr common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Deny(&_Interaction.TransactOpts, usr)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address participant, address token, uint256 dink) returns(uint256)
func (_Interaction *InteractionTransactor) Deposit(opts *bind.TransactOpts, participant common.Address, token common.Address, dink *big.Int) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "deposit", participant, token, dink)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address participant, address token, uint256 dink) returns(uint256)
func (_Interaction *InteractionSession) Deposit(participant common.Address, token common.Address, dink *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.Deposit(&_Interaction.TransactOpts, participant, token, dink)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(address participant, address token, uint256 dink) returns(uint256)
func (_Interaction *InteractionTransactorSession) Deposit(participant common.Address, token common.Address, dink *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.Deposit(&_Interaction.TransactOpts, participant, token, dink)
}

// DisableAuctionWhitelist is a paid mutator transaction binding the contract method 0x2ad12d87.
//
// Solidity: function disableAuctionWhitelist() returns()
func (_Interaction *InteractionTransactor) DisableAuctionWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "disableAuctionWhitelist")
}

// DisableAuctionWhitelist is a paid mutator transaction binding the contract method 0x2ad12d87.
//
// Solidity: function disableAuctionWhitelist() returns()
func (_Interaction *InteractionSession) DisableAuctionWhitelist() (*types.Transaction, error) {
	return _Interaction.Contract.DisableAuctionWhitelist(&_Interaction.TransactOpts)
}

// DisableAuctionWhitelist is a paid mutator transaction binding the contract method 0x2ad12d87.
//
// Solidity: function disableAuctionWhitelist() returns()
func (_Interaction *InteractionTransactorSession) DisableAuctionWhitelist() (*types.Transaction, error) {
	return _Interaction.Contract.DisableAuctionWhitelist(&_Interaction.TransactOpts)
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_Interaction *InteractionTransactor) DisableWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "disableWhitelist")
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_Interaction *InteractionSession) DisableWhitelist() (*types.Transaction, error) {
	return _Interaction.Contract.DisableWhitelist(&_Interaction.TransactOpts)
}

// DisableWhitelist is a paid mutator transaction binding the contract method 0xd6b0f484.
//
// Solidity: function disableWhitelist() returns()
func (_Interaction *InteractionTransactorSession) DisableWhitelist() (*types.Transaction, error) {
	return _Interaction.Contract.DisableWhitelist(&_Interaction.TransactOpts)
}

// Drip is a paid mutator transaction binding the contract method 0x67a5cd06.
//
// Solidity: function drip(address token) returns()
func (_Interaction *InteractionTransactor) Drip(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "drip", token)
}

// Drip is a paid mutator transaction binding the contract method 0x67a5cd06.
//
// Solidity: function drip(address token) returns()
func (_Interaction *InteractionSession) Drip(token common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Drip(&_Interaction.TransactOpts, token)
}

// Drip is a paid mutator transaction binding the contract method 0x67a5cd06.
//
// Solidity: function drip(address token) returns()
func (_Interaction *InteractionTransactorSession) Drip(token common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Drip(&_Interaction.TransactOpts, token)
}

// EnableAuctionWhitelist is a paid mutator transaction binding the contract method 0x75dd94ec.
//
// Solidity: function enableAuctionWhitelist() returns()
func (_Interaction *InteractionTransactor) EnableAuctionWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "enableAuctionWhitelist")
}

// EnableAuctionWhitelist is a paid mutator transaction binding the contract method 0x75dd94ec.
//
// Solidity: function enableAuctionWhitelist() returns()
func (_Interaction *InteractionSession) EnableAuctionWhitelist() (*types.Transaction, error) {
	return _Interaction.Contract.EnableAuctionWhitelist(&_Interaction.TransactOpts)
}

// EnableAuctionWhitelist is a paid mutator transaction binding the contract method 0x75dd94ec.
//
// Solidity: function enableAuctionWhitelist() returns()
func (_Interaction *InteractionTransactorSession) EnableAuctionWhitelist() (*types.Transaction, error) {
	return _Interaction.Contract.EnableAuctionWhitelist(&_Interaction.TransactOpts)
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_Interaction *InteractionTransactor) EnableWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "enableWhitelist")
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_Interaction *InteractionSession) EnableWhitelist() (*types.Transaction, error) {
	return _Interaction.Contract.EnableWhitelist(&_Interaction.TransactOpts)
}

// EnableWhitelist is a paid mutator transaction binding the contract method 0xcdfb2b4e.
//
// Solidity: function enableWhitelist() returns()
func (_Interaction *InteractionTransactorSession) EnableWhitelist() (*types.Transaction, error) {
	return _Interaction.Contract.EnableWhitelist(&_Interaction.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address vat_, address spot_, address hay_, address hayJoin_, address jug_, address dog_) returns()
func (_Interaction *InteractionTransactor) Initialize(opts *bind.TransactOpts, vat_ common.Address, spot_ common.Address, hay_ common.Address, hayJoin_ common.Address, jug_ common.Address, dog_ common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "initialize", vat_, spot_, hay_, hayJoin_, jug_, dog_)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address vat_, address spot_, address hay_, address hayJoin_, address jug_, address dog_) returns()
func (_Interaction *InteractionSession) Initialize(vat_ common.Address, spot_ common.Address, hay_ common.Address, hayJoin_ common.Address, jug_ common.Address, dog_ common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Initialize(&_Interaction.TransactOpts, vat_, spot_, hay_, hayJoin_, jug_, dog_)
}

// Initialize is a paid mutator transaction binding the contract method 0xcc2a9a5b.
//
// Solidity: function initialize(address vat_, address spot_, address hay_, address hayJoin_, address jug_, address dog_) returns()
func (_Interaction *InteractionTransactorSession) Initialize(vat_ common.Address, spot_ common.Address, hay_ common.Address, hayJoin_ common.Address, jug_ common.Address, dog_ common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Initialize(&_Interaction.TransactOpts, vat_, spot_, hay_, hayJoin_, jug_, dog_)
}

// Payback is a paid mutator transaction binding the contract method 0x35ed8ab8.
//
// Solidity: function payback(address token, uint256 hayAmount) returns(int256)
func (_Interaction *InteractionTransactor) Payback(opts *bind.TransactOpts, token common.Address, hayAmount *big.Int) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "payback", token, hayAmount)
}

// Payback is a paid mutator transaction binding the contract method 0x35ed8ab8.
//
// Solidity: function payback(address token, uint256 hayAmount) returns(int256)
func (_Interaction *InteractionSession) Payback(token common.Address, hayAmount *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.Payback(&_Interaction.TransactOpts, token, hayAmount)
}

// Payback is a paid mutator transaction binding the contract method 0x35ed8ab8.
//
// Solidity: function payback(address token, uint256 hayAmount) returns(int256)
func (_Interaction *InteractionTransactorSession) Payback(token common.Address, hayAmount *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.Payback(&_Interaction.TransactOpts, token, hayAmount)
}

// PaybackFor is a paid mutator transaction binding the contract method 0x3b862af1.
//
// Solidity: function paybackFor(address token, uint256 hayAmount, address borrower) returns(int256)
func (_Interaction *InteractionTransactor) PaybackFor(opts *bind.TransactOpts, token common.Address, hayAmount *big.Int, borrower common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "paybackFor", token, hayAmount, borrower)
}

// PaybackFor is a paid mutator transaction binding the contract method 0x3b862af1.
//
// Solidity: function paybackFor(address token, uint256 hayAmount, address borrower) returns(int256)
func (_Interaction *InteractionSession) PaybackFor(token common.Address, hayAmount *big.Int, borrower common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.PaybackFor(&_Interaction.TransactOpts, token, hayAmount, borrower)
}

// PaybackFor is a paid mutator transaction binding the contract method 0x3b862af1.
//
// Solidity: function paybackFor(address token, uint256 hayAmount, address borrower) returns(int256)
func (_Interaction *InteractionTransactorSession) PaybackFor(token common.Address, hayAmount *big.Int, borrower common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.PaybackFor(&_Interaction.TransactOpts, token, hayAmount, borrower)
}

// Poke is a paid mutator transaction binding the contract method 0xb1a997ac.
//
// Solidity: function poke(address token) returns()
func (_Interaction *InteractionTransactor) Poke(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "poke", token)
}

// Poke is a paid mutator transaction binding the contract method 0xb1a997ac.
//
// Solidity: function poke(address token) returns()
func (_Interaction *InteractionSession) Poke(token common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Poke(&_Interaction.TransactOpts, token)
}

// Poke is a paid mutator transaction binding the contract method 0xb1a997ac.
//
// Solidity: function poke(address token) returns()
func (_Interaction *InteractionTransactorSession) Poke(token common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Poke(&_Interaction.TransactOpts, token)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Interaction *InteractionTransactor) Rely(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "rely", usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Interaction *InteractionSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Rely(&_Interaction.TransactOpts, usr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address usr) returns()
func (_Interaction *InteractionTransactorSession) Rely(usr common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.Rely(&_Interaction.TransactOpts, usr)
}

// RemoveCollateralType is a paid mutator transaction binding the contract method 0xcba315bd.
//
// Solidity: function removeCollateralType(address token) returns()
func (_Interaction *InteractionTransactor) RemoveCollateralType(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "removeCollateralType", token)
}

// RemoveCollateralType is a paid mutator transaction binding the contract method 0xcba315bd.
//
// Solidity: function removeCollateralType(address token) returns()
func (_Interaction *InteractionSession) RemoveCollateralType(token common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.RemoveCollateralType(&_Interaction.TransactOpts, token)
}

// RemoveCollateralType is a paid mutator transaction binding the contract method 0xcba315bd.
//
// Solidity: function removeCollateralType(address token) returns()
func (_Interaction *InteractionTransactorSession) RemoveCollateralType(token common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.RemoveCollateralType(&_Interaction.TransactOpts, token)
}

// RemoveFromAuctionWhitelist is a paid mutator transaction binding the contract method 0x0d2095ac.
//
// Solidity: function removeFromAuctionWhitelist(address[] usrs) returns()
func (_Interaction *InteractionTransactor) RemoveFromAuctionWhitelist(opts *bind.TransactOpts, usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "removeFromAuctionWhitelist", usrs)
}

// RemoveFromAuctionWhitelist is a paid mutator transaction binding the contract method 0x0d2095ac.
//
// Solidity: function removeFromAuctionWhitelist(address[] usrs) returns()
func (_Interaction *InteractionSession) RemoveFromAuctionWhitelist(usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.RemoveFromAuctionWhitelist(&_Interaction.TransactOpts, usrs)
}

// RemoveFromAuctionWhitelist is a paid mutator transaction binding the contract method 0x0d2095ac.
//
// Solidity: function removeFromAuctionWhitelist(address[] usrs) returns()
func (_Interaction *InteractionTransactorSession) RemoveFromAuctionWhitelist(usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.RemoveFromAuctionWhitelist(&_Interaction.TransactOpts, usrs)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] tokens) returns()
func (_Interaction *InteractionTransactor) RemoveFromBlacklist(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "removeFromBlacklist", tokens)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] tokens) returns()
func (_Interaction *InteractionSession) RemoveFromBlacklist(tokens []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.RemoveFromBlacklist(&_Interaction.TransactOpts, tokens)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x89daf799.
//
// Solidity: function removeFromBlacklist(address[] tokens) returns()
func (_Interaction *InteractionTransactorSession) RemoveFromBlacklist(tokens []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.RemoveFromBlacklist(&_Interaction.TransactOpts, tokens)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] usrs) returns()
func (_Interaction *InteractionTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "removeFromWhitelist", usrs)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] usrs) returns()
func (_Interaction *InteractionSession) RemoveFromWhitelist(usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.RemoveFromWhitelist(&_Interaction.TransactOpts, usrs)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x548db174.
//
// Solidity: function removeFromWhitelist(address[] usrs) returns()
func (_Interaction *InteractionTransactorSession) RemoveFromWhitelist(usrs []common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.RemoveFromWhitelist(&_Interaction.TransactOpts, usrs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Interaction *InteractionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Interaction *InteractionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Interaction.Contract.RenounceOwnership(&_Interaction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Interaction *InteractionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Interaction.Contract.RenounceOwnership(&_Interaction.TransactOpts)
}

// ResetAuction is a paid mutator transaction binding the contract method 0x07251016.
//
// Solidity: function resetAuction(address token, uint256 auctionId, address keeper) returns()
func (_Interaction *InteractionTransactor) ResetAuction(opts *bind.TransactOpts, token common.Address, auctionId *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "resetAuction", token, auctionId, keeper)
}

// ResetAuction is a paid mutator transaction binding the contract method 0x07251016.
//
// Solidity: function resetAuction(address token, uint256 auctionId, address keeper) returns()
func (_Interaction *InteractionSession) ResetAuction(token common.Address, auctionId *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.ResetAuction(&_Interaction.TransactOpts, token, auctionId, keeper)
}

// ResetAuction is a paid mutator transaction binding the contract method 0x07251016.
//
// Solidity: function resetAuction(address token, uint256 auctionId, address keeper) returns()
func (_Interaction *InteractionTransactorSession) ResetAuction(token common.Address, auctionId *big.Int, keeper common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.ResetAuction(&_Interaction.TransactOpts, token, auctionId, keeper)
}

// SetCollateralDuty is a paid mutator transaction binding the contract method 0x06fbed54.
//
// Solidity: function setCollateralDuty(address token, uint256 duty) returns()
func (_Interaction *InteractionTransactor) SetCollateralDuty(opts *bind.TransactOpts, token common.Address, duty *big.Int) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "setCollateralDuty", token, duty)
}

// SetCollateralDuty is a paid mutator transaction binding the contract method 0x06fbed54.
//
// Solidity: function setCollateralDuty(address token, uint256 duty) returns()
func (_Interaction *InteractionSession) SetCollateralDuty(token common.Address, duty *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.SetCollateralDuty(&_Interaction.TransactOpts, token, duty)
}

// SetCollateralDuty is a paid mutator transaction binding the contract method 0x06fbed54.
//
// Solidity: function setCollateralDuty(address token, uint256 duty) returns()
func (_Interaction *InteractionTransactorSession) SetCollateralDuty(token common.Address, duty *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.SetCollateralDuty(&_Interaction.TransactOpts, token, duty)
}

// SetCollateralType is a paid mutator transaction binding the contract method 0x2049ad8e.
//
// Solidity: function setCollateralType(address token, address gemJoin, bytes32 ilk, address clip, uint256 mat) returns()
func (_Interaction *InteractionTransactor) SetCollateralType(opts *bind.TransactOpts, token common.Address, gemJoin common.Address, ilk [32]byte, clip common.Address, mat *big.Int) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "setCollateralType", token, gemJoin, ilk, clip, mat)
}

// SetCollateralType is a paid mutator transaction binding the contract method 0x2049ad8e.
//
// Solidity: function setCollateralType(address token, address gemJoin, bytes32 ilk, address clip, uint256 mat) returns()
func (_Interaction *InteractionSession) SetCollateralType(token common.Address, gemJoin common.Address, ilk [32]byte, clip common.Address, mat *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.SetCollateralType(&_Interaction.TransactOpts, token, gemJoin, ilk, clip, mat)
}

// SetCollateralType is a paid mutator transaction binding the contract method 0x2049ad8e.
//
// Solidity: function setCollateralType(address token, address gemJoin, bytes32 ilk, address clip, uint256 mat) returns()
func (_Interaction *InteractionTransactorSession) SetCollateralType(token common.Address, gemJoin common.Address, ilk [32]byte, clip common.Address, mat *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.SetCollateralType(&_Interaction.TransactOpts, token, gemJoin, ilk, clip, mat)
}

// SetHelioProvider is a paid mutator transaction binding the contract method 0x7b7a4137.
//
// Solidity: function setHelioProvider(address token, address helioProvider, bool compatibilityMode) returns()
func (_Interaction *InteractionTransactor) SetHelioProvider(opts *bind.TransactOpts, token common.Address, helioProvider common.Address, compatibilityMode bool) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "setHelioProvider", token, helioProvider, compatibilityMode)
}

// SetHelioProvider is a paid mutator transaction binding the contract method 0x7b7a4137.
//
// Solidity: function setHelioProvider(address token, address helioProvider, bool compatibilityMode) returns()
func (_Interaction *InteractionSession) SetHelioProvider(token common.Address, helioProvider common.Address, compatibilityMode bool) (*types.Transaction, error) {
	return _Interaction.Contract.SetHelioProvider(&_Interaction.TransactOpts, token, helioProvider, compatibilityMode)
}

// SetHelioProvider is a paid mutator transaction binding the contract method 0x7b7a4137.
//
// Solidity: function setHelioProvider(address token, address helioProvider, bool compatibilityMode) returns()
func (_Interaction *InteractionTransactorSession) SetHelioProvider(token common.Address, helioProvider common.Address, compatibilityMode bool) (*types.Transaction, error) {
	return _Interaction.Contract.SetHelioProvider(&_Interaction.TransactOpts, token, helioProvider, compatibilityMode)
}

// SetListaDistributor is a paid mutator transaction binding the contract method 0x823a6c79.
//
// Solidity: function setListaDistributor(address distributor) returns()
func (_Interaction *InteractionTransactor) SetListaDistributor(opts *bind.TransactOpts, distributor common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "setListaDistributor", distributor)
}

// SetListaDistributor is a paid mutator transaction binding the contract method 0x823a6c79.
//
// Solidity: function setListaDistributor(address distributor) returns()
func (_Interaction *InteractionSession) SetListaDistributor(distributor common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.SetListaDistributor(&_Interaction.TransactOpts, distributor)
}

// SetListaDistributor is a paid mutator transaction binding the contract method 0x823a6c79.
//
// Solidity: function setListaDistributor(address distributor) returns()
func (_Interaction *InteractionTransactorSession) SetListaDistributor(distributor common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.SetListaDistributor(&_Interaction.TransactOpts, distributor)
}

// SetWhitelistOperator is a paid mutator transaction binding the contract method 0xeddaa0e9.
//
// Solidity: function setWhitelistOperator(address usr) returns()
func (_Interaction *InteractionTransactor) SetWhitelistOperator(opts *bind.TransactOpts, usr common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "setWhitelistOperator", usr)
}

// SetWhitelistOperator is a paid mutator transaction binding the contract method 0xeddaa0e9.
//
// Solidity: function setWhitelistOperator(address usr) returns()
func (_Interaction *InteractionSession) SetWhitelistOperator(usr common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.SetWhitelistOperator(&_Interaction.TransactOpts, usr)
}

// SetWhitelistOperator is a paid mutator transaction binding the contract method 0xeddaa0e9.
//
// Solidity: function setWhitelistOperator(address usr) returns()
func (_Interaction *InteractionTransactorSession) SetWhitelistOperator(usr common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.SetWhitelistOperator(&_Interaction.TransactOpts, usr)
}

// StartAuction is a paid mutator transaction binding the contract method 0x953790b1.
//
// Solidity: function startAuction(address token, address user, address keeper) returns(uint256)
func (_Interaction *InteractionTransactor) StartAuction(opts *bind.TransactOpts, token common.Address, user common.Address, keeper common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "startAuction", token, user, keeper)
}

// StartAuction is a paid mutator transaction binding the contract method 0x953790b1.
//
// Solidity: function startAuction(address token, address user, address keeper) returns(uint256)
func (_Interaction *InteractionSession) StartAuction(token common.Address, user common.Address, keeper common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.StartAuction(&_Interaction.TransactOpts, token, user, keeper)
}

// StartAuction is a paid mutator transaction binding the contract method 0x953790b1.
//
// Solidity: function startAuction(address token, address user, address keeper) returns(uint256)
func (_Interaction *InteractionTransactorSession) StartAuction(token common.Address, user common.Address, keeper common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.StartAuction(&_Interaction.TransactOpts, token, user, keeper)
}

// SyncSnapshot is a paid mutator transaction binding the contract method 0x0bce4a76.
//
// Solidity: function syncSnapshot(address token, address user) returns()
func (_Interaction *InteractionTransactor) SyncSnapshot(opts *bind.TransactOpts, token common.Address, user common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "syncSnapshot", token, user)
}

// SyncSnapshot is a paid mutator transaction binding the contract method 0x0bce4a76.
//
// Solidity: function syncSnapshot(address token, address user) returns()
func (_Interaction *InteractionSession) SyncSnapshot(token common.Address, user common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.SyncSnapshot(&_Interaction.TransactOpts, token, user)
}

// SyncSnapshot is a paid mutator transaction binding the contract method 0x0bce4a76.
//
// Solidity: function syncSnapshot(address token, address user) returns()
func (_Interaction *InteractionTransactorSession) SyncSnapshot(token common.Address, user common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.SyncSnapshot(&_Interaction.TransactOpts, token, user)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Interaction *InteractionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Interaction *InteractionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.TransferOwnership(&_Interaction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Interaction *InteractionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Interaction.Contract.TransferOwnership(&_Interaction.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address participant, address token, uint256 dink) returns(uint256)
func (_Interaction *InteractionTransactor) Withdraw(opts *bind.TransactOpts, participant common.Address, token common.Address, dink *big.Int) (*types.Transaction, error) {
	return _Interaction.contract.Transact(opts, "withdraw", participant, token, dink)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address participant, address token, uint256 dink) returns(uint256)
func (_Interaction *InteractionSession) Withdraw(participant common.Address, token common.Address, dink *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.Withdraw(&_Interaction.TransactOpts, participant, token, dink)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address participant, address token, uint256 dink) returns(uint256)
func (_Interaction *InteractionTransactorSession) Withdraw(participant common.Address, token common.Address, dink *big.Int) (*types.Transaction, error) {
	return _Interaction.Contract.Withdraw(&_Interaction.TransactOpts, participant, token, dink)
}

// InteractionAuctionFinishedIterator is returned from FilterAuctionFinished and is used to iterate over the raw logs and unpacked data for AuctionFinished events raised by the Interaction contract.
type InteractionAuctionFinishedIterator struct {
	Event *InteractionAuctionFinished // Event containing the contract specifics and raw log

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
func (it *InteractionAuctionFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionAuctionFinished)
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
		it.Event = new(InteractionAuctionFinished)
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
func (it *InteractionAuctionFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionAuctionFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionAuctionFinished represents a AuctionFinished event raised by the Interaction contract.
type InteractionAuctionFinished struct {
	Token  common.Address
	Keeper common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionFinished is a free log retrieval operation binding the contract event 0x88a732d7774b0799e041c44ae40c27a8616dd75d611a9be3cf1d082c759a0f8e.
//
// Solidity: event AuctionFinished(address indexed token, address keeper, uint256 amount)
func (_Interaction *InteractionFilterer) FilterAuctionFinished(opts *bind.FilterOpts, token []common.Address) (*InteractionAuctionFinishedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "AuctionFinished", tokenRule)
	if err != nil {
		return nil, err
	}
	return &InteractionAuctionFinishedIterator{contract: _Interaction.contract, event: "AuctionFinished", logs: logs, sub: sub}, nil
}

// WatchAuctionFinished is a free log subscription operation binding the contract event 0x88a732d7774b0799e041c44ae40c27a8616dd75d611a9be3cf1d082c759a0f8e.
//
// Solidity: event AuctionFinished(address indexed token, address keeper, uint256 amount)
func (_Interaction *InteractionFilterer) WatchAuctionFinished(opts *bind.WatchOpts, sink chan<- *InteractionAuctionFinished, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "AuctionFinished", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionAuctionFinished)
				if err := _Interaction.contract.UnpackLog(event, "AuctionFinished", log); err != nil {
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

// ParseAuctionFinished is a log parse operation binding the contract event 0x88a732d7774b0799e041c44ae40c27a8616dd75d611a9be3cf1d082c759a0f8e.
//
// Solidity: event AuctionFinished(address indexed token, address keeper, uint256 amount)
func (_Interaction *InteractionFilterer) ParseAuctionFinished(log types.Log) (*InteractionAuctionFinished, error) {
	event := new(InteractionAuctionFinished)
	if err := _Interaction.contract.UnpackLog(event, "AuctionFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionAuctionStartedIterator is returned from FilterAuctionStarted and is used to iterate over the raw logs and unpacked data for AuctionStarted events raised by the Interaction contract.
type InteractionAuctionStartedIterator struct {
	Event *InteractionAuctionStarted // Event containing the contract specifics and raw log

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
func (it *InteractionAuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionAuctionStarted)
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
		it.Event = new(InteractionAuctionStarted)
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
func (it *InteractionAuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionAuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionAuctionStarted represents a AuctionStarted event raised by the Interaction contract.
type InteractionAuctionStarted struct {
	Token  common.Address
	User   common.Address
	Amount *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionStarted is a free log retrieval operation binding the contract event 0xa39454816a8e1e974f3d7de0cc0ba57840ba7d4cf111f280726b1f687e91d42b.
//
// Solidity: event AuctionStarted(address indexed token, address user, uint256 amount, uint256 price)
func (_Interaction *InteractionFilterer) FilterAuctionStarted(opts *bind.FilterOpts, token []common.Address) (*InteractionAuctionStartedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "AuctionStarted", tokenRule)
	if err != nil {
		return nil, err
	}
	return &InteractionAuctionStartedIterator{contract: _Interaction.contract, event: "AuctionStarted", logs: logs, sub: sub}, nil
}

// WatchAuctionStarted is a free log subscription operation binding the contract event 0xa39454816a8e1e974f3d7de0cc0ba57840ba7d4cf111f280726b1f687e91d42b.
//
// Solidity: event AuctionStarted(address indexed token, address user, uint256 amount, uint256 price)
func (_Interaction *InteractionFilterer) WatchAuctionStarted(opts *bind.WatchOpts, sink chan<- *InteractionAuctionStarted, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "AuctionStarted", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionAuctionStarted)
				if err := _Interaction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
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

// ParseAuctionStarted is a log parse operation binding the contract event 0xa39454816a8e1e974f3d7de0cc0ba57840ba7d4cf111f280726b1f687e91d42b.
//
// Solidity: event AuctionStarted(address indexed token, address user, uint256 amount, uint256 price)
func (_Interaction *InteractionFilterer) ParseAuctionStarted(log types.Log) (*InteractionAuctionStarted, error) {
	event := new(InteractionAuctionStarted)
	if err := _Interaction.contract.UnpackLog(event, "AuctionStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Interaction contract.
type InteractionBorrowIterator struct {
	Event *InteractionBorrow // Event containing the contract specifics and raw log

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
func (it *InteractionBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionBorrow)
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
		it.Event = new(InteractionBorrow)
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
func (it *InteractionBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionBorrow represents a Borrow event raised by the Interaction contract.
type InteractionBorrow struct {
	User             common.Address
	Collateral       common.Address
	CollateralAmount *big.Int
	Amount           *big.Int
	LiquidationPrice *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x10a0132d3bf8c82a7fb93a86160f3074ca5c3e5706fa2bcdf0e2b5fd495af09b.
//
// Solidity: event Borrow(address indexed user, address collateral, uint256 collateralAmount, uint256 amount, uint256 liquidationPrice)
func (_Interaction *InteractionFilterer) FilterBorrow(opts *bind.FilterOpts, user []common.Address) (*InteractionBorrowIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "Borrow", userRule)
	if err != nil {
		return nil, err
	}
	return &InteractionBorrowIterator{contract: _Interaction.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x10a0132d3bf8c82a7fb93a86160f3074ca5c3e5706fa2bcdf0e2b5fd495af09b.
//
// Solidity: event Borrow(address indexed user, address collateral, uint256 collateralAmount, uint256 amount, uint256 liquidationPrice)
func (_Interaction *InteractionFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *InteractionBorrow, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "Borrow", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionBorrow)
				if err := _Interaction.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x10a0132d3bf8c82a7fb93a86160f3074ca5c3e5706fa2bcdf0e2b5fd495af09b.
//
// Solidity: event Borrow(address indexed user, address collateral, uint256 collateralAmount, uint256 amount, uint256 liquidationPrice)
func (_Interaction *InteractionFilterer) ParseBorrow(log types.Log) (*InteractionBorrow, error) {
	event := new(InteractionBorrow)
	if err := _Interaction.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionCollateralDisabledIterator is returned from FilterCollateralDisabled and is used to iterate over the raw logs and unpacked data for CollateralDisabled events raised by the Interaction contract.
type InteractionCollateralDisabledIterator struct {
	Event *InteractionCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *InteractionCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionCollateralDisabled)
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
		it.Event = new(InteractionCollateralDisabled)
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
func (it *InteractionCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionCollateralDisabled represents a CollateralDisabled event raised by the Interaction contract.
type InteractionCollateralDisabled struct {
	Token common.Address
	Ilk   [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCollateralDisabled is a free log retrieval operation binding the contract event 0x9fc793554ab8af2084d6005e4d61d847f73406e2f552476dfb2a0d4987ef4fa0.
//
// Solidity: event CollateralDisabled(address token, bytes32 ilk)
func (_Interaction *InteractionFilterer) FilterCollateralDisabled(opts *bind.FilterOpts) (*InteractionCollateralDisabledIterator, error) {

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "CollateralDisabled")
	if err != nil {
		return nil, err
	}
	return &InteractionCollateralDisabledIterator{contract: _Interaction.contract, event: "CollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchCollateralDisabled is a free log subscription operation binding the contract event 0x9fc793554ab8af2084d6005e4d61d847f73406e2f552476dfb2a0d4987ef4fa0.
//
// Solidity: event CollateralDisabled(address token, bytes32 ilk)
func (_Interaction *InteractionFilterer) WatchCollateralDisabled(opts *bind.WatchOpts, sink chan<- *InteractionCollateralDisabled) (event.Subscription, error) {

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "CollateralDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionCollateralDisabled)
				if err := _Interaction.contract.UnpackLog(event, "CollateralDisabled", log); err != nil {
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

// ParseCollateralDisabled is a log parse operation binding the contract event 0x9fc793554ab8af2084d6005e4d61d847f73406e2f552476dfb2a0d4987ef4fa0.
//
// Solidity: event CollateralDisabled(address token, bytes32 ilk)
func (_Interaction *InteractionFilterer) ParseCollateralDisabled(log types.Log) (*InteractionCollateralDisabled, error) {
	event := new(InteractionCollateralDisabled)
	if err := _Interaction.contract.UnpackLog(event, "CollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionCollateralEnabledIterator is returned from FilterCollateralEnabled and is used to iterate over the raw logs and unpacked data for CollateralEnabled events raised by the Interaction contract.
type InteractionCollateralEnabledIterator struct {
	Event *InteractionCollateralEnabled // Event containing the contract specifics and raw log

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
func (it *InteractionCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionCollateralEnabled)
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
		it.Event = new(InteractionCollateralEnabled)
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
func (it *InteractionCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionCollateralEnabled represents a CollateralEnabled event raised by the Interaction contract.
type InteractionCollateralEnabled struct {
	Token common.Address
	Ilk   [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCollateralEnabled is a free log retrieval operation binding the contract event 0x5d8d8db1d1b082fc4f893dcf587266b61ac1079e3f5e959cd6e76cfe84274f47.
//
// Solidity: event CollateralEnabled(address token, bytes32 ilk)
func (_Interaction *InteractionFilterer) FilterCollateralEnabled(opts *bind.FilterOpts) (*InteractionCollateralEnabledIterator, error) {

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "CollateralEnabled")
	if err != nil {
		return nil, err
	}
	return &InteractionCollateralEnabledIterator{contract: _Interaction.contract, event: "CollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchCollateralEnabled is a free log subscription operation binding the contract event 0x5d8d8db1d1b082fc4f893dcf587266b61ac1079e3f5e959cd6e76cfe84274f47.
//
// Solidity: event CollateralEnabled(address token, bytes32 ilk)
func (_Interaction *InteractionFilterer) WatchCollateralEnabled(opts *bind.WatchOpts, sink chan<- *InteractionCollateralEnabled) (event.Subscription, error) {

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "CollateralEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionCollateralEnabled)
				if err := _Interaction.contract.UnpackLog(event, "CollateralEnabled", log); err != nil {
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

// ParseCollateralEnabled is a log parse operation binding the contract event 0x5d8d8db1d1b082fc4f893dcf587266b61ac1079e3f5e959cd6e76cfe84274f47.
//
// Solidity: event CollateralEnabled(address token, bytes32 ilk)
func (_Interaction *InteractionFilterer) ParseCollateralEnabled(log types.Log) (*InteractionCollateralEnabled, error) {
	event := new(InteractionCollateralEnabled)
	if err := _Interaction.contract.UnpackLog(event, "CollateralEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Interaction contract.
type InteractionDepositIterator struct {
	Event *InteractionDeposit // Event containing the contract specifics and raw log

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
func (it *InteractionDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionDeposit)
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
		it.Event = new(InteractionDeposit)
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
func (it *InteractionDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionDeposit represents a Deposit event raised by the Interaction contract.
type InteractionDeposit struct {
	User        common.Address
	Collateral  common.Address
	Amount      *big.Int
	TotalAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed user, address collateral, uint256 amount, uint256 totalAmount)
func (_Interaction *InteractionFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*InteractionDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &InteractionDepositIterator{contract: _Interaction.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed user, address collateral, uint256 amount, uint256 totalAmount)
func (_Interaction *InteractionFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *InteractionDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionDeposit)
				if err := _Interaction.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed user, address collateral, uint256 amount, uint256 totalAmount)
func (_Interaction *InteractionFilterer) ParseDeposit(log types.Log) (*InteractionDeposit, error) {
	event := new(InteractionDeposit)
	if err := _Interaction.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Interaction contract.
type InteractionInitializedIterator struct {
	Event *InteractionInitialized // Event containing the contract specifics and raw log

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
func (it *InteractionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionInitialized)
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
		it.Event = new(InteractionInitialized)
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
func (it *InteractionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionInitialized represents a Initialized event raised by the Interaction contract.
type InteractionInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Interaction *InteractionFilterer) FilterInitialized(opts *bind.FilterOpts) (*InteractionInitializedIterator, error) {

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InteractionInitializedIterator{contract: _Interaction.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Interaction *InteractionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InteractionInitialized) (event.Subscription, error) {

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionInitialized)
				if err := _Interaction.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Interaction *InteractionFilterer) ParseInitialized(log types.Log) (*InteractionInitialized, error) {
	event := new(InteractionInitialized)
	if err := _Interaction.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionLiquidationIterator is returned from FilterLiquidation and is used to iterate over the raw logs and unpacked data for Liquidation events raised by the Interaction contract.
type InteractionLiquidationIterator struct {
	Event *InteractionLiquidation // Event containing the contract specifics and raw log

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
func (it *InteractionLiquidationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionLiquidation)
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
		it.Event = new(InteractionLiquidation)
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
func (it *InteractionLiquidationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionLiquidationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionLiquidation represents a Liquidation event raised by the Interaction contract.
type InteractionLiquidation struct {
	User       common.Address
	Collateral common.Address
	Amount     *big.Int
	Leftover   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLiquidation is a free log retrieval operation binding the contract event 0x4ecae3269f800df64b16cb9f6f8b0b507018888521d1cff0841823e44bc0b00d.
//
// Solidity: event Liquidation(address indexed user, address indexed collateral, uint256 amount, uint256 leftover)
func (_Interaction *InteractionFilterer) FilterLiquidation(opts *bind.FilterOpts, user []common.Address, collateral []common.Address) (*InteractionLiquidationIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "Liquidation", userRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return &InteractionLiquidationIterator{contract: _Interaction.contract, event: "Liquidation", logs: logs, sub: sub}, nil
}

// WatchLiquidation is a free log subscription operation binding the contract event 0x4ecae3269f800df64b16cb9f6f8b0b507018888521d1cff0841823e44bc0b00d.
//
// Solidity: event Liquidation(address indexed user, address indexed collateral, uint256 amount, uint256 leftover)
func (_Interaction *InteractionFilterer) WatchLiquidation(opts *bind.WatchOpts, sink chan<- *InteractionLiquidation, user []common.Address, collateral []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "Liquidation", userRule, collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionLiquidation)
				if err := _Interaction.contract.UnpackLog(event, "Liquidation", log); err != nil {
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

// ParseLiquidation is a log parse operation binding the contract event 0x4ecae3269f800df64b16cb9f6f8b0b507018888521d1cff0841823e44bc0b00d.
//
// Solidity: event Liquidation(address indexed user, address indexed collateral, uint256 amount, uint256 leftover)
func (_Interaction *InteractionFilterer) ParseLiquidation(log types.Log) (*InteractionLiquidation, error) {
	event := new(InteractionLiquidation)
	if err := _Interaction.contract.UnpackLog(event, "Liquidation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Interaction contract.
type InteractionOwnershipTransferredIterator struct {
	Event *InteractionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *InteractionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionOwnershipTransferred)
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
		it.Event = new(InteractionOwnershipTransferred)
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
func (it *InteractionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionOwnershipTransferred represents a OwnershipTransferred event raised by the Interaction contract.
type InteractionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Interaction *InteractionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*InteractionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &InteractionOwnershipTransferredIterator{contract: _Interaction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Interaction *InteractionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *InteractionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionOwnershipTransferred)
				if err := _Interaction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Interaction *InteractionFilterer) ParseOwnershipTransferred(log types.Log) (*InteractionOwnershipTransferred, error) {
	event := new(InteractionOwnershipTransferred)
	if err := _Interaction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionPaybackIterator is returned from FilterPayback and is used to iterate over the raw logs and unpacked data for Payback events raised by the Interaction contract.
type InteractionPaybackIterator struct {
	Event *InteractionPayback // Event containing the contract specifics and raw log

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
func (it *InteractionPaybackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionPayback)
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
		it.Event = new(InteractionPayback)
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
func (it *InteractionPaybackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionPaybackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionPayback represents a Payback event raised by the Interaction contract.
type InteractionPayback struct {
	User             common.Address
	Collateral       common.Address
	Amount           *big.Int
	Debt             *big.Int
	LiquidationPrice *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPayback is a free log retrieval operation binding the contract event 0xee7777293f0d992d01f7d63f6f0e7762ee0df9d22267140e41090174a70e6f72.
//
// Solidity: event Payback(address indexed user, address collateral, uint256 amount, uint256 debt, uint256 liquidationPrice)
func (_Interaction *InteractionFilterer) FilterPayback(opts *bind.FilterOpts, user []common.Address) (*InteractionPaybackIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "Payback", userRule)
	if err != nil {
		return nil, err
	}
	return &InteractionPaybackIterator{contract: _Interaction.contract, event: "Payback", logs: logs, sub: sub}, nil
}

// WatchPayback is a free log subscription operation binding the contract event 0xee7777293f0d992d01f7d63f6f0e7762ee0df9d22267140e41090174a70e6f72.
//
// Solidity: event Payback(address indexed user, address collateral, uint256 amount, uint256 debt, uint256 liquidationPrice)
func (_Interaction *InteractionFilterer) WatchPayback(opts *bind.WatchOpts, sink chan<- *InteractionPayback, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "Payback", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionPayback)
				if err := _Interaction.contract.UnpackLog(event, "Payback", log); err != nil {
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

// ParsePayback is a log parse operation binding the contract event 0xee7777293f0d992d01f7d63f6f0e7762ee0df9d22267140e41090174a70e6f72.
//
// Solidity: event Payback(address indexed user, address collateral, uint256 amount, uint256 debt, uint256 liquidationPrice)
func (_Interaction *InteractionFilterer) ParsePayback(log types.Log) (*InteractionPayback, error) {
	event := new(InteractionPayback)
	if err := _Interaction.contract.UnpackLog(event, "Payback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InteractionWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Interaction contract.
type InteractionWithdrawIterator struct {
	Event *InteractionWithdraw // Event containing the contract specifics and raw log

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
func (it *InteractionWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InteractionWithdraw)
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
		it.Event = new(InteractionWithdraw)
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
func (it *InteractionWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InteractionWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InteractionWithdraw represents a Withdraw event raised by the Interaction contract.
type InteractionWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Interaction *InteractionFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*InteractionWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Interaction.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &InteractionWithdrawIterator{contract: _Interaction.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Interaction *InteractionFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *InteractionWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Interaction.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InteractionWithdraw)
				if err := _Interaction.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Interaction *InteractionFilterer) ParseWithdraw(log types.Log) (*InteractionWithdraw, error) {
	event := new(InteractionWithdraw)
	if err := _Interaction.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
