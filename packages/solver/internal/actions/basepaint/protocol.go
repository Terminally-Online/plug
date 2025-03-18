package basepaint

import (
	"solver/internal/actions"
	basepaint_actions "solver/internal/actions/basepaint/actions"
	basepaint_options "solver/internal/actions/basepaint/options"
	"solver/internal/bindings/references"
)

func New() actions.Protocol {
	return actions.NewProtocol(actions.Protocol{
		Name:   "BasePaint",
		Icon:   "https://cdn.onplug.io/protocols/basepaint.png",
		Tags:   []string{"lending", "defi"},
		Chains: []*references.Network{references.Base},
		Actions: map[string]actions.ActionDefinitionInterface{
			actions.ActionMint: actions.NewActionDefinition(
				"Mint {0<count:uint64>} of the latest canvas to {1<recipient:address>}",
				basepaint_actions.MintLatest,
				basepaint_options.MintLatestOptions,
				actions.IsStatic,
				actions.IsGlobal,
				actions.IsEmptyOnchainFunc,
			),
		},
	})
}
