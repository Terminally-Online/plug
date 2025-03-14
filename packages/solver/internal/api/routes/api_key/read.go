package api_key

import (
	"encoding/json"
	"net/http"
	"solver/internal/api/routes"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"solver/internal/utils"

	"github.com/gorilla/mux"
	"github.com/swaggest/openapi-go"
)

type ApiKeyResponse struct {
	ID        string  `json:"id" description:"Unique identifier for the API key"`
	Key       string  `json:"key" description:"The API key value"`
	SocketID  *string `json:"socketId,omitempty" description:"ID of the associated socket"`
	RateLimit int     `json:"rateLimit" description:"Rate limit for the API key"`
	Role      string  `json:"role" description:"Role of the API key (user or admin)"`
}

func ReadContext(oc openapi.OperationContext) error {
	oc.SetTags("API Keys")
	oc.SetSummary("Get API Key")
	oc.SetDescription("Retrieves an API key by ID. Requires admin privileges.")

	oc.AddReqStructure(struct {
		ID string `path:"id" description:"API key ID"`
	}{})

	oc.AddRespStructure(models.ApiKey{}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "failed to find api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)

	return nil
}

func ReadRequest(w http.ResponseWriter, r *http.Request, _ *solver.Solver) {
	vars := mux.Vars(r)
	id := vars["id"]

	var apiKey models.ApiKey
	if err := database.DB.First(&apiKey, "id = ?", id).Error; err != nil {
		utils.MakeHttpError(w, "failed to find api key: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(apiKey); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Read() *routes.RouteHandler {
	return routes.NewRouteHandler(ReadRequest, ReadContext, nil)
}
