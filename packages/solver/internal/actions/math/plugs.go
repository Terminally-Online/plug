package math

import (
	"encoding/json"
	"fmt"
	"math"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"strings"
)

// Input structs for the different math operations
type CalculateInputs struct {
	X         float64 `json:"x"`
	Operation string  `json:"operation"`
	Y         float64 `json:"y"`
}

type MinMaxInputs struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type AbsInputs struct {
	Value float64 `json:"value"`
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

func createResultPlug(_ float64) []signature.Plug {
	return []signature.Plug{}
}

// HandleCalculate implements a single calculation action that supports various operations
func HandleCalculate(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs CalculateInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal calculate inputs: %w", err)
	}

	// Perform the calculation based on the operation
	var result float64

	// Support both the exact operation from the options and lowercase variations
	switch inputs.Operation {
	case "+":
		result = inputs.X + inputs.Y
	case "-":
		result = inputs.X - inputs.Y
	case "*":
		result = inputs.X * inputs.Y
	case "÷", "/":
		if inputs.Y == 0 {
			return nil, fmt.Errorf("divide by zero error")
		}
		result = inputs.X / inputs.Y
	case "%":
		if inputs.Y == 0 {
			return nil, fmt.Errorf("modulo by zero error")
		}
		result = math.Mod(inputs.X, inputs.Y)
	default:
		// Try lowercase operation for text-based operations
		operation := strings.ToLower(inputs.Operation)
		switch operation {
		case "add", "+":
			result = inputs.X + inputs.Y
		case "subtract", "sub", "-":
			result = inputs.X - inputs.Y
		case "multiply", "mul", "*":
			result = inputs.X * inputs.Y
		case "divide", "div", "/", "÷":
			if inputs.Y == 0 {
				return nil, fmt.Errorf("divide by zero error")
			}
			result = inputs.X / inputs.Y
		case "modulo", "mod", "%":
			if inputs.Y == 0 {
				return nil, fmt.Errorf("modulo by zero error")
			}
			result = math.Mod(inputs.X, inputs.Y)
		default:
			return nil, fmt.Errorf("unsupported operation: %s", inputs.Operation)
		}
	}

	return createResultPlug(result), nil
}

// HandleMin returns the minimum of two numbers
func HandleMin(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs MinMaxInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal min inputs: %w", err)
	}

	result := math.Min(inputs.A, inputs.B)
	return createResultPlug(result), nil
}

// HandleMax returns the maximum of two numbers
func HandleMax(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs MinMaxInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal max inputs: %w", err)
	}

	result := math.Max(inputs.A, inputs.B)
	return createResultPlug(result), nil
}

// HandlePower raises a base to an exponent
func HandlePower(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs PowerInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal power inputs: %w", err)
	}

	result := math.Pow(inputs.Base, inputs.Exponent)
	return createResultPlug(result), nil
}

// HandleClamp clamps a value between min and max
func HandleClamp(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs ClampInputs
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal clamp inputs: %w", err)
	}

	if inputs.Min > inputs.Max {
		return nil, fmt.Errorf("min value exceeds max value")
	}

	// Clamp the value between min and max
	result := inputs.Value
	if result < inputs.Min {
		result = inputs.Min
	} else if result > inputs.Max {
		result = inputs.Max
	}

	return createResultPlug(result), nil
}
