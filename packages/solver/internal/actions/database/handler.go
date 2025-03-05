package database

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Database"
	icon = "https://cdn.onplug.io/protocols/database.png"
	tags = []string{"database", "storage", "state"}

	SetValue    = "setValue"
	GetValue    = "getValue"
	RemoveValue = "removeValue"

	chains  = []*references.Network{references.Mainnet, references.Base}
	schemas = map[string]actions.ActionDefinition{
		SetValue: {
			Sentence: "Set {0<key:string>} to {1<value:string>}",
			Handler:  HandleSetValue,
		},
		GetValue: {
			Sentence: "Get the value of {0<key:string>}",
			Handler:  HandleGetValue,
		},
		RemoveValue: {
			Sentence: "Remove the value of {0<key:string>}",
			Handler:  HandleRemoveValue,
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
		&DatabaseOptionsProvider{},
	)
}
