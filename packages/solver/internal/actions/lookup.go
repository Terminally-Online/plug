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
	PreviousActionDefinition ActionDefinitionInterface
}

func NewSchemaLookup[T any](chainId uint64, from common.Address, search map[int]string, raw *json.RawMessage, previousActionDefinition ActionDefinitionInterface) (*SchemaLookup[T], error) {
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

func GetAndUpdate[I any, O any](
	input *coil.CoilInput[I, O],
	valueFunc func() (O, error),
	coilFunc ActionOnchainFunctionInterface,
	param string,
	updates []coil.Update,
	definition ActionDefinitionInterface,
) (O, []coil.Update, error) {
	response, err := valueFunc()
	if err != nil || !input.GetIsLinked() || definition == nil {
		return response, nil, err
	}

	// TODO MASON: is there any reason why we have I as any instead of constraining it to a string?
	// It seems to me like the key of a coil input will always be a string to match the function signature of the linked coil
	inValue := fmt.Sprintf("%v", input.GetValue())
	if update, err := coilFunc.GetCoilUpdate(param, inValue, definition); update != nil {
		updates = append(updates, *update)
		return response, updates, err
	} else if err != nil {
		return response, nil, err
	}

	return response, nil, err
}
