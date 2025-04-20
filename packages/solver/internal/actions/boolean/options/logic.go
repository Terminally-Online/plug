package options

import "solver/internal/actions"

var (
	booleanOptions = []actions.Option{
		{Label: "True", Value: "true"},
		{Label: "False", Value: "false"},
	}

	logicalOperationOptions = []actions.Option{
		{Name: "And", Label: "and", Value: "isAnd"},
		{Name: "Or", Label: "or", Value: "isOr"},
		{Name: "Not", Label: "not", Value: "isNot"},
		{Name: "Xor", Label: "xor", Value: "isXor"},
		{Name: "Nand", Label: "nand", Value: "isNand"},
		{Name: "Nor", Label: "nor", Value: "isNor"},
		{Name: "Implies", Label: "implies", Value: "isImplies"},
	}
)

func LogicOperationOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		0: {Simple: booleanOptions},
		1: {Simple: logicalOperationOptions},
		2: {Simple: booleanOptions},
	}, nil
}
