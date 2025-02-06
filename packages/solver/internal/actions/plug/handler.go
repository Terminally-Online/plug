package plug

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Plug"
	icon = "https://cdn.onplug.io/protocols/plug.png"
	tags = []string{"defi"}

	chains = append(references.Mainnet.ChainIds, references.Base.ChainIds...)
	schemas = map[string]actions.ActionDefinition{
		actions.ActionTransfer: {
			Sentence:       "Transfer {0<amount:[(1.1)==721?1:float]>} {1<token:address:uint64:uint64>} {2<id:[(1.1)>20?uint256:null]>} to {3<recipient:address>}.",
			Handler:        HandleTransfer,
			IsUserSpecific: true,
		},
		actions.ActionSwap: {
			Sentence: "Swap {0<amount:float>} {1<tokenIn:address:uint64:uint64>} for {2<tokenOut:address>}.",
			Handler:  HandleSwap,
			IsUserSpecific: true,
		},
		actions.ConstraintPrice: {
			Sentence:       "Price of {0<token:address:uint64:uint64>} is {1<operator:int8>} than {2<threshold:float>}.",
			IsUserSpecific: true,
		},
		actions.ConstraintBalance: {
			Sentence:       "Balance of {0<token:address:uint64:uint64>} is {1<operator:int8>} than {2<threshold:float>}.",
			IsUserSpecific: true,
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
