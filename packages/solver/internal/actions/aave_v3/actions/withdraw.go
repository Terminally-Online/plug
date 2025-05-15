package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/aave_v3_pool"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawRequest struct {
	Token  string                           `json:"token"`
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
}

var WithdrawFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     aave_v3_pool.AaveV3PoolMetaData,
	FunctionName: "withdraw",
}

func Withdraw(lookup *actions.SchemaLookup[WithdrawRequest]) ([]signature.Plug, error) {
	tokenOut, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&WithdrawFunc,
		"amount",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	calldata, err := WithdrawFunc.GetCalldata(
		tokenOut,
		amount,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["aave"]["pool"]),
		Data:    calldata,
		Updates: updates,
	}}, nil
}
