package actions

import (
	"encoding/json"
	"fmt"
)

type ProtocolMetadata struct {
	Icon   string   `json:"icon"`
	Tags   []string `json:"tags"`
	Chains []int    `json:"chains"`
}

type ProtocolSchema struct {
	Metadata ProtocolMetadata  `json:"metadata"`
	Schema   map[string]Schema `json:"schema"`
}

type ChainSchema struct {
	Type   string `default:"action" json:"type"`
	Schema Schema `json:"schema"`
}

type Schema struct {
	Type     string                `default:"action" json:"type"`
	Sentence string                `json:"sentence"`
	Options  map[int]Options `json:"options,omitempty"`
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
