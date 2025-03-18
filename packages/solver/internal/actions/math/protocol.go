package math

import (
	"solver/internal/actions"
	math_actions "solver/internal/actions/math/actions"
	math_options "solver/internal/actions/math/options"
	"solver/internal/bindings/references"
)

var (
	IS_GLOBAL = false
	IS_USER   = true

	IS_STATIC     = false
	IS_SEARCHABLE = true
)

var (
	Calculate = "calculate"
	Min       = "min"
	Max       = "max"
	Power     = "power"
	Clamp     = "clamp"
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Math",
			Icon:   "https://cdn.onplug.io/protocols/math.png",
			Tags:   []string{"scripting"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				Calculate: actions.NewActionDefinition(
					"Calculate {0<x:uint256>} {1<operation:string>} {2<y:uint256>}",
					math_actions.Calculate,
					math_options.CalculateOptions,
					IS_GLOBAL,
					IS_STATIC,
				),
				Min: actions.NewActionDefinition(
					"Get the minimum of {0<a:uint256>} and {1<b:uint256>}",
					math_actions.Min,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
				Max: actions.NewActionDefinition(
					"Get the maximum of {0<a:uint256>} and {1<b:uint256>}",
					math_actions.Max,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
				Power: actions.NewActionDefinition(
					"Raise {0<base:uint256>} to the power of {1<exponent:uint256>}",
					math_actions.Power,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
				Clamp: actions.NewActionDefinition(
					"Clamp {0<value:uint256>} between {1<min:uint256>} and {2<max:uint256>}",
					math_actions.HandleClamp,
					nil,
					IS_GLOBAL,
					IS_STATIC,
				),
			},
		},
	)
}
