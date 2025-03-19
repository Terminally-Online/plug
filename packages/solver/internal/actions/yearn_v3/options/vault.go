package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/yearn_v3/reads"
	"strings"
)

func GetVaultOptions(chainId uint64) ([]actions.Option, error) {
	vaults, err := reads.GetVaults(chainId)
	if err != nil {
		return nil, err
	}

	var options []actions.Option
	for _, vault := range vaults {
		options = append(options, actions.Option{
			Value: vault.Address,
			Name:  vault.DisplayName,
			Label: vault.FormattedSymbol,
			Info: &actions.OptionInfo{
				Label: "APR",
				Value: fmt.Sprintf("%.2f%%", vault.APR.ForwardAPR.NetAPR*100+vault.Extra.StakingRewardsAPR*100),
			},
			Icon: &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, strings.ToLower(vault.Token.Address))},
		})
	}

	return options, nil
}

func GetUnderlyingAssetToVaultOptions(chainId uint64) (map[string][]actions.Option, error) {
	vaults, err := reads.GetVaults(chainId)
	if err != nil {
		return nil, err
	}

	tokenMap := make(map[string][]actions.Option)
	for _, vault := range vaults {
		tokenAddress := fmt.Sprintf("%s:%d", vault.Token.Address, vault.Token.Decimals)
		tokenMap[tokenAddress] = append(tokenMap[tokenAddress], actions.Option{
			Value: vault.Address,
			Name:  vault.DisplayName,
			Label: vault.FormattedSymbol,
			Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, strings.ToLower(vault.Token.Address))},
		})
	}

	return tokenMap, nil
}

func UnderlyingAssetToVaultOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	underlyingAssetOptions, err := GetUnderlyingAssetOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	underlyingAssetToVaultOptions, err := GetUnderlyingAssetToVaultOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: underlyingAssetOptions},
		2: {Complex: underlyingAssetToVaultOptions},
	}, nil
}
