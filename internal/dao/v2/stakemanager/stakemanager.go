// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakemanager

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

// IStakeManagerBotUndelegateRequest is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerBotUndelegateRequest struct {
	StartTime     *big.Int
	EndTime       *big.Int
	Amount        *big.Int
	AmountInSnBnb *big.Int
}

// IStakeManagerWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerWithdrawalRequest struct {
	Uuid          *big.Int
	AmountInSnBnb *big.Int
	StartTime     *big.Int
}

// StakemanagerMetaData contains all meta data concerning the Stakemanager contract.
var StakemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ClaimAllWithdrawals\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_uuid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ClaimUndelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_uuid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ClaimUndelegatedFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ClaimWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Delegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_delegateVotePower\",\"type\":\"bool\"}],\"name\":\"DelegateTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_delegateTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_votesChange\",\"type\":\"uint256\"}],\"name\":\"DelegateVoteTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"DisableValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"ProposeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_src\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_dest\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"ReDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardsId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Redelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bnbAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_dailySlisBnb\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_days\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_remainingSlisBnb\",\"type\":\"uint256\"}],\"name\":\"RefundCommission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"RemoveValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountInSlisBnb\",\"type\":\"uint256\"}],\"name\":\"RequestWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"RewardsCompounded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_annualRate\",\"type\":\"uint256\"}],\"name\":\"SetAnnualRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"SetBSCValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"SetManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minBnb\",\"type\":\"uint256\"}],\"name\":\"SetMinBnb\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"SetRedirectAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"SetReserveAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"SetRevenuePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_synFee\",\"type\":\"uint256\"}],\"name\":\"SetSynFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_credit\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"toRemove\",\"type\":\"bool\"}],\"name\":\"SyncCreditContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nextUndelegatedRequestIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bnbAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"Undelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_bnbAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"UndelegateFrom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateReserve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"WhitelistValidator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARDIAN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEN_DECIMALS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptNewManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"amountToDelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"annualRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bnbAmount\",\"type\":\"uint256\"}],\"name\":\"binarySearchCoveredMaxIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"claimUndelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_uuid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"claimWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"claimWithdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"compoundRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bnbAmount\",\"type\":\"uint256\"}],\"name\":\"convertBnbToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"convertBnbToSnBnb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertSharesToBnb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInSlisBnb\",\"type\":\"uint256\"}],\"name\":\"convertSnBnbToBnb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creditContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditStates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"delegateTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegateVotePower\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_delegateTo\",\"type\":\"address\"}],\"name\":\"delegateVoteTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositReserve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"disableValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAmountToUndelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountToUndelegate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_uuid\",\"type\":\"uint256\"}],\"name\":\"getBotUndelegateRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInSnBnb\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.BotUndelegateRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getClaimableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_slisBnb\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bscValidator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"getDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getRedelegateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlisBnbWithdrawLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_slisBnbWithdrawLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBnbInValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledBnb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"getUserRequestStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isClaimable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserWithdrawalRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"uuid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInSnBnb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.WithdrawalRequest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slisBnb\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_synFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_revenuePool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBnb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextConfirmedRequestUUID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"placeholder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"proposeNewManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcValidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstValidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redirectAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dailySlisBnb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingSlisBnb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastBurnTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_days\",\"type\":\"uint256\"}],\"name\":\"refundCommission\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestIndexMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestUUID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInSlisBnb\",\"type\":\"uint256\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserveAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revenuePool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"revokeBotRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_annualRate\",\"type\":\"uint256\"}],\"name\":\"setAnnualRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setBSCValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setBotRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setMinBnb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRedirectAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setReserveAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRevenuePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_synFee\",\"type\":\"uint256\"}],\"name\":\"setSynFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"togglePause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"toggleVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalReserveAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unbondingBnb\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undelegate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_uuid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"undelegateFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_actualBnbAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undelegatedQuota\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"whitelistValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakemanagerMetaData.ABI instead.
var StakemanagerABI = StakemanagerMetaData.ABI

// Stakemanager is an auto generated Go binding around an Ethereum contract.
type Stakemanager struct {
	StakemanagerCaller     // Read-only binding to the contract
	StakemanagerTransactor // Write-only binding to the contract
	StakemanagerFilterer   // Log filterer for contract events
}

// StakemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakemanagerSession struct {
	Contract     *Stakemanager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakemanagerCallerSession struct {
	Contract *StakemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakemanagerTransactorSession struct {
	Contract     *StakemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakemanagerRaw struct {
	Contract *Stakemanager // Generic contract binding to access the raw methods on
}

// StakemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakemanagerCallerRaw struct {
	Contract *StakemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakemanagerTransactorRaw struct {
	Contract *StakemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakemanager creates a new instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanager(address common.Address, backend bind.ContractBackend) (*Stakemanager, error) {
	contract, err := bindStakemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakemanager{StakemanagerCaller: StakemanagerCaller{contract: contract}, StakemanagerTransactor: StakemanagerTransactor{contract: contract}, StakemanagerFilterer: StakemanagerFilterer{contract: contract}}, nil
}

// NewStakemanagerCaller creates a new read-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerCaller(address common.Address, caller bind.ContractCaller) (*StakemanagerCaller, error) {
	contract, err := bindStakemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerCaller{contract: contract}, nil
}

// NewStakemanagerTransactor creates a new write-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakemanagerTransactor, error) {
	contract, err := bindStakemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerTransactor{contract: contract}, nil
}

// NewStakemanagerFilterer creates a new log filterer instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakemanagerFilterer, error) {
	contract, err := bindStakemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakemanagerFilterer{contract: contract}, nil
}

// bindStakemanager binds a generic wrapper to an already deployed contract.
func bindStakemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakemanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.StakemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transact(opts, method, params...)
}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Stakemanager *StakemanagerCaller) BOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "BOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Stakemanager *StakemanagerSession) BOT() ([32]byte, error) {
	return _Stakemanager.Contract.BOT(&_Stakemanager.CallOpts)
}

// BOT is a free data retrieval call binding the contract method 0x486277f6.
//
// Solidity: function BOT() view returns(bytes32)
func (_Stakemanager *StakemanagerCallerSession) BOT() ([32]byte, error) {
	return _Stakemanager.Contract.BOT(&_Stakemanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stakemanager *StakemanagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stakemanager *StakemanagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stakemanager.Contract.DEFAULTADMINROLE(&_Stakemanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stakemanager *StakemanagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stakemanager.Contract.DEFAULTADMINROLE(&_Stakemanager.CallOpts)
}

// GUARDIAN is a free data retrieval call binding the contract method 0x724c184c.
//
// Solidity: function GUARDIAN() view returns(bytes32)
func (_Stakemanager *StakemanagerCaller) GUARDIAN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "GUARDIAN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDIAN is a free data retrieval call binding the contract method 0x724c184c.
//
// Solidity: function GUARDIAN() view returns(bytes32)
func (_Stakemanager *StakemanagerSession) GUARDIAN() ([32]byte, error) {
	return _Stakemanager.Contract.GUARDIAN(&_Stakemanager.CallOpts)
}

// GUARDIAN is a free data retrieval call binding the contract method 0x724c184c.
//
// Solidity: function GUARDIAN() view returns(bytes32)
func (_Stakemanager *StakemanagerCallerSession) GUARDIAN() ([32]byte, error) {
	return _Stakemanager.Contract.GUARDIAN(&_Stakemanager.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Stakemanager *StakemanagerCaller) MANAGER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "MANAGER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Stakemanager *StakemanagerSession) MANAGER() ([32]byte, error) {
	return _Stakemanager.Contract.MANAGER(&_Stakemanager.CallOpts)
}

// MANAGER is a free data retrieval call binding the contract method 0x1b2df850.
//
// Solidity: function MANAGER() view returns(bytes32)
func (_Stakemanager *StakemanagerCallerSession) MANAGER() ([32]byte, error) {
	return _Stakemanager.Contract.MANAGER(&_Stakemanager.CallOpts)
}

// TENDECIMALS is a free data retrieval call binding the contract method 0x5d499b1b.
//
// Solidity: function TEN_DECIMALS() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) TENDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "TEN_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TENDECIMALS is a free data retrieval call binding the contract method 0x5d499b1b.
//
// Solidity: function TEN_DECIMALS() view returns(uint256)
func (_Stakemanager *StakemanagerSession) TENDECIMALS() (*big.Int, error) {
	return _Stakemanager.Contract.TENDECIMALS(&_Stakemanager.CallOpts)
}

// TENDECIMALS is a free data retrieval call binding the contract method 0x5d499b1b.
//
// Solidity: function TEN_DECIMALS() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) TENDECIMALS() (*big.Int, error) {
	return _Stakemanager.Contract.TENDECIMALS(&_Stakemanager.CallOpts)
}

// AmountToDelegate is a free data retrieval call binding the contract method 0x2401c68b.
//
// Solidity: function amountToDelegate() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) AmountToDelegate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "amountToDelegate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountToDelegate is a free data retrieval call binding the contract method 0x2401c68b.
//
// Solidity: function amountToDelegate() view returns(uint256)
func (_Stakemanager *StakemanagerSession) AmountToDelegate() (*big.Int, error) {
	return _Stakemanager.Contract.AmountToDelegate(&_Stakemanager.CallOpts)
}

// AmountToDelegate is a free data retrieval call binding the contract method 0x2401c68b.
//
// Solidity: function amountToDelegate() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) AmountToDelegate() (*big.Int, error) {
	return _Stakemanager.Contract.AmountToDelegate(&_Stakemanager.CallOpts)
}

// AnnualRate is a free data retrieval call binding the contract method 0xc0f9b0fb.
//
// Solidity: function annualRate() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) AnnualRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "annualRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AnnualRate is a free data retrieval call binding the contract method 0xc0f9b0fb.
//
// Solidity: function annualRate() view returns(uint256)
func (_Stakemanager *StakemanagerSession) AnnualRate() (*big.Int, error) {
	return _Stakemanager.Contract.AnnualRate(&_Stakemanager.CallOpts)
}

// AnnualRate is a free data retrieval call binding the contract method 0xc0f9b0fb.
//
// Solidity: function annualRate() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) AnnualRate() (*big.Int, error) {
	return _Stakemanager.Contract.AnnualRate(&_Stakemanager.CallOpts)
}

// BinarySearchCoveredMaxIndex is a free data retrieval call binding the contract method 0x65abfc7b.
//
// Solidity: function binarySearchCoveredMaxIndex(uint256 _bnbAmount) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) BinarySearchCoveredMaxIndex(opts *bind.CallOpts, _bnbAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "binarySearchCoveredMaxIndex", _bnbAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BinarySearchCoveredMaxIndex is a free data retrieval call binding the contract method 0x65abfc7b.
//
// Solidity: function binarySearchCoveredMaxIndex(uint256 _bnbAmount) view returns(uint256)
func (_Stakemanager *StakemanagerSession) BinarySearchCoveredMaxIndex(_bnbAmount *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.BinarySearchCoveredMaxIndex(&_Stakemanager.CallOpts, _bnbAmount)
}

// BinarySearchCoveredMaxIndex is a free data retrieval call binding the contract method 0x65abfc7b.
//
// Solidity: function binarySearchCoveredMaxIndex(uint256 _bnbAmount) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) BinarySearchCoveredMaxIndex(_bnbAmount *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.BinarySearchCoveredMaxIndex(&_Stakemanager.CallOpts, _bnbAmount)
}

// ConvertBnbToShares is a free data retrieval call binding the contract method 0xe0a69d0a.
//
// Solidity: function convertBnbToShares(address _operator, uint256 _bnbAmount) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) ConvertBnbToShares(opts *bind.CallOpts, _operator common.Address, _bnbAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "convertBnbToShares", _operator, _bnbAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertBnbToShares is a free data retrieval call binding the contract method 0xe0a69d0a.
//
// Solidity: function convertBnbToShares(address _operator, uint256 _bnbAmount) view returns(uint256)
func (_Stakemanager *StakemanagerSession) ConvertBnbToShares(_operator common.Address, _bnbAmount *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.ConvertBnbToShares(&_Stakemanager.CallOpts, _operator, _bnbAmount)
}

// ConvertBnbToShares is a free data retrieval call binding the contract method 0xe0a69d0a.
//
// Solidity: function convertBnbToShares(address _operator, uint256 _bnbAmount) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) ConvertBnbToShares(_operator common.Address, _bnbAmount *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.ConvertBnbToShares(&_Stakemanager.CallOpts, _operator, _bnbAmount)
}

// ConvertBnbToSnBnb is a free data retrieval call binding the contract method 0x29d8a2eb.
//
// Solidity: function convertBnbToSnBnb(uint256 _amount) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) ConvertBnbToSnBnb(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "convertBnbToSnBnb", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertBnbToSnBnb is a free data retrieval call binding the contract method 0x29d8a2eb.
//
// Solidity: function convertBnbToSnBnb(uint256 _amount) view returns(uint256)
func (_Stakemanager *StakemanagerSession) ConvertBnbToSnBnb(_amount *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.ConvertBnbToSnBnb(&_Stakemanager.CallOpts, _amount)
}

// ConvertBnbToSnBnb is a free data retrieval call binding the contract method 0x29d8a2eb.
//
// Solidity: function convertBnbToSnBnb(uint256 _amount) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) ConvertBnbToSnBnb(_amount *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.ConvertBnbToSnBnb(&_Stakemanager.CallOpts, _amount)
}

// ConvertSharesToBnb is a free data retrieval call binding the contract method 0xea756a44.
//
// Solidity: function convertSharesToBnb(address _operator, uint256 _shares) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) ConvertSharesToBnb(opts *bind.CallOpts, _operator common.Address, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "convertSharesToBnb", _operator, _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertSharesToBnb is a free data retrieval call binding the contract method 0xea756a44.
//
// Solidity: function convertSharesToBnb(address _operator, uint256 _shares) view returns(uint256)
func (_Stakemanager *StakemanagerSession) ConvertSharesToBnb(_operator common.Address, _shares *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.ConvertSharesToBnb(&_Stakemanager.CallOpts, _operator, _shares)
}

// ConvertSharesToBnb is a free data retrieval call binding the contract method 0xea756a44.
//
// Solidity: function convertSharesToBnb(address _operator, uint256 _shares) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) ConvertSharesToBnb(_operator common.Address, _shares *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.ConvertSharesToBnb(&_Stakemanager.CallOpts, _operator, _shares)
}

// ConvertSnBnbToBnb is a free data retrieval call binding the contract method 0xce6298e1.
//
// Solidity: function convertSnBnbToBnb(uint256 _amountInSlisBnb) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) ConvertSnBnbToBnb(opts *bind.CallOpts, _amountInSlisBnb *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "convertSnBnbToBnb", _amountInSlisBnb)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertSnBnbToBnb is a free data retrieval call binding the contract method 0xce6298e1.
//
// Solidity: function convertSnBnbToBnb(uint256 _amountInSlisBnb) view returns(uint256)
func (_Stakemanager *StakemanagerSession) ConvertSnBnbToBnb(_amountInSlisBnb *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.ConvertSnBnbToBnb(&_Stakemanager.CallOpts, _amountInSlisBnb)
}

// ConvertSnBnbToBnb is a free data retrieval call binding the contract method 0xce6298e1.
//
// Solidity: function convertSnBnbToBnb(uint256 _amountInSlisBnb) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) ConvertSnBnbToBnb(_amountInSlisBnb *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.ConvertSnBnbToBnb(&_Stakemanager.CallOpts, _amountInSlisBnb)
}

// CreditContracts is a free data retrieval call binding the contract method 0xe5f61c1c.
//
// Solidity: function creditContracts(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCaller) CreditContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "creditContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditContracts is a free data retrieval call binding the contract method 0xe5f61c1c.
//
// Solidity: function creditContracts(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerSession) CreditContracts(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.CreditContracts(&_Stakemanager.CallOpts, arg0)
}

// CreditContracts is a free data retrieval call binding the contract method 0xe5f61c1c.
//
// Solidity: function creditContracts(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) CreditContracts(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.CreditContracts(&_Stakemanager.CallOpts, arg0)
}

// CreditStates is a free data retrieval call binding the contract method 0x0dd98027.
//
// Solidity: function creditStates(address ) view returns(bool)
func (_Stakemanager *StakemanagerCaller) CreditStates(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "creditStates", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreditStates is a free data retrieval call binding the contract method 0x0dd98027.
//
// Solidity: function creditStates(address ) view returns(bool)
func (_Stakemanager *StakemanagerSession) CreditStates(arg0 common.Address) (bool, error) {
	return _Stakemanager.Contract.CreditStates(&_Stakemanager.CallOpts, arg0)
}

// CreditStates is a free data retrieval call binding the contract method 0x0dd98027.
//
// Solidity: function creditStates(address ) view returns(bool)
func (_Stakemanager *StakemanagerCallerSession) CreditStates(arg0 common.Address) (bool, error) {
	return _Stakemanager.Contract.CreditStates(&_Stakemanager.CallOpts, arg0)
}

// DelegateVotePower is a free data retrieval call binding the contract method 0x6aaaf6e5.
//
// Solidity: function delegateVotePower() view returns(bool)
func (_Stakemanager *StakemanagerCaller) DelegateVotePower(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "delegateVotePower")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DelegateVotePower is a free data retrieval call binding the contract method 0x6aaaf6e5.
//
// Solidity: function delegateVotePower() view returns(bool)
func (_Stakemanager *StakemanagerSession) DelegateVotePower() (bool, error) {
	return _Stakemanager.Contract.DelegateVotePower(&_Stakemanager.CallOpts)
}

// DelegateVotePower is a free data retrieval call binding the contract method 0x6aaaf6e5.
//
// Solidity: function delegateVotePower() view returns(bool)
func (_Stakemanager *StakemanagerCallerSession) DelegateVotePower() (bool, error) {
	return _Stakemanager.Contract.DelegateVotePower(&_Stakemanager.CallOpts)
}

// GetAmountToUndelegate is a free data retrieval call binding the contract method 0x77ab033d.
//
// Solidity: function getAmountToUndelegate() view returns(uint256 _amountToUndelegate)
func (_Stakemanager *StakemanagerCaller) GetAmountToUndelegate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getAmountToUndelegate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountToUndelegate is a free data retrieval call binding the contract method 0x77ab033d.
//
// Solidity: function getAmountToUndelegate() view returns(uint256 _amountToUndelegate)
func (_Stakemanager *StakemanagerSession) GetAmountToUndelegate() (*big.Int, error) {
	return _Stakemanager.Contract.GetAmountToUndelegate(&_Stakemanager.CallOpts)
}

// GetAmountToUndelegate is a free data retrieval call binding the contract method 0x77ab033d.
//
// Solidity: function getAmountToUndelegate() view returns(uint256 _amountToUndelegate)
func (_Stakemanager *StakemanagerCallerSession) GetAmountToUndelegate() (*big.Int, error) {
	return _Stakemanager.Contract.GetAmountToUndelegate(&_Stakemanager.CallOpts)
}

// GetBotUndelegateRequest is a free data retrieval call binding the contract method 0xca13fb41.
//
// Solidity: function getBotUndelegateRequest(uint256 _uuid) view returns((uint256,uint256,uint256,uint256))
func (_Stakemanager *StakemanagerCaller) GetBotUndelegateRequest(opts *bind.CallOpts, _uuid *big.Int) (IStakeManagerBotUndelegateRequest, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getBotUndelegateRequest", _uuid)

	if err != nil {
		return *new(IStakeManagerBotUndelegateRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IStakeManagerBotUndelegateRequest)).(*IStakeManagerBotUndelegateRequest)

	return out0, err

}

// GetBotUndelegateRequest is a free data retrieval call binding the contract method 0xca13fb41.
//
// Solidity: function getBotUndelegateRequest(uint256 _uuid) view returns((uint256,uint256,uint256,uint256))
func (_Stakemanager *StakemanagerSession) GetBotUndelegateRequest(_uuid *big.Int) (IStakeManagerBotUndelegateRequest, error) {
	return _Stakemanager.Contract.GetBotUndelegateRequest(&_Stakemanager.CallOpts, _uuid)
}

// GetBotUndelegateRequest is a free data retrieval call binding the contract method 0xca13fb41.
//
// Solidity: function getBotUndelegateRequest(uint256 _uuid) view returns((uint256,uint256,uint256,uint256))
func (_Stakemanager *StakemanagerCallerSession) GetBotUndelegateRequest(_uuid *big.Int) (IStakeManagerBotUndelegateRequest, error) {
	return _Stakemanager.Contract.GetBotUndelegateRequest(&_Stakemanager.CallOpts, _uuid)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address _validator) view returns(uint256 _amount)
func (_Stakemanager *StakemanagerCaller) GetClaimableAmount(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getClaimableAmount", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address _validator) view returns(uint256 _amount)
func (_Stakemanager *StakemanagerSession) GetClaimableAmount(_validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.GetClaimableAmount(&_Stakemanager.CallOpts, _validator)
}

// GetClaimableAmount is a free data retrieval call binding the contract method 0xe12f3a61.
//
// Solidity: function getClaimableAmount(address _validator) view returns(uint256 _amount)
func (_Stakemanager *StakemanagerCallerSession) GetClaimableAmount(_validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.GetClaimableAmount(&_Stakemanager.CallOpts, _validator)
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _manager, address _slisBnb, address _bscValidator)
func (_Stakemanager *StakemanagerCaller) GetContracts(opts *bind.CallOpts) (struct {
	Manager      common.Address
	SlisBnb      common.Address
	BscValidator common.Address
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getContracts")

	outstruct := new(struct {
		Manager      common.Address
		SlisBnb      common.Address
		BscValidator common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Manager = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SlisBnb = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BscValidator = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _manager, address _slisBnb, address _bscValidator)
func (_Stakemanager *StakemanagerSession) GetContracts() (struct {
	Manager      common.Address
	SlisBnb      common.Address
	BscValidator common.Address
}, error) {
	return _Stakemanager.Contract.GetContracts(&_Stakemanager.CallOpts)
}

// GetContracts is a free data retrieval call binding the contract method 0xc3a2a93a.
//
// Solidity: function getContracts() view returns(address _manager, address _slisBnb, address _bscValidator)
func (_Stakemanager *StakemanagerCallerSession) GetContracts() (struct {
	Manager      common.Address
	SlisBnb      common.Address
	BscValidator common.Address
}, error) {
	return _Stakemanager.Contract.GetContracts(&_Stakemanager.CallOpts)
}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address _validator) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) GetDelegated(opts *bind.CallOpts, _validator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getDelegated", _validator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address _validator) view returns(uint256)
func (_Stakemanager *StakemanagerSession) GetDelegated(_validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.GetDelegated(&_Stakemanager.CallOpts, _validator)
}

// GetDelegated is a free data retrieval call binding the contract method 0x56f0b57e.
//
// Solidity: function getDelegated(address _validator) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) GetDelegated(_validator common.Address) (*big.Int, error) {
	return _Stakemanager.Contract.GetDelegated(&_Stakemanager.CallOpts, _validator)
}

// GetRedelegateFee is a free data retrieval call binding the contract method 0xfb0ddd20.
//
// Solidity: function getRedelegateFee(uint256 _amount) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) GetRedelegateFee(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getRedelegateFee", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedelegateFee is a free data retrieval call binding the contract method 0xfb0ddd20.
//
// Solidity: function getRedelegateFee(uint256 _amount) view returns(uint256)
func (_Stakemanager *StakemanagerSession) GetRedelegateFee(_amount *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetRedelegateFee(&_Stakemanager.CallOpts, _amount)
}

// GetRedelegateFee is a free data retrieval call binding the contract method 0xfb0ddd20.
//
// Solidity: function getRedelegateFee(uint256 _amount) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) GetRedelegateFee(_amount *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetRedelegateFee(&_Stakemanager.CallOpts, _amount)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stakemanager *StakemanagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stakemanager *StakemanagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stakemanager.Contract.GetRoleAdmin(&_Stakemanager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stakemanager *StakemanagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stakemanager.Contract.GetRoleAdmin(&_Stakemanager.CallOpts, role)
}

// GetSlisBnbWithdrawLimit is a free data retrieval call binding the contract method 0xfe5510a6.
//
// Solidity: function getSlisBnbWithdrawLimit() view returns(uint256 _slisBnbWithdrawLimit)
func (_Stakemanager *StakemanagerCaller) GetSlisBnbWithdrawLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getSlisBnbWithdrawLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlisBnbWithdrawLimit is a free data retrieval call binding the contract method 0xfe5510a6.
//
// Solidity: function getSlisBnbWithdrawLimit() view returns(uint256 _slisBnbWithdrawLimit)
func (_Stakemanager *StakemanagerSession) GetSlisBnbWithdrawLimit() (*big.Int, error) {
	return _Stakemanager.Contract.GetSlisBnbWithdrawLimit(&_Stakemanager.CallOpts)
}

// GetSlisBnbWithdrawLimit is a free data retrieval call binding the contract method 0xfe5510a6.
//
// Solidity: function getSlisBnbWithdrawLimit() view returns(uint256 _slisBnbWithdrawLimit)
func (_Stakemanager *StakemanagerCallerSession) GetSlisBnbWithdrawLimit() (*big.Int, error) {
	return _Stakemanager.Contract.GetSlisBnbWithdrawLimit(&_Stakemanager.CallOpts)
}

// GetTotalBnbInValidators is a free data retrieval call binding the contract method 0x786a56c3.
//
// Solidity: function getTotalBnbInValidators() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) GetTotalBnbInValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getTotalBnbInValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBnbInValidators is a free data retrieval call binding the contract method 0x786a56c3.
//
// Solidity: function getTotalBnbInValidators() view returns(uint256)
func (_Stakemanager *StakemanagerSession) GetTotalBnbInValidators() (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalBnbInValidators(&_Stakemanager.CallOpts)
}

// GetTotalBnbInValidators is a free data retrieval call binding the contract method 0x786a56c3.
//
// Solidity: function getTotalBnbInValidators() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) GetTotalBnbInValidators() (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalBnbInValidators(&_Stakemanager.CallOpts)
}

// GetTotalPooledBnb is a free data retrieval call binding the contract method 0x2be7fef5.
//
// Solidity: function getTotalPooledBnb() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) GetTotalPooledBnb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getTotalPooledBnb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledBnb is a free data retrieval call binding the contract method 0x2be7fef5.
//
// Solidity: function getTotalPooledBnb() view returns(uint256)
func (_Stakemanager *StakemanagerSession) GetTotalPooledBnb() (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalPooledBnb(&_Stakemanager.CallOpts)
}

// GetTotalPooledBnb is a free data retrieval call binding the contract method 0x2be7fef5.
//
// Solidity: function getTotalPooledBnb() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) GetTotalPooledBnb() (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalPooledBnb(&_Stakemanager.CallOpts)
}

// GetUserRequestStatus is a free data retrieval call binding the contract method 0x026e3e7b.
//
// Solidity: function getUserRequestStatus(address _user, uint256 _idx) view returns(bool _isClaimable, uint256 _amount)
func (_Stakemanager *StakemanagerCaller) GetUserRequestStatus(opts *bind.CallOpts, _user common.Address, _idx *big.Int) (struct {
	IsClaimable bool
	Amount      *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getUserRequestStatus", _user, _idx)

	outstruct := new(struct {
		IsClaimable bool
		Amount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsClaimable = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserRequestStatus is a free data retrieval call binding the contract method 0x026e3e7b.
//
// Solidity: function getUserRequestStatus(address _user, uint256 _idx) view returns(bool _isClaimable, uint256 _amount)
func (_Stakemanager *StakemanagerSession) GetUserRequestStatus(_user common.Address, _idx *big.Int) (struct {
	IsClaimable bool
	Amount      *big.Int
}, error) {
	return _Stakemanager.Contract.GetUserRequestStatus(&_Stakemanager.CallOpts, _user, _idx)
}

// GetUserRequestStatus is a free data retrieval call binding the contract method 0x026e3e7b.
//
// Solidity: function getUserRequestStatus(address _user, uint256 _idx) view returns(bool _isClaimable, uint256 _amount)
func (_Stakemanager *StakemanagerCallerSession) GetUserRequestStatus(_user common.Address, _idx *big.Int) (struct {
	IsClaimable bool
	Amount      *big.Int
}, error) {
	return _Stakemanager.Contract.GetUserRequestStatus(&_Stakemanager.CallOpts, _user, _idx)
}

// GetUserWithdrawalRequests is a free data retrieval call binding the contract method 0x6c930228.
//
// Solidity: function getUserWithdrawalRequests(address _address) view returns((uint256,uint256,uint256)[])
func (_Stakemanager *StakemanagerCaller) GetUserWithdrawalRequests(opts *bind.CallOpts, _address common.Address) ([]IStakeManagerWithdrawalRequest, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getUserWithdrawalRequests", _address)

	if err != nil {
		return *new([]IStakeManagerWithdrawalRequest), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStakeManagerWithdrawalRequest)).(*[]IStakeManagerWithdrawalRequest)

	return out0, err

}

// GetUserWithdrawalRequests is a free data retrieval call binding the contract method 0x6c930228.
//
// Solidity: function getUserWithdrawalRequests(address _address) view returns((uint256,uint256,uint256)[])
func (_Stakemanager *StakemanagerSession) GetUserWithdrawalRequests(_address common.Address) ([]IStakeManagerWithdrawalRequest, error) {
	return _Stakemanager.Contract.GetUserWithdrawalRequests(&_Stakemanager.CallOpts, _address)
}

// GetUserWithdrawalRequests is a free data retrieval call binding the contract method 0x6c930228.
//
// Solidity: function getUserWithdrawalRequests(address _address) view returns((uint256,uint256,uint256)[])
func (_Stakemanager *StakemanagerCallerSession) GetUserWithdrawalRequests(_address common.Address) ([]IStakeManagerWithdrawalRequest, error) {
	return _Stakemanager.Contract.GetUserWithdrawalRequests(&_Stakemanager.CallOpts, _address)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stakemanager *StakemanagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stakemanager *StakemanagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stakemanager.Contract.HasRole(&_Stakemanager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stakemanager *StakemanagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stakemanager.Contract.HasRole(&_Stakemanager.CallOpts, role, account)
}

// MinBnb is a free data retrieval call binding the contract method 0x6bde8666.
//
// Solidity: function minBnb() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) MinBnb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "minBnb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBnb is a free data retrieval call binding the contract method 0x6bde8666.
//
// Solidity: function minBnb() view returns(uint256)
func (_Stakemanager *StakemanagerSession) MinBnb() (*big.Int, error) {
	return _Stakemanager.Contract.MinBnb(&_Stakemanager.CallOpts)
}

// MinBnb is a free data retrieval call binding the contract method 0x6bde8666.
//
// Solidity: function minBnb() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) MinBnb() (*big.Int, error) {
	return _Stakemanager.Contract.MinBnb(&_Stakemanager.CallOpts)
}

// NextConfirmedRequestUUID is a free data retrieval call binding the contract method 0x642a80c7.
//
// Solidity: function nextConfirmedRequestUUID() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) NextConfirmedRequestUUID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "nextConfirmedRequestUUID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextConfirmedRequestUUID is a free data retrieval call binding the contract method 0x642a80c7.
//
// Solidity: function nextConfirmedRequestUUID() view returns(uint256)
func (_Stakemanager *StakemanagerSession) NextConfirmedRequestUUID() (*big.Int, error) {
	return _Stakemanager.Contract.NextConfirmedRequestUUID(&_Stakemanager.CallOpts)
}

// NextConfirmedRequestUUID is a free data retrieval call binding the contract method 0x642a80c7.
//
// Solidity: function nextConfirmedRequestUUID() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) NextConfirmedRequestUUID() (*big.Int, error) {
	return _Stakemanager.Contract.NextConfirmedRequestUUID(&_Stakemanager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Stakemanager *StakemanagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Stakemanager *StakemanagerSession) Paused() (bool, error) {
	return _Stakemanager.Contract.Paused(&_Stakemanager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Stakemanager *StakemanagerCallerSession) Paused() (bool, error) {
	return _Stakemanager.Contract.Paused(&_Stakemanager.CallOpts)
}

// Placeholder is a free data retrieval call binding the contract method 0x9e75143c.
//
// Solidity: function placeholder() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) Placeholder(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "placeholder")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Placeholder is a free data retrieval call binding the contract method 0x9e75143c.
//
// Solidity: function placeholder() view returns(uint256)
func (_Stakemanager *StakemanagerSession) Placeholder() (*big.Int, error) {
	return _Stakemanager.Contract.Placeholder(&_Stakemanager.CallOpts)
}

// Placeholder is a free data retrieval call binding the contract method 0x9e75143c.
//
// Solidity: function placeholder() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) Placeholder() (*big.Int, error) {
	return _Stakemanager.Contract.Placeholder(&_Stakemanager.CallOpts)
}

// RedirectAddress is a free data retrieval call binding the contract method 0x2d788a5b.
//
// Solidity: function redirectAddress() view returns(address)
func (_Stakemanager *StakemanagerCaller) RedirectAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "redirectAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RedirectAddress is a free data retrieval call binding the contract method 0x2d788a5b.
//
// Solidity: function redirectAddress() view returns(address)
func (_Stakemanager *StakemanagerSession) RedirectAddress() (common.Address, error) {
	return _Stakemanager.Contract.RedirectAddress(&_Stakemanager.CallOpts)
}

// RedirectAddress is a free data retrieval call binding the contract method 0x2d788a5b.
//
// Solidity: function redirectAddress() view returns(address)
func (_Stakemanager *StakemanagerCallerSession) RedirectAddress() (common.Address, error) {
	return _Stakemanager.Contract.RedirectAddress(&_Stakemanager.CallOpts)
}

// Refund is a free data retrieval call binding the contract method 0x590e1ae3.
//
// Solidity: function refund() view returns(uint256 dailySlisBnb, uint256 remainingSlisBnb, uint256 lastBurnTime)
func (_Stakemanager *StakemanagerCaller) Refund(opts *bind.CallOpts) (struct {
	DailySlisBnb     *big.Int
	RemainingSlisBnb *big.Int
	LastBurnTime     *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "refund")

	outstruct := new(struct {
		DailySlisBnb     *big.Int
		RemainingSlisBnb *big.Int
		LastBurnTime     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DailySlisBnb = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RemainingSlisBnb = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastBurnTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Refund is a free data retrieval call binding the contract method 0x590e1ae3.
//
// Solidity: function refund() view returns(uint256 dailySlisBnb, uint256 remainingSlisBnb, uint256 lastBurnTime)
func (_Stakemanager *StakemanagerSession) Refund() (struct {
	DailySlisBnb     *big.Int
	RemainingSlisBnb *big.Int
	LastBurnTime     *big.Int
}, error) {
	return _Stakemanager.Contract.Refund(&_Stakemanager.CallOpts)
}

// Refund is a free data retrieval call binding the contract method 0x590e1ae3.
//
// Solidity: function refund() view returns(uint256 dailySlisBnb, uint256 remainingSlisBnb, uint256 lastBurnTime)
func (_Stakemanager *StakemanagerCallerSession) Refund() (struct {
	DailySlisBnb     *big.Int
	RemainingSlisBnb *big.Int
	LastBurnTime     *big.Int
}, error) {
	return _Stakemanager.Contract.Refund(&_Stakemanager.CallOpts)
}

// RequestIndexMap is a free data retrieval call binding the contract method 0xd94a7174.
//
// Solidity: function requestIndexMap(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerCaller) RequestIndexMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "requestIndexMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIndexMap is a free data retrieval call binding the contract method 0xd94a7174.
//
// Solidity: function requestIndexMap(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerSession) RequestIndexMap(arg0 *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.RequestIndexMap(&_Stakemanager.CallOpts, arg0)
}

// RequestIndexMap is a free data retrieval call binding the contract method 0xd94a7174.
//
// Solidity: function requestIndexMap(uint256 ) view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) RequestIndexMap(arg0 *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.RequestIndexMap(&_Stakemanager.CallOpts, arg0)
}

// RequestUUID is a free data retrieval call binding the contract method 0x1bfabf86.
//
// Solidity: function requestUUID() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) RequestUUID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "requestUUID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestUUID is a free data retrieval call binding the contract method 0x1bfabf86.
//
// Solidity: function requestUUID() view returns(uint256)
func (_Stakemanager *StakemanagerSession) RequestUUID() (*big.Int, error) {
	return _Stakemanager.Contract.RequestUUID(&_Stakemanager.CallOpts)
}

// RequestUUID is a free data retrieval call binding the contract method 0x1bfabf86.
//
// Solidity: function requestUUID() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) RequestUUID() (*big.Int, error) {
	return _Stakemanager.Contract.RequestUUID(&_Stakemanager.CallOpts)
}

// ReserveAmount is a free data retrieval call binding the contract method 0x4b09b72a.
//
// Solidity: function reserveAmount() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) ReserveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "reserveAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveAmount is a free data retrieval call binding the contract method 0x4b09b72a.
//
// Solidity: function reserveAmount() view returns(uint256)
func (_Stakemanager *StakemanagerSession) ReserveAmount() (*big.Int, error) {
	return _Stakemanager.Contract.ReserveAmount(&_Stakemanager.CallOpts)
}

// ReserveAmount is a free data retrieval call binding the contract method 0x4b09b72a.
//
// Solidity: function reserveAmount() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) ReserveAmount() (*big.Int, error) {
	return _Stakemanager.Contract.ReserveAmount(&_Stakemanager.CallOpts)
}

// RevenuePool is a free data retrieval call binding the contract method 0x7f753de6.
//
// Solidity: function revenuePool() view returns(address)
func (_Stakemanager *StakemanagerCaller) RevenuePool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "revenuePool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevenuePool is a free data retrieval call binding the contract method 0x7f753de6.
//
// Solidity: function revenuePool() view returns(address)
func (_Stakemanager *StakemanagerSession) RevenuePool() (common.Address, error) {
	return _Stakemanager.Contract.RevenuePool(&_Stakemanager.CallOpts)
}

// RevenuePool is a free data retrieval call binding the contract method 0x7f753de6.
//
// Solidity: function revenuePool() view returns(address)
func (_Stakemanager *StakemanagerCallerSession) RevenuePool() (common.Address, error) {
	return _Stakemanager.Contract.RevenuePool(&_Stakemanager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stakemanager *StakemanagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stakemanager *StakemanagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stakemanager.Contract.SupportsInterface(&_Stakemanager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stakemanager *StakemanagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stakemanager.Contract.SupportsInterface(&_Stakemanager.CallOpts, interfaceId)
}

// SynFee is a free data retrieval call binding the contract method 0x0fce8a53.
//
// Solidity: function synFee() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) SynFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "synFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SynFee is a free data retrieval call binding the contract method 0x0fce8a53.
//
// Solidity: function synFee() view returns(uint256)
func (_Stakemanager *StakemanagerSession) SynFee() (*big.Int, error) {
	return _Stakemanager.Contract.SynFee(&_Stakemanager.CallOpts)
}

// SynFee is a free data retrieval call binding the contract method 0x0fce8a53.
//
// Solidity: function synFee() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) SynFee() (*big.Int, error) {
	return _Stakemanager.Contract.SynFee(&_Stakemanager.CallOpts)
}

// TotalDelegated is a free data retrieval call binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) TotalDelegated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "totalDelegated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDelegated is a free data retrieval call binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() view returns(uint256)
func (_Stakemanager *StakemanagerSession) TotalDelegated() (*big.Int, error) {
	return _Stakemanager.Contract.TotalDelegated(&_Stakemanager.CallOpts)
}

// TotalDelegated is a free data retrieval call binding the contract method 0x80d04de8.
//
// Solidity: function totalDelegated() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) TotalDelegated() (*big.Int, error) {
	return _Stakemanager.Contract.TotalDelegated(&_Stakemanager.CallOpts)
}

// TotalReserveAmount is a free data retrieval call binding the contract method 0xc585cb08.
//
// Solidity: function totalReserveAmount() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) TotalReserveAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "totalReserveAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserveAmount is a free data retrieval call binding the contract method 0xc585cb08.
//
// Solidity: function totalReserveAmount() view returns(uint256)
func (_Stakemanager *StakemanagerSession) TotalReserveAmount() (*big.Int, error) {
	return _Stakemanager.Contract.TotalReserveAmount(&_Stakemanager.CallOpts)
}

// TotalReserveAmount is a free data retrieval call binding the contract method 0xc585cb08.
//
// Solidity: function totalReserveAmount() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) TotalReserveAmount() (*big.Int, error) {
	return _Stakemanager.Contract.TotalReserveAmount(&_Stakemanager.CallOpts)
}

// UnbondingBnb is a free data retrieval call binding the contract method 0x2fc503b6.
//
// Solidity: function unbondingBnb() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) UnbondingBnb(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "unbondingBnb")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingBnb is a free data retrieval call binding the contract method 0x2fc503b6.
//
// Solidity: function unbondingBnb() view returns(uint256)
func (_Stakemanager *StakemanagerSession) UnbondingBnb() (*big.Int, error) {
	return _Stakemanager.Contract.UnbondingBnb(&_Stakemanager.CallOpts)
}

// UnbondingBnb is a free data retrieval call binding the contract method 0x2fc503b6.
//
// Solidity: function unbondingBnb() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) UnbondingBnb() (*big.Int, error) {
	return _Stakemanager.Contract.UnbondingBnb(&_Stakemanager.CallOpts)
}

// UndelegatedQuota is a free data retrieval call binding the contract method 0x5f9edfca.
//
// Solidity: function undelegatedQuota() view returns(uint256)
func (_Stakemanager *StakemanagerCaller) UndelegatedQuota(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "undelegatedQuota")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UndelegatedQuota is a free data retrieval call binding the contract method 0x5f9edfca.
//
// Solidity: function undelegatedQuota() view returns(uint256)
func (_Stakemanager *StakemanagerSession) UndelegatedQuota() (*big.Int, error) {
	return _Stakemanager.Contract.UndelegatedQuota(&_Stakemanager.CallOpts)
}

// UndelegatedQuota is a free data retrieval call binding the contract method 0x5f9edfca.
//
// Solidity: function undelegatedQuota() view returns(uint256)
func (_Stakemanager *StakemanagerCallerSession) UndelegatedQuota() (*big.Int, error) {
	return _Stakemanager.Contract.UndelegatedQuota(&_Stakemanager.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Stakemanager *StakemanagerCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validators", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Stakemanager *StakemanagerSession) Validators(arg0 common.Address) (bool, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(bool)
func (_Stakemanager *StakemanagerCallerSession) Validators(arg0 common.Address) (bool, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// AcceptNewManager is a paid mutator transaction binding the contract method 0x827fd513.
//
// Solidity: function acceptNewManager() returns()
func (_Stakemanager *StakemanagerTransactor) AcceptNewManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "acceptNewManager")
}

// AcceptNewManager is a paid mutator transaction binding the contract method 0x827fd513.
//
// Solidity: function acceptNewManager() returns()
func (_Stakemanager *StakemanagerSession) AcceptNewManager() (*types.Transaction, error) {
	return _Stakemanager.Contract.AcceptNewManager(&_Stakemanager.TransactOpts)
}

// AcceptNewManager is a paid mutator transaction binding the contract method 0x827fd513.
//
// Solidity: function acceptNewManager() returns()
func (_Stakemanager *StakemanagerTransactorSession) AcceptNewManager() (*types.Transaction, error) {
	return _Stakemanager.Contract.AcceptNewManager(&_Stakemanager.TransactOpts)
}

// ClaimUndelegated is a paid mutator transaction binding the contract method 0x1b3133fb.
//
// Solidity: function claimUndelegated(address _validator) returns(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerTransactor) ClaimUndelegated(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimUndelegated", _validator)
}

// ClaimUndelegated is a paid mutator transaction binding the contract method 0x1b3133fb.
//
// Solidity: function claimUndelegated(address _validator) returns(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerSession) ClaimUndelegated(_validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimUndelegated(&_Stakemanager.TransactOpts, _validator)
}

// ClaimUndelegated is a paid mutator transaction binding the contract method 0x1b3133fb.
//
// Solidity: function claimUndelegated(address _validator) returns(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerTransactorSession) ClaimUndelegated(_validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimUndelegated(&_Stakemanager.TransactOpts, _validator)
}

// ClaimWithdraw is a paid mutator transaction binding the contract method 0xb13acedd.
//
// Solidity: function claimWithdraw(uint256 _idx) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimWithdraw(opts *bind.TransactOpts, _idx *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimWithdraw", _idx)
}

// ClaimWithdraw is a paid mutator transaction binding the contract method 0xb13acedd.
//
// Solidity: function claimWithdraw(uint256 _idx) returns()
func (_Stakemanager *StakemanagerSession) ClaimWithdraw(_idx *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimWithdraw(&_Stakemanager.TransactOpts, _idx)
}

// ClaimWithdraw is a paid mutator transaction binding the contract method 0xb13acedd.
//
// Solidity: function claimWithdraw(uint256 _idx) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimWithdraw(_idx *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimWithdraw(&_Stakemanager.TransactOpts, _idx)
}

// ClaimWithdrawFor is a paid mutator transaction binding the contract method 0x9a53d5af.
//
// Solidity: function claimWithdrawFor(address _user, uint256 _idx) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimWithdrawFor(opts *bind.TransactOpts, _user common.Address, _idx *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimWithdrawFor", _user, _idx)
}

// ClaimWithdrawFor is a paid mutator transaction binding the contract method 0x9a53d5af.
//
// Solidity: function claimWithdrawFor(address _user, uint256 _idx) returns()
func (_Stakemanager *StakemanagerSession) ClaimWithdrawFor(_user common.Address, _idx *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimWithdrawFor(&_Stakemanager.TransactOpts, _user, _idx)
}

// ClaimWithdrawFor is a paid mutator transaction binding the contract method 0x9a53d5af.
//
// Solidity: function claimWithdrawFor(address _user, uint256 _idx) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimWithdrawFor(_user common.Address, _idx *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimWithdrawFor(&_Stakemanager.TransactOpts, _user, _idx)
}

// CompoundRewards is a paid mutator transaction binding the contract method 0xc00c9f7f.
//
// Solidity: function compoundRewards() returns()
func (_Stakemanager *StakemanagerTransactor) CompoundRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "compoundRewards")
}

// CompoundRewards is a paid mutator transaction binding the contract method 0xc00c9f7f.
//
// Solidity: function compoundRewards() returns()
func (_Stakemanager *StakemanagerSession) CompoundRewards() (*types.Transaction, error) {
	return _Stakemanager.Contract.CompoundRewards(&_Stakemanager.TransactOpts)
}

// CompoundRewards is a paid mutator transaction binding the contract method 0xc00c9f7f.
//
// Solidity: function compoundRewards() returns()
func (_Stakemanager *StakemanagerTransactorSession) CompoundRewards() (*types.Transaction, error) {
	return _Stakemanager.Contract.CompoundRewards(&_Stakemanager.TransactOpts)
}

// DelegateTo is a paid mutator transaction binding the contract method 0x8bcd4620.
//
// Solidity: function delegateTo(address _validator, uint256 _amount) returns()
func (_Stakemanager *StakemanagerTransactor) DelegateTo(opts *bind.TransactOpts, _validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "delegateTo", _validator, _amount)
}

// DelegateTo is a paid mutator transaction binding the contract method 0x8bcd4620.
//
// Solidity: function delegateTo(address _validator, uint256 _amount) returns()
func (_Stakemanager *StakemanagerSession) DelegateTo(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.DelegateTo(&_Stakemanager.TransactOpts, _validator, _amount)
}

// DelegateTo is a paid mutator transaction binding the contract method 0x8bcd4620.
//
// Solidity: function delegateTo(address _validator, uint256 _amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) DelegateTo(_validator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.DelegateTo(&_Stakemanager.TransactOpts, _validator, _amount)
}

// DelegateVoteTo is a paid mutator transaction binding the contract method 0x9823004f.
//
// Solidity: function delegateVoteTo(address _delegateTo) returns()
func (_Stakemanager *StakemanagerTransactor) DelegateVoteTo(opts *bind.TransactOpts, _delegateTo common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "delegateVoteTo", _delegateTo)
}

// DelegateVoteTo is a paid mutator transaction binding the contract method 0x9823004f.
//
// Solidity: function delegateVoteTo(address _delegateTo) returns()
func (_Stakemanager *StakemanagerSession) DelegateVoteTo(_delegateTo common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.DelegateVoteTo(&_Stakemanager.TransactOpts, _delegateTo)
}

// DelegateVoteTo is a paid mutator transaction binding the contract method 0x9823004f.
//
// Solidity: function delegateVoteTo(address _delegateTo) returns()
func (_Stakemanager *StakemanagerTransactorSession) DelegateVoteTo(_delegateTo common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.DelegateVoteTo(&_Stakemanager.TransactOpts, _delegateTo)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Stakemanager *StakemanagerTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Stakemanager *StakemanagerSession) Deposit() (*types.Transaction, error) {
	return _Stakemanager.Contract.Deposit(&_Stakemanager.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Stakemanager *StakemanagerTransactorSession) Deposit() (*types.Transaction, error) {
	return _Stakemanager.Contract.Deposit(&_Stakemanager.TransactOpts)
}

// DepositReserve is a paid mutator transaction binding the contract method 0xf251c381.
//
// Solidity: function depositReserve() payable returns()
func (_Stakemanager *StakemanagerTransactor) DepositReserve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "depositReserve")
}

// DepositReserve is a paid mutator transaction binding the contract method 0xf251c381.
//
// Solidity: function depositReserve() payable returns()
func (_Stakemanager *StakemanagerSession) DepositReserve() (*types.Transaction, error) {
	return _Stakemanager.Contract.DepositReserve(&_Stakemanager.TransactOpts)
}

// DepositReserve is a paid mutator transaction binding the contract method 0xf251c381.
//
// Solidity: function depositReserve() payable returns()
func (_Stakemanager *StakemanagerTransactorSession) DepositReserve() (*types.Transaction, error) {
	return _Stakemanager.Contract.DepositReserve(&_Stakemanager.TransactOpts)
}

// DisableValidator is a paid mutator transaction binding the contract method 0x1fe97684.
//
// Solidity: function disableValidator(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) DisableValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "disableValidator", _address)
}

// DisableValidator is a paid mutator transaction binding the contract method 0x1fe97684.
//
// Solidity: function disableValidator(address _address) returns()
func (_Stakemanager *StakemanagerSession) DisableValidator(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.DisableValidator(&_Stakemanager.TransactOpts, _address)
}

// DisableValidator is a paid mutator transaction binding the contract method 0x1fe97684.
//
// Solidity: function disableValidator(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) DisableValidator(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.DisableValidator(&_Stakemanager.TransactOpts, _address)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.GrantRole(&_Stakemanager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.GrantRole(&_Stakemanager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x0eaf43f2.
//
// Solidity: function initialize(address _slisBnb, address _admin, address _manager, address _bot, uint256 _synFee, address _revenuePool, address _validator) returns()
func (_Stakemanager *StakemanagerTransactor) Initialize(opts *bind.TransactOpts, _slisBnb common.Address, _admin common.Address, _manager common.Address, _bot common.Address, _synFee *big.Int, _revenuePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "initialize", _slisBnb, _admin, _manager, _bot, _synFee, _revenuePool, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0x0eaf43f2.
//
// Solidity: function initialize(address _slisBnb, address _admin, address _manager, address _bot, uint256 _synFee, address _revenuePool, address _validator) returns()
func (_Stakemanager *StakemanagerSession) Initialize(_slisBnb common.Address, _admin common.Address, _manager common.Address, _bot common.Address, _synFee *big.Int, _revenuePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Initialize(&_Stakemanager.TransactOpts, _slisBnb, _admin, _manager, _bot, _synFee, _revenuePool, _validator)
}

// Initialize is a paid mutator transaction binding the contract method 0x0eaf43f2.
//
// Solidity: function initialize(address _slisBnb, address _admin, address _manager, address _bot, uint256 _synFee, address _revenuePool, address _validator) returns()
func (_Stakemanager *StakemanagerTransactorSession) Initialize(_slisBnb common.Address, _admin common.Address, _manager common.Address, _bot common.Address, _synFee *big.Int, _revenuePool common.Address, _validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Initialize(&_Stakemanager.TransactOpts, _slisBnb, _admin, _manager, _bot, _synFee, _revenuePool, _validator)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stakemanager *StakemanagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stakemanager *StakemanagerSession) Pause() (*types.Transaction, error) {
	return _Stakemanager.Contract.Pause(&_Stakemanager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Stakemanager *StakemanagerTransactorSession) Pause() (*types.Transaction, error) {
	return _Stakemanager.Contract.Pause(&_Stakemanager.TransactOpts)
}

// ProposeNewManager is a paid mutator transaction binding the contract method 0x53edafba.
//
// Solidity: function proposeNewManager(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) ProposeNewManager(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "proposeNewManager", _address)
}

// ProposeNewManager is a paid mutator transaction binding the contract method 0x53edafba.
//
// Solidity: function proposeNewManager(address _address) returns()
func (_Stakemanager *StakemanagerSession) ProposeNewManager(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ProposeNewManager(&_Stakemanager.TransactOpts, _address)
}

// ProposeNewManager is a paid mutator transaction binding the contract method 0x53edafba.
//
// Solidity: function proposeNewManager(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) ProposeNewManager(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ProposeNewManager(&_Stakemanager.TransactOpts, _address)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 _amount) returns()
func (_Stakemanager *StakemanagerTransactor) Redelegate(opts *bind.TransactOpts, srcValidator common.Address, dstValidator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "redelegate", srcValidator, dstValidator, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 _amount) returns()
func (_Stakemanager *StakemanagerSession) Redelegate(srcValidator common.Address, dstValidator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Redelegate(&_Stakemanager.TransactOpts, srcValidator, dstValidator, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address srcValidator, address dstValidator, uint256 _amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) Redelegate(srcValidator common.Address, dstValidator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Redelegate(&_Stakemanager.TransactOpts, srcValidator, dstValidator, _amount)
}

// RefundCommission is a paid mutator transaction binding the contract method 0x7d5c59eb.
//
// Solidity: function refundCommission(uint256 _days) payable returns()
func (_Stakemanager *StakemanagerTransactor) RefundCommission(opts *bind.TransactOpts, _days *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "refundCommission", _days)
}

// RefundCommission is a paid mutator transaction binding the contract method 0x7d5c59eb.
//
// Solidity: function refundCommission(uint256 _days) payable returns()
func (_Stakemanager *StakemanagerSession) RefundCommission(_days *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.RefundCommission(&_Stakemanager.TransactOpts, _days)
}

// RefundCommission is a paid mutator transaction binding the contract method 0x7d5c59eb.
//
// Solidity: function refundCommission(uint256 _days) payable returns()
func (_Stakemanager *StakemanagerTransactorSession) RefundCommission(_days *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.RefundCommission(&_Stakemanager.TransactOpts, _days)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) RemoveValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "removeValidator", _address)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _address) returns()
func (_Stakemanager *StakemanagerSession) RemoveValidator(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.RemoveValidator(&_Stakemanager.TransactOpts, _address)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) RemoveValidator(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.RemoveValidator(&_Stakemanager.TransactOpts, _address)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.RenounceRole(&_Stakemanager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.RenounceRole(&_Stakemanager.TransactOpts, role, account)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amountInSlisBnb) returns()
func (_Stakemanager *StakemanagerTransactor) RequestWithdraw(opts *bind.TransactOpts, _amountInSlisBnb *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "requestWithdraw", _amountInSlisBnb)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amountInSlisBnb) returns()
func (_Stakemanager *StakemanagerSession) RequestWithdraw(_amountInSlisBnb *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.RequestWithdraw(&_Stakemanager.TransactOpts, _amountInSlisBnb)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0x745400c9.
//
// Solidity: function requestWithdraw(uint256 _amountInSlisBnb) returns()
func (_Stakemanager *StakemanagerTransactorSession) RequestWithdraw(_amountInSlisBnb *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.RequestWithdraw(&_Stakemanager.TransactOpts, _amountInSlisBnb)
}

// RevokeBotRole is a paid mutator transaction binding the contract method 0xa2468f07.
//
// Solidity: function revokeBotRole(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) RevokeBotRole(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "revokeBotRole", _address)
}

// RevokeBotRole is a paid mutator transaction binding the contract method 0xa2468f07.
//
// Solidity: function revokeBotRole(address _address) returns()
func (_Stakemanager *StakemanagerSession) RevokeBotRole(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.RevokeBotRole(&_Stakemanager.TransactOpts, _address)
}

// RevokeBotRole is a paid mutator transaction binding the contract method 0xa2468f07.
//
// Solidity: function revokeBotRole(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) RevokeBotRole(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.RevokeBotRole(&_Stakemanager.TransactOpts, _address)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.RevokeRole(&_Stakemanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stakemanager *StakemanagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.RevokeRole(&_Stakemanager.TransactOpts, role, account)
}

// SetAnnualRate is a paid mutator transaction binding the contract method 0x20953122.
//
// Solidity: function setAnnualRate(uint256 _annualRate) returns()
func (_Stakemanager *StakemanagerTransactor) SetAnnualRate(opts *bind.TransactOpts, _annualRate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setAnnualRate", _annualRate)
}

// SetAnnualRate is a paid mutator transaction binding the contract method 0x20953122.
//
// Solidity: function setAnnualRate(uint256 _annualRate) returns()
func (_Stakemanager *StakemanagerSession) SetAnnualRate(_annualRate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetAnnualRate(&_Stakemanager.TransactOpts, _annualRate)
}

// SetAnnualRate is a paid mutator transaction binding the contract method 0x20953122.
//
// Solidity: function setAnnualRate(uint256 _annualRate) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetAnnualRate(_annualRate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetAnnualRate(&_Stakemanager.TransactOpts, _annualRate)
}

// SetBSCValidator is a paid mutator transaction binding the contract method 0x4c5e6a35.
//
// Solidity: function setBSCValidator(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) SetBSCValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setBSCValidator", _address)
}

// SetBSCValidator is a paid mutator transaction binding the contract method 0x4c5e6a35.
//
// Solidity: function setBSCValidator(address _address) returns()
func (_Stakemanager *StakemanagerSession) SetBSCValidator(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetBSCValidator(&_Stakemanager.TransactOpts, _address)
}

// SetBSCValidator is a paid mutator transaction binding the contract method 0x4c5e6a35.
//
// Solidity: function setBSCValidator(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetBSCValidator(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetBSCValidator(&_Stakemanager.TransactOpts, _address)
}

// SetBotRole is a paid mutator transaction binding the contract method 0x74db4598.
//
// Solidity: function setBotRole(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) SetBotRole(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setBotRole", _address)
}

// SetBotRole is a paid mutator transaction binding the contract method 0x74db4598.
//
// Solidity: function setBotRole(address _address) returns()
func (_Stakemanager *StakemanagerSession) SetBotRole(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetBotRole(&_Stakemanager.TransactOpts, _address)
}

// SetBotRole is a paid mutator transaction binding the contract method 0x74db4598.
//
// Solidity: function setBotRole(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetBotRole(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetBotRole(&_Stakemanager.TransactOpts, _address)
}

// SetMinBnb is a paid mutator transaction binding the contract method 0xe4600014.
//
// Solidity: function setMinBnb(uint256 _amount) returns()
func (_Stakemanager *StakemanagerTransactor) SetMinBnb(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setMinBnb", _amount)
}

// SetMinBnb is a paid mutator transaction binding the contract method 0xe4600014.
//
// Solidity: function setMinBnb(uint256 _amount) returns()
func (_Stakemanager *StakemanagerSession) SetMinBnb(_amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetMinBnb(&_Stakemanager.TransactOpts, _amount)
}

// SetMinBnb is a paid mutator transaction binding the contract method 0xe4600014.
//
// Solidity: function setMinBnb(uint256 _amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetMinBnb(_amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetMinBnb(&_Stakemanager.TransactOpts, _amount)
}

// SetRedirectAddress is a paid mutator transaction binding the contract method 0xb3a589e4.
//
// Solidity: function setRedirectAddress(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) SetRedirectAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setRedirectAddress", _address)
}

// SetRedirectAddress is a paid mutator transaction binding the contract method 0xb3a589e4.
//
// Solidity: function setRedirectAddress(address _address) returns()
func (_Stakemanager *StakemanagerSession) SetRedirectAddress(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetRedirectAddress(&_Stakemanager.TransactOpts, _address)
}

// SetRedirectAddress is a paid mutator transaction binding the contract method 0xb3a589e4.
//
// Solidity: function setRedirectAddress(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetRedirectAddress(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetRedirectAddress(&_Stakemanager.TransactOpts, _address)
}

// SetReserveAmount is a paid mutator transaction binding the contract method 0x8c4a3815.
//
// Solidity: function setReserveAmount(uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) SetReserveAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setReserveAmount", amount)
}

// SetReserveAmount is a paid mutator transaction binding the contract method 0x8c4a3815.
//
// Solidity: function setReserveAmount(uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) SetReserveAmount(amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetReserveAmount(&_Stakemanager.TransactOpts, amount)
}

// SetReserveAmount is a paid mutator transaction binding the contract method 0x8c4a3815.
//
// Solidity: function setReserveAmount(uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetReserveAmount(amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetReserveAmount(&_Stakemanager.TransactOpts, amount)
}

// SetRevenuePool is a paid mutator transaction binding the contract method 0x884957e6.
//
// Solidity: function setRevenuePool(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) SetRevenuePool(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setRevenuePool", _address)
}

// SetRevenuePool is a paid mutator transaction binding the contract method 0x884957e6.
//
// Solidity: function setRevenuePool(address _address) returns()
func (_Stakemanager *StakemanagerSession) SetRevenuePool(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetRevenuePool(&_Stakemanager.TransactOpts, _address)
}

// SetRevenuePool is a paid mutator transaction binding the contract method 0x884957e6.
//
// Solidity: function setRevenuePool(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetRevenuePool(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetRevenuePool(&_Stakemanager.TransactOpts, _address)
}

// SetSynFee is a paid mutator transaction binding the contract method 0x6dda1f38.
//
// Solidity: function setSynFee(uint256 _synFee) returns()
func (_Stakemanager *StakemanagerTransactor) SetSynFee(opts *bind.TransactOpts, _synFee *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "setSynFee", _synFee)
}

// SetSynFee is a paid mutator transaction binding the contract method 0x6dda1f38.
//
// Solidity: function setSynFee(uint256 _synFee) returns()
func (_Stakemanager *StakemanagerSession) SetSynFee(_synFee *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetSynFee(&_Stakemanager.TransactOpts, _synFee)
}

// SetSynFee is a paid mutator transaction binding the contract method 0x6dda1f38.
//
// Solidity: function setSynFee(uint256 _synFee) returns()
func (_Stakemanager *StakemanagerTransactorSession) SetSynFee(_synFee *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.SetSynFee(&_Stakemanager.TransactOpts, _synFee)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Stakemanager *StakemanagerTransactor) TogglePause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "togglePause")
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Stakemanager *StakemanagerSession) TogglePause() (*types.Transaction, error) {
	return _Stakemanager.Contract.TogglePause(&_Stakemanager.TransactOpts)
}

// TogglePause is a paid mutator transaction binding the contract method 0xc4ae3168.
//
// Solidity: function togglePause() returns()
func (_Stakemanager *StakemanagerTransactorSession) TogglePause() (*types.Transaction, error) {
	return _Stakemanager.Contract.TogglePause(&_Stakemanager.TransactOpts)
}

// ToggleVote is a paid mutator transaction binding the contract method 0x20b8ab2d.
//
// Solidity: function toggleVote() returns()
func (_Stakemanager *StakemanagerTransactor) ToggleVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "toggleVote")
}

// ToggleVote is a paid mutator transaction binding the contract method 0x20b8ab2d.
//
// Solidity: function toggleVote() returns()
func (_Stakemanager *StakemanagerSession) ToggleVote() (*types.Transaction, error) {
	return _Stakemanager.Contract.ToggleVote(&_Stakemanager.TransactOpts)
}

// ToggleVote is a paid mutator transaction binding the contract method 0x20b8ab2d.
//
// Solidity: function toggleVote() returns()
func (_Stakemanager *StakemanagerTransactorSession) ToggleVote() (*types.Transaction, error) {
	return _Stakemanager.Contract.ToggleVote(&_Stakemanager.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerTransactor) Undelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "undelegate")
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerSession) Undelegate() (*types.Transaction, error) {
	return _Stakemanager.Contract.Undelegate(&_Stakemanager.TransactOpts)
}

// Undelegate is a paid mutator transaction binding the contract method 0x92ab89bb.
//
// Solidity: function undelegate() returns(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerTransactorSession) Undelegate() (*types.Transaction, error) {
	return _Stakemanager.Contract.Undelegate(&_Stakemanager.TransactOpts)
}

// UndelegateFrom is a paid mutator transaction binding the contract method 0x252b7edc.
//
// Solidity: function undelegateFrom(address _operator, uint256 _amount) returns(uint256 _actualBnbAmount)
func (_Stakemanager *StakemanagerTransactor) UndelegateFrom(opts *bind.TransactOpts, _operator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "undelegateFrom", _operator, _amount)
}

// UndelegateFrom is a paid mutator transaction binding the contract method 0x252b7edc.
//
// Solidity: function undelegateFrom(address _operator, uint256 _amount) returns(uint256 _actualBnbAmount)
func (_Stakemanager *StakemanagerSession) UndelegateFrom(_operator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UndelegateFrom(&_Stakemanager.TransactOpts, _operator, _amount)
}

// UndelegateFrom is a paid mutator transaction binding the contract method 0x252b7edc.
//
// Solidity: function undelegateFrom(address _operator, uint256 _amount) returns(uint256 _actualBnbAmount)
func (_Stakemanager *StakemanagerTransactorSession) UndelegateFrom(_operator common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UndelegateFrom(&_Stakemanager.TransactOpts, _operator, _amount)
}

// WhitelistValidator is a paid mutator transaction binding the contract method 0x85c48a0b.
//
// Solidity: function whitelistValidator(address _address) returns()
func (_Stakemanager *StakemanagerTransactor) WhitelistValidator(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "whitelistValidator", _address)
}

// WhitelistValidator is a paid mutator transaction binding the contract method 0x85c48a0b.
//
// Solidity: function whitelistValidator(address _address) returns()
func (_Stakemanager *StakemanagerSession) WhitelistValidator(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.WhitelistValidator(&_Stakemanager.TransactOpts, _address)
}

// WhitelistValidator is a paid mutator transaction binding the contract method 0x85c48a0b.
//
// Solidity: function whitelistValidator(address _address) returns()
func (_Stakemanager *StakemanagerTransactorSession) WhitelistValidator(_address common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.WhitelistValidator(&_Stakemanager.TransactOpts, _address)
}

// WithdrawReserve is a paid mutator transaction binding the contract method 0x3e696f38.
//
// Solidity: function withdrawReserve(uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) WithdrawReserve(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "withdrawReserve", amount)
}

// WithdrawReserve is a paid mutator transaction binding the contract method 0x3e696f38.
//
// Solidity: function withdrawReserve(uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) WithdrawReserve(amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.WithdrawReserve(&_Stakemanager.TransactOpts, amount)
}

// WithdrawReserve is a paid mutator transaction binding the contract method 0x3e696f38.
//
// Solidity: function withdrawReserve(uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) WithdrawReserve(amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.WithdrawReserve(&_Stakemanager.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakemanager *StakemanagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakemanager *StakemanagerSession) Receive() (*types.Transaction, error) {
	return _Stakemanager.Contract.Receive(&_Stakemanager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakemanager *StakemanagerTransactorSession) Receive() (*types.Transaction, error) {
	return _Stakemanager.Contract.Receive(&_Stakemanager.TransactOpts)
}

// StakemanagerClaimAllWithdrawalsIterator is returned from FilterClaimAllWithdrawals and is used to iterate over the raw logs and unpacked data for ClaimAllWithdrawals events raised by the Stakemanager contract.
type StakemanagerClaimAllWithdrawalsIterator struct {
	Event *StakemanagerClaimAllWithdrawals // Event containing the contract specifics and raw log

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
func (it *StakemanagerClaimAllWithdrawalsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerClaimAllWithdrawals)
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
		it.Event = new(StakemanagerClaimAllWithdrawals)
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
func (it *StakemanagerClaimAllWithdrawalsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerClaimAllWithdrawalsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerClaimAllWithdrawals represents a ClaimAllWithdrawals event raised by the Stakemanager contract.
type StakemanagerClaimAllWithdrawals struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimAllWithdrawals is a free log retrieval operation binding the contract event 0xc4922d5cac1ce2bed206a0d31812035d888c9351379b9e3555481601623c2a4f.
//
// Solidity: event ClaimAllWithdrawals(address indexed _account, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterClaimAllWithdrawals(opts *bind.FilterOpts, _account []common.Address) (*StakemanagerClaimAllWithdrawalsIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ClaimAllWithdrawals", _accountRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerClaimAllWithdrawalsIterator{contract: _Stakemanager.contract, event: "ClaimAllWithdrawals", logs: logs, sub: sub}, nil
}

// WatchClaimAllWithdrawals is a free log subscription operation binding the contract event 0xc4922d5cac1ce2bed206a0d31812035d888c9351379b9e3555481601623c2a4f.
//
// Solidity: event ClaimAllWithdrawals(address indexed _account, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchClaimAllWithdrawals(opts *bind.WatchOpts, sink chan<- *StakemanagerClaimAllWithdrawals, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ClaimAllWithdrawals", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerClaimAllWithdrawals)
				if err := _Stakemanager.contract.UnpackLog(event, "ClaimAllWithdrawals", log); err != nil {
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

// ParseClaimAllWithdrawals is a log parse operation binding the contract event 0xc4922d5cac1ce2bed206a0d31812035d888c9351379b9e3555481601623c2a4f.
//
// Solidity: event ClaimAllWithdrawals(address indexed _account, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseClaimAllWithdrawals(log types.Log) (*StakemanagerClaimAllWithdrawals, error) {
	event := new(StakemanagerClaimAllWithdrawals)
	if err := _Stakemanager.contract.UnpackLog(event, "ClaimAllWithdrawals", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerClaimUndelegatedIterator is returned from FilterClaimUndelegated and is used to iterate over the raw logs and unpacked data for ClaimUndelegated events raised by the Stakemanager contract.
type StakemanagerClaimUndelegatedIterator struct {
	Event *StakemanagerClaimUndelegated // Event containing the contract specifics and raw log

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
func (it *StakemanagerClaimUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerClaimUndelegated)
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
		it.Event = new(StakemanagerClaimUndelegated)
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
func (it *StakemanagerClaimUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerClaimUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerClaimUndelegated represents a ClaimUndelegated event raised by the Stakemanager contract.
type StakemanagerClaimUndelegated struct {
	Uuid   *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaimUndelegated is a free log retrieval operation binding the contract event 0x0879b5e0b459383bae8f6162fd1952d41d72937c429a32478fb1e133895b71ee.
//
// Solidity: event ClaimUndelegated(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterClaimUndelegated(opts *bind.FilterOpts) (*StakemanagerClaimUndelegatedIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ClaimUndelegated")
	if err != nil {
		return nil, err
	}
	return &StakemanagerClaimUndelegatedIterator{contract: _Stakemanager.contract, event: "ClaimUndelegated", logs: logs, sub: sub}, nil
}

// WatchClaimUndelegated is a free log subscription operation binding the contract event 0x0879b5e0b459383bae8f6162fd1952d41d72937c429a32478fb1e133895b71ee.
//
// Solidity: event ClaimUndelegated(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchClaimUndelegated(opts *bind.WatchOpts, sink chan<- *StakemanagerClaimUndelegated) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ClaimUndelegated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerClaimUndelegated)
				if err := _Stakemanager.contract.UnpackLog(event, "ClaimUndelegated", log); err != nil {
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

// ParseClaimUndelegated is a log parse operation binding the contract event 0x0879b5e0b459383bae8f6162fd1952d41d72937c429a32478fb1e133895b71ee.
//
// Solidity: event ClaimUndelegated(uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseClaimUndelegated(log types.Log) (*StakemanagerClaimUndelegated, error) {
	event := new(StakemanagerClaimUndelegated)
	if err := _Stakemanager.contract.UnpackLog(event, "ClaimUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerClaimUndelegatedFromIterator is returned from FilterClaimUndelegatedFrom and is used to iterate over the raw logs and unpacked data for ClaimUndelegatedFrom events raised by the Stakemanager contract.
type StakemanagerClaimUndelegatedFromIterator struct {
	Event *StakemanagerClaimUndelegatedFrom // Event containing the contract specifics and raw log

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
func (it *StakemanagerClaimUndelegatedFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerClaimUndelegatedFrom)
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
		it.Event = new(StakemanagerClaimUndelegatedFrom)
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
func (it *StakemanagerClaimUndelegatedFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerClaimUndelegatedFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerClaimUndelegatedFrom represents a ClaimUndelegatedFrom event raised by the Stakemanager contract.
type StakemanagerClaimUndelegatedFrom struct {
	Validator common.Address
	Uuid      *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimUndelegatedFrom is a free log retrieval operation binding the contract event 0x1031206b6983c708045c4fe17ec7933f2dc7c3e600724e00fe15a9fe383d9ddd.
//
// Solidity: event ClaimUndelegatedFrom(address indexed _validator, uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterClaimUndelegatedFrom(opts *bind.FilterOpts, _validator []common.Address) (*StakemanagerClaimUndelegatedFromIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ClaimUndelegatedFrom", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerClaimUndelegatedFromIterator{contract: _Stakemanager.contract, event: "ClaimUndelegatedFrom", logs: logs, sub: sub}, nil
}

// WatchClaimUndelegatedFrom is a free log subscription operation binding the contract event 0x1031206b6983c708045c4fe17ec7933f2dc7c3e600724e00fe15a9fe383d9ddd.
//
// Solidity: event ClaimUndelegatedFrom(address indexed _validator, uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchClaimUndelegatedFrom(opts *bind.WatchOpts, sink chan<- *StakemanagerClaimUndelegatedFrom, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ClaimUndelegatedFrom", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerClaimUndelegatedFrom)
				if err := _Stakemanager.contract.UnpackLog(event, "ClaimUndelegatedFrom", log); err != nil {
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

// ParseClaimUndelegatedFrom is a log parse operation binding the contract event 0x1031206b6983c708045c4fe17ec7933f2dc7c3e600724e00fe15a9fe383d9ddd.
//
// Solidity: event ClaimUndelegatedFrom(address indexed _validator, uint256 _uuid, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseClaimUndelegatedFrom(log types.Log) (*StakemanagerClaimUndelegatedFrom, error) {
	event := new(StakemanagerClaimUndelegatedFrom)
	if err := _Stakemanager.contract.UnpackLog(event, "ClaimUndelegatedFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerClaimWithdrawalIterator is returned from FilterClaimWithdrawal and is used to iterate over the raw logs and unpacked data for ClaimWithdrawal events raised by the Stakemanager contract.
type StakemanagerClaimWithdrawalIterator struct {
	Event *StakemanagerClaimWithdrawal // Event containing the contract specifics and raw log

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
func (it *StakemanagerClaimWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerClaimWithdrawal)
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
		it.Event = new(StakemanagerClaimWithdrawal)
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
func (it *StakemanagerClaimWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerClaimWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerClaimWithdrawal represents a ClaimWithdrawal event raised by the Stakemanager contract.
type StakemanagerClaimWithdrawal struct {
	Account common.Address
	Idx     *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClaimWithdrawal is a free log retrieval operation binding the contract event 0x63bfb3a58e0713d68e49dda62c223fab04fb534eeef8ac6356cec78e691c092a.
//
// Solidity: event ClaimWithdrawal(address indexed _account, uint256 _idx, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterClaimWithdrawal(opts *bind.FilterOpts, _account []common.Address) (*StakemanagerClaimWithdrawalIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ClaimWithdrawal", _accountRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerClaimWithdrawalIterator{contract: _Stakemanager.contract, event: "ClaimWithdrawal", logs: logs, sub: sub}, nil
}

// WatchClaimWithdrawal is a free log subscription operation binding the contract event 0x63bfb3a58e0713d68e49dda62c223fab04fb534eeef8ac6356cec78e691c092a.
//
// Solidity: event ClaimWithdrawal(address indexed _account, uint256 _idx, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchClaimWithdrawal(opts *bind.WatchOpts, sink chan<- *StakemanagerClaimWithdrawal, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ClaimWithdrawal", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerClaimWithdrawal)
				if err := _Stakemanager.contract.UnpackLog(event, "ClaimWithdrawal", log); err != nil {
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

// ParseClaimWithdrawal is a log parse operation binding the contract event 0x63bfb3a58e0713d68e49dda62c223fab04fb534eeef8ac6356cec78e691c092a.
//
// Solidity: event ClaimWithdrawal(address indexed _account, uint256 _idx, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseClaimWithdrawal(log types.Log) (*StakemanagerClaimWithdrawal, error) {
	event := new(StakemanagerClaimWithdrawal)
	if err := _Stakemanager.contract.UnpackLog(event, "ClaimWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the Stakemanager contract.
type StakemanagerDelegateIterator struct {
	Event *StakemanagerDelegate // Event containing the contract specifics and raw log

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
func (it *StakemanagerDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDelegate)
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
		it.Event = new(StakemanagerDelegate)
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
func (it *StakemanagerDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDelegate represents a Delegate event raised by the Stakemanager contract.
type StakemanagerDelegate struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0xd324a1b76a4e493f0f13d6035169a65d19edfafb054c2561313e9cce3481ce39.
//
// Solidity: event Delegate(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterDelegate(opts *bind.FilterOpts) (*StakemanagerDelegateIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Delegate")
	if err != nil {
		return nil, err
	}
	return &StakemanagerDelegateIterator{contract: _Stakemanager.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0xd324a1b76a4e493f0f13d6035169a65d19edfafb054c2561313e9cce3481ce39.
//
// Solidity: event Delegate(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakemanagerDelegate) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Delegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDelegate)
				if err := _Stakemanager.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0xd324a1b76a4e493f0f13d6035169a65d19edfafb054c2561313e9cce3481ce39.
//
// Solidity: event Delegate(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseDelegate(log types.Log) (*StakemanagerDelegate, error) {
	event := new(StakemanagerDelegate)
	if err := _Stakemanager.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerDelegateToIterator is returned from FilterDelegateTo and is used to iterate over the raw logs and unpacked data for DelegateTo events raised by the Stakemanager contract.
type StakemanagerDelegateToIterator struct {
	Event *StakemanagerDelegateTo // Event containing the contract specifics and raw log

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
func (it *StakemanagerDelegateToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDelegateTo)
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
		it.Event = new(StakemanagerDelegateTo)
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
func (it *StakemanagerDelegateToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDelegateToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDelegateTo represents a DelegateTo event raised by the Stakemanager contract.
type StakemanagerDelegateTo struct {
	Validator         common.Address
	Amount            *big.Int
	DelegateVotePower bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelegateTo is a free log retrieval operation binding the contract event 0x5e71c05bb10b927aa8842ebdca9a130a7f3333b640f0c3cb1380fbef84c34749.
//
// Solidity: event DelegateTo(address _validator, uint256 _amount, bool _delegateVotePower)
func (_Stakemanager *StakemanagerFilterer) FilterDelegateTo(opts *bind.FilterOpts) (*StakemanagerDelegateToIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "DelegateTo")
	if err != nil {
		return nil, err
	}
	return &StakemanagerDelegateToIterator{contract: _Stakemanager.contract, event: "DelegateTo", logs: logs, sub: sub}, nil
}

// WatchDelegateTo is a free log subscription operation binding the contract event 0x5e71c05bb10b927aa8842ebdca9a130a7f3333b640f0c3cb1380fbef84c34749.
//
// Solidity: event DelegateTo(address _validator, uint256 _amount, bool _delegateVotePower)
func (_Stakemanager *StakemanagerFilterer) WatchDelegateTo(opts *bind.WatchOpts, sink chan<- *StakemanagerDelegateTo) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "DelegateTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDelegateTo)
				if err := _Stakemanager.contract.UnpackLog(event, "DelegateTo", log); err != nil {
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

// ParseDelegateTo is a log parse operation binding the contract event 0x5e71c05bb10b927aa8842ebdca9a130a7f3333b640f0c3cb1380fbef84c34749.
//
// Solidity: event DelegateTo(address _validator, uint256 _amount, bool _delegateVotePower)
func (_Stakemanager *StakemanagerFilterer) ParseDelegateTo(log types.Log) (*StakemanagerDelegateTo, error) {
	event := new(StakemanagerDelegateTo)
	if err := _Stakemanager.contract.UnpackLog(event, "DelegateTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerDelegateVoteToIterator is returned from FilterDelegateVoteTo and is used to iterate over the raw logs and unpacked data for DelegateVoteTo events raised by the Stakemanager contract.
type StakemanagerDelegateVoteToIterator struct {
	Event *StakemanagerDelegateVoteTo // Event containing the contract specifics and raw log

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
func (it *StakemanagerDelegateVoteToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDelegateVoteTo)
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
		it.Event = new(StakemanagerDelegateVoteTo)
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
func (it *StakemanagerDelegateVoteToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDelegateVoteToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDelegateVoteTo represents a DelegateVoteTo event raised by the Stakemanager contract.
type StakemanagerDelegateVoteTo struct {
	DelegateTo  common.Address
	VotesChange *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegateVoteTo is a free log retrieval operation binding the contract event 0xca0e65ea013cee72b2994c3d3416c266db4bde421f56eb6b05081d1cbd186530.
//
// Solidity: event DelegateVoteTo(address _delegateTo, uint256 _votesChange)
func (_Stakemanager *StakemanagerFilterer) FilterDelegateVoteTo(opts *bind.FilterOpts) (*StakemanagerDelegateVoteToIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "DelegateVoteTo")
	if err != nil {
		return nil, err
	}
	return &StakemanagerDelegateVoteToIterator{contract: _Stakemanager.contract, event: "DelegateVoteTo", logs: logs, sub: sub}, nil
}

// WatchDelegateVoteTo is a free log subscription operation binding the contract event 0xca0e65ea013cee72b2994c3d3416c266db4bde421f56eb6b05081d1cbd186530.
//
// Solidity: event DelegateVoteTo(address _delegateTo, uint256 _votesChange)
func (_Stakemanager *StakemanagerFilterer) WatchDelegateVoteTo(opts *bind.WatchOpts, sink chan<- *StakemanagerDelegateVoteTo) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "DelegateVoteTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDelegateVoteTo)
				if err := _Stakemanager.contract.UnpackLog(event, "DelegateVoteTo", log); err != nil {
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

// ParseDelegateVoteTo is a log parse operation binding the contract event 0xca0e65ea013cee72b2994c3d3416c266db4bde421f56eb6b05081d1cbd186530.
//
// Solidity: event DelegateVoteTo(address _delegateTo, uint256 _votesChange)
func (_Stakemanager *StakemanagerFilterer) ParseDelegateVoteTo(log types.Log) (*StakemanagerDelegateVoteTo, error) {
	event := new(StakemanagerDelegateVoteTo)
	if err := _Stakemanager.contract.UnpackLog(event, "DelegateVoteTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Stakemanager contract.
type StakemanagerDepositIterator struct {
	Event *StakemanagerDeposit // Event containing the contract specifics and raw log

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
func (it *StakemanagerDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDeposit)
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
		it.Event = new(StakemanagerDeposit)
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
func (it *StakemanagerDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDeposit represents a Deposit event raised by the Stakemanager contract.
type StakemanagerDeposit struct {
	Src    common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address _src, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterDeposit(opts *bind.FilterOpts) (*StakemanagerDepositIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &StakemanagerDepositIterator{contract: _Stakemanager.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address _src, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StakemanagerDeposit) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDeposit)
				if err := _Stakemanager.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address _src, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseDeposit(log types.Log) (*StakemanagerDeposit, error) {
	event := new(StakemanagerDeposit)
	if err := _Stakemanager.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerDisableValidatorIterator is returned from FilterDisableValidator and is used to iterate over the raw logs and unpacked data for DisableValidator events raised by the Stakemanager contract.
type StakemanagerDisableValidatorIterator struct {
	Event *StakemanagerDisableValidator // Event containing the contract specifics and raw log

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
func (it *StakemanagerDisableValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerDisableValidator)
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
		it.Event = new(StakemanagerDisableValidator)
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
func (it *StakemanagerDisableValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerDisableValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerDisableValidator represents a DisableValidator event raised by the Stakemanager contract.
type StakemanagerDisableValidator struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisableValidator is a free log retrieval operation binding the contract event 0x36a773811aa40c53a642b0596c94174b588d0d1b2b9b6fa6e31c3c30d686163f.
//
// Solidity: event DisableValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) FilterDisableValidator(opts *bind.FilterOpts, _address []common.Address) (*StakemanagerDisableValidatorIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "DisableValidator", _addressRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerDisableValidatorIterator{contract: _Stakemanager.contract, event: "DisableValidator", logs: logs, sub: sub}, nil
}

// WatchDisableValidator is a free log subscription operation binding the contract event 0x36a773811aa40c53a642b0596c94174b588d0d1b2b9b6fa6e31c3c30d686163f.
//
// Solidity: event DisableValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) WatchDisableValidator(opts *bind.WatchOpts, sink chan<- *StakemanagerDisableValidator, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "DisableValidator", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerDisableValidator)
				if err := _Stakemanager.contract.UnpackLog(event, "DisableValidator", log); err != nil {
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

// ParseDisableValidator is a log parse operation binding the contract event 0x36a773811aa40c53a642b0596c94174b588d0d1b2b9b6fa6e31c3c30d686163f.
//
// Solidity: event DisableValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) ParseDisableValidator(log types.Log) (*StakemanagerDisableValidator, error) {
	event := new(StakemanagerDisableValidator)
	if err := _Stakemanager.contract.UnpackLog(event, "DisableValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Stakemanager contract.
type StakemanagerInitializedIterator struct {
	Event *StakemanagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakemanagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerInitialized)
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
		it.Event = new(StakemanagerInitialized)
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
func (it *StakemanagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerInitialized represents a Initialized event raised by the Stakemanager contract.
type StakemanagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stakemanager *StakemanagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakemanagerInitializedIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakemanagerInitializedIterator{contract: _Stakemanager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Stakemanager *StakemanagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakemanagerInitialized) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerInitialized)
				if err := _Stakemanager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Stakemanager *StakemanagerFilterer) ParseInitialized(log types.Log) (*StakemanagerInitialized, error) {
	event := new(StakemanagerInitialized)
	if err := _Stakemanager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Stakemanager contract.
type StakemanagerPausedIterator struct {
	Event *StakemanagerPaused // Event containing the contract specifics and raw log

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
func (it *StakemanagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerPaused)
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
		it.Event = new(StakemanagerPaused)
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
func (it *StakemanagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerPaused represents a Paused event raised by the Stakemanager contract.
type StakemanagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Stakemanager *StakemanagerFilterer) FilterPaused(opts *bind.FilterOpts) (*StakemanagerPausedIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakemanagerPausedIterator{contract: _Stakemanager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Stakemanager *StakemanagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakemanagerPaused) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerPaused)
				if err := _Stakemanager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Stakemanager *StakemanagerFilterer) ParsePaused(log types.Log) (*StakemanagerPaused, error) {
	event := new(StakemanagerPaused)
	if err := _Stakemanager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerProposeManagerIterator is returned from FilterProposeManager and is used to iterate over the raw logs and unpacked data for ProposeManager events raised by the Stakemanager contract.
type StakemanagerProposeManagerIterator struct {
	Event *StakemanagerProposeManager // Event containing the contract specifics and raw log

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
func (it *StakemanagerProposeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerProposeManager)
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
		it.Event = new(StakemanagerProposeManager)
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
func (it *StakemanagerProposeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerProposeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerProposeManager represents a ProposeManager event raised by the Stakemanager contract.
type StakemanagerProposeManager struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProposeManager is a free log retrieval operation binding the contract event 0x00e90865ba774da1b7638b592f529a31eb753da0d7cd3a1396d69404b4e2e4e2.
//
// Solidity: event ProposeManager(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) FilterProposeManager(opts *bind.FilterOpts, _address []common.Address) (*StakemanagerProposeManagerIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ProposeManager", _addressRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerProposeManagerIterator{contract: _Stakemanager.contract, event: "ProposeManager", logs: logs, sub: sub}, nil
}

// WatchProposeManager is a free log subscription operation binding the contract event 0x00e90865ba774da1b7638b592f529a31eb753da0d7cd3a1396d69404b4e2e4e2.
//
// Solidity: event ProposeManager(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) WatchProposeManager(opts *bind.WatchOpts, sink chan<- *StakemanagerProposeManager, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ProposeManager", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerProposeManager)
				if err := _Stakemanager.contract.UnpackLog(event, "ProposeManager", log); err != nil {
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

// ParseProposeManager is a log parse operation binding the contract event 0x00e90865ba774da1b7638b592f529a31eb753da0d7cd3a1396d69404b4e2e4e2.
//
// Solidity: event ProposeManager(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) ParseProposeManager(log types.Log) (*StakemanagerProposeManager, error) {
	event := new(StakemanagerProposeManager)
	if err := _Stakemanager.contract.UnpackLog(event, "ProposeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerReDelegateIterator is returned from FilterReDelegate and is used to iterate over the raw logs and unpacked data for ReDelegate events raised by the Stakemanager contract.
type StakemanagerReDelegateIterator struct {
	Event *StakemanagerReDelegate // Event containing the contract specifics and raw log

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
func (it *StakemanagerReDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerReDelegate)
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
		it.Event = new(StakemanagerReDelegate)
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
func (it *StakemanagerReDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerReDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerReDelegate represents a ReDelegate event raised by the Stakemanager contract.
type StakemanagerReDelegate struct {
	Src    common.Address
	Dest   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReDelegate is a free log retrieval operation binding the contract event 0x0f23ce3b62792113baf786ff98a77f4309d51920fcb7d88bab3f7e6f53d2287a.
//
// Solidity: event ReDelegate(address _src, address _dest, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterReDelegate(opts *bind.FilterOpts) (*StakemanagerReDelegateIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ReDelegate")
	if err != nil {
		return nil, err
	}
	return &StakemanagerReDelegateIterator{contract: _Stakemanager.contract, event: "ReDelegate", logs: logs, sub: sub}, nil
}

// WatchReDelegate is a free log subscription operation binding the contract event 0x0f23ce3b62792113baf786ff98a77f4309d51920fcb7d88bab3f7e6f53d2287a.
//
// Solidity: event ReDelegate(address _src, address _dest, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchReDelegate(opts *bind.WatchOpts, sink chan<- *StakemanagerReDelegate) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ReDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerReDelegate)
				if err := _Stakemanager.contract.UnpackLog(event, "ReDelegate", log); err != nil {
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

// ParseReDelegate is a log parse operation binding the contract event 0x0f23ce3b62792113baf786ff98a77f4309d51920fcb7d88bab3f7e6f53d2287a.
//
// Solidity: event ReDelegate(address _src, address _dest, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseReDelegate(log types.Log) (*StakemanagerReDelegate, error) {
	event := new(StakemanagerReDelegate)
	if err := _Stakemanager.contract.UnpackLog(event, "ReDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerRedelegateIterator is returned from FilterRedelegate and is used to iterate over the raw logs and unpacked data for Redelegate events raised by the Stakemanager contract.
type StakemanagerRedelegateIterator struct {
	Event *StakemanagerRedelegate // Event containing the contract specifics and raw log

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
func (it *StakemanagerRedelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRedelegate)
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
		it.Event = new(StakemanagerRedelegate)
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
func (it *StakemanagerRedelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRedelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRedelegate represents a Redelegate event raised by the Stakemanager contract.
type StakemanagerRedelegate struct {
	RewardsId *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedelegate is a free log retrieval operation binding the contract event 0x18112e7eb30e4ec7cbd0751caa279e3f69a962e8bcfddb6919087597cc473ec3.
//
// Solidity: event Redelegate(uint256 _rewardsId, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterRedelegate(opts *bind.FilterOpts) (*StakemanagerRedelegateIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Redelegate")
	if err != nil {
		return nil, err
	}
	return &StakemanagerRedelegateIterator{contract: _Stakemanager.contract, event: "Redelegate", logs: logs, sub: sub}, nil
}

// WatchRedelegate is a free log subscription operation binding the contract event 0x18112e7eb30e4ec7cbd0751caa279e3f69a962e8bcfddb6919087597cc473ec3.
//
// Solidity: event Redelegate(uint256 _rewardsId, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchRedelegate(opts *bind.WatchOpts, sink chan<- *StakemanagerRedelegate) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Redelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRedelegate)
				if err := _Stakemanager.contract.UnpackLog(event, "Redelegate", log); err != nil {
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

// ParseRedelegate is a log parse operation binding the contract event 0x18112e7eb30e4ec7cbd0751caa279e3f69a962e8bcfddb6919087597cc473ec3.
//
// Solidity: event Redelegate(uint256 _rewardsId, uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseRedelegate(log types.Log) (*StakemanagerRedelegate, error) {
	event := new(StakemanagerRedelegate)
	if err := _Stakemanager.contract.UnpackLog(event, "Redelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerRefundCommissionIterator is returned from FilterRefundCommission and is used to iterate over the raw logs and unpacked data for RefundCommission events raised by the Stakemanager contract.
type StakemanagerRefundCommissionIterator struct {
	Event *StakemanagerRefundCommission // Event containing the contract specifics and raw log

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
func (it *StakemanagerRefundCommissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRefundCommission)
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
		it.Event = new(StakemanagerRefundCommission)
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
func (it *StakemanagerRefundCommissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRefundCommissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRefundCommission represents a RefundCommission event raised by the Stakemanager contract.
type StakemanagerRefundCommission struct {
	BnbAmount        *big.Int
	DailySlisBnb     *big.Int
	Days             *big.Int
	RemainingSlisBnb *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRefundCommission is a free log retrieval operation binding the contract event 0xd586d49d4f659031867ca21e9a1c82179f2ab7946a8f75670218d8e494b5e5d5.
//
// Solidity: event RefundCommission(uint256 _bnbAmount, uint256 _dailySlisBnb, uint256 _days, uint256 _remainingSlisBnb)
func (_Stakemanager *StakemanagerFilterer) FilterRefundCommission(opts *bind.FilterOpts) (*StakemanagerRefundCommissionIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RefundCommission")
	if err != nil {
		return nil, err
	}
	return &StakemanagerRefundCommissionIterator{contract: _Stakemanager.contract, event: "RefundCommission", logs: logs, sub: sub}, nil
}

// WatchRefundCommission is a free log subscription operation binding the contract event 0xd586d49d4f659031867ca21e9a1c82179f2ab7946a8f75670218d8e494b5e5d5.
//
// Solidity: event RefundCommission(uint256 _bnbAmount, uint256 _dailySlisBnb, uint256 _days, uint256 _remainingSlisBnb)
func (_Stakemanager *StakemanagerFilterer) WatchRefundCommission(opts *bind.WatchOpts, sink chan<- *StakemanagerRefundCommission) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RefundCommission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRefundCommission)
				if err := _Stakemanager.contract.UnpackLog(event, "RefundCommission", log); err != nil {
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

// ParseRefundCommission is a log parse operation binding the contract event 0xd586d49d4f659031867ca21e9a1c82179f2ab7946a8f75670218d8e494b5e5d5.
//
// Solidity: event RefundCommission(uint256 _bnbAmount, uint256 _dailySlisBnb, uint256 _days, uint256 _remainingSlisBnb)
func (_Stakemanager *StakemanagerFilterer) ParseRefundCommission(log types.Log) (*StakemanagerRefundCommission, error) {
	event := new(StakemanagerRefundCommission)
	if err := _Stakemanager.contract.UnpackLog(event, "RefundCommission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerRemoveValidatorIterator is returned from FilterRemoveValidator and is used to iterate over the raw logs and unpacked data for RemoveValidator events raised by the Stakemanager contract.
type StakemanagerRemoveValidatorIterator struct {
	Event *StakemanagerRemoveValidator // Event containing the contract specifics and raw log

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
func (it *StakemanagerRemoveValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRemoveValidator)
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
		it.Event = new(StakemanagerRemoveValidator)
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
func (it *StakemanagerRemoveValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRemoveValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRemoveValidator represents a RemoveValidator event raised by the Stakemanager contract.
type StakemanagerRemoveValidator struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveValidator is a free log retrieval operation binding the contract event 0x1af60f72d206709ac9c5fd393b54381507af4df1feb01708422ef8498c57aa57.
//
// Solidity: event RemoveValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) FilterRemoveValidator(opts *bind.FilterOpts, _address []common.Address) (*StakemanagerRemoveValidatorIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RemoveValidator", _addressRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerRemoveValidatorIterator{contract: _Stakemanager.contract, event: "RemoveValidator", logs: logs, sub: sub}, nil
}

// WatchRemoveValidator is a free log subscription operation binding the contract event 0x1af60f72d206709ac9c5fd393b54381507af4df1feb01708422ef8498c57aa57.
//
// Solidity: event RemoveValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) WatchRemoveValidator(opts *bind.WatchOpts, sink chan<- *StakemanagerRemoveValidator, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RemoveValidator", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRemoveValidator)
				if err := _Stakemanager.contract.UnpackLog(event, "RemoveValidator", log); err != nil {
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

// ParseRemoveValidator is a log parse operation binding the contract event 0x1af60f72d206709ac9c5fd393b54381507af4df1feb01708422ef8498c57aa57.
//
// Solidity: event RemoveValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) ParseRemoveValidator(log types.Log) (*StakemanagerRemoveValidator, error) {
	event := new(StakemanagerRemoveValidator)
	if err := _Stakemanager.contract.UnpackLog(event, "RemoveValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerRequestWithdrawIterator is returned from FilterRequestWithdraw and is used to iterate over the raw logs and unpacked data for RequestWithdraw events raised by the Stakemanager contract.
type StakemanagerRequestWithdrawIterator struct {
	Event *StakemanagerRequestWithdraw // Event containing the contract specifics and raw log

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
func (it *StakemanagerRequestWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRequestWithdraw)
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
		it.Event = new(StakemanagerRequestWithdraw)
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
func (it *StakemanagerRequestWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRequestWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRequestWithdraw represents a RequestWithdraw event raised by the Stakemanager contract.
type StakemanagerRequestWithdraw struct {
	Account         common.Address
	AmountInSlisBnb *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdraw is a free log retrieval operation binding the contract event 0x0afd74a2a0a78f6c15e41029f44995ee023fe49276f44a4b2b2cf674829362e6.
//
// Solidity: event RequestWithdraw(address indexed _account, uint256 _amountInSlisBnb)
func (_Stakemanager *StakemanagerFilterer) FilterRequestWithdraw(opts *bind.FilterOpts, _account []common.Address) (*StakemanagerRequestWithdrawIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RequestWithdraw", _accountRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerRequestWithdrawIterator{contract: _Stakemanager.contract, event: "RequestWithdraw", logs: logs, sub: sub}, nil
}

// WatchRequestWithdraw is a free log subscription operation binding the contract event 0x0afd74a2a0a78f6c15e41029f44995ee023fe49276f44a4b2b2cf674829362e6.
//
// Solidity: event RequestWithdraw(address indexed _account, uint256 _amountInSlisBnb)
func (_Stakemanager *StakemanagerFilterer) WatchRequestWithdraw(opts *bind.WatchOpts, sink chan<- *StakemanagerRequestWithdraw, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RequestWithdraw", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRequestWithdraw)
				if err := _Stakemanager.contract.UnpackLog(event, "RequestWithdraw", log); err != nil {
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

// ParseRequestWithdraw is a log parse operation binding the contract event 0x0afd74a2a0a78f6c15e41029f44995ee023fe49276f44a4b2b2cf674829362e6.
//
// Solidity: event RequestWithdraw(address indexed _account, uint256 _amountInSlisBnb)
func (_Stakemanager *StakemanagerFilterer) ParseRequestWithdraw(log types.Log) (*StakemanagerRequestWithdraw, error) {
	event := new(StakemanagerRequestWithdraw)
	if err := _Stakemanager.contract.UnpackLog(event, "RequestWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerRewardsCompoundedIterator is returned from FilterRewardsCompounded and is used to iterate over the raw logs and unpacked data for RewardsCompounded events raised by the Stakemanager contract.
type StakemanagerRewardsCompoundedIterator struct {
	Event *StakemanagerRewardsCompounded // Event containing the contract specifics and raw log

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
func (it *StakemanagerRewardsCompoundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRewardsCompounded)
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
		it.Event = new(StakemanagerRewardsCompounded)
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
func (it *StakemanagerRewardsCompoundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRewardsCompoundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRewardsCompounded represents a RewardsCompounded event raised by the Stakemanager contract.
type StakemanagerRewardsCompounded struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsCompounded is a free log retrieval operation binding the contract event 0x33222d90f0305cdd84f37927a34b36077e5328c8318dc017be0eed3e0305fa2e.
//
// Solidity: event RewardsCompounded(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterRewardsCompounded(opts *bind.FilterOpts) (*StakemanagerRewardsCompoundedIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RewardsCompounded")
	if err != nil {
		return nil, err
	}
	return &StakemanagerRewardsCompoundedIterator{contract: _Stakemanager.contract, event: "RewardsCompounded", logs: logs, sub: sub}, nil
}

// WatchRewardsCompounded is a free log subscription operation binding the contract event 0x33222d90f0305cdd84f37927a34b36077e5328c8318dc017be0eed3e0305fa2e.
//
// Solidity: event RewardsCompounded(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchRewardsCompounded(opts *bind.WatchOpts, sink chan<- *StakemanagerRewardsCompounded) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RewardsCompounded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRewardsCompounded)
				if err := _Stakemanager.contract.UnpackLog(event, "RewardsCompounded", log); err != nil {
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

// ParseRewardsCompounded is a log parse operation binding the contract event 0x33222d90f0305cdd84f37927a34b36077e5328c8318dc017be0eed3e0305fa2e.
//
// Solidity: event RewardsCompounded(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseRewardsCompounded(log types.Log) (*StakemanagerRewardsCompounded, error) {
	event := new(StakemanagerRewardsCompounded)
	if err := _Stakemanager.contract.UnpackLog(event, "RewardsCompounded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Stakemanager contract.
type StakemanagerRoleAdminChangedIterator struct {
	Event *StakemanagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakemanagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRoleAdminChanged)
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
		it.Event = new(StakemanagerRoleAdminChanged)
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
func (it *StakemanagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRoleAdminChanged represents a RoleAdminChanged event raised by the Stakemanager contract.
type StakemanagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stakemanager *StakemanagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakemanagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerRoleAdminChangedIterator{contract: _Stakemanager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stakemanager *StakemanagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakemanagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRoleAdminChanged)
				if err := _Stakemanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Stakemanager *StakemanagerFilterer) ParseRoleAdminChanged(log types.Log) (*StakemanagerRoleAdminChanged, error) {
	event := new(StakemanagerRoleAdminChanged)
	if err := _Stakemanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Stakemanager contract.
type StakemanagerRoleGrantedIterator struct {
	Event *StakemanagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakemanagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRoleGranted)
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
		it.Event = new(StakemanagerRoleGranted)
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
func (it *StakemanagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRoleGranted represents a RoleGranted event raised by the Stakemanager contract.
type StakemanagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stakemanager *StakemanagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakemanagerRoleGrantedIterator, error) {

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

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerRoleGrantedIterator{contract: _Stakemanager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stakemanager *StakemanagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakemanagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRoleGranted)
				if err := _Stakemanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Stakemanager *StakemanagerFilterer) ParseRoleGranted(log types.Log) (*StakemanagerRoleGranted, error) {
	event := new(StakemanagerRoleGranted)
	if err := _Stakemanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Stakemanager contract.
type StakemanagerRoleRevokedIterator struct {
	Event *StakemanagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakemanagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerRoleRevoked)
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
		it.Event = new(StakemanagerRoleRevoked)
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
func (it *StakemanagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerRoleRevoked represents a RoleRevoked event raised by the Stakemanager contract.
type StakemanagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stakemanager *StakemanagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakemanagerRoleRevokedIterator, error) {

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

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerRoleRevokedIterator{contract: _Stakemanager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stakemanager *StakemanagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakemanagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerRoleRevoked)
				if err := _Stakemanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Stakemanager *StakemanagerFilterer) ParseRoleRevoked(log types.Log) (*StakemanagerRoleRevoked, error) {
	event := new(StakemanagerRoleRevoked)
	if err := _Stakemanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSetAnnualRateIterator is returned from FilterSetAnnualRate and is used to iterate over the raw logs and unpacked data for SetAnnualRate events raised by the Stakemanager contract.
type StakemanagerSetAnnualRateIterator struct {
	Event *StakemanagerSetAnnualRate // Event containing the contract specifics and raw log

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
func (it *StakemanagerSetAnnualRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSetAnnualRate)
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
		it.Event = new(StakemanagerSetAnnualRate)
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
func (it *StakemanagerSetAnnualRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSetAnnualRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSetAnnualRate represents a SetAnnualRate event raised by the Stakemanager contract.
type StakemanagerSetAnnualRate struct {
	AnnualRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetAnnualRate is a free log retrieval operation binding the contract event 0x12c5bb090b58777f7a8be049e6c34fd42ec1a34e585fc5e5747169309b56da3d.
//
// Solidity: event SetAnnualRate(uint256 _annualRate)
func (_Stakemanager *StakemanagerFilterer) FilterSetAnnualRate(opts *bind.FilterOpts) (*StakemanagerSetAnnualRateIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SetAnnualRate")
	if err != nil {
		return nil, err
	}
	return &StakemanagerSetAnnualRateIterator{contract: _Stakemanager.contract, event: "SetAnnualRate", logs: logs, sub: sub}, nil
}

// WatchSetAnnualRate is a free log subscription operation binding the contract event 0x12c5bb090b58777f7a8be049e6c34fd42ec1a34e585fc5e5747169309b56da3d.
//
// Solidity: event SetAnnualRate(uint256 _annualRate)
func (_Stakemanager *StakemanagerFilterer) WatchSetAnnualRate(opts *bind.WatchOpts, sink chan<- *StakemanagerSetAnnualRate) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SetAnnualRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSetAnnualRate)
				if err := _Stakemanager.contract.UnpackLog(event, "SetAnnualRate", log); err != nil {
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

// ParseSetAnnualRate is a log parse operation binding the contract event 0x12c5bb090b58777f7a8be049e6c34fd42ec1a34e585fc5e5747169309b56da3d.
//
// Solidity: event SetAnnualRate(uint256 _annualRate)
func (_Stakemanager *StakemanagerFilterer) ParseSetAnnualRate(log types.Log) (*StakemanagerSetAnnualRate, error) {
	event := new(StakemanagerSetAnnualRate)
	if err := _Stakemanager.contract.UnpackLog(event, "SetAnnualRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSetBSCValidatorIterator is returned from FilterSetBSCValidator and is used to iterate over the raw logs and unpacked data for SetBSCValidator events raised by the Stakemanager contract.
type StakemanagerSetBSCValidatorIterator struct {
	Event *StakemanagerSetBSCValidator // Event containing the contract specifics and raw log

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
func (it *StakemanagerSetBSCValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSetBSCValidator)
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
		it.Event = new(StakemanagerSetBSCValidator)
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
func (it *StakemanagerSetBSCValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSetBSCValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSetBSCValidator represents a SetBSCValidator event raised by the Stakemanager contract.
type StakemanagerSetBSCValidator struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBSCValidator is a free log retrieval operation binding the contract event 0xd673999808ca5be155f42a945c117092db962771c6f2bcfd96c7746d75f2d472.
//
// Solidity: event SetBSCValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) FilterSetBSCValidator(opts *bind.FilterOpts, _address []common.Address) (*StakemanagerSetBSCValidatorIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SetBSCValidator", _addressRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerSetBSCValidatorIterator{contract: _Stakemanager.contract, event: "SetBSCValidator", logs: logs, sub: sub}, nil
}

// WatchSetBSCValidator is a free log subscription operation binding the contract event 0xd673999808ca5be155f42a945c117092db962771c6f2bcfd96c7746d75f2d472.
//
// Solidity: event SetBSCValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) WatchSetBSCValidator(opts *bind.WatchOpts, sink chan<- *StakemanagerSetBSCValidator, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SetBSCValidator", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSetBSCValidator)
				if err := _Stakemanager.contract.UnpackLog(event, "SetBSCValidator", log); err != nil {
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

// ParseSetBSCValidator is a log parse operation binding the contract event 0xd673999808ca5be155f42a945c117092db962771c6f2bcfd96c7746d75f2d472.
//
// Solidity: event SetBSCValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) ParseSetBSCValidator(log types.Log) (*StakemanagerSetBSCValidator, error) {
	event := new(StakemanagerSetBSCValidator)
	if err := _Stakemanager.contract.UnpackLog(event, "SetBSCValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSetManagerIterator is returned from FilterSetManager and is used to iterate over the raw logs and unpacked data for SetManager events raised by the Stakemanager contract.
type StakemanagerSetManagerIterator struct {
	Event *StakemanagerSetManager // Event containing the contract specifics and raw log

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
func (it *StakemanagerSetManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSetManager)
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
		it.Event = new(StakemanagerSetManager)
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
func (it *StakemanagerSetManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSetManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSetManager represents a SetManager event raised by the Stakemanager contract.
type StakemanagerSetManager struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetManager is a free log retrieval operation binding the contract event 0x54a6385aa0292b04e1ef8513253c17d1863f7cdfc87029d77fd55cc4c2e717e2.
//
// Solidity: event SetManager(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) FilterSetManager(opts *bind.FilterOpts, _address []common.Address) (*StakemanagerSetManagerIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SetManager", _addressRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerSetManagerIterator{contract: _Stakemanager.contract, event: "SetManager", logs: logs, sub: sub}, nil
}

// WatchSetManager is a free log subscription operation binding the contract event 0x54a6385aa0292b04e1ef8513253c17d1863f7cdfc87029d77fd55cc4c2e717e2.
//
// Solidity: event SetManager(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) WatchSetManager(opts *bind.WatchOpts, sink chan<- *StakemanagerSetManager, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SetManager", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSetManager)
				if err := _Stakemanager.contract.UnpackLog(event, "SetManager", log); err != nil {
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

// ParseSetManager is a log parse operation binding the contract event 0x54a6385aa0292b04e1ef8513253c17d1863f7cdfc87029d77fd55cc4c2e717e2.
//
// Solidity: event SetManager(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) ParseSetManager(log types.Log) (*StakemanagerSetManager, error) {
	event := new(StakemanagerSetManager)
	if err := _Stakemanager.contract.UnpackLog(event, "SetManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSetMinBnbIterator is returned from FilterSetMinBnb and is used to iterate over the raw logs and unpacked data for SetMinBnb events raised by the Stakemanager contract.
type StakemanagerSetMinBnbIterator struct {
	Event *StakemanagerSetMinBnb // Event containing the contract specifics and raw log

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
func (it *StakemanagerSetMinBnbIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSetMinBnb)
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
		it.Event = new(StakemanagerSetMinBnb)
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
func (it *StakemanagerSetMinBnbIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSetMinBnbIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSetMinBnb represents a SetMinBnb event raised by the Stakemanager contract.
type StakemanagerSetMinBnb struct {
	MinBnb *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetMinBnb is a free log retrieval operation binding the contract event 0xf34636ae263adb220f7ac9c16562e882a5a644e97b5de1df187d92108659920a.
//
// Solidity: event SetMinBnb(uint256 _minBnb)
func (_Stakemanager *StakemanagerFilterer) FilterSetMinBnb(opts *bind.FilterOpts) (*StakemanagerSetMinBnbIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SetMinBnb")
	if err != nil {
		return nil, err
	}
	return &StakemanagerSetMinBnbIterator{contract: _Stakemanager.contract, event: "SetMinBnb", logs: logs, sub: sub}, nil
}

// WatchSetMinBnb is a free log subscription operation binding the contract event 0xf34636ae263adb220f7ac9c16562e882a5a644e97b5de1df187d92108659920a.
//
// Solidity: event SetMinBnb(uint256 _minBnb)
func (_Stakemanager *StakemanagerFilterer) WatchSetMinBnb(opts *bind.WatchOpts, sink chan<- *StakemanagerSetMinBnb) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SetMinBnb")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSetMinBnb)
				if err := _Stakemanager.contract.UnpackLog(event, "SetMinBnb", log); err != nil {
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

// ParseSetMinBnb is a log parse operation binding the contract event 0xf34636ae263adb220f7ac9c16562e882a5a644e97b5de1df187d92108659920a.
//
// Solidity: event SetMinBnb(uint256 _minBnb)
func (_Stakemanager *StakemanagerFilterer) ParseSetMinBnb(log types.Log) (*StakemanagerSetMinBnb, error) {
	event := new(StakemanagerSetMinBnb)
	if err := _Stakemanager.contract.UnpackLog(event, "SetMinBnb", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSetRedirectAddressIterator is returned from FilterSetRedirectAddress and is used to iterate over the raw logs and unpacked data for SetRedirectAddress events raised by the Stakemanager contract.
type StakemanagerSetRedirectAddressIterator struct {
	Event *StakemanagerSetRedirectAddress // Event containing the contract specifics and raw log

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
func (it *StakemanagerSetRedirectAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSetRedirectAddress)
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
		it.Event = new(StakemanagerSetRedirectAddress)
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
func (it *StakemanagerSetRedirectAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSetRedirectAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSetRedirectAddress represents a SetRedirectAddress event raised by the Stakemanager contract.
type StakemanagerSetRedirectAddress struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetRedirectAddress is a free log retrieval operation binding the contract event 0xefe3f2a813355e81025c75f6f977355ddd6f34f1fc3a15bd53562869af27c89c.
//
// Solidity: event SetRedirectAddress(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) FilterSetRedirectAddress(opts *bind.FilterOpts, _address []common.Address) (*StakemanagerSetRedirectAddressIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SetRedirectAddress", _addressRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerSetRedirectAddressIterator{contract: _Stakemanager.contract, event: "SetRedirectAddress", logs: logs, sub: sub}, nil
}

// WatchSetRedirectAddress is a free log subscription operation binding the contract event 0xefe3f2a813355e81025c75f6f977355ddd6f34f1fc3a15bd53562869af27c89c.
//
// Solidity: event SetRedirectAddress(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) WatchSetRedirectAddress(opts *bind.WatchOpts, sink chan<- *StakemanagerSetRedirectAddress, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SetRedirectAddress", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSetRedirectAddress)
				if err := _Stakemanager.contract.UnpackLog(event, "SetRedirectAddress", log); err != nil {
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

// ParseSetRedirectAddress is a log parse operation binding the contract event 0xefe3f2a813355e81025c75f6f977355ddd6f34f1fc3a15bd53562869af27c89c.
//
// Solidity: event SetRedirectAddress(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) ParseSetRedirectAddress(log types.Log) (*StakemanagerSetRedirectAddress, error) {
	event := new(StakemanagerSetRedirectAddress)
	if err := _Stakemanager.contract.UnpackLog(event, "SetRedirectAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSetReserveAmountIterator is returned from FilterSetReserveAmount and is used to iterate over the raw logs and unpacked data for SetReserveAmount events raised by the Stakemanager contract.
type StakemanagerSetReserveAmountIterator struct {
	Event *StakemanagerSetReserveAmount // Event containing the contract specifics and raw log

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
func (it *StakemanagerSetReserveAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSetReserveAmount)
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
		it.Event = new(StakemanagerSetReserveAmount)
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
func (it *StakemanagerSetReserveAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSetReserveAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSetReserveAmount represents a SetReserveAmount event raised by the Stakemanager contract.
type StakemanagerSetReserveAmount struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetReserveAmount is a free log retrieval operation binding the contract event 0xeebfaa20317a50dfda61faa034ecb2c646ab3e114ccbbb2f691f5dcef0687f66.
//
// Solidity: event SetReserveAmount(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterSetReserveAmount(opts *bind.FilterOpts) (*StakemanagerSetReserveAmountIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SetReserveAmount")
	if err != nil {
		return nil, err
	}
	return &StakemanagerSetReserveAmountIterator{contract: _Stakemanager.contract, event: "SetReserveAmount", logs: logs, sub: sub}, nil
}

// WatchSetReserveAmount is a free log subscription operation binding the contract event 0xeebfaa20317a50dfda61faa034ecb2c646ab3e114ccbbb2f691f5dcef0687f66.
//
// Solidity: event SetReserveAmount(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchSetReserveAmount(opts *bind.WatchOpts, sink chan<- *StakemanagerSetReserveAmount) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SetReserveAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSetReserveAmount)
				if err := _Stakemanager.contract.UnpackLog(event, "SetReserveAmount", log); err != nil {
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

// ParseSetReserveAmount is a log parse operation binding the contract event 0xeebfaa20317a50dfda61faa034ecb2c646ab3e114ccbbb2f691f5dcef0687f66.
//
// Solidity: event SetReserveAmount(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseSetReserveAmount(log types.Log) (*StakemanagerSetReserveAmount, error) {
	event := new(StakemanagerSetReserveAmount)
	if err := _Stakemanager.contract.UnpackLog(event, "SetReserveAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSetRevenuePoolIterator is returned from FilterSetRevenuePool and is used to iterate over the raw logs and unpacked data for SetRevenuePool events raised by the Stakemanager contract.
type StakemanagerSetRevenuePoolIterator struct {
	Event *StakemanagerSetRevenuePool // Event containing the contract specifics and raw log

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
func (it *StakemanagerSetRevenuePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSetRevenuePool)
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
		it.Event = new(StakemanagerSetRevenuePool)
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
func (it *StakemanagerSetRevenuePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSetRevenuePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSetRevenuePool represents a SetRevenuePool event raised by the Stakemanager contract.
type StakemanagerSetRevenuePool struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetRevenuePool is a free log retrieval operation binding the contract event 0xf980d53d1bddd28d9ae609f333a2a294fa8f486a11bb8ddb03b71f5f17429011.
//
// Solidity: event SetRevenuePool(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) FilterSetRevenuePool(opts *bind.FilterOpts, _address []common.Address) (*StakemanagerSetRevenuePoolIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SetRevenuePool", _addressRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerSetRevenuePoolIterator{contract: _Stakemanager.contract, event: "SetRevenuePool", logs: logs, sub: sub}, nil
}

// WatchSetRevenuePool is a free log subscription operation binding the contract event 0xf980d53d1bddd28d9ae609f333a2a294fa8f486a11bb8ddb03b71f5f17429011.
//
// Solidity: event SetRevenuePool(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) WatchSetRevenuePool(opts *bind.WatchOpts, sink chan<- *StakemanagerSetRevenuePool, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SetRevenuePool", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSetRevenuePool)
				if err := _Stakemanager.contract.UnpackLog(event, "SetRevenuePool", log); err != nil {
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

// ParseSetRevenuePool is a log parse operation binding the contract event 0xf980d53d1bddd28d9ae609f333a2a294fa8f486a11bb8ddb03b71f5f17429011.
//
// Solidity: event SetRevenuePool(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) ParseSetRevenuePool(log types.Log) (*StakemanagerSetRevenuePool, error) {
	event := new(StakemanagerSetRevenuePool)
	if err := _Stakemanager.contract.UnpackLog(event, "SetRevenuePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSetSynFeeIterator is returned from FilterSetSynFee and is used to iterate over the raw logs and unpacked data for SetSynFee events raised by the Stakemanager contract.
type StakemanagerSetSynFeeIterator struct {
	Event *StakemanagerSetSynFee // Event containing the contract specifics and raw log

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
func (it *StakemanagerSetSynFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSetSynFee)
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
		it.Event = new(StakemanagerSetSynFee)
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
func (it *StakemanagerSetSynFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSetSynFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSetSynFee represents a SetSynFee event raised by the Stakemanager contract.
type StakemanagerSetSynFee struct {
	SynFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetSynFee is a free log retrieval operation binding the contract event 0xc42638714e4dadf46d57ed62b1ef9805f061d57fc97d4d5bed85d6d420363545.
//
// Solidity: event SetSynFee(uint256 _synFee)
func (_Stakemanager *StakemanagerFilterer) FilterSetSynFee(opts *bind.FilterOpts) (*StakemanagerSetSynFeeIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SetSynFee")
	if err != nil {
		return nil, err
	}
	return &StakemanagerSetSynFeeIterator{contract: _Stakemanager.contract, event: "SetSynFee", logs: logs, sub: sub}, nil
}

// WatchSetSynFee is a free log subscription operation binding the contract event 0xc42638714e4dadf46d57ed62b1ef9805f061d57fc97d4d5bed85d6d420363545.
//
// Solidity: event SetSynFee(uint256 _synFee)
func (_Stakemanager *StakemanagerFilterer) WatchSetSynFee(opts *bind.WatchOpts, sink chan<- *StakemanagerSetSynFee) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SetSynFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSetSynFee)
				if err := _Stakemanager.contract.UnpackLog(event, "SetSynFee", log); err != nil {
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

// ParseSetSynFee is a log parse operation binding the contract event 0xc42638714e4dadf46d57ed62b1ef9805f061d57fc97d4d5bed85d6d420363545.
//
// Solidity: event SetSynFee(uint256 _synFee)
func (_Stakemanager *StakemanagerFilterer) ParseSetSynFee(log types.Log) (*StakemanagerSetSynFee, error) {
	event := new(StakemanagerSetSynFee)
	if err := _Stakemanager.contract.UnpackLog(event, "SetSynFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerSyncCreditContractIterator is returned from FilterSyncCreditContract and is used to iterate over the raw logs and unpacked data for SyncCreditContract events raised by the Stakemanager contract.
type StakemanagerSyncCreditContractIterator struct {
	Event *StakemanagerSyncCreditContract // Event containing the contract specifics and raw log

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
func (it *StakemanagerSyncCreditContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerSyncCreditContract)
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
		it.Event = new(StakemanagerSyncCreditContract)
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
func (it *StakemanagerSyncCreditContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerSyncCreditContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerSyncCreditContract represents a SyncCreditContract event raised by the Stakemanager contract.
type StakemanagerSyncCreditContract struct {
	Validator common.Address
	Credit    common.Address
	ToRemove  bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSyncCreditContract is a free log retrieval operation binding the contract event 0xd3a00d1560ea1bb8a1e52e816d05216a00e3b40fe584676b33f5f04b6b01a84a.
//
// Solidity: event SyncCreditContract(address indexed _validator, address _credit, bool toRemove)
func (_Stakemanager *StakemanagerFilterer) FilterSyncCreditContract(opts *bind.FilterOpts, _validator []common.Address) (*StakemanagerSyncCreditContractIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "SyncCreditContract", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerSyncCreditContractIterator{contract: _Stakemanager.contract, event: "SyncCreditContract", logs: logs, sub: sub}, nil
}

// WatchSyncCreditContract is a free log subscription operation binding the contract event 0xd3a00d1560ea1bb8a1e52e816d05216a00e3b40fe584676b33f5f04b6b01a84a.
//
// Solidity: event SyncCreditContract(address indexed _validator, address _credit, bool toRemove)
func (_Stakemanager *StakemanagerFilterer) WatchSyncCreditContract(opts *bind.WatchOpts, sink chan<- *StakemanagerSyncCreditContract, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "SyncCreditContract", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerSyncCreditContract)
				if err := _Stakemanager.contract.UnpackLog(event, "SyncCreditContract", log); err != nil {
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

// ParseSyncCreditContract is a log parse operation binding the contract event 0xd3a00d1560ea1bb8a1e52e816d05216a00e3b40fe584676b33f5f04b6b01a84a.
//
// Solidity: event SyncCreditContract(address indexed _validator, address _credit, bool toRemove)
func (_Stakemanager *StakemanagerFilterer) ParseSyncCreditContract(log types.Log) (*StakemanagerSyncCreditContract, error) {
	event := new(StakemanagerSyncCreditContract)
	if err := _Stakemanager.contract.UnpackLog(event, "SyncCreditContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerUndelegateIterator is returned from FilterUndelegate and is used to iterate over the raw logs and unpacked data for Undelegate events raised by the Stakemanager contract.
type StakemanagerUndelegateIterator struct {
	Event *StakemanagerUndelegate // Event containing the contract specifics and raw log

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
func (it *StakemanagerUndelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUndelegate)
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
		it.Event = new(StakemanagerUndelegate)
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
func (it *StakemanagerUndelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUndelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUndelegate represents a Undelegate event raised by the Stakemanager contract.
type StakemanagerUndelegate struct {
	NextUndelegatedRequestIndex *big.Int
	BnbAmount                   *big.Int
	Shares                      *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterUndelegate is a free log retrieval operation binding the contract event 0xd519d4398e08424de0b208ab46cf899a372bb182bdd87e7d4b49e7cc1fcacd2c.
//
// Solidity: event Undelegate(uint256 _nextUndelegatedRequestIndex, uint256 _bnbAmount, uint256 _shares)
func (_Stakemanager *StakemanagerFilterer) FilterUndelegate(opts *bind.FilterOpts) (*StakemanagerUndelegateIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Undelegate")
	if err != nil {
		return nil, err
	}
	return &StakemanagerUndelegateIterator{contract: _Stakemanager.contract, event: "Undelegate", logs: logs, sub: sub}, nil
}

// WatchUndelegate is a free log subscription operation binding the contract event 0xd519d4398e08424de0b208ab46cf899a372bb182bdd87e7d4b49e7cc1fcacd2c.
//
// Solidity: event Undelegate(uint256 _nextUndelegatedRequestIndex, uint256 _bnbAmount, uint256 _shares)
func (_Stakemanager *StakemanagerFilterer) WatchUndelegate(opts *bind.WatchOpts, sink chan<- *StakemanagerUndelegate) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Undelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUndelegate)
				if err := _Stakemanager.contract.UnpackLog(event, "Undelegate", log); err != nil {
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

// ParseUndelegate is a log parse operation binding the contract event 0xd519d4398e08424de0b208ab46cf899a372bb182bdd87e7d4b49e7cc1fcacd2c.
//
// Solidity: event Undelegate(uint256 _nextUndelegatedRequestIndex, uint256 _bnbAmount, uint256 _shares)
func (_Stakemanager *StakemanagerFilterer) ParseUndelegate(log types.Log) (*StakemanagerUndelegate, error) {
	event := new(StakemanagerUndelegate)
	if err := _Stakemanager.contract.UnpackLog(event, "Undelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerUndelegateFromIterator is returned from FilterUndelegateFrom and is used to iterate over the raw logs and unpacked data for UndelegateFrom events raised by the Stakemanager contract.
type StakemanagerUndelegateFromIterator struct {
	Event *StakemanagerUndelegateFrom // Event containing the contract specifics and raw log

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
func (it *StakemanagerUndelegateFromIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUndelegateFrom)
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
		it.Event = new(StakemanagerUndelegateFrom)
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
func (it *StakemanagerUndelegateFromIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUndelegateFromIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUndelegateFrom represents a UndelegateFrom event raised by the Stakemanager contract.
type StakemanagerUndelegateFrom struct {
	Operator  common.Address
	BnbAmount *big.Int
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateFrom is a free log retrieval operation binding the contract event 0x0190f399b2c3dcb444d981796081189b4b1f466824f5a82f2bf53713ac5b427b.
//
// Solidity: event UndelegateFrom(address indexed _operator, uint256 _bnbAmount, uint256 _shares)
func (_Stakemanager *StakemanagerFilterer) FilterUndelegateFrom(opts *bind.FilterOpts, _operator []common.Address) (*StakemanagerUndelegateFromIterator, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "UndelegateFrom", _operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerUndelegateFromIterator{contract: _Stakemanager.contract, event: "UndelegateFrom", logs: logs, sub: sub}, nil
}

// WatchUndelegateFrom is a free log subscription operation binding the contract event 0x0190f399b2c3dcb444d981796081189b4b1f466824f5a82f2bf53713ac5b427b.
//
// Solidity: event UndelegateFrom(address indexed _operator, uint256 _bnbAmount, uint256 _shares)
func (_Stakemanager *StakemanagerFilterer) WatchUndelegateFrom(opts *bind.WatchOpts, sink chan<- *StakemanagerUndelegateFrom, _operator []common.Address) (event.Subscription, error) {

	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "UndelegateFrom", _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUndelegateFrom)
				if err := _Stakemanager.contract.UnpackLog(event, "UndelegateFrom", log); err != nil {
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

// ParseUndelegateFrom is a log parse operation binding the contract event 0x0190f399b2c3dcb444d981796081189b4b1f466824f5a82f2bf53713ac5b427b.
//
// Solidity: event UndelegateFrom(address indexed _operator, uint256 _bnbAmount, uint256 _shares)
func (_Stakemanager *StakemanagerFilterer) ParseUndelegateFrom(log types.Log) (*StakemanagerUndelegateFrom, error) {
	event := new(StakemanagerUndelegateFrom)
	if err := _Stakemanager.contract.UnpackLog(event, "UndelegateFrom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerUndelegateReserveIterator is returned from FilterUndelegateReserve and is used to iterate over the raw logs and unpacked data for UndelegateReserve events raised by the Stakemanager contract.
type StakemanagerUndelegateReserveIterator struct {
	Event *StakemanagerUndelegateReserve // Event containing the contract specifics and raw log

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
func (it *StakemanagerUndelegateReserveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUndelegateReserve)
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
		it.Event = new(StakemanagerUndelegateReserve)
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
func (it *StakemanagerUndelegateReserveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUndelegateReserveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUndelegateReserve represents a UndelegateReserve event raised by the Stakemanager contract.
type StakemanagerUndelegateReserve struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUndelegateReserve is a free log retrieval operation binding the contract event 0xa7e208ae65c6023c9e9428ec1501e02e7ef7f82d904c41dcb9f8a066d21b1cd8.
//
// Solidity: event UndelegateReserve(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) FilterUndelegateReserve(opts *bind.FilterOpts) (*StakemanagerUndelegateReserveIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "UndelegateReserve")
	if err != nil {
		return nil, err
	}
	return &StakemanagerUndelegateReserveIterator{contract: _Stakemanager.contract, event: "UndelegateReserve", logs: logs, sub: sub}, nil
}

// WatchUndelegateReserve is a free log subscription operation binding the contract event 0xa7e208ae65c6023c9e9428ec1501e02e7ef7f82d904c41dcb9f8a066d21b1cd8.
//
// Solidity: event UndelegateReserve(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) WatchUndelegateReserve(opts *bind.WatchOpts, sink chan<- *StakemanagerUndelegateReserve) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "UndelegateReserve")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUndelegateReserve)
				if err := _Stakemanager.contract.UnpackLog(event, "UndelegateReserve", log); err != nil {
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

// ParseUndelegateReserve is a log parse operation binding the contract event 0xa7e208ae65c6023c9e9428ec1501e02e7ef7f82d904c41dcb9f8a066d21b1cd8.
//
// Solidity: event UndelegateReserve(uint256 _amount)
func (_Stakemanager *StakemanagerFilterer) ParseUndelegateReserve(log types.Log) (*StakemanagerUndelegateReserve, error) {
	event := new(StakemanagerUndelegateReserve)
	if err := _Stakemanager.contract.UnpackLog(event, "UndelegateReserve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Stakemanager contract.
type StakemanagerUnpausedIterator struct {
	Event *StakemanagerUnpaused // Event containing the contract specifics and raw log

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
func (it *StakemanagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUnpaused)
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
		it.Event = new(StakemanagerUnpaused)
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
func (it *StakemanagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUnpaused represents a Unpaused event raised by the Stakemanager contract.
type StakemanagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Stakemanager *StakemanagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakemanagerUnpausedIterator, error) {

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakemanagerUnpausedIterator{contract: _Stakemanager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Stakemanager *StakemanagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakemanagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUnpaused)
				if err := _Stakemanager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Stakemanager *StakemanagerFilterer) ParseUnpaused(log types.Log) (*StakemanagerUnpaused, error) {
	event := new(StakemanagerUnpaused)
	if err := _Stakemanager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerWhitelistValidatorIterator is returned from FilterWhitelistValidator and is used to iterate over the raw logs and unpacked data for WhitelistValidator events raised by the Stakemanager contract.
type StakemanagerWhitelistValidatorIterator struct {
	Event *StakemanagerWhitelistValidator // Event containing the contract specifics and raw log

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
func (it *StakemanagerWhitelistValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerWhitelistValidator)
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
		it.Event = new(StakemanagerWhitelistValidator)
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
func (it *StakemanagerWhitelistValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerWhitelistValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerWhitelistValidator represents a WhitelistValidator event raised by the Stakemanager contract.
type StakemanagerWhitelistValidator struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWhitelistValidator is a free log retrieval operation binding the contract event 0x5ce1ed13ab7bf8de942f6a4d7ef859be025626a486b20a66ea2e20d771a5d0ee.
//
// Solidity: event WhitelistValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) FilterWhitelistValidator(opts *bind.FilterOpts, _address []common.Address) (*StakemanagerWhitelistValidatorIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "WhitelistValidator", _addressRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerWhitelistValidatorIterator{contract: _Stakemanager.contract, event: "WhitelistValidator", logs: logs, sub: sub}, nil
}

// WatchWhitelistValidator is a free log subscription operation binding the contract event 0x5ce1ed13ab7bf8de942f6a4d7ef859be025626a486b20a66ea2e20d771a5d0ee.
//
// Solidity: event WhitelistValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) WatchWhitelistValidator(opts *bind.WatchOpts, sink chan<- *StakemanagerWhitelistValidator, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "WhitelistValidator", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerWhitelistValidator)
				if err := _Stakemanager.contract.UnpackLog(event, "WhitelistValidator", log); err != nil {
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

// ParseWhitelistValidator is a log parse operation binding the contract event 0x5ce1ed13ab7bf8de942f6a4d7ef859be025626a486b20a66ea2e20d771a5d0ee.
//
// Solidity: event WhitelistValidator(address indexed _address)
func (_Stakemanager *StakemanagerFilterer) ParseWhitelistValidator(log types.Log) (*StakemanagerWhitelistValidator, error) {
	event := new(StakemanagerWhitelistValidator)
	if err := _Stakemanager.contract.UnpackLog(event, "WhitelistValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
