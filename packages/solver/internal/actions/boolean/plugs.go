package boolean

import (
	"encoding/json"
	"fmt"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"strings"
	"time"
)

type LogicOperationInput struct {
	A         bool   `json:"a"`
	Operation string `json:"operation"`
	B         bool   `json:"b"`
}

type NumberComparisonInput struct {
	Value      int64  `json:"value"`
	Comparison string `json:"comparison"`
	Threshold  int64  `json:"threshold"`
}

type TimeCompareInput struct {
	Time       int64  `json:"time"`
	Comparison string `json:"comparison"`
	Threshold  int64  `json:"threshold"`
}

type TimePropertyInput struct {
	Timestamp int64  `json:"timestamp"`
	Property  string `json:"property"`
}

func createBooleanResultPlug(result bool) []signature.Plug {
	stateVars := make(map[string]interface{})
	stateVars["result"] = result

	plug := signature.Plug{}

	return []signature.Plug{plug}
}

func HandleLogicOperation(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs LogicOperationInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal logic operation inputs: %w", err)
	}

	operation := strings.ToLower(inputs.Operation)

	var result bool
	switch operation {
	case "and":
		result = inputs.A && inputs.B
	case "or":
		result = inputs.A || inputs.B
	case "not":
		result = !inputs.A
	case "xor":
		result = inputs.A != inputs.B
	case "nand":
		result = !(inputs.A && inputs.B)
	case "nor":
		result = !(inputs.A || inputs.B)
	case "implies":
		result = !inputs.A || inputs.B
	default:
		return nil, fmt.Errorf("unsupported logical operation: %s", inputs.Operation)
	}

	return createBooleanResultPlug(result), nil
}

func HandleCompareNumbers(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs NumberComparisonInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal number comparison inputs: %w", err)
	}

	comparison := strings.ToLower(inputs.Comparison)

	var result bool
	switch comparison {
	case "equal":
		result = inputs.Value == inputs.Threshold
	case "notequal":
		result = inputs.Value != inputs.Threshold
	case "greaterthan":
		result = inputs.Value > inputs.Threshold
	case "greaterthanorequal":
		result = inputs.Value >= inputs.Threshold
	case "lessthan":
		result = inputs.Value < inputs.Threshold
	case "lessthanorequal":
		result = inputs.Value <= inputs.Threshold
	case "between":
		min := inputs.Value / 2 // Default fallback
		max := inputs.Threshold
		result = inputs.Value >= min && inputs.Value <= max
	default:
		return nil, fmt.Errorf("unsupported number comparison: %s", inputs.Comparison)
	}

	return createBooleanResultPlug(result), nil
}

func HandleCompareTimes(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs TimeCompareInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal time comparison inputs: %w", err)
	}

	comparison := strings.ToLower(inputs.Comparison)

	var result bool
	switch comparison {
	case "before":
		result = inputs.Time < inputs.Threshold
	case "after":
		result = inputs.Time > inputs.Threshold
	case "between":
		start := inputs.Time / 2
		end := inputs.Threshold
		result = inputs.Time >= start && inputs.Time <= end
	case "sameday":
		time1 := time.Unix(inputs.Time, 0)
		time2 := time.Unix(inputs.Threshold, 0)
		year1, month1, day1 := time1.Date()
		year2, month2, day2 := time2.Date()
		result = year1 == year2 && month1 == month2 && day1 == day2
	default:
		return nil, fmt.Errorf("unsupported time comparison: %s", inputs.Comparison)
	}

	return createBooleanResultPlug(result), nil
}

func HandleCheckTimeProperty(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs TimePropertyInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal time property inputs: %w", err)
	}

	property := strings.ToLower(inputs.Property)
	t := time.Unix(inputs.Timestamp, 0)
	weekday := t.Weekday()

	var result bool
	switch property {
	case "weekday":
		result = weekday >= time.Monday && weekday <= time.Friday
	case "weekend":
		result = weekday == time.Saturday || weekday == time.Sunday
	default:
		return nil, fmt.Errorf("unsupported time property: %s", inputs.Property)
	}

	return createBooleanResultPlug(result), nil
}
