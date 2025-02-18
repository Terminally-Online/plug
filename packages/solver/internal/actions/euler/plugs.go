package euler

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/erc_20"
	"solver/bindings/euler_account_lens"
	"solver/bindings/euler_evault_implementation"
	"solver/bindings/euler_evc"
	"solver/bindings/plug_router"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/utils"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func HandleEarn(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal earn inputs: %v", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert earn amount to uint: %w", err)
	}

	vault, err := GetVault(inputs.Vault, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("Erc20")
	}

	subAccountAddress := GetSubAccountAddress(common.HexToAddress(params.From), inputs.SubAccountIndex)

	approveCalldata, err := erc20Abi.Pack(
		"approve",
		vault.Vault, // Who is the operator here, the evc or the vault?
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}
	depositCall, err := WrapEVCCall(
		params.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		depositCalldata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap deposit call: %w", err)
	}

	return []plug_router.PlugTypesLibPlug{{
		To:   *token,
		Data: approveCalldata,
	}, depositCall}, nil
}

func HandleDepositCollateral(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit collateral inputs: %v", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit collateral amount to uint: %w", err)
	}

	vault, err := GetVault(inputs.Vault, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("Erc20")
	}

	approveCalldata, err := erc20Abi.Pack(
		"approve",
		vault.Vault, // Who is the operator here, the evc or the vault?
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	evc, err := euler_evc.EulerEvcMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvc")
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	subAccountAddress := GetSubAccountAddress(common.HexToAddress(params.From), inputs.SubAccountIndex)

	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	depositCall, err := WrapEVCCall(
		params.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		depositCalldata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap deposit call: %w", err)
	}

	enableCollateralCalldata, err := evc.Pack(
		"enableCollateral",
		subAccountAddress,
		vault.Vault,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	enableCollateralCall, err := WrapEVCCall(
		params.ChainId,
		common.HexToAddress(references.Networks[params.ChainId].References["euler"]["evc"]),
		subAccountAddress,
		big.NewInt(0),
		enableCollateralCalldata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap enable collateral call: %w", err)
	}

	return []plug_router.PlugTypesLibPlug{{
		To:   *token,
		Data: approveCalldata,
	}, depositCall, enableCollateralCall}, nil
}

func HandleWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw inputs: %w", err)
	}

	_, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert withdraw amount to uint: %w", err)
	}

	vault, err := GetVault(inputs.Vault, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	subAccountAddress := GetSubAccountAddress(common.HexToAddress(params.From), inputs.SubAccountIndex)

	calldata, err := vaultAbi.Pack(
		"withdraw",
		amount,
		common.HexToAddress(params.From),
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	call, err := WrapEVCCall(
		params.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		calldata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap withdraw call: %w", err)
	}

	return []plug_router.PlugTypesLibPlug{call}, nil
}

func HandleBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal borrow inputs: %w", err)
	}

	_, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert borrow amount to uint: %w", err)
	}

	vault, err := GetVault(inputs.Vault, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	subAccountAddress := GetSubAccountAddress(common.HexToAddress(params.From), inputs.SubAccountIndex)

	calldata, err := vaultAbi.Pack(
		"borrow",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	call, err := WrapEVCCall(
		params.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		calldata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap borrow call: %w", err)
	}

	return []plug_router.PlugTypesLibPlug{call}, nil
}

func HandleRepay(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay inputs: %w", err)
	}

	_, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert repay amount to uint: %w", err)
	}

	vault, err := GetVault(inputs.Vault, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("Erc20")
	}

	subAccountAddress := GetSubAccountAddress(common.HexToAddress(params.From), inputs.SubAccountIndex)

	approveCalldata, err := erc20Abi.Pack(
		"approve",
		vault.Vault,
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	repayCalldata, err := vaultAbi.Pack(
		"repay",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	repayCall, err := WrapEVCCall(
		params.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		repayCalldata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap repay call: %w", err)
	}

	return []plug_router.PlugTypesLibPlug{{
		To:   vault.Vault,
		Data: approveCalldata,
	}, repayCall}, nil
}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Direction int    `json:"direction"` // -1 for borrow, 1 for supply
		Vault     string `json:"vault"`
		Operator  int    `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal constraint APY inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}

	borrowApy, supplyApy, err := GetVaultApy(inputs.Vault, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault APY: %w", err)
	}

	var currentRate float64
	switch inputs.Direction {
	case -1:
		currentRate = utils.UintToFloat(borrowApy, 20)
	case 1:
		currentRate = utils.UintToFloat(supplyApy, 20)
	default:
		return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (supply), got %d", inputs.Direction)
	}

	switch inputs.Operator {
	case -1:
		if currentRate >= thresholdFloat {
			return nil, fmt.Errorf("current APY %.2f%% is not less than threshold %.2f%%", currentRate, thresholdFloat)
		}
	case 1:
		if currentRate <= thresholdFloat {
			return nil, fmt.Errorf("current APY %.2f%% is not greater than threshold %.2f%%", currentRate, thresholdFloat)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}

func HandleConstraintHealthFactor(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Vault           string `json:"vault"`
		Operator        int    `json:"operator"`
		Threshold       string `json:"threshold"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal constraint health factor inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}

	vault, err := GetVault(inputs.Vault, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	accountLens, err := euler_account_lens.NewEulerAccountLens(
		common.HexToAddress(references.Networks[params.ChainId].References["euler"]["account_lens"]),
		params.Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account lens: %w", err)
	}

	subAccountAddress := GetSubAccountAddress(common.HexToAddress(params.From), inputs.SubAccountIndex)

	vaultAccountInfo, err := accountLens.GetVaultAccountInfo(
		nil,
		subAccountAddress,
		vault.Vault,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault account info: %w", err)
	}

	// eurc has a lltv of .9
	// weth has a lltv of .85
	// I have .5 eurc and .001 weth deposited
	// I'm borrowing 0.4 usdc
	// my health factor is currently 6.97

	totalValueBorrowed := vaultAccountInfo.LiquidityInfo.LiabilityValue
	lltvValueCollateral := vaultAccountInfo.LiquidityInfo.CollateralValueLiquidation

	healthFactorFloat := new(big.Float).SetInt(lltvValueCollateral)
	healthFactorFloat.Quo(healthFactorFloat, new(big.Float).SetInt(totalValueBorrowed))
	healthFactorF64, _ := healthFactorFloat.Float64()

	switch inputs.Operator {
	case -1:
		if healthFactorF64 >= thresholdFloat {
			return nil, fmt.Errorf("current health factor %.2f is not less than threshold %.2f", healthFactorF64, thresholdFloat)
		}
	case 1:
		if healthFactorF64 <= thresholdFloat {
			return nil, fmt.Errorf("current health factor %.2f is not greater than threshold %.2f", healthFactorF64, thresholdFloat)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}

func HandleConstraintTimeToLiquidation(rawInputs json.RawMessage, params actions.HandlerParams) ([]plug_router.PlugTypesLibPlug, error) {
	var inputs struct {
		Vault           string `json:"vault"`
		Operator        int    `json:"operator"`
		Threshold       string `json:"threshold"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal constraint time to liquidation inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}

	vault, err := GetVault(inputs.Vault, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	accountLens, err := euler_account_lens.NewEulerAccountLens(
		common.HexToAddress(references.Networks[params.ChainId].References["euler"]["account_lens"]),
		params.Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account lens: %w", err)
	}

	subAccountAddress := GetSubAccountAddress(common.HexToAddress(params.From), inputs.SubAccountIndex)

	vaultAccountInfo, err := accountLens.GetVaultAccountInfo(
		nil,
		subAccountAddress,
		vault.Vault,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault account info: %w", err)
	}

	ttlFloat := new(big.Float).SetInt(vaultAccountInfo.LiquidityInfo.TimeToLiquidation)
	ttlFloat.Quo(ttlFloat, new(big.Float).SetInt64(60))
	ttlMinutes, _ := ttlFloat.Float64()

	switch inputs.Operator {
	case -1:
		if ttlMinutes >= thresholdFloat {
			return nil, fmt.Errorf("time to liquidation %.2f minutes is not less than threshold %.2f minutes", ttlMinutes, thresholdFloat)
		}
	case 1:
		if ttlMinutes <= thresholdFloat {
			return nil, fmt.Errorf("time to liquidation %.2f minutes is not greater than threshold %.2f minutes", ttlMinutes, thresholdFloat)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}

func WrapEVCCall(chainId uint64, targetContract common.Address, onBehalfOfAccount common.Address, value *big.Int, calldata []byte) (plug_router.PlugTypesLibPlug, error) {
	evc, err := euler_evc.EulerEvcMetaData.GetAbi()
	if err != nil {
		return plug_router.PlugTypesLibPlug{}, utils.ErrABI("EulerEvc")
	}

	callCalldata, err := evc.Pack(
		"call",
		targetContract,
		onBehalfOfAccount,
		value,
		calldata,
	)
	if err != nil {
		fmt.Printf("WrapEVCCall pack error: %v\n", err)
		return plug_router.PlugTypesLibPlug{}, utils.ErrTransaction(err.Error())
	}
	fmt.Printf("WrapEVCCall callCalldata: %v\n", callCalldata)

	return plug_router.PlugTypesLibPlug{
		To:   common.HexToAddress(references.Networks[chainId].References["euler"]["evc"]),
		Data: callCalldata,
	}, nil
}
