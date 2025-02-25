package middleware

import (
	"net/http"
	"solver/internal/solver"
)

type Handler struct {
	solver      solver.Solver
	rateLimiter *RateLimiter
}

func New(s solver.Solver) *Handler {
	rateLimiter := NewRateLimiter()

	return &Handler{
		solver:      s,
		rateLimiter: rateLimiter,
	}
}

func (h *Handler) Json(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) KillSwitch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
