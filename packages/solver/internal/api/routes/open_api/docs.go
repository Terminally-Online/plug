package open_api

import (
	"net/http"

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

func DocsRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
			<!DOCTYPE html>
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
						dom_id: '#swagger-ui',
						deepLinking: true,
						presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
						plugins: [SwaggerUIBundle.plugins.DownloadUrl],
						layout: "StandaloneLayout"
					});
					};
				</script>
				</body>
			</html>
		`
	w.Write([]byte(html))
}
