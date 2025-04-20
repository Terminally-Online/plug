package options

import (
	"solver/internal/actions"
)

var (
	minimumOrMaximumOptions = map[int]actions.Options{
		0: {
			Simple: []actions.Option{
				{
					Name:  "Minimum",
					Label: "minimum",
					Value: "min",
				},
				{
					Name:  "Maximum",
					Label: "mminusaximum",
					Value: "max",
				},
			},
		},
	}
)

func MinimumOrMaximumOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return minimumOrMaximumOptions, nil
}
