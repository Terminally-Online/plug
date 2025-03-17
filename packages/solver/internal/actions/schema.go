package actions

import (
	"solver/internal/bindings/references"
	"solver/internal/solver/coil"
)

type Chain struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type ProtocolMetadata struct {
	Icon   string                `json:"icon"`
	Tags   []string              `json:"tags"`
	Chains []*references.Network `json:"chains"`
}

type ProtocolSchema struct {
	Metadata ProtocolMetadata  `json:"metadata"`
	Schema   map[string]Schema `json:"schema"`
}

type ChainSchema struct {
	Type           string        `default:"action" json:"type"`
	Schema         Schema        `json:"schema"`
	IsUserSpecific bool          `json:"-"`
	LinkedInputs   []coil.Update `json:"-"` // Added to support linked inputs directly at schema level
}

type Schema struct {
	Type           string          `default:"action" json:"type"`
	Sentence       string          `json:"sentence"`
	Options        map[int]Options `json:"options,omitempty"`
	Coils          []coil.Update   `json:"coils,omitempty"`
	IsUserSpecific bool            `json:"-"`
}
