package coil

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	TypeIDStatic      uint8 = 0 // Static types (uint256, address, etc.)
	TypeIDArray       uint8 = 1 // Dynamic arrays (uint256[], address[], etc.)
	TypeIDString      uint8 = 2 // String and bytes (identical in typeId)
	TypeIDStruct      uint8 = 3 // Structs with dynamic fields
	TypeIDNestedArray uint8 = 4 // Nested arrays (uint256[][], etc.)
	TypeIDMapping     uint8 = 5 // Mapping-like structures

	WordSize uint = 32 // Standard EVM word size (32 bytes)
)

// GetCoils analyzes a function's outputs and creates the appropriate slices for data extraction
// from contract call results. This enables reactive transaction building by making outputs
// available as inputs to subsequent transactions.
//
// Parameters:
//   - abi: The parsed ABI containing the function definition
//   - functionName: The name of the function whose outputs we want to analyze
//   - outputName: Optional filter to only return a specific named output
//   - outputIndex: Optional filter to only return a specific output by index
//
// Returns:
//   - []Update: A list of updates that can be used to extract data from contract call results
//   - error: Any error encountered during analysis
func GetCoilSlices(abi *abi.ABI, functionName string, outputName *string, outputIndex *string) ([]Slice, error) {
	fn, exists := abi.Methods[functionName]
	if !exists {
		return nil, fmt.Errorf("function %s not found in ABI", functionName)
	}

	slices := make([]Slice, 0, len(fn.Outputs))
	var currentPos uint = 0

	for i, output := range fn.Outputs {
		typeInfo := GetTypeInfo(output.Type)

		if (outputName != nil && output.Name != *outputName) ||
			(outputIndex != nil && fmt.Sprint(i) != *outputIndex) {
			currentPos += typeInfo.Length
			continue
		}

		name := output.Name
		if name == "" {
			name = fmt.Sprintf("output%d", i)
		}

		typeId := typeInfo.TypeId
		slices = append(slices, Slice{
			Name:   &name,
			Index:  uint8(i),
			Start:  big.NewInt(int64(currentPos)),
			Length: big.NewInt(int64(typeInfo.Length)),
			Type:   output.Type.String(),
			TypeId: &typeId,
		})

		currentPos += typeInfo.Length
	}

	return slices, nil
}

// GetInputPosition analyzes a function's inputs and returns the position where a specific
// input parameter should be placed in the transaction data.
//
// The position calculation automatically accounts for the function selector that
// will always come from the ABI that is defined inside the action function. We DO NOT
// slide the position over to account for the encoding that takes place when building the
// final transaction data.
//
// Parameters:
//   - abi: The parsed ABI containing the function definition
//   - functionName: The name of the function whose inputs we want to analyze
//   - inputName: Optional filter to find position by parameter name
//   - inputIndex: Optional filter to find position by parameter index
//
// Returns:
//   - *big.Int: The starting position for the specified input
//   - error: Any error encountered during analysis
func GetABICoilPosition(abi *abi.ABI, functionName string, inputName *string, inputIndex *string) (*big.Int, error) {
	fn, exists := abi.Methods[functionName]
	if !exists {
		return nil, fmt.Errorf("function %s not found in ABI", functionName)
	}

	var currentPos uint = 4

	for i, input := range fn.Inputs {
		if (inputName != nil && input.Name == *inputName) ||
			(inputIndex != nil && fmt.Sprint(i) == *inputIndex) {
			return big.NewInt(int64(currentPos)), nil
		}

		typeInfo := GetTypeInfo(input.Type)
		currentPos += typeInfo.Length
	}

	return nil, fmt.Errorf("input parameter not found")
}

func GetArgumentsCoilPosition(arguments *abi.Arguments, inputName *string, inputIndex *string) (*big.Int, error) {
	if arguments == nil {
		return nil, fmt.Errorf("arguments cannot be nil")
	}

	var currentPos uint = 0

	for i, input := range *arguments {
		if (inputName != nil && input.Name == *inputName) ||
			(inputIndex != nil && fmt.Sprint(i) == *inputIndex) {
			return big.NewInt(int64(currentPos)), nil
		}

		typeInfo := GetTypeInfo(input.Type)
		currentPos += typeInfo.Length
	}

	return nil, fmt.Errorf("input parameter not found")
}
