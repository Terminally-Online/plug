package actions

import (
	"fmt"
	"math/big"
	"solver/internal/actions"
	"solver/internal/helpers/llama"
	"solver/internal/solver/signature"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type PriceRequest struct {
	Token string `json:"token"`
}

var PriceFunc = actions.ActionOnchainFunctionResponse{
	Arguments: &abi.Arguments{
		{Name: "price", Type: abi.Type{T: abi.UintTy, Size: 256}},
	},
}

func Price(lookup *actions.SchemaLookup[PriceRequest]) ([]signature.Plug, error) {
	chainName, err := llama.GetChainName(lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain name: %w", err)
	}

	parts := strings.Split(lookup.Inputs.Token, ":")
	key := fmt.Sprintf("%s:%s", chainName, parts[0])
	prices, err := llama.GetPrices([]string{key})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch token price: %w", err)
	}

	price, exists := prices[key]
	if !exists {
		return nil, fmt.Errorf("price not found for token: %s", lookup.Inputs.Token)
	}

	priceDecimal := new(big.Float).SetFloat64(price.Price)
	priceScaled := new(big.Float).Mul(priceDecimal, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))
	priceInt := new(big.Int)
	priceScaled.Int(priceInt)

	packedData, err := PriceFunc.Arguments.Pack(priceInt)
	if err != nil {
		return nil, fmt.Errorf("failed to pack price data: %w", err)
	}

	return []signature.Plug{{
		Selector: signature.ForwardedCall,
		Data:     packedData,
	}}, nil
}
