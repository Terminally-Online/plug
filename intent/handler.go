package intent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/protocols"
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

	actionSupported := false
	for _, a := range handler.SupportedActions() {
		if a == types.Action(action) {
			actionSupported = true
			break
		}
	}
	if !actionSupported {
		utils.MakeHttpError(w, fmt.Sprintf("action %s not supported by protocol", action), http.StatusBadRequest)
		return
	}

	var schema types.ActionSchema
	switch types.Action(action) {
	case types.ActionDeposit:
		if depositHandler, ok := handler.(protocols.DepositHandler); ok {
			schema = depositHandler.HandleGetDeposit()
		} else {
			utils.MakeHttpError(w, "protocol does not implement deposit handler", http.StatusBadRequest)
			return
		}
	case types.ActionBorrow:
		if borrowHandler, ok := handler.(protocols.BorrowHandler); ok {
			schema = borrowHandler.HandleGetBorrow()
		} else {
			utils.MakeHttpError(w, "protocol does not implement borrow handler", http.StatusBadRequest)
			return
		}

	case types.ActionRedeem:
		if redeemHandler, ok := handler.(protocols.RedeemHandler); ok {
			schema = redeemHandler.HandleGetRedeem()
		} else {
			utils.MakeHttpError(w, "protocol does not implement borrow handler", http.StatusBadRequest)
			return
		}

	case types.ActionRepay:
		if repayHandler, ok := handler.(protocols.RepayHandler); ok {
			schema = repayHandler.HandleGetRepay()
		} else {
			utils.MakeHttpError(w, "protocol does not implement borrow handler", http.StatusBadRequest)
			return
		}

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

	case types.ActionRedeem:
		var redeemInputs types.RedeemInputs
		if err := json.Unmarshal(req.Inputs, &redeemInputs); err != nil {
			utils.MakeHttpError(w, "invalid borrow inputs: "+err.Error(), http.StatusBadRequest)
			return
		}
		actionInputs = &redeemInputs

	case types.ActionRepay:
		var repayInputs types.RepayInputs
		if err := json.Unmarshal(req.Inputs, &repayInputs); err != nil {
			utils.MakeHttpError(w, "invalid borrow inputs: "+err.Error(), http.StatusBadRequest)
			return
		}
		actionInputs = &repayInputs

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

func (h *Handler) IsActionSupported(protocol protocols.BaseProtocolHandler, action types.Action) bool {
	for _, a := range protocol.SupportedActions() {
		if a == action {
			return true
		}
	}
	return false
}
