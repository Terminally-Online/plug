package options

import (
	"solver/internal/actions"
)

var (
	numberComparisonOptions = []actions.Option{
		{Name: "Equals", Label: "=", Value: "isEqual"},
		{Name: "Does Not Equal", Label: "≠", Value: "isNotEqual"},
		{Name: "Greater Than", Label: ">", Value: "isGreaterThan"},
		{Name: "Greater Than or Equal", Label: "≥", Value: "isGreaterThanOrEqual"},
		{Name: "Less Than", Label: "<", Value: "isLessThan"},
		{Name: "Less Than or Equal", Label: "≤", Value: "isLessThanOrEqual"},
	}
)

func CompareNumbersOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		1: {Simple: numberComparisonOptions},
	}, nil
}
