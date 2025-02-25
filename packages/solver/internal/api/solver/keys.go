package solver

import (
	"encoding/json"
	"net/http"
	"solver/internal/database"
	"solver/internal/database/models"
	"solver/internal/utils"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateApiKey(w http.ResponseWriter, r *http.Request) {
	var apiKey models.ApiKey
	if err := json.NewDecoder(r.Body).Decode(&apiKey); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.DB.Create(&apiKey).Error; err != nil {
		utils.MakeHttpError(w, "failed to save api key: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(apiKey); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ReadApiKey(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) UpdateApiKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var apiKey models.ApiKey
	if err := database.DB.First(&apiKey, "id = ?", id).Error; err != nil {
		utils.MakeHttpError(w, "failed to find api key: "+err.Error(), http.StatusNotFound)
		return
	}

	var newApiKey models.ApiKey
	if err := json.NewDecoder(r.Body).Decode(&newApiKey); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.DB.Model(&apiKey).Updates(newApiKey).Error; err != nil {
		utils.MakeHttpError(w, "failed to update api key: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteApiKey(w http.ResponseWriter, r *http.Request) {
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
