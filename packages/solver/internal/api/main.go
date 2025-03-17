package api

import (
	"net/http"
	"solver/internal/api/middleware"
	"solver/internal/api/routes/api_key"
	"solver/internal/api/routes/health"
	"solver/internal/api/routes/intent"
	"solver/internal/api/routes/kill"
	"solver/internal/api/routes/open_api"
	"solver/internal/api/routes/save"
	"solver/internal/solver"

	"github.com/gorilla/mux"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/gorillamux"
)

func SetupOpenAPIRoutes(r *mux.Router) http.Handler {
	refl := openapi3.NewReflector()

	refl.SpecSchema().SetTitle("Plug Solver API")
	refl.SpecSchema().SetVersion("v1.0.0")
	refl.SpecSchema().SetDescription("API for Plug Solver to build Ethereum transactions.")

	collector := gorillamux.NewOpenAPICollector(refl)

	if err := r.Walk(collector.Walker); err != nil {
		panic(err)
	}

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

func SetupRouter(s *solver.Solver) *mux.Router {
	m := middleware.New(*s)
	r := mux.NewRouter()

	r.Use(m.Json)
	r.Use(m.Timeout(middleware.DefaultTimeout))

	r.Handle("/health", health.GetHealth()).Methods(http.MethodGet)

	admin := r.PathPrefix("").Subrouter()
	admin.Use(m.AdminApiKey)

	admin.Handle("/api-key", api_key.Create()).Methods(http.MethodPost)
	admin.Handle("/api-key/{id}", api_key.Read()).Methods(http.MethodGet)
	admin.Handle("/api-key/{id}", api_key.Update()).Methods(http.MethodPost)
	admin.Handle("/api-key/{id}", api_key.Delete()).Methods(http.MethodDelete)

	admin.Handle("/solver/kill", kill.Get(s)).Methods(http.MethodGet)
	admin.Handle("/solver/kill", kill.Post(s)).Methods(http.MethodPost)

	protected := r.PathPrefix("").Subrouter()
	protected.Use(m.ApiKey)

	protected.Handle("/solver/save", save.Create()).Methods(http.MethodPost)
	protected.Handle("/solver/save", NewOpenAPIHandler(s.ReadIntents, s.SetupOpenAPIForReadIntents)).Methods("GET")
	protected.Handle("/solver/save/{id}", save.Read()).Methods(http.MethodGet)
	protected.Handle("/solver/save/{id}", save.ToggleSaved()).Methods(http.MethodPost)
	protected.Handle("/solver/save/{id}", save.Delete()).Methods(http.MethodDelete)
	protected.Handle("/solver/save/{id}/status", save.ToggleStatus()).Methods(http.MethodPost)

	killable := protected.PathPrefix("").Subrouter()
	killable.Use(m.KillSwitch)

	killable.Handle("/solver", intent.Get(s)).Methods(http.MethodGet)
	killable.Handle("/solver", intent.Post(s)).Methods(http.MethodPost)

	specHandler := SetupOpenAPIRoutes(r)
	r.Handle("/openapi.json", specHandler).Methods(http.MethodGet)
	r.Handle("/openapi", specHandler).Methods(http.MethodGet).Queries("format", "{format}")
	r.HandleFunc("/docs", open_api.DocsRequest).Methods(http.MethodGet)

	return r
}
