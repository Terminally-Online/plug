package morpho

import (
	"fmt"
	"solver/types"
)

func GetSupplyTokenToVaultOptions() ([]types.Option, map[string][]types.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, nil, err
	}

	seenToken := make(map[string]bool)
	tokenOptions := make([]types.Option, 0)
	tokenToVaultOptions := make(map[string][]types.Option)

	for _, vault := range vaults {
		if !seenToken[vault.Asset.Address] {
			tokenOptions = append(tokenOptions, types.Option{
				Label: vault.Asset.Symbol,
				Name:  vault.Asset.Name,
				Value: vault.Asset.Address,
				Icon:  vault.Asset.LogoURI,
			})
			seenToken[vault.Asset.Address] = true
		}

		tokenToVaultOptions[vault.Asset.Address] = append(tokenToVaultOptions[vault.Asset.Address], types.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.Metadata.Image,
			Info:  fmt.Sprintf("%.2f%%", vault.DailyApys.NetApy*100),
		})
	}

	return tokenOptions, tokenToVaultOptions, nil
}

func GetMarketAndVaultOptions() ([]types.Option, []types.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, nil, err
	}

	markets, err := GetMarkets()
	if err != nil {
		return nil, nil, err
	}

	marketOptions := make([]types.Option, 0)
	for _, market := range markets {
		marketOptions = append(marketOptions, types.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
		})
	}

	marketAndVaultOptions := make([]types.Option, len(marketOptions))
	copy(marketAndVaultOptions, marketOptions)
	for _, vault := range vaults {
		marketAndVaultOptions = append(marketAndVaultOptions, types.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.Metadata.Image,
		})
	}

	return marketOptions, marketAndVaultOptions, nil
}

func GetCollateralTokenToMarketOptions() (
	[]types.Option,
	map[string][]types.Option,
	error,
) {
	markets, err := GetMarkets()
	if err != nil {
		return nil, nil, err
	}

	seenCollateral := make(map[string]bool)
	tokenOptions := make([]types.Option, 0)
	tokenToMarketOptions := make(map[string][]types.Option)

	for _, market := range markets {
		if !seenCollateral[market.CollateralAsset.Address] {
			tokenOptions = append(tokenOptions, types.Option{
				Label: market.CollateralAsset.Symbol,
				Name:  market.CollateralAsset.Name,
				Value: market.CollateralAsset.Address,
				Icon:  market.CollateralAsset.LogoURI,
			})
			seenCollateral[market.CollateralAsset.Address] = true
		}

		tokenToMarketOptions[market.CollateralAsset.Address] = append(
			tokenToMarketOptions[market.CollateralAsset.Address],
			types.Option{
				Label: market.Metadata.Name,
				Name:  market.Metadata.Name,
				Value: market.UniqueKey,
				Icon:  market.Metadata.Icon,
				Info:  fmt.Sprintf("%.2f%%", market.State.DailySupplyApy*100),
			},
		)
	}

	return tokenOptions, tokenToMarketOptions, nil
}

func GetBorrowTokenToMarketOptions() ([]types.Option, map[string][]types.Option, error) {
	markets, err := GetMarkets()
	if err != nil {
		return nil, nil, err
	}

	seenLoanAssets := make(map[string]bool)
	tokenOptions := make([]types.Option, 0)
	tokenToMarketOptions := make(map[string][]types.Option)

	for _, market := range markets {
		if !seenLoanAssets[market.LoanAsset.Address] {
			tokenOptions = append(tokenOptions, types.Option{
				Label: market.LoanAsset.Symbol,
				Name:  market.LoanAsset.Name,
				Value: market.LoanAsset.Address,
				Icon:  market.LoanAsset.LogoURI,
			})
			seenLoanAssets[market.LoanAsset.Address] = true
		}

		tokenToMarketOptions[market.LoanAsset.Address] = append(tokenToMarketOptions[market.LoanAsset.Address], types.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
			Info:  fmt.Sprintf("%.2f%%", market.State.DailyBorrowApy*100),
		})
	}

	return tokenOptions, tokenToMarketOptions, nil
}

func GetSupplyAndCollateralTokenToMarketOptions() ([]types.Option, map[string][]types.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, nil, err
	}

	markets, err := GetMarkets()
	if err != nil {
		return nil, nil, err
	}

	seenToken := make(map[string]bool)
	tokenOptions := make([]types.Option, 0)
	tokenToMarketAndVaultOptions := make(map[string][]types.Option)

	for _, market := range markets {
		if !seenToken[market.CollateralAsset.Address] {
			tokenOptions = append(tokenOptions, types.Option{
				Label: market.CollateralAsset.Symbol,
				Name:  market.CollateralAsset.Name,
				Value: market.CollateralAsset.Address,
				Icon:  market.CollateralAsset.LogoURI,
			})
			seenToken[market.CollateralAsset.Address] = true
		}

		tokenToMarketAndVaultOptions[market.CollateralAsset.Address] = append(tokenToMarketAndVaultOptions[market.LoanAsset.Address], types.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
			Info:  fmt.Sprintf("%.2f%%", market.State.DailySupplyApy*100),
		})
	}

	for _, vault := range vaults {
		if !seenToken[vault.Asset.Address] {
			tokenOptions = append(tokenOptions, types.Option{
				Label: vault.Asset.Symbol,
				Name:  vault.Asset.Name,
				Value: vault.Asset.Address,
				Icon:  vault.Asset.LogoURI,
			})
			seenToken[vault.Asset.Address] = true
		}

		tokenToMarketAndVaultOptions[vault.Asset.Address] = append(tokenToMarketAndVaultOptions[vault.Asset.Address], types.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.Metadata.Image,
		})
	}

	return tokenOptions, tokenToMarketAndVaultOptions, nil
}
