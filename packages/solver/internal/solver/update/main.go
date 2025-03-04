package update

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/onplug/plug/packages/solver/internal/solver/types"
// )

// // SliceForReturnValue represents a slice specification for a specific return value
// type SliceForReturnValue struct {
// 	// Index of the plug in the sequence that produces this return value
// 	PlugIndex uint8
// 	// Function name that produces the return value
// 	FunctionName string
// 	// Position of the return value in the function's return values (0-based)
// 	ReturnPosition int
// 	// Resulting slice specification that can be used in an update
// 	Slice types.Slice
// }

// // ExtractSlicesFromABI takes an ABI and function name and returns slice specifications
// // for each return value of the specified function
// func ExtractSlicesFromABI(abiJSON string, functionName string, plugIndex uint8) ([]SliceForReturnValue, error) {
// 	// Parse the ABI
// 	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse ABI: %w", err)
// 	}

// 	// Find the specified function
// 	method, exists := parsedABI.Methods[functionName]
// 	if !exists {
// 		return nil, fmt.Errorf("function %s not found in ABI", functionName)
// 	}

// 	// Collect slice specifications for each return value
// 	result := make([]SliceForReturnValue, len(method.Outputs))

// 	// Current byte position in the return data
// 	currentOffset := 0

// 	// Process each return value
// 	for i, output := range method.Outputs {
// 		// Calculate the length of this return value in bytes
// 		typeSize, err := getTypeSize(output.Type)
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Create the slice specification
// 		result[i] = SliceForReturnValue{
// 			PlugIndex:      plugIndex,
// 			FunctionName:   functionName,
// 			ReturnPosition: i,
// 			Slice: types.Slice{
// 				Index:  plugIndex,
// 				Start:  uint16(currentOffset),
// 				Length: uint16(typeSize),
// 			},
// 		}

// 		// Update the offset for the next return value
// 		currentOffset += typeSize
// 	}

// 	return result, nil
// }

// // getTypeSize returns the size in bytes of a Solidity type
// func getTypeSize(t abi.Type) (int, error) {
// 	switch t.T {
// 	case abi.IntTy, abi.UintTy:
// 		return 32, nil // All integers are padded to 32 bytes
// 	case abi.BoolTy:
// 		return 32, nil // Booleans are padded to 32 bytes
// 	case abi.AddressTy:
// 		return 32, nil // Addresses are padded to 32 bytes
// 	case abi.HashTy:
// 		return 32, nil // Bytes32/hash types are 32 bytes
// 	case abi.BytesTy, abi.StringTy, abi.SliceTy, abi.ArrayTy:
// 		// These are dynamic types that are encoded differently
// 		// For dynamic types, the actual data is stored elsewhere, and a 32-byte pointer is stored
// 		// in the return value. For simplicity, we'll return 32 here, but this would need refinement
// 		// for actual dynamic data extraction
// 		return 32, nil
// 	case abi.FixedBytesTy:
// 		// For fixed bytes, the size depends on the byte length, but is padded to 32 bytes
// 		return 32, nil
// 	case abi.TupleTy:
// 		// For tuples (structs), we need to calculate the size of all components
// 		// This is simplified here but would need more detailed implementation
// 		size := 0
// 		for _, comp := range t.TupleElems {
// 			compSize, err := getTypeSize(*comp)
// 			if err != nil {
// 				return 0, err
// 			}
// 			size += compSize
// 		}
// 		return size, nil
// 	default:
// 		return 0, fmt.Errorf("unsupported type: %v", t.String())
// 	}
// }

// // GenerateUpdateFromSlices creates an update specification that uses the specified slices
// // to update a target function's parameters at specific positions
// func GenerateUpdateFromSlices(
// 	targetSlices []SliceForReturnValue,
// 	targetFunctionABI string,
// 	targetFunctionName string,
// 	targetParamIndexes []int,
// ) ([]types.Update, error) {
// 	// Validate input
// 	if len(targetSlices) < len(targetParamIndexes) {
// 		return nil, fmt.Errorf("not enough slices for requested parameter updates")
// 	}

// 	// Parse the target function ABI
// 	parsedABI, err := abi.JSON(strings.NewReader(targetFunctionABI))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to parse target ABI: %w", err)
// 	}

// 	// Find the target function
// 	method, exists := parsedABI.Methods[targetFunctionName]
// 	if !exists {
// 		return nil, fmt.Errorf("target function %s not found in ABI", targetFunctionName)
// 	}

// 	// Calculate parameter positions
// 	paramOffsets := make([]int, len(method.Inputs))
// 	currentOffset := 4 // Start after the 4-byte function selector

// 	for i, input := range method.Inputs {
// 		paramOffsets[i] = currentOffset
// 		typeSize, err := getTypeSize(input.Type)
// 		if err != nil {
// 			return nil, err
// 		}
// 		currentOffset += typeSize
// 	}

// 	// Create updates
// 	updates := make([]types.Update, len(targetParamIndexes))

// 	for i, paramIndex := range targetParamIndexes {
// 		if paramIndex >= len(paramOffsets) {
// 			return nil, fmt.Errorf("parameter index %d out of bounds", paramIndex)
// 		}

// 		// Create update with the slice and correct start position
// 		// Adjust for the first byte being stripped in _plug function
// 		updates[i] = types.Update{
// 			Slice: targetSlices[i].Slice,
// 			Start: uint16(paramOffsets[paramIndex] - 1), // -1 to account for first byte stripping
// 		}
// 	}

// 	return updates, nil
// }

// // SimulateSlice provides a way to visualize what data would be extracted by a slice
// // This is useful for debugging and validation
// func SimulateSlice(returnData []byte, slice types.Slice) ([]byte, error) {
// 	if int(slice.Start)+int(slice.Length) > len(returnData) {
// 		return nil, fmt.Errorf("slice out of bounds: start %d + length %d exceeds data length %d",
// 			slice.Start, slice.Length, len(returnData))
// 	}

// 	return returnData[slice.Start : slice.Start+slice.Length], nil
// }
