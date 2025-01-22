package aave_v3

import (
	"math/big"
	"solver/internal/actions"
	"solver/internal/references"
)

var (
	name = "Aave V3"
	icon = "https://cdn.onplug.io/protocols/aave.png"
	tags = []string{"lending", "defi"}

	chains  = append(references.Mainnet.ChainIds, references.Base.ChainIds...)
	schemas = map[string]actions.ActionDefinition{
		actions.ActionDeposit: {
			Sentence: "Deposit {0<amount:uint256>} {1<token:address>}",
			Handler:  HandleActionDeposit,
		},
		actions.ActionBorrow: {
			Sentence: "Borrow {0<tokenOut:address>} {1<amountOut:uint256>}",
			Handler:  HandleActionBorrow,
		},
		actions.ActionRepay: {
			Sentence: "Repay {0<tokenIn:address>} {1<amountIn:uint256>}",
			Handler:  HandleActionRepay,
		},
		actions.ActionWithdraw: {
			Sentence: "Withdraw {0<tokenOut:address>} {1<amountOut:uint256>}",
			Handler:  HandleActionWithdraw,
		},
		actions.ConstraintHealthFactor: {
			Type:     actions.TypeConstraint,
			Sentence: "Health factor is {0<operator:int8>} than {1<threshold:uint256>}",
			Handler:  HandleConstraintHealthFactor,
		},
		actions.ConstraintAPY: {
			Type:     actions.TypeConstraint,
			Sentence: "{0<direction:int8>} APY of {1<token:address>} is {2<operator:int8>} than {3<threshold:uint256>}%",
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
