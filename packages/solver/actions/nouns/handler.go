package nouns

import (
	"solver/actions"
	"solver/cmd/references"
)

var (
	name = "Nouns"
	icon = "https://cdn.onplug.io/protocols/nouns.png"
	tags = []string{"nft"}

	chains = references.Mainnet.ChainIds

	IncreaseBid           = "increase_bid"
	HasTrait              = "has_trait"
	IsTokenId             = "is_token_id"
	CurrentBidWithinRange = "current_bid_within_range"

	schemas = map[string]actions.ActionDefinition{
		actions.ActionBid: {
			Sentence: "Bid on noun with {0<amount:uint256>} ETH.",
			Handler:  HandleActionBid,
		},
		IncreaseBid: {
			Sentence: "Outbid the current bid by {0<percent:uint256>}%.",
			Handler:  HandleActionIncreaseBid,
		},
		HasTrait: {
			Type:     actions.TypeConstraint,
			Sentence: "Noun that has a {0<traitType:string>} of {0=>1<trait:string>}.",
			Handler:  HandleConstraintHasTrait,
		},
		IsTokenId: {
			Type:     actions.TypeConstraint,
			Sentence: "Current Noun action is for token id {0<id:uint256>}.",
			Handler:  HandleConstraintIsTokenId,
		},
		CurrentBidWithinRange: {
			Type:     actions.TypeConstraint,
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
