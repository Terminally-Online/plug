package health

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/routes"
	"solver/internal/solver"

	"github.com/go-redis/redis/v8"
	"github.com/swaggest/openapi-go"
)

type HealthResponse struct {
	Status string `json:"status" description:"The health status of the API"`
}

func GetContext(oc openapi.OperationContext) error {
	oc.SetTags("System")
	oc.SetSummary("Health Check")
	oc.SetDescription("Returns the health status of the API")

	oc.AddRespStructure(HealthResponse{
		Status: "healthy",
	}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "Internal server error"},
		openapi.WithContentType("text/plain"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func GetRequest(w http.ResponseWriter, r *http.Request, _ *redis.Client, _ *solver.Solver) {
	if err := json.NewEncoder(w).Encode(HealthResponse{
		Status: "healthy",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func GetHealth() *routes.RouteHandler {
	return routes.NewRouteHandler(GetRequest, GetContext, nil, nil)
}
