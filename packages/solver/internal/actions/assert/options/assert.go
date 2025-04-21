package options

import "solver/internal/actions"

func AssertOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	booleanOptions := []actions.Option{
		{Name: "True", Label: "true", Value: "true"},
		{Name: "False", Label: "false", Value: "false"},
	}

	return map[int]actions.Options{
		1: {Simple: booleanOptions},
	}, nil
}
