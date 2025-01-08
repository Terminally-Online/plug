package simulate

import (
	"net/http"
	"solver/solver"
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

}
