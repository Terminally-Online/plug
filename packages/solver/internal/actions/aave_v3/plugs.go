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
		common.HexToAddress(params.From),
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
		common.HexToAddress(params.From),
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
		Token  string `json:"token"`  // Address of the token to repay.
		Amount string `json:"amount"` // Raw amount of tokens to repay.
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
		common.HexToAddress(params.From),
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
		Token  string `json:"token"`  // Address of the token to receive (redeeming for).
		Amount string `json:"amount"` // Raw amount of tokens to send.
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
		common.HexToAddress(params.From),
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
		Operator  int    `json:"operator"`  // The operator to use for the threshold comparison.
		Threshold string `json:"threshold"` // The threshold value to compare against.
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
		Direction int    `json:"direction"` // -1 for borrow, 1 for deposit
		Token     string `json:"token"`     // Underlying token address
		Operator  int    `json:"operator"`  // -1 for less than, 1 for greater than
		Threshold string `json:"threshold"` // Percentage
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs")
	}

	// Convert threshold percentage to RAY units (27 decimals)
	threshold, err := utils.StringToUint(inputs.Threshold, 27)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	// NOTE: We pass in `true` to force a cache update because we want the latest APY results.
	reserves, err := getReserves(params.ChainId)
	if err != nil {
		return nil, err
	}

	var targetReserve *aave_v3_ui_pool_data_provider.IUiPoolDataProviderV3AggregatedReserveData
	for _, reserve := range reserves {
		if reserve.UnderlyingAsset == common.HexToAddress(inputs.Token) {
			targetReserve = &reserve
			break
		}
	}
	if targetReserve == nil {
		return nil, fmt.Errorf("token %s not supported", inputs.Token)
	}

	var rate *big.Int
	switch inputs.Direction {
	case -1:
		rate = targetReserve.VariableBorrowRate
	case 1:
		rate = targetReserve.LiquidityRate
	default:
		return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (deposit), got %d", inputs.Direction)
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
