package solver

import (
	"solver/internal/database/models"
	"solver/internal/solver/signature"
	"solver/internal/solver/simulation"
)

type SolutionStatus struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type Solution struct {
	Transactions []signature.Plug        `json:"transactions"`
	LivePlugs    *signature.LivePlugs    `json:"livePlugs,omitempty"`
	Intent       *models.Intent          `json:"intent,omitempty"`
	Run          *models.Run             `json:"simulation,omitempty"`
	Transaction  *simulation.Transaction `json:"transaction,omitempty"`
}
