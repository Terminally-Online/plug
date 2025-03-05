// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_math

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

// PlugMathMetaData contains all meta data concerning the PlugMath contract.
var PlugMathMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"add\",\"inputs\":[{\"name\":\"a\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"b\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"clamp\",\"inputs\":[{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"minValue\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxValue\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"divide\",\"inputs\":[{\"name\":\"a\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"b\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"max\",\"inputs\":[{\"name\":\"a\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"b\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"min\",\"inputs\":[{\"name\":\"a\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"b\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"modulo\",\"inputs\":[{\"name\":\"a\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"b\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"multiply\",\"inputs\":[{\"name\":\"a\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"b\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"power\",\"inputs\":[{\"name\":\"base\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"exponent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"subtract\",\"inputs\":[{\"name\":\"a\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"b\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"}]",
}

// PlugMathABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugMathMetaData.ABI instead.
var PlugMathABI = PlugMathMetaData.ABI

// PlugMath is an auto generated Go binding around an Ethereum contract.
type PlugMath struct {
	PlugMathCaller     // Read-only binding to the contract
	PlugMathTransactor // Write-only binding to the contract
	PlugMathFilterer   // Log filterer for contract events
}

// PlugMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugMathSession struct {
	Contract     *PlugMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugMathCallerSession struct {
	Contract *PlugMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PlugMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugMathTransactorSession struct {
	Contract     *PlugMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PlugMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugMathRaw struct {
	Contract *PlugMath // Generic contract binding to access the raw methods on
}

// PlugMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugMathCallerRaw struct {
	Contract *PlugMathCaller // Generic read-only contract binding to access the raw methods on
}

// PlugMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugMathTransactorRaw struct {
	Contract *PlugMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugMath creates a new instance of PlugMath, bound to a specific deployed contract.
func NewPlugMath(address common.Address, backend bind.ContractBackend) (*PlugMath, error) {
	contract, err := bindPlugMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugMath{PlugMathCaller: PlugMathCaller{contract: contract}, PlugMathTransactor: PlugMathTransactor{contract: contract}, PlugMathFilterer: PlugMathFilterer{contract: contract}}, nil
}

// NewPlugMathCaller creates a new read-only instance of PlugMath, bound to a specific deployed contract.
func NewPlugMathCaller(address common.Address, caller bind.ContractCaller) (*PlugMathCaller, error) {
	contract, err := bindPlugMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugMathCaller{contract: contract}, nil
}

// NewPlugMathTransactor creates a new write-only instance of PlugMath, bound to a specific deployed contract.
func NewPlugMathTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugMathTransactor, error) {
	contract, err := bindPlugMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugMathTransactor{contract: contract}, nil
}

// NewPlugMathFilterer creates a new log filterer instance of PlugMath, bound to a specific deployed contract.
func NewPlugMathFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugMathFilterer, error) {
	contract, err := bindPlugMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugMathFilterer{contract: contract}, nil
}

// bindPlugMath binds a generic wrapper to an already deployed contract.
func bindPlugMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugMath *PlugMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugMath.Contract.PlugMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugMath *PlugMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugMath.Contract.PlugMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugMath *PlugMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugMath.Contract.PlugMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugMath *PlugMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugMath *PlugMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugMath *PlugMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugMath.Contract.contract.Transact(opts, method, params...)
}

// Add is a free data retrieval call binding the contract method 0xa5f3c23b.
//
// Solidity: function add(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Add(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "add", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0xa5f3c23b.
//
// Solidity: function add(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Add(&_PlugMath.CallOpts, a, b)
}

// Add is a free data retrieval call binding the contract method 0xa5f3c23b.
//
// Solidity: function add(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Add(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Add(&_PlugMath.CallOpts, a, b)
}

// Clamp is a free data retrieval call binding the contract method 0x7b8d0f0c.
//
// Solidity: function clamp(int256 value, int256 minValue, int256 maxValue) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Clamp(opts *bind.CallOpts, value *big.Int, minValue *big.Int, maxValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "clamp", value, minValue, maxValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clamp is a free data retrieval call binding the contract method 0x7b8d0f0c.
//
// Solidity: function clamp(int256 value, int256 minValue, int256 maxValue) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Clamp(value *big.Int, minValue *big.Int, maxValue *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Clamp(&_PlugMath.CallOpts, value, minValue, maxValue)
}

// Clamp is a free data retrieval call binding the contract method 0x7b8d0f0c.
//
// Solidity: function clamp(int256 value, int256 minValue, int256 maxValue) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Clamp(value *big.Int, minValue *big.Int, maxValue *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Clamp(&_PlugMath.CallOpts, value, minValue, maxValue)
}

// Divide is a free data retrieval call binding the contract method 0xf5984236.
//
// Solidity: function divide(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Divide(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "divide", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Divide is a free data retrieval call binding the contract method 0xf5984236.
//
// Solidity: function divide(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Divide(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Divide(&_PlugMath.CallOpts, a, b)
}

// Divide is a free data retrieval call binding the contract method 0xf5984236.
//
// Solidity: function divide(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Divide(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Divide(&_PlugMath.CallOpts, a, b)
}

// Max is a free data retrieval call binding the contract method 0x81fe5786.
//
// Solidity: function max(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Max(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "max", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Max is a free data retrieval call binding the contract method 0x81fe5786.
//
// Solidity: function max(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Max(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Max(&_PlugMath.CallOpts, a, b)
}

// Max is a free data retrieval call binding the contract method 0x81fe5786.
//
// Solidity: function max(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Max(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Max(&_PlugMath.CallOpts, a, b)
}

// Min is a free data retrieval call binding the contract method 0x29aa9cbe.
//
// Solidity: function min(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Min(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "min", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Min is a free data retrieval call binding the contract method 0x29aa9cbe.
//
// Solidity: function min(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Min(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Min(&_PlugMath.CallOpts, a, b)
}

// Min is a free data retrieval call binding the contract method 0x29aa9cbe.
//
// Solidity: function min(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Min(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Min(&_PlugMath.CallOpts, a, b)
}

// Modulo is a free data retrieval call binding the contract method 0x121c3169.
//
// Solidity: function modulo(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Modulo(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "modulo", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Modulo is a free data retrieval call binding the contract method 0x121c3169.
//
// Solidity: function modulo(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Modulo(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Modulo(&_PlugMath.CallOpts, a, b)
}

// Modulo is a free data retrieval call binding the contract method 0x121c3169.
//
// Solidity: function modulo(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Modulo(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Modulo(&_PlugMath.CallOpts, a, b)
}

// Multiply is a free data retrieval call binding the contract method 0x3c4308a8.
//
// Solidity: function multiply(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Multiply(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "multiply", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiply is a free data retrieval call binding the contract method 0x3c4308a8.
//
// Solidity: function multiply(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Multiply(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Multiply(&_PlugMath.CallOpts, a, b)
}

// Multiply is a free data retrieval call binding the contract method 0x3c4308a8.
//
// Solidity: function multiply(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Multiply(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Multiply(&_PlugMath.CallOpts, a, b)
}

// Power is a free data retrieval call binding the contract method 0x529f36a9.
//
// Solidity: function power(int256 base, uint256 exponent) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Power(opts *bind.CallOpts, base *big.Int, exponent *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "power", base, exponent)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Power is a free data retrieval call binding the contract method 0x529f36a9.
//
// Solidity: function power(int256 base, uint256 exponent) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Power(base *big.Int, exponent *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Power(&_PlugMath.CallOpts, base, exponent)
}

// Power is a free data retrieval call binding the contract method 0x529f36a9.
//
// Solidity: function power(int256 base, uint256 exponent) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Power(base *big.Int, exponent *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Power(&_PlugMath.CallOpts, base, exponent)
}

// Subtract is a free data retrieval call binding the contract method 0xb93ea812.
//
// Solidity: function subtract(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCaller) Subtract(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "subtract", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Subtract is a free data retrieval call binding the contract method 0xb93ea812.
//
// Solidity: function subtract(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathSession) Subtract(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Subtract(&_PlugMath.CallOpts, a, b)
}

// Subtract is a free data retrieval call binding the contract method 0xb93ea812.
//
// Solidity: function subtract(int256 a, int256 b) pure returns(int256 result)
func (_PlugMath *PlugMathCallerSession) Subtract(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Subtract(&_PlugMath.CallOpts, a, b)
}
