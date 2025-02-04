package euler

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Euler"
	icon = "https://cdn.onplug.io/protocols/euler.png"
	tags = []string{"lending", "defi"}

	chains = append(references.Mainnet.ChainIds, references.Base.ChainIds...)

	ActionEarn 		   		= "earn"
	ActionDepositCollateral	= "deposit"
	ActionWithdraw         	= "withdraw"
	ActionBorrow           	= "borrow"
	ActionRepay            	= "repay"
	ConstraintAPY          	= "apy"
	ConstraintHealthFactor 	= "health_factor"
	ConstraintTimeToLiq    	= "time_to_liquidation"

	schemas = map[string]actions.ActionDefinition{
		ActionEarn: {
			Sentence: "Earn by depositing {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} using sub-account {3<subaccount-index:uint8>}.",
			Handler:  HandleEarn,
		},
		ActionDepositCollateral: {
			Sentence: "Deposit {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} using sub-account {3<subaccount-index:uint8>}.",
			Handler:  HandleDepositCollateral,
		},
		ActionWithdraw: {
			Sentence: "Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}.",
			Handler:  HandleWithdraw,
		},
		ActionBorrow: {
			Sentence: "Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}.",
			Handler:  HandleBorrow,
		},
		ActionRepay: {
			Sentence: "Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>}.", // apy is the only sentence to not have a subaccount.
			Handler:  HandleRepay,
		},
		ConstraintHealthFactor: {
			Type:     actions.TypeConstraint,
			Sentence: "Health factor in {0<market:string>} is {1<operator:int8>} than {2<threshold:float>}.",
			Handler:  HandleConstraintHealthFactor,
		},
		ConstraintAPY: {
			Type:     actions.TypeConstraint,
			Sentence: "{0<action:int8>} borrow APY in {1<target:string>} is {2<operator:int8>} than {3<threshold:float>}%.",
			Handler:  HandleConstraintAPY,
		},
		ConstraintTimeToLiq: {
			Type:     actions.TypeConstraint,
			Sentence: "Time to liquidation in {0<vault:address>} is {1<operator:int8>} than {2<threshold:float>} hours.",
			Handler:  HandleConstraintTimeToLiquidation,
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
		&EulerOptionsProvider{},
	)
}
