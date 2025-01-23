package api

import (
	"encoding/json"
	"net/http"
	"os"
	"solver/internal/solver"
)

type StatusResponse struct {
	Status string `json:"status"`
}

type AdminHandler struct {
	solver *solver.Solver
}

func NewAdminHandler(solver *solver.Solver) *AdminHandler {
	return &AdminHandler{
		solver: solver,
	}
}

func (h *AdminHandler) authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey == "" {
			http.Error(w, "Missing API key", http.StatusUnauthorized)
			return
		}

		if apiKey != os.Getenv("PLUG_APP_API_KEY") {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (h *AdminHandler) HandleSolverStatus(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		status := "running"
		if h.solver.IsStopped() {
			status = "stopped"
		}
		response := StatusResponse{Status: status}
		json.NewEncoder(w).Encode(response)

	case http.MethodPost:
		var request struct {
			Action string `json:"action"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		switch request.Action {
		case "start":
			h.solver.Start()
		case "stop":
			h.solver.Stop()
		default:
			http.Error(w, "Invalid action", http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"status": request.Action + "ed"})
	}
}
