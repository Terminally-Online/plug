# Monitoring Setup for Solver API

This directory contains configuration files for monitoring the solver API using Prometheus and Grafana.

## Components

- **Prometheus**: Collects and stores metrics from the solver API
- **Grafana**: Visualizes the metrics collected by Prometheus

## Metrics Collected

The solver API exposes the following metrics:

### HTTP Metrics

- `http_requests_total`: Total number of HTTP requests processed (labels: method, endpoint, status)
- `http_request_duration_seconds`: HTTP request duration in seconds (labels: method, endpoint)
- `http_request_size_bytes`: HTTP request size in bytes (labels: method, endpoint)
- `http_response_size_bytes`: HTTP response size in bytes (labels: method, endpoint)

### API Key Metrics

- `api_key_rate_limits`: Current rate limit values for API keys (labels: api_key_id)

### Solver Metrics

- `solver_requests_total`: Total number of solver requests processed
- `solver_request_duration_seconds`: Solver request duration in seconds
- `solver_errors_total`: Total number of solver errors

## Dashboard

The Grafana dashboard provides visualizations for:

1. HTTP request rate
2. Request duration (p95)
3. Solver requests and errors
4. Solver duration (p50 and p95)

## Setup & Usage

### Option 1: Start with monitoring enabled

```bash
pnpm dev:with-monitoring
```

This command will start both the monitoring stack and your solver API.

### Option 2: Add monitoring to a running API

If your API is already running:

```bash
pnpm monitoring:start
```

### Access the monitoring interfaces

- [Prometheus UI](http://localhost:9090)
- [Grafana Dashboard](http://localhost:3100) (default login: admin/admin)
- [Solver API Metrics](http://localhost:8080/metrics)

### Stop monitoring

To stop the monitoring services:

```bash
pnpm monitoring:stop
```

## Configuration Notes

- Prometheus is configured to connect to your locally running API using `host.docker.internal:8080`