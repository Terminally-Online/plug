package database

import (
	"encoding/json"
	"fmt"
	"solver/bindings/plug_database"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SetValueInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetValueInput struct {
	Key     string         `json:"key"`
	Address common.Address `json:"address"`
}

type RemoveValueInput struct {
	Key string `json:"key"`
}

func stringNameToKey(str string) [32]byte {
	hash := crypto.Keccak256([]byte(str))

	var result [32]byte
	copy(result[:], hash)
	return result
}

func HandleSetValue(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs SetValueInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal set value inputs: %w", err)
	}

	key := stringNameToKey(inputs.Key)
	value := stringNameToKey(inputs.Value)

	databaseContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["database"])
	databaseAbi, err := plug_database.PlugDatabaseMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugDatabase ABI: %w", err)
	}

	calldata, err := databaseAbi.Pack("set", key, value)
	if err != nil {
		return nil, fmt.Errorf("failed to pack set calldata: %w", err)
	}

	plug := signature.Plug{
		To:    databaseContract,
		Data:  calldata,
		Value: nil,
	}

	return []signature.Plug{plug}, nil
}

func HandleGetValue(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs GetValueInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal get value inputs: %w", err)
	}

	key := stringNameToKey(inputs.Key)

	if inputs.Address == (common.Address{}) {
		return nil, fmt.Errorf("invalid address: address cannot be zero address")
	}

	databaseContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["database"])
	databaseAbi, err := plug_database.PlugDatabaseMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugDatabase ABI: %w", err)
	}

	calldata, err := databaseAbi.Pack("get", inputs.Address, key)
	if err != nil {
		return nil, fmt.Errorf("failed to pack get calldata: %w", err)
	}

	plug := signature.Plug{
		To:    databaseContract,
		Data:  calldata,
		Value: nil,
	}

	return []signature.Plug{plug}, nil
}

func HandleRemoveValue(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs RemoveValueInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal remove value inputs: %w", err)
	}

	key := stringNameToKey(inputs.Key)

	databaseContract := common.HexToAddress(references.Networks[params.ChainId].References["plug"]["database"])
	databaseAbi, err := plug_database.PlugDatabaseMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get PlugDatabase ABI: %w", err)
	}

	calldata, err := databaseAbi.Pack("remove", key)
	if err != nil {
		return nil, fmt.Errorf("failed to pack remove calldata: %w", err)
	}

	plug := signature.Plug{
		To:    databaseContract,
		Data:  calldata,
		Value: nil,
	}

	return []signature.Plug{plug}, nil
}
