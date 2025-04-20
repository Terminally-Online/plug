package boolean

import (
	"solver/internal/actions"
	boolean_actions "solver/internal/actions/boolean/actions"
	boolean_options "solver/internal/actions/boolean/options"
	"solver/internal/bindings/references"
)

var (
	LogicOperation = "logic_peration"

	ActionCompareNumbersKey = "compare_numbers"

	ActionCompareNumbers = actions.NewActionDefinition(
		"Check if {0<a:uint256>} {1<comparison:string>} {2<b:uint256>}",
		boolean_actions.CompareNumbers,
		boolean_options.CompareNumbersOptions,
		actions.IsGlobal,
		actions.IsStatic,
		actions.IsEmptyOnchainFunc,
	)
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Boolean",
			Icon:   "https://cdn.onplug.io/protocols/boolean.png",
			Tags:   []string{"logic", "condition", "comparison"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				ActionCompareNumbersKey: ActionCompareNumbers,
			},
		},
	)
}
