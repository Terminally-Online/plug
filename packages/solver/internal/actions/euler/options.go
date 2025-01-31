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

	switch action {
		case ActionSupply:
			return map[int]actions.Options{
				1: {Simple: supplyTokenOptions},
				2: {Complex: supplyTokenToVaultOptions},
			}, nil
	default:
		return nil, nil
	}
}

func GetSupplyTokenToVaultOptions(chainId uint64) ([]actions.Option, map[string][]actions.Option, error) {
	vaults, err := GetVerifiedVaults(chainId)
	if err != nil {
		return nil, nil, err
	}

	seenToken := make(map[string]bool)
	tokenOptions := make([]actions.Option, 0)
	tokenToVaultOptions := make(map[string][]actions.Option)

	for _, vault := range vaults {
		tokenAddress := fmt.Sprintf("%s:%d", vault.Asset, vault.AssetDecimals)
		if !seenToken[tokenAddress] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: vault.AssetSymbol,
				Name:  vault.AssetName,
				Value: tokenAddress,
				Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", 1, vault.Asset.String()),
			})
			seenToken[tokenAddress] = true
		}

		tokenToVaultOptions[tokenAddress] = append(tokenToVaultOptions[tokenAddress], actions.Option{
			Label: vault.VaultSymbol,
			Name:  vault.VaultName,
			Value: vault.Vault.String(),
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", 1, vault.Asset.String()),
			Info:  fmt.Sprintf("%.2f%%", 0.0022), // TODO MASON -- This should be the APY I guess? Will need to calc in the GetVaults() function
		})
	}

	return tokenOptions, tokenToVaultOptions, nil
}