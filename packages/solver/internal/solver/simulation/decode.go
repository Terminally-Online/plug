package simulation

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"solver/bindings/plug_socket"
)

func FindRevertError(trace *Trace) (string, string) {
	// ... existing code ...
	for _, call := range trace.Calls {
		if call.Error != "" {
			if decodedError := decodeRevertError(call.RawOutput); decodedError != "" {
				return decodedError, call.RawOutput
			}
			return call.Error, call.RawOutput
		}
	}
	return "", ""
}

// TODO MASON: How can I programatically fetch the ABI by contract address in our repo?
func decodeRevertError(output string) string {
	// If output is less than 4 bytes (8 hex chars + "0x"), it's not a valid error
	if len(output) < 10 {
		return "unknown error"
	}

	// Strip "0x" prefix if present
	if output[:2] == "0x" {
		output = output[2:]
	}

	// Convert hex string to bytes
	outputBytes, err := hex.DecodeString(output)
	if err != nil {
		return "invalid hex string"
	}

	// Get the ABI
	plugSocketAbi, err := plug_socket.PlugSocketMetaData.GetAbi()
	if err != nil {
		return "failed to parse error ABI"
	}

	// Try all known errors in the ABI
	selector := outputBytes[:4]
	errorData := outputBytes[4:]

	for name, abiError := range plugSocketAbi.Errors {
		// Compare error selectors
		if bytes.Equal(selector, abiError.ID.Bytes()[:4]) {
			decoded, err := abiError.Inputs.Unpack(errorData)
			if err == nil && len(decoded) > 0 {
				return fmt.Sprintf("%s: %v", name, decoded)
			}
			return name
		}
	}

	// Fallback to string decoding if no matching error found
	if len(errorData) > 32 {
		strLen := new(big.Int).SetBytes(errorData[32:64]).Uint64()
		if strLen > 0 && len(errorData) >= int(64+strLen) {
			return string(errorData[64 : 64+strLen])
		}
	}

	return fmt.Sprintf("error selector: 0x%x", selector)

}

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
