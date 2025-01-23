package api

import (
	"solver/internal/api/middleware"
	"solver/internal/api/solver"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	s := solver.New()
	m := middleware.New(s.Solver)

	r := mux.NewRouter()
	r.Use(m.Json)

	protected := r.PathPrefix("").Subrouter()
	protected.Use(m.ApiKey)
	protected.Use(m.KillSwitch)

	r.HandleFunc("/solver", s.GetIntent).Methods("GET")
	r.HandleFunc("/solver", s.PostIntent).Methods("POST")
	r.HandleFunc("/solver/kill", s.GetKill).Methods("GET")
	r.HandleFunc("/solver/kill", s.PostKill).Methods("POST")

	return r
}

