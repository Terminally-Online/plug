package math

import (
	"solver/internal/actions"
	math_actions "solver/internal/actions/math/actions"
	math_options "solver/internal/actions/math/options"
	"solver/internal/bindings/references"
)

var (
	ActionCalculateKey        = "calculate"
	ActionMinimumOrMaximumKey = "minimum_or_maximum"

	ActionCalculateSentence        = "Calculate {0<x:uint256>} {1<operation:string>} {2<y:uint256>}"
	ActionMinimumOrMaximumSentence = "Get the {0<operation:string>} of {1<x:uint256>} and {2<y:uint256>}"

	ActionCalculate = actions.NewActionDefinition(
		ActionCalculateSentence,
		math_actions.Calculate,
		math_options.CalculateOptions,
		nil,
		&math_actions.CalculateFunc,
	)
	ActionMinimumOrMaximum = actions.NewActionDefinition(
		ActionMinimumOrMaximumSentence,
		math_actions.MinimumOrMaximum,
		math_options.MinimumOrMaximumOptions,
		nil,
		&math_actions.MinimumOrMaximumFunc,
	)
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Math",
			Icon:   "https://cdn.onplug.io/protocols/math.png",
			Tags:   []string{"scripting"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				ActionCalculateKey:        ActionCalculate,
				ActionMinimumOrMaximumKey: ActionMinimumOrMaximum,
			},
		},
	)
}
