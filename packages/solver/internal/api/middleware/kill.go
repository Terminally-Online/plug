package middleware

import "net/http"

func (m *Middleware) KillSwitch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m.solver.IsKilled {
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
