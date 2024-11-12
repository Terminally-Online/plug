package yearn_v3

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
)

// var (
// 	poolAddress      = utils.Mainnet.References["yearn_v3"]["pool"]
// )

type Handler struct {
	schemas map[types.Action]types.Schema
	actions.Protocol
}

func New() actions.BaseProtocolHandler {
	h := &Handler{
		schemas: make(map[types.Action]types.Schema),
		Protocol: actions.Protocol{
			Name:   "Yearn V3",
			Icon:   "https://yearn.fi/favicons/favicon.ico",
			Tags:   []string{"yield", "defi"},
			Chains: []int{1},
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	underlyingAssetOptions, err := GetUnderlyingAssetOptions()
	if err != nil {
		return nil
	}

	h.schemas[types.ActionDeposit] = types.Schema{
		Sentence: "Deposit {0} {1}.",
		Fields: []types.SchemaField{
			{
				Name:    "tokenIn",
				Type:    "address",
				Options: underlyingAssetOptions,
			},
			{
				Name: "amountIn",
				Type: "uint256",
			},
		},
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
	// var calldata []byte
	// var err error
	switch action {
	// case types.ActionDeposit:
	// 	calldata, err = HandleActionDeposit(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
	// if err != nil {
	// 	return nil, utils.ErrTransactionFailed(err.Error())
	// }

	// return []*types.Transaction{{
	// 	To:   poolAddress,
	// 	Data: "0x" + common.Bytes2Hex(calldata),
	// }}, nil
}
