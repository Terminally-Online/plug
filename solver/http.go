package solver

import (
	"encoding/json"
	"net/http"
	"solver/types"
)

type Server struct {
	solver *Solver
}

func NewServer(solver *Solver) *Server {
	return &Server{solver: solver}
}

func (s *Server) GetActionSchema(w http.ResponseWriter, r *http.Request) {
	action := types.Action(r.URL.Query().Get("action"))
	protocol := types.Protocol(r.URL.Query().Get("protocol"))

	if protocol == "" {
		http.Error(w, "protocol must be specified", http.StatusBadRequest)
		return
	}

	handler, exists := s.solver.protocols[protocol]
	if !exists {
		http.Error(w, "unsupported protocol", http.StatusBadRequest)
		return
	}

	var schema types.ActionSchema
	switch action {
	case types.ActionDeposit:
		schema = handler.HandleGetDeposit()
	case types.ActionBorrow:
		schema = handler.HandleGetBorrow()
	default:
		http.Error(w, "unsupported action", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(schema)
}

func (s *Server) BuildTransaction(w http.ResponseWriter, r *http.Request) {
	var inputs struct {
		Action types.Action    `json:"action"`
		Inputs json.RawMessage `json:"inputs"`
	}

	if err := json.NewDecoder(r.Body).Decode(&inputs); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var actionInputs types.ActionInputs
	switch inputs.Action {
	case types.ActionDeposit:
		var depositInputs types.DepositInputs
		if err := json.Unmarshal(inputs.Inputs, &depositInputs); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		actionInputs = &depositInputs
	case types.ActionBorrow:
		var borrowInputs types.BorrowInputs
		if err := json.Unmarshal(inputs.Inputs, &borrowInputs); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		actionInputs = &borrowInputs
	default:
		http.Error(w, "unsupported action", http.StatusBadRequest)
		return
	}

	tx, err := s.solver.BuildTransaction(inputs.Action, actionInputs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(tx)
}
