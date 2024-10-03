// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multicall_primary

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

// Multicall3Call is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call struct {
	Target   common.Address
	CallData []byte
}

// Multicall3Call3 is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3 struct {
	Target       common.Address
	AllowFailure bool
	CallData     []byte
}

// Multicall3Call3Value is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Call3Value struct {
	Target       common.Address
	AllowFailure bool
	Value        *big.Int
	CallData     []byte
}

// Multicall3Result is an auto generated low-level Go binding around an user-defined struct.
type Multicall3Result struct {
	Success    bool
	ReturnData []byte
}

// MulticallPrimaryMetaData contains all meta data concerning the MulticallPrimary contract.
var MulticallPrimaryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call3[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowFailure\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call3Value[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"aggregate3Value\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"blockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBasefee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"basefee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"difficulty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gaslimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEthBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryAggregate\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"requireSuccess\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Call[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"tryBlockAndAggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticall3.Result[]\",\"name\":\"returnData\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MulticallPrimaryABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallPrimaryMetaData.ABI instead.
var MulticallPrimaryABI = MulticallPrimaryMetaData.ABI

// MulticallPrimary is an auto generated Go binding around an Ethereum contract.
type MulticallPrimary struct {
	MulticallPrimaryCaller     // Read-only binding to the contract
	MulticallPrimaryTransactor // Write-only binding to the contract
	MulticallPrimaryFilterer   // Log filterer for contract events
}

// MulticallPrimaryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallPrimaryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallPrimaryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallPrimaryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallPrimaryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallPrimaryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallPrimarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallPrimarySession struct {
	Contract     *MulticallPrimary // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallPrimaryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallPrimaryCallerSession struct {
	Contract *MulticallPrimaryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MulticallPrimaryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallPrimaryTransactorSession struct {
	Contract     *MulticallPrimaryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MulticallPrimaryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallPrimaryRaw struct {
	Contract *MulticallPrimary // Generic contract binding to access the raw methods on
}

// MulticallPrimaryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallPrimaryCallerRaw struct {
	Contract *MulticallPrimaryCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallPrimaryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallPrimaryTransactorRaw struct {
	Contract *MulticallPrimaryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticallPrimary creates a new instance of MulticallPrimary, bound to a specific deployed contract.
func NewMulticallPrimary(address common.Address, backend bind.ContractBackend) (*MulticallPrimary, error) {
	contract, err := bindMulticallPrimary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MulticallPrimary{MulticallPrimaryCaller: MulticallPrimaryCaller{contract: contract}, MulticallPrimaryTransactor: MulticallPrimaryTransactor{contract: contract}, MulticallPrimaryFilterer: MulticallPrimaryFilterer{contract: contract}}, nil
}

// NewMulticallPrimaryCaller creates a new read-only instance of MulticallPrimary, bound to a specific deployed contract.
func NewMulticallPrimaryCaller(address common.Address, caller bind.ContractCaller) (*MulticallPrimaryCaller, error) {
	contract, err := bindMulticallPrimary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallPrimaryCaller{contract: contract}, nil
}

// NewMulticallPrimaryTransactor creates a new write-only instance of MulticallPrimary, bound to a specific deployed contract.
func NewMulticallPrimaryTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallPrimaryTransactor, error) {
	contract, err := bindMulticallPrimary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallPrimaryTransactor{contract: contract}, nil
}

// NewMulticallPrimaryFilterer creates a new log filterer instance of MulticallPrimary, bound to a specific deployed contract.
func NewMulticallPrimaryFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallPrimaryFilterer, error) {
	contract, err := bindMulticallPrimary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallPrimaryFilterer{contract: contract}, nil
}

// bindMulticallPrimary binds a generic wrapper to an already deployed contract.
func bindMulticallPrimary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MulticallPrimaryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallPrimary *MulticallPrimaryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallPrimary.Contract.MulticallPrimaryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallPrimary *MulticallPrimaryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.MulticallPrimaryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallPrimary *MulticallPrimaryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.MulticallPrimaryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallPrimary *MulticallPrimaryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallPrimary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallPrimary *MulticallPrimaryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallPrimary *MulticallPrimaryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.contract.Transact(opts, method, params...)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_MulticallPrimary *MulticallPrimaryCaller) GetBasefee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getBasefee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_MulticallPrimary *MulticallPrimarySession) GetBasefee() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetBasefee(&_MulticallPrimary.CallOpts)
}

// GetBasefee is a free data retrieval call binding the contract method 0x3e64a696.
//
// Solidity: function getBasefee() view returns(uint256 basefee)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetBasefee() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetBasefee(&_MulticallPrimary.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallPrimary *MulticallPrimaryCaller) GetBlockHash(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getBlockHash", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallPrimary *MulticallPrimarySession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MulticallPrimary.Contract.GetBlockHash(&_MulticallPrimary.CallOpts, blockNumber)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 blockNumber) view returns(bytes32 blockHash)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetBlockHash(blockNumber *big.Int) ([32]byte, error) {
	return _MulticallPrimary.Contract.GetBlockHash(&_MulticallPrimary.CallOpts, blockNumber)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallPrimary *MulticallPrimaryCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallPrimary *MulticallPrimarySession) GetBlockNumber() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetBlockNumber(&_MulticallPrimary.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 blockNumber)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetBlockNumber() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetBlockNumber(&_MulticallPrimary.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_MulticallPrimary *MulticallPrimaryCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_MulticallPrimary *MulticallPrimarySession) GetChainId() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetChainId(&_MulticallPrimary.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256 chainid)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetChainId() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetChainId(&_MulticallPrimary.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallPrimary *MulticallPrimaryCaller) GetCurrentBlockCoinbase(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getCurrentBlockCoinbase")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallPrimary *MulticallPrimarySession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MulticallPrimary.Contract.GetCurrentBlockCoinbase(&_MulticallPrimary.CallOpts)
}

// GetCurrentBlockCoinbase is a free data retrieval call binding the contract method 0xa8b0574e.
//
// Solidity: function getCurrentBlockCoinbase() view returns(address coinbase)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetCurrentBlockCoinbase() (common.Address, error) {
	return _MulticallPrimary.Contract.GetCurrentBlockCoinbase(&_MulticallPrimary.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallPrimary *MulticallPrimaryCaller) GetCurrentBlockDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getCurrentBlockDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallPrimary *MulticallPrimarySession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetCurrentBlockDifficulty(&_MulticallPrimary.CallOpts)
}

// GetCurrentBlockDifficulty is a free data retrieval call binding the contract method 0x72425d9d.
//
// Solidity: function getCurrentBlockDifficulty() view returns(uint256 difficulty)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetCurrentBlockDifficulty() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetCurrentBlockDifficulty(&_MulticallPrimary.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallPrimary *MulticallPrimaryCaller) GetCurrentBlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getCurrentBlockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallPrimary *MulticallPrimarySession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetCurrentBlockGasLimit(&_MulticallPrimary.CallOpts)
}

// GetCurrentBlockGasLimit is a free data retrieval call binding the contract method 0x86d516e8.
//
// Solidity: function getCurrentBlockGasLimit() view returns(uint256 gaslimit)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetCurrentBlockGasLimit() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetCurrentBlockGasLimit(&_MulticallPrimary.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallPrimary *MulticallPrimaryCaller) GetCurrentBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getCurrentBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallPrimary *MulticallPrimarySession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetCurrentBlockTimestamp(&_MulticallPrimary.CallOpts)
}

// GetCurrentBlockTimestamp is a free data retrieval call binding the contract method 0x0f28c97d.
//
// Solidity: function getCurrentBlockTimestamp() view returns(uint256 timestamp)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetCurrentBlockTimestamp() (*big.Int, error) {
	return _MulticallPrimary.Contract.GetCurrentBlockTimestamp(&_MulticallPrimary.CallOpts)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallPrimary *MulticallPrimaryCaller) GetEthBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getEthBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallPrimary *MulticallPrimarySession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MulticallPrimary.Contract.GetEthBalance(&_MulticallPrimary.CallOpts, addr)
}

// GetEthBalance is a free data retrieval call binding the contract method 0x4d2301cc.
//
// Solidity: function getEthBalance(address addr) view returns(uint256 balance)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetEthBalance(addr common.Address) (*big.Int, error) {
	return _MulticallPrimary.Contract.GetEthBalance(&_MulticallPrimary.CallOpts, addr)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallPrimary *MulticallPrimaryCaller) GetLastBlockHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MulticallPrimary.contract.Call(opts, &out, "getLastBlockHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallPrimary *MulticallPrimarySession) GetLastBlockHash() ([32]byte, error) {
	return _MulticallPrimary.Contract.GetLastBlockHash(&_MulticallPrimary.CallOpts)
}

// GetLastBlockHash is a free data retrieval call binding the contract method 0x27e86d6e.
//
// Solidity: function getLastBlockHash() view returns(bytes32 blockHash)
func (_MulticallPrimary *MulticallPrimaryCallerSession) GetLastBlockHash() ([32]byte, error) {
	return _MulticallPrimary.Contract.GetLastBlockHash(&_MulticallPrimary.CallOpts)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactor) Aggregate(opts *bind.TransactOpts, calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.contract.Transact(opts, "aggregate", calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallPrimary *MulticallPrimarySession) Aggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.Aggregate(&_MulticallPrimary.TransactOpts, calls)
}

// Aggregate is a paid mutator transaction binding the contract method 0x252dba42.
//
// Solidity: function aggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactorSession) Aggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.Aggregate(&_MulticallPrimary.TransactOpts, calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactor) Aggregate3(opts *bind.TransactOpts, calls []Multicall3Call3) (*types.Transaction, error) {
	return _MulticallPrimary.contract.Transact(opts, "aggregate3", calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimarySession) Aggregate3(calls []Multicall3Call3) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.Aggregate3(&_MulticallPrimary.TransactOpts, calls)
}

// Aggregate3 is a paid mutator transaction binding the contract method 0x82ad56cb.
//
// Solidity: function aggregate3((address,bool,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactorSession) Aggregate3(calls []Multicall3Call3) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.Aggregate3(&_MulticallPrimary.TransactOpts, calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactor) Aggregate3Value(opts *bind.TransactOpts, calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _MulticallPrimary.contract.Transact(opts, "aggregate3Value", calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimarySession) Aggregate3Value(calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.Aggregate3Value(&_MulticallPrimary.TransactOpts, calls)
}

// Aggregate3Value is a paid mutator transaction binding the contract method 0x174dea71.
//
// Solidity: function aggregate3Value((address,bool,uint256,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactorSession) Aggregate3Value(calls []Multicall3Call3Value) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.Aggregate3Value(&_MulticallPrimary.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactor) BlockAndAggregate(opts *bind.TransactOpts, calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.contract.Transact(opts, "blockAndAggregate", calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimarySession) BlockAndAggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.BlockAndAggregate(&_MulticallPrimary.TransactOpts, calls)
}

// BlockAndAggregate is a paid mutator transaction binding the contract method 0xc3077fa9.
//
// Solidity: function blockAndAggregate((address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactorSession) BlockAndAggregate(calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.BlockAndAggregate(&_MulticallPrimary.TransactOpts, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactor) TryAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.contract.Transact(opts, "tryAggregate", requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimarySession) TryAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.TryAggregate(&_MulticallPrimary.TransactOpts, requireSuccess, calls)
}

// TryAggregate is a paid mutator transaction binding the contract method 0xbce38bd7.
//
// Solidity: function tryAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns((bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactorSession) TryAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.TryAggregate(&_MulticallPrimary.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactor) TryBlockAndAggregate(opts *bind.TransactOpts, requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.contract.Transact(opts, "tryBlockAndAggregate", requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimarySession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.TryBlockAndAggregate(&_MulticallPrimary.TransactOpts, requireSuccess, calls)
}

// TryBlockAndAggregate is a paid mutator transaction binding the contract method 0x399542e9.
//
// Solidity: function tryBlockAndAggregate(bool requireSuccess, (address,bytes)[] calls) payable returns(uint256 blockNumber, bytes32 blockHash, (bool,bytes)[] returnData)
func (_MulticallPrimary *MulticallPrimaryTransactorSession) TryBlockAndAggregate(requireSuccess bool, calls []Multicall3Call) (*types.Transaction, error) {
	return _MulticallPrimary.Contract.TryBlockAndAggregate(&_MulticallPrimary.TransactOpts, requireSuccess, calls)
}
