package assert

import (
	"solver/bindings/plug_assert"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	Assert = "assert"
)

func New() actions.Protocol {
	return actions.New(
		actions.Protocol{
			Name:   "Assert",
			Icon:   "https://cdn.onplug.io/protocols/assert.png",
			Tags:   []string{"validation", "assert", "condition"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinition{
				Assert: {
					Sentence:     "Assert that {0<condition:bool>} is {1<assertion:bool>}",
					Handler:      HandleAssert,
					Options:      AssertOptions,
					Metadata:     plug_assert.PlugAssertMetaData,
					FunctionName: "assertTrue",
				},
			},
		},
	)
}
