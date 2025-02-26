package simulation

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"solver/bindings/plug_router"
	"solver/internal/client"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func SimulateRaw(transaction Transaction, calldata []byte, ABI *string) (*models.Run, error) {
	ctx := context.Background()

	rpcUrl, err := client.GetQuicknodeUrl(transaction.ChainId)
	if err != nil {
		return nil, err
	}

	rpcClient, err := rpc.DialContext(ctx, rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client: %v", err)
	}
	defer rpcClient.Close()

	tx := map[string]interface{}{
		"from": transaction.From,
		"to":   transaction.To,
	}

	value := new(big.Int)
	if transaction.Value != "" {
		if _, ok := value.SetString(transaction.Value, 16); !ok {
			return nil, fmt.Errorf("failed to parse value: %v", transaction.Value)
		}

		if value.Sign() > 0 {
			tx["value"] = transaction.Value
		}
	}

	if len(calldata) > 0 {
		tx["data"] = hexutil.Bytes.String(calldata)
	}

	if transaction.Gas != nil {
		tx["gas"] = transaction.Gas
	}
	if len(transaction.AccessList) > 0 {
		tx["accessList"] = transaction.AccessList
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
		Status: status,
		ResultData: models.RunOutputData{
			Raw: trace.Output,
		},
	}

	if trace.GasUsed != "" {
		gasUsed := new(big.Int)
		if _, ok := gasUsed.SetString(trace.GasUsed, 16); ok {
			run.GasEstimate = gasUsed.Uint64()
		}
	}

	if trace.Error != "" {
		run.Error = &trace.Error
	}

	if ABI != nil && len(trace.Output) > 0 && len(calldata) >= 4 {
		parsedABI, err := abi.JSON(strings.NewReader(*ABI))
		if err != nil {
			errorStr := fmt.Sprintf("failed to parse ABI: %v", err)
			run.Error = &errorStr
			return run, nil
		}

		methodID := calldata[:4]
		var method *abi.Method
		for _, m := range parsedABI.Methods {
			if bytes.Equal(m.ID, methodID) {
				method = &m
				break
			}
		}

		if method == nil {
			errorStr := "method not found in ABI"
			run.Error = &errorStr
			return run, nil
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

	if err := database.DB.Create(run).Error; err != nil {
		return nil, fmt.Errorf("failed to save simulation run: %v", err)
	}

	return run, nil
}

func Simulate(transaction Transaction, plugs signature.LivePlugs) (*models.Run, error) {
	routerAbi, err := plug_router.PlugRouterMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("PlugRouter")
	}

	plugCalldata, err := routerAbi.Pack("plug", plugs)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	runResponse, err := SimulateRaw(transaction, plugCalldata, &plug_router.PlugRouterMetaData.ABI)
	if err != nil {
		return nil, err
	}

	return runResponse, nil
}
