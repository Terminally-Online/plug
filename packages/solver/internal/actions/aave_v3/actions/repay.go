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

type RepayRequest struct {
	Token  string                           `json:"token"`
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
}

var RepayFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     aave_v3_pool.AaveV3PoolMetaData,
	FunctionName: "repay",
}

func Repay(lookup *actions.SchemaLookup[RepayRequest]) ([]signature.Plug, error) {
	tokenIn, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	var approvalUpdates []coil.Update
	approvalAmount, approvalUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&actions.Erc20ApprovalFunc,
		"_value",
		approvalUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	approveCalldata, err := actions.Erc20ApprovalFunc.GetCalldata(
		common.HexToAddress(references.Networks[lookup.ChainId].References["aave_v3"]["pool"]),
		approvalAmount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	repayAmount, repayUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&RepayFunc,
		"amount",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	repayCalldata, err := RepayFunc.GetCalldata(
		tokenIn,
		repayAmount,
		aave_utils.InterestRateMode,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:      *tokenIn,
		Data:    approveCalldata,
		Updates: approvalUpdates,
	}, {
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["aave_v3"]["pool"]),
		Data:    repayCalldata,
		Updates: repayUpdates,
	}}, nil
}
