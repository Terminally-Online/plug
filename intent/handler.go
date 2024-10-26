// File: intent/handler.go
package intent

import (
	"encoding/json"
	"net/http"
	"solver/solver"
	"solver/types"
	"solver/utils"
)

type Handler struct {
	solver *solver.Solver
}

func NewHandler(solver *solver.Solver) *Handler {
	return &Handler{
		solver: solver,
	}
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")
	if action == "" {
		utils.MakeHttpError(w, "action parameter is required", http.StatusBadRequest)
		return
	}

	protocol := r.URL.Query().Get("protocol")
	if protocol == "" {
		utils.MakeHttpError(w, "protocol parameter is required", http.StatusBadRequest)
		return
	}

	handler, exists := h.solver.GetProtocolHandler(types.Protocol(protocol))
	if !exists {
		utils.MakeHttpError(w, "unsupported protocol", http.StatusBadRequest)
		return
	}

	var schema types.ActionSchema
	switch types.Action(action) {
	case types.ActionDeposit:
		schema = handler.HandleGetDeposit()
	case types.ActionBorrow:
		schema = handler.HandleGetBorrow()
	default:
		utils.MakeHttpError(w, "unsupported action", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(schema); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

type IntentRequest struct {
	Action  types.Action    `json:"action"`
	ChainId int             `json:"chainId"`
	From    string          `json:"from"`
	Inputs  json.RawMessage `json:"inputs"`
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var req IntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.MakeHttpError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

    if req.ChainId == 0 {
        utils.MakeHttpError(w, "chainId is required", http.StatusBadRequest)
        return
    }
    if req.From == "" {
        utils.MakeHttpError(w, "from address is required", http.StatusBadRequest)
        return
    }

	var actionInputs types.ActionInputs
	switch req.Action {
	case types.ActionDeposit:
		var depositInputs types.DepositInputs
		if err := json.Unmarshal(req.Inputs, &depositInputs); err != nil {
			utils.MakeHttpError(w, "invalid deposit inputs: "+err.Error(), http.StatusBadRequest)
			return
		}
		actionInputs = &depositInputs
	case types.ActionBorrow:
		var borrowInputs types.BorrowInputs
		if err := json.Unmarshal(req.Inputs, &borrowInputs); err != nil {
			utils.MakeHttpError(w, "invalid borrow inputs: "+err.Error(), http.StatusBadRequest)
			return
		}
		actionInputs = &borrowInputs

	default:
		utils.MakeHttpError(w, "unsupported action", http.StatusBadRequest)
		return
	}

	if err := actionInputs.Validate(); err != nil {
		utils.MakeHttpError(w, "invalid inputs: "+err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := h.solver.BuildTransaction(req.Action, actionInputs, req.ChainId, req.From)
	if err != nil {
		utils.MakeHttpError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(transaction); err != nil {
		utils.MakeHttpError(w, "failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

