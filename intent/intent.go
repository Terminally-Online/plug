package intent

import (
	"slices"
	"solver/utils"
)

type IntentRequest struct {
	ChainId int      `json:"chainId"`
	From    string   `json:"from"`
	Actions []Action `json:"actions"`
}

var (
	SupportedChains = []int{1}
	MinActions      = 1
	MaxActions      = 10
)

func (i IntentRequest) Validate() error {
	if !slices.Contains(SupportedChains, i.ChainId) {
		return utils.ErrInvalidChainId("chainId", i.ChainId)
	}

	if !utils.IsAddress(i.From) {
		return utils.ErrInvalidAddress("from", i.From)
	}

	if len(i.Actions) < MinActions || len(i.Actions) > MaxActions {
		return utils.ErrInvalidArrayLength("actions", &MinActions, &MaxActions)
	}

	return nil
}

type IntentResponse struct {
	Request      IntentRequest       `json:"request"`
	Transactions []utils.Transaction `json:"transactions"`
}
