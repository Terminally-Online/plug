package simulation

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"os"
	"solver/internal/utils"

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

	resp, err := s.Simulate(req)
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

func (s *Simulator) Simulate(req SimulationRequest) (*SimulationResponse, error) {
	ctx := context.Background()

	rpcUrl, err := utils.GetProviderUrl(req.ChainId)
	if err != nil { return nil, err }

	rpcClient, err := rpc.DialContext(ctx, rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %v", err)
	}
	defer rpcClient.Close()

	var (
		blockNumber string
		gasEstimate string
		callResult  string
	)

	valueHex := "0x0"
	if req.Value != nil && req.Value.Sign() > 0 {
		valueHex = hexutil.EncodeBig(req.Value)
	}

	tx := map[string]interface{}{
		"from":  req.From.Hex(),
		"to":    req.To.Hex(),
		"value": valueHex,
	}
	if len(req.Data) > 0 {
		tx["data"] = req.Data.String()
	}

	batch := []rpc.BatchElem{
		{
			Method: "eth_blockNumber",
			Result: &blockNumber,
		},
		{
			Method: "eth_estimateGas",
			Args:   []interface{}{tx},
			Result: &gasEstimate,
		},
	}

	if len(req.Data) > 0 {
		batch = append(batch, rpc.BatchElem{
			Method: "eth_call",
			Args:   []interface{}{tx, "latest"},
			Result: &callResult,
		})
	}

	err = rpcClient.BatchCall(batch)
	if err != nil {
		return nil, fmt.Errorf("batch call failed: %v", err)
	}

	if batch[0].Error != nil {
		return nil, fmt.Errorf("failed to get block number: %v", batch[0].Error)
	}
	if batch[1].Error != nil {
		return &SimulationResponse{
			Success:      false,
			ErrorMessage: batch[1].Error.Error(),
		}, nil
	}

	blockNum := new(big.Int)
	blockNum.SetString(blockNumber[2:], 16)

	gasUsed := new(big.Int)
	gasUsed.SetString(gasEstimate[2:], 16)

	if req.GasLimit != nil && *req.GasLimit < gasUsed.Uint64() {
		gasUsed.SetUint64(*req.GasLimit)
	}

	if len(req.Data) > 0 && len(batch) > 2 && batch[2].Error != nil {
		return &SimulationResponse{
			GasUsed:      gasUsed.Uint64(),
			BlockNumber:  blockNum.Uint64(),
			Success:      false,
			ErrorMessage: batch[2].Error.Error(),
		}, nil
	}

	var returnData []byte
	if len(callResult) > 0 {
		returnData, err = hexutil.Decode(callResult)
		if err != nil {
			return &SimulationResponse{
				GasUsed:      gasUsed.Uint64(),
				BlockNumber:  blockNum.Uint64(),
				Success:      false,
				ErrorMessage: fmt.Sprintf("failed to decode return data: %v", err),
			}, nil
		}
	}

	resp := &SimulationResponse{
		GasUsed:      gasUsed.Uint64(),
		BlockNumber:  blockNum.Uint64(),
		Success:      true,
		ReturnData:   returnData,
	}

	return resp, nil
}
