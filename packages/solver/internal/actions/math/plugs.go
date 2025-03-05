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
	X         float64 `json:"x"`
	Operation string  `json:"operation"`
	Y         float64 `json:"y"`
}

type MinMaxInputs struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type PowerInputs struct {
	Base     float64 `json:"base"`
	Exponent float64 `json:"exponent"`
}

type ClampInputs struct {
	Value float64 `json:"value"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
}

func HandleCalculate(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs CalculateInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal calculate inputs: %w", err)
	}

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	a := big.NewInt(int64(inputs.X))
	b := big.NewInt(int64(inputs.Y))
	
	var functionName string
	switch inputs.Operation {
	case "+", "add":
		functionName = "add"
	case "-", "subtract", "sub":
		functionName = "subtract"
	case "*", "multiply", "mul":
		functionName = "multiply"
	case "÷", "/", "divide", "div":
		if inputs.Y == 0 {
			return nil, fmt.Errorf("divide by zero error")
		}
		functionName = "divide"
	case "%", "modulo", "mod":
		if inputs.Y == 0 {
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
			if inputs.Y == 0 {
				return nil, fmt.Errorf("divide by zero error")
			}
			functionName = "divide"
		case "modulo", "mod", "%":
			if inputs.Y == 0 {
				return nil, fmt.Errorf("modulo by zero error")
			}
			functionName = "modulo"
		default:
			return nil, fmt.Errorf("unsupported operation: %s", inputs.Operation)
		}
	}
	
	calldata, err := mathAbi.Pack(functionName, a, b)
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

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	a := big.NewInt(int64(inputs.A))
	b := big.NewInt(int64(inputs.B))
	
	calldata, err := mathAbi.Pack("min", a, b)
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

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	a := big.NewInt(int64(inputs.A))
	b := big.NewInt(int64(inputs.B))
	
	calldata, err := mathAbi.Pack("max", a, b)
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

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	base := big.NewInt(int64(inputs.Base))
	exponent := big.NewInt(int64(inputs.Exponent))
	
	calldata, err := mathAbi.Pack("power", base, exponent)
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

	if inputs.Min > inputs.Max {
		return nil, fmt.Errorf("min value exceeds max value")
	}

	mathContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["math"])
	mathAbi, err := plug_math.PlugMathMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugMath ABI: %w", err)
	}
	
	value := big.NewInt(int64(inputs.Value))
	min := big.NewInt(int64(inputs.Min))
	max := big.NewInt(int64(inputs.Max))
	
	calldata, err := mathAbi.Pack("clamp", value, min, max)
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
