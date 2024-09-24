package intent

import (
	"solver/utils"
)

type IntentRequest struct {
	ChainId int      `json:"chainId"`
	From    string   `json:"from"`
	Actions []Action `json:"actions"`
}

func (i IntentRequest) Validate() error {
	if !utils.IsSupportedChain(i.ChainId) {
		return utils.ErrInvalidChainId("chainId", i.ChainId)
	}

	if !utils.IsAddress(i.From) {
		return utils.ErrInvalidAddress("from", i.From)
	}

	if len(i.Actions) < utils.MinActions || len(i.Actions) > utils.MaxActions {
		return utils.ErrInvalidArrayLength("actions", &utils.MinActions, &utils.MaxActions)
	}

	return nil
}

type IntentResponse struct {
	Request      IntentRequest       `json:"request"`
	Transactions []utils.Transaction `json:"transactions"`
}
