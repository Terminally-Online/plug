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
	LivePlugs    *signature.LivePlugs     `json:"-"`
	Transactions []*signature.MinimalPlug `json:"transactions,omitempty"`
	Run          *models.Run              `json:"run,omitempty"`
}
