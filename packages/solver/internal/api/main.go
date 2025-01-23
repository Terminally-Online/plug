package api

import (
	"net/http"
	"solver/internal/api/middleware"
	"solver/internal/api/solver"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(JsonContentTypeMiddleware)

	middleware := middleware.New()
	solver := solver.New()

	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.ApiKeyRequired)

	r.HandleFunc("/solver", solver.GetIntent).Methods("GET")
	r.HandleFunc("/solver", solver.PostIntent).Methods("POST")
	r.HandleFunc("/solver/kill", solver.GetKill).Methods("GET")
	r.HandleFunc("/solver/kill", solver.PostKill).Methods("GET")

	return r
}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
