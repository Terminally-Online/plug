package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/plug_boolean"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type NumberComparisonRequest struct {
	A          *big.Int `json:"a"`
	Comparison string   `json:"comparison"`
	B          *big.Int `json:"b"`
}


func CompareNumbers(lookup *actions.SchemaLookup[NumberComparisonRequest]) ([]signature.Plug, error) {
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

