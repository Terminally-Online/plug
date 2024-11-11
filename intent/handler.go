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
	action := types.Action(r.URL.Query().Get("action"))

	// Case 1: No protocol - return all schemas for all protocols
	if protocol == "" {
		allSchemas := make(map[types.Protocol]types.ProtocolSchema)

		for protocol, handler := range h.solver.GetProtocols() {
			protocolSchema := types.ProtocolSchema{
				Metadata: types.ProtocolMetadata{
					Icon: handler.GetIcon(),
					Tags: handler.GetTags(),
				},
				Schema: make(map[types.Action]types.Schema),
			}

			for _, supportedAction := range handler.GetActions() {
				schema, err := handler.GetSchema(supportedAction)
				if err != nil {
					utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
					return
				}
				protocolSchema.Schema[supportedAction] = *schema
			}
			allSchemas[protocol] = protocolSchema
		}

		if err := json.NewEncoder(w).Encode(allSchemas); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	handler, exists := h.solver.GetProtocolHandler(types.Protocol(protocol))
	if !exists {
		utils.MakeHttpError(w, fmt.Sprintf("unsupported protocol: %s", protocol), http.StatusBadRequest)
		return
	}

	// Case 2: Protocol only - return all schemas for that protocol
	if action == "" {
		protocolSchema := types.ProtocolSchema{
			Metadata: types.ProtocolMetadata{
				Icon: handler.GetIcon(),
			},
			Schema: make(map[types.Action]types.Schema),
		}

		for _, supportedAction := range handler.GetActions() {
			schema, err := handler.GetSchema(supportedAction)
			if err != nil {
				utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
				return
			}
			protocolSchema.Schema[supportedAction] = *schema
		}

		response := map[types.Protocol]types.ProtocolSchema{
			types.Protocol(protocol): protocolSchema,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Case 3: Protocol and action - return specific schema
	if !h.solver.SupportsAction(handler, action) {
		utils.MakeHttpError(w, fmt.Sprintf("action %s not supported by protocol %s", action, protocol), http.StatusBadRequest)
		return
	}

	schema, err := handler.GetSchema(action)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	protocolSchema := types.ProtocolSchema{
		Metadata: types.ProtocolMetadata{
			Icon: handler.GetIcon(),
		},
		Schema: map[types.Action]types.Schema{
			action: *schema,
		},
	}

	response := map[types.Protocol]types.ProtocolSchema{
		types.Protocol(protocol): protocolSchema,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
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
