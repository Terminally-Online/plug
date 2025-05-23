package api_key

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/routes"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"solver/internal/utils"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/swaggest/openapi-go"
)

func UpdateContext(oc openapi.OperationContext) error {
	oc.SetTags("API Keys")
	oc.SetSummary("Update API Key")
	oc.SetDescription("Updates an existing API key with the provided parameters. Requires admin privileges.")

	oc.AddReqStructure(struct {
		ID string `path:"id" description:"API key ID"`
	}{})

	oc.AddReqStructure(ApiKeyCreateRequest{})

	oc.AddRespStructure(nil, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "failed to find api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)
	oc.AddRespStructure(
		map[string]string{"error": "invalid request body"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)
	oc.AddRespStructure(
		map[string]string{"error": "failed to update api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func UpdateRequest(w http.ResponseWriter, r *http.Request, _ *redis.Client, s *solver.Solver) {
	vars := mux.Vars(r)
	id := vars["id"]

	var apiKey models.ApiKey
	if err := database.DB.First(&apiKey, "id = ?", id).Error; err != nil {
		utils.RespondWithError(w, utils.ErrUnauthorized("api key does not exist"))
		return
	}

	var newApiKey models.ApiKey
	if err := json.NewDecoder(r.Body).Decode(&newApiKey); err != nil {
		utils.RespondWithError(w, utils.ErrInvalidRequestBody(err))
		return
	}

	if err := database.DB.Model(&apiKey).Updates(newApiKey).Error; err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to update api key: "+err.Error()))
		return
	}
}

func Update() *routes.RouteHandler {
	return routes.NewRouteHandler(UpdateRequest, UpdateContext, nil, nil)
}
