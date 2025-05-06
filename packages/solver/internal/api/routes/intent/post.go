package intent

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/middleware"
	"solver/internal/api/routes"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/redis"
	"solver/internal/solver"
	"solver/internal/utils"

	redisv8 "github.com/go-redis/redis/v8"
	"github.com/swaggest/openapi-go"
)

func PostContext(oc openapi.OperationContext) error {
	oc.SetTags("Solver")
	oc.SetSummary("Solve Intent")
	oc.SetDescription("Processes an intent and returns a solution (transaction data). Requires API key.")

	oc.AddReqStructure(models.Intent{})

	oc.AddRespStructure(map[string]any{}, openapi.WithHTTPStatus(http.StatusOK))
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

func PostRequest(w http.ResponseWriter, r *http.Request, c *redisv8.Client, s *solver.Solver) {
	var intent *models.Intent
	if err := json.NewDecoder(r.Body).Decode(&intent); err != nil {
		utils.RespondWithError(w, utils.ErrInvalidRequestBody(err))
		return
	}

	intent.ApiKeyId = r.Header.Get("X-Api-Key-Id")
	var err error

	intent, err = intent.GetOrCreate(database.DB)
	if err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to initialize intent: "+err.Error()))
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
		utils.RespondWithError(w, err)
	} else {
		if err := json.NewEncoder(w).Encode(solution); err != nil {
			utils.RespondWithError(w, utils.ErrInternal("failed to encode response: "+err.Error()))
		}
	}
}

func Post(s *solver.Solver) *routes.RouteHandler {
	return routes.NewRouteHandler(PostRequest, PostContext, redis.CacheRedis, s)
}
