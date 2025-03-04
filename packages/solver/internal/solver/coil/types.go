package coil

import (
	"math/big"
	"solver/bindings/plug_router"
)

// Slice represents a specification for extracting a portion of a function's output data.
// This maps directly to the Slice struct in Plug.Socket.sol.
type Slice struct {
	Name   *string  `json:"name,omitempty"`   // Optional name for reference
	Index  uint8    `json:"index"`            // Index of the plug in the sequence that produces this data
	Start  *big.Int `json:"start"`            // Starting byte position within the result data
	Length *big.Int `json:"length"`           // Length of bytes to extract
	Type   string   `json:"type"`             // Solidity type of the data (for documentation)
	TypeId *uint8   `json:"typeId,omitempty"` // Type ID matching Plug.Socket.sol typeId values
}

// Wrap converts the Slice to the format expected by the plug_router bindings
func (s Slice) Wrap() plug_router.PlugTypesLibSlice {
	var typeId uint8 = 0 // Default to static type if none provided
	if s.TypeId != nil {
		typeId = *s.TypeId
	}

	return plug_router.PlugTypesLibSlice{
		Index:  s.Index,
		Start:  s.Start,
		Length: s.Length,
		TypeId: typeId,
	}
}

// Update represents an instruction to modify a transaction's calldata
// by inserting data from a previous transaction's result.
type Update struct {
	Start *big.Int `json:"start"` // Starting position where the slice should be inserted
	Slice Slice    `json:"slice"` // The slice specification
}

// Wrap converts the Update to the format expected by the plug_router bindings
func (u Update) Wrap() plug_router.PlugTypesLibUpdate {
	return plug_router.PlugTypesLibUpdate{
		Start: u.Start,
		Slice: u.Slice.Wrap(),
	}
}

// ABIFunction represents a function definition in an ABI
// Used for JSON parsing/serialization when working with raw ABI data
type ABIFunction struct {
	Name            string         `json:"name,omitempty"`
	Inputs          []ABIParameter `json:"inputs"`
	Outputs         []ABIParameter `json:"outputs"`
	StateMutability string         `json:"stateMutability"`
	Type            string         `json:"type"`
}

// ABIParameter represents an input or output parameter in an ABI function
// Used for JSON parsing/serialization when working with raw ABI data
type ABIParameter struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	InternalType string `json:"internalType,omitempty"`
	Indexed      bool   `json:"indexed,omitempty"`
}
