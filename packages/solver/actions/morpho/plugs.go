package morpho

import (
	"encoding/json"
	"fmt"
	"solver/actions"
	"solver/types"
)

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Direction    int     `json:"direction"`    // -1 for borrow, 1 for deposit
		Address   string  `json:"address"`   // Underlying market or vault
		Operator  int     `json:"operator"`  // -1 for less than, 1 for greater than
		Threshold float64 `json:"threshold"` // Percentage
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal apy constraint inputs: %w", err)
	}

	markets, err := GetMarkets()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch markets: %w", err)
	}

	var market *Market
	for _, m := range markets {
		if m.UniqueKey == inputs.Address {
			market = &m
			break
		}
	}

	if market == nil {
		return nil, fmt.Errorf("market not found for address: %s", inputs.Address)
	}

	var currentRate float64
	switch inputs.Direction {
	case -1:
		currentRate = market.State.BorrowApy * 100 // Convert to percentage
	case 1:
		currentRate = market.State.SupplyApy * 100 // Convert to percentage
	default:
		return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (deposit), got %d", inputs.Direction)
	}

	switch inputs.Operator {
	case -1:
		if currentRate >= inputs.Threshold {
			return nil, fmt.Errorf("current rate %.2f%% is not less than threshold %.2f%%", currentRate, inputs.Threshold)
		}
	case 1:
		if currentRate <= inputs.Threshold {
			return nil, fmt.Errorf("current rate %.2f%% is not greater than threshold %.2f%%", currentRate, inputs.Threshold)
		}
	default:
		return nil, fmt.Errorf("invalid operator: must be either -1 (less than) or 1 (greater than), got %d", inputs.Operator)
	}

	return nil, nil
}
