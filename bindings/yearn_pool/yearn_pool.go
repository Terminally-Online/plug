// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearn_pool

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

// YearnPoolMetaData contains all meta data concerning the YearnPool contract.
var YearnPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"change_type\",\"type\":\"uint256\"}],\"name\":\"StrategyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"gain\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loss\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"current_debt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"protocol_fees\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"total_fees\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"total_refunds\",\"type\":\"uint256\"}],\"name\":\"StrategyReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"current_debt\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"new_debt\",\"type\":\"uint256\"}],\"name\":\"DebtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"RoleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"role_manager\",\"type\":\"address\"}],\"name\":\"UpdateRoleManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"accountant\",\"type\":\"address\"}],\"name\":\"UpdateAccountant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"deposit_limit_module\",\"type\":\"address\"}],\"name\":\"UpdateDepositLimitModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"withdraw_limit_module\",\"type\":\"address\"}],\"name\":\"UpdateWithdrawLimitModule\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"new_default_queue\",\"type\":\"address[]\"}],\"name\":\"UpdateDefaultQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"use_default_queue\",\"type\":\"bool\"}],\"name\":\"UpdateUseDefaultQueue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"new_debt\",\"type\":\"uint256\"}],\"name\":\"UpdatedMaxDebtForStrategy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"deposit_limit\",\"type\":\"uint256\"}],\"name\":\"UpdateDepositLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"minimum_total_idle\",\"type\":\"uint256\"}],\"name\":\"UpdateMinimumTotalIdle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\"}],\"name\":\"UpdateProfitMaxUnlockTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DebtPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Shutdown\",\"type\":\"event\"},{\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"name\":\"asset\",\"type\":\"address\"},{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"role_manager\",\"type\":\"address\"},{\"name\":\"profit_max_unlock_time\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_accountant\",\"type\":\"address\"}],\"name\":\"set_accountant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_default_queue\",\"type\":\"address[]\"}],\"name\":\"set_default_queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"use_default_queue\",\"type\":\"bool\"}],\"name\":\"set_use_default_queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\"}],\"name\":\"set_deposit_limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"deposit_limit\",\"type\":\"uint256\"},{\"name\":\"override\",\"type\":\"bool\"}],\"name\":\"set_deposit_limit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\"}],\"name\":\"set_deposit_limit_module\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"deposit_limit_module\",\"type\":\"address\"},{\"name\":\"override\",\"type\":\"bool\"}],\"name\":\"set_deposit_limit_module\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"withdraw_limit_module\",\"type\":\"address\"}],\"name\":\"set_withdraw_limit_module\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"minimum_total_idle\",\"type\":\"uint256\"}],\"name\":\"set_minimum_total_idle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_profit_max_unlock_time\",\"type\":\"uint256\"}],\"name\":\"setProfitMaxUnlockTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"set_role\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"add_role\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"role\",\"type\":\"uint256\"}],\"name\":\"remove_role\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"role_manager\",\"type\":\"address\"}],\"name\":\"transfer_role_manager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accept_role_manager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockedShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pricePerShare\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_default_queue\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"process_report\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buy_debt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_strategy\",\"type\":\"address\"}],\"name\":\"add_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"new_strategy\",\"type\":\"address\"},{\"name\":\"add_to_queue\",\"type\":\"bool\"}],\"name\":\"add_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"revoke_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"}],\"name\":\"force_revoke_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"new_max_debt\",\"type\":\"uint256\"}],\"name\":\"update_max_debt_for_strategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"target_debt\",\"type\":\"uint256\"}],\"name\":\"update_debt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"target_debt\",\"type\":\"uint256\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"update_debt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdown_vault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalIdle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDebt\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"max_loss\",\"type\":\"uint256\"},{\"name\":\"strategies\",\"type\":\"address[]\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"apiVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"strategy\",\"type\":\"address\"},{\"name\":\"assets_needed\",\"type\":\"uint256\"}],\"name\":\"assess_share_of_unrealised_losses\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitMaxUnlockTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fullProfitUnlockDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"profitUnlockingRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastProfitUpdate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"strategies\",\"outputs\":[{\"components\":[{\"name\":\"activation\",\"type\":\"uint256\"},{\"name\":\"last_report\",\"type\":\"uint256\"},{\"name\":\"current_debt\",\"type\":\"uint256\"},{\"name\":\"max_debt\",\"type\":\"uint256\"}],\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"uint256\"}],\"name\":\"default_queue\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"use_default_queue\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"},{\"name\":\"arg1\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimum_total_idle\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit_limit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accountant\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit_limit_module\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw_limit_module\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"role_manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"future_role_manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"arg0\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YearnPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnPoolMetaData.ABI instead.
var YearnPoolABI = YearnPoolMetaData.ABI

// YearnPool is an auto generated Go binding around an Ethereum contract.
type YearnPool struct {
	YearnPoolCaller     // Read-only binding to the contract
	YearnPoolTransactor // Write-only binding to the contract
	YearnPoolFilterer   // Log filterer for contract events
}

// YearnPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnPoolSession struct {
	Contract     *YearnPool        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnPoolCallerSession struct {
	Contract *YearnPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// YearnPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnPoolTransactorSession struct {
	Contract     *YearnPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// YearnPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnPoolRaw struct {
	Contract *YearnPool // Generic contract binding to access the raw methods on
}

// YearnPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnPoolCallerRaw struct {
	Contract *YearnPoolCaller // Generic read-only contract binding to access the raw methods on
}

// YearnPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnPoolTransactorRaw struct {
	Contract *YearnPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnPool creates a new instance of YearnPool, bound to a specific deployed contract.
func NewYearnPool(address common.Address, backend bind.ContractBackend) (*YearnPool, error) {
	contract, err := bindYearnPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnPool{YearnPoolCaller: YearnPoolCaller{contract: contract}, YearnPoolTransactor: YearnPoolTransactor{contract: contract}, YearnPoolFilterer: YearnPoolFilterer{contract: contract}}, nil
}

// NewYearnPoolCaller creates a new read-only instance of YearnPool, bound to a specific deployed contract.
func NewYearnPoolCaller(address common.Address, caller bind.ContractCaller) (*YearnPoolCaller, error) {
	contract, err := bindYearnPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnPoolCaller{contract: contract}, nil
}

// NewYearnPoolTransactor creates a new write-only instance of YearnPool, bound to a specific deployed contract.
func NewYearnPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnPoolTransactor, error) {
	contract, err := bindYearnPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnPoolTransactor{contract: contract}, nil
}

// NewYearnPoolFilterer creates a new log filterer instance of YearnPool, bound to a specific deployed contract.
func NewYearnPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnPoolFilterer, error) {
	contract, err := bindYearnPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnPoolFilterer{contract: contract}, nil
}

// bindYearnPool binds a generic wrapper to an already deployed contract.
func bindYearnPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YearnPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnPool *YearnPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnPool.Contract.YearnPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnPool *YearnPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnPool.Contract.YearnPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnPool *YearnPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnPool.Contract.YearnPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnPool *YearnPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnPool *YearnPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnPool *YearnPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnPool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnPool *YearnPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnPool *YearnPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YearnPool.Contract.DOMAINSEPARATOR(&_YearnPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_YearnPool *YearnPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _YearnPool.Contract.DOMAINSEPARATOR(&_YearnPool.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YearnPool *YearnPoolCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YearnPool *YearnPoolSession) FACTORY() (common.Address, error) {
	return _YearnPool.Contract.FACTORY(&_YearnPool.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_YearnPool *YearnPoolCallerSession) FACTORY() (common.Address, error) {
	return _YearnPool.Contract.FACTORY(&_YearnPool.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_YearnPool *YearnPoolCaller) Accountant(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "accountant")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_YearnPool *YearnPoolSession) Accountant() (common.Address, error) {
	return _YearnPool.Contract.Accountant(&_YearnPool.CallOpts)
}

// Accountant is a free data retrieval call binding the contract method 0x4fb3ccc5.
//
// Solidity: function accountant() view returns(address)
func (_YearnPool *YearnPoolCallerSession) Accountant() (common.Address, error) {
	return _YearnPool.Contract.Accountant(&_YearnPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnPool *YearnPoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnPool *YearnPoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YearnPool.Contract.Allowance(&_YearnPool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address arg0, address arg1) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YearnPool.Contract.Allowance(&_YearnPool.CallOpts, arg0, arg1)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YearnPool *YearnPoolCaller) ApiVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "apiVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YearnPool *YearnPoolSession) ApiVersion() (string, error) {
	return _YearnPool.Contract.ApiVersion(&_YearnPool.CallOpts)
}

// ApiVersion is a free data retrieval call binding the contract method 0x25829410.
//
// Solidity: function apiVersion() view returns(string)
func (_YearnPool *YearnPoolCallerSession) ApiVersion() (string, error) {
	return _YearnPool.Contract.ApiVersion(&_YearnPool.CallOpts)
}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_YearnPool *YearnPoolCaller) AssessShareOfUnrealisedLosses(opts *bind.CallOpts, strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "assess_share_of_unrealised_losses", strategy, assets_needed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_YearnPool *YearnPoolSession) AssessShareOfUnrealisedLosses(strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.AssessShareOfUnrealisedLosses(&_YearnPool.CallOpts, strategy, assets_needed)
}

// AssessShareOfUnrealisedLosses is a free data retrieval call binding the contract method 0x66d3ae57.
//
// Solidity: function assess_share_of_unrealised_losses(address strategy, uint256 assets_needed) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) AssessShareOfUnrealisedLosses(strategy common.Address, assets_needed *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.AssessShareOfUnrealisedLosses(&_YearnPool.CallOpts, strategy, assets_needed)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnPool *YearnPoolCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnPool *YearnPoolSession) Asset() (common.Address, error) {
	return _YearnPool.Contract.Asset(&_YearnPool.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnPool *YearnPoolCallerSession) Asset() (common.Address, error) {
	return _YearnPool.Contract.Asset(&_YearnPool.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_YearnPool *YearnPoolCaller) BalanceOf(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "balanceOf", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_YearnPool *YearnPoolSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _YearnPool.Contract.BalanceOf(&_YearnPool.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address addr) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) BalanceOf(addr common.Address) (*big.Int, error) {
	return _YearnPool.Contract.BalanceOf(&_YearnPool.CallOpts, addr)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.ConvertToAssets(&_YearnPool.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.ConvertToAssets(&_YearnPool.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.ConvertToShares(&_YearnPool.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.ConvertToShares(&_YearnPool.CallOpts, assets)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnPool *YearnPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnPool *YearnPoolSession) Decimals() (uint8, error) {
	return _YearnPool.Contract.Decimals(&_YearnPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnPool *YearnPoolCallerSession) Decimals() (uint8, error) {
	return _YearnPool.Contract.Decimals(&_YearnPool.CallOpts)
}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_YearnPool *YearnPoolCaller) DefaultQueue(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "default_queue", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_YearnPool *YearnPoolSession) DefaultQueue(arg0 *big.Int) (common.Address, error) {
	return _YearnPool.Contract.DefaultQueue(&_YearnPool.CallOpts, arg0)
}

// DefaultQueue is a free data retrieval call binding the contract method 0x8bf03b9e.
//
// Solidity: function default_queue(uint256 arg0) view returns(address)
func (_YearnPool *YearnPoolCallerSession) DefaultQueue(arg0 *big.Int) (common.Address, error) {
	return _YearnPool.Contract.DefaultQueue(&_YearnPool.CallOpts, arg0)
}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_YearnPool *YearnPoolCaller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "deposit_limit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_YearnPool *YearnPoolSession) DepositLimit() (*big.Int, error) {
	return _YearnPool.Contract.DepositLimit(&_YearnPool.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xe46a5797.
//
// Solidity: function deposit_limit() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) DepositLimit() (*big.Int, error) {
	return _YearnPool.Contract.DepositLimit(&_YearnPool.CallOpts)
}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_YearnPool *YearnPoolCaller) DepositLimitModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "deposit_limit_module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_YearnPool *YearnPoolSession) DepositLimitModule() (common.Address, error) {
	return _YearnPool.Contract.DepositLimitModule(&_YearnPool.CallOpts)
}

// DepositLimitModule is a free data retrieval call binding the contract method 0x61c2ccf4.
//
// Solidity: function deposit_limit_module() view returns(address)
func (_YearnPool *YearnPoolCallerSession) DepositLimitModule() (common.Address, error) {
	return _YearnPool.Contract.DepositLimitModule(&_YearnPool.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YearnPool *YearnPoolCaller) FullProfitUnlockDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "fullProfitUnlockDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YearnPool *YearnPoolSession) FullProfitUnlockDate() (*big.Int, error) {
	return _YearnPool.Contract.FullProfitUnlockDate(&_YearnPool.CallOpts)
}

// FullProfitUnlockDate is a free data retrieval call binding the contract method 0x2d632692.
//
// Solidity: function fullProfitUnlockDate() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) FullProfitUnlockDate() (*big.Int, error) {
	return _YearnPool.Contract.FullProfitUnlockDate(&_YearnPool.CallOpts)
}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_YearnPool *YearnPoolCaller) FutureRoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "future_role_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_YearnPool *YearnPoolSession) FutureRoleManager() (common.Address, error) {
	return _YearnPool.Contract.FutureRoleManager(&_YearnPool.CallOpts)
}

// FutureRoleManager is a free data retrieval call binding the contract method 0x9a98f418.
//
// Solidity: function future_role_manager() view returns(address)
func (_YearnPool *YearnPoolCallerSession) FutureRoleManager() (common.Address, error) {
	return _YearnPool.Contract.FutureRoleManager(&_YearnPool.CallOpts)
}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_YearnPool *YearnPoolCaller) GetDefaultQueue(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "get_default_queue")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_YearnPool *YearnPoolSession) GetDefaultQueue() ([]common.Address, error) {
	return _YearnPool.Contract.GetDefaultQueue(&_YearnPool.CallOpts)
}

// GetDefaultQueue is a free data retrieval call binding the contract method 0xa9bbf1cc.
//
// Solidity: function get_default_queue() view returns(address[])
func (_YearnPool *YearnPoolCallerSession) GetDefaultQueue() ([]common.Address, error) {
	return _YearnPool.Contract.GetDefaultQueue(&_YearnPool.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YearnPool *YearnPoolCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YearnPool *YearnPoolSession) IsShutdown() (bool, error) {
	return _YearnPool.Contract.IsShutdown(&_YearnPool.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_YearnPool *YearnPoolCallerSession) IsShutdown() (bool, error) {
	return _YearnPool.Contract.IsShutdown(&_YearnPool.CallOpts)
}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_YearnPool *YearnPoolCaller) LastProfitUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "lastProfitUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_YearnPool *YearnPoolSession) LastProfitUpdate() (*big.Int, error) {
	return _YearnPool.Contract.LastProfitUpdate(&_YearnPool.CallOpts)
}

// LastProfitUpdate is a free data retrieval call binding the contract method 0x8afca8f0.
//
// Solidity: function lastProfitUpdate() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) LastProfitUpdate() (*big.Int, error) {
	return _YearnPool.Contract.LastProfitUpdate(&_YearnPool.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_YearnPool *YearnPoolCaller) MaxDeposit(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "maxDeposit", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_YearnPool *YearnPoolSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxDeposit(&_YearnPool.CallOpts, receiver)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address receiver) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MaxDeposit(receiver common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxDeposit(&_YearnPool.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_YearnPool *YearnPoolCaller) MaxMint(opts *bind.CallOpts, receiver common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "maxMint", receiver)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_YearnPool *YearnPoolSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxMint(&_YearnPool.CallOpts, receiver)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address receiver) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MaxMint(receiver common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxMint(&_YearnPool.CallOpts, receiver)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_YearnPool *YearnPoolCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_YearnPool *YearnPoolSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxRedeem(&_YearnPool.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxRedeem(&_YearnPool.CallOpts, owner)
}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_YearnPool *YearnPoolCaller) MaxRedeem0(opts *bind.CallOpts, owner common.Address, max_loss *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "maxRedeem0", owner, max_loss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_YearnPool *YearnPoolSession) MaxRedeem0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.MaxRedeem0(&_YearnPool.CallOpts, owner, max_loss)
}

// MaxRedeem0 is a free data retrieval call binding the contract method 0x4abe4137.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MaxRedeem0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.MaxRedeem0(&_YearnPool.CallOpts, owner, max_loss)
}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnPool *YearnPoolCaller) MaxRedeem1(opts *bind.CallOpts, owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "maxRedeem1", owner, max_loss, strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnPool *YearnPoolSession) MaxRedeem1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxRedeem1(&_YearnPool.CallOpts, owner, max_loss, strategies)
}

// MaxRedeem1 is a free data retrieval call binding the contract method 0x34b5fab6.
//
// Solidity: function maxRedeem(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MaxRedeem1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxRedeem1(&_YearnPool.CallOpts, owner, max_loss, strategies)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_YearnPool *YearnPoolCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_YearnPool *YearnPoolSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxWithdraw(&_YearnPool.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxWithdraw(&_YearnPool.CallOpts, owner)
}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_YearnPool *YearnPoolCaller) MaxWithdraw0(opts *bind.CallOpts, owner common.Address, max_loss *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "maxWithdraw0", owner, max_loss)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_YearnPool *YearnPoolSession) MaxWithdraw0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.MaxWithdraw0(&_YearnPool.CallOpts, owner, max_loss)
}

// MaxWithdraw0 is a free data retrieval call binding the contract method 0x85b68756.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MaxWithdraw0(owner common.Address, max_loss *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.MaxWithdraw0(&_YearnPool.CallOpts, owner, max_loss)
}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnPool *YearnPoolCaller) MaxWithdraw1(opts *bind.CallOpts, owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "maxWithdraw1", owner, max_loss, strategies)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnPool *YearnPoolSession) MaxWithdraw1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxWithdraw1(&_YearnPool.CallOpts, owner, max_loss, strategies)
}

// MaxWithdraw1 is a free data retrieval call binding the contract method 0x65cb6765.
//
// Solidity: function maxWithdraw(address owner, uint256 max_loss, address[] strategies) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MaxWithdraw1(owner common.Address, max_loss *big.Int, strategies []common.Address) (*big.Int, error) {
	return _YearnPool.Contract.MaxWithdraw1(&_YearnPool.CallOpts, owner, max_loss, strategies)
}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_YearnPool *YearnPoolCaller) MinimumTotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "minimum_total_idle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_YearnPool *YearnPoolSession) MinimumTotalIdle() (*big.Int, error) {
	return _YearnPool.Contract.MinimumTotalIdle(&_YearnPool.CallOpts)
}

// MinimumTotalIdle is a free data retrieval call binding the contract method 0x356d6409.
//
// Solidity: function minimum_total_idle() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) MinimumTotalIdle() (*big.Int, error) {
	return _YearnPool.Contract.MinimumTotalIdle(&_YearnPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnPool *YearnPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnPool *YearnPoolSession) Name() (string, error) {
	return _YearnPool.Contract.Name(&_YearnPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnPool *YearnPoolCallerSession) Name() (string, error) {
	return _YearnPool.Contract.Name(&_YearnPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnPool *YearnPoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnPool *YearnPoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _YearnPool.Contract.Nonces(&_YearnPool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address arg0) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _YearnPool.Contract.Nonces(&_YearnPool.CallOpts, arg0)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.PreviewDeposit(&_YearnPool.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.PreviewDeposit(&_YearnPool.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.PreviewMint(&_YearnPool.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.PreviewMint(&_YearnPool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.PreviewRedeem(&_YearnPool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.PreviewRedeem(&_YearnPool.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.PreviewWithdraw(&_YearnPool.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _YearnPool.Contract.PreviewWithdraw(&_YearnPool.CallOpts, assets)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnPool *YearnPoolCaller) PricePerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "pricePerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnPool *YearnPoolSession) PricePerShare() (*big.Int, error) {
	return _YearnPool.Contract.PricePerShare(&_YearnPool.CallOpts)
}

// PricePerShare is a free data retrieval call binding the contract method 0x99530b06.
//
// Solidity: function pricePerShare() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) PricePerShare() (*big.Int, error) {
	return _YearnPool.Contract.PricePerShare(&_YearnPool.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YearnPool *YearnPoolCaller) ProfitMaxUnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "profitMaxUnlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YearnPool *YearnPoolSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _YearnPool.Contract.ProfitMaxUnlockTime(&_YearnPool.CallOpts)
}

// ProfitMaxUnlockTime is a free data retrieval call binding the contract method 0x0952864e.
//
// Solidity: function profitMaxUnlockTime() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) ProfitMaxUnlockTime() (*big.Int, error) {
	return _YearnPool.Contract.ProfitMaxUnlockTime(&_YearnPool.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YearnPool *YearnPoolCaller) ProfitUnlockingRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "profitUnlockingRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YearnPool *YearnPoolSession) ProfitUnlockingRate() (*big.Int, error) {
	return _YearnPool.Contract.ProfitUnlockingRate(&_YearnPool.CallOpts)
}

// ProfitUnlockingRate is a free data retrieval call binding the contract method 0x5141eebb.
//
// Solidity: function profitUnlockingRate() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) ProfitUnlockingRate() (*big.Int, error) {
	return _YearnPool.Contract.ProfitUnlockingRate(&_YearnPool.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_YearnPool *YearnPoolCaller) RoleManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "role_manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_YearnPool *YearnPoolSession) RoleManager() (common.Address, error) {
	return _YearnPool.Contract.RoleManager(&_YearnPool.CallOpts)
}

// RoleManager is a free data retrieval call binding the contract method 0x79b98917.
//
// Solidity: function role_manager() view returns(address)
func (_YearnPool *YearnPoolCallerSession) RoleManager() (common.Address, error) {
	return _YearnPool.Contract.RoleManager(&_YearnPool.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_YearnPool *YearnPoolCaller) Roles(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "roles", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_YearnPool *YearnPoolSession) Roles(arg0 common.Address) (*big.Int, error) {
	return _YearnPool.Contract.Roles(&_YearnPool.CallOpts, arg0)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(address arg0) view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) Roles(arg0 common.Address) (*big.Int, error) {
	return _YearnPool.Contract.Roles(&_YearnPool.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_YearnPool *YearnPoolCaller) Strategies(opts *bind.CallOpts, arg0 common.Address) (Struct0, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "strategies", arg0)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_YearnPool *YearnPoolSession) Strategies(arg0 common.Address) (Struct0, error) {
	return _YearnPool.Contract.Strategies(&_YearnPool.CallOpts, arg0)
}

// Strategies is a free data retrieval call binding the contract method 0x39ebf823.
//
// Solidity: function strategies(address arg0) view returns((uint256,uint256,uint256,uint256))
func (_YearnPool *YearnPoolCallerSession) Strategies(arg0 common.Address) (Struct0, error) {
	return _YearnPool.Contract.Strategies(&_YearnPool.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnPool *YearnPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnPool *YearnPoolSession) Symbol() (string, error) {
	return _YearnPool.Contract.Symbol(&_YearnPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnPool *YearnPoolCallerSession) Symbol() (string, error) {
	return _YearnPool.Contract.Symbol(&_YearnPool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnPool *YearnPoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnPool *YearnPoolSession) TotalAssets() (*big.Int, error) {
	return _YearnPool.Contract.TotalAssets(&_YearnPool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) TotalAssets() (*big.Int, error) {
	return _YearnPool.Contract.TotalAssets(&_YearnPool.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnPool *YearnPoolCaller) TotalDebt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "totalDebt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnPool *YearnPoolSession) TotalDebt() (*big.Int, error) {
	return _YearnPool.Contract.TotalDebt(&_YearnPool.CallOpts)
}

// TotalDebt is a free data retrieval call binding the contract method 0xfc7b9c18.
//
// Solidity: function totalDebt() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) TotalDebt() (*big.Int, error) {
	return _YearnPool.Contract.TotalDebt(&_YearnPool.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YearnPool *YearnPoolCaller) TotalIdle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "totalIdle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YearnPool *YearnPoolSession) TotalIdle() (*big.Int, error) {
	return _YearnPool.Contract.TotalIdle(&_YearnPool.CallOpts)
}

// TotalIdle is a free data retrieval call binding the contract method 0x9aa7df94.
//
// Solidity: function totalIdle() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) TotalIdle() (*big.Int, error) {
	return _YearnPool.Contract.TotalIdle(&_YearnPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnPool *YearnPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnPool *YearnPoolSession) TotalSupply() (*big.Int, error) {
	return _YearnPool.Contract.TotalSupply(&_YearnPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _YearnPool.Contract.TotalSupply(&_YearnPool.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_YearnPool *YearnPoolCaller) UnlockedShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "unlockedShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_YearnPool *YearnPoolSession) UnlockedShares() (*big.Int, error) {
	return _YearnPool.Contract.UnlockedShares(&_YearnPool.CallOpts)
}

// UnlockedShares is a free data retrieval call binding the contract method 0xd9a0e97a.
//
// Solidity: function unlockedShares() view returns(uint256)
func (_YearnPool *YearnPoolCallerSession) UnlockedShares() (*big.Int, error) {
	return _YearnPool.Contract.UnlockedShares(&_YearnPool.CallOpts)
}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_YearnPool *YearnPoolCaller) UseDefaultQueue(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "use_default_queue")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_YearnPool *YearnPoolSession) UseDefaultQueue() (bool, error) {
	return _YearnPool.Contract.UseDefaultQueue(&_YearnPool.CallOpts)
}

// UseDefaultQueue is a free data retrieval call binding the contract method 0x1e56558d.
//
// Solidity: function use_default_queue() view returns(bool)
func (_YearnPool *YearnPoolCallerSession) UseDefaultQueue() (bool, error) {
	return _YearnPool.Contract.UseDefaultQueue(&_YearnPool.CallOpts)
}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_YearnPool *YearnPoolCaller) WithdrawLimitModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnPool.contract.Call(opts, &out, "withdraw_limit_module")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_YearnPool *YearnPoolSession) WithdrawLimitModule() (common.Address, error) {
	return _YearnPool.Contract.WithdrawLimitModule(&_YearnPool.CallOpts)
}

// WithdrawLimitModule is a free data retrieval call binding the contract method 0xf5ba68f3.
//
// Solidity: function withdraw_limit_module() view returns(address)
func (_YearnPool *YearnPoolCallerSession) WithdrawLimitModule() (common.Address, error) {
	return _YearnPool.Contract.WithdrawLimitModule(&_YearnPool.CallOpts)
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_YearnPool *YearnPoolTransactor) AcceptRoleManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "accept_role_manager")
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_YearnPool *YearnPoolSession) AcceptRoleManager() (*types.Transaction, error) {
	return _YearnPool.Contract.AcceptRoleManager(&_YearnPool.TransactOpts)
}

// AcceptRoleManager is a paid mutator transaction binding the contract method 0xf776bf1f.
//
// Solidity: function accept_role_manager() returns()
func (_YearnPool *YearnPoolTransactorSession) AcceptRoleManager() (*types.Transaction, error) {
	return _YearnPool.Contract.AcceptRoleManager(&_YearnPool.TransactOpts)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolTransactor) AddRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "add_role", account, role)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolSession) AddRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.AddRole(&_YearnPool.TransactOpts, account, role)
}

// AddRole is a paid mutator transaction binding the contract method 0xa97cefa2.
//
// Solidity: function add_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolTransactorSession) AddRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.AddRole(&_YearnPool.TransactOpts, account, role)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_YearnPool *YearnPoolTransactor) AddStrategy(opts *bind.TransactOpts, new_strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "add_strategy", new_strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_YearnPool *YearnPoolSession) AddStrategy(new_strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.AddStrategy(&_YearnPool.TransactOpts, new_strategy)
}

// AddStrategy is a paid mutator transaction binding the contract method 0xde7aeb41.
//
// Solidity: function add_strategy(address new_strategy) returns()
func (_YearnPool *YearnPoolTransactorSession) AddStrategy(new_strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.AddStrategy(&_YearnPool.TransactOpts, new_strategy)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_YearnPool *YearnPoolTransactor) AddStrategy0(opts *bind.TransactOpts, new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "add_strategy0", new_strategy, add_to_queue)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_YearnPool *YearnPoolSession) AddStrategy0(new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _YearnPool.Contract.AddStrategy0(&_YearnPool.TransactOpts, new_strategy, add_to_queue)
}

// AddStrategy0 is a paid mutator transaction binding the contract method 0xc2e73cca.
//
// Solidity: function add_strategy(address new_strategy, bool add_to_queue) returns()
func (_YearnPool *YearnPoolTransactorSession) AddStrategy0(new_strategy common.Address, add_to_queue bool) (*types.Transaction, error) {
	return _YearnPool.Contract.AddStrategy0(&_YearnPool.TransactOpts, new_strategy, add_to_queue)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Approve(&_YearnPool.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Approve(&_YearnPool.TransactOpts, spender, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_YearnPool *YearnPoolTransactor) BuyDebt(opts *bind.TransactOpts, strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "buy_debt", strategy, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_YearnPool *YearnPoolSession) BuyDebt(strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.BuyDebt(&_YearnPool.TransactOpts, strategy, amount)
}

// BuyDebt is a paid mutator transaction binding the contract method 0xe5e91818.
//
// Solidity: function buy_debt(address strategy, uint256 amount) returns()
func (_YearnPool *YearnPoolTransactorSession) BuyDebt(strategy common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.BuyDebt(&_YearnPool.TransactOpts, strategy, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_YearnPool *YearnPoolTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_YearnPool *YearnPoolSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Deposit(&_YearnPool.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Deposit(&_YearnPool.TransactOpts, assets, receiver)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_YearnPool *YearnPoolTransactor) ForceRevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "force_revoke_strategy", strategy)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_YearnPool *YearnPoolSession) ForceRevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.ForceRevokeStrategy(&_YearnPool.TransactOpts, strategy)
}

// ForceRevokeStrategy is a paid mutator transaction binding the contract method 0xfd129e63.
//
// Solidity: function force_revoke_strategy(address strategy) returns()
func (_YearnPool *YearnPoolTransactorSession) ForceRevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.ForceRevokeStrategy(&_YearnPool.TransactOpts, strategy)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_YearnPool *YearnPoolTransactor) Initialize(opts *bind.TransactOpts, asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "initialize", asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_YearnPool *YearnPoolSession) Initialize(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Initialize(&_YearnPool.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Initialize is a paid mutator transaction binding the contract method 0x75b30be6.
//
// Solidity: function initialize(address asset, string name, string symbol, address role_manager, uint256 profit_max_unlock_time) returns()
func (_YearnPool *YearnPoolTransactorSession) Initialize(asset common.Address, name string, symbol string, role_manager common.Address, profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Initialize(&_YearnPool.TransactOpts, asset, name, symbol, role_manager, profit_max_unlock_time)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_YearnPool *YearnPoolTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_YearnPool *YearnPoolSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Mint(&_YearnPool.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Mint(&_YearnPool.TransactOpts, shares, receiver)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_YearnPool *YearnPoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "permit", owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_YearnPool *YearnPoolSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnPool.Contract.Permit(&_YearnPool.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns(bool)
func (_YearnPool *YearnPoolTransactorSession) Permit(owner common.Address, spender common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnPool.Contract.Permit(&_YearnPool.TransactOpts, owner, spender, amount, deadline, v, r, s)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_YearnPool *YearnPoolTransactor) ProcessReport(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "process_report", strategy)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_YearnPool *YearnPoolSession) ProcessReport(strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.ProcessReport(&_YearnPool.TransactOpts, strategy)
}

// ProcessReport is a paid mutator transaction binding the contract method 0x6ec2b8d4.
//
// Solidity: function process_report(address strategy) returns(uint256, uint256)
func (_YearnPool *YearnPoolTransactorSession) ProcessReport(strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.ProcessReport(&_YearnPool.TransactOpts, strategy)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YearnPool *YearnPoolTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YearnPool *YearnPoolSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Redeem(&_YearnPool.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Redeem(&_YearnPool.TransactOpts, shares, receiver, owner)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolTransactor) Redeem0(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "redeem0", shares, receiver, owner, max_loss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Redeem0(&_YearnPool.TransactOpts, shares, receiver, owner, max_loss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) Redeem0(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Redeem0(&_YearnPool.TransactOpts, shares, receiver, owner, max_loss)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnPool *YearnPoolTransactor) Redeem1(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "redeem1", shares, receiver, owner, max_loss, strategies)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnPool *YearnPoolSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Redeem1(&_YearnPool.TransactOpts, shares, receiver, owner, max_loss, strategies)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x06580f2d.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) Redeem1(shares *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Redeem1(&_YearnPool.TransactOpts, shares, receiver, owner, max_loss, strategies)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolTransactor) RemoveRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "remove_role", account, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolSession) RemoveRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.RemoveRole(&_YearnPool.TransactOpts, account, role)
}

// RemoveRole is a paid mutator transaction binding the contract method 0xe2bf56dd.
//
// Solidity: function remove_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolTransactorSession) RemoveRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.RemoveRole(&_YearnPool.TransactOpts, account, role)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_YearnPool *YearnPoolTransactor) RevokeStrategy(opts *bind.TransactOpts, strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "revoke_strategy", strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_YearnPool *YearnPoolSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.RevokeStrategy(&_YearnPool.TransactOpts, strategy)
}

// RevokeStrategy is a paid mutator transaction binding the contract method 0x577db316.
//
// Solidity: function revoke_strategy(address strategy) returns()
func (_YearnPool *YearnPoolTransactorSession) RevokeStrategy(strategy common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.RevokeStrategy(&_YearnPool.TransactOpts, strategy)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_YearnPool *YearnPoolTransactor) SetProfitMaxUnlockTime(opts *bind.TransactOpts, new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "setProfitMaxUnlockTime", new_profit_max_unlock_time)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_YearnPool *YearnPoolSession) SetProfitMaxUnlockTime(new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.SetProfitMaxUnlockTime(&_YearnPool.TransactOpts, new_profit_max_unlock_time)
}

// SetProfitMaxUnlockTime is a paid mutator transaction binding the contract method 0xdf69b22a.
//
// Solidity: function setProfitMaxUnlockTime(uint256 new_profit_max_unlock_time) returns()
func (_YearnPool *YearnPoolTransactorSession) SetProfitMaxUnlockTime(new_profit_max_unlock_time *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.SetProfitMaxUnlockTime(&_YearnPool.TransactOpts, new_profit_max_unlock_time)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_YearnPool *YearnPoolTransactor) SetAccountant(opts *bind.TransactOpts, new_accountant common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_accountant", new_accountant)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_YearnPool *YearnPoolSession) SetAccountant(new_accountant common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.SetAccountant(&_YearnPool.TransactOpts, new_accountant)
}

// SetAccountant is a paid mutator transaction binding the contract method 0x71da8a8d.
//
// Solidity: function set_accountant(address new_accountant) returns()
func (_YearnPool *YearnPoolTransactorSession) SetAccountant(new_accountant common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.SetAccountant(&_YearnPool.TransactOpts, new_accountant)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_YearnPool *YearnPoolTransactor) SetDefaultQueue(opts *bind.TransactOpts, new_default_queue []common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_default_queue", new_default_queue)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_YearnPool *YearnPoolSession) SetDefaultQueue(new_default_queue []common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDefaultQueue(&_YearnPool.TransactOpts, new_default_queue)
}

// SetDefaultQueue is a paid mutator transaction binding the contract method 0x2d9caa4e.
//
// Solidity: function set_default_queue(address[] new_default_queue) returns()
func (_YearnPool *YearnPoolTransactorSession) SetDefaultQueue(new_default_queue []common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDefaultQueue(&_YearnPool.TransactOpts, new_default_queue)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_YearnPool *YearnPoolTransactor) SetDepositLimit(opts *bind.TransactOpts, deposit_limit *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_deposit_limit", deposit_limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_YearnPool *YearnPoolSession) SetDepositLimit(deposit_limit *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDepositLimit(&_YearnPool.TransactOpts, deposit_limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0x6fe01d1e.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit) returns()
func (_YearnPool *YearnPoolTransactorSession) SetDepositLimit(deposit_limit *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDepositLimit(&_YearnPool.TransactOpts, deposit_limit)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_YearnPool *YearnPoolTransactor) SetDepositLimit0(opts *bind.TransactOpts, deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_deposit_limit0", deposit_limit, override)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_YearnPool *YearnPoolSession) SetDepositLimit0(deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDepositLimit0(&_YearnPool.TransactOpts, deposit_limit, override)
}

// SetDepositLimit0 is a paid mutator transaction binding the contract method 0x81685796.
//
// Solidity: function set_deposit_limit(uint256 deposit_limit, bool override) returns()
func (_YearnPool *YearnPoolTransactorSession) SetDepositLimit0(deposit_limit *big.Int, override bool) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDepositLimit0(&_YearnPool.TransactOpts, deposit_limit, override)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_YearnPool *YearnPoolTransactor) SetDepositLimitModule(opts *bind.TransactOpts, deposit_limit_module common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_deposit_limit_module", deposit_limit_module)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_YearnPool *YearnPoolSession) SetDepositLimitModule(deposit_limit_module common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDepositLimitModule(&_YearnPool.TransactOpts, deposit_limit_module)
}

// SetDepositLimitModule is a paid mutator transaction binding the contract method 0xbb435466.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module) returns()
func (_YearnPool *YearnPoolTransactorSession) SetDepositLimitModule(deposit_limit_module common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDepositLimitModule(&_YearnPool.TransactOpts, deposit_limit_module)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_YearnPool *YearnPoolTransactor) SetDepositLimitModule0(opts *bind.TransactOpts, deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_deposit_limit_module0", deposit_limit_module, override)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_YearnPool *YearnPoolSession) SetDepositLimitModule0(deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDepositLimitModule0(&_YearnPool.TransactOpts, deposit_limit_module, override)
}

// SetDepositLimitModule0 is a paid mutator transaction binding the contract method 0x9823dd78.
//
// Solidity: function set_deposit_limit_module(address deposit_limit_module, bool override) returns()
func (_YearnPool *YearnPoolTransactorSession) SetDepositLimitModule0(deposit_limit_module common.Address, override bool) (*types.Transaction, error) {
	return _YearnPool.Contract.SetDepositLimitModule0(&_YearnPool.TransactOpts, deposit_limit_module, override)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_YearnPool *YearnPoolTransactor) SetMinimumTotalIdle(opts *bind.TransactOpts, minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_minimum_total_idle", minimum_total_idle)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_YearnPool *YearnPoolSession) SetMinimumTotalIdle(minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.SetMinimumTotalIdle(&_YearnPool.TransactOpts, minimum_total_idle)
}

// SetMinimumTotalIdle is a paid mutator transaction binding the contract method 0xbdd81c01.
//
// Solidity: function set_minimum_total_idle(uint256 minimum_total_idle) returns()
func (_YearnPool *YearnPoolTransactorSession) SetMinimumTotalIdle(minimum_total_idle *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.SetMinimumTotalIdle(&_YearnPool.TransactOpts, minimum_total_idle)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolTransactor) SetRole(opts *bind.TransactOpts, account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_role", account, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolSession) SetRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.SetRole(&_YearnPool.TransactOpts, account, role)
}

// SetRole is a paid mutator transaction binding the contract method 0x2cf7fd85.
//
// Solidity: function set_role(address account, uint256 role) returns()
func (_YearnPool *YearnPoolTransactorSession) SetRole(account common.Address, role *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.SetRole(&_YearnPool.TransactOpts, account, role)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_YearnPool *YearnPoolTransactor) SetUseDefaultQueue(opts *bind.TransactOpts, use_default_queue bool) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_use_default_queue", use_default_queue)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_YearnPool *YearnPoolSession) SetUseDefaultQueue(use_default_queue bool) (*types.Transaction, error) {
	return _YearnPool.Contract.SetUseDefaultQueue(&_YearnPool.TransactOpts, use_default_queue)
}

// SetUseDefaultQueue is a paid mutator transaction binding the contract method 0x29c8a33b.
//
// Solidity: function set_use_default_queue(bool use_default_queue) returns()
func (_YearnPool *YearnPoolTransactorSession) SetUseDefaultQueue(use_default_queue bool) (*types.Transaction, error) {
	return _YearnPool.Contract.SetUseDefaultQueue(&_YearnPool.TransactOpts, use_default_queue)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_YearnPool *YearnPoolTransactor) SetWithdrawLimitModule(opts *bind.TransactOpts, withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "set_withdraw_limit_module", withdraw_limit_module)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_YearnPool *YearnPoolSession) SetWithdrawLimitModule(withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.SetWithdrawLimitModule(&_YearnPool.TransactOpts, withdraw_limit_module)
}

// SetWithdrawLimitModule is a paid mutator transaction binding the contract method 0x7b675894.
//
// Solidity: function set_withdraw_limit_module(address withdraw_limit_module) returns()
func (_YearnPool *YearnPoolTransactorSession) SetWithdrawLimitModule(withdraw_limit_module common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.SetWithdrawLimitModule(&_YearnPool.TransactOpts, withdraw_limit_module)
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_YearnPool *YearnPoolTransactor) ShutdownVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "shutdown_vault")
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_YearnPool *YearnPoolSession) ShutdownVault() (*types.Transaction, error) {
	return _YearnPool.Contract.ShutdownVault(&_YearnPool.TransactOpts)
}

// ShutdownVault is a paid mutator transaction binding the contract method 0x36a55450.
//
// Solidity: function shutdown_vault() returns()
func (_YearnPool *YearnPoolTransactorSession) ShutdownVault() (*types.Transaction, error) {
	return _YearnPool.Contract.ShutdownVault(&_YearnPool.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Transfer(&_YearnPool.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolTransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Transfer(&_YearnPool.TransactOpts, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "transferFrom", sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.TransferFrom(&_YearnPool.TransactOpts, sender, receiver, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 amount) returns(bool)
func (_YearnPool *YearnPoolTransactorSession) TransferFrom(sender common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.TransferFrom(&_YearnPool.TransactOpts, sender, receiver, amount)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_YearnPool *YearnPoolTransactor) TransferRoleManager(opts *bind.TransactOpts, role_manager common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "transfer_role_manager", role_manager)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_YearnPool *YearnPoolSession) TransferRoleManager(role_manager common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.TransferRoleManager(&_YearnPool.TransactOpts, role_manager)
}

// TransferRoleManager is a paid mutator transaction binding the contract method 0xef54cefd.
//
// Solidity: function transfer_role_manager(address role_manager) returns()
func (_YearnPool *YearnPoolTransactorSession) TransferRoleManager(role_manager common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.TransferRoleManager(&_YearnPool.TransactOpts, role_manager)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_YearnPool *YearnPoolTransactor) UpdateDebt(opts *bind.TransactOpts, strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "update_debt", strategy, target_debt)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_YearnPool *YearnPoolSession) UpdateDebt(strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.UpdateDebt(&_YearnPool.TransactOpts, strategy, target_debt)
}

// UpdateDebt is a paid mutator transaction binding the contract method 0x0aeebf55.
//
// Solidity: function update_debt(address strategy, uint256 target_debt) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) UpdateDebt(strategy common.Address, target_debt *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.UpdateDebt(&_YearnPool.TransactOpts, strategy, target_debt)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolTransactor) UpdateDebt0(opts *bind.TransactOpts, strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "update_debt0", strategy, target_debt, max_loss)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolSession) UpdateDebt0(strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.UpdateDebt0(&_YearnPool.TransactOpts, strategy, target_debt, max_loss)
}

// UpdateDebt0 is a paid mutator transaction binding the contract method 0xba54971f.
//
// Solidity: function update_debt(address strategy, uint256 target_debt, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) UpdateDebt0(strategy common.Address, target_debt *big.Int, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.UpdateDebt0(&_YearnPool.TransactOpts, strategy, target_debt, max_loss)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_YearnPool *YearnPoolTransactor) UpdateMaxDebtForStrategy(opts *bind.TransactOpts, strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "update_max_debt_for_strategy", strategy, new_max_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_YearnPool *YearnPoolSession) UpdateMaxDebtForStrategy(strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.UpdateMaxDebtForStrategy(&_YearnPool.TransactOpts, strategy, new_max_debt)
}

// UpdateMaxDebtForStrategy is a paid mutator transaction binding the contract method 0xb9ddcd68.
//
// Solidity: function update_max_debt_for_strategy(address strategy, uint256 new_max_debt) returns()
func (_YearnPool *YearnPoolTransactorSession) UpdateMaxDebtForStrategy(strategy common.Address, new_max_debt *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.UpdateMaxDebtForStrategy(&_YearnPool.TransactOpts, strategy, new_max_debt)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_YearnPool *YearnPoolTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_YearnPool *YearnPoolSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Withdraw(&_YearnPool.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Withdraw(&_YearnPool.TransactOpts, assets, receiver, owner)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolTransactor) Withdraw0(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "withdraw0", assets, receiver, owner, max_loss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Withdraw0(&_YearnPool.TransactOpts, assets, receiver, owner, max_loss)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) Withdraw0(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int) (*types.Transaction, error) {
	return _YearnPool.Contract.Withdraw0(&_YearnPool.TransactOpts, assets, receiver, owner, max_loss)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnPool *YearnPoolTransactor) Withdraw1(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnPool.contract.Transact(opts, "withdraw1", assets, receiver, owner, max_loss, strategies)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnPool *YearnPoolSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Withdraw1(&_YearnPool.TransactOpts, assets, receiver, owner, max_loss, strategies)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xd81a09f6.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 max_loss, address[] strategies) returns(uint256)
func (_YearnPool *YearnPoolTransactorSession) Withdraw1(assets *big.Int, receiver common.Address, owner common.Address, max_loss *big.Int, strategies []common.Address) (*types.Transaction, error) {
	return _YearnPool.Contract.Withdraw1(&_YearnPool.TransactOpts, assets, receiver, owner, max_loss, strategies)
}

// YearnPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YearnPool contract.
type YearnPoolApprovalIterator struct {
	Event *YearnPoolApproval // Event containing the contract specifics and raw log

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
func (it *YearnPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolApproval)
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
		it.Event = new(YearnPoolApproval)
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
func (it *YearnPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolApproval represents a Approval event raised by the YearnPool contract.
type YearnPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnPool *YearnPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YearnPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolApprovalIterator{contract: _YearnPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnPool *YearnPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YearnPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolApproval)
				if err := _YearnPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseApproval(log types.Log) (*YearnPoolApproval, error) {
	event := new(YearnPoolApproval)
	if err := _YearnPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolDebtPurchasedIterator is returned from FilterDebtPurchased and is used to iterate over the raw logs and unpacked data for DebtPurchased events raised by the YearnPool contract.
type YearnPoolDebtPurchasedIterator struct {
	Event *YearnPoolDebtPurchased // Event containing the contract specifics and raw log

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
func (it *YearnPoolDebtPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolDebtPurchased)
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
		it.Event = new(YearnPoolDebtPurchased)
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
func (it *YearnPoolDebtPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolDebtPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolDebtPurchased represents a DebtPurchased event raised by the YearnPool contract.
type YearnPoolDebtPurchased struct {
	Strategy common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDebtPurchased is a free log retrieval operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_YearnPool *YearnPoolFilterer) FilterDebtPurchased(opts *bind.FilterOpts, strategy []common.Address) (*YearnPoolDebtPurchasedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "DebtPurchased", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolDebtPurchasedIterator{contract: _YearnPool.contract, event: "DebtPurchased", logs: logs, sub: sub}, nil
}

// WatchDebtPurchased is a free log subscription operation binding the contract event 0xe94e7f88819f66c19b097748cb754149f63b1a176ed425dee1f1ee933e6d09b0.
//
// Solidity: event DebtPurchased(address indexed strategy, uint256 amount)
func (_YearnPool *YearnPoolFilterer) WatchDebtPurchased(opts *bind.WatchOpts, sink chan<- *YearnPoolDebtPurchased, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "DebtPurchased", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolDebtPurchased)
				if err := _YearnPool.contract.UnpackLog(event, "DebtPurchased", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseDebtPurchased(log types.Log) (*YearnPoolDebtPurchased, error) {
	event := new(YearnPoolDebtPurchased)
	if err := _YearnPool.contract.UnpackLog(event, "DebtPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolDebtUpdatedIterator is returned from FilterDebtUpdated and is used to iterate over the raw logs and unpacked data for DebtUpdated events raised by the YearnPool contract.
type YearnPoolDebtUpdatedIterator struct {
	Event *YearnPoolDebtUpdated // Event containing the contract specifics and raw log

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
func (it *YearnPoolDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolDebtUpdated)
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
		it.Event = new(YearnPoolDebtUpdated)
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
func (it *YearnPoolDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolDebtUpdated represents a DebtUpdated event raised by the YearnPool contract.
type YearnPoolDebtUpdated struct {
	Strategy    common.Address
	CurrentDebt *big.Int
	NewDebt     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDebtUpdated is a free log retrieval operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_YearnPool *YearnPoolFilterer) FilterDebtUpdated(opts *bind.FilterOpts, strategy []common.Address) (*YearnPoolDebtUpdatedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolDebtUpdatedIterator{contract: _YearnPool.contract, event: "DebtUpdated", logs: logs, sub: sub}, nil
}

// WatchDebtUpdated is a free log subscription operation binding the contract event 0x5e2b8821ad6e0e26207e0cb4d242d07eeb1cbb1cfd853e645bdcd27cc5484f95.
//
// Solidity: event DebtUpdated(address indexed strategy, uint256 current_debt, uint256 new_debt)
func (_YearnPool *YearnPoolFilterer) WatchDebtUpdated(opts *bind.WatchOpts, sink chan<- *YearnPoolDebtUpdated, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "DebtUpdated", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolDebtUpdated)
				if err := _YearnPool.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseDebtUpdated(log types.Log) (*YearnPoolDebtUpdated, error) {
	event := new(YearnPoolDebtUpdated)
	if err := _YearnPool.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the YearnPool contract.
type YearnPoolDepositIterator struct {
	Event *YearnPoolDeposit // Event containing the contract specifics and raw log

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
func (it *YearnPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolDeposit)
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
		it.Event = new(YearnPoolDeposit)
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
func (it *YearnPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolDeposit represents a Deposit event raised by the YearnPool contract.
type YearnPoolDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_YearnPool *YearnPoolFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*YearnPoolDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolDepositIterator{contract: _YearnPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_YearnPool *YearnPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *YearnPoolDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolDeposit)
				if err := _YearnPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseDeposit(log types.Log) (*YearnPoolDeposit, error) {
	event := new(YearnPoolDeposit)
	if err := _YearnPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolRoleSetIterator is returned from FilterRoleSet and is used to iterate over the raw logs and unpacked data for RoleSet events raised by the YearnPool contract.
type YearnPoolRoleSetIterator struct {
	Event *YearnPoolRoleSet // Event containing the contract specifics and raw log

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
func (it *YearnPoolRoleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolRoleSet)
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
		it.Event = new(YearnPoolRoleSet)
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
func (it *YearnPoolRoleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolRoleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolRoleSet represents a RoleSet event raised by the YearnPool contract.
type YearnPoolRoleSet struct {
	Account common.Address
	Role    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleSet is a free log retrieval operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_YearnPool *YearnPoolFilterer) FilterRoleSet(opts *bind.FilterOpts, account []common.Address, role []*big.Int) (*YearnPoolRoleSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "RoleSet", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolRoleSetIterator{contract: _YearnPool.contract, event: "RoleSet", logs: logs, sub: sub}, nil
}

// WatchRoleSet is a free log subscription operation binding the contract event 0x78557646b1d8efa2cd49740d66df5aca39eb610ca8ca0e1ccac08979b6b2c46e.
//
// Solidity: event RoleSet(address indexed account, uint256 indexed role)
func (_YearnPool *YearnPoolFilterer) WatchRoleSet(opts *bind.WatchOpts, sink chan<- *YearnPoolRoleSet, account []common.Address, role []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "RoleSet", accountRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolRoleSet)
				if err := _YearnPool.contract.UnpackLog(event, "RoleSet", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseRoleSet(log types.Log) (*YearnPoolRoleSet, error) {
	event := new(YearnPoolRoleSet)
	if err := _YearnPool.contract.UnpackLog(event, "RoleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the YearnPool contract.
type YearnPoolShutdownIterator struct {
	Event *YearnPoolShutdown // Event containing the contract specifics and raw log

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
func (it *YearnPoolShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolShutdown)
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
		it.Event = new(YearnPoolShutdown)
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
func (it *YearnPoolShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolShutdown represents a Shutdown event raised by the YearnPool contract.
type YearnPoolShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_YearnPool *YearnPoolFilterer) FilterShutdown(opts *bind.FilterOpts) (*YearnPoolShutdownIterator, error) {

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &YearnPoolShutdownIterator{contract: _YearnPool.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_YearnPool *YearnPoolFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *YearnPoolShutdown) (event.Subscription, error) {

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolShutdown)
				if err := _YearnPool.contract.UnpackLog(event, "Shutdown", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseShutdown(log types.Log) (*YearnPoolShutdown, error) {
	event := new(YearnPoolShutdown)
	if err := _YearnPool.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolStrategyChangedIterator is returned from FilterStrategyChanged and is used to iterate over the raw logs and unpacked data for StrategyChanged events raised by the YearnPool contract.
type YearnPoolStrategyChangedIterator struct {
	Event *YearnPoolStrategyChanged // Event containing the contract specifics and raw log

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
func (it *YearnPoolStrategyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolStrategyChanged)
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
		it.Event = new(YearnPoolStrategyChanged)
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
func (it *YearnPoolStrategyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolStrategyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolStrategyChanged represents a StrategyChanged event raised by the YearnPool contract.
type YearnPoolStrategyChanged struct {
	Strategy   common.Address
	ChangeType *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterStrategyChanged is a free log retrieval operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_YearnPool *YearnPoolFilterer) FilterStrategyChanged(opts *bind.FilterOpts, strategy []common.Address, change_type []*big.Int) (*YearnPoolStrategyChangedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var change_typeRule []interface{}
	for _, change_typeItem := range change_type {
		change_typeRule = append(change_typeRule, change_typeItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "StrategyChanged", strategyRule, change_typeRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolStrategyChangedIterator{contract: _YearnPool.contract, event: "StrategyChanged", logs: logs, sub: sub}, nil
}

// WatchStrategyChanged is a free log subscription operation binding the contract event 0xde8ff765a5c5dad48d27bc9faa99836fb81f3b07c9dc62cfe005475d6b83a2ca.
//
// Solidity: event StrategyChanged(address indexed strategy, uint256 indexed change_type)
func (_YearnPool *YearnPoolFilterer) WatchStrategyChanged(opts *bind.WatchOpts, sink chan<- *YearnPoolStrategyChanged, strategy []common.Address, change_type []*big.Int) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}
	var change_typeRule []interface{}
	for _, change_typeItem := range change_type {
		change_typeRule = append(change_typeRule, change_typeItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "StrategyChanged", strategyRule, change_typeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolStrategyChanged)
				if err := _YearnPool.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseStrategyChanged(log types.Log) (*YearnPoolStrategyChanged, error) {
	event := new(YearnPoolStrategyChanged)
	if err := _YearnPool.contract.UnpackLog(event, "StrategyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolStrategyReportedIterator is returned from FilterStrategyReported and is used to iterate over the raw logs and unpacked data for StrategyReported events raised by the YearnPool contract.
type YearnPoolStrategyReportedIterator struct {
	Event *YearnPoolStrategyReported // Event containing the contract specifics and raw log

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
func (it *YearnPoolStrategyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolStrategyReported)
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
		it.Event = new(YearnPoolStrategyReported)
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
func (it *YearnPoolStrategyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolStrategyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolStrategyReported represents a StrategyReported event raised by the YearnPool contract.
type YearnPoolStrategyReported struct {
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
func (_YearnPool *YearnPoolFilterer) FilterStrategyReported(opts *bind.FilterOpts, strategy []common.Address) (*YearnPoolStrategyReportedIterator, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolStrategyReportedIterator{contract: _YearnPool.contract, event: "StrategyReported", logs: logs, sub: sub}, nil
}

// WatchStrategyReported is a free log subscription operation binding the contract event 0x7f2ad1d3ba35276f35ef140f83e3e0f17b23064fd710113d3f7a5ab30d267811.
//
// Solidity: event StrategyReported(address indexed strategy, uint256 gain, uint256 loss, uint256 current_debt, uint256 protocol_fees, uint256 total_fees, uint256 total_refunds)
func (_YearnPool *YearnPoolFilterer) WatchStrategyReported(opts *bind.WatchOpts, sink chan<- *YearnPoolStrategyReported, strategy []common.Address) (event.Subscription, error) {

	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "StrategyReported", strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolStrategyReported)
				if err := _YearnPool.contract.UnpackLog(event, "StrategyReported", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseStrategyReported(log types.Log) (*YearnPoolStrategyReported, error) {
	event := new(YearnPoolStrategyReported)
	if err := _YearnPool.contract.UnpackLog(event, "StrategyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YearnPool contract.
type YearnPoolTransferIterator struct {
	Event *YearnPoolTransfer // Event containing the contract specifics and raw log

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
func (it *YearnPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolTransfer)
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
		it.Event = new(YearnPoolTransfer)
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
func (it *YearnPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolTransfer represents a Transfer event raised by the YearnPool contract.
type YearnPoolTransfer struct {
	Sender   common.Address
	Receiver common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_YearnPool *YearnPoolFilterer) FilterTransfer(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*YearnPoolTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolTransferIterator{contract: _YearnPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed sender, address indexed receiver, uint256 value)
func (_YearnPool *YearnPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YearnPoolTransfer, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "Transfer", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolTransfer)
				if err := _YearnPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseTransfer(log types.Log) (*YearnPoolTransfer, error) {
	event := new(YearnPoolTransfer)
	if err := _YearnPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateAccountantIterator is returned from FilterUpdateAccountant and is used to iterate over the raw logs and unpacked data for UpdateAccountant events raised by the YearnPool contract.
type YearnPoolUpdateAccountantIterator struct {
	Event *YearnPoolUpdateAccountant // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateAccountantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateAccountant)
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
		it.Event = new(YearnPoolUpdateAccountant)
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
func (it *YearnPoolUpdateAccountantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateAccountantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateAccountant represents a UpdateAccountant event raised by the YearnPool contract.
type YearnPoolUpdateAccountant struct {
	Accountant common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdateAccountant is a free log retrieval operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_YearnPool *YearnPoolFilterer) FilterUpdateAccountant(opts *bind.FilterOpts, accountant []common.Address) (*YearnPoolUpdateAccountantIterator, error) {

	var accountantRule []interface{}
	for _, accountantItem := range accountant {
		accountantRule = append(accountantRule, accountantItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateAccountant", accountantRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateAccountantIterator{contract: _YearnPool.contract, event: "UpdateAccountant", logs: logs, sub: sub}, nil
}

// WatchUpdateAccountant is a free log subscription operation binding the contract event 0x28709a2dab2a5d5e8688e96159011151c51644ab21839a8a45b449634d7c8b2b.
//
// Solidity: event UpdateAccountant(address indexed accountant)
func (_YearnPool *YearnPoolFilterer) WatchUpdateAccountant(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateAccountant, accountant []common.Address) (event.Subscription, error) {

	var accountantRule []interface{}
	for _, accountantItem := range accountant {
		accountantRule = append(accountantRule, accountantItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateAccountant", accountantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateAccountant)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateAccountant", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateAccountant(log types.Log) (*YearnPoolUpdateAccountant, error) {
	event := new(YearnPoolUpdateAccountant)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateAccountant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateDefaultQueueIterator is returned from FilterUpdateDefaultQueue and is used to iterate over the raw logs and unpacked data for UpdateDefaultQueue events raised by the YearnPool contract.
type YearnPoolUpdateDefaultQueueIterator struct {
	Event *YearnPoolUpdateDefaultQueue // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateDefaultQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateDefaultQueue)
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
		it.Event = new(YearnPoolUpdateDefaultQueue)
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
func (it *YearnPoolUpdateDefaultQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateDefaultQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateDefaultQueue represents a UpdateDefaultQueue event raised by the YearnPool contract.
type YearnPoolUpdateDefaultQueue struct {
	NewDefaultQueue []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateDefaultQueue is a free log retrieval operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_YearnPool *YearnPoolFilterer) FilterUpdateDefaultQueue(opts *bind.FilterOpts) (*YearnPoolUpdateDefaultQueueIterator, error) {

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateDefaultQueue")
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateDefaultQueueIterator{contract: _YearnPool.contract, event: "UpdateDefaultQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateDefaultQueue is a free log subscription operation binding the contract event 0x0bc0cb8c5ccee13e6a2fd26a699f57ad7ff6e454e6aae97ec41cd2eb9ebd63a5.
//
// Solidity: event UpdateDefaultQueue(address[] new_default_queue)
func (_YearnPool *YearnPoolFilterer) WatchUpdateDefaultQueue(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateDefaultQueue) (event.Subscription, error) {

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateDefaultQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateDefaultQueue)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateDefaultQueue", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateDefaultQueue(log types.Log) (*YearnPoolUpdateDefaultQueue, error) {
	event := new(YearnPoolUpdateDefaultQueue)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateDefaultQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateDepositLimitIterator is returned from FilterUpdateDepositLimit and is used to iterate over the raw logs and unpacked data for UpdateDepositLimit events raised by the YearnPool contract.
type YearnPoolUpdateDepositLimitIterator struct {
	Event *YearnPoolUpdateDepositLimit // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateDepositLimit)
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
		it.Event = new(YearnPoolUpdateDepositLimit)
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
func (it *YearnPoolUpdateDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateDepositLimit represents a UpdateDepositLimit event raised by the YearnPool contract.
type YearnPoolUpdateDepositLimit struct {
	DepositLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimit is a free log retrieval operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_YearnPool *YearnPoolFilterer) FilterUpdateDepositLimit(opts *bind.FilterOpts) (*YearnPoolUpdateDepositLimitIterator, error) {

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateDepositLimitIterator{contract: _YearnPool.contract, event: "UpdateDepositLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimit is a free log subscription operation binding the contract event 0xae565aab888bca5e19e25a13db7b0c9144305bf55cb0f3f4d724f730e5acdd62.
//
// Solidity: event UpdateDepositLimit(uint256 deposit_limit)
func (_YearnPool *YearnPoolFilterer) WatchUpdateDepositLimit(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateDepositLimit) (event.Subscription, error) {

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateDepositLimit)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateDepositLimit(log types.Log) (*YearnPoolUpdateDepositLimit, error) {
	event := new(YearnPoolUpdateDepositLimit)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateDepositLimitModuleIterator is returned from FilterUpdateDepositLimitModule and is used to iterate over the raw logs and unpacked data for UpdateDepositLimitModule events raised by the YearnPool contract.
type YearnPoolUpdateDepositLimitModuleIterator struct {
	Event *YearnPoolUpdateDepositLimitModule // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateDepositLimitModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateDepositLimitModule)
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
		it.Event = new(YearnPoolUpdateDepositLimitModule)
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
func (it *YearnPoolUpdateDepositLimitModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateDepositLimitModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateDepositLimitModule represents a UpdateDepositLimitModule event raised by the YearnPool contract.
type YearnPoolUpdateDepositLimitModule struct {
	DepositLimitModule common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUpdateDepositLimitModule is a free log retrieval operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_YearnPool *YearnPoolFilterer) FilterUpdateDepositLimitModule(opts *bind.FilterOpts, deposit_limit_module []common.Address) (*YearnPoolUpdateDepositLimitModuleIterator, error) {

	var deposit_limit_moduleRule []interface{}
	for _, deposit_limit_moduleItem := range deposit_limit_module {
		deposit_limit_moduleRule = append(deposit_limit_moduleRule, deposit_limit_moduleItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateDepositLimitModule", deposit_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateDepositLimitModuleIterator{contract: _YearnPool.contract, event: "UpdateDepositLimitModule", logs: logs, sub: sub}, nil
}

// WatchUpdateDepositLimitModule is a free log subscription operation binding the contract event 0x777d215db24fb9fee4ded85c66b422abd7162a1caa6ed3ec4c031f6cd29ada52.
//
// Solidity: event UpdateDepositLimitModule(address indexed deposit_limit_module)
func (_YearnPool *YearnPoolFilterer) WatchUpdateDepositLimitModule(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateDepositLimitModule, deposit_limit_module []common.Address) (event.Subscription, error) {

	var deposit_limit_moduleRule []interface{}
	for _, deposit_limit_moduleItem := range deposit_limit_module {
		deposit_limit_moduleRule = append(deposit_limit_moduleRule, deposit_limit_moduleItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateDepositLimitModule", deposit_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateDepositLimitModule)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateDepositLimitModule", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateDepositLimitModule(log types.Log) (*YearnPoolUpdateDepositLimitModule, error) {
	event := new(YearnPoolUpdateDepositLimitModule)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateDepositLimitModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateMinimumTotalIdleIterator is returned from FilterUpdateMinimumTotalIdle and is used to iterate over the raw logs and unpacked data for UpdateMinimumTotalIdle events raised by the YearnPool contract.
type YearnPoolUpdateMinimumTotalIdleIterator struct {
	Event *YearnPoolUpdateMinimumTotalIdle // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateMinimumTotalIdleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateMinimumTotalIdle)
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
		it.Event = new(YearnPoolUpdateMinimumTotalIdle)
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
func (it *YearnPoolUpdateMinimumTotalIdleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateMinimumTotalIdleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateMinimumTotalIdle represents a UpdateMinimumTotalIdle event raised by the YearnPool contract.
type YearnPoolUpdateMinimumTotalIdle struct {
	MinimumTotalIdle *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinimumTotalIdle is a free log retrieval operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_YearnPool *YearnPoolFilterer) FilterUpdateMinimumTotalIdle(opts *bind.FilterOpts) (*YearnPoolUpdateMinimumTotalIdleIterator, error) {

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateMinimumTotalIdle")
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateMinimumTotalIdleIterator{contract: _YearnPool.contract, event: "UpdateMinimumTotalIdle", logs: logs, sub: sub}, nil
}

// WatchUpdateMinimumTotalIdle is a free log subscription operation binding the contract event 0x01a4494beed88920b88742cc58f2744e198f55ff192635a1fbabc6be8ffade81.
//
// Solidity: event UpdateMinimumTotalIdle(uint256 minimum_total_idle)
func (_YearnPool *YearnPoolFilterer) WatchUpdateMinimumTotalIdle(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateMinimumTotalIdle) (event.Subscription, error) {

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateMinimumTotalIdle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateMinimumTotalIdle)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateMinimumTotalIdle", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateMinimumTotalIdle(log types.Log) (*YearnPoolUpdateMinimumTotalIdle, error) {
	event := new(YearnPoolUpdateMinimumTotalIdle)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateMinimumTotalIdle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateProfitMaxUnlockTimeIterator is returned from FilterUpdateProfitMaxUnlockTime and is used to iterate over the raw logs and unpacked data for UpdateProfitMaxUnlockTime events raised by the YearnPool contract.
type YearnPoolUpdateProfitMaxUnlockTimeIterator struct {
	Event *YearnPoolUpdateProfitMaxUnlockTime // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateProfitMaxUnlockTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateProfitMaxUnlockTime)
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
		it.Event = new(YearnPoolUpdateProfitMaxUnlockTime)
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
func (it *YearnPoolUpdateProfitMaxUnlockTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateProfitMaxUnlockTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateProfitMaxUnlockTime represents a UpdateProfitMaxUnlockTime event raised by the YearnPool contract.
type YearnPoolUpdateProfitMaxUnlockTime struct {
	ProfitMaxUnlockTime *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateProfitMaxUnlockTime is a free log retrieval operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_YearnPool *YearnPoolFilterer) FilterUpdateProfitMaxUnlockTime(opts *bind.FilterOpts) (*YearnPoolUpdateProfitMaxUnlockTimeIterator, error) {

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateProfitMaxUnlockTimeIterator{contract: _YearnPool.contract, event: "UpdateProfitMaxUnlockTime", logs: logs, sub: sub}, nil
}

// WatchUpdateProfitMaxUnlockTime is a free log subscription operation binding the contract event 0xf361aed463da6fa20358e45c6209f1d3e16d4eca706e6eab0b0aeb338729c77a.
//
// Solidity: event UpdateProfitMaxUnlockTime(uint256 profit_max_unlock_time)
func (_YearnPool *YearnPoolFilterer) WatchUpdateProfitMaxUnlockTime(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateProfitMaxUnlockTime) (event.Subscription, error) {

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateProfitMaxUnlockTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateProfitMaxUnlockTime)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateProfitMaxUnlockTime(log types.Log) (*YearnPoolUpdateProfitMaxUnlockTime, error) {
	event := new(YearnPoolUpdateProfitMaxUnlockTime)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateProfitMaxUnlockTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateRoleManagerIterator is returned from FilterUpdateRoleManager and is used to iterate over the raw logs and unpacked data for UpdateRoleManager events raised by the YearnPool contract.
type YearnPoolUpdateRoleManagerIterator struct {
	Event *YearnPoolUpdateRoleManager // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateRoleManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateRoleManager)
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
		it.Event = new(YearnPoolUpdateRoleManager)
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
func (it *YearnPoolUpdateRoleManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateRoleManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateRoleManager represents a UpdateRoleManager event raised by the YearnPool contract.
type YearnPoolUpdateRoleManager struct {
	RoleManager common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateRoleManager is a free log retrieval operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_YearnPool *YearnPoolFilterer) FilterUpdateRoleManager(opts *bind.FilterOpts, role_manager []common.Address) (*YearnPoolUpdateRoleManagerIterator, error) {

	var role_managerRule []interface{}
	for _, role_managerItem := range role_manager {
		role_managerRule = append(role_managerRule, role_managerItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateRoleManager", role_managerRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateRoleManagerIterator{contract: _YearnPool.contract, event: "UpdateRoleManager", logs: logs, sub: sub}, nil
}

// WatchUpdateRoleManager is a free log subscription operation binding the contract event 0xce93baa0b608a7d420822b6b90cfcccb70574363ba4fd26ef5ac17dd465016c4.
//
// Solidity: event UpdateRoleManager(address indexed role_manager)
func (_YearnPool *YearnPoolFilterer) WatchUpdateRoleManager(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateRoleManager, role_manager []common.Address) (event.Subscription, error) {

	var role_managerRule []interface{}
	for _, role_managerItem := range role_manager {
		role_managerRule = append(role_managerRule, role_managerItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateRoleManager", role_managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateRoleManager)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateRoleManager", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateRoleManager(log types.Log) (*YearnPoolUpdateRoleManager, error) {
	event := new(YearnPoolUpdateRoleManager)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateRoleManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateUseDefaultQueueIterator is returned from FilterUpdateUseDefaultQueue and is used to iterate over the raw logs and unpacked data for UpdateUseDefaultQueue events raised by the YearnPool contract.
type YearnPoolUpdateUseDefaultQueueIterator struct {
	Event *YearnPoolUpdateUseDefaultQueue // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateUseDefaultQueueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateUseDefaultQueue)
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
		it.Event = new(YearnPoolUpdateUseDefaultQueue)
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
func (it *YearnPoolUpdateUseDefaultQueueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateUseDefaultQueueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateUseDefaultQueue represents a UpdateUseDefaultQueue event raised by the YearnPool contract.
type YearnPoolUpdateUseDefaultQueue struct {
	UseDefaultQueue bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateUseDefaultQueue is a free log retrieval operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_YearnPool *YearnPoolFilterer) FilterUpdateUseDefaultQueue(opts *bind.FilterOpts) (*YearnPoolUpdateUseDefaultQueueIterator, error) {

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateUseDefaultQueue")
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateUseDefaultQueueIterator{contract: _YearnPool.contract, event: "UpdateUseDefaultQueue", logs: logs, sub: sub}, nil
}

// WatchUpdateUseDefaultQueue is a free log subscription operation binding the contract event 0x1f88e73ebc721f227812938fe07a069ec1f7136aafacb397ed460bd15dee13f1.
//
// Solidity: event UpdateUseDefaultQueue(bool use_default_queue)
func (_YearnPool *YearnPoolFilterer) WatchUpdateUseDefaultQueue(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateUseDefaultQueue) (event.Subscription, error) {

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateUseDefaultQueue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateUseDefaultQueue)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateUseDefaultQueue", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateUseDefaultQueue(log types.Log) (*YearnPoolUpdateUseDefaultQueue, error) {
	event := new(YearnPoolUpdateUseDefaultQueue)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateUseDefaultQueue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdateWithdrawLimitModuleIterator is returned from FilterUpdateWithdrawLimitModule and is used to iterate over the raw logs and unpacked data for UpdateWithdrawLimitModule events raised by the YearnPool contract.
type YearnPoolUpdateWithdrawLimitModuleIterator struct {
	Event *YearnPoolUpdateWithdrawLimitModule // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdateWithdrawLimitModuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdateWithdrawLimitModule)
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
		it.Event = new(YearnPoolUpdateWithdrawLimitModule)
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
func (it *YearnPoolUpdateWithdrawLimitModuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdateWithdrawLimitModuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdateWithdrawLimitModule represents a UpdateWithdrawLimitModule event raised by the YearnPool contract.
type YearnPoolUpdateWithdrawLimitModule struct {
	WithdrawLimitModule common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUpdateWithdrawLimitModule is a free log retrieval operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_YearnPool *YearnPoolFilterer) FilterUpdateWithdrawLimitModule(opts *bind.FilterOpts, withdraw_limit_module []common.Address) (*YearnPoolUpdateWithdrawLimitModuleIterator, error) {

	var withdraw_limit_moduleRule []interface{}
	for _, withdraw_limit_moduleItem := range withdraw_limit_module {
		withdraw_limit_moduleRule = append(withdraw_limit_moduleRule, withdraw_limit_moduleItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdateWithdrawLimitModule", withdraw_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdateWithdrawLimitModuleIterator{contract: _YearnPool.contract, event: "UpdateWithdrawLimitModule", logs: logs, sub: sub}, nil
}

// WatchUpdateWithdrawLimitModule is a free log subscription operation binding the contract event 0xce6e3f8beda82a13c441d76efd4a6335760f219f38c60502e6680060874e109d.
//
// Solidity: event UpdateWithdrawLimitModule(address indexed withdraw_limit_module)
func (_YearnPool *YearnPoolFilterer) WatchUpdateWithdrawLimitModule(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdateWithdrawLimitModule, withdraw_limit_module []common.Address) (event.Subscription, error) {

	var withdraw_limit_moduleRule []interface{}
	for _, withdraw_limit_moduleItem := range withdraw_limit_module {
		withdraw_limit_moduleRule = append(withdraw_limit_moduleRule, withdraw_limit_moduleItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdateWithdrawLimitModule", withdraw_limit_moduleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdateWithdrawLimitModule)
				if err := _YearnPool.contract.UnpackLog(event, "UpdateWithdrawLimitModule", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdateWithdrawLimitModule(log types.Log) (*YearnPoolUpdateWithdrawLimitModule, error) {
	event := new(YearnPoolUpdateWithdrawLimitModule)
	if err := _YearnPool.contract.UnpackLog(event, "UpdateWithdrawLimitModule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolUpdatedMaxDebtForStrategyIterator is returned from FilterUpdatedMaxDebtForStrategy and is used to iterate over the raw logs and unpacked data for UpdatedMaxDebtForStrategy events raised by the YearnPool contract.
type YearnPoolUpdatedMaxDebtForStrategyIterator struct {
	Event *YearnPoolUpdatedMaxDebtForStrategy // Event containing the contract specifics and raw log

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
func (it *YearnPoolUpdatedMaxDebtForStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolUpdatedMaxDebtForStrategy)
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
		it.Event = new(YearnPoolUpdatedMaxDebtForStrategy)
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
func (it *YearnPoolUpdatedMaxDebtForStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolUpdatedMaxDebtForStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolUpdatedMaxDebtForStrategy represents a UpdatedMaxDebtForStrategy event raised by the YearnPool contract.
type YearnPoolUpdatedMaxDebtForStrategy struct {
	Sender   common.Address
	Strategy common.Address
	NewDebt  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMaxDebtForStrategy is a free log retrieval operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_YearnPool *YearnPoolFilterer) FilterUpdatedMaxDebtForStrategy(opts *bind.FilterOpts, sender []common.Address, strategy []common.Address) (*YearnPoolUpdatedMaxDebtForStrategyIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "UpdatedMaxDebtForStrategy", senderRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolUpdatedMaxDebtForStrategyIterator{contract: _YearnPool.contract, event: "UpdatedMaxDebtForStrategy", logs: logs, sub: sub}, nil
}

// WatchUpdatedMaxDebtForStrategy is a free log subscription operation binding the contract event 0xb3eef2123fec1523a6bbc90aceb203000154c1a4974335fe06b544c7534d4b89.
//
// Solidity: event UpdatedMaxDebtForStrategy(address indexed sender, address indexed strategy, uint256 new_debt)
func (_YearnPool *YearnPoolFilterer) WatchUpdatedMaxDebtForStrategy(opts *bind.WatchOpts, sink chan<- *YearnPoolUpdatedMaxDebtForStrategy, sender []common.Address, strategy []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "UpdatedMaxDebtForStrategy", senderRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolUpdatedMaxDebtForStrategy)
				if err := _YearnPool.contract.UnpackLog(event, "UpdatedMaxDebtForStrategy", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseUpdatedMaxDebtForStrategy(log types.Log) (*YearnPoolUpdatedMaxDebtForStrategy, error) {
	event := new(YearnPoolUpdatedMaxDebtForStrategy)
	if err := _YearnPool.contract.UnpackLog(event, "UpdatedMaxDebtForStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the YearnPool contract.
type YearnPoolWithdrawIterator struct {
	Event *YearnPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *YearnPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnPoolWithdraw)
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
		it.Event = new(YearnPoolWithdraw)
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
func (it *YearnPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnPoolWithdraw represents a Withdraw event raised by the YearnPool contract.
type YearnPoolWithdraw struct {
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
func (_YearnPool *YearnPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*YearnPoolWithdrawIterator, error) {

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

	logs, sub, err := _YearnPool.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YearnPoolWithdrawIterator{contract: _YearnPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YearnPool *YearnPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *YearnPoolWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _YearnPool.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnPoolWithdraw)
				if err := _YearnPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_YearnPool *YearnPoolFilterer) ParseWithdraw(log types.Log) (*YearnPoolWithdraw, error) {
	event := new(YearnPoolWithdraw)
	if err := _YearnPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
