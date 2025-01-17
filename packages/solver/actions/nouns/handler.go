package nouns

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	IncreaseBid           = "increase_bid"
	HasTrait              = "has_trait"
	IsTokenId             = "is_token_id"
	CurrentBidWithinRange = "current_bid_within_range"

	sentences = map[types.Action]string{
		types.ActionBid:                     "Bid on noun with {0<amount:uint256>} ETH.",
		types.Action(IncreaseBid):           "Outbid the current bid by {0<percent:uint256>}%.",
		types.Action(HasTrait):              "Bid on noun that has a {0<traitType:string>} of {0=>1<trait:string>}.",
		types.Action(IsTokenId):             "Bid on noun when token id is {0<id:uint256>}.",
		types.Action(CurrentBidWithinRange): "Current bid for noun is greater than {0<min:uint256>} ETH and less than {1<max:uint256>} ETH.",
	}
)

type Nouns struct {
	*actions.BaseHandler
}

func New() actions.BaseProtocolHandler {
	return &Nouns{
		BaseHandler: actions.NewBaseHandler(
			"Nouns",
			"https://onplug.io/protocols/nouns.png",
			[]string{"nft"},
			utils.Mainnet.ChainIds,
			sentences,
			&NounsOptionsProvider{},
		),
	}
}

func (nouns *Nouns) GetTransaction(action types.Action, rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
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
