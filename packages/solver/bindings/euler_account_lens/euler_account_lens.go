// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_account_lens

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

// AccountInfo is an auto generated low-level Go binding around an user-defined struct.
type AccountInfo struct {
	EvcAccountInfo    EVCAccountInfo
	VaultAccountInfo  VaultAccountInfo
	AccountRewardInfo AccountRewardInfo
}

// AccountLiquidityInfo is an auto generated low-level Go binding around an user-defined struct.
type AccountLiquidityInfo struct {
	QueryFailure                       bool
	QueryFailureReason                 []byte
	TimeToLiquidation                  *big.Int
	LiabilityValue                     *big.Int
	CollateralValueBorrowing           *big.Int
	CollateralValueLiquidation         *big.Int
	CollateralValueRaw                 *big.Int
	CollateralLiquidityBorrowingInfo   []CollateralLiquidityInfo
	CollateralLiquidityLiquidationInfo []CollateralLiquidityInfo
	CollateralLiquidityRawInfo         []CollateralLiquidityInfo
}

// AccountMultipleVaultsInfo is an auto generated low-level Go binding around an user-defined struct.
type AccountMultipleVaultsInfo struct {
	EvcAccountInfo    EVCAccountInfo
	VaultAccountInfo  []VaultAccountInfo
	AccountRewardInfo []AccountRewardInfo
}

// AccountRewardInfo is an auto generated low-level Go binding around an user-defined struct.
type AccountRewardInfo struct {
	Timestamp               *big.Int
	Account                 common.Address
	Vault                   common.Address
	BalanceTracker          common.Address
	BalanceForwarderEnabled bool
	Balance                 *big.Int
	EnabledRewardsInfo      []EnabledRewardInfo
}

// CollateralLiquidityInfo is an auto generated low-level Go binding around an user-defined struct.
type CollateralLiquidityInfo struct {
	Collateral      common.Address
	CollateralValue *big.Int
}

// EVCAccountInfo is an auto generated low-level Go binding around an user-defined struct.
type EVCAccountInfo struct {
	Timestamp                       *big.Int
	Evc                             common.Address
	Account                         common.Address
	AddressPrefix                   [19]byte
	Owner                           common.Address
	IsLockdownMode                  bool
	IsPermitDisabledMode            bool
	LastAccountStatusCheckTimestamp *big.Int
	EnabledControllers              []common.Address
	EnabledCollaterals              []common.Address
}

// EnabledRewardInfo is an auto generated low-level Go binding around an user-defined struct.
type EnabledRewardInfo struct {
	Reward                    common.Address
	EarnedReward              *big.Int
	EarnedRewardRecentIgnored *big.Int
}

// VaultAccountInfo is an auto generated low-level Go binding around an user-defined struct.
type VaultAccountInfo struct {
	Timestamp                            *big.Int
	Account                              common.Address
	Vault                                common.Address
	Asset                                common.Address
	AssetsAccount                        *big.Int
	Shares                               *big.Int
	Assets                               *big.Int
	Borrowed                             *big.Int
	AssetAllowanceVault                  *big.Int
	AssetAllowanceVaultPermit2           *big.Int
	AssetAllowanceExpirationVaultPermit2 *big.Int
	AssetAllowancePermit2                *big.Int
	BalanceForwarderEnabled              bool
	IsController                         bool
	IsCollateral                         bool
	LiquidityInfo                        AccountLiquidityInfo
}

// EulerAccountLensMetaData contains all meta data concerning the EulerAccountLens contract.
var EulerAccountLensMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"TTL_ERROR\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_INFINITY\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_LIQUIDATION\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TTL_MORE_THAN_ONE_YEAR\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountEnabledVaultsInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLockdownMode\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPermitDisabledMode\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lastAccountStatusCheckTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"enabledControllers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"enabledCollaterals\",\"type\":\"address[]\"}],\"internalType\":\"structEVCAccountInfo\",\"name\":\"evcAccountInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetsAccount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceVault\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceVaultPermit2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceExpirationVaultPermit2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowancePermit2\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"balanceForwarderEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCollateral\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"int256\",\"name\":\"timeToLiquidation\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueBorrowing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueLiquidation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueRaw\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityBorrowingInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityLiquidationInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityRawInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structAccountLiquidityInfo\",\"name\":\"liquidityInfo\",\"type\":\"tuple\"}],\"internalType\":\"structVaultAccountInfo[]\",\"name\":\"vaultAccountInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"balanceForwarderEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"earnedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earnedRewardRecentIgnored\",\"type\":\"uint256\"}],\"internalType\":\"structEnabledRewardInfo[]\",\"name\":\"enabledRewardsInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structAccountRewardInfo[]\",\"name\":\"accountRewardInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structAccountMultipleVaultsInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getAccountInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLockdownMode\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPermitDisabledMode\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lastAccountStatusCheckTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"enabledControllers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"enabledCollaterals\",\"type\":\"address[]\"}],\"internalType\":\"structEVCAccountInfo\",\"name\":\"evcAccountInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetsAccount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceVault\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceVaultPermit2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceExpirationVaultPermit2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowancePermit2\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"balanceForwarderEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCollateral\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"int256\",\"name\":\"timeToLiquidation\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueBorrowing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueLiquidation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueRaw\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityBorrowingInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityLiquidationInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityRawInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structAccountLiquidityInfo\",\"name\":\"liquidityInfo\",\"type\":\"tuple\"}],\"internalType\":\"structVaultAccountInfo\",\"name\":\"vaultAccountInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"balanceForwarderEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"earnedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earnedRewardRecentIgnored\",\"type\":\"uint256\"}],\"internalType\":\"structEnabledRewardInfo[]\",\"name\":\"enabledRewardsInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structAccountRewardInfo\",\"name\":\"accountRewardInfo\",\"type\":\"tuple\"}],\"internalType\":\"structAccountInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getEVCAccountInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes19\",\"name\":\"addressPrefix\",\"type\":\"bytes19\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLockdownMode\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPermitDisabledMode\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"lastAccountStatusCheckTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"enabledControllers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"enabledCollaterals\",\"type\":\"address[]\"}],\"internalType\":\"structEVCAccountInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getRewardAccountInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"balanceForwarderEnabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"reward\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"earnedReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"earnedRewardRecentIgnored\",\"type\":\"uint256\"}],\"internalType\":\"structEnabledRewardInfo[]\",\"name\":\"enabledRewardsInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structAccountRewardInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getTimeToLiquidation\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getVaultAccountInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetsAccount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceVault\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceVaultPermit2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowanceExpirationVaultPermit2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetAllowancePermit2\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"balanceForwarderEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isController\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCollateral\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"queryFailure\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"queryFailureReason\",\"type\":\"bytes\"},{\"internalType\":\"int256\",\"name\":\"timeToLiquidation\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueBorrowing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueLiquidation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValueRaw\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityBorrowingInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityLiquidationInfo\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"internalType\":\"structCollateralLiquidityInfo[]\",\"name\":\"collateralLiquidityRawInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structAccountLiquidityInfo\",\"name\":\"liquidityInfo\",\"type\":\"tuple\"}],\"internalType\":\"structVaultAccountInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EulerAccountLensABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerAccountLensMetaData.ABI instead.
var EulerAccountLensABI = EulerAccountLensMetaData.ABI

// EulerAccountLens is an auto generated Go binding around an Ethereum contract.
type EulerAccountLens struct {
	EulerAccountLensCaller     // Read-only binding to the contract
	EulerAccountLensTransactor // Write-only binding to the contract
	EulerAccountLensFilterer   // Log filterer for contract events
}

// EulerAccountLensCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerAccountLensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerAccountLensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerAccountLensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerAccountLensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerAccountLensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerAccountLensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerAccountLensSession struct {
	Contract     *EulerAccountLens // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EulerAccountLensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerAccountLensCallerSession struct {
	Contract *EulerAccountLensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// EulerAccountLensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerAccountLensTransactorSession struct {
	Contract     *EulerAccountLensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// EulerAccountLensRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerAccountLensRaw struct {
	Contract *EulerAccountLens // Generic contract binding to access the raw methods on
}

// EulerAccountLensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerAccountLensCallerRaw struct {
	Contract *EulerAccountLensCaller // Generic read-only contract binding to access the raw methods on
}

// EulerAccountLensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerAccountLensTransactorRaw struct {
	Contract *EulerAccountLensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerAccountLens creates a new instance of EulerAccountLens, bound to a specific deployed contract.
func NewEulerAccountLens(address common.Address, backend bind.ContractBackend) (*EulerAccountLens, error) {
	contract, err := bindEulerAccountLens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerAccountLens{EulerAccountLensCaller: EulerAccountLensCaller{contract: contract}, EulerAccountLensTransactor: EulerAccountLensTransactor{contract: contract}, EulerAccountLensFilterer: EulerAccountLensFilterer{contract: contract}}, nil
}

// NewEulerAccountLensCaller creates a new read-only instance of EulerAccountLens, bound to a specific deployed contract.
func NewEulerAccountLensCaller(address common.Address, caller bind.ContractCaller) (*EulerAccountLensCaller, error) {
	contract, err := bindEulerAccountLens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerAccountLensCaller{contract: contract}, nil
}

// NewEulerAccountLensTransactor creates a new write-only instance of EulerAccountLens, bound to a specific deployed contract.
func NewEulerAccountLensTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerAccountLensTransactor, error) {
	contract, err := bindEulerAccountLens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerAccountLensTransactor{contract: contract}, nil
}

// NewEulerAccountLensFilterer creates a new log filterer instance of EulerAccountLens, bound to a specific deployed contract.
func NewEulerAccountLensFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerAccountLensFilterer, error) {
	contract, err := bindEulerAccountLens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerAccountLensFilterer{contract: contract}, nil
}

// bindEulerAccountLens binds a generic wrapper to an already deployed contract.
func bindEulerAccountLens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerAccountLensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerAccountLens *EulerAccountLensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerAccountLens.Contract.EulerAccountLensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerAccountLens *EulerAccountLensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerAccountLens.Contract.EulerAccountLensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerAccountLens *EulerAccountLensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerAccountLens.Contract.EulerAccountLensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerAccountLens *EulerAccountLensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerAccountLens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerAccountLens *EulerAccountLensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerAccountLens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerAccountLens *EulerAccountLensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerAccountLens.Contract.contract.Transact(opts, method, params...)
}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerAccountLens *EulerAccountLensCaller) TTLERROR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "TTL_ERROR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerAccountLens *EulerAccountLensSession) TTLERROR() (*big.Int, error) {
	return _EulerAccountLens.Contract.TTLERROR(&_EulerAccountLens.CallOpts)
}

// TTLERROR is a free data retrieval call binding the contract method 0x6410b792.
//
// Solidity: function TTL_ERROR() view returns(int256)
func (_EulerAccountLens *EulerAccountLensCallerSession) TTLERROR() (*big.Int, error) {
	return _EulerAccountLens.Contract.TTLERROR(&_EulerAccountLens.CallOpts)
}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerAccountLens *EulerAccountLensCaller) TTLINFINITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "TTL_INFINITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerAccountLens *EulerAccountLensSession) TTLINFINITY() (*big.Int, error) {
	return _EulerAccountLens.Contract.TTLINFINITY(&_EulerAccountLens.CallOpts)
}

// TTLINFINITY is a free data retrieval call binding the contract method 0x900bb8a6.
//
// Solidity: function TTL_INFINITY() view returns(int256)
func (_EulerAccountLens *EulerAccountLensCallerSession) TTLINFINITY() (*big.Int, error) {
	return _EulerAccountLens.Contract.TTLINFINITY(&_EulerAccountLens.CallOpts)
}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerAccountLens *EulerAccountLensCaller) TTLLIQUIDATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "TTL_LIQUIDATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerAccountLens *EulerAccountLensSession) TTLLIQUIDATION() (*big.Int, error) {
	return _EulerAccountLens.Contract.TTLLIQUIDATION(&_EulerAccountLens.CallOpts)
}

// TTLLIQUIDATION is a free data retrieval call binding the contract method 0x72537d9a.
//
// Solidity: function TTL_LIQUIDATION() view returns(int256)
func (_EulerAccountLens *EulerAccountLensCallerSession) TTLLIQUIDATION() (*big.Int, error) {
	return _EulerAccountLens.Contract.TTLLIQUIDATION(&_EulerAccountLens.CallOpts)
}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerAccountLens *EulerAccountLensCaller) TTLMORETHANONEYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "TTL_MORE_THAN_ONE_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerAccountLens *EulerAccountLensSession) TTLMORETHANONEYEAR() (*big.Int, error) {
	return _EulerAccountLens.Contract.TTLMORETHANONEYEAR(&_EulerAccountLens.CallOpts)
}

// TTLMORETHANONEYEAR is a free data retrieval call binding the contract method 0x4abee02a.
//
// Solidity: function TTL_MORE_THAN_ONE_YEAR() view returns(int256)
func (_EulerAccountLens *EulerAccountLensCallerSession) TTLMORETHANONEYEAR() (*big.Int, error) {
	return _EulerAccountLens.Contract.TTLMORETHANONEYEAR(&_EulerAccountLens.CallOpts)
}

// GetAccountEnabledVaultsInfo is a free data retrieval call binding the contract method 0xa90b248a.
//
// Solidity: function getAccountEnabledVaultsInfo(address evc, address account) view returns(((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]),(uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[]))[],(uint256,address,address,address,bool,uint256,(address,uint256,uint256)[])[]))
func (_EulerAccountLens *EulerAccountLensCaller) GetAccountEnabledVaultsInfo(opts *bind.CallOpts, evc common.Address, account common.Address) (AccountMultipleVaultsInfo, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "getAccountEnabledVaultsInfo", evc, account)

	if err != nil {
		return *new(AccountMultipleVaultsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountMultipleVaultsInfo)).(*AccountMultipleVaultsInfo)

	return out0, err

}

// GetAccountEnabledVaultsInfo is a free data retrieval call binding the contract method 0xa90b248a.
//
// Solidity: function getAccountEnabledVaultsInfo(address evc, address account) view returns(((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]),(uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[]))[],(uint256,address,address,address,bool,uint256,(address,uint256,uint256)[])[]))
func (_EulerAccountLens *EulerAccountLensSession) GetAccountEnabledVaultsInfo(evc common.Address, account common.Address) (AccountMultipleVaultsInfo, error) {
	return _EulerAccountLens.Contract.GetAccountEnabledVaultsInfo(&_EulerAccountLens.CallOpts, evc, account)
}

// GetAccountEnabledVaultsInfo is a free data retrieval call binding the contract method 0xa90b248a.
//
// Solidity: function getAccountEnabledVaultsInfo(address evc, address account) view returns(((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]),(uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[]))[],(uint256,address,address,address,bool,uint256,(address,uint256,uint256)[])[]))
func (_EulerAccountLens *EulerAccountLensCallerSession) GetAccountEnabledVaultsInfo(evc common.Address, account common.Address) (AccountMultipleVaultsInfo, error) {
	return _EulerAccountLens.Contract.GetAccountEnabledVaultsInfo(&_EulerAccountLens.CallOpts, evc, account)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x6332fef6.
//
// Solidity: function getAccountInfo(address account, address vault) view returns(((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]),(uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[])),(uint256,address,address,address,bool,uint256,(address,uint256,uint256)[])))
func (_EulerAccountLens *EulerAccountLensCaller) GetAccountInfo(opts *bind.CallOpts, account common.Address, vault common.Address) (AccountInfo, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "getAccountInfo", account, vault)

	if err != nil {
		return *new(AccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountInfo)).(*AccountInfo)

	return out0, err

}

// GetAccountInfo is a free data retrieval call binding the contract method 0x6332fef6.
//
// Solidity: function getAccountInfo(address account, address vault) view returns(((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]),(uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[])),(uint256,address,address,address,bool,uint256,(address,uint256,uint256)[])))
func (_EulerAccountLens *EulerAccountLensSession) GetAccountInfo(account common.Address, vault common.Address) (AccountInfo, error) {
	return _EulerAccountLens.Contract.GetAccountInfo(&_EulerAccountLens.CallOpts, account, vault)
}

// GetAccountInfo is a free data retrieval call binding the contract method 0x6332fef6.
//
// Solidity: function getAccountInfo(address account, address vault) view returns(((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]),(uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[])),(uint256,address,address,address,bool,uint256,(address,uint256,uint256)[])))
func (_EulerAccountLens *EulerAccountLensCallerSession) GetAccountInfo(account common.Address, vault common.Address) (AccountInfo, error) {
	return _EulerAccountLens.Contract.GetAccountInfo(&_EulerAccountLens.CallOpts, account, vault)
}

// GetEVCAccountInfo is a free data retrieval call binding the contract method 0xf859d53c.
//
// Solidity: function getEVCAccountInfo(address evc, address account) view returns((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]))
func (_EulerAccountLens *EulerAccountLensCaller) GetEVCAccountInfo(opts *bind.CallOpts, evc common.Address, account common.Address) (EVCAccountInfo, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "getEVCAccountInfo", evc, account)

	if err != nil {
		return *new(EVCAccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(EVCAccountInfo)).(*EVCAccountInfo)

	return out0, err

}

// GetEVCAccountInfo is a free data retrieval call binding the contract method 0xf859d53c.
//
// Solidity: function getEVCAccountInfo(address evc, address account) view returns((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]))
func (_EulerAccountLens *EulerAccountLensSession) GetEVCAccountInfo(evc common.Address, account common.Address) (EVCAccountInfo, error) {
	return _EulerAccountLens.Contract.GetEVCAccountInfo(&_EulerAccountLens.CallOpts, evc, account)
}

// GetEVCAccountInfo is a free data retrieval call binding the contract method 0xf859d53c.
//
// Solidity: function getEVCAccountInfo(address evc, address account) view returns((uint256,address,address,bytes19,address,bool,bool,uint256,address[],address[]))
func (_EulerAccountLens *EulerAccountLensCallerSession) GetEVCAccountInfo(evc common.Address, account common.Address) (EVCAccountInfo, error) {
	return _EulerAccountLens.Contract.GetEVCAccountInfo(&_EulerAccountLens.CallOpts, evc, account)
}

// GetRewardAccountInfo is a free data retrieval call binding the contract method 0xe5e2f2da.
//
// Solidity: function getRewardAccountInfo(address account, address vault) view returns((uint256,address,address,address,bool,uint256,(address,uint256,uint256)[]))
func (_EulerAccountLens *EulerAccountLensCaller) GetRewardAccountInfo(opts *bind.CallOpts, account common.Address, vault common.Address) (AccountRewardInfo, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "getRewardAccountInfo", account, vault)

	if err != nil {
		return *new(AccountRewardInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AccountRewardInfo)).(*AccountRewardInfo)

	return out0, err

}

// GetRewardAccountInfo is a free data retrieval call binding the contract method 0xe5e2f2da.
//
// Solidity: function getRewardAccountInfo(address account, address vault) view returns((uint256,address,address,address,bool,uint256,(address,uint256,uint256)[]))
func (_EulerAccountLens *EulerAccountLensSession) GetRewardAccountInfo(account common.Address, vault common.Address) (AccountRewardInfo, error) {
	return _EulerAccountLens.Contract.GetRewardAccountInfo(&_EulerAccountLens.CallOpts, account, vault)
}

// GetRewardAccountInfo is a free data retrieval call binding the contract method 0xe5e2f2da.
//
// Solidity: function getRewardAccountInfo(address account, address vault) view returns((uint256,address,address,address,bool,uint256,(address,uint256,uint256)[]))
func (_EulerAccountLens *EulerAccountLensCallerSession) GetRewardAccountInfo(account common.Address, vault common.Address) (AccountRewardInfo, error) {
	return _EulerAccountLens.Contract.GetRewardAccountInfo(&_EulerAccountLens.CallOpts, account, vault)
}

// GetTimeToLiquidation is a free data retrieval call binding the contract method 0x03c9ab53.
//
// Solidity: function getTimeToLiquidation(address account, address vault) view returns(int256)
func (_EulerAccountLens *EulerAccountLensCaller) GetTimeToLiquidation(opts *bind.CallOpts, account common.Address, vault common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "getTimeToLiquidation", account, vault)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimeToLiquidation is a free data retrieval call binding the contract method 0x03c9ab53.
//
// Solidity: function getTimeToLiquidation(address account, address vault) view returns(int256)
func (_EulerAccountLens *EulerAccountLensSession) GetTimeToLiquidation(account common.Address, vault common.Address) (*big.Int, error) {
	return _EulerAccountLens.Contract.GetTimeToLiquidation(&_EulerAccountLens.CallOpts, account, vault)
}

// GetTimeToLiquidation is a free data retrieval call binding the contract method 0x03c9ab53.
//
// Solidity: function getTimeToLiquidation(address account, address vault) view returns(int256)
func (_EulerAccountLens *EulerAccountLensCallerSession) GetTimeToLiquidation(account common.Address, vault common.Address) (*big.Int, error) {
	return _EulerAccountLens.Contract.GetTimeToLiquidation(&_EulerAccountLens.CallOpts, account, vault)
}

// GetVaultAccountInfo is a free data retrieval call binding the contract method 0x53e44f47.
//
// Solidity: function getVaultAccountInfo(address account, address vault) view returns((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[])))
func (_EulerAccountLens *EulerAccountLensCaller) GetVaultAccountInfo(opts *bind.CallOpts, account common.Address, vault common.Address) (VaultAccountInfo, error) {
	var out []interface{}
	err := _EulerAccountLens.contract.Call(opts, &out, "getVaultAccountInfo", account, vault)

	if err != nil {
		return *new(VaultAccountInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultAccountInfo)).(*VaultAccountInfo)

	return out0, err

}

// GetVaultAccountInfo is a free data retrieval call binding the contract method 0x53e44f47.
//
// Solidity: function getVaultAccountInfo(address account, address vault) view returns((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[])))
func (_EulerAccountLens *EulerAccountLensSession) GetVaultAccountInfo(account common.Address, vault common.Address) (VaultAccountInfo, error) {
	return _EulerAccountLens.Contract.GetVaultAccountInfo(&_EulerAccountLens.CallOpts, account, vault)
}

// GetVaultAccountInfo is a free data retrieval call binding the contract method 0x53e44f47.
//
// Solidity: function getVaultAccountInfo(address account, address vault) view returns((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,(bool,bytes,int256,uint256,uint256,uint256,uint256,(address,uint256)[],(address,uint256)[],(address,uint256)[])))
func (_EulerAccountLens *EulerAccountLensCallerSession) GetVaultAccountInfo(account common.Address, vault common.Address) (VaultAccountInfo, error) {
	return _EulerAccountLens.Contract.GetVaultAccountInfo(&_EulerAccountLens.CallOpts, account, vault)
}
