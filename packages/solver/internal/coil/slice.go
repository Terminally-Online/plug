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
