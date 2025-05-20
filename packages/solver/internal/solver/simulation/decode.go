package simulation

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"solver/bindings/plug_socket"
)

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
