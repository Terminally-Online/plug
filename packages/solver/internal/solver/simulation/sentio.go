package simulation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"time"

	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	SentioBaseURL  = "https://app.sentio.xyz/api/v1/solidity"
	DefaultTimeout = 30 * time.Second
	ProjectOwner   = "mason"
	ProjectSlug    = "plugsims"
)

var Sentio *SentioClient

type SentioClient struct {
	APIKey       string
	ProjectOwner string
	ProjectSlug  string
	BaseURL      string
	HttpClient   *http.Client
}

type SentioSimulationResponse struct {
	Simulation struct {
		ID     string `json:"id"`
		Result struct {
			TransactionReceipt struct {
				EffectiveGasPrice string `json:"effectiveGasPrice"`
			} `json:"transactionReceipt"`
		} `json:"result"`
	} `json:"simulation"`
}

func init() {
	Sentio = &SentioClient{
		ProjectOwner: ProjectOwner,
		ProjectSlug:  ProjectSlug,
		APIKey:       os.Getenv("SENTIO_API_KEY"),
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
		gas = hexutil.EncodeBig(big.NewInt(20000000)) // 20M gas
	}
	// 1 gwei, sentio also requires gasPrice
	gasPrice := hexutil.EncodeBig(big.NewInt(1000000000))

	reqBody := map[string]any{
		"simulation": map[string]any{
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

	headers := map[string]string{
		"Content-Type": "application/json",
		"api-key":      c.APIKey,
	}

	simResponse, err := utils.MakeHTTPRequest(
		url,
		"POST",
		headers,
		nil,
		bytes.NewBuffer(jsonData),
		SentioSimulationResponse{},
	)
	if err != nil {
		return nil, err
	}

	// Retrieved the simulation ID, now time to fetch the traces from it.
	traceURL := fmt.Sprintf("%s/%s/%s/%s/simulation/%s/call_trace",
		c.BaseURL, c.ProjectOwner, c.ProjectSlug, tx.ChainId, simResponse.Simulation.ID)

	traceHeaders := map[string]string{
		"api-key": c.APIKey,
	}

	trace, err := utils.MakeHTTPRequest(
		traceURL,
		"GET",
		traceHeaders,
		nil,
		nil,
		Trace{},
	)
	if err != nil {
		return nil, fmt.Errorf("trace request failed: %v", err)
	}

	trace.GasPrice = simResponse.Simulation.Result.TransactionReceipt.EffectiveGasPrice

	return &trace, nil
}
