package solver

import (
	"solver/internal/database/models"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/core/types"
)

type SolutionStatus struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type Transaction struct {
	From       string           `json:"from"`
	To         string           `json:"to"`
	ChainId    uint64           `json:"chainId"`
	Value      string           `json:"value"`
	Data       string           `json:"data"`
	Gas        string           `json:"gas"`
	AccessList types.AccessList `json:"accessList"`
}

type Solution struct {
	Transactions []signature.Plug     `json:"transactions"`
	LivePlugs    *signature.LivePlugs `json:"livePlugs,omitempty"`
	Intent       *models.Intent       `json:"intent,omitempty"`
	Run          *models.Run          `json:"simulation,omitempty"`
	Transaction  *Transaction         `json:"transaction,omitempty"`
}
