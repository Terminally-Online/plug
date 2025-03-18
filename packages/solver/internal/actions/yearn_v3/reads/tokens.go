package reads

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"solver/internal/actions/yearn_v3/types"
	"time"
)

var (
	tokensCache     []types.Token
	tokensUpdatedAt int64
	cacheDuration   int64 = 300
)

func FetchTokenList(chainId uint64) ([]types.Token, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/SmolDapp/tokenLists/main/lists/%d.json", chainId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("could not fetch token list")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenList types.TokenList
	if err := json.Unmarshal(body, &tokenList); err != nil {
		return nil, err
	}

	return tokenList.Tokens, nil
}

func GenerateTokenList(force ...bool) ([]types.Token, error) {
	currentTime := time.Now().Unix()
	if !((len(force) > 0 && force[0]) || tokensCache == nil || (currentTime-tokensUpdatedAt) >= cacheDuration) {
		return tokensCache, nil
	}

	chainIDs := []uint64{1, 10, 8453}
	var allTokens []types.Token

	for _, chainID := range chainIDs {
		tokens, err := FetchTokenList(chainID)
		if err != nil {
			return nil, err
		}
		allTokens = append(allTokens, tokens...)
	}

	tokensCache = allTokens
	tokensUpdatedAt = currentTime

	return allTokens, nil
}
