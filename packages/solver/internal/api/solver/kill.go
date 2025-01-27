package solver

import (
	"encoding/json"
	"net/http"
)

type KillResponse struct {
	Killed bool `json:"killed"`
}

func (h *Handler) GetKill(w http.ResponseWriter, r *http.Request) {
	response := KillResponse{Killed: h.Solver.IsKilled}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) PostKill(w http.ResponseWriter, r *http.Request) {
	h.Solver.IsKilled = !h.Solver.IsKilled

	response := KillResponse{Killed: h.Solver.IsKilled}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
