// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave_v2_pool

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

// AaveV2PoolMetaData contains all meta data concerning the AaveV2Pool contract.
var AaveV2PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AaveV2PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveV2PoolMetaData.ABI instead.
var AaveV2PoolABI = AaveV2PoolMetaData.ABI

// AaveV2Pool is an auto generated Go binding around an Ethereum contract.
type AaveV2Pool struct {
	AaveV2PoolCaller     // Read-only binding to the contract
	AaveV2PoolTransactor // Write-only binding to the contract
	AaveV2PoolFilterer   // Log filterer for contract events
}

// AaveV2PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV2PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV2PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV2PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV2PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV2PoolSession struct {
	Contract     *AaveV2Pool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveV2PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV2PoolCallerSession struct {
	Contract *AaveV2PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AaveV2PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV2PoolTransactorSession struct {
	Contract     *AaveV2PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AaveV2PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV2PoolRaw struct {
	Contract *AaveV2Pool // Generic contract binding to access the raw methods on
}

// AaveV2PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV2PoolCallerRaw struct {
	Contract *AaveV2PoolCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV2PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV2PoolTransactorRaw struct {
	Contract *AaveV2PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV2Pool creates a new instance of AaveV2Pool, bound to a specific deployed contract.
func NewAaveV2Pool(address common.Address, backend bind.ContractBackend) (*AaveV2Pool, error) {
	contract, err := bindAaveV2Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV2Pool{AaveV2PoolCaller: AaveV2PoolCaller{contract: contract}, AaveV2PoolTransactor: AaveV2PoolTransactor{contract: contract}, AaveV2PoolFilterer: AaveV2PoolFilterer{contract: contract}}, nil
}

// NewAaveV2PoolCaller creates a new read-only instance of AaveV2Pool, bound to a specific deployed contract.
func NewAaveV2PoolCaller(address common.Address, caller bind.ContractCaller) (*AaveV2PoolCaller, error) {
	contract, err := bindAaveV2Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2PoolCaller{contract: contract}, nil
}

// NewAaveV2PoolTransactor creates a new write-only instance of AaveV2Pool, bound to a specific deployed contract.
func NewAaveV2PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV2PoolTransactor, error) {
	contract, err := bindAaveV2Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV2PoolTransactor{contract: contract}, nil
}

// NewAaveV2PoolFilterer creates a new log filterer instance of AaveV2Pool, bound to a specific deployed contract.
func NewAaveV2PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV2PoolFilterer, error) {
	contract, err := bindAaveV2Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV2PoolFilterer{contract: contract}, nil
}

// bindAaveV2Pool binds a generic wrapper to an already deployed contract.
func bindAaveV2Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveV2PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2Pool *AaveV2PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV2Pool.Contract.AaveV2PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2Pool *AaveV2PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.AaveV2PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2Pool *AaveV2PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.AaveV2PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV2Pool *AaveV2PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV2Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV2Pool *AaveV2PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV2Pool *AaveV2PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AaveV2Pool *AaveV2PoolTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2Pool.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AaveV2Pool *AaveV2PoolSession) Admin() (*types.Transaction, error) {
	return _AaveV2Pool.Contract.Admin(&_AaveV2Pool.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address)
func (_AaveV2Pool *AaveV2PoolTransactorSession) Admin() (*types.Transaction, error) {
	return _AaveV2Pool.Contract.Admin(&_AaveV2Pool.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AaveV2Pool *AaveV2PoolTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV2Pool.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AaveV2Pool *AaveV2PoolSession) Implementation() (*types.Transaction, error) {
	return _AaveV2Pool.Contract.Implementation(&_AaveV2Pool.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address)
func (_AaveV2Pool *AaveV2PoolTransactorSession) Implementation() (*types.Transaction, error) {
	return _AaveV2Pool.Contract.Implementation(&_AaveV2Pool.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_AaveV2Pool *AaveV2PoolTransactor) Initialize(opts *bind.TransactOpts, _logic common.Address, _data []byte) (*types.Transaction, error) {
	return _AaveV2Pool.contract.Transact(opts, "initialize", _logic, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_AaveV2Pool *AaveV2PoolSession) Initialize(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.Initialize(&_AaveV2Pool.TransactOpts, _logic, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _logic, bytes _data) payable returns()
func (_AaveV2Pool *AaveV2PoolTransactorSession) Initialize(_logic common.Address, _data []byte) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.Initialize(&_AaveV2Pool.TransactOpts, _logic, _data)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AaveV2Pool *AaveV2PoolTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AaveV2Pool.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AaveV2Pool *AaveV2PoolSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.UpgradeTo(&_AaveV2Pool.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AaveV2Pool *AaveV2PoolTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.UpgradeTo(&_AaveV2Pool.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AaveV2Pool *AaveV2PoolTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AaveV2Pool.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AaveV2Pool *AaveV2PoolSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.UpgradeToAndCall(&_AaveV2Pool.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AaveV2Pool *AaveV2PoolTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.UpgradeToAndCall(&_AaveV2Pool.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AaveV2Pool *AaveV2PoolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AaveV2Pool.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AaveV2Pool *AaveV2PoolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.Fallback(&_AaveV2Pool.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AaveV2Pool *AaveV2PoolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AaveV2Pool.Contract.Fallback(&_AaveV2Pool.TransactOpts, calldata)
}

// AaveV2PoolUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AaveV2Pool contract.
type AaveV2PoolUpgradedIterator struct {
	Event *AaveV2PoolUpgraded // Event containing the contract specifics and raw log

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
func (it *AaveV2PoolUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV2PoolUpgraded)
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
		it.Event = new(AaveV2PoolUpgraded)
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
func (it *AaveV2PoolUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV2PoolUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV2PoolUpgraded represents a Upgraded event raised by the AaveV2Pool contract.
type AaveV2PoolUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AaveV2Pool *AaveV2PoolFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AaveV2PoolUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AaveV2Pool.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AaveV2PoolUpgradedIterator{contract: _AaveV2Pool.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AaveV2Pool *AaveV2PoolFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AaveV2PoolUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AaveV2Pool.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV2PoolUpgraded)
				if err := _AaveV2Pool.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AaveV2Pool *AaveV2PoolFilterer) ParseUpgraded(log types.Log) (*AaveV2PoolUpgraded, error) {
	event := new(AaveV2PoolUpgraded)
	if err := _AaveV2Pool.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
