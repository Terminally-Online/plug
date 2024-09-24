package utils

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

var (
	ErrInvalidPayload = func(field string, value string) ValidationError {
		return ValidationError{Field: field, Message: fmt.Sprintf("%s is not a valid payload", value)}
	}

	ErrInvalidChainId = func(field string, value int) ValidationError {
		return ValidationError{Field: field, Message: fmt.Sprintf("%d is not a valid chainId", value)}
	}

	ErrInvalidProviderId = func(field string, value int) ValidationError {
		return ValidationError{Field: field, Message: fmt.Sprintf("%d chainId does not have a supported provider", value)}
	}

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

	ErrInvalidArrayLength = func(field string, min *int, max *int) ValidationError {
		if min == nil && max == nil {
			return ValidationError{
				Field:   field,
				Message: "minimum or maximum array length are required",
			}
		}

		var message string
		if min != nil && max != nil {
			message = fmt.Sprintf("not a valid array length between %d and %d", *min, *max)
		} else if min != nil {
			message = fmt.Sprintf("not a valid array length (minimum %d)", *min)
		} else {
			message = fmt.Sprintf("not a valid array length (maximum %d)", *max)
		}
		return ValidationError{Field: field, Message: message}
	}

	ErrMissingField = func(field string) ValidationError {
		return ValidationError{Field: field, Message: "is required"}
	}
)
