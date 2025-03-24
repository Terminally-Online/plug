package intent

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/middleware"
	"solver/internal/api/routes"
	"solver/internal/cache"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"solver/internal/utils"

	"github.com/go-redis/redis/v8"
	"github.com/swaggest/openapi-go"
)

func PostContext(oc openapi.OperationContext) error {
	oc.SetTags("Solver")
	oc.SetSummary("Solve Intent")
	oc.SetDescription("Processes an intent and returns a solution (transaction data). Requires API key.")

	oc.AddReqStructure(models.Intent{})

	oc.AddRespStructure(map[string]interface{}{}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "invalid request body"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)
	oc.AddRespStructure(
		map[string]string{"error": "failed to initialize intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)
	oc.AddRespStructure(
		map[string]string{"error": "failed to solve intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func PostRequest(w http.ResponseWriter, r *http.Request, c *redis.Client, s *solver.Solver) {
	var intent *models.Intent
	if err := json.NewDecoder(r.Body).Decode(&intent); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	intent.ApiKeyId = r.Header.Get("X-Api-Key-Id")
	var err error

	intent, err = intent.GetOrCreate(database.DB)
	if err != nil {
		utils.MakeHttpError(w, "failed to initialize intent: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Wrap solver call with metrics tracking
	var solution interface{}
	err = middleware.TrackSolverMetrics(func() error {
		var solveErr error
		solution, solveErr = s.Solve(intent, true, false)
		return solveErr
	})
	
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(solution); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
	}
}

func Post(s *solver.Solver) *routes.RouteHandler {
	return routes.NewRouteHandler(PostRequest, PostContext, cache.Redis, s)
}
