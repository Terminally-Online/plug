package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type APYRequest struct {
	Direction int    `json:"direction"` // -1 for borrow, 1 for supply
	Vault     string `json:"vault"`
}

func APY(lookup *actions.SchemaLookup[APYRequest]) ([]signature.Plug, error) {
	// borrowApy, supplyApy, err := reads.GetVaultApy(lookup.Inputs.Vault, lookup.ChainId)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to get vault APY: %w", err)
	// }
	//
	// var currentRate float64
	// switch lookup.Inputs.Direction {
	// case -1:
	// 	currentRate = utils.UintToFloat(borrowApy, 20)
	// case 1:
	// 	currentRate = utils.UintToFloat(supplyApy, 20)
	// default:
	// 	return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (supply), got %d", lookup.Inputs.Direction)
	// }

	return nil, nil
}
