package nouns

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
)

var (
	IncreaseBid           = "increase_bid"
	HasTrait              = "has_trait"
	IsTokenId             = "is_token_id"
	CurrentBidWithinRange = "current_bid_within_range"
)

type Handler struct {
	schemas map[types.Action]types.Schema
	actions.Protocol
}

func New() actions.BaseProtocolHandler {
	h := &Handler{
		schemas: make(map[types.Action]types.Schema),
		Protocol: actions.Protocol{
			Name:   "Nouns",
			Icon:   "https://nouns.wtf/favicon.ico",
			Tags:   []string{"nft"},
			Chains: []int{1},
		},
	}
	h.Protocol.SchemaProvider = h
	return h.init()
}

func (h *Handler) init() *Handler {
	traitTypeOptions, traitOptions, err := GetTraitOptions()
	if err != nil {
		return nil
	}

	h.schemas[types.ActionBid] = types.Schema{
		Sentence: "Bid on noun with {0<amount:uint256>} ETH.",
	}
	h.schemas[types.Action(IncreaseBid)] = types.Schema{
		Sentence: "Outbid the current bid by {0<percent:uint256>}%.",
	}
	h.schemas[types.Action(HasTrait)] = types.Schema{
		Sentence: "Bid on noun that has a {0<traitType:string>} of {0=>1<trait:uint256>}.",
		Options: map[int]types.SchemaOptions{
			0: {Simple: traitTypeOptions},
			1: {Complex: traitOptions},
		},
	}
	h.schemas[types.Action(IsTokenId)] = types.Schema{
		Sentence: "Bid on noun when token id is {0<id:uint256>}.",
	}
	h.schemas[types.Action(CurrentBidWithinRange)] = types.Schema{
		Sentence: "Current bid for noun is greater than {0<min:uint256>} ETH and less than {1<max:uint256>} ETH.",
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
	case types.ActionBid:
		return HandleActionBid(rawInputs, params)
	case types.Action(IncreaseBid):
		return HandleActionIncreaseBid(rawInputs, params)
	case types.Action(HasTrait):
		return HandleConstraintHasTrait(rawInputs, params)
	case types.Action(IsTokenId):
		return HandleConstraintIsTokenId(rawInputs, params)
	case types.Action(CurrentBidWithinRange):
		return HandleConstraintCurrentBidWithinRange(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
