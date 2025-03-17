// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basepaint_referral

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

// BasepaintReferralMetaData contains all meta data concerning the BasepaintReferral contract.
var BasepaintReferralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIBasePaint\",\"name\":\"_basepaint\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughContractFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basepaint\",\"outputs\":[{\"internalType\":\"contractIBasePaint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cashOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"cashOutBatched\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sendMintsTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sendRewardsTo\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sendMintsTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sendRewardsTo\",\"type\":\"address\"}],\"name\":\"mintLatest\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bips\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bips\",\"type\":\"uint256\"}],\"name\":\"setRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// BasepaintReferralABI is the input ABI used to generate the binding from.
// Deprecated: Use BasepaintReferralMetaData.ABI instead.
var BasepaintReferralABI = BasepaintReferralMetaData.ABI

// BasepaintReferral is an auto generated Go binding around an Ethereum contract.
type BasepaintReferral struct {
	BasepaintReferralCaller     // Read-only binding to the contract
	BasepaintReferralTransactor // Write-only binding to the contract
	BasepaintReferralFilterer   // Log filterer for contract events
}

// BasepaintReferralCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasepaintReferralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasepaintReferralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasepaintReferralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasepaintReferralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasepaintReferralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasepaintReferralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasepaintReferralSession struct {
	Contract     *BasepaintReferral // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BasepaintReferralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasepaintReferralCallerSession struct {
	Contract *BasepaintReferralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BasepaintReferralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasepaintReferralTransactorSession struct {
	Contract     *BasepaintReferralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BasepaintReferralRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasepaintReferralRaw struct {
	Contract *BasepaintReferral // Generic contract binding to access the raw methods on
}

// BasepaintReferralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasepaintReferralCallerRaw struct {
	Contract *BasepaintReferralCaller // Generic read-only contract binding to access the raw methods on
}

// BasepaintReferralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasepaintReferralTransactorRaw struct {
	Contract *BasepaintReferralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasepaintReferral creates a new instance of BasepaintReferral, bound to a specific deployed contract.
func NewBasepaintReferral(address common.Address, backend bind.ContractBackend) (*BasepaintReferral, error) {
	contract, err := bindBasepaintReferral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasepaintReferral{BasepaintReferralCaller: BasepaintReferralCaller{contract: contract}, BasepaintReferralTransactor: BasepaintReferralTransactor{contract: contract}, BasepaintReferralFilterer: BasepaintReferralFilterer{contract: contract}}, nil
}

// NewBasepaintReferralCaller creates a new read-only instance of BasepaintReferral, bound to a specific deployed contract.
func NewBasepaintReferralCaller(address common.Address, caller bind.ContractCaller) (*BasepaintReferralCaller, error) {
	contract, err := bindBasepaintReferral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasepaintReferralCaller{contract: contract}, nil
}

// NewBasepaintReferralTransactor creates a new write-only instance of BasepaintReferral, bound to a specific deployed contract.
func NewBasepaintReferralTransactor(address common.Address, transactor bind.ContractTransactor) (*BasepaintReferralTransactor, error) {
	contract, err := bindBasepaintReferral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasepaintReferralTransactor{contract: contract}, nil
}

// NewBasepaintReferralFilterer creates a new log filterer instance of BasepaintReferral, bound to a specific deployed contract.
func NewBasepaintReferralFilterer(address common.Address, filterer bind.ContractFilterer) (*BasepaintReferralFilterer, error) {
	contract, err := bindBasepaintReferral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasepaintReferralFilterer{contract: contract}, nil
}

// bindBasepaintReferral binds a generic wrapper to an already deployed contract.
func bindBasepaintReferral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasepaintReferralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasepaintReferral *BasepaintReferralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasepaintReferral.Contract.BasepaintReferralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasepaintReferral *BasepaintReferralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.BasepaintReferralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasepaintReferral *BasepaintReferralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.BasepaintReferralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasepaintReferral *BasepaintReferralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasepaintReferral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasepaintReferral *BasepaintReferralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasepaintReferral *BasepaintReferralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BasepaintReferral *BasepaintReferralCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BasepaintReferral *BasepaintReferralSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BasepaintReferral.Contract.Allowance(&_BasepaintReferral.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BasepaintReferral *BasepaintReferralCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BasepaintReferral.Contract.Allowance(&_BasepaintReferral.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BasepaintReferral *BasepaintReferralCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BasepaintReferral *BasepaintReferralSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BasepaintReferral.Contract.BalanceOf(&_BasepaintReferral.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BasepaintReferral *BasepaintReferralCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BasepaintReferral.Contract.BalanceOf(&_BasepaintReferral.CallOpts, account)
}

// Basepaint is a free data retrieval call binding the contract method 0x44a308c3.
//
// Solidity: function basepaint() view returns(address)
func (_BasepaintReferral *BasepaintReferralCaller) Basepaint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "basepaint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Basepaint is a free data retrieval call binding the contract method 0x44a308c3.
//
// Solidity: function basepaint() view returns(address)
func (_BasepaintReferral *BasepaintReferralSession) Basepaint() (common.Address, error) {
	return _BasepaintReferral.Contract.Basepaint(&_BasepaintReferral.CallOpts)
}

// Basepaint is a free data retrieval call binding the contract method 0x44a308c3.
//
// Solidity: function basepaint() view returns(address)
func (_BasepaintReferral *BasepaintReferralCallerSession) Basepaint() (common.Address, error) {
	return _BasepaintReferral.Contract.Basepaint(&_BasepaintReferral.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BasepaintReferral *BasepaintReferralCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BasepaintReferral *BasepaintReferralSession) Decimals() (uint8, error) {
	return _BasepaintReferral.Contract.Decimals(&_BasepaintReferral.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BasepaintReferral *BasepaintReferralCallerSession) Decimals() (uint8, error) {
	return _BasepaintReferral.Contract.Decimals(&_BasepaintReferral.CallOpts)
}

// DefaultRewardRate is a free data retrieval call binding the contract method 0xc42f8e6e.
//
// Solidity: function defaultRewardRate() view returns(uint256)
func (_BasepaintReferral *BasepaintReferralCaller) DefaultRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "defaultRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultRewardRate is a free data retrieval call binding the contract method 0xc42f8e6e.
//
// Solidity: function defaultRewardRate() view returns(uint256)
func (_BasepaintReferral *BasepaintReferralSession) DefaultRewardRate() (*big.Int, error) {
	return _BasepaintReferral.Contract.DefaultRewardRate(&_BasepaintReferral.CallOpts)
}

// DefaultRewardRate is a free data retrieval call binding the contract method 0xc42f8e6e.
//
// Solidity: function defaultRewardRate() view returns(uint256)
func (_BasepaintReferral *BasepaintReferralCallerSession) DefaultRewardRate() (*big.Int, error) {
	return _BasepaintReferral.Contract.DefaultRewardRate(&_BasepaintReferral.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasepaintReferral *BasepaintReferralCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasepaintReferral *BasepaintReferralSession) Name() (string, error) {
	return _BasepaintReferral.Contract.Name(&_BasepaintReferral.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasepaintReferral *BasepaintReferralCallerSession) Name() (string, error) {
	return _BasepaintReferral.Contract.Name(&_BasepaintReferral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BasepaintReferral *BasepaintReferralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BasepaintReferral *BasepaintReferralSession) Owner() (common.Address, error) {
	return _BasepaintReferral.Contract.Owner(&_BasepaintReferral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BasepaintReferral *BasepaintReferralCallerSession) Owner() (common.Address, error) {
	return _BasepaintReferral.Contract.Owner(&_BasepaintReferral.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x221ca18c.
//
// Solidity: function rewardRate(address referrer) view returns(uint256 bips)
func (_BasepaintReferral *BasepaintReferralCaller) RewardRate(opts *bind.CallOpts, referrer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "rewardRate", referrer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x221ca18c.
//
// Solidity: function rewardRate(address referrer) view returns(uint256 bips)
func (_BasepaintReferral *BasepaintReferralSession) RewardRate(referrer common.Address) (*big.Int, error) {
	return _BasepaintReferral.Contract.RewardRate(&_BasepaintReferral.CallOpts, referrer)
}

// RewardRate is a free data retrieval call binding the contract method 0x221ca18c.
//
// Solidity: function rewardRate(address referrer) view returns(uint256 bips)
func (_BasepaintReferral *BasepaintReferralCallerSession) RewardRate(referrer common.Address) (*big.Int, error) {
	return _BasepaintReferral.Contract.RewardRate(&_BasepaintReferral.CallOpts, referrer)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BasepaintReferral *BasepaintReferralCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BasepaintReferral *BasepaintReferralSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BasepaintReferral.Contract.SupportsInterface(&_BasepaintReferral.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BasepaintReferral *BasepaintReferralCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BasepaintReferral.Contract.SupportsInterface(&_BasepaintReferral.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasepaintReferral *BasepaintReferralCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasepaintReferral *BasepaintReferralSession) Symbol() (string, error) {
	return _BasepaintReferral.Contract.Symbol(&_BasepaintReferral.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasepaintReferral *BasepaintReferralCallerSession) Symbol() (string, error) {
	return _BasepaintReferral.Contract.Symbol(&_BasepaintReferral.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasepaintReferral *BasepaintReferralCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasepaintReferral.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasepaintReferral *BasepaintReferralSession) TotalSupply() (*big.Int, error) {
	return _BasepaintReferral.Contract.TotalSupply(&_BasepaintReferral.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasepaintReferral *BasepaintReferralCallerSession) TotalSupply() (*big.Int, error) {
	return _BasepaintReferral.Contract.TotalSupply(&_BasepaintReferral.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Approve(&_BasepaintReferral.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Approve(&_BasepaintReferral.TransactOpts, spender, value)
}

// CashOut is a paid mutator transaction binding the contract method 0xd8a71364.
//
// Solidity: function cashOut(address account) returns()
func (_BasepaintReferral *BasepaintReferralTransactor) CashOut(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "cashOut", account)
}

// CashOut is a paid mutator transaction binding the contract method 0xd8a71364.
//
// Solidity: function cashOut(address account) returns()
func (_BasepaintReferral *BasepaintReferralSession) CashOut(account common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.CashOut(&_BasepaintReferral.TransactOpts, account)
}

// CashOut is a paid mutator transaction binding the contract method 0xd8a71364.
//
// Solidity: function cashOut(address account) returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) CashOut(account common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.CashOut(&_BasepaintReferral.TransactOpts, account)
}

// CashOutBatched is a paid mutator transaction binding the contract method 0x96e10c8e.
//
// Solidity: function cashOutBatched(address[] accounts) returns()
func (_BasepaintReferral *BasepaintReferralTransactor) CashOutBatched(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "cashOutBatched", accounts)
}

// CashOutBatched is a paid mutator transaction binding the contract method 0x96e10c8e.
//
// Solidity: function cashOutBatched(address[] accounts) returns()
func (_BasepaintReferral *BasepaintReferralSession) CashOutBatched(accounts []common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.CashOutBatched(&_BasepaintReferral.TransactOpts, accounts)
}

// CashOutBatched is a paid mutator transaction binding the contract method 0x96e10c8e.
//
// Solidity: function cashOutBatched(address[] accounts) returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) CashOutBatched(accounts []common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.CashOutBatched(&_BasepaintReferral.TransactOpts, accounts)
}

// Mint is a paid mutator transaction binding the contract method 0xe9eb7008.
//
// Solidity: function mint(uint256 tokenId, address sendMintsTo, uint256 count, address sendRewardsTo) payable returns()
func (_BasepaintReferral *BasepaintReferralTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, sendMintsTo common.Address, count *big.Int, sendRewardsTo common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "mint", tokenId, sendMintsTo, count, sendRewardsTo)
}

// Mint is a paid mutator transaction binding the contract method 0xe9eb7008.
//
// Solidity: function mint(uint256 tokenId, address sendMintsTo, uint256 count, address sendRewardsTo) payable returns()
func (_BasepaintReferral *BasepaintReferralSession) Mint(tokenId *big.Int, sendMintsTo common.Address, count *big.Int, sendRewardsTo common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Mint(&_BasepaintReferral.TransactOpts, tokenId, sendMintsTo, count, sendRewardsTo)
}

// Mint is a paid mutator transaction binding the contract method 0xe9eb7008.
//
// Solidity: function mint(uint256 tokenId, address sendMintsTo, uint256 count, address sendRewardsTo) payable returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) Mint(tokenId *big.Int, sendMintsTo common.Address, count *big.Int, sendRewardsTo common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Mint(&_BasepaintReferral.TransactOpts, tokenId, sendMintsTo, count, sendRewardsTo)
}

// MintLatest is a paid mutator transaction binding the contract method 0xb9e253bf.
//
// Solidity: function mintLatest(address sendMintsTo, uint256 count, address sendRewardsTo) payable returns()
func (_BasepaintReferral *BasepaintReferralTransactor) MintLatest(opts *bind.TransactOpts, sendMintsTo common.Address, count *big.Int, sendRewardsTo common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "mintLatest", sendMintsTo, count, sendRewardsTo)
}

// MintLatest is a paid mutator transaction binding the contract method 0xb9e253bf.
//
// Solidity: function mintLatest(address sendMintsTo, uint256 count, address sendRewardsTo) payable returns()
func (_BasepaintReferral *BasepaintReferralSession) MintLatest(sendMintsTo common.Address, count *big.Int, sendRewardsTo common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.MintLatest(&_BasepaintReferral.TransactOpts, sendMintsTo, count, sendRewardsTo)
}

// MintLatest is a paid mutator transaction binding the contract method 0xb9e253bf.
//
// Solidity: function mintLatest(address sendMintsTo, uint256 count, address sendRewardsTo) payable returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) MintLatest(sendMintsTo common.Address, count *big.Int, sendRewardsTo common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.MintLatest(&_BasepaintReferral.TransactOpts, sendMintsTo, count, sendRewardsTo)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BasepaintReferral *BasepaintReferralTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BasepaintReferral *BasepaintReferralSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.OnERC1155BatchReceived(&_BasepaintReferral.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_BasepaintReferral *BasepaintReferralTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.OnERC1155BatchReceived(&_BasepaintReferral.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BasepaintReferral *BasepaintReferralTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BasepaintReferral *BasepaintReferralSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.OnERC1155Received(&_BasepaintReferral.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_BasepaintReferral *BasepaintReferralTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.OnERC1155Received(&_BasepaintReferral.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BasepaintReferral *BasepaintReferralTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BasepaintReferral *BasepaintReferralSession) RenounceOwnership() (*types.Transaction, error) {
	return _BasepaintReferral.Contract.RenounceOwnership(&_BasepaintReferral.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BasepaintReferral.Contract.RenounceOwnership(&_BasepaintReferral.TransactOpts)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x812a71d5.
//
// Solidity: function setRewardRate(address referrer, uint256 bips) returns()
func (_BasepaintReferral *BasepaintReferralTransactor) SetRewardRate(opts *bind.TransactOpts, referrer common.Address, bips *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "setRewardRate", referrer, bips)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x812a71d5.
//
// Solidity: function setRewardRate(address referrer, uint256 bips) returns()
func (_BasepaintReferral *BasepaintReferralSession) SetRewardRate(referrer common.Address, bips *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.SetRewardRate(&_BasepaintReferral.TransactOpts, referrer, bips)
}

// SetRewardRate is a paid mutator transaction binding the contract method 0x812a71d5.
//
// Solidity: function setRewardRate(address referrer, uint256 bips) returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) SetRewardRate(referrer common.Address, bips *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.SetRewardRate(&_BasepaintReferral.TransactOpts, referrer, bips)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Transfer(&_BasepaintReferral.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Transfer(&_BasepaintReferral.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.TransferFrom(&_BasepaintReferral.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_BasepaintReferral *BasepaintReferralTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.TransferFrom(&_BasepaintReferral.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BasepaintReferral *BasepaintReferralTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BasepaintReferral *BasepaintReferralSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.TransferOwnership(&_BasepaintReferral.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.TransferOwnership(&_BasepaintReferral.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_BasepaintReferral *BasepaintReferralTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_BasepaintReferral *BasepaintReferralSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Withdraw(&_BasepaintReferral.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 value) returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Withdraw(&_BasepaintReferral.TransactOpts, value)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasepaintReferral *BasepaintReferralTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasepaintReferral.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasepaintReferral *BasepaintReferralSession) Receive() (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Receive(&_BasepaintReferral.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_BasepaintReferral *BasepaintReferralTransactorSession) Receive() (*types.Transaction, error) {
	return _BasepaintReferral.Contract.Receive(&_BasepaintReferral.TransactOpts)
}

// BasepaintReferralApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BasepaintReferral contract.
type BasepaintReferralApprovalIterator struct {
	Event *BasepaintReferralApproval // Event containing the contract specifics and raw log

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
func (it *BasepaintReferralApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasepaintReferralApproval)
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
		it.Event = new(BasepaintReferralApproval)
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
func (it *BasepaintReferralApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasepaintReferralApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasepaintReferralApproval represents a Approval event raised by the BasepaintReferral contract.
type BasepaintReferralApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BasepaintReferral *BasepaintReferralFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BasepaintReferralApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BasepaintReferral.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BasepaintReferralApprovalIterator{contract: _BasepaintReferral.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BasepaintReferral *BasepaintReferralFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BasepaintReferralApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BasepaintReferral.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasepaintReferralApproval)
				if err := _BasepaintReferral.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BasepaintReferral *BasepaintReferralFilterer) ParseApproval(log types.Log) (*BasepaintReferralApproval, error) {
	event := new(BasepaintReferralApproval)
	if err := _BasepaintReferral.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasepaintReferralOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BasepaintReferral contract.
type BasepaintReferralOwnershipTransferredIterator struct {
	Event *BasepaintReferralOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BasepaintReferralOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasepaintReferralOwnershipTransferred)
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
		it.Event = new(BasepaintReferralOwnershipTransferred)
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
func (it *BasepaintReferralOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasepaintReferralOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasepaintReferralOwnershipTransferred represents a OwnershipTransferred event raised by the BasepaintReferral contract.
type BasepaintReferralOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BasepaintReferral *BasepaintReferralFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BasepaintReferralOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BasepaintReferral.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BasepaintReferralOwnershipTransferredIterator{contract: _BasepaintReferral.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BasepaintReferral *BasepaintReferralFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BasepaintReferralOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BasepaintReferral.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasepaintReferralOwnershipTransferred)
				if err := _BasepaintReferral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BasepaintReferral *BasepaintReferralFilterer) ParseOwnershipTransferred(log types.Log) (*BasepaintReferralOwnershipTransferred, error) {
	event := new(BasepaintReferralOwnershipTransferred)
	if err := _BasepaintReferral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasepaintReferralToppedUpIterator is returned from FilterToppedUp and is used to iterate over the raw logs and unpacked data for ToppedUp events raised by the BasepaintReferral contract.
type BasepaintReferralToppedUpIterator struct {
	Event *BasepaintReferralToppedUp // Event containing the contract specifics and raw log

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
func (it *BasepaintReferralToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasepaintReferralToppedUp)
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
		it.Event = new(BasepaintReferralToppedUp)
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
func (it *BasepaintReferralToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasepaintReferralToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasepaintReferralToppedUp represents a ToppedUp event raised by the BasepaintReferral contract.
type BasepaintReferralToppedUp struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterToppedUp is a free log retrieval operation binding the contract event 0x60c6b9906c64e9e2aa5fc860e22f00a1cb2c959162c4a26eb3b36c9d50829306.
//
// Solidity: event ToppedUp(uint256 amount)
func (_BasepaintReferral *BasepaintReferralFilterer) FilterToppedUp(opts *bind.FilterOpts) (*BasepaintReferralToppedUpIterator, error) {

	logs, sub, err := _BasepaintReferral.contract.FilterLogs(opts, "ToppedUp")
	if err != nil {
		return nil, err
	}
	return &BasepaintReferralToppedUpIterator{contract: _BasepaintReferral.contract, event: "ToppedUp", logs: logs, sub: sub}, nil
}

// WatchToppedUp is a free log subscription operation binding the contract event 0x60c6b9906c64e9e2aa5fc860e22f00a1cb2c959162c4a26eb3b36c9d50829306.
//
// Solidity: event ToppedUp(uint256 amount)
func (_BasepaintReferral *BasepaintReferralFilterer) WatchToppedUp(opts *bind.WatchOpts, sink chan<- *BasepaintReferralToppedUp) (event.Subscription, error) {

	logs, sub, err := _BasepaintReferral.contract.WatchLogs(opts, "ToppedUp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasepaintReferralToppedUp)
				if err := _BasepaintReferral.contract.UnpackLog(event, "ToppedUp", log); err != nil {
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

// ParseToppedUp is a log parse operation binding the contract event 0x60c6b9906c64e9e2aa5fc860e22f00a1cb2c959162c4a26eb3b36c9d50829306.
//
// Solidity: event ToppedUp(uint256 amount)
func (_BasepaintReferral *BasepaintReferralFilterer) ParseToppedUp(log types.Log) (*BasepaintReferralToppedUp, error) {
	event := new(BasepaintReferralToppedUp)
	if err := _BasepaintReferral.contract.UnpackLog(event, "ToppedUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasepaintReferralTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasepaintReferral contract.
type BasepaintReferralTransferIterator struct {
	Event *BasepaintReferralTransfer // Event containing the contract specifics and raw log

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
func (it *BasepaintReferralTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasepaintReferralTransfer)
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
		it.Event = new(BasepaintReferralTransfer)
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
func (it *BasepaintReferralTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasepaintReferralTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasepaintReferralTransfer represents a Transfer event raised by the BasepaintReferral contract.
type BasepaintReferralTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasepaintReferral *BasepaintReferralFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasepaintReferralTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasepaintReferral.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasepaintReferralTransferIterator{contract: _BasepaintReferral.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasepaintReferral *BasepaintReferralFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasepaintReferralTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasepaintReferral.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasepaintReferralTransfer)
				if err := _BasepaintReferral.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BasepaintReferral *BasepaintReferralFilterer) ParseTransfer(log types.Log) (*BasepaintReferralTransfer, error) {
	event := new(BasepaintReferralTransfer)
	if err := _BasepaintReferral.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
