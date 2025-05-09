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
	ABI: "[{\"type\":\"function\",\"name\":\"add\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"divide\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"max\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"min\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"modulo\",\"inputs\":[{\"name\":\"a\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"b\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"multiply\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"power\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"subtract\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"}]",
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

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCaller) Add(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "add", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathSession) Add(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Add(&_PlugMath.CallOpts, x, y)
}

// Add is a free data retrieval call binding the contract method 0x771602f7.
//
// Solidity: function add(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCallerSession) Add(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Add(&_PlugMath.CallOpts, x, y)
}

// Divide is a free data retrieval call binding the contract method 0xf88e9fbf.
//
// Solidity: function divide(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCaller) Divide(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "divide", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Divide is a free data retrieval call binding the contract method 0xf88e9fbf.
//
// Solidity: function divide(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathSession) Divide(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Divide(&_PlugMath.CallOpts, x, y)
}

// Divide is a free data retrieval call binding the contract method 0xf88e9fbf.
//
// Solidity: function divide(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCallerSession) Divide(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Divide(&_PlugMath.CallOpts, x, y)
}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCaller) Max(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "max", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathSession) Max(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Max(&_PlugMath.CallOpts, x, y)
}

// Max is a free data retrieval call binding the contract method 0x6d5433e6.
//
// Solidity: function max(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCallerSession) Max(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Max(&_PlugMath.CallOpts, x, y)
}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCaller) Min(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "min", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathSession) Min(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Min(&_PlugMath.CallOpts, x, y)
}

// Min is a free data retrieval call binding the contract method 0x7ae2b5c7.
//
// Solidity: function min(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCallerSession) Min(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Min(&_PlugMath.CallOpts, x, y)
}

// Modulo is a free data retrieval call binding the contract method 0xbaaf073d.
//
// Solidity: function modulo(uint256 a, uint256 b) pure returns(uint256 result)
func (_PlugMath *PlugMathCaller) Modulo(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "modulo", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Modulo is a free data retrieval call binding the contract method 0xbaaf073d.
//
// Solidity: function modulo(uint256 a, uint256 b) pure returns(uint256 result)
func (_PlugMath *PlugMathSession) Modulo(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Modulo(&_PlugMath.CallOpts, a, b)
}

// Modulo is a free data retrieval call binding the contract method 0xbaaf073d.
//
// Solidity: function modulo(uint256 a, uint256 b) pure returns(uint256 result)
func (_PlugMath *PlugMathCallerSession) Modulo(a *big.Int, b *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Modulo(&_PlugMath.CallOpts, a, b)
}

// Multiply is a free data retrieval call binding the contract method 0x165c4a16.
//
// Solidity: function multiply(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCaller) Multiply(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "multiply", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Multiply is a free data retrieval call binding the contract method 0x165c4a16.
//
// Solidity: function multiply(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathSession) Multiply(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Multiply(&_PlugMath.CallOpts, x, y)
}

// Multiply is a free data retrieval call binding the contract method 0x165c4a16.
//
// Solidity: function multiply(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCallerSession) Multiply(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Multiply(&_PlugMath.CallOpts, x, y)
}

// Power is a free data retrieval call binding the contract method 0xc04f01fc.
//
// Solidity: function power(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCaller) Power(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "power", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Power is a free data retrieval call binding the contract method 0xc04f01fc.
//
// Solidity: function power(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathSession) Power(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Power(&_PlugMath.CallOpts, x, y)
}

// Power is a free data retrieval call binding the contract method 0xc04f01fc.
//
// Solidity: function power(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCallerSession) Power(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Power(&_PlugMath.CallOpts, x, y)
}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCaller) Subtract(opts *bind.CallOpts, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugMath.contract.Call(opts, &out, "subtract", x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathSession) Subtract(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Subtract(&_PlugMath.CallOpts, x, y)
}

// Subtract is a free data retrieval call binding the contract method 0x3ef5e445.
//
// Solidity: function subtract(uint256 x, uint256 y) pure returns(uint256 result)
func (_PlugMath *PlugMathCallerSession) Subtract(x *big.Int, y *big.Int) (*big.Int, error) {
	return _PlugMath.Contract.Subtract(&_PlugMath.CallOpts, x, y)
}
