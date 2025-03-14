package yearn_v3

import (
	"fmt"
	"solver/internal/actions"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type YearnV3OptionsProvider struct{}

func UnderlyingAssetToVaultOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	underlyingAssetOptions, err := GetUnderlyingAssetOptions(chainId)
	if err != nil {
		return nil, err
	}
	underlyingAssetToVaultOptions, err := GetUnderlyingAssetToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		1: {Simple: underlyingAssetOptions},
		2: {Complex: underlyingAssetToVaultOptions},
	}, nil
}

func AvailableStakingGaugeOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	availableStakingGaugeOptions, err := GetAvailableStakingGaugeOptions(chainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{1: {Simple: availableStakingGaugeOptions}}, nil
}

func APYOptions(chainId uint64, _ common.Address, _ map[int]string, _ string) (map[int]actions.Options, error) {
	vaultOptions, err := GetVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{
		0: {Simple: vaultOptions},
		1: {Simple: actions.BaseThresholdFields},
	}, nil
}

func (p *YearnV3OptionsProvider) GetOptions(chainId uint64, _ common.Address, _ map[int]string, action string) (map[int]actions.Options, error) {
	switch action {
	case actions.ActionDeposit, actions.ActionWithdraw:
		return UnderlyingAssetToVaultOptions(chainId, common.Address{}, nil, action)
	case actions.ActionStake, actions.ActionRedeem:
		return AvailableStakingGaugeOptions(chainId, common.Address{}, nil, action)
	case actions.ConstraintAPY:
		return APYOptions(chainId, common.Address{}, nil, action)
	default:
		return nil, fmt.Errorf("unsupported action for options: %s", action)
	}
}

func GetUnderlyingAssetOptions(chainId uint64) ([]actions.Option, error) {
	tokens, err := GenerateTokenList()
	if err != nil {
		return nil, err
	}

	tokenDetails := make(map[string]Token)
	for _, token := range tokens {
		tokenDetails[strings.ToLower(token.Address)] = token
	}
	vaults, err := GetVaults(chainId)
	if err != nil {
		return nil, err
	}

	tokenMap := make(map[string]actions.Option)
	for _, vault := range vaults {
		lowerAddr := strings.ToLower(vault.Token.Address)

		if _, exists := tokenMap[lowerAddr]; !exists {
			if token, ok := tokenDetails[lowerAddr]; ok {
				tokenMap[lowerAddr] = actions.Option{
					Value: fmt.Sprintf("%s:%d", token.Address, token.Decimals),
					Name:  token.Name,
					Label: token.Symbol,
					Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, lowerAddr)},
				}
			}
		}
	}

	options := make([]actions.Option, 0, len(tokenMap))
	for _, option := range tokenMap {
		options = append(options, option)
	}

	return options, nil
}

func GetUnderlyingAssetToVaultOptions(chainId uint64) (map[string][]actions.Option, error) {
	vaults, err := GetVaults(chainId)
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

func GetAvailableStakingGaugeOptions(chainId uint64) ([]actions.Option, error) {
	vaults, err := GetVaults(chainId)
	if err != nil {
		return nil, err
	}

	var options []actions.Option
	for _, vault := range vaults {
		if !vault.Staking.Available || vault.Staking.Address == "" {
			continue
		}

		options = append(options, actions.Option{
			Value: fmt.Sprintf("%s:%d", vault.Address, vault.Decimals),
			Name:  vault.DisplayName,
			Label: vault.FormattedSymbol,
			Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, strings.ToLower(vault.Token.Address))},
		})
	}

	return options, nil
}

func GetVaultOptions(chainId uint64) ([]actions.Option, error) {
	vaults, err := GetVaults(chainId)
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
