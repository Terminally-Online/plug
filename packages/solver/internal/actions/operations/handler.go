package operations

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Operations"
	icon = "https://cdn.onplug.io/protocols/operations.png"
	tags = []string{"operations"}

	ActionNumber      = "number"
	ActionSetVariable = "set_variable"
	ActionGetVariable = "get_variable"

	chains  = []*references.Network{references.Mainnet}
	schemas = map[string]actions.ActionDefinition{
		ActionNumber: {
			Sentence: "{0<number:float>}",
			// Handler:  HandleNumber,
		},
		ActionSetVariable: {
			Sentence: "Set {0<variable:string>} to {1<value:float>}",
			// Handler:  HandleSetVariable,
		},
		ActionGetVariable: {
			Sentence: "Get {0<variable:string>}",
			// Handler:  HandleGetVariable,
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
		&OperationsOptionsProvider{},
	)
}
