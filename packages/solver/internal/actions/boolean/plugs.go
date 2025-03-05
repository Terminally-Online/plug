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
	Value      int64  `json:"value"`
	Comparison string `json:"comparison"`
	Threshold  int64  `json:"threshold"`
}

type TimeCompareInput struct {
	Time       int64  `json:"time"`
	Comparison string `json:"comparison"`
	Threshold  int64  `json:"threshold"`
}

type TimePropertyInput struct {
	Timestamp int64  `json:"timestamp"`
	Property  string `json:"property"`
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
	
	var functionName string
	switch operation {
	case "and":
		functionName = "isAnd"
	case "or":
		functionName = "isOr"
	case "not":
		functionName = "isNot"
	case "xor":
		functionName = "isXor"
	case "nand":
		functionName = "isNand"
	case "nor":
		functionName = "isNor"
	case "implies":
		functionName = "isImplies"
	default:
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

	comparison := strings.ToLower(inputs.Comparison)
	booleanContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}
	
	var functionName string
	switch comparison {
	case "equal":
		functionName = "isEqual"
	case "notequal":
		functionName = "isNotEqual"
	case "greaterthan":
		functionName = "isGreaterThan"
	case "greaterthanorequal":
		functionName = "isGreaterThanOrEqual"
	case "lessthan":
		functionName = "isLessThan"
	case "lessthanorequal":
		functionName = "isLessThanOrEqual"
	case "between":
		functionName = "isBetween"
	default:
		return nil, fmt.Errorf("unsupported number comparison: %s", inputs.Comparison)
	}
	
	var calldata []byte
	valueInt := big.NewInt(inputs.Value)
	thresholdInt := big.NewInt(inputs.Threshold)
	
	if comparison == "between" {
		minInt := big.NewInt(inputs.Value) // Using value as min
		calldata, err = booleanAbi.Pack(functionName, valueInt, minInt, thresholdInt)
	} else {
		calldata, err = booleanAbi.Pack(functionName, valueInt, thresholdInt)
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

	comparison := strings.ToLower(inputs.Comparison)
	booleanContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}
	
	var functionName string
	switch comparison {
	case "before":
		functionName = "isBeforeTime"
	case "after":
		functionName = "isAfterTime"
	case "between":
		functionName = "isBetweenTimes"
	case "sameday":
		functionName = "isSameDay"
	default:
		return nil, fmt.Errorf("unsupported time comparison: %s", inputs.Comparison)
	}
	
	var calldata []byte
	timeInt := big.NewInt(inputs.Time)
	thresholdInt := big.NewInt(inputs.Threshold)
	
	if comparison == "between" {
		startInt := big.NewInt(inputs.Time)
		calldata, err = booleanAbi.Pack(functionName, timeInt, startInt, thresholdInt)
	} else {
		calldata, err = booleanAbi.Pack(functionName, timeInt, thresholdInt)
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

	property := strings.ToLower(inputs.Property)
	booleanContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}
	
	var functionName string
	switch property {
	case "weekday":
		functionName = "isWeekday"
	case "weekend":
		functionName = "isWeekend"
	default:
		return nil, fmt.Errorf("unsupported time property: %s", inputs.Property)
	}
	
	timestampInt := big.NewInt(inputs.Timestamp)
	calldata, err := booleanAbi.Pack(functionName, timestampInt)
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
