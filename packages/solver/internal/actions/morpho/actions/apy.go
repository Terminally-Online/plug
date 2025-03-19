package actions

import (
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

type APYRequest struct {
	Direction int    `json:"direction"` // -1 for borrow, 1 for deposit
	Target    string `json:"target"`    // Underlying market or vault
}

func APY(lookup *actions.SchemaLookup[APYRequest]) ([]signature.Plug, error) {
	// var currentRate float64
	// if len(lookup.Inputs.Target) == 42 {
	// 	vault, err := reads.GetVault(lookup.Inputs.Target, lookup.ChainId)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to fetch vault: %w", err)
	// 	}

	// 	if lookup.Inputs.Direction != 1 {
	// 		return nil, fmt.Errorf("vaults only support deposit direction (1), got %d", lookup.Inputs.Direction)
	// 	}

	// 	currentRate = vault.DailyApys.NetApy * 100
	// } else {
	// 	market, err := reads.GetMarket(lookup.Inputs.Target, lookup.ChainId)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to fetch market: %w", err)
	// 	}

	// 	switch lookup.Inputs.Direction {
	// 	case -1:
	// 		currentRate = market.DailyApys.BorrowApy * 100
	// 	case 1:
	// 		currentRate = market.DailyApys.SupplyApy * 100
	// 	default:
	// 		return nil, fmt.Errorf("invalid direction: must be either -1 (borrow) or 1 (deposit), got %d", lookup.Inputs.Direction)
	// 	}
	// }

	// TODO: Get this value onchain.

	return nil, nil
}
