package types

import (
	"encoding/json"
	"fmt"
)

type ProtocolMetadata struct {
	Icon string   `json:"icon"`
	Tags []string `json:"tags"`
}

type ProtocolSchema struct {
	Metadata ProtocolMetadata  `json:"metadata"`
	Schema   map[Action]Schema `json:"schema"`
}

type Option struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Info  string `json:"info,omitempty"`
	Icon  string `json:"icon,omitempty"`
}

type Schema struct {
	Sentence string                `json:"sentence"`
	Options  map[int]SchemaOptions `json:"options,omitempty"`
}

type SchemaOptions struct {
	Simple  []Option            `json:"-"`
	Complex map[string][]Option `json:"-"`
}

func (o SchemaOptions) MarshalJSON() ([]byte, error) {
	if o.Simple != nil {
		return json.Marshal(o.Simple)
	}
	return json.Marshal(o.Complex)
}

func (o *SchemaOptions) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &o.Simple); err == nil {
		return nil
	}

	if err := json.Unmarshal(data, &o.Complex); err == nil {
		return nil
	}

	return fmt.Errorf("invalid options format")
}

var (
	BaseBorowFields = []Option{
		{Label: "Borrow", Name: "Borrow", Value: "-1"},
		{Label: "Deposit", Name: "Deposit", Value: "1"},
	}

	BaseThresholdFields = []Option{
		{Label: "less than", Name: "Less Than", Value: "-1"},
		{Label: "greater than", Name: "Greater Than", Value: "1"},
	}
)
