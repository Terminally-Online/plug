package coil

import (
	"math/big"
	"solver/bindings/plug_router"
)

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
