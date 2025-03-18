package actions

import (
	"fmt"
	"solver/bindings/aave_v3_pool"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawRequest struct {
	Token  string `json:"token"`
	Amount string `json:"amount"`
}

func Withdraw(lookup *actions.SchemaLookup[WithdrawRequest]) ([]signature.Plug, error) {
	tokenOut, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}
	amountOut, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
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
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(references.Networks[lookup.ChainId].References["aave"]["pool"]),
		Data: calldata,
	}}, nil
}
