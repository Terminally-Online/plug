package options

import (
	"solver/internal/actions"
)

var (
	numberComparisonOptions = []actions.Option{
		{Label: "Equals (=)", Value: "equal"},
		{Label: "Does Not Equal (≠)", Value: "notEqual"},
		{Label: "Greater Than (>)", Value: "greaterThan"},
		{Label: "Greater Than or Equal (≥)", Value: "greaterThanOrEqual"},
		{Label: "Less Than (<)", Value: "lessThan"},
		{Label: "Less Than or Equal (≤)", Value: "lessThanOrEqual"},
	}
)


func CompareNumbersOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	return map[int]actions.Options{
		1: {Simple: numberComparisonOptions},
	}, nil
}
