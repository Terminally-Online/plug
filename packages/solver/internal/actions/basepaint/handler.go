package basepaint

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

func New() actions.Protocol {
	return actions.New(actions.Protocol{
		Name:   "BasePaint",
		Icon:   "https://cdn.onplug.io/protocols/basepaint.png",
		Tags:   []string{"lending", "defi"},
		Chains: []*references.Network{references.Base},
		Actions: map[string]actions.ActionDefinition{
			actions.ActionMint: {
				Sentence: "Mint {0<count:uint64>} of the latest canvas to {1<recipient:address>}",
				Handler:  MintLatest,
				Options:  MintLatestOptions,
			},
		},
	})
}
