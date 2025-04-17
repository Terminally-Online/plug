package actions

import (
	"fmt"
	"solver/internal/coil"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ActionOnchainFunctionInterface interface {
	GetCoilUpdate(string, string, ActionDefinitionInterface) (*coil.Update, error)
	GetCalldata(...any) ([]byte, error)
	GetFunctionName() string
}

type ActionOnchainFunctionResponse struct {
	ActionOnchainFunctionInterface
	Metadata     *bind.MetaData
	FunctionName string
	Arguments    *abi.Arguments
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

func (r *ActionOnchainFunctionResponse) GetInputs(paramToReplace string, linkedCoilKey string, definition ActionDefinitionInterface) (*coil.Update, error) {
	if r.Arguments != nil {
		position, err := coil.GetArgumentsCoilPosition(r.Arguments, &paramToReplace, nil)
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

	return nil, nil
}

func (r *ActionOnchainFunctionResponse) GetOutputs(paramToReplace string, linkedCoilKey string, definition ActionDefinitionInterface) (*coil.Update, error) {
	abi, err := r.Metadata.GetAbi()
	if err != nil {
		return nil, err
	}

	position, err := coil.GetABICoilPosition(abi, r.FunctionName, &paramToReplace, nil)
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

func (r *ActionOnchainFunctionResponse) GetCoilUpdate(paramToReplace string, linkedCoilKey string, definition ActionDefinitionInterface) (*coil.Update, error) {
	if definition == nil {
		return nil, fmt.Errorf("update reference to linked input definition is nil")
	}

	if r.Arguments != nil {
		output, err := r.GetInputs(paramToReplace, linkedCoilKey, definition)
		if err != nil {
			return nil, err
		}

		return output, nil
	}

	if r.Metadata != nil && r.FunctionName != "" {
		output, err := r.GetOutputs(paramToReplace, linkedCoilKey, definition)
		if err != nil {
			return nil, err
		}

		return output, nil
	}

	return nil, fmt.Errorf("neither Arguments nor Metadata+FunctionName provided")
}

func (r *ActionOnchainFunctionResponse) GetFunctionName() string {
	return r.FunctionName
}
