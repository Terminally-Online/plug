package assert

import (
	"solver/internal/actions"
	assert_actions "solver/internal/actions/assert/actions"
	assert_options "solver/internal/actions/assert/options"
	"solver/internal/bindings/references"
)

var (
	Assert = "assert"
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Assert",
			Icon:   "https://cdn.onplug.io/protocols/assert.png",
			Tags:   []string{"validation", "assert", "condition"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]any{
				Assert: actions.NewActionDefinition(
					"Assert that {0<condition:bool>} is {1<assertion:bool>}",
					assert_actions.HandleAssert,
					assert_options.AssertOptions,
					false,
					false,
				),
			},
		},
	)
}
