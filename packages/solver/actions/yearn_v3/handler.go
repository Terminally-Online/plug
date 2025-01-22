package yearn_v3

import (
	"solver/actions"
	"solver/utils"
)

var (
	name = "Yearn V3"
	icon = "https://cdn.onplug.io/protocols/yearn.png"
	tags = []string{"yield", "defi"}

	chains  = utils.Mainnet.ChainIds
	schemas = map[string]actions.ActionDefinition{
		actions.ActionDeposit: {
			Sentence: "Deposit {0<amount:uint256>} {1<token:address>} into {1=>2<vault:address>}.",
			Handler:  HandleActionDeposit,
		},
		actions.ActionWithdraw: {
			Sentence: "Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<vault:address>}.",
			Handler:  HandleActionWithdraw,
		},
		// actions.ActionWithdrawMax: {
		// 	Sentence: "Withdraw max {0<token:address>} from {0=>1<vault:address>}",
		// 	Handler:  HandleActionWithdrawMax,
		// },
		actions.ActionStake: {
			Sentence: "Stake {0<amount:uint256>} {1<token:address>}",
			Handler:  HandleActionStake,
		},
		actions.ActionStakeMax: {
			Sentence: "Stake max {0<token:address>}",
			Handler:  HandleActionStakeMax,
		},
		actions.ActionRedeem: {
			Sentence: "Redeem {0<amount:uint256>} {1<token:address>}",
			Handler:  HandleActionRedeem,
		},
		actions.ActionRedeemMax: {
			Sentence: "Redeem max staking rewards for {0<token:address>}",
			Handler:  HandleActionRedeemMax,
		},
		actions.ConstraintAPY: {
			Type:     actions.TypeConstraint,
			Sentence: "APY of {0<vault:address>} is {1<operator:int8>} than {2<threshold:uint256>}%.",
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
