package plug

import (
	"solver/internal/actions"
	plug_actions "solver/internal/actions/plug/actions"
	plug_options "solver/internal/actions/plug/options"
	"solver/internal/bindings/references"
)

var (
	ActionTransferSentence = "Transfer {0<amount:float>} {1<token:address:uint256:uint256>} to {2<recipient:string>}"
	ActionSwapSentence     = "Swap {0<amount:float>} {1<token:address:uint256:uint256>} for {2<tokenIn:address:uint256:uint256>}"
	ActionDeploySentence   = "Deploy Socket on {0<factory:address>} with {1<nonce:uint64>} for {2<admin:address>} with a delegate of {3<delegate:address>} with {4<implementation:address>}"
	ReadBalanceSentence    = "Get balance of {0<token:address:uint256:uint256>} held by {1<holder:string>}"
	ReadPriceSentence      = "Get price of {0<token:string>}"

	ActionTransfer = actions.NewActionDefinition(
		ActionTransferSentence,
		plug_actions.Transfer,
		plug_options.TransferOptions,
		actions.IsUser,
		actions.IsDynamic,
		actions.IsEmptyOnchainFunc,
	)
	ActionSwap = actions.NewActionDefinition(
		ActionSwapSentence,
		plug_actions.Swap,
		plug_options.SwapOptions,
		actions.IsUser,
		actions.IsDynamic,
		actions.IsEmptyOnchainFunc,
	)
	ActionDeploy = actions.NewActionDefinition(
		ActionDeploySentence,
		plug_actions.Deploy,
		nil,
		actions.IsGlobal,
		actions.IsStatic,
		actions.IsEmptyOnchainFunc,
	)
	ReadBalance = actions.NewActionDefinition(
		ReadBalanceSentence,
		plug_actions.Balance,
		plug_options.BalanceOptions,
		actions.IsUser,
		actions.IsDynamic,
		&plug_actions.Erc20BalanceFunc,
	)
	ReadPrice = actions.NewActionDefinition(
		ReadPriceSentence,
		plug_actions.Price,
		plug_options.PriceOptions,
		actions.IsUser,
		actions.IsDynamic,
		&plug_actions.PriceFunc,
	)
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Plug",
			Icon:   "https://cdn.onplug.io/protocols/plug.png",
			Tags:   []string{"defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ActionSwap:     ActionSwap,
				actions.ActionTransfer: ActionTransfer,
				actions.ActionDeploy:   ActionDeploy,
				actions.ReadBalance:    ReadBalance,
				actions.ReadPrice:      ReadPrice,
			},
		},
	)
}
