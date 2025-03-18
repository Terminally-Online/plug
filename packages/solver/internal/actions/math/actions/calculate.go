package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/plug_math"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type CalculateRequest struct {
	X         *big.Int `json:"x"`
	Operation string   `json:"operation"`
	Y         *big.Int `json:"y"`
}

func Calculate(lookup *actions.SchemaLookup[CalculateRequest]) ([]signature.Plug, error) {
	if lookup.Inputs.X == nil || lookup.Inputs.Y == nil {
		return nil, fmt.Errorf("x and y values must be provided")
	}

	mathContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}

	operation := strings.ToLower(lookup.Inputs.Operation)

	var functionName string
	switch operation {
	case "+", "add":
		functionName = "add"
	case "-", "sub", "subtract":
		functionName = "subtract"
	case "*", "mul", "multiply":
		functionName = "multiply"
	case "÷", "/", "div", "divide":
		if lookup.Inputs.Y.Cmp(big.NewInt(0)) == 0 {
			return nil, fmt.Errorf("divide by zero error")
		}
		functionName = "divide"
	case "%", "mod", "modulo":
		if lookup.Inputs.Y.Cmp(big.NewInt(0)) == 0 {
			return nil, fmt.Errorf("modulo by zero error")
		}
		functionName = "modulo"
	default:
		return nil, fmt.Errorf("unsupported operation: %s", lookup.Inputs.Operation)
	}

	calldata, err := mathAbi.Pack(functionName, lookup.Inputs.X, lookup.Inputs.Y)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", functionName, err)
	}

	plug := signature.Plug{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}

	return []signature.Plug{plug}, nil
}
