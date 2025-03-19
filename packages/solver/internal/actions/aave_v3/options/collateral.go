package options

import (
	"fmt"
	"math/big"
	"solver/internal/actions"
	"solver/internal/actions/aave_v3/reads"
	"strings"
)

func GetCollateralOptions(chainId uint64) ([]actions.Option, error) {
	reserves, err := reads.GetReserves(chainId)
	if err != nil {
		return nil, err
	}

	options := make([]actions.Option, 0)
	for _, reserve := range reserves {
		// TODO: (#12) Does not include 'Isolated' assets as optional collateral due to the nuance of
		// variable allowance. For this to be supported, we need arrow function support across
		// multiple actions because the depositing of a collateral asset that is isolated cannot
		// be used across all forms of borrowable assets -- The exposure is limited.
		//
		// NOTE: Right now they are filtered out by checking the `DebtCeiling` value because isolated
		// assets are the only ones that have a non-zero value. This is a bit of a hack and should
		// be revisited in the future.
		//
		// NOTE: Realistically, this is not something we will probably ever support though so the only
		// other option is to just return the isolated collateral assets and let a user
		// figure it out themselves which seems less ideal. The isolated assets however are the
		// exogenous bases while non-isolated tend to be native and stable assets.
		if !reserve.UsageAsCollateralEnabled || reserve.DebtCeiling.Cmp(big.NewInt(0)) > 0 {
			continue
		}

		rateFloat := new(big.Float).Quo(
			new(big.Float).SetInt(reserve.LiquidityRate),
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
			Info:  &actions.OptionInfo{Label: "Rate", Value: rate},
			Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(reserve.UnderlyingAsset.String()))},
		})
	}

	return options, nil
}

func CollateralOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	collateralOptions, err := GetCollateralOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: collateralOptions},
	}, nil
}
