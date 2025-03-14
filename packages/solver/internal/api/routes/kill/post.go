package kill

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/routes"
	"solver/internal/solver"

	"github.com/swaggest/openapi-go"
)

func PostContext(oc openapi.OperationContext) error {
	oc.SetTags("System")
	oc.SetSummary("Toggle Kill Switch")
	oc.SetDescription("Toggles the state of the kill switch that controls whether new solver operations are allowed")

	oc.AddRespStructure(KillResponse{}, openapi.WithHTTPStatus(http.StatusOK))

	oc.AddRespStructure(
		map[string]string{"error": "Internal server error"},
		openapi.WithContentType("text/plain"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func PostRequest(w http.ResponseWriter, r *http.Request, s *solver.Solver) {
	s.IsKilled = !s.IsKilled

	response := KillResponse{Killed: s.IsKilled}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Post(s *solver.Solver) *routes.RouteHandler {
	return routes.NewRouteHandler(PostRequest, PostContext, s)
}
