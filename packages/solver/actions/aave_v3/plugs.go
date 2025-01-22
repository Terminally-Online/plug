package aave_v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/bindings/aave_v3_pool"
	"solver/bindings/aave_v3_ui_pool_data_provider"
	"solver/bindings/erc_20"
	"solver/internal/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleActionDeposit(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string   `json:"token"`
		Amount *big.Int `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}
	if inputs.Amount.Cmp(big.NewInt(0)) >= 0 && inputs.Amount.Cmp(utils.Uint256Max) > 0 {
		return nil, utils.ErrField("amount", inputs.Amount.String())
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		inputs.Amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	depositCalldata, err := poolAbi.Pack("deposit",
		common.HexToAddress(inputs.Token),
		inputs.Amount,
		common.HexToAddress(params.From),
		uint16(0),
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(inputs.Token),
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		Data: depositCalldata,
	}}, nil
}

func HandleActionBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string   `json:"token"`
		Amount *big.Int `json:"amount"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal borrow inputs: %v", err)
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	calldata, err := poolAbi.Pack("borrow",
		common.HexToAddress(inputs.Token),
		inputs.Amount,
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
		Token  string  `json:"token"`  // Address of the token to repay.
		Amount big.Int `json:"amount"` // Raw amount of tokens to repay.
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay inputs: %v", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		inputs.Amount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	repayCalldata, err := poolAbi.Pack("repay",
		common.HexToAddress(inputs.Token),
		inputs.Amount,
		interestRateMode,
		common.HexToAddress(params.From),
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(inputs.Token),
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(references.Networks[params.ChainId].References["aave_v3"]["pool"]),
		Data: repayCalldata,
	}}, nil
}

func HandleActionWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Token  string   `json:"token"`  // Address of the token to receive (redeeming for).
		Amount *big.Int `json:"amount"` // Raw amount of tokens to send.
	}

	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw inputs: %v", err)
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("AaveV3Pool")
	}
	calldata, err := poolAbi.Pack("withdraw",
		common.HexToAddress(inputs.Token),
		inputs.Amount,
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
		Operator  int        `json:"operator"`  // The operator to use for the threshold comparison.
		Threshold *big.Float `json:"threshold"` // The threshold value to compare against.
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal health factor inputs")
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

	return nil, nil
}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
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

	return nil, nil
}
