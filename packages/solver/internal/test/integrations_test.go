package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"solver/internal/test/utils"

	"github.com/joho/godotenv"
)

// ANSI color codes for terminal output
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

// Helper functions for colored logging
func successLog(t *testing.T, format string, args ...interface{}) {
	t.Logf(colorGreen+format+colorReset, args...)
}

func errorLog(t *testing.T, format string, args ...interface{}) {
	t.Logf(colorRed+format+colorReset, args...)
}

func infoLog(t *testing.T, format string, args ...interface{}) {
	t.Logf(colorBlue+format+colorReset, args...)
}

func warningLog(t *testing.T, format string, args ...interface{}) {
	t.Logf(colorYellow+format+colorReset, args...)
}

// Console output helpers (for non-test contexts)
func successPrint(format string, args ...interface{}) {
	fmt.Printf(colorGreen+format+colorReset+"\n", args...)
}

func errorPrint(format string, args ...interface{}) {
	fmt.Printf(colorRed+format+colorReset+"\n", args...)
}

func infoPrint(format string, args ...interface{}) {
	fmt.Printf(colorBlue+format+colorReset+"\n", args...)
}

func warningPrint(format string, args ...interface{}) {
	fmt.Printf(colorYellow+format+colorReset+"\n", args...)
}

var testAddress = "0x50701f4f523766bFb5C195F93333107d1cB8cD90"

// Test intent struct matching the expected API format
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

// Test case definition for protocol tests
type TestCase struct {
	Name     string     `json:"name"`
	Intent   TestIntent `json:"intent"`
	ExpectOk bool       `json:"expectOk"`
}

// Check if database is available by trying to connect
// This function only checks TCP connectivity, not actual database access
func isDatabaseAvailable() bool {
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

// TestMain sets up common test environment and provides a test summary at the end
func TestMain(m *testing.M) {
	// Load environment from .env file if available and set testing defaults
	setupEnvironment()

	// Check if server is available
	serverRunning := false
	_, err := http.Get("http://localhost:8080/health")
	if err == nil {
		serverRunning = true
	}

	// Get database availability status
	dbAvailable := isDatabaseAvailable()

	// Print status with colors
	if !serverRunning {
		errorPrint("❌ ERROR: Solver server is not running on localhost:8080 - tests will fail")
	} else {
		successPrint("✓ Solver server is available")
	}

	if !dbAvailable {
		errorPrint("❌ ERROR: Database is not available - tests will fail")
	} else {
		successPrint("✓ Database is available")
	}

	// Check for specific chain ID filter
	chainID := os.Getenv("TEST_CHAIN_ID")

	if chainID != "" {
		infoPrint("Focusing tests on chain ID: %s", chainID)
	}

	// Mark that we're in test mode
	os.Setenv("GO_TEST_MODE", "true")

	// Track test statistics
	var testStats testStatistics

	// Create a buffer to capture test output
	var buf bytes.Buffer
	writer := io.MultiWriter(os.Stdout, &buf)

	// Redirect test output
	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run tests
	code := m.Run()

	// Close pipe and restore stdout
	w.Close()
	io.Copy(writer, r)
	os.Stdout = origStdout

	// Generate and print test summary
	printTestSummary(&buf, &testStats)

	// Exit with test status code
	os.Exit(code)
}

// Load environment variables needed for testing
func setupEnvironment() {
	// Find the root solver directory
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	rootDir := filepath.Join(dir, "../../")

	// Attempt to load .env file if it exists - ignore errors
	envPath := filepath.Join(rootDir, ".env")
	if _, err := os.Stat(envPath); err == nil {
		// Ignore error - just try to load it
		_ = godotenv.Load(envPath)
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
func makeTestRequest(url, method string, body interface{}) (*http.Response, []byte, error) {
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
		infoPrint("REQUEST: %s", reqInfo)
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
		infoPrint("RESPONSE: %s", respInfo)
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

// TestHealthEndpoint tests the /health endpoint
func TestHealthEndpoint(t *testing.T) {
	// Try health endpoint
	resp, body, err := makeTestRequest("http://localhost:8080/health", http.MethodGet, nil)
	if err != nil {
		errorLog(t, "ERROR: %v", err)
		t.Fatalf("Health endpoint test failed: server is not running or encountered an error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		errorLog(t, "Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))
		t.Errorf("Health endpoint returned unexpected status code: %d", resp.StatusCode)
		return
	}

	var healthResponse map[string]interface{}
	err = json.Unmarshal(body, &healthResponse)
	if err != nil {
		errorLog(t, "Failed to unmarshal response: %v. Response body: %s", err, string(body))
		t.Errorf("Failed to parse health response: %v", err)
		return
	}

	// Check for "status" field - accept both "ok" or "healthy" values
	status, ok := healthResponse["status"].(string)
	if !ok {
		errorLog(t, "Missing or invalid 'status' field in health response: %v", healthResponse)
		t.Errorf("Missing or invalid 'status' field in health response")
		return
	}

	if status != "ok" && status != "healthy" {
		errorLog(t, "Expected status 'ok' or 'healthy', got '%v'", status)
		t.Errorf("Health endpoint returned unexpected status: %s", status)
	} else {
		successLog(t, "Health endpoint returned status: %s", status)
	}
}

// TestGetSchemaEndpoint tests the /solver endpoint for schema retrieval
func TestGetSchemaEndpoint(t *testing.T) {
	// Make the request
	resp, body, err := makeTestRequest("http://localhost:8080/solver?chainId=1", http.MethodGet, nil)
	if err != nil {
		errorLog(t, "ERROR: %v", err)
		t.Fatalf("Schema endpoint test failed: %v", err)
	}

	utils.CheckEqual(t, http.StatusOK, resp.StatusCode)

	var schemaResponse map[string]interface{}
	err = json.Unmarshal(body, &schemaResponse)
	utils.RequireNoError(t, err)

	// Check that we have some protocols returned
	utils.RequireGreater(t, len(schemaResponse), 0)
	successLog(t, "Successfully retrieved schema with %d protocols", len(schemaResponse))
}

// TestGetSchemaForProtocol tests the /solver endpoint for a specific protocol
func TestGetSchemaForProtocol(t *testing.T) {
	protocols := []string{
		"aave_v3",
		"euler",
		"morpho",
		"nouns",
		"yearn_v3",
		"assert",
		"boolean",
		"math",
	}

	for _, protocol := range protocols {
		t.Run(fmt.Sprintf("Protocol_%s", protocol), func(t *testing.T) {
			resp, body, err := makeTestRequest(
				fmt.Sprintf("http://localhost:8080/solver?chainId=8453&protocol=%s", protocol),
				http.MethodGet,
				nil,
			)
			if err != nil {
				errorLog(t, "ERROR: %v", err)
				t.Fatalf("Failed to make request for protocol %s: %v", protocol, err)
			}

			// Some protocols might not be available on base, so we'll skip if we get 400
			if resp.StatusCode == http.StatusBadRequest {
				warningLog(t, "Protocol %s not available on base", protocol)
				t.Skipf("Protocol %s not available on base", protocol)
				return
			}

			utils.CheckEqual(t, http.StatusOK, resp.StatusCode)

			var schemaResponse map[string]interface{}
			err = json.Unmarshal(body, &schemaResponse)
			utils.RequireNoError(t, err)

			// Check that our protocol is in the response
			_, ok := schemaResponse[protocol]
			utils.RequireTrue(t, ok, "Protocol %s should be in the response", protocol)
			successLog(t, "Successfully retrieved schema for protocol %s", protocol)
		})
	}
}

// TestGetSchemaForActions tests the /solver endpoint for specific actions
func TestGetSchemaForActions(t *testing.T) {
	// Base chain ID - we'll prioritize this for testing
	const baseChainID = 8453

	// Use the protocols that we have test case files for
	protocols := []string{
		"euler",
		"aave_v3",
		"morpho",
		"yearn_v3",
		"plug",
	}

	// Load and test schema for each protocol
	for _, protocol := range protocols {
		testCases, err := loadTestCasesForProtocol(protocol)
		if err != nil {
			warningLog(t, "Warning: %v", err)
			// Continue with empty test cases rather than skipping entirely
			testCases = []TestCase{}
		}

		// If we have no test cases, log a warning but don't skip
		if len(testCases) == 0 {
			warningLog(t, "No test cases found for protocol %s", protocol)
		}

		// Create a map to track which actions we've already tested
		// so we don't test the same action multiple times
		testedActions := make(map[string]bool)

		// Create maps to track which actions are available on each chain
		baseChainActions := make(map[string]bool)
		otherChainActions := make(map[string]map[uint64]bool)

		// First pass: categorize actions by chain
		for _, tc := range testCases {
			// Extract the action from the test case
			var action string
			if len(tc.Intent.Inputs) > 0 {
				if a, ok := tc.Intent.Inputs[0]["action"].(string); ok {
					action = a
				} else {
					continue // Skip if we can't determine the action
				}
			} else {
				continue // Skip if there are no inputs
			}

			// Check which chain this test case is for
			if tc.Intent.ChainId == baseChainID {
				baseChainActions[action] = true
			} else {
				if otherChainActions[action] == nil {
					otherChainActions[action] = make(map[uint64]bool)
				}
				otherChainActions[action][tc.Intent.ChainId] = true
			}
		}

		// Second pass: run tests, prioritizing Base chain
		for _, tc := range testCases {
			// Extract the action from the test case
			var action string
			if len(tc.Intent.Inputs) > 0 {
				if a, ok := tc.Intent.Inputs[0]["action"].(string); ok {
					action = a
				} else {
					continue // Skip if we can't determine the action
				}
			} else {
				continue // Skip if there are no inputs
			}

			// Skip if we've already tested this action on any chain
			if testedActions[action] {
				continue
			}

			// If this action is available on Base, but this test case isn't for Base,
			// skip it as we'll use the Base version instead
			if baseChainActions[action] && tc.Intent.ChainId != baseChainID {
				continue
			}

			// We'll use this test case
			testedActions[action] = true

			// Run the test for this protocol/action combination
			t.Run(fmt.Sprintf("%s_%s_schema_chain_%d", protocol, action, tc.Intent.ChainId), func(t *testing.T) {
				resp, body, err := makeTestRequest(
					fmt.Sprintf("http://localhost:8080/solver?chainId=%d&protocol=%s&action=%s&from=%s",
						tc.Intent.ChainId, protocol, action, testAddress),
					http.MethodGet,
					nil,
				)
				if err != nil {
					errorLog(t, "ERROR: %v", err)
					t.Fatalf("Failed to make request for protocol %s action %s: %v", protocol, action, err)
				}

				// All schema endpoints should work with status 200 for valid inputs
				if resp.StatusCode != http.StatusOK {
					errorLog(t, "Action %s should be available for protocol %s on chain %d. Got status %d",
						action, protocol, tc.Intent.ChainId, resp.StatusCode)
					t.Errorf("Action %s should be available for protocol %s on chain %d. Got status %d",
						action, protocol, tc.Intent.ChainId, resp.StatusCode)
					return
				}

				var schemaResponse map[string]interface{}
				err = json.Unmarshal(body, &schemaResponse)
				utils.RequireNoError(t, err)

				// Check that our protocol is in the response
				protocolData, ok := schemaResponse[protocol]
				utils.RequireTrue(t, ok, "Protocol %s should be in the response", protocol)

				// Extract schema and check for the action
				protocolMap, ok := protocolData.(map[string]interface{})
				utils.RequireTrue(t, ok)

				schema, ok := protocolMap["schema"].(map[string]interface{})
				utils.RequireTrue(t, ok)

				_, ok = schema[action]
				utils.RequireTrue(t, ok, "Action %s should be in the schema", action)
				successLog(t, "Successfully retrieved schema for %s.%s on chain %d", protocol, action, tc.Intent.ChainId)
			})
		}
	}

	// Add tests for protocols that don't have test case files yet
	additionalTests := []struct {
		protocol string
		action   string
		chainId  uint64
	}{
		{"nouns", "bid", 1},
	}

	for _, tc := range additionalTests {
		t.Run(fmt.Sprintf("%s_%s_schema_chain_%d", tc.protocol, tc.action, tc.chainId), func(t *testing.T) {
			resp, body, err := makeTestRequest(
				fmt.Sprintf("http://localhost:8080/solver?chainId=%d&protocol=%s&action=%s", tc.chainId, tc.protocol, tc.action),
				http.MethodGet,
				nil,
			)
			if err != nil {
				errorLog(t, "ERROR: %v", err)
				t.Fatalf("Failed to make request for protocol %s action %s: %v", tc.protocol, tc.action, err)
			}

			// All schema endpoints should work with status 200 for valid inputs
			if resp.StatusCode != http.StatusOK {
				errorLog(t, "Action %s should be available for protocol %s on chain %d. Got status %d",
					tc.action, tc.protocol, tc.chainId, resp.StatusCode)
				t.Errorf("Action %s should be available for protocol %s on chain %d. Got status %d",
					tc.action, tc.protocol, tc.chainId, resp.StatusCode)
				return
			}

			var schemaResponse map[string]interface{}
			err = json.Unmarshal(body, &schemaResponse)
			utils.RequireNoError(t, err)

			// Check that our protocol is in the response
			protocolData, ok := schemaResponse[tc.protocol]
			utils.RequireTrue(t, ok, "Protocol %s should be in the response", tc.protocol)

			// Extract schema and check for the action
			protocolMap, ok := protocolData.(map[string]interface{})
			utils.RequireTrue(t, ok)

			schema, ok := protocolMap["schema"].(map[string]interface{})
			utils.RequireTrue(t, ok)

			_, ok = schema[tc.action]
			utils.RequireTrue(t, ok, "Action %s should be in the schema", tc.action)
			successLog(t, "Successfully retrieved schema for %s.%s on chain %d", tc.protocol, tc.action, tc.chainId)
		})
	}
}

// Test statistics structure to collect result metrics
type testStatistics struct {
	TotalTests     int
	PassedTests    int
	FailedTests    int
	SkippedTests   int
	Protocols      map[string]bool
	Actions        map[string]map[string]bool
	ChainTests     map[uint64]int
	ChainPassed    map[uint64]int
	ErrorsByType   map[string]int
	PanicsDetected int
}

// Prints a summarized test report similar to the bash script
func printTestSummary(buf *bytes.Buffer, stats *testStatistics) {
	output := buf.String()

	// Print header
	fmt.Printf("\n%s╔════════════════════════════════════════╗%s\n", colorBlue, colorReset)
	fmt.Printf("%s║              Test Summary              ║%s\n", colorBlue, colorReset)
	fmt.Printf("%s╚════════════════════════════════════════╝%s\n\n", colorBlue, colorReset)

	// Extract tested protocols
	protocols := extractProtocolsTested(output)
	fmt.Printf("%sProtocols tested:%s\n", colorBlue, colorReset)
	if len(protocols) == 0 {
		fmt.Println("    None detected")
	} else {
		for _, p := range protocols {
			fmt.Printf("    ✓ %s\n", p)
		}
	}

	// Extract actions tested
	actions := extractActionsTested(output)
	fmt.Printf("\n%sActions tested:%s\n", colorBlue, colorReset)
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

	fmt.Printf("\n%sChain-specific test statistics:%s\n", colorBlue, colorReset)
	if baseTests == 0 && mainnetTests == 0 {
		fmt.Println("    No chain-specific tests detected")
	} else {
		if baseTests > 0 {
			fmt.Printf("    Base (Chain 8453): Ran %d tests, %s%d passed%s\n",
				baseTests, colorGreen, basePassed, colorReset)
		}
		if mainnetTests > 0 {
			fmt.Printf("    Ethereum (Chain 1): Ran %d tests, %s%d passed%s\n",
				mainnetTests, colorGreen, mainnetPassed, colorReset)
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
			colorRed, panics, colorReset)
	}

	// Check for rate limits
	rateLimits := countMatches(output, "rate limit exceeded")
	if rateLimits > 0 {
		fmt.Printf("\n%s⚠️  Warning: %d rate limit exceeded errors detected.%s\n",
			colorYellow, rateLimits, colorReset)
		fmt.Println("The test may be hitting the API too frequently. Consider:")
		fmt.Println("  1. Increasing the rate limit in the database for the test API key")
		fmt.Println("  2. Adding delays between API calls")
		fmt.Println("  3. Running fewer tests in parallel")
	}

	// Print test result summary
	fmt.Printf("\n📊 %sPassed: %d%s | %sFailed: %d%s | %sSkipped: %d%s\n",
		colorGreen, passes, colorReset,
		colorRed, failures, colorReset,
		colorYellow, skips, colorReset)
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

// TestGetSolution tests the POST /solver endpoint with various protocol/action combinations
func TestGetSolution(t *testing.T) {
	// This test is especially important as it tests the core functionality with real inputs
	infoLog(t, "Testing solution generation for all supported protocol/action combinations")
	testEOAAddress := "0x50701f4f523766bFb5C195F93333107d1cB8cD90"

	// Base chain ID - we'll prioritize this for testing
	const baseChainID = 8453

	// Define protocols to test, with preferred chains
	baseProtocols := []string{"morpho", "euler", "yearn_v3", "aave_v3"} // Protocols to test on Base chain
	mainnetProtocols := []string{"nouns"}                               // Protocols to test on Ethereum mainnet
	utilityProtocols := []string{"assert", "boolean", "math"}           // Chain-agnostic utility protocols

	// Run tests for each group
	for _, protocol := range baseProtocols {
		runProtocolTests(t, protocol, testEOAAddress, baseChainID)
	}

	for _, protocol := range mainnetProtocols {
		runProtocolTests(t, protocol, testEOAAddress, 1)
	}

	for _, protocol := range utilityProtocols {
		runProtocolTests(t, protocol, testEOAAddress, baseChainID)
	}
}

// Helper function to run all tests for a specific protocol with preferred chain ID
func runProtocolTests(t *testing.T, protocol string, testAddress string, preferredChainID uint64) {
	// Check if we're filtering tests by chain ID
	forceSpecificChain := os.Getenv("TEST_CHAIN_ID")
	var forceChainID uint64
	if forceSpecificChain != "" {
		// Parse the chain ID from the environment variable
		id, err := strconv.ParseUint(forceSpecificChain, 10, 64)
		if err == nil {
			forceChainID = id
			infoLog(t, "Forcing tests to run only on chain ID %d based on TEST_CHAIN_ID", forceChainID)
		} else {
			warningLog(t, "Invalid TEST_CHAIN_ID value: %s", forceSpecificChain)
		}
	}

	// Load test cases for this protocol from its JSON file
	testCases, err := loadTestCasesForProtocol(protocol)
	if err != nil {
		warningLog(t, "Warning: %v", err)
		// Continue with empty test cases rather than skipping entirely
		testCases = []TestCase{}
	}

	// If we have no test cases, log a warning but don't skip
	if len(testCases) == 0 {
		warningLog(t, "No test cases found for protocol %s", protocol)
	}

	// Filter test cases by chain
	preferredChainTestCases := make([]TestCase, 0)
	otherChainTestCases := make([]TestCase, 0)

	// Sort test cases by chain
	for _, tc := range testCases {
		// If forcing a specific chain, only include those test cases
		if forceSpecificChain != "" && tc.Intent.ChainId != forceChainID {
			continue
		}

		if tc.Intent.ChainId == preferredChainID {
			preferredChainTestCases = append(preferredChainTestCases, tc)
		} else {
			otherChainTestCases = append(otherChainTestCases, tc)
		}
	}

	// Determine which test cases to run
	finalTestCases := preferredChainTestCases
	if len(finalTestCases) == 0 && forceSpecificChain == "" {
		// Only fall back to other chains if we're not forcing a specific chain
		infoLog(t, "No test cases for protocol %s on chain %d, using alternative chains", protocol, preferredChainID)
		finalTestCases = otherChainTestCases
	} else if len(finalTestCases) == 0 && forceSpecificChain != "" {
		// If forcing a specific chain and no test cases, skip this protocol
		infoLog(t, "No test cases for protocol %s on chain %d (from TEST_CHAIN_ID), skipping", protocol, forceChainID)
		return
	} else {
		infoLog(t, "Using %d test cases for protocol %s on chain %d", len(finalTestCases), protocol, preferredChainID)
	}

	// Track actions already tested to avoid duplicates
	testedActions := make(map[string]bool)

	// Run each test case
	for _, tc := range finalTestCases {
		// Extract the action to avoid duplicates
		var action string
		if len(tc.Intent.Inputs) > 0 {
			if a, ok := tc.Intent.Inputs[0]["action"].(string); ok {
				action = a
			}
		}

		// Skip if we've already tested this action (unless it's empty or we can't determine it)
		if action != "" && testedActions[action] {
			continue
		}
		if action != "" {
			testedActions[action] = true
		}

		// Run the test with context about which chain we're using
		testName := fmt.Sprintf("%s_%s_chain_%d", protocol, action, tc.Intent.ChainId)
		if action == "" {
			testName = fmt.Sprintf("%s_chain_%d", tc.Name, tc.Intent.ChainId)
		}

		t.Run(testName, func(t *testing.T) {
			// Prepare request info for better error reporting
			protocol := ""
			action := ""
			if len(tc.Intent.Inputs) > 0 {
				if p, ok := tc.Intent.Inputs[0]["protocol"].(string); ok {
					protocol = p
				}
				if a, ok := tc.Intent.Inputs[0]["action"].(string); ok {
					action = a
				}
			}
			infoLog(t, "Testing protocol=%s action=%s on chain=%d", protocol, action, tc.Intent.ChainId)

			// Override the 'From' address with our test address if one wasn't provided
			if tc.Intent.From == "" {
				tc.Intent.From = testAddress
			}

			// Make the request
			resp, body, err := makeTestRequest("http://localhost:8080/solver", http.MethodPost, tc.Intent)
			if err != nil {
				errorLog(t, "ERROR: %v", err)
				t.Fatalf("Failed to make request for protocol %s action %s: %v", protocol, action, err)
			}

			// Check response status
			if resp.StatusCode != http.StatusOK {
				errorMsg := fmt.Sprintf("Expected status code %d, got %d.",
					http.StatusOK, resp.StatusCode)

				// Try to extract error message from response body
				var errorObj map[string]interface{}
				if err := json.Unmarshal(body, &errorObj); err == nil {
					if errMsg, ok := errorObj["error"].(string); ok {
						errorMsg += fmt.Sprintf(" Error: %s", errMsg)

						// Special case: Skip "no transactions to execute" errors for now
						// TODO MASON: get rid of this test skip logic once we've figured out how to handle read methods
						if strings.Contains(errMsg, "no transactions to execute") {
							warningLog(t, "Skipping test case %s due to 'no transactions to execute' error", tc.Name)
							t.Skipf("Skipping test case with read-only method: %s", tc.Name)
							return
						}
					} else if errorMsg, ok := errorObj["message"].(string); ok {
						errorMsg += fmt.Sprintf(" Message: %s", errorMsg)

						// Special case: Skip "no transactions to execute" errors for now
						// TODO MASON: get rid of this test skip logic once we've figured out how to handle read methods
						if strings.Contains(errorMsg, "no transactions to execute") {
							warningLog(t, "Skipping test case %s due to 'no transactions to execute' error", tc.Name)
							t.Skipf("Skipping test case with read-only method: %s", tc.Name)
							return
						}
					} else {
						errorMsg += fmt.Sprintf(" Response: %s", string(body))
					}
				} else {
					errorMsg += fmt.Sprintf(" Raw response: %s", string(body))
				}

				// Print full request for debugging
				reqJSON, _ := json.MarshalIndent(tc.Intent, "", "  ")

				// Handle the failure based on whether we expect it to succeed
				if tc.ExpectOk {
					// Get more specific error categorization to help debugging
					failureCategory := categorizeFailure(string(body), tc.Intent.ChainId, protocol, action)

					// Special case: Skip "no transactions to execute" errors for now
					// TODO MASON: get rid of this test skip logic once we've figured out how to handle read methods
					if strings.Contains(string(body), "no transactions to execute") {
						warningLog(t, "Skipping test case %s due to 'no transactions to execute' error", tc.Name)
						t.Skipf("Skipping test case with read-only method: %s", tc.Name)
						return
					}

					errorLog(t, "%s\nFailure category: %s\nRequest: %s", errorMsg, failureCategory, string(reqJSON))

					// Mark as failing but continue tests - don't stop everything for one failing test
					t.Errorf("Test case %s failed: %s", tc.Name, errorMsg)
				} else {
					// If we expected an error, log it as a success
					successLog(t, "Got expected error status code %d for %s.%s on chain %d",
						resp.StatusCode, protocol, action, tc.Intent.ChainId)
				}
				return
			}

			// Parse the response
			var solution map[string]interface{}
			err = json.Unmarshal(body, &solution)
			if err != nil {
				errorLog(t, "Failed to unmarshal response: %v. Response body: %s", err, string(body))
				t.Errorf("Failed to parse response JSON: %v", err)
				return
			}

			// Validate the essential fields
			missingFields := []string{}
			for _, field := range []string{"intentId", "run"} {
				if _, ok := solution[field]; !ok {
					missingFields = append(missingFields, field)
				}
			}

			if len(missingFields) > 0 {
				errorLog(t, "Response missing required fields: %v.\nFull response: %s",
					missingFields, string(body))
				t.Errorf("Response missing required fields: %v", missingFields)
				return
			}

			// More detailed validation
			if intents, ok := solution["intents"].([]interface{}); ok {
				// Check that we have at least one intent
				if len(intents) == 0 {
					errorLog(t, "Solution contains empty intents array")
					t.Errorf("Solution contains empty intents array")
					return
				}

				// Log the number of intents for debugging
				infoLog(t, "Solution contains %d intents", len(intents))
			}

			// Validate simulation results if present
			if sim, ok := solution["simulationResults"].(map[string]interface{}); ok {
				// Check if we have simulation data
				if status, exists := sim["status"].(string); exists {
					infoLog(t, "Simulation status: %s", status)

					// If simulation status indicates failure but we expected success, that's an issue
					if status != "success" && tc.ExpectOk {
						warningLog(t, "Simulation returned status %q but test case expected success", status)
					}
				}
			}

			successLog(t, "✓ Successfully generated solution for %s.%s on chain %d", protocol, action, tc.Intent.ChainId)
		})
	}
}

// Helper function to categorize failures for better debugging
func categorizeFailure(responseBody string, chainId uint64, protocol, action string) string {
	// Look for common error patterns
	switch {
	case contains(responseBody, "rate limit"):
		return "RATE_LIMIT_EXCEEDED"
	case contains(responseBody, "not supported") || contains(responseBody, "not implemented"):
		return "FEATURE_NOT_SUPPORTED"
	case contains(responseBody, "not found") || contains(responseBody, "couldn't find"):
		return "RESOURCE_NOT_FOUND"
	case contains(responseBody, "chain") && contains(responseBody, "not supported"):
		return fmt.Sprintf("CHAIN_NOT_SUPPORTED (Chain ID: %d)", chainId)
	case contains(responseBody, "validation") || contains(responseBody, "invalid"):
		return "INPUT_VALIDATION_ERROR"
	case contains(responseBody, "rpc error") || contains(responseBody, "connection"):
		return "RPC_CONNECTION_ERROR"
	case contains(responseBody, "contract") && contains(responseBody, "error"):
		return "CONTRACT_INTERACTION_ERROR"
	case contains(responseBody, "timeout") || contains(responseBody, "timed out"):
		return "TIMEOUT"
	default:
		return "UNKNOWN_ERROR"
	}
}

// Helper function to check if a string contains another string
func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// Helper function to load test cases for a specific protocol
func loadTestCasesForProtocol(protocol string) ([]TestCase, error) {
	// Find the current directory
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	// Look for the protocol's test file
	testFilePath := filepath.Join(dir, "cases", fmt.Sprintf("%s_test_cases.json", protocol))

	// Check if the file exists
	if _, err := os.Stat(testFilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("test cases file not found: %s", testFilePath)
	}

	// Read and parse the test file
	fileData, err := os.ReadFile(testFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading test cases file: %w", err)
	}

	var testCases []TestCase
	err = json.Unmarshal(fileData, &testCases)
	if err != nil {
		return nil, fmt.Errorf("error parsing test cases: %w", err)
	}

	return testCases, nil
}
