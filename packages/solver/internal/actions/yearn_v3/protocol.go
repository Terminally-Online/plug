package yearn_v3

import (
	"solver/internal/actions"
	yearn_actions "solver/internal/actions/yearn_v3/actions"
	yearn_options "solver/internal/actions/yearn_v3/options"
	"solver/internal/bindings/references"
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Yearn V3",
			Icon:   "https://cdn.onplug.io/protocols/yearn.png",
			Tags:   []string{"yield", "defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ReadAPY: actions.NewActionDefinition(
					"APY of {0<vault:address>} is {1<operator:int8>} than {2<threshold:float>} %",
					yearn_actions.APY,
					yearn_options.APYOptions,
					actions.IsGlobal,
					actions.IsDynamic,
					actions.IsEmptyOnchainFunc,
				),
				actions.ActionDeposit: actions.NewActionDefinition(
					"Deposit {0<amount:float>} {1<token:address:uint8>} into {1=>2<vault:address>}",
					yearn_actions.Deposit,
					yearn_options.UnderlyingAssetToVaultOptions,
					actions.IsGlobal,
					actions.IsDynamic,
					&yearn_actions.DepositFunc,
				),
				actions.ActionUnstake: actions.NewActionDefinition(
					"Unstake {0<amount:float>} {1<token:address:uint8>} from {1=>2<gauge:address>}",
					yearn_actions.Unstake,
					yearn_options.AvailableStakingGaugeOptions,
					actions.IsGlobal,
					actions.IsDynamic,
					&yearn_actions.UnstakeFunc,
				),
				actions.ActionStake: actions.NewActionDefinition(
					"Stake {0<amount:float>} {1<token:address:uint8>} into {1=>2<gauge:address>}",
					yearn_actions.Stake,
					yearn_options.AvailableStakingGaugeOptions,
					actions.IsGlobal,
					actions.IsDynamic,
					&yearn_actions.StakeFunc,
				),
				actions.ActionWithdraw: actions.NewActionDefinition(
					"Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}",
					yearn_actions.Withdraw,
					yearn_options.UnderlyingAssetToVaultOptions,
					actions.IsGlobal,
					actions.IsDynamic,
					&yearn_actions.WithdrawFunc,
				),
			},
		},
	)
}
