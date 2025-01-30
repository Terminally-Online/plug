// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_factory

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

// PlugFactoryMetaData contains all meta data concerning the PlugFactory contract.
var PlugFactoryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"deploy\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"alreadyDeployed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"socketAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initCodeHash\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"initCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"SocketDeployed\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"SaltInvalid\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// PlugFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugFactoryMetaData.ABI instead.
var PlugFactoryABI = PlugFactoryMetaData.ABI

// PlugFactory is an auto generated Go binding around an Ethereum contract.
type PlugFactory struct {
	PlugFactoryCaller     // Read-only binding to the contract
	PlugFactoryTransactor // Write-only binding to the contract
	PlugFactoryFilterer   // Log filterer for contract events
}

// PlugFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugFactorySession struct {
	Contract     *PlugFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugFactoryCallerSession struct {
	Contract *PlugFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PlugFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugFactoryTransactorSession struct {
	Contract     *PlugFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PlugFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugFactoryRaw struct {
	Contract *PlugFactory // Generic contract binding to access the raw methods on
}

// PlugFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugFactoryCallerRaw struct {
	Contract *PlugFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PlugFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugFactoryTransactorRaw struct {
	Contract *PlugFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugFactory creates a new instance of PlugFactory, bound to a specific deployed contract.
func NewPlugFactory(address common.Address, backend bind.ContractBackend) (*PlugFactory, error) {
	contract, err := bindPlugFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugFactory{PlugFactoryCaller: PlugFactoryCaller{contract: contract}, PlugFactoryTransactor: PlugFactoryTransactor{contract: contract}, PlugFactoryFilterer: PlugFactoryFilterer{contract: contract}}, nil
}

// NewPlugFactoryCaller creates a new read-only instance of PlugFactory, bound to a specific deployed contract.
func NewPlugFactoryCaller(address common.Address, caller bind.ContractCaller) (*PlugFactoryCaller, error) {
	contract, err := bindPlugFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugFactoryCaller{contract: contract}, nil
}

// NewPlugFactoryTransactor creates a new write-only instance of PlugFactory, bound to a specific deployed contract.
func NewPlugFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugFactoryTransactor, error) {
	contract, err := bindPlugFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugFactoryTransactor{contract: contract}, nil
}

// NewPlugFactoryFilterer creates a new log filterer instance of PlugFactory, bound to a specific deployed contract.
func NewPlugFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugFactoryFilterer, error) {
	contract, err := bindPlugFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugFactoryFilterer{contract: contract}, nil
}

// bindPlugFactory binds a generic wrapper to an already deployed contract.
func bindPlugFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugFactory *PlugFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugFactory.Contract.PlugFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugFactory *PlugFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugFactory.Contract.PlugFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugFactory *PlugFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugFactory.Contract.PlugFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugFactory *PlugFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugFactory *PlugFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugFactory *PlugFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugFactory.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address implementation, bytes32 salt) view returns(address vault)
func (_PlugFactory *PlugFactoryCaller) GetAddress(opts *bind.CallOpts, implementation common.Address, salt [32]byte) (common.Address, error) {
	var out []interface{}
	err := _PlugFactory.contract.Call(opts, &out, "getAddress", implementation, salt)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address implementation, bytes32 salt) view returns(address vault)
func (_PlugFactory *PlugFactorySession) GetAddress(implementation common.Address, salt [32]byte) (common.Address, error) {
	return _PlugFactory.Contract.GetAddress(&_PlugFactory.CallOpts, implementation, salt)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address implementation, bytes32 salt) view returns(address vault)
func (_PlugFactory *PlugFactoryCallerSession) GetAddress(implementation common.Address, salt [32]byte) (common.Address, error) {
	return _PlugFactory.Contract.GetAddress(&_PlugFactory.CallOpts, implementation, salt)
}

// InitCodeHash is a free data retrieval call binding the contract method 0x75fd9f28.
//
// Solidity: function initCodeHash(address implementation) view returns(bytes32 initCodeHash)
func (_PlugFactory *PlugFactoryCaller) InitCodeHash(opts *bind.CallOpts, implementation common.Address) ([32]byte, error) {
	var out []interface{}
	err := _PlugFactory.contract.Call(opts, &out, "initCodeHash", implementation)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InitCodeHash is a free data retrieval call binding the contract method 0x75fd9f28.
//
// Solidity: function initCodeHash(address implementation) view returns(bytes32 initCodeHash)
func (_PlugFactory *PlugFactorySession) InitCodeHash(implementation common.Address) ([32]byte, error) {
	return _PlugFactory.Contract.InitCodeHash(&_PlugFactory.CallOpts, implementation)
}

// InitCodeHash is a free data retrieval call binding the contract method 0x75fd9f28.
//
// Solidity: function initCodeHash(address implementation) view returns(bytes32 initCodeHash)
func (_PlugFactory *PlugFactoryCallerSession) InitCodeHash(implementation common.Address) ([32]byte, error) {
	return _PlugFactory.Contract.InitCodeHash(&_PlugFactory.CallOpts, implementation)
}

// Deploy is a paid mutator transaction binding the contract method 0x00774360.
//
// Solidity: function deploy(bytes salt) payable returns(bool alreadyDeployed, address socketAddress)
func (_PlugFactory *PlugFactoryTransactor) Deploy(opts *bind.TransactOpts, salt []byte) (*types.Transaction, error) {
	return _PlugFactory.contract.Transact(opts, "deploy", salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x00774360.
//
// Solidity: function deploy(bytes salt) payable returns(bool alreadyDeployed, address socketAddress)
func (_PlugFactory *PlugFactorySession) Deploy(salt []byte) (*types.Transaction, error) {
	return _PlugFactory.Contract.Deploy(&_PlugFactory.TransactOpts, salt)
}

// Deploy is a paid mutator transaction binding the contract method 0x00774360.
//
// Solidity: function deploy(bytes salt) payable returns(bool alreadyDeployed, address socketAddress)
func (_PlugFactory *PlugFactoryTransactorSession) Deploy(salt []byte) (*types.Transaction, error) {
	return _PlugFactory.Contract.Deploy(&_PlugFactory.TransactOpts, salt)
}

// PlugFactorySocketDeployedIterator is returned from FilterSocketDeployed and is used to iterate over the raw logs and unpacked data for SocketDeployed events raised by the PlugFactory contract.
type PlugFactorySocketDeployedIterator struct {
	Event *PlugFactorySocketDeployed // Event containing the contract specifics and raw log

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
func (it *PlugFactorySocketDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlugFactorySocketDeployed)
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
		it.Event = new(PlugFactorySocketDeployed)
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
func (it *PlugFactorySocketDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlugFactorySocketDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlugFactorySocketDeployed represents a SocketDeployed event raised by the PlugFactory contract.
type PlugFactorySocketDeployed struct {
	Implementation common.Address
	Vault          common.Address
	Salt           [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSocketDeployed is a free log retrieval operation binding the contract event 0x32c7564be3bbeabe66360c08d019367b0d744fcb948046d92552c00c9743dd10.
//
// Solidity: event SocketDeployed(address indexed implementation, address indexed vault, bytes32 salt)
func (_PlugFactory *PlugFactoryFilterer) FilterSocketDeployed(opts *bind.FilterOpts, implementation []common.Address, vault []common.Address) (*PlugFactorySocketDeployedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _PlugFactory.contract.FilterLogs(opts, "SocketDeployed", implementationRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return &PlugFactorySocketDeployedIterator{contract: _PlugFactory.contract, event: "SocketDeployed", logs: logs, sub: sub}, nil
}

// WatchSocketDeployed is a free log subscription operation binding the contract event 0x32c7564be3bbeabe66360c08d019367b0d744fcb948046d92552c00c9743dd10.
//
// Solidity: event SocketDeployed(address indexed implementation, address indexed vault, bytes32 salt)
func (_PlugFactory *PlugFactoryFilterer) WatchSocketDeployed(opts *bind.WatchOpts, sink chan<- *PlugFactorySocketDeployed, implementation []common.Address, vault []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}
	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _PlugFactory.contract.WatchLogs(opts, "SocketDeployed", implementationRule, vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlugFactorySocketDeployed)
				if err := _PlugFactory.contract.UnpackLog(event, "SocketDeployed", log); err != nil {
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

// ParseSocketDeployed is a log parse operation binding the contract event 0x32c7564be3bbeabe66360c08d019367b0d744fcb948046d92552c00c9743dd10.
//
// Solidity: event SocketDeployed(address indexed implementation, address indexed vault, bytes32 salt)
func (_PlugFactory *PlugFactoryFilterer) ParseSocketDeployed(log types.Log) (*PlugFactorySocketDeployed, error) {
	event := new(PlugFactorySocketDeployed)
	if err := _PlugFactory.contract.UnpackLog(event, "SocketDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
