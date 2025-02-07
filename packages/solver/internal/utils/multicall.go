package utils

import (
	"context"
	"fmt"
	"math/big"

	"solver/bindings/multicall_primary"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type MulticallCalldata struct {
	Target common.Address
	Method string
	Args   []interface{}
	ABI    *abi.ABI
}

func ExecuteMulticall(chainId uint64, multicallAddress common.Address, calls []MulticallCalldata) ([][]byte, error) {
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

	return returnData, nil
}
