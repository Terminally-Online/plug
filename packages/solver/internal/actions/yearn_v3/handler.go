package yearn_v3

import (
	"solver/internal/actions"
	"solver/internal/references"
)

var (
	name = "Yearn V3"
	icon = "https://cdn.onplug.io/protocols/yearn.png"
	tags = []string{"yield", "defi"}

	chains  = references.Mainnet.ChainIds
	schemas = map[string]actions.ActionDefinition{
		actions.ActionDeposit: {
			Sentence: "Deposit {0<amount:float>} {1<token:address:uint8>} into {1=>2<vault:address>}.",
			Handler:  HandleActionDeposit,
		},
		actions.ActionWithdraw: {
			Sentence: "Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}.",
			Handler:  HandleActionWithdraw,
		},
		// actions.ActionWithdrawMax: {
		// 	Sentence: "Withdraw max {0<token:address>} from {0=>1<vault:address>}",
		// 	Handler:  HandleActionWithdrawMax,
		// },
		actions.ActionStake: {
			Sentence: "Stake {0<amount:float>} {1<token:address:uint8>}",
			Handler:  HandleActionStake,
		},
		actions.ActionStakeMax: {
			Sentence: "Stake max {0<token:address:uint8>}",
			Handler:  HandleActionStakeMax,
		},
		actions.ActionRedeem: {
			Sentence: "Redeem {0<amount:float>} {1<token:address:uint8>}",
			Handler:  HandleActionRedeem,
		},
		actions.ActionRedeemMax: {
			Sentence: "Redeem max staking rewards for {0<token:address:uint8>}",
			Handler:  HandleActionRedeemMax,
		},
		actions.ConstraintAPY: {
			Type:     actions.TypeConstraint,
			Sentence: "APY of {0<vault:address>} is {1<operator:int8>} than {2<threshold:float>}%.",
			Handler:  HandleConstraintAPY,
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
		&YearnV3OptionsProvider{},
	)
}
