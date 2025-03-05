package database

import (
	"solver/bindings/plug_database"
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
			Sentence:     "Set {0<key:string>} to {1<value:string>}",
			Handler:      HandleSetValue,
			Metadata:     plug_database.PlugDatabaseMetaData,
			FunctionName: "set",
		},
		GetValue: {
			Sentence:     "Get the value of {0<key:string>} from {1<address:address>}",
			Handler:      HandleGetValue,
			Metadata:     plug_database.PlugDatabaseMetaData,
			FunctionName: "get",
		},
		RemoveValue: {
			Sentence:     "Remove the value of {0<key:string>}",
			Handler:      HandleRemoveValue,
			Metadata:     plug_database.PlugDatabaseMetaData,
			FunctionName: "remove",
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
