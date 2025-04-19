package simulation

import (
	"context"
	"fmt"
	"math/big"
	"solver/internal/client"
	"solver/internal/database/models"
	"solver/internal/database/types"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// SimulateLivePlugs performs a local simulation of a LivePlugs execution
// using the node's debug_traceCall API. It creates a Run record with the
// simulation results including gas estimates and error information.
func SimulateLivePlugs(livePlugs *signature.LivePlugs) (*models.Run, error) {
	ctx := context.Background()

	rpcUrl, err := client.GetQuicknodeUrl(livePlugs.ChainId)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.DialContext(ctx, rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %v", err)
	}
	defer rpcClient.Close()

	routerAddress := livePlugs.GetRouterAddress()

	tx := map[string]any{
		"from": livePlugs.From,
		"to":   routerAddress.Hex(),
	}

	if livePlugs.Data != "" {
		tx["data"] = livePlugs.Data
	} else {
		callData, err := livePlugs.GetCallData()
		if err != nil {
			return nil, err
		}
		tx["data"] = hexutil.Bytes(callData).String()
	}

	var blockNumber string
	if err := rpcClient.CallContext(ctx, &blockNumber, "eth_blockNumber"); err != nil {
		return nil, fmt.Errorf("failed to get block number: %v", err)
	}

	var baseFee struct {
		BaseFeePerGas string `json:"baseFeePerGas"`
	}
	if err := rpcClient.CallContext(ctx, &baseFee, "eth_getBlockByNumber", blockNumber, false); err != nil {
		return nil, fmt.Errorf("failed to get base fee: %v", err)
	}

	tx["gasPrice"] = baseFee.BaseFeePerGas

	callTraceConfig := map[string]any{
		"tracer": "callTracer",
	}

	var trace struct {
		Type     string         `json:"type"`
		From     common.Address `json:"from"`
		To       common.Address `json:"to"`
		Value    string         `json:"value"`
		Gas      string         `json:"gas"`
		GasUsed  string         `json:"gasUsed"`
		GasPrice string         `json:"gasPrice"`
		Input    hexutil.Bytes  `json:"input"`
		Output   hexutil.Bytes  `json:"output"`
		Error    string         `json:"error"`
	}

	if err := rpcClient.CallContext(ctx, &trace, "debug_traceCall", tx, "latest", callTraceConfig); err != nil {
		return nil, utils.ErrSimulationFailed(err.Error())
	}

	// Create run object with results
	status := "success"
	if trace.Error != "" {
		status = "failed"
	}

	run := &models.Run{
		LivePlugsId: livePlugs.Id,
		IntentId:    livePlugs.IntentId,
		From:        livePlugs.From,
		To:          routerAddress.Hex(),
		Status:      status,
		Data: models.RunOutputData{
			Raw: trace.Output,
		},
	}

	if trace.GasUsed != "" {
		gasUsed := new(big.Int)
		if _, ok := gasUsed.SetString(trace.GasUsed, 0); ok {
			run.GasUsed = gasUsed.Uint64()
		}
	}

	if trace.Error != "" {
		run.Error = &trace.Error
	}

	return run, nil
}

// SimulateEOATx simulates direct EOA transactions that bypass the Plug router contract.
// Unlike SimulateLivePlugs, this simulates a transaction sent directly from an EOA to
// the target contract, making it suitable for wallets that don't support Plug socket
// signatures. It still creates a Run record linked to the original LivePlugs.
func SimulateEOATx(tx *signature.Transaction, livePlugsId *string, chainId uint64) (*models.Run, error) {
	ctx := context.Background()

	rpcUrl, err := client.GetQuicknodeUrl(chainId)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.DialContext(ctx, rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %v", err)
	}
	defer rpcClient.Close()

	simTx := map[string]any{
		"from": tx.From.Hex(),
		"to":   tx.To.Hex(),
	}

	if len(tx.Data) > 0 {
		simTx["data"] = hexutil.Bytes(tx.Data).String()
	}

	if tx.Value != nil {
		simTx["value"] = hexutil.EncodeBig(tx.Value)
	}

	if tx.Gas != nil {
		simTx["gas"] = hexutil.EncodeBig(tx.Gas)
	}

	// Get block metadata for simulation
	var blockNumber string
	if err := rpcClient.CallContext(ctx, &blockNumber, "eth_blockNumber"); err != nil {
		return nil, fmt.Errorf("failed to get block number: %v", err)
	}

	var baseFee struct {
		BaseFeePerGas string `json:"baseFeePerGas"`
	}
	if err := rpcClient.CallContext(ctx, &baseFee, "eth_getBlockByNumber", blockNumber, false); err != nil {
		return nil, fmt.Errorf("failed to get base fee: %v", err)
	}

	simTx["gasPrice"] = baseFee.BaseFeePerGas
	callTraceConfig := map[string]any{
		"tracer": "callTracer",
	}

	var trace struct {
		Type     string         `json:"type"`
		From     common.Address `json:"from"`
		To       common.Address `json:"to"`
		Value    string         `json:"value"`
		Gas      string         `json:"gas"`
		GasUsed  string         `json:"gasUsed"`
		GasPrice string         `json:"gasPrice"`
		Input    hexutil.Bytes  `json:"input"`
		Output   hexutil.Bytes  `json:"output"`
		Error    string         `json:"error"`
	}

	if err := rpcClient.CallContext(ctx, &trace, "debug_traceCall", simTx, "latest", callTraceConfig); err != nil {
		return nil, utils.ErrSimulationFailed(err.Error())
	}

	// Create run object with results
	status := "success"
	if trace.Error != "" {
		status = "failed"
	}

	run := &models.Run{
		From:   tx.From.Hex(),
		To:     tx.To.Hex(),
		Value:  types.NewBigInt(tx.Value),
		Status: status,
		Data: models.RunOutputData{
			Raw: trace.Output,
		},
	}

	if livePlugsId != nil {
		run.LivePlugsId = *livePlugsId
	}

	if trace.GasUsed != "" {
		gasUsed := new(big.Int)
		if _, ok := gasUsed.SetString(trace.GasUsed, 0); ok {
			run.GasUsed = gasUsed.Uint64()
		}
	}

	if trace.Error != "" {
		run.Error = &trace.Error
	}

	return run, nil
}
