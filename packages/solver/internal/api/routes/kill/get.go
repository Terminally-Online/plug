package kill

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/routes"
	"solver/internal/solver"

	"github.com/swaggest/openapi-go"
)

type KillResponse struct {
	Killed bool `json:"killed" description:"Whether the kill switch is active"`
}

func GetContext(oc openapi.OperationContext) error {
	oc.SetTags("System")
	oc.SetSummary("Get Kill Switch Status")
	oc.SetDescription("Returns the current status of the kill switch, which controls whether new solver operations are allowed")

	oc.AddRespStructure(KillResponse{}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "Internal server error"},
		openapi.WithContentType("text/plain"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func GetRequest(w http.ResponseWriter, r *http.Request, s *solver.Solver) {
	response := KillResponse{Killed: s.IsKilled}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Get(s *solver.Solver) *routes.RouteHandler {
	return routes.NewRouteHandler(GetRequest, GetContext, s)
}
