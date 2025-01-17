package yearn_v3

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	name = "Yearn V3"
	icon = "https://onplug.io/protocols/yearn.png"
	tags = []string{"yield", "defi"}

	chains    = utils.Mainnet.ChainIds
	sentences = map[types.Action]string{
		types.ActionDeposit:     "Deposit {0<amount:uint256>} {1<token:address>} into {1=>2<vault:address>}.",
		types.ActionWithdraw:    "Withdraw {0<amount:uint256>} {1<token:address>} from {1=>2<vault:address>}.",
		types.ActionWithdrawMax: "Withdraw max {0<token:address>} from {0=>1<vault:address>}",
		types.ActionStake:       "Stake {0<amount:uint256>} {1<token:address>}",
		types.ActionStakeMax:    "Stake max {0<token:address>}",
		types.ActionRedeem:      "Redeem {0<amount:uint256>} {1<token:address>}",
		types.ActionRedeemMax:   "Redeem max staking rewards for {0<token:address>}",
		types.ConstraintAPY:     "APY of {0<vault:address>} is {1<operator:int8>} than {2<threshold:uint256>}%.",
	}
)

type YearnV3 struct {
	*actions.BaseHandler
}

func New() actions.BaseProtocolHandler {
	return &YearnV3{
		BaseHandler: actions.NewBaseHandler(
			name,
			icon,
			tags,
			chains,
			sentences,
			&YearnV3OptionsProvider{},
		),
	}
}

func (yearnV3 *YearnV3) GetTransaction(action types.Action, rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	switch action {
	case types.ActionDeposit:
		return HandleActionDeposit(rawInputs, params)
	case types.ActionWithdraw:
		return HandleActionWithdraw(rawInputs, params)
	case types.ActionStake:
		return HandleActionStake(rawInputs, params)
	case types.ActionStakeMax:
		return HandleActionStakeMax(rawInputs, params)
	case types.ActionRedeem:
		return HandleActionRedeem(rawInputs, params)
	case types.ActionRedeemMax:
		return HandleActionRedeemMax(rawInputs, params)
	case types.ConstraintAPY:
		return HandleConstraintAPY(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
