// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_database

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

// PlugDatabaseMetaData contains all meta data concerning the PlugDatabase contract.
var PlugDatabaseMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"database\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"remove\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"set\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"}]",
}

// PlugDatabaseABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugDatabaseMetaData.ABI instead.
var PlugDatabaseABI = PlugDatabaseMetaData.ABI

// PlugDatabase is an auto generated Go binding around an Ethereum contract.
type PlugDatabase struct {
	PlugDatabaseCaller     // Read-only binding to the contract
	PlugDatabaseTransactor // Write-only binding to the contract
	PlugDatabaseFilterer   // Log filterer for contract events
}

// PlugDatabaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugDatabaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugDatabaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugDatabaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugDatabaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugDatabaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugDatabaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugDatabaseSession struct {
	Contract     *PlugDatabase     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugDatabaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugDatabaseCallerSession struct {
	Contract *PlugDatabaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PlugDatabaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugDatabaseTransactorSession struct {
	Contract     *PlugDatabaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PlugDatabaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugDatabaseRaw struct {
	Contract *PlugDatabase // Generic contract binding to access the raw methods on
}

// PlugDatabaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugDatabaseCallerRaw struct {
	Contract *PlugDatabaseCaller // Generic read-only contract binding to access the raw methods on
}

// PlugDatabaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugDatabaseTransactorRaw struct {
	Contract *PlugDatabaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugDatabase creates a new instance of PlugDatabase, bound to a specific deployed contract.
func NewPlugDatabase(address common.Address, backend bind.ContractBackend) (*PlugDatabase, error) {
	contract, err := bindPlugDatabase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugDatabase{PlugDatabaseCaller: PlugDatabaseCaller{contract: contract}, PlugDatabaseTransactor: PlugDatabaseTransactor{contract: contract}, PlugDatabaseFilterer: PlugDatabaseFilterer{contract: contract}}, nil
}

// NewPlugDatabaseCaller creates a new read-only instance of PlugDatabase, bound to a specific deployed contract.
func NewPlugDatabaseCaller(address common.Address, caller bind.ContractCaller) (*PlugDatabaseCaller, error) {
	contract, err := bindPlugDatabase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugDatabaseCaller{contract: contract}, nil
}

// NewPlugDatabaseTransactor creates a new write-only instance of PlugDatabase, bound to a specific deployed contract.
func NewPlugDatabaseTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugDatabaseTransactor, error) {
	contract, err := bindPlugDatabase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugDatabaseTransactor{contract: contract}, nil
}

// NewPlugDatabaseFilterer creates a new log filterer instance of PlugDatabase, bound to a specific deployed contract.
func NewPlugDatabaseFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugDatabaseFilterer, error) {
	contract, err := bindPlugDatabase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugDatabaseFilterer{contract: contract}, nil
}

// bindPlugDatabase binds a generic wrapper to an already deployed contract.
func bindPlugDatabase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugDatabaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugDatabase *PlugDatabaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugDatabase.Contract.PlugDatabaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugDatabase *PlugDatabaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugDatabase.Contract.PlugDatabaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugDatabase *PlugDatabaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugDatabase.Contract.PlugDatabaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugDatabase *PlugDatabaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugDatabase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugDatabase *PlugDatabaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugDatabase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugDatabase *PlugDatabaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugDatabase.Contract.contract.Transact(opts, method, params...)
}

// Database is a free data retrieval call binding the contract method 0x70da1dd0.
//
// Solidity: function database(address , bytes32 ) view returns(bytes32)
func (_PlugDatabase *PlugDatabaseCaller) Database(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "database", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Database is a free data retrieval call binding the contract method 0x70da1dd0.
//
// Solidity: function database(address , bytes32 ) view returns(bytes32)
func (_PlugDatabase *PlugDatabaseSession) Database(arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	return _PlugDatabase.Contract.Database(&_PlugDatabase.CallOpts, arg0, arg1)
}

// Database is a free data retrieval call binding the contract method 0x70da1dd0.
//
// Solidity: function database(address , bytes32 ) view returns(bytes32)
func (_PlugDatabase *PlugDatabaseCallerSession) Database(arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	return _PlugDatabase.Contract.Database(&_PlugDatabase.CallOpts, arg0, arg1)
}

// Get is a free data retrieval call binding the contract method 0x7b82d74e.
//
// Solidity: function get(address sender, bytes32 key) view returns(bytes32 result)
func (_PlugDatabase *PlugDatabaseCaller) Get(opts *bind.CallOpts, sender common.Address, key [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "get", sender, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x7b82d74e.
//
// Solidity: function get(address sender, bytes32 key) view returns(bytes32 result)
func (_PlugDatabase *PlugDatabaseSession) Get(sender common.Address, key [32]byte) ([32]byte, error) {
	return _PlugDatabase.Contract.Get(&_PlugDatabase.CallOpts, sender, key)
}

// Get is a free data retrieval call binding the contract method 0x7b82d74e.
//
// Solidity: function get(address sender, bytes32 key) view returns(bytes32 result)
func (_PlugDatabase *PlugDatabaseCallerSession) Get(sender common.Address, key [32]byte) ([32]byte, error) {
	return _PlugDatabase.Contract.Get(&_PlugDatabase.CallOpts, sender, key)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 key) returns()
func (_PlugDatabase *PlugDatabaseTransactor) Remove(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "remove", key)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 key) returns()
func (_PlugDatabase *PlugDatabaseSession) Remove(key [32]byte) (*types.Transaction, error) {
	return _PlugDatabase.Contract.Remove(&_PlugDatabase.TransactOpts, key)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 key) returns()
func (_PlugDatabase *PlugDatabaseTransactorSession) Remove(key [32]byte) (*types.Transaction, error) {
	return _PlugDatabase.Contract.Remove(&_PlugDatabase.TransactOpts, key)
}

// Set is a paid mutator transaction binding the contract method 0xf71f7a25.
//
// Solidity: function set(bytes32 key, bytes32 value) returns(bytes32 result)
func (_PlugDatabase *PlugDatabaseTransactor) Set(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "set", key, value)
}

// Set is a paid mutator transaction binding the contract method 0xf71f7a25.
//
// Solidity: function set(bytes32 key, bytes32 value) returns(bytes32 result)
func (_PlugDatabase *PlugDatabaseSession) Set(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _PlugDatabase.Contract.Set(&_PlugDatabase.TransactOpts, key, value)
}

// Set is a paid mutator transaction binding the contract method 0xf71f7a25.
//
// Solidity: function set(bytes32 key, bytes32 value) returns(bytes32 result)
func (_PlugDatabase *PlugDatabaseTransactorSession) Set(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _PlugDatabase.Contract.Set(&_PlugDatabase.TransactOpts, key, value)
}
