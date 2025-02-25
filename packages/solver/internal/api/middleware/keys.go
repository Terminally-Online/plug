package middleware

import (
	"net/http"
	"solver/internal/database"
	"solver/internal/database/models"
)

func (h *Handler) AdminApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		if apiKey == "" {
			http.Error(w, "missing api key", http.StatusUnauthorized)
			return
		}

		var dbApiKey models.ApiKey
		if err := database.DB.Where("key = ?", apiKey).First(&dbApiKey).Error; err != nil {
			http.Error(w, "invalid api key", http.StatusUnauthorized)
			return
		}
		if dbApiKey.Role != "admin" {
			http.Error(w, "invalid api key role", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// TODO: Implement rate limiting middleware
func (h *Handler) ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-Api-Key")
		if apiKey == "" {
			http.Error(w, "missing api key", http.StatusUnauthorized)
			return
		}

		var dbApiKey models.ApiKey
		if err := database.DB.Where("key = ?", apiKey).First(&dbApiKey).Error; err != nil {
			http.Error(w, "invalid api key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
