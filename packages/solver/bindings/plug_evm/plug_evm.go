// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package plug_evm

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

// PlugEvmMetaData contains all meta data concerning the PlugEvm contract.
var PlugEvmMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockHash\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockHash\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCaller\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCodeHash\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGasPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRemainingGas\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isContract\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"}]",
}

// PlugEvmABI is the input ABI used to generate the binding from.
// Deprecated: Use PlugEvmMetaData.ABI instead.
var PlugEvmABI = PlugEvmMetaData.ABI

// PlugEvm is an auto generated Go binding around an Ethereum contract.
type PlugEvm struct {
	PlugEvmCaller     // Read-only binding to the contract
	PlugEvmTransactor // Write-only binding to the contract
	PlugEvmFilterer   // Log filterer for contract events
}

// PlugEvmCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlugEvmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugEvmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlugEvmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugEvmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlugEvmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlugEvmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlugEvmSession struct {
	Contract     *PlugEvm          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlugEvmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlugEvmCallerSession struct {
	Contract *PlugEvmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PlugEvmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlugEvmTransactorSession struct {
	Contract     *PlugEvmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PlugEvmRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlugEvmRaw struct {
	Contract *PlugEvm // Generic contract binding to access the raw methods on
}

// PlugEvmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlugEvmCallerRaw struct {
	Contract *PlugEvmCaller // Generic read-only contract binding to access the raw methods on
}

// PlugEvmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlugEvmTransactorRaw struct {
	Contract *PlugEvmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlugEvm creates a new instance of PlugEvm, bound to a specific deployed contract.
func NewPlugEvm(address common.Address, backend bind.ContractBackend) (*PlugEvm, error) {
	contract, err := bindPlugEvm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlugEvm{PlugEvmCaller: PlugEvmCaller{contract: contract}, PlugEvmTransactor: PlugEvmTransactor{contract: contract}, PlugEvmFilterer: PlugEvmFilterer{contract: contract}}, nil
}

// NewPlugEvmCaller creates a new read-only instance of PlugEvm, bound to a specific deployed contract.
func NewPlugEvmCaller(address common.Address, caller bind.ContractCaller) (*PlugEvmCaller, error) {
	contract, err := bindPlugEvm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlugEvmCaller{contract: contract}, nil
}

// NewPlugEvmTransactor creates a new write-only instance of PlugEvm, bound to a specific deployed contract.
func NewPlugEvmTransactor(address common.Address, transactor bind.ContractTransactor) (*PlugEvmTransactor, error) {
	contract, err := bindPlugEvm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlugEvmTransactor{contract: contract}, nil
}

// NewPlugEvmFilterer creates a new log filterer instance of PlugEvm, bound to a specific deployed contract.
func NewPlugEvmFilterer(address common.Address, filterer bind.ContractFilterer) (*PlugEvmFilterer, error) {
	contract, err := bindPlugEvm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlugEvmFilterer{contract: contract}, nil
}

// bindPlugEvm binds a generic wrapper to an already deployed contract.
func bindPlugEvm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlugEvmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugEvm *PlugEvmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugEvm.Contract.PlugEvmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugEvm *PlugEvmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugEvm.Contract.PlugEvmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugEvm *PlugEvmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugEvm.Contract.PlugEvmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlugEvm *PlugEvmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PlugEvm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlugEvm *PlugEvmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugEvm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlugEvm *PlugEvmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlugEvm.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256 result)
func (_PlugEvm *PlugEvmCaller) GetBalance(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getBalance", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256 result)
func (_PlugEvm *PlugEvmSession) GetBalance(account common.Address) (*big.Int, error) {
	return _PlugEvm.Contract.GetBalance(&_PlugEvm.CallOpts, account)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address account) view returns(uint256 result)
func (_PlugEvm *PlugEvmCallerSession) GetBalance(account common.Address) (*big.Int, error) {
	return _PlugEvm.Contract.GetBalance(&_PlugEvm.CallOpts, account)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint256 result)
func (_PlugEvm *PlugEvmCaller) GetBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint256 result)
func (_PlugEvm *PlugEvmSession) GetBlockGasLimit() (*big.Int, error) {
	return _PlugEvm.Contract.GetBlockGasLimit(&_PlugEvm.CallOpts)
}

// GetBlockGasLimit is a free data retrieval call binding the contract method 0x2cc8377d.
//
// Solidity: function getBlockGasLimit() view returns(uint256 result)
func (_PlugEvm *PlugEvmCallerSession) GetBlockGasLimit() (*big.Int, error) {
	return _PlugEvm.Contract.GetBlockGasLimit(&_PlugEvm.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0x9663f88f.
//
// Solidity: function getBlockHash() view returns(bytes32 result)
func (_PlugEvm *PlugEvmCaller) GetBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0x9663f88f.
//
// Solidity: function getBlockHash() view returns(bytes32 result)
func (_PlugEvm *PlugEvmSession) GetBlockHash() ([32]byte, error) {
	return _PlugEvm.Contract.GetBlockHash(&_PlugEvm.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0x9663f88f.
//
// Solidity: function getBlockHash() view returns(bytes32 result)
func (_PlugEvm *PlugEvmCallerSession) GetBlockHash() ([32]byte, error) {
	return _PlugEvm.Contract.GetBlockHash(&_PlugEvm.CallOpts)
}

// GetBlockHash0 is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 result)
func (_PlugEvm *PlugEvmCaller) GetBlockHash0(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getBlockHash0", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash0 is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 result)
func (_PlugEvm *PlugEvmSession) GetBlockHash0(blockNumber *big.Int) ([32]byte, error) {
	return _PlugEvm.Contract.GetBlockHash0(&_PlugEvm.CallOpts, blockNumber)
}

// GetBlockHash0 is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 result)
func (_PlugEvm *PlugEvmCallerSession) GetBlockHash0(blockNumber *big.Int) ([32]byte, error) {
	return _PlugEvm.Contract.GetBlockHash0(&_PlugEvm.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 result)
func (_PlugEvm *PlugEvmCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 result)
func (_PlugEvm *PlugEvmSession) GetBlockNumber() (*big.Int, error) {
	return _PlugEvm.Contract.GetBlockNumber(&_PlugEvm.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 result)
func (_PlugEvm *PlugEvmCallerSession) GetBlockNumber() (*big.Int, error) {
	return _PlugEvm.Contract.GetBlockNumber(&_PlugEvm.CallOpts)
}

// GetCaller is a free data retrieval call binding the contract method 0xab470f05.
//
// Solidity: function getCaller() view returns(address result)
func (_PlugEvm *PlugEvmCaller) GetCaller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getCaller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCaller is a free data retrieval call binding the contract method 0xab470f05.
//
// Solidity: function getCaller() view returns(address result)
func (_PlugEvm *PlugEvmSession) GetCaller() (common.Address, error) {
	return _PlugEvm.Contract.GetCaller(&_PlugEvm.CallOpts)
}

// GetCaller is a free data retrieval call binding the contract method 0xab470f05.
//
// Solidity: function getCaller() view returns(address result)
func (_PlugEvm *PlugEvmCallerSession) GetCaller() (common.Address, error) {
	return _PlugEvm.Contract.GetCaller(&_PlugEvm.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 result)
func (_PlugEvm *PlugEvmCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 result)
func (_PlugEvm *PlugEvmSession) GetChainId() (*big.Int, error) {
	return _PlugEvm.Contract.GetChainId(&_PlugEvm.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 result)
func (_PlugEvm *PlugEvmCallerSession) GetChainId() (*big.Int, error) {
	return _PlugEvm.Contract.GetChainId(&_PlugEvm.CallOpts)
}

// GetCodeHash is a free data retrieval call binding the contract method 0x81ea4408.
//
// Solidity: function getCodeHash(address account) view returns(bytes32 result)
func (_PlugEvm *PlugEvmCaller) GetCodeHash(opts *bind.CallOpts, account common.Address) ([32]byte, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getCodeHash", account)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCodeHash is a free data retrieval call binding the contract method 0x81ea4408.
//
// Solidity: function getCodeHash(address account) view returns(bytes32 result)
func (_PlugEvm *PlugEvmSession) GetCodeHash(account common.Address) ([32]byte, error) {
	return _PlugEvm.Contract.GetCodeHash(&_PlugEvm.CallOpts, account)
}

// GetCodeHash is a free data retrieval call binding the contract method 0x81ea4408.
//
// Solidity: function getCodeHash(address account) view returns(bytes32 result)
func (_PlugEvm *PlugEvmCallerSession) GetCodeHash(account common.Address) ([32]byte, error) {
	return _PlugEvm.Contract.GetCodeHash(&_PlugEvm.CallOpts, account)
}

// GetGasPrice is a free data retrieval call binding the contract method 0x455259cb.
//
// Solidity: function getGasPrice() view returns(uint256 result)
func (_PlugEvm *PlugEvmCaller) GetGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGasPrice is a free data retrieval call binding the contract method 0x455259cb.
//
// Solidity: function getGasPrice() view returns(uint256 result)
func (_PlugEvm *PlugEvmSession) GetGasPrice() (*big.Int, error) {
	return _PlugEvm.Contract.GetGasPrice(&_PlugEvm.CallOpts)
}

// GetGasPrice is a free data retrieval call binding the contract method 0x455259cb.
//
// Solidity: function getGasPrice() view returns(uint256 result)
func (_PlugEvm *PlugEvmCallerSession) GetGasPrice() (*big.Int, error) {
	return _PlugEvm.Contract.GetGasPrice(&_PlugEvm.CallOpts)
}

// GetRemainingGas is a free data retrieval call binding the contract method 0x53868490.
//
// Solidity: function getRemainingGas() view returns(uint256 result)
func (_PlugEvm *PlugEvmCaller) GetRemainingGas(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getRemainingGas")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingGas is a free data retrieval call binding the contract method 0x53868490.
//
// Solidity: function getRemainingGas() view returns(uint256 result)
func (_PlugEvm *PlugEvmSession) GetRemainingGas() (*big.Int, error) {
	return _PlugEvm.Contract.GetRemainingGas(&_PlugEvm.CallOpts)
}

// GetRemainingGas is a free data retrieval call binding the contract method 0x53868490.
//
// Solidity: function getRemainingGas() view returns(uint256 result)
func (_PlugEvm *PlugEvmCallerSession) GetRemainingGas() (*big.Int, error) {
	return _PlugEvm.Contract.GetRemainingGas(&_PlugEvm.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256 result)
func (_PlugEvm *PlugEvmCaller) GetTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "getTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256 result)
func (_PlugEvm *PlugEvmSession) GetTimestamp() (*big.Int, error) {
	return _PlugEvm.Contract.GetTimestamp(&_PlugEvm.CallOpts)
}

// GetTimestamp is a free data retrieval call binding the contract method 0x188ec356.
//
// Solidity: function getTimestamp() view returns(uint256 result)
func (_PlugEvm *PlugEvmCallerSession) GetTimestamp() (*big.Int, error) {
	return _PlugEvm.Contract.GetTimestamp(&_PlugEvm.CallOpts)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address account) view returns(bool result)
func (_PlugEvm *PlugEvmCaller) IsContract(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _PlugEvm.contract.Call(opts, &out, "isContract", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address account) view returns(bool result)
func (_PlugEvm *PlugEvmSession) IsContract(account common.Address) (bool, error) {
	return _PlugEvm.Contract.IsContract(&_PlugEvm.CallOpts, account)
}

// IsContract is a free data retrieval call binding the contract method 0x16279055.
//
// Solidity: function isContract(address account) view returns(bool result)
func (_PlugEvm *PlugEvmCallerSession) IsContract(account common.Address) (bool, error) {
	return _PlugEvm.Contract.IsContract(&_PlugEvm.CallOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address recipient, uint256 amount, uint256 gasLimit) payable returns(bool success, bytes returnData)
func (_PlugEvm *PlugEvmTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _PlugEvm.contract.Transact(opts, "transfer", recipient, amount, gasLimit)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address recipient, uint256 amount, uint256 gasLimit) payable returns(bool success, bytes returnData)
func (_PlugEvm *PlugEvmSession) Transfer(recipient common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _PlugEvm.Contract.Transfer(&_PlugEvm.TransactOpts, recipient, amount, gasLimit)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address recipient, uint256 amount, uint256 gasLimit) payable returns(bool success, bytes returnData)
func (_PlugEvm *PlugEvmTransactorSession) Transfer(recipient common.Address, amount *big.Int, gasLimit *big.Int) (*types.Transaction, error) {
	return _PlugEvm.Contract.Transfer(&_PlugEvm.TransactOpts, recipient, amount, gasLimit)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) payable returns()
func (_PlugEvm *PlugEvmTransactor) Transfer0(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlugEvm.contract.Transact(opts, "transfer0", recipient, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) payable returns()
func (_PlugEvm *PlugEvmSession) Transfer0(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlugEvm.Contract.Transfer0(&_PlugEvm.TransactOpts, recipient, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) payable returns()
func (_PlugEvm *PlugEvmTransactorSession) Transfer0(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PlugEvm.Contract.Transfer0(&_PlugEvm.TransactOpts, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PlugEvm *PlugEvmTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlugEvm.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PlugEvm *PlugEvmSession) Receive() (*types.Transaction, error) {
	return _PlugEvm.Contract.Receive(&_PlugEvm.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PlugEvm *PlugEvmTransactorSession) Receive() (*types.Transaction, error) {
	return _PlugEvm.Contract.Receive(&_PlugEvm.TransactOpts)
}
