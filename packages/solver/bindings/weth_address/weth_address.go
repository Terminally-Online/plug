// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package weth_address

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

// WethAddressMetaData contains all meta data concerning the WethAddress contract.
var WethAddressMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"}]",
}

// WethAddressABI is the input ABI used to generate the binding from.
// Deprecated: Use WethAddressMetaData.ABI instead.
var WethAddressABI = WethAddressMetaData.ABI

// WethAddress is an auto generated Go binding around an Ethereum contract.
type WethAddress struct {
	WethAddressCaller     // Read-only binding to the contract
	WethAddressTransactor // Write-only binding to the contract
	WethAddressFilterer   // Log filterer for contract events
}

// WethAddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type WethAddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethAddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WethAddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethAddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WethAddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WethAddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WethAddressSession struct {
	Contract     *WethAddress      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WethAddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WethAddressCallerSession struct {
	Contract *WethAddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WethAddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WethAddressTransactorSession struct {
	Contract     *WethAddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WethAddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type WethAddressRaw struct {
	Contract *WethAddress // Generic contract binding to access the raw methods on
}

// WethAddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WethAddressCallerRaw struct {
	Contract *WethAddressCaller // Generic read-only contract binding to access the raw methods on
}

// WethAddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WethAddressTransactorRaw struct {
	Contract *WethAddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWethAddress creates a new instance of WethAddress, bound to a specific deployed contract.
func NewWethAddress(address common.Address, backend bind.ContractBackend) (*WethAddress, error) {
	contract, err := bindWethAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WethAddress{WethAddressCaller: WethAddressCaller{contract: contract}, WethAddressTransactor: WethAddressTransactor{contract: contract}, WethAddressFilterer: WethAddressFilterer{contract: contract}}, nil
}

// NewWethAddressCaller creates a new read-only instance of WethAddress, bound to a specific deployed contract.
func NewWethAddressCaller(address common.Address, caller bind.ContractCaller) (*WethAddressCaller, error) {
	contract, err := bindWethAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WethAddressCaller{contract: contract}, nil
}

// NewWethAddressTransactor creates a new write-only instance of WethAddress, bound to a specific deployed contract.
func NewWethAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*WethAddressTransactor, error) {
	contract, err := bindWethAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WethAddressTransactor{contract: contract}, nil
}

// NewWethAddressFilterer creates a new log filterer instance of WethAddress, bound to a specific deployed contract.
func NewWethAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*WethAddressFilterer, error) {
	contract, err := bindWethAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WethAddressFilterer{contract: contract}, nil
}

// bindWethAddress binds a generic wrapper to an already deployed contract.
func bindWethAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WethAddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WethAddress *WethAddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WethAddress.Contract.WethAddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WethAddress *WethAddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WethAddress.Contract.WethAddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WethAddress *WethAddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WethAddress.Contract.WethAddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WethAddress *WethAddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WethAddress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WethAddress *WethAddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WethAddress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WethAddress *WethAddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WethAddress.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WethAddress *WethAddressCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WethAddress.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WethAddress *WethAddressSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WethAddress.Contract.Allowance(&_WethAddress.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_WethAddress *WethAddressCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _WethAddress.Contract.Allowance(&_WethAddress.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WethAddress *WethAddressCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WethAddress.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WethAddress *WethAddressSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WethAddress.Contract.BalanceOf(&_WethAddress.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_WethAddress *WethAddressCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _WethAddress.Contract.BalanceOf(&_WethAddress.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WethAddress *WethAddressCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WethAddress.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WethAddress *WethAddressSession) Decimals() (uint8, error) {
	return _WethAddress.Contract.Decimals(&_WethAddress.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WethAddress *WethAddressCallerSession) Decimals() (uint8, error) {
	return _WethAddress.Contract.Decimals(&_WethAddress.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WethAddress *WethAddressCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WethAddress.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WethAddress *WethAddressSession) Name() (string, error) {
	return _WethAddress.Contract.Name(&_WethAddress.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WethAddress *WethAddressCallerSession) Name() (string, error) {
	return _WethAddress.Contract.Name(&_WethAddress.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WethAddress *WethAddressCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WethAddress.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WethAddress *WethAddressSession) Symbol() (string, error) {
	return _WethAddress.Contract.Symbol(&_WethAddress.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WethAddress *WethAddressCallerSession) Symbol() (string, error) {
	return _WethAddress.Contract.Symbol(&_WethAddress.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WethAddress *WethAddressCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WethAddress.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WethAddress *WethAddressSession) TotalSupply() (*big.Int, error) {
	return _WethAddress.Contract.TotalSupply(&_WethAddress.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WethAddress *WethAddressCallerSession) TotalSupply() (*big.Int, error) {
	return _WethAddress.Contract.TotalSupply(&_WethAddress.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WethAddress *WethAddressTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WethAddress *WethAddressSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.Contract.Approve(&_WethAddress.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address guy, uint256 wad) returns(bool)
func (_WethAddress *WethAddressTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.Contract.Approve(&_WethAddress.TransactOpts, guy, wad)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WethAddress *WethAddressTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WethAddress.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WethAddress *WethAddressSession) Deposit() (*types.Transaction, error) {
	return _WethAddress.Contract.Deposit(&_WethAddress.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WethAddress *WethAddressTransactorSession) Deposit() (*types.Transaction, error) {
	return _WethAddress.Contract.Deposit(&_WethAddress.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WethAddress *WethAddressTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WethAddress *WethAddressSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.Contract.Transfer(&_WethAddress.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 wad) returns(bool)
func (_WethAddress *WethAddressTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.Contract.Transfer(&_WethAddress.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WethAddress *WethAddressTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WethAddress *WethAddressSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.Contract.TransferFrom(&_WethAddress.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 wad) returns(bool)
func (_WethAddress *WethAddressTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.Contract.TransferFrom(&_WethAddress.TransactOpts, src, dst, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WethAddress *WethAddressTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WethAddress *WethAddressSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.Contract.Withdraw(&_WethAddress.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WethAddress *WethAddressTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WethAddress.Contract.Withdraw(&_WethAddress.TransactOpts, wad)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WethAddress *WethAddressTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WethAddress.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WethAddress *WethAddressSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WethAddress.Contract.Fallback(&_WethAddress.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WethAddress *WethAddressTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WethAddress.Contract.Fallback(&_WethAddress.TransactOpts, calldata)
}

// WethAddressApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WethAddress contract.
type WethAddressApprovalIterator struct {
	Event *WethAddressApproval // Event containing the contract specifics and raw log

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
func (it *WethAddressApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethAddressApproval)
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
		it.Event = new(WethAddressApproval)
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
func (it *WethAddressApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethAddressApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethAddressApproval represents a Approval event raised by the WethAddress contract.
type WethAddressApproval struct {
	Src common.Address
	Guy common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WethAddress *WethAddressFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, guy []common.Address) (*WethAddressApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WethAddress.contract.FilterLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return &WethAddressApprovalIterator{contract: _WethAddress.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WethAddress *WethAddressFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WethAddressApproval, src []common.Address, guy []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var guyRule []interface{}
	for _, guyItem := range guy {
		guyRule = append(guyRule, guyItem)
	}

	logs, sub, err := _WethAddress.contract.WatchLogs(opts, "Approval", srcRule, guyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethAddressApproval)
				if err := _WethAddress.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed guy, uint256 wad)
func (_WethAddress *WethAddressFilterer) ParseApproval(log types.Log) (*WethAddressApproval, error) {
	event := new(WethAddressApproval)
	if err := _WethAddress.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethAddressDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the WethAddress contract.
type WethAddressDepositIterator struct {
	Event *WethAddressDeposit // Event containing the contract specifics and raw log

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
func (it *WethAddressDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethAddressDeposit)
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
		it.Event = new(WethAddressDeposit)
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
func (it *WethAddressDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethAddressDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethAddressDeposit represents a Deposit event raised by the WethAddress contract.
type WethAddressDeposit struct {
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WethAddress *WethAddressFilterer) FilterDeposit(opts *bind.FilterOpts, dst []common.Address) (*WethAddressDepositIterator, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WethAddress.contract.FilterLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return &WethAddressDepositIterator{contract: _WethAddress.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WethAddress *WethAddressFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *WethAddressDeposit, dst []common.Address) (event.Subscription, error) {

	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WethAddress.contract.WatchLogs(opts, "Deposit", dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethAddressDeposit)
				if err := _WethAddress.contract.UnpackLog(event, "Deposit", log); err != nil {
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
// Solidity: event Deposit(address indexed dst, uint256 wad)
func (_WethAddress *WethAddressFilterer) ParseDeposit(log types.Log) (*WethAddressDeposit, error) {
	event := new(WethAddressDeposit)
	if err := _WethAddress.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethAddressTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WethAddress contract.
type WethAddressTransferIterator struct {
	Event *WethAddressTransfer // Event containing the contract specifics and raw log

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
func (it *WethAddressTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethAddressTransfer)
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
		it.Event = new(WethAddressTransfer)
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
func (it *WethAddressTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethAddressTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethAddressTransfer represents a Transfer event raised by the WethAddress contract.
type WethAddressTransfer struct {
	Src common.Address
	Dst common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WethAddress *WethAddressFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*WethAddressTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WethAddress.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &WethAddressTransferIterator{contract: _WethAddress.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WethAddress *WethAddressFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WethAddressTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _WethAddress.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethAddressTransfer)
				if err := _WethAddress.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 wad)
func (_WethAddress *WethAddressFilterer) ParseTransfer(log types.Log) (*WethAddressTransfer, error) {
	event := new(WethAddressTransfer)
	if err := _WethAddress.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WethAddressWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the WethAddress contract.
type WethAddressWithdrawalIterator struct {
	Event *WethAddressWithdrawal // Event containing the contract specifics and raw log

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
func (it *WethAddressWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WethAddressWithdrawal)
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
		it.Event = new(WethAddressWithdrawal)
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
func (it *WethAddressWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WethAddressWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WethAddressWithdrawal represents a Withdrawal event raised by the WethAddress contract.
type WethAddressWithdrawal struct {
	Src common.Address
	Wad *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WethAddress *WethAddressFilterer) FilterWithdrawal(opts *bind.FilterOpts, src []common.Address) (*WethAddressWithdrawalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WethAddress.contract.FilterLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return &WethAddressWithdrawalIterator{contract: _WethAddress.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WethAddress *WethAddressFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *WethAddressWithdrawal, src []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}

	logs, sub, err := _WethAddress.contract.WatchLogs(opts, "Withdrawal", srcRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WethAddressWithdrawal)
				if err := _WethAddress.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed src, uint256 wad)
func (_WethAddress *WethAddressFilterer) ParseWithdrawal(log types.Log) (*WethAddressWithdrawal, error) {
	event := new(WethAddressWithdrawal)
	if err := _WethAddress.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
