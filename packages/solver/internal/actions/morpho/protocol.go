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

var (
	IS_GLOBAL = false
	IS_USER   = true

	IS_STATIC     = false
	IS_SEARCHABLE = true
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Morpho",
			Icon:   "https://cdn.onplug.io/protocols/morpho.png",
			Tags:   []string{"lending", "defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]any{
				ActionEarn: actions.NewActionDefinition(
					"Earn by depositing {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:string>}",
					morpho_actions.Earn,
					morpho_options.SupplyTokenToVaultOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
				ActionSupplyCollateral: actions.NewActionDefinition(
					"Supply {0<amount:float>} {1<token:address:uint8>} as collateral to {1=>2<market:string>}",
					morpho_actions.SupplyCollateral,
					morpho_options.CollateralTokenToMarketOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
				ActionWithdraw: actions.NewActionDefinition(
					"Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<target:string>}",
					morpho_actions.Withdraw,
					morpho_options.SupplyAndCollateralTokenToMarketOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
				ActionBorrow: actions.NewActionDefinition(
					"Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<market:string>}",
					morpho_actions.Borrow,
					morpho_options.BorrowTokenToMarketOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
				ActionRepay: actions.NewActionDefinition(
					"Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<market:string>}",
					morpho_actions.Repay,
					morpho_options.BorrowTokenToMarketOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
				ActionClaimRewards: actions.NewActionDefinition(
					"Claim all reward distributions",
					morpho_actions.ClaimRewards,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
				actions.ReadHealthFactor: actions.NewActionDefinition(
					"Get health factor in {0<market:string>}",
					morpho_actions.HealthFactor,
					morpho_options.HealthFactorOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
				actions.ReadAPY: actions.NewActionDefinition(
					"Get {0<action:int8>} APY in {1<target:string>}",
					morpho_actions.APY,
					morpho_options.APYOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
			},
		},
	)
}
