package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/morpho/reads"
	"solver/internal/actions/morpho/types"
)

func GetCollateralTokenToMarketOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
	markets, err := reads.GetMarkets(chainId)
	if err != nil {
		return nil, nil, err
	}

	seenCollateral := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToMarketOptions := make(map[string][]actions.Option)

	for _, market := range markets {
		collateralAddress := fmt.Sprintf("%s:%d", market.CollateralAsset.Address, market.CollateralAsset.Decimals)

		if !seenCollateral[collateralAddress] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: market.CollateralAsset.Symbol,
				Name:  market.CollateralAsset.Name,
				Value: collateralAddress,
				Icon:  &actions.OptionIcon{Default: market.CollateralAsset.LogoURI},
			})
			seenCollateral[collateralAddress] = true
		}

		targetParams := types.MorphoTargetParams{
			TargetId:     market.UniqueKey,
			MarketParams: market.Params,
		}
		targetParamsString := targetParams.SerializeToCompactString()

		tokenToMarketOptions[collateralAddress] = append(
			tokenToMarketOptions[collateralAddress],
			actions.Option{
				Label: market.Metadata.Name,
				Name:  market.Metadata.Name,
				Value: targetParamsString,
				Icon:  &actions.OptionIcon{Default: market.Metadata.Icon},
				Info:  &actions.OptionInfo{Label: "Supply APY", Value: fmt.Sprintf("%.2f%%", market.State.DailySupplyApy*100)},
			},
		)
	}

	return tokenOptions, tokenToMarketOptions, nil
}

func CollateralTokenToMarketOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	collateralOptions, collateralToMarketOptions, err := GetCollateralTokenToMarketOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		1: {Simple: collateralOptions},
		2: {Complex: collateralToMarketOptions},
	}, nil
}

func SupplyAndCollateralTokenToMarketOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	supplyAndCollateralTokenOptions, supplyAndCollateralTokenToMarketOptions, err := GetSupplyAndCollateralTokenToMarketOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: supplyAndCollateralTokenOptions},
		2: {Complex: supplyAndCollateralTokenToMarketOptions},
	}, nil
}

func GetBorrowTokenToMarketOptions[T any](lookup *actions.SchemaLookup[T]) ([]actions.Option, map[string][]actions.Option, error) {
	markets, err := reads.GetMarkets(lookup.ChainId)
	if err != nil {
		return nil, nil, err
	}

	seenLoanAssets := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToMarketOptions := make(map[string][]actions.Option)

	for _, market := range markets {
		loanAssetAddress := fmt.Sprintf("%s:%d", market.LoanAsset.Address, market.LoanAsset.Decimals)
		if !seenLoanAssets[loanAssetAddress] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: market.LoanAsset.Symbol,
				Name:  market.LoanAsset.Name,
				Value: loanAssetAddress,
				Icon:  &actions.OptionIcon{Default: market.LoanAsset.LogoURI},
			})
			seenLoanAssets[loanAssetAddress] = true
		}

		targetParams := types.MorphoTargetParams{
			TargetId:     market.UniqueKey,
			MarketParams: market.Params,
		}
		targetParamsString := targetParams.SerializeToCompactString()

		tokenToMarketOptions[loanAssetAddress] = append(tokenToMarketOptions[loanAssetAddress], actions.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: targetParamsString,
			Icon:  &actions.OptionIcon{Default: market.Metadata.Icon},
			Info:  &actions.OptionInfo{Label: "Borrow APY", Value: fmt.Sprintf("%.2f%%", market.State.DailyBorrowApy*100)},
		})
	}

	return tokenOptions, tokenToMarketOptions, nil
}

func BorrowTokenToMarketOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	borrowOptions, borrowToMarketOptions, err := GetBorrowTokenToMarketOptions(lookup)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: borrowOptions},
		2: {Complex: borrowToMarketOptions},
	}, nil
}

func GetSupplyAndCollateralTokenToMarketOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
	vaults, err := reads.GetVaults(chainId)
	if err != nil {
		return nil, nil, err
	}

	markets, err := reads.GetMarkets(chainId)
	if err != nil {
		return nil, nil, err
	}

	seenToken := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToMarketAndVaultOptions := make(map[string][]actions.Option)

	for _, market := range markets {
		collateralAssetAddress := fmt.Sprintf("%s:%d", market.CollateralAsset.Address, market.CollateralAsset.Decimals)
		if !seenToken[collateralAssetAddress] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: market.CollateralAsset.Symbol,
				Name:  market.CollateralAsset.Name,
				Value: collateralAssetAddress,
				Icon:  &actions.OptionIcon{Default: market.CollateralAsset.LogoURI},
			})
			seenToken[collateralAssetAddress] = true
		}

		targetParams := types.MorphoTargetParams{
			TargetId:     market.UniqueKey,
			MarketParams: market.Params,
		}
		targetParamsString := targetParams.SerializeToCompactString()

		tokenToMarketAndVaultOptions[collateralAssetAddress] = append(
			tokenToMarketAndVaultOptions[collateralAssetAddress],
			actions.Option{
				Label: market.Metadata.Name,
				Name:  market.Metadata.Name,
				Value: targetParamsString,
				Icon:  &actions.OptionIcon{Default: market.Metadata.Icon},
				Info:  &actions.OptionInfo{Label: "Supply APY", Value: fmt.Sprintf("%.2f%%", market.State.DailySupplyApy*100)},
			},
		)
	}

	for _, vault := range vaults {
		vaultAssetAddress := fmt.Sprintf("%s:%d", vault.Asset.Address, vault.Asset.Decimals)
		if !seenToken[vault.Asset.Address] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: vault.Asset.Symbol,
				Name:  vault.Asset.Name,
				Value: vaultAssetAddress,
				Icon:  &actions.OptionIcon{Default: vault.Asset.LogoURI},
			})
			seenToken[vaultAssetAddress] = true
		}

		tokenToMarketAndVaultOptions[vaultAssetAddress] = append(tokenToMarketAndVaultOptions[vaultAssetAddress], actions.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: fmt.Sprintf("%s:%s", vault.Address, "{}"),
			Icon:  &actions.OptionIcon{Default: vault.Metadata.Image},
		})
	}

	return tokenOptions, tokenToMarketAndVaultOptions, nil
}
