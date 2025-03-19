package options

import (
	"fmt"
	"solver/internal/actions"
	"solver/internal/actions/yearn_v3/reads"
	"solver/internal/actions/yearn_v3/types"
	"strings"
)

func GetUnderlyingAssetOptions(chainId uint64) ([]actions.Option, error) {
	tokens, err := reads.GenerateTokenList()
	if err != nil {
		return nil, err
	}

	tokenDetails := make(map[string]types.Token)
	for _, token := range tokens {
		tokenDetails[strings.ToLower(token.Address)] = token
	}
	vaults, err := reads.GetVaults(chainId)
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
