package options

import "solver/internal/actions"

var (
	booleanOptions = []actions.Option{
		{Label: "True", Value: "true"},
		{Label: "False", Value: "false"},
	}

	logicalOperationOptions = []actions.Option{
		{Label: "AND", Value: "and"},
		{Label: "OR", Value: "or"},
		{Label: "NOT", Value: "not"},
		{Label: "XOR", Value: "xor"},
		{Label: "NAND", Value: "nand"},
		{Label: "NOR", Value: "nor"},
		{Label: "IMPLIES", Value: "implies"},
	}
)

func LogicOperationOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		0: {Simple: booleanOptions},
		1: {Simple: logicalOperationOptions},
		2: {Simple: booleanOptions},
	}, nil
}
