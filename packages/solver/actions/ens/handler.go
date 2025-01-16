package ens

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	SetPrimary   = "set_primary"
	GracePeriod  = "grace_period"
	TimeLeft     = "time_left"
	RenewalPrice = "renewal_price"
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
			Icon:   "/protocols/ens.png",
			Tags:   []string{"naming", "web3"},
			Chains: utils.Mainnet.ChainIds,
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	durationOptions, err := GetDurationOptions()
	if err != nil {
		return nil
	}

	h.schemas[types.ActionBuy] = types.Schema{
		Sentence: "Buy ENS {0<name:string>} with a max price of {1<maxPrice:uint256>} ETH.",
	}

	h.schemas[types.ActionRenew] = types.Schema{
		Sentence: "Renew ENS {0<name:string>} for {1<duration:uint256>} years.",
		Options: map[int]types.SchemaOptions{
			1: {Simple: durationOptions},
		},
	}

	h.schemas[types.Action(RenewalPrice)] = types.Schema{
		Sentence: "Price to renew ENS {0<name:string>} for {1<duration:uint256>} is less than {2<price:uint256>} ETH.",
		Options: map[int]types.SchemaOptions{
			1: {Simple: durationOptions},
		},
	}

	h.schemas[types.Action(GracePeriod)] = types.Schema{
		Sentence: "ENS {0<name:string>} is in renewal grace period.",
	}

	h.schemas[types.Action(TimeLeft)] = types.Schema{
		Sentence: "Time left in ENS {0<name:string>} is less than {1<duration:uint256>}.",
		Options: map[int]types.SchemaOptions{
			1: {Simple: durationOptions},
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
	switch action {
	case types.ActionBuy:
		return HandleActionBuy(rawInputs, params)
	case types.ActionRenew:
		return HandleActionRenew(rawInputs, params)
	case types.Action(RenewalPrice):
		return HandleConstraintRenewalPrice(rawInputs, params)
	case types.Action(GracePeriod):
		return HandleConstraintGracePeriod(rawInputs, params)
	case types.Action(TimeLeft):
		return HandleConstraintTimeLeft(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
