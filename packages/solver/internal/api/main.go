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

	// requires the .env api key
	admin := r.PathPrefix("").Subrouter()
	admin.Use(m.AdminApiKey)
	admin.HandleFunc("/solver/kill", s.PostKill).Methods("POST")

	// requires a valid api key
	protected := r.PathPrefix("").Subrouter()
	protected.Use(m.ApiKey)
	protected.HandleFunc("/solver/kill", s.GetKill).Methods("GET")

	// can be short circuited by the kill switch
	killable := protected.PathPrefix("").Subrouter()
	killable.Use(m.KillSwitch)
	killable.HandleFunc("/solver", s.GetSchema).Methods("GET")
	killable.HandleFunc("/solver", s.GetSolution).Methods("POST")

	return r
}
