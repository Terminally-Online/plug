package morpho

import (
	"solver/internal/actions"
	"solver/internal/references"
)

var (
	name = "Morpho"
	icon = "https://cdn.onplug.io/protocols/morpho.png"
	tags = []string{"lending", "defi"}

	chains = references.Mainnet.ChainIds

	ActionEarn             = "earn"
	ActionSupplyCollateral = "supply_collateral"
	ActionWithdraw         = "withdraw"
	ActionWithdrawAll      = "withdraw_all"
	ActionBorrow           = "borrow"
	ActionRepay            = "repay"
	ActionRepayAll         = "repay_all"
	ActionClaimRewards     = "claim_rewards"
	ConstraintLLTV         = "lltv"
	ConstraintAPY          = "apy"

	schemas = map[string]actions.ActionDefinition{
		ActionEarn: {
			Sentence: "Earn by depositing {0<amount:uint256>} {1<token:address>} to {1=>2<vault:string>}.",
			Handler:  HandleEarn,
		},
		ActionSupplyCollateral: {
			Sentence: "Supply {0<amount:uint256>} {1<token:address>} as collateral to {1=>2<market:string>}.",
			Handler:  HandleSupplyCollateral,
		},
		ActionWithdraw: {
			Sentence: "Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<target:string>}.",
			Handler:  HandleWithdraw,
		},
		ActionWithdrawAll: {
			Sentence: "Withdraw all {0<token:address>} from {0=>1<target:string>}.",
			Handler:  HandleWithdrawAll,
		},
		ActionBorrow: {
			Sentence: "Borrow {0<amount:uint256>} {1<token:address>} from {1=>2<market:string>}.",
			Handler:  HandleBorrow,
		},
		ActionRepay: {
			Sentence: "Repay {0<amount:uint256>} {1<token:address>} to {1=>2<market:string>}.",
			Handler:  HandleRepay,
		},
		ActionRepayAll: {
			Sentence: "Repay all {0<token:address>} to {0=>1<market:string>}.",
			Handler:  HandleRepayAll,
		},
		ActionClaimRewards: {
			Sentence: "Claim all reward distributions.",
			Handler:  HandleClaimRewards,
		},
		actions.ConstraintHealthFactor: {
			Type:     actions.TypeConstraint,
			Sentence: "Health factor in {0<market:string>} is {1<operator:int8>} than {2<threshold:uint256>}.",
			Handler:  HandleConstraintHealthFactor,
		},
		actions.ConstraintAPY: {
			Type:     actions.TypeConstraint,
			Sentence: "{0<action:int8>} APY in {1<target:string>} is {2<operator:int8>} than {3<threshold:uint256>}%.",
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
		&MorphoOptionsProvider{},
	)
}
