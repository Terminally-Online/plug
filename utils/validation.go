package utils

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s %s", e.Field, e.Message)
}

var (
	ErrInvalidAddress = func(field string, value string) ValidationError {
		return ValidationError{Field: field, Message: fmt.Sprintf("%s is not a valid Ethereum address", value)}
	}

	ErrInvalidHex = func(field string, value string) ValidationError {
		return ValidationError{Field: field, Message: fmt.Sprintf("%s is not a valid hex value", value)}
	}

	ErrInvalidUint = func(field string, value string, size int) ValidationError {
		return ValidationError{Field: field, Message: fmt.Sprintf("%s is not a valid uint%d value", value, size)}
	}

	ErrInvalidBytes = func(field string, value string, size int) ValidationError {
		return ValidationError{Field: field, Message: fmt.Sprintf("%s is not a valid bytes%d value", value, size)}
	}

	ErrMissingField = func(field string) ValidationError {
		return ValidationError{Field: field, Message: "is required"}
	}
)
