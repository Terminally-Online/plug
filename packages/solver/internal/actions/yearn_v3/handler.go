package yearn_v3

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

func New() actions.Protocol {
	return actions.New(
		actions.Protocol{
			Name:   "Yearn V3",
			Icon:   "https://cdn.onplug.io/protocols/yearn.png",
			Tags:   []string{"yield", "defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinition{
				actions.ActionDeposit: {
					Sentence: "Deposit {0<amount:float>} {1<token:address:uint8>} into {1=>2<vault:address>}",
					Handler:  HandleActionDeposit,
					Options:  UnderlyingAssetToVaultOptions,
				},
				actions.ActionWithdraw: {
					Sentence: "Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}",
					Handler:  HandleActionWithdraw,
					Options:  UnderlyingAssetToVaultOptions,
				},
				actions.ActionStake: {
					Sentence: "Stake {0<amount:float>} {1<token:address:uint8>}",
					Handler:  HandleActionStake,
					Options:  AvailableStakingGaugeOptions,
				},
				actions.ActionRedeem: {
					Sentence: "Redeem {0<amount:float>} {1<token:address:uint8>}",
					Handler:  HandleActionRedeem,
					Options:  AvailableStakingGaugeOptions,
				},
				actions.ConstraintAPY: {
					Type:     actions.TypeConstraint,
					Sentence: "APY of {0<vault:address>} is {1<operator:int8>} than {2<threshold:float>} %",
					Handler:  HandleConstraintAPY,
					Options:  APYOptions,
				},
			},
		},
	)
}
