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
	GetIcon() string
	SupportedActions() []types.Action
	SupportedChains() []int
	GetSchema(action types.Action) (types.Schema, error)
	GetTransaction(action types.Action, rawInputs json.RawMessage, params HandlerParams) ([]*types.Transaction, error)
}

type Protocol struct {
	Name            string
	Icon            string
	SupportedChains []int
}

func (p Protocol) GetIcon() string {
	return p.Icon
}
