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

type ClampRequest struct {
	Value coil.CoilInput[*big.Int, *big.Int] `json:"value"`
	Min   coil.CoilInput[*big.Int, *big.Int] `json:"min"`
	Max   coil.CoilInput[*big.Int, *big.Int] `json:"max"`
}

var ClampFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_math.PlugMathMetaData,
	FunctionName: "clamp",
}

func HandleClamp(lookup *actions.SchemaLookup[ClampRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.Value.GetValue() == nil || lookup.Inputs.Min.GetValue() == nil || lookup.Inputs.Max.GetValue() == nil {
		return nil, fmt.Errorf("value, min, and max values must be provided")
	}

	var valueUpdates []coil.Update
	value, valueUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Value,
		lookup.Inputs.Value.GetValueWithError,
		&ClampFunc,
		"value",
		valueUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get value: %w", err)
	}

	var minUpdates []coil.Update
	min, minUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Min,
		lookup.Inputs.Min.GetValueWithError,
		&ClampFunc,
		"minValue",
		minUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get min value: %w", err)
	}

	var maxUpdates []coil.Update
	max, maxUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Max,
		lookup.Inputs.Max.GetValueWithError,
		&ClampFunc,
		"maxValue",
		maxUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get max value: %w", err)
	}
	updates := append(valueUpdates, minUpdates...)
	updates = append(updates, maxUpdates...)

	if min.Cmp(max) > 0 {
		return nil, fmt.Errorf("min value exceeds max value")
	}

	calldata, err := ClampFunc.GetCalldata(
		value,
		min,
		max,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack clamp calldata: %w", err)
	}

	plug := signature.Plug{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"]),
		Data:    calldata,
		Value:   nil,
		Updates: updates,
	}

	return []signature.Plug{plug}, nil
}
