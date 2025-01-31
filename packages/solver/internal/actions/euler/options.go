package euler

import (
	"fmt"
	"solver/internal/actions"
)

type EulerOptionsProvider struct{}

func (p *EulerOptionsProvider) GetOptions(chainId uint64, action string) (map[int]actions.Options, error) {
	supplyTokenOptions, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	marketOptions, marketAndVaultOptions, err := GetMarketAndVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	borrowTokenOptions, borrowTokenToVaultOptions, err := GetBorrowTokenToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}

	switch action {
	case ActionSupply:
		return map[int]actions.Options{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
		}, nil
	case ActionWithdraw:
		return map[int]actions.Options{
			1: {Simple: supplyTokenOptions},
			2: {Complex: supplyTokenToVaultOptions},
		}, nil
	case ActionWithdrawAll:
		return map[int]actions.Options{
			0: {Simple: supplyTokenOptions},
			1: {Complex: supplyTokenToVaultOptions},
		}, nil
	case ActionBorrow:
		return map[int]actions.Options{
			1: {Simple: borrowTokenOptions},
			2: {Complex: borrowTokenToVaultOptions},
		}, nil
	case ActionRepay:
		return map[int]actions.Options{
			1: {Simple: borrowTokenOptions},
			2: {Complex: borrowTokenToVaultOptions},
		}, nil
	case ActionRepayWithShares:
		return map[int]actions.Options{
			1: {Simple: borrowTokenOptions},
			2: {Complex: borrowTokenToVaultOptions},
		}, nil
	case ConstraintHealthFactor:
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
	case ConstraintTimeToLiq:
		return map[int]actions.Options{
			0: {Simple: marketOptions},
			1: {Simple: actions.BaseThresholdFields},
		}, nil
	default:
		return nil, nil
	}
}

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
				Icon:  vault.Asset.LogoURI,
			})
			seenToken[assetAddress] = true
		}

		tokenToVaultOptions[assetAddress] = append(tokenToVaultOptions[assetAddress], actions.Option{
			Label: vault.Name,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.LogoURI,
			Info:  fmt.Sprintf("%.2f%%", vault.APY*100),
		})
	}

	return tokenOptions, tokenToVaultOptions, nil
}

func GetMarketAndVaultOptions(chainId uint64) ([]actions.Option, []actions.Option, error) {
	vaults, err := GetVaults(chainId)
	if err != nil {
		return nil, nil, err
	}

	marketOptions := make([]actions.Option, 0)
	marketAndVaultOptions := make([]actions.Option, 0)

	for _, vault := range vaults {
		option := actions.Option{
			Label: vault.Name,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.LogoURI,
		}
		marketOptions = append(marketOptions, option)
		marketAndVaultOptions = append(marketAndVaultOptions, option)
	}

	return marketOptions, marketAndVaultOptions, nil
}

func GetBorrowTokenToVaultOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
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
			Label: vault.Name,
			Name:  vault.Name,
			Value: vault.Address,
			Icon:  vault.LogoURI,
			Info:  fmt.Sprintf("%.2f%%", vault.BorrowAPY*100),
		})
	}

	return tokenOptions, tokenToVaultOptions, nil
}