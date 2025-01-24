package middleware

import (
	"net/http"
	"os"
	"solver/internal/solver"
)

type Handler struct {
	apiKey string
	solver *solver.Solver
}

func New(s *solver.Solver) *Handler {
	apiKey := os.Getenv("PLUG_APP_API_KEY")
	if apiKey == "" {
		panic("PLUG_APP_API_KEY environment variable is not set")
	}

	return &Handler{
		apiKey: apiKey,
		solver: s,
	}
}

func (h *Handler) Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-api-key") != h.apiKey {
			http.Error(w, "invalid api key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) KillSwitch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow kill switch endpoints to pass through
		if r.URL.Path == "/solver/kill" {
			next.ServeHTTP(w, r)
			return
		}

		if h.solver.IsKilled {
			http.Error(
				w,
				"solver is temporarily killed and not processing orders",
				http.StatusUnauthorized,
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
