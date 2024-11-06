// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave_v3_ui_pool_data_provider

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

// IUiPoolDataProviderV3AggregatedReserveData is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV3AggregatedReserveData struct {
	UnderlyingAsset                common.Address
	Name                           string
	Symbol                         string
	Decimals                       *big.Int
	BaseLTVasCollateral            *big.Int
	ReserveLiquidationThreshold    *big.Int
	ReserveLiquidationBonus        *big.Int
	ReserveFactor                  *big.Int
	UsageAsCollateralEnabled       bool
	BorrowingEnabled               bool
	StableBorrowRateEnabled        bool
	IsActive                       bool
	IsFrozen                       bool
	LiquidityIndex                 *big.Int
	VariableBorrowIndex            *big.Int
	LiquidityRate                  *big.Int
	VariableBorrowRate             *big.Int
	StableBorrowRate               *big.Int
	LastUpdateTimestamp            *big.Int
	ATokenAddress                  common.Address
	StableDebtTokenAddress         common.Address
	VariableDebtTokenAddress       common.Address
	InterestRateStrategyAddress    common.Address
	AvailableLiquidity             *big.Int
	TotalPrincipalStableDebt       *big.Int
	AverageStableRate              *big.Int
	StableDebtLastUpdateTimestamp  *big.Int
	TotalScaledVariableDebt        *big.Int
	PriceInMarketReferenceCurrency *big.Int
	PriceOracle                    common.Address
	VariableRateSlope1             *big.Int
	VariableRateSlope2             *big.Int
	StableRateSlope1               *big.Int
	StableRateSlope2               *big.Int
	BaseStableBorrowRate           *big.Int
	BaseVariableBorrowRate         *big.Int
	OptimalUsageRatio              *big.Int
	IsPaused                       bool
	IsSiloedBorrowing              bool
	AccruedToTreasury              *big.Int
	Unbacked                       *big.Int
	IsolationModeTotalDebt         *big.Int
	FlashLoanEnabled               bool
	DebtCeiling                    *big.Int
	DebtCeilingDecimals            *big.Int
	EModeCategoryId                uint8
	BorrowCap                      *big.Int
	SupplyCap                      *big.Int
	EModeLtv                       uint16
	EModeLiquidationThreshold      uint16
	EModeLiquidationBonus          uint16
	EModePriceSource               common.Address
	EModeLabel                     string
	BorrowableInIsolation          bool
	VirtualAccActive               bool
	VirtualUnderlyingBalance       *big.Int
}

// IUiPoolDataProviderV3BaseCurrencyInfo is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV3BaseCurrencyInfo struct {
	MarketReferenceCurrencyUnit       *big.Int
	MarketReferenceCurrencyPriceInUsd *big.Int
	NetworkBaseTokenPriceInUsd        *big.Int
	NetworkBaseTokenPriceDecimals     uint8
}

// IUiPoolDataProviderV3UserReserveData is an auto generated low-level Go binding around an user-defined struct.
type IUiPoolDataProviderV3UserReserveData struct {
	UnderlyingAsset                 common.Address
	ScaledATokenBalance             *big.Int
	UsageAsCollateralEnabledOnUser  bool
	StableBorrowRate                *big.Int
	ScaledVariableDebt              *big.Int
	PrincipalStableDebt             *big.Int
	StableBorrowLastUpdateTimestamp *big.Int
}

// AaveV3UiPoolDataProviderMetaData contains all meta data concerning the AaveV3UiPoolDataProvider contract.
var AaveV3UiPoolDataProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEACAggregatorProxy\",\"name\":\"_networkBaseTokenPriceInUsdProxyAggregator\",\"type\":\"address\"},{\"internalType\":\"contractIEACAggregatorProxy\",\"name\":\"_marketReferenceCurrencyPriceInUsdProxyAggregator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETH_CURRENCY_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MKR_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseLTVasCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveLiquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"liquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"stableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrincipalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableDebtLastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalScaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInMarketReferenceCurrency\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimalUsageRatio\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSiloedBorrowing\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"flashLoanEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"debtCeiling\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtCeilingDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"eModeCategoryId\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"eModeLtv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"eModeLiquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"eModePriceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"eModeLabel\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"borrowableInIsolation\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"virtualAccActive\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"virtualUnderlyingBalance\",\"type\":\"uint128\"}],\"internalType\":\"structIUiPoolDataProviderV3.AggregatedReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marketReferenceCurrencyUnit\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"marketReferenceCurrencyPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"networkBaseTokenPriceInUsd\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"networkBaseTokenPriceDecimals\",\"type\":\"uint8\"}],\"internalType\":\"structIUiPoolDataProviderV3.BaseCurrencyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReservesData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"underlyingAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"scaledATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabledOnUser\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowLastUpdateTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIUiPoolDataProviderV3.UserReserveData[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketReferenceCurrencyPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"contractIEACAggregatorProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"networkBaseTokenPriceInUsdProxyAggregator\",\"outputs\":[{\"internalType\":\"contractIEACAggregatorProxy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AaveV3UiPoolDataProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveV3UiPoolDataProviderMetaData.ABI instead.
var AaveV3UiPoolDataProviderABI = AaveV3UiPoolDataProviderMetaData.ABI

// AaveV3UiPoolDataProvider is an auto generated Go binding around an Ethereum contract.
type AaveV3UiPoolDataProvider struct {
	AaveV3UiPoolDataProviderCaller     // Read-only binding to the contract
	AaveV3UiPoolDataProviderTransactor // Write-only binding to the contract
	AaveV3UiPoolDataProviderFilterer   // Log filterer for contract events
}

// AaveV3UiPoolDataProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV3UiPoolDataProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3UiPoolDataProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV3UiPoolDataProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3UiPoolDataProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV3UiPoolDataProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3UiPoolDataProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV3UiPoolDataProviderSession struct {
	Contract     *AaveV3UiPoolDataProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AaveV3UiPoolDataProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV3UiPoolDataProviderCallerSession struct {
	Contract *AaveV3UiPoolDataProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AaveV3UiPoolDataProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV3UiPoolDataProviderTransactorSession struct {
	Contract     *AaveV3UiPoolDataProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AaveV3UiPoolDataProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV3UiPoolDataProviderRaw struct {
	Contract *AaveV3UiPoolDataProvider // Generic contract binding to access the raw methods on
}

// AaveV3UiPoolDataProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV3UiPoolDataProviderCallerRaw struct {
	Contract *AaveV3UiPoolDataProviderCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV3UiPoolDataProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV3UiPoolDataProviderTransactorRaw struct {
	Contract *AaveV3UiPoolDataProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV3UiPoolDataProvider creates a new instance of AaveV3UiPoolDataProvider, bound to a specific deployed contract.
func NewAaveV3UiPoolDataProvider(address common.Address, backend bind.ContractBackend) (*AaveV3UiPoolDataProvider, error) {
	contract, err := bindAaveV3UiPoolDataProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolDataProvider{AaveV3UiPoolDataProviderCaller: AaveV3UiPoolDataProviderCaller{contract: contract}, AaveV3UiPoolDataProviderTransactor: AaveV3UiPoolDataProviderTransactor{contract: contract}, AaveV3UiPoolDataProviderFilterer: AaveV3UiPoolDataProviderFilterer{contract: contract}}, nil
}

// NewAaveV3UiPoolDataProviderCaller creates a new read-only instance of AaveV3UiPoolDataProvider, bound to a specific deployed contract.
func NewAaveV3UiPoolDataProviderCaller(address common.Address, caller bind.ContractCaller) (*AaveV3UiPoolDataProviderCaller, error) {
	contract, err := bindAaveV3UiPoolDataProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolDataProviderCaller{contract: contract}, nil
}

// NewAaveV3UiPoolDataProviderTransactor creates a new write-only instance of AaveV3UiPoolDataProvider, bound to a specific deployed contract.
func NewAaveV3UiPoolDataProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV3UiPoolDataProviderTransactor, error) {
	contract, err := bindAaveV3UiPoolDataProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolDataProviderTransactor{contract: contract}, nil
}

// NewAaveV3UiPoolDataProviderFilterer creates a new log filterer instance of AaveV3UiPoolDataProvider, bound to a specific deployed contract.
func NewAaveV3UiPoolDataProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV3UiPoolDataProviderFilterer, error) {
	contract, err := bindAaveV3UiPoolDataProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV3UiPoolDataProviderFilterer{contract: contract}, nil
}

// bindAaveV3UiPoolDataProvider binds a generic wrapper to an already deployed contract.
func bindAaveV3UiPoolDataProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveV3UiPoolDataProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3UiPoolDataProvider.Contract.AaveV3UiPoolDataProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3UiPoolDataProvider.Contract.AaveV3UiPoolDataProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3UiPoolDataProvider.Contract.AaveV3UiPoolDataProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3UiPoolDataProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3UiPoolDataProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3UiPoolDataProvider.Contract.contract.Transact(opts, method, params...)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCaller) ETHCURRENCYUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3UiPoolDataProvider.contract.Call(opts, &out, "ETH_CURRENCY_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _AaveV3UiPoolDataProvider.Contract.ETHCURRENCYUNIT(&_AaveV3UiPoolDataProvider.CallOpts)
}

// ETHCURRENCYUNIT is a free data retrieval call binding the contract method 0x0496f53a.
//
// Solidity: function ETH_CURRENCY_UNIT() view returns(uint256)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerSession) ETHCURRENCYUNIT() (*big.Int, error) {
	return _AaveV3UiPoolDataProvider.Contract.ETHCURRENCYUNIT(&_AaveV3UiPoolDataProvider.CallOpts)
}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCaller) MKRADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolDataProvider.contract.Call(opts, &out, "MKR_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderSession) MKRADDRESS() (common.Address, error) {
	return _AaveV3UiPoolDataProvider.Contract.MKRADDRESS(&_AaveV3UiPoolDataProvider.CallOpts)
}

// MKRADDRESS is a free data retrieval call binding the contract method 0x825ffd92.
//
// Solidity: function MKR_ADDRESS() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerSession) MKRADDRESS() (common.Address, error) {
	return _AaveV3UiPoolDataProvider.Contract.MKRADDRESS(&_AaveV3UiPoolDataProvider.CallOpts)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCaller) Bytes32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _AaveV3UiPoolDataProvider.contract.Call(opts, &out, "bytes32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _AaveV3UiPoolDataProvider.Contract.Bytes32ToString(&_AaveV3UiPoolDataProvider.CallOpts, _bytes32)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 _bytes32) pure returns(string)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerSession) Bytes32ToString(_bytes32 [32]byte) (string, error) {
	return _AaveV3UiPoolDataProvider.Contract.Bytes32ToString(&_AaveV3UiPoolDataProvider.CallOpts, _bytes32)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool,bool,uint128)[], (uint256,int256,int256,uint8))
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCaller) GetReservesData(opts *bind.CallOpts, provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	var out []interface{}
	err := _AaveV3UiPoolDataProvider.contract.Call(opts, &out, "getReservesData", provider)

	if err != nil {
		return *new([]IUiPoolDataProviderV3AggregatedReserveData), *new(IUiPoolDataProviderV3BaseCurrencyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiPoolDataProviderV3AggregatedReserveData)).(*[]IUiPoolDataProviderV3AggregatedReserveData)
	out1 := *abi.ConvertType(out[1], new(IUiPoolDataProviderV3BaseCurrencyInfo)).(*IUiPoolDataProviderV3BaseCurrencyInfo)

	return out0, out1, err

}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool,bool,uint128)[], (uint256,int256,int256,uint8))
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderSession) GetReservesData(provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	return _AaveV3UiPoolDataProvider.Contract.GetReservesData(&_AaveV3UiPoolDataProvider.CallOpts, provider)
}

// GetReservesData is a free data retrieval call binding the contract method 0xec489c21.
//
// Solidity: function getReservesData(address provider) view returns((address,string,string,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,bool,bool,uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint128,uint128,uint128,bool,uint256,uint256,uint8,uint256,uint256,uint16,uint16,uint16,address,string,bool,bool,uint128)[], (uint256,int256,int256,uint8))
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerSession) GetReservesData(provider common.Address) ([]IUiPoolDataProviderV3AggregatedReserveData, IUiPoolDataProviderV3BaseCurrencyInfo, error) {
	return _AaveV3UiPoolDataProvider.Contract.GetReservesData(&_AaveV3UiPoolDataProvider.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCaller) GetReservesList(opts *bind.CallOpts, provider common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolDataProvider.contract.Call(opts, &out, "getReservesList", provider)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _AaveV3UiPoolDataProvider.Contract.GetReservesList(&_AaveV3UiPoolDataProvider.CallOpts, provider)
}

// GetReservesList is a free data retrieval call binding the contract method 0x586c1442.
//
// Solidity: function getReservesList(address provider) view returns(address[])
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerSession) GetReservesList(provider common.Address) ([]common.Address, error) {
	return _AaveV3UiPoolDataProvider.Contract.GetReservesList(&_AaveV3UiPoolDataProvider.CallOpts, provider)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCaller) GetUserReservesData(opts *bind.CallOpts, provider common.Address, user common.Address) ([]IUiPoolDataProviderV3UserReserveData, uint8, error) {
	var out []interface{}
	err := _AaveV3UiPoolDataProvider.contract.Call(opts, &out, "getUserReservesData", provider, user)

	if err != nil {
		return *new([]IUiPoolDataProviderV3UserReserveData), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]IUiPoolDataProviderV3UserReserveData)).(*[]IUiPoolDataProviderV3UserReserveData)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderSession) GetUserReservesData(provider common.Address, user common.Address) ([]IUiPoolDataProviderV3UserReserveData, uint8, error) {
	return _AaveV3UiPoolDataProvider.Contract.GetUserReservesData(&_AaveV3UiPoolDataProvider.CallOpts, provider, user)
}

// GetUserReservesData is a free data retrieval call binding the contract method 0x51974cc0.
//
// Solidity: function getUserReservesData(address provider, address user) view returns((address,uint256,bool,uint256,uint256,uint256,uint256)[], uint8)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerSession) GetUserReservesData(provider common.Address, user common.Address) ([]IUiPoolDataProviderV3UserReserveData, uint8, error) {
	return _AaveV3UiPoolDataProvider.Contract.GetUserReservesData(&_AaveV3UiPoolDataProvider.CallOpts, provider, user)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCaller) MarketReferenceCurrencyPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolDataProvider.contract.Call(opts, &out, "marketReferenceCurrencyPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveV3UiPoolDataProvider.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_AaveV3UiPoolDataProvider.CallOpts)
}

// MarketReferenceCurrencyPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0xd22cf68a.
//
// Solidity: function marketReferenceCurrencyPriceInUsdProxyAggregator() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerSession) MarketReferenceCurrencyPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveV3UiPoolDataProvider.Contract.MarketReferenceCurrencyPriceInUsdProxyAggregator(&_AaveV3UiPoolDataProvider.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCaller) NetworkBaseTokenPriceInUsdProxyAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3UiPoolDataProvider.contract.Call(opts, &out, "networkBaseTokenPriceInUsdProxyAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveV3UiPoolDataProvider.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_AaveV3UiPoolDataProvider.CallOpts)
}

// NetworkBaseTokenPriceInUsdProxyAggregator is a free data retrieval call binding the contract method 0x3c1740ed.
//
// Solidity: function networkBaseTokenPriceInUsdProxyAggregator() view returns(address)
func (_AaveV3UiPoolDataProvider *AaveV3UiPoolDataProviderCallerSession) NetworkBaseTokenPriceInUsdProxyAggregator() (common.Address, error) {
	return _AaveV3UiPoolDataProvider.Contract.NetworkBaseTokenPriceInUsdProxyAggregator(&_AaveV3UiPoolDataProvider.CallOpts)
}
