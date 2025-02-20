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
	LivePlugs    *signature.LivePlugs            `json:"livePlugs,omitempty"`
	Transaction  *simulation.SimulationRequest  `json:"transaction,omitempty"`
	Simulation   *simulation.SimulationResponse `json:"simulation,omitempty"`
}
