package aave_v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/bindings/aave_v3_pool"
	"solver/bindings/aave_v3_ui_pool_data_provider"
	"solver/bindings/erc_20"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionDeposit(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs types.DepositInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	tokenIn, decimals, err := utils.ParseAddressAndDecimals(inputs.TokenIn)
	if err != nil {
		return nil, err
	}
	amountIn, err := utils.FloatToUint(inputs.AmountIn, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %v", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(poolAddress), amountIn)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV3Pool")
	}
	depositCalldata, err := poolAbi.Pack("deposit",
		common.HexToAddress(tokenIn),
		amountIn,
		common.HexToAddress(params.From),
		uint16(0),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   tokenIn,
		Data: "0x" + common.Bytes2Hex(approveCalldata),
	}, {
		To:   poolAddress,
		Data: "0x" + common.Bytes2Hex(depositCalldata),
	}}, nil
}

func HandleActionBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs types.BorrowInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal borrow inputs: %v", err)
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	tokenOut, decimals, err := utils.ParseAddressAndDecimals(inputs.TokenOut)
	if err != nil {
		return nil, err
	}
	amountOut, err := utils.FloatToUint(inputs.AmountOut, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert borrow amount to uint: %v", err)
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV3Pool")
	}
	calldata, err := poolAbi.Pack("borrow",
		common.HexToAddress(tokenOut),
		amountOut,
		interestRateMode,
		uint16(0),
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   poolAddress,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, nil
}

func HandleActionRepay(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs types.RepayInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay inputs: %v", err)
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	tokenIn, decimals, err := utils.ParseAddressAndDecimals(inputs.TokenIn)
	if err != nil {
		return nil, err
	}
	amountIn, err := utils.FloatToUint(inputs.AmountIn, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert repayment amount to uint: %v", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack("approve", common.HexToAddress(poolAddress), amountIn)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV3Pool")
	}
	repayCalldata, err := poolAbi.Pack("repay",
		common.HexToAddress(tokenIn),
		amountIn,
		interestRateMode,
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   tokenIn,
		Data: "0x" + common.Bytes2Hex(approveCalldata),
	}, {
		To:   poolAddress,
		Data: "0x" + common.Bytes2Hex(repayCalldata),
	}}, nil
}

func HandleActionWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs types.WithdrawInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw inputs: %v", err)
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	tokenOut, decimals, err := utils.ParseAddressAndDecimals(inputs.TokenOut)
	if err != nil {
		return nil, err
	}
	amountOut, err := utils.FloatToUint(inputs.AmountOut, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert withdraw amount to uint: %v", err)
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV3Pool")
	}
	calldata, err := poolAbi.Pack("withdraw",
		common.HexToAddress(tokenOut),
		amountOut,
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   poolAddress,
		Data: "0x" + common.Bytes2Hex(calldata),
	}}, nil
}

func HandleConstraintHealthFactor(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs types.ThresholdInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal health factor inputs")
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	// aavev3 uses 18 decimals for their health factor https://github.com/aave/aave-v3-core/blob/782f51917056a53a2c228701058a6c3fb233684a/test-suites/emode.spec.ts#L555
	threshold, err := utils.FloatToUint(inputs.Threshold, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert threshold to uint: %v", err)
	}

	healthFactor, err := getHealthFactor(params.ChainId, params.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get health factor: %v", err)
	}

	switch inputs.Operator {
	case -1:
		if healthFactor.Cmp(threshold) >= 0 {
			return nil, fmt.Errorf("current health factor %.2f is not less than threshold %.2f", healthFactor, inputs.Threshold)
		}
	case 1:
		if healthFactor.Cmp(threshold) <= 0 {
			return nil, fmt.Errorf("current health factor %.2f is not greater than threshold %.2f", healthFactor, inputs.Threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return []*types.Transaction{}, nil
}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Direction int        `json:"direction"` // -1 for borrow, 1 for deposit
		Token     string     `json:"token"`     // Underlying token address
		Operator  int        `json:"operator"`  // -1 for less than, 1 for greater than
		Threshold float64     `json:"threshold"` // Percentage
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs")
	}

    // Convert threshold percentage to RAY units (27 decimals)
	threshold, err := utils.FloatToUint(inputs.Threshold, 27)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %v", err)
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

	return []*types.Transaction{}, nil
}
