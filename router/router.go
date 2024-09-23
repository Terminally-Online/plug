package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"solver/api"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.Use(jsonContentTypeMiddleware)

	intent := r.PathPrefix("/intent").Subrouter()
	intent.HandleFunc("", api.CreateIntent).Methods("POST")

	return r
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
