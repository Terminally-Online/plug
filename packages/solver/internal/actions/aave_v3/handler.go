package aave_v3

import (
	"math/big"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Aave V3"
	icon = "https://cdn.onplug.io/protocols/aave.png"
	tags = []string{"lending", "defi"}

	chains  = []*references.Network{references.Mainnet, references.Base}
	schemas = map[string]actions.ActionDefinition{
		actions.ActionDeposit: {
			Sentence: "Deposit {0<amount:float>} {1<token:address:uint8>}",
			Handler:  HandleActionDeposit,
		},
		actions.ActionBorrow: {
			Sentence: "Borrow {0<amount:float>} {1<token:address:uint8>}",
			Handler:  HandleActionBorrow,
		},
		actions.ActionRepay: {
			Sentence: "Repay {0<amount:float>} {1<token:address:uint8>}",
			Handler:  HandleActionRepay,
		},
		actions.ActionWithdraw: {
			Sentence: "Withdraw {0<amount:float>} {1<token:address:uint8>}",
			Handler:  HandleActionWithdraw,
		},
		actions.ConstraintHealthFactor: {
			Type:     actions.TypeConstraint,
			Sentence: "Health factor is {0<operator:int8>} than {1<threshold:float>}",
			Handler:  HandleConstraintHealthFactor,
		},
		actions.ConstraintAPY: {
			Type:     actions.TypeConstraint,
			Sentence: "{0<direction:int8>} APY of {1<token:address:uint8>} is {2<operator:int8>} than {3<threshold:float>}%",
			Handler:  HandleConstraintAPY,
		},
	}

	interestRateMode = new(big.Int).SetUint64(2)
)

func New() actions.BaseProtocolHandler {
	return actions.NewBaseHandler(
		name,
		icon,
		tags,
		chains,
		schemas,
		&AaveOptionsProvider{},
	)
}
