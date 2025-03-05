package boolean

import (
	"solver/bindings/plug_boolean"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Boolean"
	icon = "https://cdn.onplug.io/protocols/boolean.png"
	tags = []string{"logic", "condition", "comparison"}

	LogicOperation    = "logicOperation"
	CompareNumbers    = "compareNumbers"
	CompareTimes      = "compareTimes"
	CheckTimeProperty = "checkTimeProperty"

	chains  = []*references.Network{references.Mainnet, references.Base}
	schemas = map[string]actions.ActionDefinition{
		LogicOperation: {
			Sentence:     "Check if {0<a:boolean>} {1<operation:string>} {2<b:boolean>}",
			Handler:      HandleLogicOperation,
			Metadata:     plug_boolean.PlugBooleanMetaData,
		},
		CompareNumbers: {
			Sentence:     "Check if {0<value:integer>} {1<comparison:string>} {2<threshold:integer>}",
			Handler:      HandleCompareNumbers,
			Metadata:     plug_boolean.PlugBooleanMetaData,
		},
		CompareTimes: {
			Sentence:     "Check if {0<time:timestamp>} {1<comparison:string>} {2<threshold:timestamp>}",
			Handler:      HandleCompareTimes,
			Metadata:     plug_boolean.PlugBooleanMetaData,
		},
		CheckTimeProperty: {
			Sentence:     "Check if {0<timestamp:timestamp>} is a {1<property:string>}",
			Handler:      HandleCheckTimeProperty,
			Metadata:     plug_boolean.PlugBooleanMetaData,
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
		&BooleanOptionsProvider{},
	)
}
