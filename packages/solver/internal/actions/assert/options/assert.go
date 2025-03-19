package options

import "solver/internal/actions"

func AssertOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	booleanOptions := []actions.Option{
		{Label: "True", Value: "true"},
		{Label: "False", Value: "false"},
	}

	return map[int]actions.Options{
		1: {Simple: booleanOptions},
	}, nil
}
