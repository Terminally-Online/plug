package aave_v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/aave_v3_pool"
	"solver/bindings/aave_v3_ui_pool_data_provider"
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionDeposit(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string `json:"token"`
		Amount string `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %w", err)
	}

	token, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}
	amount, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	depositCalldata, err := poolAbi.Pack("deposit",
		token,
		amount,
		params.From,
		uint16(0),
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		Data: depositCalldata,
	}}, nil
}

func HandleActionBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string `json:"token"`
		Amount string `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal borrow inputs: %w", err)
	}

	tokenOut, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}
	amountOut, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert borrow amount to uint: %w", err)
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	calldata, err := poolAbi.Pack("borrow",
		tokenOut,
		amountOut,
		interestRateMode,
		uint16(0),
		params.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		Data: calldata,
	}}, nil
}

func HandleActionRepay(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string `json:"token"`  
		Amount string `json:"amount"` 
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay inputs: %w", err)
	}

	tokenIn, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}
	amountIn, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert repayment amount to uint: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		amountIn,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	repayCalldata, err := poolAbi.Pack("repay",
		tokenIn,
		amountIn,
		interestRateMode,
		params.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   *tokenIn,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		Data: repayCalldata,
	}}, nil
}

func HandleActionWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string `json:"token"`  
		Amount string `json:"amount"` 
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw inputs: %w", err)
	}

	tokenOut, decimals, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}
	amountOut, err := utils.StringToUint(inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert withdraw amount to uint: %w", err)
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	calldata, err := poolAbi.Pack("withdraw",
		tokenOut,
		amountOut,
		params.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(references.Networks[params.ChainId].References["aave"]["pool"]),
		Data: calldata,
	}}, nil
}

func HandleConstraintHealthFactor(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Operator  int    `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal health factor inputs: %w", err)
	}

	// NOTE: Aave v3 uses 18 decimals for their health factor.
	//      https://github.com/aave/aave-v3-core/blob/782f51917056a53a2c228701058a6c3fb233684a/test-suites/emode.spec.ts#L555
	threshold, err := utils.StringToUint(inputs.Threshold, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert threshold to uint: %w", err)
	}

	healthFactor, err := getHealthFactor(params.ChainId, params.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get health factor: %w", err)
	}

	switch inputs.Operator {
	case -1:
		if healthFactor.Cmp(threshold) >= 0 {
			return nil, fmt.Errorf("current health factor %.2f is not less than threshold %.2f", healthFactor, threshold)
		}
	case 1:
		if healthFactor.Cmp(threshold) <= 0 {
			return nil, fmt.Errorf("current health factor %.2f is not greater than threshold %.2f", healthFactor, threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Action    int    `json:"action"`
		Token     string `json:"token"`
		Operator  int    `json:"operator"`
		Threshold string `json:"threshold"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs")
	}

	token, _, err := utils.ParseAddressAndDecimals(inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	threshold, err := utils.StringToUint(inputs.Threshold, 27)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	reserves, err := getReserves(params.ChainId)
	if err != nil {
		return nil, err
	}

	var targetReserve *aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData
	for _, reserve := range reserves {
		if reserve.UnderlyingAsset == *token {
			targetReserve = &reserve
			break
		}
	}
	if targetReserve == nil {
		return nil, fmt.Errorf("token %s not supported", inputs.Token)
	}

	var rate *big.Int
	switch inputs.Action {
	case -1:
		rate = targetReserve.VariableBorrowRate
	case 1:
		rate = targetReserve.LiquidityRate
	default:
		return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (deposit), got %d", inputs.Action)
	}

	switch inputs.Operator {
	case -1:
		if rate.Cmp(threshold) >= 0 {
			return nil, fmt.Errorf("current rate %.2f%% is not less than threshold %.2f%%", rate, threshold)
		}
	case 1:
		if rate.Cmp(threshold) <= 0 {
			return nil, fmt.Errorf("current rate %.2f%% is not greater than threshold %.2f%%", rate, threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}
