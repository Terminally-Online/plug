package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/plug_math"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/coil"
	"solver/internal/solver/signature"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type CalculateRequest struct {
	X         coil.CoilInput[*big.Int, *big.Int] `json:"x"` // TODO Mason: should the T be a pointer?
	Operation string                             `json:"operation"`
	Y         coil.CoilInput[*big.Int, *big.Int] `json:"y"`
}

func Calculate(lookup *actions.SchemaLookup[CalculateRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.X.GetValue() == nil || lookup.Inputs.Y.GetValue() == nil {
		return nil, fmt.Errorf("x and y values must be provided")
	}

	var functionName string
	switch strings.ToLower(lookup.Inputs.Operation) {
	case "+", "add":
		functionName = "add"
	case "-", "sub", "subtract":
		functionName = "subtract"
	case "*", "mul", "multiply":
		functionName = "multiply"
	case "÷", "/", "div", "divide":
		functionName = "divide"
	case "%", "mod", "modulo":
		functionName = "modulo"
	default:
		return nil, fmt.Errorf("unsupported operation: %s", lookup.Inputs.Operation)
	}

	CalculateFunc := actions.ActionOnchainFunctionResponse{
		Metadata:     plug_math.PlugMathMetaData,
		FunctionName: functionName,
	}

	var xUpdates []coil.Update
	xResult, xUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.X,
		lookup.Inputs.X.GetValueWithError,
		&CalculateFunc,
		"a",
		xUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get x value: %w", err)
	}

	var yUpdates []coil.Update
	yResult, yUpdates, err := actions.GetAndUpdate(
		&lookup.Inputs.Y,
		lookup.Inputs.Y.GetValueWithError,
		&CalculateFunc,
		"b",
		yUpdates,
		lookup.PreviousActionDefinition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get y value: %w", err)
	}

	// Check if the divisor is zero for divide and modulo operations
	if yResult.Cmp(big.NewInt(0)) == 0 && (functionName == "divide" || functionName == "modulo") {
		return nil, fmt.Errorf("divide by zero error")
	}

	calldata, err := CalculateFunc.GetCalldata(
		xResult,
		yResult,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", functionName, err)
	}

	plug := signature.Plug{
		To:      common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"]),
		Data:    calldata,
		Value:   nil,
		Updates: append(xUpdates, yUpdates...),
	}

	return []signature.Plug{plug}, nil
}
