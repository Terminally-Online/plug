package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

// ANSI color codes for terminal output
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

// Standard test address (common for all tests)
var TestAddress = "0x50701f4f523766bFb5C195F93333107d1cB8cD90"

// TestIntent struct matching the expected API format
type TestIntent struct {
	Id      string           `json:"id,omitempty"`
	ChainId uint64           `json:"chainId"`
	From    string           `json:"from,omitempty"`
	Inputs  []map[string]any `json:"inputs,omitempty"`
	Options struct {
		IsEOA    bool `json:"isEOA"`
		Simulate bool `json:"simulate"`
		Submit   bool `json:"submit"`
	} `json:"options,omitempty"`
}

// Helper functions for colored logging
func SuccessLog(t *testing.T, format string, args ...interface{}) {
	t.Logf(ColorGreen+format+ColorReset, args...)
}

func ErrorLog(t *testing.T, format string, args ...interface{}) {
	t.Logf(ColorRed+format+ColorReset, args...)
}

func InfoLog(t *testing.T, format string, args ...interface{}) {
	t.Logf(ColorBlue+format+ColorReset, args...)
}

func WarningLog(t *testing.T, format string, args ...interface{}) {
	t.Logf(ColorYellow+format+ColorReset, args...)
}

// RedText simply wraps text in red for visual highlighting
func RedText(text string) string {
	return ColorRed + text + ColorReset
}

// ErrorEqual is a wrapper around a standard equality check that shows assertion failures in red
func ErrorEqual(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if expected == actual {
		return true
	}
	
	errorMsg := fmt.Sprintf("Expected: %v, got: %v", expected, actual)
	if len(msgAndArgs) > 0 {
		if msg, ok := msgAndArgs[0].(string); ok {
			if len(msgAndArgs) > 1 {
				errorMsg += "\n" + fmt.Sprintf(msg, msgAndArgs[1:]...)
			} else {
				errorMsg += "\n" + msg
			}
		}
	}
	
	t.Error(RedText(errorMsg))
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
	
	t.Error(RedText(errorMsg))
	return false
}

// Console output helpers (for non-test contexts)
func SuccessPrint(format string, args ...interface{}) {
	fmt.Printf(ColorGreen+format+ColorReset+"\n", args...)
}

func ErrorPrint(format string, args ...interface{}) {
	fmt.Printf(ColorRed+format+ColorReset+"\n", args...)
}

func InfoPrint(format string, args ...interface{}) {
	fmt.Printf(ColorBlue+format+ColorReset+"\n", args...)
}

func WarningPrint(format string, args ...interface{}) {
	fmt.Printf(ColorYellow+format+ColorReset+"\n", args...)
}

// Check if database is available by trying to connect
// This function only checks TCP connectivity, not actual database access
func IsDatabaseAvailable() bool {
	// Just check if the database is running - don't actually connect
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	if host == "" || port == "" {
		return false
	}

	// Try to connect to the database server without actually connecting to a specific database
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 1*time.Second)
	if err != nil {
		fmt.Printf("Database connection check failed: %v\n", err)
		return false
	}
	conn.Close()
	return true
}

// Load environment variables needed for testing
func SetupEnvironment() {
	// Attempt to load .env file if it exists - ignore errors
	if _, err := os.Stat(".env"); err == nil {
		// Ignore error - just try to load it
		_ = godotenv.Load()
	}

	// Make sure essential environment variables for tests are set
	// These values are used when running tests in isolation and ensure
	// the tests can run even without a fully configured environment
	envDefaults := map[string]string{
		// Database configuration
		"DATABASE_HOST":     "localhost",
		"DATABASE_USER":     "plug",
		"DATABASE_PASSWORD": "plugdev",
		"DATABASE_NAME":     "plug_solver",
		"DATABASE_PORT":     "6432",
		"DATABASE_SSLMODE":  "disable",

		// API keys and authentication
		"ADMIN_API_KEY": "test-admin-key",
		"TEST_API_KEY":  "testing", // Special testing API key that doesn't get rate limited

		// Encryption (without actually decrypting anything)
		"ENCRYPTION_KEY": "test-encryption-key",

		// Blockchain connections
		"RPC_URL":  "http://localhost:8545", // Default RPC URL for local testing
		"CHAIN_ID": "1",                     // Default to Ethereum mainnet

		// Testing flags
		"GO_TEST_MODE": "true", // Indicates we're in test mode
	}

	for key, defaultValue := range envDefaults {
		if os.Getenv(key) == "" {
			os.Setenv(key, defaultValue)
		}
	}

	// Print a message that we're using test environment values
	fmt.Println("Using test environment configuration")
}

// Minimum delay between requests to avoid overloading the server
var minRequestDelay = 50 * time.Millisecond
var lastRequestTime = time.Now()
var requestMutex sync.Mutex

// Helper function to make test HTTP requests with detailed error reporting
func MakeTestRequest(url, method string, body interface{}) (*http.Response, []byte, error) {
	var reqBody io.Reader
	var jsonData []byte
	var err error

	// Apply a small delay between requests to avoid overwhelming the server
	requestMutex.Lock()
	elapsed := time.Since(lastRequestTime)
	if elapsed < minRequestDelay {
		time.Sleep(minRequestDelay - elapsed)
	}
	lastRequestTime = time.Now()
	requestMutex.Unlock()

	if body != nil {
		jsonData, err = json.Marshal(body)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}

	// Use the dedicated testing API key that doesn't get rate limited
	testApiKey := os.Getenv("TEST_API_KEY")
	if testApiKey == "" {
		testApiKey = "testing" // Default testing API key with no rate limits
	}
	req.Header.Set("X-Api-Key", testApiKey)

	// Add User-Agent header to identify test requests
	req.Header.Set("User-Agent", "Plug-Solver-TestSuite/1.0")

	client := http.Client{
		Timeout: 10 * time.Second, // Set a reasonable timeout
	}

	// Always log the request in test failures
	reqInfo := fmt.Sprintf("%s %s", method, url)
	if body != nil {
		// For brevity in logs, trim long request bodies
		bodyStr := string(jsonData)
		if len(bodyStr) > 500 {
			bodyStr = bodyStr[:500] + "... [truncated]"
		}
		reqInfo += fmt.Sprintf("\nRequest body: %s", bodyStr)
	}

	// Detailed logging in debug mode
	if os.Getenv("TEST_DEBUG") == "true" {
		InfoPrint("REQUEST: %s", reqInfo)
	}

	resp, err := client.Do(req)
	if err != nil {
		errMsg := fmt.Sprintf("Request failed: %s\n%v", reqInfo, err)
		if os.IsTimeout(err) {
			return nil, nil, fmt.Errorf("request timed out: %s", errMsg)
		}
		return nil, nil, fmt.Errorf("failed to execute request: %s", errMsg)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	// Create response info for logging
	respInfo := fmt.Sprintf("Status: %d", resp.StatusCode)
	respBodyStr := string(responseBody)
	if len(respBodyStr) > 1000 {
		respBodyStr = respBodyStr[:1000] + "... [truncated]"
	}
	respInfo += fmt.Sprintf("\nBody: %s", respBodyStr)

	// Detailed logging in debug mode
	if os.Getenv("TEST_DEBUG") == "true" {
		InfoPrint("RESPONSE: %s", respInfo)
	}

	// For all responses (including errors), attach context information to help with debugging
	return &http.Response{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       io.NopCloser(bytes.NewReader(responseBody)),
		Request: &http.Request{
			Method: method,
			URL:    req.URL,
		},
	}, responseBody, nil
}

// Helper function to categorize failures for better debugging
func CategorizeFailure(responseBody string, chainId uint64, protocol, action string) string {
	// Look for common error patterns
	switch {
	case Contains(responseBody, "rate limit"):
		return "RATE_LIMIT_EXCEEDED"
	case Contains(responseBody, "not supported") || Contains(responseBody, "not implemented"):
		return "FEATURE_NOT_SUPPORTED"
	case Contains(responseBody, "not found") || Contains(responseBody, "couldn't find"):
		return "RESOURCE_NOT_FOUND"
	case Contains(responseBody, "chain") && Contains(responseBody, "not supported"):
		return fmt.Sprintf("CHAIN_NOT_SUPPORTED (Chain ID: %d)", chainId)
	case Contains(responseBody, "validation") || Contains(responseBody, "invalid"):
		return "INPUT_VALIDATION_ERROR"
	case Contains(responseBody, "rpc error") || Contains(responseBody, "connection"):
		return "RPC_CONNECTION_ERROR"
	case Contains(responseBody, "contract") && Contains(responseBody, "error"):
		return "CONTRACT_INTERACTION_ERROR"
	case Contains(responseBody, "timeout") || Contains(responseBody, "timed out"):
		return "TIMEOUT"
	default:
		return "UNKNOWN_ERROR"
	}
}

// Helper function to check if a string contains another string
func Contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// Print a summarized test report
func PrintTestSummary(buf *bytes.Buffer) {
	output := buf.String()

	// Print header
	fmt.Printf("\n%s╔════════════════════════════════════════╗%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s║              Test Summary              ║%s\n", ColorBlue, ColorReset)
	fmt.Printf("%s╚════════════════════════════════════════╝%s\n\n", ColorBlue, ColorReset)

	// Extract tested protocols
	protocols := extractProtocolsTested(output)
	fmt.Printf("%sProtocols tested:%s\n", ColorBlue, ColorReset)
	if len(protocols) == 0 {
		fmt.Println("    None detected")
	} else {
		for _, p := range protocols {
			fmt.Printf("    ✓ %s\n", p)
		}
	}

	// Extract actions tested
	actions := extractActionsTested(output)
	fmt.Printf("\n%sActions tested:%s\n", ColorBlue, ColorReset)
	if len(actions) == 0 {
		fmt.Println("    None detected")
	} else {
		for _, a := range actions {
			fmt.Printf("    ✓ %s\n", a)
		}
	}

	// Extract chain-specific statistics
	baseTests, basePassed := countChainTests(output, 8453)
	mainnetTests, mainnetPassed := countChainTests(output, 1)

	fmt.Printf("\n%sChain-specific test statistics:%s\n", ColorBlue, ColorReset)
	if baseTests == 0 && mainnetTests == 0 {
		fmt.Println("    No chain-specific tests detected")
	} else {
		if baseTests > 0 {
			fmt.Printf("    Base (Chain 8453): Ran %d tests, %s%d passed%s\n",
				baseTests, ColorGreen, basePassed, ColorReset)
		}
		if mainnetTests > 0 {
			fmt.Printf("    Ethereum (Chain 1): Ran %d tests, %s%d passed%s\n",
				mainnetTests, ColorGreen, mainnetPassed, ColorReset)
		}
	}

	// Count test results
	passes := countMatches(output, "--- PASS:")
	failures := countMatches(output, "--- FAIL:")
	skips := countMatches(output, "--- SKIP:")

	// Check for panics
	panics := countMatches(output, "panic: ")
	if panics > 0 {
		fmt.Printf("\n%s⚠️  Test execution had %d panic(s). Check the test output for details.%s\n",
			ColorRed, panics, ColorReset)
	}

	// Check for rate limits
	rateLimits := countMatches(output, "rate limit exceeded")
	if rateLimits > 0 {
		fmt.Printf("\n%s⚠️  Warning: %d rate limit exceeded errors detected.%s\n",
			ColorYellow, rateLimits, ColorReset)
		fmt.Println("The test may be hitting the API too frequently. Consider:")
		fmt.Println("  1. Increasing the rate limit in the database for the test API key")
		fmt.Println("  2. Adding delays between API calls")
		fmt.Println("  3. Running fewer tests in parallel")
	}

	// Print test result summary
	fmt.Printf("\n📊 %sPassed: %d%s | %sFailed: %d%s | %sSkipped: %d%s\n",
		ColorGreen, passes, ColorReset,
		ColorRed, failures, ColorReset,
		ColorYellow, skips, ColorReset)
}

// Helper to count occurrences of a pattern in text
func countMatches(text, pattern string) int {
	return strings.Count(text, pattern)
}

// Extract protocols tested from output
func extractProtocolsTested(output string) []string {
	var protocols []string
	matches := make(map[string]bool)

	// Match protocol names that follow a specific pattern
	r := regexp.MustCompile(`Protocol_([a-z0-9_]+)`)
	for _, match := range r.FindAllStringSubmatch(output, -1) {
		if len(match) > 1 {
			matches[match[1]] = true
		}
	}

	for protocol := range matches {
		protocols = append(protocols, protocol)
	}

	sort.Strings(protocols)
	return protocols
}

// Extract actions tested from output
func extractActionsTested(output string) []string {
	var actions []string
	matches := make(map[string]bool)

	// Match action patterns like "Testing protocol=x action=y"
	r := regexp.MustCompile(`Testing protocol=([a-z0-9_]+) action=([a-z0-9_]+)`)
	for _, match := range r.FindAllStringSubmatch(output, -1) {
		if len(match) > 2 {
			actionKey := fmt.Sprintf("%s: %s", match[1], match[2])
			matches[actionKey] = true
		}
	}

	for action := range matches {
		actions = append(actions, action)
	}

	sort.Strings(actions)
	return actions
}

// Count tests for a specific chain
func countChainTests(output string, chainID uint64) (tests int, passed int) {
	// Count test runs
	runPattern := fmt.Sprintf("RUN.*chain_%d", chainID)
	tests = countMatches(output, runPattern)

	// Count passed tests
	passPatterns := []string{
		fmt.Sprintf("PASS:.*chain_%d", chainID),
		fmt.Sprintf("--- PASS: .*chain_%d", chainID),
	}

	for _, pattern := range passPatterns {
		passed += countMatches(output, pattern)
	}

	return tests, passed
}