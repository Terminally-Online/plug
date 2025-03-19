package options

import "solver/internal/actions"

var (
	riskOptions = []actions.Option{
		{
			Name:  "Liquidatable",
			Label: "liquidatable",
			Value: "1.0",
			Info:  &actions.OptionInfo{Value: "1.0"},
		}, {
			Name:  "Risky",
			Label: "risky",
			Value: "1.25",
			Info:  &actions.OptionInfo{Value: "1.25"},
		}, {
			Name:  "Safe",
			Label: "safe",
			Value: "2.0",
			Info:  &actions.OptionInfo{Value: "2.0"},
		},
	}
)

func HealthFactorOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		0: {Simple: actions.BaseThresholdFields},
		1: {Simple: riskOptions},
	}, nil
}
