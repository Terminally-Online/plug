package routes

import (
	"net/http"
	"solver/internal/solver"

	"github.com/swaggest/openapi-go"
)

type RouteFunc func(http.ResponseWriter, *http.Request, *solver.Solver)
type OpenAPIFunc func(oc openapi.OperationContext) error

type RouteHandler struct {
	route        RouteFunc
	setupOpenAPI OpenAPIFunc
	solver       *solver.Solver
}

func NewRouteHandler(route RouteFunc, setupOpenAPI OpenAPIFunc, solver *solver.Solver) *RouteHandler {
	return &RouteHandler{
		route:        route,
		setupOpenAPI: setupOpenAPI,
		solver:       solver,
	}
}

func (h *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.route(w, r, h.solver)
}

func (h *RouteHandler) SetupOpenAPIOperation(oc openapi.OperationContext) error {
	return h.setupOpenAPI(oc)
}
