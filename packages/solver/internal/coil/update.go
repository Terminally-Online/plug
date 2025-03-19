package coil

import (
	"math/big"
	"solver/bindings/plug_router"
)

type Update struct {
	Start *big.Int `json:"start"` // Starting position where the slice should be inserted
	Slice Slice    `json:"slice"` // The slice specification
}

func (u Update) Wrap() plug_router.PlugTypesLibUpdate {
	return plug_router.PlugTypesLibUpdate{
		Start: u.Start,
		Slice: u.Slice.Wrap(),
	}
}
