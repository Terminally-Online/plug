package cache

import (
	"context"
	"encoding/json"
	"net/http"
	"solver/internal/api/routes"
	redisInternal "solver/internal/redis"
	"solver/internal/solver"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/swaggest/openapi-go"
)

type CacheClearResponse struct {
	Status  string `json:"status" description:"The status of the operation"`
	Message string `json:"message" description:"A message describing the result"`
}

// ClearContext sets up the OpenAPI context for the cache clear endpoint
func ClearContext(oc openapi.OperationContext) error {
	oc.SetTags("System")
	oc.SetSummary("Clear Redis Cache")
	oc.SetDescription("Clears all keys from the Redis cache database (development use only)")

	oc.AddRespStructure(CacheClearResponse{
		Status:  "success",
		Message: "Cache cleared successfully",
	}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "Internal server error"},
		openapi.WithContentType("text/plain"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

// ClearRequest handles the actual request to clear the cache
func ClearRequest(w http.ResponseWriter, r *http.Request, redisClient *redis.Client, _ *solver.Solver) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Clear the cache
	err := redisInternal.ClearCache(ctx)
	if err != nil {
		http.Error(w, "Failed to clear cache: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	response := CacheClearResponse{
		Status:  "success",
		Message: "Cache cleared successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Clear returns a RouteHandler for clearing the Redis cache
func Clear() *routes.RouteHandler {
	return routes.NewRouteHandler(ClearRequest, ClearContext, nil, nil)
}
