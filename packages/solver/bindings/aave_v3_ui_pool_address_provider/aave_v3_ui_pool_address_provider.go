// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave_v3_ui_pool_address_provider

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

// AaveV3UiPoolAddressProviderMetaData contains all meta data concerning the AaveV3UiPoolAddressProvider contract.
var AaveV3UiPoolAddressProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ACLAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ACLManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementationAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementationAddress\",\"type\":\"address\"}],\"name\":\"AddressSetAsProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"oldMarketId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolDataProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleSentinelUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getACLAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACLManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolDataProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracleSentinel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAclAdmin\",\"type\":\"address\"}],\"name\":\"setACLAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAclManager\",\"type\":\"address\"}],\"name\":\"setACLManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newImplementationAddress\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolConfiguratorImpl\",\"type\":\"address\"}],\"name\":\"setPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDataProvider\",\"type\":\"address\"}],\"name\":\"setPoolDataProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolImpl\",\"type\":\"address\"}],\"name\":\"setPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracleSentinel\",\"type\":\"address\"}],\"name\":\"setPriceOracleSentinel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveV3UiPoolAddressProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveV3UiPoolAddressProviderMetaData.ABI instead.
var AaveV3UiPoolAddressProviderABI = AaveV3UiPoolAddressProviderMetaData.ABI

// AaveV3UiPoolAddressProvider is an auto generated Go binding around an Ethereum contract.
type AaveV3UiPoolAddressProvider struct {
	AaveV3UiPoolAddressProviderCaller     // Read-only binding to the contract
	AaveV3UiPoolAddressProviderTransactor // Write-only binding to the contract
	AaveV3UiPoolAddressProviderFilterer   // Log filterer for contract events
}

// AaveV3UiPoolAddressProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV3UiPoolAddressProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3UiPoolAddressProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV3UiPoolAddressProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3UiPoolAddressProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV3UiPoolAddressProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3UiPoolAddressProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV3UiPoolAddressProviderSession struct {
	Contract     *AaveV3UiPoolAddressProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AaveV3UiPoolAddressProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV3UiPoolAddressProviderCallerSession struct {
	Contract *AaveV3UiPoolAddressProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// AaveV3UiPoolAddressProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV3UiPoolAddressProviderTransactorSession struct {
	Contract     *AaveV3UiPoolAddressProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// AaveV3UiPoolAddressProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV3UiPoolAddressProviderRaw struct {
	Contract *AaveV3UiPoolAddressProvider // Generic contract binding to access the raw methods on
}

// AaveV3UiPoolAddressProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV3UiPoolAddressProviderCallerRaw struct {
	Contract *AaveV3UiPoolAddressProviderCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV3UiPoolAddressProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV3UiPoolAddressProviderTransactorRaw struct {
	Contract *AaveV3UiPoolAddressProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV3UiPoolAddressProvider creates a new instance of AaveV3UiPoolAddressProvider, bound to a specific deployed contract.
func NewAaveV3UiPoolAddressProvider(address common.Address, backend bind.ContractBackend) (*AaveV3UiPoolAddressProvider, error) {
	contract, err := bindAaveV3UiPoolAddressProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProvider{AaveV3UiPoolAddressProviderCaller: AaveV3UiPoolAddressProviderCaller{contract: contract}, AaveV3UiPoolAddressProviderTransactor: AaveV3UiPoolAddressProviderTransactor{contract: contract}, AaveV3UiPoolAddressProviderFilterer: AaveV3UiPoolAddressProviderFilterer{contract: contract}}, nil
}

// NewAaveV3UiPoolAddressProviderCaller creates a new read-only instance of AaveV3UiPoolAddressProvider, bound to a specific deployed contract.
func NewAaveV3UiPoolAddressProviderCaller(address common.Address, caller bind.ContractCaller) (*AaveV3UiPoolAddressProviderCaller, error) {
	contract, err := bindAaveV3UiPoolAddressProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderCaller{contract: contract}, nil
}

// NewAaveV3UiPoolAddressProviderTransactor creates a new write-only instance of AaveV3UiPoolAddressProvider, bound to a specific deployed contract.
func NewAaveV3UiPoolAddressProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV3UiPoolAddressProviderTransactor, error) {
	contract, err := bindAaveV3UiPoolAddressProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderTransactor{contract: contract}, nil
}

// NewAaveV3UiPoolAddressProviderFilterer creates a new log filterer instance of AaveV3UiPoolAddressProvider, bound to a specific deployed contract.
func NewAaveV3UiPoolAddressProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV3UiPoolAddressProviderFilterer, error) {
	contract, err := bindAaveV3UiPoolAddressProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderFilterer{contract: contract}, nil
}

// bindAaveV3UiPoolAddressProvider binds a generic wrapper to an already deployed contract.
func bindAaveV3UiPoolAddressProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveV3UiPoolAddressProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3UiPoolAddressProvider.Contract.AaveV3UiPoolAddressProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.AaveV3UiPoolAddressProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.AaveV3UiPoolAddressProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3UiPoolAddressProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.contract.Transact(opts, method, params...)
}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetACLAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getACLAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetACLAdmin() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetACLAdmin(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetACLAdmin() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetACLAdmin(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetACLManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getACLManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetACLManager() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetACLManager(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetACLManager() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetACLManager(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetAddress(id [32]byte) (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetAddress(&_AaveV3UiPoolAddressProvider.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetAddress(&_AaveV3UiPoolAddressProvider.CallOpts, id)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetMarketId() (string, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetMarketId(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetMarketId() (string, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetMarketId(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetPool() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPool(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetPool() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPool(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetPoolConfigurator() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPoolConfigurator(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetPoolConfigurator() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPoolConfigurator(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetPoolDataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getPoolDataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetPoolDataProvider() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPoolDataProvider(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetPoolDataProvider() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPoolDataProvider(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetPriceOracle() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPriceOracle(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPriceOracle(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) GetPriceOracleSentinel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "getPriceOracleSentinel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) GetPriceOracleSentinel() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPriceOracleSentinel(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) GetPriceOracleSentinel() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.GetPriceOracleSentinel(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolAddressProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) Owner() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.Owner(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderCallerSession) Owner() (common.Address, error) {
	return _AaveV3UiPoolAddressProvider.Contract.Owner(&_AaveV3UiPoolAddressProvider.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.RenounceOwnership(&_AaveV3UiPoolAddressProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.RenounceOwnership(&_AaveV3UiPoolAddressProvider.TransactOpts)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetACLAdmin(opts *bind.TransactOpts, newAclAdmin common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setACLAdmin", newAclAdmin)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetACLAdmin(newAclAdmin common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetACLAdmin(&_AaveV3UiPoolAddressProvider.TransactOpts, newAclAdmin)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetACLAdmin(newAclAdmin common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetACLAdmin(&_AaveV3UiPoolAddressProvider.TransactOpts, newAclAdmin)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetACLManager(opts *bind.TransactOpts, newAclManager common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setACLManager", newAclManager)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetACLManager(newAclManager common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetACLManager(&_AaveV3UiPoolAddressProvider.TransactOpts, newAclManager)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetACLManager(newAclManager common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetACLManager(&_AaveV3UiPoolAddressProvider.TransactOpts, newAclManager)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetAddress(&_AaveV3UiPoolAddressProvider.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetAddress(&_AaveV3UiPoolAddressProvider.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setAddressAsProxy", id, newImplementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetAddressAsProxy(id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetAddressAsProxy(&_AaveV3UiPoolAddressProvider.TransactOpts, id, newImplementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetAddressAsProxy(id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetAddressAsProxy(&_AaveV3UiPoolAddressProvider.TransactOpts, id, newImplementationAddress)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetMarketId(opts *bind.TransactOpts, newMarketId string) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setMarketId", newMarketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetMarketId(newMarketId string) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetMarketId(&_AaveV3UiPoolAddressProvider.TransactOpts, newMarketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetMarketId(newMarketId string) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetMarketId(&_AaveV3UiPoolAddressProvider.TransactOpts, newMarketId)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetPoolConfiguratorImpl(opts *bind.TransactOpts, newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setPoolConfiguratorImpl", newPoolConfiguratorImpl)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetPoolConfiguratorImpl(newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPoolConfiguratorImpl(&_AaveV3UiPoolAddressProvider.TransactOpts, newPoolConfiguratorImpl)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetPoolConfiguratorImpl(newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPoolConfiguratorImpl(&_AaveV3UiPoolAddressProvider.TransactOpts, newPoolConfiguratorImpl)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetPoolDataProvider(opts *bind.TransactOpts, newDataProvider common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setPoolDataProvider", newDataProvider)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetPoolDataProvider(newDataProvider common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPoolDataProvider(&_AaveV3UiPoolAddressProvider.TransactOpts, newDataProvider)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetPoolDataProvider(newDataProvider common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPoolDataProvider(&_AaveV3UiPoolAddressProvider.TransactOpts, newDataProvider)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetPoolImpl(opts *bind.TransactOpts, newPoolImpl common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setPoolImpl", newPoolImpl)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetPoolImpl(newPoolImpl common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPoolImpl(&_AaveV3UiPoolAddressProvider.TransactOpts, newPoolImpl)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetPoolImpl(newPoolImpl common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPoolImpl(&_AaveV3UiPoolAddressProvider.TransactOpts, newPoolImpl)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetPriceOracle(opts *bind.TransactOpts, newPriceOracle common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setPriceOracle", newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPriceOracle(&_AaveV3UiPoolAddressProvider.TransactOpts, newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPriceOracle(&_AaveV3UiPoolAddressProvider.TransactOpts, newPriceOracle)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) SetPriceOracleSentinel(opts *bind.TransactOpts, newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "setPriceOracleSentinel", newPriceOracleSentinel)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) SetPriceOracleSentinel(newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPriceOracleSentinel(&_AaveV3UiPoolAddressProvider.TransactOpts, newPriceOracleSentinel)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) SetPriceOracleSentinel(newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.SetPriceOracleSentinel(&_AaveV3UiPoolAddressProvider.TransactOpts, newPriceOracleSentinel)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.TransferOwnership(&_AaveV3UiPoolAddressProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveV3UiPoolAddressProvider.Contract.TransferOwnership(&_AaveV3UiPoolAddressProvider.TransactOpts, newOwner)
}

// AaveV3UiPoolAddressProviderACLAdminUpdatedIterator is returned from FilterACLAdminUpdated and is used to iterate over the raw logs and unpacked data for ACLAdminUpdated events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderACLAdminUpdatedIterator struct {
	Event *AaveV3UiPoolAddressProviderACLAdminUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderACLAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderACLAdminUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderACLAdminUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderACLAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderACLAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderACLAdminUpdated represents a ACLAdminUpdated event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderACLAdminUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterACLAdminUpdated is a free log retrieval operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterACLAdminUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3UiPoolAddressProviderACLAdminUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "ACLAdminUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderACLAdminUpdatedIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "ACLAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchACLAdminUpdated is a free log subscription operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchACLAdminUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderACLAdminUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "ACLAdminUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderACLAdminUpdated)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "ACLAdminUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseACLAdminUpdated is a log parse operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParseACLAdminUpdated(log types.Log) (*AaveV3UiPoolAddressProviderACLAdminUpdated, error) {
	event := new(AaveV3UiPoolAddressProviderACLAdminUpdated)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "ACLAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderACLManagerUpdatedIterator is returned from FilterACLManagerUpdated and is used to iterate over the raw logs and unpacked data for ACLManagerUpdated events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderACLManagerUpdatedIterator struct {
	Event *AaveV3UiPoolAddressProviderACLManagerUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderACLManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderACLManagerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderACLManagerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderACLManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderACLManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderACLManagerUpdated represents a ACLManagerUpdated event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderACLManagerUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterACLManagerUpdated is a free log retrieval operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterACLManagerUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3UiPoolAddressProviderACLManagerUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "ACLManagerUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderACLManagerUpdatedIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "ACLManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchACLManagerUpdated is a free log subscription operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchACLManagerUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderACLManagerUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "ACLManagerUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderACLManagerUpdated)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "ACLManagerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseACLManagerUpdated is a log parse operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParseACLManagerUpdated(log types.Log) (*AaveV3UiPoolAddressProviderACLManagerUpdated, error) {
	event := new(AaveV3UiPoolAddressProviderACLManagerUpdated)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "ACLManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderAddressSetIterator struct {
	Event *AaveV3UiPoolAddressProviderAddressSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderAddressSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderAddressSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderAddressSet represents a AddressSet event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderAddressSet struct {
	Id         [32]byte
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterAddressSet(opts *bind.FilterOpts, id [][32]byte, oldAddress []common.Address, newAddress []common.Address) (*AaveV3UiPoolAddressProviderAddressSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "AddressSet", idRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderAddressSetIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderAddressSet, id [][32]byte, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "AddressSet", idRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderAddressSet)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressSet is a log parse operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParseAddressSet(log types.Log) (*AaveV3UiPoolAddressProviderAddressSet, error) {
	event := new(AaveV3UiPoolAddressProviderAddressSet)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderAddressSetAsProxyIterator is returned from FilterAddressSetAsProxy and is used to iterate over the raw logs and unpacked data for AddressSetAsProxy events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderAddressSetAsProxyIterator struct {
	Event *AaveV3UiPoolAddressProviderAddressSetAsProxy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderAddressSetAsProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderAddressSetAsProxy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderAddressSetAsProxy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderAddressSetAsProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderAddressSetAsProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderAddressSetAsProxy represents a AddressSetAsProxy event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderAddressSetAsProxy struct {
	Id                       [32]byte
	ProxyAddress             common.Address
	OldImplementationAddress common.Address
	NewImplementationAddress common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAddressSetAsProxy is a free log retrieval operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterAddressSetAsProxy(opts *bind.FilterOpts, id [][32]byte, proxyAddress []common.Address, newImplementationAddress []common.Address) (*AaveV3UiPoolAddressProviderAddressSetAsProxyIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}

	var newImplementationAddressRule []interface{}
	for _, newImplementationAddressItem := range newImplementationAddress {
		newImplementationAddressRule = append(newImplementationAddressRule, newImplementationAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "AddressSetAsProxy", idRule, proxyAddressRule, newImplementationAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderAddressSetAsProxyIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "AddressSetAsProxy", logs: logs, sub: sub}, nil
}

// WatchAddressSetAsProxy is a free log subscription operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchAddressSetAsProxy(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderAddressSetAsProxy, id [][32]byte, proxyAddress []common.Address, newImplementationAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}

	var newImplementationAddressRule []interface{}
	for _, newImplementationAddressItem := range newImplementationAddress {
		newImplementationAddressRule = append(newImplementationAddressRule, newImplementationAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "AddressSetAsProxy", idRule, proxyAddressRule, newImplementationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderAddressSetAsProxy)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "AddressSetAsProxy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddressSetAsProxy is a log parse operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParseAddressSetAsProxy(log types.Log) (*AaveV3UiPoolAddressProviderAddressSetAsProxy, error) {
	event := new(AaveV3UiPoolAddressProviderAddressSetAsProxy)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "AddressSetAsProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderMarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderMarketIdSetIterator struct {
	Event *AaveV3UiPoolAddressProviderMarketIdSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderMarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderMarketIdSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderMarketIdSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderMarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderMarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderMarketIdSet represents a MarketIdSet event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderMarketIdSet struct {
	OldMarketId common.Hash
	NewMarketId common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterMarketIdSet(opts *bind.FilterOpts, oldMarketId []string, newMarketId []string) (*AaveV3UiPoolAddressProviderMarketIdSetIterator, error) {

	var oldMarketIdRule []interface{}
	for _, oldMarketIdItem := range oldMarketId {
		oldMarketIdRule = append(oldMarketIdRule, oldMarketIdItem)
	}
	var newMarketIdRule []interface{}
	for _, newMarketIdItem := range newMarketId {
		newMarketIdRule = append(newMarketIdRule, newMarketIdItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "MarketIdSet", oldMarketIdRule, newMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderMarketIdSetIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderMarketIdSet, oldMarketId []string, newMarketId []string) (event.Subscription, error) {

	var oldMarketIdRule []interface{}
	for _, oldMarketIdItem := range oldMarketId {
		oldMarketIdRule = append(oldMarketIdRule, oldMarketIdItem)
	}
	var newMarketIdRule []interface{}
	for _, newMarketIdItem := range newMarketId {
		newMarketIdRule = append(newMarketIdRule, newMarketIdItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "MarketIdSet", oldMarketIdRule, newMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderMarketIdSet)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketIdSet is a log parse operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParseMarketIdSet(log types.Log) (*AaveV3UiPoolAddressProviderMarketIdSet, error) {
	event := new(AaveV3UiPoolAddressProviderMarketIdSet)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderOwnershipTransferredIterator struct {
	Event *AaveV3UiPoolAddressProviderOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderOwnershipTransferred represents a OwnershipTransferred event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AaveV3UiPoolAddressProviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderOwnershipTransferredIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderOwnershipTransferred)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParseOwnershipTransferred(log types.Log) (*AaveV3UiPoolAddressProviderOwnershipTransferred, error) {
	event := new(AaveV3UiPoolAddressProviderOwnershipTransferred)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderPoolConfiguratorUpdatedIterator is returned from FilterPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for PoolConfiguratorUpdated events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPoolConfiguratorUpdatedIterator struct {
	Event *AaveV3UiPoolAddressProviderPoolConfiguratorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderPoolConfiguratorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderPoolConfiguratorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderPoolConfiguratorUpdated represents a PoolConfiguratorUpdated event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPoolConfiguratorUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterPoolConfiguratorUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3UiPoolAddressProviderPoolConfiguratorUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "PoolConfiguratorUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderPoolConfiguratorUpdatedIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "PoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderPoolConfiguratorUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "PoolConfiguratorUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderPoolConfiguratorUpdated)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PoolConfiguratorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolConfiguratorUpdated is a log parse operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParsePoolConfiguratorUpdated(log types.Log) (*AaveV3UiPoolAddressProviderPoolConfiguratorUpdated, error) {
	event := new(AaveV3UiPoolAddressProviderPoolConfiguratorUpdated)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderPoolDataProviderUpdatedIterator is returned from FilterPoolDataProviderUpdated and is used to iterate over the raw logs and unpacked data for PoolDataProviderUpdated events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPoolDataProviderUpdatedIterator struct {
	Event *AaveV3UiPoolAddressProviderPoolDataProviderUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderPoolDataProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderPoolDataProviderUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderPoolDataProviderUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderPoolDataProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderPoolDataProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderPoolDataProviderUpdated represents a PoolDataProviderUpdated event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPoolDataProviderUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolDataProviderUpdated is a free log retrieval operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterPoolDataProviderUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3UiPoolAddressProviderPoolDataProviderUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "PoolDataProviderUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderPoolDataProviderUpdatedIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "PoolDataProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolDataProviderUpdated is a free log subscription operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchPoolDataProviderUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderPoolDataProviderUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "PoolDataProviderUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderPoolDataProviderUpdated)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PoolDataProviderUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolDataProviderUpdated is a log parse operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParsePoolDataProviderUpdated(log types.Log) (*AaveV3UiPoolAddressProviderPoolDataProviderUpdated, error) {
	event := new(AaveV3UiPoolAddressProviderPoolDataProviderUpdated)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PoolDataProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderPoolUpdatedIterator is returned from FilterPoolUpdated and is used to iterate over the raw logs and unpacked data for PoolUpdated events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPoolUpdatedIterator struct {
	Event *AaveV3UiPoolAddressProviderPoolUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderPoolUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderPoolUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderPoolUpdated represents a PoolUpdated event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPoolUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolUpdated is a free log retrieval operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterPoolUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3UiPoolAddressProviderPoolUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "PoolUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderPoolUpdatedIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "PoolUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolUpdated is a free log subscription operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchPoolUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderPoolUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "PoolUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderPoolUpdated)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolUpdated is a log parse operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParsePoolUpdated(log types.Log) (*AaveV3UiPoolAddressProviderPoolUpdated, error) {
	event := new(AaveV3UiPoolAddressProviderPoolUpdated)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderPriceOracleSentinelUpdatedIterator is returned from FilterPriceOracleSentinelUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleSentinelUpdated events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPriceOracleSentinelUpdatedIterator struct {
	Event *AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderPriceOracleSentinelUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderPriceOracleSentinelUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderPriceOracleSentinelUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated represents a PriceOracleSentinelUpdated event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleSentinelUpdated is a free log retrieval operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterPriceOracleSentinelUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3UiPoolAddressProviderPriceOracleSentinelUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "PriceOracleSentinelUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderPriceOracleSentinelUpdatedIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "PriceOracleSentinelUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleSentinelUpdated is a free log subscription operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchPriceOracleSentinelUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "PriceOracleSentinelUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PriceOracleSentinelUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceOracleSentinelUpdated is a log parse operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParsePriceOracleSentinelUpdated(log types.Log) (*AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated, error) {
	event := new(AaveV3UiPoolAddressProviderPriceOracleSentinelUpdated)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PriceOracleSentinelUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPriceOracleUpdatedIterator struct {
	Event *AaveV3UiPoolAddressProviderPriceOracleUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderPriceOracleUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderPriceOracleUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderPriceOracleUpdated represents a PriceOracleUpdated event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderPriceOracleUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3UiPoolAddressProviderPriceOracleUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "PriceOracleUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderPriceOracleUpdatedIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderPriceOracleUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "PriceOracleUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderPriceOracleUpdated)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParsePriceOracleUpdated(log types.Log) (*AaveV3UiPoolAddressProviderPriceOracleUpdated, error) {
	event := new(AaveV3UiPoolAddressProviderPriceOracleUpdated)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3UiPoolAddressProviderProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderProxyCreatedIterator struct {
	Event *AaveV3UiPoolAddressProviderProxyCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AaveV3UiPoolAddressProviderProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3UiPoolAddressProviderProxyCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AaveV3UiPoolAddressProviderProxyCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AaveV3UiPoolAddressProviderProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3UiPoolAddressProviderProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3UiPoolAddressProviderProxyCreated represents a ProxyCreated event raised by the AaveV3UiPoolAddressProvider contract.
type AaveV3UiPoolAddressProviderProxyCreated struct {
	Id                    [32]byte
	ProxyAddress          common.Address
	ImplementationAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) FilterProxyCreated(opts *bind.FilterOpts, id [][32]byte, proxyAddress []common.Address, implementationAddress []common.Address) (*AaveV3UiPoolAddressProviderProxyCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}
	var implementationAddressRule []interface{}
	for _, implementationAddressItem := range implementationAddress {
		implementationAddressRule = append(implementationAddressRule, implementationAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.FilterLogs(opts, "ProxyCreated", idRule, proxyAddressRule, implementationAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolAddressProviderProxyCreatedIterator{contract: _AaveV3UiPoolAddressProvider.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *AaveV3UiPoolAddressProviderProxyCreated, id [][32]byte, proxyAddress []common.Address, implementationAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}
	var implementationAddressRule []interface{}
	for _, implementationAddressItem := range implementationAddress {
		implementationAddressRule = append(implementationAddressRule, implementationAddressItem)
	}

	logs, sub, err := _AaveV3UiPoolAddressProvider.contract.WatchLogs(opts, "ProxyCreated", idRule, proxyAddressRule, implementationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3UiPoolAddressProviderProxyCreated)
				if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProxyCreated is a log parse operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_AaveV3UiPoolAddressProvider *AaveV3UiPoolAddressProviderFilterer) ParseProxyCreated(log types.Log) (*AaveV3UiPoolAddressProviderProxyCreated, error) {
	event := new(AaveV3UiPoolAddressProviderProxyCreated)
	if err := _AaveV3UiPoolAddressProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
