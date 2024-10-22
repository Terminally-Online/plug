package router

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"solver/api"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(JsonContentTypeMiddleware)

	r.HandleFunc("/intent", api.GetIntent).Methods("POST")
	r.HandleFunc("/payment", api.GetPayment).Methods("GET")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Catch-all hit: %s %s", r.Method, r.URL.Path)
		http.Error(w, "Not found", http.StatusNotFound)
	})

	return r
}

func JsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

