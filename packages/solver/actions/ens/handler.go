package ens

import (
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
			Chains: []int{1}, // Ethereum mainnet only
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	// Buy ENS name
	h.schemas[types.Action("buy_ens")] = types.Schema{
		Sentence: "Buy ENS name {0<name:string>} for {1<price:uint256>} ETH.",
		Options:  map[int]types.SchemaOptions{},
	}

	// Renew ENS for specific years
	h.schemas[types.Action("renew_ens")] = types.Schema{
		Sentence: "Renew ENS name {0<name:string>} for {1<years:uint256>} years.",
		Options:  map[int]types.SchemaOptions{},
	}

	// Renew ENS for maximum duration
	h.schemas[types.Action("renew_ens_max")] = types.Schema{
		Sentence: "Renew ENS name {0<name:string>} for maximum duration.",
		Options:  map[int]types.SchemaOptions{},
	}

	return h
}

func (h *Handler) GetSchemas() map[types.Action]types.Schema {
	return h.schemas
}
