package plug

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Plug"
	icon = "https://cdn.onplug.io/protocols/plug.png"
	tags = []string{"defi"}

	chains  = []*references.Network{references.Mainnet, references.Base}
	schemas = map[string]actions.ActionDefinition{
		actions.ActionTransfer: {
			Sentence:       "Transfer {0<amount:float>} {1<token:address:uint8:uint256>} to {2<recipient:string>}",
			Handler:        HandleTransfer,
			IsUserSpecific: true,
			IsSearchable:   true,
		},
		actions.ActionSwap: {
			Sentence:       "Swap {0<amount:float>} {1<token:address:uint256:uint256>} for {2<tokenIn:address:uint256:uint256>}",
			Handler:        HandleSwap,
			IsUserSpecific: true,
		},
		actions.ConstraintPrice: {
			Sentence:       "Price of {0<token:address:uint256:uint256>} is {1<operator:int8>} than {2<threshold:float>}",
			Handler:        HandleConstraintPrice,
			IsUserSpecific: true,
		},
		actions.ConstraintBalance: {
			Sentence:       "Balance of {0<token:address:uint256:uint256>} at {1<address:address>} is {2<operator:int8>} than {3<threshold:float>}",
			Handler:        HandleConstraintBalance,
			IsUserSpecific: true,
			IsSearchable:   true,
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
