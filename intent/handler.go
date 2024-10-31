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

type IntentRequest struct {
	Action  types.Action    `json:"action"`
	ChainId int             `json:"chainId"`
	From    string          `json:"from"`
	Inputs  json.RawMessage `json:"inputs"`
}

type ActionHandler struct {
	getSchema       func(interface{}) types.ActionSchema
	unmarshalInputs func(json.RawMessage) (types.ActionInputs, error)
	assertHandler   func(interface{}) bool
	handlerErrorMsg string
}

type Handler struct {
	solver         *solver.Solver
	actionHandlers map[types.Action]ActionHandler
}

func NewHandler(solver *solver.Solver) *Handler {
	h := &Handler{
		solver:         solver,
		actionHandlers: make(map[types.Action]ActionHandler),
	}

	h.actionHandlers[types.ActionDeposit] = ActionHandler{
		getSchema: func(h interface{}) types.ActionSchema {
			return h.(protocols.DepositHandler).HandleGetDeposit()
		},
		unmarshalInputs: func(data json.RawMessage) (types.ActionInputs, error) {
			var inputs types.DepositInputs
			err := json.Unmarshal(data, &inputs)
			return &inputs, err
		},
		assertHandler: func(h interface{}) bool {
			_, ok := h.(protocols.DepositHandler)
			return ok
		},
		handlerErrorMsg: "protocol does not implement deposit handler",
	}

	h.actionHandlers[types.ActionBorrow] = ActionHandler{
		getSchema: func(h interface{}) types.ActionSchema {
			return h.(protocols.BorrowHandler).HandleGetBorrow()
		},
		unmarshalInputs: func(data json.RawMessage) (types.ActionInputs, error) {
			var inputs types.BorrowInputs
			err := json.Unmarshal(data, &inputs)
			return &inputs, err
		},
		assertHandler: func(h interface{}) bool {
			_, ok := h.(protocols.BorrowHandler)
			return ok
		},
		handlerErrorMsg: "protocol does not implement borrow handler",
	}

	h.actionHandlers[types.ActionRedeem] = ActionHandler{
		getSchema: func(h interface{}) types.ActionSchema {
			return h.(protocols.RedeemHandler).HandleGetRedeem()
		},
		unmarshalInputs: func(data json.RawMessage) (types.ActionInputs, error) {
			var inputs types.RedeemInputs
			err := json.Unmarshal(data, &inputs)
			return &inputs, err
		},
		assertHandler: func(h interface{}) bool {
			_, ok := h.(protocols.RedeemHandler)
			return ok
		},
		handlerErrorMsg: "protocol does not implement redeem handler",
	}

	h.actionHandlers[types.ActionRepay] = ActionHandler{
		getSchema: func(h interface{}) types.ActionSchema {
			return h.(protocols.RepayHandler).HandleGetRepay()
		},
		unmarshalInputs: func(data json.RawMessage) (types.ActionInputs, error) {
			var inputs types.RepayInputs
			err := json.Unmarshal(data, &inputs)
			return &inputs, err
		},
		assertHandler: func(h interface{}) bool {
			_, ok := h.(protocols.RepayHandler)
			return ok
		},
		handlerErrorMsg: "protocol does not implement repay handler",
	}

	return h
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	action := types.Action(r.URL.Query().Get("action"))
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
		utils.MakeHttpError(w, fmt.Sprintf("unsupported protocol: %s", protocol), http.StatusBadRequest)
		return
	}

	actionHandler, exists := h.actionHandlers[action]
	if !exists || !h.IsActionSupported(handler, action) {
		utils.MakeHttpError(w, fmt.Sprintf("unsupported action: %s", action), http.StatusBadRequest)
		return
	}

	if !actionHandler.assertHandler(handler) {
		utils.MakeHttpError(w, actionHandler.handlerErrorMsg, http.StatusBadRequest)
		return
	}

	schema := actionHandler.getSchema(handler)
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

	if req.ChainId == 0 {
		utils.MakeHttpError(w, "chainId is required", http.StatusBadRequest)
		return
	}

	if req.From == "" {
		utils.MakeHttpError(w, "from address is required", http.StatusBadRequest)
		return
	}

	actionHandler, exists := h.actionHandlers[req.Action]
	if !exists {
		utils.MakeHttpError(w, fmt.Sprintf("unsupported action: %s", req.Action), http.StatusBadRequest)
		return
	}

	actionInputs, err := actionHandler.unmarshalInputs(req.Inputs)
	if err != nil {
		utils.MakeHttpError(w, fmt.Sprintf("invalid %s inputs: %v", req.Action, err), http.StatusBadRequest)
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
