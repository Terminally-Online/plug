package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"solver/internal/test/utils"
)

// TestMain sets up common test environment and provides a test summary at the end
func TestMain(m *testing.M) {
	// Load environment from .env file if available and set testing defaults
	utils.SetupEnvironment()

	// Check if server is available
	serverRunning := false
	_, err := http.Get("http://localhost:8080/health")
	if err == nil {
		serverRunning = true
	}

	// Get database availability status
	dbAvailable := utils.IsDatabaseAvailable()

	// Print status with colors
	if !serverRunning {
		utils.ErrorPrint("❌ ERROR: Solver server is not running on localhost:8080 - tests will fail")
	} else {
		utils.SuccessPrint("✓ Solver server is available")
	}

	if !dbAvailable {
		utils.ErrorPrint("❌ ERROR: Database is not available - tests will fail")
	} else {
		utils.SuccessPrint("✓ Database is available")
	}

	// Check for specific chain ID filter
	chainID := os.Getenv("TEST_CHAIN_ID")

	if chainID != "" {
		utils.InfoPrint("Focusing tests on chain ID: %s", chainID)
	}

	// Mark that we're in test mode
	os.Setenv("GO_TEST_MODE", "true")

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

	// Print separator between test output and summary
	fmt.Println("\n" + utils.ColorBlue + strings.Repeat("=", 80) + utils.ColorReset)
	fmt.Println(utils.ColorBlue + "TEST RESULTS SUMMARY" + utils.ColorReset)
	fmt.Println(utils.ColorBlue + strings.Repeat("=", 80) + utils.ColorReset + "\n")

	// Generate and print test summary
	utils.PrintTestSummary(&buf)

	// Exit with test status code
	os.Exit(code)
}

// TestHealthEndpoint tests the /health endpoint
func TestHealthEndpoint(t *testing.T) {
	// Try health endpoint
	resp, body, err := utils.MakeTestRequest("http://localhost:8080/health", http.MethodGet, nil)
	if err != nil {
		utils.ErrorLog(t, "ERROR: %v", err)
		t.Fatalf("Health endpoint test failed: server is not running or encountered an error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		utils.ErrorLog(t, "Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))
		t.Errorf("Health endpoint returned unexpected status code: %d", resp.StatusCode)
		return
	}

	var healthResponse map[string]interface{}
	err = json.Unmarshal(body, &healthResponse)
	if err != nil {
		utils.ErrorLog(t, "Failed to unmarshal response: %v. Response body: %s", err, string(body))
		t.Errorf("Failed to parse health response: %v", err)
		return
	}

	// Check for "status" field - accept both "ok" or "healthy" values
	status, ok := healthResponse["status"].(string)
	if !ok {
		utils.ErrorLog(t, "Missing or invalid 'status' field in health response: %v", healthResponse)
		t.Errorf("Missing or invalid 'status' field in health response")
		return
	}

	if status != "ok" && status != "healthy" {
		utils.ErrorLog(t, "Expected status 'ok' or 'healthy', got '%v'", status)
		t.Errorf("Health endpoint returned unexpected status: %s", status)
	} else {
		utils.SuccessLog(t, "Health endpoint returned status: %s", status)
	}
}

// TestGetSchemaEndpoint tests the /solver endpoint for schema retrieval
func TestGetSchemaEndpoint(t *testing.T) {
	// Make the request
	resp, body, err := utils.MakeTestRequest("http://localhost:8080/solver?chainId=1", http.MethodGet, nil)
	if err != nil {
		utils.ErrorLog(t, "ERROR: %v", err)
		t.Fatalf("Schema endpoint test failed: %v", err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var schemaResponse map[string]interface{}
	err = json.Unmarshal(body, &schemaResponse)
	require.NoError(t, err)

	// Check that we have some protocols returned
	assert.Greater(t, len(schemaResponse), 0)
	utils.SuccessLog(t, "Successfully retrieved schema with %d protocols", len(schemaResponse))
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
			resp, body, err := utils.MakeTestRequest(
				fmt.Sprintf("http://localhost:8080/solver?chainId=1&protocol=%s", protocol),
				http.MethodGet,
				nil,
			)
			if err != nil {
				utils.ErrorLog(t, "ERROR: %v", err)
				t.Fatalf("Failed to make request for protocol %s: %v", protocol, err)
			}

			// Some protocols might not be available on chain 1, so we'll skip if we get 400
			if resp.StatusCode == http.StatusBadRequest {
				utils.WarningLog(t, "Protocol %s not available on chain 1", protocol)
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
			utils.SuccessLog(t, "Successfully retrieved schema for protocol %s", protocol)
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
			utils.WarningLog(t, "Warning: %v", err)
			// Continue with empty test cases rather than skipping entirely
			testCases = []TestCase{}
		}

		// If we have no test cases, log a warning but don't skip
		if len(testCases) == 0 {
			utils.WarningLog(t, "No test cases found for protocol %s", protocol)
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
				resp, body, err := utils.MakeTestRequest(
					fmt.Sprintf("http://localhost:8080/solver?chainId=%d&protocol=%s&action=%s&from=%s",
						tc.Intent.ChainId, protocol, action, utils.TestAddress),
					http.MethodGet,
					nil,
				)
				if err != nil {
					utils.ErrorLog(t, "ERROR: %v", err)
					t.Fatalf("Failed to make request for protocol %s action %s: %v", protocol, action, err)
				}

				// All schema endpoints should work with status 200 for valid inputs
				if resp.StatusCode != http.StatusOK {
					utils.ErrorLog(t, "Action %s should be available for protocol %s on chain %d. Got status %d",
						action, protocol, tc.Intent.ChainId, resp.StatusCode)
					t.Errorf("Action %s should be available for protocol %s on chain %d. Got status %d",
						action, protocol, tc.Intent.ChainId, resp.StatusCode)
					return
				}

				var schemaResponse map[string]interface{}
				err = json.Unmarshal(body, &schemaResponse)
				require.NoError(t, err)

				// Check that our protocol is in the response
				protocolData, ok := schemaResponse[protocol]
				assert.True(t, ok, "Protocol %s should be in the response", protocol)

				// Extract schema and check for the action
				protocolMap, ok := protocolData.(map[string]interface{})
				assert.True(t, ok)

				schema, ok := protocolMap["schema"].(map[string]interface{})
				assert.True(t, ok)

				_, ok = schema[action]
				assert.True(t, ok, "Action %s should be in the schema", action)
				utils.SuccessLog(t, "Successfully retrieved schema for %s.%s on chain %d", protocol, action, tc.Intent.ChainId)
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
		{"ens", "buy", 1},
		{"ens", "renew", 1},
	}

	for _, tc := range additionalTests {
		t.Run(fmt.Sprintf("%s_%s_schema_chain_%d", tc.protocol, tc.action, tc.chainId), func(t *testing.T) {
			resp, body, err := utils.MakeTestRequest(
				fmt.Sprintf("http://localhost:8080/solver?chainId=%d&protocol=%s&action=%s", tc.chainId, tc.protocol, tc.action),
				http.MethodGet,
				nil,
			)
			if err != nil {
				utils.ErrorLog(t, "ERROR: %v", err)
				t.Fatalf("Failed to make request for protocol %s action %s: %v", tc.protocol, tc.action, err)
			}

			// All schema endpoints should work with status 200 for valid inputs
			if resp.StatusCode != http.StatusOK {
				utils.ErrorLog(t, "Action %s should be available for protocol %s on chain %d. Got status %d",
					tc.action, tc.protocol, tc.chainId, resp.StatusCode)
				t.Errorf("Action %s should be available for protocol %s on chain %d. Got status %d",
					tc.action, tc.protocol, tc.chainId, resp.StatusCode)
				return
			}

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
			utils.SuccessLog(t, "Successfully retrieved schema for %s.%s on chain %d", tc.protocol, tc.action, tc.chainId)
		})
	}
}

// Test case definition for protocol tests
type TestCase struct {
	Name     string           `json:"name"`
	Intent   utils.TestIntent `json:"intent"`
	ExpectOk bool             `json:"expectOk"`
}

// TestGetSolution tests the POST /solver endpoint with various protocol/action combinations
func TestGetSolution(t *testing.T) {
	// This test is especially important as it tests the core functionality with real inputs
	utils.InfoLog(t, "Testing solution generation for all supported protocol/action combinations")

	// Base chain ID - we'll prioritize this for testing
	const baseChainID = 8453

	// Define protocols to test, with preferred chains
	baseProtocols := []string{"morpho", "euler"}                          // Protocols to test on Base chain
	mainnetProtocols := []string{"aave_v3", "yearn_v3", "nouns", "ens"}   // Protocols to test on Ethereum mainnet
	utilityProtocols := []string{"assert", "boolean", "math", "database"} // Chain-agnostic utility protocols

	// Run tests for each group
	for _, protocol := range baseProtocols {
		runProtocolTests(t, protocol, utils.TestAddress, baseChainID)
	}

	for _, protocol := range mainnetProtocols {
		runProtocolTests(t, protocol, utils.TestAddress, 1) // Chain ID 1 for Ethereum mainnet
	}

	for _, protocol := range utilityProtocols {
		runProtocolTests(t, protocol, utils.TestAddress, baseChainID) // Test utilities on Base
	}
}

// Helper function to run all tests for a specific protocol with preferred chain ID
func runProtocolTests(t *testing.T, protocol string, testAddress string, preferredChainID uint64) {
	// Check if we're filtering tests by chain ID
	forceSpecificChain := os.Getenv("TEST_CHAIN_ID")
	var forceChainID uint64
	if forceSpecificChain != "" {
		// Parse the chain ID from the environment variable
		var id uint64
		_, err := fmt.Sscanf(forceSpecificChain, "%d", &id)
		if err == nil {
			forceChainID = id
			utils.InfoLog(t, "Forcing tests to run only on chain ID %d based on TEST_CHAIN_ID", forceChainID)
		} else {
			utils.WarningLog(t, "Invalid TEST_CHAIN_ID value: %s", forceSpecificChain)
		}
	}

	// Load test cases for this protocol from its JSON file
	testCases, err := loadTestCasesForProtocol(protocol)
	if err != nil {
		utils.WarningLog(t, "Warning: %v", err)
		// Continue with empty test cases rather than skipping entirely
		testCases = []TestCase{}
	}

	// If we have no test cases, log a warning but don't skip
	if len(testCases) == 0 {
		utils.WarningLog(t, "No test cases found for protocol %s", protocol)
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
		utils.InfoLog(t, "No test cases for protocol %s on chain %d, using alternative chains", protocol, preferredChainID)
		finalTestCases = otherChainTestCases
	} else if len(finalTestCases) == 0 && forceSpecificChain != "" {
		// If forcing a specific chain and no test cases, skip this protocol
		utils.InfoLog(t, "No test cases for protocol %s on chain %d (from TEST_CHAIN_ID), skipping", protocol, forceChainID)
		return
	} else {
		utils.InfoLog(t, "Using %d test cases for protocol %s on chain %d", len(finalTestCases), protocol, preferredChainID)
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
			utils.InfoLog(t, "Testing protocol=%s action=%s on chain=%d", protocol, action, tc.Intent.ChainId)

			// Override the 'From' address with our test address if one wasn't provided
			if tc.Intent.From == "" {
				tc.Intent.From = testAddress
			}

			// Make the request
			resp, body, err := utils.MakeTestRequest("http://localhost:8080/solver", http.MethodPost, tc.Intent)
			if err != nil {
				utils.ErrorLog(t, "ERROR: %v", err)
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
						if utils.Contains(errMsg, "no transactions to execute") {
							utils.WarningLog(t, "Skipping test case %s due to 'no transactions to execute' error", tc.Name)
							t.Skipf("Skipping test case with read-only method: %s", tc.Name)
							return
						}
					} else if errorMsg, ok := errorObj["message"].(string); ok {
						errorMsg += fmt.Sprintf(" Message: %s", errorMsg)

						// Special case: Skip "no transactions to execute" errors for now
						// TODO MASON: get rid of this test skip logic once we've figured out how to handle read methods
						if utils.Contains(errorMsg, "no transactions to execute") {
							utils.WarningLog(t, "Skipping test case %s due to 'no transactions to execute' error", tc.Name)
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
					failureCategory := utils.CategorizeFailure(string(body), tc.Intent.ChainId, protocol, action)

					// Special case: Skip "no transactions to execute" errors for now
					// TODO MASON: get rid of this test skip logic once we've figured out how to handle read methods
					if utils.Contains(string(body), "no transactions to execute") {
						utils.WarningLog(t, "Skipping test case %s due to 'no transactions to execute' error", tc.Name)
						t.Skipf("Skipping test case with read-only method: %s", tc.Name)
						return
					}

					utils.ErrorLog(t, "%s\nFailure category: %s\nRequest: %s", errorMsg, failureCategory, string(reqJSON))

					// Mark as failing but continue tests - don't stop everything for one failing test
					t.Errorf("Test case %s failed: %s", tc.Name, errorMsg)
				} else {
					// If we expected an error, log it as a success
					utils.SuccessLog(t, "Got expected error status code %d for %s.%s on chain %d",
						resp.StatusCode, protocol, action, tc.Intent.ChainId)
				}
				return
			}

			// Parse the response
			var solution map[string]interface{}
			err = json.Unmarshal(body, &solution)
			if err != nil {
				utils.ErrorLog(t, "Failed to unmarshal response: %v. Response body: %s", err, string(body))
				t.Errorf("Failed to parse response JSON: %v", err)
				return
			}

			// Validate the essential fields
			missingFields := []string{}
			for _, field := range []string{"intentId", "transactions"} {
				if _, ok := solution[field]; !ok {
					missingFields = append(missingFields, field)
				}
			}

			if len(missingFields) > 0 {
				utils.ErrorLog(t, "Response missing required fields: %v.\nFull response: %s",
					missingFields, string(body))
				t.Errorf("Response missing required fields: %v", missingFields)
				return
			}

			// More detailed validation
			if intents, ok := solution["intents"].([]interface{}); ok {
				// Check that we have at least one intent
				if len(intents) == 0 {
					utils.ErrorLog(t, "Solution contains empty intents array")
					t.Errorf("Solution contains empty intents array")
					return
				}

				// Log the number of intents for debugging
				utils.InfoLog(t, "Solution contains %d intents", len(intents))
			}

			// Validate simulation results if present
			if sim, ok := solution["simulationResults"].(map[string]interface{}); ok {
				// Check if we have simulation data
				if status, exists := sim["status"].(string); exists {
					utils.InfoLog(t, "Simulation status: %s", status)

					// If simulation status indicates failure but we expected success, that's an issue
					if status != "success" && tc.ExpectOk {
						utils.WarningLog(t, "Simulation returned status %q but test case expected success", status)
					}
				}
			}

			utils.SuccessLog(t, "✓ Successfully generated solution for %s.%s on chain %d", protocol, action, tc.Intent.ChainId)
		})
	}
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
