package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/plug_math"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type MinMaxRequest struct {
	A coil.CoilInput[*big.Int, *big.Int] `json:"a"`
	B coil.CoilInput[*big.Int, *big.Int] `json:"b"`
}

var MinFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_math.PlugMathMetaData,
	FunctionName: "min",
}

var MaxFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_math.PlugMathMetaData,
	FunctionName: "max",
}

func Min(lookup *actions.SchemaLookup[MinMaxRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.A.GetValue() == nil || lookup.Inputs.B.GetValue() == nil {
		return nil, fmt.Errorf("a and b values must be provided")
	}

	var aUpdates []coil.Update
	a, aUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.A,
		lookup.Inputs.A.GetValueWithError,
		&MinFunc,
		"a",
		aUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get a value: %w", err)
	}

	var bUpdates []coil.Update
	b, bUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.B,
		lookup.Inputs.B.GetValueWithError,
		&MinFunc,
		"b",
		bUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get b value: %w", err)
	}

	calldata, err := MinFunc.GetCalldata(a, b)
	if err != nil {
		return nil, fmt.Errorf("failed to get calldata for min function: %w", err)
	}

	return []signature.Plug{{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"]),
		Data:    calldata,
		Value:   nil,
		Updates: append(aUpdates, bUpdates...),
	}}, nil
}

func Max(lookup *actions.SchemaLookup[MinMaxRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.A.GetValue() == nil || lookup.Inputs.B.GetValue() == nil {
		return nil, fmt.Errorf("a and b values must be provided")
	}

	var aUpdates []coil.Update
	a, aUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.A,
		lookup.Inputs.A.GetValueWithError,
		&MaxFunc,
		"a",
		aUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get a value: %w", err)
	}

	var bUpdates []coil.Update
	b, bUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.B,
		lookup.Inputs.B.GetValueWithError,
		&MaxFunc,
		"b",
		bUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get b value: %w", err)
	}

	calldata, err := MaxFunc.GetCalldata(a, b)
	if err != nil {
		return nil, fmt.Errorf("failed to get calldata for max function: %w", err)
	}

	return []signature.Plug{{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"]),
		Data:    calldata,
		Value:   nil,
		Updates: append(aUpdates, bUpdates...),
	}}, nil
}
