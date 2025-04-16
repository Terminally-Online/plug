package plug

import (
	"solver/internal/actions"
	plug_actions "solver/internal/actions/plug/actions"
	plug_options "solver/internal/actions/plug/options"
	"solver/internal/bindings/references"
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Plug",
			Icon:   "https://cdn.onplug.io/protocols/plug.png",
			Tags:   []string{"defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ReadBalance: actions.NewActionDefinition(
					"Get balance of {0<token:address:uint256:uint256>} held by {1<holder:string>}",
					plug_actions.Balance,
					plug_options.BalanceOptions,
					actions.IsUser,
					actions.IsDynamic,
					&plug_actions.BalanceFunc,
				),
				actions.ReadPrice: actions.NewActionDefinition(
					"Get price of {0<token:address:uint256:uint256>}",
					plug_actions.Price,
					plug_options.PriceOptions,
					actions.IsUser,
					actions.IsDynamic,
					&plug_actions.PriceFunc, // TODO: Need to figure out how to handle offchain stuff.
				),
				actions.ActionSwap: actions.NewActionDefinition(
					"Swap {0<amount:float>} {1<token:address:uint256:uint256>} for {2<tokenIn:address:uint256:uint256>}",
					plug_actions.Swap,
					plug_options.SwapOptions,
					actions.IsUser,
					actions.IsDynamic,
					actions.IsEmptyOnchainFunc,
				),
				actions.ActionTransfer: actions.NewActionDefinition(
					"Transfer {0<amount:float>} {1<token:address:uint256:uint256>} to {2<recipient:string>}",
					plug_actions.Transfer,
					plug_options.TransferOptions,
					actions.IsUser,
					actions.IsDynamic,
					&plug_actions.TransferFunc,
				),
				actions.ActionDeploy: actions.NewActionDefinition(
					"Deploy Socket on {0<factory:address>} with {1<nonce:uint64>} for {2<admin:address>} with a delegate of {3<delegate:address>} with {4<implementation:address>}",
					plug_actions.Deploy,
					nil,
					actions.IsGlobal,
					actions.IsStatic,
					actions.IsEmptyOnchainFunc,
				),
			},
		},
	)
}
