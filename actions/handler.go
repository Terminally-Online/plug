package actions

import (
	"encoding/json"
	"solver/types"

	"github.com/ethereum/go-ethereum/ethclient"
)

type HandlerParams struct {
	Provider *ethclient.Client
	ChainId  int
	From     string
}

type BaseProtocolHandler interface {
	SupportedActions() []types.Action
	SupportedChains() []int
	GetSchema(action types.Action) (types.ActionSchema, error)
	GetTransaction(action types.Action, rawInputs json.RawMessage, params HandlerParams) ([]*types.Transaction, error)
}

type Protocol struct {
	Name            string
	SupportedChains []int
}
