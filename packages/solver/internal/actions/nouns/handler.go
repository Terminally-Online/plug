package nouns

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	IncreaseBid           = "increase_bid"
	HasTrait              = "has_trait"
	IsTokenId             = "is_token_id"
	CurrentBidWithinRange = "current_bid_within_range"

	name = "Nouns"
	icon = "https://cdn.onplug.io/protocols/nouns.png"
	tags = []string{"nft"}

	chains  = []*references.Network{references.Mainnet}
	schemas = map[string]actions.ActionDefinition{
		actions.ActionBid: {
			Sentence: "Bid on noun with {0<amount:string>} ETH",
			Handler:  HandleActionBid,
		},
		IncreaseBid: {
			Sentence: "Outbid the current bid by {0<percent:string>} %",
			Handler:  HandleActionIncreaseBid,
		},
		HasTrait: {
			Type:     actions.TypeConstraint,
			Sentence: "Noun that has a {0<traitType:string>} of {0=>1<trait:string>}",
			Handler:  HandleConstraintHasTrait,
		},
		IsTokenId: {
			Type:     actions.TypeConstraint,
			Sentence: "Current Noun action is for token id {0<id:uint256>}",
			Handler:  HandleConstraintIsTokenId,
		},
		CurrentBidWithinRange: {
			Type:     actions.TypeConstraint,
			Sentence: "Bid for Noun is greater than {0<min:string>} ETH and less than {1<max:string>} ETH",
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
