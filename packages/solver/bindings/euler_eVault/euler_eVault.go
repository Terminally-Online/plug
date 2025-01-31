// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package euler_eVault

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

// EulerEVaultMetaData contains all meta data concerning the EulerEVault contract.
var EulerEVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"evc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"protocolConfig\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sequenceRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceTracker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"}],\"internalType\":\"structBase.Integrations\",\"name\":\"integrations\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"initialize\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowing\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"riskManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"balanceForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"internalType\":\"structDispatch.DeployedModules\",\"name\":\"modules\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"E_AccountLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_AmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadAssetReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadBorrowCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadCollateral\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadMaxLiquidationDiscount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSharesOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSharesReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BadSupplyCap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_BorrowCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_CheckUnauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_CollateralDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ConfigAmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ControllerDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_DebtAmountTooLargeToEncode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_EmptyError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ExcessiveRepayAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_FlashLoanNotRepaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientCash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InsufficientDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_InvalidLTVAsset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LTVBorrow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LTVLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_LiquidationCoolOff\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_MinYield\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NoLiability\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NoPriceOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotHookTarget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_OperationDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_OutstandingDebt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ProxyMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Reentrancy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_RepayTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SelfLiquidation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SelfTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_SupplyCapExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_TransientState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ViolatorLiquidityDeferred\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ZeroAssets\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"E_ZeroShares\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"BalanceForwarderStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolReceiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governorReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"governorShares\",\"type\":\"uint256\"}],\"name\":\"ConvertFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"DebtSocialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dToken\",\"type\":\"address\"}],\"name\":\"EVaultCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newSupplyCap\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newBorrowCap\",\"type\":\"uint16\"}],\"name\":\"GovSetCaps\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newConfigFlags\",\"type\":\"uint32\"}],\"name\":\"GovSetConfigFlags\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newFeeReceiver\",\"type\":\"address\"}],\"name\":\"GovSetFeeReceiver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernorAdmin\",\"type\":\"address\"}],\"name\":\"GovSetGovernorAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newHookTarget\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newHookedOps\",\"type\":\"uint32\"}],\"name\":\"GovSetHookConfig\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"GovSetInterestFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"GovSetInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint48\",\"name\":\"targetTimestamp\",\"type\":\"uint48\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"name\":\"GovSetLTV\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newCoolOffTime\",\"type\":\"uint16\"}],\"name\":\"GovSetLiquidationCoolOffTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newDiscount\",\"type\":\"uint16\"}],\"name\":\"GovSetMaxLiquidationDiscount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"InterestAccrued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAssets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yieldBalance\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"PullDebt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accumulatedFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"VaultStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"EVC\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVBorrow\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVFull\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"initialLiquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint48\",\"name\":\"targetTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"LTVLiquidation\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LTVList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_BALANCE_FORWARDER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_BORROWING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_GOVERNANCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_INITIALIZE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_LIQUIDATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_RISKMANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_VAULT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"liquidation\",\"type\":\"bool\"}],\"name\":\"accountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"liquidation\",\"type\":\"bool\"}],\"name\":\"accountLiquidityFull\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"collateralValues\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"liabilityValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accumulatedFeesAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceForwarderEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balanceTrackerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"caps\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"supplyCap\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"borrowCap\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"collaterals\",\"type\":\"address[]\"}],\"name\":\"checkAccountStatus\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"}],\"name\":\"checkLiquidation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRepay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxYield\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkVaultStatus\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configFlags\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convertFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"debtOfExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBalanceForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governorAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hookConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxyCreator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestAccumulator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"violator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAssets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minYieldBalance\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationCoolOffTime\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLiquidationDiscount\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"permit2Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolConfigAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeReceiver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"pullDebt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"repayWithShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"supplyCap\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"borrowCap\",\"type\":\"uint16\"}],\"name\":\"setCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newConfigFlags\",\"type\":\"uint32\"}],\"name\":\"setConfigFlags\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newFeeReceiver\",\"type\":\"address\"}],\"name\":\"setFeeReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernorAdmin\",\"type\":\"address\"}],\"name\":\"setGovernorAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newHookTarget\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"newHookedOps\",\"type\":\"uint32\"}],\"name\":\"setHookConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newFee\",\"type\":\"uint16\"}],\"name\":\"setInterestFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newModel\",\"type\":\"address\"}],\"name\":\"setInterestRateModel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"borrowLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"liquidationLTV\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rampDuration\",\"type\":\"uint32\"}],\"name\":\"setLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newCoolOffTime\",\"type\":\"uint16\"}],\"name\":\"setLiquidationCoolOffTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newDiscount\",\"type\":\"uint16\"}],\"name\":\"setMaxLiquidationDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowsExact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"touch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferFromMax\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unitOfAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewDelegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EulerEVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use EulerEVaultMetaData.ABI instead.
var EulerEVaultABI = EulerEVaultMetaData.ABI

// EulerEVault is an auto generated Go binding around an Ethereum contract.
type EulerEVault struct {
	EulerEVaultCaller     // Read-only binding to the contract
	EulerEVaultTransactor // Write-only binding to the contract
	EulerEVaultFilterer   // Log filterer for contract events
}

// EulerEVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type EulerEVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EulerEVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EulerEVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EulerEVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EulerEVaultSession struct {
	Contract     *EulerEVault      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EulerEVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EulerEVaultCallerSession struct {
	Contract *EulerEVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EulerEVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EulerEVaultTransactorSession struct {
	Contract     *EulerEVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EulerEVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type EulerEVaultRaw struct {
	Contract *EulerEVault // Generic contract binding to access the raw methods on
}

// EulerEVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EulerEVaultCallerRaw struct {
	Contract *EulerEVaultCaller // Generic read-only contract binding to access the raw methods on
}

// EulerEVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EulerEVaultTransactorRaw struct {
	Contract *EulerEVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEulerEVault creates a new instance of EulerEVault, bound to a specific deployed contract.
func NewEulerEVault(address common.Address, backend bind.ContractBackend) (*EulerEVault, error) {
	contract, err := bindEulerEVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EulerEVault{EulerEVaultCaller: EulerEVaultCaller{contract: contract}, EulerEVaultTransactor: EulerEVaultTransactor{contract: contract}, EulerEVaultFilterer: EulerEVaultFilterer{contract: contract}}, nil
}

// NewEulerEVaultCaller creates a new read-only instance of EulerEVault, bound to a specific deployed contract.
func NewEulerEVaultCaller(address common.Address, caller bind.ContractCaller) (*EulerEVaultCaller, error) {
	contract, err := bindEulerEVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultCaller{contract: contract}, nil
}

// NewEulerEVaultTransactor creates a new write-only instance of EulerEVault, bound to a specific deployed contract.
func NewEulerEVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*EulerEVaultTransactor, error) {
	contract, err := bindEulerEVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultTransactor{contract: contract}, nil
}

// NewEulerEVaultFilterer creates a new log filterer instance of EulerEVault, bound to a specific deployed contract.
func NewEulerEVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*EulerEVaultFilterer, error) {
	contract, err := bindEulerEVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultFilterer{contract: contract}, nil
}

// bindEulerEVault binds a generic wrapper to an already deployed contract.
func bindEulerEVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EulerEVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEVault *EulerEVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEVault.Contract.EulerEVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEVault *EulerEVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.Contract.EulerEVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEVault *EulerEVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEVault.Contract.EulerEVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EulerEVault *EulerEVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EulerEVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EulerEVault *EulerEVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EulerEVault *EulerEVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EulerEVault.Contract.contract.Transact(opts, method, params...)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEVault *EulerEVaultCaller) EVC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "EVC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEVault *EulerEVaultSession) EVC() (common.Address, error) {
	return _EulerEVault.Contract.EVC(&_EulerEVault.CallOpts)
}

// EVC is a free data retrieval call binding the contract method 0xa70354a1.
//
// Solidity: function EVC() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) EVC() (common.Address, error) {
	return _EulerEVault.Contract.EVC(&_EulerEVault.CallOpts)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEVault *EulerEVaultCaller) LTVBorrow(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "LTVBorrow", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEVault *EulerEVaultSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EulerEVault.Contract.LTVBorrow(&_EulerEVault.CallOpts, collateral)
}

// LTVBorrow is a free data retrieval call binding the contract method 0xbf58094d.
//
// Solidity: function LTVBorrow(address collateral) view returns(uint16)
func (_EulerEVault *EulerEVaultCallerSession) LTVBorrow(collateral common.Address) (uint16, error) {
	return _EulerEVault.Contract.LTVBorrow(&_EulerEVault.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEVault *EulerEVaultCaller) LTVFull(opts *bind.CallOpts, collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "LTVFull", collateral)

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
func (_EulerEVault *EulerEVaultSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EulerEVault.Contract.LTVFull(&_EulerEVault.CallOpts, collateral)
}

// LTVFull is a free data retrieval call binding the contract method 0x33708d0c.
//
// Solidity: function LTVFull(address collateral) view returns(uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEVault *EulerEVaultCallerSession) LTVFull(collateral common.Address) (struct {
	BorrowLTV             uint16
	LiquidationLTV        uint16
	InitialLiquidationLTV uint16
	TargetTimestamp       *big.Int
	RampDuration          uint32
}, error) {
	return _EulerEVault.Contract.LTVFull(&_EulerEVault.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEVault *EulerEVaultCaller) LTVLiquidation(opts *bind.CallOpts, collateral common.Address) (uint16, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "LTVLiquidation", collateral)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEVault *EulerEVaultSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EulerEVault.Contract.LTVLiquidation(&_EulerEVault.CallOpts, collateral)
}

// LTVLiquidation is a free data retrieval call binding the contract method 0xaf5aaeeb.
//
// Solidity: function LTVLiquidation(address collateral) view returns(uint16)
func (_EulerEVault *EulerEVaultCallerSession) LTVLiquidation(collateral common.Address) (uint16, error) {
	return _EulerEVault.Contract.LTVLiquidation(&_EulerEVault.CallOpts, collateral)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEVault *EulerEVaultCaller) LTVList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "LTVList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEVault *EulerEVaultSession) LTVList() ([]common.Address, error) {
	return _EulerEVault.Contract.LTVList(&_EulerEVault.CallOpts)
}

// LTVList is a free data retrieval call binding the contract method 0x6a16ef84.
//
// Solidity: function LTVList() view returns(address[])
func (_EulerEVault *EulerEVaultCallerSession) LTVList() ([]common.Address, error) {
	return _EulerEVault.Contract.LTVList(&_EulerEVault.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEVault *EulerEVaultCaller) MODULEBALANCEFORWARDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "MODULE_BALANCE_FORWARDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEVault *EulerEVaultSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EulerEVault.Contract.MODULEBALANCEFORWARDER(&_EulerEVault.CallOpts)
}

// MODULEBALANCEFORWARDER is a free data retrieval call binding the contract method 0x883e3875.
//
// Solidity: function MODULE_BALANCE_FORWARDER() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) MODULEBALANCEFORWARDER() (common.Address, error) {
	return _EulerEVault.Contract.MODULEBALANCEFORWARDER(&_EulerEVault.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEVault *EulerEVaultCaller) MODULEBORROWING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "MODULE_BORROWING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEVault *EulerEVaultSession) MODULEBORROWING() (common.Address, error) {
	return _EulerEVault.Contract.MODULEBORROWING(&_EulerEVault.CallOpts)
}

// MODULEBORROWING is a free data retrieval call binding the contract method 0x14c054bc.
//
// Solidity: function MODULE_BORROWING() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) MODULEBORROWING() (common.Address, error) {
	return _EulerEVault.Contract.MODULEBORROWING(&_EulerEVault.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEVault *EulerEVaultCaller) MODULEGOVERNANCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "MODULE_GOVERNANCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEVault *EulerEVaultSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EulerEVault.Contract.MODULEGOVERNANCE(&_EulerEVault.CallOpts)
}

// MODULEGOVERNANCE is a free data retrieval call binding the contract method 0xb4cd541b.
//
// Solidity: function MODULE_GOVERNANCE() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) MODULEGOVERNANCE() (common.Address, error) {
	return _EulerEVault.Contract.MODULEGOVERNANCE(&_EulerEVault.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEVault *EulerEVaultCaller) MODULEINITIALIZE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "MODULE_INITIALIZE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEVault *EulerEVaultSession) MODULEINITIALIZE() (common.Address, error) {
	return _EulerEVault.Contract.MODULEINITIALIZE(&_EulerEVault.CallOpts)
}

// MODULEINITIALIZE is a free data retrieval call binding the contract method 0xad80ad0b.
//
// Solidity: function MODULE_INITIALIZE() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) MODULEINITIALIZE() (common.Address, error) {
	return _EulerEVault.Contract.MODULEINITIALIZE(&_EulerEVault.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEVault *EulerEVaultCaller) MODULELIQUIDATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "MODULE_LIQUIDATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEVault *EulerEVaultSession) MODULELIQUIDATION() (common.Address, error) {
	return _EulerEVault.Contract.MODULELIQUIDATION(&_EulerEVault.CallOpts)
}

// MODULELIQUIDATION is a free data retrieval call binding the contract method 0x42895567.
//
// Solidity: function MODULE_LIQUIDATION() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) MODULELIQUIDATION() (common.Address, error) {
	return _EulerEVault.Contract.MODULELIQUIDATION(&_EulerEVault.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEVault *EulerEVaultCaller) MODULERISKMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "MODULE_RISKMANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEVault *EulerEVaultSession) MODULERISKMANAGER() (common.Address, error) {
	return _EulerEVault.Contract.MODULERISKMANAGER(&_EulerEVault.CallOpts)
}

// MODULERISKMANAGER is a free data retrieval call binding the contract method 0x7d5f2e4e.
//
// Solidity: function MODULE_RISKMANAGER() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) MODULERISKMANAGER() (common.Address, error) {
	return _EulerEVault.Contract.MODULERISKMANAGER(&_EulerEVault.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEVault *EulerEVaultCaller) MODULETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "MODULE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEVault *EulerEVaultSession) MODULETOKEN() (common.Address, error) {
	return _EulerEVault.Contract.MODULETOKEN(&_EulerEVault.CallOpts)
}

// MODULETOKEN is a free data retrieval call binding the contract method 0x5fa23055.
//
// Solidity: function MODULE_TOKEN() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) MODULETOKEN() (common.Address, error) {
	return _EulerEVault.Contract.MODULETOKEN(&_EulerEVault.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEVault *EulerEVaultCaller) MODULEVAULT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "MODULE_VAULT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEVault *EulerEVaultSession) MODULEVAULT() (common.Address, error) {
	return _EulerEVault.Contract.MODULEVAULT(&_EulerEVault.CallOpts)
}

// MODULEVAULT is a free data retrieval call binding the contract method 0xe2f206e5.
//
// Solidity: function MODULE_VAULT() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) MODULEVAULT() (common.Address, error) {
	return _EulerEVault.Contract.MODULEVAULT(&_EulerEVault.CallOpts)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerEVault *EulerEVaultCaller) AccountLiquidity(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "accountLiquidity", account, liquidation)

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
func (_EulerEVault *EulerEVaultSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EulerEVault.Contract.AccountLiquidity(&_EulerEVault.CallOpts, account, liquidation)
}

// AccountLiquidity is a free data retrieval call binding the contract method 0xa824bf67.
//
// Solidity: function accountLiquidity(address account, bool liquidation) view returns(uint256 collateralValue, uint256 liabilityValue)
func (_EulerEVault *EulerEVaultCallerSession) AccountLiquidity(account common.Address, liquidation bool) (struct {
	CollateralValue *big.Int
	LiabilityValue  *big.Int
}, error) {
	return _EulerEVault.Contract.AccountLiquidity(&_EulerEVault.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerEVault *EulerEVaultCaller) AccountLiquidityFull(opts *bind.CallOpts, account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "accountLiquidityFull", account, liquidation)

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
func (_EulerEVault *EulerEVaultSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EulerEVault.Contract.AccountLiquidityFull(&_EulerEVault.CallOpts, account, liquidation)
}

// AccountLiquidityFull is a free data retrieval call binding the contract method 0xc7b0e3a3.
//
// Solidity: function accountLiquidityFull(address account, bool liquidation) view returns(address[] collaterals, uint256[] collateralValues, uint256 liabilityValue)
func (_EulerEVault *EulerEVaultCallerSession) AccountLiquidityFull(account common.Address, liquidation bool) (struct {
	Collaterals      []common.Address
	CollateralValues []*big.Int
	LiabilityValue   *big.Int
}, error) {
	return _EulerEVault.Contract.AccountLiquidityFull(&_EulerEVault.CallOpts, account, liquidation)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) AccumulatedFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "accumulatedFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) AccumulatedFees() (*big.Int, error) {
	return _EulerEVault.Contract.AccumulatedFees(&_EulerEVault.CallOpts)
}

// AccumulatedFees is a free data retrieval call binding the contract method 0x587f5ed7.
//
// Solidity: function accumulatedFees() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) AccumulatedFees() (*big.Int, error) {
	return _EulerEVault.Contract.AccumulatedFees(&_EulerEVault.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) AccumulatedFeesAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "accumulatedFeesAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EulerEVault.Contract.AccumulatedFeesAssets(&_EulerEVault.CallOpts)
}

// AccumulatedFeesAssets is a free data retrieval call binding the contract method 0xf6e50f58.
//
// Solidity: function accumulatedFeesAssets() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) AccumulatedFeesAssets() (*big.Int, error) {
	return _EulerEVault.Contract.AccumulatedFeesAssets(&_EulerEVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) Allowance(opts *bind.CallOpts, holder common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "allowance", holder, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.Allowance(&_EulerEVault.CallOpts, holder, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address holder, address spender) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) Allowance(holder common.Address, spender common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.Allowance(&_EulerEVault.CallOpts, holder, spender)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEVault *EulerEVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEVault *EulerEVaultSession) Asset() (common.Address, error) {
	return _EulerEVault.Contract.Asset(&_EulerEVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) Asset() (common.Address, error) {
	return _EulerEVault.Contract.Asset(&_EulerEVault.CallOpts)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEVault *EulerEVaultCaller) BalanceForwarderEnabled(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "balanceForwarderEnabled", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEVault *EulerEVaultSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EulerEVault.Contract.BalanceForwarderEnabled(&_EulerEVault.CallOpts, account)
}

// BalanceForwarderEnabled is a free data retrieval call binding the contract method 0xe15c82ec.
//
// Solidity: function balanceForwarderEnabled(address account) view returns(bool)
func (_EulerEVault *EulerEVaultCallerSession) BalanceForwarderEnabled(account common.Address) (bool, error) {
	return _EulerEVault.Contract.BalanceForwarderEnabled(&_EulerEVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.BalanceOf(&_EulerEVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.BalanceOf(&_EulerEVault.CallOpts, account)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEVault *EulerEVaultCaller) BalanceTrackerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "balanceTrackerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEVault *EulerEVaultSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEVault.Contract.BalanceTrackerAddress(&_EulerEVault.CallOpts)
}

// BalanceTrackerAddress is a free data retrieval call binding the contract method 0xece6a7fa.
//
// Solidity: function balanceTrackerAddress() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) BalanceTrackerAddress() (common.Address, error) {
	return _EulerEVault.Contract.BalanceTrackerAddress(&_EulerEVault.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerEVault *EulerEVaultCaller) Caps(opts *bind.CallOpts) (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "caps")

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
func (_EulerEVault *EulerEVaultSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EulerEVault.Contract.Caps(&_EulerEVault.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x18e22d98.
//
// Solidity: function caps() view returns(uint16 supplyCap, uint16 borrowCap)
func (_EulerEVault *EulerEVaultCallerSession) Caps() (struct {
	SupplyCap uint16
	BorrowCap uint16
}, error) {
	return _EulerEVault.Contract.Caps(&_EulerEVault.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) Cash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "cash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) Cash() (*big.Int, error) {
	return _EulerEVault.Contract.Cash(&_EulerEVault.CallOpts)
}

// Cash is a free data retrieval call binding the contract method 0x961be391.
//
// Solidity: function cash() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) Cash() (*big.Int, error) {
	return _EulerEVault.Contract.Cash(&_EulerEVault.CallOpts)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEVault *EulerEVaultCaller) CheckAccountStatus(opts *bind.CallOpts, account common.Address, collaterals []common.Address) ([4]byte, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "checkAccountStatus", account, collaterals)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEVault *EulerEVaultSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EulerEVault.Contract.CheckAccountStatus(&_EulerEVault.CallOpts, account, collaterals)
}

// CheckAccountStatus is a free data retrieval call binding the contract method 0xb168c58f.
//
// Solidity: function checkAccountStatus(address account, address[] collaterals) view returns(bytes4)
func (_EulerEVault *EulerEVaultCallerSession) CheckAccountStatus(account common.Address, collaterals []common.Address) ([4]byte, error) {
	return _EulerEVault.Contract.CheckAccountStatus(&_EulerEVault.CallOpts, account, collaterals)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerEVault *EulerEVaultCaller) CheckLiquidation(opts *bind.CallOpts, liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "checkLiquidation", liquidator, violator, collateral)

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
func (_EulerEVault *EulerEVaultSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EulerEVault.Contract.CheckLiquidation(&_EulerEVault.CallOpts, liquidator, violator, collateral)
}

// CheckLiquidation is a free data retrieval call binding the contract method 0x88aa6f12.
//
// Solidity: function checkLiquidation(address liquidator, address violator, address collateral) view returns(uint256 maxRepay, uint256 maxYield)
func (_EulerEVault *EulerEVaultCallerSession) CheckLiquidation(liquidator common.Address, violator common.Address, collateral common.Address) (struct {
	MaxRepay *big.Int
	MaxYield *big.Int
}, error) {
	return _EulerEVault.Contract.CheckLiquidation(&_EulerEVault.CallOpts, liquidator, violator, collateral)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEVault *EulerEVaultCaller) ConfigFlags(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "configFlags")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEVault *EulerEVaultSession) ConfigFlags() (uint32, error) {
	return _EulerEVault.Contract.ConfigFlags(&_EulerEVault.CallOpts)
}

// ConfigFlags is a free data retrieval call binding the contract method 0x2b38a367.
//
// Solidity: function configFlags() view returns(uint32)
func (_EulerEVault *EulerEVaultCallerSession) ConfigFlags() (uint32, error) {
	return _EulerEVault.Contract.ConfigFlags(&_EulerEVault.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.ConvertToAssets(&_EulerEVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.ConvertToAssets(&_EulerEVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.ConvertToShares(&_EulerEVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.ConvertToShares(&_EulerEVault.CallOpts, assets)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEVault *EulerEVaultCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEVault *EulerEVaultSession) Creator() (common.Address, error) {
	return _EulerEVault.Contract.Creator(&_EulerEVault.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) Creator() (common.Address, error) {
	return _EulerEVault.Contract.Creator(&_EulerEVault.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEVault *EulerEVaultCaller) DToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "dToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEVault *EulerEVaultSession) DToken() (common.Address, error) {
	return _EulerEVault.Contract.DToken(&_EulerEVault.CallOpts)
}

// DToken is a free data retrieval call binding the contract method 0xd9d7858a.
//
// Solidity: function dToken() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) DToken() (common.Address, error) {
	return _EulerEVault.Contract.DToken(&_EulerEVault.CallOpts)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) DebtOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "debtOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.DebtOf(&_EulerEVault.CallOpts, account)
}

// DebtOf is a free data retrieval call binding the contract method 0xd283e75f.
//
// Solidity: function debtOf(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) DebtOf(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.DebtOf(&_EulerEVault.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) DebtOfExact(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "debtOfExact", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.DebtOfExact(&_EulerEVault.CallOpts, account)
}

// DebtOfExact is a free data retrieval call binding the contract method 0xab49b7f1.
//
// Solidity: function debtOfExact(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) DebtOfExact(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.DebtOfExact(&_EulerEVault.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEVault *EulerEVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEVault *EulerEVaultSession) Decimals() (uint8, error) {
	return _EulerEVault.Contract.Decimals(&_EulerEVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EulerEVault *EulerEVaultCallerSession) Decimals() (uint8, error) {
	return _EulerEVault.Contract.Decimals(&_EulerEVault.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEVault *EulerEVaultCaller) FeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "feeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEVault *EulerEVaultSession) FeeReceiver() (common.Address, error) {
	return _EulerEVault.Contract.FeeReceiver(&_EulerEVault.CallOpts)
}

// FeeReceiver is a free data retrieval call binding the contract method 0xb3f00674.
//
// Solidity: function feeReceiver() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) FeeReceiver() (common.Address, error) {
	return _EulerEVault.Contract.FeeReceiver(&_EulerEVault.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEVault *EulerEVaultCaller) GovernorAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "governorAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEVault *EulerEVaultSession) GovernorAdmin() (common.Address, error) {
	return _EulerEVault.Contract.GovernorAdmin(&_EulerEVault.CallOpts)
}

// GovernorAdmin is a free data retrieval call binding the contract method 0x6ce98c29.
//
// Solidity: function governorAdmin() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) GovernorAdmin() (common.Address, error) {
	return _EulerEVault.Contract.GovernorAdmin(&_EulerEVault.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerEVault *EulerEVaultCaller) HookConfig(opts *bind.CallOpts) (common.Address, uint32, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "hookConfig")

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
func (_EulerEVault *EulerEVaultSession) HookConfig() (common.Address, uint32, error) {
	return _EulerEVault.Contract.HookConfig(&_EulerEVault.CallOpts)
}

// HookConfig is a free data retrieval call binding the contract method 0xcf349b7d.
//
// Solidity: function hookConfig() view returns(address, uint32)
func (_EulerEVault *EulerEVaultCallerSession) HookConfig() (common.Address, uint32, error) {
	return _EulerEVault.Contract.HookConfig(&_EulerEVault.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) InterestAccumulator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "interestAccumulator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) InterestAccumulator() (*big.Int, error) {
	return _EulerEVault.Contract.InterestAccumulator(&_EulerEVault.CallOpts)
}

// InterestAccumulator is a free data retrieval call binding the contract method 0x087a6007.
//
// Solidity: function interestAccumulator() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) InterestAccumulator() (*big.Int, error) {
	return _EulerEVault.Contract.InterestAccumulator(&_EulerEVault.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEVault *EulerEVaultCaller) InterestFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "interestFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEVault *EulerEVaultSession) InterestFee() (uint16, error) {
	return _EulerEVault.Contract.InterestFee(&_EulerEVault.CallOpts)
}

// InterestFee is a free data retrieval call binding the contract method 0xa75df498.
//
// Solidity: function interestFee() view returns(uint16)
func (_EulerEVault *EulerEVaultCallerSession) InterestFee() (uint16, error) {
	return _EulerEVault.Contract.InterestFee(&_EulerEVault.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) InterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "interestRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) InterestRate() (*big.Int, error) {
	return _EulerEVault.Contract.InterestRate(&_EulerEVault.CallOpts)
}

// InterestRate is a free data retrieval call binding the contract method 0x7c3a00fd.
//
// Solidity: function interestRate() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) InterestRate() (*big.Int, error) {
	return _EulerEVault.Contract.InterestRate(&_EulerEVault.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEVault *EulerEVaultCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEVault *EulerEVaultSession) InterestRateModel() (common.Address, error) {
	return _EulerEVault.Contract.InterestRateModel(&_EulerEVault.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) InterestRateModel() (common.Address, error) {
	return _EulerEVault.Contract.InterestRateModel(&_EulerEVault.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEVault *EulerEVaultCaller) LiquidationCoolOffTime(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "liquidationCoolOffTime")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEVault *EulerEVaultSession) LiquidationCoolOffTime() (uint16, error) {
	return _EulerEVault.Contract.LiquidationCoolOffTime(&_EulerEVault.CallOpts)
}

// LiquidationCoolOffTime is a free data retrieval call binding the contract method 0x4abdb959.
//
// Solidity: function liquidationCoolOffTime() view returns(uint16)
func (_EulerEVault *EulerEVaultCallerSession) LiquidationCoolOffTime() (uint16, error) {
	return _EulerEVault.Contract.LiquidationCoolOffTime(&_EulerEVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) MaxDeposit(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "maxDeposit", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.MaxDeposit(&_EulerEVault.CallOpts, account)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) MaxDeposit(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.MaxDeposit(&_EulerEVault.CallOpts, account)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEVault *EulerEVaultCaller) MaxLiquidationDiscount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "maxLiquidationDiscount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEVault *EulerEVaultSession) MaxLiquidationDiscount() (uint16, error) {
	return _EulerEVault.Contract.MaxLiquidationDiscount(&_EulerEVault.CallOpts)
}

// MaxLiquidationDiscount is a free data retrieval call binding the contract method 0x4f7e43df.
//
// Solidity: function maxLiquidationDiscount() view returns(uint16)
func (_EulerEVault *EulerEVaultCallerSession) MaxLiquidationDiscount() (uint16, error) {
	return _EulerEVault.Contract.MaxLiquidationDiscount(&_EulerEVault.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) MaxMint(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "maxMint", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.MaxMint(&_EulerEVault.CallOpts, account)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address account) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) MaxMint(account common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.MaxMint(&_EulerEVault.CallOpts, account)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.MaxRedeem(&_EulerEVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.MaxRedeem(&_EulerEVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.MaxWithdraw(&_EulerEVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _EulerEVault.Contract.MaxWithdraw(&_EulerEVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEVault *EulerEVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEVault *EulerEVaultSession) Name() (string, error) {
	return _EulerEVault.Contract.Name(&_EulerEVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EulerEVault *EulerEVaultCallerSession) Name() (string, error) {
	return _EulerEVault.Contract.Name(&_EulerEVault.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEVault *EulerEVaultCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEVault *EulerEVaultSession) Oracle() (common.Address, error) {
	return _EulerEVault.Contract.Oracle(&_EulerEVault.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) Oracle() (common.Address, error) {
	return _EulerEVault.Contract.Oracle(&_EulerEVault.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEVault *EulerEVaultCaller) Permit2Address(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "permit2Address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEVault *EulerEVaultSession) Permit2Address() (common.Address, error) {
	return _EulerEVault.Contract.Permit2Address(&_EulerEVault.CallOpts)
}

// Permit2Address is a free data retrieval call binding the contract method 0xc5224983.
//
// Solidity: function permit2Address() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) Permit2Address() (common.Address, error) {
	return _EulerEVault.Contract.Permit2Address(&_EulerEVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.PreviewDeposit(&_EulerEVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.PreviewDeposit(&_EulerEVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.PreviewMint(&_EulerEVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.PreviewMint(&_EulerEVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.PreviewRedeem(&_EulerEVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.PreviewRedeem(&_EulerEVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.PreviewWithdraw(&_EulerEVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _EulerEVault.Contract.PreviewWithdraw(&_EulerEVault.CallOpts, assets)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEVault *EulerEVaultCaller) ProtocolConfigAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "protocolConfigAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEVault *EulerEVaultSession) ProtocolConfigAddress() (common.Address, error) {
	return _EulerEVault.Contract.ProtocolConfigAddress(&_EulerEVault.CallOpts)
}

// ProtocolConfigAddress is a free data retrieval call binding the contract method 0x539bd5bf.
//
// Solidity: function protocolConfigAddress() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) ProtocolConfigAddress() (common.Address, error) {
	return _EulerEVault.Contract.ProtocolConfigAddress(&_EulerEVault.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEVault *EulerEVaultCaller) ProtocolFeeReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "protocolFeeReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEVault *EulerEVaultSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EulerEVault.Contract.ProtocolFeeReceiver(&_EulerEVault.CallOpts)
}

// ProtocolFeeReceiver is a free data retrieval call binding the contract method 0x39a51be5.
//
// Solidity: function protocolFeeReceiver() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) ProtocolFeeReceiver() (common.Address, error) {
	return _EulerEVault.Contract.ProtocolFeeReceiver(&_EulerEVault.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) ProtocolFeeShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "protocolFeeShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) ProtocolFeeShare() (*big.Int, error) {
	return _EulerEVault.Contract.ProtocolFeeShare(&_EulerEVault.CallOpts)
}

// ProtocolFeeShare is a free data retrieval call binding the contract method 0x960b26a2.
//
// Solidity: function protocolFeeShare() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) ProtocolFeeShare() (*big.Int, error) {
	return _EulerEVault.Contract.ProtocolFeeShare(&_EulerEVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEVault *EulerEVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEVault *EulerEVaultSession) Symbol() (string, error) {
	return _EulerEVault.Contract.Symbol(&_EulerEVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EulerEVault *EulerEVaultCallerSession) Symbol() (string, error) {
	return _EulerEVault.Contract.Symbol(&_EulerEVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) TotalAssets() (*big.Int, error) {
	return _EulerEVault.Contract.TotalAssets(&_EulerEVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _EulerEVault.Contract.TotalAssets(&_EulerEVault.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) TotalBorrows() (*big.Int, error) {
	return _EulerEVault.Contract.TotalBorrows(&_EulerEVault.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) TotalBorrows() (*big.Int, error) {
	return _EulerEVault.Contract.TotalBorrows(&_EulerEVault.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) TotalBorrowsExact(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "totalBorrowsExact")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) TotalBorrowsExact() (*big.Int, error) {
	return _EulerEVault.Contract.TotalBorrowsExact(&_EulerEVault.CallOpts)
}

// TotalBorrowsExact is a free data retrieval call binding the contract method 0xe388be7b.
//
// Solidity: function totalBorrowsExact() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) TotalBorrowsExact() (*big.Int, error) {
	return _EulerEVault.Contract.TotalBorrowsExact(&_EulerEVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEVault *EulerEVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEVault *EulerEVaultSession) TotalSupply() (*big.Int, error) {
	return _EulerEVault.Contract.TotalSupply(&_EulerEVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EulerEVault *EulerEVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _EulerEVault.Contract.TotalSupply(&_EulerEVault.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEVault *EulerEVaultCaller) UnitOfAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EulerEVault.contract.Call(opts, &out, "unitOfAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEVault *EulerEVaultSession) UnitOfAccount() (common.Address, error) {
	return _EulerEVault.Contract.UnitOfAccount(&_EulerEVault.CallOpts)
}

// UnitOfAccount is a free data retrieval call binding the contract method 0x3e833364.
//
// Solidity: function unitOfAccount() view returns(address)
func (_EulerEVault *EulerEVaultCallerSession) UnitOfAccount() (common.Address, error) {
	return _EulerEVault.Contract.UnitOfAccount(&_EulerEVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.Contract.Approve(&_EulerEVault.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.Contract.Approve(&_EulerEVault.TransactOpts, spender, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "borrow", amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Borrow(&_EulerEVault.TransactOpts, amount, receiver)
}

// Borrow is a paid mutator transaction binding the contract method 0x4b3fd148.
//
// Solidity: function borrow(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactorSession) Borrow(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Borrow(&_EulerEVault.TransactOpts, amount, receiver)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEVault *EulerEVaultTransactor) CheckVaultStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "checkVaultStatus")
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEVault *EulerEVaultSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EulerEVault.Contract.CheckVaultStatus(&_EulerEVault.TransactOpts)
}

// CheckVaultStatus is a paid mutator transaction binding the contract method 0x4b3d1223.
//
// Solidity: function checkVaultStatus() returns(bytes4)
func (_EulerEVault *EulerEVaultTransactorSession) CheckVaultStatus() (*types.Transaction, error) {
	return _EulerEVault.Contract.CheckVaultStatus(&_EulerEVault.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEVault *EulerEVaultTransactor) ConvertFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "convertFees")
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEVault *EulerEVaultSession) ConvertFees() (*types.Transaction, error) {
	return _EulerEVault.Contract.ConvertFees(&_EulerEVault.TransactOpts)
}

// ConvertFees is a paid mutator transaction binding the contract method 0x2b5335c3.
//
// Solidity: function convertFees() returns()
func (_EulerEVault *EulerEVaultTransactorSession) ConvertFees() (*types.Transaction, error) {
	return _EulerEVault.Contract.ConvertFees(&_EulerEVault.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "deposit", amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Deposit(&_EulerEVault.TransactOpts, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactorSession) Deposit(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Deposit(&_EulerEVault.TransactOpts, amount, receiver)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEVault *EulerEVaultTransactor) DisableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "disableBalanceForwarder")
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEVault *EulerEVaultSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEVault.Contract.DisableBalanceForwarder(&_EulerEVault.TransactOpts)
}

// DisableBalanceForwarder is a paid mutator transaction binding the contract method 0x41233a98.
//
// Solidity: function disableBalanceForwarder() returns()
func (_EulerEVault *EulerEVaultTransactorSession) DisableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEVault.Contract.DisableBalanceForwarder(&_EulerEVault.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEVault *EulerEVaultTransactor) DisableController(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "disableController")
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEVault *EulerEVaultSession) DisableController() (*types.Transaction, error) {
	return _EulerEVault.Contract.DisableController(&_EulerEVault.TransactOpts)
}

// DisableController is a paid mutator transaction binding the contract method 0x869e50c7.
//
// Solidity: function disableController() returns()
func (_EulerEVault *EulerEVaultTransactorSession) DisableController() (*types.Transaction, error) {
	return _EulerEVault.Contract.DisableController(&_EulerEVault.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEVault *EulerEVaultTransactor) EnableBalanceForwarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "enableBalanceForwarder")
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEVault *EulerEVaultSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEVault.Contract.EnableBalanceForwarder(&_EulerEVault.TransactOpts)
}

// EnableBalanceForwarder is a paid mutator transaction binding the contract method 0x64b1cdd6.
//
// Solidity: function enableBalanceForwarder() returns()
func (_EulerEVault *EulerEVaultTransactorSession) EnableBalanceForwarder() (*types.Transaction, error) {
	return _EulerEVault.Contract.EnableBalanceForwarder(&_EulerEVault.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEVault *EulerEVaultTransactor) FlashLoan(opts *bind.TransactOpts, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "flashLoan", amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEVault *EulerEVaultSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEVault.Contract.FlashLoan(&_EulerEVault.TransactOpts, amount, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5296a431.
//
// Solidity: function flashLoan(uint256 amount, bytes data) returns()
func (_EulerEVault *EulerEVaultTransactorSession) FlashLoan(amount *big.Int, data []byte) (*types.Transaction, error) {
	return _EulerEVault.Contract.FlashLoan(&_EulerEVault.TransactOpts, amount, data)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEVault *EulerEVaultTransactor) Initialize(opts *bind.TransactOpts, proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "initialize", proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEVault *EulerEVaultSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Initialize(&_EulerEVault.TransactOpts, proxyCreator)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address proxyCreator) returns()
func (_EulerEVault *EulerEVaultTransactorSession) Initialize(proxyCreator common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Initialize(&_EulerEVault.TransactOpts, proxyCreator)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEVault *EulerEVaultTransactor) Liquidate(opts *bind.TransactOpts, violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "liquidate", violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEVault *EulerEVaultSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEVault.Contract.Liquidate(&_EulerEVault.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Liquidate is a paid mutator transaction binding the contract method 0xc1342574.
//
// Solidity: function liquidate(address violator, address collateral, uint256 repayAssets, uint256 minYieldBalance) returns()
func (_EulerEVault *EulerEVaultTransactorSession) Liquidate(violator common.Address, collateral common.Address, repayAssets *big.Int, minYieldBalance *big.Int) (*types.Transaction, error) {
	return _EulerEVault.Contract.Liquidate(&_EulerEVault.TransactOpts, violator, collateral, repayAssets, minYieldBalance)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactor) Mint(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "mint", amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Mint(&_EulerEVault.TransactOpts, amount, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactorSession) Mint(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Mint(&_EulerEVault.TransactOpts, amount, receiver)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEVault *EulerEVaultTransactor) PullDebt(opts *bind.TransactOpts, amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "pullDebt", amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEVault *EulerEVaultSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.PullDebt(&_EulerEVault.TransactOpts, amount, from)
}

// PullDebt is a paid mutator transaction binding the contract method 0xaebde56b.
//
// Solidity: function pullDebt(uint256 amount, address from) returns()
func (_EulerEVault *EulerEVaultTransactorSession) PullDebt(amount *big.Int, from common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.PullDebt(&_EulerEVault.TransactOpts, amount, from)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEVault *EulerEVaultTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "redeem", amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEVault *EulerEVaultSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Redeem(&_EulerEVault.TransactOpts, amount, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEVault *EulerEVaultTransactorSession) Redeem(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Redeem(&_EulerEVault.TransactOpts, amount, receiver, owner)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactor) Repay(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "repay", amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Repay(&_EulerEVault.TransactOpts, amount, receiver)
}

// Repay is a paid mutator transaction binding the contract method 0xacb70815.
//
// Solidity: function repay(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactorSession) Repay(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Repay(&_EulerEVault.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEVault *EulerEVaultTransactor) RepayWithShares(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "repayWithShares", amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEVault *EulerEVaultSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.RepayWithShares(&_EulerEVault.TransactOpts, amount, receiver)
}

// RepayWithShares is a paid mutator transaction binding the contract method 0xa9c8eb7e.
//
// Solidity: function repayWithShares(uint256 amount, address receiver) returns(uint256 shares, uint256 debt)
func (_EulerEVault *EulerEVaultTransactorSession) RepayWithShares(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.RepayWithShares(&_EulerEVault.TransactOpts, amount, receiver)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEVault *EulerEVaultTransactor) SetCaps(opts *bind.TransactOpts, supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setCaps", supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEVault *EulerEVaultSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetCaps(&_EulerEVault.TransactOpts, supplyCap, borrowCap)
}

// SetCaps is a paid mutator transaction binding the contract method 0xd87f780f.
//
// Solidity: function setCaps(uint16 supplyCap, uint16 borrowCap) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetCaps(supplyCap uint16, borrowCap uint16) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetCaps(&_EulerEVault.TransactOpts, supplyCap, borrowCap)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEVault *EulerEVaultTransactor) SetConfigFlags(opts *bind.TransactOpts, newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setConfigFlags", newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEVault *EulerEVaultSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetConfigFlags(&_EulerEVault.TransactOpts, newConfigFlags)
}

// SetConfigFlags is a paid mutator transaction binding the contract method 0xada3d56f.
//
// Solidity: function setConfigFlags(uint32 newConfigFlags) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetConfigFlags(newConfigFlags uint32) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetConfigFlags(&_EulerEVault.TransactOpts, newConfigFlags)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEVault *EulerEVaultTransactor) SetFeeReceiver(opts *bind.TransactOpts, newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setFeeReceiver", newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEVault *EulerEVaultSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetFeeReceiver(&_EulerEVault.TransactOpts, newFeeReceiver)
}

// SetFeeReceiver is a paid mutator transaction binding the contract method 0xefdcd974.
//
// Solidity: function setFeeReceiver(address newFeeReceiver) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetFeeReceiver(newFeeReceiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetFeeReceiver(&_EulerEVault.TransactOpts, newFeeReceiver)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEVault *EulerEVaultTransactor) SetGovernorAdmin(opts *bind.TransactOpts, newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setGovernorAdmin", newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEVault *EulerEVaultSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetGovernorAdmin(&_EulerEVault.TransactOpts, newGovernorAdmin)
}

// SetGovernorAdmin is a paid mutator transaction binding the contract method 0x82ebd674.
//
// Solidity: function setGovernorAdmin(address newGovernorAdmin) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetGovernorAdmin(newGovernorAdmin common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetGovernorAdmin(&_EulerEVault.TransactOpts, newGovernorAdmin)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEVault *EulerEVaultTransactor) SetHookConfig(opts *bind.TransactOpts, newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setHookConfig", newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEVault *EulerEVaultSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetHookConfig(&_EulerEVault.TransactOpts, newHookTarget, newHookedOps)
}

// SetHookConfig is a paid mutator transaction binding the contract method 0xd1a3a308.
//
// Solidity: function setHookConfig(address newHookTarget, uint32 newHookedOps) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetHookConfig(newHookTarget common.Address, newHookedOps uint32) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetHookConfig(&_EulerEVault.TransactOpts, newHookTarget, newHookedOps)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEVault *EulerEVaultTransactor) SetInterestFee(opts *bind.TransactOpts, newFee uint16) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setInterestFee", newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEVault *EulerEVaultSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetInterestFee(&_EulerEVault.TransactOpts, newFee)
}

// SetInterestFee is a paid mutator transaction binding the contract method 0x60cb90ef.
//
// Solidity: function setInterestFee(uint16 newFee) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetInterestFee(newFee uint16) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetInterestFee(&_EulerEVault.TransactOpts, newFee)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEVault *EulerEVaultTransactor) SetInterestRateModel(opts *bind.TransactOpts, newModel common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setInterestRateModel", newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEVault *EulerEVaultSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetInterestRateModel(&_EulerEVault.TransactOpts, newModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0x8bcd4016.
//
// Solidity: function setInterestRateModel(address newModel) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetInterestRateModel(newModel common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetInterestRateModel(&_EulerEVault.TransactOpts, newModel)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEVault *EulerEVaultTransactor) SetLTV(opts *bind.TransactOpts, collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setLTV", collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEVault *EulerEVaultSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetLTV(&_EulerEVault.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLTV is a paid mutator transaction binding the contract method 0x4bca3d5b.
//
// Solidity: function setLTV(address collateral, uint16 borrowLTV, uint16 liquidationLTV, uint32 rampDuration) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetLTV(collateral common.Address, borrowLTV uint16, liquidationLTV uint16, rampDuration uint32) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetLTV(&_EulerEVault.TransactOpts, collateral, borrowLTV, liquidationLTV, rampDuration)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEVault *EulerEVaultTransactor) SetLiquidationCoolOffTime(opts *bind.TransactOpts, newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setLiquidationCoolOffTime", newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEVault *EulerEVaultSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetLiquidationCoolOffTime(&_EulerEVault.TransactOpts, newCoolOffTime)
}

// SetLiquidationCoolOffTime is a paid mutator transaction binding the contract method 0xaf06d3cf.
//
// Solidity: function setLiquidationCoolOffTime(uint16 newCoolOffTime) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetLiquidationCoolOffTime(newCoolOffTime uint16) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetLiquidationCoolOffTime(&_EulerEVault.TransactOpts, newCoolOffTime)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEVault *EulerEVaultTransactor) SetMaxLiquidationDiscount(opts *bind.TransactOpts, newDiscount uint16) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "setMaxLiquidationDiscount", newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEVault *EulerEVaultSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetMaxLiquidationDiscount(&_EulerEVault.TransactOpts, newDiscount)
}

// SetMaxLiquidationDiscount is a paid mutator transaction binding the contract method 0xb4113ba7.
//
// Solidity: function setMaxLiquidationDiscount(uint16 newDiscount) returns()
func (_EulerEVault *EulerEVaultTransactorSession) SetMaxLiquidationDiscount(newDiscount uint16) (*types.Transaction, error) {
	return _EulerEVault.Contract.SetMaxLiquidationDiscount(&_EulerEVault.TransactOpts, newDiscount)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactor) Skim(opts *bind.TransactOpts, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "skim", amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Skim(&_EulerEVault.TransactOpts, amount, receiver)
}

// Skim is a paid mutator transaction binding the contract method 0x8d56c639.
//
// Solidity: function skim(uint256 amount, address receiver) returns(uint256)
func (_EulerEVault *EulerEVaultTransactorSession) Skim(amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Skim(&_EulerEVault.TransactOpts, amount, receiver)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEVault *EulerEVaultTransactor) Touch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "touch")
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEVault *EulerEVaultSession) Touch() (*types.Transaction, error) {
	return _EulerEVault.Contract.Touch(&_EulerEVault.TransactOpts)
}

// Touch is a paid mutator transaction binding the contract method 0xa55526db.
//
// Solidity: function touch() returns()
func (_EulerEVault *EulerEVaultTransactorSession) Touch() (*types.Transaction, error) {
	return _EulerEVault.Contract.Touch(&_EulerEVault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.Contract.Transfer(&_EulerEVault.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.Contract.Transfer(&_EulerEVault.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.Contract.TransferFrom(&_EulerEVault.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_EulerEVault *EulerEVaultTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EulerEVault.Contract.TransferFrom(&_EulerEVault.TransactOpts, from, to, amount)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEVault *EulerEVaultTransactor) TransferFromMax(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "transferFromMax", from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEVault *EulerEVaultSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.TransferFromMax(&_EulerEVault.TransactOpts, from, to)
}

// TransferFromMax is a paid mutator transaction binding the contract method 0xcbfdd7e1.
//
// Solidity: function transferFromMax(address from, address to) returns(bool)
func (_EulerEVault *EulerEVaultTransactorSession) TransferFromMax(from common.Address, to common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.TransferFromMax(&_EulerEVault.TransactOpts, from, to)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEVault *EulerEVaultTransactor) ViewDelegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "viewDelegate")
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEVault *EulerEVaultSession) ViewDelegate() (*types.Transaction, error) {
	return _EulerEVault.Contract.ViewDelegate(&_EulerEVault.TransactOpts)
}

// ViewDelegate is a paid mutator transaction binding the contract method 0x1fe8b953.
//
// Solidity: function viewDelegate() payable returns()
func (_EulerEVault *EulerEVaultTransactorSession) ViewDelegate() (*types.Transaction, error) {
	return _EulerEVault.Contract.ViewDelegate(&_EulerEVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEVault *EulerEVaultTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEVault.contract.Transact(opts, "withdraw", amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEVault *EulerEVaultSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Withdraw(&_EulerEVault.TransactOpts, amount, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 amount, address receiver, address owner) returns(uint256)
func (_EulerEVault *EulerEVaultTransactorSession) Withdraw(amount *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _EulerEVault.Contract.Withdraw(&_EulerEVault.TransactOpts, amount, receiver, owner)
}

// EulerEVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EulerEVault contract.
type EulerEVaultApprovalIterator struct {
	Event *EulerEVaultApproval // Event containing the contract specifics and raw log

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
func (it *EulerEVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultApproval)
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
		it.Event = new(EulerEVaultApproval)
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
func (it *EulerEVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultApproval represents a Approval event raised by the EulerEVault contract.
type EulerEVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEVault *EulerEVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EulerEVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultApprovalIterator{contract: _EulerEVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EulerEVault *EulerEVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EulerEVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultApproval)
				if err := _EulerEVault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseApproval(log types.Log) (*EulerEVaultApproval, error) {
	event := new(EulerEVaultApproval)
	if err := _EulerEVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultBalanceForwarderStatusIterator is returned from FilterBalanceForwarderStatus and is used to iterate over the raw logs and unpacked data for BalanceForwarderStatus events raised by the EulerEVault contract.
type EulerEVaultBalanceForwarderStatusIterator struct {
	Event *EulerEVaultBalanceForwarderStatus // Event containing the contract specifics and raw log

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
func (it *EulerEVaultBalanceForwarderStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultBalanceForwarderStatus)
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
		it.Event = new(EulerEVaultBalanceForwarderStatus)
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
func (it *EulerEVaultBalanceForwarderStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultBalanceForwarderStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultBalanceForwarderStatus represents a BalanceForwarderStatus event raised by the EulerEVault contract.
type EulerEVaultBalanceForwarderStatus struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceForwarderStatus is a free log retrieval operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerEVault *EulerEVaultFilterer) FilterBalanceForwarderStatus(opts *bind.FilterOpts, account []common.Address) (*EulerEVaultBalanceForwarderStatusIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultBalanceForwarderStatusIterator{contract: _EulerEVault.contract, event: "BalanceForwarderStatus", logs: logs, sub: sub}, nil
}

// WatchBalanceForwarderStatus is a free log subscription operation binding the contract event 0xc3e011ddce6181dafb5798a536341c7c601913626c31d31744f91b77b7e2412d.
//
// Solidity: event BalanceForwarderStatus(address indexed account, bool status)
func (_EulerEVault *EulerEVaultFilterer) WatchBalanceForwarderStatus(opts *bind.WatchOpts, sink chan<- *EulerEVaultBalanceForwarderStatus, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "BalanceForwarderStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultBalanceForwarderStatus)
				if err := _EulerEVault.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseBalanceForwarderStatus(log types.Log) (*EulerEVaultBalanceForwarderStatus, error) {
	event := new(EulerEVaultBalanceForwarderStatus)
	if err := _EulerEVault.contract.UnpackLog(event, "BalanceForwarderStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the EulerEVault contract.
type EulerEVaultBorrowIterator struct {
	Event *EulerEVaultBorrow // Event containing the contract specifics and raw log

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
func (it *EulerEVaultBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultBorrow)
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
		it.Event = new(EulerEVaultBorrow)
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
func (it *EulerEVaultBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultBorrow represents a Borrow event raised by the EulerEVault contract.
type EulerEVaultBorrow struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) FilterBorrow(opts *bind.FilterOpts, account []common.Address) (*EulerEVaultBorrowIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultBorrowIterator{contract: _EulerEVault.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcbc04eca7e9da35cb1393a6135a199ca52e450d5e9251cbd99f7847d33a36750.
//
// Solidity: event Borrow(address indexed account, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *EulerEVaultBorrow, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "Borrow", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultBorrow)
				if err := _EulerEVault.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseBorrow(log types.Log) (*EulerEVaultBorrow, error) {
	event := new(EulerEVaultBorrow)
	if err := _EulerEVault.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultConvertFeesIterator is returned from FilterConvertFees and is used to iterate over the raw logs and unpacked data for ConvertFees events raised by the EulerEVault contract.
type EulerEVaultConvertFeesIterator struct {
	Event *EulerEVaultConvertFees // Event containing the contract specifics and raw log

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
func (it *EulerEVaultConvertFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultConvertFees)
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
		it.Event = new(EulerEVaultConvertFees)
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
func (it *EulerEVaultConvertFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultConvertFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultConvertFees represents a ConvertFees event raised by the EulerEVault contract.
type EulerEVaultConvertFees struct {
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
func (_EulerEVault *EulerEVaultFilterer) FilterConvertFees(opts *bind.FilterOpts, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (*EulerEVaultConvertFeesIterator, error) {

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

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultConvertFeesIterator{contract: _EulerEVault.contract, event: "ConvertFees", logs: logs, sub: sub}, nil
}

// WatchConvertFees is a free log subscription operation binding the contract event 0x4e16b07cac5fe5604af487e07b1b62efc8bd47477b18839f4688d2cae957f965.
//
// Solidity: event ConvertFees(address indexed sender, address indexed protocolReceiver, address indexed governorReceiver, uint256 protocolShares, uint256 governorShares)
func (_EulerEVault *EulerEVaultFilterer) WatchConvertFees(opts *bind.WatchOpts, sink chan<- *EulerEVaultConvertFees, sender []common.Address, protocolReceiver []common.Address, governorReceiver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "ConvertFees", senderRule, protocolReceiverRule, governorReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultConvertFees)
				if err := _EulerEVault.contract.UnpackLog(event, "ConvertFees", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseConvertFees(log types.Log) (*EulerEVaultConvertFees, error) {
	event := new(EulerEVaultConvertFees)
	if err := _EulerEVault.contract.UnpackLog(event, "ConvertFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultDebtSocializedIterator is returned from FilterDebtSocialized and is used to iterate over the raw logs and unpacked data for DebtSocialized events raised by the EulerEVault contract.
type EulerEVaultDebtSocializedIterator struct {
	Event *EulerEVaultDebtSocialized // Event containing the contract specifics and raw log

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
func (it *EulerEVaultDebtSocializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultDebtSocialized)
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
		it.Event = new(EulerEVaultDebtSocialized)
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
func (it *EulerEVaultDebtSocializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultDebtSocializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultDebtSocialized represents a DebtSocialized event raised by the EulerEVault contract.
type EulerEVaultDebtSocialized struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebtSocialized is a free log retrieval operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) FilterDebtSocialized(opts *bind.FilterOpts, account []common.Address) (*EulerEVaultDebtSocializedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultDebtSocializedIterator{contract: _EulerEVault.contract, event: "DebtSocialized", logs: logs, sub: sub}, nil
}

// WatchDebtSocialized is a free log subscription operation binding the contract event 0xe786d0bc2e83bf230ed9895a9c4d7756ab0c6e22eb8a4ff69c161ece76bd36df.
//
// Solidity: event DebtSocialized(address indexed account, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) WatchDebtSocialized(opts *bind.WatchOpts, sink chan<- *EulerEVaultDebtSocialized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "DebtSocialized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultDebtSocialized)
				if err := _EulerEVault.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseDebtSocialized(log types.Log) (*EulerEVaultDebtSocialized, error) {
	event := new(EulerEVaultDebtSocialized)
	if err := _EulerEVault.contract.UnpackLog(event, "DebtSocialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EulerEVault contract.
type EulerEVaultDepositIterator struct {
	Event *EulerEVaultDeposit // Event containing the contract specifics and raw log

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
func (it *EulerEVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultDeposit)
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
		it.Event = new(EulerEVaultDeposit)
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
func (it *EulerEVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultDeposit represents a Deposit event raised by the EulerEVault contract.
type EulerEVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEVault *EulerEVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*EulerEVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultDepositIterator{contract: _EulerEVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEVault *EulerEVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EulerEVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultDeposit)
				if err := _EulerEVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseDeposit(log types.Log) (*EulerEVaultDeposit, error) {
	event := new(EulerEVaultDeposit)
	if err := _EulerEVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultEVaultCreatedIterator is returned from FilterEVaultCreated and is used to iterate over the raw logs and unpacked data for EVaultCreated events raised by the EulerEVault contract.
type EulerEVaultEVaultCreatedIterator struct {
	Event *EulerEVaultEVaultCreated // Event containing the contract specifics and raw log

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
func (it *EulerEVaultEVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultEVaultCreated)
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
		it.Event = new(EulerEVaultEVaultCreated)
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
func (it *EulerEVaultEVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultEVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultEVaultCreated represents a EVaultCreated event raised by the EulerEVault contract.
type EulerEVaultEVaultCreated struct {
	Creator common.Address
	Asset   common.Address
	DToken  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEVaultCreated is a free log retrieval operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerEVault *EulerEVaultFilterer) FilterEVaultCreated(opts *bind.FilterOpts, creator []common.Address, asset []common.Address) (*EulerEVaultEVaultCreatedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultEVaultCreatedIterator{contract: _EulerEVault.contract, event: "EVaultCreated", logs: logs, sub: sub}, nil
}

// WatchEVaultCreated is a free log subscription operation binding the contract event 0x0cd345140b9008a43f99a999a328ece572a0193e8c8bf5f5755585e6f293b85e.
//
// Solidity: event EVaultCreated(address indexed creator, address indexed asset, address dToken)
func (_EulerEVault *EulerEVaultFilterer) WatchEVaultCreated(opts *bind.WatchOpts, sink chan<- *EulerEVaultEVaultCreated, creator []common.Address, asset []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "EVaultCreated", creatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultEVaultCreated)
				if err := _EulerEVault.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseEVaultCreated(log types.Log) (*EulerEVaultEVaultCreated, error) {
	event := new(EulerEVaultEVaultCreated)
	if err := _EulerEVault.contract.UnpackLog(event, "EVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetCapsIterator is returned from FilterGovSetCaps and is used to iterate over the raw logs and unpacked data for GovSetCaps events raised by the EulerEVault contract.
type EulerEVaultGovSetCapsIterator struct {
	Event *EulerEVaultGovSetCaps // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetCapsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetCaps)
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
		it.Event = new(EulerEVaultGovSetCaps)
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
func (it *EulerEVaultGovSetCapsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetCapsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetCaps represents a GovSetCaps event raised by the EulerEVault contract.
type EulerEVaultGovSetCaps struct {
	NewSupplyCap uint16
	NewBorrowCap uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterGovSetCaps is a free log retrieval operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetCaps(opts *bind.FilterOpts) (*EulerEVaultGovSetCapsIterator, error) {

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetCapsIterator{contract: _EulerEVault.contract, event: "GovSetCaps", logs: logs, sub: sub}, nil
}

// WatchGovSetCaps is a free log subscription operation binding the contract event 0xadbdcd178dfddc478805a3703b6cf3b72ca5e78ecebacffe1aad03188cc1cbf4.
//
// Solidity: event GovSetCaps(uint16 newSupplyCap, uint16 newBorrowCap)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetCaps(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetCaps) (event.Subscription, error) {

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetCaps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetCaps)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetCaps(log types.Log) (*EulerEVaultGovSetCaps, error) {
	event := new(EulerEVaultGovSetCaps)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetCaps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetConfigFlagsIterator is returned from FilterGovSetConfigFlags and is used to iterate over the raw logs and unpacked data for GovSetConfigFlags events raised by the EulerEVault contract.
type EulerEVaultGovSetConfigFlagsIterator struct {
	Event *EulerEVaultGovSetConfigFlags // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetConfigFlagsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetConfigFlags)
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
		it.Event = new(EulerEVaultGovSetConfigFlags)
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
func (it *EulerEVaultGovSetConfigFlagsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetConfigFlagsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetConfigFlags represents a GovSetConfigFlags event raised by the EulerEVault contract.
type EulerEVaultGovSetConfigFlags struct {
	NewConfigFlags uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetConfigFlags is a free log retrieval operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetConfigFlags(opts *bind.FilterOpts) (*EulerEVaultGovSetConfigFlagsIterator, error) {

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetConfigFlagsIterator{contract: _EulerEVault.contract, event: "GovSetConfigFlags", logs: logs, sub: sub}, nil
}

// WatchGovSetConfigFlags is a free log subscription operation binding the contract event 0xe7f84c52c0ef295afe77de8cb30516d6f28d50306f979b45776dd1b619ae5ffc.
//
// Solidity: event GovSetConfigFlags(uint32 newConfigFlags)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetConfigFlags(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetConfigFlags) (event.Subscription, error) {

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetConfigFlags")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetConfigFlags)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetConfigFlags(log types.Log) (*EulerEVaultGovSetConfigFlags, error) {
	event := new(EulerEVaultGovSetConfigFlags)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetConfigFlags", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetFeeReceiverIterator is returned from FilterGovSetFeeReceiver and is used to iterate over the raw logs and unpacked data for GovSetFeeReceiver events raised by the EulerEVault contract.
type EulerEVaultGovSetFeeReceiverIterator struct {
	Event *EulerEVaultGovSetFeeReceiver // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetFeeReceiverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetFeeReceiver)
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
		it.Event = new(EulerEVaultGovSetFeeReceiver)
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
func (it *EulerEVaultGovSetFeeReceiverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetFeeReceiverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetFeeReceiver represents a GovSetFeeReceiver event raised by the EulerEVault contract.
type EulerEVaultGovSetFeeReceiver struct {
	NewFeeReceiver common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetFeeReceiver is a free log retrieval operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetFeeReceiver(opts *bind.FilterOpts, newFeeReceiver []common.Address) (*EulerEVaultGovSetFeeReceiverIterator, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetFeeReceiverIterator{contract: _EulerEVault.contract, event: "GovSetFeeReceiver", logs: logs, sub: sub}, nil
}

// WatchGovSetFeeReceiver is a free log subscription operation binding the contract event 0x836a1afef2bc89de2cb4713cc8d312fccf2ff835230721c5f41f13374707413a.
//
// Solidity: event GovSetFeeReceiver(address indexed newFeeReceiver)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetFeeReceiver(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetFeeReceiver, newFeeReceiver []common.Address) (event.Subscription, error) {

	var newFeeReceiverRule []interface{}
	for _, newFeeReceiverItem := range newFeeReceiver {
		newFeeReceiverRule = append(newFeeReceiverRule, newFeeReceiverItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetFeeReceiver", newFeeReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetFeeReceiver)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetFeeReceiver(log types.Log) (*EulerEVaultGovSetFeeReceiver, error) {
	event := new(EulerEVaultGovSetFeeReceiver)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetFeeReceiver", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetGovernorAdminIterator is returned from FilterGovSetGovernorAdmin and is used to iterate over the raw logs and unpacked data for GovSetGovernorAdmin events raised by the EulerEVault contract.
type EulerEVaultGovSetGovernorAdminIterator struct {
	Event *EulerEVaultGovSetGovernorAdmin // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetGovernorAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetGovernorAdmin)
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
		it.Event = new(EulerEVaultGovSetGovernorAdmin)
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
func (it *EulerEVaultGovSetGovernorAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetGovernorAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetGovernorAdmin represents a GovSetGovernorAdmin event raised by the EulerEVault contract.
type EulerEVaultGovSetGovernorAdmin struct {
	NewGovernorAdmin common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGovSetGovernorAdmin is a free log retrieval operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetGovernorAdmin(opts *bind.FilterOpts, newGovernorAdmin []common.Address) (*EulerEVaultGovSetGovernorAdminIterator, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetGovernorAdminIterator{contract: _EulerEVault.contract, event: "GovSetGovernorAdmin", logs: logs, sub: sub}, nil
}

// WatchGovSetGovernorAdmin is a free log subscription operation binding the contract event 0x1c145a4cd16d4148579b0f2296884ac4aa47536e4ef10a32e1cdc0dc3dd20ea4.
//
// Solidity: event GovSetGovernorAdmin(address indexed newGovernorAdmin)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetGovernorAdmin(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetGovernorAdmin, newGovernorAdmin []common.Address) (event.Subscription, error) {

	var newGovernorAdminRule []interface{}
	for _, newGovernorAdminItem := range newGovernorAdmin {
		newGovernorAdminRule = append(newGovernorAdminRule, newGovernorAdminItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetGovernorAdmin", newGovernorAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetGovernorAdmin)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetGovernorAdmin(log types.Log) (*EulerEVaultGovSetGovernorAdmin, error) {
	event := new(EulerEVaultGovSetGovernorAdmin)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetGovernorAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetHookConfigIterator is returned from FilterGovSetHookConfig and is used to iterate over the raw logs and unpacked data for GovSetHookConfig events raised by the EulerEVault contract.
type EulerEVaultGovSetHookConfigIterator struct {
	Event *EulerEVaultGovSetHookConfig // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetHookConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetHookConfig)
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
		it.Event = new(EulerEVaultGovSetHookConfig)
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
func (it *EulerEVaultGovSetHookConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetHookConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetHookConfig represents a GovSetHookConfig event raised by the EulerEVault contract.
type EulerEVaultGovSetHookConfig struct {
	NewHookTarget common.Address
	NewHookedOps  uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGovSetHookConfig is a free log retrieval operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetHookConfig(opts *bind.FilterOpts, newHookTarget []common.Address) (*EulerEVaultGovSetHookConfigIterator, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetHookConfigIterator{contract: _EulerEVault.contract, event: "GovSetHookConfig", logs: logs, sub: sub}, nil
}

// WatchGovSetHookConfig is a free log subscription operation binding the contract event 0xabadffb695acdb6863cd1324a91e5c359712b9110a55f9103774e2fb67dedb6a.
//
// Solidity: event GovSetHookConfig(address indexed newHookTarget, uint32 newHookedOps)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetHookConfig(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetHookConfig, newHookTarget []common.Address) (event.Subscription, error) {

	var newHookTargetRule []interface{}
	for _, newHookTargetItem := range newHookTarget {
		newHookTargetRule = append(newHookTargetRule, newHookTargetItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetHookConfig", newHookTargetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetHookConfig)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetHookConfig(log types.Log) (*EulerEVaultGovSetHookConfig, error) {
	event := new(EulerEVaultGovSetHookConfig)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetHookConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetInterestFeeIterator is returned from FilterGovSetInterestFee and is used to iterate over the raw logs and unpacked data for GovSetInterestFee events raised by the EulerEVault contract.
type EulerEVaultGovSetInterestFeeIterator struct {
	Event *EulerEVaultGovSetInterestFee // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetInterestFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetInterestFee)
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
		it.Event = new(EulerEVaultGovSetInterestFee)
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
func (it *EulerEVaultGovSetInterestFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetInterestFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetInterestFee represents a GovSetInterestFee event raised by the EulerEVault contract.
type EulerEVaultGovSetInterestFee struct {
	NewFee uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestFee is a free log retrieval operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetInterestFee(opts *bind.FilterOpts) (*EulerEVaultGovSetInterestFeeIterator, error) {

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetInterestFeeIterator{contract: _EulerEVault.contract, event: "GovSetInterestFee", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestFee is a free log subscription operation binding the contract event 0x634a58674e370383703eff32d9d4e4b3d1add94d50e8bcb631b04995d8e47341.
//
// Solidity: event GovSetInterestFee(uint16 newFee)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetInterestFee(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetInterestFee) (event.Subscription, error) {

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetInterestFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetInterestFee)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetInterestFee(log types.Log) (*EulerEVaultGovSetInterestFee, error) {
	event := new(EulerEVaultGovSetInterestFee)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetInterestFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetInterestRateModelIterator is returned from FilterGovSetInterestRateModel and is used to iterate over the raw logs and unpacked data for GovSetInterestRateModel events raised by the EulerEVault contract.
type EulerEVaultGovSetInterestRateModelIterator struct {
	Event *EulerEVaultGovSetInterestRateModel // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetInterestRateModel)
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
		it.Event = new(EulerEVaultGovSetInterestRateModel)
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
func (it *EulerEVaultGovSetInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetInterestRateModel represents a GovSetInterestRateModel event raised by the EulerEVault contract.
type EulerEVaultGovSetInterestRateModel struct {
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGovSetInterestRateModel is a free log retrieval operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetInterestRateModel(opts *bind.FilterOpts) (*EulerEVaultGovSetInterestRateModelIterator, error) {

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetInterestRateModelIterator{contract: _EulerEVault.contract, event: "GovSetInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchGovSetInterestRateModel is a free log subscription operation binding the contract event 0xe5f2a795fc5f8baf1b05659293834c88859298226d87422c88624b4c9f4d3a43.
//
// Solidity: event GovSetInterestRateModel(address newInterestRateModel)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetInterestRateModel(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetInterestRateModel)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetInterestRateModel(log types.Log) (*EulerEVaultGovSetInterestRateModel, error) {
	event := new(EulerEVaultGovSetInterestRateModel)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetLTVIterator is returned from FilterGovSetLTV and is used to iterate over the raw logs and unpacked data for GovSetLTV events raised by the EulerEVault contract.
type EulerEVaultGovSetLTVIterator struct {
	Event *EulerEVaultGovSetLTV // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetLTVIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetLTV)
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
		it.Event = new(EulerEVaultGovSetLTV)
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
func (it *EulerEVaultGovSetLTVIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetLTVIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetLTV represents a GovSetLTV event raised by the EulerEVault contract.
type EulerEVaultGovSetLTV struct {
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
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetLTV(opts *bind.FilterOpts, collateral []common.Address) (*EulerEVaultGovSetLTVIterator, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetLTVIterator{contract: _EulerEVault.contract, event: "GovSetLTV", logs: logs, sub: sub}, nil
}

// WatchGovSetLTV is a free log subscription operation binding the contract event 0xc69392046c26324e9eee913208811542aabcbde6a41ce9ee3b45473b18eb3c76.
//
// Solidity: event GovSetLTV(address indexed collateral, uint16 borrowLTV, uint16 liquidationLTV, uint16 initialLiquidationLTV, uint48 targetTimestamp, uint32 rampDuration)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetLTV(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetLTV, collateral []common.Address) (event.Subscription, error) {

	var collateralRule []interface{}
	for _, collateralItem := range collateral {
		collateralRule = append(collateralRule, collateralItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetLTV", collateralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetLTV)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetLTV(log types.Log) (*EulerEVaultGovSetLTV, error) {
	event := new(EulerEVaultGovSetLTV)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetLTV", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetLiquidationCoolOffTimeIterator is returned from FilterGovSetLiquidationCoolOffTime and is used to iterate over the raw logs and unpacked data for GovSetLiquidationCoolOffTime events raised by the EulerEVault contract.
type EulerEVaultGovSetLiquidationCoolOffTimeIterator struct {
	Event *EulerEVaultGovSetLiquidationCoolOffTime // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetLiquidationCoolOffTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetLiquidationCoolOffTime)
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
		it.Event = new(EulerEVaultGovSetLiquidationCoolOffTime)
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
func (it *EulerEVaultGovSetLiquidationCoolOffTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetLiquidationCoolOffTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetLiquidationCoolOffTime represents a GovSetLiquidationCoolOffTime event raised by the EulerEVault contract.
type EulerEVaultGovSetLiquidationCoolOffTime struct {
	NewCoolOffTime uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovSetLiquidationCoolOffTime is a free log retrieval operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetLiquidationCoolOffTime(opts *bind.FilterOpts) (*EulerEVaultGovSetLiquidationCoolOffTimeIterator, error) {

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetLiquidationCoolOffTimeIterator{contract: _EulerEVault.contract, event: "GovSetLiquidationCoolOffTime", logs: logs, sub: sub}, nil
}

// WatchGovSetLiquidationCoolOffTime is a free log subscription operation binding the contract event 0xdf4edc1d288e7b3306b287d03fd77b2070b8b308c702bf7297f72d928175dfa5.
//
// Solidity: event GovSetLiquidationCoolOffTime(uint16 newCoolOffTime)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetLiquidationCoolOffTime(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetLiquidationCoolOffTime) (event.Subscription, error) {

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetLiquidationCoolOffTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetLiquidationCoolOffTime)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetLiquidationCoolOffTime(log types.Log) (*EulerEVaultGovSetLiquidationCoolOffTime, error) {
	event := new(EulerEVaultGovSetLiquidationCoolOffTime)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetLiquidationCoolOffTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultGovSetMaxLiquidationDiscountIterator is returned from FilterGovSetMaxLiquidationDiscount and is used to iterate over the raw logs and unpacked data for GovSetMaxLiquidationDiscount events raised by the EulerEVault contract.
type EulerEVaultGovSetMaxLiquidationDiscountIterator struct {
	Event *EulerEVaultGovSetMaxLiquidationDiscount // Event containing the contract specifics and raw log

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
func (it *EulerEVaultGovSetMaxLiquidationDiscountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultGovSetMaxLiquidationDiscount)
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
		it.Event = new(EulerEVaultGovSetMaxLiquidationDiscount)
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
func (it *EulerEVaultGovSetMaxLiquidationDiscountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultGovSetMaxLiquidationDiscountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultGovSetMaxLiquidationDiscount represents a GovSetMaxLiquidationDiscount event raised by the EulerEVault contract.
type EulerEVaultGovSetMaxLiquidationDiscount struct {
	NewDiscount uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGovSetMaxLiquidationDiscount is a free log retrieval operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerEVault *EulerEVaultFilterer) FilterGovSetMaxLiquidationDiscount(opts *bind.FilterOpts) (*EulerEVaultGovSetMaxLiquidationDiscountIterator, error) {

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return &EulerEVaultGovSetMaxLiquidationDiscountIterator{contract: _EulerEVault.contract, event: "GovSetMaxLiquidationDiscount", logs: logs, sub: sub}, nil
}

// WatchGovSetMaxLiquidationDiscount is a free log subscription operation binding the contract event 0x558a63d245d08220a137de3573129d3921e70e806adccf3a068c4723b9b3322d.
//
// Solidity: event GovSetMaxLiquidationDiscount(uint16 newDiscount)
func (_EulerEVault *EulerEVaultFilterer) WatchGovSetMaxLiquidationDiscount(opts *bind.WatchOpts, sink chan<- *EulerEVaultGovSetMaxLiquidationDiscount) (event.Subscription, error) {

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "GovSetMaxLiquidationDiscount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultGovSetMaxLiquidationDiscount)
				if err := _EulerEVault.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseGovSetMaxLiquidationDiscount(log types.Log) (*EulerEVaultGovSetMaxLiquidationDiscount, error) {
	event := new(EulerEVaultGovSetMaxLiquidationDiscount)
	if err := _EulerEVault.contract.UnpackLog(event, "GovSetMaxLiquidationDiscount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultInterestAccruedIterator is returned from FilterInterestAccrued and is used to iterate over the raw logs and unpacked data for InterestAccrued events raised by the EulerEVault contract.
type EulerEVaultInterestAccruedIterator struct {
	Event *EulerEVaultInterestAccrued // Event containing the contract specifics and raw log

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
func (it *EulerEVaultInterestAccruedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultInterestAccrued)
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
		it.Event = new(EulerEVaultInterestAccrued)
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
func (it *EulerEVaultInterestAccruedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultInterestAccruedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultInterestAccrued represents a InterestAccrued event raised by the EulerEVault contract.
type EulerEVaultInterestAccrued struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInterestAccrued is a free log retrieval operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) FilterInterestAccrued(opts *bind.FilterOpts, account []common.Address) (*EulerEVaultInterestAccruedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultInterestAccruedIterator{contract: _EulerEVault.contract, event: "InterestAccrued", logs: logs, sub: sub}, nil
}

// WatchInterestAccrued is a free log subscription operation binding the contract event 0x5e804d42ae3b860f881d11cb44a4bb1f2f0d5b3d081f5539a32d6f97b629d978.
//
// Solidity: event InterestAccrued(address indexed account, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) WatchInterestAccrued(opts *bind.WatchOpts, sink chan<- *EulerEVaultInterestAccrued, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "InterestAccrued", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultInterestAccrued)
				if err := _EulerEVault.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseInterestAccrued(log types.Log) (*EulerEVaultInterestAccrued, error) {
	event := new(EulerEVaultInterestAccrued)
	if err := _EulerEVault.contract.UnpackLog(event, "InterestAccrued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the EulerEVault contract.
type EulerEVaultLiquidateIterator struct {
	Event *EulerEVaultLiquidate // Event containing the contract specifics and raw log

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
func (it *EulerEVaultLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultLiquidate)
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
		it.Event = new(EulerEVaultLiquidate)
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
func (it *EulerEVaultLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultLiquidate represents a Liquidate event raised by the EulerEVault contract.
type EulerEVaultLiquidate struct {
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
func (_EulerEVault *EulerEVaultFilterer) FilterLiquidate(opts *bind.FilterOpts, liquidator []common.Address, violator []common.Address) (*EulerEVaultLiquidateIterator, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultLiquidateIterator{contract: _EulerEVault.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0x8246cc71ab01533b5bebc672a636df812f10637ad720797319d5741d5ebb3962.
//
// Solidity: event Liquidate(address indexed liquidator, address indexed violator, address collateral, uint256 repayAssets, uint256 yieldBalance)
func (_EulerEVault *EulerEVaultFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *EulerEVaultLiquidate, liquidator []common.Address, violator []common.Address) (event.Subscription, error) {

	var liquidatorRule []interface{}
	for _, liquidatorItem := range liquidator {
		liquidatorRule = append(liquidatorRule, liquidatorItem)
	}
	var violatorRule []interface{}
	for _, violatorItem := range violator {
		violatorRule = append(violatorRule, violatorItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "Liquidate", liquidatorRule, violatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultLiquidate)
				if err := _EulerEVault.contract.UnpackLog(event, "Liquidate", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseLiquidate(log types.Log) (*EulerEVaultLiquidate, error) {
	event := new(EulerEVaultLiquidate)
	if err := _EulerEVault.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultPullDebtIterator is returned from FilterPullDebt and is used to iterate over the raw logs and unpacked data for PullDebt events raised by the EulerEVault contract.
type EulerEVaultPullDebtIterator struct {
	Event *EulerEVaultPullDebt // Event containing the contract specifics and raw log

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
func (it *EulerEVaultPullDebtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultPullDebt)
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
		it.Event = new(EulerEVaultPullDebt)
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
func (it *EulerEVaultPullDebtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultPullDebtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultPullDebt represents a PullDebt event raised by the EulerEVault contract.
type EulerEVaultPullDebt struct {
	From   common.Address
	To     common.Address
	Assets *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPullDebt is a free log retrieval operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) FilterPullDebt(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerEVaultPullDebtIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultPullDebtIterator{contract: _EulerEVault.contract, event: "PullDebt", logs: logs, sub: sub}, nil
}

// WatchPullDebt is a free log subscription operation binding the contract event 0xe6d0bfd9025bf59969101a13cf02e3ba2811b533816c47d7155546c7c8a1048f.
//
// Solidity: event PullDebt(address indexed from, address indexed to, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) WatchPullDebt(opts *bind.WatchOpts, sink chan<- *EulerEVaultPullDebt, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "PullDebt", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultPullDebt)
				if err := _EulerEVault.contract.UnpackLog(event, "PullDebt", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParsePullDebt(log types.Log) (*EulerEVaultPullDebt, error) {
	event := new(EulerEVaultPullDebt)
	if err := _EulerEVault.contract.UnpackLog(event, "PullDebt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the EulerEVault contract.
type EulerEVaultRepayIterator struct {
	Event *EulerEVaultRepay // Event containing the contract specifics and raw log

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
func (it *EulerEVaultRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultRepay)
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
		it.Event = new(EulerEVaultRepay)
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
func (it *EulerEVaultRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultRepay represents a Repay event raised by the EulerEVault contract.
type EulerEVaultRepay struct {
	Account common.Address
	Assets  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) FilterRepay(opts *bind.FilterOpts, account []common.Address) (*EulerEVaultRepayIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultRepayIterator{contract: _EulerEVault.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x5c16de4f8b59bd9caf0f49a545f25819a895ed223294290b408242e72a594231.
//
// Solidity: event Repay(address indexed account, uint256 assets)
func (_EulerEVault *EulerEVaultFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *EulerEVaultRepay, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "Repay", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultRepay)
				if err := _EulerEVault.contract.UnpackLog(event, "Repay", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseRepay(log types.Log) (*EulerEVaultRepay, error) {
	event := new(EulerEVaultRepay)
	if err := _EulerEVault.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EulerEVault contract.
type EulerEVaultTransferIterator struct {
	Event *EulerEVaultTransfer // Event containing the contract specifics and raw log

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
func (it *EulerEVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultTransfer)
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
		it.Event = new(EulerEVaultTransfer)
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
func (it *EulerEVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultTransfer represents a Transfer event raised by the EulerEVault contract.
type EulerEVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEVault *EulerEVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EulerEVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultTransferIterator{contract: _EulerEVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EulerEVault *EulerEVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EulerEVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultTransfer)
				if err := _EulerEVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseTransfer(log types.Log) (*EulerEVaultTransfer, error) {
	event := new(EulerEVaultTransfer)
	if err := _EulerEVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultVaultStatusIterator is returned from FilterVaultStatus and is used to iterate over the raw logs and unpacked data for VaultStatus events raised by the EulerEVault contract.
type EulerEVaultVaultStatusIterator struct {
	Event *EulerEVaultVaultStatus // Event containing the contract specifics and raw log

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
func (it *EulerEVaultVaultStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultVaultStatus)
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
		it.Event = new(EulerEVaultVaultStatus)
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
func (it *EulerEVaultVaultStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultVaultStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultVaultStatus represents a VaultStatus event raised by the EulerEVault contract.
type EulerEVaultVaultStatus struct {
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
func (_EulerEVault *EulerEVaultFilterer) FilterVaultStatus(opts *bind.FilterOpts) (*EulerEVaultVaultStatusIterator, error) {

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return &EulerEVaultVaultStatusIterator{contract: _EulerEVault.contract, event: "VaultStatus", logs: logs, sub: sub}, nil
}

// WatchVaultStatus is a free log subscription operation binding the contract event 0x80b61abbfc5f73cfe5cf93cec97a69ed20643dc6c6f1833b05a1560aa164e24c.
//
// Solidity: event VaultStatus(uint256 totalShares, uint256 totalBorrows, uint256 accumulatedFees, uint256 cash, uint256 interestAccumulator, uint256 interestRate, uint256 timestamp)
func (_EulerEVault *EulerEVaultFilterer) WatchVaultStatus(opts *bind.WatchOpts, sink chan<- *EulerEVaultVaultStatus) (event.Subscription, error) {

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "VaultStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultVaultStatus)
				if err := _EulerEVault.contract.UnpackLog(event, "VaultStatus", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseVaultStatus(log types.Log) (*EulerEVaultVaultStatus, error) {
	event := new(EulerEVaultVaultStatus)
	if err := _EulerEVault.contract.UnpackLog(event, "VaultStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EulerEVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EulerEVault contract.
type EulerEVaultWithdrawIterator struct {
	Event *EulerEVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *EulerEVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EulerEVaultWithdraw)
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
		it.Event = new(EulerEVaultWithdraw)
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
func (it *EulerEVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EulerEVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EulerEVaultWithdraw represents a Withdraw event raised by the EulerEVault contract.
type EulerEVaultWithdraw struct {
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
func (_EulerEVault *EulerEVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*EulerEVaultWithdrawIterator, error) {

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

	logs, sub, err := _EulerEVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EulerEVaultWithdrawIterator{contract: _EulerEVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_EulerEVault *EulerEVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EulerEVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _EulerEVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EulerEVaultWithdraw)
				if err := _EulerEVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_EulerEVault *EulerEVaultFilterer) ParseWithdraw(log types.Log) (*EulerEVaultWithdraw, error) {
	event := new(EulerEVaultWithdraw)
	if err := _EulerEVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
