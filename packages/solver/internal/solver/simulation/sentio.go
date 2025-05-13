package simulation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	SentioBaseURL  = "https://app.sentio.xyz/api/v1/solidity"
	DefaultTimeout = 30 * time.Second
)

var Sentio *SentioClient

type SentioClient struct {
	APIKey       string
	ProjectOwner string
	ProjectSlug  string
	BaseURL      string
	HttpClient   *http.Client
}

func init() {
	Sentio = &SentioClient{
		ProjectOwner: "mason",
		ProjectSlug:  "testing",
		APIKey:       "6hyo9WIlX1As5wG6NWLfMTu8H6Y5BS9Wr",
		BaseURL:      SentioBaseURL,
		HttpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

/**
 * SimulateTransaction simulates a transaction using the Sentio API and fetches the traces.
 * It sends a POST request to the Sentio API with the transaction details, then uses the simulation ID to make
 * another request to retrieve the trace data simulation result.
 */
func (c *SentioClient) SimulateTransaction(tx SimulationRequest) (*Trace, error) {
	// https://app.sentio.xyz/api/v1/solidity/{owner}/{slug}/{chainId}/simulation
	url := fmt.Sprintf("%s/%s/%s/%s/simulation", c.BaseURL, c.ProjectOwner, c.ProjectSlug, tx.ChainId)

	// debug_traceCall does not require a gas limit, but sentio requires it.
	var gas string
	if tx.Gas != nil {
		gas = hexutil.EncodeBig(tx.Gas)
	} else {
		gas = "0x30D40"
	}
	// 1 gwei, sentio also requires gasPrice (boo)
	gasPrice := "0xF4240"

	reqBody := map[string]interface{}{
		"simulation": map[string]interface{}{
			"networkId":        tx.ChainId,
			"chainId":          tx.ChainId,
			"to":               tx.To,
			"from":             tx.From,
			"input":            tx.Data,
			"blockNumber":      "latest",
			"gas":              gas,
			"gasPrice":         gasPrice,
			"value":            hexutil.EncodeBig(tx.Value),
			"transactionIndex": "0", // this does not matter unless we're hitting existing blocks
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal simulation request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create simulation request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("api-key", c.APIKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("simulation request failed: %v", err)
	}
	defer resp.Body.Close()

	var simResponse struct {
		Simulation struct {
			ID     string `json:"id"`
			Result struct {
				TransactionReceipt struct {
					EffectiveGasPrice string `json:"effectiveGasPrice"`
				} `json:"transactionReceipt"`
			} `json:"result"`
		} `json:"simulation"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&simResponse); err != nil {
		return nil, fmt.Errorf("failed to parse simulation response: %v", err)
	}

	// Retrieved the simulation ID, now time to fetch the traces from it.
	traceURL := fmt.Sprintf("%s/%s/%s/%s/simulation/%s/trace",
		c.BaseURL, c.ProjectOwner, c.ProjectSlug, tx.ChainId, simResponse.Simulation.ID)

	traceReq, err := http.NewRequest("GET", traceURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create trace request: %v", err)
	}
	traceReq.Header.Set("api-key", c.APIKey)

	traceResp, err := c.HttpClient.Do(traceReq)
	if err != nil {
		return nil, fmt.Errorf("trace request failed: %v", err)
	}
	defer traceResp.Body.Close()

	var trace Trace
	if err := json.NewDecoder(traceResp.Body).Decode(&trace); err != nil {
		return nil, fmt.Errorf("failed to parse trace response: %v", err)
	}

	fmt.Printf("trace response: %+v\n", trace)

	trace.GasPrice = simResponse.Simulation.Result.TransactionReceipt.EffectiveGasPrice

	return &trace, nil
}
