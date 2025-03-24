package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// CheckNoError checks if error is nil and fails test with visible formatting if not
// Use instead of require.NoError
func CheckNoError(t *testing.T, err error, msgAndArgs ...interface{}) bool {
	if err == nil {
		return true
	}

	// Get file and line information for the error
	_, file, line, _ := runtime.Caller(1)

	// Format error message
	var msg string
	if len(msgAndArgs) > 0 {
		if fmtMsg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				msg = fmt.Sprintf(fmtMsg, msgAndArgs[1:]...) + ": "
			} else {
				msg = fmtMsg + ": "
			}
		}
	}

	// Build error message with file/line info
	errorPrefix := fmt.Sprintf("%s❌ ERROR [%s:%d] %s: ", ColorRed, filepath.Base(file), line, t.Name())

	// Log the error and fail the test
	t.Errorf("%s%s%v%s", errorPrefix, msg, err, ColorReset)

	return false
}

// CheckEqual checks equality with visible failure message
// Use instead of require.Equal
func CheckEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if expected == actual {
		return true
	}

	// Get file and line information
	_, file, line, _ := runtime.Caller(1)

	// Format message
	var msg string
	if len(msgAndArgs) > 0 {
		if fmtMsg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				msg = fmt.Sprintf(fmtMsg, msgAndArgs[1:]...) + ": "
			} else {
				msg = fmtMsg + ": "
			}
		}
	}

	// Build error message
	errorPrefix := fmt.Sprintf("%s❌ FAILED [%s:%d] %s: ", ColorRed, filepath.Base(file), line, t.Name())
	errorMsg := fmt.Sprintf("%sExpected: %v, got: %v", msg, expected, actual)

	// Log and fail
	t.Errorf("%s%s%s", errorPrefix, errorMsg, ColorReset)

	return false
}

// ErrorContains checks if a container contains an item and shows failure in red
func ErrorContains(t *testing.T, container, item interface{}, msgAndArgs ...interface{}) bool {
	// Check if container contains item based on type
	contains := false

	if container == nil {
		contains = false
	} else if str, ok := container.(string); ok {
		if itemStr, ok := item.(string); ok {
			contains = strings.Contains(str, itemStr)
		}
	} else if m, ok := container.(map[string]interface{}); ok {
		if key, ok := item.(string); ok {
			_, contains = m[key]
		}
	}

	if contains {
		return true
	}

	errorMsg := fmt.Sprintf("Expected %v to contain %v", container, item)
	if len(msgAndArgs) > 0 {
		if msg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				errorMsg += "\n" + fmt.Sprintf(msg, msgAndArgs[1:]...)
			} else {
				errorMsg += "\n" + msg
			}
		}
	}

	t.Error(ColorRed + "FAILED ASSERTION: " + errorMsg + ColorReset)
	return false
}

// FatalOnError fails the test immediately if error is not nil
// Use instead of t.Fatal or require.NoError when you need to stop the test
func FatalOnError(t *testing.T, err error, msgAndArgs ...interface{}) {
	if err == nil {
		return
	}

	// Get file and line information
	_, file, line, _ := runtime.Caller(1)

	// Format message
	var msg string
	if len(msgAndArgs) > 0 {
		if fmtMsg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				msg = fmt.Sprintf(fmtMsg, msgAndArgs[1:]...) + ": "
			} else {
				msg = fmtMsg + ": "
			}
		}
	}

	// Build error message
	errorPrefix := fmt.Sprintf("%s❌ FATAL ERROR [%s:%d] %s: ", ColorRed, filepath.Base(file), line, t.Name())

	// Log and fail immediately
	t.Fatalf("%s%s%v%s", errorPrefix, msg, err, ColorReset)
}
