package solver

import (
	"solver/internal/solver"
	"solver/internal/solver/simulation"
)

type Handler struct {
	Solver solver.Solver
	Simulator simulation.Simulator
}

func New() Handler {
	return Handler{
		Solver: solver.New(),
		Simulator: simulation.New(),
	}
}

