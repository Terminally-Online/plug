package nouns

import (
	"solver/internal/actions"
	nouns_actions "solver/internal/actions/nouns/actions"
	nouns_options "solver/internal/actions/nouns/options"
	"solver/internal/bindings/references"
)

var (
	IS_GLOBAL = false
	IS_USER   = true

	IS_STATIC  = false
	IS_DYNAMIC = true
)

var (
	IncreaseBid    = "increase_bid"
	HasTrait       = "has_trait"
	IsTokenId      = "is_token_id"
	CurrentAuction = "current_auction"
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Nouns",
			Icon:   "https://cdn.onplug.io/protocols/nouns.png",
			Tags:   []string{"nft"},
			Chains: []*references.Network{references.Mainnet},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ActionBid: actions.NewActionDefinition(
					"Bid on noun with {0<amount:string>} ETH",
					nouns_actions.Bid,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
				IncreaseBid: actions.NewActionDefinition(
					"Outbid the current bid by {0<percent:string>} %",
					nouns_actions.IncreaseBid,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
				HasTrait: actions.NewActionDefinition(
					"Noun that has a {0<traitType:string>} of {0=>1<trait:string>}",
					nouns_actions.HasTrait,
					nouns_options.HasTraitOptions,
					IS_GLOBAL,
					IS_DYNAMIC,
				),
				IsTokenId: actions.NewActionDefinition(
					"Current Noun action is for token id {0<id:uint256>}",
					nouns_actions.IsTokenId,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
				CurrentAuction: actions.NewActionDefinition(
					"Get current auction",
					nouns_actions.CurrentAuction,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
			},
		},
	)
}
