// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_evault

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

// EulerEvaultMetaData contains all meta data concerning the EulerEvault contract.
var EulerEvaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"protocolConfig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequenceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"}],\"internalType\":\"structBase.Integrations\",\"name\":\"integrations\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"initialize\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowing\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"riskManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"internalType\":\"structDispatch.DeployedModules\",\"name\":\"modules\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"E_AccountLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_AmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadAssetReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadBorrowCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadMaxLiquidationDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSharesOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSharesReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSupplyCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BorrowCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_CheckUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_CollateralDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ConfigAmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ControllerDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_DebtAmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_EmptyError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ExcessiveRepayAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_FlashLoanNotRepaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientCash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InvalidLTVAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LTVBorrow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LTVLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LiquidationCoolOff\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_MinYield\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NoLiability\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NoPriceOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotHookTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_OperationDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_OutstandingDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ProxyMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_RepayTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SelfLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_TransientState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ViolatorLiquidityDeferred\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ZeroAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ZeroShares\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"BalanceForwarderStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolReceiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governorReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"governorShares\",\"type\":\"uint256\"}],\"name\":\"ConvertFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"DebtSocialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dToken\",\"type\":\"address\"}],\"name\":\"EVaultCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newSupplyCap\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newBorrowCap\",\"type\":\"uint16\"}],\"name\":\"GovSetCaps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newConfigFlags\",\"type\":\"uint32\"}],\"name\":\"GovSetConfigFlags\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeReceiver\",\"type\":\"address\"}],\"name\":\"GovSetFeeReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernorAdmin\",\"type\":\"address\"}],\"name\":\"GovSetGovernorAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHookTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newHookedOps\",\"type\":\"uint32\"}],\"name\":\"GovSetHookConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"GovSetInterestFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"GovSetInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"targetTimestamp\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"name\":\"GovSetLTV\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newCoolOffTime\",\"type\":\"uint16\"}],\"name\":\"GovSetLiquidationCoolOffTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newDiscount\",\"type\":\"uint16\"}],\"name\":\"GovSetMaxLiquidationDiscount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"InterestAccrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldBalance\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"PullDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accumulatedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"VaultStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EVC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVBorrow\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVFull\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"targetTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVLiquidation\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LTVList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_BALANCE_FORWARDER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_BORROWING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_GOVERNANCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_INITIALIZE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_LIQUIDATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_RISKMANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"liquidation\",\"type\":\"bool\"}],\"name\":\"accountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"liquidation\",\"type\":\"bool\"}],\"name\":\"accountLiquidityFull\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"collateralValues\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFeesAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceForwarderEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceTrackerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"caps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"supplyCap\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"borrowCap\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"}],\"name\":\"checkAccountStatus\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"checkLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRepay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxYield\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkVaultStatus\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configFlags\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convertFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOfExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governorAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hookConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyCreator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestAccumulator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minYieldBalance\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationCoolOffTime\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidationDiscount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolConfigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"pullDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"repayWithShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"supplyCap\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"borrowCap\",\"type\":\"uint16\"}],\"name\":\"setCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newConfigFlags\",\"type\":\"uint32\"}],\"name\":\"setConfigFlags\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernorAdmin\",\"type\":\"address\"}],\"name\":\"setGovernorAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHookTarget\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newHookedOps\",\"type\":\"uint32\"}],\"name\":\"setHookConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"setInterestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"setInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"name\":\"setLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newCoolOffTime\",\"type\":\"uint16\"}],\"name\":\"setLiquidationCoolOffTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newDiscount\",\"type\":\"uint16\"}],\"name\":\"setMaxLiquidationDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"touch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferFromMax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unitOfAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewDelegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EulerEvaultABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerEvaultMetaData.ABI instead.
var EulerEvaultABI = EulerEvaultMetaData.ABI

// EulerEvault is an auto generated Go binding around an Ethereum contract.
type EulerEvault struct {
	EulerEvaultCaller     // Read-only binding to the contract
	EulerEvaultTransactor // Write-only binding to the contract
	EulerEvaultFilterer   // Log filterer for contract events
}

// EulerEvaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerEvaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerEvaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerEvaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEvaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerEvaultSession struct {
	Contract     *EulerEvault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EulerEvaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerEvaultCallerSession struct {
	Contract *EulerEvaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EulerEvaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerEvaultTransactorSession struct {
	Contract     *EulerEvaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EulerEvaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerEvaultRaw struct {
	Contract *EulerEvault // Generic contract binding to access the raw methods on
}

// EulerEvaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerEvaultCallerRaw struct {
	Contract *EulerEvaultCaller // Generic read-only contract binding to access the raw methods on
}

// EulerEvaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerEvaultTransactorRaw struct {
	Contract *EulerEvaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerEvault creates a new instance of EulerEvault, bound to a specific deployed contract.
func NewEulerEvault(address common.Address, backend bind.ContractBackend) (*EulerEvault, error) {
	contract, err := bindEulerEvault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerEvault{EulerEvaultCaller: EulerEvaultCaller{contract: contract}, EulerEvaultTransactor: EulerEvaultTransactor{contract: contract}, EulerEvaultFilterer: EulerEvaultFilterer{contract: contract}}, nil
}

// NewEulerEvaultCaller creates a new read-only instance of EulerEvault, bound to a specific deployed contract.
func NewEulerEvaultCaller(address common.Address, caller bind.ContractCaller) (*EulerEvaultCaller, error) {
	contract, err := bindEulerEvault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultCaller{contract: contract}, nil
}

// NewEulerEvaultTransactor creates a new write-only instance of EulerEvault, bound to a specific deployed contract.
func NewEulerEvaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerEvaultTransactor, error) {
	contract, err := bindEulerEvault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultTransactor{contract: contract}, nil
}

// NewEulerEvaultFilterer creates a new log filterer instance of EulerEvault, bound to a specific deployed contract.
func NewEulerEvaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerEvaultFilterer, error) {
	contract, err := bindEulerEvault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultFilterer{contract: contract}, nil
}

// bindEulerEvault binds a generic wrapper to an already deployed contract.
func bindEulerEvault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerEvaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEvault *EulerEvaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEvault.Contract.EulerEvaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEvault *EulerEvaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.Contract.EulerEvaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEvault *EulerEvaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEvault.Contract.EulerEvaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEvault *EulerEvaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEvault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEvault *EulerEvaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEvault *EulerEvaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEvault.Contract.contract.Transact(opts, method, params...)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEvault *EulerEvaultCaller) EVC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "EVC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEvault *EulerEvaultSession) EVC() (common.Address, error) {
	return _EulerEvault.Contract.EVC(&_EulerEvault.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) EVC() (common.Address, error) {
	return _EulerEvault.Contract.EVC(&_EulerEvault.CallOpts)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEvault *EulerEvaultCaller) LTVBorrow(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "LTVBorrow", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEvault *EulerEvaultSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EulerEvault.Contract.LTVBorrow(&_EulerEvault.CallOpts, collateral)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEvault *EulerEvaultCallerSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EulerEvault.Contract.LTVBorrow(&_EulerEvault.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvault *EulerEvaultCaller) LTVFull(opts *bind.CallOpts, collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "LTVFull", collateral)

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
func (_EulerEvault *EulerEvaultSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EulerEvault.Contract.LTVFull(&_EulerEvault.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvault *EulerEvaultCallerSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EulerEvault.Contract.LTVFull(&_EulerEvault.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEvault *EulerEvaultCaller) LTVLiquidation(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "LTVLiquidation", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEvault *EulerEvaultSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EulerEvault.Contract.LTVLiquidation(&_EulerEvault.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEvault *EulerEvaultCallerSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EulerEvault.Contract.LTVLiquidation(&_EulerEvault.CallOpts, collateral)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEvault *EulerEvaultCaller) LTVList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "LTVList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEvault *EulerEvaultSession) LTVList() ([]common.Address, error) {
	return _EulerEvault.Contract.LTVList(&_EulerEvault.CallOpts)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEvault *EulerEvaultCallerSession) LTVList() ([]common.Address, error) {
	return _EulerEvault.Contract.LTVList(&_EulerEvault.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEvault *EulerEvaultCaller) MODULEBALANCEFORWARDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "MODULE_BALANCE_FORWARDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEvault *EulerEvaultSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EulerEvault.Contract.MODULEBALANCEFORWARDER(&_EulerEvault.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EulerEvault.Contract.MODULEBALANCEFORWARDER(&_EulerEvault.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEvault *EulerEvaultCaller) MODULEBORROWING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "MODULE_BORROWING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEvault *EulerEvaultSession) MODULEBORROWING() (common.Address, error) {
	return _EulerEvault.Contract.MODULEBORROWING(&_EulerEvault.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) MODULEBORROWING() (common.Address, error) {
	return _EulerEvault.Contract.MODULEBORROWING(&_EulerEvault.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEvault *EulerEvaultCaller) MODULEGOVERNANCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "MODULE_GOVERNANCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEvault *EulerEvaultSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EulerEvault.Contract.MODULEGOVERNANCE(&_EulerEvault.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EulerEvault.Contract.MODULEGOVERNANCE(&_EulerEvault.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEvault *EulerEvaultCaller) MODULEINITIALIZE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "MODULE_INITIALIZE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEvault *EulerEvaultSession) MODULEINITIALIZE() (common.Address, error) {
	return _EulerEvault.Contract.MODULEINITIALIZE(&_EulerEvault.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) MODULEINITIALIZE() (common.Address, error) {
	return _EulerEvault.Contract.MODULEINITIALIZE(&_EulerEvault.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEvault *EulerEvaultCaller) MODULELIQUIDATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "MODULE_LIQUIDATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEvault *EulerEvaultSession) MODULELIQUIDATION() (common.Address, error) {
	return _EulerEvault.Contract.MODULELIQUIDATION(&_EulerEvault.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) MODULELIQUIDATION() (common.Address, error) {
	return _EulerEvault.Contract.MODULELIQUIDATION(&_EulerEvault.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEvault *EulerEvaultCaller) MODULERISKMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "MODULE_RISKMANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEvault *EulerEvaultSession) MODULERISKMANAGER() (common.Address, error) {
	return _EulerEvault.Contract.MODULERISKMANAGER(&_EulerEvault.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) MODULERISKMANAGER() (common.Address, error) {
	return _EulerEvault.Contract.MODULERISKMANAGER(&_EulerEvault.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEvault *EulerEvaultCaller) MODULETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "MODULE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEvault *EulerEvaultSession) MODULETOKEN() (common.Address, error) {
	return _EulerEvault.Contract.MODULETOKEN(&_EulerEvault.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) MODULETOKEN() (common.Address, error) {
	return _EulerEvault.Contract.MODULETOKEN(&_EulerEvault.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEvault *EulerEvaultCaller) MODULEVAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "MODULE_VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEvault *EulerEvaultSession) MODULEVAULT() (common.Address, error) {
	return _EulerEvault.Contract.MODULEVAULT(&_EulerEvault.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) MODULEVAULT() (common.Address, error) {
	return _EulerEvault.Contract.MODULEVAULT(&_EulerEvault.CallOpts)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerEvault *EulerEvaultCaller) AccountLiquidity(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "accountLiquidity", account, liquidation)

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
func (_EulerEvault *EulerEvaultSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EulerEvault.Contract.AccountLiquidity(&_EulerEvault.CallOpts, account, liquidation)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerEvault *EulerEvaultCallerSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EulerEvault.Contract.AccountLiquidity(&_EulerEvault.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerEvault *EulerEvaultCaller) AccountLiquidityFull(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "accountLiquidityFull", account, liquidation)

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
func (_EulerEvault *EulerEvaultSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EulerEvault.Contract.AccountLiquidityFull(&_EulerEvault.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerEvault *EulerEvaultCallerSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EulerEvault.Contract.AccountLiquidityFull(&_EulerEvault.CallOpts, account, liquidation)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) AccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "accumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) AccumulatedFees() (*big.Int, error) {
	return _EulerEvault.Contract.AccumulatedFees(&_EulerEvault.CallOpts)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) AccumulatedFees() (*big.Int, error) {
	return _EulerEvault.Contract.AccumulatedFees(&_EulerEvault.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) AccumulatedFeesAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "accumulatedFeesAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EulerEvault.Contract.AccumulatedFeesAssets(&_EulerEvault.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EulerEvault.Contract.AccumulatedFeesAssets(&_EulerEvault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.Allowance(&_EulerEvault.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.Allowance(&_EulerEvault.CallOpts, holder, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEvault *EulerEvaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEvault *EulerEvaultSession) Asset() (common.Address, error) {
	return _EulerEvault.Contract.Asset(&_EulerEvault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) Asset() (common.Address, error) {
	return _EulerEvault.Contract.Asset(&_EulerEvault.CallOpts)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEvault *EulerEvaultCaller) BalanceForwarderEnabled(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "balanceForwarderEnabled", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEvault *EulerEvaultSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EulerEvault.Contract.BalanceForwarderEnabled(&_EulerEvault.CallOpts, account)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEvault *EulerEvaultCallerSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EulerEvault.Contract.BalanceForwarderEnabled(&_EulerEvault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.BalanceOf(&_EulerEvault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.BalanceOf(&_EulerEvault.CallOpts, account)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEvault *EulerEvaultCaller) BalanceTrackerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "balanceTrackerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEvault *EulerEvaultSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEvault.Contract.BalanceTrackerAddress(&_EulerEvault.CallOpts)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEvault.Contract.BalanceTrackerAddress(&_EulerEvault.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerEvault *EulerEvaultCaller) Caps(opts *bind.CallOpts) (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "caps")

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
func (_EulerEvault *EulerEvaultSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EulerEvault.Contract.Caps(&_EulerEvault.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerEvault *EulerEvaultCallerSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EulerEvault.Contract.Caps(&_EulerEvault.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) Cash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) Cash() (*big.Int, error) {
	return _EulerEvault.Contract.Cash(&_EulerEvault.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) Cash() (*big.Int, error) {
	return _EulerEvault.Contract.Cash(&_EulerEvault.CallOpts)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEvault *EulerEvaultCaller) CheckAccountStatus(opts *bind.CallOpts, account common.Address, collaterals []common.Address) ([4]byte, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "checkAccountStatus", account, collaterals)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEvault *EulerEvaultSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EulerEvault.Contract.CheckAccountStatus(&_EulerEvault.CallOpts, account, collaterals)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEvault *EulerEvaultCallerSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EulerEvault.Contract.CheckAccountStatus(&_EulerEvault.CallOpts, account, collaterals)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerEvault *EulerEvaultCaller) CheckLiquidation(opts *bind.CallOpts, liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "checkLiquidation", liquidator, violator, collateral)

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
func (_EulerEvault *EulerEvaultSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EulerEvault.Contract.CheckLiquidation(&_EulerEvault.CallOpts, liquidator, violator, collateral)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerEvault *EulerEvaultCallerSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EulerEvault.Contract.CheckLiquidation(&_EulerEvault.CallOpts, liquidator, violator, collateral)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEvault *EulerEvaultCaller) ConfigFlags(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "configFlags")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEvault *EulerEvaultSession) ConfigFlags() (uint32, error) {
	return _EulerEvault.Contract.ConfigFlags(&_EulerEvault.CallOpts)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEvault *EulerEvaultCallerSession) ConfigFlags() (uint32, error) {
	return _EulerEvault.Contract.ConfigFlags(&_EulerEvault.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.ConvertToAssets(&_EulerEvault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.ConvertToAssets(&_EulerEvault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.ConvertToShares(&_EulerEvault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.ConvertToShares(&_EulerEvault.CallOpts, assets)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEvault *EulerEvaultCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEvault *EulerEvaultSession) Creator() (common.Address, error) {
	return _EulerEvault.Contract.Creator(&_EulerEvault.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) Creator() (common.Address, error) {
	return _EulerEvault.Contract.Creator(&_EulerEvault.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEvault *EulerEvaultCaller) DToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "dToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEvault *EulerEvaultSession) DToken() (common.Address, error) {
	return _EulerEvault.Contract.DToken(&_EulerEvault.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) DToken() (common.Address, error) {
	return _EulerEvault.Contract.DToken(&_EulerEvault.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) DebtOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "debtOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.DebtOf(&_EulerEvault.CallOpts, account)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.DebtOf(&_EulerEvault.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) DebtOfExact(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "debtOfExact", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.DebtOfExact(&_EulerEvault.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.DebtOfExact(&_EulerEvault.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEvault *EulerEvaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEvault *EulerEvaultSession) Decimals() (uint8, error) {
	return _EulerEvault.Contract.Decimals(&_EulerEvault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEvault *EulerEvaultCallerSession) Decimals() (uint8, error) {
	return _EulerEvault.Contract.Decimals(&_EulerEvault.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEvault *EulerEvaultCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEvault *EulerEvaultSession) FeeReceiver() (common.Address, error) {
	return _EulerEvault.Contract.FeeReceiver(&_EulerEvault.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) FeeReceiver() (common.Address, error) {
	return _EulerEvault.Contract.FeeReceiver(&_EulerEvault.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEvault *EulerEvaultCaller) GovernorAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "governorAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEvault *EulerEvaultSession) GovernorAdmin() (common.Address, error) {
	return _EulerEvault.Contract.GovernorAdmin(&_EulerEvault.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) GovernorAdmin() (common.Address, error) {
	return _EulerEvault.Contract.GovernorAdmin(&_EulerEvault.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerEvault *EulerEvaultCaller) HookConfig(opts *bind.CallOpts) (common.Address, uint32, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "hookConfig")

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
func (_EulerEvault *EulerEvaultSession) HookConfig() (common.Address, uint32, error) {
	return _EulerEvault.Contract.HookConfig(&_EulerEvault.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerEvault *EulerEvaultCallerSession) HookConfig() (common.Address, uint32, error) {
	return _EulerEvault.Contract.HookConfig(&_EulerEvault.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) InterestAccumulator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "interestAccumulator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) InterestAccumulator() (*big.Int, error) {
	return _EulerEvault.Contract.InterestAccumulator(&_EulerEvault.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) InterestAccumulator() (*big.Int, error) {
	return _EulerEvault.Contract.InterestAccumulator(&_EulerEvault.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEvault *EulerEvaultCaller) InterestFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "interestFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEvault *EulerEvaultSession) InterestFee() (uint16, error) {
	return _EulerEvault.Contract.InterestFee(&_EulerEvault.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEvault *EulerEvaultCallerSession) InterestFee() (uint16, error) {
	return _EulerEvault.Contract.InterestFee(&_EulerEvault.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) InterestRate() (*big.Int, error) {
	return _EulerEvault.Contract.InterestRate(&_EulerEvault.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) InterestRate() (*big.Int, error) {
	return _EulerEvault.Contract.InterestRate(&_EulerEvault.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEvault *EulerEvaultCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEvault *EulerEvaultSession) InterestRateModel() (common.Address, error) {
	return _EulerEvault.Contract.InterestRateModel(&_EulerEvault.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) InterestRateModel() (common.Address, error) {
	return _EulerEvault.Contract.InterestRateModel(&_EulerEvault.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEvault *EulerEvaultCaller) LiquidationCoolOffTime(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "liquidationCoolOffTime")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEvault *EulerEvaultSession) LiquidationCoolOffTime() (uint16, error) {
	return _EulerEvault.Contract.LiquidationCoolOffTime(&_EulerEvault.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEvault *EulerEvaultCallerSession) LiquidationCoolOffTime() (uint16, error) {
	return _EulerEvault.Contract.LiquidationCoolOffTime(&_EulerEvault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) MaxDeposit(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "maxDeposit", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.MaxDeposit(&_EulerEvault.CallOpts, account)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.MaxDeposit(&_EulerEvault.CallOpts, account)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEvault *EulerEvaultCaller) MaxLiquidationDiscount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "maxLiquidationDiscount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEvault *EulerEvaultSession) MaxLiquidationDiscount() (uint16, error) {
	return _EulerEvault.Contract.MaxLiquidationDiscount(&_EulerEvault.CallOpts)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEvault *EulerEvaultCallerSession) MaxLiquidationDiscount() (uint16, error) {
	return _EulerEvault.Contract.MaxLiquidationDiscount(&_EulerEvault.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) MaxMint(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "maxMint", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.MaxMint(&_EulerEvault.CallOpts, account)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.MaxMint(&_EulerEvault.CallOpts, account)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.MaxRedeem(&_EulerEvault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.MaxRedeem(&_EulerEvault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.MaxWithdraw(&_EulerEvault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EulerEvault.Contract.MaxWithdraw(&_EulerEvault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvault *EulerEvaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvault *EulerEvaultSession) Name() (string, error) {
	return _EulerEvault.Contract.Name(&_EulerEvault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEvault *EulerEvaultCallerSession) Name() (string, error) {
	return _EulerEvault.Contract.Name(&_EulerEvault.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEvault *EulerEvaultCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEvault *EulerEvaultSession) Oracle() (common.Address, error) {
	return _EulerEvault.Contract.Oracle(&_EulerEvault.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) Oracle() (common.Address, error) {
	return _EulerEvault.Contract.Oracle(&_EulerEvault.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEvault *EulerEvaultCaller) Permit2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "permit2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEvault *EulerEvaultSession) Permit2Address() (common.Address, error) {
	return _EulerEvault.Contract.Permit2Address(&_EulerEvault.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) Permit2Address() (common.Address, error) {
	return _EulerEvault.Contract.Permit2Address(&_EulerEvault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.PreviewDeposit(&_EulerEvault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.PreviewDeposit(&_EulerEvault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.PreviewMint(&_EulerEvault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.PreviewMint(&_EulerEvault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.PreviewRedeem(&_EulerEvault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.PreviewRedeem(&_EulerEvault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.PreviewWithdraw(&_EulerEvault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EulerEvault.Contract.PreviewWithdraw(&_EulerEvault.CallOpts, assets)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEvault *EulerEvaultCaller) ProtocolConfigAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "protocolConfigAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEvault *EulerEvaultSession) ProtocolConfigAddress() (common.Address, error) {
	return _EulerEvault.Contract.ProtocolConfigAddress(&_EulerEvault.CallOpts)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) ProtocolConfigAddress() (common.Address, error) {
	return _EulerEvault.Contract.ProtocolConfigAddress(&_EulerEvault.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEvault *EulerEvaultCaller) ProtocolFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "protocolFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEvault *EulerEvaultSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EulerEvault.Contract.ProtocolFeeReceiver(&_EulerEvault.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EulerEvault.Contract.ProtocolFeeReceiver(&_EulerEvault.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) ProtocolFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "protocolFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) ProtocolFeeShare() (*big.Int, error) {
	return _EulerEvault.Contract.ProtocolFeeShare(&_EulerEvault.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) ProtocolFeeShare() (*big.Int, error) {
	return _EulerEvault.Contract.ProtocolFeeShare(&_EulerEvault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEvault *EulerEvaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEvault *EulerEvaultSession) Symbol() (string, error) {
	return _EulerEvault.Contract.Symbol(&_EulerEvault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEvault *EulerEvaultCallerSession) Symbol() (string, error) {
	return _EulerEvault.Contract.Symbol(&_EulerEvault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) TotalAssets() (*big.Int, error) {
	return _EulerEvault.Contract.TotalAssets(&_EulerEvault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) TotalAssets() (*big.Int, error) {
	return _EulerEvault.Contract.TotalAssets(&_EulerEvault.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) TotalBorrows() (*big.Int, error) {
	return _EulerEvault.Contract.TotalBorrows(&_EulerEvault.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) TotalBorrows() (*big.Int, error) {
	return _EulerEvault.Contract.TotalBorrows(&_EulerEvault.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) TotalBorrowsExact(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "totalBorrowsExact")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) TotalBorrowsExact() (*big.Int, error) {
	return _EulerEvault.Contract.TotalBorrowsExact(&_EulerEvault.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) TotalBorrowsExact() (*big.Int, error) {
	return _EulerEvault.Contract.TotalBorrowsExact(&_EulerEvault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEvault *EulerEvaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEvault *EulerEvaultSession) TotalSupply() (*big.Int, error) {
	return _EulerEvault.Contract.TotalSupply(&_EulerEvault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEvault *EulerEvaultCallerSession) TotalSupply() (*big.Int, error) {
	return _EulerEvault.Contract.TotalSupply(&_EulerEvault.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEvault *EulerEvaultCaller) UnitOfAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEvault.contract.Call(opts, &out, "unitOfAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEvault *EulerEvaultSession) UnitOfAccount() (common.Address, error) {
	return _EulerEvault.Contract.UnitOfAccount(&_EulerEvault.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEvault *EulerEvaultCallerSession) UnitOfAccount() (common.Address, error) {
	return _EulerEvault.Contract.UnitOfAccount(&_EulerEvault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.Contract.Approve(&_EulerEvault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.Contract.Approve(&_EulerEvault.TransactOpts, spender, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "borrow", amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Borrow(&_EulerEvault.TransactOpts, amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactorSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Borrow(&_EulerEvault.TransactOpts, amount, receiver)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEvault *EulerEvaultTransactor) CheckVaultStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "checkVaultStatus")
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEvault *EulerEvaultSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EulerEvault.Contract.CheckVaultStatus(&_EulerEvault.TransactOpts)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEvault *EulerEvaultTransactorSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EulerEvault.Contract.CheckVaultStatus(&_EulerEvault.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEvault *EulerEvaultTransactor) ConvertFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "convertFees")
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEvault *EulerEvaultSession) ConvertFees() (*types.Transaction, error) {
	return _EulerEvault.Contract.ConvertFees(&_EulerEvault.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEvault *EulerEvaultTransactorSession) ConvertFees() (*types.Transaction, error) {
	return _EulerEvault.Contract.ConvertFees(&_EulerEvault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "deposit", amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Deposit(&_EulerEvault.TransactOpts, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactorSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Deposit(&_EulerEvault.TransactOpts, amount, receiver)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEvault *EulerEvaultTransactor) DisableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "disableBalanceForwarder")
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEvault *EulerEvaultSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEvault.Contract.DisableBalanceForwarder(&_EulerEvault.TransactOpts)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEvault *EulerEvaultTransactorSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEvault.Contract.DisableBalanceForwarder(&_EulerEvault.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEvault *EulerEvaultTransactor) DisableController(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "disableController")
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEvault *EulerEvaultSession) DisableController() (*types.Transaction, error) {
	return _EulerEvault.Contract.DisableController(&_EulerEvault.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEvault *EulerEvaultTransactorSession) DisableController() (*types.Transaction, error) {
	return _EulerEvault.Contract.DisableController(&_EulerEvault.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEvault *EulerEvaultTransactor) EnableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "enableBalanceForwarder")
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEvault *EulerEvaultSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEvault.Contract.EnableBalanceForwarder(&_EulerEvault.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEvault *EulerEvaultTransactorSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEvault.Contract.EnableBalanceForwarder(&_EulerEvault.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEvault *EulerEvaultTransactor) FlashLoan(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "flashLoan", amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEvault *EulerEvaultSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvault.Contract.FlashLoan(&_EulerEvault.TransactOpts, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEvault *EulerEvaultTransactorSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEvault.Contract.FlashLoan(&_EulerEvault.TransactOpts, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEvault *EulerEvaultTransactor) Initialize(opts *bind.TransactOpts, proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "initialize", proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEvault *EulerEvaultSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Initialize(&_EulerEvault.TransactOpts, proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEvault *EulerEvaultTransactorSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Initialize(&_EulerEvault.TransactOpts, proxyCreator)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEvault *EulerEvaultTransactor) Liquidate(opts *bind.TransactOpts, violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "liquidate", violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEvault *EulerEvaultSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEvault.Contract.Liquidate(&_EulerEvault.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEvault *EulerEvaultTransactorSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEvault.Contract.Liquidate(&_EulerEvault.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactor) Mint(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "mint", amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Mint(&_EulerEvault.TransactOpts, amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactorSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Mint(&_EulerEvault.TransactOpts, amount, receiver)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEvault *EulerEvaultTransactor) PullDebt(opts *bind.TransactOpts, amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "pullDebt", amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEvault *EulerEvaultSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.PullDebt(&_EulerEvault.TransactOpts, amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEvault *EulerEvaultTransactorSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.PullDebt(&_EulerEvault.TransactOpts, amount, from)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvault *EulerEvaultTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "redeem", amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvault *EulerEvaultSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Redeem(&_EulerEvault.TransactOpts, amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvault *EulerEvaultTransactorSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Redeem(&_EulerEvault.TransactOpts, amount, receiver, owner)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactor) Repay(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "repay", amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Repay(&_EulerEvault.TransactOpts, amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactorSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Repay(&_EulerEvault.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEvault *EulerEvaultTransactor) RepayWithShares(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "repayWithShares", amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEvault *EulerEvaultSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.RepayWithShares(&_EulerEvault.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEvault *EulerEvaultTransactorSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.RepayWithShares(&_EulerEvault.TransactOpts, amount, receiver)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEvault *EulerEvaultTransactor) SetCaps(opts *bind.TransactOpts, supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setCaps", supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEvault *EulerEvaultSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetCaps(&_EulerEvault.TransactOpts, supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetCaps(&_EulerEvault.TransactOpts, supplyCap, borrowCap)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEvault *EulerEvaultTransactor) SetConfigFlags(opts *bind.TransactOpts, newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setConfigFlags", newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEvault *EulerEvaultSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetConfigFlags(&_EulerEvault.TransactOpts, newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetConfigFlags(&_EulerEvault.TransactOpts, newConfigFlags)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEvault *EulerEvaultTransactor) SetFeeReceiver(opts *bind.TransactOpts, newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setFeeReceiver", newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEvault *EulerEvaultSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetFeeReceiver(&_EulerEvault.TransactOpts, newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetFeeReceiver(&_EulerEvault.TransactOpts, newFeeReceiver)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEvault *EulerEvaultTransactor) SetGovernorAdmin(opts *bind.TransactOpts, newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setGovernorAdmin", newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEvault *EulerEvaultSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetGovernorAdmin(&_EulerEvault.TransactOpts, newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetGovernorAdmin(&_EulerEvault.TransactOpts, newGovernorAdmin)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEvault *EulerEvaultTransactor) SetHookConfig(opts *bind.TransactOpts, newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setHookConfig", newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEvault *EulerEvaultSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetHookConfig(&_EulerEvault.TransactOpts, newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetHookConfig(&_EulerEvault.TransactOpts, newHookTarget, newHookedOps)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEvault *EulerEvaultTransactor) SetInterestFee(opts *bind.TransactOpts, newFee uint16) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setInterestFee", newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEvault *EulerEvaultSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetInterestFee(&_EulerEvault.TransactOpts, newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetInterestFee(&_EulerEvault.TransactOpts, newFee)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEvault *EulerEvaultTransactor) SetInterestRateModel(opts *bind.TransactOpts, newModel common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setInterestRateModel", newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEvault *EulerEvaultSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetInterestRateModel(&_EulerEvault.TransactOpts, newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetInterestRateModel(&_EulerEvault.TransactOpts, newModel)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEvault *EulerEvaultTransactor) SetLTV(opts *bind.TransactOpts, collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setLTV", collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEvault *EulerEvaultSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetLTV(&_EulerEvault.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetLTV(&_EulerEvault.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEvault *EulerEvaultTransactor) SetLiquidationCoolOffTime(opts *bind.TransactOpts, newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setLiquidationCoolOffTime", newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEvault *EulerEvaultSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetLiquidationCoolOffTime(&_EulerEvault.TransactOpts, newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetLiquidationCoolOffTime(&_EulerEvault.TransactOpts, newCoolOffTime)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEvault *EulerEvaultTransactor) SetMaxLiquidationDiscount(opts *bind.TransactOpts, newDiscount uint16) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "setMaxLiquidationDiscount", newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEvault *EulerEvaultSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetMaxLiquidationDiscount(&_EulerEvault.TransactOpts, newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEvault *EulerEvaultTransactorSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EulerEvault.Contract.SetMaxLiquidationDiscount(&_EulerEvault.TransactOpts, newDiscount)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactor) Skim(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "skim", amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Skim(&_EulerEvault.TransactOpts, amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEvault *EulerEvaultTransactorSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Skim(&_EulerEvault.TransactOpts, amount, receiver)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEvault *EulerEvaultTransactor) Touch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "touch")
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEvault *EulerEvaultSession) Touch() (*types.Transaction, error) {
	return _EulerEvault.Contract.Touch(&_EulerEvault.TransactOpts)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEvault *EulerEvaultTransactorSession) Touch() (*types.Transaction, error) {
	return _EulerEvault.Contract.Touch(&_EulerEvault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.Contract.Transfer(&_EulerEvault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.Contract.Transfer(&_EulerEvault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.Contract.TransferFrom(&_EulerEvault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEvault *EulerEvaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEvault.Contract.TransferFrom(&_EulerEvault.TransactOpts, from, to, amount)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEvault *EulerEvaultTransactor) TransferFromMax(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "transferFromMax", from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEvault *EulerEvaultSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.TransferFromMax(&_EulerEvault.TransactOpts, from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEvault *EulerEvaultTransactorSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.TransferFromMax(&_EulerEvault.TransactOpts, from, to)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEvault *EulerEvaultTransactor) ViewDelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "viewDelegate")
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEvault *EulerEvaultSession) ViewDelegate() (*types.Transaction, error) {
	return _EulerEvault.Contract.ViewDelegate(&_EulerEvault.TransactOpts)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEvault *EulerEvaultTransactorSession) ViewDelegate() (*types.Transaction, error) {
	return _EulerEvault.Contract.ViewDelegate(&_EulerEvault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvault *EulerEvaultTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvault.contract.Transact(opts, "withdraw", amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvault *EulerEvaultSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Withdraw(&_EulerEvault.TransactOpts, amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEvault *EulerEvaultTransactorSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEvault.Contract.Withdraw(&_EulerEvault.TransactOpts, amount, receiver, owner)
}

// EulerEvaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EulerEvault contract.
type EulerEvaultApprovalIterator struct {
	Event *EulerEvaultApproval // Event containing the contract specifics and raw log

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
func (it *EulerEvaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultApproval)
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
		it.Event = new(EulerEvaultApproval)
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
func (it *EulerEvaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultApproval represents a Approval event raised by the EulerEvault contract.
type EulerEvaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEvault *EulerEvaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EulerEvaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultApprovalIterator{contract: _EulerEvault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEvault *EulerEvaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EulerEvaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultApproval)
				if err := _EulerEvault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseApproval(log types.Log) (*EulerEvaultApproval, error) {
	event := new(EulerEvaultApproval)
	if err := _EulerEvault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultBalanceForwarderStatusIterator is returned from FilterBalanceForwarderStatus and is used to iterate over the raw logs and unpacked data for BalanceForwarderStatus events raised by the EulerEvault contract.
type EulerEvaultBalanceForwarderStatusIterator struct {
	Event *EulerEvaultBalanceForwarderStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvaultBalanceForwarderStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultBalanceForwarderStatus)
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
		it.Event = new(EulerEvaultBalanceForwarderStatus)
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
func (it *EulerEvaultBalanceForwarderStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultBalanceForwarderStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultBalanceForwarderStatus represents a BalanceForwarderStatus event raised by the EulerEvault contract.
type EulerEvaultBalanceForwarderStatus struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceForwarderStatus is a free log retrieval operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerEvault *EulerEvaultFilterer) FilterBalanceForwarderStatus(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultBalanceForwarderStatusIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultBalanceForwarderStatusIterator{contract: _EulerEvault.contract, event: "BalanceForwarderStatus", logs: logs, sub: sub}, nil
}

// WatchBalanceForwarderStatus is a free log subscription operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerEvault *EulerEvaultFilterer) WatchBalanceForwarderStatus(opts *bind.WatchOpts, sink chan<- *EulerEvaultBalanceForwarderStatus, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultBalanceForwarderStatus)
				if err := _EulerEvault.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseBalanceForwarderStatus(log types.Log) (*EulerEvaultBalanceForwarderStatus, error) {
	event := new(EulerEvaultBalanceForwarderStatus)
	if err := _EulerEvault.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the EulerEvault contract.
type EulerEvaultBorrowIterator struct {
	Event *EulerEvaultBorrow // Event containing the contract specifics and raw log

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
func (it *EulerEvaultBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultBorrow)
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
		it.Event = new(EulerEvaultBorrow)
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
func (it *EulerEvaultBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultBorrow represents a Borrow event raised by the EulerEvault contract.
type EulerEvaultBorrow struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) FilterBorrow(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultBorrowIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultBorrowIterator{contract: _EulerEvault.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *EulerEvaultBorrow, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultBorrow)
				if err := _EulerEvault.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseBorrow(log types.Log) (*EulerEvaultBorrow, error) {
	event := new(EulerEvaultBorrow)
	if err := _EulerEvault.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultConvertFeesIterator is returned from FilterConvertFees and is used to iterate over the raw logs and unpacked data for ConvertFees events raised by the EulerEvault contract.
type EulerEvaultConvertFeesIterator struct {
	Event *EulerEvaultConvertFees // Event containing the contract specifics and raw log

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
func (it *EulerEvaultConvertFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultConvertFees)
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
		it.Event = new(EulerEvaultConvertFees)
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
func (it *EulerEvaultConvertFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultConvertFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultConvertFees represents a ConvertFees event raised by the EulerEvault contract.
type EulerEvaultConvertFees struct {
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
func (_EulerEvault *EulerEvaultFilterer) FilterConvertFees(opts *bind.FilterOpts, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (*EulerEvaultConvertFeesIterator, error) {

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

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultConvertFeesIterator{contract: _EulerEvault.contract, event: "ConvertFees", logs: logs, sub: sub}, nil
}

// WatchConvertFees is a free log subscription operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EulerEvault *EulerEvaultFilterer) WatchConvertFees(opts *bind.WatchOpts, sink chan<- *EulerEvaultConvertFees, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultConvertFees)
				if err := _EulerEvault.contract.UnpackLog(event, "ConvertFees", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseConvertFees(log types.Log) (*EulerEvaultConvertFees, error) {
	event := new(EulerEvaultConvertFees)
	if err := _EulerEvault.contract.UnpackLog(event, "ConvertFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultDebtSocializedIterator is returned from FilterDebtSocialized and is used to iterate over the raw logs and unpacked data for DebtSocialized events raised by the EulerEvault contract.
type EulerEvaultDebtSocializedIterator struct {
	Event *EulerEvaultDebtSocialized // Event containing the contract specifics and raw log

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
func (it *EulerEvaultDebtSocializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultDebtSocialized)
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
		it.Event = new(EulerEvaultDebtSocialized)
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
func (it *EulerEvaultDebtSocializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultDebtSocializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultDebtSocialized represents a DebtSocialized event raised by the EulerEvault contract.
type EulerEvaultDebtSocialized struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtSocialized is a free log retrieval operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) FilterDebtSocialized(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultDebtSocializedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultDebtSocializedIterator{contract: _EulerEvault.contract, event: "DebtSocialized", logs: logs, sub: sub}, nil
}

// WatchDebtSocialized is a free log subscription operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) WatchDebtSocialized(opts *bind.WatchOpts, sink chan<- *EulerEvaultDebtSocialized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultDebtSocialized)
				if err := _EulerEvault.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseDebtSocialized(log types.Log) (*EulerEvaultDebtSocialized, error) {
	event := new(EulerEvaultDebtSocialized)
	if err := _EulerEvault.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EulerEvault contract.
type EulerEvaultDepositIterator struct {
	Event *EulerEvaultDeposit // Event containing the contract specifics and raw log

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
func (it *EulerEvaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultDeposit)
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
		it.Event = new(EulerEvaultDeposit)
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
func (it *EulerEvaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultDeposit represents a Deposit event raised by the EulerEvault contract.
type EulerEvaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvault *EulerEvaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*EulerEvaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultDepositIterator{contract: _EulerEvault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvault *EulerEvaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EulerEvaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultDeposit)
				if err := _EulerEvault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseDeposit(log types.Log) (*EulerEvaultDeposit, error) {
	event := new(EulerEvaultDeposit)
	if err := _EulerEvault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultEVaultCreatedIterator is returned from FilterEVaultCreated and is used to iterate over the raw logs and unpacked data for EVaultCreated events raised by the EulerEvault contract.
type EulerEvaultEVaultCreatedIterator struct {
	Event *EulerEvaultEVaultCreated // Event containing the contract specifics and raw log

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
func (it *EulerEvaultEVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultEVaultCreated)
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
		it.Event = new(EulerEvaultEVaultCreated)
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
func (it *EulerEvaultEVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultEVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultEVaultCreated represents a EVaultCreated event raised by the EulerEvault contract.
type EulerEvaultEVaultCreated struct {
	Creator common.Address
	Asset   common.Address
	DToken  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEVaultCreated is a free log retrieval operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerEvault *EulerEvaultFilterer) FilterEVaultCreated(opts *bind.FilterOpts, creator []common.Address, asset []common.Address) (*EulerEvaultEVaultCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultEVaultCreatedIterator{contract: _EulerEvault.contract, event: "EVaultCreated", logs: logs, sub: sub}, nil
}

// WatchEVaultCreated is a free log subscription operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerEvault *EulerEvaultFilterer) WatchEVaultCreated(opts *bind.WatchOpts, sink chan<- *EulerEvaultEVaultCreated, creator []common.Address, asset []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultEVaultCreated)
				if err := _EulerEvault.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseEVaultCreated(log types.Log) (*EulerEvaultEVaultCreated, error) {
	event := new(EulerEvaultEVaultCreated)
	if err := _EulerEvault.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetCapsIterator is returned from FilterGovSetCaps and is used to iterate over the raw logs and unpacked data for GovSetCaps events raised by the EulerEvault contract.
type EulerEvaultGovSetCapsIterator struct {
	Event *EulerEvaultGovSetCaps // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetCapsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetCaps)
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
		it.Event = new(EulerEvaultGovSetCaps)
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
func (it *EulerEvaultGovSetCapsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetCapsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetCaps represents a GovSetCaps event raised by the EulerEvault contract.
type EulerEvaultGovSetCaps struct {
	NewSupplyCap uint16
	NewBorrowCap uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGovSetCaps is a free log retrieval operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetCaps(opts *bind.FilterOpts) (*EulerEvaultGovSetCapsIterator, error) {

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetCapsIterator{contract: _EulerEvault.contract, event: "GovSetCaps", logs: logs, sub: sub}, nil
}

// WatchGovSetCaps is a free log subscription operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetCaps(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetCaps) (event.Subscription, error) {

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetCaps)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetCaps(log types.Log) (*EulerEvaultGovSetCaps, error) {
	event := new(EulerEvaultGovSetCaps)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetConfigFlagsIterator is returned from FilterGovSetConfigFlags and is used to iterate over the raw logs and unpacked data for GovSetConfigFlags events raised by the EulerEvault contract.
type EulerEvaultGovSetConfigFlagsIterator struct {
	Event *EulerEvaultGovSetConfigFlags // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetConfigFlagsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetConfigFlags)
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
		it.Event = new(EulerEvaultGovSetConfigFlags)
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
func (it *EulerEvaultGovSetConfigFlagsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetConfigFlagsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetConfigFlags represents a GovSetConfigFlags event raised by the EulerEvault contract.
type EulerEvaultGovSetConfigFlags struct {
	NewConfigFlags uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetConfigFlags is a free log retrieval operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetConfigFlags(opts *bind.FilterOpts) (*EulerEvaultGovSetConfigFlagsIterator, error) {

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetConfigFlagsIterator{contract: _EulerEvault.contract, event: "GovSetConfigFlags", logs: logs, sub: sub}, nil
}

// WatchGovSetConfigFlags is a free log subscription operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetConfigFlags(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetConfigFlags) (event.Subscription, error) {

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetConfigFlags)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetConfigFlags(log types.Log) (*EulerEvaultGovSetConfigFlags, error) {
	event := new(EulerEvaultGovSetConfigFlags)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetFeeReceiverIterator is returned from FilterGovSetFeeReceiver and is used to iterate over the raw logs and unpacked data for GovSetFeeReceiver events raised by the EulerEvault contract.
type EulerEvaultGovSetFeeReceiverIterator struct {
	Event *EulerEvaultGovSetFeeReceiver // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetFeeReceiver)
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
		it.Event = new(EulerEvaultGovSetFeeReceiver)
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
func (it *EulerEvaultGovSetFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetFeeReceiver represents a GovSetFeeReceiver event raised by the EulerEvault contract.
type EulerEvaultGovSetFeeReceiver struct {
	NewFeeReceiver common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetFeeReceiver is a free log retrieval operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetFeeReceiver(opts *bind.FilterOpts, newFeeReceiver []common.Address) (*EulerEvaultGovSetFeeReceiverIterator, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetFeeReceiverIterator{contract: _EulerEvault.contract, event: "GovSetFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchGovSetFeeReceiver is a free log subscription operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetFeeReceiver(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetFeeReceiver, newFeeReceiver []common.Address) (event.Subscription, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetFeeReceiver)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetFeeReceiver(log types.Log) (*EulerEvaultGovSetFeeReceiver, error) {
	event := new(EulerEvaultGovSetFeeReceiver)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetGovernorAdminIterator is returned from FilterGovSetGovernorAdmin and is used to iterate over the raw logs and unpacked data for GovSetGovernorAdmin events raised by the EulerEvault contract.
type EulerEvaultGovSetGovernorAdminIterator struct {
	Event *EulerEvaultGovSetGovernorAdmin // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetGovernorAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetGovernorAdmin)
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
		it.Event = new(EulerEvaultGovSetGovernorAdmin)
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
func (it *EulerEvaultGovSetGovernorAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetGovernorAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetGovernorAdmin represents a GovSetGovernorAdmin event raised by the EulerEvault contract.
type EulerEvaultGovSetGovernorAdmin struct {
	NewGovernorAdmin common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovSetGovernorAdmin is a free log retrieval operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetGovernorAdmin(opts *bind.FilterOpts, newGovernorAdmin []common.Address) (*EulerEvaultGovSetGovernorAdminIterator, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetGovernorAdminIterator{contract: _EulerEvault.contract, event: "GovSetGovernorAdmin", logs: logs, sub: sub}, nil
}

// WatchGovSetGovernorAdmin is a free log subscription operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetGovernorAdmin(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetGovernorAdmin, newGovernorAdmin []common.Address) (event.Subscription, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetGovernorAdmin)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetGovernorAdmin(log types.Log) (*EulerEvaultGovSetGovernorAdmin, error) {
	event := new(EulerEvaultGovSetGovernorAdmin)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetHookConfigIterator is returned from FilterGovSetHookConfig and is used to iterate over the raw logs and unpacked data for GovSetHookConfig events raised by the EulerEvault contract.
type EulerEvaultGovSetHookConfigIterator struct {
	Event *EulerEvaultGovSetHookConfig // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetHookConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetHookConfig)
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
		it.Event = new(EulerEvaultGovSetHookConfig)
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
func (it *EulerEvaultGovSetHookConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetHookConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetHookConfig represents a GovSetHookConfig event raised by the EulerEvault contract.
type EulerEvaultGovSetHookConfig struct {
	NewHookTarget common.Address
	NewHookedOps  uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGovSetHookConfig is a free log retrieval operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetHookConfig(opts *bind.FilterOpts, newHookTarget []common.Address) (*EulerEvaultGovSetHookConfigIterator, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetHookConfigIterator{contract: _EulerEvault.contract, event: "GovSetHookConfig", logs: logs, sub: sub}, nil
}

// WatchGovSetHookConfig is a free log subscription operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetHookConfig(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetHookConfig, newHookTarget []common.Address) (event.Subscription, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetHookConfig)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetHookConfig(log types.Log) (*EulerEvaultGovSetHookConfig, error) {
	event := new(EulerEvaultGovSetHookConfig)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetInterestFeeIterator is returned from FilterGovSetInterestFee and is used to iterate over the raw logs and unpacked data for GovSetInterestFee events raised by the EulerEvault contract.
type EulerEvaultGovSetInterestFeeIterator struct {
	Event *EulerEvaultGovSetInterestFee // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetInterestFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetInterestFee)
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
		it.Event = new(EulerEvaultGovSetInterestFee)
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
func (it *EulerEvaultGovSetInterestFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetInterestFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetInterestFee represents a GovSetInterestFee event raised by the EulerEvault contract.
type EulerEvaultGovSetInterestFee struct {
	NewFee uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestFee is a free log retrieval operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetInterestFee(opts *bind.FilterOpts) (*EulerEvaultGovSetInterestFeeIterator, error) {

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetInterestFeeIterator{contract: _EulerEvault.contract, event: "GovSetInterestFee", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestFee is a free log subscription operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetInterestFee(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetInterestFee) (event.Subscription, error) {

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetInterestFee)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetInterestFee(log types.Log) (*EulerEvaultGovSetInterestFee, error) {
	event := new(EulerEvaultGovSetInterestFee)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetInterestRateModelIterator is returned from FilterGovSetInterestRateModel and is used to iterate over the raw logs and unpacked data for GovSetInterestRateModel events raised by the EulerEvault contract.
type EulerEvaultGovSetInterestRateModelIterator struct {
	Event *EulerEvaultGovSetInterestRateModel // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetInterestRateModel)
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
		it.Event = new(EulerEvaultGovSetInterestRateModel)
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
func (it *EulerEvaultGovSetInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetInterestRateModel represents a GovSetInterestRateModel event raised by the EulerEvault contract.
type EulerEvaultGovSetInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestRateModel is a free log retrieval operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetInterestRateModel(opts *bind.FilterOpts) (*EulerEvaultGovSetInterestRateModelIterator, error) {

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetInterestRateModelIterator{contract: _EulerEvault.contract, event: "GovSetInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestRateModel is a free log subscription operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetInterestRateModel(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetInterestRateModel)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetInterestRateModel(log types.Log) (*EulerEvaultGovSetInterestRateModel, error) {
	event := new(EulerEvaultGovSetInterestRateModel)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetLTVIterator is returned from FilterGovSetLTV and is used to iterate over the raw logs and unpacked data for GovSetLTV events raised by the EulerEvault contract.
type EulerEvaultGovSetLTVIterator struct {
	Event *EulerEvaultGovSetLTV // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetLTVIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetLTV)
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
		it.Event = new(EulerEvaultGovSetLTV)
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
func (it *EulerEvaultGovSetLTVIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetLTVIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetLTV represents a GovSetLTV event raised by the EulerEvault contract.
type EulerEvaultGovSetLTV struct {
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
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetLTV(opts *bind.FilterOpts, collateral []common.Address) (*EulerEvaultGovSetLTVIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetLTVIterator{contract: _EulerEvault.contract, event: "GovSetLTV", logs: logs, sub: sub}, nil
}

// WatchGovSetLTV is a free log subscription operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetLTV(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetLTV, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetLTV)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetLTV(log types.Log) (*EulerEvaultGovSetLTV, error) {
	event := new(EulerEvaultGovSetLTV)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetLiquidationCoolOffTimeIterator is returned from FilterGovSetLiquidationCoolOffTime and is used to iterate over the raw logs and unpacked data for GovSetLiquidationCoolOffTime events raised by the EulerEvault contract.
type EulerEvaultGovSetLiquidationCoolOffTimeIterator struct {
	Event *EulerEvaultGovSetLiquidationCoolOffTime // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetLiquidationCoolOffTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetLiquidationCoolOffTime)
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
		it.Event = new(EulerEvaultGovSetLiquidationCoolOffTime)
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
func (it *EulerEvaultGovSetLiquidationCoolOffTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetLiquidationCoolOffTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetLiquidationCoolOffTime represents a GovSetLiquidationCoolOffTime event raised by the EulerEvault contract.
type EulerEvaultGovSetLiquidationCoolOffTime struct {
	NewCoolOffTime uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetLiquidationCoolOffTime is a free log retrieval operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetLiquidationCoolOffTime(opts *bind.FilterOpts) (*EulerEvaultGovSetLiquidationCoolOffTimeIterator, error) {

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetLiquidationCoolOffTimeIterator{contract: _EulerEvault.contract, event: "GovSetLiquidationCoolOffTime", logs: logs, sub: sub}, nil
}

// WatchGovSetLiquidationCoolOffTime is a free log subscription operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetLiquidationCoolOffTime(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetLiquidationCoolOffTime) (event.Subscription, error) {

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetLiquidationCoolOffTime)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetLiquidationCoolOffTime(log types.Log) (*EulerEvaultGovSetLiquidationCoolOffTime, error) {
	event := new(EulerEvaultGovSetLiquidationCoolOffTime)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultGovSetMaxLiquidationDiscountIterator is returned from FilterGovSetMaxLiquidationDiscount and is used to iterate over the raw logs and unpacked data for GovSetMaxLiquidationDiscount events raised by the EulerEvault contract.
type EulerEvaultGovSetMaxLiquidationDiscountIterator struct {
	Event *EulerEvaultGovSetMaxLiquidationDiscount // Event containing the contract specifics and raw log

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
func (it *EulerEvaultGovSetMaxLiquidationDiscountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultGovSetMaxLiquidationDiscount)
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
		it.Event = new(EulerEvaultGovSetMaxLiquidationDiscount)
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
func (it *EulerEvaultGovSetMaxLiquidationDiscountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultGovSetMaxLiquidationDiscountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultGovSetMaxLiquidationDiscount represents a GovSetMaxLiquidationDiscount event raised by the EulerEvault contract.
type EulerEvaultGovSetMaxLiquidationDiscount struct {
	NewDiscount uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGovSetMaxLiquidationDiscount is a free log retrieval operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerEvault *EulerEvaultFilterer) FilterGovSetMaxLiquidationDiscount(opts *bind.FilterOpts) (*EulerEvaultGovSetMaxLiquidationDiscountIterator, error) {

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultGovSetMaxLiquidationDiscountIterator{contract: _EulerEvault.contract, event: "GovSetMaxLiquidationDiscount", logs: logs, sub: sub}, nil
}

// WatchGovSetMaxLiquidationDiscount is a free log subscription operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerEvault *EulerEvaultFilterer) WatchGovSetMaxLiquidationDiscount(opts *bind.WatchOpts, sink chan<- *EulerEvaultGovSetMaxLiquidationDiscount) (event.Subscription, error) {

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultGovSetMaxLiquidationDiscount)
				if err := _EulerEvault.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseGovSetMaxLiquidationDiscount(log types.Log) (*EulerEvaultGovSetMaxLiquidationDiscount, error) {
	event := new(EulerEvaultGovSetMaxLiquidationDiscount)
	if err := _EulerEvault.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultInterestAccruedIterator is returned from FilterInterestAccrued and is used to iterate over the raw logs and unpacked data for InterestAccrued events raised by the EulerEvault contract.
type EulerEvaultInterestAccruedIterator struct {
	Event *EulerEvaultInterestAccrued // Event containing the contract specifics and raw log

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
func (it *EulerEvaultInterestAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultInterestAccrued)
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
		it.Event = new(EulerEvaultInterestAccrued)
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
func (it *EulerEvaultInterestAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultInterestAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultInterestAccrued represents a InterestAccrued event raised by the EulerEvault contract.
type EulerEvaultInterestAccrued struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInterestAccrued is a free log retrieval operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) FilterInterestAccrued(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultInterestAccruedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultInterestAccruedIterator{contract: _EulerEvault.contract, event: "InterestAccrued", logs: logs, sub: sub}, nil
}

// WatchInterestAccrued is a free log subscription operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) WatchInterestAccrued(opts *bind.WatchOpts, sink chan<- *EulerEvaultInterestAccrued, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultInterestAccrued)
				if err := _EulerEvault.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseInterestAccrued(log types.Log) (*EulerEvaultInterestAccrued, error) {
	event := new(EulerEvaultInterestAccrued)
	if err := _EulerEvault.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the EulerEvault contract.
type EulerEvaultLiquidateIterator struct {
	Event *EulerEvaultLiquidate // Event containing the contract specifics and raw log

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
func (it *EulerEvaultLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultLiquidate)
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
		it.Event = new(EulerEvaultLiquidate)
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
func (it *EulerEvaultLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultLiquidate represents a Liquidate event raised by the EulerEvault contract.
type EulerEvaultLiquidate struct {
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
func (_EulerEvault *EulerEvaultFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, violator []common.Address) (*EulerEvaultLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultLiquidateIterator{contract: _EulerEvault.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EulerEvault *EulerEvaultFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *EulerEvaultLiquidate, liquidator []common.Address, violator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultLiquidate)
				if err := _EulerEvault.contract.UnpackLog(event, "Liquidate", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseLiquidate(log types.Log) (*EulerEvaultLiquidate, error) {
	event := new(EulerEvaultLiquidate)
	if err := _EulerEvault.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultPullDebtIterator is returned from FilterPullDebt and is used to iterate over the raw logs and unpacked data for PullDebt events raised by the EulerEvault contract.
type EulerEvaultPullDebtIterator struct {
	Event *EulerEvaultPullDebt // Event containing the contract specifics and raw log

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
func (it *EulerEvaultPullDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultPullDebt)
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
		it.Event = new(EulerEvaultPullDebt)
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
func (it *EulerEvaultPullDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultPullDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultPullDebt represents a PullDebt event raised by the EulerEvault contract.
type EulerEvaultPullDebt struct {
	From   common.Address
	To     common.Address
	Assets *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPullDebt is a free log retrieval operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) FilterPullDebt(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerEvaultPullDebtIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultPullDebtIterator{contract: _EulerEvault.contract, event: "PullDebt", logs: logs, sub: sub}, nil
}

// WatchPullDebt is a free log subscription operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) WatchPullDebt(opts *bind.WatchOpts, sink chan<- *EulerEvaultPullDebt, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultPullDebt)
				if err := _EulerEvault.contract.UnpackLog(event, "PullDebt", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParsePullDebt(log types.Log) (*EulerEvaultPullDebt, error) {
	event := new(EulerEvaultPullDebt)
	if err := _EulerEvault.contract.UnpackLog(event, "PullDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the EulerEvault contract.
type EulerEvaultRepayIterator struct {
	Event *EulerEvaultRepay // Event containing the contract specifics and raw log

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
func (it *EulerEvaultRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultRepay)
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
		it.Event = new(EulerEvaultRepay)
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
func (it *EulerEvaultRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultRepay represents a Repay event raised by the EulerEvault contract.
type EulerEvaultRepay struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) FilterRepay(opts *bind.FilterOpts, account []common.Address) (*EulerEvaultRepayIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultRepayIterator{contract: _EulerEvault.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerEvault *EulerEvaultFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *EulerEvaultRepay, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultRepay)
				if err := _EulerEvault.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseRepay(log types.Log) (*EulerEvaultRepay, error) {
	event := new(EulerEvaultRepay)
	if err := _EulerEvault.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EulerEvault contract.
type EulerEvaultTransferIterator struct {
	Event *EulerEvaultTransfer // Event containing the contract specifics and raw log

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
func (it *EulerEvaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultTransfer)
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
		it.Event = new(EulerEvaultTransfer)
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
func (it *EulerEvaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultTransfer represents a Transfer event raised by the EulerEvault contract.
type EulerEvaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEvault *EulerEvaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerEvaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultTransferIterator{contract: _EulerEvault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEvault *EulerEvaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EulerEvaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultTransfer)
				if err := _EulerEvault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseTransfer(log types.Log) (*EulerEvaultTransfer, error) {
	event := new(EulerEvaultTransfer)
	if err := _EulerEvault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultVaultStatusIterator is returned from FilterVaultStatus and is used to iterate over the raw logs and unpacked data for VaultStatus events raised by the EulerEvault contract.
type EulerEvaultVaultStatusIterator struct {
	Event *EulerEvaultVaultStatus // Event containing the contract specifics and raw log

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
func (it *EulerEvaultVaultStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultVaultStatus)
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
		it.Event = new(EulerEvaultVaultStatus)
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
func (it *EulerEvaultVaultStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultVaultStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultVaultStatus represents a VaultStatus event raised by the EulerEvault contract.
type EulerEvaultVaultStatus struct {
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
func (_EulerEvault *EulerEvaultFilterer) FilterVaultStatus(opts *bind.FilterOpts) (*EulerEvaultVaultStatusIterator, error) {

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return &EulerEvaultVaultStatusIterator{contract: _EulerEvault.contract, event: "VaultStatus", logs: logs, sub: sub}, nil
}

// WatchVaultStatus is a free log subscription operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EulerEvault *EulerEvaultFilterer) WatchVaultStatus(opts *bind.WatchOpts, sink chan<- *EulerEvaultVaultStatus) (event.Subscription, error) {

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultVaultStatus)
				if err := _EulerEvault.contract.UnpackLog(event, "VaultStatus", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseVaultStatus(log types.Log) (*EulerEvaultVaultStatus, error) {
	event := new(EulerEvaultVaultStatus)
	if err := _EulerEvault.contract.UnpackLog(event, "VaultStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEvaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EulerEvault contract.
type EulerEvaultWithdrawIterator struct {
	Event *EulerEvaultWithdraw // Event containing the contract specifics and raw log

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
func (it *EulerEvaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEvaultWithdraw)
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
		it.Event = new(EulerEvaultWithdraw)
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
func (it *EulerEvaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEvaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEvaultWithdraw represents a Withdraw event raised by the EulerEvault contract.
type EulerEvaultWithdraw struct {
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
func (_EulerEvault *EulerEvaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*EulerEvaultWithdrawIterator, error) {

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

	logs, sub, err := _EulerEvault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEvaultWithdrawIterator{contract: _EulerEvault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEvault *EulerEvaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EulerEvaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEvault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEvaultWithdraw)
				if err := _EulerEvault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_EulerEvault *EulerEvaultFilterer) ParseWithdraw(log types.Log) (*EulerEvaultWithdraw, error) {
	event := new(EulerEvaultWithdraw)
	if err := _EulerEvault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
