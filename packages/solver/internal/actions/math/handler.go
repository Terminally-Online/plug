package math

import (
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	name = "Math"
	icon = "https://cdn.onplug.io/protocols/math.png"
	tags = []string{"scripting"}

	Calculate = "calculate"
	Min       = "min"
	Max       = "max"
	Power     = "power"
	Clamp     = "clamp"

	chains  = []*references.Network{references.Mainnet, references.Base}
	schemas = map[string]actions.ActionDefinition{
		Calculate: {
			Sentence: "Calculate {0<x:float>} {1<operation:string>} {2<y:float>}",
			Handler:  HandleCalculate,
		},
		Min: {
			Sentence: "Get the minimum of {0<a:float>} and {1<b:float>}",
			Handler:  HandleMin,
		},
		Max: {
			Sentence: "Get the maximum of {0<a:float>} and {1<b:float>}",
			Handler:  HandleMax,
		},
		Power: {
			Sentence: "Raise {0<base:float>} to the power of {1<exponent:float>}",
			Handler:  HandlePower,
		},
		Clamp: {
			Sentence: "Clamp {0<value:float>} between {1<min:float>} and {2<max:float>}",
			Handler:  HandleClamp,
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
		&MathOptionsProvider{},
	)
}
