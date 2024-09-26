// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave_v3_pool

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

// DataTypesEModeCategory is an auto generated low-level Go binding around an user-defined struct.
type DataTypesEModeCategory struct {
	Ltv                  uint16
	LiquidationThreshold uint16
	LiquidationBonus     uint16
	PriceSource          common.Address
	Label                string
}

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveData struct {
	Configuration               DataTypesReserveConfigurationMap
	LiquidityIndex              *big.Int
	CurrentLiquidityRate        *big.Int
	VariableBorrowIndex         *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	Id                          uint16
	LiquidationGracePeriodUntil *big.Int
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	AccruedToTreasury           *big.Int
	Unbacked                    *big.Int
	IsolationModeTotalDebt      *big.Int
	VirtualUnderlyingBalance    *big.Int
}

// DataTypesReserveDataLegacy is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveDataLegacy struct {
	Configuration               DataTypesReserveConfigurationMap
	LiquidityIndex              *big.Int
	CurrentLiquidityRate        *big.Int
	VariableBorrowIndex         *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	Id                          uint16
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	AccruedToTreasury           *big.Int
	Unbacked                    *big.Int
	IsolationModeTotalDebt      *big.Int
}

// DataTypesUserConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUserConfigurationMap struct {
	Data *big.Int
}

// AaveV3PoolMetaData contains all meta data concerning the AaveV3Pool contract.
var AaveV3PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"backer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"BackUnbacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"}],\"name\":\"IsolationModeTotalDebtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"MintUnbacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountMinted\",\"type\":\"uint256\"}],\"name\":\"MintedToTreasury\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RebalanceStableBorrowRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useATokens\",\"type\":\"bool\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"Supply\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumDataTypes.InterestRateMode\",\"name\":\"interestRateMode\",\"type\":\"uint8\"}],\"name\":\"SwapBorrowRateMode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"}],\"name\":\"UserEModeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BRIDGE_PROTOCOL_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TOTAL\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLASHLOAN_PREMIUM_TO_PROTOCOL\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUMBER_RESERVES\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STABLE_RATE_BORROW_SIZE_PERCENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"backUnbacked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.EModeCategory\",\"name\":\"category\",\"type\":\"tuple\"}],\"name\":\"configureEModeCategory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"dropReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceFromBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"interestRateModes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoanSimple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"getEModeCategoryData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ltv\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationBonus\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"priceSource\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"label\",\"type\":\"string\"}],\"internalType\":\"structDataTypes.EModeCategory\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEModeLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFlashLoanLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getLiquidationGracePeriod\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidationLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"}],\"name\":\"getReserveAddressById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"}],\"internalType\":\"structDataTypes.ReserveDataLegacy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveDataExtended\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"id\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"liquidationGracePeriodUntil\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"accruedToTreasury\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"unbacked\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"isolationModeTotalDebt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"virtualUnderlyingBalance\",\"type\":\"uint128\"}],\"internalType\":\"structDataTypes.ReserveData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupplyLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsBase\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserEMode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getVirtualUnderlyingBalance\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"name\":\"mintToTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"mintUnbacked\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rebalanceStableBorrowRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"}],\"name\":\"repayWithATokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"repayWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"resetIsolationModeTotalDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"}],\"name\":\"setConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"until\",\"type\":\"uint40\"}],\"name\":\"setLiquidationGracePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rateStrategyAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateStrategyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"categoryId\",\"type\":\"uint8\"}],\"name\":\"setUserEMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"supply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"supplyWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"swapToVariable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"syncIndexesState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"syncRatesState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"}],\"name\":\"updateBridgeProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"flashLoanPremiumTotal\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"flashLoanPremiumToProtocol\",\"type\":\"uint128\"}],\"name\":\"updateFlashloanPremiums\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveV3PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveV3PoolMetaData.ABI instead.
var AaveV3PoolABI = AaveV3PoolMetaData.ABI

// AaveV3Pool is an auto generated Go binding around an Ethereum contract.
type AaveV3Pool struct {
	AaveV3PoolCaller     // Read-only binding to the contract
	AaveV3PoolTransactor // Write-only binding to the contract
	AaveV3PoolFilterer   // Log filterer for contract events
}

// AaveV3PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV3PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV3PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV3PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV3PoolSession struct {
	Contract     *AaveV3Pool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AaveV3PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV3PoolCallerSession struct {
	Contract *AaveV3PoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AaveV3PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV3PoolTransactorSession struct {
	Contract     *AaveV3PoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AaveV3PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV3PoolRaw struct {
	Contract *AaveV3Pool // Generic contract binding to access the raw methods on
}

// AaveV3PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV3PoolCallerRaw struct {
	Contract *AaveV3PoolCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV3PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV3PoolTransactorRaw struct {
	Contract *AaveV3PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV3Pool creates a new instance of AaveV3Pool, bound to a specific deployed contract.
func NewAaveV3Pool(address common.Address, backend bind.ContractBackend) (*AaveV3Pool, error) {
	contract, err := bindAaveV3Pool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV3Pool{AaveV3PoolCaller: AaveV3PoolCaller{contract: contract}, AaveV3PoolTransactor: AaveV3PoolTransactor{contract: contract}, AaveV3PoolFilterer: AaveV3PoolFilterer{contract: contract}}, nil
}

// NewAaveV3PoolCaller creates a new read-only instance of AaveV3Pool, bound to a specific deployed contract.
func NewAaveV3PoolCaller(address common.Address, caller bind.ContractCaller) (*AaveV3PoolCaller, error) {
	contract, err := bindAaveV3Pool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolCaller{contract: contract}, nil
}

// NewAaveV3PoolTransactor creates a new write-only instance of AaveV3Pool, bound to a specific deployed contract.
func NewAaveV3PoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV3PoolTransactor, error) {
	contract, err := bindAaveV3Pool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolTransactor{contract: contract}, nil
}

// NewAaveV3PoolFilterer creates a new log filterer instance of AaveV3Pool, bound to a specific deployed contract.
func NewAaveV3PoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV3PoolFilterer, error) {
	contract, err := bindAaveV3Pool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolFilterer{contract: contract}, nil
}

// bindAaveV3Pool binds a generic wrapper to an already deployed contract.
func bindAaveV3Pool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3Pool *AaveV3PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3Pool.Contract.AaveV3PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3Pool *AaveV3PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.AaveV3PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3Pool *AaveV3PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.AaveV3PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3Pool *AaveV3PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3Pool *AaveV3PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3Pool *AaveV3PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV3Pool *AaveV3PoolSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveV3Pool.Contract.ADDRESSESPROVIDER(&_AaveV3Pool.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveV3Pool.Contract.ADDRESSESPROVIDER(&_AaveV3Pool.CallOpts)
}

// BRIDGEPROTOCOLFEE is a free data retrieval call binding the contract method 0x272d9072.
//
// Solidity: function BRIDGE_PROTOCOL_FEE() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCaller) BRIDGEPROTOCOLFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "BRIDGE_PROTOCOL_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BRIDGEPROTOCOLFEE is a free data retrieval call binding the contract method 0x272d9072.
//
// Solidity: function BRIDGE_PROTOCOL_FEE() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) BRIDGEPROTOCOLFEE() (*big.Int, error) {
	return _AaveV3Pool.Contract.BRIDGEPROTOCOLFEE(&_AaveV3Pool.CallOpts)
}

// BRIDGEPROTOCOLFEE is a free data retrieval call binding the contract method 0x272d9072.
//
// Solidity: function BRIDGE_PROTOCOL_FEE() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCallerSession) BRIDGEPROTOCOLFEE() (*big.Int, error) {
	return _AaveV3Pool.Contract.BRIDGEPROTOCOLFEE(&_AaveV3Pool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AaveV3Pool *AaveV3PoolCaller) FLASHLOANPREMIUMTOTAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "FLASHLOAN_PREMIUM_TOTAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AaveV3Pool *AaveV3PoolSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _AaveV3Pool.Contract.FLASHLOANPREMIUMTOTAL(&_AaveV3Pool.CallOpts)
}

// FLASHLOANPREMIUMTOTAL is a free data retrieval call binding the contract method 0x074b2e43.
//
// Solidity: function FLASHLOAN_PREMIUM_TOTAL() view returns(uint128)
func (_AaveV3Pool *AaveV3PoolCallerSession) FLASHLOANPREMIUMTOTAL() (*big.Int, error) {
	return _AaveV3Pool.Contract.FLASHLOANPREMIUMTOTAL(&_AaveV3Pool.CallOpts)
}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AaveV3Pool *AaveV3PoolCaller) FLASHLOANPREMIUMTOPROTOCOL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "FLASHLOAN_PREMIUM_TO_PROTOCOL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AaveV3Pool *AaveV3PoolSession) FLASHLOANPREMIUMTOPROTOCOL() (*big.Int, error) {
	return _AaveV3Pool.Contract.FLASHLOANPREMIUMTOPROTOCOL(&_AaveV3Pool.CallOpts)
}

// FLASHLOANPREMIUMTOPROTOCOL is a free data retrieval call binding the contract method 0x6a99c036.
//
// Solidity: function FLASHLOAN_PREMIUM_TO_PROTOCOL() view returns(uint128)
func (_AaveV3Pool *AaveV3PoolCallerSession) FLASHLOANPREMIUMTOPROTOCOL() (*big.Int, error) {
	return _AaveV3Pool.Contract.FLASHLOANPREMIUMTOPROTOCOL(&_AaveV3Pool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AaveV3Pool *AaveV3PoolCaller) MAXNUMBERRESERVES(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "MAX_NUMBER_RESERVES")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AaveV3Pool *AaveV3PoolSession) MAXNUMBERRESERVES() (uint16, error) {
	return _AaveV3Pool.Contract.MAXNUMBERRESERVES(&_AaveV3Pool.CallOpts)
}

// MAXNUMBERRESERVES is a free data retrieval call binding the contract method 0xf8119d51.
//
// Solidity: function MAX_NUMBER_RESERVES() view returns(uint16)
func (_AaveV3Pool *AaveV3PoolCallerSession) MAXNUMBERRESERVES() (uint16, error) {
	return _AaveV3Pool.Contract.MAXNUMBERRESERVES(&_AaveV3Pool.CallOpts)
}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCaller) MAXSTABLERATEBORROWSIZEPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "MAX_STABLE_RATE_BORROW_SIZE_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) MAXSTABLERATEBORROWSIZEPERCENT() (*big.Int, error) {
	return _AaveV3Pool.Contract.MAXSTABLERATEBORROWSIZEPERCENT(&_AaveV3Pool.CallOpts)
}

// MAXSTABLERATEBORROWSIZEPERCENT is a free data retrieval call binding the contract method 0xe82fec2f.
//
// Solidity: function MAX_STABLE_RATE_BORROW_SIZE_PERCENT() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCallerSession) MAXSTABLERATEBORROWSIZEPERCENT() (*big.Int, error) {
	return _AaveV3Pool.Contract.MAXSTABLERATEBORROWSIZEPERCENT(&_AaveV3Pool.CallOpts)
}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCaller) POOLREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "POOL_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) POOLREVISION() (*big.Int, error) {
	return _AaveV3Pool.Contract.POOLREVISION(&_AaveV3Pool.CallOpts)
}

// POOLREVISION is a free data retrieval call binding the contract method 0x0148170e.
//
// Solidity: function POOL_REVISION() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCallerSession) POOLREVISION() (*big.Int, error) {
	return _AaveV3Pool.Contract.POOLREVISION(&_AaveV3Pool.CallOpts)
}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) GetBorrowLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getBorrowLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolSession) GetBorrowLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetBorrowLogic(&_AaveV3Pool.CallOpts)
}

// GetBorrowLogic is a free data retrieval call binding the contract method 0x2be29fa7.
//
// Solidity: function getBorrowLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetBorrowLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetBorrowLogic(&_AaveV3Pool.CallOpts)
}

// GetBridgeLogic is a free data retrieval call binding the contract method 0xdf374c36.
//
// Solidity: function getBridgeLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) GetBridgeLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getBridgeLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeLogic is a free data retrieval call binding the contract method 0xdf374c36.
//
// Solidity: function getBridgeLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolSession) GetBridgeLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetBridgeLogic(&_AaveV3Pool.CallOpts)
}

// GetBridgeLogic is a free data retrieval call binding the contract method 0xdf374c36.
//
// Solidity: function getBridgeLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetBridgeLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetBridgeLogic(&_AaveV3Pool.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AaveV3Pool *AaveV3PoolCaller) GetConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getConfiguration", asset)

	if err != nil {
		return *new(DataTypesReserveConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveConfigurationMap)).(*DataTypesReserveConfigurationMap)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AaveV3Pool *AaveV3PoolSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _AaveV3Pool.Contract.GetConfiguration(&_AaveV3Pool.CallOpts, asset)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_AaveV3Pool *AaveV3PoolCallerSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _AaveV3Pool.Contract.GetConfiguration(&_AaveV3Pool.CallOpts, asset)
}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AaveV3Pool *AaveV3PoolCaller) GetEModeCategoryData(opts *bind.CallOpts, id uint8) (DataTypesEModeCategory, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getEModeCategoryData", id)

	if err != nil {
		return *new(DataTypesEModeCategory), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesEModeCategory)).(*DataTypesEModeCategory)

	return out0, err

}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AaveV3Pool *AaveV3PoolSession) GetEModeCategoryData(id uint8) (DataTypesEModeCategory, error) {
	return _AaveV3Pool.Contract.GetEModeCategoryData(&_AaveV3Pool.CallOpts, id)
}

// GetEModeCategoryData is a free data retrieval call binding the contract method 0x6c6f6ae1.
//
// Solidity: function getEModeCategoryData(uint8 id) view returns((uint16,uint16,uint16,address,string))
func (_AaveV3Pool *AaveV3PoolCallerSession) GetEModeCategoryData(id uint8) (DataTypesEModeCategory, error) {
	return _AaveV3Pool.Contract.GetEModeCategoryData(&_AaveV3Pool.CallOpts, id)
}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) GetEModeLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getEModeLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolSession) GetEModeLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetEModeLogic(&_AaveV3Pool.CallOpts)
}

// GetEModeLogic is a free data retrieval call binding the contract method 0xf32b9a73.
//
// Solidity: function getEModeLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetEModeLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetEModeLogic(&_AaveV3Pool.CallOpts)
}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) GetFlashLoanLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getFlashLoanLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolSession) GetFlashLoanLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetFlashLoanLogic(&_AaveV3Pool.CallOpts)
}

// GetFlashLoanLogic is a free data retrieval call binding the contract method 0x348fde0f.
//
// Solidity: function getFlashLoanLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetFlashLoanLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetFlashLoanLogic(&_AaveV3Pool.CallOpts)
}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) GetLiquidationLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getLiquidationLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolSession) GetLiquidationLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetLiquidationLogic(&_AaveV3Pool.CallOpts)
}

// GetLiquidationLogic is a free data retrieval call binding the contract method 0x911a3413.
//
// Solidity: function getLiquidationLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetLiquidationLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetLiquidationLogic(&_AaveV3Pool.CallOpts)
}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) GetPoolLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getPoolLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolSession) GetPoolLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetPoolLogic(&_AaveV3Pool.CallOpts)
}

// GetPoolLogic is a free data retrieval call binding the contract method 0xd3350155.
//
// Solidity: function getPoolLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetPoolLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetPoolLogic(&_AaveV3Pool.CallOpts)
}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) GetReserveAddressById(opts *bind.CallOpts, id uint16) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getReserveAddressById", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AaveV3Pool *AaveV3PoolSession) GetReserveAddressById(id uint16) (common.Address, error) {
	return _AaveV3Pool.Contract.GetReserveAddressById(&_AaveV3Pool.CallOpts, id)
}

// GetReserveAddressById is a free data retrieval call binding the contract method 0x52751797.
//
// Solidity: function getReserveAddressById(uint16 id) view returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetReserveAddressById(id uint16) (common.Address, error) {
	return _AaveV3Pool.Contract.GetReserveAddressById(&_AaveV3Pool.CallOpts, id)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (_AaveV3Pool *AaveV3PoolCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveDataLegacy, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveDataLegacy), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveDataLegacy)).(*DataTypesReserveDataLegacy)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (_AaveV3Pool *AaveV3PoolSession) GetReserveData(asset common.Address) (DataTypesReserveDataLegacy, error) {
	return _AaveV3Pool.Contract.GetReserveData(&_AaveV3Pool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,address,address,address,address,uint128,uint128,uint128))
func (_AaveV3Pool *AaveV3PoolCallerSession) GetReserveData(asset common.Address) (DataTypesReserveDataLegacy, error) {
	return _AaveV3Pool.Contract.GetReserveData(&_AaveV3Pool.CallOpts, asset)
}

// GetReserveDataExtended is a free data retrieval call binding the contract method 0x8381995f.
//
// Solidity: function getReserveDataExtended(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,uint40,address,address,address,address,uint128,uint128,uint128,uint128))
func (_AaveV3Pool *AaveV3PoolCaller) GetReserveDataExtended(opts *bind.CallOpts, asset common.Address) (DataTypesReserveData, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getReserveDataExtended", asset)

	if err != nil {
		return *new(DataTypesReserveData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveData)).(*DataTypesReserveData)

	return out0, err

}

// GetReserveDataExtended is a free data retrieval call binding the contract method 0x8381995f.
//
// Solidity: function getReserveDataExtended(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,uint40,address,address,address,address,uint128,uint128,uint128,uint128))
func (_AaveV3Pool *AaveV3PoolSession) GetReserveDataExtended(asset common.Address) (DataTypesReserveData, error) {
	return _AaveV3Pool.Contract.GetReserveDataExtended(&_AaveV3Pool.CallOpts, asset)
}

// GetReserveDataExtended is a free data retrieval call binding the contract method 0x8381995f.
//
// Solidity: function getReserveDataExtended(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,uint16,uint40,address,address,address,address,uint128,uint128,uint128,uint128))
func (_AaveV3Pool *AaveV3PoolCallerSession) GetReserveDataExtended(asset common.Address) (DataTypesReserveData, error) {
	return _AaveV3Pool.Contract.GetReserveDataExtended(&_AaveV3Pool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getReserveNormalizedIncome", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _AaveV3Pool.Contract.GetReserveNormalizedIncome(&_AaveV3Pool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _AaveV3Pool.Contract.GetReserveNormalizedIncome(&_AaveV3Pool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getReserveNormalizedVariableDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _AaveV3Pool.Contract.GetReserveNormalizedVariableDebt(&_AaveV3Pool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _AaveV3Pool.Contract.GetReserveNormalizedVariableDebt(&_AaveV3Pool.CallOpts, asset)
}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCaller) GetReservesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getReservesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) GetReservesCount() (*big.Int, error) {
	return _AaveV3Pool.Contract.GetReservesCount(&_AaveV3Pool.CallOpts)
}

// GetReservesCount is a free data retrieval call binding the contract method 0x72218d04.
//
// Solidity: function getReservesCount() view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetReservesCount() (*big.Int, error) {
	return _AaveV3Pool.Contract.GetReservesCount(&_AaveV3Pool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveV3Pool *AaveV3PoolCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveV3Pool *AaveV3PoolSession) GetReservesList() ([]common.Address, error) {
	return _AaveV3Pool.Contract.GetReservesList(&_AaveV3Pool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_AaveV3Pool *AaveV3PoolCallerSession) GetReservesList() ([]common.Address, error) {
	return _AaveV3Pool.Contract.GetReservesList(&_AaveV3Pool.CallOpts)
}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCaller) GetSupplyLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getSupplyLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolSession) GetSupplyLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetSupplyLogic(&_AaveV3Pool.CallOpts)
}

// GetSupplyLogic is a free data retrieval call binding the contract method 0x870e7744.
//
// Solidity: function getSupplyLogic() pure returns(address)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetSupplyLogic() (common.Address, error) {
	return _AaveV3Pool.Contract.GetSupplyLogic(&_AaveV3Pool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveV3Pool *AaveV3PoolCaller) GetUserAccountData(opts *bind.CallOpts, user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getUserAccountData", user)

	outstruct := new(struct {
		TotalCollateralBase         *big.Int
		TotalDebtBase               *big.Int
		AvailableBorrowsBase        *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralBase = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDebtBase = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsBase = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentLiquidationThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveV3Pool *AaveV3PoolSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AaveV3Pool.Contract.GetUserAccountData(&_AaveV3Pool.CallOpts, user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralBase, uint256 totalDebtBase, uint256 availableBorrowsBase, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralBase         *big.Int
	TotalDebtBase               *big.Int
	AvailableBorrowsBase        *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _AaveV3Pool.Contract.GetUserAccountData(&_AaveV3Pool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveV3Pool *AaveV3PoolCaller) GetUserConfiguration(opts *bind.CallOpts, user common.Address) (DataTypesUserConfigurationMap, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getUserConfiguration", user)

	if err != nil {
		return *new(DataTypesUserConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesUserConfigurationMap)).(*DataTypesUserConfigurationMap)

	return out0, err

}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveV3Pool *AaveV3PoolSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AaveV3Pool.Contract.GetUserConfiguration(&_AaveV3Pool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_AaveV3Pool *AaveV3PoolCallerSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _AaveV3Pool.Contract.GetUserConfiguration(&_AaveV3Pool.CallOpts, user)
}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCaller) GetUserEMode(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getUserEMode", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) GetUserEMode(user common.Address) (*big.Int, error) {
	return _AaveV3Pool.Contract.GetUserEMode(&_AaveV3Pool.CallOpts, user)
}

// GetUserEMode is a free data retrieval call binding the contract method 0xeddf1b79.
//
// Solidity: function getUserEMode(address user) view returns(uint256)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetUserEMode(user common.Address) (*big.Int, error) {
	return _AaveV3Pool.Contract.GetUserEMode(&_AaveV3Pool.CallOpts, user)
}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_AaveV3Pool *AaveV3PoolCaller) GetVirtualUnderlyingBalance(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AaveV3Pool.contract.Call(opts, &out, "getVirtualUnderlyingBalance", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_AaveV3Pool *AaveV3PoolSession) GetVirtualUnderlyingBalance(asset common.Address) (*big.Int, error) {
	return _AaveV3Pool.Contract.GetVirtualUnderlyingBalance(&_AaveV3Pool.CallOpts, asset)
}

// GetVirtualUnderlyingBalance is a free data retrieval call binding the contract method 0x6fb07f96.
//
// Solidity: function getVirtualUnderlyingBalance(address asset) view returns(uint128)
func (_AaveV3Pool *AaveV3PoolCallerSession) GetVirtualUnderlyingBalance(asset common.Address) (*big.Int, error) {
	return _AaveV3Pool.Contract.GetVirtualUnderlyingBalance(&_AaveV3Pool.CallOpts, asset)
}

// BackUnbacked is a paid mutator transaction binding the contract method 0xd65dc7a1.
//
// Solidity: function backUnbacked(address asset, uint256 amount, uint256 fee) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactor) BackUnbacked(opts *bind.TransactOpts, asset common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "backUnbacked", asset, amount, fee)
}

// BackUnbacked is a paid mutator transaction binding the contract method 0xd65dc7a1.
//
// Solidity: function backUnbacked(address asset, uint256 amount, uint256 fee) returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) BackUnbacked(asset common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.BackUnbacked(&_AaveV3Pool.TransactOpts, asset, amount, fee)
}

// BackUnbacked is a paid mutator transaction binding the contract method 0xd65dc7a1.
//
// Solidity: function backUnbacked(address asset, uint256 amount, uint256 fee) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactorSession) BackUnbacked(asset common.Address, amount *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.BackUnbacked(&_AaveV3Pool.TransactOpts, asset, amount, fee)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AaveV3Pool *AaveV3PoolSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Borrow(&_AaveV3Pool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Borrow(&_AaveV3Pool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0xd579ea7d.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,address,string) category) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) ConfigureEModeCategory(opts *bind.TransactOpts, id uint8, category DataTypesEModeCategory) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "configureEModeCategory", id, category)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0xd579ea7d.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,address,string) category) returns()
func (_AaveV3Pool *AaveV3PoolSession) ConfigureEModeCategory(id uint8, category DataTypesEModeCategory) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.ConfigureEModeCategory(&_AaveV3Pool.TransactOpts, id, category)
}

// ConfigureEModeCategory is a paid mutator transaction binding the contract method 0xd579ea7d.
//
// Solidity: function configureEModeCategory(uint8 id, (uint16,uint16,uint16,address,string) category) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) ConfigureEModeCategory(id uint8, category DataTypesEModeCategory) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.ConfigureEModeCategory(&_AaveV3Pool.TransactOpts, id, category)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Deposit(&_AaveV3Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Deposit(&_AaveV3Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) DropReserve(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "dropReserve", asset)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AaveV3Pool *AaveV3PoolSession) DropReserve(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.DropReserve(&_AaveV3Pool.TransactOpts, asset)
}

// DropReserve is a paid mutator transaction binding the contract method 0x63c9b860.
//
// Solidity: function dropReserve(address asset) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) DropReserve(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.DropReserve(&_AaveV3Pool.TransactOpts, asset)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) FinalizeTransfer(opts *bind.TransactOpts, asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "finalizeTransfer", asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AaveV3Pool *AaveV3PoolSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.FinalizeTransfer(&_AaveV3Pool.TransactOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.FinalizeTransfer(&_AaveV3Pool.TransactOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) FlashLoan(opts *bind.TransactOpts, receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "flashLoan", receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.FlashLoan(&_AaveV3Pool.TransactOpts, receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] interestRateModes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, interestRateModes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.FlashLoan(&_AaveV3Pool.TransactOpts, receiverAddress, assets, amounts, interestRateModes, onBehalfOf, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) FlashLoanSimple(opts *bind.TransactOpts, receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "flashLoanSimple", receiverAddress, asset, amount, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolSession) FlashLoanSimple(receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.FlashLoanSimple(&_AaveV3Pool.TransactOpts, receiverAddress, asset, amount, params, referralCode)
}

// FlashLoanSimple is a paid mutator transaction binding the contract method 0x42b0b77c.
//
// Solidity: function flashLoanSimple(address receiverAddress, address asset, uint256 amount, bytes params, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) FlashLoanSimple(receiverAddress common.Address, asset common.Address, amount *big.Int, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.FlashLoanSimple(&_AaveV3Pool.TransactOpts, receiverAddress, asset, amount, params, referralCode)
}

// GetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) returns(uint40)
func (_AaveV3Pool *AaveV3PoolTransactor) GetLiquidationGracePeriod(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "getLiquidationGracePeriod", asset)
}

// GetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) returns(uint40)
func (_AaveV3Pool *AaveV3PoolSession) GetLiquidationGracePeriod(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.GetLiquidationGracePeriod(&_AaveV3Pool.TransactOpts, asset)
}

// GetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0x5c9a8b18.
//
// Solidity: function getLiquidationGracePeriod(address asset) returns(uint40)
func (_AaveV3Pool *AaveV3PoolTransactorSession) GetLiquidationGracePeriod(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.GetLiquidationGracePeriod(&_AaveV3Pool.TransactOpts, asset)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) InitReserve(opts *bind.TransactOpts, asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "initReserve", asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AaveV3Pool *AaveV3PoolSession) InitReserve(asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.InitReserve(&_AaveV3Pool.TransactOpts, asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address asset, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) InitReserve(asset common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.InitReserve(&_AaveV3Pool.TransactOpts, asset, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) Initialize(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "initialize", provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AaveV3Pool *AaveV3PoolSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Initialize(&_AaveV3Pool.TransactOpts, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Initialize(&_AaveV3Pool.TransactOpts, provider)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) LiquidationCall(opts *bind.TransactOpts, collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "liquidationCall", collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AaveV3Pool *AaveV3PoolSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.LiquidationCall(&_AaveV3Pool.TransactOpts, collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.LiquidationCall(&_AaveV3Pool.TransactOpts, collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) MintToTreasury(opts *bind.TransactOpts, assets []common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "mintToTreasury", assets)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AaveV3Pool *AaveV3PoolSession) MintToTreasury(assets []common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.MintToTreasury(&_AaveV3Pool.TransactOpts, assets)
}

// MintToTreasury is a paid mutator transaction binding the contract method 0x9cd19996.
//
// Solidity: function mintToTreasury(address[] assets) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) MintToTreasury(assets []common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.MintToTreasury(&_AaveV3Pool.TransactOpts, assets)
}

// MintUnbacked is a paid mutator transaction binding the contract method 0x69a933a5.
//
// Solidity: function mintUnbacked(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) MintUnbacked(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "mintUnbacked", asset, amount, onBehalfOf, referralCode)
}

// MintUnbacked is a paid mutator transaction binding the contract method 0x69a933a5.
//
// Solidity: function mintUnbacked(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolSession) MintUnbacked(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.MintUnbacked(&_AaveV3Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// MintUnbacked is a paid mutator transaction binding the contract method 0x69a933a5.
//
// Solidity: function mintUnbacked(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) MintUnbacked(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.MintUnbacked(&_AaveV3Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) RebalanceStableBorrowRate(opts *bind.TransactOpts, asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "rebalanceStableBorrowRate", asset, user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AaveV3Pool *AaveV3PoolSession) RebalanceStableBorrowRate(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.RebalanceStableBorrowRate(&_AaveV3Pool.TransactOpts, asset, user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) RebalanceStableBorrowRate(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.RebalanceStableBorrowRate(&_AaveV3Pool.TransactOpts, asset, user)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactor) Repay(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "repay", asset, amount, interestRateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) Repay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Repay(&_AaveV3Pool.TransactOpts, asset, amount, interestRateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactorSession) Repay(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Repay(&_AaveV3Pool.TransactOpts, asset, amount, interestRateMode, onBehalfOf)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactor) RepayWithATokens(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "repayWithATokens", asset, amount, interestRateMode)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) RepayWithATokens(asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.RepayWithATokens(&_AaveV3Pool.TransactOpts, asset, amount, interestRateMode)
}

// RepayWithATokens is a paid mutator transaction binding the contract method 0x2dad97d4.
//
// Solidity: function repayWithATokens(address asset, uint256 amount, uint256 interestRateMode) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactorSession) RepayWithATokens(asset common.Address, amount *big.Int, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.RepayWithATokens(&_AaveV3Pool.TransactOpts, asset, amount, interestRateMode)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactor) RepayWithPermit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "repayWithPermit", asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) RepayWithPermit(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.RepayWithPermit(&_AaveV3Pool.TransactOpts, asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RepayWithPermit is a paid mutator transaction binding the contract method 0xee3e210b.
//
// Solidity: function repayWithPermit(address asset, uint256 amount, uint256 interestRateMode, address onBehalfOf, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactorSession) RepayWithPermit(asset common.Address, amount *big.Int, interestRateMode *big.Int, onBehalfOf common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.RepayWithPermit(&_AaveV3Pool.TransactOpts, asset, amount, interestRateMode, onBehalfOf, deadline, permitV, permitR, permitS)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) RescueTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "rescueTokens", token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AaveV3Pool *AaveV3PoolSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.RescueTokens(&_AaveV3Pool.TransactOpts, token, to, amount)
}

// RescueTokens is a paid mutator transaction binding the contract method 0xcea9d26f.
//
// Solidity: function rescueTokens(address token, address to, uint256 amount) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) RescueTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.RescueTokens(&_AaveV3Pool.TransactOpts, token, to, amount)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) ResetIsolationModeTotalDebt(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "resetIsolationModeTotalDebt", asset)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AaveV3Pool *AaveV3PoolSession) ResetIsolationModeTotalDebt(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.ResetIsolationModeTotalDebt(&_AaveV3Pool.TransactOpts, asset)
}

// ResetIsolationModeTotalDebt is a paid mutator transaction binding the contract method 0xe43e88a1.
//
// Solidity: function resetIsolationModeTotalDebt(address asset) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) ResetIsolationModeTotalDebt(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.ResetIsolationModeTotalDebt(&_AaveV3Pool.TransactOpts, asset)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SetConfiguration(opts *bind.TransactOpts, asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "setConfiguration", asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AaveV3Pool *AaveV3PoolSession) SetConfiguration(asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetConfiguration(&_AaveV3Pool.TransactOpts, asset, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xf51e435b.
//
// Solidity: function setConfiguration(address asset, (uint256) configuration) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SetConfiguration(asset common.Address, configuration DataTypesReserveConfigurationMap) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetConfiguration(&_AaveV3Pool.TransactOpts, asset, configuration)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SetLiquidationGracePeriod(opts *bind.TransactOpts, asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "setLiquidationGracePeriod", asset, until)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_AaveV3Pool *AaveV3PoolSession) SetLiquidationGracePeriod(asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetLiquidationGracePeriod(&_AaveV3Pool.TransactOpts, asset, until)
}

// SetLiquidationGracePeriod is a paid mutator transaction binding the contract method 0xb1a99e26.
//
// Solidity: function setLiquidationGracePeriod(address asset, uint40 until) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SetLiquidationGracePeriod(asset common.Address, until *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetLiquidationGracePeriod(&_AaveV3Pool.TransactOpts, asset, until)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SetReserveInterestRateStrategyAddress(opts *bind.TransactOpts, asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "setReserveInterestRateStrategyAddress", asset, rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AaveV3Pool *AaveV3PoolSession) SetReserveInterestRateStrategyAddress(asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetReserveInterestRateStrategyAddress(&_AaveV3Pool.TransactOpts, asset, rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address asset, address rateStrategyAddress) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SetReserveInterestRateStrategyAddress(asset common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetReserveInterestRateStrategyAddress(&_AaveV3Pool.TransactOpts, asset, rateStrategyAddress)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SetUserEMode(opts *bind.TransactOpts, categoryId uint8) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "setUserEMode", categoryId)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AaveV3Pool *AaveV3PoolSession) SetUserEMode(categoryId uint8) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetUserEMode(&_AaveV3Pool.TransactOpts, categoryId)
}

// SetUserEMode is a paid mutator transaction binding the contract method 0x28530a47.
//
// Solidity: function setUserEMode(uint8 categoryId) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SetUserEMode(categoryId uint8) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetUserEMode(&_AaveV3Pool.TransactOpts, categoryId)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "setUserUseReserveAsCollateral", asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AaveV3Pool *AaveV3PoolSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetUserUseReserveAsCollateral(&_AaveV3Pool.TransactOpts, asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SetUserUseReserveAsCollateral(&_AaveV3Pool.TransactOpts, asset, useAsCollateral)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) Supply(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "supply", asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Supply(&_AaveV3Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Supply is a paid mutator transaction binding the contract method 0x617ba037.
//
// Solidity: function supply(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) Supply(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Supply(&_AaveV3Pool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SupplyWithPermit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "supplyWithPermit", asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AaveV3Pool *AaveV3PoolSession) SupplyWithPermit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SupplyWithPermit(&_AaveV3Pool.TransactOpts, asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SupplyWithPermit is a paid mutator transaction binding the contract method 0x02c205f0.
//
// Solidity: function supplyWithPermit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SupplyWithPermit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SupplyWithPermit(&_AaveV3Pool.TransactOpts, asset, amount, onBehalfOf, referralCode, deadline, permitV, permitR, permitS)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 interestRateMode) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SwapBorrowRateMode(opts *bind.TransactOpts, asset common.Address, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "swapBorrowRateMode", asset, interestRateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 interestRateMode) returns()
func (_AaveV3Pool *AaveV3PoolSession) SwapBorrowRateMode(asset common.Address, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SwapBorrowRateMode(&_AaveV3Pool.TransactOpts, asset, interestRateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 interestRateMode) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SwapBorrowRateMode(asset common.Address, interestRateMode *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SwapBorrowRateMode(&_AaveV3Pool.TransactOpts, asset, interestRateMode)
}

// SwapToVariable is a paid mutator transaction binding the contract method 0x2520d5ee.
//
// Solidity: function swapToVariable(address asset, address user) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SwapToVariable(opts *bind.TransactOpts, asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "swapToVariable", asset, user)
}

// SwapToVariable is a paid mutator transaction binding the contract method 0x2520d5ee.
//
// Solidity: function swapToVariable(address asset, address user) returns()
func (_AaveV3Pool *AaveV3PoolSession) SwapToVariable(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SwapToVariable(&_AaveV3Pool.TransactOpts, asset, user)
}

// SwapToVariable is a paid mutator transaction binding the contract method 0x2520d5ee.
//
// Solidity: function swapToVariable(address asset, address user) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SwapToVariable(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SwapToVariable(&_AaveV3Pool.TransactOpts, asset, user)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SyncIndexesState(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "syncIndexesState", asset)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_AaveV3Pool *AaveV3PoolSession) SyncIndexesState(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SyncIndexesState(&_AaveV3Pool.TransactOpts, asset)
}

// SyncIndexesState is a paid mutator transaction binding the contract method 0xab2b51f6.
//
// Solidity: function syncIndexesState(address asset) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SyncIndexesState(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SyncIndexesState(&_AaveV3Pool.TransactOpts, asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) SyncRatesState(opts *bind.TransactOpts, asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "syncRatesState", asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_AaveV3Pool *AaveV3PoolSession) SyncRatesState(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SyncRatesState(&_AaveV3Pool.TransactOpts, asset)
}

// SyncRatesState is a paid mutator transaction binding the contract method 0x98c7da4e.
//
// Solidity: function syncRatesState(address asset) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) SyncRatesState(asset common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.SyncRatesState(&_AaveV3Pool.TransactOpts, asset)
}

// UpdateBridgeProtocolFee is a paid mutator transaction binding the contract method 0x3036b439.
//
// Solidity: function updateBridgeProtocolFee(uint256 protocolFee) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) UpdateBridgeProtocolFee(opts *bind.TransactOpts, protocolFee *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "updateBridgeProtocolFee", protocolFee)
}

// UpdateBridgeProtocolFee is a paid mutator transaction binding the contract method 0x3036b439.
//
// Solidity: function updateBridgeProtocolFee(uint256 protocolFee) returns()
func (_AaveV3Pool *AaveV3PoolSession) UpdateBridgeProtocolFee(protocolFee *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.UpdateBridgeProtocolFee(&_AaveV3Pool.TransactOpts, protocolFee)
}

// UpdateBridgeProtocolFee is a paid mutator transaction binding the contract method 0x3036b439.
//
// Solidity: function updateBridgeProtocolFee(uint256 protocolFee) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) UpdateBridgeProtocolFee(protocolFee *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.UpdateBridgeProtocolFee(&_AaveV3Pool.TransactOpts, protocolFee)
}

// UpdateFlashloanPremiums is a paid mutator transaction binding the contract method 0xbcb6e522.
//
// Solidity: function updateFlashloanPremiums(uint128 flashLoanPremiumTotal, uint128 flashLoanPremiumToProtocol) returns()
func (_AaveV3Pool *AaveV3PoolTransactor) UpdateFlashloanPremiums(opts *bind.TransactOpts, flashLoanPremiumTotal *big.Int, flashLoanPremiumToProtocol *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "updateFlashloanPremiums", flashLoanPremiumTotal, flashLoanPremiumToProtocol)
}

// UpdateFlashloanPremiums is a paid mutator transaction binding the contract method 0xbcb6e522.
//
// Solidity: function updateFlashloanPremiums(uint128 flashLoanPremiumTotal, uint128 flashLoanPremiumToProtocol) returns()
func (_AaveV3Pool *AaveV3PoolSession) UpdateFlashloanPremiums(flashLoanPremiumTotal *big.Int, flashLoanPremiumToProtocol *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.UpdateFlashloanPremiums(&_AaveV3Pool.TransactOpts, flashLoanPremiumTotal, flashLoanPremiumToProtocol)
}

// UpdateFlashloanPremiums is a paid mutator transaction binding the contract method 0xbcb6e522.
//
// Solidity: function updateFlashloanPremiums(uint128 flashLoanPremiumTotal, uint128 flashLoanPremiumToProtocol) returns()
func (_AaveV3Pool *AaveV3PoolTransactorSession) UpdateFlashloanPremiums(flashLoanPremiumTotal *big.Int, flashLoanPremiumToProtocol *big.Int) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.UpdateFlashloanPremiums(&_AaveV3Pool.TransactOpts, flashLoanPremiumTotal, flashLoanPremiumToProtocol)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AaveV3Pool *AaveV3PoolSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Withdraw(&_AaveV3Pool.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_AaveV3Pool *AaveV3PoolTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _AaveV3Pool.Contract.Withdraw(&_AaveV3Pool.TransactOpts, asset, amount, to)
}

// AaveV3PoolBackUnbackedIterator is returned from FilterBackUnbacked and is used to iterate over the raw logs and unpacked data for BackUnbacked events raised by the AaveV3Pool contract.
type AaveV3PoolBackUnbackedIterator struct {
	Event *AaveV3PoolBackUnbacked // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolBackUnbackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolBackUnbacked)
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
		it.Event = new(AaveV3PoolBackUnbacked)
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
func (it *AaveV3PoolBackUnbackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolBackUnbackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolBackUnbacked represents a BackUnbacked event raised by the AaveV3Pool contract.
type AaveV3PoolBackUnbacked struct {
	Reserve common.Address
	Backer  common.Address
	Amount  *big.Int
	Fee     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBackUnbacked is a free log retrieval operation binding the contract event 0x281596e92b2d974beb7d4f124df30a0b39067b096893e95011ce4bdad798b759.
//
// Solidity: event BackUnbacked(address indexed reserve, address indexed backer, uint256 amount, uint256 fee)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterBackUnbacked(opts *bind.FilterOpts, reserve []common.Address, backer []common.Address) (*AaveV3PoolBackUnbackedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var backerRule []interface{}
	for _, backerItem := range backer {
		backerRule = append(backerRule, backerItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "BackUnbacked", reserveRule, backerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolBackUnbackedIterator{contract: _AaveV3Pool.contract, event: "BackUnbacked", logs: logs, sub: sub}, nil
}

// WatchBackUnbacked is a free log subscription operation binding the contract event 0x281596e92b2d974beb7d4f124df30a0b39067b096893e95011ce4bdad798b759.
//
// Solidity: event BackUnbacked(address indexed reserve, address indexed backer, uint256 amount, uint256 fee)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchBackUnbacked(opts *bind.WatchOpts, sink chan<- *AaveV3PoolBackUnbacked, reserve []common.Address, backer []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var backerRule []interface{}
	for _, backerItem := range backer {
		backerRule = append(backerRule, backerItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "BackUnbacked", reserveRule, backerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolBackUnbacked)
				if err := _AaveV3Pool.contract.UnpackLog(event, "BackUnbacked", log); err != nil {
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

// ParseBackUnbacked is a log parse operation binding the contract event 0x281596e92b2d974beb7d4f124df30a0b39067b096893e95011ce4bdad798b759.
//
// Solidity: event BackUnbacked(address indexed reserve, address indexed backer, uint256 amount, uint256 fee)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseBackUnbacked(log types.Log) (*AaveV3PoolBackUnbacked, error) {
	event := new(AaveV3PoolBackUnbacked)
	if err := _AaveV3Pool.contract.UnpackLog(event, "BackUnbacked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the AaveV3Pool contract.
type AaveV3PoolBorrowIterator struct {
	Event *AaveV3PoolBorrow // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolBorrow)
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
		it.Event = new(AaveV3PoolBorrow)
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
func (it *AaveV3PoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolBorrow represents a Borrow event raised by the AaveV3Pool contract.
type AaveV3PoolBorrow struct {
	Reserve          common.Address
	User             common.Address
	OnBehalfOf       common.Address
	Amount           *big.Int
	InterestRateMode uint8
	BorrowRate       *big.Int
	ReferralCode     uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*AaveV3PoolBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolBorrowIterator{contract: _AaveV3Pool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *AaveV3PoolBorrow, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolBorrow)
				if err := _AaveV3Pool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xb3d084820fb1a9decffb176436bd02558d15fac9b0ddfed8c465bc7359d7dce0.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint8 interestRateMode, uint256 borrowRate, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseBorrow(log types.Log) (*AaveV3PoolBorrow, error) {
	event := new(AaveV3PoolBorrow)
	if err := _AaveV3Pool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the AaveV3Pool contract.
type AaveV3PoolFlashLoanIterator struct {
	Event *AaveV3PoolFlashLoan // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolFlashLoan)
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
		it.Event = new(AaveV3PoolFlashLoan)
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
func (it *AaveV3PoolFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolFlashLoan represents a FlashLoan event raised by the AaveV3Pool contract.
type AaveV3PoolFlashLoan struct {
	Target           common.Address
	Initiator        common.Address
	Asset            common.Address
	Amount           *big.Int
	InterestRateMode uint8
	Premium          *big.Int
	ReferralCode     uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterFlashLoan(opts *bind.FilterOpts, target []common.Address, asset []common.Address, referralCode []uint16) (*AaveV3PoolFlashLoanIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "FlashLoan", targetRule, assetRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolFlashLoanIterator{contract: _AaveV3Pool.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *AaveV3PoolFlashLoan, target []common.Address, asset []common.Address, referralCode []uint16) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "FlashLoan", targetRule, assetRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolFlashLoan)
				if err := _AaveV3Pool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0xefefaba5e921573100900a3ad9cf29f222d995fb3b6045797eaea7521bd8d6f0.
//
// Solidity: event FlashLoan(address indexed target, address initiator, address indexed asset, uint256 amount, uint8 interestRateMode, uint256 premium, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseFlashLoan(log types.Log) (*AaveV3PoolFlashLoan, error) {
	event := new(AaveV3PoolFlashLoan)
	if err := _AaveV3Pool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolIsolationModeTotalDebtUpdatedIterator is returned from FilterIsolationModeTotalDebtUpdated and is used to iterate over the raw logs and unpacked data for IsolationModeTotalDebtUpdated events raised by the AaveV3Pool contract.
type AaveV3PoolIsolationModeTotalDebtUpdatedIterator struct {
	Event *AaveV3PoolIsolationModeTotalDebtUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolIsolationModeTotalDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolIsolationModeTotalDebtUpdated)
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
		it.Event = new(AaveV3PoolIsolationModeTotalDebtUpdated)
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
func (it *AaveV3PoolIsolationModeTotalDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolIsolationModeTotalDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolIsolationModeTotalDebtUpdated represents a IsolationModeTotalDebtUpdated event raised by the AaveV3Pool contract.
type AaveV3PoolIsolationModeTotalDebtUpdated struct {
	Asset     common.Address
	TotalDebt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIsolationModeTotalDebtUpdated is a free log retrieval operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterIsolationModeTotalDebtUpdated(opts *bind.FilterOpts, asset []common.Address) (*AaveV3PoolIsolationModeTotalDebtUpdatedIterator, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "IsolationModeTotalDebtUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolIsolationModeTotalDebtUpdatedIterator{contract: _AaveV3Pool.contract, event: "IsolationModeTotalDebtUpdated", logs: logs, sub: sub}, nil
}

// WatchIsolationModeTotalDebtUpdated is a free log subscription operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchIsolationModeTotalDebtUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3PoolIsolationModeTotalDebtUpdated, asset []common.Address) (event.Subscription, error) {

	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "IsolationModeTotalDebtUpdated", assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolIsolationModeTotalDebtUpdated)
				if err := _AaveV3Pool.contract.UnpackLog(event, "IsolationModeTotalDebtUpdated", log); err != nil {
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

// ParseIsolationModeTotalDebtUpdated is a log parse operation binding the contract event 0xaef84d3b40895fd58c561f3998000f0583abb992a52fbdc99ace8e8de4d676a5.
//
// Solidity: event IsolationModeTotalDebtUpdated(address indexed asset, uint256 totalDebt)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseIsolationModeTotalDebtUpdated(log types.Log) (*AaveV3PoolIsolationModeTotalDebtUpdated, error) {
	event := new(AaveV3PoolIsolationModeTotalDebtUpdated)
	if err := _AaveV3Pool.contract.UnpackLog(event, "IsolationModeTotalDebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the AaveV3Pool contract.
type AaveV3PoolLiquidationCallIterator struct {
	Event *AaveV3PoolLiquidationCall // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolLiquidationCall)
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
		it.Event = new(AaveV3PoolLiquidationCall)
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
func (it *AaveV3PoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolLiquidationCall represents a LiquidationCall event raised by the AaveV3Pool contract.
type AaveV3PoolLiquidationCall struct {
	CollateralAsset            common.Address
	DebtAsset                  common.Address
	User                       common.Address
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (*AaveV3PoolLiquidationCallIterator, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolLiquidationCallIterator{contract: _AaveV3Pool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *AaveV3PoolLiquidationCall, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (event.Subscription, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolLiquidationCall)
				if err := _AaveV3Pool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseLiquidationCall(log types.Log) (*AaveV3PoolLiquidationCall, error) {
	event := new(AaveV3PoolLiquidationCall)
	if err := _AaveV3Pool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolMintUnbackedIterator is returned from FilterMintUnbacked and is used to iterate over the raw logs and unpacked data for MintUnbacked events raised by the AaveV3Pool contract.
type AaveV3PoolMintUnbackedIterator struct {
	Event *AaveV3PoolMintUnbacked // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolMintUnbackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolMintUnbacked)
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
		it.Event = new(AaveV3PoolMintUnbacked)
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
func (it *AaveV3PoolMintUnbackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolMintUnbackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolMintUnbacked represents a MintUnbacked event raised by the AaveV3Pool contract.
type AaveV3PoolMintUnbacked struct {
	Reserve      common.Address
	User         common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintUnbacked is a free log retrieval operation binding the contract event 0xf25af37b3d3ec226063dc9bdc103ece7eb110a50f340fe854bb7bc1b0676d7d0.
//
// Solidity: event MintUnbacked(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterMintUnbacked(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*AaveV3PoolMintUnbackedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "MintUnbacked", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolMintUnbackedIterator{contract: _AaveV3Pool.contract, event: "MintUnbacked", logs: logs, sub: sub}, nil
}

// WatchMintUnbacked is a free log subscription operation binding the contract event 0xf25af37b3d3ec226063dc9bdc103ece7eb110a50f340fe854bb7bc1b0676d7d0.
//
// Solidity: event MintUnbacked(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchMintUnbacked(opts *bind.WatchOpts, sink chan<- *AaveV3PoolMintUnbacked, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "MintUnbacked", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolMintUnbacked)
				if err := _AaveV3Pool.contract.UnpackLog(event, "MintUnbacked", log); err != nil {
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

// ParseMintUnbacked is a log parse operation binding the contract event 0xf25af37b3d3ec226063dc9bdc103ece7eb110a50f340fe854bb7bc1b0676d7d0.
//
// Solidity: event MintUnbacked(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseMintUnbacked(log types.Log) (*AaveV3PoolMintUnbacked, error) {
	event := new(AaveV3PoolMintUnbacked)
	if err := _AaveV3Pool.contract.UnpackLog(event, "MintUnbacked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolMintedToTreasuryIterator is returned from FilterMintedToTreasury and is used to iterate over the raw logs and unpacked data for MintedToTreasury events raised by the AaveV3Pool contract.
type AaveV3PoolMintedToTreasuryIterator struct {
	Event *AaveV3PoolMintedToTreasury // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolMintedToTreasuryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolMintedToTreasury)
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
		it.Event = new(AaveV3PoolMintedToTreasury)
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
func (it *AaveV3PoolMintedToTreasuryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolMintedToTreasuryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolMintedToTreasury represents a MintedToTreasury event raised by the AaveV3Pool contract.
type AaveV3PoolMintedToTreasury struct {
	Reserve      common.Address
	AmountMinted *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintedToTreasury is a free log retrieval operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterMintedToTreasury(opts *bind.FilterOpts, reserve []common.Address) (*AaveV3PoolMintedToTreasuryIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "MintedToTreasury", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolMintedToTreasuryIterator{contract: _AaveV3Pool.contract, event: "MintedToTreasury", logs: logs, sub: sub}, nil
}

// WatchMintedToTreasury is a free log subscription operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchMintedToTreasury(opts *bind.WatchOpts, sink chan<- *AaveV3PoolMintedToTreasury, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "MintedToTreasury", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolMintedToTreasury)
				if err := _AaveV3Pool.contract.UnpackLog(event, "MintedToTreasury", log); err != nil {
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

// ParseMintedToTreasury is a log parse operation binding the contract event 0xbfa21aa5d5f9a1f0120a95e7c0749f389863cbdbfff531aa7339077a5bc919de.
//
// Solidity: event MintedToTreasury(address indexed reserve, uint256 amountMinted)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseMintedToTreasury(log types.Log) (*AaveV3PoolMintedToTreasury, error) {
	event := new(AaveV3PoolMintedToTreasury)
	if err := _AaveV3Pool.contract.UnpackLog(event, "MintedToTreasury", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolRebalanceStableBorrowRateIterator is returned from FilterRebalanceStableBorrowRate and is used to iterate over the raw logs and unpacked data for RebalanceStableBorrowRate events raised by the AaveV3Pool contract.
type AaveV3PoolRebalanceStableBorrowRateIterator struct {
	Event *AaveV3PoolRebalanceStableBorrowRate // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolRebalanceStableBorrowRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolRebalanceStableBorrowRate)
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
		it.Event = new(AaveV3PoolRebalanceStableBorrowRate)
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
func (it *AaveV3PoolRebalanceStableBorrowRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolRebalanceStableBorrowRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolRebalanceStableBorrowRate represents a RebalanceStableBorrowRate event raised by the AaveV3Pool contract.
type AaveV3PoolRebalanceStableBorrowRate struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRebalanceStableBorrowRate is a free log retrieval operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterRebalanceStableBorrowRate(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AaveV3PoolRebalanceStableBorrowRateIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "RebalanceStableBorrowRate", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolRebalanceStableBorrowRateIterator{contract: _AaveV3Pool.contract, event: "RebalanceStableBorrowRate", logs: logs, sub: sub}, nil
}

// WatchRebalanceStableBorrowRate is a free log subscription operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchRebalanceStableBorrowRate(opts *bind.WatchOpts, sink chan<- *AaveV3PoolRebalanceStableBorrowRate, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "RebalanceStableBorrowRate", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolRebalanceStableBorrowRate)
				if err := _AaveV3Pool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
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

// ParseRebalanceStableBorrowRate is a log parse operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseRebalanceStableBorrowRate(log types.Log) (*AaveV3PoolRebalanceStableBorrowRate, error) {
	event := new(AaveV3PoolRebalanceStableBorrowRate)
	if err := _AaveV3Pool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the AaveV3Pool contract.
type AaveV3PoolRepayIterator struct {
	Event *AaveV3PoolRepay // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolRepay)
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
		it.Event = new(AaveV3PoolRepay)
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
func (it *AaveV3PoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolRepay represents a Repay event raised by the AaveV3Pool contract.
type AaveV3PoolRepay struct {
	Reserve    common.Address
	User       common.Address
	Repayer    common.Address
	Amount     *big.Int
	UseATokens bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, repayer []common.Address) (*AaveV3PoolRepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolRepayIterator{contract: _AaveV3Pool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *AaveV3PoolRepay, reserve []common.Address, user []common.Address, repayer []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolRepay)
				if err := _AaveV3Pool.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0xa534c8dbe71f871f9f3530e97a74601fea17b426cae02e1c5aee42c96c784051.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount, bool useATokens)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseRepay(log types.Log) (*AaveV3PoolRepay, error) {
	event := new(AaveV3PoolRepay)
	if err := _AaveV3Pool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the AaveV3Pool contract.
type AaveV3PoolReserveDataUpdatedIterator struct {
	Event *AaveV3PoolReserveDataUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolReserveDataUpdated)
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
		it.Event = new(AaveV3PoolReserveDataUpdated)
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
func (it *AaveV3PoolReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolReserveDataUpdated represents a ReserveDataUpdated event raised by the AaveV3Pool contract.
type AaveV3PoolReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*AaveV3PoolReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolReserveDataUpdatedIterator{contract: _AaveV3Pool.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3PoolReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolReserveDataUpdated)
				if err := _AaveV3Pool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
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

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseReserveDataUpdated(log types.Log) (*AaveV3PoolReserveDataUpdated, error) {
	event := new(AaveV3PoolReserveDataUpdated)
	if err := _AaveV3Pool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the AaveV3Pool contract.
type AaveV3PoolReserveUsedAsCollateralDisabledIterator struct {
	Event *AaveV3PoolReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolReserveUsedAsCollateralDisabled)
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
		it.Event = new(AaveV3PoolReserveUsedAsCollateralDisabled)
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
func (it *AaveV3PoolReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the AaveV3Pool contract.
type AaveV3PoolReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AaveV3PoolReserveUsedAsCollateralDisabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolReserveUsedAsCollateralDisabledIterator{contract: _AaveV3Pool.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *AaveV3PoolReserveUsedAsCollateralDisabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolReserveUsedAsCollateralDisabled)
				if err := _AaveV3Pool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
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

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*AaveV3PoolReserveUsedAsCollateralDisabled, error) {
	event := new(AaveV3PoolReserveUsedAsCollateralDisabled)
	if err := _AaveV3Pool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the AaveV3Pool contract.
type AaveV3PoolReserveUsedAsCollateralEnabledIterator struct {
	Event *AaveV3PoolReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolReserveUsedAsCollateralEnabled)
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
		it.Event = new(AaveV3PoolReserveUsedAsCollateralEnabled)
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
func (it *AaveV3PoolReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the AaveV3Pool contract.
type AaveV3PoolReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AaveV3PoolReserveUsedAsCollateralEnabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolReserveUsedAsCollateralEnabledIterator{contract: _AaveV3Pool.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *AaveV3PoolReserveUsedAsCollateralEnabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolReserveUsedAsCollateralEnabled)
				if err := _AaveV3Pool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
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

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*AaveV3PoolReserveUsedAsCollateralEnabled, error) {
	event := new(AaveV3PoolReserveUsedAsCollateralEnabled)
	if err := _AaveV3Pool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolSupplyIterator is returned from FilterSupply and is used to iterate over the raw logs and unpacked data for Supply events raised by the AaveV3Pool contract.
type AaveV3PoolSupplyIterator struct {
	Event *AaveV3PoolSupply // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolSupplyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolSupply)
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
		it.Event = new(AaveV3PoolSupply)
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
func (it *AaveV3PoolSupplyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolSupplyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolSupply represents a Supply event raised by the AaveV3Pool contract.
type AaveV3PoolSupply struct {
	Reserve      common.Address
	User         common.Address
	OnBehalfOf   common.Address
	Amount       *big.Int
	ReferralCode uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSupply is a free log retrieval operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterSupply(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (*AaveV3PoolSupplyIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "Supply", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolSupplyIterator{contract: _AaveV3Pool.contract, event: "Supply", logs: logs, sub: sub}, nil
}

// WatchSupply is a free log subscription operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchSupply(opts *bind.WatchOpts, sink chan<- *AaveV3PoolSupply, reserve []common.Address, onBehalfOf []common.Address, referralCode []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralCodeRule []interface{}
	for _, referralCodeItem := range referralCode {
		referralCodeRule = append(referralCodeRule, referralCodeItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "Supply", reserveRule, onBehalfOfRule, referralCodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolSupply)
				if err := _AaveV3Pool.contract.UnpackLog(event, "Supply", log); err != nil {
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

// ParseSupply is a log parse operation binding the contract event 0x2b627736bca15cd5381dcf80b0bf11fd197d01a037c52b927a881a10fb73ba61.
//
// Solidity: event Supply(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referralCode)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseSupply(log types.Log) (*AaveV3PoolSupply, error) {
	event := new(AaveV3PoolSupply)
	if err := _AaveV3Pool.contract.UnpackLog(event, "Supply", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolSwapBorrowRateModeIterator is returned from FilterSwapBorrowRateMode and is used to iterate over the raw logs and unpacked data for SwapBorrowRateMode events raised by the AaveV3Pool contract.
type AaveV3PoolSwapBorrowRateModeIterator struct {
	Event *AaveV3PoolSwapBorrowRateMode // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolSwapBorrowRateModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolSwapBorrowRateMode)
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
		it.Event = new(AaveV3PoolSwapBorrowRateMode)
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
func (it *AaveV3PoolSwapBorrowRateModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolSwapBorrowRateModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolSwapBorrowRateMode represents a SwapBorrowRateMode event raised by the AaveV3Pool contract.
type AaveV3PoolSwapBorrowRateMode struct {
	Reserve          common.Address
	User             common.Address
	InterestRateMode uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwapBorrowRateMode is a free log retrieval operation binding the contract event 0x7962b394d85a534033ba2efcf43cd36de57b7ebeb3de0ca4428965d9b3ddc481.
//
// Solidity: event SwapBorrowRateMode(address indexed reserve, address indexed user, uint8 interestRateMode)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterSwapBorrowRateMode(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*AaveV3PoolSwapBorrowRateModeIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "SwapBorrowRateMode", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolSwapBorrowRateModeIterator{contract: _AaveV3Pool.contract, event: "SwapBorrowRateMode", logs: logs, sub: sub}, nil
}

// WatchSwapBorrowRateMode is a free log subscription operation binding the contract event 0x7962b394d85a534033ba2efcf43cd36de57b7ebeb3de0ca4428965d9b3ddc481.
//
// Solidity: event SwapBorrowRateMode(address indexed reserve, address indexed user, uint8 interestRateMode)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchSwapBorrowRateMode(opts *bind.WatchOpts, sink chan<- *AaveV3PoolSwapBorrowRateMode, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "SwapBorrowRateMode", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolSwapBorrowRateMode)
				if err := _AaveV3Pool.contract.UnpackLog(event, "SwapBorrowRateMode", log); err != nil {
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

// ParseSwapBorrowRateMode is a log parse operation binding the contract event 0x7962b394d85a534033ba2efcf43cd36de57b7ebeb3de0ca4428965d9b3ddc481.
//
// Solidity: event SwapBorrowRateMode(address indexed reserve, address indexed user, uint8 interestRateMode)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseSwapBorrowRateMode(log types.Log) (*AaveV3PoolSwapBorrowRateMode, error) {
	event := new(AaveV3PoolSwapBorrowRateMode)
	if err := _AaveV3Pool.contract.UnpackLog(event, "SwapBorrowRateMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolUserEModeSetIterator is returned from FilterUserEModeSet and is used to iterate over the raw logs and unpacked data for UserEModeSet events raised by the AaveV3Pool contract.
type AaveV3PoolUserEModeSetIterator struct {
	Event *AaveV3PoolUserEModeSet // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolUserEModeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolUserEModeSet)
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
		it.Event = new(AaveV3PoolUserEModeSet)
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
func (it *AaveV3PoolUserEModeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolUserEModeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolUserEModeSet represents a UserEModeSet event raised by the AaveV3Pool contract.
type AaveV3PoolUserEModeSet struct {
	User       common.Address
	CategoryId uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUserEModeSet is a free log retrieval operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterUserEModeSet(opts *bind.FilterOpts, user []common.Address) (*AaveV3PoolUserEModeSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "UserEModeSet", userRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolUserEModeSetIterator{contract: _AaveV3Pool.contract, event: "UserEModeSet", logs: logs, sub: sub}, nil
}

// WatchUserEModeSet is a free log subscription operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchUserEModeSet(opts *bind.WatchOpts, sink chan<- *AaveV3PoolUserEModeSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "UserEModeSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolUserEModeSet)
				if err := _AaveV3Pool.contract.UnpackLog(event, "UserEModeSet", log); err != nil {
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

// ParseUserEModeSet is a log parse operation binding the contract event 0xd728da875fc88944cbf17638bcbe4af0eedaef63becd1d1c57cc097eb4608d84.
//
// Solidity: event UserEModeSet(address indexed user, uint8 categoryId)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseUserEModeSet(log types.Log) (*AaveV3PoolUserEModeSet, error) {
	event := new(AaveV3PoolUserEModeSet)
	if err := _AaveV3Pool.contract.UnpackLog(event, "UserEModeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3PoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AaveV3Pool contract.
type AaveV3PoolWithdrawIterator struct {
	Event *AaveV3PoolWithdraw // Event containing the contract specifics and raw log

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
func (it *AaveV3PoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3PoolWithdraw)
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
		it.Event = new(AaveV3PoolWithdraw)
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
func (it *AaveV3PoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3PoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3PoolWithdraw represents a Withdraw event raised by the AaveV3Pool contract.
type AaveV3PoolWithdraw struct {
	Reserve common.Address
	User    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AaveV3Pool *AaveV3PoolFilterer) FilterWithdraw(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, to []common.Address) (*AaveV3PoolWithdrawIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveV3Pool.contract.FilterLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3PoolWithdrawIterator{contract: _AaveV3Pool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AaveV3Pool *AaveV3PoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AaveV3PoolWithdraw, reserve []common.Address, user []common.Address, to []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AaveV3Pool.contract.WatchLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3PoolWithdraw)
				if err := _AaveV3Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_AaveV3Pool *AaveV3PoolFilterer) ParseWithdraw(log types.Log) (*AaveV3PoolWithdraw, error) {
	event := new(AaveV3PoolWithdraw)
	if err := _AaveV3Pool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
