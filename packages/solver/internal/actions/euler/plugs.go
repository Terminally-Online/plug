package euler

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/euler_account_lens"
	"solver/bindings/euler_evault_implementation"
	"solver/bindings/euler_evc"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func HandleEarn(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token string `json:"token"`
		Target string `json:"target"`
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
		params.From,
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

func HandleDepositCollateral(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token string `json:"token"`
		Target string `json:"target"`
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

	evc, err := euler_evc.EulerEvcMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvc")
	}

	// TODO: All of these calls need to go through the evc call method and not to the vaults directly.
	// The subaccounts need to be passed along as the receiver parameter and in the onBehalfOfAccount parameter on the call() function.
	// Oh my fuck we have to manage their stupid ass virtual accounts here...
	// I believe we have to make all calls through the evc call method in order for it manage sub accounts for us.
	// enableCollateralCalldata, err := evc.Pack(
	// 	""


	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		params.From,
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

func HandleWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount string `json:"amount"`
		Token  string  `json:"token"`
		Target string `json:"target"`
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

	vault, err := GetVault(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	calldata, err := vaultAbi.Pack(
		"withdraw",
		amount,
		params.From,
		params.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To: vault.Vault,
		Data: calldata,
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

	_, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
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
		params.From,
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
	var inputs struct {
		Amount string `json:"amount"`
		Token string  `json:"token"`
		Target string `json:"target"`
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

	vault, err := GetVault(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	calldata, err := vaultAbi.Pack(
		"repay",
		amount,
		params.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To: vault.Vault,
		Data: calldata,
	}}, nil
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
	var inputs struct {
		Target   	string `json:"target"`
		Operator 	int    `json:"operator"`
		Threshold 	string `json:"threshold"`
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal constraint health factor inputs: %w", err)
	}

	thresholdFloat, err := strconv.ParseFloat(inputs.Threshold, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse threshold: %w", err)
	}

	vault, err := GetVault(inputs.Target, params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	provider, err := utils.GetProvider(params.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	accountLens, err := euler_account_lens.NewEulerAccountLens(
		common.HexToAddress(references.Networks[params.ChainId].References["euler"]["account_lens"]),
		provider,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get account lens: %w", err)
	}

	vaultAccountInfo, err := accountLens.GetVaultAccountInfo(
		nil,
		common.HexToAddress(params.From),
		vault.Vault,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault account info: %w", err)
	}

	fmt.Printf("vaultAccountInfo: %v\n", vaultAccountInfo)
	fmt.Printf("vaultAccountInfo.LiquidityInfo %v\n", vaultAccountInfo.LiquidityInfo)
	fmt.Printf("vaultAccountInfo.LiquidityInfo.CollateralValueBorrowing %v\n", vaultAccountInfo.LiquidityInfo.CollateralValueBorrowing)
	fmt.Printf("vaultAccountInfo.LiquidityInfo.CollateralValueLiquidation %v\n", vaultAccountInfo.LiquidityInfo.CollateralValueLiquidation)
	fmt.Printf("vaultAccountInfo.LiquidityInfo.CollateralValueRaw %v\n", vaultAccountInfo.LiquidityInfo.CollateralValueRaw)
	fmt.Printf("vaultAccountInfo.LiquidityInfo.LiabilityValue %v\n", vaultAccountInfo.LiquidityInfo.LiabilityValue)

	// eurc has a lltv of .9
	// weth has a lltv of .85
	// I have .5 eurc and .001 weth deposited
	// I'm borrowing 0.4 usdc
	// my health factor is currently 6.97

	// This is guesswork rn, need to log this stuff out first to make sure the values I'm using here match up with what they're using on their UI
	return nil, nil
}

func HandleConstraintTimeToLiquidation(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	return nil, nil
}