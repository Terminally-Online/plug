package math

import (
	"solver/bindings/plug_math"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	Calculate = "calculate"
	Min       = "min"
	Max       = "max"
	Power     = "power"
	Clamp     = "clamp"
)

func New() actions.Protocol {
	return actions.New(
		actions.Protocol{
			Name:   "Math",
			Icon:   "https://cdn.onplug.io/protocols/math.png",
			Tags:   []string{"scripting"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinition{
				Calculate: {
					Sentence:     "Calculate {0<x:uint256>} {1<operation:string>} {2<y:uint256>}",
					Handler:      HandleCalculate,
					Options:      CalculateOptions,
					Metadata:     plug_math.PlugMathMetaData,
					FunctionName: "add",
				},
				Min: {
					Sentence:     "Get the minimum of {0<a:uint256>} and {1<b:uint256>}",
					Handler:      HandleMin,
					Metadata:     plug_math.PlugMathMetaData,
					FunctionName: "min",
				},
				Max: {
					Sentence:     "Get the maximum of {0<a:uint256>} and {1<b:uint256>}",
					Handler:      HandleMax,
					Metadata:     plug_math.PlugMathMetaData,
					FunctionName: "max",
				},
				Power: {
					Sentence:     "Raise {0<base:uint256>} to the power of {1<exponent:uint256>}",
					Handler:      HandlePower,
					Metadata:     plug_math.PlugMathMetaData,
					FunctionName: "power",
				},
				Clamp: {
					Sentence:     "Clamp {0<value:uint256>} between {1<min:uint256>} and {2<max:uint256>}",
					Handler:      HandleClamp,
					Metadata:     plug_math.PlugMathMetaData,
					FunctionName: "clamp",
				},
			},
		},
	)
}
