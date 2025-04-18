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
				"Mint {0<count:uint64>} copy of the current canvas to {1<recipient:string>}",
				basepaint_actions.MintLatest,
				basepaint_options.MintLatestOptions,
				actions.IsUser,
				actions.IsDynamic,
				&basepaint_actions.MintLatestFunc,
			),
		},
	})
}
