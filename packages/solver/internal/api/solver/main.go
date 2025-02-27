package solver

import (
	"solver/internal/solver"
)

type Handler struct {
	Solver solver.Solver
}

func New() Handler {
	return Handler{
		Solver: solver.New(),
	}
}

