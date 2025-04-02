package actions

import (
	"fmt"
	"solver/internal/coil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ActionOnchainFunctionInterface interface {
	GetCoilUpdate(string, string, ActionDefinitionInterface) (*coil.Update, error)
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

func (r *ActionOnchainFunctionResponse) GetCoilUpdate(paramToReplace string, linkedCoilKey string, definition ActionDefinitionInterface) (*coil.Update, error) {
	if definition == nil {
		return nil, fmt.Errorf("update reference to linked input definition is nil")
	}

	abi, err := r.Metadata.GetAbi()
	if err != nil {
		return nil, err
	}

	position, err := coil.GetCoilPosition(abi, r.FunctionName, &paramToReplace, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to find coils: %w", err)
	}

	slice, err := definition.GetCoilSlice(linkedCoilKey)
	if err != nil {
		return nil, fmt.Errorf("failed to find coils: %w", err)
	}

	return &coil.Update{
		Start: position,
		Slice: *slice,
	}, nil
}
