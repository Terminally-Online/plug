package coil

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// Returns all return values from the given function in the ABI
func FindCoils(abi *abi.ABI, functionName string, outputName *string, outputIndex *string) ([]Update, error) {
	fn := abi.Methods[functionName]
	if fn.Name == "" {
		return nil, fmt.Errorf("function %s not found", functionName)
	}

	coils := make([]Update, len(fn.Outputs))
	var currentPos uint = 0
	for i, output := range fn.Outputs {

		length := getTypeLength(output.Type.String())

		// Skip if outputName is provided and doesn't match
		if outputName != nil && output.Name != *outputName {
			currentPos += length
			continue
		}

		// Skip if outputIndex is provided and doesn't match
		if outputIndex != nil && fmt.Sprint(i) != *outputIndex {
			currentPos += length
			continue
		}

		coils[i] = Update{
			Start: big.NewInt(int64(currentPos)),
			Slice: Slice{
				Name:   &output.Name,
				Index:  uint8(i),
				Start:  big.NewInt(int64(currentPos)),
				Length: big.NewInt(int64(length)),
			},
		}

		currentPos += length
	}

	return coils, nil
}

func getTypeLength(typeName string) uint {
	switch typeName {
	case "address":
		return 20 // 20 bytes for address
	case "uint256", "int256":
		return 32 // 32 bytes for 256-bit integers
	case "uint8", "int8":
		return 1 // 1 byte for 8-bit integers
	case "bool":
		return 1 // 1 byte for boolean
	case "bytes32":
		return 32 // 32 bytes for fixed-size byte array
	default:
		return 32 // Default to 32 bytes for unknown types
	}
}
