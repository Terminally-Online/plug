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
	ABI: "[{\"type\":\"function\",\"name\":\"isAnd\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isBetween\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isEqual\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isGreaterThan\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isGreaterThanOrEqual\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isImplies\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isLessThan\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isLessThanOrEqual\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isNand\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isNor\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isNot\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isNotEqual\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isOr\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isXor\",\"inputs\":[{\"name\":\"a\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"b\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"}]",
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

// IsEqual is a free data retrieval call binding the contract method 0x3ced224a.
//
// Solidity: function isEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsEqual(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isEqual", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEqual is a free data retrieval call binding the contract method 0x3ced224a.
//
// Solidity: function isEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsEqual(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsEqual(&_PlugBoolean.CallOpts, a, b)
}

// IsEqual is a free data retrieval call binding the contract method 0x3ced224a.
//
// Solidity: function isEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsEqual(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsEqual(&_PlugBoolean.CallOpts, a, b)
}

// IsGreaterThan is a free data retrieval call binding the contract method 0xdcf6a592.
//
// Solidity: function isGreaterThan(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsGreaterThan(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isGreaterThan", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGreaterThan is a free data retrieval call binding the contract method 0xdcf6a592.
//
// Solidity: function isGreaterThan(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsGreaterThan(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsGreaterThan(&_PlugBoolean.CallOpts, a, b)
}

// IsGreaterThan is a free data retrieval call binding the contract method 0xdcf6a592.
//
// Solidity: function isGreaterThan(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsGreaterThan(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsGreaterThan(&_PlugBoolean.CallOpts, a, b)
}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x09aa2a75.
//
// Solidity: function isGreaterThanOrEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsGreaterThanOrEqual(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isGreaterThanOrEqual", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x09aa2a75.
//
// Solidity: function isGreaterThanOrEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsGreaterThanOrEqual(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsGreaterThanOrEqual(&_PlugBoolean.CallOpts, a, b)
}

// IsGreaterThanOrEqual is a free data retrieval call binding the contract method 0x09aa2a75.
//
// Solidity: function isGreaterThanOrEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsGreaterThanOrEqual(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsGreaterThanOrEqual(&_PlugBoolean.CallOpts, a, b)
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
// Solidity: function isLessThan(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsLessThan(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isLessThan", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLessThan is a free data retrieval call binding the contract method 0xe9970b6c.
//
// Solidity: function isLessThan(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsLessThan(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsLessThan(&_PlugBoolean.CallOpts, a, b)
}

// IsLessThan is a free data retrieval call binding the contract method 0xe9970b6c.
//
// Solidity: function isLessThan(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsLessThan(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsLessThan(&_PlugBoolean.CallOpts, a, b)
}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x5f3e9a58.
//
// Solidity: function isLessThanOrEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsLessThanOrEqual(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isLessThanOrEqual", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x5f3e9a58.
//
// Solidity: function isLessThanOrEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsLessThanOrEqual(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsLessThanOrEqual(&_PlugBoolean.CallOpts, a, b)
}

// IsLessThanOrEqual is a free data retrieval call binding the contract method 0x5f3e9a58.
//
// Solidity: function isLessThanOrEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsLessThanOrEqual(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsLessThanOrEqual(&_PlugBoolean.CallOpts, a, b)
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
// Solidity: function isNotEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCaller) IsNotEqual(opts *bind.CallOpts, a *big.Int, b *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugBoolean.contract.Call(opts, &out, "isNotEqual", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNotEqual is a free data retrieval call binding the contract method 0x7de3e9c1.
//
// Solidity: function isNotEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanSession) IsNotEqual(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsNotEqual(&_PlugBoolean.CallOpts, a, b)
}

// IsNotEqual is a free data retrieval call binding the contract method 0x7de3e9c1.
//
// Solidity: function isNotEqual(uint256 a, uint256 b) pure returns(bool result)
func (_PlugBoolean *PlugBooleanCallerSession) IsNotEqual(a *big.Int, b *big.Int) (bool, error) {
	return _PlugBoolean.Contract.IsNotEqual(&_PlugBoolean.CallOpts, a, b)
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
