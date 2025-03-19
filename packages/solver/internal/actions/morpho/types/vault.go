package types

type Vault struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Metadata struct {
		Image       string `json:"image"`
		Description string `json:"description"`
	} `json:"metadata"`
	Asset     Asset `json:"asset"`
	DailyApys struct {
		Apy    float64 `json:"apy"`
		NetApy float64 `json:"netApy"`
	} `json:"dailyApys"`
	State struct {
		Allocation []struct {
			Enabled bool   `json:"enabled"`
			Market  Market `json:"market"`
		} `json:"allocation"`
	} `json:"state"`
}
