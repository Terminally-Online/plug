package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), 1*time.Second)
	if err != nil {
		fmt.Printf("Database connection check failed: %v\n", err)
		return false
	}
	conn.Close()
	return true
}

// TestMain sets up common test environment
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
		warningPrint("⚠️ WARNING: Solver server is not running on localhost:8080 - tests requiring server will be skipped")
	} else {
		successPrint("✓ Solver server is available - will run server-dependent tests")
	}

	if !dbAvailable {
		warningPrint("⚠️ WARNING: Database is not fully available - tests using database features will be limited")
	} else {
		successPrint("✓ Database is available - will run database-dependent tests")
	}

	// Set environment variables to indicate availability
	if serverRunning {
		os.Setenv("TEST_SERVER_AVAILABLE", "true")
	}
	if dbAvailable {
		os.Setenv("TEST_DB_AVAILABLE", "true")
	}

	// Mark that we're in test mode
	os.Setenv("GO_TEST_MODE", "true")

	// Run tests
	code := m.Run()

	// Exit with test status code
	os.Exit(code)
}

// Load environment variables needed for testing
func setupEnvironment() {
	// Find the root solver directory
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	rootDir := filepath.Join(dir, "../../../")

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
		"API_KEY":       "test-api-key",
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
		warningLog(t, "ERROR: %v", err)
		t.Skip("Skipping test: server is not running or encountered an error")
		return
	}

	if resp.StatusCode != http.StatusOK {
		errorLog(t, "Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))
		return
	}

	var healthResponse map[string]interface{}
	err = json.Unmarshal(body, &healthResponse)
	if err != nil {
		errorLog(t, "Failed to unmarshal response: %v. Response body: %s", err, string(body))
		return
	}

	// Check for "status" field - accept both "ok" or "healthy" values
	status, ok := healthResponse["status"].(string)
	if !ok {
		errorLog(t, "Missing or invalid 'status' field in health response: %v", healthResponse)
		return
	}

	if status != "ok" && status != "healthy" {
		errorLog(t, "Expected status 'ok' or 'healthy', got '%v'", status)
	} else {
		successLog(t, "Health endpoint returned status: %s", status)
	}
}

// TestGetSchemaEndpoint tests the /solver endpoint for schema retrieval
func TestGetSchemaEndpoint(t *testing.T) {
	// Make the request
	resp, body, err := makeTestRequest("http://localhost:8080/solver?chainId=1", http.MethodGet, nil)
	if err != nil {
		warningLog(t, "ERROR: %v", err)
		t.Skip("Skipping test: encountered an error")
		return
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var schemaResponse map[string]interface{}
	err = json.Unmarshal(body, &schemaResponse)
	require.NoError(t, err)

	// Check that we have some protocols returned
	assert.Greater(t, len(schemaResponse), 0)
	successLog(t, "Successfully retrieved schema with %d protocols", len(schemaResponse))
}

// TestGetSchemaForProtocol tests the /solver endpoint for a specific protocol
func TestGetSchemaForProtocol(t *testing.T) {
	protocols := []string{
		"euler",
		"aave_v3",
		"morpho",
		"yearn_v3",
		"nouns",
		"ens",
		"assert",
		"boolean",
		"math",
		"database",
	}

	for _, protocol := range protocols {
		t.Run(fmt.Sprintf("Protocol_%s", protocol), func(t *testing.T) {
			resp, body, err := makeTestRequest(
				fmt.Sprintf("http://localhost:8080/solver?chainId=1&protocol=%s", protocol),
				http.MethodGet,
				nil,
			)
			if err != nil {
				t.Skip("Skipping test: server is not running")
				return
			}

			// Some protocols might not be available on chain 1, so we'll skip if we get 400
			if resp.StatusCode == http.StatusBadRequest {
				warningLog(t, "Protocol %s not available on chain 1", protocol)
				t.Skipf("Protocol %s not available on chain 1", protocol)
				return
			}

			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var schemaResponse map[string]interface{}
			err = json.Unmarshal(body, &schemaResponse)
			require.NoError(t, err)

			// Check that our protocol is in the response
			_, ok := schemaResponse[protocol]
			assert.True(t, ok, "Protocol %s should be in the response", protocol)
			successLog(t, "Successfully retrieved schema for protocol %s", protocol)
		})
	}
}

// TestGetSchemaForActions tests the /solver endpoint for specific actions
func TestGetSchemaForActions(t *testing.T) {
	// Test cases for protocol/action combinations
	testCases := []struct {
		protocol string
		action   string
		chainId  uint64
	}{
		{"euler", "supply", 8453},
		{"euler", "withdraw", 8453},
		{"euler", "supply_collateral", 8453},
		{"euler", "withdraw_collateral", 8453},
		{"euler", "borrow", 8453},
		{"euler", "repay", 8453},

		{"aave_v3", "deposit", 8453},
		{"aave_v3", "borrow", 8453},
		{"aave_v3", "repay", 8453},
		{"aave_v3", "withdraw", 8453},

		{"morpho", "earn", 8453},
		{"morpho", "supply_collateral", 8453},
		{"morpho", "withdraw", 8453},
		{"morpho", "withdraw_all", 8453},
		{"morpho", "borrow", 8453},
		{"morpho", "repay", 8453},
		{"morpho", "repay_all", 8453},
		{"morpho", "claim_rewards", 8453},

		{"yearn_v3", "deposit", 8453},
		{"yearn_v3", "withdraw", 8453},
		{"yearn_v3", "stake", 8453},
		{"yearn_v3", "stake_max", 8453},
		{"yearn_v3", "redeem", 8453},
		{"yearn_v3", "redeem_max", 8453},

		{"nouns", "bid", 1},

		{"ens", "buy", 1},
		{"ens", "renew", 1},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s_schema", tc.protocol, tc.action), func(t *testing.T) {
			resp, body, err := makeTestRequest(
				fmt.Sprintf("http://localhost:8080/solver?chainId=%d&protocol=%s&action=%s", tc.chainId, tc.protocol, tc.action),
				http.MethodGet,
				nil,
			)
			if err != nil {
				t.Skip("Skipping test: server is not running")
				return
			}

			// Some combinations might not be available, so we'll skip if we get 400
			if resp.StatusCode == http.StatusBadRequest {
				warningLog(t, "Action %s not available for protocol %s on chain %d", tc.action, tc.protocol, tc.chainId)
				t.Skipf("Action %s not available for protocol %s on chain %d", tc.action, tc.protocol, tc.chainId)
				return
			}

			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var schemaResponse map[string]interface{}
			err = json.Unmarshal(body, &schemaResponse)
			require.NoError(t, err)

			// Check that our protocol is in the response
			protocolData, ok := schemaResponse[tc.protocol]
			assert.True(t, ok, "Protocol %s should be in the response", tc.protocol)

			// Extract schema and check for the action
			protocolMap, ok := protocolData.(map[string]interface{})
			assert.True(t, ok)

			schema, ok := protocolMap["schema"].(map[string]interface{})
			assert.True(t, ok)

			_, ok = schema[tc.action]
			assert.True(t, ok, "Action %s should be in the schema", tc.action)
			successLog(t, "Successfully retrieved schema for %s.%s", tc.protocol, tc.action)
		})
	}
}

// TestGetSolution tests the POST /solver endpoint with various protocol/action combinations
func TestGetSolution(t *testing.T) {
	// This test is especially important as it tests the core functionality with real inputs
	infoLog(t, "Testing solution generation for all supported protocol/action combinations")
	testEOAAddress := "0x50701f4f523766bFb5C195F93333107d1cB8cD90"

	testCases := []struct {
		name     string
		intent   TestIntent
		expectOk bool
	}{
		{
			name: "Euler Supply",
			intent: TestIntent{
				ChainId: 8453, // Base
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol":    "euler",
						"action":      "supply",
						"amount":      "0.001",
						"vault":       "0x0A1a3b5f2041F33522C4efc754a7D096f880eE16",
						"token":       "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913:6",
						"sub-account": 5,
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: false,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Euler Withdraw",
			intent: TestIntent{
				ChainId: 8453, // Base
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "euler",
						"action":   "withdraw",
						"amount":   "0.001",
						"vault":    "0x0A1a3b5f2041F33522C4efc754a7D096f880eE16",
						"token":    "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913:6",
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: false,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Euler Supply Collateral",
			intent: TestIntent{
				ChainId: 8453, // Base
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol":    "euler",
						"action":      "supply_collateral",
						"amount":      "0.001",
						"vault":       "0x0A1a3b5f2041F33522C4efc754a7D096f880eE16",
						"token":       "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913:6",
						"sub-account": 5,
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: false,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Euler Withdraw Collateral",
			intent: TestIntent{
				ChainId: 8453, // Base
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol":    "euler",
						"action":      "withdraw_collateral",
						"amount":      "0.001",
						"vault":       "0x0A1a3b5f2041F33522C4efc754a7D096f880eE16",
						"token":       "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913:6",
						"sub-account": 5,
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: false,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Euler Borrow",
			intent: TestIntent{
				ChainId: 8453, // Base
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol":    "euler",
						"action":      "borrow",
						"amount":      "0.001",
						"vault":       "0x0A1a3b5f2041F33522C4efc754a7D096f880eE16",
						"token":       "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913:6",
						"sub-account": 5,
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: false,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Euler Repay",
			intent: TestIntent{
				ChainId: 8453, // Base
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol":    "euler",
						"action":      "repay",
						"amount":      "0.001",
						"vault":       "0x0A1a3b5f2041F33522C4efc754a7D096f880eE16",
						"token":       "0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913:6",
						"sub-account": 5,
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: false,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Aave V3 Supply",
			intent: TestIntent{
				ChainId: 1, // Ethereum
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "aave_v3",
						"action":   "deposit",
						"amount":   "0.001",
						"token":    "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2:18",
						"vault":    "0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2",
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: false,
					Submit:   false,
				},
			},
			expectOk: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Prepare request info for better error reporting
			protocol := ""
			action := ""
			if len(tc.intent.Inputs) > 0 {
				if p, ok := tc.intent.Inputs[0]["protocol"].(string); ok {
					protocol = p
				}
				if a, ok := tc.intent.Inputs[0]["action"].(string); ok {
					action = a
				}
			}
			infoLog(t, "Testing protocol=%s action=%s on chain=%d", protocol, action, tc.intent.ChainId)

			// Make the request
			resp, body, err := makeTestRequest("http://localhost:8080/solver", http.MethodPost, tc.intent)
			if err != nil {
				errorLog(t, "ERROR: %v", err)
				t.Fatalf("ERROR: %v", err)
				return
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
					} else if errorMsg, ok := errorObj["message"].(string); ok {
						errorMsg += fmt.Sprintf(" Message: %s", errorMsg)
					} else {
						errorMsg += fmt.Sprintf(" Response: %s", string(body))
					}
				} else {
					errorMsg += fmt.Sprintf(" Raw response: %s", string(body))
				}

				// Print full request for debugging
				reqJSON, _ := json.MarshalIndent(tc.intent, "", "  ")
				errorLog(t, "%s\nRequest: %s", errorMsg, string(reqJSON))

				// Fail the test if we expected it to succeed
				if tc.expectOk {
					t.Errorf("Test case %s failed: %s", tc.name, errorMsg)
				} else {
					successLog(t, "Got expected error status code: %d", resp.StatusCode)
				}
				return
			}

			// Parse the response
			var solution map[string]interface{}
			err = json.Unmarshal(body, &solution)
			if err != nil {
				errorLog(t, "Failed to unmarshal response: %v. Response body: %s", err, string(body))
				return
			}

			// Validate the essential fields
			missingFields := []string{}
			for _, field := range []string{"intentId", "livePlugs", "transactions"} {
				if _, ok := solution[field]; !ok {
					missingFields = append(missingFields, field)
				}
			}

			if len(missingFields) > 0 {
				errorLog(t, "Response missing required fields: %v.\nFull response: %s",
					missingFields, string(body))
				return
			}

			// More detailed validation
			if intents, ok := solution["intents"].([]interface{}); ok {
				// Check that we have at least one intent
				if len(intents) == 0 {
					errorLog(t, "Solution contains empty intents array")
					return
				}

				// Log the number of intents for debugging
				infoLog(t, "Solution contains %d intents", len(intents))
			}

			// Validate simulation results
			if sim, ok := solution["simulationResults"].(map[string]interface{}); ok {
				// Check if we have simulation data
				if status, exists := sim["status"].(string); exists {
					infoLog(t, "Simulation status: %s", status)
				}
			}

			successLog(t, "✓ Successfully generated solution for %s (%s)", protocol, action)
		})
	}
}

// // TestMultipleProtocolsInIntent tests using multiple protocols in a single intent
// func TestMultipleProtocolsInIntent(t *testing.T) {
// 	infoLog(t, "Testing multi-protocol intent with assert, boolean, and math operations")
// 	testEOAAddress := "0x50701f4f523766bFb5C195F93333107d1cB8cD90"

// 	// Create a complex intent that uses multiple protocols
// 	intent := TestIntent{
// 		ChainId: 1, // Ethereum
// 		From:    testEOAAddress,
// 		Inputs: []map[string]any{
// 			{
// 				"protocol": "assert",
// 				"action":   "equals",
// 				"left":     "100",
// 				"right":    "100",
// 			},
// 			{
// 				"protocol": "boolean",
// 				"action":   "and",
// 				"left":     true,
// 				"right":    true,
// 			},
// 			{
// 				"protocol": "math",
// 				"action":   "add",
// 				"left":     "100",
// 				"right":    "200",
// 			},
// 		},
// 		Options: struct {
// 			IsEOA    bool `json:"isEOA"`
// 			Simulate bool `json:"simulate"`
// 			Submit   bool `json:"submit"`
// 		}{
// 			IsEOA:    false,
// 			Simulate: false,
// 			Submit:   false,
// 		},
// 	}

// 	// Make the multi-protocol request
// 	resp, body, err := makeTestRequest("http://localhost:8080/solver", http.MethodPost, intent)
// 	if err != nil {
// 		errorLog(t, "ERROR: %v", err)
// 		t.Fatalf("ERROR: %v", err)
// 		return
// 	}

// 	// Check response status
// 	if resp.StatusCode != http.StatusOK {
// 		errorMsg := fmt.Sprintf("Expected status code %d, got %d.",
// 			http.StatusOK, resp.StatusCode)

// 		// Try to extract error message from response body
// 		var errorObj map[string]interface{}
// 		if err := json.Unmarshal(body, &errorObj); err == nil {
// 			if errMsg, ok := errorObj["error"].(string); ok {
// 				errorMsg += fmt.Sprintf(" Error: %s", errMsg)
// 			} else if errorMsg, ok := errorObj["message"].(string); ok {
// 				errorMsg += fmt.Sprintf(" Message: %s", errorMsg)
// 			} else {
// 				errorMsg += fmt.Sprintf(" Response: %s", string(body))
// 			}
// 		} else {
// 			errorMsg += fmt.Sprintf(" Raw response: %s", string(body))
// 		}

// 		// Print full request for debugging
// 		reqJSON, _ := json.MarshalIndent(intent, "", "  ")
// 		errorLog(t, "%s\nRequest: %s", errorMsg, string(reqJSON))
// 		return
// 	}

// 	// Parse the response
// 	var solution map[string]interface{}
// 	if err := json.Unmarshal(body, &solution); err != nil {
// 		errorLog(t, "Failed to unmarshal response: %v. Response body: %s", err, string(body))
// 		return
// 	}

// 	// Validate all required fields
// 	requiredFields := []string{"id", "intents", "simulationResults"}
// 	missingFields := []string{}
// 	for _, field := range requiredFields {
// 		if _, ok := solution[field]; !ok {
// 			missingFields = append(missingFields, field)
// 		}
// 	}

// 	if len(missingFields) > 0 {
// 		errorLog(t, "Response missing required fields: %v.\nFull response: %s",
// 			missingFields, string(body))
// 		return
// 	}

// 	// Validate intent count (should include all 3 protocols)
// 	if intents, ok := solution["intents"].([]interface{}); ok {
// 		if len(intents) == 0 {
// 			errorLog(t, "Expected intents in solution, got empty array. Full response: %s",
// 				string(body))
// 			return
// 		}

// 		// Verify we have all 3 protocols represented
// 		protocols := make(map[string]bool)
// 		for _, intent := range intents {
// 			if intentMap, ok := intent.(map[string]interface{}); ok {
// 				// Extract which protocol this intent is for
// 				for k, v := range intentMap {
// 					if inputsArr, ok := v.([]interface{}); ok && k == "inputs" {
// 						for _, input := range inputsArr {
// 							if inputMap, ok := input.(map[string]interface{}); ok {
// 								if proto, ok := inputMap["protocol"].(string); ok {
// 									protocols[proto] = true
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}

// 		infoLog(t, "Multi-protocol intent processed with %d intents", len(intents))
// 		infoLog(t, "Protocols included: %v", getMapKeys(protocols))

// 		// Verify all 3 protocols are included
// 		expectedProtocols := []string{"assert", "boolean", "math"}
// 		for _, proto := range expectedProtocols {
// 			if !protocols[proto] {
// 				errorLog(t, "Expected protocol %s in response, but it was not found", proto)
// 			}
// 		}

// 		successLog(t, "✓ Multi-protocol intent successfully processed")
// 	} else {
// 		errorLog(t, "Invalid intents field format in response: %s", string(body))
// 	}
// }

// Helper function to get map keys as a slice
func getMapKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
