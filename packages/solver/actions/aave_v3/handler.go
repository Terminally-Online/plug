package aave_v3

import (
	"math/big"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	name = "Aave V3"
	icon = "https://cdn.onplug.io/protocols/aave.png"
	tags = []string{"lending", "defi"}

	chains  = append(utils.Mainnet.ChainIds, utils.Base.ChainIds...)
	schemas = map[types.Action]actions.ActionDefinition{
		types.ActionDeposit: {
			Sentence: "Deposit {0<amount:uint256>} {1<token:address>}",
			Handler:  HandleActionDeposit,
		},
		types.ActionBorrow: {
			Sentence: "Borrow {0<tokenOut:address>} {1<amountOut:uint256>}",
			Handler:  HandleActionBorrow,
		},
		types.ActionRepay: {
			Sentence: "Repay {0<tokenIn:address>} {1<amountIn:uint256>}",
			Handler:  HandleActionRepay,
		},
		types.ActionWithdraw: {
			Sentence: "Withdraw {0<tokenOut:address>} {1<amountOut:uint256>}",
			Handler:  HandleActionWithdraw,
		},
		types.ConstraintHealthFactor: {
			Type:     types.TypeConstraint,
			Sentence: "Health factor is {0<operator:int8>} than {1<threshold:uint256>}",
			Handler:  HandleConstraintHealthFactor,
		},
		types.ConstraintAPY: {
			Type:     types.TypeConstraint,
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
