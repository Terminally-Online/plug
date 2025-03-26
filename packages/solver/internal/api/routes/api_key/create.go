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
	"github.com/swaggest/openapi-go"
)

type ApiKeyCreateRequest struct {
	SocketID  *string `json:"socketId,omitempty" description:"ID of the socket to associate with this key"`
	RateLimit int     `json:"rateLimit" description:"Rate limit for this API key"`
	Role      string  `json:"role" description:"Role for this API key (user or admin)"`
}

func CreateContext(oc openapi.OperationContext) error {
	oc.SetTags("API Keys")
	oc.SetSummary("Create API Key")
	oc.SetDescription("Creates a new API key with the specified parameters. Requires admin privileges.")

	oc.AddReqStructure(ApiKeyCreateRequest{})
	oc.AddRespStructure(models.ApiKey{}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "invalid request body"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusBadRequest),
	)
	oc.AddRespStructure(
		map[string]string{"error": "failed to save api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func CreateRequest(w http.ResponseWriter, r *http.Request, _ *redis.Client, s *solver.Solver) {
	var apiKey models.ApiKey
	if err := json.NewDecoder(r.Body).Decode(&apiKey); err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to decode request body: "+err.Error()))
		return
	}
	if err := database.DB.Create(&apiKey).Error; err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to save api key: "+err.Error()))
		return
	}
	if err := json.NewEncoder(w).Encode(apiKey); err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to encode response: "+err.Error()))
		return
	}
}

func Create() *routes.RouteHandler {
	return routes.NewRouteHandler(CreateRequest, CreateContext, nil, nil)
}
