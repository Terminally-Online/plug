package open_api

import "net/http"

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
