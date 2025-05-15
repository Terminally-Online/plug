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

type CalculateRequest struct {
	X         coil.CoilInput[*big.Int, *big.Int] `json:"x"` // TODO Mason: should the T be a pointer?
	Operation string                             `json:"operation"`
	Y         coil.CoilInput[*big.Int, *big.Int] `json:"y"`
}

var CalculateFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     plug_math.PlugMathMetaData,
	FunctionName: "add",
}

func Calculate(lookup *actions.SchemaLookup[CalculateRequest]) ([]signature.Plug, error) {
	CalculateFunc.FunctionName = lookup.Inputs.Operation

	x, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.X,
		lookup.Inputs.X.GetValueWithError,
		&CalculateFunc,
		"x",
		nil,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get x value: %w", err)
	}

	y, updates, err := actions.GetAndUpdate(
		&lookup.Inputs.Y,
		lookup.Inputs.Y.GetValueWithError,
		&CalculateFunc,
		"y",
		updates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get y value: %w", err)
	}

	calldata, err := CalculateFunc.GetCalldata(x, y)
	if err != nil {
		return nil, err
	}

	mathContract := common.HexToAddress(references.Plug["math"])
	return []signature.Plug{{
		To:      mathContract,
		Data:    calldata,
		Updates: updates,
	}}, nil
}
