package solver

import (
	"encoding/json"
	"net/http"
	"strings"

	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/utils"

	"github.com/gorilla/mux"
)

func (h *Handler) GetIntent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if strings.HasPrefix(id, "0x") {
	}
}

// TODO: Make upsert inset.
func (h *Handler) CreateIntent(w http.ResponseWriter, r *http.Request) {
	var inputs models.Intent
	if err := json.NewDecoder(r.Body).Decode(&inputs); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Omit("id", "nextSimulationAt", "periodEndAt").Create(&inputs).Error; err != nil {
		utils.MakeHttpError(w, "failed to save intent: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(inputs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}


func (h *Handler) ReadIntent(w http.ResponseWriter, r *http.Request) {
	var inputs models.Intent
	if err := json.NewDecoder(r.Body).Decode(&inputs); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Omit("id", "nextSimulationAt", "periodEndAt").Create(&inputs).Error; err != nil {
		utils.MakeHttpError(w, "failed to save intent: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(inputs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}