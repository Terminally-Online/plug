package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

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
	if resp.StatusCode != http.StatusOK {
		utils.FailTest(t, "Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))
		return
	}

	// Parse the response to get the intent ID
	var createdIntent map[string]interface{}
	err = json.Unmarshal(body, &createdIntent)
	if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
		return
	}

	// Verify that we received an intent ID
	intentId, ok := createdIntent["id"].(string)
	if !ok {
		utils.FailTest(t, "Intent ID not found in response")
		return
	}
	if intentId == "" {
		utils.FailTest(t, "Intent ID should not be empty")
		return
	}

	utils.SuccessLog(t, "Successfully created intent with ID: %s", intentId)

	// Test intent retrieval
	t.Run("GetIntent", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodGet, nil)
		if !utils.CheckNoError(t, err, "Failed to get intent") {
			return
		}
		if resp.StatusCode != http.StatusOK {
			utils.FailTest(t, "Expected status code %d, got %d. Response: %s",
				http.StatusOK, resp.StatusCode, string(body))
			return
		}

		// The API returns an array of intents
		var retrievedIntents []map[string]interface{}
		err = json.Unmarshal(body, &retrievedIntents)
		if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
			return
		}
		if len(retrievedIntents) <= 0 {
			utils.FailTest(t, "Expected at least one intent in response")
			return
		}

		retrievedIntent := retrievedIntents[0]
		// Verify the intent data matches what we created
		if retrievedIntent["id"] != intentId {
			utils.FailTest(t, "Intent ID mismatch. Expected: %s, got: %v", intentId, retrievedIntent["id"])
			return
		}

		if retrievedIntent["chainId"] != float64(1) {
			utils.FailTest(t, "Chain ID mismatch. Expected: %v, got: %v", float64(1), retrievedIntent["chainId"])
			return
		}

		if retrievedIntent["from"] != utils.TestAddress {
			utils.FailTest(t, "From address mismatch. Expected: %s, got: %v", utils.TestAddress, retrievedIntent["from"])
			return
		}

		utils.SuccessLog(t, "Successfully retrieved intent with ID: %s", intentId)
	})

	// Test updating intent status
	t.Run("ToggleIntentStatus", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s/status", intentId), http.MethodPost, nil)
		if !utils.CheckNoError(t, err, "Failed to toggle intent status") {
			return
		}
		if resp.StatusCode != http.StatusOK {
			utils.FailTest(t, "Expected status code %d, got %d. Response: %s",
				http.StatusOK, resp.StatusCode, string(body))
			return
		}

		var updatedIntent map[string]interface{}
		err = json.Unmarshal(body, &updatedIntent)
		if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
			return
		}

		// Verify the status was toggled from "active" to "paused"
		if updatedIntent["status"] != "paused" {
			utils.FailTest(t, "Intent status should be paused. Expected: %s, got: %v", "paused", updatedIntent["status"])
			return
		}

		// Toggle it back to active
		resp, body, err = utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s/status", intentId), http.MethodPost, nil)
		if !utils.CheckNoError(t, err, "Failed to toggle intent status back") {
			return
		}
		if !utils.CheckEqual(t, http.StatusOK, resp.StatusCode, "Equality check failed") {
			return
		}

		err = json.Unmarshal(body, &updatedIntent)
		if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
			return
		}
		if updatedIntent["status"] != "active" {
			utils.FailTest(t, "Intent status should be active. Expected: %s, got: %v", "active", updatedIntent["status"])
			return
		}

		utils.SuccessLog(t, "Successfully toggled intent status")
	})

	// Test toggling saved status
	t.Run("ToggleIntentSaved", func(t *testing.T) {
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodPost, nil)
		if !utils.CheckNoError(t, err, "Failed to toggle intent saved status") {
			return
		}
		if resp.StatusCode != http.StatusOK {
			utils.FailTest(t, "Expected status code %d, got %d. Response: %s",
				http.StatusOK, resp.StatusCode, string(body))
			return
		}

		var updatedIntent map[string]interface{}
		err = json.Unmarshal(body, &updatedIntent)
		if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
			return
		}

		if updatedIntent["saved"] != true {
			utils.FailTest(t, "Intent saved status should be true. Expected: %v, got: %v", true, updatedIntent["saved"])
			return
		}

		// Toggle it back
		resp, body, err = utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodPost, nil)
		if !utils.CheckNoError(t, err, "Failed to toggle intent saved status back") {
			return
		}
		if !utils.CheckEqual(t, http.StatusOK, resp.StatusCode, "Equality check failed") {
			return
		}

		err = json.Unmarshal(body, &updatedIntent)
		if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
			return
		}
		if updatedIntent["saved"] != false {
			utils.FailTest(t, "Intent saved status should be false. Expected: %v, got: %v", false, updatedIntent["saved"])
			return
		}

		utils.SuccessLog(t, "Successfully toggled intent saved status")
	})

	// Test intent deletion
	t.Run("DeleteIntent", func(t *testing.T) {
		resp, _, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodDelete, nil)
		if !utils.CheckNoError(t, err, "Failed to delete intent") {
			return
		}
		if resp.StatusCode != http.StatusNoContent {
			utils.FailTest(t, "Expected status code %d, got %d",
				http.StatusNoContent, resp.StatusCode)
			return
		}

		// Verify the intent is no longer retrievable
		resp, body, err := utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodGet, nil)
		if !utils.CheckNoError(t, err, "Request to get deleted intent failed") {
			return
		}

		var retrievedIntents []map[string]interface{}
		err = json.Unmarshal(body, &retrievedIntents)
		if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
			return
		}

		// API returns empty array for nonexistent intent
		if len(retrievedIntents) != 0 {
			utils.FailTest(t, "Expected no intents to be returned after deletion. Expected: 0, got: %d", len(retrievedIntents))
			return
		}

		utils.SuccessLog(t, "Successfully deleted intent with ID: %s", intentId)
	})

	t.Run("NextSimulationAt", func(t *testing.T) {
		// Create a new intent with a frequency
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
		if resp.StatusCode != http.StatusOK {
			utils.FailTest(t, "Expected status code %d, got %d. Response: %s",
				http.StatusOK, resp.StatusCode, string(body))
			return
		}

		// Parse the response to get the intent ID
		var createdIntent map[string]interface{}
		err = json.Unmarshal(body, &createdIntent)
		if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
			return
		}

		// Verify that we received an intent ID
		intentId, ok := createdIntent["id"].(string)
		if !ok {
			utils.FailTest(t, "Intent ID not found in response")
			return
		}
		if intentId == "" {
			utils.FailTest(t, "Intent ID should not be empty")
			return
		}

		// Verify that scheduling parameters were saved correctly
		startAtStr, ok := createdIntent["startAt"].(string)
		if !ok {
			utils.FailTest(t, "startAt not found in response")
			return
		}
		parsedStartAt, err := time.Parse(time.RFC3339, startAtStr)
		if !utils.CheckNoError(t, err, "Failed to parse startAt time") {
			return
		}
		if parsedStartAt.Sub(startAt) > time.Second || startAt.Sub(parsedStartAt) > time.Second {
			utils.FailTest(t, "startAt time mismatch. Expected within 1 second of: %v, got: %v",
				startAt.Format(time.RFC3339), parsedStartAt.Format(time.RFC3339))
			return
		}

		endAtStr, ok := createdIntent["endAt"].(string)
		if !ok {
			utils.FailTest(t, "endAt not found in response")
			return
		}
		parsedEndAt, err := time.Parse(time.RFC3339, endAtStr)
		if !utils.CheckNoError(t, err, "Failed to parse endAt time") {
			return
		}
		if parsedEndAt.Sub(endAt) > time.Second || endAt.Sub(parsedEndAt) > time.Second {
			utils.FailTest(t, "endAt time mismatch. Expected within 1 second of: %v, got: %v",
				endAt.Format(time.RFC3339), parsedEndAt.Format(time.RFC3339))
			return
		}

		// Check if periodEndAt is set correctly (should be startAt + frequency)
		periodEndAtStr, ok := createdIntent["periodEndAt"].(string)
		if !ok {
			utils.FailTest(t, "periodEndAt not found in response")
			return
		}
		parsedPeriodEndAt, err := time.Parse(time.RFC3339, periodEndAtStr)
		if !utils.CheckNoError(t, err, "Failed to parse periodEndAt") {
			return
		}
		expectedPeriodEndAt := startAt.Add(24 * time.Hour) // 1 day frequency
		if parsedPeriodEndAt.Sub(expectedPeriodEndAt) > time.Second || expectedPeriodEndAt.Sub(parsedPeriodEndAt) > time.Second {
			utils.FailTest(t, "periodEndAt time mismatch. Expected within 1 second of: %v, got: %v",
				expectedPeriodEndAt.Format(time.RFC3339), parsedPeriodEndAt.Format(time.RFC3339))
			return
		}

		// ensure next simulation at is set correctly
		nextSimulationAtStr, ok := createdIntent["nextSimulationAt"].(string)
		if !ok {
			utils.FailTest(t, "nextSimulationAt not found in response")
			return
		}
		parsedNextSimulationAt, err := time.Parse(time.RFC3339, nextSimulationAtStr)
		if !utils.CheckNoError(t, err, "Failed to parse nextSimulationAt") {
			return
		}
		if parsedNextSimulationAt.Sub(startAt) > time.Second || startAt.Sub(parsedNextSimulationAt) > time.Second {
			utils.FailTest(t, "nextSimulationAt time mismatch. Expected within 1 second of: %v, got: %v",
				startAt.Format(time.RFC3339), parsedNextSimulationAt.Format(time.RFC3339))
			return
		}

		utils.SuccessLog(t, "Successfully created and verified intent with schedule, ID: %s", intentId)
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
	if !utils.CheckNoError(t, err, "Failed to create intent") {
		return
	}
	if !utils.CheckEqual(t, http.StatusOK, resp.StatusCode, "Equality check failed") {
		return
	}

	var createdIntent map[string]interface{}
	err = json.Unmarshal(body, &createdIntent)
	if !utils.CheckNoError(t, err, "Unexpected error") {
		return
	}

	intentId, ok := createdIntent["id"].(string)
	if !ok {
		utils.FailTest(t, "Intent ID not found in response")
		return
	}
	if intentId == "" {
		utils.FailTest(t, "Value should not be empty")
		return
	}

	// Clean up after the test
	defer func() {
		utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodDelete, nil)
	}()

	// Confirm intent was created with saved=true
	if createdIntent["saved"] != true {
		utils.FailTest(t, "Intent should be saved. Expected: %v, got: %v", true, createdIntent["saved"])
		return
	}

	// Query by user address
	resp, body, err = utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", utils.TestAddress), http.MethodGet, nil)
	if !utils.CheckNoError(t, err, "Failed to query intents by address") {
		return
	}
	if !utils.CheckEqual(t, http.StatusOK, resp.StatusCode, "Equality check failed") {
		return
	}

	var retrievedIntents []map[string]interface{}
	err = json.Unmarshal(body, &retrievedIntents)
	if !utils.CheckNoError(t, err, "Unexpected error") {
		return
	}

	// Verify our intent is in the results
	found := false
	for _, retrievedIntent := range retrievedIntents {
		if retrievedIntent["id"].(string) == intentId {
			found = true
			break
		}
	}
	if !found {
		utils.FailTest(t, "Created intent should be found when querying by address")
		return
	}

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
	if resp.StatusCode != http.StatusOK {
		utils.FailTest(t, "Expected status code %d, got %d. Response: %s",
			http.StatusOK, resp.StatusCode, string(body))
		return
	}

	// Parse the response to get the intent ID
	var createdIntent map[string]interface{}
	err = json.Unmarshal(body, &createdIntent)
	if !utils.CheckNoError(t, err, "Failed to parse response JSON") {
		return
	}

	// Verify that we received an intent ID
	intentId, ok := createdIntent["id"].(string)
	if !ok {
		utils.FailTest(t, "Intent ID not found in response")
		return
	}
	if intentId == "" {
		utils.FailTest(t, "Intent ID should not be empty")
		return
	}

	// Verify that scheduling parameters were saved correctly
	startAtStr, ok := createdIntent["startAt"].(string)
	if !ok {
		utils.FailTest(t, "startAt not found in response")
		return
	}
	parsedStartAt, err := time.Parse(time.RFC3339, startAtStr)
	if !utils.CheckNoError(t, err, "Failed to parse startAt time") {
		return
	}
	if parsedStartAt.Sub(startAt) > time.Second || startAt.Sub(parsedStartAt) > time.Second {
		utils.FailTest(t, "startAt time mismatch. Expected within 1 second of: %v, got: %v",
			startAt.Format(time.RFC3339), parsedStartAt.Format(time.RFC3339))
		return
	}

	endAtStr, ok := createdIntent["endAt"].(string)
	if !ok {
		utils.FailTest(t, "endAt not found in response")
		return
	}
	parsedEndAt, err := time.Parse(time.RFC3339, endAtStr)
	if !utils.CheckNoError(t, err, "Failed to parse endAt time") {
		return
	}
	if parsedEndAt.Sub(endAt) > time.Second || endAt.Sub(parsedEndAt) > time.Second {
		utils.FailTest(t, "endAt time mismatch. Expected within 1 second of: %v, got: %v",
			endAt.Format(time.RFC3339), parsedEndAt.Format(time.RFC3339))
		return
	}

	// Check if periodEndAt is set correctly (should be startAt + frequency)
	periodEndAtStr, ok := createdIntent["periodEndAt"].(string)
	if !ok {
		utils.FailTest(t, "periodEndAt not found in response")
		return
	}
	parsedPeriodEndAt, err := time.Parse(time.RFC3339, periodEndAtStr)
	if !utils.CheckNoError(t, err, "Failed to parse periodEndAt") {
		return
	}
	expectedPeriodEndAt := startAt.Add(24 * time.Hour) // 1 day frequency
	if parsedPeriodEndAt.Sub(expectedPeriodEndAt) > time.Second || expectedPeriodEndAt.Sub(parsedPeriodEndAt) > time.Second {
		utils.FailTest(t, "periodEndAt time mismatch. Expected within 1 second of: %v, got: %v",
			expectedPeriodEndAt.Format(time.RFC3339), parsedPeriodEndAt.Format(time.RFC3339))
		return
	}

	utils.SuccessLog(t, "Successfully created and verified intent with schedule, ID: %s", intentId)

	// Clean up by deleting the intent
	resp, _, err = utils.MakeTestRequest(fmt.Sprintf("http://localhost:8080/solver/save/%s", intentId), http.MethodDelete, nil)
	if !utils.CheckNoError(t, err, "Failed to delete scheduled intent") {
		return
	}
	if !utils.CheckEqual(t, http.StatusNoContent, resp.StatusCode, "Equality check failed") {
		return
	}
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
				utils.FailTest(t, "Failed to parse error response: %v", err)
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
		if !utils.CheckNoError(t, err, "Unexpected error") {
			return
		}

		// Use our custom error assertions that highlight failures in red
		if !utils.ErrorEqual(t, http.StatusNotFound, resp.StatusCode, "Should return not found status code") {
			return
		}

		var errorResponse map[string]interface{}
		err = json.Unmarshal(body, &errorResponse)
		if !utils.CheckNoError(t, err, "Unexpected error") {
			return
		}

		if !utils.ErrorContains(t, errorResponse, "error", "Error response missing 'error' field") {
			return
		}

		errorMsg, ok := errorResponse["error"].(string)
		if !ok {
			utils.FailTest(t, "Error field is not a string")
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
		if !utils.CheckNoError(t, err, "Unexpected error") {
			return
		}

		// Use our custom error assertions that highlight failures in red
		if !utils.ErrorEqual(t, http.StatusNotFound, resp.StatusCode, "Should return not found status code") {
			return
		}

		var errorResponse map[string]interface{}
		err = json.Unmarshal(body, &errorResponse)
		if !utils.CheckNoError(t, err, "Unexpected error") {
			return
		}

		if !utils.ErrorContains(t, errorResponse, "error", "Error response missing 'error' field") {
			return
		}

		errorMsg, ok := errorResponse["error"].(string)
		if !ok {
			utils.FailTest(t, "Error field is not a string")
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
		if !utils.CheckNoError(t, err, "Unexpected error") {
			return
		}

		// Use our custom error assertions that highlight failures in red
		if !utils.ErrorEqual(t, http.StatusNotFound, resp.StatusCode, "Should return not found status code") {
			return
		}

		var errorResponse map[string]interface{}
		err = json.Unmarshal(body, &errorResponse)
		if !utils.CheckNoError(t, err, "Unexpected error") {
			return
		}

		if !utils.ErrorContains(t, errorResponse, "error", "Error response missing 'error' field") {
			return
		}

		errorMsg, ok := errorResponse["error"].(string)
		if !ok {
			utils.FailTest(t, "Error field is not a string")
			return
		}

		if !utils.ErrorContains(t, errorMsg, "failed to find intent", "Unexpected error message") {
			return
		}

		utils.SuccessLog(t, "Correctly returned 404 for deleting nonexistent intent")
	})
}
