package options

import (
	"solver/internal/actions"
)

var (
	numberComparisonOptions = []actions.Option{
		{Name: "Equals", Label: "is equal to", Value: "isEqual"},
		{Name: "Does Not Equal", Label: "is not equal to", Value: "isNotEqual"},
		{Name: "Greater Than", Label: "is greater than", Value: "isGreaterThan"},
		{Name: "Greater Than or Equal", Label: "is greater than or equal to", Value: "isGreaterThanOrEqual"},
		{Name: "Less Than", Label: "is less than", Value: "isLessThan"},
		{Name: "Less Than or Equal", Label: "is less than or equal to", Value: "isLessThanOrEqual"},
	}
)

func NumberComparisonOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		1: {Simple: numberComparisonOptions},
	}, nil
}
