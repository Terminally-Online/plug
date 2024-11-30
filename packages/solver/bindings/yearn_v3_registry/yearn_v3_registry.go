// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yearn_v3_registry

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

// YearnV3RegistryMetaData contains all meta data concerning the YearnV3Registry contract.
var YearnV3RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_releaseRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultType\",\"type\":\"uint256\"}],\"name\":\"NewEndorsedVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseVersion\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vaultType\",\"type\":\"uint256\"}],\"name\":\"RemovedVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateEndorser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"UpdateTagger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"VaultTagged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MULTI_STRATEGY_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SINGLE_STRATEGY_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"assetIsUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"endorseMultiStrategyVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"endorseSingleStrategyVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_vaultType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deploymentTimestamp\",\"type\":\"uint256\"}],\"name\":\"endorseVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"endorsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllEndorsedVaults\",\"outputs\":[{\"internalType\":\"address[][]\",\"name\":\"allEndorsedVaults\",\"type\":\"address[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"getEndorsedVaults\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"isEndorsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_roleManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_profitMaxUnlockTime\",\"type\":\"uint256\"}],\"name\":\"newEndorsedVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_roleManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_profitMaxUnlockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_releaseDelta\",\"type\":\"uint256\"}],\"name\":\"newEndorsedVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"numEndorsedVaults\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releaseRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"removeAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"removeVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_canEndorse\",\"type\":\"bool\"}],\"name\":\"setEndorser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_canTag\",\"type\":\"bool\"}],\"name\":\"setTagger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tag\",\"type\":\"string\"}],\"name\":\"tagVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"taggers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaultInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"releaseVersion\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"vaultType\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"deploymentTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// YearnV3RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use YearnV3RegistryMetaData.ABI instead.
var YearnV3RegistryABI = YearnV3RegistryMetaData.ABI

// YearnV3Registry is an auto generated Go binding around an Ethereum contract.
type YearnV3Registry struct {
	YearnV3RegistryCaller     // Read-only binding to the contract
	YearnV3RegistryTransactor // Write-only binding to the contract
	YearnV3RegistryFilterer   // Log filterer for contract events
}

// YearnV3RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type YearnV3RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YearnV3RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YearnV3RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YearnV3RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YearnV3RegistrySession struct {
	Contract     *YearnV3Registry  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YearnV3RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YearnV3RegistryCallerSession struct {
	Contract *YearnV3RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// YearnV3RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YearnV3RegistryTransactorSession struct {
	Contract     *YearnV3RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// YearnV3RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type YearnV3RegistryRaw struct {
	Contract *YearnV3Registry // Generic contract binding to access the raw methods on
}

// YearnV3RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YearnV3RegistryCallerRaw struct {
	Contract *YearnV3RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// YearnV3RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YearnV3RegistryTransactorRaw struct {
	Contract *YearnV3RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYearnV3Registry creates a new instance of YearnV3Registry, bound to a specific deployed contract.
func NewYearnV3Registry(address common.Address, backend bind.ContractBackend) (*YearnV3Registry, error) {
	contract, err := bindYearnV3Registry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YearnV3Registry{YearnV3RegistryCaller: YearnV3RegistryCaller{contract: contract}, YearnV3RegistryTransactor: YearnV3RegistryTransactor{contract: contract}, YearnV3RegistryFilterer: YearnV3RegistryFilterer{contract: contract}}, nil
}

// NewYearnV3RegistryCaller creates a new read-only instance of YearnV3Registry, bound to a specific deployed contract.
func NewYearnV3RegistryCaller(address common.Address, caller bind.ContractCaller) (*YearnV3RegistryCaller, error) {
	contract, err := bindYearnV3Registry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryCaller{contract: contract}, nil
}

// NewYearnV3RegistryTransactor creates a new write-only instance of YearnV3Registry, bound to a specific deployed contract.
func NewYearnV3RegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*YearnV3RegistryTransactor, error) {
	contract, err := bindYearnV3Registry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryTransactor{contract: contract}, nil
}

// NewYearnV3RegistryFilterer creates a new log filterer instance of YearnV3Registry, bound to a specific deployed contract.
func NewYearnV3RegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*YearnV3RegistryFilterer, error) {
	contract, err := bindYearnV3Registry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryFilterer{contract: contract}, nil
}

// bindYearnV3Registry binds a generic wrapper to an already deployed contract.
func bindYearnV3Registry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YearnV3RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnV3Registry *YearnV3RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnV3Registry.Contract.YearnV3RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnV3Registry *YearnV3RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.YearnV3RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnV3Registry *YearnV3RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.YearnV3RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YearnV3Registry *YearnV3RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YearnV3Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YearnV3Registry *YearnV3RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YearnV3Registry *YearnV3RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.contract.Transact(opts, method, params...)
}

// MULTISTRATEGYTYPE is a free data retrieval call binding the contract method 0x18d1dd83.
//
// Solidity: function MULTI_STRATEGY_TYPE() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistryCaller) MULTISTRATEGYTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "MULTI_STRATEGY_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MULTISTRATEGYTYPE is a free data retrieval call binding the contract method 0x18d1dd83.
//
// Solidity: function MULTI_STRATEGY_TYPE() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistrySession) MULTISTRATEGYTYPE() (*big.Int, error) {
	return _YearnV3Registry.Contract.MULTISTRATEGYTYPE(&_YearnV3Registry.CallOpts)
}

// MULTISTRATEGYTYPE is a free data retrieval call binding the contract method 0x18d1dd83.
//
// Solidity: function MULTI_STRATEGY_TYPE() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistryCallerSession) MULTISTRATEGYTYPE() (*big.Int, error) {
	return _YearnV3Registry.Contract.MULTISTRATEGYTYPE(&_YearnV3Registry.CallOpts)
}

// SINGLESTRATEGYTYPE is a free data retrieval call binding the contract method 0x98a5e07b.
//
// Solidity: function SINGLE_STRATEGY_TYPE() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistryCaller) SINGLESTRATEGYTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "SINGLE_STRATEGY_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SINGLESTRATEGYTYPE is a free data retrieval call binding the contract method 0x98a5e07b.
//
// Solidity: function SINGLE_STRATEGY_TYPE() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistrySession) SINGLESTRATEGYTYPE() (*big.Int, error) {
	return _YearnV3Registry.Contract.SINGLESTRATEGYTYPE(&_YearnV3Registry.CallOpts)
}

// SINGLESTRATEGYTYPE is a free data retrieval call binding the contract method 0x98a5e07b.
//
// Solidity: function SINGLE_STRATEGY_TYPE() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistryCallerSession) SINGLESTRATEGYTYPE() (*big.Int, error) {
	return _YearnV3Registry.Contract.SINGLESTRATEGYTYPE(&_YearnV3Registry.CallOpts)
}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistryCaller) AssetIsUsed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "assetIsUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistrySession) AssetIsUsed(arg0 common.Address) (bool, error) {
	return _YearnV3Registry.Contract.AssetIsUsed(&_YearnV3Registry.CallOpts, arg0)
}

// AssetIsUsed is a free data retrieval call binding the contract method 0xac01762a.
//
// Solidity: function assetIsUsed(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistryCallerSession) AssetIsUsed(arg0 common.Address) (bool, error) {
	return _YearnV3Registry.Contract.AssetIsUsed(&_YearnV3Registry.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YearnV3Registry *YearnV3RegistryCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "assets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YearnV3Registry *YearnV3RegistrySession) Assets(arg0 *big.Int) (common.Address, error) {
	return _YearnV3Registry.Contract.Assets(&_YearnV3Registry.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets(uint256 ) view returns(address)
func (_YearnV3Registry *YearnV3RegistryCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _YearnV3Registry.Contract.Assets(&_YearnV3Registry.CallOpts, arg0)
}

// Endorsers is a free data retrieval call binding the contract method 0x3515a20b.
//
// Solidity: function endorsers(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistryCaller) Endorsers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "endorsers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Endorsers is a free data retrieval call binding the contract method 0x3515a20b.
//
// Solidity: function endorsers(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistrySession) Endorsers(arg0 common.Address) (bool, error) {
	return _YearnV3Registry.Contract.Endorsers(&_YearnV3Registry.CallOpts, arg0)
}

// Endorsers is a free data retrieval call binding the contract method 0x3515a20b.
//
// Solidity: function endorsers(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistryCallerSession) Endorsers(arg0 common.Address) (bool, error) {
	return _YearnV3Registry.Contract.Endorsers(&_YearnV3Registry.CallOpts, arg0)
}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YearnV3Registry *YearnV3RegistryCaller) GetAllEndorsedVaults(opts *bind.CallOpts) ([][]common.Address, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "getAllEndorsedVaults")

	if err != nil {
		return *new([][]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([][]common.Address)).(*[][]common.Address)

	return out0, err

}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YearnV3Registry *YearnV3RegistrySession) GetAllEndorsedVaults() ([][]common.Address, error) {
	return _YearnV3Registry.Contract.GetAllEndorsedVaults(&_YearnV3Registry.CallOpts)
}

// GetAllEndorsedVaults is a free data retrieval call binding the contract method 0x70df8ba7.
//
// Solidity: function getAllEndorsedVaults() view returns(address[][] allEndorsedVaults)
func (_YearnV3Registry *YearnV3RegistryCallerSession) GetAllEndorsedVaults() ([][]common.Address, error) {
	return _YearnV3Registry.Contract.GetAllEndorsedVaults(&_YearnV3Registry.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YearnV3Registry *YearnV3RegistryCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YearnV3Registry *YearnV3RegistrySession) GetAssets() ([]common.Address, error) {
	return _YearnV3Registry.Contract.GetAssets(&_YearnV3Registry.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[])
func (_YearnV3Registry *YearnV3RegistryCallerSession) GetAssets() ([]common.Address, error) {
	return _YearnV3Registry.Contract.GetAssets(&_YearnV3Registry.CallOpts)
}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YearnV3Registry *YearnV3RegistryCaller) GetEndorsedVaults(opts *bind.CallOpts, _asset common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "getEndorsedVaults", _asset)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YearnV3Registry *YearnV3RegistrySession) GetEndorsedVaults(_asset common.Address) ([]common.Address, error) {
	return _YearnV3Registry.Contract.GetEndorsedVaults(&_YearnV3Registry.CallOpts, _asset)
}

// GetEndorsedVaults is a free data retrieval call binding the contract method 0x53d2e949.
//
// Solidity: function getEndorsedVaults(address _asset) view returns(address[])
func (_YearnV3Registry *YearnV3RegistryCallerSession) GetEndorsedVaults(_asset common.Address) ([]common.Address, error) {
	return _YearnV3Registry.Contract.GetEndorsedVaults(&_YearnV3Registry.CallOpts, _asset)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnV3Registry *YearnV3RegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnV3Registry *YearnV3RegistrySession) Governance() (common.Address, error) {
	return _YearnV3Registry.Contract.Governance(&_YearnV3Registry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_YearnV3Registry *YearnV3RegistryCallerSession) Governance() (common.Address, error) {
	return _YearnV3Registry.Contract.Governance(&_YearnV3Registry.CallOpts)
}

// IsEndorsed is a free data retrieval call binding the contract method 0xa237e94d.
//
// Solidity: function isEndorsed(address _vault) view returns(bool)
func (_YearnV3Registry *YearnV3RegistryCaller) IsEndorsed(opts *bind.CallOpts, _vault common.Address) (bool, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "isEndorsed", _vault)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEndorsed is a free data retrieval call binding the contract method 0xa237e94d.
//
// Solidity: function isEndorsed(address _vault) view returns(bool)
func (_YearnV3Registry *YearnV3RegistrySession) IsEndorsed(_vault common.Address) (bool, error) {
	return _YearnV3Registry.Contract.IsEndorsed(&_YearnV3Registry.CallOpts, _vault)
}

// IsEndorsed is a free data retrieval call binding the contract method 0xa237e94d.
//
// Solidity: function isEndorsed(address _vault) view returns(bool)
func (_YearnV3Registry *YearnV3RegistryCallerSession) IsEndorsed(_vault common.Address) (bool, error) {
	return _YearnV3Registry.Contract.IsEndorsed(&_YearnV3Registry.CallOpts, _vault)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Registry *YearnV3RegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Registry *YearnV3RegistrySession) Name() (string, error) {
	return _YearnV3Registry.Contract.Name(&_YearnV3Registry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YearnV3Registry *YearnV3RegistryCallerSession) Name() (string, error) {
	return _YearnV3Registry.Contract.Name(&_YearnV3Registry.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistryCaller) NumAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "numAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistrySession) NumAssets() (*big.Int, error) {
	return _YearnV3Registry.Contract.NumAssets(&_YearnV3Registry.CallOpts)
}

// NumAssets is a free data retrieval call binding the contract method 0xa46fe83b.
//
// Solidity: function numAssets() view returns(uint256)
func (_YearnV3Registry *YearnV3RegistryCallerSession) NumAssets() (*big.Int, error) {
	return _YearnV3Registry.Contract.NumAssets(&_YearnV3Registry.CallOpts)
}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YearnV3Registry *YearnV3RegistryCaller) NumEndorsedVaults(opts *bind.CallOpts, _asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "numEndorsedVaults", _asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YearnV3Registry *YearnV3RegistrySession) NumEndorsedVaults(_asset common.Address) (*big.Int, error) {
	return _YearnV3Registry.Contract.NumEndorsedVaults(&_YearnV3Registry.CallOpts, _asset)
}

// NumEndorsedVaults is a free data retrieval call binding the contract method 0xb2c6161c.
//
// Solidity: function numEndorsedVaults(address _asset) view returns(uint256)
func (_YearnV3Registry *YearnV3RegistryCallerSession) NumEndorsedVaults(_asset common.Address) (*big.Int, error) {
	return _YearnV3Registry.Contract.NumEndorsedVaults(&_YearnV3Registry.CallOpts, _asset)
}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YearnV3Registry *YearnV3RegistryCaller) ReleaseRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "releaseRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YearnV3Registry *YearnV3RegistrySession) ReleaseRegistry() (common.Address, error) {
	return _YearnV3Registry.Contract.ReleaseRegistry(&_YearnV3Registry.CallOpts)
}

// ReleaseRegistry is a free data retrieval call binding the contract method 0x19ee073e.
//
// Solidity: function releaseRegistry() view returns(address)
func (_YearnV3Registry *YearnV3RegistryCallerSession) ReleaseRegistry() (common.Address, error) {
	return _YearnV3Registry.Contract.ReleaseRegistry(&_YearnV3Registry.CallOpts)
}

// Taggers is a free data retrieval call binding the contract method 0x5b25d2c8.
//
// Solidity: function taggers(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistryCaller) Taggers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "taggers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Taggers is a free data retrieval call binding the contract method 0x5b25d2c8.
//
// Solidity: function taggers(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistrySession) Taggers(arg0 common.Address) (bool, error) {
	return _YearnV3Registry.Contract.Taggers(&_YearnV3Registry.CallOpts, arg0)
}

// Taggers is a free data retrieval call binding the contract method 0x5b25d2c8.
//
// Solidity: function taggers(address ) view returns(bool)
func (_YearnV3Registry *YearnV3RegistryCallerSession) Taggers(arg0 common.Address) (bool, error) {
	return _YearnV3Registry.Contract.Taggers(&_YearnV3Registry.CallOpts, arg0)
}

// VaultInfo is a free data retrieval call binding the contract method 0x9164359a.
//
// Solidity: function vaultInfo(address ) view returns(address asset, uint96 releaseVersion, uint64 vaultType, uint128 deploymentTimestamp, uint64 index, string tag)
func (_YearnV3Registry *YearnV3RegistryCaller) VaultInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	VaultType           uint64
	DeploymentTimestamp *big.Int
	Index               uint64
	Tag                 string
}, error) {
	var out []interface{}
	err := _YearnV3Registry.contract.Call(opts, &out, "vaultInfo", arg0)

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
func (_YearnV3Registry *YearnV3RegistrySession) VaultInfo(arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	VaultType           uint64
	DeploymentTimestamp *big.Int
	Index               uint64
	Tag                 string
}, error) {
	return _YearnV3Registry.Contract.VaultInfo(&_YearnV3Registry.CallOpts, arg0)
}

// VaultInfo is a free data retrieval call binding the contract method 0x9164359a.
//
// Solidity: function vaultInfo(address ) view returns(address asset, uint96 releaseVersion, uint64 vaultType, uint128 deploymentTimestamp, uint64 index, string tag)
func (_YearnV3Registry *YearnV3RegistryCallerSession) VaultInfo(arg0 common.Address) (struct {
	Asset               common.Address
	ReleaseVersion      *big.Int
	VaultType           uint64
	DeploymentTimestamp *big.Int
	Index               uint64
	Tag                 string
}, error) {
	return _YearnV3Registry.Contract.VaultInfo(&_YearnV3Registry.CallOpts, arg0)
}

// EndorseMultiStrategyVault is a paid mutator transaction binding the contract method 0x0ab322d9.
//
// Solidity: function endorseMultiStrategyVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) EndorseMultiStrategyVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "endorseMultiStrategyVault", _vault)
}

// EndorseMultiStrategyVault is a paid mutator transaction binding the contract method 0x0ab322d9.
//
// Solidity: function endorseMultiStrategyVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistrySession) EndorseMultiStrategyVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.EndorseMultiStrategyVault(&_YearnV3Registry.TransactOpts, _vault)
}

// EndorseMultiStrategyVault is a paid mutator transaction binding the contract method 0x0ab322d9.
//
// Solidity: function endorseMultiStrategyVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) EndorseMultiStrategyVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.EndorseMultiStrategyVault(&_YearnV3Registry.TransactOpts, _vault)
}

// EndorseSingleStrategyVault is a paid mutator transaction binding the contract method 0x0f7872cc.
//
// Solidity: function endorseSingleStrategyVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) EndorseSingleStrategyVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "endorseSingleStrategyVault", _vault)
}

// EndorseSingleStrategyVault is a paid mutator transaction binding the contract method 0x0f7872cc.
//
// Solidity: function endorseSingleStrategyVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistrySession) EndorseSingleStrategyVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.EndorseSingleStrategyVault(&_YearnV3Registry.TransactOpts, _vault)
}

// EndorseSingleStrategyVault is a paid mutator transaction binding the contract method 0x0f7872cc.
//
// Solidity: function endorseSingleStrategyVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) EndorseSingleStrategyVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.EndorseSingleStrategyVault(&_YearnV3Registry.TransactOpts, _vault)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x89c6acec.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _vaultType, uint256 _deploymentTimestamp) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) EndorseVault(opts *bind.TransactOpts, _vault common.Address, _releaseDelta *big.Int, _vaultType *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "endorseVault", _vault, _releaseDelta, _vaultType, _deploymentTimestamp)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x89c6acec.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _vaultType, uint256 _deploymentTimestamp) returns()
func (_YearnV3Registry *YearnV3RegistrySession) EndorseVault(_vault common.Address, _releaseDelta *big.Int, _vaultType *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.EndorseVault(&_YearnV3Registry.TransactOpts, _vault, _releaseDelta, _vaultType, _deploymentTimestamp)
}

// EndorseVault is a paid mutator transaction binding the contract method 0x89c6acec.
//
// Solidity: function endorseVault(address _vault, uint256 _releaseDelta, uint256 _vaultType, uint256 _deploymentTimestamp) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) EndorseVault(_vault common.Address, _releaseDelta *big.Int, _vaultType *big.Int, _deploymentTimestamp *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.EndorseVault(&_YearnV3Registry.TransactOpts, _vault, _releaseDelta, _vaultType, _deploymentTimestamp)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x17bdd312.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime) returns(address _vault)
func (_YearnV3Registry *YearnV3RegistryTransactor) NewEndorsedVault(opts *bind.TransactOpts, _asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "newEndorsedVault", _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x17bdd312.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime) returns(address _vault)
func (_YearnV3Registry *YearnV3RegistrySession) NewEndorsedVault(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.NewEndorsedVault(&_YearnV3Registry.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime)
}

// NewEndorsedVault is a paid mutator transaction binding the contract method 0x17bdd312.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime) returns(address _vault)
func (_YearnV3Registry *YearnV3RegistryTransactorSession) NewEndorsedVault(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.NewEndorsedVault(&_YearnV3Registry.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime)
}

// NewEndorsedVault0 is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YearnV3Registry *YearnV3RegistryTransactor) NewEndorsedVault0(opts *bind.TransactOpts, _asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "newEndorsedVault0", _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// NewEndorsedVault0 is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YearnV3Registry *YearnV3RegistrySession) NewEndorsedVault0(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.NewEndorsedVault0(&_YearnV3Registry.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// NewEndorsedVault0 is a paid mutator transaction binding the contract method 0x7be7b20b.
//
// Solidity: function newEndorsedVault(address _asset, string _name, string _symbol, address _roleManager, uint256 _profitMaxUnlockTime, uint256 _releaseDelta) returns(address _vault)
func (_YearnV3Registry *YearnV3RegistryTransactorSession) NewEndorsedVault0(_asset common.Address, _name string, _symbol string, _roleManager common.Address, _profitMaxUnlockTime *big.Int, _releaseDelta *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.NewEndorsedVault0(&_YearnV3Registry.TransactOpts, _asset, _name, _symbol, _roleManager, _profitMaxUnlockTime, _releaseDelta)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) RemoveAsset(opts *bind.TransactOpts, _asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "removeAsset", _asset, _index)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YearnV3Registry *YearnV3RegistrySession) RemoveAsset(_asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.RemoveAsset(&_YearnV3Registry.TransactOpts, _asset, _index)
}

// RemoveAsset is a paid mutator transaction binding the contract method 0x2317ef67.
//
// Solidity: function removeAsset(address _asset, uint256 _index) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) RemoveAsset(_asset common.Address, _index *big.Int) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.RemoveAsset(&_YearnV3Registry.TransactOpts, _asset, _index)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) RemoveVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "removeVault", _vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistrySession) RemoveVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.RemoveVault(&_YearnV3Registry.TransactOpts, _vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(address _vault) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) RemoveVault(_vault common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.RemoveVault(&_YearnV3Registry.TransactOpts, _vault)
}

// SetEndorser is a paid mutator transaction binding the contract method 0x2c2a72d5.
//
// Solidity: function setEndorser(address _account, bool _canEndorse) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) SetEndorser(opts *bind.TransactOpts, _account common.Address, _canEndorse bool) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "setEndorser", _account, _canEndorse)
}

// SetEndorser is a paid mutator transaction binding the contract method 0x2c2a72d5.
//
// Solidity: function setEndorser(address _account, bool _canEndorse) returns()
func (_YearnV3Registry *YearnV3RegistrySession) SetEndorser(_account common.Address, _canEndorse bool) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.SetEndorser(&_YearnV3Registry.TransactOpts, _account, _canEndorse)
}

// SetEndorser is a paid mutator transaction binding the contract method 0x2c2a72d5.
//
// Solidity: function setEndorser(address _account, bool _canEndorse) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) SetEndorser(_account common.Address, _canEndorse bool) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.SetEndorser(&_YearnV3Registry.TransactOpts, _account, _canEndorse)
}

// SetTagger is a paid mutator transaction binding the contract method 0x2aa59c92.
//
// Solidity: function setTagger(address _account, bool _canTag) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) SetTagger(opts *bind.TransactOpts, _account common.Address, _canTag bool) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "setTagger", _account, _canTag)
}

// SetTagger is a paid mutator transaction binding the contract method 0x2aa59c92.
//
// Solidity: function setTagger(address _account, bool _canTag) returns()
func (_YearnV3Registry *YearnV3RegistrySession) SetTagger(_account common.Address, _canTag bool) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.SetTagger(&_YearnV3Registry.TransactOpts, _account, _canTag)
}

// SetTagger is a paid mutator transaction binding the contract method 0x2aa59c92.
//
// Solidity: function setTagger(address _account, bool _canTag) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) SetTagger(_account common.Address, _canTag bool) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.SetTagger(&_YearnV3Registry.TransactOpts, _account, _canTag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) TagVault(opts *bind.TransactOpts, _vault common.Address, _tag string) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "tagVault", _vault, _tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YearnV3Registry *YearnV3RegistrySession) TagVault(_vault common.Address, _tag string) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.TagVault(&_YearnV3Registry.TransactOpts, _vault, _tag)
}

// TagVault is a paid mutator transaction binding the contract method 0x60bd68f8.
//
// Solidity: function tagVault(address _vault, string _tag) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) TagVault(_vault common.Address, _tag string) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.TagVault(&_YearnV3Registry.TransactOpts, _vault, _tag)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YearnV3Registry *YearnV3RegistryTransactor) TransferGovernance(opts *bind.TransactOpts, _newGovernance common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.contract.Transact(opts, "transferGovernance", _newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YearnV3Registry *YearnV3RegistrySession) TransferGovernance(_newGovernance common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.TransferGovernance(&_YearnV3Registry.TransactOpts, _newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address _newGovernance) returns()
func (_YearnV3Registry *YearnV3RegistryTransactorSession) TransferGovernance(_newGovernance common.Address) (*types.Transaction, error) {
	return _YearnV3Registry.Contract.TransferGovernance(&_YearnV3Registry.TransactOpts, _newGovernance)
}

// YearnV3RegistryGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the YearnV3Registry contract.
type YearnV3RegistryGovernanceTransferredIterator struct {
	Event *YearnV3RegistryGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *YearnV3RegistryGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3RegistryGovernanceTransferred)
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
		it.Event = new(YearnV3RegistryGovernanceTransferred)
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
func (it *YearnV3RegistryGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3RegistryGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3RegistryGovernanceTransferred represents a GovernanceTransferred event raised by the YearnV3Registry contract.
type YearnV3RegistryGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_YearnV3Registry *YearnV3RegistryFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*YearnV3RegistryGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _YearnV3Registry.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryGovernanceTransferredIterator{contract: _YearnV3Registry.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_YearnV3Registry *YearnV3RegistryFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *YearnV3RegistryGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _YearnV3Registry.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3RegistryGovernanceTransferred)
				if err := _YearnV3Registry.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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
func (_YearnV3Registry *YearnV3RegistryFilterer) ParseGovernanceTransferred(log types.Log) (*YearnV3RegistryGovernanceTransferred, error) {
	event := new(YearnV3RegistryGovernanceTransferred)
	if err := _YearnV3Registry.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3RegistryNewEndorsedVaultIterator is returned from FilterNewEndorsedVault and is used to iterate over the raw logs and unpacked data for NewEndorsedVault events raised by the YearnV3Registry contract.
type YearnV3RegistryNewEndorsedVaultIterator struct {
	Event *YearnV3RegistryNewEndorsedVault // Event containing the contract specifics and raw log

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
func (it *YearnV3RegistryNewEndorsedVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3RegistryNewEndorsedVault)
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
		it.Event = new(YearnV3RegistryNewEndorsedVault)
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
func (it *YearnV3RegistryNewEndorsedVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3RegistryNewEndorsedVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3RegistryNewEndorsedVault represents a NewEndorsedVault event raised by the YearnV3Registry contract.
type YearnV3RegistryNewEndorsedVault struct {
	Vault          common.Address
	Asset          common.Address
	ReleaseVersion *big.Int
	VaultType      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewEndorsedVault is a free log retrieval operation binding the contract event 0xa9a7c68f108b706e545bc75ac8590730afa49f639d2e48f367105c9801c18fd2.
//
// Solidity: event NewEndorsedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnV3Registry *YearnV3RegistryFilterer) FilterNewEndorsedVault(opts *bind.FilterOpts, vault []common.Address, asset []common.Address) (*YearnV3RegistryNewEndorsedVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YearnV3Registry.contract.FilterLogs(opts, "NewEndorsedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryNewEndorsedVaultIterator{contract: _YearnV3Registry.contract, event: "NewEndorsedVault", logs: logs, sub: sub}, nil
}

// WatchNewEndorsedVault is a free log subscription operation binding the contract event 0xa9a7c68f108b706e545bc75ac8590730afa49f639d2e48f367105c9801c18fd2.
//
// Solidity: event NewEndorsedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnV3Registry *YearnV3RegistryFilterer) WatchNewEndorsedVault(opts *bind.WatchOpts, sink chan<- *YearnV3RegistryNewEndorsedVault, vault []common.Address, asset []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YearnV3Registry.contract.WatchLogs(opts, "NewEndorsedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3RegistryNewEndorsedVault)
				if err := _YearnV3Registry.contract.UnpackLog(event, "NewEndorsedVault", log); err != nil {
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
func (_YearnV3Registry *YearnV3RegistryFilterer) ParseNewEndorsedVault(log types.Log) (*YearnV3RegistryNewEndorsedVault, error) {
	event := new(YearnV3RegistryNewEndorsedVault)
	if err := _YearnV3Registry.contract.UnpackLog(event, "NewEndorsedVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3RegistryRemovedVaultIterator is returned from FilterRemovedVault and is used to iterate over the raw logs and unpacked data for RemovedVault events raised by the YearnV3Registry contract.
type YearnV3RegistryRemovedVaultIterator struct {
	Event *YearnV3RegistryRemovedVault // Event containing the contract specifics and raw log

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
func (it *YearnV3RegistryRemovedVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3RegistryRemovedVault)
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
		it.Event = new(YearnV3RegistryRemovedVault)
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
func (it *YearnV3RegistryRemovedVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3RegistryRemovedVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3RegistryRemovedVault represents a RemovedVault event raised by the YearnV3Registry contract.
type YearnV3RegistryRemovedVault struct {
	Vault          common.Address
	Asset          common.Address
	ReleaseVersion *big.Int
	VaultType      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemovedVault is a free log retrieval operation binding the contract event 0xb8d23ba050f8f00e22675f82cf3786ade63b12a46d4ea236927baf4d173c3092.
//
// Solidity: event RemovedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnV3Registry *YearnV3RegistryFilterer) FilterRemovedVault(opts *bind.FilterOpts, vault []common.Address, asset []common.Address) (*YearnV3RegistryRemovedVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YearnV3Registry.contract.FilterLogs(opts, "RemovedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryRemovedVaultIterator{contract: _YearnV3Registry.contract, event: "RemovedVault", logs: logs, sub: sub}, nil
}

// WatchRemovedVault is a free log subscription operation binding the contract event 0xb8d23ba050f8f00e22675f82cf3786ade63b12a46d4ea236927baf4d173c3092.
//
// Solidity: event RemovedVault(address indexed vault, address indexed asset, uint256 releaseVersion, uint256 vaultType)
func (_YearnV3Registry *YearnV3RegistryFilterer) WatchRemovedVault(opts *bind.WatchOpts, sink chan<- *YearnV3RegistryRemovedVault, vault []common.Address, asset []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _YearnV3Registry.contract.WatchLogs(opts, "RemovedVault", vaultRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3RegistryRemovedVault)
				if err := _YearnV3Registry.contract.UnpackLog(event, "RemovedVault", log); err != nil {
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
func (_YearnV3Registry *YearnV3RegistryFilterer) ParseRemovedVault(log types.Log) (*YearnV3RegistryRemovedVault, error) {
	event := new(YearnV3RegistryRemovedVault)
	if err := _YearnV3Registry.contract.UnpackLog(event, "RemovedVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3RegistryUpdateEndorserIterator is returned from FilterUpdateEndorser and is used to iterate over the raw logs and unpacked data for UpdateEndorser events raised by the YearnV3Registry contract.
type YearnV3RegistryUpdateEndorserIterator struct {
	Event *YearnV3RegistryUpdateEndorser // Event containing the contract specifics and raw log

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
func (it *YearnV3RegistryUpdateEndorserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3RegistryUpdateEndorser)
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
		it.Event = new(YearnV3RegistryUpdateEndorser)
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
func (it *YearnV3RegistryUpdateEndorserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3RegistryUpdateEndorserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3RegistryUpdateEndorser represents a UpdateEndorser event raised by the YearnV3Registry contract.
type YearnV3RegistryUpdateEndorser struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateEndorser is a free log retrieval operation binding the contract event 0xc93ec0e3c82bbe3866d85f7d0915cda166df7c76944b9fae88bcf11608f791bf.
//
// Solidity: event UpdateEndorser(address indexed account, bool status)
func (_YearnV3Registry *YearnV3RegistryFilterer) FilterUpdateEndorser(opts *bind.FilterOpts, account []common.Address) (*YearnV3RegistryUpdateEndorserIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnV3Registry.contract.FilterLogs(opts, "UpdateEndorser", accountRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryUpdateEndorserIterator{contract: _YearnV3Registry.contract, event: "UpdateEndorser", logs: logs, sub: sub}, nil
}

// WatchUpdateEndorser is a free log subscription operation binding the contract event 0xc93ec0e3c82bbe3866d85f7d0915cda166df7c76944b9fae88bcf11608f791bf.
//
// Solidity: event UpdateEndorser(address indexed account, bool status)
func (_YearnV3Registry *YearnV3RegistryFilterer) WatchUpdateEndorser(opts *bind.WatchOpts, sink chan<- *YearnV3RegistryUpdateEndorser, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnV3Registry.contract.WatchLogs(opts, "UpdateEndorser", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3RegistryUpdateEndorser)
				if err := _YearnV3Registry.contract.UnpackLog(event, "UpdateEndorser", log); err != nil {
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
func (_YearnV3Registry *YearnV3RegistryFilterer) ParseUpdateEndorser(log types.Log) (*YearnV3RegistryUpdateEndorser, error) {
	event := new(YearnV3RegistryUpdateEndorser)
	if err := _YearnV3Registry.contract.UnpackLog(event, "UpdateEndorser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3RegistryUpdateTaggerIterator is returned from FilterUpdateTagger and is used to iterate over the raw logs and unpacked data for UpdateTagger events raised by the YearnV3Registry contract.
type YearnV3RegistryUpdateTaggerIterator struct {
	Event *YearnV3RegistryUpdateTagger // Event containing the contract specifics and raw log

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
func (it *YearnV3RegistryUpdateTaggerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3RegistryUpdateTagger)
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
		it.Event = new(YearnV3RegistryUpdateTagger)
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
func (it *YearnV3RegistryUpdateTaggerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3RegistryUpdateTaggerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3RegistryUpdateTagger represents a UpdateTagger event raised by the YearnV3Registry contract.
type YearnV3RegistryUpdateTagger struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTagger is a free log retrieval operation binding the contract event 0xcd202ea0907016dd42e40faedc5cf5c3e6368644993f46c95990dfa7f84bfaa9.
//
// Solidity: event UpdateTagger(address indexed account, bool status)
func (_YearnV3Registry *YearnV3RegistryFilterer) FilterUpdateTagger(opts *bind.FilterOpts, account []common.Address) (*YearnV3RegistryUpdateTaggerIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnV3Registry.contract.FilterLogs(opts, "UpdateTagger", accountRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryUpdateTaggerIterator{contract: _YearnV3Registry.contract, event: "UpdateTagger", logs: logs, sub: sub}, nil
}

// WatchUpdateTagger is a free log subscription operation binding the contract event 0xcd202ea0907016dd42e40faedc5cf5c3e6368644993f46c95990dfa7f84bfaa9.
//
// Solidity: event UpdateTagger(address indexed account, bool status)
func (_YearnV3Registry *YearnV3RegistryFilterer) WatchUpdateTagger(opts *bind.WatchOpts, sink chan<- *YearnV3RegistryUpdateTagger, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _YearnV3Registry.contract.WatchLogs(opts, "UpdateTagger", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3RegistryUpdateTagger)
				if err := _YearnV3Registry.contract.UnpackLog(event, "UpdateTagger", log); err != nil {
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
func (_YearnV3Registry *YearnV3RegistryFilterer) ParseUpdateTagger(log types.Log) (*YearnV3RegistryUpdateTagger, error) {
	event := new(YearnV3RegistryUpdateTagger)
	if err := _YearnV3Registry.contract.UnpackLog(event, "UpdateTagger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YearnV3RegistryVaultTaggedIterator is returned from FilterVaultTagged and is used to iterate over the raw logs and unpacked data for VaultTagged events raised by the YearnV3Registry contract.
type YearnV3RegistryVaultTaggedIterator struct {
	Event *YearnV3RegistryVaultTagged // Event containing the contract specifics and raw log

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
func (it *YearnV3RegistryVaultTaggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YearnV3RegistryVaultTagged)
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
		it.Event = new(YearnV3RegistryVaultTagged)
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
func (it *YearnV3RegistryVaultTaggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YearnV3RegistryVaultTaggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YearnV3RegistryVaultTagged represents a VaultTagged event raised by the YearnV3Registry contract.
type YearnV3RegistryVaultTagged struct {
	Vault common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVaultTagged is a free log retrieval operation binding the contract event 0x9831985c7ff4c3d7dbd921d753150cc58a3a2a21c93795bbaac5bbf32baab3bb.
//
// Solidity: event VaultTagged(address indexed vault)
func (_YearnV3Registry *YearnV3RegistryFilterer) FilterVaultTagged(opts *bind.FilterOpts, vault []common.Address) (*YearnV3RegistryVaultTaggedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YearnV3Registry.contract.FilterLogs(opts, "VaultTagged", vaultRule)
	if err != nil {
		return nil, err
	}
	return &YearnV3RegistryVaultTaggedIterator{contract: _YearnV3Registry.contract, event: "VaultTagged", logs: logs, sub: sub}, nil
}

// WatchVaultTagged is a free log subscription operation binding the contract event 0x9831985c7ff4c3d7dbd921d753150cc58a3a2a21c93795bbaac5bbf32baab3bb.
//
// Solidity: event VaultTagged(address indexed vault)
func (_YearnV3Registry *YearnV3RegistryFilterer) WatchVaultTagged(opts *bind.WatchOpts, sink chan<- *YearnV3RegistryVaultTagged, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _YearnV3Registry.contract.WatchLogs(opts, "VaultTagged", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YearnV3RegistryVaultTagged)
				if err := _YearnV3Registry.contract.UnpackLog(event, "VaultTagged", log); err != nil {
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
func (_YearnV3Registry *YearnV3RegistryFilterer) ParseVaultTagged(log types.Log) (*YearnV3RegistryVaultTagged, error) {
	event := new(YearnV3RegistryVaultTagged)
	if err := _YearnV3Registry.contract.UnpackLog(event, "VaultTagged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
