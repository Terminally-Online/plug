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

func DeleteContext(oc openapi.OperationContext) error {
	oc.SetTags("API Keys")
	oc.SetSummary("Delete API Key")
	oc.SetDescription("Deletes an API key by ID. Requires admin privileges.")

	oc.AddReqStructure(struct {
		ID string `path:"id" description:"API key ID"`
	}{})

	oc.AddRespStructure(models.ApiKey{}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "failed to find api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)
	oc.AddRespStructure(
		map[string]string{"error": "failed to delete api key"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func DeleteRequest(w http.ResponseWriter, r *http.Request, _ *redis.Client, s *solver.Solver) {
	vars := mux.Vars(r)
	id := vars["id"]

	var apiKey models.ApiKey
	if err := database.DB.First(&apiKey, "id = ?", id).Error; err != nil {
		utils.MakeHttpError(w, "failed to find api key: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&apiKey).Error; err != nil {
		utils.MakeHttpError(w, "failed to delete api key: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(apiKey); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func Delete() *routes.RouteHandler {
	return routes.NewRouteHandler(DeleteRequest, DeleteContext, nil, nil)
}
