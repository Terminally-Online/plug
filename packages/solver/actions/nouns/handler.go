package nouns

import (
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	name = "Nouns"
	icon = "https://cdn.onplug.io/protocols/nouns.png"
	tags = []string{"nft"}

	chains = utils.Mainnet.ChainIds

	IncreaseBid           = "increase_bid"
	HasTrait              = "has_trait"
	IsTokenId             = "is_token_id"
	CurrentBidWithinRange = "current_bid_within_range"

	schemas = map[types.Action]actions.ActionDefinition{
		types.ActionBid: {
			Sentence: "Bid on noun with {0<amount:uint256>} ETH.",
			Handler:  HandleActionBid,
		},
		types.Action(IncreaseBid): {
			Sentence: "Outbid the current bid by {0<percent:uint256>}%.",
			Handler:  HandleActionIncreaseBid,
		},
		types.Action(HasTrait): {
			Type:     types.TypeConstraint,
			Sentence: "Noun that has a {0<traitType:string>} of {0=>1<trait:string>}.",
			Handler:  HandleConstraintHasTrait,
		},
		types.Action(IsTokenId): {
			Type:     types.TypeConstraint,
			Sentence: "Current Noun action is for token id {0<id:uint256>}.",
			Handler:  HandleConstraintIsTokenId,
		},
		types.Action(CurrentBidWithinRange): {
			Type:     types.TypeConstraint,
			Sentence: "Bid for Noun is greater than {0<min:uint256>} ETH and less than {1<max:uint256>} ETH.",
			Handler:  HandleConstraintCurrentBidWithinRange,
		},
	}
)

func New() actions.BaseProtocolHandler {
	return actions.NewBaseHandler(
		name,
		icon,
		tags,
		chains,
		schemas,
		&NounsOptionsProvider{},
	)
}
