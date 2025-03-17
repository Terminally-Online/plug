package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type TimeoutResponseWriter struct {
	http.ResponseWriter
	wroteHeader *bool
}

func (tw *TimeoutResponseWriter) WriteHeader(code int) {
	*tw.wroteHeader = true
	tw.ResponseWriter.WriteHeader(code)
}

func (tw *TimeoutResponseWriter) Write(b []byte) (int, error) {
	*tw.wroteHeader = true
	return tw.ResponseWriter.Write(b)
}

func (m *Middleware) Timeout(timeout time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			done := make(chan struct{})

			var wroteHeader bool

			wrappedWriter := &TimeoutResponseWriter{
				ResponseWriter: w,
				wroteHeader:    &wroteHeader,
			}

			go func() {
				next.ServeHTTP(wrappedWriter, r.WithContext(ctx))
				close(done)
			}()

			select {
			case <-done:
				return
			case <-ctx.Done():
				if ctx.Err() == context.DeadlineExceeded && !wroteHeader {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusRequestTimeout)
					json.NewEncoder(w).Encode(map[string]string{
						"error": "Request timed out",
					})
				}
				return
			}
		})
	}
}
