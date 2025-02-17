package simulation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"solver/bindings/plug_router"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"strings"

	"bytes"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type Simulator struct {
	signer common.Address
}

func New() Simulator {
	return Simulator{
		signer: common.HexToAddress(os.Getenv("SOLVER_ADDRESS")),
	}
}

func (s *Simulator) GetNext() (SimulationDefinitions, error) {
	url := fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.simulation.simulateNext")
	response, err := utils.MakeHTTPRequest(
		url,
		"POST",
		map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    os.Getenv("PLUG_APP_API_KEY"),
		},
		nil,
		nil,
		SimulationDefinitions{},
	)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (s *Simulator) GetSimulationRequest(
	executionId string, chainId uint64, plugs signature.LivePlugs,
) (*SimulationRequest, error) {
	routerAbi, err := plug_router.PlugRouterMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("PlugRouter")
	}

	plugCalldata, err := routerAbi.Pack("plug", plugs)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}
	return &SimulationRequest{
		ExecutionId: executionId,
		ChainId:     chainId,
		From:        common.HexToAddress(os.Getenv("SOLVER_ADDRESS")),
		To:          common.HexToAddress(references.Networks[chainId].References["plug"]["router"]),
		Data:        plugCalldata,
		Value:       big.NewInt(0),
		ABI:         plug_router.PlugRouterMetaData.ABI,
	}, nil
}

func (s *Simulator) GetSimulationResponse(
	executionId string, chainId uint64, plugs signature.LivePlugs,
) (
	*SimulationRequest, *SimulationResponse, error,
) {
	simulationRequest, err := s.GetSimulationRequest(executionId, chainId, plugs)
	if err != nil {
		return nil, nil, err
	}

	simulationResponse, err := s.Simulate(simulationRequest)
	if err != nil {
		return nil, nil, err
	}

	return simulationRequest, simulationResponse, nil
}

func (s *Simulator) PostSimulation(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req SimulationRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request: %v", err), http.StatusBadRequest)
		return
	}

	if req.ChainId == 0 {
		http.Error(w, "Chain ID is required", http.StatusBadRequest)
		return
	}

	resp, err := s.Simulate(&req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Simulation failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}

func (s *Simulator) PostSimulations(simulations []SimulationResponse) error {
	response := SimulationResponses{
		Json: simulations,
	}
	body, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = utils.MakeHTTPRequest(
		fmt.Sprintf("%s%s", os.Getenv("PLUG_APP_API_URL"), "jobs.simulation.simulated"),
		"POST",
		map[string]string{
			"Content-Type": "application/json",
			"X-API-Key":    os.Getenv("PLUG_APP_API_KEY"),
		},
		nil,
		bytes.NewReader(body),
		SimulationResponses{},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Simulator) Simulate(req *SimulationRequest) (*SimulationResponse, error) {
	ctx := context.Background()

	rpcUrl, err := utils.GetProviderUrl(req.ChainId)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.DialContext(ctx, rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %v", err)
	}
	defer rpcClient.Close()

	tx := map[string]interface{}{
		"from": req.From.Hex(),
		"to":   req.To.Hex(),
	}

	if req.Value != nil && req.Value.Sign() > 0 {
		tx["value"] = hexutil.EncodeBig(req.Value)
	}
	if len(req.Data) > 0 {
		tx["data"] = req.Data.String()
	}
	if req.GasLimit != nil {
		tx["gas"] = hexutil.EncodeUint64(*req.GasLimit)
	}
	if len(req.AccessList) > 0 {
		tx["accessList"] = req.AccessList
	}

	callTraceConfig := map[string]interface{}{
		"tracer": "callTracer",
	}

	var trace struct {
		Type    string         `json:"type"`
		From    common.Address `json:"from"`
		To      common.Address `json:"to"`
		Value   string         `json:"value"`
		Gas     string         `json:"gas"`
		GasUsed string         `json:"gasUsed"`
		Input   hexutil.Bytes  `json:"input"`
		Output  hexutil.Bytes  `json:"output"`
		Error   string         `json:"error"`
	}

	if err := rpcClient.CallContext(ctx, &trace, "debug_traceCall", tx, "latest", callTraceConfig); err != nil {
		return nil, fmt.Errorf("trace call failed: %v", err)
	}

	resp := &SimulationResponse{
		ExecutionId: req.ExecutionId,
		Success:     trace.Error == "",
		Data: OutputData{
			Raw: trace.Output,
		},
	}

	if trace.GasUsed != "" {
		gasUsed := new(big.Int)
		if _, ok := gasUsed.SetString(trace.GasUsed[2:], 16); ok {
			resp.GasUsed = gasUsed.Uint64()
		}
	}

	if trace.Error != "" {
		resp.ErrorMessage = trace.Error
	}

	if req.ABI != "" && len(trace.Output) > 0 && len(req.Data) >= 4 {
		parsedABI, err := abi.JSON(strings.NewReader(req.ABI))
		if err != nil {
			resp.ErrorMessage = fmt.Sprintf("failed to parse ABI: %v", err)
			return resp, nil
		}

		methodID := req.Data[:4]
		var method *abi.Method
		for _, m := range parsedABI.Methods {
			if bytes.Equal(m.ID, methodID) {
				method = &m
				break
			}
		}

		if method == nil {
			resp.ErrorMessage = "method not found in ABI"
			return resp, nil
		}

		// decoded, err := method.Outputs.Unpack(trace.Output)
		// if err != nil {
		// 	resp.ErrorMessage = fmt.Sprintf("failed to decode return data: %v", err)
		// 	return resp, nil
		// }
		//
		// resp.Data.Decoded = decoded
	}

	return resp, nil
}
