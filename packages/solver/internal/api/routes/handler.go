package routes

import (
	"net/http"
	"solver/internal/solver"

	"github.com/go-redis/redis/v8"
	"github.com/swaggest/openapi-go"
)

type RouteFunc func(http.ResponseWriter, *http.Request, *redis.Client, *solver.Solver)
type OpenAPIFunc func(oc openapi.OperationContext) error

type RouteHandler struct {
	route        RouteFunc
	setupOpenAPI OpenAPIFunc
	redis        *redis.Client
	solver       *solver.Solver
}

func NewRouteHandler(route RouteFunc, setupOpenAPI OpenAPIFunc, redis *redis.Client, solver *solver.Solver) *RouteHandler {
	return &RouteHandler{
		route:        route,
		setupOpenAPI: setupOpenAPI,
		redis:        redis,
		solver:       solver,
	}
}

func (h *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.route(w, r, h.redis, h.solver)
}

func (h *RouteHandler) SetupOpenAPIOperation(oc openapi.OperationContext) error {
	return h.setupOpenAPI(oc)
}
