package coil

import (
	"fmt"
	"math/big"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ABIFunction represents a function definition in an ABI
type ABIFunction struct {
	Name            string         `json:"name"`
	Inputs          []ABIParameter `json:"inputs"`
	Outputs         []ABIParameter `json:"outputs"`
	StateMutability string         `json:"stateMutability"`
	Type            string         `json:"type"`
}

// ABIParameter represents an input or output parameter in an ABI function
type ABIParameter struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	InternalType string `json:"internalType,omitempty"`
	Indexed      bool   `json:"indexed,omitempty"`
}

// Coil represents a function's return value parsing configuration
type Coil struct {
	Function *ABIFunction
	Updates  []signature.Update
}

// GetFunctionReturnTypes takes a contract ABI and function name and returns the output parameters
func GetFunctionReturnTypes(abi *abi.ABI, functionName string) ([]ABIParameter, error) {
	fn := abi.Methods[functionName]
	if fn.Name == "" {
		return nil, fmt.Errorf("function %s not found in ABI", functionName)
	}
	return convertABIArgs(fn.Outputs), nil
}

// GetFunctionByName returns the full function definition from the ABI
func GetFunctionByName(abi *abi.ABI, name string) (*ABIFunction, error) {
	fn := abi.Methods[name]
	if fn.Name == "" {
		return nil, fmt.Errorf("function %s not found", name)
	}

	return &ABIFunction{
		Name:            fn.Name,
		Inputs:          convertABIArgs(fn.Inputs),
		Outputs:         convertABIArgs(fn.Outputs),
		StateMutability: string(fn.StateMutability),
		Type:            "function",
	}, nil
}

// HasNamedOutputs checks if the function has named return values
func (f ABIFunction) HasNamedOutputs() bool {
	for _, output := range f.Outputs {
		if output.Name != "" {
			return true
		}
	}
	return false
}

// GetOutputNames returns a slice of output parameter names
func (f ABIFunction) GetOutputNames() []string {
	names := make([]string, len(f.Outputs))
	for i, output := range f.Outputs {
		names[i] = output.Name
		if names[i] == "" {
			names[i] = fmt.Sprintf("output%d", i)
		}
	}
	return names
}

// NewCoil creates a new Coil for parsing return values from the given function
func NewCoil(abi *abi.ABI, functionName string) (*Coil, error) {
	fn := abi.Methods[functionName]
	if fn.Name == "" {
		return nil, fmt.Errorf("function %s not found", functionName)
	}

	coil := &Coil{
		Function: &ABIFunction{
			Name:            fn.Name,
			Inputs:          convertABIArgs(fn.Inputs),
			Outputs:         convertABIArgs(fn.Outputs),
			StateMutability: string(fn.StateMutability),
			Type:            "function",
		},
		Updates: make([]signature.Update, len(fn.Outputs)),
	}

	var currentPos uint = 0
	for i, output := range fn.Outputs {
		length := getTypeLength(output.Type.String())

		coil.Updates[i] = signature.Update{
			Start: big.NewInt(int64(currentPos)),
			Slice: signature.Slice{
				Name:   &output.Name,
				Index:  uint8(i),
				Start:  big.NewInt(int64(currentPos)),
				Length: big.NewInt(int64(length)),
			},
		}

		currentPos += length
	}

	return coil, nil
}

func convertABIArgs(args abi.Arguments) []ABIParameter {
	params := make([]ABIParameter, len(args))
	for i, arg := range args {
		params[i] = ABIParameter{
			Name:         arg.Name,
			Type:         arg.Type.String(),
			InternalType: arg.Type.String(),
		}
	}
	return params
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

// GetUpdates returns the slice of Updates
func (c *Coil) GetUpdates() []signature.Update {
	return c.Updates
}
