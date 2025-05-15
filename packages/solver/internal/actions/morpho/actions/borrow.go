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

type BorrowRequest struct {
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
	Token  string                           `json:"token"`
	Target string                           `json:"target"`
}

var BorrowFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     morpho_router.MorphoRouterMetaData,
	FunctionName: "borrow",
}

func Borrow(lookup *actions.SchemaLookup[BorrowRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	actionParams, err := types.DeserializeFromCompactString(lookup.Inputs.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize market params: %w", err)
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

	borrowCalldata, err := BorrowFunc.GetCalldata(
		actionParams.MarketParams,
		amount,
		big.NewInt(0),
		lookup.From,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	morphoRouterContract := common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"])
	return []signature.Plug{{
		To:      morphoRouterContract,
		Data:    borrowCalldata,
		Updates: updates,
	}}, nil
}
