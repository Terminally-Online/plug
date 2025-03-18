package actions

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/helpers/llama"
	"solver/internal/solver/signature"
)

type PriceRequest struct {
	Token string `json:"token"`
}

func Price(lookup *actions.SchemaLookup[PriceRequest]) ([]signature.Plug, error) {
	tokenId := fmt.Sprintf("ethereum:%s", lookup.Inputs.Token)
	prices, err := llama.GetPrices([]string{tokenId})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token price: %w", err)
	}

	_, exists := prices[tokenId]
	if !exists {
		return nil, fmt.Errorf("price not found for token: %s", lookup.Inputs.Token)
	}

	return nil, nil
}
