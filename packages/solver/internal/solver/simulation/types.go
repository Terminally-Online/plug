package simulation

import "github.com/ethereum/go-ethereum/common"

type SimulationResult struct {
	Success      bool             `json:"success"`
	GasUsed      uint64           `json:"gasUsed"`
	ReturnData   []byte           `json:"returnData"`
	StateChanges []StateChange    `json:"stateChanges"`
	Error        *SimulationError `json:"error,omitempty"`
}

type StateChange struct {
	Address common.Address `json:"address"`
	Key     string         `json:"key"`
	OldVal  string         `json:"oldVal"`
	NewVal  string         `json:"newVal"`
}

type SimulationError struct {
	Message      string `json:"message"`
	RevertData   []byte `json:"revertData,omitempty"`
	RevertReason string `json:"revertReason,omitempty"`
}
