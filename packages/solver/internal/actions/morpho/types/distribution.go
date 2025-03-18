package types

type Distribution struct {
	Claimable   string   `json:"claimable"`
	Proof       []string `json:"proof"`
	TxData      string   `json:"txData"`
	Distributor struct {
		Id      string `json:"id"`
		Address string `json:"address"`
		ChainId int    `json:"chain_id"`
	} `json:"distributor"`
	Asset struct {
		Id      string `json:"id"`
		Address string `json:"address"`
		ChainId int    `json:"chain_id"`
	} `json:"asset"`
}

type DistributionResponse struct {
	Timestamp  string `json:"timestamp"`
	Pagination struct {
		PerPage    int    `json:"per_page"`
		Page       int    `json:"page"`
		TotalPages int    `json:"total_pages"`
		Next       string `json:"next,omitempty"`
		Prev       string `json:"prev,omitempty"`
	} `json:"pagination"`
	Data []Distribution `json:"data"`
}
