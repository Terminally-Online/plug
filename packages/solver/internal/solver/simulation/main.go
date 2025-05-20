package simulation

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"solver/internal/client"
	"solver/internal/database/models"
	"solver/internal/database/types"
	"solver/internal/solver/signature"
	"solver/internal/utils"

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

	var callData string
	if livePlugs.Data != "" {
		callData = livePlugs.Data
	} else {
		callDataBytes, err := livePlugs.GetCallData()
		if err != nil {
			return nil, err
		}
		callData = hexutil.Bytes(callDataBytes).String()
	}

	var blockNumber string
	if err := rpcClient.CallContext(ctx, &blockNumber, "eth_blockNumber"); err != nil {
		return nil, fmt.Errorf("failed to get block number: %v", err)
	}

	// Calculate total value by summing up values from CallWithValue plugs
	// TODO MASON: do we need to do this? Why haven't we done it before?
	totalValue := new(big.Int)
	for _, plug := range livePlugs.Plugs.Plugs {
		if plug.Value != nil {
			totalValue.Add(totalValue, plug.Value)
		}
	}

	simRequest := SimulationRequest{
		ChainId: fmt.Sprintf("%d", livePlugs.ChainId),
		From:    os.Getenv("SOLVER_ADDRESS"),
		To:      routerAddress.Hex(),
		Data:    callData,
		Value:   totalValue,
		Gas:     big.NewInt(20_000_000),
	}

	fmt.Printf("SimulateLivePlugs::simRequest: %+v\n", simRequest)

	trace, err := Sentio.SimulateTransaction(simRequest)
	if err != nil {
		return nil, fmt.Errorf("sentio simulation failed: %v", err)
	}
	utils.LogObject("SimulateLivePlugs::trace", trace)

	status := "success"
	var errorMsg *string
	revertReason, _ := FindRevertError(trace)
	if revertReason != "" {
		status = "failed"
		errorMsg = &revertReason
	}

	run := &models.Run{
		LivePlugsId: livePlugs.Id,
		IntentId:    livePlugs.IntentId,
		From:        os.Getenv("SOLVER_ADDRESS"),
		To:          routerAddress.Hex(),
		Status:      status,
		Error:       errorMsg,
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

	simRequest := SimulationRequest{
		ChainId: fmt.Sprintf("%d", chainId),
		From:    tx.From.Hex(),
		To:      tx.To.Hex(),
		Data:    hexutil.Bytes(tx.Data).String(),
		Value:   tx.Value,
	}

	if len(tx.Data) > 0 {
		simRequest.Data = hexutil.Bytes(tx.Data).String()
	}

	trace, err := Sentio.SimulateTransaction(simRequest)
	if err != nil {
		return nil, fmt.Errorf("sentio simulation failed: %v", err)
	}
	utils.LogObject("SimulateEOATx::trace", trace)

	status := "success"
	var errorMsg *string
	revertReason, _ := FindRevertError(trace)
	if revertReason != "" {
		status = "failed"
		errorMsg = &revertReason
	}

	run := &models.Run{
		From:   tx.From.Hex(),
		To:     tx.To.Hex(),
		Value:  types.NewBigInt(tx.Value),
		Status: status,
		Error:  errorMsg,
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
