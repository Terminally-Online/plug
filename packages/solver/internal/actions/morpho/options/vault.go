package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
)

func GetSupplyTokenToVaultOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
	vaults, err := reads.GetVaults(chainId)
	if err != nil {
		return nil, nil, err
	}

	seenToken := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToVaultOptions := make(map[string][]actions.Option)

	for _, vault := range vaults {
		assetAddress := fmt.Sprintf("%s:%d", vault.Asset.Address, vault.Asset.Decimals)
		if !seenToken[assetAddress] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: vault.Asset.Symbol,
				Name:  vault.Asset.Name,
				Value: assetAddress,
				Icon:  &actions.OptionIcon{Default: vault.Asset.LogoURI},
			})
			seenToken[assetAddress] = true
		}

		tokenToVaultOptions[assetAddress] = append(tokenToVaultOptions[assetAddress], actions.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: assetAddress,
			Icon:  &actions.OptionIcon{Default: vault.Metadata.Image},
			Info: &actions.OptionInfo{
				Label: "Net APY",
				Value: fmt.Sprintf("%.2f%%", vault.DailyApys.NetApy*100),
			},
		})
	}

	return tokenOptions, tokenToVaultOptions, nil
}

func SupplyTokenToVaultOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	supplyTokenOptions, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: supplyTokenOptions},
		2: {Complex: supplyTokenToVaultOptions},
	}, nil
}

func GetMarketAndVaultOptions(chainId uint64) ([]actions.Option, []actions.Option, error) {
	vaults, err := reads.GetVaults(chainId)
	if err != nil {
		return nil, nil, err
	}

	markets, err := reads.GetMarkets(chainId)
	if err != nil {
		return nil, nil, err
	}

	marketOptions := make([]actions.Option, 0)
	for _, market := range markets {
		marketOptions = append(marketOptions, actions.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  &actions.OptionIcon{Default: market.Metadata.Icon},
		})
	}

	marketAndVaultOptions := make([]actions.Option, len(marketOptions))
	copy(marketAndVaultOptions, marketOptions)
	for _, vault := range vaults {
		marketAndVaultOptions = append(marketAndVaultOptions, actions.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  &actions.OptionIcon{Default: vault.Metadata.Image},
		})
	}

	return marketOptions, marketAndVaultOptions, nil
}
