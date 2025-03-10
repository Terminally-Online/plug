// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_assert

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

// PlugAssertMetaData contains all meta data concerning the PlugAssert contract.
var PlugAssertMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"assertFalse\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertFalse\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertTrue\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertTrue\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"fail\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"}]",
}

// PlugAssertABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugAssertMetaData.ABI instead.
var PlugAssertABI = PlugAssertMetaData.ABI

// PlugAssert is an auto generated Go binding around an Ethereum contract.
type PlugAssert struct {
	PlugAssertCaller     // Read-only binding to the contract
	PlugAssertTransactor // Write-only binding to the contract
	PlugAssertFilterer   // Log filterer for contract events
}

// PlugAssertCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugAssertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugAssertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugAssertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugAssertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugAssertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugAssertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugAssertSession struct {
	Contract     *PlugAssert       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugAssertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugAssertCallerSession struct {
	Contract *PlugAssertCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PlugAssertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugAssertTransactorSession struct {
	Contract     *PlugAssertTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PlugAssertRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugAssertRaw struct {
	Contract *PlugAssert // Generic contract binding to access the raw methods on
}

// PlugAssertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugAssertCallerRaw struct {
	Contract *PlugAssertCaller // Generic read-only contract binding to access the raw methods on
}

// PlugAssertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugAssertTransactorRaw struct {
	Contract *PlugAssertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugAssert creates a new instance of PlugAssert, bound to a specific deployed contract.
func NewPlugAssert(address common.Address, backend bind.ContractBackend) (*PlugAssert, error) {
	contract, err := bindPlugAssert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugAssert{PlugAssertCaller: PlugAssertCaller{contract: contract}, PlugAssertTransactor: PlugAssertTransactor{contract: contract}, PlugAssertFilterer: PlugAssertFilterer{contract: contract}}, nil
}

// NewPlugAssertCaller creates a new read-only instance of PlugAssert, bound to a specific deployed contract.
func NewPlugAssertCaller(address common.Address, caller bind.ContractCaller) (*PlugAssertCaller, error) {
	contract, err := bindPlugAssert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugAssertCaller{contract: contract}, nil
}

// NewPlugAssertTransactor creates a new write-only instance of PlugAssert, bound to a specific deployed contract.
func NewPlugAssertTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugAssertTransactor, error) {
	contract, err := bindPlugAssert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugAssertTransactor{contract: contract}, nil
}

// NewPlugAssertFilterer creates a new log filterer instance of PlugAssert, bound to a specific deployed contract.
func NewPlugAssertFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugAssertFilterer, error) {
	contract, err := bindPlugAssert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugAssertFilterer{contract: contract}, nil
}

// bindPlugAssert binds a generic wrapper to an already deployed contract.
func bindPlugAssert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugAssertMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugAssert *PlugAssertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugAssert.Contract.PlugAssertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugAssert *PlugAssertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugAssert.Contract.PlugAssertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugAssert *PlugAssertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugAssert.Contract.PlugAssertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugAssert *PlugAssertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugAssert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugAssert *PlugAssertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugAssert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugAssert *PlugAssertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugAssert.Contract.contract.Transact(opts, method, params...)
}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string message) pure returns()
func (_PlugAssert *PlugAssertCaller) AssertFalse(opts *bind.CallOpts, condition bool, message string) error {
	var out []interface{}
	err := _PlugAssert.contract.Call(opts, &out, "assertFalse", condition, message)

	if err != nil {
		return err
	}

	return err

}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string message) pure returns()
func (_PlugAssert *PlugAssertSession) AssertFalse(condition bool, message string) error {
	return _PlugAssert.Contract.AssertFalse(&_PlugAssert.CallOpts, condition, message)
}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string message) pure returns()
func (_PlugAssert *PlugAssertCallerSession) AssertFalse(condition bool, message string) error {
	return _PlugAssert.Contract.AssertFalse(&_PlugAssert.CallOpts, condition, message)
}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_PlugAssert *PlugAssertCaller) AssertFalse0(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _PlugAssert.contract.Call(opts, &out, "assertFalse0", condition)

	if err != nil {
		return err
	}

	return err

}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_PlugAssert *PlugAssertSession) AssertFalse0(condition bool) error {
	return _PlugAssert.Contract.AssertFalse0(&_PlugAssert.CallOpts, condition)
}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_PlugAssert *PlugAssertCallerSession) AssertFalse0(condition bool) error {
	return _PlugAssert.Contract.AssertFalse0(&_PlugAssert.CallOpts, condition)
}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_PlugAssert *PlugAssertCaller) AssertTrue(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _PlugAssert.contract.Call(opts, &out, "assertTrue", condition)

	if err != nil {
		return err
	}

	return err

}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_PlugAssert *PlugAssertSession) AssertTrue(condition bool) error {
	return _PlugAssert.Contract.AssertTrue(&_PlugAssert.CallOpts, condition)
}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_PlugAssert *PlugAssertCallerSession) AssertTrue(condition bool) error {
	return _PlugAssert.Contract.AssertTrue(&_PlugAssert.CallOpts, condition)
}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string message) pure returns()
func (_PlugAssert *PlugAssertCaller) AssertTrue0(opts *bind.CallOpts, condition bool, message string) error {
	var out []interface{}
	err := _PlugAssert.contract.Call(opts, &out, "assertTrue0", condition, message)

	if err != nil {
		return err
	}

	return err

}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string message) pure returns()
func (_PlugAssert *PlugAssertSession) AssertTrue0(condition bool, message string) error {
	return _PlugAssert.Contract.AssertTrue0(&_PlugAssert.CallOpts, condition, message)
}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string message) pure returns()
func (_PlugAssert *PlugAssertCallerSession) AssertTrue0(condition bool, message string) error {
	return _PlugAssert.Contract.AssertTrue0(&_PlugAssert.CallOpts, condition, message)
}

// Fail is a free data retrieval call binding the contract method 0x78122f3a.
//
// Solidity: function fail(string message) pure returns()
func (_PlugAssert *PlugAssertCaller) Fail(opts *bind.CallOpts, message string) error {
	var out []interface{}
	err := _PlugAssert.contract.Call(opts, &out, "fail", message)

	if err != nil {
		return err
	}

	return err

}

// Fail is a free data retrieval call binding the contract method 0x78122f3a.
//
// Solidity: function fail(string message) pure returns()
func (_PlugAssert *PlugAssertSession) Fail(message string) error {
	return _PlugAssert.Contract.Fail(&_PlugAssert.CallOpts, message)
}

// Fail is a free data retrieval call binding the contract method 0x78122f3a.
//
// Solidity: function fail(string message) pure returns()
func (_PlugAssert *PlugAssertCallerSession) Fail(message string) error {
	return _PlugAssert.Contract.Fail(&_PlugAssert.CallOpts, message)
}
