package boolean

import (
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

type BooleanOptionsProvider struct{}

func (m *BooleanOptionsProvider) GetOptions(chainId uint64, _ common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	booleanOptions := []actions.Option{
		{Label: "True", Value: "true"},
		{Label: "False", Value: "false"},
	}

	logicalOperationOptions := []actions.Option{
		{Label: "AND", Value: "and"},
		{Label: "OR", Value: "or"},
		{Label: "NOT", Value: "not"},
		{Label: "XOR", Value: "xor"},
		{Label: "NAND", Value: "nand"},
		{Label: "NOR", Value: "nor"},
		{Label: "IMPLIES", Value: "implies"},
	}

	numberComparisonOptions := []actions.Option{
		{Label: "Equals (=)", Value: "equal"},
		{Label: "Does Not Equal (≠)", Value: "notEqual"},
		{Label: "Greater Than (>)", Value: "greaterThan"},
		{Label: "Greater Than or Equal (≥)", Value: "greaterThanOrEqual"},
		{Label: "Less Than (<)", Value: "lessThan"},
		{Label: "Less Than or Equal (≤)", Value: "lessThanOrEqual"},
		{Label: "Is Between", Value: "between"},
	}
	
	timeComparisonOptions := []actions.Option{
		{Label: "Is Before", Value: "before"},
		{Label: "Is After", Value: "after"},
		{Label: "Is Between", Value: "between"},
		{Label: "Is On Same Day As", Value: "sameDay"},
	}
	
	timePropertyOptions := []actions.Option{
		{Label: "Weekday", Value: "weekday"},
		{Label: "Weekend", Value: "weekend"},
	}

	switch action {
	case LogicOperation:
		return map[int]actions.Options{
			0: {Simple: booleanOptions},
			1: {Simple: logicalOperationOptions},
			2: {Simple: booleanOptions},
		}, nil
	case CompareNumbers:
		return map[int]actions.Options{
			1: {Simple: numberComparisonOptions},
		}, nil
	case CompareTimes:
		return map[int]actions.Options{
			1: {Simple: timeComparisonOptions},
		}, nil
	case CheckTimeProperty:
		return map[int]actions.Options{
			1: {Simple: timePropertyOptions},
		}, nil
	default:
		return nil, nil
	}
}
