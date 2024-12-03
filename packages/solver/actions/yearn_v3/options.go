package yearn_v3

import (
	"fmt"
	"solver/types"
	"strings"
)

func GetUnderlyingAssetOptions() ([]types.Option, error) {
	tokens, err := GenerateTokenList()
	if err != nil {
		return nil, err
	}

	tokenDetails := make(map[string]Token)
	for _, token := range tokens {
		tokenDetails[strings.ToLower(token.Address)] = token
	}

	vaults, err := GetVaults()
	if err != nil {
		return nil, err
	}

	tokenMap := make(map[string]types.Option)

	for _, vault := range vaults {
		lowerAddr := strings.ToLower(vault.Token.Address)

		if _, exists := tokenMap[lowerAddr]; !exists {
			if token, ok := tokenDetails[lowerAddr]; ok {
				tokenMap[lowerAddr] = types.Option{
					Value: token.Address,
					Name:  token.Name,
					Label: token.Symbol,
					Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/1/%s", vault.Token.Address),
				}
			}
		}
	}

	options := make([]types.Option, 0, len(tokenMap))
	for _, option := range tokenMap {
		options = append(options, option)
	}

	return options, nil
}

func GetUnderlyingAssetToVaultOptions() (map[string][]types.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, err
	}

	tokenMap := make(map[string][]types.Option)
	for _, vault := range vaults {
		tokenMap[vault.Token.Address] = append(tokenMap[vault.Token.Address], types.Option{
			Value: vault.Address,
			Name:  vault.DisplayName,
			Label: vault.FormattedSymbol,
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/1/%s", vault.Address),
		})
	}

	return tokenMap, nil
}

func GetAvailableStakingGuageOptions() ([]types.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, err
	}

	var options []types.Option
	for _, vault := range vaults {
		if !vault.Staking.Available || vault.Staking.Address == "" {
			continue
		}

		options = append(options, types.Option{
			Value: vault.Address,
			Name:  vault.DisplayName,
			Label: vault.FormattedSymbol,
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/1/%s", vault.Address),
		})
	}

	return options, nil
}

func GetVaultOptions() ([]types.Option, error) {
	vaults, err := GetVaults()
	if err != nil {
		return nil, err
	}

	var options []types.Option
	for _, vault := range vaults {
		options = append(options, types.Option{
			Value: vault.Address,
			Name:  vault.DisplayName,
			Label: vault.FormattedSymbol,
			Info:  fmt.Sprintf("%.2f%%", vault.APR.ForwardAPR.NetAPR*100+vault.Extra.StakingRewardsAPR*100),
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/1/%s", vault.Address),
		})
	}

	return options, nil
}
