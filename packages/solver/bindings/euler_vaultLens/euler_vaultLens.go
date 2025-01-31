// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_vaultLens

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

// InterestRateInfo is an auto generated low-level Go binding around an user-defined struct.
type InterestRateInfo struct {
	Cash      *big.Int
	Borrows   *big.Int
	BorrowSPY *big.Int
	BorrowAPY *big.Int
	SupplyAPY *big.Int
}

// InterestRateModelDetailedInfo is an auto generated low-level Go binding around an user-defined struct.
type InterestRateModelDetailedInfo struct {
	InterestRateModel       common.Address
	InterestRateModelType   uint8
	InterestRateModelParams []byte
}

// LTVInfo is an auto generated low-level Go binding around an user-defined struct.
type LTVInfo struct {
	Collateral            common.Address
	BorrowLTV             *big.Int
	LiquidationLTV        *big.Int
	InitialLiquidationLTV *big.Int
	TargetTimestamp       *big.Int
	RampDuration          *big.Int
}

// OracleDetailedInfo is an auto generated low-level Go binding around an user-defined struct.
type OracleDetailedInfo struct {
	Oracle     common.Address
	Name       string
	OracleInfo []byte
}

// RewardAmountInfo is an auto generated low-level Go binding around an user-defined struct.
type RewardAmountInfo struct {
	Epoch        *big.Int
	EpochStart   *big.Int
	EpochEnd     *big.Int
	RewardAmount *big.Int
}

// VaultInfoFull is an auto generated low-level Go binding around an user-defined struct.
type VaultInfoFull struct {
	Timestamp              *big.Int
	Vault                  common.Address
	VaultName              string
	VaultSymbol            string
	VaultDecimals          *big.Int
	Asset                  common.Address
	AssetName              string
	AssetSymbol            string
	AssetDecimals          *big.Int
	UnitOfAccount          common.Address
	UnitOfAccountName      string
	UnitOfAccountSymbol    string
	UnitOfAccountDecimals  *big.Int
	TotalShares            *big.Int
	TotalCash              *big.Int
	TotalBorrowed          *big.Int
	TotalAssets            *big.Int
	AccumulatedFeesShares  *big.Int
	AccumulatedFeesAssets  *big.Int
	GovernorFeeReceiver    common.Address
	ProtocolFeeReceiver    common.Address
	ProtocolFeeShare       *big.Int
	InterestFee            *big.Int
	HookedOperations       *big.Int
	ConfigFlags            *big.Int
	SupplyCap              *big.Int
	BorrowCap              *big.Int
	MaxLiquidationDiscount *big.Int
	LiquidationCoolOffTime *big.Int
	DToken                 common.Address
	Oracle                 common.Address
	InterestRateModel      common.Address
	HookTarget             common.Address
	Evc                    common.Address
	ProtocolConfig         common.Address
	BalanceTracker         common.Address
	Permit2                common.Address
	Creator                common.Address
	GovernorAdmin          common.Address
	IrmInfo                VaultInterestRateModelInfo
	CollateralLTVInfo      []LTVInfo
	LiabilityPriceInfo     AssetPriceInfo
	CollateralPriceInfo    []AssetPriceInfo
	OracleInfo             OracleDetailedInfo
	BackupAssetPriceInfo   AssetPriceInfo
	BackupAssetOracleInfo  OracleDetailedInfo
}

// VaultInterestRateModelInfo is an auto generated low-level Go binding around an user-defined struct.
type VaultInterestRateModelInfo struct {
	QueryFailure          bool
	QueryFailureReason    []byte
	Vault                 common.Address
	InterestRateModel     common.Address
	InterestRateInfo      []InterestRateInfo
	InterestRateModelInfo InterestRateModelDetailedInfo
}

// VaultRewardInfo is an auto generated low-level Go binding around an user-defined struct.
type VaultRewardInfo struct {
	Timestamp             *big.Int
	Vault                 common.Address
	Reward                common.Address
	RewardName            string
	RewardSymbol          string
	RewardDecimals        uint8
	BalanceTracker        common.Address
	EpochDuration         *big.Int
	CurrentEpoch          *big.Int
	TotalRewardedEligible *big.Int
	TotalRewardRegistered *big.Int
	TotalRewardClaimed    *big.Int
	EpochInfoPrevious     []RewardAmountInfo
	EpochInfoUpcoming     []RewardAmountInfo
}

// EulerVaultLensMetaData contains all meta data concerning the EulerVaultLens contract.
var EulerVaultLensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleLens\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_utilsLens\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_irmLens\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TTL_ERROR\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_INFINITY\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_LIQUIDATION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_MORE_THAN_ONE_YEAR\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getRecognizedCollateralsLTVInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowLTV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationLTV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rampDuration\",\"type\":\"uint256\"}],\"internalType\":\"structLTVInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numberOfEpochs\",\"type\":\"uint256\"}],\"name\":\"getRewardVaultInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"rewardName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"rewardSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"rewardDecimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardedEligible\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardRegistered\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRewardClaimed\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structRewardAmountInfo[]\",\"name\":\"epochInfoPrevious\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"internalType\":\"structRewardAmountInfo[]\",\"name\":\"epochInfoUpcoming\",\"type\":\"tuple[]\"}],\"internalType\":\"structVaultRewardInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getVaultInfoFull\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"vaultName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"vaultSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"vaultDecimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"assetSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"assetDecimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"unitOfAccount\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"unitOfAccountName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unitOfAccountSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"unitOfAccountDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedFeesShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accumulatedFeesAssets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"governorFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"protocolFeeReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeeShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hookedOperations\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"configFlags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLiquidationDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationCoolOffTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hookTarget\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"protocolConfig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governorAdmin\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSPY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAPY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAPY\",\"type\":\"uint256\"}],\"internalType\":\"structInterestRateInfo[]\",\"name\":\"interestRateInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"},{\"internalType\":\"enumInterestRateModelType\",\"name\":\"interestRateModelType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"interestRateModelParams\",\"type\":\"bytes\"}],\"internalType\":\"structInterestRateModelDetailedInfo\",\"name\":\"interestRateModelInfo\",\"type\":\"tuple\"}],\"internalType\":\"structVaultInterestRateModelInfo\",\"name\":\"irmInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowLTV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationLTV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rampDuration\",\"type\":\"uint256\"}],\"internalType\":\"structLTVInfo[]\",\"name\":\"collateralLTVInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unitOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutAsk\",\"type\":\"uint256\"}],\"internalType\":\"structAssetPriceInfo\",\"name\":\"liabilityPriceInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unitOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutAsk\",\"type\":\"uint256\"}],\"internalType\":\"structAssetPriceInfo[]\",\"name\":\"collateralPriceInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"oracleInfo\",\"type\":\"bytes\"}],\"internalType\":\"structOracleDetailedInfo\",\"name\":\"oracleInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"unitOfAccount\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutAsk\",\"type\":\"uint256\"}],\"internalType\":\"structAssetPriceInfo\",\"name\":\"backupAssetPriceInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"oracleInfo\",\"type\":\"bytes\"}],\"internalType\":\"structOracleDetailedInfo\",\"name\":\"backupAssetOracleInfo\",\"type\":\"tuple\"}],\"internalType\":\"structVaultInfoFull\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"cash\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"borrows\",\"type\":\"uint256[]\"}],\"name\":\"getVaultInterestRateModelInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSPY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAPY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAPY\",\"type\":\"uint256\"}],\"internalType\":\"structInterestRateInfo[]\",\"name\":\"interestRateInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"},{\"internalType\":\"enumInterestRateModelType\",\"name\":\"interestRateModelType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"interestRateModelParams\",\"type\":\"bytes\"}],\"internalType\":\"structInterestRateModelDetailedInfo\",\"name\":\"interestRateModelInfo\",\"type\":\"tuple\"}],\"internalType\":\"structVaultInterestRateModelInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getVaultKinkInterestRateModelInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowSPY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAPY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyAPY\",\"type\":\"uint256\"}],\"internalType\":\"structInterestRateInfo[]\",\"name\":\"interestRateInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"interestRateModel\",\"type\":\"address\"},{\"internalType\":\"enumInterestRateModelType\",\"name\":\"interestRateModelType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"interestRateModelParams\",\"type\":\"bytes\"}],\"internalType\":\"structInterestRateModelDetailedInfo\",\"name\":\"interestRateModelInfo\",\"type\":\"tuple\"}],\"internalType\":\"structVaultInterestRateModelInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"irmLens\",\"outputs\":[{\"internalType\":\"contractIRMLens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleLens\",\"outputs\":[{\"internalType\":\"contractOracleLens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"utilsLens\",\"outputs\":[{\"internalType\":\"contractUtilsLens\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EulerVaultLensABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerVaultLensMetaData.ABI instead.
var EulerVaultLensABI = EulerVaultLensMetaData.ABI

// EulerVaultLens is an auto generated Go binding around an Ethereum contract.
type EulerVaultLens struct {
	EulerVaultLensCaller     // Read-only binding to the contract
	EulerVaultLensTransactor // Write-only binding to the contract
	EulerVaultLensFilterer   // Log filterer for contract events
}

// EulerVaultLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerVaultLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerVaultLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerVaultLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerVaultLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerVaultLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerVaultLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerVaultLensSession struct {
	Contract     *EulerVaultLens   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EulerVaultLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerVaultLensCallerSession struct {
	Contract *EulerVaultLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EulerVaultLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerVaultLensTransactorSession struct {
	Contract     *EulerVaultLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EulerVaultLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerVaultLensRaw struct {
	Contract *EulerVaultLens // Generic contract binding to access the raw methods on
}

// EulerVaultLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerVaultLensCallerRaw struct {
	Contract *EulerVaultLensCaller // Generic read-only contract binding to access the raw methods on
}

// EulerVaultLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerVaultLensTransactorRaw struct {
	Contract *EulerVaultLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerVaultLens creates a new instance of EulerVaultLens, bound to a specific deployed contract.
func NewEulerVaultLens(address common.Address, backend bind.ContractBackend) (*EulerVaultLens, error) {
	contract, err := bindEulerVaultLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerVaultLens{EulerVaultLensCaller: EulerVaultLensCaller{contract: contract}, EulerVaultLensTransactor: EulerVaultLensTransactor{contract: contract}, EulerVaultLensFilterer: EulerVaultLensFilterer{contract: contract}}, nil
}

// NewEulerVaultLensCaller creates a new read-only instance of EulerVaultLens, bound to a specific deployed contract.
func NewEulerVaultLensCaller(address common.Address, caller bind.ContractCaller) (*EulerVaultLensCaller, error) {
	contract, err := bindEulerVaultLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerVaultLensCaller{contract: contract}, nil
}

// NewEulerVaultLensTransactor creates a new write-only instance of EulerVaultLens, bound to a specific deployed contract.
func NewEulerVaultLensTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerVaultLensTransactor, error) {
	contract, err := bindEulerVaultLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerVaultLensTransactor{contract: contract}, nil
}

// NewEulerVaultLensFilterer creates a new log filterer instance of EulerVaultLens, bound to a specific deployed contract.
func NewEulerVaultLensFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerVaultLensFilterer, error) {
	contract, err := bindEulerVaultLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerVaultLensFilterer{contract: contract}, nil
}

// bindEulerVaultLens binds a generic wrapper to an already deployed contract.
func bindEulerVaultLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerVaultLensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerVaultLens *EulerVaultLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerVaultLens.Contract.EulerVaultLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerVaultLens *EulerVaultLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVaultLens.Contract.EulerVaultLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerVaultLens *EulerVaultLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerVaultLens.Contract.EulerVaultLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerVaultLens *EulerVaultLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerVaultLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerVaultLens *EulerVaultLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerVaultLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerVaultLens *EulerVaultLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerVaultLens.Contract.contract.Transact(opts, method, params...)
}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerVaultLens *EulerVaultLensCaller) TTLERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "TTL_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerVaultLens *EulerVaultLensSession) TTLERROR() (*big.Int, error) {
	return _EulerVaultLens.Contract.TTLERROR(&_EulerVaultLens.CallOpts)
}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerVaultLens *EulerVaultLensCallerSession) TTLERROR() (*big.Int, error) {
	return _EulerVaultLens.Contract.TTLERROR(&_EulerVaultLens.CallOpts)
}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerVaultLens *EulerVaultLensCaller) TTLINFINITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "TTL_INFINITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerVaultLens *EulerVaultLensSession) TTLINFINITY() (*big.Int, error) {
	return _EulerVaultLens.Contract.TTLINFINITY(&_EulerVaultLens.CallOpts)
}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerVaultLens *EulerVaultLensCallerSession) TTLINFINITY() (*big.Int, error) {
	return _EulerVaultLens.Contract.TTLINFINITY(&_EulerVaultLens.CallOpts)
}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerVaultLens *EulerVaultLensCaller) TTLLIQUIDATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "TTL_LIQUIDATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerVaultLens *EulerVaultLensSession) TTLLIQUIDATION() (*big.Int, error) {
	return _EulerVaultLens.Contract.TTLLIQUIDATION(&_EulerVaultLens.CallOpts)
}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerVaultLens *EulerVaultLensCallerSession) TTLLIQUIDATION() (*big.Int, error) {
	return _EulerVaultLens.Contract.TTLLIQUIDATION(&_EulerVaultLens.CallOpts)
}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerVaultLens *EulerVaultLensCaller) TTLMORETHANONEYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "TTL_MORE_THAN_ONE_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerVaultLens *EulerVaultLensSession) TTLMORETHANONEYEAR() (*big.Int, error) {
	return _EulerVaultLens.Contract.TTLMORETHANONEYEAR(&_EulerVaultLens.CallOpts)
}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerVaultLens *EulerVaultLensCallerSession) TTLMORETHANONEYEAR() (*big.Int, error) {
	return _EulerVaultLens.Contract.TTLMORETHANONEYEAR(&_EulerVaultLens.CallOpts)
}

// GetRecognizedCollateralsLTVInfo is a free data retrieval call binding the contract method 0x6369fedb.
//
// Solidity: function getRecognizedCollateralsLTVInfo(address vault) view returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_EulerVaultLens *EulerVaultLensCaller) GetRecognizedCollateralsLTVInfo(opts *bind.CallOpts, vault common.Address) ([]LTVInfo, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "getRecognizedCollateralsLTVInfo", vault)

	if err != nil {
		return *new([]LTVInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]LTVInfo)).(*[]LTVInfo)

	return out0, err

}

// GetRecognizedCollateralsLTVInfo is a free data retrieval call binding the contract method 0x6369fedb.
//
// Solidity: function getRecognizedCollateralsLTVInfo(address vault) view returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_EulerVaultLens *EulerVaultLensSession) GetRecognizedCollateralsLTVInfo(vault common.Address) ([]LTVInfo, error) {
	return _EulerVaultLens.Contract.GetRecognizedCollateralsLTVInfo(&_EulerVaultLens.CallOpts, vault)
}

// GetRecognizedCollateralsLTVInfo is a free data retrieval call binding the contract method 0x6369fedb.
//
// Solidity: function getRecognizedCollateralsLTVInfo(address vault) view returns((address,uint256,uint256,uint256,uint256,uint256)[])
func (_EulerVaultLens *EulerVaultLensCallerSession) GetRecognizedCollateralsLTVInfo(vault common.Address) ([]LTVInfo, error) {
	return _EulerVaultLens.Contract.GetRecognizedCollateralsLTVInfo(&_EulerVaultLens.CallOpts, vault)
}

// GetRewardVaultInfo is a free data retrieval call binding the contract method 0xe74b9632.
//
// Solidity: function getRewardVaultInfo(address vault, address reward, uint256 numberOfEpochs) view returns((uint256,address,address,string,string,uint8,address,uint256,uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)[]))
func (_EulerVaultLens *EulerVaultLensCaller) GetRewardVaultInfo(opts *bind.CallOpts, vault common.Address, reward common.Address, numberOfEpochs *big.Int) (VaultRewardInfo, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "getRewardVaultInfo", vault, reward, numberOfEpochs)

	if err != nil {
		return *new(VaultRewardInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultRewardInfo)).(*VaultRewardInfo)

	return out0, err

}

// GetRewardVaultInfo is a free data retrieval call binding the contract method 0xe74b9632.
//
// Solidity: function getRewardVaultInfo(address vault, address reward, uint256 numberOfEpochs) view returns((uint256,address,address,string,string,uint8,address,uint256,uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)[]))
func (_EulerVaultLens *EulerVaultLensSession) GetRewardVaultInfo(vault common.Address, reward common.Address, numberOfEpochs *big.Int) (VaultRewardInfo, error) {
	return _EulerVaultLens.Contract.GetRewardVaultInfo(&_EulerVaultLens.CallOpts, vault, reward, numberOfEpochs)
}

// GetRewardVaultInfo is a free data retrieval call binding the contract method 0xe74b9632.
//
// Solidity: function getRewardVaultInfo(address vault, address reward, uint256 numberOfEpochs) view returns((uint256,address,address,string,string,uint8,address,uint256,uint256,uint256,uint256,uint256,(uint256,uint256,uint256,uint256)[],(uint256,uint256,uint256,uint256)[]))
func (_EulerVaultLens *EulerVaultLensCallerSession) GetRewardVaultInfo(vault common.Address, reward common.Address, numberOfEpochs *big.Int) (VaultRewardInfo, error) {
	return _EulerVaultLens.Contract.GetRewardVaultInfo(&_EulerVaultLens.CallOpts, vault, reward, numberOfEpochs)
}

// GetVaultInfoFull is a free data retrieval call binding the contract method 0x116e1d2e.
//
// Solidity: function getVaultInfoFull(address vault) view returns((uint256,address,string,string,uint256,address,string,string,uint256,address,string,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address,address,address,address,address,address,(bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)),(address,uint256,uint256,uint256,uint256,uint256)[],(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256),(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256)[],(address,string,bytes),(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256),(address,string,bytes)))
func (_EulerVaultLens *EulerVaultLensCaller) GetVaultInfoFull(opts *bind.CallOpts, vault common.Address) (VaultInfoFull, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "getVaultInfoFull", vault)

	if err != nil {
		return *new(VaultInfoFull), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultInfoFull)).(*VaultInfoFull)

	return out0, err

}

// GetVaultInfoFull is a free data retrieval call binding the contract method 0x116e1d2e.
//
// Solidity: function getVaultInfoFull(address vault) view returns((uint256,address,string,string,uint256,address,string,string,uint256,address,string,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address,address,address,address,address,address,(bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)),(address,uint256,uint256,uint256,uint256,uint256)[],(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256),(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256)[],(address,string,bytes),(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256),(address,string,bytes)))
func (_EulerVaultLens *EulerVaultLensSession) GetVaultInfoFull(vault common.Address) (VaultInfoFull, error) {
	return _EulerVaultLens.Contract.GetVaultInfoFull(&_EulerVaultLens.CallOpts, vault)
}

// GetVaultInfoFull is a free data retrieval call binding the contract method 0x116e1d2e.
//
// Solidity: function getVaultInfoFull(address vault) view returns((uint256,address,string,string,uint256,address,string,string,uint256,address,string,string,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address,address,address,address,address,address,address,address,address,(bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)),(address,uint256,uint256,uint256,uint256,uint256)[],(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256),(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256)[],(address,string,bytes),(bool,bytes,uint256,address,address,address,uint256,uint256,uint256,uint256),(address,string,bytes)))
func (_EulerVaultLens *EulerVaultLensCallerSession) GetVaultInfoFull(vault common.Address) (VaultInfoFull, error) {
	return _EulerVaultLens.Contract.GetVaultInfoFull(&_EulerVaultLens.CallOpts, vault)
}

// GetVaultInterestRateModelInfo is a free data retrieval call binding the contract method 0xd1dc6e3b.
//
// Solidity: function getVaultInterestRateModelInfo(address vault, uint256[] cash, uint256[] borrows) view returns((bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)))
func (_EulerVaultLens *EulerVaultLensCaller) GetVaultInterestRateModelInfo(opts *bind.CallOpts, vault common.Address, cash []*big.Int, borrows []*big.Int) (VaultInterestRateModelInfo, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "getVaultInterestRateModelInfo", vault, cash, borrows)

	if err != nil {
		return *new(VaultInterestRateModelInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultInterestRateModelInfo)).(*VaultInterestRateModelInfo)

	return out0, err

}

// GetVaultInterestRateModelInfo is a free data retrieval call binding the contract method 0xd1dc6e3b.
//
// Solidity: function getVaultInterestRateModelInfo(address vault, uint256[] cash, uint256[] borrows) view returns((bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)))
func (_EulerVaultLens *EulerVaultLensSession) GetVaultInterestRateModelInfo(vault common.Address, cash []*big.Int, borrows []*big.Int) (VaultInterestRateModelInfo, error) {
	return _EulerVaultLens.Contract.GetVaultInterestRateModelInfo(&_EulerVaultLens.CallOpts, vault, cash, borrows)
}

// GetVaultInterestRateModelInfo is a free data retrieval call binding the contract method 0xd1dc6e3b.
//
// Solidity: function getVaultInterestRateModelInfo(address vault, uint256[] cash, uint256[] borrows) view returns((bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)))
func (_EulerVaultLens *EulerVaultLensCallerSession) GetVaultInterestRateModelInfo(vault common.Address, cash []*big.Int, borrows []*big.Int) (VaultInterestRateModelInfo, error) {
	return _EulerVaultLens.Contract.GetVaultInterestRateModelInfo(&_EulerVaultLens.CallOpts, vault, cash, borrows)
}

// GetVaultKinkInterestRateModelInfo is a free data retrieval call binding the contract method 0xca4ef6d6.
//
// Solidity: function getVaultKinkInterestRateModelInfo(address vault) view returns((bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)))
func (_EulerVaultLens *EulerVaultLensCaller) GetVaultKinkInterestRateModelInfo(opts *bind.CallOpts, vault common.Address) (VaultInterestRateModelInfo, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "getVaultKinkInterestRateModelInfo", vault)

	if err != nil {
		return *new(VaultInterestRateModelInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultInterestRateModelInfo)).(*VaultInterestRateModelInfo)

	return out0, err

}

// GetVaultKinkInterestRateModelInfo is a free data retrieval call binding the contract method 0xca4ef6d6.
//
// Solidity: function getVaultKinkInterestRateModelInfo(address vault) view returns((bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)))
func (_EulerVaultLens *EulerVaultLensSession) GetVaultKinkInterestRateModelInfo(vault common.Address) (VaultInterestRateModelInfo, error) {
	return _EulerVaultLens.Contract.GetVaultKinkInterestRateModelInfo(&_EulerVaultLens.CallOpts, vault)
}

// GetVaultKinkInterestRateModelInfo is a free data retrieval call binding the contract method 0xca4ef6d6.
//
// Solidity: function getVaultKinkInterestRateModelInfo(address vault) view returns((bool,bytes,address,address,(uint256,uint256,uint256,uint256,uint256)[],(address,uint8,bytes)))
func (_EulerVaultLens *EulerVaultLensCallerSession) GetVaultKinkInterestRateModelInfo(vault common.Address) (VaultInterestRateModelInfo, error) {
	return _EulerVaultLens.Contract.GetVaultKinkInterestRateModelInfo(&_EulerVaultLens.CallOpts, vault)
}

// IrmLens is a free data retrieval call binding the contract method 0xdfe31a13.
//
// Solidity: function irmLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensCaller) IrmLens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "irmLens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IrmLens is a free data retrieval call binding the contract method 0xdfe31a13.
//
// Solidity: function irmLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensSession) IrmLens() (common.Address, error) {
	return _EulerVaultLens.Contract.IrmLens(&_EulerVaultLens.CallOpts)
}

// IrmLens is a free data retrieval call binding the contract method 0xdfe31a13.
//
// Solidity: function irmLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensCallerSession) IrmLens() (common.Address, error) {
	return _EulerVaultLens.Contract.IrmLens(&_EulerVaultLens.CallOpts)
}

// OracleLens is a free data retrieval call binding the contract method 0xc90be1e4.
//
// Solidity: function oracleLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensCaller) OracleLens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "oracleLens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleLens is a free data retrieval call binding the contract method 0xc90be1e4.
//
// Solidity: function oracleLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensSession) OracleLens() (common.Address, error) {
	return _EulerVaultLens.Contract.OracleLens(&_EulerVaultLens.CallOpts)
}

// OracleLens is a free data retrieval call binding the contract method 0xc90be1e4.
//
// Solidity: function oracleLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensCallerSession) OracleLens() (common.Address, error) {
	return _EulerVaultLens.Contract.OracleLens(&_EulerVaultLens.CallOpts)
}

// UtilsLens is a free data retrieval call binding the contract method 0x0f80efc3.
//
// Solidity: function utilsLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensCaller) UtilsLens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerVaultLens.contract.Call(opts, &out, "utilsLens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UtilsLens is a free data retrieval call binding the contract method 0x0f80efc3.
//
// Solidity: function utilsLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensSession) UtilsLens() (common.Address, error) {
	return _EulerVaultLens.Contract.UtilsLens(&_EulerVaultLens.CallOpts)
}

// UtilsLens is a free data retrieval call binding the contract method 0x0f80efc3.
//
// Solidity: function utilsLens() view returns(address)
func (_EulerVaultLens *EulerVaultLensCallerSession) UtilsLens() (common.Address, error) {
	return _EulerVaultLens.Contract.UtilsLens(&_EulerVaultLens.CallOpts)
}
