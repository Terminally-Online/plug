# Monitoring Setup for Solver API

This directory contains configuration files for monitoring the solver API using Prometheus and Grafana.

## Components

- **Prometheus**: Collects and stores metrics from the solver API
- **Grafana**: Visualizes the metrics collected by Prometheus

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

- Prometheus is configured to connect to the locally running API using `host.docker.internal:8080`