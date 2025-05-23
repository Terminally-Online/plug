package utils

import (
	"fmt"
	"net/http"
	"strings"
)

type SolverError struct {
	Message    string `json:"message"`
	Field      string `json:"status,omitempty"`
	Type       string `json:"type,omitempty"`
	StatusCode int    `json:"-"`
}

func (e SolverError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s", e.Field, e.Message)
	}
	return e.Message
}

func (e SolverError) GetStatusCode() int {
	if e.StatusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.StatusCode
}

var (
	ErrEnvironmentNotInitialized = func(error string) SolverError {
		return SolverError{
			Type:    "EnvironmentNotInitialized",
			Message: fmt.Sprintf("loading environment (.env) failed: %s", error),
		}
	}

	ErrEnvironmentVarNotSet = func(envVar string) SolverError {
		return SolverError{
			Type:    "EnvironmentVarNotSet",
			Message: fmt.Sprintf("%s environment variable not set.", envVar),
		}
	}

	ErrEthClient = func(error string) SolverError {
		return SolverError{
			Type:    "EthClient",
			Message: fmt.Sprintf("failed to connect to Ethereum node: %s", error),
		}
	}

	ErrExplorer = func(error string) SolverError {
		return SolverError{
			Type:    "Explorer",
			Message: fmt.Sprintf("failed to connect to Etherscan: %s", error),
		}
	}

	ErrABI = func(contractName string) SolverError {
		return SolverError{
			Type:    "ABI",
			Message: fmt.Sprintf("Contract %s could not be interfaced with.", contractName),
		}
	}

	ErrContract = func(address string) SolverError {
		return SolverError{
			Type:    "Contract",
			Message: fmt.Sprintf("Contract at %s could not be interfaced with.", address),
		}
	}

	ErrBuild = func(error string) SolverError {
		return SolverError{
			Type:    "Build",
			Message: fmt.Sprintf("Building transaction failed: %s", error),
		}
	}

	ErrOptions = func(error string) SolverError {
		return SolverError{
			Type:    "Options",
			Message: fmt.Sprintf("Building options failed: %s", error),
		}
	}

	ErrTransaction = func(error string) SolverError {
		return SolverError{
			Type:    "Transaction",
			Message: fmt.Sprintf("Building transaction failed: %s", error),
		}
	}

	ErrChainId = func(field string, value uint64) SolverError {
		return SolverError{
			Type:       "ChainId",
			Field:      field,
			Message:    fmt.Sprintf("%d is not a valid chainId", value),
			StatusCode: http.StatusBadRequest,
		}
	}

	// note, kind of guessing with most of these error messages. We can remove or cleanup them as we start to see them in the wild
	ErrSimulationFailed = func(error string) SolverError {
		switch {
		case strings.Contains(error, "insufficient funds for gas"):
			return SolverError{
				Type:       "SimulationFailed",
				Message:    "insufficient funds for gas",
				StatusCode: http.StatusConflict,
			}
		case strings.Contains(error, "execution reverted"):
			return SolverError{
				Type:       "SimulationFailed",
				Message:    "execution reverted",
				StatusCode: http.StatusConflict,
			}
		case strings.Contains(error, "out of gas"):
			return SolverError{
				Type:       "SimulationFailed",
				Message:    "transaction ran out of gas",
				StatusCode: http.StatusConflict,
			}
		case strings.Contains(error, "gas limit reached"):
			return SolverError{
				Type:       "SimulationFailed",
				Message:    "gas limit exceeded",
				StatusCode: http.StatusConflict,
			}
		case strings.Contains(error, "insufficient balance"):
			return SolverError{
				Type:       "SimulationFailed",
				Message:    "insufficient balance for transfer",
				StatusCode: http.StatusConflict,
			}
		case strings.Contains(error, "insufficient allowance"):
			return SolverError{
				Type:       "SimulationFailed",
				Message:    "insufficient token allowance",
				StatusCode: http.StatusConflict,
			}
		default:
			return SolverError{
				Type:       "SimulationFailed",
				Message:    fmt.Sprintf("simulation failed: %s", error),
				StatusCode: http.StatusInternalServerError,
			}
		}
	}

	ErrNotImplemented = func(message string) SolverError {
		return SolverError{
			Type:       "NotImplemented",
			Message:    message,
			StatusCode: http.StatusNotImplemented,
		}
	}

	ErrMissingField = func(field string) SolverError {
		return SolverError{
			Type:       "MissingField",
			Field:      field,
			Message:    fmt.Sprintf("request missing required field: %s", field),
			StatusCode: http.StatusBadRequest,
		}
	}

	ErrUnauthorized = func(message string) SolverError {
		return SolverError{
			Type:       "Unauthorized",
			Message:    message,
			StatusCode: http.StatusUnauthorized,
		}
	}

	ErrInvalidRequestBody = func(err error) SolverError {
		return SolverError{
			Type:       "InvalidRequestBody",
			Message:    fmt.Sprintf("invalid request body: %v", err),
			StatusCode: http.StatusBadRequest,
		}
	}

	ErrNotFound = func(message string) SolverError {
		return SolverError{
			Type:       "NotFound",
			Message:    message,
			StatusCode: http.StatusNotFound,
		}
	}

	ErrInvalidField = func(field string, value string) SolverError {
		return SolverError{
			Type:       "InvalidField",
			Field:      field,
			Message:    fmt.Sprintf("invalid value for field %s: %s", field, value),
			StatusCode: http.StatusBadRequest,
		}
	}

	ErrInvalidParameters = func(err error) SolverError {
		return SolverError{
			Type:       "InvalidParameters",
			Message:    fmt.Sprintf("invalid parameters: %v", err),
			StatusCode: http.StatusBadRequest,
		}
	}

	ErrInternal = func(message string) SolverError {
		return SolverError{
			Type:       "InternalServerError",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
		}
	}
)

// Helper function to check if an error is a SolverError and of a specific type
func IsErrorOfType(err error, errorType string) bool {
	if serverErr, ok := err.(SolverError); ok {
		return serverErr.Type == errorType
	}
	return false
}

// Helper function to get the status code from any error
func GetStatusCodeFromError(err error) int {
	if serverErr, ok := err.(SolverError); ok {
		return serverErr.GetStatusCode()
	}
	return http.StatusInternalServerError // Default to 500
}
