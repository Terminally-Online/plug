package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/yearn_v3/reads"
	"strings"
)

func GetAvailableTokenToStakingGaugeOptions(chainId uint64) (tokenOptions []actions.Option, gaugeOptions []actions.Option, tokenToGaugeOptions map[string][]actions.Option, err error) {
	vaults, err := reads.GetVaults(chainId)
	if err != nil {
		return nil, nil, nil, err
	}

	tokenOptions = make([]actions.Option, 0)
	gaugeOptions = make([]actions.Option, 0)
	tokenToGaugeOptions = make(map[string][]actions.Option)
	seenToken := make(map[string]bool)

	for _, vault := range vaults {
		if !vault.Staking.Available || vault.Staking.Address == "" {
			continue
		}

		if !seenToken[vault.Token.Address] {
			tokenOptions = append(tokenOptions, actions.Option{
				Label: vault.Token.DisplayName,
				Name:  vault.Token.Name,
				Value: fmt.Sprintf("%s:%d", vault.Token.Address, vault.Decimals),
				Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, strings.ToLower(vault.Token.Address))},
			})
			seenToken[vault.Token.Address] = true
		}

		var gaugeApr string
		if len(vault.Staking.Rewards) > 0 {
			gaugeAprFloat := vault.Staking.Rewards[0].APR
			gaugeApr = fmt.Sprintf("%.2f%%", gaugeAprFloat*10)
		} else {
			gaugeApr = "0.0%"
		}

		gaugeOption := actions.Option{
			Label: vault.Name,
			Name:  vault.DisplayName,
			Value: vault.Staking.Address,
			Info:  &actions.OptionInfo{Label: "Gauge APR", Value: gaugeApr},
			Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, strings.ToLower(vault.Token.Address))},
		}

		tokenKey := fmt.Sprintf("%s:%d", vault.Token.Address, vault.Decimals)
		gaugeOptions = append(gaugeOptions, gaugeOption)
		tokenToGaugeOptions[tokenKey] = append(tokenToGaugeOptions[tokenKey], gaugeOption)
	}

	return tokenOptions, gaugeOptions, tokenToGaugeOptions, nil
}

func AvailableStakingGaugeOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	tokenOptions, _, tokenToGaugeOptions, err := GetAvailableTokenToStakingGaugeOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}

	return map[int]actions.Options{
		1: {Simple: tokenOptions},
		2: {Complex: tokenToGaugeOptions},
	}, nil
}
