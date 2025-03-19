package actions

import (
	"fmt"
	"solver/internal/coil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ActionOnchainFunctionInterface interface {
	GetCoilUpdate(string) (*coil.Update, error)
}

type ActionOnchainFunctionResponse struct {
	ActionOnchainFunctionInterface
	Metadata     *bind.MetaData
	FunctionName string
}

func (r *ActionOnchainFunctionResponse) GetCalldata(inputs ...any) ([]byte, error) {
	abi, err := r.Metadata.GetAbi()
	if err != nil {
		return nil, err
	}

	calldata, err := abi.Pack(r.FunctionName, inputs...)
	if err != nil {
		return nil, err
	}

	return calldata, nil
}

func (r *ActionOnchainFunctionResponse) GetCoilUpdate(param string) (*coil.Update, error) {
	abi, err := r.Metadata.GetAbi()
	if err != nil {
		return nil, err
	}

	slices, err := coil.GetCoilSlices(abi, r.FunctionName, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find coils: %w", err)
	}

	position, err := coil.GetCoilPosition(abi, r.FunctionName, &param, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find coils: %w", err)
	}

	return &coil.Update{
		Start: position,
		Slice: slices[0],
	}, nil
}
