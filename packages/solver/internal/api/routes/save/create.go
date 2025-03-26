package save

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/routes"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"solver/internal/utils"

	"github.com/go-redis/redis/v8"
	"github.com/swaggest/openapi-go"
)

func CreateContext(oc openapi.OperationContext) error {
	oc.SetTags("Intents")
	oc.SetSummary("Create Intent")
	oc.SetDescription("Creates a new intent with the provided parameters")

	oc.AddReqStructure(models.Intent{})

	oc.AddRespStructure(models.Intent{}, openapi.WithHTTPStatus(http.StatusOK))
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

func CreateRequest(w http.ResponseWriter, r *http.Request, _ *redis.Client, s *solver.Solver) {
	var inputs models.Intent
	if err := json.NewDecoder(r.Body).Decode(&inputs); err != nil {
		utils.RespondWithError(w, utils.ErrInvalidRequestBody(err))
		return
	}

	err := inputs.ValidateFields()
	if err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	var apiKey models.ApiKey
	if err := database.DB.First(&apiKey, "key = ?", r.Header.Get("X-Api-Key")).Error; err != nil {
		utils.RespondWithError(w, utils.ErrUnauthorized("api key does not exist"))
		return
	}
	inputs.ApiKeyId = apiKey.Id
	if val, exists := inputs.Options["save"]; exists {
		inputs.Saved = val.(bool)
	}
	if err := database.DB.Omit("nextSimulationAt", "periodEndAt").Create(&inputs).Error; err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to save intent: "+err.Error()))
		return
	}
	if err := json.NewEncoder(w).Encode(inputs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Create() *routes.RouteHandler {
	return routes.NewRouteHandler(CreateRequest, CreateContext, nil, nil)
}
