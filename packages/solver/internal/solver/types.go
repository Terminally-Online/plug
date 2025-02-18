package solver

import (
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
)

type SolutionStatus struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type Solution struct {
	Transactions []signature.Plug               `json:"transactions"`
	Plug         *simulation.SimulationRequest  `json:"plug,omitempty"`
	Simulation   *simulation.SimulationResponse `json:"simulation,omitempty"`
}
