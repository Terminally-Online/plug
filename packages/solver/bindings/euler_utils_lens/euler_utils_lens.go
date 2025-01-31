// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_utils_lens

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

// AssetPriceInfo is an auto generated low-level Go binding around an user-defined struct.
type AssetPriceInfo struct {
	QueryFailure       bool
	QueryFailureReason []byte
	Timestamp          *big.Int
	Oracle             common.Address
	Asset              common.Address
	UnitOfAccount      common.Address
	AmountIn           *big.Int
	AmountOutMid       *big.Int
	AmountOutBid       *big.Int
	AmountOutAsk       *big.Int
}

// EulerUtilsLensMetaData contains all meta data concerning the EulerUtilsLens contract.
var EulerUtilsLensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleLens\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TTL_ERROR\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_INFINITY\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_LIQUIDATION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_MORE_THAN_ONE_YEAR\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liabilityVault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"collateralValues\",\"type\":\"uint256[]\"}],\"name\":\"calculateTimeToLiquidation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowSPY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestFee\",\"type\":\"uint256\"}],\"name\":\"computeAPYs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAPY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAPY\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unitOfAccount\",\"type\":\"address\"}],\"name\":\"getAssetPriceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unitOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutAsk\",\"type\":\"uint256\"}],\"internalType\":\"structAssetPriceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getControllerAssetPriceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unitOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutAsk\",\"type\":\"uint256\"}],\"internalType\":\"structAssetPriceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleLens\",\"outputs\":[{\"internalType\":\"contractOracleLens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"tokenAllowances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"tokenBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EulerUtilsLensABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerUtilsLensMetaData.ABI instead.
var EulerUtilsLensABI = EulerUtilsLensMetaData.ABI

// EulerUtilsLens is an auto generated Go binding around an Ethereum contract.
type EulerUtilsLens struct {
	EulerUtilsLensCaller     // Read-only binding to the contract
	EulerUtilsLensTransactor // Write-only binding to the contract
	EulerUtilsLensFilterer   // Log filterer for contract events
}

// EulerUtilsLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerUtilsLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerUtilsLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerUtilsLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerUtilsLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerUtilsLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerUtilsLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerUtilsLensSession struct {
	Contract     *EulerUtilsLens   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EulerUtilsLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerUtilsLensCallerSession struct {
	Contract *EulerUtilsLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EulerUtilsLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerUtilsLensTransactorSession struct {
	Contract     *EulerUtilsLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EulerUtilsLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerUtilsLensRaw struct {
	Contract *EulerUtilsLens // Generic contract binding to access the raw methods on
}

// EulerUtilsLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerUtilsLensCallerRaw struct {
	Contract *EulerUtilsLensCaller // Generic read-only contract binding to access the raw methods on
}

// EulerUtilsLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerUtilsLensTransactorRaw struct {
	Contract *EulerUtilsLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerUtilsLens creates a new instance of EulerUtilsLens, bound to a specific deployed contract.
func NewEulerUtilsLens(address common.Address, backend bind.ContractBackend) (*EulerUtilsLens, error) {
	contract, err := bindEulerUtilsLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerUtilsLens{EulerUtilsLensCaller: EulerUtilsLensCaller{contract: contract}, EulerUtilsLensTransactor: EulerUtilsLensTransactor{contract: contract}, EulerUtilsLensFilterer: EulerUtilsLensFilterer{contract: contract}}, nil
}

// NewEulerUtilsLensCaller creates a new read-only instance of EulerUtilsLens, bound to a specific deployed contract.
func NewEulerUtilsLensCaller(address common.Address, caller bind.ContractCaller) (*EulerUtilsLensCaller, error) {
	contract, err := bindEulerUtilsLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerUtilsLensCaller{contract: contract}, nil
}

// NewEulerUtilsLensTransactor creates a new write-only instance of EulerUtilsLens, bound to a specific deployed contract.
func NewEulerUtilsLensTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerUtilsLensTransactor, error) {
	contract, err := bindEulerUtilsLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerUtilsLensTransactor{contract: contract}, nil
}

// NewEulerUtilsLensFilterer creates a new log filterer instance of EulerUtilsLens, bound to a specific deployed contract.
func NewEulerUtilsLensFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerUtilsLensFilterer, error) {
	contract, err := bindEulerUtilsLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerUtilsLensFilterer{contract: contract}, nil
}

// bindEulerUtilsLens binds a generic wrapper to an already deployed contract.
func bindEulerUtilsLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerUtilsLensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerUtilsLens *EulerUtilsLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerUtilsLens.Contract.EulerUtilsLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerUtilsLens *EulerUtilsLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerUtilsLens.Contract.EulerUtilsLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerUtilsLens *EulerUtilsLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerUtilsLens.Contract.EulerUtilsLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerUtilsLens *EulerUtilsLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerUtilsLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerUtilsLens *EulerUtilsLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerUtilsLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerUtilsLens *EulerUtilsLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerUtilsLens.Contract.contract.Transact(opts, method, params...)
}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCaller) TTLERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "TTL_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensSession) TTLERROR() (*big.Int, error) {
	return _EulerUtilsLens.Contract.TTLERROR(&_EulerUtilsLens.CallOpts)
}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCallerSession) TTLERROR() (*big.Int, error) {
	return _EulerUtilsLens.Contract.TTLERROR(&_EulerUtilsLens.CallOpts)
}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCaller) TTLINFINITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "TTL_INFINITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensSession) TTLINFINITY() (*big.Int, error) {
	return _EulerUtilsLens.Contract.TTLINFINITY(&_EulerUtilsLens.CallOpts)
}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCallerSession) TTLINFINITY() (*big.Int, error) {
	return _EulerUtilsLens.Contract.TTLINFINITY(&_EulerUtilsLens.CallOpts)
}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCaller) TTLLIQUIDATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "TTL_LIQUIDATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensSession) TTLLIQUIDATION() (*big.Int, error) {
	return _EulerUtilsLens.Contract.TTLLIQUIDATION(&_EulerUtilsLens.CallOpts)
}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCallerSession) TTLLIQUIDATION() (*big.Int, error) {
	return _EulerUtilsLens.Contract.TTLLIQUIDATION(&_EulerUtilsLens.CallOpts)
}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCaller) TTLMORETHANONEYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "TTL_MORE_THAN_ONE_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensSession) TTLMORETHANONEYEAR() (*big.Int, error) {
	return _EulerUtilsLens.Contract.TTLMORETHANONEYEAR(&_EulerUtilsLens.CallOpts)
}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCallerSession) TTLMORETHANONEYEAR() (*big.Int, error) {
	return _EulerUtilsLens.Contract.TTLMORETHANONEYEAR(&_EulerUtilsLens.CallOpts)
}

// CalculateTimeToLiquidation is a free data retrieval call binding the contract method 0x99876855.
//
// Solidity: function calculateTimeToLiquidation(address liabilityVault, uint256 liabilityValue, address[] collaterals, uint256[] collateralValues) view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCaller) CalculateTimeToLiquidation(opts *bind.CallOpts, liabilityVault common.Address, liabilityValue *big.Int, collaterals []common.Address, collateralValues []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "calculateTimeToLiquidation", liabilityVault, liabilityValue, collaterals, collateralValues)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTimeToLiquidation is a free data retrieval call binding the contract method 0x99876855.
//
// Solidity: function calculateTimeToLiquidation(address liabilityVault, uint256 liabilityValue, address[] collaterals, uint256[] collateralValues) view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensSession) CalculateTimeToLiquidation(liabilityVault common.Address, liabilityValue *big.Int, collaterals []common.Address, collateralValues []*big.Int) (*big.Int, error) {
	return _EulerUtilsLens.Contract.CalculateTimeToLiquidation(&_EulerUtilsLens.CallOpts, liabilityVault, liabilityValue, collaterals, collateralValues)
}

// CalculateTimeToLiquidation is a free data retrieval call binding the contract method 0x99876855.
//
// Solidity: function calculateTimeToLiquidation(address liabilityVault, uint256 liabilityValue, address[] collaterals, uint256[] collateralValues) view returns(int256)
func (_EulerUtilsLens *EulerUtilsLensCallerSession) CalculateTimeToLiquidation(liabilityVault common.Address, liabilityValue *big.Int, collaterals []common.Address, collateralValues []*big.Int) (*big.Int, error) {
	return _EulerUtilsLens.Contract.CalculateTimeToLiquidation(&_EulerUtilsLens.CallOpts, liabilityVault, liabilityValue, collaterals, collateralValues)
}

// ComputeAPYs is a free data retrieval call binding the contract method 0x6a88b0de.
//
// Solidity: function computeAPYs(uint256 borrowSPY, uint256 cash, uint256 borrows, uint256 interestFee) pure returns(uint256 borrowAPY, uint256 supplyAPY)
func (_EulerUtilsLens *EulerUtilsLensCaller) ComputeAPYs(opts *bind.CallOpts, borrowSPY *big.Int, cash *big.Int, borrows *big.Int, interestFee *big.Int) (struct {
	BorrowAPY *big.Int
	SupplyAPY *big.Int
}, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "computeAPYs", borrowSPY, cash, borrows, interestFee)

	outstruct := new(struct {
		BorrowAPY *big.Int
		SupplyAPY *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowAPY = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SupplyAPY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ComputeAPYs is a free data retrieval call binding the contract method 0x6a88b0de.
//
// Solidity: function computeAPYs(uint256 borrowSPY, uint256 cash, uint256 borrows, uint256 interestFee) pure returns(uint256 borrowAPY, uint256 supplyAPY)
func (_EulerUtilsLens *EulerUtilsLensSession) ComputeAPYs(borrowSPY *big.Int, cash *big.Int, borrows *big.Int, interestFee *big.Int) (struct {
	BorrowAPY *big.Int
	SupplyAPY *big.Int
}, error) {
	return _EulerUtilsLens.Contract.ComputeAPYs(&_EulerUtilsLens.CallOpts, borrowSPY, cash, borrows, interestFee)
}

// ComputeAPYs is a free data retrieval call binding the contract method 0x6a88b0de.
//
// Solidity: function computeAPYs(uint256 borrowSPY, uint256 cash, uint256 borrows, uint256 interestFee) pure returns(uint256 borrowAPY, uint256 supplyAPY)
func (_EulerUtilsLens *EulerUtilsLensCallerSession) ComputeAPYs(borrowSPY *big.Int, cash *big.Int, borrows *big.Int, interestFee *big.Int) (struct {
	BorrowAPY *big.Int
	SupplyAPY *big.Int
}, error) {
	return _EulerUtilsLens.Contract.ComputeAPYs(&_EulerUtilsLens.CallOpts, borrowSPY, cash, borrows, interestFee)
}

// GetAssetPriceInfo is a free data retrieval call binding the contract method 0x222911d0.
//
// Solidity: function getAssetPriceInfo(address asset, address unitOfAccount) view returns((bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256))
func (_EulerUtilsLens *EulerUtilsLensCaller) GetAssetPriceInfo(opts *bind.CallOpts, asset common.Address, unitOfAccount common.Address) (AssetPriceInfo, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "getAssetPriceInfo", asset, unitOfAccount)

	if err != nil {
		return *new(AssetPriceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AssetPriceInfo)).(*AssetPriceInfo)

	return out0, err

}

// GetAssetPriceInfo is a free data retrieval call binding the contract method 0x222911d0.
//
// Solidity: function getAssetPriceInfo(address asset, address unitOfAccount) view returns((bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256))
func (_EulerUtilsLens *EulerUtilsLensSession) GetAssetPriceInfo(asset common.Address, unitOfAccount common.Address) (AssetPriceInfo, error) {
	return _EulerUtilsLens.Contract.GetAssetPriceInfo(&_EulerUtilsLens.CallOpts, asset, unitOfAccount)
}

// GetAssetPriceInfo is a free data retrieval call binding the contract method 0x222911d0.
//
// Solidity: function getAssetPriceInfo(address asset, address unitOfAccount) view returns((bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256))
func (_EulerUtilsLens *EulerUtilsLensCallerSession) GetAssetPriceInfo(asset common.Address, unitOfAccount common.Address) (AssetPriceInfo, error) {
	return _EulerUtilsLens.Contract.GetAssetPriceInfo(&_EulerUtilsLens.CallOpts, asset, unitOfAccount)
}

// GetControllerAssetPriceInfo is a free data retrieval call binding the contract method 0xb0d4d6bb.
//
// Solidity: function getControllerAssetPriceInfo(address controller, address asset) view returns((bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256))
func (_EulerUtilsLens *EulerUtilsLensCaller) GetControllerAssetPriceInfo(opts *bind.CallOpts, controller common.Address, asset common.Address) (AssetPriceInfo, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "getControllerAssetPriceInfo", controller, asset)

	if err != nil {
		return *new(AssetPriceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AssetPriceInfo)).(*AssetPriceInfo)

	return out0, err

}

// GetControllerAssetPriceInfo is a free data retrieval call binding the contract method 0xb0d4d6bb.
//
// Solidity: function getControllerAssetPriceInfo(address controller, address asset) view returns((bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256))
func (_EulerUtilsLens *EulerUtilsLensSession) GetControllerAssetPriceInfo(controller common.Address, asset common.Address) (AssetPriceInfo, error) {
	return _EulerUtilsLens.Contract.GetControllerAssetPriceInfo(&_EulerUtilsLens.CallOpts, controller, asset)
}

// GetControllerAssetPriceInfo is a free data retrieval call binding the contract method 0xb0d4d6bb.
//
// Solidity: function getControllerAssetPriceInfo(address controller, address asset) view returns((bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256))
func (_EulerUtilsLens *EulerUtilsLensCallerSession) GetControllerAssetPriceInfo(controller common.Address, asset common.Address) (AssetPriceInfo, error) {
	return _EulerUtilsLens.Contract.GetControllerAssetPriceInfo(&_EulerUtilsLens.CallOpts, controller, asset)
}

// OracleLens is a free data retrieval call binding the contract method 0xc90be1e4.
//
// Solidity: function oracleLens() view returns(address)
func (_EulerUtilsLens *EulerUtilsLensCaller) OracleLens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "oracleLens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleLens is a free data retrieval call binding the contract method 0xc90be1e4.
//
// Solidity: function oracleLens() view returns(address)
func (_EulerUtilsLens *EulerUtilsLensSession) OracleLens() (common.Address, error) {
	return _EulerUtilsLens.Contract.OracleLens(&_EulerUtilsLens.CallOpts)
}

// OracleLens is a free data retrieval call binding the contract method 0xc90be1e4.
//
// Solidity: function oracleLens() view returns(address)
func (_EulerUtilsLens *EulerUtilsLensCallerSession) OracleLens() (common.Address, error) {
	return _EulerUtilsLens.Contract.OracleLens(&_EulerUtilsLens.CallOpts)
}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spender, address account, address[] tokens) view returns(uint256[])
func (_EulerUtilsLens *EulerUtilsLensCaller) TokenAllowances(opts *bind.CallOpts, spender common.Address, account common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "tokenAllowances", spender, account, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spender, address account, address[] tokens) view returns(uint256[])
func (_EulerUtilsLens *EulerUtilsLensSession) TokenAllowances(spender common.Address, account common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _EulerUtilsLens.Contract.TokenAllowances(&_EulerUtilsLens.CallOpts, spender, account, tokens)
}

// TokenAllowances is a free data retrieval call binding the contract method 0xed583506.
//
// Solidity: function tokenAllowances(address spender, address account, address[] tokens) view returns(uint256[])
func (_EulerUtilsLens *EulerUtilsLensCallerSession) TokenAllowances(spender common.Address, account common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _EulerUtilsLens.Contract.TokenAllowances(&_EulerUtilsLens.CallOpts, spender, account, tokens)
}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address account, address[] tokens) view returns(uint256[])
func (_EulerUtilsLens *EulerUtilsLensCaller) TokenBalances(opts *bind.CallOpts, account common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _EulerUtilsLens.contract.Call(opts, &out, "tokenBalances", account, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address account, address[] tokens) view returns(uint256[])
func (_EulerUtilsLens *EulerUtilsLensSession) TokenBalances(account common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _EulerUtilsLens.Contract.TokenBalances(&_EulerUtilsLens.CallOpts, account, tokens)
}

// TokenBalances is a free data retrieval call binding the contract method 0x3ad206cc.
//
// Solidity: function tokenBalances(address account, address[] tokens) view returns(uint256[])
func (_EulerUtilsLens *EulerUtilsLensCallerSession) TokenBalances(account common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _EulerUtilsLens.Contract.TokenBalances(&_EulerUtilsLens.CallOpts, account, tokens)
}
