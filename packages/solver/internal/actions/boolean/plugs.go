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
	A          *big.Int `json:"a"`
	Comparison string   `json:"comparison"`
	B          *big.Int `json:"b"`
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

func HandleLogicOperation(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs LogicOperationInput
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal logic operation inputs: %w", err)
	}

	operation := strings.ToLower(inputs.Operation)
	booleanContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["boolean"])
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

func HandleCompareNumbers(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs NumberComparisonInput
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal number comparison inputs: %w", err)
	}

	comparison := strings.ToLower(inputs.Comparison)
	booleanContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}

	comparisonFunctions := map[string]string{
		"equal":              "isEqual",
		"notequal":           "isNotEqual",
		"greaterthan":        "isGreaterThan",
		"greaterthanorequal": "isGreaterThanOrEqual",
		"lessthan":           "isLessThan",
		"lessthanorequal":    "isLessThanOrEqual",
	}

	functionName, supported := comparisonFunctions[comparison]
	if !supported {
		return nil, fmt.Errorf("unsupported number comparison: %s", inputs.Comparison)
	}

	calldata, err := booleanAbi.Pack(functionName, inputs.A, inputs.B)

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

func HandleCompareTimes(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs TimeCompareInput
	if err := json.Unmarshal(raw, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal time comparison inputs: %w", err)
	}

	comparison := strings.ToLower(inputs.Comparison)
	booleanContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}

	comparisonFunctions := map[string]string{
		"before":  "isBeforeTime",
		"after":   "isAfterTime",
		"sameday": "isSameDay",
	}

	functionName, supported := comparisonFunctions[comparison]
	if !supported {
		return nil, fmt.Errorf("unsupported time comparison: %s", inputs.Comparison)
	}

	calldata, err := booleanAbi.Pack(functionName, inputs.Time, inputs.Threshold)

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

func HandleCheckTimeProperty(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
	var inputs TimePropertyInput
	if err := json.Unmarshal(raw, &inputs); err != nil {
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
	booleanContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["boolean"])
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
