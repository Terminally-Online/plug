// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_router

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

// PlugTypesLibLivePlugs is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibLivePlugs struct {
	Plugs     PlugTypesLibPlugs
	Signature []byte
}

// PlugTypesLibPlug is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibPlug struct {
	To    common.Address
	Data  []byte
	Value *big.Int
	Gas   *big.Int
}

// PlugTypesLibPlugs is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibPlugs struct {
	Socket common.Address
	Plugs  []PlugTypesLibPlug
	Solver []byte
	Salt   []byte
}

// PlugTypesLibResult is an auto generated low-level Go binding around an user-defined struct.
type PlugTypesLibResult struct {
	Index uint8
	Error string
}

// PlugRouterMetaData contains all meta data concerning the PlugRouter contract.
var PlugRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"plug\",\"inputs\":[{\"name\":\"livePlugs\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.LivePlugs\",\"components\":[{\"name\":\"plugs\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plugs\",\"components\":[{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"plugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"solver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Result\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"plug\",\"inputs\":[{\"name\":\"livePlugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.LivePlugs[]\",\"components\":[{\"name\":\"plugs\",\"type\":\"tuple\",\"internalType\":\"structPlugTypesLib.Plugs\",\"components\":[{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"plugs\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Plug[]\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"solver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"results\",\"type\":\"tuple[]\",\"internalType\":\"structPlugTypesLib.Result[]\",\"components\":[{\"name\":\"index\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"error\",\"name\":\"SocketAddressInvalid\",\"inputs\":[{\"name\":\"intended\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// PlugRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugRouterMetaData.ABI instead.
var PlugRouterABI = PlugRouterMetaData.ABI

// PlugRouter is an auto generated Go binding around an Ethereum contract.
type PlugRouter struct {
	PlugRouterCaller     // Read-only binding to the contract
	PlugRouterTransactor // Write-only binding to the contract
	PlugRouterFilterer   // Log filterer for contract events
}

// PlugRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugRouterSession struct {
	Contract     *PlugRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugRouterCallerSession struct {
	Contract *PlugRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PlugRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugRouterTransactorSession struct {
	Contract     *PlugRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PlugRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugRouterRaw struct {
	Contract *PlugRouter // Generic contract binding to access the raw methods on
}

// PlugRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugRouterCallerRaw struct {
	Contract *PlugRouterCaller // Generic read-only contract binding to access the raw methods on
}

// PlugRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugRouterTransactorRaw struct {
	Contract *PlugRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugRouter creates a new instance of PlugRouter, bound to a specific deployed contract.
func NewPlugRouter(address common.Address, backend bind.ContractBackend) (*PlugRouter, error) {
	contract, err := bindPlugRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugRouter{PlugRouterCaller: PlugRouterCaller{contract: contract}, PlugRouterTransactor: PlugRouterTransactor{contract: contract}, PlugRouterFilterer: PlugRouterFilterer{contract: contract}}, nil
}

// NewPlugRouterCaller creates a new read-only instance of PlugRouter, bound to a specific deployed contract.
func NewPlugRouterCaller(address common.Address, caller bind.ContractCaller) (*PlugRouterCaller, error) {
	contract, err := bindPlugRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugRouterCaller{contract: contract}, nil
}

// NewPlugRouterTransactor creates a new write-only instance of PlugRouter, bound to a specific deployed contract.
func NewPlugRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugRouterTransactor, error) {
	contract, err := bindPlugRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugRouterTransactor{contract: contract}, nil
}

// NewPlugRouterFilterer creates a new log filterer instance of PlugRouter, bound to a specific deployed contract.
func NewPlugRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugRouterFilterer, error) {
	contract, err := bindPlugRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugRouterFilterer{contract: contract}, nil
}

// bindPlugRouter binds a generic wrapper to an already deployed contract.
func bindPlugRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugRouter *PlugRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugRouter.Contract.PlugRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugRouter *PlugRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugRouter.Contract.PlugRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugRouter *PlugRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugRouter.Contract.PlugRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugRouter *PlugRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugRouter *PlugRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugRouter *PlugRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugRouter.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string name)
func (_PlugRouter *PlugRouterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlugRouter.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string name)
func (_PlugRouter *PlugRouterSession) Name() (string, error) {
	return _PlugRouter.Contract.Name(&_PlugRouter.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string name)
func (_PlugRouter *PlugRouterCallerSession) Name() (string, error) {
	return _PlugRouter.Contract.Name(&_PlugRouter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string version)
func (_PlugRouter *PlugRouterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PlugRouter.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string version)
func (_PlugRouter *PlugRouterSession) Symbol() (string, error) {
	return _PlugRouter.Contract.Symbol(&_PlugRouter.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string version)
func (_PlugRouter *PlugRouterCallerSession) Symbol() (string, error) {
	return _PlugRouter.Contract.Symbol(&_PlugRouter.CallOpts)
}

// Plug is a paid mutator transaction binding the contract method 0x5365d433.
//
// Solidity: function plug(((address,(address,bytes,uint256,uint256)[],bytes,bytes),bytes) livePlugs) payable returns((uint8,string) results)
func (_PlugRouter *PlugRouterTransactor) Plug(opts *bind.TransactOpts, livePlugs PlugTypesLibLivePlugs) (*types.Transaction, error) {
	return _PlugRouter.contract.Transact(opts, "plug", livePlugs)
}

// Plug is a paid mutator transaction binding the contract method 0x5365d433.
//
// Solidity: function plug(((address,(address,bytes,uint256,uint256)[],bytes,bytes),bytes) livePlugs) payable returns((uint8,string) results)
func (_PlugRouter *PlugRouterSession) Plug(livePlugs PlugTypesLibLivePlugs) (*types.Transaction, error) {
	return _PlugRouter.Contract.Plug(&_PlugRouter.TransactOpts, livePlugs)
}

// Plug is a paid mutator transaction binding the contract method 0x5365d433.
//
// Solidity: function plug(((address,(address,bytes,uint256,uint256)[],bytes,bytes),bytes) livePlugs) payable returns((uint8,string) results)
func (_PlugRouter *PlugRouterTransactorSession) Plug(livePlugs PlugTypesLibLivePlugs) (*types.Transaction, error) {
	return _PlugRouter.Contract.Plug(&_PlugRouter.TransactOpts, livePlugs)
}

// Plug0 is a paid mutator transaction binding the contract method 0xf17ae13f.
//
// Solidity: function plug(((address,(address,bytes,uint256,uint256)[],bytes,bytes),bytes)[] livePlugs) payable returns((uint8,string)[] results)
func (_PlugRouter *PlugRouterTransactor) Plug0(opts *bind.TransactOpts, livePlugs []PlugTypesLibLivePlugs) (*types.Transaction, error) {
	return _PlugRouter.contract.Transact(opts, "plug0", livePlugs)
}

// Plug0 is a paid mutator transaction binding the contract method 0xf17ae13f.
//
// Solidity: function plug(((address,(address,bytes,uint256,uint256)[],bytes,bytes),bytes)[] livePlugs) payable returns((uint8,string)[] results)
func (_PlugRouter *PlugRouterSession) Plug0(livePlugs []PlugTypesLibLivePlugs) (*types.Transaction, error) {
	return _PlugRouter.Contract.Plug0(&_PlugRouter.TransactOpts, livePlugs)
}

// Plug0 is a paid mutator transaction binding the contract method 0xf17ae13f.
//
// Solidity: function plug(((address,(address,bytes,uint256,uint256)[],bytes,bytes),bytes)[] livePlugs) payable returns((uint8,string)[] results)
func (_PlugRouter *PlugRouterTransactorSession) Plug0(livePlugs []PlugTypesLibLivePlugs) (*types.Transaction, error) {
	return _PlugRouter.Contract.Plug0(&_PlugRouter.TransactOpts, livePlugs)
}
