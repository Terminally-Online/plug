package morpho

import (
	"fmt"
	"solver/types"
)

type MorphoOptionsProvider struct{}

func (p *MorphoOptionsProvider) GetOptions(chainId int, action types.Action) (map[int]types.SchemaOptions, error) {
	supplyTokenOptions, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions()
	if err != nil {
		return nil, err
	}
	marketOptions, marketAndVaultOptions, err := GetMarketAndVaultOptions()
	if err != nil {
		return nil, err
	}
	collateralOptions, collateralToMarketOptions, err := GetCollateralTokenToMarketOptions()
	if err != nil {
		return nil, err
	}
	borrowOptions, borrowToMarketOptions, err := GetBorrowTokenToMarketOptions()
	if err != nil {
		return nil, err
	}
	supplyAndCollateralTokenOptions, supplyAndCollateralTokenToMarketOptions, err := GetSupplyAndCollateralTokenToMarketOptions()
	if err != nil {
		return nil, err
	}

	switch action {
	case types.Action(ActionEarn):
		return map[int]types.SchemaOptions{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
		}, nil
	case types.Action(ActionSupplyCollateral):
		return map[int]types.SchemaOptions{
			1: {Simple: collateralOptions},
			2: {Complex: collateralToMarketOptions},
		}, nil
	case types.Action(ActionWithdraw):
		return map[int]types.SchemaOptions{
			1: {Simple: supplyAndCollateralTokenOptions},
			2: {Complex: supplyAndCollateralTokenToMarketOptions},
		}, nil
	case types.Action(ActionWithdrawAll):
		return map[int]types.SchemaOptions{
			0: {Simple: supplyAndCollateralTokenOptions},
			1: {Complex: supplyAndCollateralTokenToMarketOptions},
		}, nil
	case types.Action(ActionBorrow):
		return map[int]types.SchemaOptions{
			1: {Simple: borrowOptions},
			2: {Complex: borrowToMarketOptions},
		}, nil
	case types.Action(ActionRepay):
		return map[int]types.SchemaOptions{
			1: {Simple: borrowOptions},
			2: {Complex: borrowToMarketOptions},
		}, nil
	case types.Action(ActionRepayAll):
		return map[int]types.SchemaOptions{
			0: {Simple: borrowOptions},
			1: {Complex: borrowToMarketOptions},
		}, nil
	case types.ConstraintHealthFactor:
		return map[int]types.SchemaOptions{
			0: {Simple: marketOptions},
			1: {Simple: types.BaseThresholdFields},
		}, nil
	case types.Action(ConstraintAPY):
		return map[int]types.SchemaOptions{
			0: {Simple: types.BaseLendActionTypeFields},
			1: {Simple: marketAndVaultOptions},
			2: {Simple: types.BaseThresholdFields},
		}, nil
	default:
		return nil, nil
	}
}

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
				Value: fmt.Sprintf("%s:%d", vault.Asset.Address, vault.Asset.Decimals),
				Icon:  vault.Asset.LogoURI,
			})
			seenToken[vault.Asset.Address] = true
		}

		tokenToVaultOptions[vault.Asset.Address] = append(tokenToVaultOptions[vault.Asset.Address], types.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: fmt.Sprintf("%s:%d", vault.Address, vault.Asset.Decimals),
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
				Value: fmt.Sprintf("%s:%d", market.CollateralAsset.Address, market.CollateralAsset.Decimals),
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
				Value: fmt.Sprintf("%s:%d", market.LoanAsset.Address, market.LoanAsset.Decimals),
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
				Value: fmt.Sprintf("%s:%d", market.CollateralAsset.Address, market.CollateralAsset.Decimals),
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
				Value: fmt.Sprintf("%s:%d", vault.Asset.Address, vault.Asset.Decimals),
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
