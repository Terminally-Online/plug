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
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter(s *solver.Solver) *mux.Router {
	m := middleware.New(*s)
	r := mux.NewRouter()

	r.Use(m.Metrics)
	r.Use(m.Json)
	r.Use(m.Timeout)

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
	// protected.Handle("/solver/save", NewOpenAPIHandler(s.ReadIntents, s.SetupOpenAPIForReadIntents)).Methods("GET")
	protected.Handle("/solver/save/{id}", save.Read()).Methods(http.MethodGet)
	protected.Handle("/solver/save/{id}", save.ToggleSaved()).Methods(http.MethodPost)
	protected.Handle("/solver/save/{id}", save.Delete()).Methods(http.MethodDelete)
	protected.Handle("/solver/save/{id}/status", save.ToggleStatus()).Methods(http.MethodPost)

	killable := protected.PathPrefix("").Subrouter()
	killable.Use(m.KillSwitch)

	killable.Handle("/solver", intent.Get(s)).Methods(http.MethodGet)
	killable.Handle("/solver", intent.Post(s)).Methods(http.MethodPost)

	openapi := open_api.SetupOpenAPIRoutes(r)
	r.Handle("/openapi.json", openapi).Methods(http.MethodGet)
	r.Handle("/openapi", openapi).Methods(http.MethodGet).Queries("format", "{format}")
	r.HandleFunc("/docs", open_api.DocsRequest).Methods(http.MethodGet)

	r.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	return r
}
