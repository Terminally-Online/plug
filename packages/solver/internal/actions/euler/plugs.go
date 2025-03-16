package euler

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/erc_20"
	"solver/bindings/euler_account_lens"
	"solver/bindings/euler_evault_implementation"
	"solver/bindings/euler_evc"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func HandleEarn(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Vault  string `json:"vault"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
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

	vault, err := GetVault(inputs.Vault, lookup.ChainId)
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

	approveCalldata, err := erc20Abi.Pack(
		"approve",
		vault.Vault,
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}
	depositCall, err := WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		lookup.From,
		big.NewInt(0),
		depositCalldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap deposit call: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, depositCall}, nil
}

// I believe this is exactly the same as WithdrawCollateral, but I don't want to rope it in with the other handler with a default sub account index of 0 in case there is some difference I'm not aware of yet. On the options side, this one does not have a sub account index option to remove complexity around earn deposits not being indexed in a contract like borrow and lending positions.
func HandleWithdraw(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string `json:"token"`
		Vault  string `json:"vault"`
	}

	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw inputs: %w", err)
	}

	collateralInputs := struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}{
		Amount:          inputs.Amount,
		Token:           inputs.Token,
		Vault:           inputs.Vault,
		SubAccountIndex: 0,
	}

	collateralRawInputs, err := json.Marshal(collateralInputs)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal collateral inputs: %w", err)
	}

	return HandleWithdrawCollateral(lookup, collateralRawInputs)
}

func HandleDepositCollateral(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
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

	vault, err := GetVault(inputs.Vault, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("Erc20")
	}

	approveCalldata, err := erc20Abi.Pack(
		"approve",
		vault.Vault,
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

	subAccountAddress := GetSubAccountAddress(lookup.From, inputs.SubAccountIndex)

	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	depositCall, err := WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		depositCalldata,
		nil,
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
		lookup.ChainId,
		common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["evc"]),
		subAccountAddress,
		big.NewInt(0),
		enableCollateralCalldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap enable collateral call: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, depositCall, enableCollateralCall}, nil
}

func HandleWithdrawCollateral(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(raw, &inputs); err != nil {
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

	vault, err := GetVault(inputs.Vault, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	subAccountAddress := GetSubAccountAddress(lookup.From, inputs.SubAccountIndex)

	calldata, err := vaultAbi.Pack(
		"withdraw",
		amount,
		lookup.From,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	call, err := WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		calldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap withdraw collateral call: %w", err)
	}

	return []signature.Plug{call}, nil
}

func HandleBorrow(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(raw, &inputs); err != nil {
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

	vault, err := GetVault(inputs.Vault, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	subAccountAddress := GetSubAccountAddress(lookup.From, inputs.SubAccountIndex)

	calldata, err := vaultAbi.Pack(
		"borrow",
		amount,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	call, err := WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		calldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap borrow call: %w", err)
	}

	return []signature.Plug{call}, nil
}

func HandleRepay(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Amount          string `json:"amount"`
		Token           string `json:"token"`
		Vault           string `json:"vault"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(raw, &inputs); err != nil {
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

	vault, err := GetVault(inputs.Vault, lookup.ChainId)
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

	subAccountAddress := GetSubAccountAddress(lookup.From, inputs.SubAccountIndex)

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
		lookup.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		repayCalldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap repay call: %w", err)
	}

	return []signature.Plug{{
		To:   vault.Vault,
		Data: approveCalldata,
	}, repayCall}, nil
}

func HandleConstraintAPY(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Direction int    `json:"direction"` // -1 for borrow, 1 for supply
		Vault     string `json:"vault"`
		Operator  int    `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal constraint APY inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}

	borrowApy, supplyApy, err := GetVaultApy(inputs.Vault, lookup.ChainId)
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

func HandleConstraintHealthFactor(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Operator        int    `json:"operator"`
		Threshold       string `json:"threshold"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal constraint health factor inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}

	accountLens, err := euler_account_lens.NewEulerAccountLens(
		common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["account_lens"]),
		lookup.Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account lens: %w", err)
	}

	subAccountAddress := GetSubAccountAddress(lookup.From, inputs.SubAccountIndex)
	evcAddress := common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["evc"])

	vaultAccountInfos, err := accountLens.GetAccountEnabledVaultsInfo(
		nil,
		evcAddress,
		subAccountAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account vaults info: %w", err)
	}

	var borrowVault *euler_account_lens.VaultAccountInfo
	for _, vaultAccountInfo := range vaultAccountInfos.VaultAccountInfo {
		if !vaultAccountInfo.LiquidityInfo.QueryFailure {
			borrowVault = &vaultAccountInfo
			break
		}
	}

	if borrowVault == nil {
		return nil, fmt.Errorf("no vault found with valid liquidity info")
	}

	// eurc has a lltv of .9
	// weth has a lltv of .85
	// I have .5 eurc and .001 weth deposited
	// I'm borrowing 0.4 usdc
	// my health factor is currently 6.97

	totalValueBorrowed := borrowVault.LiquidityInfo.LiabilityValue
	lltvValueCollateral := borrowVault.LiquidityInfo.CollateralValueLiquidation

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

func HandleConstraintTimeToLiquidation(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs struct {
		Operator        int    `json:"operator"`
		Threshold       string `json:"threshold"`
		SubAccountIndex uint8  `json:"sub-account"`
	}

	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal constraint time to liquidation inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}

	accountLens, err := euler_account_lens.NewEulerAccountLens(
		common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["account_lens"]),
		lookup.Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account lens: %w", err)
	}

	subAccountAddress := GetSubAccountAddress(lookup.From, inputs.SubAccountIndex)
	evcAddress := common.HexToAddress(references.Networks[lookup.ChainId].References["euler"]["evc"])

	vaultAccountInfos, err := accountLens.GetAccountEnabledVaultsInfo(
		nil,
		evcAddress,
		subAccountAddress,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account vaults info: %w", err)
	}

	var borrowVault *euler_account_lens.VaultAccountInfo
	for _, vaultAccountInfo := range vaultAccountInfos.VaultAccountInfo {
		if !vaultAccountInfo.LiquidityInfo.QueryFailure {
			borrowVault = &vaultAccountInfo
			break
		}
	}

	if borrowVault == nil {
		return nil, fmt.Errorf("no vault found with valid liquidity info")
	}

	ttlFloat := new(big.Float).SetInt(borrowVault.LiquidityInfo.TimeToLiquidation)
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

func WrapEVCCall(chainId uint64, targetContract common.Address, onBehalfOfAccount common.Address, value *big.Int, calldata []byte, updates *[]coil.Update) (signature.Plug, error) {
	evc, err := euler_evc.EulerEvcMetaData.GetAbi()
	if err != nil {
		return signature.Plug{}, utils.ErrABI("EulerEvc")
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
		return signature.Plug{}, utils.ErrTransaction(err.Error())
	}

	return signature.Plug{
		To:      common.HexToAddress(references.Networks[chainId].References["euler"]["evc"]),
		Data:    callCalldata,
		Updates: *updates,
	}, nil
}
