package assert

import (
	"solver/bindings/plug_assert"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Assert"
	icon = "https://cdn.onplug.io/protocols/assert.png"
	color = "#F3908A" // red
	tags = []string{"validation", "assert", "condition"}

	AssertTrue  = "assertTrue"
	AssertFalse = "assertFalse"

	chains  = []*references.Network{references.Mainnet, references.Base}
	schemas = map[string]actions.ActionDefinition{
		AssertTrue: {
			Sentence:     "Assert that {0<condition:boolean>} is true {1<message:string>}",
			Handler:      HandleAssertTrue,
			Metadata:     plug_assert.PlugAssertMetaData,
			FunctionName: "assertTrue",
		},
		AssertFalse: {
			Sentence:     "Assert that {0<condition:boolean>} is false {1<message:string>}",
			Handler:      HandleAssertFalse,
			Metadata:     plug_assert.PlugAssertMetaData,
			FunctionName: "assertFalse",
		},
	}
)

func New() actions.BaseProtocolHandler {
	return actions.NewBaseHandler(
		name,
		icon,
		color,
		tags,
		chains,
		schemas,
		&AssertOptionsProvider{},
	)
}

