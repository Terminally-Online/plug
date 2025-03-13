package solver

import (
	"solver/internal/solver"

	"github.com/swaggest/openapi-go"
)

type Handler struct {
	Solver solver.Solver
}

func New() Handler {
	return Handler{
		Solver: solver.New(),
	}
}

// OpenAPI setup methods for health
func (h *Handler) SetupOpenAPIForHealth(oc openapi.OperationContext) error {
	return SetupOpenAPIForHealth(oc)
}

// OpenAPI setup methods for kill
func (h *Handler) SetupOpenAPIForGetKill(oc openapi.OperationContext) error {
	return SetupOpenAPIForGetKill(oc)
}

func (h *Handler) SetupOpenAPIForPostKill(oc openapi.OperationContext) error {
	return SetupOpenAPIForPostKill(oc)
}

// OpenAPI setup methods for intents
func (h *Handler) SetupOpenAPIForCreateIntent(oc openapi.OperationContext) error {
	return SetupOpenAPIForCreateIntent(oc)
}

func (h *Handler) SetupOpenAPIForReadIntents(oc openapi.OperationContext) error {
	return SetupOpenAPIForReadIntents(oc)
}

func (h *Handler) SetupOpenAPIForReadIntent(oc openapi.OperationContext) error {
	return SetupOpenAPIForReadIntent(oc)
}

func (h *Handler) SetupOpenAPIForToggleIntentStatus(oc openapi.OperationContext) error {
	return SetupOpenAPIForToggleIntentStatus(oc)
}

func (h *Handler) SetupOpenAPIForToggleIntentSaved(oc openapi.OperationContext) error {
	return SetupOpenAPIForToggleIntentSaved(oc)
}

func (h *Handler) SetupOpenAPIForDeleteIntent(oc openapi.OperationContext) error {
	return SetupOpenAPIForDeleteIntent(oc)
}

// OpenAPI setup methods for API keys
func (h *Handler) SetupOpenAPIForCreateApiKey(oc openapi.OperationContext) error {
	return SetupOpenAPIForCreateApiKey(oc)
}

func (h *Handler) SetupOpenAPIForReadApiKey(oc openapi.OperationContext) error {
	return SetupOpenAPIForReadApiKey(oc)
}

func (h *Handler) SetupOpenAPIForUpdateApiKey(oc openapi.OperationContext) error {
	return SetupOpenAPIForUpdateApiKey(oc)
}

func (h *Handler) SetupOpenAPIForDeleteApiKey(oc openapi.OperationContext) error {
	return SetupOpenAPIForDeleteApiKey(oc)
}

// OpenAPI setup methods for solver
func (h *Handler) SetupOpenAPIForGetSchema(oc openapi.OperationContext) error {
	return SetupOpenAPIForGetSchema(oc)
}

func (h *Handler) SetupOpenAPIForGetSolution(oc openapi.OperationContext) error {
	return SetupOpenAPIForGetSolution(oc)
}