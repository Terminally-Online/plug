package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ServerError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Details string `json:"details"`
}

func (e ServerError) Error() string {
	return e.Message
}

func Error(w http.ResponseWriter, error error, statusCode int) {
	w.WriteHeader(statusCode)

	err := ServerError{
		Status:  statusCode,
		Details: error.Error(),
	}

	switch error.(type) {
	case ValidationError:
		err.Message = "Validation error"
	case *json.SyntaxError, *json.UnmarshalTypeError:
		err.Message = "Invalid JSON"
	default:
		if statusCode >= 500 {
			err.Message = "Internal server error"
		} else {
			err.Message = "Request error"
		}
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var (
	ErrNotImplemented = func(functionName string) ServerError {
		return ServerError{Message: fmt.Sprintf("%s is not implemented.", functionName)}
	}

	ErrEnvironmentNotInitialized = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("loading environment (.env) failed: %s", error)}
	}

	ErrEnvironmentVarNotSet = func(envVar string) ServerError {
		return ServerError{Message: fmt.Sprintf("%s environment variable not set.", envVar)}
	}

	ErrEthClientFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("failed to connect to Ethereum node: %s", error)}
	}

	ErrNetworkRequestFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("connecting to downstream provider failed: %s", error)}
	}

	ErrAPIRequestFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("downstream provider request failed: %s", error)}
	}

	ErrExplorerFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("failed to connect to Etherscan: %s", error)}
	}

	ErrABIFailed = func(contractName string) ServerError {
		return ServerError{Message: fmt.Sprintf("Contract %s could not be interfaced with.", contractName)}
	}

	ErrContractFailed = func(address string) ServerError {
		return ServerError{Message: fmt.Sprintf("Contract at %s could not be interfaced with.", address)}
	}

	ErrCallFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("Calling contract failed: %s", error)}
	}

	ErrBuildFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("Building transaction failed: %s", error)}
	}

	ErrOptionsFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("Building options failed: %s", error)}
	}

	ErrTransactionFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("Building transaction failed: %s", error)}
	}

	ErrHttpFailed = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("Building transaction failed: %s", error)}
	}
)
