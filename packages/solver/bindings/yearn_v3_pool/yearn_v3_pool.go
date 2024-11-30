// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearn_v3_pool

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Activation  *big.Int
	LastReport  *big.Int
	CurrentDebt *big.Int
	MaxDebt     *big.Int
}

// YearnV3PoolMetaData contains all meta data concerning the YearnV3Pool contract.
var YearnV3PoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"change_type\",\"type\":\"uint256\"}],\"name\":\"StrategyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"gain\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"current_debt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"protocol_fees\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"total_fees\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"total_refunds\",\"type\":\"uint256\"}],\"name\":\"StrategyReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"current_debt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"new_debt\",\"type\":\"uint256\"}],\"name\":\"DebtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"RoleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"role_manager\",\"type\":\"address\"}],\"name\":\"UpdateRoleManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"accountant\",\"type\":\"address\"}],\"name\":\"UpdateAccountant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"deposit_limit_module\",\"type\":\"address\"}],\"name\":\"UpdateDepositLimitModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"withdraw_limit_module\",\"type\":\"address\"}],\"name\":\"UpdateWithdrawLimitModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"new_default_queue\",\"type\":\"address[]\"}],\"name\":\"UpdateDefaultQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"use_default_queue\",\"type\":\"bool\"}],\"name\":\"UpdateUseDefaultQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"new_debt\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxDebtForStrategy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"deposit_limit\",\"type\":\"uint256\"}],\"name\":\"UpdateDepositLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minimum_total_idle\",\"type\":\"uint256\"}],\"name\":\"UpdateMinimumTotalIdle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\"}],\"name\":\"UpdateProfitMaxUnlockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebtPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Shutdown\",\"type\":\"event\"},{\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"role_manager\",\"type\":\"address\"},{\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_accountant\",\"type\":\"address\"}],\"name\":\"set_accountant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_default_queue\",\"type\":\"address[]\"}],\"name\":\"set_default_queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"use_default_queue\",\"type\":\"bool\"}],\"name\":\"set_use_default_queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\"}],\"name\":\"set_deposit_limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\"},{\"name\":\"override\",\"type\":\"bool\"}],\"name\":\"set_deposit_limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\"}],\"name\":\"set_deposit_limit_module\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\"},{\"name\":\"override\",\"type\":\"bool\"}],\"name\":\"set_deposit_limit_module\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"withdraw_limit_module\",\"type\":\"address\"}],\"name\":\"set_withdraw_limit_module\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"minimum_total_idle\",\"type\":\"uint256\"}],\"name\":\"set_minimum_total_idle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_profit_max_unlock_time\",\"type\":\"uint256\"}],\"name\":\"setProfitMaxUnlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"set_role\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"add_role\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"remove_role\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"role_manager\",\"type\":\"address\"}],\"name\":\"transfer_role_manager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accept_role_manager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockedShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_default_queue\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"process_report\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy_debt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_strategy\",\"type\":\"address\"}],\"name\":\"add_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_strategy\",\"type\":\"address\"},{\"name\":\"add_to_queue\",\"type\":\"bool\"}],\"name\":\"add_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"revoke_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"force_revoke_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"new_max_debt\",\"type\":\"uint256\"}],\"name\":\"update_max_debt_for_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"target_debt\",\"type\":\"uint256\"}],\"name\":\"update_debt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"target_debt\",\"type\":\"uint256\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"update_debt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdown_vault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalIdle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDebt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"assets_needed\",\"type\":\"uint256\"}],\"name\":\"assess_share_of_unrealised_losses\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitMaxUnlockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fullProfitUnlockDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitUnlockingRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProfitUpdate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"strategies\",\"outputs\":[{\"components\":[{\"name\":\"activation\",\"type\":\"uint256\"},{\"name\":\"last_report\",\"type\":\"uint256\"},{\"name\":\"current_debt\",\"type\":\"uint256\"},{\"name\":\"max_debt\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"default_queue\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"use_default_queue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimum_total_idle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit_limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountant\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit_limit_module\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw_limit_module\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"role_manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_role_manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YearnV3PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnV3PoolMetaData.ABI instead.
var YearnV3PoolABI = YearnV3PoolMetaData.ABI

// YearnV3Pool is an auto generated Go binding around an Ethereum contract.
type YearnV3Pool struct {
	YearnV3PoolCaller     // Read-only binding to the contract
	YearnV3PoolTransactor // Write-only binding to the contract
	YearnV3PoolFilterer   // Log filterer for contract events
}

// YearnV3PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnV3PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnV3PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnV3PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnV3PoolSession struct {
	Contract     *YearnV3Pool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnV3PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnV3PoolCallerSession struct {
	Contract *YearnV3PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// YearnV3PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnV3PoolTransactorSession struct {
	Contract     *YearnV3PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YearnV3PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnV3PoolRaw struct {
	Contract *YearnV3Pool // Generic contract binding to access the raw methods on
}

// YearnV3PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnV3PoolCallerRaw struct {
	Contract *YearnV3PoolCaller // Generic read-only contract binding to access the raw methods on
}

// YearnV3PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnV3PoolTransactorRaw struct {
	Contract *YearnV3PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnV3Pool creates a new instance of YearnV3Pool, bound to a specific deployed contract.
func NewYearnV3Pool(address common.Address, backend bind.ContractBackend) (*YearnV3Pool, error) {
	contract, err := bindYearnV3Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnV3Pool{YearnV3PoolCaller: YearnV3PoolCaller{contract: contract}, YearnV3PoolTransactor: YearnV3PoolTransactor{contract: contract}, YearnV3PoolFilterer: YearnV3PoolFilterer{contract: contract}}, nil
}

// NewYearnV3PoolCaller creates a new read-only instance of YearnV3Pool, bound to a specific deployed contract.
func NewYearnV3PoolCaller(address common.Address, caller bind.ContractCaller) (*YearnV3PoolCaller, error) {
	contract, err := bindYearnV3Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolCaller{contract: contract}, nil
}

// NewYearnV3PoolTransactor creates a new write-only instance of YearnV3Pool, bound to a specific deployed contract.
func NewYearnV3PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnV3PoolTransactor, error) {
	contract, err := bindYearnV3Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolTransactor{contract: contract}, nil
}

// NewYearnV3PoolFilterer creates a new log filterer instance of YearnV3Pool, bound to a specific deployed contract.
func NewYearnV3PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnV3PoolFilterer, error) {
	contract, err := bindYearnV3Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolFilterer{contract: contract}, nil
}

// bindYearnV3Pool binds a generic wrapper to an already deployed contract.
func bindYearnV3Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YearnV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnV3Pool *YearnV3PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnV3Pool.Contract.YearnV3PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnV3Pool *YearnV3PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.YearnV3PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnV3Pool *YearnV3PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.YearnV3PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnV3Pool *YearnV3PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnV3Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnV3Pool *YearnV3PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnV3Pool *YearnV3PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnV3Pool *YearnV3PoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnV3Pool *YearnV3PoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YearnV3Pool.Contract.DOMAINSEPARATOR(&_YearnV3Pool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnV3Pool *YearnV3PoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YearnV3Pool.Contract.DOMAINSEPARATOR(&_YearnV3Pool.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YearnV3Pool *YearnV3PoolCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YearnV3Pool *YearnV3PoolSession) FACTORY() (common.Address, error) {
	return _YearnV3Pool.Contract.FACTORY(&_YearnV3Pool.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YearnV3Pool *YearnV3PoolCallerSession) FACTORY() (common.Address, error) {
	return _YearnV3Pool.Contract.FACTORY(&_YearnV3Pool.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_YearnV3Pool *YearnV3PoolCaller) Accountant(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "accountant")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_YearnV3Pool *YearnV3PoolSession) Accountant() (common.Address, error) {
	return _YearnV3Pool.Contract.Accountant(&_YearnV3Pool.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_YearnV3Pool *YearnV3PoolCallerSession) Accountant() (common.Address, error) {
	return _YearnV3Pool.Contract.Accountant(&_YearnV3Pool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.Allowance(&_YearnV3Pool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.Allowance(&_YearnV3Pool.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YearnV3Pool *YearnV3PoolCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YearnV3Pool *YearnV3PoolSession) ApiVersion() (string, error) {
	return _YearnV3Pool.Contract.ApiVersion(&_YearnV3Pool.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YearnV3Pool *YearnV3PoolCallerSession) ApiVersion() (string, error) {
	return _YearnV3Pool.Contract.ApiVersion(&_YearnV3Pool.CallOpts)
}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) AssessShareOfUnrealisedLosses(opts *bind.CallOpts, strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "assess_share_of_unrealised_losses", strategy, assets_needed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) AssessShareOfUnrealisedLosses(strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.AssessShareOfUnrealisedLosses(&_YearnV3Pool.CallOpts, strategy, assets_needed)
}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) AssessShareOfUnrealisedLosses(strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.AssessShareOfUnrealisedLosses(&_YearnV3Pool.CallOpts, strategy, assets_needed)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnV3Pool *YearnV3PoolCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnV3Pool *YearnV3PoolSession) Asset() (common.Address, error) {
	return _YearnV3Pool.Contract.Asset(&_YearnV3Pool.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnV3Pool *YearnV3PoolCallerSession) Asset() (common.Address, error) {
	return _YearnV3Pool.Contract.Asset(&_YearnV3Pool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.BalanceOf(&_YearnV3Pool.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.BalanceOf(&_YearnV3Pool.CallOpts, addr)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.ConvertToAssets(&_YearnV3Pool.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.ConvertToAssets(&_YearnV3Pool.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.ConvertToShares(&_YearnV3Pool.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.ConvertToShares(&_YearnV3Pool.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnV3Pool *YearnV3PoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnV3Pool *YearnV3PoolSession) Decimals() (uint8, error) {
	return _YearnV3Pool.Contract.Decimals(&_YearnV3Pool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnV3Pool *YearnV3PoolCallerSession) Decimals() (uint8, error) {
	return _YearnV3Pool.Contract.Decimals(&_YearnV3Pool.CallOpts)
}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_YearnV3Pool *YearnV3PoolCaller) DefaultQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "default_queue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_YearnV3Pool *YearnV3PoolSession) DefaultQueue(arg0 *big.Int) (common.Address, error) {
	return _YearnV3Pool.Contract.DefaultQueue(&_YearnV3Pool.CallOpts, arg0)
}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_YearnV3Pool *YearnV3PoolCallerSession) DefaultQueue(arg0 *big.Int) (common.Address, error) {
	return _YearnV3Pool.Contract.DefaultQueue(&_YearnV3Pool.CallOpts, arg0)
}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "deposit_limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) DepositLimit() (*big.Int, error) {
	return _YearnV3Pool.Contract.DepositLimit(&_YearnV3Pool.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) DepositLimit() (*big.Int, error) {
	return _YearnV3Pool.Contract.DepositLimit(&_YearnV3Pool.CallOpts)
}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_YearnV3Pool *YearnV3PoolCaller) DepositLimitModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "deposit_limit_module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_YearnV3Pool *YearnV3PoolSession) DepositLimitModule() (common.Address, error) {
	return _YearnV3Pool.Contract.DepositLimitModule(&_YearnV3Pool.CallOpts)
}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_YearnV3Pool *YearnV3PoolCallerSession) DepositLimitModule() (common.Address, error) {
	return _YearnV3Pool.Contract.DepositLimitModule(&_YearnV3Pool.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) FullProfitUnlockDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "fullProfitUnlockDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) FullProfitUnlockDate() (*big.Int, error) {
	return _YearnV3Pool.Contract.FullProfitUnlockDate(&_YearnV3Pool.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) FullProfitUnlockDate() (*big.Int, error) {
	return _YearnV3Pool.Contract.FullProfitUnlockDate(&_YearnV3Pool.CallOpts)
}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_YearnV3Pool *YearnV3PoolCaller) FutureRoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "future_role_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_YearnV3Pool *YearnV3PoolSession) FutureRoleManager() (common.Address, error) {
	return _YearnV3Pool.Contract.FutureRoleManager(&_YearnV3Pool.CallOpts)
}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_YearnV3Pool *YearnV3PoolCallerSession) FutureRoleManager() (common.Address, error) {
	return _YearnV3Pool.Contract.FutureRoleManager(&_YearnV3Pool.CallOpts)
}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_YearnV3Pool *YearnV3PoolCaller) GetDefaultQueue(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "get_default_queue")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_YearnV3Pool *YearnV3PoolSession) GetDefaultQueue() ([]common.Address, error) {
	return _YearnV3Pool.Contract.GetDefaultQueue(&_YearnV3Pool.CallOpts)
}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_YearnV3Pool *YearnV3PoolCallerSession) GetDefaultQueue() ([]common.Address, error) {
	return _YearnV3Pool.Contract.GetDefaultQueue(&_YearnV3Pool.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YearnV3Pool *YearnV3PoolCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YearnV3Pool *YearnV3PoolSession) IsShutdown() (bool, error) {
	return _YearnV3Pool.Contract.IsShutdown(&_YearnV3Pool.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YearnV3Pool *YearnV3PoolCallerSession) IsShutdown() (bool, error) {
	return _YearnV3Pool.Contract.IsShutdown(&_YearnV3Pool.CallOpts)
}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) LastProfitUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "lastProfitUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) LastProfitUpdate() (*big.Int, error) {
	return _YearnV3Pool.Contract.LastProfitUpdate(&_YearnV3Pool.CallOpts)
}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) LastProfitUpdate() (*big.Int, error) {
	return _YearnV3Pool.Contract.LastProfitUpdate(&_YearnV3Pool.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MaxDeposit(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "maxDeposit", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxDeposit(&_YearnV3Pool.CallOpts, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxDeposit(&_YearnV3Pool.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxMint(&_YearnV3Pool.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxMint(&_YearnV3Pool.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxRedeem(&_YearnV3Pool.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxRedeem(&_YearnV3Pool.CallOpts, owner)
}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MaxRedeem0(opts *bind.CallOpts, owner common.Address, max_loss *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "maxRedeem0", owner, max_loss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MaxRedeem0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxRedeem0(&_YearnV3Pool.CallOpts, owner, max_loss)
}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MaxRedeem0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxRedeem0(&_YearnV3Pool.CallOpts, owner, max_loss)
}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MaxRedeem1(opts *bind.CallOpts, owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "maxRedeem1", owner, max_loss, strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MaxRedeem1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxRedeem1(&_YearnV3Pool.CallOpts, owner, max_loss, strategies)
}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MaxRedeem1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxRedeem1(&_YearnV3Pool.CallOpts, owner, max_loss, strategies)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxWithdraw(&_YearnV3Pool.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxWithdraw(&_YearnV3Pool.CallOpts, owner)
}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MaxWithdraw0(opts *bind.CallOpts, owner common.Address, max_loss *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "maxWithdraw0", owner, max_loss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MaxWithdraw0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxWithdraw0(&_YearnV3Pool.CallOpts, owner, max_loss)
}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MaxWithdraw0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxWithdraw0(&_YearnV3Pool.CallOpts, owner, max_loss)
}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MaxWithdraw1(opts *bind.CallOpts, owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "maxWithdraw1", owner, max_loss, strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MaxWithdraw1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxWithdraw1(&_YearnV3Pool.CallOpts, owner, max_loss, strategies)
}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MaxWithdraw1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.MaxWithdraw1(&_YearnV3Pool.CallOpts, owner, max_loss, strategies)
}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) MinimumTotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "minimum_total_idle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) MinimumTotalIdle() (*big.Int, error) {
	return _YearnV3Pool.Contract.MinimumTotalIdle(&_YearnV3Pool.CallOpts)
}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) MinimumTotalIdle() (*big.Int, error) {
	return _YearnV3Pool.Contract.MinimumTotalIdle(&_YearnV3Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Pool *YearnV3PoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Pool *YearnV3PoolSession) Name() (string, error) {
	return _YearnV3Pool.Contract.Name(&_YearnV3Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Pool *YearnV3PoolCallerSession) Name() (string, error) {
	return _YearnV3Pool.Contract.Name(&_YearnV3Pool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.Nonces(&_YearnV3Pool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.Nonces(&_YearnV3Pool.CallOpts, arg0)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.PreviewDeposit(&_YearnV3Pool.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.PreviewDeposit(&_YearnV3Pool.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.PreviewMint(&_YearnV3Pool.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.PreviewMint(&_YearnV3Pool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.PreviewRedeem(&_YearnV3Pool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.PreviewRedeem(&_YearnV3Pool.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.PreviewWithdraw(&_YearnV3Pool.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _YearnV3Pool.Contract.PreviewWithdraw(&_YearnV3Pool.CallOpts, assets)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) PricePerShare() (*big.Int, error) {
	return _YearnV3Pool.Contract.PricePerShare(&_YearnV3Pool.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) PricePerShare() (*big.Int, error) {
	return _YearnV3Pool.Contract.PricePerShare(&_YearnV3Pool.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) ProfitMaxUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "profitMaxUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _YearnV3Pool.Contract.ProfitMaxUnlockTime(&_YearnV3Pool.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _YearnV3Pool.Contract.ProfitMaxUnlockTime(&_YearnV3Pool.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) ProfitUnlockingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "profitUnlockingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) ProfitUnlockingRate() (*big.Int, error) {
	return _YearnV3Pool.Contract.ProfitUnlockingRate(&_YearnV3Pool.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) ProfitUnlockingRate() (*big.Int, error) {
	return _YearnV3Pool.Contract.ProfitUnlockingRate(&_YearnV3Pool.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_YearnV3Pool *YearnV3PoolCaller) RoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "role_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_YearnV3Pool *YearnV3PoolSession) RoleManager() (common.Address, error) {
	return _YearnV3Pool.Contract.RoleManager(&_YearnV3Pool.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_YearnV3Pool *YearnV3PoolCallerSession) RoleManager() (common.Address, error) {
	return _YearnV3Pool.Contract.RoleManager(&_YearnV3Pool.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) Roles(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "roles", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Roles(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.Roles(&_YearnV3Pool.CallOpts, arg0)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) Roles(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Pool.Contract.Roles(&_YearnV3Pool.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_YearnV3Pool *YearnV3PoolCaller) Strategies(opts *bind.CallOpts, arg0 common.Address) (Struct0, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "strategies", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_YearnV3Pool *YearnV3PoolSession) Strategies(arg0 common.Address) (Struct0, error) {
	return _YearnV3Pool.Contract.Strategies(&_YearnV3Pool.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_YearnV3Pool *YearnV3PoolCallerSession) Strategies(arg0 common.Address) (Struct0, error) {
	return _YearnV3Pool.Contract.Strategies(&_YearnV3Pool.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnV3Pool *YearnV3PoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnV3Pool *YearnV3PoolSession) Symbol() (string, error) {
	return _YearnV3Pool.Contract.Symbol(&_YearnV3Pool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnV3Pool *YearnV3PoolCallerSession) Symbol() (string, error) {
	return _YearnV3Pool.Contract.Symbol(&_YearnV3Pool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) TotalAssets() (*big.Int, error) {
	return _YearnV3Pool.Contract.TotalAssets(&_YearnV3Pool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) TotalAssets() (*big.Int, error) {
	return _YearnV3Pool.Contract.TotalAssets(&_YearnV3Pool.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) TotalDebt() (*big.Int, error) {
	return _YearnV3Pool.Contract.TotalDebt(&_YearnV3Pool.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) TotalDebt() (*big.Int, error) {
	return _YearnV3Pool.Contract.TotalDebt(&_YearnV3Pool.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) TotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "totalIdle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) TotalIdle() (*big.Int, error) {
	return _YearnV3Pool.Contract.TotalIdle(&_YearnV3Pool.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) TotalIdle() (*big.Int, error) {
	return _YearnV3Pool.Contract.TotalIdle(&_YearnV3Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) TotalSupply() (*big.Int, error) {
	return _YearnV3Pool.Contract.TotalSupply(&_YearnV3Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) TotalSupply() (*big.Int, error) {
	return _YearnV3Pool.Contract.TotalSupply(&_YearnV3Pool.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCaller) UnlockedShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "unlockedShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) UnlockedShares() (*big.Int, error) {
	return _YearnV3Pool.Contract.UnlockedShares(&_YearnV3Pool.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_YearnV3Pool *YearnV3PoolCallerSession) UnlockedShares() (*big.Int, error) {
	return _YearnV3Pool.Contract.UnlockedShares(&_YearnV3Pool.CallOpts)
}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_YearnV3Pool *YearnV3PoolCaller) UseDefaultQueue(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "use_default_queue")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_YearnV3Pool *YearnV3PoolSession) UseDefaultQueue() (bool, error) {
	return _YearnV3Pool.Contract.UseDefaultQueue(&_YearnV3Pool.CallOpts)
}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_YearnV3Pool *YearnV3PoolCallerSession) UseDefaultQueue() (bool, error) {
	return _YearnV3Pool.Contract.UseDefaultQueue(&_YearnV3Pool.CallOpts)
}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_YearnV3Pool *YearnV3PoolCaller) WithdrawLimitModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Pool.contract.Call(opts, &out, "withdraw_limit_module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_YearnV3Pool *YearnV3PoolSession) WithdrawLimitModule() (common.Address, error) {
	return _YearnV3Pool.Contract.WithdrawLimitModule(&_YearnV3Pool.CallOpts)
}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_YearnV3Pool *YearnV3PoolCallerSession) WithdrawLimitModule() (common.Address, error) {
	return _YearnV3Pool.Contract.WithdrawLimitModule(&_YearnV3Pool.CallOpts)
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_YearnV3Pool *YearnV3PoolTransactor) AcceptRoleManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "accept_role_manager")
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_YearnV3Pool *YearnV3PoolSession) AcceptRoleManager() (*types.Transaction, error) {
	return _YearnV3Pool.Contract.AcceptRoleManager(&_YearnV3Pool.TransactOpts)
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) AcceptRoleManager() (*types.Transaction, error) {
	return _YearnV3Pool.Contract.AcceptRoleManager(&_YearnV3Pool.TransactOpts)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) AddRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "add_role", account, role)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolSession) AddRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.AddRole(&_YearnV3Pool.TransactOpts, account, role)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) AddRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.AddRole(&_YearnV3Pool.TransactOpts, account, role)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) AddStrategy(opts *bind.TransactOpts, new_strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "add_strategy", new_strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_YearnV3Pool *YearnV3PoolSession) AddStrategy(new_strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.AddStrategy(&_YearnV3Pool.TransactOpts, new_strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) AddStrategy(new_strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.AddStrategy(&_YearnV3Pool.TransactOpts, new_strategy)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) AddStrategy0(opts *bind.TransactOpts, new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "add_strategy0", new_strategy, add_to_queue)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_YearnV3Pool *YearnV3PoolSession) AddStrategy0(new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.AddStrategy0(&_YearnV3Pool.TransactOpts, new_strategy, add_to_queue)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) AddStrategy0(new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.AddStrategy0(&_YearnV3Pool.TransactOpts, new_strategy, add_to_queue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Approve(&_YearnV3Pool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Approve(&_YearnV3Pool.TransactOpts, spender, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) BuyDebt(opts *bind.TransactOpts, strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "buy_debt", strategy, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_YearnV3Pool *YearnV3PoolSession) BuyDebt(strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.BuyDebt(&_YearnV3Pool.TransactOpts, strategy, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) BuyDebt(strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.BuyDebt(&_YearnV3Pool.TransactOpts, strategy, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Deposit(&_YearnV3Pool.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Deposit(&_YearnV3Pool.TransactOpts, assets, receiver)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) ForceRevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "force_revoke_strategy", strategy)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_YearnV3Pool *YearnV3PoolSession) ForceRevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.ForceRevokeStrategy(&_YearnV3Pool.TransactOpts, strategy)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) ForceRevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.ForceRevokeStrategy(&_YearnV3Pool.TransactOpts, strategy)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) Initialize(opts *bind.TransactOpts, asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "initialize", asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_YearnV3Pool *YearnV3PoolSession) Initialize(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Initialize(&_YearnV3Pool.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) Initialize(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Initialize(&_YearnV3Pool.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Mint(&_YearnV3Pool.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Mint(&_YearnV3Pool.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_YearnV3Pool *YearnV3PoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "permit", owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_YearnV3Pool *YearnV3PoolSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Permit(&_YearnV3Pool.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Permit(&_YearnV3Pool.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) ProcessReport(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "process_report", strategy)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_YearnV3Pool *YearnV3PoolSession) ProcessReport(strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.ProcessReport(&_YearnV3Pool.TransactOpts, strategy)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) ProcessReport(strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.ProcessReport(&_YearnV3Pool.TransactOpts, strategy)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Redeem(&_YearnV3Pool.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Redeem(&_YearnV3Pool.TransactOpts, shares, receiver, owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) Redeem0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "redeem0", shares, receiver, owner, max_loss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Redeem0(&_YearnV3Pool.TransactOpts, shares, receiver, owner, max_loss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Redeem0(&_YearnV3Pool.TransactOpts, shares, receiver, owner, max_loss)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) Redeem1(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "redeem1", shares, receiver, owner, max_loss, strategies)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Redeem1(&_YearnV3Pool.TransactOpts, shares, receiver, owner, max_loss, strategies)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Redeem1(&_YearnV3Pool.TransactOpts, shares, receiver, owner, max_loss, strategies)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) RemoveRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "remove_role", account, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolSession) RemoveRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.RemoveRole(&_YearnV3Pool.TransactOpts, account, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) RemoveRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.RemoveRole(&_YearnV3Pool.TransactOpts, account, role)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) RevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "revoke_strategy", strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_YearnV3Pool *YearnV3PoolSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.RevokeStrategy(&_YearnV3Pool.TransactOpts, strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.RevokeStrategy(&_YearnV3Pool.TransactOpts, strategy)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetProfitMaxUnlockTime(opts *bind.TransactOpts, new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "setProfitMaxUnlockTime", new_profit_max_unlock_time)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetProfitMaxUnlockTime(new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetProfitMaxUnlockTime(&_YearnV3Pool.TransactOpts, new_profit_max_unlock_time)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetProfitMaxUnlockTime(new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetProfitMaxUnlockTime(&_YearnV3Pool.TransactOpts, new_profit_max_unlock_time)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetAccountant(opts *bind.TransactOpts, new_accountant common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_accountant", new_accountant)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetAccountant(new_accountant common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetAccountant(&_YearnV3Pool.TransactOpts, new_accountant)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetAccountant(new_accountant common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetAccountant(&_YearnV3Pool.TransactOpts, new_accountant)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetDefaultQueue(opts *bind.TransactOpts, new_default_queue []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_default_queue", new_default_queue)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetDefaultQueue(new_default_queue []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDefaultQueue(&_YearnV3Pool.TransactOpts, new_default_queue)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetDefaultQueue(new_default_queue []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDefaultQueue(&_YearnV3Pool.TransactOpts, new_default_queue)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetDepositLimit(opts *bind.TransactOpts, deposit_limit *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_deposit_limit", deposit_limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetDepositLimit(deposit_limit *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDepositLimit(&_YearnV3Pool.TransactOpts, deposit_limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetDepositLimit(deposit_limit *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDepositLimit(&_YearnV3Pool.TransactOpts, deposit_limit)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetDepositLimit0(opts *bind.TransactOpts, deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_deposit_limit0", deposit_limit, override)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetDepositLimit0(deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDepositLimit0(&_YearnV3Pool.TransactOpts, deposit_limit, override)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetDepositLimit0(deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDepositLimit0(&_YearnV3Pool.TransactOpts, deposit_limit, override)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetDepositLimitModule(opts *bind.TransactOpts, deposit_limit_module common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_deposit_limit_module", deposit_limit_module)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetDepositLimitModule(deposit_limit_module common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDepositLimitModule(&_YearnV3Pool.TransactOpts, deposit_limit_module)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetDepositLimitModule(deposit_limit_module common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDepositLimitModule(&_YearnV3Pool.TransactOpts, deposit_limit_module)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetDepositLimitModule0(opts *bind.TransactOpts, deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_deposit_limit_module0", deposit_limit_module, override)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetDepositLimitModule0(deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDepositLimitModule0(&_YearnV3Pool.TransactOpts, deposit_limit_module, override)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetDepositLimitModule0(deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetDepositLimitModule0(&_YearnV3Pool.TransactOpts, deposit_limit_module, override)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetMinimumTotalIdle(opts *bind.TransactOpts, minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_minimum_total_idle", minimum_total_idle)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetMinimumTotalIdle(minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetMinimumTotalIdle(&_YearnV3Pool.TransactOpts, minimum_total_idle)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetMinimumTotalIdle(minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetMinimumTotalIdle(&_YearnV3Pool.TransactOpts, minimum_total_idle)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_role", account, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetRole(&_YearnV3Pool.TransactOpts, account, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetRole(&_YearnV3Pool.TransactOpts, account, role)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetUseDefaultQueue(opts *bind.TransactOpts, use_default_queue bool) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_use_default_queue", use_default_queue)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetUseDefaultQueue(use_default_queue bool) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetUseDefaultQueue(&_YearnV3Pool.TransactOpts, use_default_queue)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetUseDefaultQueue(use_default_queue bool) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetUseDefaultQueue(&_YearnV3Pool.TransactOpts, use_default_queue)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) SetWithdrawLimitModule(opts *bind.TransactOpts, withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "set_withdraw_limit_module", withdraw_limit_module)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_YearnV3Pool *YearnV3PoolSession) SetWithdrawLimitModule(withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetWithdrawLimitModule(&_YearnV3Pool.TransactOpts, withdraw_limit_module)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) SetWithdrawLimitModule(withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.SetWithdrawLimitModule(&_YearnV3Pool.TransactOpts, withdraw_limit_module)
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_YearnV3Pool *YearnV3PoolTransactor) ShutdownVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "shutdown_vault")
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_YearnV3Pool *YearnV3PoolSession) ShutdownVault() (*types.Transaction, error) {
	return _YearnV3Pool.Contract.ShutdownVault(&_YearnV3Pool.TransactOpts)
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) ShutdownVault() (*types.Transaction, error) {
	return _YearnV3Pool.Contract.ShutdownVault(&_YearnV3Pool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Transfer(&_YearnV3Pool.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Transfer(&_YearnV3Pool.TransactOpts, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "transferFrom", sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.TransferFrom(&_YearnV3Pool.TransactOpts, sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnV3Pool *YearnV3PoolTransactorSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.TransferFrom(&_YearnV3Pool.TransactOpts, sender, receiver, amount)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) TransferRoleManager(opts *bind.TransactOpts, role_manager common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "transfer_role_manager", role_manager)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_YearnV3Pool *YearnV3PoolSession) TransferRoleManager(role_manager common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.TransferRoleManager(&_YearnV3Pool.TransactOpts, role_manager)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) TransferRoleManager(role_manager common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.TransferRoleManager(&_YearnV3Pool.TransactOpts, role_manager)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) UpdateDebt(opts *bind.TransactOpts, strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "update_debt", strategy, target_debt)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) UpdateDebt(strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.UpdateDebt(&_YearnV3Pool.TransactOpts, strategy, target_debt)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) UpdateDebt(strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.UpdateDebt(&_YearnV3Pool.TransactOpts, strategy, target_debt)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) UpdateDebt0(opts *bind.TransactOpts, strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "update_debt0", strategy, target_debt, max_loss)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) UpdateDebt0(strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.UpdateDebt0(&_YearnV3Pool.TransactOpts, strategy, target_debt, max_loss)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) UpdateDebt0(strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.UpdateDebt0(&_YearnV3Pool.TransactOpts, strategy, target_debt, max_loss)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_YearnV3Pool *YearnV3PoolTransactor) UpdateMaxDebtForStrategy(opts *bind.TransactOpts, strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "update_max_debt_for_strategy", strategy, new_max_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_YearnV3Pool *YearnV3PoolSession) UpdateMaxDebtForStrategy(strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.UpdateMaxDebtForStrategy(&_YearnV3Pool.TransactOpts, strategy, new_max_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_YearnV3Pool *YearnV3PoolTransactorSession) UpdateMaxDebtForStrategy(strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.UpdateMaxDebtForStrategy(&_YearnV3Pool.TransactOpts, strategy, new_max_debt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Withdraw(&_YearnV3Pool.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Withdraw(&_YearnV3Pool.TransactOpts, assets, receiver, owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) Withdraw0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "withdraw0", assets, receiver, owner, max_loss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Withdraw0(&_YearnV3Pool.TransactOpts, assets, receiver, owner, max_loss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Withdraw0(&_YearnV3Pool.TransactOpts, assets, receiver, owner, max_loss)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactor) Withdraw1(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.contract.Transact(opts, "withdraw1", assets, receiver, owner, max_loss, strategies)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnV3Pool *YearnV3PoolSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Withdraw1(&_YearnV3Pool.TransactOpts, assets, receiver, owner, max_loss, strategies)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnV3Pool *YearnV3PoolTransactorSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnV3Pool.Contract.Withdraw1(&_YearnV3Pool.TransactOpts, assets, receiver, owner, max_loss, strategies)
}

// YearnV3PoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YearnV3Pool contract.
type YearnV3PoolApprovalIterator struct {
	Event *YearnV3PoolApproval // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolApproval)
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
		it.Event = new(YearnV3PoolApproval)
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
func (it *YearnV3PoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolApproval represents a Approval event raised by the YearnV3Pool contract.
type YearnV3PoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YearnV3PoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolApprovalIterator{contract: _YearnV3Pool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YearnV3PoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolApproval)
				if err := _YearnV3Pool.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseApproval(log types.Log) (*YearnV3PoolApproval, error) {
	event := new(YearnV3PoolApproval)
	if err := _YearnV3Pool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolDebtPurchasedIterator is returned from FilterDebtPurchased and is used to iterate over the raw logs and unpacked data for DebtPurchased events raised by the YearnV3Pool contract.
type YearnV3PoolDebtPurchasedIterator struct {
	Event *YearnV3PoolDebtPurchased // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolDebtPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolDebtPurchased)
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
		it.Event = new(YearnV3PoolDebtPurchased)
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
func (it *YearnV3PoolDebtPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolDebtPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolDebtPurchased represents a DebtPurchased event raised by the YearnV3Pool contract.
type YearnV3PoolDebtPurchased struct {
	Strategy common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDebtPurchased is a free log retrieval operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterDebtPurchased(opts *bind.FilterOpts, strategy []common.Address) (*YearnV3PoolDebtPurchasedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "DebtPurchased", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolDebtPurchasedIterator{contract: _YearnV3Pool.contract, event: "DebtPurchased", logs: logs, sub: sub}, nil
}

// WatchDebtPurchased is a free log subscription operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchDebtPurchased(opts *bind.WatchOpts, sink chan<- *YearnV3PoolDebtPurchased, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "DebtPurchased", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolDebtPurchased)
				if err := _YearnV3Pool.contract.UnpackLog(event, "DebtPurchased", log); err != nil {
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

// ParseDebtPurchased is a log parse operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseDebtPurchased(log types.Log) (*YearnV3PoolDebtPurchased, error) {
	event := new(YearnV3PoolDebtPurchased)
	if err := _YearnV3Pool.contract.UnpackLog(event, "DebtPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolDebtUpdatedIterator is returned from FilterDebtUpdated and is used to iterate over the raw logs and unpacked data for DebtUpdated events raised by the YearnV3Pool contract.
type YearnV3PoolDebtUpdatedIterator struct {
	Event *YearnV3PoolDebtUpdated // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolDebtUpdated)
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
		it.Event = new(YearnV3PoolDebtUpdated)
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
func (it *YearnV3PoolDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolDebtUpdated represents a DebtUpdated event raised by the YearnV3Pool contract.
type YearnV3PoolDebtUpdated struct {
	Strategy    common.Address
	CurrentDebt *big.Int
	NewDebt     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDebtUpdated is a free log retrieval operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterDebtUpdated(opts *bind.FilterOpts, strategy []common.Address) (*YearnV3PoolDebtUpdatedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolDebtUpdatedIterator{contract: _YearnV3Pool.contract, event: "DebtUpdated", logs: logs, sub: sub}, nil
}

// WatchDebtUpdated is a free log subscription operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchDebtUpdated(opts *bind.WatchOpts, sink chan<- *YearnV3PoolDebtUpdated, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolDebtUpdated)
				if err := _YearnV3Pool.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
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

// ParseDebtUpdated is a log parse operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseDebtUpdated(log types.Log) (*YearnV3PoolDebtUpdated, error) {
	event := new(YearnV3PoolDebtUpdated)
	if err := _YearnV3Pool.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the YearnV3Pool contract.
type YearnV3PoolDepositIterator struct {
	Event *YearnV3PoolDeposit // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolDeposit)
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
		it.Event = new(YearnV3PoolDeposit)
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
func (it *YearnV3PoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolDeposit represents a Deposit event raised by the YearnV3Pool contract.
type YearnV3PoolDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*YearnV3PoolDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolDepositIterator{contract: _YearnV3Pool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *YearnV3PoolDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolDeposit)
				if err := _YearnV3Pool.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseDeposit(log types.Log) (*YearnV3PoolDeposit, error) {
	event := new(YearnV3PoolDeposit)
	if err := _YearnV3Pool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolRoleSetIterator is returned from FilterRoleSet and is used to iterate over the raw logs and unpacked data for RoleSet events raised by the YearnV3Pool contract.
type YearnV3PoolRoleSetIterator struct {
	Event *YearnV3PoolRoleSet // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolRoleSet)
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
		it.Event = new(YearnV3PoolRoleSet)
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
func (it *YearnV3PoolRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolRoleSet represents a RoleSet event raised by the YearnV3Pool contract.
type YearnV3PoolRoleSet struct {
	Account common.Address
	Role    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleSet is a free log retrieval operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterRoleSet(opts *bind.FilterOpts, account []common.Address, role []*big.Int) (*YearnV3PoolRoleSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "RoleSet", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolRoleSetIterator{contract: _YearnV3Pool.contract, event: "RoleSet", logs: logs, sub: sub}, nil
}

// WatchRoleSet is a free log subscription operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchRoleSet(opts *bind.WatchOpts, sink chan<- *YearnV3PoolRoleSet, account []common.Address, role []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "RoleSet", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolRoleSet)
				if err := _YearnV3Pool.contract.UnpackLog(event, "RoleSet", log); err != nil {
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

// ParseRoleSet is a log parse operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseRoleSet(log types.Log) (*YearnV3PoolRoleSet, error) {
	event := new(YearnV3PoolRoleSet)
	if err := _YearnV3Pool.contract.UnpackLog(event, "RoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the YearnV3Pool contract.
type YearnV3PoolShutdownIterator struct {
	Event *YearnV3PoolShutdown // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolShutdown)
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
		it.Event = new(YearnV3PoolShutdown)
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
func (it *YearnV3PoolShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolShutdown represents a Shutdown event raised by the YearnV3Pool contract.
type YearnV3PoolShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_YearnV3Pool *YearnV3PoolFilterer) FilterShutdown(opts *bind.FilterOpts) (*YearnV3PoolShutdownIterator, error) {

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolShutdownIterator{contract: _YearnV3Pool.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_YearnV3Pool *YearnV3PoolFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *YearnV3PoolShutdown) (event.Subscription, error) {

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolShutdown)
				if err := _YearnV3Pool.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_YearnV3Pool *YearnV3PoolFilterer) ParseShutdown(log types.Log) (*YearnV3PoolShutdown, error) {
	event := new(YearnV3PoolShutdown)
	if err := _YearnV3Pool.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolStrategyChangedIterator is returned from FilterStrategyChanged and is used to iterate over the raw logs and unpacked data for StrategyChanged events raised by the YearnV3Pool contract.
type YearnV3PoolStrategyChangedIterator struct {
	Event *YearnV3PoolStrategyChanged // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolStrategyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolStrategyChanged)
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
		it.Event = new(YearnV3PoolStrategyChanged)
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
func (it *YearnV3PoolStrategyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolStrategyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolStrategyChanged represents a StrategyChanged event raised by the YearnV3Pool contract.
type YearnV3PoolStrategyChanged struct {
	Strategy   common.Address
	ChangeType *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategyChanged is a free log retrieval operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterStrategyChanged(opts *bind.FilterOpts, strategy []common.Address, change_type []*big.Int) (*YearnV3PoolStrategyChangedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var change_typeRule []interface{}
	for _, change_typeItem := range change_type {
		change_typeRule = append(change_typeRule, change_typeItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "StrategyChanged", strategyRule, change_typeRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolStrategyChangedIterator{contract: _YearnV3Pool.contract, event: "StrategyChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyChanged is a free log subscription operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchStrategyChanged(opts *bind.WatchOpts, sink chan<- *YearnV3PoolStrategyChanged, strategy []common.Address, change_type []*big.Int) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var change_typeRule []interface{}
	for _, change_typeItem := range change_type {
		change_typeRule = append(change_typeRule, change_typeItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "StrategyChanged", strategyRule, change_typeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolStrategyChanged)
				if err := _YearnV3Pool.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
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

// ParseStrategyChanged is a log parse operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseStrategyChanged(log types.Log) (*YearnV3PoolStrategyChanged, error) {
	event := new(YearnV3PoolStrategyChanged)
	if err := _YearnV3Pool.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolStrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the YearnV3Pool contract.
type YearnV3PoolStrategyReportedIterator struct {
	Event *YearnV3PoolStrategyReported // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolStrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolStrategyReported)
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
		it.Event = new(YearnV3PoolStrategyReported)
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
func (it *YearnV3PoolStrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolStrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolStrategyReported represents a StrategyReported event raised by the YearnV3Pool contract.
type YearnV3PoolStrategyReported struct {
	Strategy     common.Address
	Gain         *big.Int
	Loss         *big.Int
	CurrentDebt  *big.Int
	ProtocolFees *big.Int
	TotalFees    *big.Int
	TotalRefunds *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStrategyReported is a free log retrieval operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*YearnV3PoolStrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolStrategyReportedIterator{contract: _YearnV3Pool.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *YearnV3PoolStrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolStrategyReported)
				if err := _YearnV3Pool.contract.UnpackLog(event, "StrategyReported", log); err != nil {
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

// ParseStrategyReported is a log parse operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseStrategyReported(log types.Log) (*YearnV3PoolStrategyReported, error) {
	event := new(YearnV3PoolStrategyReported)
	if err := _YearnV3Pool.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YearnV3Pool contract.
type YearnV3PoolTransferIterator struct {
	Event *YearnV3PoolTransfer // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolTransfer)
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
		it.Event = new(YearnV3PoolTransfer)
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
func (it *YearnV3PoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolTransfer represents a Transfer event raised by the YearnV3Pool contract.
type YearnV3PoolTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*YearnV3PoolTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolTransferIterator{contract: _YearnV3Pool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YearnV3PoolTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolTransfer)
				if err := _YearnV3Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseTransfer(log types.Log) (*YearnV3PoolTransfer, error) {
	event := new(YearnV3PoolTransfer)
	if err := _YearnV3Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateAccountantIterator is returned from FilterUpdateAccountant and is used to iterate over the raw logs and unpacked data for UpdateAccountant events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateAccountantIterator struct {
	Event *YearnV3PoolUpdateAccountant // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateAccountantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateAccountant)
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
		it.Event = new(YearnV3PoolUpdateAccountant)
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
func (it *YearnV3PoolUpdateAccountantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateAccountantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateAccountant represents a UpdateAccountant event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateAccountant struct {
	Accountant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAccountant is a free log retrieval operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateAccountant(opts *bind.FilterOpts, accountant []common.Address) (*YearnV3PoolUpdateAccountantIterator, error) {

	var accountantRule []interface{}
	for _, accountantItem := range accountant {
		accountantRule = append(accountantRule, accountantItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateAccountant", accountantRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateAccountantIterator{contract: _YearnV3Pool.contract, event: "UpdateAccountant", logs: logs, sub: sub}, nil
}

// WatchUpdateAccountant is a free log subscription operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateAccountant(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateAccountant, accountant []common.Address) (event.Subscription, error) {

	var accountantRule []interface{}
	for _, accountantItem := range accountant {
		accountantRule = append(accountantRule, accountantItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateAccountant", accountantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateAccountant)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateAccountant", log); err != nil {
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

// ParseUpdateAccountant is a log parse operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateAccountant(log types.Log) (*YearnV3PoolUpdateAccountant, error) {
	event := new(YearnV3PoolUpdateAccountant)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateAccountant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateDefaultQueueIterator is returned from FilterUpdateDefaultQueue and is used to iterate over the raw logs and unpacked data for UpdateDefaultQueue events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateDefaultQueueIterator struct {
	Event *YearnV3PoolUpdateDefaultQueue // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateDefaultQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateDefaultQueue)
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
		it.Event = new(YearnV3PoolUpdateDefaultQueue)
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
func (it *YearnV3PoolUpdateDefaultQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateDefaultQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateDefaultQueue represents a UpdateDefaultQueue event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateDefaultQueue struct {
	NewDefaultQueue []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultQueue is a free log retrieval operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateDefaultQueue(opts *bind.FilterOpts) (*YearnV3PoolUpdateDefaultQueueIterator, error) {

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateDefaultQueue")
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateDefaultQueueIterator{contract: _YearnV3Pool.contract, event: "UpdateDefaultQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultQueue is a free log subscription operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateDefaultQueue(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateDefaultQueue) (event.Subscription, error) {

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateDefaultQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateDefaultQueue)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateDefaultQueue", log); err != nil {
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

// ParseUpdateDefaultQueue is a log parse operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateDefaultQueue(log types.Log) (*YearnV3PoolUpdateDefaultQueue, error) {
	event := new(YearnV3PoolUpdateDefaultQueue)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateDefaultQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateDepositLimitIterator is returned from FilterUpdateDepositLimit and is used to iterate over the raw logs and unpacked data for UpdateDepositLimit events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateDepositLimitIterator struct {
	Event *YearnV3PoolUpdateDepositLimit // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateDepositLimit)
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
		it.Event = new(YearnV3PoolUpdateDepositLimit)
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
func (it *YearnV3PoolUpdateDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateDepositLimit represents a UpdateDepositLimit event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateDepositLimit struct {
	DepositLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimit is a free log retrieval operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateDepositLimit(opts *bind.FilterOpts) (*YearnV3PoolUpdateDepositLimitIterator, error) {

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateDepositLimitIterator{contract: _YearnV3Pool.contract, event: "UpdateDepositLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimit is a free log subscription operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateDepositLimit(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateDepositLimit) (event.Subscription, error) {

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateDepositLimit)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
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

// ParseUpdateDepositLimit is a log parse operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateDepositLimit(log types.Log) (*YearnV3PoolUpdateDepositLimit, error) {
	event := new(YearnV3PoolUpdateDepositLimit)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateDepositLimitModuleIterator is returned from FilterUpdateDepositLimitModule and is used to iterate over the raw logs and unpacked data for UpdateDepositLimitModule events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateDepositLimitModuleIterator struct {
	Event *YearnV3PoolUpdateDepositLimitModule // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateDepositLimitModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateDepositLimitModule)
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
		it.Event = new(YearnV3PoolUpdateDepositLimitModule)
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
func (it *YearnV3PoolUpdateDepositLimitModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateDepositLimitModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateDepositLimitModule represents a UpdateDepositLimitModule event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateDepositLimitModule struct {
	DepositLimitModule common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimitModule is a free log retrieval operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateDepositLimitModule(opts *bind.FilterOpts, deposit_limit_module []common.Address) (*YearnV3PoolUpdateDepositLimitModuleIterator, error) {

	var deposit_limit_moduleRule []interface{}
	for _, deposit_limit_moduleItem := range deposit_limit_module {
		deposit_limit_moduleRule = append(deposit_limit_moduleRule, deposit_limit_moduleItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateDepositLimitModule", deposit_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateDepositLimitModuleIterator{contract: _YearnV3Pool.contract, event: "UpdateDepositLimitModule", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimitModule is a free log subscription operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateDepositLimitModule(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateDepositLimitModule, deposit_limit_module []common.Address) (event.Subscription, error) {

	var deposit_limit_moduleRule []interface{}
	for _, deposit_limit_moduleItem := range deposit_limit_module {
		deposit_limit_moduleRule = append(deposit_limit_moduleRule, deposit_limit_moduleItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateDepositLimitModule", deposit_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateDepositLimitModule)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateDepositLimitModule", log); err != nil {
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

// ParseUpdateDepositLimitModule is a log parse operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateDepositLimitModule(log types.Log) (*YearnV3PoolUpdateDepositLimitModule, error) {
	event := new(YearnV3PoolUpdateDepositLimitModule)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateDepositLimitModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateMinimumTotalIdleIterator is returned from FilterUpdateMinimumTotalIdle and is used to iterate over the raw logs and unpacked data for UpdateMinimumTotalIdle events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateMinimumTotalIdleIterator struct {
	Event *YearnV3PoolUpdateMinimumTotalIdle // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateMinimumTotalIdleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateMinimumTotalIdle)
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
		it.Event = new(YearnV3PoolUpdateMinimumTotalIdle)
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
func (it *YearnV3PoolUpdateMinimumTotalIdleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateMinimumTotalIdleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateMinimumTotalIdle represents a UpdateMinimumTotalIdle event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateMinimumTotalIdle struct {
	MinimumTotalIdle *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinimumTotalIdle is a free log retrieval operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateMinimumTotalIdle(opts *bind.FilterOpts) (*YearnV3PoolUpdateMinimumTotalIdleIterator, error) {

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateMinimumTotalIdle")
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateMinimumTotalIdleIterator{contract: _YearnV3Pool.contract, event: "UpdateMinimumTotalIdle", logs: logs, sub: sub}, nil
}

// WatchUpdateMinimumTotalIdle is a free log subscription operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateMinimumTotalIdle(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateMinimumTotalIdle) (event.Subscription, error) {

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateMinimumTotalIdle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateMinimumTotalIdle)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateMinimumTotalIdle", log); err != nil {
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

// ParseUpdateMinimumTotalIdle is a log parse operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateMinimumTotalIdle(log types.Log) (*YearnV3PoolUpdateMinimumTotalIdle, error) {
	event := new(YearnV3PoolUpdateMinimumTotalIdle)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateMinimumTotalIdle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateProfitMaxUnlockTimeIterator is returned from FilterUpdateProfitMaxUnlockTime and is used to iterate over the raw logs and unpacked data for UpdateProfitMaxUnlockTime events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateProfitMaxUnlockTimeIterator struct {
	Event *YearnV3PoolUpdateProfitMaxUnlockTime // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateProfitMaxUnlockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateProfitMaxUnlockTime)
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
		it.Event = new(YearnV3PoolUpdateProfitMaxUnlockTime)
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
func (it *YearnV3PoolUpdateProfitMaxUnlockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateProfitMaxUnlockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateProfitMaxUnlockTime represents a UpdateProfitMaxUnlockTime event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateProfitMaxUnlockTime struct {
	ProfitMaxUnlockTime *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateProfitMaxUnlockTime is a free log retrieval operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateProfitMaxUnlockTime(opts *bind.FilterOpts) (*YearnV3PoolUpdateProfitMaxUnlockTimeIterator, error) {

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateProfitMaxUnlockTimeIterator{contract: _YearnV3Pool.contract, event: "UpdateProfitMaxUnlockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateProfitMaxUnlockTime is a free log subscription operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateProfitMaxUnlockTime(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateProfitMaxUnlockTime) (event.Subscription, error) {

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateProfitMaxUnlockTime)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
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

// ParseUpdateProfitMaxUnlockTime is a log parse operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateProfitMaxUnlockTime(log types.Log) (*YearnV3PoolUpdateProfitMaxUnlockTime, error) {
	event := new(YearnV3PoolUpdateProfitMaxUnlockTime)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateRoleManagerIterator is returned from FilterUpdateRoleManager and is used to iterate over the raw logs and unpacked data for UpdateRoleManager events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateRoleManagerIterator struct {
	Event *YearnV3PoolUpdateRoleManager // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateRoleManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateRoleManager)
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
		it.Event = new(YearnV3PoolUpdateRoleManager)
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
func (it *YearnV3PoolUpdateRoleManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateRoleManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateRoleManager represents a UpdateRoleManager event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateRoleManager struct {
	RoleManager common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateRoleManager is a free log retrieval operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateRoleManager(opts *bind.FilterOpts, role_manager []common.Address) (*YearnV3PoolUpdateRoleManagerIterator, error) {

	var role_managerRule []interface{}
	for _, role_managerItem := range role_manager {
		role_managerRule = append(role_managerRule, role_managerItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateRoleManager", role_managerRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateRoleManagerIterator{contract: _YearnV3Pool.contract, event: "UpdateRoleManager", logs: logs, sub: sub}, nil
}

// WatchUpdateRoleManager is a free log subscription operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateRoleManager(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateRoleManager, role_manager []common.Address) (event.Subscription, error) {

	var role_managerRule []interface{}
	for _, role_managerItem := range role_manager {
		role_managerRule = append(role_managerRule, role_managerItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateRoleManager", role_managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateRoleManager)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateRoleManager", log); err != nil {
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

// ParseUpdateRoleManager is a log parse operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateRoleManager(log types.Log) (*YearnV3PoolUpdateRoleManager, error) {
	event := new(YearnV3PoolUpdateRoleManager)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateRoleManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateUseDefaultQueueIterator is returned from FilterUpdateUseDefaultQueue and is used to iterate over the raw logs and unpacked data for UpdateUseDefaultQueue events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateUseDefaultQueueIterator struct {
	Event *YearnV3PoolUpdateUseDefaultQueue // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateUseDefaultQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateUseDefaultQueue)
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
		it.Event = new(YearnV3PoolUpdateUseDefaultQueue)
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
func (it *YearnV3PoolUpdateUseDefaultQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateUseDefaultQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateUseDefaultQueue represents a UpdateUseDefaultQueue event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateUseDefaultQueue struct {
	UseDefaultQueue bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateUseDefaultQueue is a free log retrieval operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateUseDefaultQueue(opts *bind.FilterOpts) (*YearnV3PoolUpdateUseDefaultQueueIterator, error) {

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateUseDefaultQueue")
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateUseDefaultQueueIterator{contract: _YearnV3Pool.contract, event: "UpdateUseDefaultQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateUseDefaultQueue is a free log subscription operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateUseDefaultQueue(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateUseDefaultQueue) (event.Subscription, error) {

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateUseDefaultQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateUseDefaultQueue)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateUseDefaultQueue", log); err != nil {
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

// ParseUpdateUseDefaultQueue is a log parse operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateUseDefaultQueue(log types.Log) (*YearnV3PoolUpdateUseDefaultQueue, error) {
	event := new(YearnV3PoolUpdateUseDefaultQueue)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateUseDefaultQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdateWithdrawLimitModuleIterator is returned from FilterUpdateWithdrawLimitModule and is used to iterate over the raw logs and unpacked data for UpdateWithdrawLimitModule events raised by the YearnV3Pool contract.
type YearnV3PoolUpdateWithdrawLimitModuleIterator struct {
	Event *YearnV3PoolUpdateWithdrawLimitModule // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdateWithdrawLimitModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdateWithdrawLimitModule)
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
		it.Event = new(YearnV3PoolUpdateWithdrawLimitModule)
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
func (it *YearnV3PoolUpdateWithdrawLimitModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdateWithdrawLimitModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdateWithdrawLimitModule represents a UpdateWithdrawLimitModule event raised by the YearnV3Pool contract.
type YearnV3PoolUpdateWithdrawLimitModule struct {
	WithdrawLimitModule common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawLimitModule is a free log retrieval operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdateWithdrawLimitModule(opts *bind.FilterOpts, withdraw_limit_module []common.Address) (*YearnV3PoolUpdateWithdrawLimitModuleIterator, error) {

	var withdraw_limit_moduleRule []interface{}
	for _, withdraw_limit_moduleItem := range withdraw_limit_module {
		withdraw_limit_moduleRule = append(withdraw_limit_moduleRule, withdraw_limit_moduleItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdateWithdrawLimitModule", withdraw_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdateWithdrawLimitModuleIterator{contract: _YearnV3Pool.contract, event: "UpdateWithdrawLimitModule", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawLimitModule is a free log subscription operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdateWithdrawLimitModule(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdateWithdrawLimitModule, withdraw_limit_module []common.Address) (event.Subscription, error) {

	var withdraw_limit_moduleRule []interface{}
	for _, withdraw_limit_moduleItem := range withdraw_limit_module {
		withdraw_limit_moduleRule = append(withdraw_limit_moduleRule, withdraw_limit_moduleItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdateWithdrawLimitModule", withdraw_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdateWithdrawLimitModule)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateWithdrawLimitModule", log); err != nil {
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

// ParseUpdateWithdrawLimitModule is a log parse operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdateWithdrawLimitModule(log types.Log) (*YearnV3PoolUpdateWithdrawLimitModule, error) {
	event := new(YearnV3PoolUpdateWithdrawLimitModule)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdateWithdrawLimitModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolUpdatedMaxDebtForStrategyIterator is returned from FilterUpdatedMaxDebtForStrategy and is used to iterate over the raw logs and unpacked data for UpdatedMaxDebtForStrategy events raised by the YearnV3Pool contract.
type YearnV3PoolUpdatedMaxDebtForStrategyIterator struct {
	Event *YearnV3PoolUpdatedMaxDebtForStrategy // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolUpdatedMaxDebtForStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolUpdatedMaxDebtForStrategy)
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
		it.Event = new(YearnV3PoolUpdatedMaxDebtForStrategy)
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
func (it *YearnV3PoolUpdatedMaxDebtForStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolUpdatedMaxDebtForStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolUpdatedMaxDebtForStrategy represents a UpdatedMaxDebtForStrategy event raised by the YearnV3Pool contract.
type YearnV3PoolUpdatedMaxDebtForStrategy struct {
	Sender   common.Address
	Strategy common.Address
	NewDebt  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxDebtForStrategy is a free log retrieval operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterUpdatedMaxDebtForStrategy(opts *bind.FilterOpts, sender []common.Address, strategy []common.Address) (*YearnV3PoolUpdatedMaxDebtForStrategyIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "UpdatedMaxDebtForStrategy", senderRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolUpdatedMaxDebtForStrategyIterator{contract: _YearnV3Pool.contract, event: "UpdatedMaxDebtForStrategy", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxDebtForStrategy is a free log subscription operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchUpdatedMaxDebtForStrategy(opts *bind.WatchOpts, sink chan<- *YearnV3PoolUpdatedMaxDebtForStrategy, sender []common.Address, strategy []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "UpdatedMaxDebtForStrategy", senderRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolUpdatedMaxDebtForStrategy)
				if err := _YearnV3Pool.contract.UnpackLog(event, "UpdatedMaxDebtForStrategy", log); err != nil {
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

// ParseUpdatedMaxDebtForStrategy is a log parse operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseUpdatedMaxDebtForStrategy(log types.Log) (*YearnV3PoolUpdatedMaxDebtForStrategy, error) {
	event := new(YearnV3PoolUpdatedMaxDebtForStrategy)
	if err := _YearnV3Pool.contract.UnpackLog(event, "UpdatedMaxDebtForStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3PoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the YearnV3Pool contract.
type YearnV3PoolWithdrawIterator struct {
	Event *YearnV3PoolWithdraw // Event containing the contract specifics and raw log

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
func (it *YearnV3PoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3PoolWithdraw)
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
		it.Event = new(YearnV3PoolWithdraw)
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
func (it *YearnV3PoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3PoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3PoolWithdraw represents a Withdraw event raised by the YearnV3Pool contract.
type YearnV3PoolWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Pool *YearnV3PoolFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*YearnV3PoolWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Pool.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3PoolWithdrawIterator{contract: _YearnV3Pool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Pool *YearnV3PoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *YearnV3PoolWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Pool.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3PoolWithdraw)
				if err := _YearnV3Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Pool *YearnV3PoolFilterer) ParseWithdraw(log types.Log) (*YearnV3PoolWithdraw, error) {
	event := new(YearnV3PoolWithdraw)
	if err := _YearnV3Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
