package database

import (
	"encoding/json"
	"fmt"
	"solver/internal/actions"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/common"
)

type SetValueInput struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetValueInput struct {
	Key     string `json:"key"`
	Address string `json:"address"`
}

type RemoveValueInput struct {
	Key string `json:"key"`
}

// Helper function to convert string to bytes32
func stringToBytes32(str string) [32]byte {
	// Ensure the string is not longer than 32 bytes
	if len(str) > 32 {
		str = str[:32]
	}
	
	// Create a byte array and copy the string into it
	var result [32]byte
	copy(result[:], []byte(str))
	return result
}

// HandleSetValue implements the set function from PlugDatabase.sol
func HandleSetValue(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs SetValueInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal set value inputs: %w", err)
	}

	// Local implementation (this will be replaced with contract call after ABI generation)
	// Convert key and value to bytes32
	_ = stringToBytes32(inputs.Key)
	value := stringToBytes32(inputs.Value)

	// For local implementation, we're just creating a plug without actual state changes
	// In the future, this will be a transaction to the PlugDatabase contract
	plug := signature.Plug{
		To:    common.Address{}, // This will be the PlugDatabase contract address
		Data:  []byte{},         // This will be the encoded function call
		Value: nil,
	}

	// In a real implementation, we would return the stored value in the plug's state
	_ = fmt.Sprintf("%x", value) // Use this when we implement state
	
	return []signature.Plug{plug}, nil
}

// HandleGetValue implements the get function from PlugDatabase.sol
func HandleGetValue(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs GetValueInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal get value inputs: %w", err)
	}

	// Local implementation (this will be replaced with contract call after ABI generation)
	// Convert key to bytes32
	_ = stringToBytes32(inputs.Key)
	
	// Normalize address
	_ = common.HexToAddress(inputs.Address)

	// For local implementation, we're just creating a plug without actual state changes
	// In the future, this will be a read-only call to the PlugDatabase contract
	plug := signature.Plug{
		To:    common.Address{}, // This will be the PlugDatabase contract address
		Data:  []byte{},         // This will be the encoded function call
		Value: nil,
	}

	// In a real implementation, we would return the stored value for this address and key
	// For now, just return a placeholder value
	dummyValue := [32]byte{}
	_ = fmt.Sprintf("%x", dummyValue) // Use this when we implement state
	
	return []signature.Plug{plug}, nil
}

// HandleRemoveValue implements the remove function from PlugDatabase.sol
func HandleRemoveValue(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs RemoveValueInput
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal remove value inputs: %w", err)
	}

	// Local implementation (this will be replaced with contract call after ABI generation)
	// Convert key to bytes32
	_ = stringToBytes32(inputs.Key)

	// For local implementation, we're just creating a plug without actual state changes
	// In the future, this will be a transaction to the PlugDatabase contract
	plug := signature.Plug{
		To:    common.Address{}, // This will be the PlugDatabase contract address
		Data:  []byte{},         // This will be the encoded function call
		Value: nil,
	}
	
	return []signature.Plug{plug}, nil
}