package types

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
