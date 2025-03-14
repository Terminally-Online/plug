package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/gorillamux"
)

// SetupOpenAPIRoutes sets up the OpenAPI specification routes
func SetupOpenAPIRoutes(r *mux.Router) http.Handler {
	// Setup OpenAPI schema
	refl := openapi3.NewReflector()
	
	// Configure the API metadata
	refl.SpecSchema().SetTitle("Plug Solver API")
	refl.SpecSchema().SetVersion("v1.0.0")
	refl.SpecSchema().SetDescription("API for Plug Solver to build Ethereum transactions.")
	
	// Create OpenAPI collector
	collector := gorillamux.NewOpenAPICollector(refl)

	// Walk the router and collect all the OpenAPI operations
	if err := r.Walk(collector.Walker); err != nil {
		panic(err)
	}

	// Create the handler to serve the OpenAPI specification
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		format := r.URL.Query().Get("format")
		if format == "yaml" {
			data, err := refl.Spec.MarshalYAML()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text/yaml")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Write(data)
			return
		}

		data, err := refl.Spec.MarshalJSON()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(data)
	})
}