package utils

import (
	"fmt"
)

type Error struct {
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("Error: %s", e.Message)
}

var (
	ErrNotImplemented = func(functionName string) ValidationError {
		return ValidationError{Message: fmt.Sprintf("%s is not implemented.", functionName)}
	}

	ErrContractFailed = func(address string) ValidationError {
		return ValidationError{Message: fmt.Sprintf("Contract at %s could not be interfaced with.", address)}
	}

	ErrTransactionFailed = func(functionName string, address string) ValidationError {
		return ValidationError{Message: fmt.Sprintf("%s transaction to %s failed while being built.", functionName, address)}
	}
)
