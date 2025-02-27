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

func (h *Handler) CreateIntent(w http.ResponseWriter, r *http.Request) {
	var inputs models.Intent
	if err := json.NewDecoder(r.Body).Decode(&inputs); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	var apiKey models.ApiKey
	if err := database.DB.First(&apiKey, "key = ?", r.Header.Get("X-Api-Key")).Error; err != nil {
		utils.MakeHttpError(w, "failed to find api key: "+err.Error(), http.StatusNotFound)
		return
	}
	inputs.ApiKeyId = apiKey.Id
	if err := database.DB.Omit("nextSimulationAt", "periodEndAt").Create(&inputs).Error; err != nil {
		utils.MakeHttpError(w, "failed to save intent: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(inputs); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ReadIntent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var intents []models.Intent
	query := database.DB
	if strings.HasPrefix(id, "0x") {
		query = query.Where("\"from\" = ?", id)
	} else {
		query = query.Where("id = ?", id)
	}
	result := query.
		Order("created_at desc").
		Find(&intents)
	if result.Error != nil {
		utils.MakeHttpError(w, "database error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(intents); err != nil {
		utils.MakeHttpError(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ToggleIntent(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) UpdateIntent(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "updating intents is not supported", http.StatusNotImplemented)
}

func (h *Handler) DeleteIntent(w http.ResponseWriter, r *http.Request) {
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
