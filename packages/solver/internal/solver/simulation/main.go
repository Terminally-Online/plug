package simulation

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os"
	"solver/bindings/plug_router"
	"solver/internal/bindings/references"
	"solver/internal/client"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func getSimulationRequest(id string, chainId uint64, plugs signature.LivePlugs) (*SimulationRequest, error) {
	routerAbi, err := plug_router.PlugRouterMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("PlugRouter")
	}

	plugCalldata, err := routerAbi.Pack("plug", plugs)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}
	return &SimulationRequest{
		Id:      id,
		ChainId: chainId,
		From:    common.HexToAddress(os.Getenv("SOLVER_ADDRESS")),
		To:      common.HexToAddress(references.Networks[chainId].References["plug"]["router"]),
		Data:    plugCalldata,
		Value:   big.NewInt(0),
		ABI:     plug_router.PlugRouterMetaData.ABI,
	}, nil
}

func SimulateRaw(req *SimulationRequest) (*SimulationResponse, error) {
	ctx := context.Background()

	rpcUrl, err := client.GetQuicknodeUrl(req.ChainId)
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

	resp := &SimulationResponse{
		Id:      req.Id,
		Success: trace.Error == "",
		Data: OutputData{
			Raw: trace.Output,
		},
	}

	if trace.GasUsed != "" {
		gasUsed := new(big.Int)
		if _, ok := gasUsed.SetString(trace.GasUsed[2:], 16); ok {
			resp.Gas.Used = gasUsed.Uint64()
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

	return resp, nil
}

func Simulate(id string, chainId uint64, plugs signature.LivePlugs) (*SimulationRequest, *SimulationResponse, error) {
	simulationRequest, err := getSimulationRequest(id, chainId, plugs)
	if err != nil {
		return nil, nil, err
	}

	simulationResponse, err := SimulateRaw(simulationRequest)
	if err != nil {
		return nil, nil, err
	}

	return simulationRequest, simulationResponse, nil
}
