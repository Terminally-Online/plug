package euler

import (
	"solver/internal/actions"
	euler_actions "solver/internal/actions/euler/actions"
	euler_options "solver/internal/actions/euler/options"
	"solver/internal/bindings/references"
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Euler",
			Icon:   "https://cdn.onplug.io/protocols/euler.png",
			Tags:   []string{"lending", "defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ReadAPY: actions.NewActionDefinition(
					"Get {0<direction:int8>} APY in {1<vault:string>}",
					euler_actions.APY,
					euler_options.TokenOptions,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				actions.ActionBorrow: actions.NewActionDefinition(
					"Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>} using {3<sub-account:uint8>}",
					euler_actions.Borrow,
					euler_options.BorrowTokenToVaultOptions,
					actions.IsUser,
					actions.IsDynamic,
					&euler_actions.BorrowFunc,
				),
				actions.ActionDepositCollateral: actions.NewActionDefinition(
					"Deposit collateral {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} using {3<sub-account:uint8>}",
					euler_actions.DepositCollateral,
					euler_options.SupplyTokenToVaultToPositionsOptions,
					actions.IsUser,
					actions.IsDynamic,
					&euler_actions.DepositCollateralFunc,
				),
				actions.ActionEarn: actions.NewActionDefinition(
					"Earn by depositing {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>}",
					euler_actions.Earn,
					euler_options.SupplyTokenToVaultOptions,
					actions.IsUser,
					actions.IsDynamic,
					&euler_actions.DepositFunc,
				),
				actions.ReadHealthFactor: actions.NewActionDefinition(
					"Get health factor for {0<sub-account:uint8>}",
					euler_actions.HealthFactor,
					euler_options.PositionOptions,
					actions.IsUser,
					actions.IsDynamic,
					actions.IsEmptyOnchainFunc,
				),
				actions.ActionRepay: actions.NewActionDefinition(
					"Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} using {3<sub-account:uint8>}",
					euler_actions.Repay,
					euler_options.BorrowTokenToVaultOptions,
					actions.IsUser,
					actions.IsDynamic,
					&euler_actions.RepayFunc,
				),
				actions.ReadTimeToLiquidiation: actions.NewActionDefinition(
					"Get time to liquidation for {0<sub-account:uint8>} in minutes",
					euler_actions.TimeToLiquidation,
					euler_options.PositionOptions,
					actions.IsUser,
					actions.IsDynamic,
					actions.IsEmptyOnchainFunc,
				),
				actions.ActionWithdraw: actions.NewActionDefinition(
					"Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}",
					euler_actions.HandleWithdraw,
					euler_options.SupplyTokenToVaultOptions,
					actions.IsUser,
					actions.IsDynamic,
					&euler_actions.WithdrawFunc,
				),
				actions.ActionWithdrawCollateral: actions.NewActionDefinition(
					"Withdraw collateral {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>} using {3<sub-account:uint8>}",
					euler_actions.HandleWithdraw,
					euler_options.SupplyTokenToVaultToPositionsOptions,
					actions.IsUser,
					actions.IsDynamic,
					&euler_actions.WithdrawFunc,
				),
			},
		},
	)
}
