package plug

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

type Handler struct {
	schemas map[types.Action]types.Schema
	actions.Protocol
}

func New() actions.BaseProtocolHandler {
	h := &Handler{
		schemas: make(map[types.Action]types.Schema),
		Protocol: actions.Protocol{
			Name:   "Plug",
			Icon:   "https://onplug.io/favicon.ico",
			Tags:   []string{"defi"},
			Chains: utils.Mainnet.ChainIds,
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	h.schemas[types.ActionTransfer] = types.Schema{
		Sentence: "Transfer {0<amount:[(1.1)=721?1:uint256]>} {1<token:address:uint256>} {2<id:[(1.1)>20?uint256:null]>} to {3<recipient:address>}",
	}

	h.schemas[types.ActionTransferFrom] = types.Schema{
		Sentence: "Transfer {0<amount:uint256>} {1<token:address>} to {2<recipient:address>}.",
	}

	h.schemas[types.ActionSwap] = types.Schema{
		Sentence: "Swap {0<amount:uint256>} {1<tokenIn:address>} for {2<tokenOut:address>}.",
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
	case types.ActionTransfer:
		return HandleTransfer(rawInputs, params)
	case types.ActionTransferFrom:
		return HandleTransferFrom(rawInputs, params)
	case types.ActionSwap:
		return HandleSwap(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
