package aave_v3

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	poolAddress      = utils.Mainnet.References["aave_v3"]["pool"]
	interestRateMode = new(big.Int).SetUint64(2)
	sentences        = map[types.Action]string{
		types.ActionDeposit:          "Deposit {0<tokenIn:address>} {1<amountIn:uint256>}.",
		types.ActionBorrow:           "Borrow {0<tokenOut:address>} {1<amountOut:uint256>}.",
		types.ActionRepay:            "Repay {0<tokenIn:address>} {1<amountIn:uint256>}.",
		types.ActionWithdraw:         "Withdraw {0<tokenOut:address>} {1<amountOut:uint256>}.",
		types.ConstraintHealthFactor: "Health factor is {0<operator:int8>} than {1<threshold:uint256>}.",
		types.ConstraintAPY:          "{0<direction:int8>} APY of {1<token:address>} is {2<operator:int8>} than {3<threshold:uint256>}%.",
	}
)

type AaveV3 struct {
	*actions.BaseHandler
}

func New() actions.BaseProtocolHandler {
	return &AaveV3{
		BaseHandler: actions.NewBaseHandler(
			"Aave V3",
			"https://onplug.io/protocols/aave.png",
			[]string{"lending", "defi"},
			utils.Mainnet.ChainIds,
			sentences,
			&AaveOptionsProvider{},
		),
	}
}

func (aaveV3 *AaveV3) GetTransaction(action types.Action, rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	switch action {
	case types.ActionDeposit:
		return HandleActionDeposit(rawInputs, params)
	case types.ActionBorrow:
		return HandleActionBorrow(rawInputs, params)
	case types.ActionRepay:
		return HandleActionRepay(rawInputs, params)
	case types.ActionWithdraw:
		return HandleActionWithdraw(rawInputs, params)
	case types.ConstraintHealthFactor:
		return HandleConstraintHealthFactor(rawInputs, params)
	case types.ConstraintAPY:
		return HandleConstraintAPY(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
