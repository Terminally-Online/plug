package intent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"solver/actions"
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

type HandlerDefinition[T types.ActionInputs] struct {
    GetSchema    func(actions.BaseProtocolHandler) types.ActionSchema
    HandlerType  interface{} // Used for error messages
    InputType    T           // Used for type inference
}

func createActionHandler[T types.ActionInputs](def HandlerDefinition[T]) ActionHandler {
    return ActionHandler{
        getSchema: func(h interface{}) types.ActionSchema {
            return def.GetSchema(h.(actions.BaseProtocolHandler))
        },
        unmarshalInputs: func(data json.RawMessage) (types.ActionInputs, error) {
            var inputs T
            err := json.Unmarshal(data, &inputs)
            return inputs, err
        },
        assertHandler: func(h interface{}) bool {
            _, ok := h.(interface{ HandleGet() types.ActionSchema })
            return ok
        },
        handlerErrorMsg: fmt.Sprintf("protocol does not implement %T", def.HandlerType),
    }
}

func NewHandler(solver *solver.Solver) *Handler {
    h := &Handler{
        solver:         solver,
        actionHandlers: make(map[types.Action]ActionHandler),
    }

    // Register all action handlers using the generic function
    actionDefinitions := map[types.Action]interface{}{
        types.ActionDeposit: HandlerDefinition[*types.DepositInputs]{
            GetSchema:    func(h actions.BaseProtocolHandler) types.ActionSchema { return h.(actions.DepositHandler).HandleGetDeposit() },
            HandlerType:  (*actions.DepositHandler)(nil),
            InputType:    &types.DepositInputs{},
        },
        types.ActionBorrow: HandlerDefinition[*types.BorrowInputs]{
            GetSchema:    func(h actions.BaseProtocolHandler) types.ActionSchema { return h.(actions.BorrowHandler).HandleGetBorrow() },
            HandlerType:  (*actions.BorrowHandler)(nil),
            InputType:    &types.BorrowInputs{},
        },
        types.ActionRedeem: HandlerDefinition[*types.RedeemInputs]{
            GetSchema:    func(h actions.BaseProtocolHandler) types.ActionSchema { return h.(actions.RedeemHandler).HandleGetRedeem() },
            HandlerType:  (*actions.RedeemHandler)(nil),
            InputType:    &types.RedeemInputs{},
        },
        types.ActionRepay: HandlerDefinition[*types.RepayInputs]{
            GetSchema:    func(h actions.BaseProtocolHandler) types.ActionSchema { return h.(actions.RepayHandler).HandleGetRepay() },
            HandlerType:  (*actions.RepayHandler)(nil),
            InputType:    &types.RepayInputs{},
        },
        types.ActionHarvest: HandlerDefinition[*types.HarvestInputs]{
            GetSchema:    func(h actions.BaseProtocolHandler) types.ActionSchema { return h.(actions.HarvestHandler).HandleGetHarvest() },
            HandlerType:  (*actions.HarvestHandler)(nil),
            InputType:    &types.HarvestInputs{},
        },
    }

    for action, def := range actionDefinitions {
        switch d := def.(type) {
        case HandlerDefinition[*types.DepositInputs]:
            h.actionHandlers[action] = createActionHandler(d)
        case HandlerDefinition[*types.BorrowInputs]:
            h.actionHandlers[action] = createActionHandler(d)
        case HandlerDefinition[*types.RedeemInputs]:
            h.actionHandlers[action] = createActionHandler(d)
        case HandlerDefinition[*types.RepayInputs]:
            h.actionHandlers[action] = createActionHandler(d)
        case HandlerDefinition[*types.HarvestInputs]:
            h.actionHandlers[action] = createActionHandler(d)
        }
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

func (h *Handler) IsActionSupported(protocol actions.BaseProtocolHandler, action types.Action) bool {
	for _, a := range protocol.SupportedActions() {
		if a == action {
			return true
		}
	}
	return false
}
