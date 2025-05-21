package simulation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type SimulationRequest struct {
	ChainId string   `json:"chainId"`
	From    string   `json:"from"`
	To      string   `json:"to"`
	Data    string   `json:"data"`
	Value   *big.Int `json:"value"`
	Gas     *big.Int `json:"gas"`
}

type Location struct {
	InstructionIndex int `json:"instructionIndex"`
}

type Call struct {
	Depth        int      `json:"depth"`
	From         string   `json:"from"`
	To           string   `json:"to"`
	StartIndex   int      `json:"startIndex"`
	EndIndex     int      `json:"endIndex"`
	Value        string   `json:"value"`
	RawInput     string   `json:"rawInput"`
	RawOutput    string   `json:"rawOutput"`
	Type         string   `json:"type"`
	Calls        []Call   `json:"calls"`
	Logs         []Log    `json:"logs"`
	Location     Location `json:"location"`
	FunctionName string   `json:"functionName"`
	Gas          string   `json:"gas"`
	GasUsed      string   `json:"gasUsed"`
	Error        string   `json:"error,omitempty"`
}

type Trace struct {
	Type     string         `json:"type"`
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Value    string         `json:"value"`
	Gas      string         `json:"gas"`
	GasUsed  string         `json:"gasUsed"`
	GasPrice string         `json:"gasPrice"`
	Input    hexutil.Bytes  `json:"input"`
	Output   hexutil.Bytes  `json:"output"`
	Error    string         `json:"error"`
	Calls    []Call         `json:"calls"`
	Logs     []Log          `json:"logs"`
}

type Log struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    hexutil.Bytes  `json:"data"`
}

type EventParameter struct {
	Name    string
	Type    string
	Indexed bool
	Value   interface{}
}

type DecodedLog struct {
	Name       *string
	Parameters []EventParameter
	Raw        Log
}
