package ens

import (
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	name = "ENS"
	icon = "https://cdn.onplug.io/protocols/ens.png"
	tags = []string{"naming", "web3"}

	chains = utils.Mainnet.ChainIds

	SetPrimary   = "set_primary"
	GracePeriod  = "grace_period"
	TimeLeft     = "time_left"
	RenewalPrice = "renewal_price"

	schemas = map[types.Action]actions.ActionDefinition{
		types.ActionBuy: {
			Sentence: "Buy ENS {0<name:string>} with a max price of {1<maxPrice:float>} ETH.",
			Handler:  HandleActionBuy,
		},
		types.ActionRenew: {
			Sentence: "Renew ENS {0<name:string>} for {1<duration:uint256>} years.",
			Handler:  HandleActionRenew,
		},
		types.Action(RenewalPrice): {
			Type:     types.TypeConstraint,
			Sentence: "Price to renew ENS {0<name:string>} for {1<duration:uint256>} is less than {2<price:uint256>} ETH.",
			Handler:  HandleConstraintRenewalPrice,
		},
		types.Action(GracePeriod): {
			Type:     types.TypeConstraint,
			Sentence: "ENS {0<name:string>} is in renewal grace period.",
			Handler:  HandleConstraintGracePeriod,
		},
		types.Action(TimeLeft): {
			Type:     types.TypeConstraint,
			Sentence: "Time left in ENS {0<name:string>} is less than {1<duration:uint256>}.",
			Handler:  HandleConstraintTimeLeft,
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
		&EnsOptionsProvider{},
	)
}
