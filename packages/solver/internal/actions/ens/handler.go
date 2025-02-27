package ens

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "ENS"
	icon = "https://cdn.onplug.io/protocols/ens.png"
	tags = []string{"naming", "web3"}

	GracePeriod = "grace_period"
	TimeLeft    = "time_left"

	chains  = []*references.Network{references.Mainnet}
	schemas = map[string]actions.ActionDefinition{
		actions.ActionBuy: {
			Sentence: "Buy ENS {0<name:string>} with a max price of {1<maxPrice:float>} ETH",
			Handler:  HandleActionBuy,
		},
		actions.ActionRenew: {
			Sentence: "Renew ENS {0<name:string>} for {1<duration:uint256>} years",
			Handler:  HandleActionRenew,
		},
		actions.ConstraintPrice: {
			Type:     actions.TypeConstraint,
			Sentence: "Price to renew ENS {0<name:string>} for {1<duration:uint256>} is less than {2<price:float>} ETH",
			Handler:  HandleConstraintRenewalPrice,
		},
		GracePeriod: {
			Type:     actions.TypeConstraint,
			Sentence: "ENS {0<name:string>} is in renewal grace period",
			Handler:  HandleConstraintGracePeriod,
		},
		TimeLeft: {
			Type:     actions.TypeConstraint,
			Sentence: "Time left in ENS {0<name:string>} is less than {1<duration:uint256>}",
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
