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
	Operation string                             `json:"operation"`
	X         coil.CoilInput[*big.Int, *big.Int] `json:"x"`
	Y         coil.CoilInput[*big.Int, *big.Int] `json:"y"`
}

var MinimumOrMaximumFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_math.PlugMathMetaData,
	FunctionName: "min",
}

func MinimumOrMaximum(lookup *actions.SchemaLookup[MinMaxRequest]) ([]signature.Plug, error) {
	MinimumOrMaximumFunc.FunctionName = lookup.Inputs.Operation

	x, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.X,
		lookup.Inputs.X.GetValueWithError,
		&MinimumOrMaximumFunc,
		"x",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get a value: %w", err)
	}

	y, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Y,
		lookup.Inputs.Y.GetValueWithError,
		&MinimumOrMaximumFunc,
		"y",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get b value: %w", err)
	}

	calldata, err := MinimumOrMaximumFunc.GetCalldata(x, y)
	if err != nil {
		return nil, fmt.Errorf("failed to get calldata for max function: %w", err)
	}

	mathContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"])
	return []signature.Plug{{
		To:      mathContract,
		Data:    calldata,
		Updates: updates,
	}}, nil
}
