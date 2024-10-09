package intent

import (
	"solver/types"
	"solver/utils"
)

type IntentRequest struct {
	ChainId int            `json:"chainId"`
	From    string         `json:"from"`
	Solver  *string        `json:"solver"`
	Actions []types.Action `json:"actions"`
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

type Plugs struct {
	Socket       string              `json:"socket"`
	Transactions []types.Transaction `json:"plugs"`
	Solver       string              `json:"solver"`
	Salt         string              `json:"salt"`
}

type IntentResponse struct {
	Plugs     Plugs         `json:"plugs"`
	Signature string        `json:"signature"`
}
