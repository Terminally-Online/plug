package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/morpho_router"
	"solver/internal/actions"
	"solver/internal/actions/morpho/types"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type RepayRequest struct {
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
	Token  string                           `json:"token"`
	Target string                           `json:"target"`
}

var RepayFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     morpho_router.MorphoRouterMetaData,
	FunctionName: "repay",
}

func Repay(lookup *actions.SchemaLookup[RepayRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	actionParams, err := types.DeserializeFromCompactString(lookup.Inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize market params: %w", err)
	}

	approvalAmount, approvalUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&actions.Erc20ApprovalFunc,
		"_value",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	approveCalldata, err := actions.Erc20ApprovalFunc.GetCalldata(
		common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"]),
		approvalAmount,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	repayAmount, repayUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&RepayFunc,
		"assets",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	repayCalldata, err := RepayFunc.GetCalldata(
		actionParams.MarketParams,
		repayAmount,
		big.NewInt(0),
		lookup.From,
		[]byte{},
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:      *token,
		Data:    approveCalldata,
		Updates: approvalUpdates,
	}, {
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"]),
		Data:    repayCalldata,
		Updates: repayUpdates,
	}}, nil
}
