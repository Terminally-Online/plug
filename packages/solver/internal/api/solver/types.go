package solver

import (
	"encoding/json"
	"fmt"
)

type IntentRequest struct {
	ChainId int               `json:"chainId"`
	From    string            `json:"from"`
	Inputs  []json.RawMessage `json:"inputs"`
}

func (r *IntentRequest) Validate() error { 
	if r.ChainId == 0 {
		return fmt.Errorf("'chainId' is required")
	}

	if r.From == "" {
		return fmt.Errorf("'from' is required")
	}

	return nil
}

func (req *IntentRequest) UnmarshalJSON(data []byte) error {
	type AuxIntentRequest IntentRequest
	aux := (*AuxIntentRequest)(req)
	
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("invalid request body: %w", err)
	}
	
	if err := req.Validate(); err != nil {
		return err
	}
	
	return nil
}

type KillResponse struct {
	Killed bool `json:"killed"`
}
