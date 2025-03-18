package boolean

import (
	"solver/internal/actions"
	boolean_actions "solver/internal/actions/boolean/actions"
	boolean_options "solver/internal/actions/boolean/options"
	"solver/internal/bindings/references"
)

var (
	LogicOperation = "logic_peration"
	CompareNumbers = "compare_numbers"
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Boolean",
			Icon:   "https://cdn.onplug.io/protocols/boolean.png",
			Tags:   []string{"logic", "condition", "comparison"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				LogicOperation: actions.NewActionDefinition(
					"Check if {0<a:boolean>} {1<operation:string>} {2<b:boolean>}",
					boolean_actions.LogicOperation,
					boolean_options.LogicOperationOptions,
					false,
					false,
				),
				CompareNumbers: actions.NewActionDefinition(
					"Check if {0<a:integer>} {1<comparison:string>} {2<b:integer>}",
					boolean_actions.CompareNumbers,
					boolean_options.CompareNumbersOptions,
					false,
					false,
				),
			},
		},
	)
}
