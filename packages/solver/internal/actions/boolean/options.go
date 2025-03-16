package boolean

import (
	"solver/internal/actions"
)

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

	numberComparisonOptions = []actions.Option{
		{Label: "Equals (=)", Value: "equal"},
		{Label: "Does Not Equal (≠)", Value: "notEqual"},
		{Label: "Greater Than (>)", Value: "greaterThan"},
		{Label: "Greater Than or Equal (≥)", Value: "greaterThanOrEqual"},
		{Label: "Less Than (<)", Value: "lessThan"},
		{Label: "Less Than or Equal (≤)", Value: "lessThanOrEqual"},
	}
)

func LogicOperationOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		0: {Simple: booleanOptions},
		1: {Simple: logicalOperationOptions},
		2: {Simple: booleanOptions},
	}, nil
}

func CompareNumbersOptions(lookup *actions.SchemaLookup) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		1: {Simple: numberComparisonOptions},
	}, nil
}
