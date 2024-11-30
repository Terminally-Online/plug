// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearn_v3_router

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

// YearnV3RouterMetaData contains all meta data concerning the YearnV3Router contract.
var YearnV3RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name_\",\"type\":\"string\"},{\"internalType\":\"contractIWETH9\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WETH9\",\"outputs\":[{\"internalType\":\"contractIWETH9\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"depositToVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"depositToVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"depositToVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"depositToVault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"fromVault\",\"type\":\"address\"},{\"internalType\":\"contractIYearn4626\",\"name\":\"toVault\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"fromVault\",\"type\":\"address\"},{\"internalType\":\"contractIYearn4626\",\"name\":\"toVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"fromVault\",\"type\":\"address\"},{\"internalType\":\"contractIYearn4626\",\"name\":\"toVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"fromVault\",\"type\":\"address\"},{\"internalType\":\"contractIYearn4626\",\"name\":\"toVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearnV2\",\"name\":\"fromVault\",\"type\":\"address\"},{\"internalType\":\"contractIYearn4626\",\"name\":\"toVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"migrateFromV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearnV2\",\"name\":\"fromVault\",\"type\":\"address\"},{\"internalType\":\"contractIYearn4626\",\"name\":\"toVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"migrateFromV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearnV2\",\"name\":\"fromVault\",\"type\":\"address\"},{\"internalType\":\"contractIYearn4626\",\"name\":\"toVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minSharesOut\",\"type\":\"uint256\"}],\"name\":\"migrateFromV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearnV2\",\"name\":\"fromVault\",\"type\":\"address\"},{\"internalType\":\"contractIYearn4626\",\"name\":\"toVault\",\"type\":\"address\"}],\"name\":\"migrateFromV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"pullToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"redeemDefault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountMinimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unwrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxLoss\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIYearn4626\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSharesOut\",\"type\":\"uint256\"}],\"name\":\"withdrawDefault\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrapWETH9\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// YearnV3RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnV3RouterMetaData.ABI instead.
var YearnV3RouterABI = YearnV3RouterMetaData.ABI

// YearnV3Router is an auto generated Go binding around an Ethereum contract.
type YearnV3Router struct {
	YearnV3RouterCaller     // Read-only binding to the contract
	YearnV3RouterTransactor // Write-only binding to the contract
	YearnV3RouterFilterer   // Log filterer for contract events
}

// YearnV3RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnV3RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnV3RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnV3RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnV3RouterSession struct {
	Contract     *YearnV3Router    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnV3RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnV3RouterCallerSession struct {
	Contract *YearnV3RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// YearnV3RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnV3RouterTransactorSession struct {
	Contract     *YearnV3RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// YearnV3RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnV3RouterRaw struct {
	Contract *YearnV3Router // Generic contract binding to access the raw methods on
}

// YearnV3RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnV3RouterCallerRaw struct {
	Contract *YearnV3RouterCaller // Generic read-only contract binding to access the raw methods on
}

// YearnV3RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnV3RouterTransactorRaw struct {
	Contract *YearnV3RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnV3Router creates a new instance of YearnV3Router, bound to a specific deployed contract.
func NewYearnV3Router(address common.Address, backend bind.ContractBackend) (*YearnV3Router, error) {
	contract, err := bindYearnV3Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnV3Router{YearnV3RouterCaller: YearnV3RouterCaller{contract: contract}, YearnV3RouterTransactor: YearnV3RouterTransactor{contract: contract}, YearnV3RouterFilterer: YearnV3RouterFilterer{contract: contract}}, nil
}

// NewYearnV3RouterCaller creates a new read-only instance of YearnV3Router, bound to a specific deployed contract.
func NewYearnV3RouterCaller(address common.Address, caller bind.ContractCaller) (*YearnV3RouterCaller, error) {
	contract, err := bindYearnV3Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnV3RouterCaller{contract: contract}, nil
}

// NewYearnV3RouterTransactor creates a new write-only instance of YearnV3Router, bound to a specific deployed contract.
func NewYearnV3RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnV3RouterTransactor, error) {
	contract, err := bindYearnV3Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnV3RouterTransactor{contract: contract}, nil
}

// NewYearnV3RouterFilterer creates a new log filterer instance of YearnV3Router, bound to a specific deployed contract.
func NewYearnV3RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnV3RouterFilterer, error) {
	contract, err := bindYearnV3Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnV3RouterFilterer{contract: contract}, nil
}

// bindYearnV3Router binds a generic wrapper to an already deployed contract.
func bindYearnV3Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YearnV3RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnV3Router *YearnV3RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnV3Router.Contract.YearnV3RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnV3Router *YearnV3RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Router.Contract.YearnV3RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnV3Router *YearnV3RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnV3Router.Contract.YearnV3RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnV3Router *YearnV3RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnV3Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnV3Router *YearnV3RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnV3Router *YearnV3RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnV3Router.Contract.contract.Transact(opts, method, params...)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_YearnV3Router *YearnV3RouterCaller) WETH9(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Router.contract.Call(opts, &out, "WETH9")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_YearnV3Router *YearnV3RouterSession) WETH9() (common.Address, error) {
	return _YearnV3Router.Contract.WETH9(&_YearnV3Router.CallOpts)
}

// WETH9 is a free data retrieval call binding the contract method 0x4aa4a4fc.
//
// Solidity: function WETH9() view returns(address)
func (_YearnV3Router *YearnV3RouterCallerSession) WETH9() (common.Address, error) {
	return _YearnV3Router.Contract.WETH9(&_YearnV3Router.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Router *YearnV3RouterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnV3Router.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Router *YearnV3RouterSession) Name() (string, error) {
	return _YearnV3Router.Contract.Name(&_YearnV3Router.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Router *YearnV3RouterCallerSession) Name() (string, error) {
	return _YearnV3Router.Contract.Name(&_YearnV3Router.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address to, uint256 amount) payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) Approve(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "approve", token, to, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address to, uint256 amount) payable returns()
func (_YearnV3Router *YearnV3RouterSession) Approve(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Approve(&_YearnV3Router.TransactOpts, token, to, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address token, address to, uint256 amount) payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) Approve(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Approve(&_YearnV3Router.TransactOpts, token, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address vault, uint256 amount, address to, uint256 minSharesOut) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterTransactor) Deposit(opts *bind.TransactOpts, vault common.Address, amount *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "deposit", vault, amount, to, minSharesOut)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address vault, uint256 amount, address to, uint256 minSharesOut) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterSession) Deposit(vault common.Address, amount *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Deposit(&_YearnV3Router.TransactOpts, vault, amount, to, minSharesOut)
}

// Deposit is a paid mutator transaction binding the contract method 0x90d25074.
//
// Solidity: function deposit(address vault, uint256 amount, address to, uint256 minSharesOut) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterTransactorSession) Deposit(vault common.Address, amount *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Deposit(&_YearnV3Router.TransactOpts, vault, amount, to, minSharesOut)
}

// DepositToVault is a paid mutator transaction binding the contract method 0x30b3484f.
//
// Solidity: function depositToVault(address vault, uint256 amount, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) DepositToVault(opts *bind.TransactOpts, vault common.Address, amount *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "depositToVault", vault, amount, to, minSharesOut)
}

// DepositToVault is a paid mutator transaction binding the contract method 0x30b3484f.
//
// Solidity: function depositToVault(address vault, uint256 amount, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) DepositToVault(vault common.Address, amount *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.DepositToVault(&_YearnV3Router.TransactOpts, vault, amount, to, minSharesOut)
}

// DepositToVault is a paid mutator transaction binding the contract method 0x30b3484f.
//
// Solidity: function depositToVault(address vault, uint256 amount, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) DepositToVault(vault common.Address, amount *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.DepositToVault(&_YearnV3Router.TransactOpts, vault, amount, to, minSharesOut)
}

// DepositToVault0 is a paid mutator transaction binding the contract method 0xa3d11158.
//
// Solidity: function depositToVault(address vault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) DepositToVault0(opts *bind.TransactOpts, vault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "depositToVault0", vault, minSharesOut)
}

// DepositToVault0 is a paid mutator transaction binding the contract method 0xa3d11158.
//
// Solidity: function depositToVault(address vault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) DepositToVault0(vault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.DepositToVault0(&_YearnV3Router.TransactOpts, vault, minSharesOut)
}

// DepositToVault0 is a paid mutator transaction binding the contract method 0xa3d11158.
//
// Solidity: function depositToVault(address vault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) DepositToVault0(vault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.DepositToVault0(&_YearnV3Router.TransactOpts, vault, minSharesOut)
}

// DepositToVault1 is a paid mutator transaction binding the contract method 0xc28916f6.
//
// Solidity: function depositToVault(address vault, uint256 amount, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) DepositToVault1(opts *bind.TransactOpts, vault common.Address, amount *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "depositToVault1", vault, amount, minSharesOut)
}

// DepositToVault1 is a paid mutator transaction binding the contract method 0xc28916f6.
//
// Solidity: function depositToVault(address vault, uint256 amount, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) DepositToVault1(vault common.Address, amount *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.DepositToVault1(&_YearnV3Router.TransactOpts, vault, amount, minSharesOut)
}

// DepositToVault1 is a paid mutator transaction binding the contract method 0xc28916f6.
//
// Solidity: function depositToVault(address vault, uint256 amount, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) DepositToVault1(vault common.Address, amount *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.DepositToVault1(&_YearnV3Router.TransactOpts, vault, amount, minSharesOut)
}

// DepositToVault2 is a paid mutator transaction binding the contract method 0xefc7a861.
//
// Solidity: function depositToVault(address vault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) DepositToVault2(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "depositToVault2", vault)
}

// DepositToVault2 is a paid mutator transaction binding the contract method 0xefc7a861.
//
// Solidity: function depositToVault(address vault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) DepositToVault2(vault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.DepositToVault2(&_YearnV3Router.TransactOpts, vault)
}

// DepositToVault2 is a paid mutator transaction binding the contract method 0xefc7a861.
//
// Solidity: function depositToVault(address vault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) DepositToVault2(vault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.DepositToVault2(&_YearnV3Router.TransactOpts, vault)
}

// Migrate is a paid mutator transaction binding the contract method 0x1068361f.
//
// Solidity: function migrate(address fromVault, address toVault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Migrate(opts *bind.TransactOpts, fromVault common.Address, toVault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "migrate", fromVault, toVault)
}

// Migrate is a paid mutator transaction binding the contract method 0x1068361f.
//
// Solidity: function migrate(address fromVault, address toVault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Migrate(fromVault common.Address, toVault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Migrate(&_YearnV3Router.TransactOpts, fromVault, toVault)
}

// Migrate is a paid mutator transaction binding the contract method 0x1068361f.
//
// Solidity: function migrate(address fromVault, address toVault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Migrate(fromVault common.Address, toVault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Migrate(&_YearnV3Router.TransactOpts, fromVault, toVault)
}

// Migrate0 is a paid mutator transaction binding the contract method 0x4a250330.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 shares, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Migrate0(opts *bind.TransactOpts, fromVault common.Address, toVault common.Address, shares *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "migrate0", fromVault, toVault, shares, to, minSharesOut)
}

// Migrate0 is a paid mutator transaction binding the contract method 0x4a250330.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 shares, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Migrate0(fromVault common.Address, toVault common.Address, shares *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Migrate0(&_YearnV3Router.TransactOpts, fromVault, toVault, shares, to, minSharesOut)
}

// Migrate0 is a paid mutator transaction binding the contract method 0x4a250330.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 shares, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Migrate0(fromVault common.Address, toVault common.Address, shares *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Migrate0(&_YearnV3Router.TransactOpts, fromVault, toVault, shares, to, minSharesOut)
}

// Migrate1 is a paid mutator transaction binding the contract method 0xa79721fe.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 shares, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Migrate1(opts *bind.TransactOpts, fromVault common.Address, toVault common.Address, shares *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "migrate1", fromVault, toVault, shares, minSharesOut)
}

// Migrate1 is a paid mutator transaction binding the contract method 0xa79721fe.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 shares, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Migrate1(fromVault common.Address, toVault common.Address, shares *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Migrate1(&_YearnV3Router.TransactOpts, fromVault, toVault, shares, minSharesOut)
}

// Migrate1 is a paid mutator transaction binding the contract method 0xa79721fe.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 shares, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Migrate1(fromVault common.Address, toVault common.Address, shares *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Migrate1(&_YearnV3Router.TransactOpts, fromVault, toVault, shares, minSharesOut)
}

// Migrate2 is a paid mutator transaction binding the contract method 0xf16565ee.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Migrate2(opts *bind.TransactOpts, fromVault common.Address, toVault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "migrate2", fromVault, toVault, minSharesOut)
}

// Migrate2 is a paid mutator transaction binding the contract method 0xf16565ee.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Migrate2(fromVault common.Address, toVault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Migrate2(&_YearnV3Router.TransactOpts, fromVault, toVault, minSharesOut)
}

// Migrate2 is a paid mutator transaction binding the contract method 0xf16565ee.
//
// Solidity: function migrate(address fromVault, address toVault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Migrate2(fromVault common.Address, toVault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Migrate2(&_YearnV3Router.TransactOpts, fromVault, toVault, minSharesOut)
}

// MigrateFromV2 is a paid mutator transaction binding the contract method 0x3dff6639.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 shares, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) MigrateFromV2(opts *bind.TransactOpts, fromVault common.Address, toVault common.Address, shares *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "migrateFromV2", fromVault, toVault, shares, to, minSharesOut)
}

// MigrateFromV2 is a paid mutator transaction binding the contract method 0x3dff6639.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 shares, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) MigrateFromV2(fromVault common.Address, toVault common.Address, shares *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.MigrateFromV2(&_YearnV3Router.TransactOpts, fromVault, toVault, shares, to, minSharesOut)
}

// MigrateFromV2 is a paid mutator transaction binding the contract method 0x3dff6639.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 shares, address to, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) MigrateFromV2(fromVault common.Address, toVault common.Address, shares *big.Int, to common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.MigrateFromV2(&_YearnV3Router.TransactOpts, fromVault, toVault, shares, to, minSharesOut)
}

// MigrateFromV20 is a paid mutator transaction binding the contract method 0x9b4f09af.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 shares, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) MigrateFromV20(opts *bind.TransactOpts, fromVault common.Address, toVault common.Address, shares *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "migrateFromV20", fromVault, toVault, shares, minSharesOut)
}

// MigrateFromV20 is a paid mutator transaction binding the contract method 0x9b4f09af.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 shares, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) MigrateFromV20(fromVault common.Address, toVault common.Address, shares *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.MigrateFromV20(&_YearnV3Router.TransactOpts, fromVault, toVault, shares, minSharesOut)
}

// MigrateFromV20 is a paid mutator transaction binding the contract method 0x9b4f09af.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 shares, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) MigrateFromV20(fromVault common.Address, toVault common.Address, shares *big.Int, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.MigrateFromV20(&_YearnV3Router.TransactOpts, fromVault, toVault, shares, minSharesOut)
}

// MigrateFromV21 is a paid mutator transaction binding the contract method 0xa3ebe717.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) MigrateFromV21(opts *bind.TransactOpts, fromVault common.Address, toVault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "migrateFromV21", fromVault, toVault, minSharesOut)
}

// MigrateFromV21 is a paid mutator transaction binding the contract method 0xa3ebe717.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) MigrateFromV21(fromVault common.Address, toVault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.MigrateFromV21(&_YearnV3Router.TransactOpts, fromVault, toVault, minSharesOut)
}

// MigrateFromV21 is a paid mutator transaction binding the contract method 0xa3ebe717.
//
// Solidity: function migrateFromV2(address fromVault, address toVault, uint256 minSharesOut) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) MigrateFromV21(fromVault common.Address, toVault common.Address, minSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.MigrateFromV21(&_YearnV3Router.TransactOpts, fromVault, toVault, minSharesOut)
}

// MigrateFromV22 is a paid mutator transaction binding the contract method 0xdbae0292.
//
// Solidity: function migrateFromV2(address fromVault, address toVault) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterTransactor) MigrateFromV22(opts *bind.TransactOpts, fromVault common.Address, toVault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "migrateFromV22", fromVault, toVault)
}

// MigrateFromV22 is a paid mutator transaction binding the contract method 0xdbae0292.
//
// Solidity: function migrateFromV2(address fromVault, address toVault) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterSession) MigrateFromV22(fromVault common.Address, toVault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.MigrateFromV22(&_YearnV3Router.TransactOpts, fromVault, toVault)
}

// MigrateFromV22 is a paid mutator transaction binding the contract method 0xdbae0292.
//
// Solidity: function migrateFromV2(address fromVault, address toVault) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterTransactorSession) MigrateFromV22(fromVault common.Address, toVault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.MigrateFromV22(&_YearnV3Router.TransactOpts, fromVault, toVault)
}

// Mint is a paid mutator transaction binding the contract method 0x3c173a4f.
//
// Solidity: function mint(address vault, uint256 shares, address to, uint256 maxAmountIn) payable returns(uint256 amountIn)
func (_YearnV3Router *YearnV3RouterTransactor) Mint(opts *bind.TransactOpts, vault common.Address, shares *big.Int, to common.Address, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "mint", vault, shares, to, maxAmountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x3c173a4f.
//
// Solidity: function mint(address vault, uint256 shares, address to, uint256 maxAmountIn) payable returns(uint256 amountIn)
func (_YearnV3Router *YearnV3RouterSession) Mint(vault common.Address, shares *big.Int, to common.Address, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Mint(&_YearnV3Router.TransactOpts, vault, shares, to, maxAmountIn)
}

// Mint is a paid mutator transaction binding the contract method 0x3c173a4f.
//
// Solidity: function mint(address vault, uint256 shares, address to, uint256 maxAmountIn) payable returns(uint256 amountIn)
func (_YearnV3Router *YearnV3RouterTransactorSession) Mint(vault common.Address, shares *big.Int, to common.Address, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Mint(&_YearnV3Router.TransactOpts, vault, shares, to, maxAmountIn)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_YearnV3Router *YearnV3RouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_YearnV3Router *YearnV3RouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Multicall(&_YearnV3Router.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_YearnV3Router *YearnV3RouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Multicall(&_YearnV3Router.TransactOpts, data)
}

// PullToken is a paid mutator transaction binding the contract method 0x73d15414.
//
// Solidity: function pullToken(address token, uint256 amount, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) PullToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "pullToken", token, amount, recipient)
}

// PullToken is a paid mutator transaction binding the contract method 0x73d15414.
//
// Solidity: function pullToken(address token, uint256 amount, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterSession) PullToken(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.PullToken(&_YearnV3Router.TransactOpts, token, amount, recipient)
}

// PullToken is a paid mutator transaction binding the contract method 0x73d15414.
//
// Solidity: function pullToken(address token, uint256 amount, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) PullToken(token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.PullToken(&_YearnV3Router.TransactOpts, token, amount, recipient)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address vault, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Redeem(opts *bind.TransactOpts, vault common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "redeem", vault, maxLoss)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address vault, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Redeem(vault common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Redeem(&_YearnV3Router.TransactOpts, vault, maxLoss)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address vault, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Redeem(vault common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Redeem(&_YearnV3Router.TransactOpts, vault, maxLoss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address vault, uint256 shares, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Redeem0(opts *bind.TransactOpts, vault common.Address, shares *big.Int, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "redeem0", vault, shares, maxLoss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address vault, uint256 shares, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Redeem0(vault common.Address, shares *big.Int, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Redeem0(&_YearnV3Router.TransactOpts, vault, shares, maxLoss)
}

// Redeem0 is a paid mutator transaction binding the contract method 0x2b83cccd.
//
// Solidity: function redeem(address vault, uint256 shares, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Redeem0(vault common.Address, shares *big.Int, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Redeem0(&_YearnV3Router.TransactOpts, vault, shares, maxLoss)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(address vault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Redeem1(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "redeem1", vault)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(address vault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Redeem1(vault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Redeem1(&_YearnV3Router.TransactOpts, vault)
}

// Redeem1 is a paid mutator transaction binding the contract method 0x95a2251f.
//
// Solidity: function redeem(address vault) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Redeem1(vault common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Redeem1(&_YearnV3Router.TransactOpts, vault)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address vault, uint256 shares, address to, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Redeem2(opts *bind.TransactOpts, vault common.Address, shares *big.Int, to common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "redeem2", vault, shares, to, maxLoss)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address vault, uint256 shares, address to, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Redeem2(vault common.Address, shares *big.Int, to common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Redeem2(&_YearnV3Router.TransactOpts, vault, shares, to, maxLoss)
}

// Redeem2 is a paid mutator transaction binding the contract method 0xf3f094a1.
//
// Solidity: function redeem(address vault, uint256 shares, address to, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Redeem2(vault common.Address, shares *big.Int, to common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Redeem2(&_YearnV3Router.TransactOpts, vault, shares, to, maxLoss)
}

// RedeemDefault is a paid mutator transaction binding the contract method 0x5b9a66ee.
//
// Solidity: function redeemDefault(address vault, uint256 shares, address to, uint256 minAmountOut) payable returns(uint256 amountOut)
func (_YearnV3Router *YearnV3RouterTransactor) RedeemDefault(opts *bind.TransactOpts, vault common.Address, shares *big.Int, to common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "redeemDefault", vault, shares, to, minAmountOut)
}

// RedeemDefault is a paid mutator transaction binding the contract method 0x5b9a66ee.
//
// Solidity: function redeemDefault(address vault, uint256 shares, address to, uint256 minAmountOut) payable returns(uint256 amountOut)
func (_YearnV3Router *YearnV3RouterSession) RedeemDefault(vault common.Address, shares *big.Int, to common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.RedeemDefault(&_YearnV3Router.TransactOpts, vault, shares, to, minAmountOut)
}

// RedeemDefault is a paid mutator transaction binding the contract method 0x5b9a66ee.
//
// Solidity: function redeemDefault(address vault, uint256 shares, address to, uint256 minAmountOut) payable returns(uint256 amountOut)
func (_YearnV3Router *YearnV3RouterTransactorSession) RedeemDefault(vault common.Address, shares *big.Int, to common.Address, minAmountOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.RedeemDefault(&_YearnV3Router.TransactOpts, vault, shares, to, minAmountOut)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_YearnV3Router *YearnV3RouterSession) RefundETH() (*types.Transaction, error) {
	return _YearnV3Router.Contract.RefundETH(&_YearnV3Router.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) RefundETH() (*types.Transaction, error) {
	return _YearnV3Router.Contract.RefundETH(&_YearnV3Router.TransactOpts)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SelfPermit(&_YearnV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SelfPermit(&_YearnV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SelfPermitAllowed(&_YearnV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SelfPermitAllowed(&_YearnV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SelfPermitAllowedIfNecessary(&_YearnV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SelfPermitAllowedIfNecessary(&_YearnV3Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SelfPermitIfNecessary(&_YearnV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SelfPermitIfNecessary(&_YearnV3Router.TransactOpts, token, value, deadline, v, r, s)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) SweepToken(opts *bind.TransactOpts, token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "sweepToken", token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SweepToken(&_YearnV3Router.TransactOpts, token, amountMinimum, recipient)
}

// SweepToken is a paid mutator transaction binding the contract method 0xdf2ab5bb.
//
// Solidity: function sweepToken(address token, uint256 amountMinimum, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) SweepToken(token common.Address, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.SweepToken(&_YearnV3Router.TransactOpts, token, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) UnwrapWETH9(opts *bind.TransactOpts, amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "unwrapWETH9", amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.UnwrapWETH9(&_YearnV3Router.TransactOpts, amountMinimum, recipient)
}

// UnwrapWETH9 is a paid mutator transaction binding the contract method 0x49404b7c.
//
// Solidity: function unwrapWETH9(uint256 amountMinimum, address recipient) payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) UnwrapWETH9(amountMinimum *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _YearnV3Router.Contract.UnwrapWETH9(&_YearnV3Router.TransactOpts, amountMinimum, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4b2084e3.
//
// Solidity: function withdraw(address vault, uint256 amount, address to, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactor) Withdraw(opts *bind.TransactOpts, vault common.Address, amount *big.Int, to common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "withdraw", vault, amount, to, maxLoss)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4b2084e3.
//
// Solidity: function withdraw(address vault, uint256 amount, address to, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterSession) Withdraw(vault common.Address, amount *big.Int, to common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Withdraw(&_YearnV3Router.TransactOpts, vault, amount, to, maxLoss)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4b2084e3.
//
// Solidity: function withdraw(address vault, uint256 amount, address to, uint256 maxLoss) payable returns(uint256)
func (_YearnV3Router *YearnV3RouterTransactorSession) Withdraw(vault common.Address, amount *big.Int, to common.Address, maxLoss *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.Withdraw(&_YearnV3Router.TransactOpts, vault, amount, to, maxLoss)
}

// WithdrawDefault is a paid mutator transaction binding the contract method 0x6f63427e.
//
// Solidity: function withdrawDefault(address vault, uint256 amount, address to, uint256 maxSharesOut) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterTransactor) WithdrawDefault(opts *bind.TransactOpts, vault common.Address, amount *big.Int, to common.Address, maxSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "withdrawDefault", vault, amount, to, maxSharesOut)
}

// WithdrawDefault is a paid mutator transaction binding the contract method 0x6f63427e.
//
// Solidity: function withdrawDefault(address vault, uint256 amount, address to, uint256 maxSharesOut) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterSession) WithdrawDefault(vault common.Address, amount *big.Int, to common.Address, maxSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.WithdrawDefault(&_YearnV3Router.TransactOpts, vault, amount, to, maxSharesOut)
}

// WithdrawDefault is a paid mutator transaction binding the contract method 0x6f63427e.
//
// Solidity: function withdrawDefault(address vault, uint256 amount, address to, uint256 maxSharesOut) payable returns(uint256 sharesOut)
func (_YearnV3Router *YearnV3RouterTransactorSession) WithdrawDefault(vault common.Address, amount *big.Int, to common.Address, maxSharesOut *big.Int) (*types.Transaction, error) {
	return _YearnV3Router.Contract.WithdrawDefault(&_YearnV3Router.TransactOpts, vault, amount, to, maxSharesOut)
}

// WrapWETH9 is a paid mutator transaction binding the contract method 0x3f50fd1f.
//
// Solidity: function wrapWETH9() payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) WrapWETH9(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Router.contract.Transact(opts, "wrapWETH9")
}

// WrapWETH9 is a paid mutator transaction binding the contract method 0x3f50fd1f.
//
// Solidity: function wrapWETH9() payable returns()
func (_YearnV3Router *YearnV3RouterSession) WrapWETH9() (*types.Transaction, error) {
	return _YearnV3Router.Contract.WrapWETH9(&_YearnV3Router.TransactOpts)
}

// WrapWETH9 is a paid mutator transaction binding the contract method 0x3f50fd1f.
//
// Solidity: function wrapWETH9() payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) WrapWETH9() (*types.Transaction, error) {
	return _YearnV3Router.Contract.WrapWETH9(&_YearnV3Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_YearnV3Router *YearnV3RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_YearnV3Router *YearnV3RouterSession) Receive() (*types.Transaction, error) {
	return _YearnV3Router.Contract.Receive(&_YearnV3Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_YearnV3Router *YearnV3RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _YearnV3Router.Contract.Receive(&_YearnV3Router.TransactOpts)
}
