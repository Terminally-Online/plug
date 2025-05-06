package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/aave_v3_pool"
	"solver/internal/actions"
	aave_utils "solver/internal/actions/aave_v3/utils"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type BorrowRequest struct {
	Token  string                           `json:"token"`
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
}

var BorrowFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     aave_v3_pool.AaveV3PoolMetaData,
	FunctionName: "borrow",
}

func Borrow(lookup *actions.SchemaLookup[BorrowRequest]) ([]signature.Plug, error) {
	tokenOut, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&BorrowFunc,
		"amount",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	calldata, err := BorrowFunc.GetCalldata(
		tokenOut,
		amount,
		aave_utils.InterestRateMode,
		uint16(0),
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	aavePool := common.HexToAddress(references.Networks[lookup.ChainId].References["aave_v3"]["pool"])
	return []signature.Plug{{
		To:      aavePool,
		Data:    calldata,
		Updates: updates,
	}}, nil
}
