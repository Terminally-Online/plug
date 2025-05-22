package simulation

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"solver/bindings/plug_socket"
	"solver/internal/bindings/events"
	"solver/internal/database/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type DecodedOutput struct {
	FunctionName string
	Values       []interface{}
	RawOutput    hexutil.Bytes
}

type DecodedTrace struct {
	Logs   []types.DecodedLog
	Output *DecodedOutput
}

func FindRevertError(trace *Trace) (string, string) {
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
	// TODO: need to do the abi lookup
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

// ExtractAllLogs recursively extracts all logs from a trace or call and their nested calls
func ExtractAllLogs(node interface{}) []types.Log {
	if node == nil {
		return nil
	}

	var logs []types.Log
	var calls []Call

	switch n := node.(type) {
	case *Trace:
		logs = n.Logs
		calls = n.Calls
	case *Call:
		logs = n.Logs
		calls = n.Calls
	default:
		return nil
	}

	// Start with logs from this node
	allLogs := make([]types.Log, len(logs))
	copy(allLogs, logs)

	// Recursively get logs from nested calls
	for _, call := range calls {
		allLogs = append(allLogs, ExtractAllLogs(&call)...)
	}

	return allLogs
}

// DecodeTraceResults decodes all logs and outputs from a trace using the provided ABI provider
func DecodeTraceResults(trace *Trace) (*DecodedTrace, error) {
	if trace == nil {
		return nil, fmt.Errorf("trace is nil")
	}

	// Decode the main trace output if it exists
	var output *DecodedOutput
	if len(trace.Output) > 0 {
		// TODO: Get ABI for the target contract and decode the output
		output = &DecodedOutput{
			RawOutput: trace.Output,
		}
	}

	// Extract all logs from the trace and its calls
	allLogs := ExtractAllLogs(trace)
	decodedLogs := make([]types.DecodedLog, 0, len(allLogs))

	for _, log := range allLogs {
		signature := log.Topics[0].Hex()
		eventDef := events.EventsBySignature[signature]

		decoded := types.DecodedLog{
			Address: log.Address,
			Raw:     log,
		}

		if eventDef != nil {
			decoded.Name = eventDef.Name

			// Pre-allocate slice with correct size to maintain order
			params := make([]types.EventParameter, len(eventDef.Inputs))

			event := abi.NewEvent(eventDef.Name, eventDef.Name, false, eventDef.Inputs)

			// Decode non-indexed args
			args := make(map[string]interface{})
			if err := event.Inputs.UnpackIntoMap(args, log.Data); err != nil {
				fmt.Printf("failed to unpack simulation log data: %v\n", err)
				continue
			}

			// Process all inputs in original ABI order
			topicIndex := 1 // Skip first topic (event signature)
			for i, input := range eventDef.Inputs {
				param := types.EventParameter{
					Name:    input.Name,
					Type:    input.Type.String(),
					Indexed: input.Indexed,
				}

				if input.Indexed {
					if topicIndex < len(log.Topics) {
						param.Value = log.Topics[topicIndex]
						topicIndex++
					}
				} else {
					if v, ok := args[input.Name]; ok {
						param.Value = v
					}
				}
				params[i] = param // Place parameter in correct position
			}
			decoded.Parameters = params
		}

		decodedLogs = append(decodedLogs, decoded)
	}

	return &DecodedTrace{
		Logs:   decodedLogs,
		Output: output,
	}, nil
}
