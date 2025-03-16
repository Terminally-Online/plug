package math

import (
	"solver/internal/actions"
)

var (
	operationOptions = map[int]actions.Options{
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

func CalculateOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
	return operationOptions, nil
}
