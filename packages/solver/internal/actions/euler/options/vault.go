package options

import (
	"fmt"
	"solver/bindings/euler_vault_lens"
	"solver/internal/actions"
	"solver/internal/actions/euler/reads"
	"solver/internal/utils"
	"strings"
)

func GetSupplyTokenToVaultOptions(chainId uint64, vaults []euler_vault_lens.VaultInfoFull) (tokenOptions []actions.Option, vaultOptions []actions.Option, tokenToVaultOptions map[string][]actions.Option, err error) {
	tokenOptions = make([]actions.Option, 0)
	vaultOptions = make([]actions.Option, 0)
	tokenToVaultOptions = make(map[string][]actions.Option)
	seenToken := make(map[string]bool)
	for _, vault := range vaults {
		tokenAddress := fmt.Sprintf("%s:%d", vault.Asset, vault.AssetDecimals)
		if !seenToken[tokenAddress] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: vault.AssetSymbol,
				Name:  vault.AssetName,
				Value: tokenAddress,
				Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
			})
			seenToken[tokenAddress] = true
		}

		var supplyApy string
		if len(vault.IrmInfo.InterestRateInfo) > 0 {
			supplyApyFloat := utils.UintToFloat(vault.IrmInfo.InterestRateInfo[0].SupplyAPY, 25)
			supplyApy = fmt.Sprintf("%.2f%%", supplyApyFloat)
		} else {
			supplyApy = "0.0%"
		}

		vaultOption := actions.Option{
			Label: vault.VaultSymbol,
			Name:  vault.VaultName,
			Value: vault.Vault.String(),
			Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
			Info:  &actions.OptionInfo{Label: "Supply APY", Value: supplyApy},
		}
		vaultOptions = append(vaultOptions, vaultOption)
		tokenToVaultOptions[tokenAddress] = append(tokenToVaultOptions[tokenAddress], vaultOption)
	}

	return tokenOptions, vaultOptions, tokenToVaultOptions, nil
}

func SupplyTokenToVaultOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	vaults, err := reads.GetVerifiedVaults(lookup.ChainId)
	if err != nil {
		return nil, err
	}

	supplyTokenOptions, _, supplyTokenToVaultOptions, err := GetSupplyTokenToVaultOptions(lookup.ChainId, vaults)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: supplyTokenOptions},
		2: {Complex: supplyTokenToVaultOptions},
	}, nil
}

func GetBorrowTokenToVaultOptions(chainId uint64, vaults []euler_vault_lens.VaultInfoFull) (tokenOptions []actions.Option, vaultOptions []actions.Option, tokenToVaultOptions map[string][]actions.Option, err error) {
	tokenOptions = make([]actions.Option, 0)
	vaultOptions = make([]actions.Option, 0)
	tokenToVaultOptions = make(map[string][]actions.Option)
	seenToken := make(map[string]bool)
	for _, vault := range vaults {
		tokenAddress := fmt.Sprintf("%s:%d", vault.Asset, vault.AssetDecimals)
		if !seenToken[tokenAddress] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: vault.AssetSymbol,
				Name:  vault.AssetName,
				Value: tokenAddress,
				Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
			})
			seenToken[tokenAddress] = true
		}

		var borrowApy string
		if len(vault.IrmInfo.InterestRateInfo) > 0 {
			borrowApyFloat := utils.UintToFloat(vault.IrmInfo.InterestRateInfo[0].BorrowAPY, 25)
			borrowApy = fmt.Sprintf("%.2f%%", borrowApyFloat)
		} else {
			borrowApy = "0.0%"
		}

		vaultOption := actions.Option{
			Label: vault.VaultSymbol,
			Name:  vault.VaultName,
			Value: vault.Vault.String(),
			Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s?h=60&w=60", chainId, strings.ToLower(vault.Asset.String()))},
			Info:  &actions.OptionInfo{Label: "Borrow APY", Value: borrowApy},
		}
		vaultOptions = append(vaultOptions, vaultOption)
		tokenToVaultOptions[tokenAddress] = append(tokenToVaultOptions[tokenAddress], vaultOption)
	}

	return tokenOptions, vaultOptions, tokenToVaultOptions, nil
}

func BorrowTokenToVaultOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	vaults, err := reads.GetVerifiedVaults(lookup.ChainId)
	if err != nil {
		return nil, err
	}

	borrowTokenOptions, _, borrowTokenToVaultOptions, err := GetBorrowTokenToVaultOptions(lookup.ChainId, vaults)
	if err != nil {
		return nil, err
	}
	addressPositions, err := GetAddressPositions(lookup.ChainId, lookup.From)
	if err != nil {
		fmt.Printf("error getting address positions: %v\n", err)
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: borrowTokenOptions},
		2: {Complex: borrowTokenToVaultOptions},
		3: {Simple: addressPositions},
	}, nil
}
