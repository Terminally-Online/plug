package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/yearn_v3_gauge"
	"solver/internal/actions"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"solver/internal/utils"
)

type UnstakeRequest struct {
	Amount coil.CoilInput[string, *big.Int] `json:"amount"`
	Gauge  string                           `json:"gauge"`
}

var UnstakeFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     yearn_v3_gauge.YearnV3GaugeMetaData,
	FunctionName: "redeem",
}

func Redeem(lookup *actions.SchemaLookup[UnstakeRequest]) ([]signature.Plug, error) {
	gauge, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Gauge)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	var redeemUpdates []coil.Update
	redeemAmount, redeemUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Amount,
		lookup.Inputs.Amount.GetUintFromFloatFunc(uint8(decimals)),
		&UnstakeFunc,
		"_assets",
		redeemUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	redeemCalldata, err := UnstakeFunc.GetCalldata(
		redeemAmount,
		lookup.From,
		lookup.From,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:      *gauge,
		Data:    redeemCalldata,
		Updates: redeemUpdates,
	}}, nil
}
