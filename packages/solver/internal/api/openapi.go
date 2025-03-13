package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/gorillamux"
)

// SetupOpenAPIRoutes sets up the OpenAPI specification routes
func SetupOpenAPIRoutes(r *mux.Router) http.Handler {
	// Setup OpenAPI schema
	refl := openapi3.NewReflector()

	// Configure the API metadata
	schema := refl.SpecSchema()
	schema.SetTitle("Plug Solver API")
	schema.SetVersion("v1.0.0")
	schema.SetDescription("API for Plug Solver to build Ethereum transactions.")

	// Initialize Components
	refl.Spec.Components = &openapi3.Components{}

	// Add API key security scheme
	refl.Spec.Components.SecuritySchemes = &openapi3.ComponentsSecuritySchemes{
		MapOfSecuritySchemeOrRefValues: map[string]openapi3.SecuritySchemeOrRef{
			"ApiKeyAuth": {
				SecurityScheme: &openapi3.SecurityScheme{
					APIKeySecurityScheme: &openapi3.APIKeySecurityScheme{
						Name: "X-Api-Key",
						In:   "header",
					},
				},
			},
		},
	}

	// Set global security requirement
	refl.Spec.Security = []map[string][]string{
		{
			"ApiKeyAuth": {},
		},
	}

	// Create OpenAPI collector
	collector := gorillamux.NewOpenAPICollector(refl)

	// Walk the router and collect all the OpenAPI operations
	if err := r.Walk(collector.Walker); err != nil {
		panic(err)
	}

	// Remove admin endpoints from the spec
	paths := refl.Spec.Paths.MapOfPathItemValues
	if paths != nil {
		for path := range paths {
			if path == "/api-key" || strings.HasPrefix(path, "/api-key/") || path == "/solver/kill" {
				delete(paths, path)
			}
		}
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
