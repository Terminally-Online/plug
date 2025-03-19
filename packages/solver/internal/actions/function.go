package actions

import (
	"solver/internal/coil"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)


type ActionOnchainFunctionInterface interface {
	GetUpdate(string) (*coil.Update, error)
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

func (r *ActionOnchainFunctionResponse) GetUpdate(param string) (*coil.Update, error) {
	// TODO: Need to implement this for real
	return nil, nil
}
