package nouns

import (
	"solver/internal/actions"
	nouns_actions "solver/internal/actions/nouns/actions"
	nouns_options "solver/internal/actions/nouns/options"
	"solver/internal/bindings/references"
)

var (
	ActionIncreaseBidKey = "increase_bid"

	ActionBidSentence          = "Bid on noun with {0<amount:string>} ETH"
	ActionIncreaseBidSentence  = "Outbid the current bid by {0<percent:string>} %"
	ReadHasTraitSentence       = "Noun that has a {0<traitType:string>} of {0=>1<trait:string>}"
	ReadCurrentAuctionSentence = "Get current auction"

	ActionBid = actions.NewActionDefinition(
		ActionBidSentence,
		nouns_actions.Bid,
		nil,
		nil,
		actions.IsEmptyOnchainFunc,
	)
	ActionIncreaseBid = actions.NewActionDefinition(
		ActionIncreaseBidSentence,
		nouns_actions.IncreaseBid,
		nil,
		nil,
		actions.IsEmptyOnchainFunc,
	)
	ReadHasTrait = actions.NewActionDefinition(
		ReadHasTraitSentence,
		nouns_actions.HasTrait,
		nouns_options.HasTraitOptions,
		&actions.ActionProperties{
			IsSearchable: true,
		},
		actions.IsEmptyOnchainFunc,
	)
	ReadCurrentAuction = actions.NewActionDefinition(
		ReadCurrentAuctionSentence,
		nouns_actions.CurrentAuction,
		nil,
		nil,
		actions.IsEmptyOnchainFunc,
	)
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Nouns",
			Icon:   "https://cdn.onplug.io/protocols/nouns.png",
			Tags:   []string{"nft"},
			Chains: []*references.Network{references.Mainnet},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ActionBid:          ActionBid,
				ActionIncreaseBidKey:       ActionIncreaseBid,
				actions.ReadHasTrait:       ReadHasTrait,
				actions.ReadCurrentAuction: ReadCurrentAuction,
			},
		},
	)
}
