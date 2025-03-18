package options

import (
	"fmt"
	"math/big"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3/reads"
	"strings"
)


func GetBorrowOptions(chainId uint64) ([]actions.Option, error) {
	reserves, err := reads.GetReserves(chainId)
	if err != nil {
		return nil, err
	}

	options := make([]actions.Option, 0)
	for _, reserve := range reserves {
		if !reserve.BorrowingEnabled {
			continue
		}

		rateFloat := new(big.Float).Quo(
			new(big.Float).SetInt(reserve.VariableBorrowRate),
			new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(25), nil)),
		)

		var rate string
		if rateFloat.Cmp(big.NewFloat(0)) > 0 && rateFloat.Cmp(big.NewFloat(0.01)) < 0 {
			rate = "<0.01%"
		} else {
			rate = rateFloat.Text('f', 2) + "%"
		}

		options = append(options, actions.Option{
			Label: reserve.Symbol,
			Value: fmt.Sprintf("%s:%d", reserve.UnderlyingAsset.String(), uint8(reserve.Decimals.Uint64())),
			Name:  reserve.Name,
			Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(reserve.UnderlyingAsset.String()))},
			Info:  &actions.OptionInfo{Label: "Rate", Value: rate},
		})
	}

	return options, nil
}

func BorrowOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	borrowOptions, err := GetBorrowOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: borrowOptions},
	}, nil
}
