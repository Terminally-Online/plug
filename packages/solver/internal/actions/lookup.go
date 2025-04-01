package actions

import (
	"encoding/json"
	"fmt"
	"solver/internal/client"
	"solver/internal/coil"

	"github.com/ethereum/go-ethereum/common"
)

type SchemaLookup[T any] struct {
	ChainId                  uint64
	Client                   *client.Client
	From                     common.Address
	Search                   map[int]string
	Inputs                   *T
	PreviousActionDefinition *ActionDefinition[any]
}

func NewSchemaLookup[T any](chainId uint64, from common.Address, search map[int]string, raw *json.RawMessage, previousActionDefinition *ActionDefinition[any]) (*SchemaLookup[T], error) {
	client, err := client.New(chainId)
	if err != nil {
		return nil, err
	}

	var inputs T
	if raw != nil {
		if err := json.Unmarshal(*raw, &inputs); err != nil {
			return nil, fmt.Errorf("failed to unmarshal deposit inputs: %v", err)
		}
	}

	return &SchemaLookup[T]{
		Client:                   client,
		ChainId:                  chainId,
		From:                     from,
		Search:                   search,
		Inputs:                   &inputs,
		PreviousActionDefinition: previousActionDefinition,
	}, nil
}

type ValueFunc[R any] func() (R, error)
type FunctionResponseInterface interface {
	GetCoilUpdate(string, *SchemaLookup[any]) (*coil.Update, error)
}

func (lookup *SchemaLookup[T]) GetAndUpdate(
	input *coil.CoilInput[T, any],
	valueFunc ValueFunc[any],
	coilFunc FunctionResponseInterface,
	param string,
	updates []coil.Update,
) (any, []coil.Update, error) {
	response, err := valueFunc()
	if err != nil || !input.GetIsLinked() {
		return response, nil, err
	}

	if update, err := coilFunc.GetCoilUpdate(param, lookup); update != nil {
		updates = append(updates, *update)
		return response, updates, err
	} else if err != nil {
		return response, nil, err
	}

	return response, nil, err
}
