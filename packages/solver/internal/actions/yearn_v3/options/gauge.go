package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/yearn_v3/reads"
	"strings"
)

func GetAvailableStakingGaugeOptions(chainId uint64) ([]actions.Option, error) {
	vaults, err := reads.GetVaults(chainId)
	if err != nil {
		return nil, err
	}

	var options []actions.Option
	for _, vault := range vaults {
		if !vault.Staking.Available || vault.Staking.Address == "" {
			continue
		}

		options = append(options, actions.Option{
			Value: fmt.Sprintf("%s:%d", vault.Staking.Address, vault.Decimals),
			Name:  vault.DisplayName,
			Label: vault.FormattedSymbol,
			Icon:  &actions.OptionIcon{Default: fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/%d/%s", chainId, strings.ToLower(vault.Token.Address))},
		})
	}

	return options, nil
}

func AvailableStakingGaugeOptions[T any](lookup *actions.SchemaLookup[T]) (map[int]actions.Options, error) {
	availableStakingGaugeOptions, err := GetAvailableStakingGaugeOptions(lookup.ChainId)
	if err != nil {
		return nil, err
	}
	return map[int]actions.Options{1: {Simple: availableStakingGaugeOptions}}, nil
}
