package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PriceData struct {
	Decimals   int     `json:"decimals,omitempty"`
	Symbol     string  `json:"symbol,omitempty"`
	Price      float64 `json:"price,omitempty"`
	Change     float64 `json:"change,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
}

type PricePoint struct {
	Timestamp int64   `json:"timestamp"`
	Price     float64 `json:"price"`
}

type CoinResponse struct {
	Coins map[string]struct {
		Decimals   int          `json:"decimals,omitempty"`
		Symbol     string       `json:"symbol,omitempty"`
		Prices     []PricePoint `json:"prices"`
		Confidence float64      `json:"confidence,omitempty"`
	} `json:"coins"`
}

func GetPriceKey(chain, address string) string {
	return fmt.Sprintf("%s:%s", strings.ToLower(chain), address)
}

func GetPrices(queries []string) (map[string]PriceData, error) {
	if len(queries) == 0 {
		return nil, nil
	}

	query := strings.Join(queries, ",")
	url := fmt.Sprintf("https://coins.llama.fi/chart/%s?span=48&period=30m&searchWidth=1200", query)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prices: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response CoinResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	result := make(map[string]PriceData)
	for key, coin := range response.Coins {
		if len(coin.Prices) == 0 {
			continue
		}

		start := coin.Prices[0].Price
		end := coin.Prices[len(coin.Prices)-1].Price
		var change float64
		if start != 0 {
			change = ((end - start) / start) * 100
		}

		result[key] = PriceData{
			Decimals:   coin.Decimals,
			Symbol:     coin.Symbol,
			Price:      end,
			Change:     change,
			Confidence: coin.Confidence,
		}
	}

	return result, nil
}
