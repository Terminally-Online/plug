// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearn_v3_gauge

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

// YearnV3GaugeMetaData contains all meta data concerning the YearnV3Gauge contract.
var YearnV3GaugeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_veYfi\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dYfi\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_veYfiDYfiPool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BoostedBalanceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"}],\"name\":\"DurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"RecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"historicalRewards\",\"type\":\"uint256\"}],\"name\":\"RewardsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sweep\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transfered\",\"type\":\"uint256\"}],\"name\":\"TransferredPenalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"userRewardPerTokenPaid\",\"type\":\"uint256\"}],\"name\":\"UpdatedRewards\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOSTING_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOOST_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRECISION_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_TOKEN\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VEYFI\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VE_YFI_POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"boostedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"getReward\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"historicalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_accounts\",\"type\":\"address[]\"}],\"name\":\"kick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"nextBoostedBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodFinish\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"queueNewRewards\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerTokenStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newDuration\",\"type\":\"uint256\"}],\"name\":\"setDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"setRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_claim\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_claim\",\"type\":\"bool\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// YearnV3GaugeABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnV3GaugeMetaData.ABI instead.
var YearnV3GaugeABI = YearnV3GaugeMetaData.ABI

// YearnV3Gauge is an auto generated Go binding around an Ethereum contract.
type YearnV3Gauge struct {
	YearnV3GaugeCaller     // Read-only binding to the contract
	YearnV3GaugeTransactor // Write-only binding to the contract
	YearnV3GaugeFilterer   // Log filterer for contract events
}

// YearnV3GaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnV3GaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3GaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnV3GaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3GaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnV3GaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3GaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnV3GaugeSession struct {
	Contract     *YearnV3Gauge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnV3GaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnV3GaugeCallerSession struct {
	Contract *YearnV3GaugeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// YearnV3GaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnV3GaugeTransactorSession struct {
	Contract     *YearnV3GaugeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// YearnV3GaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnV3GaugeRaw struct {
	Contract *YearnV3Gauge // Generic contract binding to access the raw methods on
}

// YearnV3GaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnV3GaugeCallerRaw struct {
	Contract *YearnV3GaugeCaller // Generic read-only contract binding to access the raw methods on
}

// YearnV3GaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnV3GaugeTransactorRaw struct {
	Contract *YearnV3GaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnV3Gauge creates a new instance of YearnV3Gauge, bound to a specific deployed contract.
func NewYearnV3Gauge(address common.Address, backend bind.ContractBackend) (*YearnV3Gauge, error) {
	contract, err := bindYearnV3Gauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnV3Gauge{YearnV3GaugeCaller: YearnV3GaugeCaller{contract: contract}, YearnV3GaugeTransactor: YearnV3GaugeTransactor{contract: contract}, YearnV3GaugeFilterer: YearnV3GaugeFilterer{contract: contract}}, nil
}

// NewYearnV3GaugeCaller creates a new read-only instance of YearnV3Gauge, bound to a specific deployed contract.
func NewYearnV3GaugeCaller(address common.Address, caller bind.ContractCaller) (*YearnV3GaugeCaller, error) {
	contract, err := bindYearnV3Gauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeCaller{contract: contract}, nil
}

// NewYearnV3GaugeTransactor creates a new write-only instance of YearnV3Gauge, bound to a specific deployed contract.
func NewYearnV3GaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnV3GaugeTransactor, error) {
	contract, err := bindYearnV3Gauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeTransactor{contract: contract}, nil
}

// NewYearnV3GaugeFilterer creates a new log filterer instance of YearnV3Gauge, bound to a specific deployed contract.
func NewYearnV3GaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnV3GaugeFilterer, error) {
	contract, err := bindYearnV3Gauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeFilterer{contract: contract}, nil
}

// bindYearnV3Gauge binds a generic wrapper to an already deployed contract.
func bindYearnV3Gauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YearnV3GaugeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnV3Gauge *YearnV3GaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnV3Gauge.Contract.YearnV3GaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnV3Gauge *YearnV3GaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.YearnV3GaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnV3Gauge *YearnV3GaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.YearnV3GaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnV3Gauge *YearnV3GaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnV3Gauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnV3Gauge *YearnV3GaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnV3Gauge *YearnV3GaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.contract.Transact(opts, method, params...)
}

// BOOSTINGFACTOR is a free data retrieval call binding the contract method 0x980091fc.
//
// Solidity: function BOOSTING_FACTOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) BOOSTINGFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "BOOSTING_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTINGFACTOR is a free data retrieval call binding the contract method 0x980091fc.
//
// Solidity: function BOOSTING_FACTOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) BOOSTINGFACTOR() (*big.Int, error) {
	return _YearnV3Gauge.Contract.BOOSTINGFACTOR(&_YearnV3Gauge.CallOpts)
}

// BOOSTINGFACTOR is a free data retrieval call binding the contract method 0x980091fc.
//
// Solidity: function BOOSTING_FACTOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) BOOSTINGFACTOR() (*big.Int, error) {
	return _YearnV3Gauge.Contract.BOOSTINGFACTOR(&_YearnV3Gauge.CallOpts)
}

// BOOSTDENOMINATOR is a free data retrieval call binding the contract method 0x1bd32ed3.
//
// Solidity: function BOOST_DENOMINATOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) BOOSTDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "BOOST_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BOOSTDENOMINATOR is a free data retrieval call binding the contract method 0x1bd32ed3.
//
// Solidity: function BOOST_DENOMINATOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) BOOSTDENOMINATOR() (*big.Int, error) {
	return _YearnV3Gauge.Contract.BOOSTDENOMINATOR(&_YearnV3Gauge.CallOpts)
}

// BOOSTDENOMINATOR is a free data retrieval call binding the contract method 0x1bd32ed3.
//
// Solidity: function BOOST_DENOMINATOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) BOOSTDENOMINATOR() (*big.Int, error) {
	return _YearnV3Gauge.Contract.BOOSTDENOMINATOR(&_YearnV3Gauge.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) PRECISIONFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "PRECISION_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) PRECISIONFACTOR() (*big.Int, error) {
	return _YearnV3Gauge.Contract.PRECISIONFACTOR(&_YearnV3Gauge.CallOpts)
}

// PRECISIONFACTOR is a free data retrieval call binding the contract method 0xccd34cd5.
//
// Solidity: function PRECISION_FACTOR() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) PRECISIONFACTOR() (*big.Int, error) {
	return _YearnV3Gauge.Contract.PRECISIONFACTOR(&_YearnV3Gauge.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCaller) REWARDTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "REWARD_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeSession) REWARDTOKEN() (common.Address, error) {
	return _YearnV3Gauge.Contract.REWARDTOKEN(&_YearnV3Gauge.CallOpts)
}

// REWARDTOKEN is a free data retrieval call binding the contract method 0x99248ea7.
//
// Solidity: function REWARD_TOKEN() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) REWARDTOKEN() (common.Address, error) {
	return _YearnV3Gauge.Contract.REWARDTOKEN(&_YearnV3Gauge.CallOpts)
}

// VEYFI is a free data retrieval call binding the contract method 0x7d2f791d.
//
// Solidity: function VEYFI() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCaller) VEYFI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "VEYFI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEYFI is a free data retrieval call binding the contract method 0x7d2f791d.
//
// Solidity: function VEYFI() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeSession) VEYFI() (common.Address, error) {
	return _YearnV3Gauge.Contract.VEYFI(&_YearnV3Gauge.CallOpts)
}

// VEYFI is a free data retrieval call binding the contract method 0x7d2f791d.
//
// Solidity: function VEYFI() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) VEYFI() (common.Address, error) {
	return _YearnV3Gauge.Contract.VEYFI(&_YearnV3Gauge.CallOpts)
}

// VEYFIPOOL is a free data retrieval call binding the contract method 0xb5387c78.
//
// Solidity: function VE_YFI_POOL() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCaller) VEYFIPOOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "VE_YFI_POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VEYFIPOOL is a free data retrieval call binding the contract method 0xb5387c78.
//
// Solidity: function VE_YFI_POOL() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeSession) VEYFIPOOL() (common.Address, error) {
	return _YearnV3Gauge.Contract.VEYFIPOOL(&_YearnV3Gauge.CallOpts)
}

// VEYFIPOOL is a free data retrieval call binding the contract method 0xb5387c78.
//
// Solidity: function VE_YFI_POOL() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) VEYFIPOOL() (common.Address, error) {
	return _YearnV3Gauge.Contract.VEYFIPOOL(&_YearnV3Gauge.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.Allowance(&_YearnV3Gauge.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.Allowance(&_YearnV3Gauge.CallOpts, owner, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeSession) Asset() (common.Address, error) {
	return _YearnV3Gauge.Contract.Asset(&_YearnV3Gauge.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Asset() (common.Address, error) {
	return _YearnV3Gauge.Contract.Asset(&_YearnV3Gauge.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.BalanceOf(&_YearnV3Gauge.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.BalanceOf(&_YearnV3Gauge.CallOpts, account)
}

// BoostedBalanceOf is a free data retrieval call binding the contract method 0x1beabcd2.
//
// Solidity: function boostedBalanceOf(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) BoostedBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "boostedBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoostedBalanceOf is a free data retrieval call binding the contract method 0x1beabcd2.
//
// Solidity: function boostedBalanceOf(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) BoostedBalanceOf(_account common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.BoostedBalanceOf(&_YearnV3Gauge.CallOpts, _account)
}

// BoostedBalanceOf is a free data retrieval call binding the contract method 0x1beabcd2.
//
// Solidity: function boostedBalanceOf(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) BoostedBalanceOf(_account common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.BoostedBalanceOf(&_YearnV3Gauge.CallOpts, _account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) ConvertToAssets(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "convertToAssets", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.ConvertToAssets(&_YearnV3Gauge.CallOpts, _shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 _shares) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) ConvertToAssets(_shares *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.ConvertToAssets(&_YearnV3Gauge.CallOpts, _shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) ConvertToShares(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "convertToShares", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.ConvertToShares(&_YearnV3Gauge.CallOpts, _assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) ConvertToShares(_assets *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.ConvertToShares(&_YearnV3Gauge.CallOpts, _assets)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) CurrentRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "currentRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) CurrentRewards() (*big.Int, error) {
	return _YearnV3Gauge.Contract.CurrentRewards(&_YearnV3Gauge.CallOpts)
}

// CurrentRewards is a free data retrieval call binding the contract method 0x901a7d53.
//
// Solidity: function currentRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) CurrentRewards() (*big.Int, error) {
	return _YearnV3Gauge.Contract.CurrentRewards(&_YearnV3Gauge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnV3Gauge *YearnV3GaugeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnV3Gauge *YearnV3GaugeSession) Decimals() (uint8, error) {
	return _YearnV3Gauge.Contract.Decimals(&_YearnV3Gauge.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Decimals() (uint8, error) {
	return _YearnV3Gauge.Contract.Decimals(&_YearnV3Gauge.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Duration() (*big.Int, error) {
	return _YearnV3Gauge.Contract.Duration(&_YearnV3Gauge.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Duration() (*big.Int, error) {
	return _YearnV3Gauge.Contract.Duration(&_YearnV3Gauge.CallOpts)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) Earned(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "earned", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Earned(_account common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.Earned(&_YearnV3Gauge.CallOpts, _account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Earned(_account common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.Earned(&_YearnV3Gauge.CallOpts, _account)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) HistoricalRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "historicalRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) HistoricalRewards() (*big.Int, error) {
	return _YearnV3Gauge.Contract.HistoricalRewards(&_YearnV3Gauge.CallOpts)
}

// HistoricalRewards is a free data retrieval call binding the contract method 0x262d3d6d.
//
// Solidity: function historicalRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) HistoricalRewards() (*big.Int, error) {
	return _YearnV3Gauge.Contract.HistoricalRewards(&_YearnV3Gauge.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) LastTimeRewardApplicable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "lastTimeRewardApplicable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _YearnV3Gauge.Contract.LastTimeRewardApplicable(&_YearnV3Gauge.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x80faa57d.
//
// Solidity: function lastTimeRewardApplicable() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) LastTimeRewardApplicable() (*big.Int, error) {
	return _YearnV3Gauge.Contract.LastTimeRewardApplicable(&_YearnV3Gauge.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) LastUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "lastUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) LastUpdateTime() (*big.Int, error) {
	return _YearnV3Gauge.Contract.LastUpdateTime(&_YearnV3Gauge.CallOpts)
}

// LastUpdateTime is a free data retrieval call binding the contract method 0xc8f33c91.
//
// Solidity: function lastUpdateTime() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) LastUpdateTime() (*big.Int, error) {
	return _YearnV3Gauge.Contract.LastUpdateTime(&_YearnV3Gauge.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.MaxDeposit(&_YearnV3Gauge.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.MaxDeposit(&_YearnV3Gauge.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.MaxMint(&_YearnV3Gauge.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.MaxMint(&_YearnV3Gauge.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) MaxRedeem(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "maxRedeem", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.MaxRedeem(&_YearnV3Gauge.CallOpts, _owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address _owner) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) MaxRedeem(_owner common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.MaxRedeem(&_YearnV3Gauge.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) MaxWithdraw(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "maxWithdraw", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.MaxWithdraw(&_YearnV3Gauge.CallOpts, _owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address _owner) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) MaxWithdraw(_owner common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.MaxWithdraw(&_YearnV3Gauge.CallOpts, _owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Gauge *YearnV3GaugeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Gauge *YearnV3GaugeSession) Name() (string, error) {
	return _YearnV3Gauge.Contract.Name(&_YearnV3Gauge.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Name() (string, error) {
	return _YearnV3Gauge.Contract.Name(&_YearnV3Gauge.CallOpts)
}

// NextBoostedBalanceOf is a free data retrieval call binding the contract method 0xc67ffb4e.
//
// Solidity: function nextBoostedBalanceOf(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) NextBoostedBalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "nextBoostedBalanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBoostedBalanceOf is a free data retrieval call binding the contract method 0xc67ffb4e.
//
// Solidity: function nextBoostedBalanceOf(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) NextBoostedBalanceOf(_account common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.NextBoostedBalanceOf(&_YearnV3Gauge.CallOpts, _account)
}

// NextBoostedBalanceOf is a free data retrieval call binding the contract method 0xc67ffb4e.
//
// Solidity: function nextBoostedBalanceOf(address _account) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) NextBoostedBalanceOf(_account common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.NextBoostedBalanceOf(&_YearnV3Gauge.CallOpts, _account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeSession) Owner() (common.Address, error) {
	return _YearnV3Gauge.Contract.Owner(&_YearnV3Gauge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Owner() (common.Address, error) {
	return _YearnV3Gauge.Contract.Owner(&_YearnV3Gauge.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) PeriodFinish(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "periodFinish")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) PeriodFinish() (*big.Int, error) {
	return _YearnV3Gauge.Contract.PeriodFinish(&_YearnV3Gauge.CallOpts)
}

// PeriodFinish is a free data retrieval call binding the contract method 0xebe2b12b.
//
// Solidity: function periodFinish() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) PeriodFinish() (*big.Int, error) {
	return _YearnV3Gauge.Contract.PeriodFinish(&_YearnV3Gauge.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) PreviewDeposit(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "previewDeposit", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.PreviewDeposit(&_YearnV3Gauge.CallOpts, _assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) PreviewDeposit(_assets *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.PreviewDeposit(&_YearnV3Gauge.CallOpts, _assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) PreviewMint(opts *bind.CallOpts, _shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "previewMint", _shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.PreviewMint(&_YearnV3Gauge.CallOpts, _shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 _shares) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) PreviewMint(_shares *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.PreviewMint(&_YearnV3Gauge.CallOpts, _shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) PreviewRedeem(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "previewRedeem", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) PreviewRedeem(_assets *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.PreviewRedeem(&_YearnV3Gauge.CallOpts, _assets)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) PreviewRedeem(_assets *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.PreviewRedeem(&_YearnV3Gauge.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) PreviewWithdraw(opts *bind.CallOpts, _assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "previewWithdraw", _assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.PreviewWithdraw(&_YearnV3Gauge.CallOpts, _assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 _assets) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) PreviewWithdraw(_assets *big.Int) (*big.Int, error) {
	return _YearnV3Gauge.Contract.PreviewWithdraw(&_YearnV3Gauge.CallOpts, _assets)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) QueuedRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "queuedRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) QueuedRewards() (*big.Int, error) {
	return _YearnV3Gauge.Contract.QueuedRewards(&_YearnV3Gauge.CallOpts)
}

// QueuedRewards is a free data retrieval call binding the contract method 0x63d38c3b.
//
// Solidity: function queuedRewards() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) QueuedRewards() (*big.Int, error) {
	return _YearnV3Gauge.Contract.QueuedRewards(&_YearnV3Gauge.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "recipients", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_YearnV3Gauge *YearnV3GaugeSession) Recipients(arg0 common.Address) (common.Address, error) {
	return _YearnV3Gauge.Contract.Recipients(&_YearnV3Gauge.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(address)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Recipients(arg0 common.Address) (common.Address, error) {
	return _YearnV3Gauge.Contract.Recipients(&_YearnV3Gauge.CallOpts, arg0)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) RewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "rewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) RewardPerToken() (*big.Int, error) {
	return _YearnV3Gauge.Contract.RewardPerToken(&_YearnV3Gauge.CallOpts)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xcd3daf9d.
//
// Solidity: function rewardPerToken() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) RewardPerToken() (*big.Int, error) {
	return _YearnV3Gauge.Contract.RewardPerToken(&_YearnV3Gauge.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) RewardPerTokenStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "rewardPerTokenStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) RewardPerTokenStored() (*big.Int, error) {
	return _YearnV3Gauge.Contract.RewardPerTokenStored(&_YearnV3Gauge.CallOpts)
}

// RewardPerTokenStored is a free data retrieval call binding the contract method 0xdf136d65.
//
// Solidity: function rewardPerTokenStored() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) RewardPerTokenStored() (*big.Int, error) {
	return _YearnV3Gauge.Contract.RewardPerTokenStored(&_YearnV3Gauge.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) RewardRate() (*big.Int, error) {
	return _YearnV3Gauge.Contract.RewardRate(&_YearnV3Gauge.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) RewardRate() (*big.Int, error) {
	return _YearnV3Gauge.Contract.RewardRate(&_YearnV3Gauge.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.Rewards(&_YearnV3Gauge.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.Rewards(&_YearnV3Gauge.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnV3Gauge *YearnV3GaugeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnV3Gauge *YearnV3GaugeSession) Symbol() (string, error) {
	return _YearnV3Gauge.Contract.Symbol(&_YearnV3Gauge.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) Symbol() (string, error) {
	return _YearnV3Gauge.Contract.Symbol(&_YearnV3Gauge.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) TotalAssets() (*big.Int, error) {
	return _YearnV3Gauge.Contract.TotalAssets(&_YearnV3Gauge.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) TotalAssets() (*big.Int, error) {
	return _YearnV3Gauge.Contract.TotalAssets(&_YearnV3Gauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) TotalSupply() (*big.Int, error) {
	return _YearnV3Gauge.Contract.TotalSupply(&_YearnV3Gauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _YearnV3Gauge.Contract.TotalSupply(&_YearnV3Gauge.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Gauge.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.UserRewardPerTokenPaid(&_YearnV3Gauge.CallOpts, arg0)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x8b876347.
//
// Solidity: function userRewardPerTokenPaid(address ) view returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeCallerSession) UserRewardPerTokenPaid(arg0 common.Address) (*big.Int, error) {
	return _YearnV3Gauge.Contract.UserRewardPerTokenPaid(&_YearnV3Gauge.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Approve(&_YearnV3Gauge.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Approve(&_YearnV3Gauge.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.DecreaseAllowance(&_YearnV3Gauge.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.DecreaseAllowance(&_YearnV3Gauge.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Deposit(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "deposit", _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Deposit(&_YearnV3Gauge.TransactOpts, _assets, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _assets, address _receiver) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Deposit(_assets *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Deposit(&_YearnV3Gauge.TransactOpts, _assets, _receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Deposit0(opts *bind.TransactOpts, _assets *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "deposit0", _assets)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Deposit0(_assets *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Deposit0(&_YearnV3Gauge.TransactOpts, _assets)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _assets) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Deposit0(_assets *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Deposit0(&_YearnV3Gauge.TransactOpts, _assets)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Deposit1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "deposit1")
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Deposit1() (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Deposit1(&_YearnV3Gauge.TransactOpts)
}

// Deposit1 is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Deposit1() (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Deposit1(&_YearnV3Gauge.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) GetReward() (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.GetReward(&_YearnV3Gauge.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) GetReward() (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.GetReward(&_YearnV3Gauge.TransactOpts)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) GetReward0(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "getReward0", _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.GetReward0(&_YearnV3Gauge.TransactOpts, _account)
}

// GetReward0 is a paid mutator transaction binding the contract method 0xc00007b0.
//
// Solidity: function getReward(address _account) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) GetReward0(_account common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.GetReward0(&_YearnV3Gauge.TransactOpts, _account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.IncreaseAllowance(&_YearnV3Gauge.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.IncreaseAllowance(&_YearnV3Gauge.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _asset, address _owner) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactor) Initialize(opts *bind.TransactOpts, _asset common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "initialize", _asset, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _asset, address _owner) returns()
func (_YearnV3Gauge *YearnV3GaugeSession) Initialize(_asset common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Initialize(&_YearnV3Gauge.TransactOpts, _asset, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _asset, address _owner) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Initialize(_asset common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Initialize(&_YearnV3Gauge.TransactOpts, _asset, _owner)
}

// Kick is a paid mutator transaction binding the contract method 0x1530e6d8.
//
// Solidity: function kick(address[] _accounts) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactor) Kick(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "kick", _accounts)
}

// Kick is a paid mutator transaction binding the contract method 0x1530e6d8.
//
// Solidity: function kick(address[] _accounts) returns()
func (_YearnV3Gauge *YearnV3GaugeSession) Kick(_accounts []common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Kick(&_YearnV3Gauge.TransactOpts, _accounts)
}

// Kick is a paid mutator transaction binding the contract method 0x1530e6d8.
//
// Solidity: function kick(address[] _accounts) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Kick(_accounts []common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Kick(&_YearnV3Gauge.TransactOpts, _accounts)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Mint(opts *bind.TransactOpts, _shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "mint", _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Mint(&_YearnV3Gauge.TransactOpts, _shares, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 _shares, address _receiver) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Mint(_shares *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Mint(&_YearnV3Gauge.TransactOpts, _shares, _receiver)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) QueueNewRewards(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "queueNewRewards", _amount)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) QueueNewRewards(_amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.QueueNewRewards(&_YearnV3Gauge.TransactOpts, _amount)
}

// QueueNewRewards is a paid mutator transaction binding the contract method 0x590a41f5.
//
// Solidity: function queueNewRewards(uint256 _amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) QueueNewRewards(_amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.QueueNewRewards(&_YearnV3Gauge.TransactOpts, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Redeem(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "redeem", _assets, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Redeem(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Redeem(&_YearnV3Gauge.TransactOpts, _assets, _receiver, _owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Redeem(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Redeem(&_YearnV3Gauge.TransactOpts, _assets, _receiver, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YearnV3Gauge *YearnV3GaugeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YearnV3Gauge *YearnV3GaugeSession) RenounceOwnership() (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.RenounceOwnership(&_YearnV3Gauge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.RenounceOwnership(&_YearnV3Gauge.TransactOpts)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _newDuration) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactor) SetDuration(opts *bind.TransactOpts, _newDuration *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "setDuration", _newDuration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _newDuration) returns()
func (_YearnV3Gauge *YearnV3GaugeSession) SetDuration(_newDuration *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.SetDuration(&_YearnV3Gauge.TransactOpts, _newDuration)
}

// SetDuration is a paid mutator transaction binding the contract method 0xf6be71d1.
//
// Solidity: function setDuration(uint256 _newDuration) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) SetDuration(_newDuration *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.SetDuration(&_YearnV3Gauge.TransactOpts, _newDuration)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactor) SetRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "setRecipient", _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_YearnV3Gauge *YearnV3GaugeSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.SetRecipient(&_YearnV3Gauge.TransactOpts, _recipient)
}

// SetRecipient is a paid mutator transaction binding the contract method 0x3bbed4a0.
//
// Solidity: function setRecipient(address _recipient) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) SetRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.SetRecipient(&_YearnV3Gauge.TransactOpts, _recipient)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Sweep(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "sweep", _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Sweep(&_YearnV3Gauge.TransactOpts, _token)
}

// Sweep is a paid mutator transaction binding the contract method 0x01681a62.
//
// Solidity: function sweep(address _token) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Sweep(_token common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Sweep(&_YearnV3Gauge.TransactOpts, _token)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Transfer(&_YearnV3Gauge.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Transfer(&_YearnV3Gauge.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.TransferFrom(&_YearnV3Gauge.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.TransferFrom(&_YearnV3Gauge.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YearnV3Gauge *YearnV3GaugeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.TransferOwnership(&_YearnV3Gauge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.TransferOwnership(&_YearnV3Gauge.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Withdraw() (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Withdraw(&_YearnV3Gauge.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Withdraw() (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Withdraw(&_YearnV3Gauge.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa810a54c.
//
// Solidity: function withdraw(bool _claim) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Withdraw0(opts *bind.TransactOpts, _claim bool) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "withdraw0", _claim)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa810a54c.
//
// Solidity: function withdraw(bool _claim) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Withdraw0(_claim bool) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Withdraw0(&_YearnV3Gauge.TransactOpts, _claim)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0xa810a54c.
//
// Solidity: function withdraw(bool _claim) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Withdraw0(_claim bool) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Withdraw0(&_YearnV3Gauge.TransactOpts, _claim)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Withdraw1(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "withdraw1", _assets, _receiver, _owner)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Withdraw1(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Withdraw1(&_YearnV3Gauge.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw1 is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Withdraw1(_assets *big.Int, _receiver common.Address, _owner common.Address) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Withdraw1(&_YearnV3Gauge.TransactOpts, _assets, _receiver, _owner)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xd045f2c4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, bool _claim) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactor) Withdraw2(opts *bind.TransactOpts, _assets *big.Int, _receiver common.Address, _owner common.Address, _claim bool) (*types.Transaction, error) {
	return _YearnV3Gauge.contract.Transact(opts, "withdraw2", _assets, _receiver, _owner, _claim)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xd045f2c4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, bool _claim) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeSession) Withdraw2(_assets *big.Int, _receiver common.Address, _owner common.Address, _claim bool) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Withdraw2(&_YearnV3Gauge.TransactOpts, _assets, _receiver, _owner, _claim)
}

// Withdraw2 is a paid mutator transaction binding the contract method 0xd045f2c4.
//
// Solidity: function withdraw(uint256 _assets, address _receiver, address _owner, bool _claim) returns(uint256)
func (_YearnV3Gauge *YearnV3GaugeTransactorSession) Withdraw2(_assets *big.Int, _receiver common.Address, _owner common.Address, _claim bool) (*types.Transaction, error) {
	return _YearnV3Gauge.Contract.Withdraw2(&_YearnV3Gauge.TransactOpts, _assets, _receiver, _owner, _claim)
}

// YearnV3GaugeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YearnV3Gauge contract.
type YearnV3GaugeApprovalIterator struct {
	Event *YearnV3GaugeApproval // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeApproval)
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
		it.Event = new(YearnV3GaugeApproval)
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
func (it *YearnV3GaugeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeApproval represents a Approval event raised by the YearnV3Gauge contract.
type YearnV3GaugeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YearnV3GaugeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeApprovalIterator{contract: _YearnV3Gauge.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeApproval)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseApproval(log types.Log) (*YearnV3GaugeApproval, error) {
	event := new(YearnV3GaugeApproval)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeBoostedBalanceUpdatedIterator is returned from FilterBoostedBalanceUpdated and is used to iterate over the raw logs and unpacked data for BoostedBalanceUpdated events raised by the YearnV3Gauge contract.
type YearnV3GaugeBoostedBalanceUpdatedIterator struct {
	Event *YearnV3GaugeBoostedBalanceUpdated // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeBoostedBalanceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeBoostedBalanceUpdated)
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
		it.Event = new(YearnV3GaugeBoostedBalanceUpdated)
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
func (it *YearnV3GaugeBoostedBalanceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeBoostedBalanceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeBoostedBalanceUpdated represents a BoostedBalanceUpdated event raised by the YearnV3Gauge contract.
type YearnV3GaugeBoostedBalanceUpdated struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBoostedBalanceUpdated is a free log retrieval operation binding the contract event 0x291ff844d30f85bb011aca3bfccedead238b6ed2e4b283504e3c2231d134524b.
//
// Solidity: event BoostedBalanceUpdated(address account, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterBoostedBalanceUpdated(opts *bind.FilterOpts) (*YearnV3GaugeBoostedBalanceUpdatedIterator, error) {

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "BoostedBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeBoostedBalanceUpdatedIterator{contract: _YearnV3Gauge.contract, event: "BoostedBalanceUpdated", logs: logs, sub: sub}, nil
}

// WatchBoostedBalanceUpdated is a free log subscription operation binding the contract event 0x291ff844d30f85bb011aca3bfccedead238b6ed2e4b283504e3c2231d134524b.
//
// Solidity: event BoostedBalanceUpdated(address account, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchBoostedBalanceUpdated(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeBoostedBalanceUpdated) (event.Subscription, error) {

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "BoostedBalanceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeBoostedBalanceUpdated)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "BoostedBalanceUpdated", log); err != nil {
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

// ParseBoostedBalanceUpdated is a log parse operation binding the contract event 0x291ff844d30f85bb011aca3bfccedead238b6ed2e4b283504e3c2231d134524b.
//
// Solidity: event BoostedBalanceUpdated(address account, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseBoostedBalanceUpdated(log types.Log) (*YearnV3GaugeBoostedBalanceUpdated, error) {
	event := new(YearnV3GaugeBoostedBalanceUpdated)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "BoostedBalanceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the YearnV3Gauge contract.
type YearnV3GaugeDepositIterator struct {
	Event *YearnV3GaugeDeposit // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeDeposit)
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
		it.Event = new(YearnV3GaugeDeposit)
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
func (it *YearnV3GaugeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeDeposit represents a Deposit event raised by the YearnV3Gauge contract.
type YearnV3GaugeDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*YearnV3GaugeDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeDepositIterator{contract: _YearnV3Gauge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeDeposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeDeposit)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseDeposit(log types.Log) (*YearnV3GaugeDeposit, error) {
	event := new(YearnV3GaugeDeposit)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeDurationUpdatedIterator is returned from FilterDurationUpdated and is used to iterate over the raw logs and unpacked data for DurationUpdated events raised by the YearnV3Gauge contract.
type YearnV3GaugeDurationUpdatedIterator struct {
	Event *YearnV3GaugeDurationUpdated // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeDurationUpdated)
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
		it.Event = new(YearnV3GaugeDurationUpdated)
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
func (it *YearnV3GaugeDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeDurationUpdated represents a DurationUpdated event raised by the YearnV3Gauge contract.
type YearnV3GaugeDurationUpdated struct {
	Duration     *big.Int
	RewardRate   *big.Int
	PeriodFinish *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDurationUpdated is a free log retrieval operation binding the contract event 0x82cec86ca93cde0f1fd0129402cc13d8f7b7f5154320025ab47357562ba02528.
//
// Solidity: event DurationUpdated(uint256 duration, uint256 rewardRate, uint256 periodFinish)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterDurationUpdated(opts *bind.FilterOpts) (*YearnV3GaugeDurationUpdatedIterator, error) {

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "DurationUpdated")
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeDurationUpdatedIterator{contract: _YearnV3Gauge.contract, event: "DurationUpdated", logs: logs, sub: sub}, nil
}

// WatchDurationUpdated is a free log subscription operation binding the contract event 0x82cec86ca93cde0f1fd0129402cc13d8f7b7f5154320025ab47357562ba02528.
//
// Solidity: event DurationUpdated(uint256 duration, uint256 rewardRate, uint256 periodFinish)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchDurationUpdated(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "DurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeDurationUpdated)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "DurationUpdated", log); err != nil {
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

// ParseDurationUpdated is a log parse operation binding the contract event 0x82cec86ca93cde0f1fd0129402cc13d8f7b7f5154320025ab47357562ba02528.
//
// Solidity: event DurationUpdated(uint256 duration, uint256 rewardRate, uint256 periodFinish)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseDurationUpdated(log types.Log) (*YearnV3GaugeDurationUpdated, error) {
	event := new(YearnV3GaugeDurationUpdated)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "DurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the YearnV3Gauge contract.
type YearnV3GaugeInitializeIterator struct {
	Event *YearnV3GaugeInitialize // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeInitialize)
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
		it.Event = new(YearnV3GaugeInitialize)
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
func (it *YearnV3GaugeInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeInitialize represents a Initialize event raised by the YearnV3Gauge contract.
type YearnV3GaugeInitialize struct {
	Asset common.Address
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0xdc90fed0326ba91706deeac7eb34ac9f8b680734f9d782864dc29704d23bed6a.
//
// Solidity: event Initialize(address indexed asset, address indexed owner)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterInitialize(opts *bind.FilterOpts, asset []common.Address, owner []common.Address) (*YearnV3GaugeInitializeIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "Initialize", assetRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeInitializeIterator{contract: _YearnV3Gauge.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0xdc90fed0326ba91706deeac7eb34ac9f8b680734f9d782864dc29704d23bed6a.
//
// Solidity: event Initialize(address indexed asset, address indexed owner)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeInitialize, asset []common.Address, owner []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "Initialize", assetRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeInitialize)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0xdc90fed0326ba91706deeac7eb34ac9f8b680734f9d782864dc29704d23bed6a.
//
// Solidity: event Initialize(address indexed asset, address indexed owner)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseInitialize(log types.Log) (*YearnV3GaugeInitialize, error) {
	event := new(YearnV3GaugeInitialize)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the YearnV3Gauge contract.
type YearnV3GaugeInitializedIterator struct {
	Event *YearnV3GaugeInitialized // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeInitialized)
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
		it.Event = new(YearnV3GaugeInitialized)
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
func (it *YearnV3GaugeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeInitialized represents a Initialized event raised by the YearnV3Gauge contract.
type YearnV3GaugeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterInitialized(opts *bind.FilterOpts) (*YearnV3GaugeInitializedIterator, error) {

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeInitializedIterator{contract: _YearnV3Gauge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeInitialized) (event.Subscription, error) {

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeInitialized)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseInitialized(log types.Log) (*YearnV3GaugeInitialized, error) {
	event := new(YearnV3GaugeInitialized)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the YearnV3Gauge contract.
type YearnV3GaugeOwnershipTransferredIterator struct {
	Event *YearnV3GaugeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeOwnershipTransferred)
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
		it.Event = new(YearnV3GaugeOwnershipTransferred)
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
func (it *YearnV3GaugeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeOwnershipTransferred represents a OwnershipTransferred event raised by the YearnV3Gauge contract.
type YearnV3GaugeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*YearnV3GaugeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeOwnershipTransferredIterator{contract: _YearnV3Gauge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeOwnershipTransferred)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseOwnershipTransferred(log types.Log) (*YearnV3GaugeOwnershipTransferred, error) {
	event := new(YearnV3GaugeOwnershipTransferred)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeRecipientUpdatedIterator is returned from FilterRecipientUpdated and is used to iterate over the raw logs and unpacked data for RecipientUpdated events raised by the YearnV3Gauge contract.
type YearnV3GaugeRecipientUpdatedIterator struct {
	Event *YearnV3GaugeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeRecipientUpdated)
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
		it.Event = new(YearnV3GaugeRecipientUpdated)
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
func (it *YearnV3GaugeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeRecipientUpdated represents a RecipientUpdated event raised by the YearnV3Gauge contract.
type YearnV3GaugeRecipientUpdated struct {
	Account   common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRecipientUpdated is a free log retrieval operation binding the contract event 0x62e69886a5df0ba8ffcacbfc1388754e7abd9bde24b036354c561f1acd4e4593.
//
// Solidity: event RecipientUpdated(address indexed account, address indexed recipient)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterRecipientUpdated(opts *bind.FilterOpts, account []common.Address, recipient []common.Address) (*YearnV3GaugeRecipientUpdatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "RecipientUpdated", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeRecipientUpdatedIterator{contract: _YearnV3Gauge.contract, event: "RecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchRecipientUpdated is a free log subscription operation binding the contract event 0x62e69886a5df0ba8ffcacbfc1388754e7abd9bde24b036354c561f1acd4e4593.
//
// Solidity: event RecipientUpdated(address indexed account, address indexed recipient)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchRecipientUpdated(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeRecipientUpdated, account []common.Address, recipient []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "RecipientUpdated", accountRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeRecipientUpdated)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "RecipientUpdated", log); err != nil {
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

// ParseRecipientUpdated is a log parse operation binding the contract event 0x62e69886a5df0ba8ffcacbfc1388754e7abd9bde24b036354c561f1acd4e4593.
//
// Solidity: event RecipientUpdated(address indexed account, address indexed recipient)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseRecipientUpdated(log types.Log) (*YearnV3GaugeRecipientUpdated, error) {
	event := new(YearnV3GaugeRecipientUpdated)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "RecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the YearnV3Gauge contract.
type YearnV3GaugeRewardPaidIterator struct {
	Event *YearnV3GaugeRewardPaid // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeRewardPaid)
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
		it.Event = new(YearnV3GaugeRewardPaid)
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
func (it *YearnV3GaugeRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeRewardPaid represents a RewardPaid event raised by the YearnV3Gauge contract.
type YearnV3GaugeRewardPaid struct {
	User   common.Address
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address) (*YearnV3GaugeRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeRewardPaidIterator{contract: _YearnV3Gauge.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeRewardPaid, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "RewardPaid", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeRewardPaid)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0xe2403640ba68fed3a2f88b7557551d1993f84b99bb10ff833f0cf8db0c5e0486.
//
// Solidity: event RewardPaid(address indexed user, uint256 reward)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseRewardPaid(log types.Log) (*YearnV3GaugeRewardPaid, error) {
	event := new(YearnV3GaugeRewardPaid)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeRewardsAddedIterator is returned from FilterRewardsAdded and is used to iterate over the raw logs and unpacked data for RewardsAdded events raised by the YearnV3Gauge contract.
type YearnV3GaugeRewardsAddedIterator struct {
	Event *YearnV3GaugeRewardsAdded // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeRewardsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeRewardsAdded)
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
		it.Event = new(YearnV3GaugeRewardsAdded)
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
func (it *YearnV3GaugeRewardsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeRewardsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeRewardsAdded represents a RewardsAdded event raised by the YearnV3Gauge contract.
type YearnV3GaugeRewardsAdded struct {
	CurrentRewards    *big.Int
	LastUpdateTime    *big.Int
	PeriodFinish      *big.Int
	RewardRate        *big.Int
	HistoricalRewards *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRewardsAdded is a free log retrieval operation binding the contract event 0x944ffd3678415a15cbfef07dd7d9f20cdc6f36d12588a4ba7e8eb440f32c61be.
//
// Solidity: event RewardsAdded(uint256 currentRewards, uint256 lastUpdateTime, uint256 periodFinish, uint256 rewardRate, uint256 historicalRewards)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterRewardsAdded(opts *bind.FilterOpts) (*YearnV3GaugeRewardsAddedIterator, error) {

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "RewardsAdded")
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeRewardsAddedIterator{contract: _YearnV3Gauge.contract, event: "RewardsAdded", logs: logs, sub: sub}, nil
}

// WatchRewardsAdded is a free log subscription operation binding the contract event 0x944ffd3678415a15cbfef07dd7d9f20cdc6f36d12588a4ba7e8eb440f32c61be.
//
// Solidity: event RewardsAdded(uint256 currentRewards, uint256 lastUpdateTime, uint256 periodFinish, uint256 rewardRate, uint256 historicalRewards)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchRewardsAdded(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeRewardsAdded) (event.Subscription, error) {

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "RewardsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeRewardsAdded)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "RewardsAdded", log); err != nil {
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

// ParseRewardsAdded is a log parse operation binding the contract event 0x944ffd3678415a15cbfef07dd7d9f20cdc6f36d12588a4ba7e8eb440f32c61be.
//
// Solidity: event RewardsAdded(uint256 currentRewards, uint256 lastUpdateTime, uint256 periodFinish, uint256 rewardRate, uint256 historicalRewards)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseRewardsAdded(log types.Log) (*YearnV3GaugeRewardsAdded, error) {
	event := new(YearnV3GaugeRewardsAdded)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "RewardsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeRewardsQueuedIterator is returned from FilterRewardsQueued and is used to iterate over the raw logs and unpacked data for RewardsQueued events raised by the YearnV3Gauge contract.
type YearnV3GaugeRewardsQueuedIterator struct {
	Event *YearnV3GaugeRewardsQueued // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeRewardsQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeRewardsQueued)
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
		it.Event = new(YearnV3GaugeRewardsQueued)
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
func (it *YearnV3GaugeRewardsQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeRewardsQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeRewardsQueued represents a RewardsQueued event raised by the YearnV3Gauge contract.
type YearnV3GaugeRewardsQueued struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsQueued is a free log retrieval operation binding the contract event 0x1c88aa9a39b1a6357a85c97a3bd4e2b0738e74c68b92928276bc85f495b2450b.
//
// Solidity: event RewardsQueued(address indexed from, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterRewardsQueued(opts *bind.FilterOpts, from []common.Address) (*YearnV3GaugeRewardsQueuedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "RewardsQueued", fromRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeRewardsQueuedIterator{contract: _YearnV3Gauge.contract, event: "RewardsQueued", logs: logs, sub: sub}, nil
}

// WatchRewardsQueued is a free log subscription operation binding the contract event 0x1c88aa9a39b1a6357a85c97a3bd4e2b0738e74c68b92928276bc85f495b2450b.
//
// Solidity: event RewardsQueued(address indexed from, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchRewardsQueued(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeRewardsQueued, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "RewardsQueued", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeRewardsQueued)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "RewardsQueued", log); err != nil {
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

// ParseRewardsQueued is a log parse operation binding the contract event 0x1c88aa9a39b1a6357a85c97a3bd4e2b0738e74c68b92928276bc85f495b2450b.
//
// Solidity: event RewardsQueued(address indexed from, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseRewardsQueued(log types.Log) (*YearnV3GaugeRewardsQueued, error) {
	event := new(YearnV3GaugeRewardsQueued)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "RewardsQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeSweepIterator is returned from FilterSweep and is used to iterate over the raw logs and unpacked data for Sweep events raised by the YearnV3Gauge contract.
type YearnV3GaugeSweepIterator struct {
	Event *YearnV3GaugeSweep // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeSweepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeSweep)
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
		it.Event = new(YearnV3GaugeSweep)
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
func (it *YearnV3GaugeSweepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeSweepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeSweep represents a Sweep event raised by the YearnV3Gauge contract.
type YearnV3GaugeSweep struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSweep is a free log retrieval operation binding the contract event 0xab2246061d7b0dd3631d037e3f6da75782ae489eeb9f6af878a4b25df9b07c77.
//
// Solidity: event Sweep(address indexed token, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterSweep(opts *bind.FilterOpts, token []common.Address) (*YearnV3GaugeSweepIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "Sweep", tokenRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeSweepIterator{contract: _YearnV3Gauge.contract, event: "Sweep", logs: logs, sub: sub}, nil
}

// WatchSweep is a free log subscription operation binding the contract event 0xab2246061d7b0dd3631d037e3f6da75782ae489eeb9f6af878a4b25df9b07c77.
//
// Solidity: event Sweep(address indexed token, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchSweep(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeSweep, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "Sweep", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeSweep)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "Sweep", log); err != nil {
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

// ParseSweep is a log parse operation binding the contract event 0xab2246061d7b0dd3631d037e3f6da75782ae489eeb9f6af878a4b25df9b07c77.
//
// Solidity: event Sweep(address indexed token, uint256 amount)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseSweep(log types.Log) (*YearnV3GaugeSweep, error) {
	event := new(YearnV3GaugeSweep)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "Sweep", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YearnV3Gauge contract.
type YearnV3GaugeTransferIterator struct {
	Event *YearnV3GaugeTransfer // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeTransfer)
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
		it.Event = new(YearnV3GaugeTransfer)
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
func (it *YearnV3GaugeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeTransfer represents a Transfer event raised by the YearnV3Gauge contract.
type YearnV3GaugeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YearnV3GaugeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeTransferIterator{contract: _YearnV3Gauge.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeTransfer)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseTransfer(log types.Log) (*YearnV3GaugeTransfer, error) {
	event := new(YearnV3GaugeTransfer)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeTransferredPenaltyIterator is returned from FilterTransferredPenalty and is used to iterate over the raw logs and unpacked data for TransferredPenalty events raised by the YearnV3Gauge contract.
type YearnV3GaugeTransferredPenaltyIterator struct {
	Event *YearnV3GaugeTransferredPenalty // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeTransferredPenaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeTransferredPenalty)
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
		it.Event = new(YearnV3GaugeTransferredPenalty)
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
func (it *YearnV3GaugeTransferredPenaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeTransferredPenaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeTransferredPenalty represents a TransferredPenalty event raised by the YearnV3Gauge contract.
type YearnV3GaugeTransferredPenalty struct {
	Account    common.Address
	Transfered *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferredPenalty is a free log retrieval operation binding the contract event 0xfdcc759119f4a689ba608afdccb078153573a5a615700713ebb84704609694cc.
//
// Solidity: event TransferredPenalty(address indexed account, uint256 transfered)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterTransferredPenalty(opts *bind.FilterOpts, account []common.Address) (*YearnV3GaugeTransferredPenaltyIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "TransferredPenalty", accountRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeTransferredPenaltyIterator{contract: _YearnV3Gauge.contract, event: "TransferredPenalty", logs: logs, sub: sub}, nil
}

// WatchTransferredPenalty is a free log subscription operation binding the contract event 0xfdcc759119f4a689ba608afdccb078153573a5a615700713ebb84704609694cc.
//
// Solidity: event TransferredPenalty(address indexed account, uint256 transfered)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchTransferredPenalty(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeTransferredPenalty, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "TransferredPenalty", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeTransferredPenalty)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "TransferredPenalty", log); err != nil {
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

// ParseTransferredPenalty is a log parse operation binding the contract event 0xfdcc759119f4a689ba608afdccb078153573a5a615700713ebb84704609694cc.
//
// Solidity: event TransferredPenalty(address indexed account, uint256 transfered)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseTransferredPenalty(log types.Log) (*YearnV3GaugeTransferredPenalty, error) {
	event := new(YearnV3GaugeTransferredPenalty)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "TransferredPenalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeUpdatedRewardsIterator is returned from FilterUpdatedRewards and is used to iterate over the raw logs and unpacked data for UpdatedRewards events raised by the YearnV3Gauge contract.
type YearnV3GaugeUpdatedRewardsIterator struct {
	Event *YearnV3GaugeUpdatedRewards // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeUpdatedRewardsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeUpdatedRewards)
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
		it.Event = new(YearnV3GaugeUpdatedRewards)
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
func (it *YearnV3GaugeUpdatedRewardsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeUpdatedRewardsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeUpdatedRewards represents a UpdatedRewards event raised by the YearnV3Gauge contract.
type YearnV3GaugeUpdatedRewards struct {
	Account                common.Address
	RewardPerTokenStored   *big.Int
	LastUpdateTime         *big.Int
	Rewards                *big.Int
	UserRewardPerTokenPaid *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterUpdatedRewards is a free log retrieval operation binding the contract event 0xfbe590c835e1c07f8e971c36021d1be46f43f7b0b6dc5413dbd5753590569d58.
//
// Solidity: event UpdatedRewards(address indexed account, uint256 rewardPerTokenStored, uint256 lastUpdateTime, uint256 rewards, uint256 userRewardPerTokenPaid)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterUpdatedRewards(opts *bind.FilterOpts, account []common.Address) (*YearnV3GaugeUpdatedRewardsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "UpdatedRewards", accountRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeUpdatedRewardsIterator{contract: _YearnV3Gauge.contract, event: "UpdatedRewards", logs: logs, sub: sub}, nil
}

// WatchUpdatedRewards is a free log subscription operation binding the contract event 0xfbe590c835e1c07f8e971c36021d1be46f43f7b0b6dc5413dbd5753590569d58.
//
// Solidity: event UpdatedRewards(address indexed account, uint256 rewardPerTokenStored, uint256 lastUpdateTime, uint256 rewards, uint256 userRewardPerTokenPaid)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchUpdatedRewards(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeUpdatedRewards, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "UpdatedRewards", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeUpdatedRewards)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
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

// ParseUpdatedRewards is a log parse operation binding the contract event 0xfbe590c835e1c07f8e971c36021d1be46f43f7b0b6dc5413dbd5753590569d58.
//
// Solidity: event UpdatedRewards(address indexed account, uint256 rewardPerTokenStored, uint256 lastUpdateTime, uint256 rewards, uint256 userRewardPerTokenPaid)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseUpdatedRewards(log types.Log) (*YearnV3GaugeUpdatedRewards, error) {
	event := new(YearnV3GaugeUpdatedRewards)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "UpdatedRewards", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3GaugeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the YearnV3Gauge contract.
type YearnV3GaugeWithdrawIterator struct {
	Event *YearnV3GaugeWithdraw // Event containing the contract specifics and raw log

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
func (it *YearnV3GaugeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3GaugeWithdraw)
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
		it.Event = new(YearnV3GaugeWithdraw)
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
func (it *YearnV3GaugeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3GaugeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3GaugeWithdraw represents a Withdraw event raised by the YearnV3Gauge contract.
type YearnV3GaugeWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Gauge *YearnV3GaugeFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*YearnV3GaugeWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3GaugeWithdrawIterator{contract: _YearnV3Gauge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Gauge *YearnV3GaugeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *YearnV3GaugeWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _YearnV3Gauge.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3GaugeWithdraw)
				if err := _YearnV3Gauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_YearnV3Gauge *YearnV3GaugeFilterer) ParseWithdraw(log types.Log) (*YearnV3GaugeWithdraw, error) {
	event := new(YearnV3GaugeWithdraw)
	if err := _YearnV3Gauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
