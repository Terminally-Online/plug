package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"solver/internal/test/utils"
)

// Test intent CRUD operations
func TestIntentCreation(t *testing.T) {
	// Create a new intent
	intent := utils.TestIntent{
		ChainId: 1, // Ethereum mainnet
		From:    utils.TestAddress,
		Inputs: []map[string]any{
			{
				"action":   "transfer",
				"protocol": "assert",
			},
		},
		Options: struct {
			IsEOA    bool `json:"isEOA"`
			Simulate bool `json:"simulate"`
			Submit   bool `json:"submit"`
		}{
			IsEOA:    true,
			Simulate: true,
			Submit:   false,
		},
	}

	// Send the intent creation request
	resp, body, err := utils.MakeTestRequest("http://localhost:8080/solver/save", http.MethodPost, intent)
	if err != nil {
		t.Fatalf("Failed to create intent: %v", err)
	}

	// Ensure the request was successful
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code %d, got %d. Response: %s",
		http.StatusOK, resp.StatusCode, string(body))

	// Parse the response to get the intent ID
	var createdIntent map[string]interface{}
	err = json.Unmarshal(body, &createdIntent)
	require.NoError(t, err, "Failed to parse response JSON")

	// Verify that we received an intent ID
	intentId, ok := createdIntent["id"].(string)
	require.True(t, ok, "Intent ID not found in response")
	require.NotEmpty(t, intentId, "Intent ID should not be empty")

	utils.SuccessLog(t, "Successfully created intent with ID: %s", intentId)

	// Test intent retrieval
	t.Run("GetIntent", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodGet, nil)
		require.NoError(t, err, "Failed to get intent")
		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))

		// The API returns an array of intents
		var retrievedIntents []map[string]interface{}
		err = json.Unmarshal(body, &retrievedIntents)
		require.NoError(t, err, "Failed to parse response JSON")
		require.Greater(t, len(retrievedIntents), 0, "Expected at least one intent in response")

		retrievedIntent := retrievedIntents[0]
		// Verify the intent data matches what we created
		assert.Equal(t, intentId, retrievedIntent["id"], "Intent ID mismatch")
		assert.Equal(t, float64(1), retrievedIntent["chainId"], "Chain ID mismatch")
		assert.Equal(t, utils.TestAddress, retrievedIntent["from"], "From address mismatch")

		utils.SuccessLog(t, "Successfully retrieved intent with ID: %s", intentId)
	})

	// Test updating intent status
	t.Run("ToggleIntentStatus", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s/status", intentId), http.MethodPost, nil)
		require.NoError(t, err, "Failed to toggle intent status")
		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))

		var updatedIntent map[string]interface{}
		err = json.Unmarshal(body, &updatedIntent)
		require.NoError(t, err, "Failed to parse response JSON")

		// Verify the status was toggled from "active" to "paused"
		assert.Equal(t, "paused", updatedIntent["status"], "Intent status should be paused")

		// Toggle it back to active
		resp, body, err = utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s/status", intentId), http.MethodPost, nil)
		require.NoError(t, err, "Failed to toggle intent status back")
		require.Equal(t, http.StatusOK, resp.StatusCode)

		err = json.Unmarshal(body, &updatedIntent)
		require.NoError(t, err, "Failed to parse response JSON")
		assert.Equal(t, "active", updatedIntent["status"], "Intent status should be active")

		utils.SuccessLog(t, "Successfully toggled intent status")
	})

	// Test toggling saved status
	t.Run("ToggleIntentSaved", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodPost, nil)
		require.NoError(t, err, "Failed to toggle intent saved status")
		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))

		var updatedIntent map[string]interface{}
		err = json.Unmarshal(body, &updatedIntent)
		require.NoError(t, err, "Failed to parse response JSON")

		// By default, saved is false, so the first toggle sets it to true
		assert.Equal(t, true, updatedIntent["saved"], "Intent saved status should be true")

		// Toggle it back
		resp, body, err = utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodPost, nil)
		require.NoError(t, err, "Failed to toggle intent saved status back")
		require.Equal(t, http.StatusOK, resp.StatusCode)

		err = json.Unmarshal(body, &updatedIntent)
		require.NoError(t, err, "Failed to parse response JSON")
		assert.Equal(t, false, updatedIntent["saved"], "Intent saved status should be false")

		utils.SuccessLog(t, "Successfully toggled intent saved status")
	})

	// Test intent deletion
	t.Run("DeleteIntent", func(t *testing.T) {
		resp, _, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodDelete, nil)
		require.NoError(t, err, "Failed to delete intent")
		require.Equal(t, http.StatusNoContent, resp.StatusCode, "Expected status code %d, got %d",
			http.StatusNoContent, resp.StatusCode)

		// Verify the intent is no longer retrievable
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodGet, nil)
		require.NoError(t, err, "Request to get deleted intent failed")

		var retrievedIntents []map[string]interface{}
		err = json.Unmarshal(body, &retrievedIntents)
		require.NoError(t, err, "Failed to parse response JSON")

		// API returns empty array for nonexistent intent
		assert.Equal(t, 0, len(retrievedIntents), "Expected no intents to be returned after deletion")

		utils.SuccessLog(t, "Successfully deleted intent with ID: %s", intentId)
	})
}

// Test querying intents by address
func TestIntentQueryByAddress(t *testing.T) {
	// Create a test intent with a specific address
	intent := utils.TestIntent{
		ChainId: 1,
		From:    utils.TestAddress,
		Inputs: []map[string]any{
			{
				"action":   "transfer",
				"protocol": "assert",
			},
		},
		Options: struct {
			IsEOA    bool `json:"isEOA"`
			Simulate bool `json:"simulate"`
			Submit   bool `json:"submit"`
		}{
			IsEOA:    true,
			Simulate: true,
			Submit:   false,
		},
	}

	// Mark it as saved directly (options field would work too)
	intentWithSaved := map[string]interface{}{
		"chainId": intent.ChainId,
		"from":    intent.From,
		"inputs":  intent.Inputs,
		"options": map[string]interface{}{
			"isEOA":    true,
			"simulate": true,
			"submit":   false,
			"save":     true,
		},
	}

	// Create the intent
	resp, body, err := utils.MakeTestRequest("http://localhost:8080/solver/save", http.MethodPost, intentWithSaved)
	require.NoError(t, err, "Failed to create intent")
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var createdIntent map[string]interface{}
	err = json.Unmarshal(body, &createdIntent)
	require.NoError(t, err)

	intentId, ok := createdIntent["id"].(string)
	require.True(t, ok, "Intent ID not found in response")
	require.NotEmpty(t, intentId)

	// Clean up after the test
	defer func() {
		utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodDelete, nil)
	}()

	// Confirm intent was created with saved=true
	assert.Equal(t, true, createdIntent["saved"], "Intent should be saved")

	// Query by user address
	resp, body, err = utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", utils.TestAddress), http.MethodGet, nil)
	require.NoError(t, err, "Failed to query intents by address")
	require.Equal(t, http.StatusOK, resp.StatusCode)

	var retrievedIntents []map[string]interface{}
	err = json.Unmarshal(body, &retrievedIntents)
	require.NoError(t, err)

	// Verify our intent is in the results
	found := false
	for _, retrievedIntent := range retrievedIntents {
		if retrievedIntent["id"].(string) == intentId {
			found = true
			break
		}
	}
	assert.True(t, found, "Created intent should be found when querying by address")

	utils.SuccessLog(t, "Successfully retrieved intent by address")
}

// Test intent with scheduling parameters
func TestIntentWithSchedule(t *testing.T) {
	// Create a new intent with scheduling parameters
	startAt := time.Now().Add(time.Hour)
	endAt := startAt.Add(24 * time.Hour * 7) // One week from start

	intentWithSchedule := map[string]interface{}{
		"chainId":   1,
		"from":      utils.TestAddress,
		"frequency": 1, // Daily frequency
		"startAt":   startAt.Format(time.RFC3339),
		"endAt":     endAt.Format(time.RFC3339),
		"status":    "active",
		"inputs": []map[string]interface{}{
			{
				"action":   "transfer",
				"protocol": "assert",
			},
		},
		"options": map[string]interface{}{
			"isEOA":    true,
			"simulate": true,
			"submit":   false,
		},
	}

	// Send the intent creation request
	resp, body, err := utils.MakeTestRequest("http://localhost:8080/solver/save", http.MethodPost, intentWithSchedule)
	if err != nil {
		t.Fatalf("Failed to create intent with schedule: %v", err)
	}

	// Ensure the request was successful
	require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code %d, got %d. Response: %s",
		http.StatusOK, resp.StatusCode, string(body))

	// Parse the response to get the intent ID
	var createdIntent map[string]interface{}
	err = json.Unmarshal(body, &createdIntent)
	require.NoError(t, err, "Failed to parse response JSON")

	// Verify that we received an intent ID
	intentId, ok := createdIntent["id"].(string)
	require.True(t, ok, "Intent ID not found in response")
	require.NotEmpty(t, intentId, "Intent ID should not be empty")

	// Verify that scheduling parameters were saved correctly
	startAtStr, ok := createdIntent["startAt"].(string)
	require.True(t, ok, "startAt not found in response")
	parsedStartAt, err := time.Parse(time.RFC3339, startAtStr)
	require.NoError(t, err, "Failed to parse startAt time")
	assert.WithinDuration(t, startAt, parsedStartAt, time.Second, "startAt time mismatch")

	endAtStr, ok := createdIntent["endAt"].(string)
	require.True(t, ok, "endAt not found in response")
	parsedEndAt, err := time.Parse(time.RFC3339, endAtStr)
	require.NoError(t, err, "Failed to parse endAt time")
	assert.WithinDuration(t, endAt, parsedEndAt, time.Second, "endAt time mismatch")

	// Check if periodEndAt is set correctly (should be startAt + frequency)
	periodEndAtStr, ok := createdIntent["periodEndAt"].(string)
	require.True(t, ok, "periodEndAt not found in response")
	parsedPeriodEndAt, err := time.Parse(time.RFC3339, periodEndAtStr)
	require.NoError(t, err, "Failed to parse periodEndAt")
	expectedPeriodEndAt := startAt.Add(24 * time.Hour) // 1 day frequency
	assert.WithinDuration(t, expectedPeriodEndAt, parsedPeriodEndAt, time.Second, "periodEndAt time mismatch")

	utils.SuccessLog(t, "Successfully created and verified intent with schedule, ID: %s", intentId)

	// Clean up by deleting the intent
	resp, _, err = utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodDelete, nil)
	require.NoError(t, err, "Failed to delete scheduled intent")
	require.Equal(t, http.StatusNoContent, resp.StatusCode)
}

// Test creating an intent with invalid parameters
func TestInvalidIntentCreation(t *testing.T) {
	// Test cases for invalid intent creation
	testCases := []struct {
		name   string
		intent map[string]interface{}
	}{
		{
			name: "Missing ChainId",
			intent: map[string]interface{}{
				"from": utils.TestAddress,
				"inputs": []map[string]interface{}{
					{
						"action":   "transfer",
						"protocol": "assert",
					},
				},
			},
		},
		{
			name: "Invalid ChainId",
			intent: map[string]interface{}{
				"chainId": "not-a-number",
				"from":    utils.TestAddress,
				"inputs": []map[string]interface{}{
					{
						"action":   "transfer",
						"protocol": "assert",
					},
				},
			},
		},
		{
			name: "Missing From",
			intent: map[string]interface{}{
				"chainId": 1,
				"inputs": []map[string]interface{}{
					{
						"action":   "transfer",
						"protocol": "assert",
					},
				},
			},
		},
		{
			name: "Invalid StartAt Format",
			intent: map[string]interface{}{
				"chainId": 1,
				"from":    utils.TestAddress,
				"startAt": "invalid-date-format",
				"inputs": []map[string]interface{}{
					{
						"action":   "transfer",
						"protocol": "assert",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, body, err := utils.MakeTestRequest("http://localhost:8080/solver/save", http.MethodPost, tc.intent)
			if err != nil {
				t.Fatalf("Request failed: %v", err)
			}

			// Use our custom red error assertions
			if !utils.ErrorEqual(t, http.StatusBadRequest, resp.StatusCode, "Invalid intent should return 400") {
				return
			}

			// Verify an error response came back
			var errorResponse map[string]interface{}
			err = json.Unmarshal(body, &errorResponse)
			if err != nil {
				t.Error(utils.RedText(fmt.Sprintf("Failed to parse error response: %v", err)))
				return
			}

			// Check that there's an error field
			if !utils.ErrorContains(t, errorResponse, "error", "Response should contain error field") {
				return
			}

			utils.SuccessLog(t, "Correctly rejected invalid intent: %s", tc.name)
		})
	}
}

// Test nonexistent intent API error cases
func TestNonexistentIntentErrors(t *testing.T) {
	nonexistentId := "12345678-1234-1234-1234-123456789012"

	// Test toggle status on nonexistent intent
	t.Run("ToggleStatusNonexistent", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s/status", nonexistentId), http.MethodPost, nil)
		require.NoError(t, err)

		// Use our custom error assertions that highlight failures in red
		if !utils.ErrorEqual(t, http.StatusNotFound, resp.StatusCode, "Should return not found status code") {
			return
		}

		var errorResponse map[string]interface{}
		err = json.Unmarshal(body, &errorResponse)
		require.NoError(t, err)

		if !utils.ErrorContains(t, errorResponse, "error", "Error response missing 'error' field") {
			return
		}

		errorMsg, ok := errorResponse["error"].(string)
		if !ok {
			t.Error(utils.RedText("Error field is not a string"))
			return
		}

		if !utils.ErrorContains(t, errorMsg, "failed to find intent", "Unexpected error message") {
			return
		}

		utils.SuccessLog(t, "Correctly returned 404 for toggling status on nonexistent intent")
	})

	// Test toggle saved on nonexistent intent
	t.Run("ToggleSavedNonexistent", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", nonexistentId), http.MethodPost, nil)
		require.NoError(t, err)

		// Use our custom error assertions that highlight failures in red
		if !utils.ErrorEqual(t, http.StatusNotFound, resp.StatusCode, "Should return not found status code") {
			return
		}

		var errorResponse map[string]interface{}
		err = json.Unmarshal(body, &errorResponse)
		require.NoError(t, err)

		if !utils.ErrorContains(t, errorResponse, "error", "Error response missing 'error' field") {
			return
		}

		errorMsg, ok := errorResponse["error"].(string)
		if !ok {
			t.Error(utils.RedText("Error field is not a string"))
			return
		}

		if !utils.ErrorContains(t, errorMsg, "failed to find intent", "Unexpected error message") {
			return
		}

		utils.SuccessLog(t, "Correctly returned 404 for toggling saved on nonexistent intent")
	})

	// Test delete nonexistent intent
	t.Run("DeleteNonexistent", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", nonexistentId), http.MethodDelete, nil)
		require.NoError(t, err)

		// Use our custom error assertions that highlight failures in red
		if !utils.ErrorEqual(t, http.StatusNotFound, resp.StatusCode, "Should return not found status code") {
			return
		}

		var errorResponse map[string]interface{}
		err = json.Unmarshal(body, &errorResponse)
		require.NoError(t, err)

		if !utils.ErrorContains(t, errorResponse, "error", "Error response missing 'error' field") {
			return
		}

		errorMsg, ok := errorResponse["error"].(string)
		if !ok {
			t.Error(utils.RedText("Error field is not a string"))
			return
		}

		if !utils.ErrorContains(t, errorMsg, "failed to find intent", "Unexpected error message") {
			return
		}

		utils.SuccessLog(t, "Correctly returned 404 for deleting nonexistent intent")
	})
}
