package yearn_v3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"solver/utils"
	"time"
)

var (
	vaultsCache     []YearnVault
	tokensCache     []Token
	vaultsUpdatedAt int64
	tokensUpdatedAt int64
	cacheDuration   int64 = 300
)

type YearnVault struct {
	Address           string `json:"address"`
	Type              string `json:"type"`
	Kind              string `json:"kind"`
	Symbol            string `json:"symbol"`
	DisplaySymbol     string `json:"display_symbol"`
	FormattedSymbol   string `json:"formated_symbol"`
	Name              string `json:"name"`
	DisplayName       string `json:"display_name"`
	FormattedName     string `json:"formated_name"`
	Icon              string `json:"icon"`
	Version           string `json:"version"`
	Category          string `json:"category"`
	Decimals          int    `json:"decimals"`
	ChainID           int    `json:"chainID"`
	Endorsed          bool   `json:"endorsed"`
	Boosted           bool   `json:"boosted"`
	EmergencyShutdown bool   `json:"emergency_shutdown"`
	Token             struct {
		Address                   string   `json:"address"`
		UnderlyingTokensAddresses []string `json:"underlyingTokensAddresses"`
		Name                      string   `json:"name"`
		Symbol                    string   `json:"symbol"`
		Type                      string   `json:"type"`
		DisplayName               string   `json:"display_name"`
		DisplaySymbol             string   `json:"display_symbol"`
		Description               string   `json:"description"`
		Icon                      string   `json:"icon"`
		Decimals                  int      `json:"decimals"`
	} `json:"token"`
	TVL struct {
		TotalAssets string  `json:"totalAssets"`
		TVL         float64 `json:"tvl"`
		Price       float64 `json:"price"`
	} `json:"tvl"`
	Extra struct {
		StakingRewardsAPR float64 `json:"stakingRewardsAPR"`
		GammaRewardAPR    float64 `json:"gammaRewardAPR"`
	} `json:"extra"`
	APR struct {
		Type   string  `json:"type"`
		NetAPR float64 `json:"netAPR"`
		Points struct {
			WeekAgo   float64 `json:"weekAgo"`
			MonthAgo  float64 `json:"monthAgo"`
			Inception float64 `json:"inception"`
		} `json:"points"`
		ForwardAPR struct {
			Type   string  `json:"type"`
			NetAPR float64 `json:"netAPR"`
		} `json:"forwardAPR"`
	} `json:"apr"`
	Details struct {
		IsRetired     bool `json:"isRetired"`
		IsHidden      bool `json:"isHidden"`
		IsAggregator  bool `json:"isAggregator"`
		IsBoosted     bool `json:"isBoosted"`
		IsAutomated   bool `json:"isAutomated"`
		IsHighlighted bool `json:"isHighlighted"`
		IsPool        bool `json:"isPool"`
	} `json:"details"`
	Staking struct {
		Address   string `json:"address"`
		Available bool   `json:"available"`
		Source    string `json:"source"`
		Rewards   []struct {
			Address    string  `json:"address"`
			Name       string  `json:"name"`
			Symbol     string  `json:"symbol"`
			Decimals   int     `json:"decimals"`
			Price      float64 `json:"price"`
			IsFinished bool    `json:"isFinished"`
			FinishedAt int64   `json:"finishedAt"`
			APR        float64 `json:"apr"`
			PerWeek    float64 `json:"perWeek"`
		} `json:"rewards"`
	} `json:"staking"`
}

type Token struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	LogoURI  string `json:"logoURI"`
	ChainID  int    `json:"chainId"`
	Decimals int    `json:"decimals"`
}

type TokenList struct {
	Tokens []Token `json:"tokens"`
}

func GetVaults(force ...bool) ([]YearnVault, error) {
	currentTime := time.Now().Unix()
	if !((len(force) > 0 && force[0]) || vaultsCache == nil || (currentTime-vaultsUpdatedAt) >= cacheDuration) {
		return vaultsCache, nil
	}

	url := "https://ydaemon.yearn.finance/1/vaults/all"
	response, err := utils.MakeHTTPRequest(
		url,
		"GET",
		map[string]string{
			"Content-Type": "application/json",
		},
		nil,
		nil,
		[]YearnVault{},
	)
	if err != nil {
		return nil, err
	}

	endorsedVaults := make([]YearnVault, 0)
	for _, vault := range response {
		if vault.Endorsed && !vault.EmergencyShutdown && !vault.Details.IsRetired && !vault.Details.IsHidden {
			endorsedVaults = append(endorsedVaults, vault)
		}
	}

	vaultsCache = endorsedVaults
	vaultsUpdatedAt = currentTime

	return endorsedVaults, nil
}

func FetchTokenList(chainID int) ([]Token, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/SmolDapp/tokenLists/main/lists/%d.json", chainID)
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

	var tokenList TokenList
	if err := json.Unmarshal(body, &tokenList); err != nil {
		return nil, err
	}

	return tokenList.Tokens, nil
}

func GenerateTokenList(force ...bool) ([]Token, error) {
	currentTime := time.Now().Unix()
	if !((len(force) > 0 && force[0]) || tokensCache == nil || (currentTime-tokensUpdatedAt) >= cacheDuration) {
		return tokensCache, nil
	}

	chainIDs := []int{1, 10, 8453}
	var allTokens []Token

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
