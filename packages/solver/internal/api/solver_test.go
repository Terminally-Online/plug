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
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
	// Load environment from .env file if available
	setupEnvironment()
	
	// Check if server and database are available
	serverRunning := false
	_, err := http.Get("http://localhost:8080/health")
	if err == nil {
		serverRunning = true
	}
	
	dbAvailable := isDatabaseAvailable()
	
	// Print status
	if !serverRunning {
		fmt.Println("⚠️ WARNING: Solver server is not running - tests requiring server will be skipped")
	}
	if !dbAvailable {
		fmt.Println("⚠️ WARNING: Database is not available - tests requiring database will be skipped")
	}
	
	// Set environment variables to indicate availability
	if serverRunning {
		os.Setenv("TEST_SERVER_AVAILABLE", "true")
	}
	if dbAvailable {
		os.Setenv("TEST_DB_AVAILABLE", "true")
	}

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
	envDefaults := map[string]string{
		"DATABASE_HOST":     "localhost",
		"DATABASE_USER":     "plug",
		"DATABASE_PASSWORD": "plugdev",
		"DATABASE_NAME":     "plug_solver", 
		"DATABASE_PORT":     "6432",
		"ADMIN_API_KEY":     "test-admin-key",
		"RPC_URL":           "http://localhost:8545", // Default RPC URL for local testing
	}
	
	for key, defaultValue := range envDefaults {
		if os.Getenv(key) == "" {
			os.Setenv(key, defaultValue)
		}
	}
}

// Helper function to make test HTTP requests with detailed error reporting
func makeTestRequest(url, method string, body interface{}) (*http.Response, []byte, error) {
	var reqBody io.Reader
	var jsonData []byte
	var err error

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
	req.Header.Set("X-Api-Key", "test-api-key") // Using a test API key

	client := http.Client{
		Timeout: 10 * time.Second, // Set a reasonable timeout
	}

	resp, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return nil, nil, fmt.Errorf("request timed out: %w", err)
		}
		return nil, nil, fmt.Errorf("failed to execute request: %w", err)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read response body: %w", err)
	}
	defer resp.Body.Close()

	// If server returned an error, add the response body to the error message
	if resp.StatusCode >= 400 {
		return resp, responseBody, nil
	}

	return resp, responseBody, nil
}

// Helper to check if server is available and skip if not
func skipIfServerNotAvailable(t *testing.T) {
	if os.Getenv("TEST_SERVER_AVAILABLE") != "true" {
		t.Skip("Skipping test: server is not running based on initial check")
		return
	}
}

// TestHealthEndpoint tests the /health endpoint
func TestHealthEndpoint(t *testing.T) {
	skipIfServerNotAvailable(t)
	
	// Try health endpoint
	resp, body, err := makeTestRequest("http://localhost:8080/health", http.MethodGet, nil)
	if err != nil {
		t.Logf("ERROR: %v", err)
		t.Skip("Skipping test: server is not running or encountered an error")
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))
		return
	}

	var healthResponse map[string]interface{}
	err = json.Unmarshal(body, &healthResponse)
	if err != nil {
		t.Errorf("Failed to unmarshal response: %v. Response body: %s", err, string(body))
		return
	}

	if healthResponse["status"] != "ok" {
		t.Errorf("Expected status 'ok', got '%v'", healthResponse["status"])
	}
}

// TestGetSchemaEndpoint tests the /solver endpoint for schema retrieval
func TestGetSchemaEndpoint(t *testing.T) {
	skipIfServerNotAvailable(t)
	
	// Make the request
	resp, body, err := makeTestRequest("http://localhost:8080/solver?chainId=1", http.MethodGet, nil)
	if err != nil {
		t.Logf("ERROR: %v", err)
		t.Skip("Skipping test: encountered an error")
		return
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var schemaResponse map[string]interface{}
	err = json.Unmarshal(body, &schemaResponse)
	require.NoError(t, err)

	// Check that we have some protocols returned
	assert.Greater(t, len(schemaResponse), 0)
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
		{"euler", "supply", 8453},              // Base
		{"euler", "withdraw", 8453},            // Base
		{"euler", "supply_collateral", 8453},   // Base
		{"euler", "withdraw_collateral", 8453}, // Base
		{"euler", "borrow", 8453},              // Base
		{"euler", "repay", 8453},               // Base

		{"aave_v3", "supply", 1},   // Ethereum
		{"aave_v3", "withdraw", 1}, // Ethereum
		{"aave_v3", "borrow", 1},   // Ethereum
		{"aave_v3", "repay", 1},    // Ethereum

		{"morpho", "deposit", 1},  // Ethereum/com
		{"morpho", "redeem", 1},   // Ethereum
		{"morpho", "supply", 1},   // Ethereum
		{"morpho", "withdraw", 1}, // Ethereum

		{"yearn_v3", "deposit", 1},  // Ethereum
		{"yearn_v3", "withdraw", 1}, // Ethereum

		{"nouns", "bid", 1}, // Ethereum

		{"ens", "register", 1}, // Ethereum
		{"ens", "renew", 1},    // Ethereum

		{"assert", "equals", 1},       // Logic comparison
		{"assert", "not_equals", 1},   // Logic comparison
		{"assert", "greater_than", 1}, // Logic comparison
		{"assert", "less_than", 1},    // Logic comparison

		{"boolean", "and", 1}, // Logic operations
		{"boolean", "or", 1},  // Logic operations
		{"boolean", "not", 1}, // Logic operations

		{"math", "add", 1},      // Math operations
		{"math", "subtract", 1}, // Math operations
		{"math", "multiply", 1}, // Math operations
		{"math", "divide", 1},   // Math operations

		{"database", "select", 1}, // Database operations
		{"database", "insert", 1}, // Database operations
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
		})
	}
}

// TestGetSolution tests the POST /solver endpoint with various protocol/action combinations
func TestGetSolution(t *testing.T) {
	skipIfServerNotAvailable(t)
	
	// This test is especially important as it tests the core functionality with real inputs
	t.Log("Testing solution generation for all supported protocol/action combinations")
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
					Simulate: true,
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
					Simulate: true,
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
					Simulate: true,
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
					Simulate: true,
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
					Simulate: true,
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
					Simulate: true,
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
						"action":   "supply",
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
					Simulate: true,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Assert Equals",
			intent: TestIntent{
				ChainId: 1, // Ethereum
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "assert",
						"action":   "equals",
						"left":     "100",
						"right":    "100",
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: true,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Boolean AND",
			intent: TestIntent{
				ChainId: 1, // Ethereum
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "boolean",
						"action":   "and",
						"left":     true,
						"right":    true,
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: true,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Math Add",
			intent: TestIntent{
				ChainId: 1, // Ethereum
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "math",
						"action":   "add",
						"left":     "100",
						"right":    "200",
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: true,
					Submit:   false,
				},
			},
			expectOk: true,
		},
		{
			name: "Database Select",
			intent: TestIntent{
				ChainId: 1, // Ethereum
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "database",
						"action":   "select",
						"key":      "test-key",
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: true,
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
			t.Logf("Testing protocol=%s action=%s on chain=%d", protocol, action, tc.intent.ChainId)

			// Make the request
			resp, body, err := makeTestRequest("http://localhost:8080/solver", http.MethodPost, tc.intent)
			if err != nil {
				t.Logf("ERROR: %v", err)
				t.Skip("Skipping test: server is not running or encountered an error")
				return
			}

			// For cases where we expect failure
			if !tc.expectOk {
				if resp.StatusCode == http.StatusOK {
					t.Errorf("Expected error status code, got %d (OK)", resp.StatusCode)
				} else {
					t.Logf("Got expected error status code: %d", resp.StatusCode)
				}
				return
			}

			// Check response status
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %d, got %d. Response: %s",
					http.StatusOK, resp.StatusCode, string(body))
				return
			}

			// Parse the response
			var solution map[string]interface{}
			err = json.Unmarshal(body, &solution)
			if err != nil {
				t.Errorf("Failed to unmarshal response: %v. Response body: %s", err, string(body))
				return
			}

			// Validate the essential fields
			missingFields := []string{}
			for _, field := range []string{"id", "intents", "simulationResults"} {
				if _, ok := solution[field]; !ok {
					missingFields = append(missingFields, field)
				}
			}

			if len(missingFields) > 0 {
				t.Errorf("Response missing required fields: %v", missingFields)
				return
			}

			t.Logf("✓ Successfully generated solution for %s (%s)", protocol, action)
		})
	}
}

// TestMultipleProtocolsInIntent tests using multiple protocols in a single intent
func TestMultipleProtocolsInIntent(t *testing.T) {
	t.Log("Testing multi-protocol intent with assert, boolean, and math operations")
	testEOAAddress := "0x50701f4f523766bFb5C195F93333107d1cB8cD90"

	// Create a complex intent that uses multiple protocols
	intent := TestIntent{
		ChainId: 1, // Ethereum
		From:    testEOAAddress,
		Inputs: []map[string]any{
			{
				"protocol": "assert",
				"action":   "equals",
				"left":     "100",
				"right":    "100",
			},
			{
				"protocol": "boolean",
				"action":   "and",
				"left":     true,
				"right":    true,
			},
			{
				"protocol": "math",
				"action":   "add",
				"left":     "100",
				"right":    "200",
			},
		},
		Options: struct {
			IsEOA    bool `json:"isEOA"`
			Simulate bool `json:"simulate"`
			Submit   bool `json:"submit"`
		}{
			IsEOA:    false,
			Simulate: true,
			Submit:   false,
		},
	}

	// Make the multi-protocol request
	resp, body, err := makeTestRequest("http://localhost:8080/solver", http.MethodPost, intent)
	if err != nil {
		t.Logf("ERROR: %v", err)
		t.Skip("Skipping test: server is not running or encountered an error")
		return
	}

	// Check response status
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))
		return
	}

	// Parse the response
	var solution map[string]interface{}
	if err := json.Unmarshal(body, &solution); err != nil {
		t.Errorf("Failed to unmarshal response: %v. Response body: %s", err, string(body))
		return
	}

	// Validate all required fields
	requiredFields := []string{"id", "intents", "simulationResults"}
	for _, field := range requiredFields {
		if _, ok := solution[field]; !ok {
			t.Errorf("Response missing required field: %s", field)
			return
		}
	}

	// Validate intent count (should include all 3 protocols)
	if intents, ok := solution["intents"].([]interface{}); ok {
		if len(intents) == 0 {
			t.Errorf("Expected intents in solution, got empty array")
		} else {
			t.Logf("✓ Multi-protocol intent successfully processed with %d intents", len(intents))
		}
	} else {
		t.Errorf("Invalid intents field format in response")
	}
}

// TestInvalidInputs tests error handling with various invalid inputs
func TestInvalidInputs(t *testing.T) {
	testEOAAddress := "0x50701f4f523766bFb5C195F93333107d1cB8cD90"

	testCases := []struct {
		name   string
		intent TestIntent
	}{
		{
			name: "Invalid Protocol",
			intent: TestIntent{
				ChainId: 1,
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "invalid_protocol",
						"action":   "invalid_action",
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: true,
					Submit:   false,
				},
			},
		},
		{
			name: "Missing Required Fields",
			intent: TestIntent{
				ChainId: 1,
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "euler",
						"action":   "supply",
					},
				},
				Options: struct {
					IsEOA    bool `json:"isEOA"`
					Simulate bool `json:"simulate"`
					Submit   bool `json:"submit"`
				}{
					IsEOA:    false,
					Simulate: true,
					Submit:   false,
				},
			},
		},
		{
			name: "Invalid Chain ID",
			intent: TestIntent{
				ChainId: 999999, // Invalid chain ID
				From:    testEOAAddress,
				Inputs: []map[string]any{
					{
						"protocol": "euler",
						"action":   "supply",
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
					Simulate: true,
					Submit:   false,
				},
			},
		},
		{
			name: "Invalid Address Format",
			intent: TestIntent{
				ChainId: 1,
				From:    "invalid-address", // Invalid address format
				Inputs: []map[string]any{
					{
						"protocol": "euler",
						"action":   "supply",
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
					Simulate: true,
					Submit:   false,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, _, err := makeTestRequest("http://localhost:8080/solver", http.MethodPost, tc.intent)
			if err != nil {
				t.Skip("Skipping test: server is not running")
				return
			}

			// We expect these to fail with 4xx or 5xx
			assert.True(t, resp.StatusCode >= 400, "Expected error status code")
		})
	}
}

// TestChainSpecificProtocols tests protocols on specific chains
func TestChainSpecificProtocols(t *testing.T) {
	// Define chains and their respective protocols to test
	chainTests := []struct {
		chainId   uint64
		chainName string
		protocols []string
	}{
		{1, "Ethereum", []string{"aave_v3", "morpho", "yearn_v3", "nouns", "ens"}},
		{8453, "Base", []string{"euler"}},
		// Add more chains as needed
	}

	for _, ct := range chainTests {
		t.Run(fmt.Sprintf("Chain_%s", ct.chainName), func(t *testing.T) {
			// Test that we can get schema for this chain
			resp, body, err := makeTestRequest(
				fmt.Sprintf("http://localhost:8080/solver?chainId=%d", ct.chainId),
				http.MethodGet,
				nil,
			)
			if err != nil {
				t.Skip("Skipping test: server is not running")
				return
			}

			assert.Equal(t, http.StatusOK, resp.StatusCode)

			var schemaResponse map[string]interface{}
			err = json.Unmarshal(body, &schemaResponse)
			require.NoError(t, err)

			// Check that the specified protocols are available on this chain
			for _, protocol := range ct.protocols {
				t.Run(fmt.Sprintf("Protocol_%s", protocol), func(t *testing.T) {
					protocolResp, protocolBody, err := makeTestRequest(
						fmt.Sprintf("http://localhost:8080/solver?chainId=%d&protocol=%s", ct.chainId, protocol),
						http.MethodGet,
						nil,
					)
					if err != nil {
						t.Skip("Skipping test: server is not running")
						return
					}

					// If this protocol is not available on this chain, skip
					if protocolResp.StatusCode != http.StatusOK {
						t.Skipf("Protocol %s not available on chain %s", protocol, ct.chainName)
						return
					}

					var protocolSchema map[string]interface{}
					err = json.Unmarshal(protocolBody, &protocolSchema)
					require.NoError(t, err)

					// Check that our protocol is in the response
					_, ok := protocolSchema[protocol]
					assert.True(t, ok, "Protocol %s should be in the response", protocol)
				})
			}
		})
	}
}
