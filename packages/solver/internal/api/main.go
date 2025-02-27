package api

import (
	"solver/internal/api/middleware"
	"solver/internal/api/solver"

	"github.com/gorilla/mux"
)

func SetupRouter(s solver.Handler) *mux.Router {
	m := middleware.New(s.Solver)
	r := mux.NewRouter()
	r.Use(m.Json)

	// Plug app only routes
	admin := r.PathPrefix("").Subrouter()
	admin.Use(m.AdminApiKey)

	admin.HandleFunc("/api-key", s.CreateApiKey).Methods("POST")
	admin.HandleFunc("/api-key/{id}", s.ReadApiKey).Methods("GET")
	admin.HandleFunc("/api-key/{id}", s.UpdateApiKey).Methods("POST")
	admin.HandleFunc("/api-key/{id}", s.DeleteApiKey).Methods("DELETE")

	admin.HandleFunc("/solver/kill", s.GetKill).Methods("GET")
	admin.HandleFunc("/solver/kill", s.PostKill).Methods("POST")

	// API key protected routes
	protected := r.PathPrefix("").Subrouter()
	protected.Use(m.ApiKey)

	protected.HandleFunc("/solver/save", s.CreateIntent).Methods("POST")
	protected.HandleFunc("/solver/save/{id}", s.ReadIntent).Methods("GET")
	protected.HandleFunc("/solver/save/{id}", s.ToggleIntent).Methods("POST")
	protected.HandleFunc("/solver/save/{id}", s.DeleteIntent).Methods("DELETE")

	// API key protected routes that can be killed
	killable := protected.PathPrefix("").Subrouter()
	killable.Use(m.KillSwitch)
	killable.HandleFunc("/solver", s.GetSchema).Methods("GET")
	killable.HandleFunc("/solver", s.GetSolution).Methods("POST")

	return r
}
