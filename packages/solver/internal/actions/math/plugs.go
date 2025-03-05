package math

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/plug_math"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type CalculateInputs struct {
	X         *big.Int `json:"x"`
	Operation string   `json:"operation"`
	Y         *big.Int `json:"y"`
}

type MinMaxInputs struct {
	A *big.Int `json:"a"`
	B *big.Int `json:"b"`
}

type PowerInputs struct {
	Base     *big.Int `json:"base"`
	Exponent *big.Int `json:"exponent"`
}

type ClampInputs struct {
	Value *big.Int `json:"value"`
	Min   *big.Int `json:"min"`
	Max   *big.Int `json:"max"`
}

func HandleCalculate(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs CalculateInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal calculate inputs: %w", err)
	}

	// Validate that X and Y are not nil
	if inputs.X == nil || inputs.Y == nil {
		return nil, fmt.Errorf("X and Y values must be provided")
	}

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	var functionName string
	switch inputs.Operation {
	case "+", "add":
		functionName = "add"
	case "-", "subtract", "sub":
		functionName = "subtract"
	case "*", "multiply", "mul":
		functionName = "multiply"
	case "÷", "/", "divide", "div":
		if inputs.Y.Cmp(big.NewInt(0)) == 0 {
			return nil, fmt.Errorf("divide by zero error")
		}
		functionName = "divide"
	case "%", "modulo", "mod":
		if inputs.Y.Cmp(big.NewInt(0)) == 0 {
			return nil, fmt.Errorf("modulo by zero error")
		}
		functionName = "modulo"
	default:
		operation := strings.ToLower(inputs.Operation)
		switch operation {
		case "add", "+":
			functionName = "add"
		case "subtract", "sub", "-":
			functionName = "subtract"
		case "multiply", "mul", "*":
			functionName = "multiply"
		case "divide", "div", "/", "÷":
			if inputs.Y.Cmp(big.NewInt(0)) == 0 {
				return nil, fmt.Errorf("divide by zero error")
			}
			functionName = "divide"
		case "modulo", "mod", "%":
			if inputs.Y.Cmp(big.NewInt(0)) == 0 {
				return nil, fmt.Errorf("modulo by zero error")
			}
			functionName = "modulo"
		default:
			return nil, fmt.Errorf("unsupported operation: %s", inputs.Operation)
		}
	}
	
	calldata, err := mathAbi.Pack(functionName, inputs.X, inputs.Y)
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

func HandleMin(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs MinMaxInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal min inputs: %w", err)
	}

	// Validate that A and B are not nil
	if inputs.A == nil || inputs.B == nil {
		return nil, fmt.Errorf("A and B values must be provided")
	}

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	calldata, err := mathAbi.Pack("min", inputs.A, inputs.B)
	if err != nil {
		return nil, fmt.Errorf("failed to pack min calldata: %w", err)
	}

	plug := signature.Plug{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}

func HandleMax(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs MinMaxInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal max inputs: %w", err)
	}

	// Validate that A and B are not nil
	if inputs.A == nil || inputs.B == nil {
		return nil, fmt.Errorf("A and B values must be provided")
	}

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	calldata, err := mathAbi.Pack("max", inputs.A, inputs.B)
	if err != nil {
		return nil, fmt.Errorf("failed to pack max calldata: %w", err)
	}

	plug := signature.Plug{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}

func HandlePower(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs PowerInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal power inputs: %w", err)
	}

	// Validate that Base and Exponent are not nil
	if inputs.Base == nil || inputs.Exponent == nil {
		return nil, fmt.Errorf("Base and Exponent values must be provided")
	}

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	// Check for very large exponents that could cause computational issues
	if inputs.Exponent.Cmp(big.NewInt(1000)) > 0 {
		return nil, fmt.Errorf("exponent too large: maximum allowed is 1000")
	}
	
	calldata, err := mathAbi.Pack("power", inputs.Base, inputs.Exponent)
	if err != nil {
		return nil, fmt.Errorf("failed to pack power calldata: %w", err)
	}

	plug := signature.Plug{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}

func HandleClamp(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs ClampInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal clamp inputs: %w", err)
	}

	// Validate that Value, Min, and Max are not nil
	if inputs.Value == nil || inputs.Min == nil || inputs.Max == nil {
		return nil, fmt.Errorf("Value, Min, and Max values must be provided")
	}
	
	// Validate min <= max
	if inputs.Min.Cmp(inputs.Max) > 0 {
		return nil, fmt.Errorf("min value exceeds max value")
	}

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	calldata, err := mathAbi.Pack("clamp", inputs.Value, inputs.Min, inputs.Max)
	if err != nil {
		return nil, fmt.Errorf("failed to pack clamp calldata: %w", err)
	}

	plug := signature.Plug{
		To:    mathContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}