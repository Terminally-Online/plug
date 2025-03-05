package boolean

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/bindings/plug_boolean"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type LogicOperationInput struct {
	A         bool   `json:"a"`
	Operation string `json:"operation"`
	B         bool   `json:"b"`
}

type NumberComparisonInput struct {
	Value      *big.Int `json:"value"`
	Comparison string   `json:"comparison"`
	Threshold  *big.Int `json:"threshold"`
	Min        *big.Int `json:"min,omitempty"`      // Used for "between" comparison
	Max        *big.Int `json:"max,omitempty"`      // Used for "between" comparison
}

type TimeCompareInput struct {
	Time       *big.Int `json:"time"`
	Comparison string   `json:"comparison"`
	Threshold  *big.Int `json:"threshold"`
	StartTime  *big.Int `json:"startTime,omitempty"` // Used for "between" comparison
	EndTime    *big.Int `json:"endTime,omitempty"`   // Used for "between" comparison
}

type TimePropertyInput struct {
	Timestamp *big.Int `json:"timestamp"`
	Property  string   `json:"property"`
}

func HandleLogicOperation(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs LogicOperationInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal logic operation inputs: %w", err)
	}

	operation := strings.ToLower(inputs.Operation)
	booleanContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}
	
	// Map of supported logical operations to their function names
	operationFunctions := map[string]string{
		"and":     "isAnd",
		"or":      "isOr",
		"not":     "isNot",
		"xor":     "isXor",
		"nand":    "isNand",
		"nor":     "isNor",
		"implies": "isImplies",
	}
	
	functionName, supported := operationFunctions[operation]
	if !supported {
		return nil, fmt.Errorf("unsupported logical operation: %s", inputs.Operation)
	}
	
	var calldata []byte
	if operation == "not" {
		calldata, err = booleanAbi.Pack(functionName, inputs.A)
	} else {
		calldata, err = booleanAbi.Pack(functionName, inputs.A, inputs.B)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", functionName, err)
	}

	plug := signature.Plug{
		To:    booleanContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}

func HandleCompareNumbers(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs NumberComparisonInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal number comparison inputs: %w", err)
	}

	// Validate required fields are not nil
	if inputs.Value == nil {
		return nil, fmt.Errorf("value must be provided")
	}
	
	// For non-between comparisons, we need a threshold
	comparison := strings.ToLower(inputs.Comparison)
	if comparison != "between" && inputs.Threshold == nil {
		return nil, fmt.Errorf("threshold must be provided for %s comparison", comparison)
	}

	booleanContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}
	
	// Map of supported comparison operations to their function names
	comparisonFunctions := map[string]string{
		"equal":             "isEqual",
		"notequal":          "isNotEqual",
		"greaterthan":       "isGreaterThan",
		"greaterthanorequal": "isGreaterThanOrEqual",
		"lessthan":          "isLessThan",
		"lessthanorequal":   "isLessThanOrEqual",
		"between":           "isBetween",
	}
	
	functionName, supported := comparisonFunctions[comparison]
	if !supported {
		return nil, fmt.Errorf("unsupported number comparison: %s", inputs.Comparison)
	}
	
	var calldata []byte
	
	if comparison == "between" {
		// For "between" comparison, check if we have Min and Max values
		if inputs.Min == nil || inputs.Max == nil {
			return nil, fmt.Errorf("for 'between' comparison, min and max values must be provided")
		}
		
		// Validate min <= max
		if inputs.Min.Cmp(inputs.Max) > 0 {
			return nil, fmt.Errorf("min value cannot be greater than max value")
		}
		
		calldata, err = booleanAbi.Pack(functionName, inputs.Value, inputs.Min, inputs.Max)
	} else {
		calldata, err = booleanAbi.Pack(functionName, inputs.Value, inputs.Threshold)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", functionName, err)
	}

	plug := signature.Plug{
		To:    booleanContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}

func HandleCompareTimes(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs TimeCompareInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal time comparison inputs: %w", err)
	}

	// Validate required fields are not nil
	if inputs.Time == nil {
		return nil, fmt.Errorf("time value must be provided")
	}
	
	// For non-between comparisons, we need a threshold
	comparison := strings.ToLower(inputs.Comparison)
	if comparison != "between" && inputs.Threshold == nil {
		return nil, fmt.Errorf("threshold must be provided for %s comparison", comparison)
	}

	booleanContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}
	
	// Map of supported time comparison operations to their function names
	comparisonFunctions := map[string]string{
		"before":   "isBeforeTime",
		"after":    "isAfterTime",
		"between":  "isBetweenTimes",
		"sameday":  "isSameDay",
	}
	
	functionName, supported := comparisonFunctions[comparison]
	if !supported {
		return nil, fmt.Errorf("unsupported time comparison: %s", inputs.Comparison)
	}
	
	var calldata []byte
	
	if comparison == "between" {
		// For "between" comparison with time, check if we have StartTime and EndTime values
		if inputs.StartTime == nil || inputs.EndTime == nil {
			return nil, fmt.Errorf("for 'between' time comparison, startTime and endTime values must be provided")
		}
		
		// Validate startTime <= endTime
		if inputs.StartTime.Cmp(inputs.EndTime) > 0 {
			return nil, fmt.Errorf("startTime cannot be after endTime")
		}
		
		calldata, err = booleanAbi.Pack(functionName, inputs.Time, inputs.StartTime, inputs.EndTime)
	} else {
		calldata, err = booleanAbi.Pack(functionName, inputs.Time, inputs.Threshold)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", functionName, err)
	}

	plug := signature.Plug{
		To:    booleanContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}

func HandleCheckTimeProperty(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs TimePropertyInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal time property inputs: %w", err)
	}

	// Validate required fields are not nil
	if inputs.Timestamp == nil {
		return nil, fmt.Errorf("timestamp must be provided")
	}

	// Validate property is not empty
	if inputs.Property == "" {
		return nil, fmt.Errorf("property must be provided")
	}

	property := strings.ToLower(inputs.Property)
	booleanContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}
	
	// Map of supported time property checks to their function names
	propertyFunctions := map[string]string{
		"weekday": "isWeekday",
		"weekend": "isWeekend",
	}
	
	functionName, supported := propertyFunctions[property]
	if !supported {
		return nil, fmt.Errorf("unsupported time property: %s", inputs.Property)
	}
	
	calldata, err := booleanAbi.Pack(functionName, inputs.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to pack %s calldata: %w", functionName, err)
	}

	plug := signature.Plug{
		To:    booleanContract,
		Data:  calldata,
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}
