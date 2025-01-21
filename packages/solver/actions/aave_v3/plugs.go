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
	var inputs struct {
		types.BaseInputs
		Token  string   `json:"token"`
		Amount *big.Int `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}
	if inputs.Amount.Cmp(big.NewInt(0)) >= 0 && inputs.Amount.Cmp(utils.Uint256Max) > 0 {
		return nil, utils.ErrInvalidField("amount", inputs.Amount.String())
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(utils.Networks[params.ChainId].References["aave_v3"]["pool"]),
		inputs.Amount,
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV3Pool")
	}
	depositCalldata, err := poolAbi.Pack("deposit",
		common.HexToAddress(inputs.Token),
		inputs.Amount,
		common.HexToAddress(params.From),
		uint16(0),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   inputs.Token,
		Data: "0x" + common.Bytes2Hex(approveCalldata),
	}, {
		To:   utils.Networks[params.ChainId].References["aave_v3"]["pool"],
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

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV3Pool")
	}
	calldata, err := poolAbi.Pack("borrow",
		common.HexToAddress(inputs.TokenOut),
		inputs.AmountOut,
		interestRateMode,
		uint16(0),
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   utils.Networks[params.ChainId].References["aave_v3"]["pool"],
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

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(utils.Networks[params.ChainId].References["aave_v3"]["pool"]),
		inputs.AmountIn,
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV3Pool")
	}
	repayCalldata, err := poolAbi.Pack("repay",
		common.HexToAddress(inputs.TokenIn),
		inputs.AmountIn,
		interestRateMode,
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   inputs.TokenIn,
		Data: "0x" + common.Bytes2Hex(approveCalldata),
	}, {
		To:   utils.Networks[params.ChainId].References["aave_v3"]["pool"],
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

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV3Pool")
	}
	calldata, err := poolAbi.Pack("withdraw",
		common.HexToAddress(inputs.TokenOut),
		inputs.AmountOut,
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:   utils.Networks[params.ChainId].References["aave"]["pool"],
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

	healthFactor, err := getHealthFactor(params.ChainId, params.From)
	if err != nil {
		return nil, fmt.Errorf("failed to get health factor: %v", err)
	}

	healthFactorFloat := new(big.Float).SetInt(healthFactor)
	switch inputs.Operator {
	case -1:
		if healthFactorFloat.Cmp(inputs.Threshold) >= 0 {
			return nil, fmt.Errorf("current health factor %.2f is not less than threshold %.2f", healthFactorFloat, inputs.Threshold)
		}
	case 1:
		if healthFactorFloat.Cmp(inputs.Threshold) <= 0 {
			return nil, fmt.Errorf("current health factor %.2f is not greater than threshold %.2f", healthFactorFloat, inputs.Threshold)
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
		Threshold *big.Float `json:"threshold"` // Percentage
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs")
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
	rateFloat := new(big.Float).Quo(
		new(big.Float).SetInt(rate),
		new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil)),
	)

	switch inputs.Operator {
	case -1:
		if rateFloat.Cmp(inputs.Threshold) >= 0 {
			return nil, fmt.Errorf("current rate %.2f%% is not less than threshold %.2f%%", rateFloat, inputs.Threshold)
		}
	case 1:
		if rateFloat.Cmp(inputs.Threshold) <= 0 {
			return nil, fmt.Errorf("current rate %.2f%% is not greater than threshold %.2f%%", rateFloat, inputs.Threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return []*types.Transaction{}, nil
}
