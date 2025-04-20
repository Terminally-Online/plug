package boolean

import (
	"solver/internal/actions"
	boolean_actions "solver/internal/actions/boolean/actions"
	boolean_options "solver/internal/actions/boolean/options"
	"solver/internal/bindings/references"
)

var (
	ActionNumberComparisonKey = "number_comparison"
	ActionNumberLogicKey      = "logic_operation"

	ActionNumberComparisonSentence = "{0<a:uint256>} {1<comparison:string>} {2<b:uint256>}"
	ActionNumberLogicSentence      = "{0<a:bool>} {1<comparison:string>} {2<b:[(1)!=isNot?bool:null]>}"

	ActionNumberComparison = actions.NewActionDefinition(
		ActionNumberComparisonSentence,
		boolean_actions.NumberComparison,
		boolean_options.NumberComparisonOptions,
		actions.IsGlobal,
		actions.IsStatic,
		&boolean_actions.NumberComparisonFunc,
	)
	ActionNumberLogic = actions.NewActionDefinition(
		ActionNumberLogicSentence,
		boolean_actions.NumberLogic,
		boolean_options.NumberLogicOptions,
		actions.IsGlobal,
		actions.IsStatic,
		&boolean_actions.NumberLogicFunc,
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
				ActionNumberComparisonKey: ActionNumberComparison,
				ActionNumberLogicKey:      ActionNumberLogic,
			},
		},
	)
}
