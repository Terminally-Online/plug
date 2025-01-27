package morpho

import (
	"fmt"
	"solver/internal/actions"
)

type MorphoOptionsProvider struct{}

func (p *MorphoOptionsProvider) GetOptions(chainId uint64, action string) (map[int]actions.Options, error) {
	supplyTokenOptions, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	marketOptions, marketAndVaultOptions, err := GetMarketAndVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	collateralOptions, collateralToMarketOptions, err := GetCollateralTokenToMarketOptions(chainId)
	if err != nil {
		return nil, err
	}
	borrowOptions, borrowToMarketOptions, err := GetBorrowTokenToMarketOptions(chainId)
	if err != nil {
		return nil, err
	}
	supplyAndCollateralTokenOptions, supplyAndCollateralTokenToMarketOptions, err := GetSupplyAndCollateralTokenToMarketOptions(chainId)
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

func GetSupplyTokenToVaultOptions(chainId int) ([]actions.Option, map[string][]actions.Option, error) {
	vaults, err := GetVaults(chainId)
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
				Icon:  vault.Asset.LogoURI,
			})
			seenToken[assetAddress] = true
		}

		tokenToVaultOptions[assetAddress] = append(tokenToVaultOptions[assetAddress], actions.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: assetAddress,
			Icon:  vault.Metadata.Image,
			Info:  fmt.Sprintf("%.2f%%", vault.DailyApys.NetApy*100),
		})
	}

	return tokenOptions, tokenToVaultOptions, nil
}

func GetMarketAndVaultOptions(chainId int) ([]actions.Option, []actions.Option, error) {
	vaults, err := GetVaults(chainId)
	if err != nil {
		return nil, nil, err
	}

	markets, err := GetMarkets(chainId)
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

func GetCollateralTokenToMarketOptions(chainId int) ([]actions.Option, map[string][]actions.Option, error) {
	markets, err := GetMarkets(chainId)
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
				Icon:  market.CollateralAsset.LogoURI,
			})
			seenCollateral[collateralAddress] = true
		}

		tokenToMarketOptions[collateralAddress] = append(
			tokenToMarketOptions[collateralAddress],
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

func GetBorrowTokenToMarketOptions(chainId int) ([]actions.Option, map[string][]actions.Option, error) {
	markets, err := GetMarkets(chainId)
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
				Icon:  market.LoanAsset.LogoURI,
			})
			seenLoanAssets[loanAssetAddress] = true
		}

		tokenToMarketOptions[loanAssetAddress] = append(tokenToMarketOptions[loanAssetAddress], actions.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  market.Metadata.Icon,
			Info:  fmt.Sprintf("%.2f%%", market.State.DailyBorrowApy*100),
		})
	}

	return tokenOptions, tokenToMarketOptions, nil
}

func GetSupplyAndCollateralTokenToMarketOptions(chainId int) ([]actions.Option, map[string][]actions.Option, error) {
	vaults, err := GetVaults(chainId)
	if err != nil {
		return nil, nil, err
	}

	markets, err := GetMarkets(chainId)
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
				Icon:  market.CollateralAsset.LogoURI,
			})
			seenToken[collateralAssetAddress] = true
		}

		tokenToMarketAndVaultOptions[collateralAssetAddress] = append(
			tokenToMarketAndVaultOptions[collateralAssetAddress],
			actions.Option{
				Label: market.Metadata.Name,
				Name:  market.Metadata.Name,
				Value: market.UniqueKey,
				Icon:  market.Metadata.Icon,
				Info:  fmt.Sprintf("%.2f%%", market.State.DailySupplyApy*100),
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
				Icon:  vault.Asset.LogoURI,
			})
			seenToken[vaultAssetAddress] = true
		}

		tokenToMarketAndVaultOptions[vaultAssetAddress] = append(tokenToMarketAndVaultOptions[vaultAssetAddress], actions.Option{
			Label: vault.Symbol,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.Metadata.Image,
		})
	}

	return tokenOptions, tokenToMarketAndVaultOptions, nil
}
