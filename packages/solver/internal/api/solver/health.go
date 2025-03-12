package solver

import (
	"encoding/json"
	"net/http"

	"github.com/swaggest/openapi-go"
)

// HealthResponse defines the structure for health check responses.
type HealthResponse struct {
	Status string `json:"status" description:"The health status of the API"`
}

// GetHealth returns the current health status of the API
func (h *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "healthy",
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// SetupOpenAPIForHealth defines the OpenAPI documentation for the health endpoint
func SetupOpenAPIForHealth(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("System")
	oc.SetSummary("Health Check")
	oc.SetDescription("Returns the health status of the API")
	
	// Define success response
	oc.AddRespStructure(HealthResponse{
		Status: "healthy",
	}, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error response
	oc.AddRespStructure(
		map[string]string{"error": "Internal server error"},
		openapi.WithContentType("text/plain"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}