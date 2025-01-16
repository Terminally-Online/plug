package actions

import (
	"encoding/json"
	"fmt"
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
	GetTags() []string
	GetActions() []types.Action
	GetChains() []int
	GetSchema(chainId string, action types.Action) (*types.Schema, error)
	GetTransaction(action types.Action, rawInputs json.RawMessage, params HandlerParams) ([]*types.Transaction, error)
}

type SchemaProvider interface {
	GetSchemas() map[types.Action]types.Schema
}

type Protocol struct {
	Name           string
	Icon           string
	Tags           []string
	Chains         []int
	SchemaProvider SchemaProvider
}

func (p Protocol) GetName() string {
	return p.Name
}

func (p Protocol) GetIcon() string {
	return p.Icon
}

func (p Protocol) GetTags() []string {
	return p.Tags
}

func (p Protocol) GetChains() []int {
	return p.Chains
}

func (p Protocol) GetActions() []types.Action {
	schemas := p.SchemaProvider.GetSchemas()
	actions := make([]types.Action, 0, len(schemas))
	for action := range schemas {
		actions = append(actions, action)
	}
	return actions
}

func (p Protocol) GetSchemas() map[types.Action]types.Schema {
	return p.SchemaProvider.GetSchemas()
}

func (p Protocol) GetSchema(chainId string, action types.Action) (*types.Schema, error) {
	schemas := p.SchemaProvider.GetSchemas()
	schema, exists := schemas[action]
	if !exists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
	return &schema, nil
}
