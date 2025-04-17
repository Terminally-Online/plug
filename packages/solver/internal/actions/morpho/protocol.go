package morpho

import (
	"solver/internal/actions"
	morpho_actions "solver/internal/actions/morpho/actions"
	morpho_options "solver/internal/actions/morpho/options"
	"solver/internal/bindings/references"
)

var (
	ActionEarn             = "earn"
	ActionSupplyCollateral = "supply_collateral"
	ActionWithdraw         = "withdraw"
	ActionBorrow           = "borrow"
	ActionRepay            = "repay"
	ActionClaimRewards     = "claim_rewards"
	ConstraintLLTV         = "lltv"
	ConstraintAPY          = "apy"
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Morpho",
			Icon:   "https://cdn.onplug.io/protocols/morpho.png",
			Tags:   []string{"lending", "defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ReadAPY: actions.NewActionDefinition(
					"Get {0<action:int8>} APY in {1<target:string>}",
					morpho_actions.APY,
					morpho_options.APYOptions,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				ActionBorrow: actions.NewActionDefinition(
					"Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<market:string>}",
					morpho_actions.Borrow,
					morpho_options.BorrowTokenToMarketOptions,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				ActionClaimRewards: actions.NewActionDefinition(
					"Claim all reward distributions",
					morpho_actions.ClaimRewards,
					nil,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				ActionEarn: actions.NewActionDefinition(
					"Earn by depositing {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:string>}",
					morpho_actions.Earn,
					morpho_options.SupplyTokenToVaultOptions,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				actions.ReadHealthFactor: actions.NewActionDefinition(
					"Get health factor in {0<market:string>}",
					morpho_actions.HealthFactor,
					morpho_options.HealthFactorOptions,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				ActionRepay: actions.NewActionDefinition(
					"Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<market:string>}",
					morpho_actions.Repay,
					morpho_options.BorrowTokenToMarketOptions,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				ActionSupplyCollateral: actions.NewActionDefinition(
					"Supply {0<amount:float>} {1<token:address:uint8>} as collateral to {1=>2<market:string>}",
					morpho_actions.SupplyCollateral,
					morpho_options.CollateralTokenToMarketOptions,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
				ActionWithdraw: actions.NewActionDefinition(
					"Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<target:string>}",
					morpho_actions.Withdraw,
					morpho_options.SupplyAndCollateralTokenToMarketOptions,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
			},
		},
	)
}
