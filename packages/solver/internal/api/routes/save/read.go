package save

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/routes"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"solver/internal/utils"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/swaggest/openapi-go"
)

type IntentListResponse []models.Intent

func ReadContext(oc openapi.OperationContext) error {
	oc.SetTags("Intents")
	oc.SetSummary("Get Intent")
	oc.SetDescription("Retrieves an intent or list of intents by ID or address")

	oc.AddReqStructure(struct {
		ID string `path:"id" description:"Intent ID, address, or comma-separated list of addresses"`
	}{})

	oc.AddRespStructure(IntentListResponse{}, openapi.WithHTTPStatus(http.StatusOK))

	oc.AddRespStructure(
		map[string]string{"error": "database error"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func ReadRequest(w http.ResponseWriter, r *http.Request, _ *redis.Client, s *solver.Solver) {
	vars := mux.Vars(r)
	id := vars["id"]

	var intents []models.Intent
	query := database.DB
	if strings.Contains(id, ",") {
		addresses := strings.Split(id, ",")
		query = query.Where("\"from\" IN ? AND saved = ?", addresses, true)
	} else if strings.HasPrefix(id, "0x") {
		query = query.Where("\"from\" = ? AND saved = ?", id, true)
	} else {
		query = query.Where("id = ?", id)
	}
	result := query.
		Order("created_at desc").
		Find(&intents)
	if result.Error != nil {
		utils.RespondWithError(w, utils.ErrInternal("database error: "+result.Error.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(intents); err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to encode response: "+err.Error()))
		return
	}
}

func Read() *routes.RouteHandler {
	return routes.NewRouteHandler(ReadRequest, ReadContext, nil, nil)
}
