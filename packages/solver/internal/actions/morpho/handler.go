package morpho

import (
	"solver/internal/actions"
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
	return actions.New(
		actions.Protocol{
			Name:   "Morpho",
			Icon:   "https://cdn.onplug.io/protocols/morpho.png",
			Tags:   []string{"lending", "defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinition{
				ActionEarn: {
					Sentence: "Earn by depositing {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:string>}",
					Handler:  HandleEarn,
					Options:  SupplyTokenToVaultOptions,
				},
				ActionSupplyCollateral: {
					Sentence: "Supply {0<amount:float>} {1<token:address:uint8>} as collateral to {1=>2<market:string>}",
					Handler:  HandleSupplyCollateral,
					Options:  CollateralTokenToMarketOptions,
				},
				ActionWithdraw: {
					Sentence: "Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<target:string>}",
					Handler:  HandleWithdraw,
					Options:  SupplyAndCollateralTokenToMarketOptions,
				},
				ActionBorrow: {
					Sentence: "Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<market:string>}",
					Handler:  HandleBorrow,
					Options:  BorrowTokenToMarketOptions,
				},
				ActionRepay: {
					Sentence: "Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<market:string>}",
					Handler:  HandleRepay,
					Options:  BorrowTokenToMarketOptions,
				},
				ActionClaimRewards: {
					Sentence: "Claim all reward distributions",
					Handler:  HandleClaimRewards,
				},
				actions.ConstraintHealthFactor: {
					Type:     actions.TypeConstraint,
					Sentence: "Health factor in {0<market:string>} is {1<operator:int8>} than {2<threshold:float>}",
					Handler:  HandleConstraintHealthFactor,
					Options:  HealthFactorOptions,
				},
				actions.ConstraintAPY: {
					Type:     actions.TypeConstraint,
					Sentence: "{0<action:int8>} APY in {1<target:string>} is {2<operator:int8>} than {3<threshold:float>} %",
					Handler:  HandleConstraintAPY,
					Options:  APYOptions,
				},
			},
		},
	)
}
