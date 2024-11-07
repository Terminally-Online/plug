package aave_v3

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/bindings/aave_v3_pool"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
)

func HandleDeposit(rawInputs json.RawMessage, params actions.HandlerParams) ([]byte, error) {
	var inputs types.DepositInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}
	return poolAbi.Pack("deposit",
		common.HexToAddress(inputs.TokenIn),
		inputs.AmountIn,
		common.HexToAddress(params.From),
		uint16(0),
	)
}

func HandleBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]byte, error) {
	var inputs types.BorrowInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal borrow inputs: %v", err)
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}
	return poolAbi.Pack("borrow",
		common.HexToAddress(inputs.TokenOut),
		inputs.AmountOut,
		interestRateMode,
		uint16(0),
		common.HexToAddress(params.From),
	)
}

func HandleRepay(rawInputs json.RawMessage, params actions.HandlerParams) ([]byte, error) {
	var inputs types.RepayInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal repay inputs: %v", err)
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}
	return poolAbi.Pack("repay",
		common.HexToAddress(inputs.TokenIn),
		inputs.AmountIn,
		interestRateMode,
		common.HexToAddress(params.From),
	)
}

func HandleWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]byte, error) {
	var inputs types.WithdrawInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal withdraw inputs: %v", err)
	}
	if err := inputs.Validate(); err != nil {
		return nil, err
	}

	poolAbi, err := aave_v3_pool.AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("AaveV2Pool")
	}
	return poolAbi.Pack("withdraw",
		common.HexToAddress(inputs.TokenOut),
		inputs.AmountOut,
		common.HexToAddress(params.From),
	)
}
