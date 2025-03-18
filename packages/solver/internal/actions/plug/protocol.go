package plug

import (
	"solver/internal/actions"
	plug_actions "solver/internal/actions/plug/actions"
	"solver/internal/bindings/references"
)

var (
	IS_GLOBAL = false
	IS_USER   = true

	IS_STATIC     = false
	IS_SEARCHABLE = true
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Plug",
			Icon:   "https://cdn.onplug.io/protocols/plug.png",
			Tags:   []string{"defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ActionTransfer: actions.NewActionDefinition(
					"Transfer {0<amount:float>} {1<token:address:uint256:uint256>} to {2<recipient:string>}",
					plug_actions.Transfer,
					plug_actions.TransferOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				actions.ActionSwap: actions.NewActionDefinition(
					"Swap {0<amount:float>} {1<token:address:uint256:uint256>} for {2<tokenIn:address:uint256:uint256>}",
					plug_actions.Swap,
					plug_actions.SwapOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				actions.ReadPrice: actions.NewActionDefinition(
					"Price of {0<token:address:uint256:uint256>} is {1<operator:int8>} than {2<threshold:float>}",
					plug_actions.Price,
					plug_actions.PriceOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
				actions.ReadBalance: actions.NewActionDefinition(
					"Balance of {0<token:address:uint256:uint256>} held by {1<address:address>}.",
					plug_actions.Balance,
					plug_actions.BalanceOptions,
					IS_USER,
					IS_SEARCHABLE,
				),
			},
		},
	)
}
