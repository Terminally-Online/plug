package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Total number of HTTP requests processed
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests processed",
		},
		[]string{"method", "endpoint", "status"},
	)

	// Request duration
	httpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Request size
	httpRequestSize = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_size_bytes",
			Help: "HTTP request size in bytes",
		},
		[]string{"method", "endpoint"},
	)

	// Response size
	httpResponseSize = promauto.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_response_size_bytes",
			Help: "HTTP response size in bytes",
		},
		[]string{"method", "endpoint"},
	)

	// API key rate limits
	apiKeyRateLimits = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "api_key_rate_limits",
			Help: "Current rate limit values for API keys",
		},
		[]string{"api_key_id"},
	)

	// Solver metrics
	solverRequestsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "solver_requests_total",
			Help: "Total number of solver requests processed",
		},
	)

	solverRequestDuration = promauto.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "solver_request_duration_seconds",
			Help:    "Solver request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
	)

	solverErrorsTotal = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "solver_errors_total",
			Help: "Total number of solver errors",
		},
	)
)

// ResponseWriterWithMetrics is a wrapper for http.ResponseWriter that captures metrics
type ResponseWriterWithMetrics struct {
	http.ResponseWriter
	statusCode    int
	responseSize  int64
	requestMethod string
	requestPath   string
	startTime     time.Time
}

// WriteHeader captures the status code from the response
func (rwm *ResponseWriterWithMetrics) WriteHeader(statusCode int) {
	rwm.statusCode = statusCode
	rwm.ResponseWriter.WriteHeader(statusCode)
}

// Write captures the size of the response
func (rwm *ResponseWriterWithMetrics) Write(b []byte) (int, error) {
	size, err := rwm.ResponseWriter.Write(b)
	rwm.responseSize += int64(size)
	return size, err
}

// recordMetrics records metrics after the response is complete
func (rwm *ResponseWriterWithMetrics) recordMetrics() {
	duration := time.Since(rwm.startTime).Seconds()
	statusCode := strconv.Itoa(rwm.statusCode)

	httpRequestsTotal.WithLabelValues(rwm.requestMethod, rwm.requestPath, statusCode).Inc()
	httpRequestDuration.WithLabelValues(rwm.requestMethod, rwm.requestPath).Observe(duration)
	httpResponseSize.WithLabelValues(rwm.requestMethod, rwm.requestPath).Observe(float64(rwm.responseSize))
}

// Metrics is a middleware for collecting HTTP metrics
func (m *Middleware) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPath := r.URL.Path
		requestMethod := r.Method
		requestSize := r.ContentLength

		// Create a response writer with metrics
		metricsWriter := &ResponseWriterWithMetrics{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default status code
			requestMethod:  requestMethod,
			requestPath:    requestPath,
			startTime:      time.Now(),
		}

		// Record request size
		httpRequestSize.WithLabelValues(requestMethod, requestPath).Observe(float64(requestSize))

		// Track API key usage if present
		if apiKeyID := r.Header.Get("X-Api-Key-Id"); apiKeyID != "" {
			if limit, ok := m.apiKeyLimiter.GetLimit(apiKeyID); ok {
				apiKeyRateLimits.WithLabelValues(apiKeyID).Set(float64(limit.Remaining))
			}
		}

		// Call the next handler
		next.ServeHTTP(metricsWriter, r)

		// Record response metrics
		metricsWriter.recordMetrics()
	})
}

// TrackSolverMetrics records solver-specific metrics
func TrackSolverMetrics(fn func() error) error {
	solverRequestsTotal.Inc()
	
	startTime := time.Now()
	err := fn()
	
	solverRequestDuration.Observe(time.Since(startTime).Seconds())
	
	if err != nil {
		solverErrorsTotal.Inc()
	}
	
	return err
}