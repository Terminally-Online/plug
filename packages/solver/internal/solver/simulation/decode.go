package simulation

// import (
// 	"bytes"
// 	"fmt"
// 	"strings"

// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/common/hexutil"
// )

// // Store ABIs for contracts you interact with
// type TraceDecoder struct {
// 	abis map[string]*abi.ABI // contract address -> ABI
// }

// func NewTraceDecoder(abiJSON map[string]string) (*TraceDecoder, error) {
// 	decoder := &TraceDecoder{
// 		abis: make(map[string]*abi.ABI),
// 	}

// 	for addr, jsonStr := range abiJSON {
// 		parsedABI, err := abi.JSON(strings.NewReader(jsonStr))
// 		if err != nil {
// 			return nil, err
// 		}
// 		decoder.abis[addr] = &parsedABI
// 	}

// 	return decoder, nil
// }

// func (d *TraceDecoder) DecodeOutput(contractAddr string, methodName string, output hexutil.Bytes) (map[string]interface{}, error) {
// 	if abi, ok := d.abis[contractAddr]; ok {
// 		method, ok := abi.Methods[methodName]
// 		if !ok {
// 			return nil, fmt.Errorf("method %s not found", methodName)
// 		}

// 		decoded := make(map[string]interface{})
// 		if err := method.Outputs.UnpackIntoMap(decoded, output); err != nil {
// 			return nil, err
// 		}
// 		return decoded, nil
// 	}
// 	return nil, fmt.Errorf("ABI not found for contract %s", contractAddr)
// }

// func DecodeTrace(trace *CallTrace, decoder *TraceDecoder) (map[string]interface{}, error) {
// 	result := make(map[string]interface{})

// 	// Decode function call/return
// 	if len(trace.Input) >= 4 { // Has function selector
// 		// Get method signature from input
// 		selector := trace.Input[:4]
// 		method, err := decoder.GetMethodFromSelector(trace.To.Hex(), selector)
// 		if err == nil {
// 			// Decode input parameters
// 			inputs, _ := method.Inputs.Unpack(trace.Input[4:])
// 			result["inputs"] = inputs

// 			// Decode output if successful
// 			if trace.Error == "" && len(trace.Output) > 0 {
// 				outputs, _ := method.Outputs.Unpack(trace.Output)
// 				result["outputs"] = outputs
// 			}
// 		}
// 	}

// 	// Decode revert reason
// 	if trace.Error != "" {
// 		// Standard revert reason format
// 		if len(trace.Output) >= 4 && bytes.Equal(trace.Output[:4], []byte{0x08, 0xc3, 0x79, 0xa0}) {
// 			reason, _ := abi.UnpackRevert(trace.Output)
// 			result["revertReason"] = reason
// 		}
// 	}

// 	// Decode events
// 	var events []map[string]interface{}
// 	for _, log := range trace.Logs {
// 		if decoded, err := decoder.DecodeLog(log.Address.Hex(), log.Topics, log.Data); err == nil {
// 			events = append(events, decoded)
// 		}
// 	}
// 	if len(events) > 0 {
// 		result["events"] = events
// 	}

// 	// Recursively decode internal calls
// 	if len(trace.Calls) > 0 {
// 		var internalCalls []map[string]interface{}
// 		for _, call := range trace.Calls {
// 			if decoded, err := DecodeTrace(&call, decoder); err == nil {
// 				internalCalls = append(internalCalls, decoded)
// 			}
// 		}
// 		result["internalCalls"] = internalCalls
// 	}

// 	return result, nil
// }

// // Helper to decode events
// func (d *TraceDecoder) DecodeLog(contractAddr string, topics []common.Hash, data []byte) (map[string]interface{}, error) {
// 	abi := d.abis[contractAddr]
// 	if abi == nil {
// 		return nil, fmt.Errorf("ABI not found")
// 	}

// 	// Find event definition matching topic[0] (event signature)
// 	event, err := abi.EventByID(topics[0])
// 	if err != nil {
// 		return nil, err
// 	}

// 	decoded := make(map[string]interface{})
// 	if err := abi.UnpackIntoMap(decoded, event.Name, data); err != nil {
// 		return nil, err
// 	}

// 	// Add indexed parameters from topics
// 	for i, arg := range event.Inputs {
// 		if arg.Indexed {
// 			decoded[arg.Name] = topics[i+1]
// 		}
// 	}

// 	return decoded, nil
// }
