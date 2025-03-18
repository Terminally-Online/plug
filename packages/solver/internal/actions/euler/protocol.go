package euler

import (
	"solver/internal/actions"
	euler_actions "solver/internal/actions/euler/actions"
	euler_options "solver/internal/actions/euler/options"
	"solver/internal/bindings/references"
)

var (
	ActionEarn               = "supply"
	ActionWithdraw           = "withdraw"
	ActionDepositCollateral  = "supply_collateral"
	ActionWithdrawCollateral = "withdraw_collateral"
	ActionBorrow             = "borrow"
	ActionRepay              = "repay"
	ConstraintAPY            = "apy"
	ConstraintHealthFactor   = "health_factor"
	ConstraintTimeToLiq      = "time_to_liquidation"
)

var (
	IS_GLOBAL = false
	IS_USER   = true

	IS_STATIC     = false
	IS_SEARCHABLE = true
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Euler",
			Icon:   "https://cdn.onplug.io/protocols/euler.png",
			Tags:   []string{"lending", "defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				ActionEarn: actions.NewActionDefinition(
					"Earn by depositing {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>}.",
					euler_actions.Earn,
					euler_options.SupplyTokenToVaultOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				ActionWithdraw: actions.NewActionDefinition(
					"Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}.",
					euler_actions.HandleWithdraw,
					euler_options.SupplyTokenToVaultOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				ActionDepositCollateral: actions.NewActionDefinition(
					"Deposit collateral {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} using {3<sub-account:uint8>}.",
					euler_actions.DepositCollateral,
					euler_options.SupplyTokenToVaultToPositionsOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				ActionWithdrawCollateral: actions.NewActionDefinition(
					"Withdraw collateral {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>} using {3<sub-account:uint8>}.",
					euler_actions.HandleWithdrawCollateral,
					euler_options.SupplyTokenToVaultToPositionsOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				ActionBorrow: actions.NewActionDefinition(
					"Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>} using {3<sub-account:uint8>}.",
					euler_actions.Borrow,
					euler_options.BorrowTokenToVaultOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				ActionRepay: actions.NewActionDefinition(
					"Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} using {3<sub-account:uint8>}.",
					euler_actions.Repay,
					euler_options.BorrowTokenToVaultOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				ConstraintHealthFactor: actions.NewActionDefinition(
					"Health factor for {0<sub-account:uint8>} is {1<operator:int8>} than {2<threshold:float>}.",
					euler_actions.HealthFactor,
					euler_options.PositionOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				ConstraintTimeToLiq: actions.NewActionDefinition(
					"Time to liquidation in for {0<sub-account:uint8>} is {1<operator:int8>} than {2<threshold:float>} minutes.",
					euler_actions.TimeToLiquidation,
					euler_options.PositionOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				ConstraintAPY: actions.NewActionDefinition(
					"{0<direction:int8>} APY in {1<vault:string>} is {2<operator:int8>} than {3<threshold:float>}%.",
					euler_actions.APY,
					euler_options.TokenOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
			},
		},
	)
}
