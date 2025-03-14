package plug

import (
	"solver/bindings/erc_20"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	Balance = "balance"
)

func New() actions.Protocol {
	return actions.New(
		actions.Protocol{
			Name:   "Plug",
			Icon:   "https://cdn.onplug.io/protocols/plug.png",
			Tags:   []string{"defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinition{
				actions.ActionTransfer: {
					Sentence:       "Transfer {0<amount:float>} {1<token:address:uint8:uint256>} to {2<recipient:string>}",
					Handler:        HandleTransfer,
					Options:        TransferOptions,
					IsUserSpecific: true,
					IsSearchable:   true,
				},
				actions.ActionSwap: {
					Sentence:       "Swap {0<amount:float>} {1<token:address:uint256:uint256>} for {2<tokenIn:address:uint256:uint256>}",
					Handler:        HandleSwap,
					Options:        SwapOptions,
					IsUserSpecific: true,
				},
				actions.ConstraintPrice: {
					Sentence:       "Price of {0<token:address:uint256:uint256>} is {1<operator:int8>} than {2<threshold:float>}",
					Handler:        HandleConstraintPrice,
					Options:        PriceOptions,
					IsUserSpecific: true,
				},
				Balance: {
					Sentence:       "Balance of {0<token:address:uint256:uint256>} held by {1<address:address>}.",
					Handler:        HandleBalance,
					Options:        BalanceOptions,
					IsUserSpecific: true,
					IsSearchable:   true,
					Metadata:       erc_20.Erc20MetaData,
					FunctionName:   "balanceOf",
				},
			},
		},
	)
}
