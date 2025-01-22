package morpho

import (
	"fmt"
	"solver/actions"
)

type MorphoOptionsProvider struct{}

func (p *MorphoOptionsProvider) GetOptions(chainId int, action string) (map[int]actions.Options, error) {
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
	case ActionEarn:
		return map[int]actions.Options{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
		}, nil
	case ActionSupplyCollateral:
		return map[int]actions.Options{
			1: {Simple: collateralOptions},
			2: {Complex: collateralToMarketOptions},
		}, nil
	case ActionWithdraw:
		return map[int]actions.Options{
			1: {Simple: supplyAndCollateralTokenOptions},
			2: {Complex: supplyAndCollateralTokenToMarketOptions},
		}, nil
	case ActionWithdrawAll:
		return map[int]actions.Options{
			0: {Simple: supplyAndCollateralTokenOptions},
			1: {Complex: supplyAndCollateralTokenToMarketOptions},
		}, nil
	case ActionBorrow:
		return map[int]actions.Options{
			1: {Simple: borrowOptions},
			2: {Complex: borrowToMarketOptions},
		}, nil
	case ActionRepay:
		return map[int]actions.Options{
			1: {Simple: borrowOptions},
			2: {Complex: borrowToMarketOptions},
		}, nil
	case ActionRepayAll:
		return map[int]actions.Options{
			0: {Simple: borrowOptions},
			1: {Complex: borrowToMarketOptions},
		}, nil
	case actions.ConstraintHealthFactor:
		return map[int]actions.Options{
			0: {Simple: marketOptions},
			1: {Simple: actions.BaseThresholdFields},
		}, nil
	case ConstraintAPY:
		return map[int]actions.Options{
			0: {Simple: actions.BaseLendActionTypeFields},
			1: {Simple: marketAndVaultOptions},
			2: {Simple: actions.BaseThresholdFields},
		}, nil
	default:
		return nil, nil
	}
}

func GetSupplyTokenToVaultOptions() ([]actions.Option, map[string][]actions.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, nil, err
	}

	seenToken := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToVaultOptions := make(map[string][]actions.Option)

	for _, vault := range vaults {
		if !seenToken[vault.Asset.Address] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: vault.Asset.Symbol,
				Name:  vault.Asset.Name,
				Value: vault.Asset.Address,
				Icon:  vault.Asset.LogoURI,
			})
			seenToken[vault.Asset.Address] = true
		}

		tokenToVaultOptions[vault.Asset.Address] = append(tokenToVaultOptions[vault.Asset.Address], actions.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.Metadata.Image,
			Info:  fmt.Sprintf("%.2f%%", vault.DailyApys.NetApy*100),
		})
	}

	return tokenOptions, tokenToVaultOptions, nil
}

func GetMarketAndVaultOptions() ([]actions.Option, []actions.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, nil, err
	}

	markets, err := GetMarkets()
	if err != nil {
		return nil, nil, err
	}

	marketOptions := make([]actions.Option, 0)
	for _, market := range markets {
		marketOptions = append(marketOptions, actions.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
		})
	}

	marketAndVaultOptions := make([]actions.Option, len(marketOptions))
	copy(marketAndVaultOptions, marketOptions)
	for _, vault := range vaults {
		marketAndVaultOptions = append(marketAndVaultOptions, actions.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.Metadata.Image,
		})
	}

	return marketOptions, marketAndVaultOptions, nil
}

func GetCollateralTokenToMarketOptions() (
	[]actions.Option,
	map[string][]actions.Option,
	error,
) {
	markets, err := GetMarkets()
	if err != nil {
		return nil, nil, err
	}

	seenCollateral := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToMarketOptions := make(map[string][]actions.Option)

	for _, market := range markets {
		if !seenCollateral[market.CollateralAsset.Address] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: market.CollateralAsset.Symbol,
				Name:  market.CollateralAsset.Name,
				Value: market.CollateralAsset.Address,
				Icon:  market.CollateralAsset.LogoURI,
			})
			seenCollateral[market.CollateralAsset.Address] = true
		}

		tokenToMarketOptions[market.CollateralAsset.Address] = append(
			tokenToMarketOptions[market.CollateralAsset.Address],
			actions.Option{
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

func GetBorrowTokenToMarketOptions() ([]actions.Option, map[string][]actions.Option, error) {
	markets, err := GetMarkets()
	if err != nil {
		return nil, nil, err
	}

	seenLoanAssets := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToMarketOptions := make(map[string][]actions.Option)

	for _, market := range markets {
		if !seenLoanAssets[market.LoanAsset.Address] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: market.LoanAsset.Symbol,
				Name:  market.LoanAsset.Name,
				Value: market.LoanAsset.Address,
				Icon:  market.LoanAsset.LogoURI,
			})
			seenLoanAssets[market.LoanAsset.Address] = true
		}

		tokenToMarketOptions[market.LoanAsset.Address] = append(tokenToMarketOptions[market.LoanAsset.Address], actions.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
			Info:  fmt.Sprintf("%.2f%%", market.State.DailyBorrowApy*100),
		})
	}

	return tokenOptions, tokenToMarketOptions, nil
}

func GetSupplyAndCollateralTokenToMarketOptions() ([]actions.Option, map[string][]actions.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, nil, err
	}

	markets, err := GetMarkets()
	if err != nil {
		return nil, nil, err
	}

	seenToken := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToMarketAndVaultOptions := make(map[string][]actions.Option)

	for _, market := range markets {
		if !seenToken[market.CollateralAsset.Address] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: market.CollateralAsset.Symbol,
				Name:  market.CollateralAsset.Name,
				Value: market.CollateralAsset.Address,
				Icon:  market.CollateralAsset.LogoURI,
			})
			seenToken[market.CollateralAsset.Address] = true
		}

		tokenToMarketAndVaultOptions[market.CollateralAsset.Address] = append(tokenToMarketAndVaultOptions[market.LoanAsset.Address], actions.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
			Info:  fmt.Sprintf("%.2f%%", market.State.DailySupplyApy*100),
		})
	}

	for _, vault := range vaults {
		if !seenToken[vault.Asset.Address] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: vault.Asset.Symbol,
				Name:  vault.Asset.Name,
				Value: vault.Asset.Address,
				Icon:  vault.Asset.LogoURI,
			})
			seenToken[vault.Asset.Address] = true
		}

		tokenToMarketAndVaultOptions[vault.Asset.Address] = append(tokenToMarketAndVaultOptions[vault.Asset.Address], actions.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.Metadata.Image,
		})
	}

	return tokenOptions, tokenToMarketAndVaultOptions, nil
}
