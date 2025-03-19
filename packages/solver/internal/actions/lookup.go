package actions

import (
	"encoding/json"
	"fmt"
	"solver/internal/client"

	"github.com/ethereum/go-ethereum/common"
)

type SchemaLookup[T any] struct {
	ChainId uint64
	Client  *client.Client
	From    common.Address
	Search  map[int]string
	Inputs  *T
}

func NewSchemaLookup[T any](chainId uint64, from common.Address, search map[int]string, raw *json.RawMessage) (*SchemaLookup[T], error) {
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
		Client:  client,
		ChainId: chainId,
		From:    from,
		Search:  search,
		Inputs:  &inputs,
	}, nil
}
