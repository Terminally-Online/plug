package save

import (
	"fmt"
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
	oc.SetTags("Intents")
	oc.SetSummary("Delete Intent")
	oc.SetDescription("Deletes an intent by ID")

	oc.AddReqStructure(struct {
		ID string `path:"id" description:"Intent ID"`
	}{})

	oc.AddRespStructure(nil, openapi.WithHTTPStatus(http.StatusNoContent))
	oc.AddRespStructure(
		map[string]string{"error": "failed to find intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)
	oc.AddRespStructure(
		map[string]string{"error": "failed to delete intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func DeleteRequest(w http.ResponseWriter, r *http.Request, _ *redis.Client, s *solver.Solver) {
	vars := mux.Vars(r)
	id := vars["id"]

	var intent models.Intent
	if err := database.DB.First(&intent, "id = ?", id).Error; err != nil {
		utils.RespondWithError(w, utils.ErrNotFound(fmt.Sprintf("failed to find intent with id: %s", err.Error())))
		return
	}

	if err := database.DB.Delete(&intent).Error; err != nil {
		utils.RespondWithError(w, utils.ErrInternal("failed to delete intent: "+err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func Delete() *routes.RouteHandler {
	return routes.NewRouteHandler(DeleteRequest, DeleteContext, nil, nil)
}
