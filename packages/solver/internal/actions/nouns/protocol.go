package nouns

import (
	"solver/internal/actions"
	nouns_actions "solver/internal/actions/nouns/actions"
	nouns_options "solver/internal/actions/nouns/options"
	"solver/internal/bindings/references"
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
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				actions.ReadCurrentAuction: actions.NewActionDefinition(
					"Get current auction",
					nouns_actions.CurrentAuction,
					nil,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				actions.ReadHasTrait: actions.NewActionDefinition(
					"Noun that has a {0<traitType:string>} of {0=>1<trait:string>}",
					nouns_actions.HasTrait,
					nouns_options.HasTraitOptions,
					actions.IsGlobal,
					actions.IsDynamic,
					actions.IsEmptyOnchainFunc,
				),
				IncreaseBid: actions.NewActionDefinition(
					"Outbid the current bid by {0<percent:string>} %",
					nouns_actions.IncreaseBid,
					nil,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
			},
		},
	)
}
