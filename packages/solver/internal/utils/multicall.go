package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"solver/bindings/multicall_primary"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type MulticallCalldata struct {
	Target     common.Address
	Method     string
	Args       []interface{}
	ABI        *abi.ABI
	OutputType interface{}
}

func ExecuteMulticall(chainId uint64, multicallAddress common.Address, calls []MulticallCalldata) ([]interface{}, error) {
	provider, err := GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	multicallCalls := make([]multicall_primary.Multicall3Call, len(calls))
	for i, call := range calls {
		callData, err := call.ABI.Pack(call.Method, call.Args...)
		if err != nil {
			return nil, fmt.Errorf("failed to pack %s call: %w", call.Method, err)
		}

		multicallCalls[i] = multicall_primary.Multicall3Call{
			Target:   call.Target,
			CallData: callData,
		}
	}

	multicallAbi, err := multicall_primary.MulticallPrimaryMetaData.GetAbi()
	if err != nil {
		return nil, ErrABI("Multicall")
	}

	input, err := multicallAbi.Pack("aggregate", multicallCalls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack multicall aggregate: %w", err)
	}

	msg := ethereum.CallMsg{
		To:   &multicallAddress,
		Data: input,
	}

	output, err := provider.CallContract(context.Background(), msg, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make multicall: %w", err)
	}

	unpacked, err := multicallAbi.Unpack("aggregate", output)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack multicall aggregate: %w", err)
	}

	_, ok := unpacked[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("multicall result 0 was not a big.Int")
	}

	returnData, ok := unpacked[1].([][]byte)
	if !ok {
		return nil, fmt.Errorf("multicall result 1 was not a [][]byte")
	}

	results := make([]interface{}, len(returnData))
	for i, data := range returnData {
		unpacked, err := calls[i].ABI.Unpack(calls[i].Method, data)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack data for call %d: %w", i, err)
		}

		if len(unpacked) == 0 {
			return nil, fmt.Errorf("empty result for call %d", i)
		}

		jsonData, err := json.Marshal(unpacked[0])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data for call %d: %w", i, err)
		}

		result := calls[i].OutputType
		if err := json.Unmarshal(jsonData, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal data for call %d: %w", i, err)
		}

		results[i] = result
	}

	return results, nil
}

// Pow10 returns 10^n
func Pow10(n int) int64 {
	result := int64(1)
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}
