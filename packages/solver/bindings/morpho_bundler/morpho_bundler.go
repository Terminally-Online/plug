// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package morpho_bundler

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

// Call is an auto generated low-level Go binding around an user-defined struct.
type Call struct {
	To           common.Address
	Data         []byte
	Value        *big.Int
	SkipRevert   bool
	CallbackHash [32]byte
}

// MorphoBundlerMetaData contains all meta data concerning the MorphoBundler contract.
var MorphoBundlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitiated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBundle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectReenterHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingExpectedReenter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"initiator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"skipRevert\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"callbackHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCall[]\",\"name\":\"bundle\",\"type\":\"tuple[]\"}],\"name\":\"multicall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"skipRevert\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"callbackHash\",\"type\":\"bytes32\"}],\"internalType\":\"structCall[]\",\"name\":\"bundle\",\"type\":\"tuple[]\"}],\"name\":\"reenter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reenterHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MorphoBundlerABI is the input ABI used to generate the binding from.
// Deprecated: Use MorphoBundlerMetaData.ABI instead.
var MorphoBundlerABI = MorphoBundlerMetaData.ABI

// MorphoBundler is an auto generated Go binding around an Ethereum contract.
type MorphoBundler struct {
	MorphoBundlerCaller     // Read-only binding to the contract
	MorphoBundlerTransactor // Write-only binding to the contract
	MorphoBundlerFilterer   // Log filterer for contract events
}

// MorphoBundlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MorphoBundlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoBundlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MorphoBundlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoBundlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MorphoBundlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MorphoBundlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MorphoBundlerSession struct {
	Contract     *MorphoBundler    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MorphoBundlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MorphoBundlerCallerSession struct {
	Contract *MorphoBundlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MorphoBundlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MorphoBundlerTransactorSession struct {
	Contract     *MorphoBundlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MorphoBundlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MorphoBundlerRaw struct {
	Contract *MorphoBundler // Generic contract binding to access the raw methods on
}

// MorphoBundlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MorphoBundlerCallerRaw struct {
	Contract *MorphoBundlerCaller // Generic read-only contract binding to access the raw methods on
}

// MorphoBundlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MorphoBundlerTransactorRaw struct {
	Contract *MorphoBundlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMorphoBundler creates a new instance of MorphoBundler, bound to a specific deployed contract.
func NewMorphoBundler(address common.Address, backend bind.ContractBackend) (*MorphoBundler, error) {
	contract, err := bindMorphoBundler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MorphoBundler{MorphoBundlerCaller: MorphoBundlerCaller{contract: contract}, MorphoBundlerTransactor: MorphoBundlerTransactor{contract: contract}, MorphoBundlerFilterer: MorphoBundlerFilterer{contract: contract}}, nil
}

// NewMorphoBundlerCaller creates a new read-only instance of MorphoBundler, bound to a specific deployed contract.
func NewMorphoBundlerCaller(address common.Address, caller bind.ContractCaller) (*MorphoBundlerCaller, error) {
	contract, err := bindMorphoBundler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoBundlerCaller{contract: contract}, nil
}

// NewMorphoBundlerTransactor creates a new write-only instance of MorphoBundler, bound to a specific deployed contract.
func NewMorphoBundlerTransactor(address common.Address, transactor bind.ContractTransactor) (*MorphoBundlerTransactor, error) {
	contract, err := bindMorphoBundler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MorphoBundlerTransactor{contract: contract}, nil
}

// NewMorphoBundlerFilterer creates a new log filterer instance of MorphoBundler, bound to a specific deployed contract.
func NewMorphoBundlerFilterer(address common.Address, filterer bind.ContractFilterer) (*MorphoBundlerFilterer, error) {
	contract, err := bindMorphoBundler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MorphoBundlerFilterer{contract: contract}, nil
}

// bindMorphoBundler binds a generic wrapper to an already deployed contract.
func bindMorphoBundler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MorphoBundlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoBundler *MorphoBundlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoBundler.Contract.MorphoBundlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoBundler *MorphoBundlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoBundlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoBundler *MorphoBundlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoBundler.Contract.MorphoBundlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MorphoBundler *MorphoBundlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MorphoBundler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MorphoBundler *MorphoBundlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MorphoBundler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MorphoBundler *MorphoBundlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MorphoBundler.Contract.contract.Transact(opts, method, params...)
}

// Initiator is a free data retrieval call binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() view returns(address)
func (_MorphoBundler *MorphoBundlerCaller) Initiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MorphoBundler.contract.Call(opts, &out, "initiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Initiator is a free data retrieval call binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() view returns(address)
func (_MorphoBundler *MorphoBundlerSession) Initiator() (common.Address, error) {
	return _MorphoBundler.Contract.Initiator(&_MorphoBundler.CallOpts)
}

// Initiator is a free data retrieval call binding the contract method 0x5c39fcc1.
//
// Solidity: function initiator() view returns(address)
func (_MorphoBundler *MorphoBundlerCallerSession) Initiator() (common.Address, error) {
	return _MorphoBundler.Contract.Initiator(&_MorphoBundler.CallOpts)
}

// ReenterHash is a free data retrieval call binding the contract method 0xe69fe134.
//
// Solidity: function reenterHash() view returns(bytes32)
func (_MorphoBundler *MorphoBundlerCaller) ReenterHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MorphoBundler.contract.Call(opts, &out, "reenterHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReenterHash is a free data retrieval call binding the contract method 0xe69fe134.
//
// Solidity: function reenterHash() view returns(bytes32)
func (_MorphoBundler *MorphoBundlerSession) ReenterHash() ([32]byte, error) {
	return _MorphoBundler.Contract.ReenterHash(&_MorphoBundler.CallOpts)
}

// ReenterHash is a free data retrieval call binding the contract method 0xe69fe134.
//
// Solidity: function reenterHash() view returns(bytes32)
func (_MorphoBundler *MorphoBundlerCallerSession) ReenterHash() ([32]byte, error) {
	return _MorphoBundler.Contract.ReenterHash(&_MorphoBundler.CallOpts)
}

// Multicall is a paid mutator transaction binding the contract method 0x374f435d.
//
// Solidity: function multicall((address,bytes,uint256,bool,bytes32)[] bundle) payable returns()
func (_MorphoBundler *MorphoBundlerTransactor) Multicall(opts *bind.TransactOpts, bundle []Call) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "multicall", bundle)
}

// Multicall is a paid mutator transaction binding the contract method 0x374f435d.
//
// Solidity: function multicall((address,bytes,uint256,bool,bytes32)[] bundle) payable returns()
func (_MorphoBundler *MorphoBundlerSession) Multicall(bundle []Call) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Multicall(&_MorphoBundler.TransactOpts, bundle)
}

// Multicall is a paid mutator transaction binding the contract method 0x374f435d.
//
// Solidity: function multicall((address,bytes,uint256,bool,bytes32)[] bundle) payable returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Multicall(bundle []Call) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Multicall(&_MorphoBundler.TransactOpts, bundle)
}

// Reenter is a paid mutator transaction binding the contract method 0x803a7fba.
//
// Solidity: function reenter((address,bytes,uint256,bool,bytes32)[] bundle) returns()
func (_MorphoBundler *MorphoBundlerTransactor) Reenter(opts *bind.TransactOpts, bundle []Call) (*types.Transaction, error) {
	return _MorphoBundler.contract.Transact(opts, "reenter", bundle)
}

// Reenter is a paid mutator transaction binding the contract method 0x803a7fba.
//
// Solidity: function reenter((address,bytes,uint256,bool,bytes32)[] bundle) returns()
func (_MorphoBundler *MorphoBundlerSession) Reenter(bundle []Call) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Reenter(&_MorphoBundler.TransactOpts, bundle)
}

// Reenter is a paid mutator transaction binding the contract method 0x803a7fba.
//
// Solidity: function reenter((address,bytes,uint256,bool,bytes32)[] bundle) returns()
func (_MorphoBundler *MorphoBundlerTransactorSession) Reenter(bundle []Call) (*types.Transaction, error) {
	return _MorphoBundler.Contract.Reenter(&_MorphoBundler.TransactOpts, bundle)
}
