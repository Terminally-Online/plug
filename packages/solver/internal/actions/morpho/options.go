package morpho

import (
	"fmt"
	"solver/internal/actions"

	"github.com/ethereum/go-ethereum/common"
)

func GetSupplyTokenToVaultOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
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

func SupplyTokenToVaultOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	supplyTokenOptions, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: supplyTokenOptions},
		2: {Complex: supplyTokenToVaultOptions},
	}, nil
}

func GetMarketAndVaultOptions(chainId uint64) ([]actions.Option, []actions.Option, error) {
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

func HealthFactorOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	marketOptions, _, err := GetMarketAndVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: marketOptions},
		1: {Simple: actions.BaseThresholdFields},
	}, nil
}

func APYOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	_, marketAndVaultOptions, err := GetMarketAndVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: actions.BaseLendActionTypeFields},
		1: {Simple: marketAndVaultOptions},
		2: {Simple: actions.BaseThresholdFields},
	}, nil
}

func GetCollateralTokenToMarketOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
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
				Icon:  &actions.OptionIcon{Default: market.CollateralAsset.LogoURI},
			})
			seenCollateral[collateralAddress] = true
		}

		tokenToMarketOptions[collateralAddress] = append(
			tokenToMarketOptions[collateralAddress],
			actions.Option{
				Label: market.Metadata.Name,
				Name:  market.Metadata.Name,
				Value: market.UniqueKey,
				Icon:  &actions.OptionIcon{Default: market.Metadata.Icon},
				Info:  &actions.OptionInfo{Label: "Supply APY", Value: fmt.Sprintf("%.2f%%", market.State.DailySupplyApy*100)},
			},
		)
	}

	return tokenOptions, tokenToMarketOptions, nil
}

func CollateralTokenToMarketOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	collateralOptions, collateralToMarketOptions, err := GetCollateralTokenToMarketOptions(chainId)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		1: {Simple: collateralOptions},
		2: {Complex: collateralToMarketOptions},
	}, nil
}

func SupplyAndCollateralTokenToMarketOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	supplyAndCollateralTokenOptions, supplyAndCollateralTokenToMarketOptions, err := GetSupplyAndCollateralTokenToMarketOptions(chainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: supplyAndCollateralTokenOptions},
		2: {Complex: supplyAndCollateralTokenToMarketOptions},
	}, nil
}

func GetBorrowTokenToMarketOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
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
				Icon:  &actions.OptionIcon{Default: market.LoanAsset.LogoURI},
			})
			seenLoanAssets[loanAssetAddress] = true
		}

		tokenToMarketOptions[loanAssetAddress] = append(tokenToMarketOptions[loanAssetAddress], actions.Option{
			Label: market.Metadata.Name,
			Name:  market.Metadata.Name,
			Value: market.UniqueKey,
			Icon:  &actions.OptionIcon{Default: market.Metadata.Icon},
			Info:  &actions.OptionInfo{Label: "Borrow APY", Value: fmt.Sprintf("%.2f%%", market.State.DailyBorrowApy*100)},
		})
	}

	return tokenOptions, tokenToMarketOptions, nil
}

func BorrowTokenToMarketOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	borrowOptions, borrowToMarketOptions, err := GetBorrowTokenToMarketOptions(chainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: borrowOptions},
		2: {Complex: borrowToMarketOptions},
	}, nil
}

func GetSupplyAndCollateralTokenToMarketOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
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
				Icon:  &actions.OptionIcon{Default: market.CollateralAsset.LogoURI},
			})
			seenToken[collateralAssetAddress] = true
		}

		tokenToMarketAndVaultOptions[collateralAssetAddress] = append(
			tokenToMarketAndVaultOptions[collateralAssetAddress],
			actions.Option{
				Label: market.Metadata.Name,
				Name:  market.Metadata.Name,
				Value: market.UniqueKey,
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
			Value: vault.Address,
			Icon:  &actions.OptionIcon{Default: vault.Metadata.Image},
		})
	}

	return tokenOptions, tokenToMarketAndVaultOptions, nil
}
