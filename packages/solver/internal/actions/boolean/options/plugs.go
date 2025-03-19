package options

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

type LogicOperationRequest struct {
	A         bool   `json:"a"`
	Operation string `json:"operation"`
	B         bool   `json:"b"`
}

type NumberComparisonRequest struct {
	A          *big.Int `json:"a"`
	Comparison string   `json:"comparison"`
	B          *big.Int `json:"b"`
}

func LogicOperation(lookup *actions.SchemaLookup[LogicOperationRequest]) ([]signature.Plug, error) {
	operation := strings.ToLower(lookup.Inputs.Operation)
	booleanContract := common.HexToAddress(references.Networks[lookup.ChainId].References["plug"]["boolean"])
	booleanAbi, err := plug_boolean.PlugBooleanMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugBoolean ABI: %w", err)
	}

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
		return nil, fmt.Errorf("unsupported logical operation: %s", lookup.Inputs.Operation)
	}

	var calldata []byte
	if operation == "not" {
		calldata, err = booleanAbi.Pack(functionName, lookup.Inputs.A)
	} else {
		calldata, err = booleanAbi.Pack(functionName, lookup.Inputs.A, lookup.Inputs.B)
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

func CompareNumbers(lookup *actions.SchemaLookup[NumberComparisonRequest], raw json.RawMessage) ([]signature.Plug, error) {
	comparison := strings.ToLower(lookup.Inputs.Comparison)
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
		return nil, fmt.Errorf("unsupported number comparison: %s", lookup.Inputs.Comparison)
	}

	calldata, err := booleanAbi.Pack(functionName, lookup.Inputs.A, lookup.Inputs.B)
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
