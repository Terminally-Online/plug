package solver

import (
	"encoding/json"
)

type IntentRequest struct {
	ChainId int               `json:"chainId"`
	From    string            `json:"from"`
	Inputs  []json.RawMessage `json:"inputs"`
}

type KillResponse struct {
	Killed bool `json:"killed"`
}
