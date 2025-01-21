package router

import (
	"net/http"
	"solver/intent"
	"solver/solver"

	"github.com/gorilla/mux"
)

func SetupRouter(solver *solver.Solver) *mux.Router {
	r := mux.NewRouter()
	r.Use(JsonContentTypeMiddleware)

	intentHandler := intent.NewHandler(solver)

	r.HandleFunc("/intent", intentHandler.Get).Methods("GET")
	r.HandleFunc("/intent", intentHandler.Post).Methods("POST")

	return r
}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
