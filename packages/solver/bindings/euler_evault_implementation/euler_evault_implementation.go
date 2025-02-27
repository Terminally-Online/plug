// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_evault_implementation

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

// BaseIntegrations is an auto generated low-level Go binding around an user-defined struct.
type BaseIntegrations struct {
	Evc              common.Address
	ProtocolConfig   common.Address
	SequenceRegistry common.Address
	BalanceTracker   common.Address
	Permit2          common.Address
}

// DispatchDeployedModules is an auto generated low-level Go binding around an user-defined struct.
type DispatchDeployedModules struct {
	Initialize       common.Address
	Token            common.Address
	Vault            common.Address
	Borrowing        common.Address
	Liquidation      common.Address
	RiskManager      common.Address
	BalanceForwarder common.Address
	Governance       common.Address
}

// EulerEvaultImplementationMetaData contains all meta data concerning the EulerEvaultImplementation contract.
var EulerEvaultImplementationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"protocolConfig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequenceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"}],\"internalType\":\"structBase.Integrations\",\"name\":\"integrations\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"initialize\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowing\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"riskManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"internalType\":\"structDispatch.DeployedModules\",\"name\":\"modules\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"E_AccountLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_AmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadAssetReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadBorrowCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadMaxLiquidationDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSharesOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSharesReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSupplyCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BorrowCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_CheckUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_CollateralDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ConfigAmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ControllerDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_DebtAmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_EmptyError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ExcessiveRepayAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_FlashLoanNotRepaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientCash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InvalidLTVAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LTVBorrow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LTVLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LiquidationCoolOff\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_MinYield\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NoLiability\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NoPriceOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotHookTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_OperationDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_OutstandingDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ProxyMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_RepayTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SelfLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_TransientState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ViolatorLiquidityDeferred\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ZeroAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ZeroShares\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"BalanceForwarderStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolReceiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governorReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"governorShares\",\"type\":\"uint256\"}],\"name\":\"ConvertFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"DebtSocialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dToken\",\"type\":\"address\"}],\"name\":\"EVaultCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newSupplyCap\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newBorrowCap\",\"type\":\"uint16\"}],\"name\":\"GovSetCaps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newConfigFlags\",\"type\":\"uint32\"}],\"name\":\"GovSetConfigFlags\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeReceiver\",\"type\":\"address\"}],\"name\":\"GovSetFeeReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernorAdmin\",\"type\":\"address\"}],\"name\":\"GovSetGovernorAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHookTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newHookedOps\",\"type\":\"uint32\"}],\"name\":\"GovSetHookConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"GovSetInterestFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"GovSetInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"targetTimestamp\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"name\":\"GovSetLTV\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newCoolOffTime\",\"type\":\"uint16\"}],\"name\":\"GovSetLiquidationCoolOffTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newDiscount\",\"type\":\"uint16\"}],\"name\":\"GovSetMaxLiquidationDiscount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"InterestAccrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldBalance\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"PullDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accumulatedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"VaultStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EVC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVBorrow\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVFull\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"targetTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVLiquidation\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LTVList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_BALANCE_FORWARDER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_BORROWING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_GOVERNANCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_INITIALIZE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_LIQUIDATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_RISKMANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"liquidation\",\"type\":\"bool\"}],\"name\":\"accountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"liquidation\",\"type\":\"bool\"}],\"name\":\"accountLiquidityFull\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"collateralValues\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFeesAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceForwarderEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceTrackerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"caps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"supplyCap\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"borrowCap\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"}],\"name\":\"checkAccountStatus\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"checkLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRepay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxYield\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkVaultStatus\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configFlags\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convertFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOfExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governorAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hookConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyCreator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestAccumulator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minYieldBalance\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationCoolOffTime\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidationDiscount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolConfigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"pullDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"repayWithShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"supplyCap\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"borrowCap\",\"type\":\"uint16\"}],\"name\":\"setCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newConfigFlags\",\"type\":\"uint32\"}],\"name\":\"setConfigFlags\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernorAdmin\",\"type\":\"address\"}],\"name\":\"setGovernorAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHookTarget\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newHookedOps\",\"type\":\"uint32\"}],\"name\":\"setHookConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"setInterestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"setInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"name\":\"setLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newCoolOffTime\",\"type\":\"uint16\"}],\"name\":\"setLiquidationCoolOffTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newDiscount\",\"type\":\"uint16\"}],\"name\":\"setMaxLiquidationDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"touch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferFromMax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unitOfAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewDelegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EulerEvaultImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerEvaultImplementationMetaData.ABI instead.
var EulerEvaultImplementationABI = EulerEvaultImplementationMetaData.ABI

// EulerEvaultImplementation is an auto generated Go binding around an Ethereum contract.
type EulerEvaultImplementation struct {
	EulerEvaultImplementationCaller     // Read-only binding to the contract
	EulerEvaultImplementationTransactor // Write-only binding to the contract
	EulerEvaultImplementationFilterer   // Log filterer for contract events
}

// EulerEvaultImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerEvaultImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvaultImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerEvaultImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvaultImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerEvaultImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvaultImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerEvaultImplementationSession struct {
	Contract     *EulerEvaultImplementation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EulerEvaultImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerEvaultImplementationCallerSession struct {
	Contract *EulerEvaultImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// EulerEvaultImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerEvaultImplementationTransactorSession struct {
	Contract     *EulerEvaultImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// EulerEvaultImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerEvaultImplementationRaw struct {
	Contract *EulerEvaultImplementation // Generic contract binding to access the raw methods on
}

// EulerEvaultImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerEvaultImplementationCallerRaw struct {
	Contract *EulerEvaultImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// EulerEvaultImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerEvaultImplementationTransactorRaw struct {
	Contract *EulerEvaultImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerEvaultImplementation creates a new instance of EulerEvaultImplementation, bound to a specific deployed contract.
func NewEulerEvaultImplementation(address common.Address, backend bind.ContractBackend) (*EulerEvaultImplementation, error) {
	contract, err := bindEulerEvaultImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementation{EulerEvaultImplementationCaller: EulerEvaultImplementationCaller{contract: contract}, EulerEvaultImplementationTransactor: EulerEvaultImplementationTransactor{contract: contract}, EulerEvaultImplementationFilterer: EulerEvaultImplementationFilterer{contract: contract}}, nil
}

// NewEulerEvaultImplementationCaller creates a new read-only instance of EulerEvaultImplementation, bound to a specific deployed contract.
func NewEulerEvaultImplementationCaller(address common.Address, caller bind.ContractCaller) (*EulerEvaultImplementationCaller, error) {
	contract, err := bindEulerEvaultImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationCaller{contract: contract}, nil
}

// NewEulerEvaultImplementationTransactor creates a new write-only instance of EulerEvaultImplementation, bound to a specific deployed contract.
func NewEulerEvaultImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerEvaultImplementationTransactor, error) {
	contract, err := bindEulerEvaultImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationTransactor{contract: contract}, nil
}

// NewEulerEvaultImplementationFilterer creates a new log filterer instance of EulerEvaultImplementation, bound to a specific deployed contract.
func NewEulerEvaultImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerEvaultImplementationFilterer, error) {
	contract, err := bindEulerEvaultImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationFilterer{contract: contract}, nil
}

// bindEulerEvaultImplementation binds a generic wrapper to an already deployed contract.
func bindEulerEvaultImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEvaultImplementation *EulerEvaultImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEvaultImplementation.Contract.EulerEvaultImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEvaultImplementation *EulerEvaultImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.EulerEvaultImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEvaultImplementation *EulerEvaultImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.EulerEvaultImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEvaultImplementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.contract.Transact(opts, method, params...)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) EVC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "EVC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) EVC() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.EVC(&_EulerEvaultImplementation.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) EVC() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.EVC(&_EulerEvaultImplementation.CallOpts)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) LTVBorrow(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "LTVBorrow", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EulerEvaultImplementation.Contract.LTVBorrow(&_EulerEvaultImplementation.CallOpts, collateral)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EulerEvaultImplementation.Contract.LTVBorrow(&_EulerEvaultImplementation.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) LTVFull(opts *bind.CallOpts, collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "LTVFull", collateral)

	outstruct := new(struct {
		BorrowLTV             uint16
		LiquidationLTV        uint16
		InitialLiquidationLTV uint16
		TargetTimestamp       *big.Int
		RampDuration          uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowLTV = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.LiquidationLTV = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.InitialLiquidationLTV = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.TargetTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RampDuration = *abi.ConvertType(out[4], new(uint32)).(*uint32)

	return *outstruct, err

}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EulerEvaultImplementation.Contract.LTVFull(&_EulerEvaultImplementation.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EulerEvaultImplementation.Contract.LTVFull(&_EulerEvaultImplementation.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) LTVLiquidation(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "LTVLiquidation", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EulerEvaultImplementation.Contract.LTVLiquidation(&_EulerEvaultImplementation.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EulerEvaultImplementation.Contract.LTVLiquidation(&_EulerEvaultImplementation.CallOpts, collateral)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) LTVList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "LTVList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) LTVList() ([]common.Address, error) {
	return _EulerEvaultImplementation.Contract.LTVList(&_EulerEvaultImplementation.CallOpts)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) LTVList() ([]common.Address, error) {
	return _EulerEvaultImplementation.Contract.LTVList(&_EulerEvaultImplementation.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MODULEBALANCEFORWARDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "MODULE_BALANCE_FORWARDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEBALANCEFORWARDER(&_EulerEvaultImplementation.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEBALANCEFORWARDER(&_EulerEvaultImplementation.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MODULEBORROWING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "MODULE_BORROWING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MODULEBORROWING() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEBORROWING(&_EulerEvaultImplementation.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MODULEBORROWING() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEBORROWING(&_EulerEvaultImplementation.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MODULEGOVERNANCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "MODULE_GOVERNANCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEGOVERNANCE(&_EulerEvaultImplementation.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEGOVERNANCE(&_EulerEvaultImplementation.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MODULEINITIALIZE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "MODULE_INITIALIZE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MODULEINITIALIZE() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEINITIALIZE(&_EulerEvaultImplementation.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MODULEINITIALIZE() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEINITIALIZE(&_EulerEvaultImplementation.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MODULELIQUIDATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "MODULE_LIQUIDATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MODULELIQUIDATION() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULELIQUIDATION(&_EulerEvaultImplementation.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MODULELIQUIDATION() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULELIQUIDATION(&_EulerEvaultImplementation.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MODULERISKMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "MODULE_RISKMANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MODULERISKMANAGER() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULERISKMANAGER(&_EulerEvaultImplementation.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MODULERISKMANAGER() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULERISKMANAGER(&_EulerEvaultImplementation.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MODULETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "MODULE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MODULETOKEN() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULETOKEN(&_EulerEvaultImplementation.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MODULETOKEN() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULETOKEN(&_EulerEvaultImplementation.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MODULEVAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "MODULE_VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MODULEVAULT() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEVAULT(&_EulerEvaultImplementation.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MODULEVAULT() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.MODULEVAULT(&_EulerEvaultImplementation.CallOpts)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) AccountLiquidity(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "accountLiquidity", account, liquidation)

	outstruct := new(struct {
		CollateralValue *big.Int
		LiabilityValue  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiabilityValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EulerEvaultImplementation.Contract.AccountLiquidity(&_EulerEvaultImplementation.CallOpts, account, liquidation)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EulerEvaultImplementation.Contract.AccountLiquidity(&_EulerEvaultImplementation.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) AccountLiquidityFull(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "accountLiquidityFull", account, liquidation)

	outstruct := new(struct {
		Collaterals      []common.Address
		CollateralValues []*big.Int
		LiabilityValue   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Collaterals = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CollateralValues = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.LiabilityValue = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EulerEvaultImplementation.Contract.AccountLiquidityFull(&_EulerEvaultImplementation.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EulerEvaultImplementation.Contract.AccountLiquidityFull(&_EulerEvaultImplementation.CallOpts, account, liquidation)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) AccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "accumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) AccumulatedFees() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.AccumulatedFees(&_EulerEvaultImplementation.CallOpts)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) AccumulatedFees() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.AccumulatedFees(&_EulerEvaultImplementation.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) AccumulatedFeesAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "accumulatedFeesAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.AccumulatedFeesAssets(&_EulerEvaultImplementation.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.AccumulatedFeesAssets(&_EulerEvaultImplementation.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.Allowance(&_EulerEvaultImplementation.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.Allowance(&_EulerEvaultImplementation.CallOpts, holder, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Asset() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.Asset(&_EulerEvaultImplementation.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Asset() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.Asset(&_EulerEvaultImplementation.CallOpts)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) BalanceForwarderEnabled(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "balanceForwarderEnabled", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EulerEvaultImplementation.Contract.BalanceForwarderEnabled(&_EulerEvaultImplementation.CallOpts, account)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EulerEvaultImplementation.Contract.BalanceForwarderEnabled(&_EulerEvaultImplementation.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.BalanceOf(&_EulerEvaultImplementation.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.BalanceOf(&_EulerEvaultImplementation.CallOpts, account)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) BalanceTrackerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "balanceTrackerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.BalanceTrackerAddress(&_EulerEvaultImplementation.CallOpts)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.BalanceTrackerAddress(&_EulerEvaultImplementation.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Caps(opts *bind.CallOpts) (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "caps")

	outstruct := new(struct {
		SupplyCap uint16
		BorrowCap uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SupplyCap = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.BorrowCap = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EulerEvaultImplementation.Contract.Caps(&_EulerEvaultImplementation.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EulerEvaultImplementation.Contract.Caps(&_EulerEvaultImplementation.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Cash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Cash() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.Cash(&_EulerEvaultImplementation.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Cash() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.Cash(&_EulerEvaultImplementation.CallOpts)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) CheckAccountStatus(opts *bind.CallOpts, account common.Address, collaterals []common.Address) ([4]byte, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "checkAccountStatus", account, collaterals)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EulerEvaultImplementation.Contract.CheckAccountStatus(&_EulerEvaultImplementation.CallOpts, account, collaterals)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EulerEvaultImplementation.Contract.CheckAccountStatus(&_EulerEvaultImplementation.CallOpts, account, collaterals)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) CheckLiquidation(opts *bind.CallOpts, liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "checkLiquidation", liquidator, violator, collateral)

	outstruct := new(struct {
		MaxRepay *big.Int
		MaxYield *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxRepay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxYield = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EulerEvaultImplementation.Contract.CheckLiquidation(&_EulerEvaultImplementation.CallOpts, liquidator, violator, collateral)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EulerEvaultImplementation.Contract.CheckLiquidation(&_EulerEvaultImplementation.CallOpts, liquidator, violator, collateral)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) ConfigFlags(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "configFlags")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) ConfigFlags() (uint32, error) {
	return _EulerEvaultImplementation.Contract.ConfigFlags(&_EulerEvaultImplementation.CallOpts)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) ConfigFlags() (uint32, error) {
	return _EulerEvaultImplementation.Contract.ConfigFlags(&_EulerEvaultImplementation.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.ConvertToAssets(&_EulerEvaultImplementation.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.ConvertToAssets(&_EulerEvaultImplementation.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.ConvertToShares(&_EulerEvaultImplementation.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.ConvertToShares(&_EulerEvaultImplementation.CallOpts, assets)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Creator() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.Creator(&_EulerEvaultImplementation.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Creator() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.Creator(&_EulerEvaultImplementation.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) DToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "dToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) DToken() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.DToken(&_EulerEvaultImplementation.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) DToken() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.DToken(&_EulerEvaultImplementation.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) DebtOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "debtOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.DebtOf(&_EulerEvaultImplementation.CallOpts, account)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.DebtOf(&_EulerEvaultImplementation.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) DebtOfExact(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "debtOfExact", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.DebtOfExact(&_EulerEvaultImplementation.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.DebtOfExact(&_EulerEvaultImplementation.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Decimals() (uint8, error) {
	return _EulerEvaultImplementation.Contract.Decimals(&_EulerEvaultImplementation.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Decimals() (uint8, error) {
	return _EulerEvaultImplementation.Contract.Decimals(&_EulerEvaultImplementation.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) FeeReceiver() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.FeeReceiver(&_EulerEvaultImplementation.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) FeeReceiver() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.FeeReceiver(&_EulerEvaultImplementation.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) GovernorAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "governorAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) GovernorAdmin() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.GovernorAdmin(&_EulerEvaultImplementation.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) GovernorAdmin() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.GovernorAdmin(&_EulerEvaultImplementation.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) HookConfig(opts *bind.CallOpts) (common.Address, uint32, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "hookConfig")

	if err != nil {
		return *new(common.Address), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) HookConfig() (common.Address, uint32, error) {
	return _EulerEvaultImplementation.Contract.HookConfig(&_EulerEvaultImplementation.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) HookConfig() (common.Address, uint32, error) {
	return _EulerEvaultImplementation.Contract.HookConfig(&_EulerEvaultImplementation.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) InterestAccumulator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "interestAccumulator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) InterestAccumulator() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.InterestAccumulator(&_EulerEvaultImplementation.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) InterestAccumulator() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.InterestAccumulator(&_EulerEvaultImplementation.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) InterestFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "interestFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) InterestFee() (uint16, error) {
	return _EulerEvaultImplementation.Contract.InterestFee(&_EulerEvaultImplementation.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) InterestFee() (uint16, error) {
	return _EulerEvaultImplementation.Contract.InterestFee(&_EulerEvaultImplementation.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) InterestRate() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.InterestRate(&_EulerEvaultImplementation.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) InterestRate() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.InterestRate(&_EulerEvaultImplementation.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) InterestRateModel() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.InterestRateModel(&_EulerEvaultImplementation.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) InterestRateModel() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.InterestRateModel(&_EulerEvaultImplementation.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) LiquidationCoolOffTime(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "liquidationCoolOffTime")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) LiquidationCoolOffTime() (uint16, error) {
	return _EulerEvaultImplementation.Contract.LiquidationCoolOffTime(&_EulerEvaultImplementation.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) LiquidationCoolOffTime() (uint16, error) {
	return _EulerEvaultImplementation.Contract.LiquidationCoolOffTime(&_EulerEvaultImplementation.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MaxDeposit(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "maxDeposit", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.MaxDeposit(&_EulerEvaultImplementation.CallOpts, account)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.MaxDeposit(&_EulerEvaultImplementation.CallOpts, account)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MaxLiquidationDiscount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "maxLiquidationDiscount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MaxLiquidationDiscount() (uint16, error) {
	return _EulerEvaultImplementation.Contract.MaxLiquidationDiscount(&_EulerEvaultImplementation.CallOpts)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MaxLiquidationDiscount() (uint16, error) {
	return _EulerEvaultImplementation.Contract.MaxLiquidationDiscount(&_EulerEvaultImplementation.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MaxMint(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "maxMint", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.MaxMint(&_EulerEvaultImplementation.CallOpts, account)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.MaxMint(&_EulerEvaultImplementation.CallOpts, account)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.MaxRedeem(&_EulerEvaultImplementation.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.MaxRedeem(&_EulerEvaultImplementation.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.MaxWithdraw(&_EulerEvaultImplementation.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.MaxWithdraw(&_EulerEvaultImplementation.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Name() (string, error) {
	return _EulerEvaultImplementation.Contract.Name(&_EulerEvaultImplementation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Name() (string, error) {
	return _EulerEvaultImplementation.Contract.Name(&_EulerEvaultImplementation.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Oracle() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.Oracle(&_EulerEvaultImplementation.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Oracle() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.Oracle(&_EulerEvaultImplementation.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Permit2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "permit2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Permit2Address() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.Permit2Address(&_EulerEvaultImplementation.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Permit2Address() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.Permit2Address(&_EulerEvaultImplementation.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.PreviewDeposit(&_EulerEvaultImplementation.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.PreviewDeposit(&_EulerEvaultImplementation.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.PreviewMint(&_EulerEvaultImplementation.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.PreviewMint(&_EulerEvaultImplementation.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.PreviewRedeem(&_EulerEvaultImplementation.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.PreviewRedeem(&_EulerEvaultImplementation.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.PreviewWithdraw(&_EulerEvaultImplementation.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.PreviewWithdraw(&_EulerEvaultImplementation.CallOpts, assets)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) ProtocolConfigAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "protocolConfigAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) ProtocolConfigAddress() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.ProtocolConfigAddress(&_EulerEvaultImplementation.CallOpts)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) ProtocolConfigAddress() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.ProtocolConfigAddress(&_EulerEvaultImplementation.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) ProtocolFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "protocolFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.ProtocolFeeReceiver(&_EulerEvaultImplementation.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.ProtocolFeeReceiver(&_EulerEvaultImplementation.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) ProtocolFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "protocolFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) ProtocolFeeShare() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.ProtocolFeeShare(&_EulerEvaultImplementation.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) ProtocolFeeShare() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.ProtocolFeeShare(&_EulerEvaultImplementation.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Symbol() (string, error) {
	return _EulerEvaultImplementation.Contract.Symbol(&_EulerEvaultImplementation.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) Symbol() (string, error) {
	return _EulerEvaultImplementation.Contract.Symbol(&_EulerEvaultImplementation.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) TotalAssets() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.TotalAssets(&_EulerEvaultImplementation.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) TotalAssets() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.TotalAssets(&_EulerEvaultImplementation.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) TotalBorrows() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.TotalBorrows(&_EulerEvaultImplementation.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) TotalBorrows() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.TotalBorrows(&_EulerEvaultImplementation.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) TotalBorrowsExact(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "totalBorrowsExact")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) TotalBorrowsExact() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.TotalBorrowsExact(&_EulerEvaultImplementation.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) TotalBorrowsExact() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.TotalBorrowsExact(&_EulerEvaultImplementation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) TotalSupply() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.TotalSupply(&_EulerEvaultImplementation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) TotalSupply() (*big.Int, error) {
	return _EulerEvaultImplementation.Contract.TotalSupply(&_EulerEvaultImplementation.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCaller) UnitOfAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvaultImplementation.contract.Call(opts, &out, "unitOfAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) UnitOfAccount() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.UnitOfAccount(&_EulerEvaultImplementation.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEvaultImplementation *EulerEvaultImplementationCallerSession) UnitOfAccount() (common.Address, error) {
	return _EulerEvaultImplementation.Contract.UnitOfAccount(&_EulerEvaultImplementation.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Approve(&_EulerEvaultImplementation.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Approve(&_EulerEvaultImplementation.TransactOpts, spender, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "borrow", amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Borrow(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Borrow(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) CheckVaultStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "checkVaultStatus")
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.CheckVaultStatus(&_EulerEvaultImplementation.TransactOpts)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.CheckVaultStatus(&_EulerEvaultImplementation.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) ConvertFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "convertFees")
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) ConvertFees() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.ConvertFees(&_EulerEvaultImplementation.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) ConvertFees() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.ConvertFees(&_EulerEvaultImplementation.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "deposit", amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Deposit(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Deposit(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) DisableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "disableBalanceForwarder")
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.DisableBalanceForwarder(&_EulerEvaultImplementation.TransactOpts)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.DisableBalanceForwarder(&_EulerEvaultImplementation.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) DisableController(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "disableController")
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) DisableController() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.DisableController(&_EulerEvaultImplementation.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) DisableController() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.DisableController(&_EulerEvaultImplementation.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) EnableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "enableBalanceForwarder")
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.EnableBalanceForwarder(&_EulerEvaultImplementation.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.EnableBalanceForwarder(&_EulerEvaultImplementation.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) FlashLoan(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "flashLoan", amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.FlashLoan(&_EulerEvaultImplementation.TransactOpts, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.FlashLoan(&_EulerEvaultImplementation.TransactOpts, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Initialize(opts *bind.TransactOpts, proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "initialize", proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Initialize(&_EulerEvaultImplementation.TransactOpts, proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Initialize(&_EulerEvaultImplementation.TransactOpts, proxyCreator)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Liquidate(opts *bind.TransactOpts, violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "liquidate", violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Liquidate(&_EulerEvaultImplementation.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Liquidate(&_EulerEvaultImplementation.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Mint(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "mint", amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Mint(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Mint(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) PullDebt(opts *bind.TransactOpts, amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "pullDebt", amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.PullDebt(&_EulerEvaultImplementation.TransactOpts, amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.PullDebt(&_EulerEvaultImplementation.TransactOpts, amount, from)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "redeem", amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Redeem(&_EulerEvaultImplementation.TransactOpts, amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Redeem(&_EulerEvaultImplementation.TransactOpts, amount, receiver, owner)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Repay(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "repay", amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Repay(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Repay(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) RepayWithShares(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "repayWithShares", amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.RepayWithShares(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.RepayWithShares(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetCaps(opts *bind.TransactOpts, supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setCaps", supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetCaps(&_EulerEvaultImplementation.TransactOpts, supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetCaps(&_EulerEvaultImplementation.TransactOpts, supplyCap, borrowCap)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetConfigFlags(opts *bind.TransactOpts, newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setConfigFlags", newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetConfigFlags(&_EulerEvaultImplementation.TransactOpts, newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetConfigFlags(&_EulerEvaultImplementation.TransactOpts, newConfigFlags)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetFeeReceiver(opts *bind.TransactOpts, newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setFeeReceiver", newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetFeeReceiver(&_EulerEvaultImplementation.TransactOpts, newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetFeeReceiver(&_EulerEvaultImplementation.TransactOpts, newFeeReceiver)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetGovernorAdmin(opts *bind.TransactOpts, newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setGovernorAdmin", newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetGovernorAdmin(&_EulerEvaultImplementation.TransactOpts, newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetGovernorAdmin(&_EulerEvaultImplementation.TransactOpts, newGovernorAdmin)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetHookConfig(opts *bind.TransactOpts, newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setHookConfig", newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetHookConfig(&_EulerEvaultImplementation.TransactOpts, newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetHookConfig(&_EulerEvaultImplementation.TransactOpts, newHookTarget, newHookedOps)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetInterestFee(opts *bind.TransactOpts, newFee uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setInterestFee", newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetInterestFee(&_EulerEvaultImplementation.TransactOpts, newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetInterestFee(&_EulerEvaultImplementation.TransactOpts, newFee)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetInterestRateModel(opts *bind.TransactOpts, newModel common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setInterestRateModel", newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetInterestRateModel(&_EulerEvaultImplementation.TransactOpts, newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetInterestRateModel(&_EulerEvaultImplementation.TransactOpts, newModel)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetLTV(opts *bind.TransactOpts, collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setLTV", collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetLTV(&_EulerEvaultImplementation.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetLTV(&_EulerEvaultImplementation.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetLiquidationCoolOffTime(opts *bind.TransactOpts, newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setLiquidationCoolOffTime", newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetLiquidationCoolOffTime(&_EulerEvaultImplementation.TransactOpts, newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetLiquidationCoolOffTime(&_EulerEvaultImplementation.TransactOpts, newCoolOffTime)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) SetMaxLiquidationDiscount(opts *bind.TransactOpts, newDiscount uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "setMaxLiquidationDiscount", newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetMaxLiquidationDiscount(&_EulerEvaultImplementation.TransactOpts, newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.SetMaxLiquidationDiscount(&_EulerEvaultImplementation.TransactOpts, newDiscount)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Skim(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "skim", amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Skim(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Skim(&_EulerEvaultImplementation.TransactOpts, amount, receiver)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Touch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "touch")
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Touch() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Touch(&_EulerEvaultImplementation.TransactOpts)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Touch() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Touch(&_EulerEvaultImplementation.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Transfer(&_EulerEvaultImplementation.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Transfer(&_EulerEvaultImplementation.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.TransferFrom(&_EulerEvaultImplementation.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.TransferFrom(&_EulerEvaultImplementation.TransactOpts, from, to, amount)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) TransferFromMax(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "transferFromMax", from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.TransferFromMax(&_EulerEvaultImplementation.TransactOpts, from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.TransferFromMax(&_EulerEvaultImplementation.TransactOpts, from, to)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) ViewDelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "viewDelegate")
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) ViewDelegate() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.ViewDelegate(&_EulerEvaultImplementation.TransactOpts)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) ViewDelegate() (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.ViewDelegate(&_EulerEvaultImplementation.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.contract.Transact(opts, "withdraw", amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Withdraw(&_EulerEvaultImplementation.TransactOpts, amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvaultImplementation *EulerEvaultImplementationTransactorSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvaultImplementation.Contract.Withdraw(&_EulerEvaultImplementation.TransactOpts, amount, receiver, owner)
}

// EulerEvaultImplementationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationApprovalIterator struct {
	Event *EulerEvaultImplementationApproval // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationApproval)
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
		it.Event = new(EulerEvaultImplementationApproval)
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
func (it *EulerEvaultImplementationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationApproval represents a Approval event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EulerEvaultImplementationApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationApprovalIterator{contract: _EulerEvaultImplementation.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationApproval)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseApproval(log types.Log) (*EulerEvaultImplementationApproval, error) {
	event := new(EulerEvaultImplementationApproval)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationBalanceForwarderStatusIterator is returned from FilterBalanceForwarderStatus and is used to iterate over the raw logs and unpacked data for BalanceForwarderStatus events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationBalanceForwarderStatusIterator struct {
	Event *EulerEvaultImplementationBalanceForwarderStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationBalanceForwarderStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationBalanceForwarderStatus)
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
		it.Event = new(EulerEvaultImplementationBalanceForwarderStatus)
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
func (it *EulerEvaultImplementationBalanceForwarderStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationBalanceForwarderStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationBalanceForwarderStatus represents a BalanceForwarderStatus event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationBalanceForwarderStatus struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceForwarderStatus is a free log retrieval operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterBalanceForwarderStatus(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultImplementationBalanceForwarderStatusIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationBalanceForwarderStatusIterator{contract: _EulerEvaultImplementation.contract, event: "BalanceForwarderStatus", logs: logs, sub: sub}, nil
}

// WatchBalanceForwarderStatus is a free log subscription operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchBalanceForwarderStatus(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationBalanceForwarderStatus, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationBalanceForwarderStatus)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
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

// ParseBalanceForwarderStatus is a log parse operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseBalanceForwarderStatus(log types.Log) (*EulerEvaultImplementationBalanceForwarderStatus, error) {
	event := new(EulerEvaultImplementationBalanceForwarderStatus)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationBorrowIterator struct {
	Event *EulerEvaultImplementationBorrow // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationBorrow)
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
		it.Event = new(EulerEvaultImplementationBorrow)
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
func (it *EulerEvaultImplementationBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationBorrow represents a Borrow event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationBorrow struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterBorrow(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultImplementationBorrowIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationBorrowIterator{contract: _EulerEvaultImplementation.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationBorrow, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationBorrow)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseBorrow(log types.Log) (*EulerEvaultImplementationBorrow, error) {
	event := new(EulerEvaultImplementationBorrow)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationConvertFeesIterator is returned from FilterConvertFees and is used to iterate over the raw logs and unpacked data for ConvertFees events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationConvertFeesIterator struct {
	Event *EulerEvaultImplementationConvertFees // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationConvertFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationConvertFees)
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
		it.Event = new(EulerEvaultImplementationConvertFees)
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
func (it *EulerEvaultImplementationConvertFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationConvertFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationConvertFees represents a ConvertFees event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationConvertFees struct {
	Sender           common.Address
	ProtocolReceiver common.Address
	GovernorReceiver common.Address
	ProtocolShares   *big.Int
	GovernorShares   *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConvertFees is a free log retrieval operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterConvertFees(opts *bind.FilterOpts, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (*EulerEvaultImplementationConvertFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var protocolReceiverRule []interface{}
	for _, protocolReceiverItem := range protocolReceiver {
		protocolReceiverRule = append(protocolReceiverRule, protocolReceiverItem)
	}
	var governorReceiverRule []interface{}
	for _, governorReceiverItem := range governorReceiver {
		governorReceiverRule = append(governorReceiverRule, governorReceiverItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationConvertFeesIterator{contract: _EulerEvaultImplementation.contract, event: "ConvertFees", logs: logs, sub: sub}, nil
}

// WatchConvertFees is a free log subscription operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchConvertFees(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationConvertFees, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var protocolReceiverRule []interface{}
	for _, protocolReceiverItem := range protocolReceiver {
		protocolReceiverRule = append(protocolReceiverRule, protocolReceiverItem)
	}
	var governorReceiverRule []interface{}
	for _, governorReceiverItem := range governorReceiver {
		governorReceiverRule = append(governorReceiverRule, governorReceiverItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationConvertFees)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "ConvertFees", log); err != nil {
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

// ParseConvertFees is a log parse operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseConvertFees(log types.Log) (*EulerEvaultImplementationConvertFees, error) {
	event := new(EulerEvaultImplementationConvertFees)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "ConvertFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationDebtSocializedIterator is returned from FilterDebtSocialized and is used to iterate over the raw logs and unpacked data for DebtSocialized events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationDebtSocializedIterator struct {
	Event *EulerEvaultImplementationDebtSocialized // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationDebtSocializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationDebtSocialized)
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
		it.Event = new(EulerEvaultImplementationDebtSocialized)
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
func (it *EulerEvaultImplementationDebtSocializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationDebtSocializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationDebtSocialized represents a DebtSocialized event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationDebtSocialized struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtSocialized is a free log retrieval operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterDebtSocialized(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultImplementationDebtSocializedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationDebtSocializedIterator{contract: _EulerEvaultImplementation.contract, event: "DebtSocialized", logs: logs, sub: sub}, nil
}

// WatchDebtSocialized is a free log subscription operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchDebtSocialized(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationDebtSocialized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationDebtSocialized)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
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

// ParseDebtSocialized is a log parse operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseDebtSocialized(log types.Log) (*EulerEvaultImplementationDebtSocialized, error) {
	event := new(EulerEvaultImplementationDebtSocialized)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationDepositIterator struct {
	Event *EulerEvaultImplementationDeposit // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationDeposit)
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
		it.Event = new(EulerEvaultImplementationDeposit)
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
func (it *EulerEvaultImplementationDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationDeposit represents a Deposit event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*EulerEvaultImplementationDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationDepositIterator{contract: _EulerEvaultImplementation.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationDeposit)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseDeposit(log types.Log) (*EulerEvaultImplementationDeposit, error) {
	event := new(EulerEvaultImplementationDeposit)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationEVaultCreatedIterator is returned from FilterEVaultCreated and is used to iterate over the raw logs and unpacked data for EVaultCreated events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationEVaultCreatedIterator struct {
	Event *EulerEvaultImplementationEVaultCreated // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationEVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationEVaultCreated)
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
		it.Event = new(EulerEvaultImplementationEVaultCreated)
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
func (it *EulerEvaultImplementationEVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationEVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationEVaultCreated represents a EVaultCreated event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationEVaultCreated struct {
	Creator common.Address
	Asset   common.Address
	DToken  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEVaultCreated is a free log retrieval operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterEVaultCreated(opts *bind.FilterOpts, creator []common.Address, asset []common.Address) (*EulerEvaultImplementationEVaultCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationEVaultCreatedIterator{contract: _EulerEvaultImplementation.contract, event: "EVaultCreated", logs: logs, sub: sub}, nil
}

// WatchEVaultCreated is a free log subscription operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchEVaultCreated(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationEVaultCreated, creator []common.Address, asset []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationEVaultCreated)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
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

// ParseEVaultCreated is a log parse operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseEVaultCreated(log types.Log) (*EulerEvaultImplementationEVaultCreated, error) {
	event := new(EulerEvaultImplementationEVaultCreated)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetCapsIterator is returned from FilterGovSetCaps and is used to iterate over the raw logs and unpacked data for GovSetCaps events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetCapsIterator struct {
	Event *EulerEvaultImplementationGovSetCaps // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetCapsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetCaps)
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
		it.Event = new(EulerEvaultImplementationGovSetCaps)
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
func (it *EulerEvaultImplementationGovSetCapsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetCapsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetCaps represents a GovSetCaps event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetCaps struct {
	NewSupplyCap uint16
	NewBorrowCap uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGovSetCaps is a free log retrieval operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetCaps(opts *bind.FilterOpts) (*EulerEvaultImplementationGovSetCapsIterator, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetCapsIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetCaps", logs: logs, sub: sub}, nil
}

// WatchGovSetCaps is a free log subscription operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetCaps(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetCaps) (event.Subscription, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetCaps)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
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

// ParseGovSetCaps is a log parse operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetCaps(log types.Log) (*EulerEvaultImplementationGovSetCaps, error) {
	event := new(EulerEvaultImplementationGovSetCaps)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetConfigFlagsIterator is returned from FilterGovSetConfigFlags and is used to iterate over the raw logs and unpacked data for GovSetConfigFlags events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetConfigFlagsIterator struct {
	Event *EulerEvaultImplementationGovSetConfigFlags // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetConfigFlagsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetConfigFlags)
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
		it.Event = new(EulerEvaultImplementationGovSetConfigFlags)
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
func (it *EulerEvaultImplementationGovSetConfigFlagsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetConfigFlagsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetConfigFlags represents a GovSetConfigFlags event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetConfigFlags struct {
	NewConfigFlags uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetConfigFlags is a free log retrieval operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetConfigFlags(opts *bind.FilterOpts) (*EulerEvaultImplementationGovSetConfigFlagsIterator, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetConfigFlagsIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetConfigFlags", logs: logs, sub: sub}, nil
}

// WatchGovSetConfigFlags is a free log subscription operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetConfigFlags(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetConfigFlags) (event.Subscription, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetConfigFlags)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
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

// ParseGovSetConfigFlags is a log parse operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetConfigFlags(log types.Log) (*EulerEvaultImplementationGovSetConfigFlags, error) {
	event := new(EulerEvaultImplementationGovSetConfigFlags)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetFeeReceiverIterator is returned from FilterGovSetFeeReceiver and is used to iterate over the raw logs and unpacked data for GovSetFeeReceiver events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetFeeReceiverIterator struct {
	Event *EulerEvaultImplementationGovSetFeeReceiver // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetFeeReceiver)
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
		it.Event = new(EulerEvaultImplementationGovSetFeeReceiver)
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
func (it *EulerEvaultImplementationGovSetFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetFeeReceiver represents a GovSetFeeReceiver event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetFeeReceiver struct {
	NewFeeReceiver common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetFeeReceiver is a free log retrieval operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetFeeReceiver(opts *bind.FilterOpts, newFeeReceiver []common.Address) (*EulerEvaultImplementationGovSetFeeReceiverIterator, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetFeeReceiverIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchGovSetFeeReceiver is a free log subscription operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetFeeReceiver(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetFeeReceiver, newFeeReceiver []common.Address) (event.Subscription, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetFeeReceiver)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
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

// ParseGovSetFeeReceiver is a log parse operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetFeeReceiver(log types.Log) (*EulerEvaultImplementationGovSetFeeReceiver, error) {
	event := new(EulerEvaultImplementationGovSetFeeReceiver)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetGovernorAdminIterator is returned from FilterGovSetGovernorAdmin and is used to iterate over the raw logs and unpacked data for GovSetGovernorAdmin events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetGovernorAdminIterator struct {
	Event *EulerEvaultImplementationGovSetGovernorAdmin // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetGovernorAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetGovernorAdmin)
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
		it.Event = new(EulerEvaultImplementationGovSetGovernorAdmin)
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
func (it *EulerEvaultImplementationGovSetGovernorAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetGovernorAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetGovernorAdmin represents a GovSetGovernorAdmin event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetGovernorAdmin struct {
	NewGovernorAdmin common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovSetGovernorAdmin is a free log retrieval operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetGovernorAdmin(opts *bind.FilterOpts, newGovernorAdmin []common.Address) (*EulerEvaultImplementationGovSetGovernorAdminIterator, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetGovernorAdminIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetGovernorAdmin", logs: logs, sub: sub}, nil
}

// WatchGovSetGovernorAdmin is a free log subscription operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetGovernorAdmin(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetGovernorAdmin, newGovernorAdmin []common.Address) (event.Subscription, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetGovernorAdmin)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
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

// ParseGovSetGovernorAdmin is a log parse operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetGovernorAdmin(log types.Log) (*EulerEvaultImplementationGovSetGovernorAdmin, error) {
	event := new(EulerEvaultImplementationGovSetGovernorAdmin)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetHookConfigIterator is returned from FilterGovSetHookConfig and is used to iterate over the raw logs and unpacked data for GovSetHookConfig events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetHookConfigIterator struct {
	Event *EulerEvaultImplementationGovSetHookConfig // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetHookConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetHookConfig)
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
		it.Event = new(EulerEvaultImplementationGovSetHookConfig)
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
func (it *EulerEvaultImplementationGovSetHookConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetHookConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetHookConfig represents a GovSetHookConfig event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetHookConfig struct {
	NewHookTarget common.Address
	NewHookedOps  uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGovSetHookConfig is a free log retrieval operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetHookConfig(opts *bind.FilterOpts, newHookTarget []common.Address) (*EulerEvaultImplementationGovSetHookConfigIterator, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetHookConfigIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetHookConfig", logs: logs, sub: sub}, nil
}

// WatchGovSetHookConfig is a free log subscription operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetHookConfig(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetHookConfig, newHookTarget []common.Address) (event.Subscription, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetHookConfig)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
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

// ParseGovSetHookConfig is a log parse operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetHookConfig(log types.Log) (*EulerEvaultImplementationGovSetHookConfig, error) {
	event := new(EulerEvaultImplementationGovSetHookConfig)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetInterestFeeIterator is returned from FilterGovSetInterestFee and is used to iterate over the raw logs and unpacked data for GovSetInterestFee events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetInterestFeeIterator struct {
	Event *EulerEvaultImplementationGovSetInterestFee // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetInterestFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetInterestFee)
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
		it.Event = new(EulerEvaultImplementationGovSetInterestFee)
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
func (it *EulerEvaultImplementationGovSetInterestFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetInterestFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetInterestFee represents a GovSetInterestFee event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetInterestFee struct {
	NewFee uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestFee is a free log retrieval operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetInterestFee(opts *bind.FilterOpts) (*EulerEvaultImplementationGovSetInterestFeeIterator, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetInterestFeeIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetInterestFee", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestFee is a free log subscription operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetInterestFee(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetInterestFee) (event.Subscription, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetInterestFee)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
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

// ParseGovSetInterestFee is a log parse operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetInterestFee(log types.Log) (*EulerEvaultImplementationGovSetInterestFee, error) {
	event := new(EulerEvaultImplementationGovSetInterestFee)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetInterestRateModelIterator is returned from FilterGovSetInterestRateModel and is used to iterate over the raw logs and unpacked data for GovSetInterestRateModel events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetInterestRateModelIterator struct {
	Event *EulerEvaultImplementationGovSetInterestRateModel // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetInterestRateModel)
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
		it.Event = new(EulerEvaultImplementationGovSetInterestRateModel)
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
func (it *EulerEvaultImplementationGovSetInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetInterestRateModel represents a GovSetInterestRateModel event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestRateModel is a free log retrieval operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetInterestRateModel(opts *bind.FilterOpts) (*EulerEvaultImplementationGovSetInterestRateModelIterator, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetInterestRateModelIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestRateModel is a free log subscription operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetInterestRateModel(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetInterestRateModel)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
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

// ParseGovSetInterestRateModel is a log parse operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetInterestRateModel(log types.Log) (*EulerEvaultImplementationGovSetInterestRateModel, error) {
	event := new(EulerEvaultImplementationGovSetInterestRateModel)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetLTVIterator is returned from FilterGovSetLTV and is used to iterate over the raw logs and unpacked data for GovSetLTV events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetLTVIterator struct {
	Event *EulerEvaultImplementationGovSetLTV // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetLTVIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetLTV)
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
		it.Event = new(EulerEvaultImplementationGovSetLTV)
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
func (it *EulerEvaultImplementationGovSetLTVIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetLTVIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetLTV represents a GovSetLTV event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetLTV struct {
	Collateral            common.Address
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovSetLTV is a free log retrieval operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetLTV(opts *bind.FilterOpts, collateral []common.Address) (*EulerEvaultImplementationGovSetLTVIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetLTVIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetLTV", logs: logs, sub: sub}, nil
}

// WatchGovSetLTV is a free log subscription operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetLTV(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetLTV, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetLTV)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
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

// ParseGovSetLTV is a log parse operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetLTV(log types.Log) (*EulerEvaultImplementationGovSetLTV, error) {
	event := new(EulerEvaultImplementationGovSetLTV)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetLiquidationCoolOffTimeIterator is returned from FilterGovSetLiquidationCoolOffTime and is used to iterate over the raw logs and unpacked data for GovSetLiquidationCoolOffTime events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetLiquidationCoolOffTimeIterator struct {
	Event *EulerEvaultImplementationGovSetLiquidationCoolOffTime // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetLiquidationCoolOffTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetLiquidationCoolOffTime)
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
		it.Event = new(EulerEvaultImplementationGovSetLiquidationCoolOffTime)
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
func (it *EulerEvaultImplementationGovSetLiquidationCoolOffTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetLiquidationCoolOffTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetLiquidationCoolOffTime represents a GovSetLiquidationCoolOffTime event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetLiquidationCoolOffTime struct {
	NewCoolOffTime uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetLiquidationCoolOffTime is a free log retrieval operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetLiquidationCoolOffTime(opts *bind.FilterOpts) (*EulerEvaultImplementationGovSetLiquidationCoolOffTimeIterator, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetLiquidationCoolOffTimeIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetLiquidationCoolOffTime", logs: logs, sub: sub}, nil
}

// WatchGovSetLiquidationCoolOffTime is a free log subscription operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetLiquidationCoolOffTime(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetLiquidationCoolOffTime) (event.Subscription, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetLiquidationCoolOffTime)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
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

// ParseGovSetLiquidationCoolOffTime is a log parse operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetLiquidationCoolOffTime(log types.Log) (*EulerEvaultImplementationGovSetLiquidationCoolOffTime, error) {
	event := new(EulerEvaultImplementationGovSetLiquidationCoolOffTime)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationGovSetMaxLiquidationDiscountIterator is returned from FilterGovSetMaxLiquidationDiscount and is used to iterate over the raw logs and unpacked data for GovSetMaxLiquidationDiscount events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetMaxLiquidationDiscountIterator struct {
	Event *EulerEvaultImplementationGovSetMaxLiquidationDiscount // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationGovSetMaxLiquidationDiscountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationGovSetMaxLiquidationDiscount)
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
		it.Event = new(EulerEvaultImplementationGovSetMaxLiquidationDiscount)
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
func (it *EulerEvaultImplementationGovSetMaxLiquidationDiscountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationGovSetMaxLiquidationDiscountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationGovSetMaxLiquidationDiscount represents a GovSetMaxLiquidationDiscount event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationGovSetMaxLiquidationDiscount struct {
	NewDiscount uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGovSetMaxLiquidationDiscount is a free log retrieval operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterGovSetMaxLiquidationDiscount(opts *bind.FilterOpts) (*EulerEvaultImplementationGovSetMaxLiquidationDiscountIterator, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationGovSetMaxLiquidationDiscountIterator{contract: _EulerEvaultImplementation.contract, event: "GovSetMaxLiquidationDiscount", logs: logs, sub: sub}, nil
}

// WatchGovSetMaxLiquidationDiscount is a free log subscription operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchGovSetMaxLiquidationDiscount(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationGovSetMaxLiquidationDiscount) (event.Subscription, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationGovSetMaxLiquidationDiscount)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
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

// ParseGovSetMaxLiquidationDiscount is a log parse operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseGovSetMaxLiquidationDiscount(log types.Log) (*EulerEvaultImplementationGovSetMaxLiquidationDiscount, error) {
	event := new(EulerEvaultImplementationGovSetMaxLiquidationDiscount)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationInterestAccruedIterator is returned from FilterInterestAccrued and is used to iterate over the raw logs and unpacked data for InterestAccrued events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationInterestAccruedIterator struct {
	Event *EulerEvaultImplementationInterestAccrued // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationInterestAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationInterestAccrued)
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
		it.Event = new(EulerEvaultImplementationInterestAccrued)
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
func (it *EulerEvaultImplementationInterestAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationInterestAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationInterestAccrued represents a InterestAccrued event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationInterestAccrued struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInterestAccrued is a free log retrieval operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterInterestAccrued(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultImplementationInterestAccruedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationInterestAccruedIterator{contract: _EulerEvaultImplementation.contract, event: "InterestAccrued", logs: logs, sub: sub}, nil
}

// WatchInterestAccrued is a free log subscription operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchInterestAccrued(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationInterestAccrued, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationInterestAccrued)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
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

// ParseInterestAccrued is a log parse operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseInterestAccrued(log types.Log) (*EulerEvaultImplementationInterestAccrued, error) {
	event := new(EulerEvaultImplementationInterestAccrued)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationLiquidateIterator struct {
	Event *EulerEvaultImplementationLiquidate // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationLiquidate)
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
		it.Event = new(EulerEvaultImplementationLiquidate)
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
func (it *EulerEvaultImplementationLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationLiquidate represents a Liquidate event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationLiquidate struct {
	Liquidator   common.Address
	Violator     common.Address
	Collateral   common.Address
	RepayAssets  *big.Int
	YieldBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, violator []common.Address) (*EulerEvaultImplementationLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationLiquidateIterator{contract: _EulerEvaultImplementation.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationLiquidate, liquidator []common.Address, violator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationLiquidate)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseLiquidate(log types.Log) (*EulerEvaultImplementationLiquidate, error) {
	event := new(EulerEvaultImplementationLiquidate)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationPullDebtIterator is returned from FilterPullDebt and is used to iterate over the raw logs and unpacked data for PullDebt events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationPullDebtIterator struct {
	Event *EulerEvaultImplementationPullDebt // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationPullDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationPullDebt)
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
		it.Event = new(EulerEvaultImplementationPullDebt)
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
func (it *EulerEvaultImplementationPullDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationPullDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationPullDebt represents a PullDebt event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationPullDebt struct {
	From   common.Address
	To     common.Address
	Assets *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPullDebt is a free log retrieval operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterPullDebt(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerEvaultImplementationPullDebtIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationPullDebtIterator{contract: _EulerEvaultImplementation.contract, event: "PullDebt", logs: logs, sub: sub}, nil
}

// WatchPullDebt is a free log subscription operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchPullDebt(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationPullDebt, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationPullDebt)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "PullDebt", log); err != nil {
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

// ParsePullDebt is a log parse operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParsePullDebt(log types.Log) (*EulerEvaultImplementationPullDebt, error) {
	event := new(EulerEvaultImplementationPullDebt)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "PullDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationRepayIterator struct {
	Event *EulerEvaultImplementationRepay // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationRepay)
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
		it.Event = new(EulerEvaultImplementationRepay)
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
func (it *EulerEvaultImplementationRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationRepay represents a Repay event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationRepay struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterRepay(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultImplementationRepayIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationRepayIterator{contract: _EulerEvaultImplementation.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationRepay, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationRepay)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseRepay(log types.Log) (*EulerEvaultImplementationRepay, error) {
	event := new(EulerEvaultImplementationRepay)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationTransferIterator struct {
	Event *EulerEvaultImplementationTransfer // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationTransfer)
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
		it.Event = new(EulerEvaultImplementationTransfer)
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
func (it *EulerEvaultImplementationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationTransfer represents a Transfer event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerEvaultImplementationTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationTransferIterator{contract: _EulerEvaultImplementation.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationTransfer)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseTransfer(log types.Log) (*EulerEvaultImplementationTransfer, error) {
	event := new(EulerEvaultImplementationTransfer)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationVaultStatusIterator is returned from FilterVaultStatus and is used to iterate over the raw logs and unpacked data for VaultStatus events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationVaultStatusIterator struct {
	Event *EulerEvaultImplementationVaultStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationVaultStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationVaultStatus)
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
		it.Event = new(EulerEvaultImplementationVaultStatus)
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
func (it *EulerEvaultImplementationVaultStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationVaultStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationVaultStatus represents a VaultStatus event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationVaultStatus struct {
	TotalShares         *big.Int
	TotalBorrows        *big.Int
	AccumulatedFees     *big.Int
	Cash                *big.Int
	InterestAccumulator *big.Int
	InterestRate        *big.Int
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVaultStatus is a free log retrieval operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterVaultStatus(opts *bind.FilterOpts) (*EulerEvaultImplementationVaultStatusIterator, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationVaultStatusIterator{contract: _EulerEvaultImplementation.contract, event: "VaultStatus", logs: logs, sub: sub}, nil
}

// WatchVaultStatus is a free log subscription operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchVaultStatus(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationVaultStatus) (event.Subscription, error) {

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationVaultStatus)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "VaultStatus", log); err != nil {
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

// ParseVaultStatus is a log parse operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseVaultStatus(log types.Log) (*EulerEvaultImplementationVaultStatus, error) {
	event := new(EulerEvaultImplementationVaultStatus)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "VaultStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultImplementationWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationWithdrawIterator struct {
	Event *EulerEvaultImplementationWithdraw // Event containing the contract specifics and raw log

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
func (it *EulerEvaultImplementationWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultImplementationWithdraw)
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
		it.Event = new(EulerEvaultImplementationWithdraw)
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
func (it *EulerEvaultImplementationWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultImplementationWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultImplementationWithdraw represents a Withdraw event raised by the EulerEvaultImplementation contract.
type EulerEvaultImplementationWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*EulerEvaultImplementationWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultImplementationWithdrawIterator{contract: _EulerEvaultImplementation.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EulerEvaultImplementationWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEvaultImplementation.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultImplementationWithdraw)
				if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvaultImplementation *EulerEvaultImplementationFilterer) ParseWithdraw(log types.Log) (*EulerEvaultImplementationWithdraw, error) {
	event := new(EulerEvaultImplementationWithdraw)
	if err := _EulerEvaultImplementation.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
