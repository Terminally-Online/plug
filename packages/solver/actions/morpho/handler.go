package morpho

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
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

	sentences = map[types.Action]string{
		types.Action(ActionEarn):             "Earn by depositing {0<amount:uint256>} {1<token:address>} to {1=>2<vault:string>}.",
		types.Action(ActionSupplyCollateral): "Supply {0<amount:uint256>} {1<token:address>} as collateral to {1=>2<market:string>}.",
		types.Action(ActionWithdraw):         "Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<target:string>}.",
		types.Action(ActionWithdrawAll):      "Withdraw all {0<token:address>} from {0=>1<target:string>}.",
		types.Action(ActionBorrow):           "Borrow {0<amount:uint256>} {1<token:address>} from {1=>2<market:string>}.",
		types.Action(ActionRepay):            "Repay {0<amount:uint256>} {1<token:address>} to {1=>2<market:string>}.",
		types.Action(ActionRepayAll):         "Repay all {0<token:address>} to {0=>1<market:string>}.",
		types.Action(ActionClaimRewards):     "Claim all reward distributions.",
		types.ConstraintHealthFactor:         "Health factor in {0<market:string>} is {1<operator:int8>} than {2<threshold:uint256>}.",
		types.Action(ConstraintAPY):          "{0<action:int8>} APY in {1<target:string>} is {2<operator:int8>} than {3<threshold:uint256>}%.",
	}
)

type Morpho struct {
	*actions.BaseHandler
}

func New() actions.BaseProtocolHandler {
	return &Morpho{
		BaseHandler: actions.NewBaseHandler(
			"Morpho",
			"https://onplug.io/protocols/morpho.png",
			[]string{"lending", "defi"},
			utils.Mainnet.ChainIds,
			sentences,
			&MorphoOptionsProvider{},
		),
	}
}

func (morpho *Morpho) GetTransaction(action types.Action, rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	switch action {
	case types.Action(ActionEarn):
		return HandleEarn(rawInputs, params)
	case types.Action(ActionSupplyCollateral):
		return HandleSupplyCollateral(rawInputs, params)
	case types.Action(ActionWithdraw):
		return HandleWithdraw(rawInputs, params)
	case types.Action(ActionWithdrawAll):
		return HandleWithdrawAll(rawInputs, params)
	case types.Action(ActionBorrow):
		return HandleBorrow(rawInputs, params)
	case types.Action(ActionRepay):
		return HandleRepay(rawInputs, params)
	case types.Action(ActionRepayAll):
		return HandleRepayAll(rawInputs, params)
	case types.Action(ActionClaimRewards):
		return HandleClaimRewards(rawInputs, params)
	case types.Action(ConstraintLLTV):
		return HandleConstraintAPY(rawInputs, params)
	case types.ConstraintAPY:
		return HandleConstraintAPY(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
