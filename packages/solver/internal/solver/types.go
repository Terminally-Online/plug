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
	Status       SolutionStatus            `json:"status"`
	Transactions *[]models.Transaction     `json:"transactions"`
	LivePlugs    *signature.LivePlugs      `json:"livePlugs,omitempty"`
	Intent       *models.Intent            `json:"intent,omitempty"`
	Run          *models.Run               `json:"simulation,omitempty"`
	Transaction  *models.TransactionBundle `json:"transaction,omitempty"`
}
