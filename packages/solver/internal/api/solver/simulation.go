package solver

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"solver/internal/solver/simulation"
)

func (h *Handler) PostSimulation(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var req simulation.SimulationRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request: %v", err), http.StatusBadRequest)
		return
	}

	resp, err := h.Simulator.Simulate(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Simulation failed: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
		return
	}
}
