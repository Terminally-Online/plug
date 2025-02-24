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

	protected := r.PathPrefix("").Subrouter()
	protected.Use(m.ApiKey)
	protected.HandleFunc("/solver/kill", s.GetKill).Methods("GET")
	protected.HandleFunc("/solver/kill", s.PostKill).Methods("POST")

	protected.HandleFunc("/solver/save", s.CreateIntent).Methods("POST")
	protected.HandleFunc("/solver/save/{id}", s.ReadIntent).Methods("GET")
	protected.HandleFunc("/solver/save/{id}", s.UpdateIntent).Methods("POST")
	protected.HandleFunc("/solver/save/{id}", s.DeleteIntent).Methods("DELETE")
	protected.HandleFunc("/solver/save/{id}/toggle", s.ToggleIntent).Methods("POST")

	killable := protected.PathPrefix("").Subrouter()
	killable.Use(m.KillSwitch)
	killable.HandleFunc("/solver", s.GetSchema).Methods("GET")
	killable.HandleFunc("/solver", s.GetSolution).Methods("POST")

	return r
}
