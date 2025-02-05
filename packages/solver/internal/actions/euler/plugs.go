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
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

func HandleEarn(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount 			string `json:"amount"`
		Token 			string `json:"token"`
		Target 			string `json:"target"`
		SubAccountIndex uint8  `json:"subAccountIndex"` 
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
		*big.NewInt(0),
		depositCalldata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap deposit call: %w", err)
	}

	return []signature.Plug{{
		To: *token,
		Data: approveCalldata,
	}, depositCall}, nil
}

func HandleDepositCollateral(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Amount 			string `json:"amount"`
		Token 			string `json:"token"`
		Target 			string `json:"target"`
		SubAccountIndex uint8 `json:"subAccountIndex"`
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
		*big.NewInt(0),
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
		*big.NewInt(0),
		enableCollateralCalldata,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap enable collateral call: %w", err)
	}

	return []signature.Plug{{
		To: *token,
		Data: approveCalldata,
	}, depositCall, enableCollateralCall}, nil
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

	_, err := strconv.ParseFloat(inputs.Threshold, 64)
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

func WrapEVCCall(chainId uint64, targetContract common.Address, onBehalfOfAccount common.Address, value big.Int, calldata []byte) (signature.Plug, error) {
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
		return signature.Plug{}, utils.ErrTransaction(err.Error())
	}

	return signature.Plug{
		To: common.HexToAddress(references.Networks[chainId].References["euler"]["evc"]),
		Data: callCalldata,
	}, nil
}

func GetSubAccountAddress(account common.Address, index uint8) common.Address {
	// Convert owner address to big.Int for bitwise operations
    ownerInt := new(big.Int).SetBytes(account[:])
    
    // Create big.Int for accountId
    accountIdInt := new(big.Int).SetUint64(uint64(index))
    
    // XOR operation
    result := new(big.Int).Xor(ownerInt, accountIdInt)
    
    // Convert back to address
    var subAccount common.Address
    resultBytes := result.Bytes()
    
    // Ensure proper padding to 20 bytes (address length)
    copy(subAccount[20-len(resultBytes):], resultBytes)
    
    return subAccount
}