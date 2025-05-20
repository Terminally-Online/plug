package options

import (
	"solver/internal/actions"
)

var (
	calculateOptions = map[int]actions.Options{
		1: {
			Simple: []actions.Option{
				{
					Name:  "Add",
					Label: "plus",
					Value: "add",
				},
				{
					Name:  "Subtract",
					Label: "minus",
					Value: "subtract",
				},
				{
					Name:  "Multiply",
					Label: "times",
					Value: "multiply",
				},
				{
					Name:  "Divide",
					Label: "divided by",
					Value: "divide",
				},
				{
					Name:  "Modulo",
					Label: "modulo",
					Value: "modulo",
				},
				{
					Name:  "Power Of",
					Label: "to the power of",
					Value: "power",
				},
			},
		},
	}
)

func CalculateOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return calculateOptions, nil
}
