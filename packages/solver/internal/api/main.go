package api

import (
	"net/http"
	"solver/internal/api/middleware"
	"solver/internal/api/solver"

	"github.com/gorilla/mux"
)

func SetupRouter(s solver.Handler) *mux.Router {
	m := middleware.New(s.Solver)
	r := mux.NewRouter()
	r.Use(m.Json)

	// Create OpenAPI handler for health endpoint
	r.Handle("/health", NewOpenAPIHandler(s.GetHealth, s.SetupOpenAPIForHealth)).Methods("GET")

	// Plug app only routes with OpenAPI documentation
	admin := r.PathPrefix("").Subrouter()
	admin.Use(m.AdminApiKey)

	admin.Handle("/api-key", NewOpenAPIHandler(s.CreateApiKey, s.SetupOpenAPIForCreateApiKey)).Methods("POST")
	admin.Handle("/api-key/{id}", NewOpenAPIHandler(s.ReadApiKey, s.SetupOpenAPIForReadApiKey)).Methods("GET")
	admin.Handle("/api-key/{id}", NewOpenAPIHandler(s.UpdateApiKey, s.SetupOpenAPIForUpdateApiKey)).Methods("POST")
	admin.Handle("/api-key/{id}", NewOpenAPIHandler(s.DeleteApiKey, s.SetupOpenAPIForDeleteApiKey)).Methods("DELETE")

	// Kill switch endpoints with OpenAPI
	admin.Handle("/solver/kill", NewOpenAPIHandler(s.GetKill, s.SetupOpenAPIForGetKill)).Methods("GET")
	admin.Handle("/solver/kill", NewOpenAPIHandler(s.PostKill, s.SetupOpenAPIForPostKill)).Methods("POST")

	// API key protected routes
	protected := r.PathPrefix("").Subrouter()
	protected.Use(m.ApiKey)

	// Intent endpoints with OpenAPI
	protected.Handle("/solver/save", NewOpenAPIHandler(s.CreateIntent, s.SetupOpenAPIForCreateIntent)).Methods("POST")
	protected.Handle("/solver/save", NewOpenAPIHandler(s.ReadIntents, s.SetupOpenAPIForReadIntents)).Methods("GET")
	protected.Handle("/solver/save/{id}", NewOpenAPIHandler(s.ReadIntent, s.SetupOpenAPIForReadIntent)).Methods("GET")
	protected.Handle("/solver/save/{id}", NewOpenAPIHandler(s.ToggleIntentSaved, s.SetupOpenAPIForToggleIntentSaved)).Methods("POST")
	protected.Handle("/solver/save/{id}", NewOpenAPIHandler(s.DeleteIntent, s.SetupOpenAPIForDeleteIntent)).Methods("DELETE")
	protected.Handle("/solver/save/{id}/status", NewOpenAPIHandler(s.ToggleIntentStatus, s.SetupOpenAPIForToggleIntentStatus)).Methods("POST")

	// API key protected routes that can be killed
	killable := protected.PathPrefix("").Subrouter()
	killable.Use(m.KillSwitch)

	// Main solver endpoints with OpenAPI
	killable.Handle("/solver", NewOpenAPIHandler(s.GetSchema, s.SetupOpenAPIForGetSchema)).Methods("GET")
	killable.Handle("/solver", NewOpenAPIHandler(s.GetSolution, s.SetupOpenAPIForGetSolution)).Methods("POST")

	// Setup OpenAPI routes
	specHandler := SetupOpenAPIRoutes(r)

	// Add OpenAPI specification endpoints (both JSON and YAML formats)
	r.Handle("/openapi.json", specHandler).Methods("GET")
	r.Handle("/openapi", specHandler).Methods("GET").Queries("format", "{format}")

	// Add Swagger UI viewer endpoint
	r.HandleFunc("/api-docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		html := `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Plug Solver API Documentation</title>
  <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css">
  <style>
    html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
    *, *:before, *:after { box-sizing: inherit; }
    body { margin: 0; background: #fafafa; }
  </style>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js"></script>
  <script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-standalone-preset.js"></script>
  <script>
    window.onload = function() {
      window.ui = SwaggerUIBundle({
        url: "/openapi.json",
        dom_id: "#swagger-ui",
        deepLinking: true,
        presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
        plugins: [SwaggerUIBundle.plugins.DownloadUrl],
        layout: "StandaloneLayout",
        requestInterceptor: function(req) {
          var apiKey = localStorage.getItem("api_key");
          if (apiKey) {
            req.headers["X-Api-Key"] = apiKey;
          }
          return req;
        },
        onComplete: function() {
          var apiKey = localStorage.getItem("api_key");
          if (apiKey) {
            window.ui.preauthorizeApiKey("ApiKeyAuth", apiKey);
          }
        }
      });

      // Add API key input field
      var topbarEl = document.querySelector(".topbar");
      var apiKeyContainer = document.createElement("div");
      apiKeyContainer.style.cssText = "margin-right: 16px; display: flex; align-items: center;";
      apiKeyContainer.innerHTML = '<label style="color: #3b4151; margin-right: 8px;">API Key:</label><input type="text" id="api_key_input" style="padding: 4px 8px; border: 1px solid #d9d9d9; border-radius: 4px;" placeholder="Enter API key">';
      topbarEl.insertBefore(apiKeyContainer, topbarEl.firstChild);

      // Handle API key input
      var apiKeyInput = document.getElementById("api_key_input");
      apiKeyInput.value = localStorage.getItem("api_key") || "";
      apiKeyInput.addEventListener("change", function(e) {
        var key = e.target.value.trim();
        if (key) {
          localStorage.setItem("api_key", key);
          window.ui.preauthorizeApiKey("ApiKeyAuth", key);
        } else {
          localStorage.removeItem("api_key");
        }
      });
    };
  </script>
</body>
</html>`
		w.Write([]byte(html))
	}).Methods("GET")

	return r
}
