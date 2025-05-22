package simulation

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"solver/bindings/plug_socket"
	"solver/internal/bindings/events"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type DecodedOutput struct {
	FunctionName string
	Values       []interface{}
	RawOutput    hexutil.Bytes
}

type DecodedTrace struct {
	Logs   []DecodedLog
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
func ExtractAllLogs(node interface{}) []Log {
	if node == nil {
		return nil
	}

	var logs []Log
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
	allLogs := make([]Log, len(logs))
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
	decodedLogs := make([]DecodedLog, 0, len(allLogs))

	// Decode each log using the ABI provider
	for _, log := range allLogs {
		signature := log.Topics[0].Hex()
		eventDef := events.EventsBySignature[signature]

		decoded := DecodedLog{
			Address: log.Address,
			Raw:     log,
		}

		if eventDef != nil {
			// We found a matching event definition
			decoded.Name = &eventDef.Name

			// Decode the arguments
			if args, err := DecodeLog(&log); err == nil {
				fmt.Printf("Decoded log: %+v\n", args)
				// decoded.Parameters = args
			}
		}

		decodedLogs = append(decodedLogs, decoded)
	}

	return &DecodedTrace{
		Logs:   decodedLogs,
		Output: output,
	}, nil
}

func DecodeLog(log *Log) (map[string]interface{}, error) {
	if len(log.Topics) == 0 {
		return nil, fmt.Errorf("log has no topics")
	}

	signature := log.Topics[0].Hex()
	def, exists := events.EventsBySignature[signature]
	if !exists {
		return nil, fmt.Errorf("unknown event signature: %s", signature)
	}

	event := abi.Event{
		Name:   def.Name,
		Inputs: def.Inputs,
	}

	decoded := make(map[string]interface{})

	// Handle non-indexed parameters (data)
	if len(log.Data) > 0 {
		if err := event.Inputs.UnpackIntoMap(decoded, log.Data); err != nil {
			return nil, fmt.Errorf("unpacking data: %w", err)
		}
	}

	// Handle indexed parameters (topics)
	indexedInputs := 0
	for _, input := range def.Inputs {
		if input.Indexed {
			if indexedInputs+1 < len(log.Topics) { // +1 because first topic is signature
				decoded[input.Name] = log.Topics[indexedInputs+1] // +1 to skip signature
				indexedInputs++
			}
		}
	}

	return decoded, nil
}
