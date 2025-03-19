package database
//
// import (
// 	"encoding/hex"
// 	"encoding/json"
// 	"fmt"
// 	"math/big"
// 	"solver/bindings/plug_database"
// 	"solver/internal/actions"
// 	"solver/internal/bindings/references"
// 	"solver/internal/solver/signature"
// 	"strings"
//
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/crypto"
// )
//
// type SetInputBase struct {
// 	Key string `json:"key"`
// }
//
// type GetInputBase struct {
// 	Key     string         `json:"key"`
// 	Address common.Address `json:"address"`
// }
//
// type SetUint256Input struct {
// 	SetInputBase
// 	Value *big.Int `json:"value"`
// }
//
// type SetInt256Input struct {
// 	SetInputBase
// 	Value *big.Int `json:"value"`
// }
//
// type SetBytes32Input struct {
// 	SetInputBase
// 	Value [32]byte `json:"value"`
// }
//
// type SetBytesInput struct {
// 	SetInputBase
// 	Value []byte `json:"value"`
// }
//
// type SetAddressInput struct {
// 	SetInputBase
// 	Value common.Address `json:"value"`
// }
//
// type SetBoolInput struct {
// 	SetInputBase
// 	Value bool `json:"value"`
// }
//
// type SetStringInput struct {
// 	SetInputBase
// 	Value string `json:"value"`
// }
//
// type RemoveValueInput struct {
// 	Key string `json:"key"`
// }
//
// type BatchSetInput struct {
// 	Count  *big.Int `json:"count"`
// 	Keys   []string `json:"keys"`
// 	Values []string `json:"values"`
// }
//
// func stringNameToKey(str string) [32]byte {
// 	hash := crypto.Keccak256([]byte(str))
// 	var result [32]byte
// 	copy(result[:], hash)
// 	return result
// }
//
// func getDatabaseContractInfo(chainId uint64) (common.Address, *abi.ABI, error) {
// 	databaseContract := common.HexToAddress(references.Networks[chainId].References["plug"]["database"])
// 	databaseAbi, err := plug_database.PlugDatabaseMetaData.GetAbi()
// 	if err != nil {
// 		return common.Address{}, nil, fmt.Errorf("failed to get PlugDatabase ABI: %w", err)
// 	}
//
// 	return databaseContract, databaseAbi, nil
// }
//
// func HandleSetUint256(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs SetUint256Input
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal set uint256 inputs: %w", err)
// 	}
//
// 	if inputs.Value == nil {
// 		return nil, fmt.Errorf("value cannot be nil")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("setUint", key, inputs.Value)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack setUint calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleSetInt256(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs SetInt256Input
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal set int256 inputs: %w", err)
// 	}
//
// 	if inputs.Value == nil {
// 		return nil, fmt.Errorf("value cannot be nil")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("setInt", key, inputs.Value)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack setInt calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleSetBytes32(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs SetBytes32Input
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		var stringInput struct {
// 			Key   string `json:"key"`
// 			Value string `json:"value"`
// 		}
// 		if err2 := json.Unmarshal(raw, &stringInput); err2 == nil {
// 			valueStr := strings.TrimPrefix(stringInput.Value, "0x")
// 			if len(valueStr) > 64 {
// 				return nil, fmt.Errorf("bytes32 value too long: max 32 bytes (64 hex chars)")
// 			}
//
// 			valueBytes, err := hex.DecodeString(valueStr)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid bytes32 hex string: %w", err)
// 			}
//
// 			var bytes32Value [32]byte
// 			copy(bytes32Value[32-len(valueBytes):], valueBytes)
//
// 			inputs = SetBytes32Input{
// 				SetInputBase: SetInputBase{Key: stringInput.Key},
// 				Value:        bytes32Value,
// 			}
// 		} else {
// 			return nil, fmt.Errorf("failed to unmarshal set bytes32 inputs: %w", err)
// 		}
// 	}
//
// 	key := stringNameToKey(inputs.Key)
// 	value := inputs.Value
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("set", key, value)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack set calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleSetBytes(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs SetBytesInput
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		var stringInput struct {
// 			Key   string `json:"key"`
// 			Value string `json:"value"`
// 		}
// 		if err2 := json.Unmarshal(raw, &stringInput); err2 == nil {
// 			valueStr := strings.TrimPrefix(stringInput.Value, "0x")
// 			valueBytes, err := hex.DecodeString(valueStr)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid bytes hex string: %w", err)
// 			}
//
// 			inputs = SetBytesInput{
// 				SetInputBase: SetInputBase{Key: stringInput.Key},
// 				Value:        valueBytes,
// 			}
// 		} else {
// 			return nil, fmt.Errorf("failed to unmarshal set bytes inputs: %w", err)
// 		}
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("setBytes", key, inputs.Value)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack setBytes calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleSetAddress(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs SetAddressInput
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		var stringInput struct {
// 			Key   string `json:"key"`
// 			Value string `json:"value"`
// 		}
// 		if err2 := json.Unmarshal(raw, &stringInput); err2 == nil {
// 			if !common.IsHexAddress(stringInput.Value) {
// 				return nil, fmt.Errorf("invalid ethereum address: %s", stringInput.Value)
// 			}
//
// 			inputs = SetAddressInput{
// 				SetInputBase: SetInputBase{Key: stringInput.Key},
// 				Value:        common.HexToAddress(stringInput.Value),
// 			}
// 		} else {
// 			return nil, fmt.Errorf("failed to unmarshal set address inputs: %w", err)
// 		}
// 	}
//
// 	if inputs.Value == (common.Address{}) {
// 		return nil, fmt.Errorf("address cannot be zero address")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("setAddress", key, inputs.Value)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack setAddress calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleSetBool(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs SetBoolInput
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal set bool inputs: %w", err)
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("setBool", key, inputs.Value)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack setBool calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleSetString(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs SetStringInput
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal set string inputs: %w", err)
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("setString", key, inputs.Value)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack setString calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleGetUint256(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs GetInputBase
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal get uint256 inputs: %w", err)
// 	}
//
// 	if inputs.Address == (common.Address{}) {
// 		return nil, fmt.Errorf("invalid address: address cannot be zero address")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("getUint", inputs.Address, key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack getUint calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleGetInt256(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs GetInputBase
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal get int256 inputs: %w", err)
// 	}
//
// 	if inputs.Address == (common.Address{}) {
// 		return nil, fmt.Errorf("invalid address: address cannot be zero address")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("getInt", inputs.Address, key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack getInt calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleGetBytes32(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs GetInputBase
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal get bytes32 inputs: %w", err)
// 	}
//
// 	if inputs.Address == (common.Address{}) {
// 		return nil, fmt.Errorf("invalid address: address cannot be zero address")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("get", inputs.Address, key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack get calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleGetBytes(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs GetInputBase
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal get bytes inputs: %w", err)
// 	}
//
// 	if inputs.Address == (common.Address{}) {
// 		return nil, fmt.Errorf("invalid address: address cannot be zero address")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("getBytes", inputs.Address, key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack getBytes calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleGetAddress(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs GetInputBase
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal get address inputs: %w", err)
// 	}
//
// 	if inputs.Address == (common.Address{}) {
// 		return nil, fmt.Errorf("invalid address: address cannot be zero address")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("getAddress", inputs.Address, key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack getAddress calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleGetBool(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs GetInputBase
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal get bool inputs: %w", err)
// 	}
//
// 	if inputs.Address == (common.Address{}) {
// 		return nil, fmt.Errorf("invalid address: address cannot be zero address")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("getBool", inputs.Address, key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack getBool calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleGetString(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs GetInputBase
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal get string inputs: %w", err)
// 	}
//
// 	if inputs.Address == (common.Address{}) {
// 		return nil, fmt.Errorf("invalid address: address cannot be zero address")
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("getString", inputs.Address, key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack getString calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleRemoveValue(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs RemoveValueInput
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal remove value inputs: %w", err)
// 	}
//
// 	key := stringNameToKey(inputs.Key)
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("remove", key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack remove calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
//
// func HandleBatchSet(lookup *actions.SchemaLookup, raw json.RawMessage) ([]signature.Plug, error) {
// 	var inputs BatchSetInput
// 	if err := json.Unmarshal(raw, &inputs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal batch set inputs: %w", err)
// 	}
//
// 	if inputs.Count == nil || len(inputs.Keys) == 0 || len(inputs.Values) == 0 {
// 		return nil, fmt.Errorf("batch set requires count, keys, and values")
// 	}
//
// 	if inputs.Count.Cmp(big.NewInt(int64(len(inputs.Keys)))) != 0 ||
// 		len(inputs.Keys) != len(inputs.Values) {
// 		return nil, fmt.Errorf("count must match the number of keys and values")
// 	}
//
// 	keys := make([][32]byte, len(inputs.Keys))
// 	for i, key := range inputs.Keys {
// 		keys[i] = stringNameToKey(key)
// 	}
//
// 	values := make([][32]byte, len(inputs.Values))
// 	for i, value := range inputs.Values {
// 		valueStr := strings.TrimPrefix(value, "0x")
// 		if len(valueStr) > 64 {
// 			return nil, fmt.Errorf("bytes32 value too long at index %d: max 32 bytes (64 hex chars)", i)
// 		}
//
// 		var bytes32Value [32]byte
// 		if len(valueStr) > 0 {
// 			valueBytes, err := hex.DecodeString(valueStr)
// 			if err != nil {
// 				return nil, fmt.Errorf("invalid bytes32 hex string at index %d: %w", i, err)
// 			}
// 			copy(bytes32Value[32-len(valueBytes):], valueBytes)
// 		}
// 		values[i] = bytes32Value
// 	}
//
// 	databaseContract, databaseAbi, err := getDatabaseContractInfo(lookup.ChainId)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	calldata, err := databaseAbi.Pack("batchSet", keys, values, uint8(1))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to pack batchSet calldata: %w", err)
// 	}
//
// 	plug := signature.Plug{
// 		To:    databaseContract,
// 		Data:  calldata,
// 		Value: nil,
// 	}
//
// 	return []signature.Plug{plug}, nil
// }
