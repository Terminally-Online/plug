package morpho

import (
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	name = "Morpho"
	icon = "https://onplug.io/protocols/morpho.png"
	tags = []string{"lending", "defi"}

	chains = utils.Mainnet.ChainIds

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

	schemas = map[types.Action]actions.ActionDefinition{
		types.Action(ActionEarn): {
			Sentence: "Earn by depositing {0<amount:uint256>} {1<token:address>} to {1=>2<vault:string>}.",
			Handler:  HandleEarn,
		},
		types.Action(ActionSupplyCollateral): {
			Sentence: "Supply {0<amount:uint256>} {1<token:address>} as collateral to {1=>2<market:string>}.",
			Handler:  HandleSupplyCollateral,
		},
		types.Action(ActionWithdraw): {
			Sentence: "Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<target:string>}.",
			Handler:  HandleWithdraw,
		},
		types.Action(ActionWithdrawAll): {
			Sentence: "Withdraw all {0<token:address>} from {0=>1<target:string>}.",
			Handler:  HandleWithdrawAll,
		},
		types.Action(ActionBorrow): {
			Sentence: "Borrow {0<amount:uint256>} {1<token:address>} from {1=>2<market:string>}.",
			Handler:  HandleBorrow,
		},
		types.Action(ActionRepay): {
			Sentence: "Repay {0<amount:uint256>} {1<token:address>} to {1=>2<market:string>}.",
			Handler:  HandleRepay,
		},
		types.Action(ActionRepayAll): {
			Sentence: "Repay all {0<token:address>} to {0=>1<market:string>}.",
			Handler:  HandleRepayAll,
		},
		types.Action(ActionClaimRewards): {
			Sentence: "Claim all reward distributions.",
			Handler:  HandleClaimRewards,
		},
		types.ConstraintHealthFactor: {
			Sentence: "Health factor in {0<market:string>} is {1<operator:int8>} than {2<threshold:uint256>}.",
			Handler:  HandleConstraintHealthFactor,
		},
		types.ConstraintAPY: {
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
