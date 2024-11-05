package intent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/solver"
	"solver/types"
	"solver/utils"
)

type IntentRequest struct {
	Action  types.Action    `json:"action"`
	ChainId int             `json:"chainId"`
	From    string          `json:"from"`
	Inputs  json.RawMessage `json:"inputs"`
}

type Handler struct {
	solver *solver.Solver
}

func NewHandler(solver *solver.Solver) *Handler {
	return &Handler{
		solver: solver,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	protocol := r.URL.Query().Get("protocol")
	if protocol == "" {
		utils.MakeHttpError(w, "protocol is required", http.StatusBadRequest)
		return
	}

	handler, exists := h.solver.GetProtocolHandler(types.Protocol(protocol))
	if !exists {
		utils.MakeHttpError(w, fmt.Sprintf("unsupported protocol: %s", protocol), http.StatusBadRequest)
		return
	}

	action := types.Action(r.URL.Query().Get("action"))
	if action == "" {
		schemas := make(map[types.Action]types.ActionSchema)
		for _, supportedAction := range handler.SupportedActions() {
			schema, err := handler.GetSchema(supportedAction)
			if err != nil {
				utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
				return
			}
			schemas[supportedAction] = schema
		}

		if err := json.NewEncoder(w).Encode(schemas); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if !h.solver.SupportsAction(handler, action) {
		utils.MakeHttpError(w, fmt.Sprintf("action %s not supported by protocol %s", action, protocol), http.StatusBadRequest)
		return
	}

	schema, err := handler.GetSchema(action)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(schema); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var req IntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.validateRequest(&req); err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := h.solver.GetTransaction(req.Action, req.Inputs, req.ChainId, req.From)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) validateRequest(req *IntentRequest) error {
	if req.ChainId == 0 {
		return fmt.Errorf("chainId is required")
	}

	if req.From == "" {
		return fmt.Errorf("from address is required")
	}

	return nil
}
