package options

import "solver/internal/actions"

var (
	booleanOptions = []actions.Option{
		{Name: "True", Label: "true", Value: "true"},
		{Name: "False", Label: "false", Value: "false"},
	}

	logicalOperationOptions = []actions.Option{
		{Name: "And", Label: "and", Value: "isAnd"},
		{Name: "Or", Label: "or", Value: "isOr"},
		{Name: "Not", Label: "is not true", Value: "isNot"},
		{Name: "Xor", Label: "exclusive or", Value: "isXor"},
		{Name: "Nand", Label: "is not and", Value: "isNand"},
		{Name: "Nor", Label: "nor", Value: "isNor"},
		{Name: "Implies", Label: "implies", Value: "isImplies"},
	}
)

func NumberLogicOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		0: {Simple: booleanOptions},
		1: {Simple: logicalOperationOptions},
		2: {Simple: booleanOptions},
	}, nil
}
