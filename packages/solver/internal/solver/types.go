package solver

import (
	"solver/internal/database/models"
	"solver/internal/solver/signature"
)

type SolutionStatus struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type Solution struct {
	IntentId     string               `json:"intentId"`
	LivePlugs    *signature.LivePlugs `json:"-"`
	Transactions []signature.Plug     `json:"transactions,omitempty"`
	Run          *models.Run          `json:"run,omitempty"`
}
