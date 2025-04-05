package types

import "github.com/ethereum/go-ethereum/common/hexutil"

type ExecutionResponse struct {
	ProofOfTask      string         `json:"proofOfTask,omitempty"`
	Data             *hexutil.Bytes `json:"data,omitempty"`
	TaskDefinitionId int            `json:"taskDefinitionId,omitempty"`
}
