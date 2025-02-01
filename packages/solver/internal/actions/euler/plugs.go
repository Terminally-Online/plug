package euler

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/euler_evault_implementation"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"
)

func HandleSupply(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token string `json:"token"`
		Target string `json:"target"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal supply inputs: %v", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert supply collateral amount to uint: %w", err)
	}

	vault, err := GetVault(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}
	approveCalldata, err := vaultAbi.Pack(
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
		vault.Vault,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To: *token,
		Data: approveCalldata,
	}, {
		To: vault.Vault,
		Data: depositCalldata,
	}}, nil
}

func HandleBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token string  `json:"token"`
		Target string `json:"target"`
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal borrow inputs: %w", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert borrow amount to uint: %w", err)
	}

	vault, err := GetVault(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	calldata, err := vaultAbi.Pack(
		"borrow",
		amount,
		token,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To: vault.Vault,
		Data: calldata,
	}}, nil
}

func HandleRepay(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	return nil, nil
}

func HandleRepayWithShares(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	return nil, nil
}

func HandleWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	return nil, nil
}

func HandleReturn(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	return nil, nil
}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Direction int    `json:"direction"` // -1 for borrow, 1 for supply
		Target    string `json:"target"`
		Operator  int    `json:"apy"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal constraint APY inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}

	thresholdUint, err := utils.FloatToUint(thresholdFloat, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert threshold to uint: %w", err)
	}

	borrowApy, supplyApy, err := GetVaultApy(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault APY: %w", err)
	}

	var currentRate *big.Int
	switch inputs.Direction {
	case -1: 
		currentRate = borrowApy
	case 1:
		currentRate = supplyApy
	default:
		return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (supply), got %d", inputs.Direction)
	}

	switch inputs.Operator {
	case -1:
		if currentRate.Cmp(thresholdUint) >= 0 {
			return nil, fmt.Errorf("current APY %.2f is not less than threshold %.2f", currentRate, thresholdUint)
		}
	case 1:
		if currentRate.Cmp(thresholdUint) <= 0 {
			return nil, fmt.Errorf("current APY %.2f is not greater than threshold %.2f", currentRate, thresholdUint)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}

func HandleConstraintHealthFactor(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	return nil, nil
}

func HandleConstraintTimeToLiquidation(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	return nil, nil
}