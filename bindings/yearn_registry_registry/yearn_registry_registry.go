// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearn_registry_registry

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

// YearnRegistryRegistryMetaData contains all meta data concerning the YearnRegistryRegistry contract.
var YearnRegistryRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_releaseRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultType\",\"type\":\"uint256\"}],\"name\":\"NewEndorsedVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultType\",\"type\":\"uint256\"}],\"name\":\"RemovedVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateEndorser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateTagger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"VaultTagged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MULTI_STRATEGY_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SINGLE_STRATEGY_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetIsUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"endorseMultiStrategyVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"endorseSingleStrategyVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_vaultType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deploymentTimestamp\",\"type\":\"uint256\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"endorsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllEndorsedVaults\",\"outputs\":[{\"internalType\":\"address[][]\",\"name\":\"allEndorsedVaults\",\"type\":\"address[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getEndorsedVaults\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"isEndorsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_roleManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_profitMaxUnlockTime\",\"type\":\"uint256\"}],\"name\":\"newEndorsedVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_roleManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_profitMaxUnlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"}],\"name\":\"newEndorsedVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"numEndorsedVaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"removeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_canEndorse\",\"type\":\"bool\"}],\"name\":\"setEndorser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_canTag\",\"type\":\"bool\"}],\"name\":\"setTagger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"}],\"name\":\"tagVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"taggers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaultInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"releaseVersion\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"vaultType\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"deploymentTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YearnRegistryRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnRegistryRegistryMetaData.ABI instead.
var YearnRegistryRegistryABI = YearnRegistryRegistryMetaData.ABI

// YearnRegistryRegistry is an auto generated Go binding around an Ethereum contract.
type YearnRegistryRegistry struct {
	YearnRegistryRegistryCaller     // Read-only binding to the contract
	YearnRegistryRegistryTransactor // Write-only binding to the contract
	YearnRegistryRegistryFilterer   // Log filterer for contract events
}

// YearnRegistryRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnRegistryRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnRegistryRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnRegistryRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnRegistryRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnRegistryRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnRegistryRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnRegistryRegistrySession struct {
	Contract     *YearnRegistryRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// YearnRegistryRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnRegistryRegistryCallerSession struct {
	Contract *YearnRegistryRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// YearnRegistryRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnRegistryRegistryTransactorSession struct {
	Contract     *YearnRegistryRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// YearnRegistryRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnRegistryRegistryRaw struct {
	Contract *YearnRegistryRegistry // Generic contract binding to access the raw methods on
}

// YearnRegistryRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnRegistryRegistryCallerRaw struct {
	Contract *YearnRegistryRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// YearnRegistryRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnRegistryRegistryTransactorRaw struct {
	Contract *YearnRegistryRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnRegistryRegistry creates a new instance of YearnRegistryRegistry, bound to a specific deployed contract.
func NewYearnRegistryRegistry(address common.Address, backend bind.ContractBackend) (*YearnRegistryRegistry, error) {
	contract, err := bindYearnRegistryRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistry{YearnRegistryRegistryCaller: YearnRegistryRegistryCaller{contract: contract}, YearnRegistryRegistryTransactor: YearnRegistryRegistryTransactor{contract: contract}, YearnRegistryRegistryFilterer: YearnRegistryRegistryFilterer{contract: contract}}, nil
}

// NewYearnRegistryRegistryCaller creates a new read-only instance of YearnRegistryRegistry, bound to a specific deployed contract.
func NewYearnRegistryRegistryCaller(address common.Address, caller bind.ContractCaller) (*YearnRegistryRegistryCaller, error) {
	contract, err := bindYearnRegistryRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryCaller{contract: contract}, nil
}

// NewYearnRegistryRegistryTransactor creates a new write-only instance of YearnRegistryRegistry, bound to a specific deployed contract.
func NewYearnRegistryRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnRegistryRegistryTransactor, error) {
	contract, err := bindYearnRegistryRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryTransactor{contract: contract}, nil
}

// NewYearnRegistryRegistryFilterer creates a new log filterer instance of YearnRegistryRegistry, bound to a specific deployed contract.
func NewYearnRegistryRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnRegistryRegistryFilterer, error) {
	contract, err := bindYearnRegistryRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryFilterer{contract: contract}, nil
}

// bindYearnRegistryRegistry binds a generic wrapper to an already deployed contract.
func bindYearnRegistryRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YearnRegistryRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnRegistryRegistry *YearnRegistryRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnRegistryRegistry.Contract.YearnRegistryRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnRegistryRegistry *YearnRegistryRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.YearnRegistryRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnRegistryRegistry *YearnRegistryRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.YearnRegistryRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnRegistryRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.contract.Transact(opts, method, params...)
}

// MULTISTRATEGYTYPE is a free data retrieval call binding the contract method 0x18d1dd83.
//
// Solidity: function MULTI_STRATEGY_TYPE() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) MULTISTRATEGYTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "MULTI_STRATEGY_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTISTRATEGYTYPE is a free data retrieval call binding the contract method 0x18d1dd83.
//
// Solidity: function MULTI_STRATEGY_TYPE() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) MULTISTRATEGYTYPE() (*big.Int, error) {
	return _YearnRegistryRegistry.Contract.MULTISTRATEGYTYPE(&_YearnRegistryRegistry.CallOpts)
}

// MULTISTRATEGYTYPE is a free data retrieval call binding the contract method 0x18d1dd83.
//
// Solidity: function MULTI_STRATEGY_TYPE() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) MULTISTRATEGYTYPE() (*big.Int, error) {
	return _YearnRegistryRegistry.Contract.MULTISTRATEGYTYPE(&_YearnRegistryRegistry.CallOpts)
}

// SINGLESTRATEGYTYPE is a free data retrieval call binding the contract method 0x98a5e07b.
//
// Solidity: function SINGLE_STRATEGY_TYPE() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) SINGLESTRATEGYTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "SINGLE_STRATEGY_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SINGLESTRATEGYTYPE is a free data retrieval call binding the contract method 0x98a5e07b.
//
// Solidity: function SINGLE_STRATEGY_TYPE() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) SINGLESTRATEGYTYPE() (*big.Int, error) {
	return _YearnRegistryRegistry.Contract.SINGLESTRATEGYTYPE(&_YearnRegistryRegistry.CallOpts)
}

// SINGLESTRATEGYTYPE is a free data retrieval call binding the contract method 0x98a5e07b.
//
// Solidity: function SINGLE_STRATEGY_TYPE() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) SINGLESTRATEGYTYPE() (*big.Int, error) {
	return _YearnRegistryRegistry.Contract.SINGLESTRATEGYTYPE(&_YearnRegistryRegistry.CallOpts)
}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) AssetIsUsed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "assetIsUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) AssetIsUsed(arg0 common.Address) (bool, error) {
	return _YearnRegistryRegistry.Contract.AssetIsUsed(&_YearnRegistryRegistry.CallOpts, arg0)
}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) AssetIsUsed(arg0 common.Address) (bool, error) {
	return _YearnRegistryRegistry.Contract.AssetIsUsed(&_YearnRegistryRegistry.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "assets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) Assets(arg0 *big.Int) (common.Address, error) {
	return _YearnRegistryRegistry.Contract.Assets(&_YearnRegistryRegistry.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _YearnRegistryRegistry.Contract.Assets(&_YearnRegistryRegistry.CallOpts, arg0)
}

// Endorsers is a free data retrieval call binding the contract method 0x3515a20b.
//
// Solidity: function endorsers(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) Endorsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "endorsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Endorsers is a free data retrieval call binding the contract method 0x3515a20b.
//
// Solidity: function endorsers(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) Endorsers(arg0 common.Address) (bool, error) {
	return _YearnRegistryRegistry.Contract.Endorsers(&_YearnRegistryRegistry.CallOpts, arg0)
}

// Endorsers is a free data retrieval call binding the contract method 0x3515a20b.
//
// Solidity: function endorsers(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) Endorsers(arg0 common.Address) (bool, error) {
	return _YearnRegistryRegistry.Contract.Endorsers(&_YearnRegistryRegistry.CallOpts, arg0)
}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) GetAllEndorsedVaults(opts *bind.CallOpts) ([][]common.Address, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "getAllEndorsedVaults")

	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)

	return out0, err

}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) GetAllEndorsedVaults() ([][]common.Address, error) {
	return _YearnRegistryRegistry.Contract.GetAllEndorsedVaults(&_YearnRegistryRegistry.CallOpts)
}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) GetAllEndorsedVaults() ([][]common.Address, error) {
	return _YearnRegistryRegistry.Contract.GetAllEndorsedVaults(&_YearnRegistryRegistry.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) GetAssets() ([]common.Address, error) {
	return _YearnRegistryRegistry.Contract.GetAssets(&_YearnRegistryRegistry.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) GetAssets() ([]common.Address, error) {
	return _YearnRegistryRegistry.Contract.GetAssets(&_YearnRegistryRegistry.CallOpts)
}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) GetEndorsedVaults(opts *bind.CallOpts, _asset common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "getEndorsedVaults", _asset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) GetEndorsedVaults(_asset common.Address) ([]common.Address, error) {
	return _YearnRegistryRegistry.Contract.GetEndorsedVaults(&_YearnRegistryRegistry.CallOpts, _asset)
}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) GetEndorsedVaults(_asset common.Address) ([]common.Address, error) {
	return _YearnRegistryRegistry.Contract.GetEndorsedVaults(&_YearnRegistryRegistry.CallOpts, _asset)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) Governance() (common.Address, error) {
	return _YearnRegistryRegistry.Contract.Governance(&_YearnRegistryRegistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) Governance() (common.Address, error) {
	return _YearnRegistryRegistry.Contract.Governance(&_YearnRegistryRegistry.CallOpts)
}

// IsEndorsed is a free data retrieval call binding the contract method 0xa237e94d.
//
// Solidity: function isEndorsed(address _vault) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) IsEndorsed(opts *bind.CallOpts, _vault common.Address) (bool, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "isEndorsed", _vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEndorsed is a free data retrieval call binding the contract method 0xa237e94d.
//
// Solidity: function isEndorsed(address _vault) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) IsEndorsed(_vault common.Address) (bool, error) {
	return _YearnRegistryRegistry.Contract.IsEndorsed(&_YearnRegistryRegistry.CallOpts, _vault)
}

// IsEndorsed is a free data retrieval call binding the contract method 0xa237e94d.
//
// Solidity: function isEndorsed(address _vault) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) IsEndorsed(_vault common.Address) (bool, error) {
	return _YearnRegistryRegistry.Contract.IsEndorsed(&_YearnRegistryRegistry.CallOpts, _vault)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) Name() (string, error) {
	return _YearnRegistryRegistry.Contract.Name(&_YearnRegistryRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) Name() (string, error) {
	return _YearnRegistryRegistry.Contract.Name(&_YearnRegistryRegistry.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) NumAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "numAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) NumAssets() (*big.Int, error) {
	return _YearnRegistryRegistry.Contract.NumAssets(&_YearnRegistryRegistry.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) NumAssets() (*big.Int, error) {
	return _YearnRegistryRegistry.Contract.NumAssets(&_YearnRegistryRegistry.CallOpts)
}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) NumEndorsedVaults(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "numEndorsedVaults", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) NumEndorsedVaults(_asset common.Address) (*big.Int, error) {
	return _YearnRegistryRegistry.Contract.NumEndorsedVaults(&_YearnRegistryRegistry.CallOpts, _asset)
}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) NumEndorsedVaults(_asset common.Address) (*big.Int, error) {
	return _YearnRegistryRegistry.Contract.NumEndorsedVaults(&_YearnRegistryRegistry.CallOpts, _asset)
}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) ReleaseRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "releaseRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) ReleaseRegistry() (common.Address, error) {
	return _YearnRegistryRegistry.Contract.ReleaseRegistry(&_YearnRegistryRegistry.CallOpts)
}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) ReleaseRegistry() (common.Address, error) {
	return _YearnRegistryRegistry.Contract.ReleaseRegistry(&_YearnRegistryRegistry.CallOpts)
}

// Taggers is a free data retrieval call binding the contract method 0x5b25d2c8.
//
// Solidity: function taggers(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) Taggers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "taggers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Taggers is a free data retrieval call binding the contract method 0x5b25d2c8.
//
// Solidity: function taggers(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) Taggers(arg0 common.Address) (bool, error) {
	return _YearnRegistryRegistry.Contract.Taggers(&_YearnRegistryRegistry.CallOpts, arg0)
}

// Taggers is a free data retrieval call binding the contract method 0x5b25d2c8.
//
// Solidity: function taggers(address ) view returns(bool)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) Taggers(arg0 common.Address) (bool, error) {
	return _YearnRegistryRegistry.Contract.Taggers(&_YearnRegistryRegistry.CallOpts, arg0)
}

// VaultInfo is a free data retrieval call binding the contract method 0x9164359a.
//
// Solidity: function vaultInfo(address ) view returns(address asset, uint96 releaseVersion, uint64 vaultType, uint128 deploymentTimestamp, uint64 index, string tag)
func (_YearnRegistryRegistry *YearnRegistryRegistryCaller) VaultInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	VaultType           uint64
	DeploymentTimestamp *big.Int
	Index               uint64
	Tag                 string
}, error) {
	var out []interface{}
	err := _YearnRegistryRegistry.contract.Call(opts, &out, "vaultInfo", arg0)

	outstruct := new(struct {
		Asset               common.Address
		ReleaseVersion      *big.Int
		VaultType           uint64
		DeploymentTimestamp *big.Int
		Index               uint64
		Tag                 string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Asset = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ReleaseVersion = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.VaultType = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.DeploymentTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Index = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Tag = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// VaultInfo is a free data retrieval call binding the contract method 0x9164359a.
//
// Solidity: function vaultInfo(address ) view returns(address asset, uint96 releaseVersion, uint64 vaultType, uint128 deploymentTimestamp, uint64 index, string tag)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) VaultInfo(arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	VaultType           uint64
	DeploymentTimestamp *big.Int
	Index               uint64
	Tag                 string
}, error) {
	return _YearnRegistryRegistry.Contract.VaultInfo(&_YearnRegistryRegistry.CallOpts, arg0)
}

// VaultInfo is a free data retrieval call binding the contract method 0x9164359a.
//
// Solidity: function vaultInfo(address ) view returns(address asset, uint96 releaseVersion, uint64 vaultType, uint128 deploymentTimestamp, uint64 index, string tag)
func (_YearnRegistryRegistry *YearnRegistryRegistryCallerSession) VaultInfo(arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	VaultType           uint64
	DeploymentTimestamp *big.Int
	Index               uint64
	Tag                 string
}, error) {
	return _YearnRegistryRegistry.Contract.VaultInfo(&_YearnRegistryRegistry.CallOpts, arg0)
}

// EndorseMultiStrategyVault is a paid mutator transaction binding the contract method 0x0ab322d9.
//
// Solidity: function endorseMultiStrategyVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) EndorseMultiStrategyVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "endorseMultiStrategyVault", _vault)
}

// EndorseMultiStrategyVault is a paid mutator transaction binding the contract method 0x0ab322d9.
//
// Solidity: function endorseMultiStrategyVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) EndorseMultiStrategyVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.EndorseMultiStrategyVault(&_YearnRegistryRegistry.TransactOpts, _vault)
}

// EndorseMultiStrategyVault is a paid mutator transaction binding the contract method 0x0ab322d9.
//
// Solidity: function endorseMultiStrategyVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) EndorseMultiStrategyVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.EndorseMultiStrategyVault(&_YearnRegistryRegistry.TransactOpts, _vault)
}

// EndorseSingleStrategyVault is a paid mutator transaction binding the contract method 0x0f7872cc.
//
// Solidity: function endorseSingleStrategyVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) EndorseSingleStrategyVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "endorseSingleStrategyVault", _vault)
}

// EndorseSingleStrategyVault is a paid mutator transaction binding the contract method 0x0f7872cc.
//
// Solidity: function endorseSingleStrategyVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) EndorseSingleStrategyVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.EndorseSingleStrategyVault(&_YearnRegistryRegistry.TransactOpts, _vault)
}

// EndorseSingleStrategyVault is a paid mutator transaction binding the contract method 0x0f7872cc.
//
// Solidity: function endorseSingleStrategyVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) EndorseSingleStrategyVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.EndorseSingleStrategyVault(&_YearnRegistryRegistry.TransactOpts, _vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x89c6acec.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _vaultType, uint256 _deploymentTimestamp) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) EndorseVault(opts *bind.TransactOpts, _vault common.Address, _releaseDelta *big.Int, _vaultType *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "endorseVault", _vault, _releaseDelta, _vaultType, _deploymentTimestamp)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x89c6acec.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _vaultType, uint256 _deploymentTimestamp) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) EndorseVault(_vault common.Address, _releaseDelta *big.Int, _vaultType *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.EndorseVault(&_YearnRegistryRegistry.TransactOpts, _vault, _releaseDelta, _vaultType, _deploymentTimestamp)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x89c6acec.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _vaultType, uint256 _deploymentTimestamp) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) EndorseVault(_vault common.Address, _releaseDelta *big.Int, _vaultType *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.EndorseVault(&_YearnRegistryRegistry.TransactOpts, _vault, _releaseDelta, _vaultType, _deploymentTimestamp)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x17bdd312.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime) returns(address _vault)
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) NewEndorsedVault(opts *bind.TransactOpts, _asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "newEndorsedVault", _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x17bdd312.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime) returns(address _vault)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) NewEndorsedVault(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.NewEndorsedVault(&_YearnRegistryRegistry.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x17bdd312.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime) returns(address _vault)
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) NewEndorsedVault(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.NewEndorsedVault(&_YearnRegistryRegistry.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime)
}

// NewEndorsedVault0 is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) NewEndorsedVault0(opts *bind.TransactOpts, _asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "newEndorsedVault0", _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// NewEndorsedVault0 is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) NewEndorsedVault0(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.NewEndorsedVault0(&_YearnRegistryRegistry.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// NewEndorsedVault0 is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) NewEndorsedVault0(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.NewEndorsedVault0(&_YearnRegistryRegistry.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) RemoveAsset(opts *bind.TransactOpts, _asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "removeAsset", _asset, _index)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) RemoveAsset(_asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.RemoveAsset(&_YearnRegistryRegistry.TransactOpts, _asset, _index)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) RemoveAsset(_asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.RemoveAsset(&_YearnRegistryRegistry.TransactOpts, _asset, _index)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) RemoveVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "removeVault", _vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) RemoveVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.RemoveVault(&_YearnRegistryRegistry.TransactOpts, _vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) RemoveVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.RemoveVault(&_YearnRegistryRegistry.TransactOpts, _vault)
}

// SetEndorser is a paid mutator transaction binding the contract method 0x2c2a72d5.
//
// Solidity: function setEndorser(address _account, bool _canEndorse) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) SetEndorser(opts *bind.TransactOpts, _account common.Address, _canEndorse bool) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "setEndorser", _account, _canEndorse)
}

// SetEndorser is a paid mutator transaction binding the contract method 0x2c2a72d5.
//
// Solidity: function setEndorser(address _account, bool _canEndorse) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) SetEndorser(_account common.Address, _canEndorse bool) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.SetEndorser(&_YearnRegistryRegistry.TransactOpts, _account, _canEndorse)
}

// SetEndorser is a paid mutator transaction binding the contract method 0x2c2a72d5.
//
// Solidity: function setEndorser(address _account, bool _canEndorse) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) SetEndorser(_account common.Address, _canEndorse bool) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.SetEndorser(&_YearnRegistryRegistry.TransactOpts, _account, _canEndorse)
}

// SetTagger is a paid mutator transaction binding the contract method 0x2aa59c92.
//
// Solidity: function setTagger(address _account, bool _canTag) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) SetTagger(opts *bind.TransactOpts, _account common.Address, _canTag bool) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "setTagger", _account, _canTag)
}

// SetTagger is a paid mutator transaction binding the contract method 0x2aa59c92.
//
// Solidity: function setTagger(address _account, bool _canTag) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) SetTagger(_account common.Address, _canTag bool) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.SetTagger(&_YearnRegistryRegistry.TransactOpts, _account, _canTag)
}

// SetTagger is a paid mutator transaction binding the contract method 0x2aa59c92.
//
// Solidity: function setTagger(address _account, bool _canTag) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) SetTagger(_account common.Address, _canTag bool) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.SetTagger(&_YearnRegistryRegistry.TransactOpts, _account, _canTag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) TagVault(opts *bind.TransactOpts, _vault common.Address, _tag string) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "tagVault", _vault, _tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) TagVault(_vault common.Address, _tag string) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.TagVault(&_YearnRegistryRegistry.TransactOpts, _vault, _tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) TagVault(_vault common.Address, _tag string) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.TagVault(&_YearnRegistryRegistry.TransactOpts, _vault, _tag)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactor) TransferGovernance(opts *bind.TransactOpts, _newGovernance common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.contract.Transact(opts, "transferGovernance", _newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistrySession) TransferGovernance(_newGovernance common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.TransferGovernance(&_YearnRegistryRegistry.TransactOpts, _newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YearnRegistryRegistry *YearnRegistryRegistryTransactorSession) TransferGovernance(_newGovernance common.Address) (*types.Transaction, error) {
	return _YearnRegistryRegistry.Contract.TransferGovernance(&_YearnRegistryRegistry.TransactOpts, _newGovernance)
}

// YearnRegistryRegistryGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryGovernanceTransferredIterator struct {
	Event *YearnRegistryRegistryGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *YearnRegistryRegistryGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnRegistryRegistryGovernanceTransferred)
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
		it.Event = new(YearnRegistryRegistryGovernanceTransferred)
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
func (it *YearnRegistryRegistryGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnRegistryRegistryGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnRegistryRegistryGovernanceTransferred represents a GovernanceTransferred event raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*YearnRegistryRegistryGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryGovernanceTransferredIterator{contract: _YearnRegistryRegistry.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *YearnRegistryRegistryGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnRegistryRegistryGovernanceTransferred)
				if err := _YearnRegistryRegistry.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) ParseGovernanceTransferred(log types.Log) (*YearnRegistryRegistryGovernanceTransferred, error) {
	event := new(YearnRegistryRegistryGovernanceTransferred)
	if err := _YearnRegistryRegistry.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnRegistryRegistryNewEndorsedVaultIterator is returned from FilterNewEndorsedVault and is used to iterate over the raw logs and unpacked data for NewEndorsedVault events raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryNewEndorsedVaultIterator struct {
	Event *YearnRegistryRegistryNewEndorsedVault // Event containing the contract specifics and raw log

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
func (it *YearnRegistryRegistryNewEndorsedVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnRegistryRegistryNewEndorsedVault)
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
		it.Event = new(YearnRegistryRegistryNewEndorsedVault)
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
func (it *YearnRegistryRegistryNewEndorsedVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnRegistryRegistryNewEndorsedVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnRegistryRegistryNewEndorsedVault represents a NewEndorsedVault event raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryNewEndorsedVault struct {
	Vault          common.Address
	Asset          common.Address
	ReleaseVersion *big.Int
	VaultType      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewEndorsedVault is a free log retrieval operation binding the contract event 0xa9a7c68f108b706e545bc75ac8590730afa49f639d2e48f367105c9801c18fd2.
//
// Solidity: event NewEndorsedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) FilterNewEndorsedVault(opts *bind.FilterOpts, vault []common.Address, asset []common.Address) (*YearnRegistryRegistryNewEndorsedVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.FilterLogs(opts, "NewEndorsedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryNewEndorsedVaultIterator{contract: _YearnRegistryRegistry.contract, event: "NewEndorsedVault", logs: logs, sub: sub}, nil
}

// WatchNewEndorsedVault is a free log subscription operation binding the contract event 0xa9a7c68f108b706e545bc75ac8590730afa49f639d2e48f367105c9801c18fd2.
//
// Solidity: event NewEndorsedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) WatchNewEndorsedVault(opts *bind.WatchOpts, sink chan<- *YearnRegistryRegistryNewEndorsedVault, vault []common.Address, asset []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.WatchLogs(opts, "NewEndorsedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnRegistryRegistryNewEndorsedVault)
				if err := _YearnRegistryRegistry.contract.UnpackLog(event, "NewEndorsedVault", log); err != nil {
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

// ParseNewEndorsedVault is a log parse operation binding the contract event 0xa9a7c68f108b706e545bc75ac8590730afa49f639d2e48f367105c9801c18fd2.
//
// Solidity: event NewEndorsedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) ParseNewEndorsedVault(log types.Log) (*YearnRegistryRegistryNewEndorsedVault, error) {
	event := new(YearnRegistryRegistryNewEndorsedVault)
	if err := _YearnRegistryRegistry.contract.UnpackLog(event, "NewEndorsedVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnRegistryRegistryRemovedVaultIterator is returned from FilterRemovedVault and is used to iterate over the raw logs and unpacked data for RemovedVault events raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryRemovedVaultIterator struct {
	Event *YearnRegistryRegistryRemovedVault // Event containing the contract specifics and raw log

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
func (it *YearnRegistryRegistryRemovedVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnRegistryRegistryRemovedVault)
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
		it.Event = new(YearnRegistryRegistryRemovedVault)
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
func (it *YearnRegistryRegistryRemovedVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnRegistryRegistryRemovedVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnRegistryRegistryRemovedVault represents a RemovedVault event raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryRemovedVault struct {
	Vault          common.Address
	Asset          common.Address
	ReleaseVersion *big.Int
	VaultType      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemovedVault is a free log retrieval operation binding the contract event 0xb8d23ba050f8f00e22675f82cf3786ade63b12a46d4ea236927baf4d173c3092.
//
// Solidity: event RemovedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) FilterRemovedVault(opts *bind.FilterOpts, vault []common.Address, asset []common.Address) (*YearnRegistryRegistryRemovedVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.FilterLogs(opts, "RemovedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryRemovedVaultIterator{contract: _YearnRegistryRegistry.contract, event: "RemovedVault", logs: logs, sub: sub}, nil
}

// WatchRemovedVault is a free log subscription operation binding the contract event 0xb8d23ba050f8f00e22675f82cf3786ade63b12a46d4ea236927baf4d173c3092.
//
// Solidity: event RemovedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) WatchRemovedVault(opts *bind.WatchOpts, sink chan<- *YearnRegistryRegistryRemovedVault, vault []common.Address, asset []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.WatchLogs(opts, "RemovedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnRegistryRegistryRemovedVault)
				if err := _YearnRegistryRegistry.contract.UnpackLog(event, "RemovedVault", log); err != nil {
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

// ParseRemovedVault is a log parse operation binding the contract event 0xb8d23ba050f8f00e22675f82cf3786ade63b12a46d4ea236927baf4d173c3092.
//
// Solidity: event RemovedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) ParseRemovedVault(log types.Log) (*YearnRegistryRegistryRemovedVault, error) {
	event := new(YearnRegistryRegistryRemovedVault)
	if err := _YearnRegistryRegistry.contract.UnpackLog(event, "RemovedVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnRegistryRegistryUpdateEndorserIterator is returned from FilterUpdateEndorser and is used to iterate over the raw logs and unpacked data for UpdateEndorser events raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryUpdateEndorserIterator struct {
	Event *YearnRegistryRegistryUpdateEndorser // Event containing the contract specifics and raw log

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
func (it *YearnRegistryRegistryUpdateEndorserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnRegistryRegistryUpdateEndorser)
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
		it.Event = new(YearnRegistryRegistryUpdateEndorser)
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
func (it *YearnRegistryRegistryUpdateEndorserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnRegistryRegistryUpdateEndorserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnRegistryRegistryUpdateEndorser represents a UpdateEndorser event raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryUpdateEndorser struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateEndorser is a free log retrieval operation binding the contract event 0xc93ec0e3c82bbe3866d85f7d0915cda166df7c76944b9fae88bcf11608f791bf.
//
// Solidity: event UpdateEndorser(address indexed account, bool status)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) FilterUpdateEndorser(opts *bind.FilterOpts, account []common.Address) (*YearnRegistryRegistryUpdateEndorserIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.FilterLogs(opts, "UpdateEndorser", accountRule)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryUpdateEndorserIterator{contract: _YearnRegistryRegistry.contract, event: "UpdateEndorser", logs: logs, sub: sub}, nil
}

// WatchUpdateEndorser is a free log subscription operation binding the contract event 0xc93ec0e3c82bbe3866d85f7d0915cda166df7c76944b9fae88bcf11608f791bf.
//
// Solidity: event UpdateEndorser(address indexed account, bool status)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) WatchUpdateEndorser(opts *bind.WatchOpts, sink chan<- *YearnRegistryRegistryUpdateEndorser, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.WatchLogs(opts, "UpdateEndorser", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnRegistryRegistryUpdateEndorser)
				if err := _YearnRegistryRegistry.contract.UnpackLog(event, "UpdateEndorser", log); err != nil {
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

// ParseUpdateEndorser is a log parse operation binding the contract event 0xc93ec0e3c82bbe3866d85f7d0915cda166df7c76944b9fae88bcf11608f791bf.
//
// Solidity: event UpdateEndorser(address indexed account, bool status)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) ParseUpdateEndorser(log types.Log) (*YearnRegistryRegistryUpdateEndorser, error) {
	event := new(YearnRegistryRegistryUpdateEndorser)
	if err := _YearnRegistryRegistry.contract.UnpackLog(event, "UpdateEndorser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnRegistryRegistryUpdateTaggerIterator is returned from FilterUpdateTagger and is used to iterate over the raw logs and unpacked data for UpdateTagger events raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryUpdateTaggerIterator struct {
	Event *YearnRegistryRegistryUpdateTagger // Event containing the contract specifics and raw log

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
func (it *YearnRegistryRegistryUpdateTaggerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnRegistryRegistryUpdateTagger)
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
		it.Event = new(YearnRegistryRegistryUpdateTagger)
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
func (it *YearnRegistryRegistryUpdateTaggerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnRegistryRegistryUpdateTaggerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnRegistryRegistryUpdateTagger represents a UpdateTagger event raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryUpdateTagger struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTagger is a free log retrieval operation binding the contract event 0xcd202ea0907016dd42e40faedc5cf5c3e6368644993f46c95990dfa7f84bfaa9.
//
// Solidity: event UpdateTagger(address indexed account, bool status)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) FilterUpdateTagger(opts *bind.FilterOpts, account []common.Address) (*YearnRegistryRegistryUpdateTaggerIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.FilterLogs(opts, "UpdateTagger", accountRule)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryUpdateTaggerIterator{contract: _YearnRegistryRegistry.contract, event: "UpdateTagger", logs: logs, sub: sub}, nil
}

// WatchUpdateTagger is a free log subscription operation binding the contract event 0xcd202ea0907016dd42e40faedc5cf5c3e6368644993f46c95990dfa7f84bfaa9.
//
// Solidity: event UpdateTagger(address indexed account, bool status)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) WatchUpdateTagger(opts *bind.WatchOpts, sink chan<- *YearnRegistryRegistryUpdateTagger, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.WatchLogs(opts, "UpdateTagger", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnRegistryRegistryUpdateTagger)
				if err := _YearnRegistryRegistry.contract.UnpackLog(event, "UpdateTagger", log); err != nil {
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

// ParseUpdateTagger is a log parse operation binding the contract event 0xcd202ea0907016dd42e40faedc5cf5c3e6368644993f46c95990dfa7f84bfaa9.
//
// Solidity: event UpdateTagger(address indexed account, bool status)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) ParseUpdateTagger(log types.Log) (*YearnRegistryRegistryUpdateTagger, error) {
	event := new(YearnRegistryRegistryUpdateTagger)
	if err := _YearnRegistryRegistry.contract.UnpackLog(event, "UpdateTagger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnRegistryRegistryVaultTaggedIterator is returned from FilterVaultTagged and is used to iterate over the raw logs and unpacked data for VaultTagged events raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryVaultTaggedIterator struct {
	Event *YearnRegistryRegistryVaultTagged // Event containing the contract specifics and raw log

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
func (it *YearnRegistryRegistryVaultTaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnRegistryRegistryVaultTagged)
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
		it.Event = new(YearnRegistryRegistryVaultTagged)
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
func (it *YearnRegistryRegistryVaultTaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnRegistryRegistryVaultTaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnRegistryRegistryVaultTagged represents a VaultTagged event raised by the YearnRegistryRegistry contract.
type YearnRegistryRegistryVaultTagged struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultTagged is a free log retrieval operation binding the contract event 0x9831985c7ff4c3d7dbd921d753150cc58a3a2a21c93795bbaac5bbf32baab3bb.
//
// Solidity: event VaultTagged(address indexed vault)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) FilterVaultTagged(opts *bind.FilterOpts, vault []common.Address) (*YearnRegistryRegistryVaultTaggedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.FilterLogs(opts, "VaultTagged", vaultRule)
	if err != nil {
		return nil, err
	}
	return &YearnRegistryRegistryVaultTaggedIterator{contract: _YearnRegistryRegistry.contract, event: "VaultTagged", logs: logs, sub: sub}, nil
}

// WatchVaultTagged is a free log subscription operation binding the contract event 0x9831985c7ff4c3d7dbd921d753150cc58a3a2a21c93795bbaac5bbf32baab3bb.
//
// Solidity: event VaultTagged(address indexed vault)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) WatchVaultTagged(opts *bind.WatchOpts, sink chan<- *YearnRegistryRegistryVaultTagged, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YearnRegistryRegistry.contract.WatchLogs(opts, "VaultTagged", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnRegistryRegistryVaultTagged)
				if err := _YearnRegistryRegistry.contract.UnpackLog(event, "VaultTagged", log); err != nil {
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

// ParseVaultTagged is a log parse operation binding the contract event 0x9831985c7ff4c3d7dbd921d753150cc58a3a2a21c93795bbaac5bbf32baab3bb.
//
// Solidity: event VaultTagged(address indexed vault)
func (_YearnRegistryRegistry *YearnRegistryRegistryFilterer) ParseVaultTagged(log types.Log) (*YearnRegistryRegistryVaultTagged, error) {
	event := new(YearnRegistryRegistryVaultTagged)
	if err := _YearnRegistryRegistry.contract.UnpackLog(event, "VaultTagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
