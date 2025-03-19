package actions

import (
	"encoding/json"
	"fmt"
)

type Options struct {
	Simple  []Option            `json:"simple,omitempty"`
	Complex map[string][]Option `json:"complex,omitempty"`
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
