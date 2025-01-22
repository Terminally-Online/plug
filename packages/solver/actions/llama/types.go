package llama

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
