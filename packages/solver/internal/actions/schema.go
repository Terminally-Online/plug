package actions

import (
	"encoding/json"
	"fmt"
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
	IsUserSpecific bool          `json:"isUserSpecific,omitempty"`
	Schema         Schema        `json:"schema"`
	LinkedInputs   []coil.Update `json:"linkedInputs,omitempty"` // Added to support linked inputs directly at schema level
}

type Schema struct {
	Type           string          `default:"action" json:"type"`
	Sentence       string          `json:"sentence"`
	IsUserSpecific bool            `json:"isUserSpecific,omitempty"`
	Options        map[int]Options `json:"options,omitempty"`
	Coils          []coil.Update   `json:"coils,omitempty"`
}

func (o Options) MarshalJSON() ([]byte, error) {
	if o.Simple != nil {
		return json.Marshal(o.Simple)
	}
	return json.Marshal(o.Complex)
}

func (o *Options) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.Simple); err == nil {
		return nil
	}

	if err := json.Unmarshal(data, &o.Complex); err == nil {
		return nil
	}

	return fmt.Errorf("invalid options format")
}

var (
	BaseLendActionTypeFields = []Option{
		{Label: "Borrow", Name: "Borrow", Value: "-1"},
		{Label: "Deposit", Name: "Deposit", Value: "1"},
	}

	BaseThresholdFields = []Option{
		{Label: "less than", Name: "Less Than", Value: "-1"},
		{Label: "greater than", Name: "Greater Than", Value: "1"},
	}
)
