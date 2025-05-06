package basepaint

import (
	"solver/internal/actions"
	basepaint_actions "solver/internal/actions/basepaint/actions"
	basepaint_options "solver/internal/actions/basepaint/options"
	"solver/internal/bindings/references"
)

var (
	ActionMintSentence = "Mint {0<count:uint64>} copy of the current canvas to {1<recipient:string>}"

	ActionMint = actions.NewActionDefinition(
		ActionMintSentence,
		basepaint_actions.MintLatest,
		basepaint_options.MintLatestOptions,
		&actions.ActionProperties{
			IsSearchable:   true,
			IsUserSpecific: true,
		},
		&basepaint_actions.MintLatestFunc,
	)
)

func New() actions.Protocol {
	return actions.NewProtocol(actions.Protocol{
		Name:   "BasePaint",
		Icon:   "https://cdn.onplug.io/protocols/basepaint.png",
		Tags:   []string{"lending", "defi"},
		Chains: []*references.Network{references.Base},
		Actions: map[string]actions.ActionDefinitionInterface{
			actions.ActionMint: ActionMint,
		},
	})
}
