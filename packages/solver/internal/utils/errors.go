package utils

import (
	"fmt"
)

type ServerError struct {
	Message string `json:"message"`
	Field   string `json:"status,omitempty"`
}

func (e ServerError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s", e.Field, e.Message)
	}
	return e.Message
}

var (
	ErrEnvironmentNotInitialized = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("loading environment (.env) failed: %s", error)}
	}

	ErrEnvironmentVarNotSet = func(envVar string) ServerError {
		return ServerError{Message: fmt.Sprintf("%s environment variable not set.", envVar)}
	}

	ErrEthClient = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("failed to connect to Ethereum node: %s", error)}
	}

	ErrExplorer = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("failed to connect to Etherscan: %s", error)}
	}

	ErrABI = func(contractName string) ServerError {
		return ServerError{Message: fmt.Sprintf("Contract %s could not be interfaced with.", contractName)}
	}

	ErrContract = func(address string) ServerError {
		return ServerError{Message: fmt.Sprintf("Contract at %s could not be interfaced with.", address)}
	}

	ErrBuild = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("Building transaction failed: %s", error)}
	}

	ErrOptions = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("Building options failed: %s", error)}
	}

	ErrTransaction = func(error string) ServerError {
		return ServerError{Message: fmt.Sprintf("Building transaction failed: %s", error)}
	}

	ErrField = func(field string, message string) ServerError {
		return ServerError{Field: field, Message: message}
	}

	ErrChainId = func(field string, value uint64) ServerError {
		return ServerError{Field: field, Message: fmt.Sprintf("%d is not a valid chainId", value)}
	}

	ErrNotImplemented = func(message string) ServerError { 
		return ServerError{Message: message }
	}
)
