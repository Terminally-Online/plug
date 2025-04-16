package actions

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/helpers/llama"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type PriceRequest struct {
	Token string `json:"token"`
}

var PriceFunc = actions.ActionOnchainFunctionResponse{
	Arguments: abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 256}},
	},
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
