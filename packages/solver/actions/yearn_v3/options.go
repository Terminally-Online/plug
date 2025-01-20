package yearn_v3

import (
	"fmt"
	"solver/types"
	"strings"
)

type YearnV3OptionsProvider struct{}

func (p *YearnV3OptionsProvider) GetOptions(chainId int, action types.Action) (map[int]types.SchemaOptions, error) {
	underlyingAssetOptions, err := GetUnderlyingAssetOptions()
	if err != nil {
		return nil, err
	}
	underlyingAssetToVaultOptions, err := GetUnderlyingAssetToVaultOptions()
	if err != nil {
		return nil, err
	}
	availableStakingGuageOptions, err := GetAvailableStakingGuageOptions()
	if err != nil {
		return nil, err
	}
	vaultOptions, err := GetVaultOptions()
	if err != nil {
		return nil, err
	}

	switch action {
	case types.ActionDeposit:
		return map[int]types.SchemaOptions{
			1: {Simple: underlyingAssetOptions},
			2: {Complex: underlyingAssetToVaultOptions},
		}, nil
	case types.ActionWithdraw:
		return map[int]types.SchemaOptions{
			1: {Simple: underlyingAssetOptions},
			2: {Complex: underlyingAssetToVaultOptions},
		}, nil
	case types.ActionWithdrawMax:
		return map[int]types.SchemaOptions{
			0: {Simple: underlyingAssetOptions},
			1: {Complex: underlyingAssetToVaultOptions},
		}, nil
	case types.ActionStake:
		return map[int]types.SchemaOptions{
			1: {Simple: availableStakingGuageOptions},
		}, nil
	case types.ActionStakeMax:
		return map[int]types.SchemaOptions{
			0: {Simple: availableStakingGuageOptions},
		}, nil
	case types.ActionRedeem:
		return map[int]types.SchemaOptions{
			1: {Simple: availableStakingGuageOptions},
		}, nil
	case types.ActionRedeemMax:
		return map[int]types.SchemaOptions{
			0: {Simple: availableStakingGuageOptions},
		}, nil
	case types.ConstraintAPY:
		return map[int]types.SchemaOptions{
			0: {Simple: vaultOptions},
			1: {Simple: types.BaseThresholdFields},
		}, nil
	default:
		return nil, fmt.Errorf("unsupported action for options: %s", action)
	}
}

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
