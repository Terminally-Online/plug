package assert

import (
	"solver/bindings/plug_assert"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Assert"
	icon = "https://cdn.onplug.io/protocols/assert.png"
	tags = []string{"validation", "assert", "condition"}

	AssertTrue  = "assertTrue"
	AssertFalse = "assertFalse"

	chains  = []*references.Network{references.Mainnet, references.Base}
	schemas = map[string]actions.ActionDefinition{
		AssertTrue: {
			Sentence:     "Assert that {0<condition:bool>} is {1<assertion:bool>}",
			Handler:      HandleAssertTrue,
			Metadata:     plug_assert.PlugAssertMetaData,
			FunctionName: "assertTrue",
		},
	}
)

func New() actions.BaseProtocolHandler {
	return actions.NewBaseHandler(
		name,
		icon,
		tags,
		chains,
		schemas,
		&AssertOptionsProvider{},
	)
}
