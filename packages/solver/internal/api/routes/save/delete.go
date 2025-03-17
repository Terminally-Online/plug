package save

import (
	"net/http"
	"solver/internal/api/routes"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/solver"
	"solver/internal/utils"

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

func DeleteRequest(w http.ResponseWriter, r *http.Request, _ *solver.Solver) {
	vars := mux.Vars(r)
	id := vars["id"]

	var intent models.Intent
	if err := database.DB.First(&intent, "id = ?", id).Error; err != nil {
		utils.MakeHttpError(w, "failed to find intent: "+err.Error(), http.StatusNotFound)
		return
	}

	if err := database.DB.Delete(&intent).Error; err != nil {
		utils.MakeHttpError(w, "failed to delete intent: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func Delete() *routes.RouteHandler {
	return routes.NewRouteHandler(DeleteRequest, DeleteContext, nil)
}
