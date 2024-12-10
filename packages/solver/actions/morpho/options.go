package morpho

import (
	"fmt"
	"solver/types"
)

func GetMarketOptions() ([]types.Option, error) {
	markets, err := GetMarkets()
	if err != nil {
		return nil, err
	}

	options := make([]types.Option, 0)
	for _, market := range markets {
		options = append(options, types.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
		})
	}

	return options, nil
}

func GetCollateralTokenToMarketOptions() ([]types.Option, map[string][]types.Option, error) {
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
				Name:  market.CollateralAsset.Symbol,
				Value: market.CollateralAsset.Address,
			})
			seenCollateral[market.CollateralAsset.Address] = true
		}

		tokenToMarketOptions[market.CollateralAsset.Address] = append(tokenToMarketOptions[market.CollateralAsset.Address], types.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
			Info:  fmt.Sprintf("%.2f%%", market.State.DailySupplyApy*100),
		})
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
				Name:  market.LoanAsset.Symbol,
				Value: market.LoanAsset.Address,
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
