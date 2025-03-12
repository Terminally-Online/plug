package solver

import (
	"encoding/json"
	"net/http"

	"github.com/swaggest/openapi-go"
)

// KillResponse defines the structure for kill switch responses
type KillResponse struct {
	Killed bool `json:"killed" description:"Whether the kill switch is active"`
}

// GetKill returns the current status of the kill switch
func (h *Handler) GetKill(w http.ResponseWriter, r *http.Request) {
	response := KillResponse{Killed: h.Solver.IsKilled}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// PostKill toggles the kill switch status
func (h *Handler) PostKill(w http.ResponseWriter, r *http.Request) {
	h.Solver.IsKilled = !h.Solver.IsKilled

	response := KillResponse{Killed: h.Solver.IsKilled}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// SetupOpenAPIForGetKill defines the OpenAPI documentation for the GET /solver/kill endpoint
func SetupOpenAPIForGetKill(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("System")
	oc.SetSummary("Get Kill Switch Status")
	oc.SetDescription("Returns the current status of the kill switch, which controls whether new solver operations are allowed")
	
	// Define success response
	oc.AddRespStructure(KillResponse{}, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error response
	oc.AddRespStructure(
		map[string]string{"error": "Internal server error"},
		openapi.WithContentType("text/plain"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

// SetupOpenAPIForPostKill defines the OpenAPI documentation for the POST /solver/kill endpoint
func SetupOpenAPIForPostKill(oc openapi.OperationContext) error {
	// Set operation tags, summary, and description
	oc.SetTags("System")
	oc.SetSummary("Toggle Kill Switch")
	oc.SetDescription("Toggles the state of the kill switch that controls whether new solver operations are allowed")
	
	// Define success response
	oc.AddRespStructure(KillResponse{}, openapi.WithHTTPStatus(http.StatusOK))
	
	// Define error response
	oc.AddRespStructure(
		map[string]string{"error": "Internal server error"},
		openapi.WithContentType("text/plain"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}