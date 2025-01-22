package plug

import (
	"solver/actions"
	"solver/types"
	"solver/utils"
)

var (
	name = "Plug"
	icon = "https://cdn.onplug.io/protocols/plug.png"
	tags = []string{"defi"}

	chains = utils.Mainnet.ChainIds

	schemas = map[types.Action]actions.ActionDefinition{
		types.ActionTransfer: {
			Sentence: "Transfer {0<amount:[(1.1)=721?'1':float]>} {1<token:address:uint256>} {2<id:[(1.1)>20?uint256:null]>} to {3<recipient:address>}",
			Handler:  HandleTransfer,
		},
		types.ActionTransferFrom: {
			Sentence: "Transfer {0<amount:uint256>} {1<token:address>} to {2<recipient:address>}.",
			Handler:  HandleTransferFrom,
		},
		types.ActionSwap: {
			Sentence: "Swap {0<amount:float>} {1<tokenIn:address>} for {2<tokenOut:address>}.",
			Handler:  HandleSwap,
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
		&PlugOptionsProvider{},
	)
}
