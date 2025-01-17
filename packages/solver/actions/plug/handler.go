package plug

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	sentences = map[types.Action]string{
		types.ActionTransfer:     "Transfer {0<amount:[(1.1)=721?1:uint256]>} {1<token:address:uint256>} {2<id:[(1.1)>20?uint256:null]>} to {3<recipient:address>}",
		types.ActionTransferFrom: "Transfer {0<amount:uint256>} {1<token:address>} to {2<recipient:address>}.",
		types.ActionSwap:         "Swap {0<amount:uint256>} {1<tokenIn:address>} for {2<tokenOut:address>}.",
	}
)

type Plug struct {
	*actions.BaseHandler
}

func New() actions.BaseProtocolHandler {
	return &Plug{
		BaseHandler: actions.NewBaseHandler(
			"Plug",
			"https://onplug.io/protocols/plug.png",
			[]string{"defi"},
			utils.Mainnet.ChainIds,
			sentences,
			&PlugOptionsProvider{},
		),
	}
}

func (plug *Plug) GetTransaction(action types.Action, rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	switch action {
	case types.ActionTransfer:
		return HandleTransfer(rawInputs, params)
	case types.ActionTransferFrom:
		return HandleTransferFrom(rawInputs, params)
	case types.ActionSwap:
		return HandleSwap(rawInputs, params)
	default:
		return nil, fmt.Errorf("unsupported action: %s", action)
	}
}
