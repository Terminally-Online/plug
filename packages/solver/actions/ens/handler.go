package ens

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	SetPrimary   = "set_primary"
	GracePeriod  = "grace_period"
	TimeLeft     = "time_left"
	RenewalPrice = "renewal_price"

	sentences = map[types.Action]string{
		types.ActionBuy:            "Buy ENS {0<name:string>} with a max price of {1<maxPrice:uint256>} ETH.",
		types.ActionRenew:          "Renew ENS {0<name:string>} for {1<duration:uint256>} years.",
		types.Action(RenewalPrice): "Price to renew ENS {0<name:string>} for {1<duration:uint256>} is less than {2<price:uint256>} ETH.",
		types.Action(GracePeriod):  "ENS {0<name:string>} is in renewal grace period.",
		types.Action(TimeLeft):     "Time left in ENS {0<name:string>} is less than {1<duration:uint256>}.",
	}
)

type Ens struct {
	*actions.BaseHandler
}

func New() actions.BaseProtocolHandler {
	return &Ens{
		BaseHandler: actions.NewBaseHandler(
			"ENS",
			"https://onplug.io/protocols/ens.png",
			[]string{"naming", "web3"},
			utils.Mainnet.ChainIds,
			sentences,
			&EnsOptionsProvider{},
		),
	}
}

func (ens *Ens) GetTransaction(action types.Action, rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	switch action {
	case types.ActionBuy:
		return HandleActionBuy(rawInputs, params)
	case types.ActionRenew:
		return HandleActionRenew(rawInputs, params)
	case types.Action(RenewalPrice):
		return HandleConstraintRenewalPrice(rawInputs, params)
	case types.Action(GracePeriod):
		return HandleConstraintGracePeriod(rawInputs, params)
	case types.Action(TimeLeft):
		return HandleConstraintTimeLeft(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
