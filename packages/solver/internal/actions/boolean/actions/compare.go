package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/plug_boolean"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type CompareNumbersRequest struct {
	A          coil.CoilInput[*big.Int, *big.Int] `json:"a"`
	Comparison string                             `json:"comparison"`
	B          coil.CoilInput[*big.Int, *big.Int] `json:"b"`
}

var NumberComparisonFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_boolean.PlugBooleanMetaData,
	FunctionName: "isEqual",
}

func NumberComparison(lookup *actions.SchemaLookup[CompareNumbersRequest]) ([]signature.Plug, error) {
	a, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.A,
		lookup.Inputs.A.GetValueWithError,
		&NumberComparisonFunc,
		"a",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	b, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.B,
		lookup.Inputs.B.GetValueWithError,
		&NumberComparisonFunc,
		"b",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, err
	}

	NumberComparisonFunc.FunctionName = lookup.Inputs.Comparison
	calldata, err := NumberComparisonFunc.GetCalldata(a, b)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", lookup.Inputs.Comparison, err)
	}

	booleanContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["boolean"])
	return []signature.Plug{{
		Selector: signature.StaticCall,
		To:       booleanContract,
		Data:     calldata,
		Updates:  updates,
	}}, nil
}
