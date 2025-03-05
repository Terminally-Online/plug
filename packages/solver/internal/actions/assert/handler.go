package assert

import (
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
			Sentence: "Assert that {0<condition:boolean>} is true",
			Handler:  HandleAssertTrue,
		},
		AssertFalse: {
			Sentence: "Assert that {0<condition:boolean>} is false",
			Handler:  HandleAssertFalse,
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

