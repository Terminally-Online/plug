package middleware

import (
	"net/http"
	"os"
)

type MiddlewareHandler struct {
	ApiKey string
}

func New() *MiddlewareHandler {
	apiKey := os.Getenv("PLUG_APP_API_KEY")
	if apiKey == "" {
		panic("PLUG_APP_API_KEY environment variable is not set")
	}

	return &MiddlewareHandler{
		ApiKey: apiKey,
	}
}

func (h *MiddlewareHandler) ApiKeyRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-api-key") != h.ApiKey {
			http.Error(w, "Invalid API key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
