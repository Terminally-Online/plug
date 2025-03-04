package simulation

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"solver/bindings/plug_router"
	"solver/internal/client"
	"solver/internal/database/models"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func SimulateRaw(livePlug *models.LivePlug, ABI *string) (*models.Run, error) {
	ctx := context.Background()

	rpcUrl, err := client.GetQuicknodeUrl(livePlug.ChainId)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.DialContext(ctx, rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %v", err)
	}
	defer rpcClient.Close()

	tx := map[string]interface{}{
		"from": livePlug.From,
		"to":   livePlug.To,
	}

	if len(livePlug.Data) > 0 {
		tx["data"] = livePlug.Data
	}

	if livePlug.Value != nil {
		value := hexutil.EncodeBig(livePlug.Value)
		tx["value"] = value
	}

	if livePlug.Gas != nil {
		gas := hexutil.EncodeBig(livePlug.Gas)
		tx["gas"] = gas
	}
	fmt.Printf("tx: %v\n", tx)

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

	callTraceConfig := map[string]interface{}{
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
		return nil, fmt.Errorf("trace call failed: %v", err)
	}

	status := "success"
	if trace.Error != "" {
		status = "failed"
	}

	run := &models.Run{
		LivePlugId: livePlug.Id,
		IntentId:   livePlug.IntentId,
		From:       livePlug.From,
		To:         livePlug.To,
		Value:      livePlug.Value,
		Status:     status,
		ResultData: models.RunOutputData{
			Raw: trace.Output,
		},
	}

	if trace.GasUsed != "" {
		gasUsed := new(big.Int)
		if _, ok := gasUsed.SetString(trace.GasUsed, 0); ok {
			run.GasEstimate = gasUsed.Uint64()
		}
	}

	if trace.Error != "" {
		run.Error = &trace.Error
	}

	if ABI != nil && len(trace.Output) > 0 && len(livePlug.Data) >= 4 {
		parsedABI, err := abi.JSON(strings.NewReader(*ABI))
		if err != nil {
			errorStr := fmt.Sprintf("failed to parse ABI: %v", err)
			run.Error = &errorStr
			return run, nil
		}

		methodIDHex := livePlug.Data[:10]
		methodID := common.Hex2Bytes(methodIDHex[2:])
		var method *abi.Method
		for _, m := range parsedABI.Methods {
			if bytes.Equal(m.ID, methodID) {
				method = &m
				break
			}
		}

		if method == nil {
			return nil, fmt.Errorf("method not found for ID: %x", methodID)
		}

		// TODO: Do this when there is nothing left more pressing to work on. This
		//       really should not be prioritized unless there are zero tickets and
		//       we have an actual use for it somewhere in the app or idea.
		// decoded, err := method.Outputs.Unpack(trace.Output)
		// if err != nil {
		// 	resp.ErrorMessage = fmt.Sprintf("failed to decode return data: %v", err)
		// 	return resp, nil
		// }
		// resp.Data.Decoded = decoded
	}

	return run, nil
}

func Simulate(livePlug *models.LivePlug) (*models.Run, error) {
	runResponse, err := SimulateRaw(livePlug, &plug_router.PlugRouterMetaData.ABI)
	if err != nil {
		return nil, err
	}

	return runResponse, nil
}
