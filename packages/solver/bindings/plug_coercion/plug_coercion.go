// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_coercion

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

// PlugCoercionMetaData contains all meta data concerning the PlugCoercion contract.
var PlugCoercionMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addressToUint160\",\"inputs\":[{\"name\":\"y\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"boolToInt\",\"inputs\":[{\"name\":\"y\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"boolToUint\",\"inputs\":[{\"name\":\"y\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"bytes32ToBytes\",\"inputs\":[{\"name\":\"y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"bytes32ToUint\",\"inputs\":[{\"name\":\"y\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"bytesToBytes32\",\"inputs\":[{\"name\":\"y\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"intToBool\",\"inputs\":[{\"name\":\"y\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toInt128\",\"inputs\":[{\"name\":\"y\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"int128\",\"internalType\":\"int128\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toInt16\",\"inputs\":[{\"name\":\"y\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"int16\",\"internalType\":\"int16\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toInt256\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toInt32\",\"inputs\":[{\"name\":\"y\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"int32\",\"internalType\":\"int32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toInt64\",\"inputs\":[{\"name\":\"y\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toInt8\",\"inputs\":[{\"name\":\"y\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"int8\",\"internalType\":\"int8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUint128\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUint16\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUint160\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUint256\",\"inputs\":[{\"name\":\"y\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUint32\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUint64\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUint8\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUint96\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"uint160ToAddress\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint160\",\"internalType\":\"uint160\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"uintToBool\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"uintToBytes32\",\"inputs\":[{\"name\":\"y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"z\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"}]",
}

// PlugCoercionABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugCoercionMetaData.ABI instead.
var PlugCoercionABI = PlugCoercionMetaData.ABI

// PlugCoercion is an auto generated Go binding around an Ethereum contract.
type PlugCoercion struct {
	PlugCoercionCaller     // Read-only binding to the contract
	PlugCoercionTransactor // Write-only binding to the contract
	PlugCoercionFilterer   // Log filterer for contract events
}

// PlugCoercionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugCoercionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugCoercionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugCoercionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugCoercionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugCoercionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugCoercionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugCoercionSession struct {
	Contract     *PlugCoercion     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugCoercionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugCoercionCallerSession struct {
	Contract *PlugCoercionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PlugCoercionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugCoercionTransactorSession struct {
	Contract     *PlugCoercionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PlugCoercionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugCoercionRaw struct {
	Contract *PlugCoercion // Generic contract binding to access the raw methods on
}

// PlugCoercionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugCoercionCallerRaw struct {
	Contract *PlugCoercionCaller // Generic read-only contract binding to access the raw methods on
}

// PlugCoercionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugCoercionTransactorRaw struct {
	Contract *PlugCoercionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugCoercion creates a new instance of PlugCoercion, bound to a specific deployed contract.
func NewPlugCoercion(address common.Address, backend bind.ContractBackend) (*PlugCoercion, error) {
	contract, err := bindPlugCoercion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugCoercion{PlugCoercionCaller: PlugCoercionCaller{contract: contract}, PlugCoercionTransactor: PlugCoercionTransactor{contract: contract}, PlugCoercionFilterer: PlugCoercionFilterer{contract: contract}}, nil
}

// NewPlugCoercionCaller creates a new read-only instance of PlugCoercion, bound to a specific deployed contract.
func NewPlugCoercionCaller(address common.Address, caller bind.ContractCaller) (*PlugCoercionCaller, error) {
	contract, err := bindPlugCoercion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugCoercionCaller{contract: contract}, nil
}

// NewPlugCoercionTransactor creates a new write-only instance of PlugCoercion, bound to a specific deployed contract.
func NewPlugCoercionTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugCoercionTransactor, error) {
	contract, err := bindPlugCoercion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugCoercionTransactor{contract: contract}, nil
}

// NewPlugCoercionFilterer creates a new log filterer instance of PlugCoercion, bound to a specific deployed contract.
func NewPlugCoercionFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugCoercionFilterer, error) {
	contract, err := bindPlugCoercion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugCoercionFilterer{contract: contract}, nil
}

// bindPlugCoercion binds a generic wrapper to an already deployed contract.
func bindPlugCoercion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugCoercionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugCoercion *PlugCoercionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugCoercion.Contract.PlugCoercionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugCoercion *PlugCoercionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugCoercion.Contract.PlugCoercionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugCoercion *PlugCoercionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugCoercion.Contract.PlugCoercionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugCoercion *PlugCoercionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugCoercion.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugCoercion *PlugCoercionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugCoercion.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugCoercion *PlugCoercionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugCoercion.Contract.contract.Transact(opts, method, params...)
}

// AddressToUint160 is a free data retrieval call binding the contract method 0xf73ecfde.
//
// Solidity: function addressToUint160(address y) pure returns(uint160 z)
func (_PlugCoercion *PlugCoercionCaller) AddressToUint160(opts *bind.CallOpts, y common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "addressToUint160", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressToUint160 is a free data retrieval call binding the contract method 0xf73ecfde.
//
// Solidity: function addressToUint160(address y) pure returns(uint160 z)
func (_PlugCoercion *PlugCoercionSession) AddressToUint160(y common.Address) (*big.Int, error) {
	return _PlugCoercion.Contract.AddressToUint160(&_PlugCoercion.CallOpts, y)
}

// AddressToUint160 is a free data retrieval call binding the contract method 0xf73ecfde.
//
// Solidity: function addressToUint160(address y) pure returns(uint160 z)
func (_PlugCoercion *PlugCoercionCallerSession) AddressToUint160(y common.Address) (*big.Int, error) {
	return _PlugCoercion.Contract.AddressToUint160(&_PlugCoercion.CallOpts, y)
}

// BoolToInt is a free data retrieval call binding the contract method 0x5192f3c0.
//
// Solidity: function boolToInt(bool y) pure returns(int256 z)
func (_PlugCoercion *PlugCoercionCaller) BoolToInt(opts *bind.CallOpts, y bool) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "boolToInt", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoolToInt is a free data retrieval call binding the contract method 0x5192f3c0.
//
// Solidity: function boolToInt(bool y) pure returns(int256 z)
func (_PlugCoercion *PlugCoercionSession) BoolToInt(y bool) (*big.Int, error) {
	return _PlugCoercion.Contract.BoolToInt(&_PlugCoercion.CallOpts, y)
}

// BoolToInt is a free data retrieval call binding the contract method 0x5192f3c0.
//
// Solidity: function boolToInt(bool y) pure returns(int256 z)
func (_PlugCoercion *PlugCoercionCallerSession) BoolToInt(y bool) (*big.Int, error) {
	return _PlugCoercion.Contract.BoolToInt(&_PlugCoercion.CallOpts, y)
}

// BoolToUint is a free data retrieval call binding the contract method 0x22009642.
//
// Solidity: function boolToUint(bool y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionCaller) BoolToUint(opts *bind.CallOpts, y bool) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "boolToUint", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoolToUint is a free data retrieval call binding the contract method 0x22009642.
//
// Solidity: function boolToUint(bool y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionSession) BoolToUint(y bool) (*big.Int, error) {
	return _PlugCoercion.Contract.BoolToUint(&_PlugCoercion.CallOpts, y)
}

// BoolToUint is a free data retrieval call binding the contract method 0x22009642.
//
// Solidity: function boolToUint(bool y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionCallerSession) BoolToUint(y bool) (*big.Int, error) {
	return _PlugCoercion.Contract.BoolToUint(&_PlugCoercion.CallOpts, y)
}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 y) pure returns(bytes z)
func (_PlugCoercion *PlugCoercionCaller) Bytes32ToBytes(opts *bind.CallOpts, y [32]byte) ([]byte, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "bytes32ToBytes", y)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 y) pure returns(bytes z)
func (_PlugCoercion *PlugCoercionSession) Bytes32ToBytes(y [32]byte) ([]byte, error) {
	return _PlugCoercion.Contract.Bytes32ToBytes(&_PlugCoercion.CallOpts, y)
}

// Bytes32ToBytes is a free data retrieval call binding the contract method 0x4c0999c7.
//
// Solidity: function bytes32ToBytes(bytes32 y) pure returns(bytes z)
func (_PlugCoercion *PlugCoercionCallerSession) Bytes32ToBytes(y [32]byte) ([]byte, error) {
	return _PlugCoercion.Contract.Bytes32ToBytes(&_PlugCoercion.CallOpts, y)
}

// Bytes32ToUint is a free data retrieval call binding the contract method 0xe60b1424.
//
// Solidity: function bytes32ToUint(bytes32 y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionCaller) Bytes32ToUint(opts *bind.CallOpts, y [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "bytes32ToUint", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bytes32ToUint is a free data retrieval call binding the contract method 0xe60b1424.
//
// Solidity: function bytes32ToUint(bytes32 y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionSession) Bytes32ToUint(y [32]byte) (*big.Int, error) {
	return _PlugCoercion.Contract.Bytes32ToUint(&_PlugCoercion.CallOpts, y)
}

// Bytes32ToUint is a free data retrieval call binding the contract method 0xe60b1424.
//
// Solidity: function bytes32ToUint(bytes32 y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionCallerSession) Bytes32ToUint(y [32]byte) (*big.Int, error) {
	return _PlugCoercion.Contract.Bytes32ToUint(&_PlugCoercion.CallOpts, y)
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes y) pure returns(bytes32 z)
func (_PlugCoercion *PlugCoercionCaller) BytesToBytes32(opts *bind.CallOpts, y []byte) ([32]byte, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "bytesToBytes32", y)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes y) pure returns(bytes32 z)
func (_PlugCoercion *PlugCoercionSession) BytesToBytes32(y []byte) ([32]byte, error) {
	return _PlugCoercion.Contract.BytesToBytes32(&_PlugCoercion.CallOpts, y)
}

// BytesToBytes32 is a free data retrieval call binding the contract method 0xbfe370d9.
//
// Solidity: function bytesToBytes32(bytes y) pure returns(bytes32 z)
func (_PlugCoercion *PlugCoercionCallerSession) BytesToBytes32(y []byte) ([32]byte, error) {
	return _PlugCoercion.Contract.BytesToBytes32(&_PlugCoercion.CallOpts, y)
}

// IntToBool is a free data retrieval call binding the contract method 0xa3586a4c.
//
// Solidity: function intToBool(int256 y) pure returns(bool z)
func (_PlugCoercion *PlugCoercionCaller) IntToBool(opts *bind.CallOpts, y *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "intToBool", y)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IntToBool is a free data retrieval call binding the contract method 0xa3586a4c.
//
// Solidity: function intToBool(int256 y) pure returns(bool z)
func (_PlugCoercion *PlugCoercionSession) IntToBool(y *big.Int) (bool, error) {
	return _PlugCoercion.Contract.IntToBool(&_PlugCoercion.CallOpts, y)
}

// IntToBool is a free data retrieval call binding the contract method 0xa3586a4c.
//
// Solidity: function intToBool(int256 y) pure returns(bool z)
func (_PlugCoercion *PlugCoercionCallerSession) IntToBool(y *big.Int) (bool, error) {
	return _PlugCoercion.Contract.IntToBool(&_PlugCoercion.CallOpts, y)
}

// ToInt128 is a free data retrieval call binding the contract method 0xdd2a0316.
//
// Solidity: function toInt128(int256 y) pure returns(int128 z)
func (_PlugCoercion *PlugCoercionCaller) ToInt128(opts *bind.CallOpts, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toInt128", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToInt128 is a free data retrieval call binding the contract method 0xdd2a0316.
//
// Solidity: function toInt128(int256 y) pure returns(int128 z)
func (_PlugCoercion *PlugCoercionSession) ToInt128(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToInt128(&_PlugCoercion.CallOpts, y)
}

// ToInt128 is a free data retrieval call binding the contract method 0xdd2a0316.
//
// Solidity: function toInt128(int256 y) pure returns(int128 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToInt128(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToInt128(&_PlugCoercion.CallOpts, y)
}

// ToInt16 is a free data retrieval call binding the contract method 0xcf65b4d3.
//
// Solidity: function toInt16(int256 y) pure returns(int16 z)
func (_PlugCoercion *PlugCoercionCaller) ToInt16(opts *bind.CallOpts, y *big.Int) (int16, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toInt16", y)

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// ToInt16 is a free data retrieval call binding the contract method 0xcf65b4d3.
//
// Solidity: function toInt16(int256 y) pure returns(int16 z)
func (_PlugCoercion *PlugCoercionSession) ToInt16(y *big.Int) (int16, error) {
	return _PlugCoercion.Contract.ToInt16(&_PlugCoercion.CallOpts, y)
}

// ToInt16 is a free data retrieval call binding the contract method 0xcf65b4d3.
//
// Solidity: function toInt16(int256 y) pure returns(int16 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToInt16(y *big.Int) (int16, error) {
	return _PlugCoercion.Contract.ToInt16(&_PlugCoercion.CallOpts, y)
}

// ToInt256 is a free data retrieval call binding the contract method 0xdfbe873b.
//
// Solidity: function toInt256(uint256 y) pure returns(int256 z)
func (_PlugCoercion *PlugCoercionCaller) ToInt256(opts *bind.CallOpts, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toInt256", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToInt256 is a free data retrieval call binding the contract method 0xdfbe873b.
//
// Solidity: function toInt256(uint256 y) pure returns(int256 z)
func (_PlugCoercion *PlugCoercionSession) ToInt256(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToInt256(&_PlugCoercion.CallOpts, y)
}

// ToInt256 is a free data retrieval call binding the contract method 0xdfbe873b.
//
// Solidity: function toInt256(uint256 y) pure returns(int256 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToInt256(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToInt256(&_PlugCoercion.CallOpts, y)
}

// ToInt32 is a free data retrieval call binding the contract method 0x9c6f59be.
//
// Solidity: function toInt32(int256 y) pure returns(int32 z)
func (_PlugCoercion *PlugCoercionCaller) ToInt32(opts *bind.CallOpts, y *big.Int) (int32, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toInt32", y)

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// ToInt32 is a free data retrieval call binding the contract method 0x9c6f59be.
//
// Solidity: function toInt32(int256 y) pure returns(int32 z)
func (_PlugCoercion *PlugCoercionSession) ToInt32(y *big.Int) (int32, error) {
	return _PlugCoercion.Contract.ToInt32(&_PlugCoercion.CallOpts, y)
}

// ToInt32 is a free data retrieval call binding the contract method 0x9c6f59be.
//
// Solidity: function toInt32(int256 y) pure returns(int32 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToInt32(y *big.Int) (int32, error) {
	return _PlugCoercion.Contract.ToInt32(&_PlugCoercion.CallOpts, y)
}

// ToInt64 is a free data retrieval call binding the contract method 0xd6bd32aa.
//
// Solidity: function toInt64(int256 y) pure returns(int64 z)
func (_PlugCoercion *PlugCoercionCaller) ToInt64(opts *bind.CallOpts, y *big.Int) (int64, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toInt64", y)

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// ToInt64 is a free data retrieval call binding the contract method 0xd6bd32aa.
//
// Solidity: function toInt64(int256 y) pure returns(int64 z)
func (_PlugCoercion *PlugCoercionSession) ToInt64(y *big.Int) (int64, error) {
	return _PlugCoercion.Contract.ToInt64(&_PlugCoercion.CallOpts, y)
}

// ToInt64 is a free data retrieval call binding the contract method 0xd6bd32aa.
//
// Solidity: function toInt64(int256 y) pure returns(int64 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToInt64(y *big.Int) (int64, error) {
	return _PlugCoercion.Contract.ToInt64(&_PlugCoercion.CallOpts, y)
}

// ToInt8 is a free data retrieval call binding the contract method 0xf136dc02.
//
// Solidity: function toInt8(int256 y) pure returns(int8 z)
func (_PlugCoercion *PlugCoercionCaller) ToInt8(opts *bind.CallOpts, y *big.Int) (int8, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toInt8", y)

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// ToInt8 is a free data retrieval call binding the contract method 0xf136dc02.
//
// Solidity: function toInt8(int256 y) pure returns(int8 z)
func (_PlugCoercion *PlugCoercionSession) ToInt8(y *big.Int) (int8, error) {
	return _PlugCoercion.Contract.ToInt8(&_PlugCoercion.CallOpts, y)
}

// ToInt8 is a free data retrieval call binding the contract method 0xf136dc02.
//
// Solidity: function toInt8(int256 y) pure returns(int8 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToInt8(y *big.Int) (int8, error) {
	return _PlugCoercion.Contract.ToInt8(&_PlugCoercion.CallOpts, y)
}

// ToUint128 is a free data retrieval call binding the contract method 0x809fdd33.
//
// Solidity: function toUint128(uint256 y) pure returns(uint128 z)
func (_PlugCoercion *PlugCoercionCaller) ToUint128(opts *bind.CallOpts, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toUint128", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToUint128 is a free data retrieval call binding the contract method 0x809fdd33.
//
// Solidity: function toUint128(uint256 y) pure returns(uint128 z)
func (_PlugCoercion *PlugCoercionSession) ToUint128(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToUint128(&_PlugCoercion.CallOpts, y)
}

// ToUint128 is a free data retrieval call binding the contract method 0x809fdd33.
//
// Solidity: function toUint128(uint256 y) pure returns(uint128 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToUint128(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToUint128(&_PlugCoercion.CallOpts, y)
}

// ToUint16 is a free data retrieval call binding the contract method 0x9374068f.
//
// Solidity: function toUint16(uint256 y) pure returns(uint16 z)
func (_PlugCoercion *PlugCoercionCaller) ToUint16(opts *bind.CallOpts, y *big.Int) (uint16, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toUint16", y)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ToUint16 is a free data retrieval call binding the contract method 0x9374068f.
//
// Solidity: function toUint16(uint256 y) pure returns(uint16 z)
func (_PlugCoercion *PlugCoercionSession) ToUint16(y *big.Int) (uint16, error) {
	return _PlugCoercion.Contract.ToUint16(&_PlugCoercion.CallOpts, y)
}

// ToUint16 is a free data retrieval call binding the contract method 0x9374068f.
//
// Solidity: function toUint16(uint256 y) pure returns(uint16 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToUint16(y *big.Int) (uint16, error) {
	return _PlugCoercion.Contract.ToUint16(&_PlugCoercion.CallOpts, y)
}

// ToUint160 is a free data retrieval call binding the contract method 0xdfef6beb.
//
// Solidity: function toUint160(uint256 y) pure returns(uint160 z)
func (_PlugCoercion *PlugCoercionCaller) ToUint160(opts *bind.CallOpts, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toUint160", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToUint160 is a free data retrieval call binding the contract method 0xdfef6beb.
//
// Solidity: function toUint160(uint256 y) pure returns(uint160 z)
func (_PlugCoercion *PlugCoercionSession) ToUint160(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToUint160(&_PlugCoercion.CallOpts, y)
}

// ToUint160 is a free data retrieval call binding the contract method 0xdfef6beb.
//
// Solidity: function toUint160(uint256 y) pure returns(uint160 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToUint160(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToUint160(&_PlugCoercion.CallOpts, y)
}

// ToUint256 is a free data retrieval call binding the contract method 0xfdcf791b.
//
// Solidity: function toUint256(int256 y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionCaller) ToUint256(opts *bind.CallOpts, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toUint256", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToUint256 is a free data retrieval call binding the contract method 0xfdcf791b.
//
// Solidity: function toUint256(int256 y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionSession) ToUint256(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToUint256(&_PlugCoercion.CallOpts, y)
}

// ToUint256 is a free data retrieval call binding the contract method 0xfdcf791b.
//
// Solidity: function toUint256(int256 y) pure returns(uint256 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToUint256(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToUint256(&_PlugCoercion.CallOpts, y)
}

// ToUint32 is a free data retrieval call binding the contract method 0xc8193255.
//
// Solidity: function toUint32(uint256 y) pure returns(uint32 z)
func (_PlugCoercion *PlugCoercionCaller) ToUint32(opts *bind.CallOpts, y *big.Int) (uint32, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toUint32", y)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ToUint32 is a free data retrieval call binding the contract method 0xc8193255.
//
// Solidity: function toUint32(uint256 y) pure returns(uint32 z)
func (_PlugCoercion *PlugCoercionSession) ToUint32(y *big.Int) (uint32, error) {
	return _PlugCoercion.Contract.ToUint32(&_PlugCoercion.CallOpts, y)
}

// ToUint32 is a free data retrieval call binding the contract method 0xc8193255.
//
// Solidity: function toUint32(uint256 y) pure returns(uint32 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToUint32(y *big.Int) (uint32, error) {
	return _PlugCoercion.Contract.ToUint32(&_PlugCoercion.CallOpts, y)
}

// ToUint64 is a free data retrieval call binding the contract method 0x2665fad0.
//
// Solidity: function toUint64(uint256 y) pure returns(uint64 z)
func (_PlugCoercion *PlugCoercionCaller) ToUint64(opts *bind.CallOpts, y *big.Int) (uint64, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toUint64", y)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ToUint64 is a free data retrieval call binding the contract method 0x2665fad0.
//
// Solidity: function toUint64(uint256 y) pure returns(uint64 z)
func (_PlugCoercion *PlugCoercionSession) ToUint64(y *big.Int) (uint64, error) {
	return _PlugCoercion.Contract.ToUint64(&_PlugCoercion.CallOpts, y)
}

// ToUint64 is a free data retrieval call binding the contract method 0x2665fad0.
//
// Solidity: function toUint64(uint256 y) pure returns(uint64 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToUint64(y *big.Int) (uint64, error) {
	return _PlugCoercion.Contract.ToUint64(&_PlugCoercion.CallOpts, y)
}

// ToUint8 is a free data retrieval call binding the contract method 0x0cc4681e.
//
// Solidity: function toUint8(uint256 y) pure returns(uint8 z)
func (_PlugCoercion *PlugCoercionCaller) ToUint8(opts *bind.CallOpts, y *big.Int) (uint8, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toUint8", y)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ToUint8 is a free data retrieval call binding the contract method 0x0cc4681e.
//
// Solidity: function toUint8(uint256 y) pure returns(uint8 z)
func (_PlugCoercion *PlugCoercionSession) ToUint8(y *big.Int) (uint8, error) {
	return _PlugCoercion.Contract.ToUint8(&_PlugCoercion.CallOpts, y)
}

// ToUint8 is a free data retrieval call binding the contract method 0x0cc4681e.
//
// Solidity: function toUint8(uint256 y) pure returns(uint8 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToUint8(y *big.Int) (uint8, error) {
	return _PlugCoercion.Contract.ToUint8(&_PlugCoercion.CallOpts, y)
}

// ToUint96 is a free data retrieval call binding the contract method 0x1cf887fc.
//
// Solidity: function toUint96(uint256 y) pure returns(uint96 z)
func (_PlugCoercion *PlugCoercionCaller) ToUint96(opts *bind.CallOpts, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "toUint96", y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToUint96 is a free data retrieval call binding the contract method 0x1cf887fc.
//
// Solidity: function toUint96(uint256 y) pure returns(uint96 z)
func (_PlugCoercion *PlugCoercionSession) ToUint96(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToUint96(&_PlugCoercion.CallOpts, y)
}

// ToUint96 is a free data retrieval call binding the contract method 0x1cf887fc.
//
// Solidity: function toUint96(uint256 y) pure returns(uint96 z)
func (_PlugCoercion *PlugCoercionCallerSession) ToUint96(y *big.Int) (*big.Int, error) {
	return _PlugCoercion.Contract.ToUint96(&_PlugCoercion.CallOpts, y)
}

// Uint160ToAddress is a free data retrieval call binding the contract method 0xfbe706fd.
//
// Solidity: function uint160ToAddress(uint160 y) pure returns(address z)
func (_PlugCoercion *PlugCoercionCaller) Uint160ToAddress(opts *bind.CallOpts, y *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "uint160ToAddress", y)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Uint160ToAddress is a free data retrieval call binding the contract method 0xfbe706fd.
//
// Solidity: function uint160ToAddress(uint160 y) pure returns(address z)
func (_PlugCoercion *PlugCoercionSession) Uint160ToAddress(y *big.Int) (common.Address, error) {
	return _PlugCoercion.Contract.Uint160ToAddress(&_PlugCoercion.CallOpts, y)
}

// Uint160ToAddress is a free data retrieval call binding the contract method 0xfbe706fd.
//
// Solidity: function uint160ToAddress(uint160 y) pure returns(address z)
func (_PlugCoercion *PlugCoercionCallerSession) Uint160ToAddress(y *big.Int) (common.Address, error) {
	return _PlugCoercion.Contract.Uint160ToAddress(&_PlugCoercion.CallOpts, y)
}

// UintToBool is a free data retrieval call binding the contract method 0x612eb5cc.
//
// Solidity: function uintToBool(uint256 y) pure returns(bool z)
func (_PlugCoercion *PlugCoercionCaller) UintToBool(opts *bind.CallOpts, y *big.Int) (bool, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "uintToBool", y)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UintToBool is a free data retrieval call binding the contract method 0x612eb5cc.
//
// Solidity: function uintToBool(uint256 y) pure returns(bool z)
func (_PlugCoercion *PlugCoercionSession) UintToBool(y *big.Int) (bool, error) {
	return _PlugCoercion.Contract.UintToBool(&_PlugCoercion.CallOpts, y)
}

// UintToBool is a free data retrieval call binding the contract method 0x612eb5cc.
//
// Solidity: function uintToBool(uint256 y) pure returns(bool z)
func (_PlugCoercion *PlugCoercionCallerSession) UintToBool(y *big.Int) (bool, error) {
	return _PlugCoercion.Contract.UintToBool(&_PlugCoercion.CallOpts, y)
}

// UintToBytes32 is a free data retrieval call binding the contract method 0x886d3db9.
//
// Solidity: function uintToBytes32(uint256 y) pure returns(bytes32 z)
func (_PlugCoercion *PlugCoercionCaller) UintToBytes32(opts *bind.CallOpts, y *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PlugCoercion.contract.Call(opts, &out, "uintToBytes32", y)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UintToBytes32 is a free data retrieval call binding the contract method 0x886d3db9.
//
// Solidity: function uintToBytes32(uint256 y) pure returns(bytes32 z)
func (_PlugCoercion *PlugCoercionSession) UintToBytes32(y *big.Int) ([32]byte, error) {
	return _PlugCoercion.Contract.UintToBytes32(&_PlugCoercion.CallOpts, y)
}

// UintToBytes32 is a free data retrieval call binding the contract method 0x886d3db9.
//
// Solidity: function uintToBytes32(uint256 y) pure returns(bytes32 z)
func (_PlugCoercion *PlugCoercionCallerSession) UintToBytes32(y *big.Int) ([32]byte, error) {
	return _PlugCoercion.Contract.UintToBytes32(&_PlugCoercion.CallOpts, y)
}
