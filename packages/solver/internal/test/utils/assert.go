package utils

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

// ANSI color codes for test output (duplicated from utils.go for independence)
const (
	testColorReset  = "\033[0m"
	testColorRed    = "\033[31m"
	testColorGreen  = "\033[32m"
	testColorYellow = "\033[33m"
	testColorBlue   = "\033[34m"
)

// RequireEqual checks that expected and actual are equal, fails test if not
// Colored alternative to require.Equal
func RequireEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	if expected == actual {
		return
	}

	// Get file and line information
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(msgAndArgs) > 0 {
		if fmtMsg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				msg = fmt.Sprintf(fmtMsg, msgAndArgs[1:]...)
			} else {
				msg = fmtMsg
			}
		}
	}

	// Build the error message
	errorPrefix := fmt.Sprintf("%s❌ REQUIRE FAILED [%s:%d] %s: ", testColorRed, filepath.Base(file), line, t.Name())
	errorMsg := fmt.Sprintf("Expected: %v, got: %v", expected, actual)
	if msg != "" {
		errorMsg = msg + "\n" + errorMsg
	}

	// Log the error and fail the test
	t.Fatalf("%s%s%s", errorPrefix, errorMsg, testColorReset)
}

// RequireNoError checks if err is nil, fails test with visible formatting if not
// Use instead of require.NoError
func RequireNoError(t *testing.T, err error, msgAndArgs ...interface{}) {
	if err == nil {
		return
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
	errorPrefix := fmt.Sprintf("%s❌ ERROR [%s:%d] %s: ", testColorRed, filepath.Base(file), line, t.Name())

	// Log the error and fail the test
	t.Fatalf("%s%s%v%s", errorPrefix, msg, err, testColorReset)
}

// RequireTrue checks that the condition is true, fails test if not
// Colored alternative to require.True
func RequireTrue(t *testing.T, condition bool, msgAndArgs ...interface{}) {
	if condition {
		return
	}

	// Get file and line information
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(msgAndArgs) > 0 {
		if fmtMsg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				msg = fmt.Sprintf(fmtMsg, msgAndArgs[1:]...)
			} else {
				msg = fmtMsg
			}
		}
	} else {
		msg = "Expected condition to be true, but was false"
	}

	// Build the error message
	errorPrefix := fmt.Sprintf("%s❌ REQUIRE FAILED [%s:%d] %s: ", testColorRed, filepath.Base(file), line, t.Name())

	// Log the error and fail the test
	t.Fatalf("%s%s%s", errorPrefix, msg, testColorReset)
}

// RequireNotEmpty checks that the value is not empty (string, slice, map, etc)
// Colored alternative to require.NotEmpty
func RequireNotEmpty(t *testing.T, value interface{}, msgAndArgs ...interface{}) {
	empty := false

	// Check different types of emptiness
	switch v := value.(type) {
	case string:
		empty = v == ""
	case nil:
		empty = true
	default:
		// Use reflection for other types (slices, maps, etc)
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Slice, reflect.Map, reflect.Chan:
			empty = rv.Len() == 0
		case reflect.Ptr:
			empty = rv.IsNil()
		}
	}

	if !empty {
		return
	}

	// Get file and line information
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(msgAndArgs) > 0 {
		if fmtMsg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				msg = fmt.Sprintf(fmtMsg, msgAndArgs[1:]...)
			} else {
				msg = fmtMsg
			}
		}
	} else {
		msg = "Expected value to not be empty"
	}

	// Build the error message
	errorPrefix := fmt.Sprintf("%s❌ REQUIRE FAILED [%s:%d] %s: ", testColorRed, filepath.Base(file), line, t.Name())

	// Log the error and fail the test
	t.Fatalf("%s%s%s", errorPrefix, msg, testColorReset)
}

// RequireGreater checks that actual is greater than expected
// Colored alternative to require.Greater
func RequireGreater(t *testing.T, actual, expected interface{}, msgAndArgs ...interface{}) {
	greater := false

	// Compare values based on their types
	switch a := actual.(type) {
	case int:
		if e, ok := expected.(int); ok {
			greater = a > e
		}
	case int64:
		if e, ok := expected.(int64); ok {
			greater = a > e
		}
	case float64:
		if e, ok := expected.(float64); ok {
			greater = a > e
		}
	case uint64:
		if e, ok := expected.(uint64); ok {
			greater = a > e
		}
	case time.Time:
		if e, ok := expected.(time.Time); ok {
			greater = a.After(e)
		}
	}

	if greater {
		return
	}

	// Get file and line information
	_, file, line, _ := runtime.Caller(1)

	var msg string
	if len(msgAndArgs) > 0 {
		if fmtMsg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				msg = fmt.Sprintf(fmtMsg, msgAndArgs[1:]...)
			} else {
				msg = fmtMsg
			}
		}
	} else {
		msg = fmt.Sprintf("Expected %v to be greater than %v", actual, expected)
	}

	// Build the error message
	errorPrefix := fmt.Sprintf("%s❌ REQUIRE FAILED [%s:%d] %s: ", testColorRed, filepath.Base(file), line, t.Name())

	// Log the error and fail the test
	t.Fatalf("%s%s%s", errorPrefix, msg, testColorReset)
}

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
