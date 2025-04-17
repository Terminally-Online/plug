package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/morpho_router"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
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

	var amountUpdates []coil.Update
	amount, amountUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&BorrowFunc,
		"amount",
		amountUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	market, err := reads.GetMarket(lookup.Inputs.Target, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get market: %w", err)
	}

	borrowCalldata, err := BorrowFunc.GetCalldata(
		market.Params,
		amount,
		big.NewInt(0),
		lookup.From,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["morpho"]["router"]),
		Data:    borrowCalldata,
		Updates: amountUpdates,
	}}, nil
}
