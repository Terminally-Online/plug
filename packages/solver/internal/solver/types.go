package solver

type SimulationRequest struct {
	Id          string   `json:"id"`
	Status      string   `json:"status"`
	Error       string   `json:"error,omitempty"`
	Errors      []string `json:"errors,omitempty"`
	GasEstimate int      `json:"gasEstimate,omitempty"`
}

type SimulationsRequest struct {
	Json []SimulationRequest `json:"json"`
}

type SimulationsResponse struct {
	Result struct {
		Data struct {
			Json []string `json:"json"`
		} `json:"data"`
	} `json:"result"`
}

type ExecutionRequest struct {
	Id      string           `json:"id"`
	ChainId uint64           `json:"chainId"`
	From    string           `json:"from"`
	Inputs  []map[string]any `json:"inputs"`
}

type ExecutionsRequest struct {
	Result struct {
		Data struct {
			Json []ExecutionRequest `json:"json"`
		} `json:"data"`
	} `json:"result"`
}
