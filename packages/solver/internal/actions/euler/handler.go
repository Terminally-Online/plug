package euler

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Euler"
	icon = "https://cdn.onplug.io/protocols/euler.png"
	tags = []string{"lending", "defi"}

	chains = []*references.Network{references.Mainnet, references.Base}

	ActionEarn               = "earn"
	ActionWithdraw           = "withdraw"
	ActionDepositCollateral  = "supply_collateral"
	ActionWithdrawCollateral = "withdraw_collateral"
	ActionBorrow             = "borrow"
	ActionRepay              = "repay"
	ConstraintAPY            = "apy"
	ConstraintHealthFactor   = "health_factor"
	ConstraintTimeToLiq      = "time_to_liquidation"

	schemas = map[string]actions.ActionDefinition{
		ActionEarn: {
			Sentence:       "Earn by depositing {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>}.",
			Handler:        HandleEarn,
			IsUserSpecific: true,
		},
		ActionWithdraw: {
			Sentence:       "Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>}.",
			Handler:        HandleWithdraw,
			IsUserSpecific: true,
		},
		ActionDepositCollateral: {
			Sentence:       "Deposit collateral {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} using {3<sub-account:uint8>}.",
			Handler:        HandleDepositCollateral,
			IsUserSpecific: true,
		},
		ActionWithdrawCollateral: {
			Sentence:       "Withdraw collateral {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>} using {3<sub-account:uint8>}.",
			Handler:        HandleWithdrawCollateral,
			IsUserSpecific: true,
		},
		ActionBorrow: {
			Sentence:       "Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<vault:address>} using {3<sub-account:uint8>}.",
			Handler:        HandleBorrow,
			IsUserSpecific: true,
		},
		ActionRepay: {
			Sentence:       "Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:address>} using {3<sub-account:uint8>}.",
			Handler:        HandleRepay,
			IsUserSpecific: true,
		},
		ConstraintHealthFactor: {
			Type:           actions.TypeConstraint,
			Sentence:       "Health factor for {0<sub-account:uint8>} is {1<operator:int8>} than {2<threshold:float>}.",
			Handler:        HandleConstraintHealthFactor,
			IsUserSpecific: true,
		},
		ConstraintTimeToLiq: {
			Type:           actions.TypeConstraint,
			Sentence:       "Time to liquidation in for {0<sub-account:uint8>} is {1<operator:int8>} than {2<threshold:float>} minutes.",
			Handler:        HandleConstraintTimeToLiquidation,
			IsUserSpecific: true,
		},
		ConstraintAPY: {
			Type:     actions.TypeConstraint,
			Sentence: "{0<direction:int8>} APY in {1<vault:string>} is {2<operator:int8>} than {3<threshold:float>}%.",
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
		&EulerOptionsProvider{},
	)
}
