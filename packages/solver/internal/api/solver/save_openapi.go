package solver

import (
	"net/http"
	"solver/internal/database/models"

	"github.com/swaggest/openapi-go"
)

// Simple models for OpenAPI documentation
type IntentResponse struct {
	ID      string `json:"id" description:"Unique identifier of the intent"`
	Status  string `json:"status" description:"Status of the intent (active, paused, etc.)"`
	ChainID uint64 `json:"chainId" description:"Blockchain chain ID"`
	From    string `json:"from" description:"Source address for the transaction"`
	Saved   bool   `json:"saved" description:"Whether the intent is saved"`
	Locked  bool   `json:"locked" description:"Whether the intent is locked"`
}

// IntentListResponse is the response for listing intents
type IntentListResponse []models.Intent

// SetupOpenAPIForCreateIntent defines the OpenAPI documentation for the POST /solver/save endpoint
func SetupOpenAPIForCreateIntent(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("Intents")
	oc.SetSummary("Create Intent")
	oc.SetDescription("Creates a new intent with the provided parameters")

	// Define request body
	oc.AddReqStructure(models.Intent{})

	// Define success response
	oc.AddRespStructure(models.Intent{}, openapi.WithHTTPStatus(http.StatusOK))

	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "invalid request body"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)

	oc.AddRespStructure(
		map[string]string{"error": "failed to find api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)

	oc.AddRespStructure(
		map[string]string{"error": "failed to save intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func SetupOpenAPIForReadIntents(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("Intents")
	oc.SetSummary("List Intents")
	oc.SetDescription("Retrieves a list of saved intents for the current credentials")

	// Define success response
	oc.AddRespStructure(IntentListResponse{}, openapi.WithHTTPStatus(http.StatusOK))

	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "database error"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

// SetupOpenAPIForReadIntent defines the OpenAPI documentation for the GET /solver/save/{id} endpoint
func SetupOpenAPIForReadIntent(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("Intents")
	oc.SetSummary("Get Intent")
	oc.SetDescription("Retrieves an intent or list of intents by ID or address")

	// Define path parameter
	oc.AddReqStructure(struct {
		ID string `path:"id" description:"Intent ID, address, or comma-separated list of addresses"`
	}{})

	// Define success response
	oc.AddRespStructure(IntentListResponse{}, openapi.WithHTTPStatus(http.StatusOK))

	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "database error"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

// SetupOpenAPIForToggleIntentStatus defines the OpenAPI documentation for the POST /solver/save/{id}/status endpoint
func SetupOpenAPIForToggleIntentStatus(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("Intents")
	oc.SetSummary("Toggle Intent Status")
	oc.SetDescription("Toggles the status of an intent between 'active' and 'paused'")

	// Define path parameter
	oc.AddReqStructure(struct {
		ID string `path:"id" description:"Intent ID"`
	}{})

	// Define success response
	oc.AddRespStructure(models.Intent{}, openapi.WithHTTPStatus(http.StatusOK))

	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "failed to find intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)

	oc.AddRespStructure(
		map[string]string{"error": "failed to toggle intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

// SetupOpenAPIForToggleIntentSaved defines the OpenAPI documentation for the POST /solver/save/{id} endpoint
func SetupOpenAPIForToggleIntentSaved(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("Intents")
	oc.SetSummary("Toggle Intent Saved")
	oc.SetDescription("Toggles the saved status of an intent")

	// Define path parameter
	oc.AddReqStructure(struct {
		ID string `path:"id" description:"Intent ID"`
	}{})

	// Define success response
	oc.AddRespStructure(models.Intent{}, openapi.WithHTTPStatus(http.StatusOK))

	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "failed to find intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)

	oc.AddRespStructure(
		map[string]string{"error": "failed to toggle intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

// SetupOpenAPIForDeleteIntent defines the OpenAPI documentation for the DELETE /solver/save/{id} endpoint
func SetupOpenAPIForDeleteIntent(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("Intents")
	oc.SetSummary("Delete Intent")
	oc.SetDescription("Deletes an intent by ID")

	// Define path parameter
	oc.AddReqStructure(struct {
		ID string `path:"id" description:"Intent ID"`
	}{})

	// Define success response
	oc.AddRespStructure(nil, openapi.WithHTTPStatus(http.StatusNoContent))

	// Define error responses
	oc.AddRespStructure(
		map[string]string{"error": "failed to find intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)

	oc.AddRespStructure(
		map[string]string{"error": "failed to delete intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}
