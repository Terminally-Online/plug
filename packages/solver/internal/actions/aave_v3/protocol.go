package aave_v3

import (
	"solver/internal/actions"
	aave_actions "solver/internal/actions/aave_v3/actions"
	aave_options "solver/internal/actions/aave_v3/options"
	"solver/internal/bindings/references"
)


func New() actions.Protocol {
	return actions.NewProtocol(actions.Protocol{
		Name:   "Aave V3",
		Icon:   "https://cdn.onplug.io/protocols/aave.png",
		Tags:   []string{"lending", "defi"},
		Chains: []*references.Network{references.Mainnet, references.Base},
		Actions: map[string]actions.ActionDefinitionInterface{
			actions.ActionDeposit: actions.NewActionDefinition(
				"Deposit {0<amount:float>} {1<token:address:uint8>}",
				aave_actions.Deposit,
				aave_options.CollateralOptions,
				false,
				false,
			),
			actions.ActionBorrow: actions.NewActionDefinition(
				"Borrow {0<amount:float>} {1<token:address:uint8>}",
				aave_actions.Borrow,
				aave_options.BorrowOptions,
				false,
				false,
			),
			actions.ActionRepay: actions.NewActionDefinition(
				"Repay {0<amount:float>} {1<token:address:uint8>}",
				aave_actions.Repay,
				aave_options.BorrowOptions,
				false,
				false,
			),
			actions.ActionWithdraw: actions.NewActionDefinition(
				"Withdraw {0<amount:float>} {1<token:address:uint8>}",
				aave_actions.Withdraw,
				aave_options.CollateralOptions,
				false,
				false,
			),
			actions.ConstraintHealthFactor: actions.NewActionDefinition(
				"Health factor is {0<operator:int8>} than {1<threshold:float>}",
				aave_actions.HealthFactor,
				aave_options.HealthFactorOptions,
				false,
				false,
			),
			actions.ConstraintAPY: actions.NewActionDefinition(
				"{0<action:int8>} APY of {1<token:address:uint8>} is {2<operator:int8>} than {3<threshold:float>}%",
				aave_actions.APY,
				aave_options.APYOptions,
				false,
				false,
			),
		},
	})
}
