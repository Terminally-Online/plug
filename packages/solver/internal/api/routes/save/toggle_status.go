package save

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

func ToggleStatusContext(oc openapi.OperationContext) error {
	oc.SetTags("Intents")
	oc.SetSummary("Toggle Intent Status")
	oc.SetDescription("Toggles the status of an intent between 'active' and 'paused'")

	oc.AddReqStructure(struct {
		ID string `path:"id" description:"Intent ID"`
	}{})

	oc.AddRespStructure(models.Intent{}, openapi.WithHTTPStatus(http.StatusOK))
	oc.AddRespStructure(
		map[string]string{"error": "failed to find intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusNotFound),
	)
	oc.AddRespStructure(
		map[string]string{"error": "failed to toggle intent"},
		openapi.WithContentType("application/json"),
		openapi.WithHTTPStatus(http.StatusInternalServerError),
	)

	return nil
}

func ToggleStatusRequest(w http.ResponseWriter, r *http.Request, _ *solver.Solver) {
	vars := mux.Vars(r)
	id := vars["id"]

	var intent models.Intent
	if err := database.DB.First(&intent, "id = ?", id).Error; err != nil {
		utils.MakeHttpError(w, "failed to find intent: "+err.Error(), http.StatusNotFound)
		return
	}

	if intent.Status != "active" {
		intent.Status = "active"
	} else {
		intent.Status = "paused"
	}

	if err := database.DB.Select("status").Updates(&intent).Error; err != nil {
		utils.MakeHttpError(w, "failed to toggle intent: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(intent); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func ToggleStatus() *routes.RouteHandler {
	return routes.NewRouteHandler(ToggleStatusRequest, ToggleStatusContext, nil)
}
