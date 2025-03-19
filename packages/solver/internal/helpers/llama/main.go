package llama

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type LlamaPriceData struct {
	Decimals   int     `json:"decimals,omitempty"`
	Symbol     string  `json:"symbol,omitempty"`
	Price      float64 `json:"price,omitempty"`
	Change     float64 `json:"change,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
}

type LlamaPricePoint struct {
	Timestamp int64   `json:"timestamp"`
	Price     float64 `json:"price"`
}

type LlamaCoinResponse struct {
	Coins map[string]struct {
		Decimals   int               `json:"decimals,omitempty"`
		Symbol     string            `json:"symbol,omitempty"`
		Prices     []LlamaPricePoint `json:"prices"`
		Confidence float64           `json:"confidence,omitempty"`
	} `json:"coins"`
}

// duration for calculating TWAP in minutes
var TWAPDurationMinutes = 180 // currently using a 3 hour twap

func GetPriceKey(chain, address string) string {
	return fmt.Sprintf("%s:%s", strings.ToLower(chain), address)
}

func GetPrices(queries []string) (map[string]LlamaPriceData, error) {
	if len(queries) == 0 {
		return nil, nil
	}

	period := 10             // 10 minute intervals
	span := 24 * 60 / period // always 24 hours to get 1d change data

	query := strings.Join(queries, ",")
	url := fmt.Sprintf("https://coins.llama.fi/chart/%s?span=%d&period=%d", query, span, period)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch prices: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response LlamaCoinResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	result := make(map[string]LlamaPriceData)
	for key, coin := range response.Coins {
		if len(coin.Prices) == 0 {
			continue
		}

		oldestPrice := coin.Prices[0].Price

		// Calculate TWAP over the last TWAPDurationMinutes minutes
		twapPrice := calculateTWAP(coin.Prices, TWAPDurationMinutes)

		// Calculate change using the start price and TWAP
		var change float64
		if oldestPrice != 0 {
			change = ((twapPrice - oldestPrice) / oldestPrice) * 100
		}

		result[key] = LlamaPriceData{
			Decimals:   coin.Decimals,
			Symbol:     coin.Symbol,
			Price:      twapPrice,
			Change:     change,
			Confidence: coin.Confidence,
		}
	}

	return result, nil
}

// calculateTWAP calculates the time-weighted average price for the given duration in hours
func calculateTWAP(prices []LlamaPricePoint, durationMinutes int) float64 {
	if len(prices) == 0 {
		return 0
	}

	// Get the current timestamp (last price point)
	currentTime := prices[len(prices)-1].Timestamp

	// Calculate cutoff timestamp (current time - duration in seconds)
	cutoffTime := currentTime - int64(durationMinutes*60)

	var twapSum float64
	var count int

	// Go through prices in reverse to get the most recent ones first
	for i := len(prices) - 1; i >= 0; i-- {
		if prices[i].Timestamp < cutoffTime {
			break
		}

		twapSum += prices[i].Price
		count++
	}

	if count == 0 {
		// If no prices in the TWAP window, return the most recent price
		return prices[len(prices)-1].Price
	}

	return twapSum / float64(count)
}
