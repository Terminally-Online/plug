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
	ABI: "[{\"type\":\"function\",\"name\":\"TYPE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TYPE_BOOL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TYPE_BYTES\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TYPE_BYTES32\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TYPE_INT256\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TYPE_STRING\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TYPE_UINT256\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressStorage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchSet\",\"inputs\":[{\"name\":\"keys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"values\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"typeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"boolStorage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bytes32Storage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bytesStorage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exists\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAddress\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBool\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBytes\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getInt\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getString\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getType\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUint\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"intStorage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"remove\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeWithType\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"typeId\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"set\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAddress\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBool\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInt\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setString\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUint\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stringStorage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uintStorage\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
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

// TYPEADDRESS is a free data retrieval call binding the contract method 0x3e8584da.
//
// Solidity: function TYPE_ADDRESS() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCaller) TYPEADDRESS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "TYPE_ADDRESS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPEADDRESS is a free data retrieval call binding the contract method 0x3e8584da.
//
// Solidity: function TYPE_ADDRESS() view returns(uint8)
func (_PlugDatabase *PlugDatabaseSession) TYPEADDRESS() (uint8, error) {
	return _PlugDatabase.Contract.TYPEADDRESS(&_PlugDatabase.CallOpts)
}

// TYPEADDRESS is a free data retrieval call binding the contract method 0x3e8584da.
//
// Solidity: function TYPE_ADDRESS() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCallerSession) TYPEADDRESS() (uint8, error) {
	return _PlugDatabase.Contract.TYPEADDRESS(&_PlugDatabase.CallOpts)
}

// TYPEBOOL is a free data retrieval call binding the contract method 0xecdea6da.
//
// Solidity: function TYPE_BOOL() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCaller) TYPEBOOL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "TYPE_BOOL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPEBOOL is a free data retrieval call binding the contract method 0xecdea6da.
//
// Solidity: function TYPE_BOOL() view returns(uint8)
func (_PlugDatabase *PlugDatabaseSession) TYPEBOOL() (uint8, error) {
	return _PlugDatabase.Contract.TYPEBOOL(&_PlugDatabase.CallOpts)
}

// TYPEBOOL is a free data retrieval call binding the contract method 0xecdea6da.
//
// Solidity: function TYPE_BOOL() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCallerSession) TYPEBOOL() (uint8, error) {
	return _PlugDatabase.Contract.TYPEBOOL(&_PlugDatabase.CallOpts)
}

// TYPEBYTES is a free data retrieval call binding the contract method 0x074f3f3a.
//
// Solidity: function TYPE_BYTES() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCaller) TYPEBYTES(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "TYPE_BYTES")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPEBYTES is a free data retrieval call binding the contract method 0x074f3f3a.
//
// Solidity: function TYPE_BYTES() view returns(uint8)
func (_PlugDatabase *PlugDatabaseSession) TYPEBYTES() (uint8, error) {
	return _PlugDatabase.Contract.TYPEBYTES(&_PlugDatabase.CallOpts)
}

// TYPEBYTES is a free data retrieval call binding the contract method 0x074f3f3a.
//
// Solidity: function TYPE_BYTES() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCallerSession) TYPEBYTES() (uint8, error) {
	return _PlugDatabase.Contract.TYPEBYTES(&_PlugDatabase.CallOpts)
}

// TYPEBYTES32 is a free data retrieval call binding the contract method 0x5a0f05cc.
//
// Solidity: function TYPE_BYTES32() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCaller) TYPEBYTES32(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "TYPE_BYTES32")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPEBYTES32 is a free data retrieval call binding the contract method 0x5a0f05cc.
//
// Solidity: function TYPE_BYTES32() view returns(uint8)
func (_PlugDatabase *PlugDatabaseSession) TYPEBYTES32() (uint8, error) {
	return _PlugDatabase.Contract.TYPEBYTES32(&_PlugDatabase.CallOpts)
}

// TYPEBYTES32 is a free data retrieval call binding the contract method 0x5a0f05cc.
//
// Solidity: function TYPE_BYTES32() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCallerSession) TYPEBYTES32() (uint8, error) {
	return _PlugDatabase.Contract.TYPEBYTES32(&_PlugDatabase.CallOpts)
}

// TYPEINT256 is a free data retrieval call binding the contract method 0x0cdbaee9.
//
// Solidity: function TYPE_INT256() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCaller) TYPEINT256(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "TYPE_INT256")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPEINT256 is a free data retrieval call binding the contract method 0x0cdbaee9.
//
// Solidity: function TYPE_INT256() view returns(uint8)
func (_PlugDatabase *PlugDatabaseSession) TYPEINT256() (uint8, error) {
	return _PlugDatabase.Contract.TYPEINT256(&_PlugDatabase.CallOpts)
}

// TYPEINT256 is a free data retrieval call binding the contract method 0x0cdbaee9.
//
// Solidity: function TYPE_INT256() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCallerSession) TYPEINT256() (uint8, error) {
	return _PlugDatabase.Contract.TYPEINT256(&_PlugDatabase.CallOpts)
}

// TYPESTRING is a free data retrieval call binding the contract method 0xbb7a6eee.
//
// Solidity: function TYPE_STRING() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCaller) TYPESTRING(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "TYPE_STRING")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPESTRING is a free data retrieval call binding the contract method 0xbb7a6eee.
//
// Solidity: function TYPE_STRING() view returns(uint8)
func (_PlugDatabase *PlugDatabaseSession) TYPESTRING() (uint8, error) {
	return _PlugDatabase.Contract.TYPESTRING(&_PlugDatabase.CallOpts)
}

// TYPESTRING is a free data retrieval call binding the contract method 0xbb7a6eee.
//
// Solidity: function TYPE_STRING() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCallerSession) TYPESTRING() (uint8, error) {
	return _PlugDatabase.Contract.TYPESTRING(&_PlugDatabase.CallOpts)
}

// TYPEUINT256 is a free data retrieval call binding the contract method 0x035412d5.
//
// Solidity: function TYPE_UINT256() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCaller) TYPEUINT256(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "TYPE_UINT256")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TYPEUINT256 is a free data retrieval call binding the contract method 0x035412d5.
//
// Solidity: function TYPE_UINT256() view returns(uint8)
func (_PlugDatabase *PlugDatabaseSession) TYPEUINT256() (uint8, error) {
	return _PlugDatabase.Contract.TYPEUINT256(&_PlugDatabase.CallOpts)
}

// TYPEUINT256 is a free data retrieval call binding the contract method 0x035412d5.
//
// Solidity: function TYPE_UINT256() view returns(uint8)
func (_PlugDatabase *PlugDatabaseCallerSession) TYPEUINT256() (uint8, error) {
	return _PlugDatabase.Contract.TYPEUINT256(&_PlugDatabase.CallOpts)
}

// AddressStorage is a free data retrieval call binding the contract method 0xe05f1b35.
//
// Solidity: function addressStorage(address , bytes32 ) view returns(address)
func (_PlugDatabase *PlugDatabaseCaller) AddressStorage(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "addressStorage", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressStorage is a free data retrieval call binding the contract method 0xe05f1b35.
//
// Solidity: function addressStorage(address , bytes32 ) view returns(address)
func (_PlugDatabase *PlugDatabaseSession) AddressStorage(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _PlugDatabase.Contract.AddressStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// AddressStorage is a free data retrieval call binding the contract method 0xe05f1b35.
//
// Solidity: function addressStorage(address , bytes32 ) view returns(address)
func (_PlugDatabase *PlugDatabaseCallerSession) AddressStorage(arg0 common.Address, arg1 [32]byte) (common.Address, error) {
	return _PlugDatabase.Contract.AddressStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// BoolStorage is a free data retrieval call binding the contract method 0x2dea244b.
//
// Solidity: function boolStorage(address , bytes32 ) view returns(bool)
func (_PlugDatabase *PlugDatabaseCaller) BoolStorage(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "boolStorage", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BoolStorage is a free data retrieval call binding the contract method 0x2dea244b.
//
// Solidity: function boolStorage(address , bytes32 ) view returns(bool)
func (_PlugDatabase *PlugDatabaseSession) BoolStorage(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PlugDatabase.Contract.BoolStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// BoolStorage is a free data retrieval call binding the contract method 0x2dea244b.
//
// Solidity: function boolStorage(address , bytes32 ) view returns(bool)
func (_PlugDatabase *PlugDatabaseCallerSession) BoolStorage(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _PlugDatabase.Contract.BoolStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// Bytes32Storage is a free data retrieval call binding the contract method 0x7b40a17c.
//
// Solidity: function bytes32Storage(address , bytes32 ) view returns(bytes32)
func (_PlugDatabase *PlugDatabaseCaller) Bytes32Storage(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "bytes32Storage", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Bytes32Storage is a free data retrieval call binding the contract method 0x7b40a17c.
//
// Solidity: function bytes32Storage(address , bytes32 ) view returns(bytes32)
func (_PlugDatabase *PlugDatabaseSession) Bytes32Storage(arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	return _PlugDatabase.Contract.Bytes32Storage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// Bytes32Storage is a free data retrieval call binding the contract method 0x7b40a17c.
//
// Solidity: function bytes32Storage(address , bytes32 ) view returns(bytes32)
func (_PlugDatabase *PlugDatabaseCallerSession) Bytes32Storage(arg0 common.Address, arg1 [32]byte) ([32]byte, error) {
	return _PlugDatabase.Contract.Bytes32Storage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// BytesStorage is a free data retrieval call binding the contract method 0xc72cb92c.
//
// Solidity: function bytesStorage(address , bytes32 ) view returns(bytes)
func (_PlugDatabase *PlugDatabaseCaller) BytesStorage(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "bytesStorage", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// BytesStorage is a free data retrieval call binding the contract method 0xc72cb92c.
//
// Solidity: function bytesStorage(address , bytes32 ) view returns(bytes)
func (_PlugDatabase *PlugDatabaseSession) BytesStorage(arg0 common.Address, arg1 [32]byte) ([]byte, error) {
	return _PlugDatabase.Contract.BytesStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// BytesStorage is a free data retrieval call binding the contract method 0xc72cb92c.
//
// Solidity: function bytesStorage(address , bytes32 ) view returns(bytes)
func (_PlugDatabase *PlugDatabaseCallerSession) BytesStorage(arg0 common.Address, arg1 [32]byte) ([]byte, error) {
	return _PlugDatabase.Contract.BytesStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// Exists is a free data retrieval call binding the contract method 0xb13af58b.
//
// Solidity: function exists(address sender, bytes32 key, uint8 typeId) view returns(bool)
func (_PlugDatabase *PlugDatabaseCaller) Exists(opts *bind.CallOpts, sender common.Address, key [32]byte, typeId uint8) (bool, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "exists", sender, key, typeId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0xb13af58b.
//
// Solidity: function exists(address sender, bytes32 key, uint8 typeId) view returns(bool)
func (_PlugDatabase *PlugDatabaseSession) Exists(sender common.Address, key [32]byte, typeId uint8) (bool, error) {
	return _PlugDatabase.Contract.Exists(&_PlugDatabase.CallOpts, sender, key, typeId)
}

// Exists is a free data retrieval call binding the contract method 0xb13af58b.
//
// Solidity: function exists(address sender, bytes32 key, uint8 typeId) view returns(bool)
func (_PlugDatabase *PlugDatabaseCallerSession) Exists(sender common.Address, key [32]byte, typeId uint8) (bool, error) {
	return _PlugDatabase.Contract.Exists(&_PlugDatabase.CallOpts, sender, key, typeId)
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

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address sender, bytes32 key) view returns(address result)
func (_PlugDatabase *PlugDatabaseCaller) GetAddress(opts *bind.CallOpts, sender common.Address, key [32]byte) (common.Address, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "getAddress", sender, key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address sender, bytes32 key) view returns(address result)
func (_PlugDatabase *PlugDatabaseSession) GetAddress(sender common.Address, key [32]byte) (common.Address, error) {
	return _PlugDatabase.Contract.GetAddress(&_PlugDatabase.CallOpts, sender, key)
}

// GetAddress is a free data retrieval call binding the contract method 0x7ac4ed64.
//
// Solidity: function getAddress(address sender, bytes32 key) view returns(address result)
func (_PlugDatabase *PlugDatabaseCallerSession) GetAddress(sender common.Address, key [32]byte) (common.Address, error) {
	return _PlugDatabase.Contract.GetAddress(&_PlugDatabase.CallOpts, sender, key)
}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(address sender, bytes32 key) view returns(bool result)
func (_PlugDatabase *PlugDatabaseCaller) GetBool(opts *bind.CallOpts, sender common.Address, key [32]byte) (bool, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "getBool", sender, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(address sender, bytes32 key) view returns(bool result)
func (_PlugDatabase *PlugDatabaseSession) GetBool(sender common.Address, key [32]byte) (bool, error) {
	return _PlugDatabase.Contract.GetBool(&_PlugDatabase.CallOpts, sender, key)
}

// GetBool is a free data retrieval call binding the contract method 0x9d74b37d.
//
// Solidity: function getBool(address sender, bytes32 key) view returns(bool result)
func (_PlugDatabase *PlugDatabaseCallerSession) GetBool(sender common.Address, key [32]byte) (bool, error) {
	return _PlugDatabase.Contract.GetBool(&_PlugDatabase.CallOpts, sender, key)
}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(address sender, bytes32 key) view returns(bytes result)
func (_PlugDatabase *PlugDatabaseCaller) GetBytes(opts *bind.CallOpts, sender common.Address, key [32]byte) ([]byte, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "getBytes", sender, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(address sender, bytes32 key) view returns(bytes result)
func (_PlugDatabase *PlugDatabaseSession) GetBytes(sender common.Address, key [32]byte) ([]byte, error) {
	return _PlugDatabase.Contract.GetBytes(&_PlugDatabase.CallOpts, sender, key)
}

// GetBytes is a free data retrieval call binding the contract method 0x6556f767.
//
// Solidity: function getBytes(address sender, bytes32 key) view returns(bytes result)
func (_PlugDatabase *PlugDatabaseCallerSession) GetBytes(sender common.Address, key [32]byte) ([]byte, error) {
	return _PlugDatabase.Contract.GetBytes(&_PlugDatabase.CallOpts, sender, key)
}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(address sender, bytes32 key) view returns(int256 result)
func (_PlugDatabase *PlugDatabaseCaller) GetInt(opts *bind.CallOpts, sender common.Address, key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "getInt", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(address sender, bytes32 key) view returns(int256 result)
func (_PlugDatabase *PlugDatabaseSession) GetInt(sender common.Address, key [32]byte) (*big.Int, error) {
	return _PlugDatabase.Contract.GetInt(&_PlugDatabase.CallOpts, sender, key)
}

// GetInt is a free data retrieval call binding the contract method 0x95ee8bae.
//
// Solidity: function getInt(address sender, bytes32 key) view returns(int256 result)
func (_PlugDatabase *PlugDatabaseCallerSession) GetInt(sender common.Address, key [32]byte) (*big.Int, error) {
	return _PlugDatabase.Contract.GetInt(&_PlugDatabase.CallOpts, sender, key)
}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(address sender, bytes32 key) view returns(string result)
func (_PlugDatabase *PlugDatabaseCaller) GetString(opts *bind.CallOpts, sender common.Address, key [32]byte) (string, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "getString", sender, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(address sender, bytes32 key) view returns(string result)
func (_PlugDatabase *PlugDatabaseSession) GetString(sender common.Address, key [32]byte) (string, error) {
	return _PlugDatabase.Contract.GetString(&_PlugDatabase.CallOpts, sender, key)
}

// GetString is a free data retrieval call binding the contract method 0xe318de73.
//
// Solidity: function getString(address sender, bytes32 key) view returns(string result)
func (_PlugDatabase *PlugDatabaseCallerSession) GetString(sender common.Address, key [32]byte) (string, error) {
	return _PlugDatabase.Contract.GetString(&_PlugDatabase.CallOpts, sender, key)
}

// GetType is a free data retrieval call binding the contract method 0xfdd1311d.
//
// Solidity: function getType(address sender, bytes32 key) view returns(uint8)
func (_PlugDatabase *PlugDatabaseCaller) GetType(opts *bind.CallOpts, sender common.Address, key [32]byte) (uint8, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "getType", sender, key)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0xfdd1311d.
//
// Solidity: function getType(address sender, bytes32 key) view returns(uint8)
func (_PlugDatabase *PlugDatabaseSession) GetType(sender common.Address, key [32]byte) (uint8, error) {
	return _PlugDatabase.Contract.GetType(&_PlugDatabase.CallOpts, sender, key)
}

// GetType is a free data retrieval call binding the contract method 0xfdd1311d.
//
// Solidity: function getType(address sender, bytes32 key) view returns(uint8)
func (_PlugDatabase *PlugDatabaseCallerSession) GetType(sender common.Address, key [32]byte) (uint8, error) {
	return _PlugDatabase.Contract.GetType(&_PlugDatabase.CallOpts, sender, key)
}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(address sender, bytes32 key) view returns(uint256 result)
func (_PlugDatabase *PlugDatabaseCaller) GetUint(opts *bind.CallOpts, sender common.Address, key [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "getUint", sender, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(address sender, bytes32 key) view returns(uint256 result)
func (_PlugDatabase *PlugDatabaseSession) GetUint(sender common.Address, key [32]byte) (*big.Int, error) {
	return _PlugDatabase.Contract.GetUint(&_PlugDatabase.CallOpts, sender, key)
}

// GetUint is a free data retrieval call binding the contract method 0x71658552.
//
// Solidity: function getUint(address sender, bytes32 key) view returns(uint256 result)
func (_PlugDatabase *PlugDatabaseCallerSession) GetUint(sender common.Address, key [32]byte) (*big.Int, error) {
	return _PlugDatabase.Contract.GetUint(&_PlugDatabase.CallOpts, sender, key)
}

// IntStorage is a free data retrieval call binding the contract method 0x20feaf5e.
//
// Solidity: function intStorage(address , bytes32 ) view returns(int256)
func (_PlugDatabase *PlugDatabaseCaller) IntStorage(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "intStorage", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntStorage is a free data retrieval call binding the contract method 0x20feaf5e.
//
// Solidity: function intStorage(address , bytes32 ) view returns(int256)
func (_PlugDatabase *PlugDatabaseSession) IntStorage(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _PlugDatabase.Contract.IntStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// IntStorage is a free data retrieval call binding the contract method 0x20feaf5e.
//
// Solidity: function intStorage(address , bytes32 ) view returns(int256)
func (_PlugDatabase *PlugDatabaseCallerSession) IntStorage(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _PlugDatabase.Contract.IntStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// StringStorage is a free data retrieval call binding the contract method 0xef6e9be8.
//
// Solidity: function stringStorage(address , bytes32 ) view returns(string)
func (_PlugDatabase *PlugDatabaseCaller) StringStorage(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (string, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "stringStorage", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StringStorage is a free data retrieval call binding the contract method 0xef6e9be8.
//
// Solidity: function stringStorage(address , bytes32 ) view returns(string)
func (_PlugDatabase *PlugDatabaseSession) StringStorage(arg0 common.Address, arg1 [32]byte) (string, error) {
	return _PlugDatabase.Contract.StringStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// StringStorage is a free data retrieval call binding the contract method 0xef6e9be8.
//
// Solidity: function stringStorage(address , bytes32 ) view returns(string)
func (_PlugDatabase *PlugDatabaseCallerSession) StringStorage(arg0 common.Address, arg1 [32]byte) (string, error) {
	return _PlugDatabase.Contract.StringStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// UintStorage is a free data retrieval call binding the contract method 0x354d56da.
//
// Solidity: function uintStorage(address , bytes32 ) view returns(uint256)
func (_PlugDatabase *PlugDatabaseCaller) UintStorage(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PlugDatabase.contract.Call(opts, &out, "uintStorage", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UintStorage is a free data retrieval call binding the contract method 0x354d56da.
//
// Solidity: function uintStorage(address , bytes32 ) view returns(uint256)
func (_PlugDatabase *PlugDatabaseSession) UintStorage(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _PlugDatabase.Contract.UintStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// UintStorage is a free data retrieval call binding the contract method 0x354d56da.
//
// Solidity: function uintStorage(address , bytes32 ) view returns(uint256)
func (_PlugDatabase *PlugDatabaseCallerSession) UintStorage(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _PlugDatabase.Contract.UintStorage(&_PlugDatabase.CallOpts, arg0, arg1)
}

// BatchSet is a paid mutator transaction binding the contract method 0x8e94d30b.
//
// Solidity: function batchSet(bytes32[] keys, bytes32[] values, uint8 typeId) returns()
func (_PlugDatabase *PlugDatabaseTransactor) BatchSet(opts *bind.TransactOpts, keys [][32]byte, values [][32]byte, typeId uint8) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "batchSet", keys, values, typeId)
}

// BatchSet is a paid mutator transaction binding the contract method 0x8e94d30b.
//
// Solidity: function batchSet(bytes32[] keys, bytes32[] values, uint8 typeId) returns()
func (_PlugDatabase *PlugDatabaseSession) BatchSet(keys [][32]byte, values [][32]byte, typeId uint8) (*types.Transaction, error) {
	return _PlugDatabase.Contract.BatchSet(&_PlugDatabase.TransactOpts, keys, values, typeId)
}

// BatchSet is a paid mutator transaction binding the contract method 0x8e94d30b.
//
// Solidity: function batchSet(bytes32[] keys, bytes32[] values, uint8 typeId) returns()
func (_PlugDatabase *PlugDatabaseTransactorSession) BatchSet(keys [][32]byte, values [][32]byte, typeId uint8) (*types.Transaction, error) {
	return _PlugDatabase.Contract.BatchSet(&_PlugDatabase.TransactOpts, keys, values, typeId)
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

// RemoveWithType is a paid mutator transaction binding the contract method 0xf3606f4f.
//
// Solidity: function removeWithType(bytes32 key, uint8 typeId) returns()
func (_PlugDatabase *PlugDatabaseTransactor) RemoveWithType(opts *bind.TransactOpts, key [32]byte, typeId uint8) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "removeWithType", key, typeId)
}

// RemoveWithType is a paid mutator transaction binding the contract method 0xf3606f4f.
//
// Solidity: function removeWithType(bytes32 key, uint8 typeId) returns()
func (_PlugDatabase *PlugDatabaseSession) RemoveWithType(key [32]byte, typeId uint8) (*types.Transaction, error) {
	return _PlugDatabase.Contract.RemoveWithType(&_PlugDatabase.TransactOpts, key, typeId)
}

// RemoveWithType is a paid mutator transaction binding the contract method 0xf3606f4f.
//
// Solidity: function removeWithType(bytes32 key, uint8 typeId) returns()
func (_PlugDatabase *PlugDatabaseTransactorSession) RemoveWithType(key [32]byte, typeId uint8) (*types.Transaction, error) {
	return _PlugDatabase.Contract.RemoveWithType(&_PlugDatabase.TransactOpts, key, typeId)
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

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 key, address value) returns(address result)
func (_PlugDatabase *PlugDatabaseTransactor) SetAddress(opts *bind.TransactOpts, key [32]byte, value common.Address) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "setAddress", key, value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 key, address value) returns(address result)
func (_PlugDatabase *PlugDatabaseSession) SetAddress(key [32]byte, value common.Address) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetAddress(&_PlugDatabase.TransactOpts, key, value)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 key, address value) returns(address result)
func (_PlugDatabase *PlugDatabaseTransactorSession) SetAddress(key [32]byte, value common.Address) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetAddress(&_PlugDatabase.TransactOpts, key, value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 key, bool value) returns(bool result)
func (_PlugDatabase *PlugDatabaseTransactor) SetBool(opts *bind.TransactOpts, key [32]byte, value bool) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "setBool", key, value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 key, bool value) returns(bool result)
func (_PlugDatabase *PlugDatabaseSession) SetBool(key [32]byte, value bool) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetBool(&_PlugDatabase.TransactOpts, key, value)
}

// SetBool is a paid mutator transaction binding the contract method 0xabfdcced.
//
// Solidity: function setBool(bytes32 key, bool value) returns(bool result)
func (_PlugDatabase *PlugDatabaseTransactorSession) SetBool(key [32]byte, value bool) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetBool(&_PlugDatabase.TransactOpts, key, value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(bytes32 key, bytes value) returns(bytes result)
func (_PlugDatabase *PlugDatabaseTransactor) SetBytes(opts *bind.TransactOpts, key [32]byte, value []byte) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "setBytes", key, value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(bytes32 key, bytes value) returns(bytes result)
func (_PlugDatabase *PlugDatabaseSession) SetBytes(key [32]byte, value []byte) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetBytes(&_PlugDatabase.TransactOpts, key, value)
}

// SetBytes is a paid mutator transaction binding the contract method 0x2e28d084.
//
// Solidity: function setBytes(bytes32 key, bytes value) returns(bytes result)
func (_PlugDatabase *PlugDatabaseTransactorSession) SetBytes(key [32]byte, value []byte) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetBytes(&_PlugDatabase.TransactOpts, key, value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 key, int256 value) returns(int256 result)
func (_PlugDatabase *PlugDatabaseTransactor) SetInt(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "setInt", key, value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 key, int256 value) returns(int256 result)
func (_PlugDatabase *PlugDatabaseSession) SetInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetInt(&_PlugDatabase.TransactOpts, key, value)
}

// SetInt is a paid mutator transaction binding the contract method 0x3e49bed0.
//
// Solidity: function setInt(bytes32 key, int256 value) returns(int256 result)
func (_PlugDatabase *PlugDatabaseTransactorSession) SetInt(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetInt(&_PlugDatabase.TransactOpts, key, value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 key, string value) returns(string result)
func (_PlugDatabase *PlugDatabaseTransactor) SetString(opts *bind.TransactOpts, key [32]byte, value string) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "setString", key, value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 key, string value) returns(string result)
func (_PlugDatabase *PlugDatabaseSession) SetString(key [32]byte, value string) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetString(&_PlugDatabase.TransactOpts, key, value)
}

// SetString is a paid mutator transaction binding the contract method 0x6e899550.
//
// Solidity: function setString(bytes32 key, string value) returns(string result)
func (_PlugDatabase *PlugDatabaseTransactorSession) SetString(key [32]byte, value string) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetString(&_PlugDatabase.TransactOpts, key, value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 key, uint256 value) returns(uint256 result)
func (_PlugDatabase *PlugDatabaseTransactor) SetUint(opts *bind.TransactOpts, key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _PlugDatabase.contract.Transact(opts, "setUint", key, value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 key, uint256 value) returns(uint256 result)
func (_PlugDatabase *PlugDatabaseSession) SetUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetUint(&_PlugDatabase.TransactOpts, key, value)
}

// SetUint is a paid mutator transaction binding the contract method 0xe2a4853a.
//
// Solidity: function setUint(bytes32 key, uint256 value) returns(uint256 result)
func (_PlugDatabase *PlugDatabaseTransactorSession) SetUint(key [32]byte, value *big.Int) (*types.Transaction, error) {
	return _PlugDatabase.Contract.SetUint(&_PlugDatabase.TransactOpts, key, value)
}
