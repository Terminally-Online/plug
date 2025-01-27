package simulation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type SimulationRequest struct {
	ChainId        uint64           `json:"chainId"`
	From           common.Address   `json:"from"`
	To             common.Address   `json:"to"`
	Data           hexutil.Bytes    `json:"data,omitempty"`
	GasLimit       *uint64          `json:"gasLimit,omitempty"`
	Value          *big.Int         `json:"value,omitempty"`
	AccessList     types.AccessList `json:"accessList,omitempty"`
	BlockNumber    *uint64          `json:"blockNumber,omitempty"`
	BlockTimestamp *uint64          `json:"blockTimestamp,omitempty"`
}

type SimulationResponse struct {
	GasUsed      uint64       `json:"gasUsed"`
	BlockNumber  uint64       `json:"blockNumber"`
	Success      bool         `json:"success"`
	Logs         []*types.Log `json:"logs"`
	ReturnData   []byte       `json:"returnData"`
	ErrorMessage string       `json:"errorMessage,omitempty"`
}

type CallTrace struct {
	CallType string         `json:"callType"`
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Value    *big.Int       `json:"value"`
}

type CallRawRequest struct {
	From        common.Address
	To          common.Address
	Value       *big.Int
	Data        []byte
	AccessList  types.AccessList
	FormatTrace bool
}

type CallRawResult struct {
	GasUsed        uint64
	BlockNumber    uint64
	Success        bool
	Trace          []CallTrace
	Logs           []*types.Log
	ExitReason     string
	ReturnData     []byte
	FormattedTrace *string
}
