package api

import (
	"solver/internal/api/middleware"
	"solver/internal/api/solver"

	"github.com/gorilla/mux"
)

func SetupRouter(s *solver.Handler) *mux.Router {
	m := middleware.New(s.Solver)

	r := mux.NewRouter()
	r.Use(m.Json)

	protected := r.PathPrefix("").Subrouter()
	protected.Use(m.ApiKey)

	protectedKillable := r.PathPrefix("").Subrouter()
	protectedKillable.Use(m.ApiKey)
	protectedKillable.Use(m.KillSwitch)

	r.HandleFunc("/solver", s.GetIntent).Methods("GET")
	protectedKillable.HandleFunc("/solver", s.PostIntent).Methods("POST")
	protected.HandleFunc("/solver/kill", s.GetKill).Methods("GET")
	protected.HandleFunc("/solver/kill", s.PostKill).Methods("POST")

	return r
}
