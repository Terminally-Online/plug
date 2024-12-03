package ens

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
)

type Handler struct {
	schemas map[types.Action]types.Schema
	actions.Protocol
}

func New() actions.BaseProtocolHandler {
	h := &Handler{
		schemas: make(map[types.Action]types.Schema),
		Protocol: actions.Protocol{
			Name:   "ENS",
			Icon:   "https://app.ens.domains/favicon.ico",
			Tags:   []string{"naming", "web3"},
			Chains: []int{1},
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	h.schemas[types.ActionBuy] = types.Schema{
		Sentence: "Buy ENS name {0<name:string>} for {1<price:uint256>} ETH.",
		Options:  map[int]types.SchemaOptions{},
	}

	h.schemas[types.ActionRenew] = types.Schema{
		Sentence: "Renew ENS {0<name:string>} for {1<duration:uint256>} years.",
		Options:  map[int]types.SchemaOptions{},
	}

	h.schemas[types.ConstraintGracePeriod] = types.Schema{
		Sentence: "ENS {0<name:string>} is in renewal grace period.",
		Options:  map[int]types.SchemaOptions{},
	}

	h.schemas[types.ConstraintTimeLeft] = types.Schema{
		Sentence: "Time left in ENS {0<name:string>} is less than {1<duration:uint256>}.",
		Options:  map[int]types.SchemaOptions{},
	}

	return h
}

func (h *Handler) GetSchemas() map[types.Action]types.Schema {
	return h.schemas
}

func (h *Handler) GetSchema(action types.Action) (*types.Schema, error) {
	schema, exists := h.schemas[action]
	if !exists {
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
	return &schema, nil
}

func (h *Handler) GetTransaction(action types.Action, rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	switch action {
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
