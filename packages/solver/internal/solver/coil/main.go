package coil

import (
	"encoding/json"
	"fmt"
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

// ContractABI represents the full ABI of a contract
type ContractABI []ABIFunction

// GetFunctionReturnTypes takes a contract ABI and function name and returns the output parameters
func GetFunctionReturnTypes(abiJSON string, functionName string) ([]ABIParameter, error) {
	var abi ContractABI

	if err := json.Unmarshal([]byte(abiJSON), &abi); err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Find the function in the ABI
	for _, function := range abi {
		if function.Type == "function" && function.Name == functionName {
			return function.Outputs, nil
		}
	}

	return nil, fmt.Errorf("function %s not found in ABI", functionName)
}

// Additional helper methods

// GetFunctionByName returns the full function definition from the ABI
func (abi ContractABI) GetFunctionByName(name string) (*ABIFunction, error) {
	for _, function := range abi {
		if function.Type == "function" && function.Name == name {
			return &function, nil
		}
	}
	return nil, fmt.Errorf("function %s not found", name)
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

// Add these types at the top of the file after the existing types

// Coil represents a function's return value parsing configuration
type Coil struct {
	Function *ABIFunction
	Updates  []Update
}

// Update represents how to parse a return value
type Update struct {
	Start uint
	Slice Slice
}

// Slice represents the location and size of a return value in the data
type Slice struct {
	Index  uint8
	Start  uint
	Length uint
}

// Add this function after the existing functions

// NewCoil creates a new Coil for parsing return values from the given function
func NewCoil(abiString string, functionName string) (*Coil, error) {
	var abi ContractABI
	if err := json.Unmarshal([]byte(abiString), &abi); err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	fn, err := abi.GetFunctionByName(functionName)
	if err != nil {
		return nil, err
	}

	coil := &Coil{
		Function: fn,
		Updates:  make([]Update, len(fn.Outputs)),
	}

	// Track the current position in the return data
	var currentPos uint = 0

	// Create an Update and Slice for each output parameter
	for i, output := range fn.Outputs {
		length := getTypeLength(output.Type)

		coil.Updates[i] = Update{
			Start: currentPos,
			Slice: Slice{
				Index:  uint8(i),
				Start:  currentPos,
				Length: length,
			},
		}

		currentPos += length
	}

	return coil, nil
}

// getTypeLength returns the length in bytes of the given Solidity type
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
	// Add more types as needed
	default:
		return 32 // Default to 32 bytes for unknown types
	}
}

// Add a helper method to get the updates
func (c *Coil) GetUpdates() []Update {
	return c.Updates
}
