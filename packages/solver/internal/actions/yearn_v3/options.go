package yearn_v3

import (
	"fmt"
	"solver/internal/actions"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type YearnV3OptionsProvider struct{}

func (p *YearnV3OptionsProvider) GetOptions(chainId uint64, _ common.Address, action string) (map[int]actions.Options, error) {
	underlyingAssetOptions, err := GetUnderlyingAssetOptions(chainId)
	if err != nil {
		return nil, err
	}
	underlyingAssetToVaultOptions, err := GetUnderlyingAssetToVaultOptions(chainId)
	if err != nil {
		return nil, err
	}
	availableStakingGaugeOptions, err := GetAvailableStakingGaugeOptions(chainId)
	if err != nil {
		return nil, err
	}
	vaultOptions, err := GetVaultOptions(chainId)
	if err != nil {
		return nil, err
	}

	switch action {
	case actions.ActionDeposit:
		return map[int]actions.Options{
			1: {Simple: underlyingAssetOptions},
			2: {Complex: underlyingAssetToVaultOptions},
		}, nil
	case actions.ActionWithdraw:
		return map[int]actions.Options{
			1: {Simple: underlyingAssetOptions},
			2: {Complex: underlyingAssetToVaultOptions},
		}, nil
	case actions.ActionWithdrawMax:
		return map[int]actions.Options{
			0: {Simple: underlyingAssetOptions},
			1: {Complex: underlyingAssetToVaultOptions},
		}, nil
	case actions.ActionStake:
		return map[int]actions.Options{
			1: {Simple: availableStakingGaugeOptions},
		}, nil
	case actions.ActionStakeMax:
		return map[int]actions.Options{
			0: {Simple: availableStakingGaugeOptions},
		}, nil
	case actions.ActionRedeem:
		return map[int]actions.Options{
			1: {Simple: availableStakingGaugeOptions},
		}, nil
	case actions.ActionRedeemMax:
		return map[int]actions.Options{
			0: {Simple: availableStakingGaugeOptions},
		}, nil
	case actions.ConstraintAPY:
		return map[int]actions.Options{
			0: {Simple: vaultOptions},
			1: {Simple: actions.BaseThresholdFields},
		}, nil
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
					Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/1/%s", vault.Token.Address),
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
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, vault.Token.Address),
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
			Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, vault.Token.Address),
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
			Info: actions.OptionInfo{
				Label: "APR",
				Value: fmt.Sprintf("%.2f%%", vault.APR.ForwardAPR.NetAPR*100+vault.Extra.StakingRewardsAPR*100),
			},
			Icon: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, vault.Token.Address),
		})
	}

	return options, nil
}
