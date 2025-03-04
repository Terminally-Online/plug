package coil

import (
	"math/big"
	"solver/bindings/plug_router"
)

// ABIFunction represents a function definition in an ABI
type ABIFunction struct {
	Name            string         `json:"name,omitempty"`
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

type Slice struct {
	Name   *string  `json:"name,omitempty"` // Optional name for reference
	Index  uint8    `json:"index"`          // Index of the plug in the sequence that produces this data
	Start  *big.Int `json:"start"`          // Starting byte position within the result data
	Length *big.Int `json:"length"`         // Length of bytes to extract
	Type   string   `json:"type"`           // Type of the data
}

func (s Slice) Wrap() plug_router.PlugTypesLibSlice {
	return plug_router.PlugTypesLibSlice{
		Index:  s.Index,
		Start:  s.Start,
		Length: s.Length,
	}
}

type Update struct {
	Start *big.Int `json:"start"` // Starting position where the slice should be inserted
	Slice Slice    `json:"slice"` // The slice specification
}

func (p Update) Wrap() plug_router.PlugTypesLibUpdate {
	return plug_router.PlugTypesLibUpdate{
		Start: p.Start,
		Slice: p.Slice.Wrap(),
	}
}
