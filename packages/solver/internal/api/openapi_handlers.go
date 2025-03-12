package api

import (
	"net/http"

	"github.com/swaggest/openapi-go"
)

// OpenAPIHandler defines an interface for handlers that support OpenAPI documentation
type OpenAPIHandler interface {
	http.Handler
	SetupOpenAPIOperation(oc openapi.OperationContext) error
}

// GenericOpenAPIHandler wraps a standard HTTP handler with OpenAPI documentation
type GenericOpenAPIHandler struct {
	handler     http.HandlerFunc
	setupOpenAPI func(oc openapi.OperationContext) error
}

// NewOpenAPIHandler creates a new OpenAPI-compatible handler
func NewOpenAPIHandler(handler http.HandlerFunc, setupOpenAPI func(oc openapi.OperationContext) error) *GenericOpenAPIHandler {
	return &GenericOpenAPIHandler{
		handler:     handler,
		setupOpenAPI: setupOpenAPI,
	}
}

// ServeHTTP implements the http.Handler interface
func (h *GenericOpenAPIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler(w, r)
}

// SetupOpenAPIOperation implements the OpenAPIHandler interface
func (h *GenericOpenAPIHandler) SetupOpenAPIOperation(oc openapi.OperationContext) error {
	return h.setupOpenAPI(oc)
}