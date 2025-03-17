package boolean

import (
	"solver/bindings/plug_boolean"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	LogicOperation = "logic_peration"
	CompareNumbers = "compare_numbers"
)

func New() actions.Protocol {
	return actions.New(
		actions.Protocol{
			Name:   "Boolean",
			Icon:   "https://cdn.onplug.io/protocols/boolean.png",
			Tags:   []string{"logic", "condition", "comparison"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinition{
				LogicOperation: {
					Sentence:     "Check if {0<a:boolean>} {1<operation:string>} {2<b:boolean>}",
					Handler:      HandleLogicOperation,
					Options:      LogicOperationOptions,
					Metadata:     plug_boolean.PlugBooleanMetaData,
					FunctionName: "isAnd",
				},
				CompareNumbers: {
					Sentence:     "Check if {0<a:integer>} {1<comparison:string>} {2<b:integer>}",
					Handler:      HandleCompareNumbers,
					Options:      CompareNumbersOptions,
					Metadata:     plug_boolean.PlugBooleanMetaData,
					FunctionName: "isEqual",
				},
			},
		},
	)
}
