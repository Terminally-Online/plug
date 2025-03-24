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

	// API key request counts
	apiKeyRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_key_requests_total",
			Help: "Total number of requests by API key",
		},
		[]string{"api_key_id", "endpoint"},
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

	// Cache metrics
	cacheHitsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cache_hits_total",
			Help: "Total number of cache hits",
		},
		[]string{"key"},
	)

	cacheMissesTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cache_misses_total",
			Help: "Total number of cache misses",
		},
		[]string{"key"},
	)

	cachePopulateTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "cache_populate_time_milliseconds",
			Help: "Time taken to populate cache in milliseconds",
			// Custom buckets appropriate for millisecond measurements
			Buckets: []float64{5, 10, 25, 50, 100, 250, 500, 1000, 2500, 5000, 10000},
		},
		[]string{"key"},
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
	apiKeyID      string
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

		// Get API key ID if it's already set (happens when middleware ordering puts this after ApiKey middleware)
		apiKeyID := r.Header.Get("X-Api-Key-Id")

		// Create a response writer with metrics
		metricsWriter := &ResponseWriterWithMetrics{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default status code
			requestMethod:  requestMethod,
			requestPath:    requestPath,
			startTime:      time.Now(),
			apiKeyID:       apiKeyID,
		}

		// Record request size
		httpRequestSize.WithLabelValues(requestMethod, requestPath).Observe(float64(requestSize))

		// Call the next handler
		next.ServeHTTP(metricsWriter, r)

		// After the handler executes, the API key ID header is set by the keys middleware
		if apiKeyID == "" {
			apiKeyID = r.Header.Get("X-Api-Key-Id")
			metricsWriter.apiKeyID = apiKeyID
		}

		// Track request count by API key and endpoint
		if apiKeyID != "" {
			apiKeyRequestsTotal.WithLabelValues(apiKeyID, requestPath).Inc()
		}

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

// TrackCacheOperation records cache hit/miss metrics
func TrackCacheOperation(key string, hit bool, populateTimeMillis int64) {
	if hit {
		cacheHitsTotal.WithLabelValues(key).Inc()
	} else {
		cacheMissesTotal.WithLabelValues(key).Inc()
	}

	// Only record populate time if the cache was populated (miss or refresh)
	if populateTimeMillis > 0 {
		// Convert to float64 as required by Observe method
		cachePopulateTime.WithLabelValues(key).Observe(float64(populateTimeMillis))
	}
}
