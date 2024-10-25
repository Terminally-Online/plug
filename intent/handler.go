package intent

import (
	"encoding/json"
	"net/http"
	"solver/solver"
	"solver/types"
)

type Handler struct {
	solver *solver.Solver
}

func NewHandler(solver *solver.Solver) *Handler {
	return &Handler{
		solver: solver,
	}
}

// Get handles GET requests to /intent
// Returns schema for specified action and protocol
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")
	if action == "" {
		http.Error(w, "action parameter is required", http.StatusBadRequest)
		return
	}

	protocol := r.URL.Query().Get("protocol")
	if protocol == "" {
		http.Error(w, "protocol parameter is required", http.StatusBadRequest)
		return
	}

	handler, exists := h.solver.GetProtocolHandler(types.Protocol(protocol))
	if !exists {
		http.Error(w, "unsupported protocol", http.StatusBadRequest)
		return
	}

	var schema types.ActionSchema
	switch types.Action(action) {
	case types.ActionDeposit:
		schema = handler.HandleGetDeposit()
	case types.ActionBorrow:
		schema = handler.HandleGetBorrow()
	default:
		http.Error(w, "unsupported action", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(schema); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// IntentRequest represents the incoming POST request structure
type IntentRequest struct {
	Action types.Action    `json:"action"`
	Inputs json.RawMessage `json:"inputs"`
}

// Post handles POST requests to /intent
// Builds and returns a transaction based on the provided action and inputs
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var req IntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	var actionInputs types.ActionInputs
	switch req.Action {
	case types.ActionDeposit:
		var depositInputs types.DepositInputs
		if err := json.Unmarshal(req.Inputs, &depositInputs); err != nil {
			http.Error(w, "invalid deposit inputs: "+err.Error(), http.StatusBadRequest)
			return
		}
		actionInputs = &depositInputs

	case types.ActionBorrow:
		var borrowInputs types.BorrowInputs
		if err := json.Unmarshal(req.Inputs, &borrowInputs); err != nil {
			http.Error(w, "invalid borrow inputs: "+err.Error(), http.StatusBadRequest)
			return
		}
		actionInputs = &borrowInputs

	default:
		http.Error(w, "unsupported action", http.StatusBadRequest)
		return
	}

	// Validate inputs
	if err := actionInputs.Validate(); err != nil {
		http.Error(w, "invalid inputs: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Build transaction
	tx, err := h.solver.BuildTransaction(req.Action, actionInputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(tx); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
