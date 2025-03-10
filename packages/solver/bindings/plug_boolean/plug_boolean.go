// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_boolean

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

// PlugBooleanMetaData contains all meta data concerning the PlugBoolean contract.
var PlugBooleanMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"isAfterTime\",\"inputs\":[{\"name\":\"time\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isAnd\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isBeforeTime\",\"inputs\":[{\"name\":\"time\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isBetween\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isBetweenTimes\",\"inputs\":[{\"name\":\"time\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isEqual\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isFalse\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isGreaterThan\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isGreaterThanOrEqual\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isImplies\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isLessThan\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isLessThanOrEqual\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isNand\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isNor\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isNot\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isNotEqual\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isOr\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isSameDay\",\"inputs\":[{\"name\":\"timestamp1\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp2\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isTrue\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isWeekday\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isWeekend\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isXor\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"}]",
}

// PlugBooleanABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugBooleanMetaData.ABI instead.
var PlugBooleanABI = PlugBooleanMetaData.ABI

// PlugBoolean is an auto generated Go binding around an Ethereum contract.
type PlugBoolean struct {
	PlugBooleanCaller     // Read-only binding to the contract
	PlugBooleanTransactor // Write-only binding to the contract
	PlugBooleanFilterer   // Log filterer for contract events
}

// PlugBooleanCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugBooleanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugBooleanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugBooleanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugBooleanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugBooleanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugBooleanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugBooleanSession struct {
	Contract     *PlugBoolean      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugBooleanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugBooleanCallerSession struct {
	Contract *PlugBooleanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PlugBooleanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugBooleanTransactorSession struct {
	Contract     *PlugBooleanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PlugBooleanRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugBooleanRaw struct {
	Contract *PlugBoolean // Generic contract binding to access the raw methods on
}

// PlugBooleanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugBooleanCallerRaw struct {
	Contract *PlugBooleanCaller // Generic read-only contract binding to access the raw methods on
}

// PlugBooleanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugBooleanTransactorRaw struct {
	Contract *PlugBooleanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugBoolean creates a new instance of PlugBoolean, bound to a specific deployed contract.
func NewPlugBoolean(address common.Address, backend bind.ContractBackend) (*PlugBoolean, error) {
	contract, err := bindPlugBoolean(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugBoolean{PlugBooleanCaller: PlugBooleanCaller{contract: contract}, PlugBooleanTransactor: PlugBooleanTransactor{contract: contract}, PlugBooleanFilterer: PlugBooleanFilterer{contract: contract}}, nil
}

// NewPlugBooleanCaller creates a new read-only instance of PlugBoolean, bound to a specific deployed contract.
func NewPlugBooleanCaller(address common.Address, caller bind.ContractCaller) (*PlugBooleanCaller, error) {
	contract, err := bindPlugBoolean(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugBooleanCaller{contract: contract}, nil
}

// NewPlugBooleanTransactor creates a new write-only instance of PlugBoolean, bound to a specific deployed contract.
func NewPlugBooleanTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugBooleanTransactor, error) {
	contract, err := bindPlugBoolean(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugBooleanTransactor{contract: contract}, nil
}

// NewPlugBooleanFilterer creates a new log filterer instance of PlugBoolean, bound to a specific deployed contract.
func NewPlugBooleanFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugBooleanFilterer, error) {
	contract, err := bindPlugBoolean(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugBooleanFilterer{contract: contract}, nil
}

// bindPlugBoolean binds a generic wrapper to an already deployed contract.
func bindPlugBoolean(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugBoolean *PlugBooleanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugBoolean.Contract.PlugBooleanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugBoolean *PlugBooleanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugBoolean.Contract.PlugBooleanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugBoolean *PlugBooleanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugBoolean.Contract.PlugBooleanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugBoolean *PlugBooleanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugBoolean.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugBoolean *PlugBooleanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugBoolean.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugBoolean *PlugBooleanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugBoolean.Contract.contract.Transact(opts, method, params...)
}

// IsAfterTime is a free data retrieval call binding the contract method 0xb61536ae.
//
// Solidity: function isAfterTime(uint256 time, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsAfterTime(opts *bind.CallOpts, time *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isAfterTime", time, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAfterTime is a free data retrieval call binding the contract method 0xb61536ae.
//
// Solidity: function isAfterTime(uint256 time, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsAfterTime(time *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsAfterTime(&_PlugBoolean.CallOpts, time, threshold)
}

// IsAfterTime is a free data retrieval call binding the contract method 0xb61536ae.
//
// Solidity: function isAfterTime(uint256 time, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsAfterTime(time *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsAfterTime(&_PlugBoolean.CallOpts, time, threshold)
}

// IsAnd is a free data retrieval call binding the contract method 0x1d235a52.
//
// Solidity: function isAnd(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsAnd(opts *bind.CallOpts, a bool, b bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isAnd", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAnd is a free data retrieval call binding the contract method 0x1d235a52.
//
// Solidity: function isAnd(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsAnd(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsAnd(&_PlugBoolean.CallOpts, a, b)
}

// IsAnd is a free data retrieval call binding the contract method 0x1d235a52.
//
// Solidity: function isAnd(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsAnd(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsAnd(&_PlugBoolean.CallOpts, a, b)
}

// IsBeforeTime is a free data retrieval call binding the contract method 0xb2df673f.
//
// Solidity: function isBeforeTime(uint256 time, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsBeforeTime(opts *bind.CallOpts, time *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isBeforeTime", time, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBeforeTime is a free data retrieval call binding the contract method 0xb2df673f.
//
// Solidity: function isBeforeTime(uint256 time, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsBeforeTime(time *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsBeforeTime(&_PlugBoolean.CallOpts, time, threshold)
}

// IsBeforeTime is a free data retrieval call binding the contract method 0xb2df673f.
//
// Solidity: function isBeforeTime(uint256 time, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsBeforeTime(time *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsBeforeTime(&_PlugBoolean.CallOpts, time, threshold)
}

// IsBetween is a free data retrieval call binding the contract method 0x95f4ac38.
//
// Solidity: function isBetween(uint256 value, uint256 min, uint256 max) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsBetween(opts *bind.CallOpts, value *big.Int, min *big.Int, max *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isBetween", value, min, max)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBetween is a free data retrieval call binding the contract method 0x95f4ac38.
//
// Solidity: function isBetween(uint256 value, uint256 min, uint256 max) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsBetween(value *big.Int, min *big.Int, max *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsBetween(&_PlugBoolean.CallOpts, value, min, max)
}

// IsBetween is a free data retrieval call binding the contract method 0x95f4ac38.
//
// Solidity: function isBetween(uint256 value, uint256 min, uint256 max) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsBetween(value *big.Int, min *big.Int, max *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsBetween(&_PlugBoolean.CallOpts, value, min, max)
}

// IsBetweenTimes is a free data retrieval call binding the contract method 0x6d52497a.
//
// Solidity: function isBetweenTimes(uint256 time, uint256 start, uint256 end) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsBetweenTimes(opts *bind.CallOpts, time *big.Int, start *big.Int, end *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isBetweenTimes", time, start, end)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBetweenTimes is a free data retrieval call binding the contract method 0x6d52497a.
//
// Solidity: function isBetweenTimes(uint256 time, uint256 start, uint256 end) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsBetweenTimes(time *big.Int, start *big.Int, end *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsBetweenTimes(&_PlugBoolean.CallOpts, time, start, end)
}

// IsBetweenTimes is a free data retrieval call binding the contract method 0x6d52497a.
//
// Solidity: function isBetweenTimes(uint256 time, uint256 start, uint256 end) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsBetweenTimes(time *big.Int, start *big.Int, end *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsBetweenTimes(&_PlugBoolean.CallOpts, time, start, end)
}

// IsEqual is a free data retrieval call binding the contract method 0x3ced224a.
//
// Solidity: function isEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsEqual(opts *bind.CallOpts, value *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isEqual", value, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEqual is a free data retrieval call binding the contract method 0x3ced224a.
//
// Solidity: function isEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsEqual(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsEqual(&_PlugBoolean.CallOpts, value, threshold)
}

// IsEqual is a free data retrieval call binding the contract method 0x3ced224a.
//
// Solidity: function isEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsEqual(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsEqual(&_PlugBoolean.CallOpts, value, threshold)
}

// IsFalse is a free data retrieval call binding the contract method 0x410dbb8b.
//
// Solidity: function isFalse(bool value) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsFalse(opts *bind.CallOpts, value bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isFalse", value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFalse is a free data retrieval call binding the contract method 0x410dbb8b.
//
// Solidity: function isFalse(bool value) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsFalse(value bool) (bool, error) {
	return _PlugBoolean.Contract.IsFalse(&_PlugBoolean.CallOpts, value)
}

// IsFalse is a free data retrieval call binding the contract method 0x410dbb8b.
//
// Solidity: function isFalse(bool value) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsFalse(value bool) (bool, error) {
	return _PlugBoolean.Contract.IsFalse(&_PlugBoolean.CallOpts, value)
}

// IsGreaterThan is a free data retrieval call binding the contract method 0xdcf6a592.
//
// Solidity: function isGreaterThan(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsGreaterThan(opts *bind.CallOpts, value *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isGreaterThan", value, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGreaterThan is a free data retrieval call binding the contract method 0xdcf6a592.
//
// Solidity: function isGreaterThan(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsGreaterThan(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsGreaterThan(&_PlugBoolean.CallOpts, value, threshold)
}

// IsGreaterThan is a free data retrieval call binding the contract method 0xdcf6a592.
//
// Solidity: function isGreaterThan(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsGreaterThan(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsGreaterThan(&_PlugBoolean.CallOpts, value, threshold)
}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x09aa2a75.
//
// Solidity: function isGreaterThanOrEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsGreaterThanOrEqual(opts *bind.CallOpts, value *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isGreaterThanOrEqual", value, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x09aa2a75.
//
// Solidity: function isGreaterThanOrEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsGreaterThanOrEqual(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsGreaterThanOrEqual(&_PlugBoolean.CallOpts, value, threshold)
}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x09aa2a75.
//
// Solidity: function isGreaterThanOrEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsGreaterThanOrEqual(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsGreaterThanOrEqual(&_PlugBoolean.CallOpts, value, threshold)
}

// IsImplies is a free data retrieval call binding the contract method 0x3eaea839.
//
// Solidity: function isImplies(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsImplies(opts *bind.CallOpts, a bool, b bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isImplies", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsImplies is a free data retrieval call binding the contract method 0x3eaea839.
//
// Solidity: function isImplies(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsImplies(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsImplies(&_PlugBoolean.CallOpts, a, b)
}

// IsImplies is a free data retrieval call binding the contract method 0x3eaea839.
//
// Solidity: function isImplies(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsImplies(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsImplies(&_PlugBoolean.CallOpts, a, b)
}

// IsLessThan is a free data retrieval call binding the contract method 0xe9970b6c.
//
// Solidity: function isLessThan(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsLessThan(opts *bind.CallOpts, value *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isLessThan", value, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLessThan is a free data retrieval call binding the contract method 0xe9970b6c.
//
// Solidity: function isLessThan(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsLessThan(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsLessThan(&_PlugBoolean.CallOpts, value, threshold)
}

// IsLessThan is a free data retrieval call binding the contract method 0xe9970b6c.
//
// Solidity: function isLessThan(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsLessThan(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsLessThan(&_PlugBoolean.CallOpts, value, threshold)
}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x5f3e9a58.
//
// Solidity: function isLessThanOrEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsLessThanOrEqual(opts *bind.CallOpts, value *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isLessThanOrEqual", value, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x5f3e9a58.
//
// Solidity: function isLessThanOrEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsLessThanOrEqual(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsLessThanOrEqual(&_PlugBoolean.CallOpts, value, threshold)
}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x5f3e9a58.
//
// Solidity: function isLessThanOrEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsLessThanOrEqual(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsLessThanOrEqual(&_PlugBoolean.CallOpts, value, threshold)
}

// IsNand is a free data retrieval call binding the contract method 0xaf4256ba.
//
// Solidity: function isNand(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsNand(opts *bind.CallOpts, a bool, b bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isNand", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNand is a free data retrieval call binding the contract method 0xaf4256ba.
//
// Solidity: function isNand(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsNand(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsNand(&_PlugBoolean.CallOpts, a, b)
}

// IsNand is a free data retrieval call binding the contract method 0xaf4256ba.
//
// Solidity: function isNand(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsNand(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsNand(&_PlugBoolean.CallOpts, a, b)
}

// IsNor is a free data retrieval call binding the contract method 0x5a1389d7.
//
// Solidity: function isNor(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsNor(opts *bind.CallOpts, a bool, b bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isNor", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNor is a free data retrieval call binding the contract method 0x5a1389d7.
//
// Solidity: function isNor(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsNor(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsNor(&_PlugBoolean.CallOpts, a, b)
}

// IsNor is a free data retrieval call binding the contract method 0x5a1389d7.
//
// Solidity: function isNor(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsNor(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsNor(&_PlugBoolean.CallOpts, a, b)
}

// IsNot is a free data retrieval call binding the contract method 0x1c8d23cd.
//
// Solidity: function isNot(bool a) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsNot(opts *bind.CallOpts, a bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isNot", a)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNot is a free data retrieval call binding the contract method 0x1c8d23cd.
//
// Solidity: function isNot(bool a) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsNot(a bool) (bool, error) {
	return _PlugBoolean.Contract.IsNot(&_PlugBoolean.CallOpts, a)
}

// IsNot is a free data retrieval call binding the contract method 0x1c8d23cd.
//
// Solidity: function isNot(bool a) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsNot(a bool) (bool, error) {
	return _PlugBoolean.Contract.IsNot(&_PlugBoolean.CallOpts, a)
}

// IsNotEqual is a free data retrieval call binding the contract method 0x7de3e9c1.
//
// Solidity: function isNotEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsNotEqual(opts *bind.CallOpts, value *big.Int, threshold *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isNotEqual", value, threshold)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNotEqual is a free data retrieval call binding the contract method 0x7de3e9c1.
//
// Solidity: function isNotEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsNotEqual(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsNotEqual(&_PlugBoolean.CallOpts, value, threshold)
}

// IsNotEqual is a free data retrieval call binding the contract method 0x7de3e9c1.
//
// Solidity: function isNotEqual(uint256 value, uint256 threshold) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsNotEqual(value *big.Int, threshold *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsNotEqual(&_PlugBoolean.CallOpts, value, threshold)
}

// IsOr is a free data retrieval call binding the contract method 0x75bafee5.
//
// Solidity: function isOr(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsOr(opts *bind.CallOpts, a bool, b bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isOr", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOr is a free data retrieval call binding the contract method 0x75bafee5.
//
// Solidity: function isOr(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsOr(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsOr(&_PlugBoolean.CallOpts, a, b)
}

// IsOr is a free data retrieval call binding the contract method 0x75bafee5.
//
// Solidity: function isOr(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsOr(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsOr(&_PlugBoolean.CallOpts, a, b)
}

// IsSameDay is a free data retrieval call binding the contract method 0x4fa48a88.
//
// Solidity: function isSameDay(uint256 timestamp1, uint256 timestamp2) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsSameDay(opts *bind.CallOpts, timestamp1 *big.Int, timestamp2 *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isSameDay", timestamp1, timestamp2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSameDay is a free data retrieval call binding the contract method 0x4fa48a88.
//
// Solidity: function isSameDay(uint256 timestamp1, uint256 timestamp2) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsSameDay(timestamp1 *big.Int, timestamp2 *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsSameDay(&_PlugBoolean.CallOpts, timestamp1, timestamp2)
}

// IsSameDay is a free data retrieval call binding the contract method 0x4fa48a88.
//
// Solidity: function isSameDay(uint256 timestamp1, uint256 timestamp2) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsSameDay(timestamp1 *big.Int, timestamp2 *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsSameDay(&_PlugBoolean.CallOpts, timestamp1, timestamp2)
}

// IsTrue is a free data retrieval call binding the contract method 0x74f13070.
//
// Solidity: function isTrue(bool value) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsTrue(opts *bind.CallOpts, value bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isTrue", value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrue is a free data retrieval call binding the contract method 0x74f13070.
//
// Solidity: function isTrue(bool value) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsTrue(value bool) (bool, error) {
	return _PlugBoolean.Contract.IsTrue(&_PlugBoolean.CallOpts, value)
}

// IsTrue is a free data retrieval call binding the contract method 0x74f13070.
//
// Solidity: function isTrue(bool value) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsTrue(value bool) (bool, error) {
	return _PlugBoolean.Contract.IsTrue(&_PlugBoolean.CallOpts, value)
}

// IsWeekday is a free data retrieval call binding the contract method 0xe15d4b4e.
//
// Solidity: function isWeekday(uint256 timestamp) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsWeekday(opts *bind.CallOpts, timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isWeekday", timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWeekday is a free data retrieval call binding the contract method 0xe15d4b4e.
//
// Solidity: function isWeekday(uint256 timestamp) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsWeekday(timestamp *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsWeekday(&_PlugBoolean.CallOpts, timestamp)
}

// IsWeekday is a free data retrieval call binding the contract method 0xe15d4b4e.
//
// Solidity: function isWeekday(uint256 timestamp) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsWeekday(timestamp *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsWeekday(&_PlugBoolean.CallOpts, timestamp)
}

// IsWeekend is a free data retrieval call binding the contract method 0x41ff7087.
//
// Solidity: function isWeekend(uint256 timestamp) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsWeekend(opts *bind.CallOpts, timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isWeekend", timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWeekend is a free data retrieval call binding the contract method 0x41ff7087.
//
// Solidity: function isWeekend(uint256 timestamp) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsWeekend(timestamp *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsWeekend(&_PlugBoolean.CallOpts, timestamp)
}

// IsWeekend is a free data retrieval call binding the contract method 0x41ff7087.
//
// Solidity: function isWeekend(uint256 timestamp) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsWeekend(timestamp *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsWeekend(&_PlugBoolean.CallOpts, timestamp)
}

// IsXor is a free data retrieval call binding the contract method 0x3182925b.
//
// Solidity: function isXor(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsXor(opts *bind.CallOpts, a bool, b bool) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isXor", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsXor is a free data retrieval call binding the contract method 0x3182925b.
//
// Solidity: function isXor(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsXor(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsXor(&_PlugBoolean.CallOpts, a, b)
}

// IsXor is a free data retrieval call binding the contract method 0x3182925b.
//
// Solidity: function isXor(bool a, bool b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsXor(a bool, b bool) (bool, error) {
	return _PlugBoolean.Contract.IsXor(&_PlugBoolean.CallOpts, a, b)
}
