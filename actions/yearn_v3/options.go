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
		for _, underlyingAddr := range vault.Token.UnderlyingTokensAddresses {
			if underlyingAddr == "" {
				continue
			}

			lowerAddr := strings.ToLower(underlyingAddr)

			if _, exists := tokenMap[lowerAddr]; !exists {
				if token, ok := tokenDetails[lowerAddr]; ok {
					tokenMap[lowerAddr] = types.Option{
						Value: fmt.Sprintf("%s:%d", token.Address, token.Decimals),
						Name:  token.Name,
						Label: token.Symbol,
						Icon:  fmt.Sprintf("https://token-icons.llamao.fi/icons/tokens/1/%s", underlyingAddr),
					}
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
