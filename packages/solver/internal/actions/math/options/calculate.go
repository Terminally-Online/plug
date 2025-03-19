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
					Label: "+",
					Value: "+",
				},
				{
					Name:  "Subtract",
					Label: "-",
					Value: "-",
				},
				{
					Name:  "Multiply",
					Label: "*",
					Value: "*",
				},
				{
					Name:  "Divide",
					Label: "÷",
					Value: "÷",
				},
				{
					Name:  "Modulo",
					Label: "%",
					Value: "%",
				},
			},
		},
	}
)

func CalculateOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return calculateOptions, nil
}
