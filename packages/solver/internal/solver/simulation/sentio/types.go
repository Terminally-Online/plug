package sentio

// CallTrace represents a call trace from Sentio API
type CallTrace struct {
	TxHash           string      `json:"txHash,omitempty"`
	SimulationId     string      `json:"simulationId,omitempty"`
	BlockNumber      int64       `json:"blockNumber"`
	From             string      `json:"from"`
	To               string      `json:"to"`
	Value            string      `json:"value,omitempty"`
	Gas              int64       `json:"gas"`
	GasUsed          int64       `json:"gasUsed"`
	Input            string      `json:"input"`
	Output           string      `json:"output"`
	Error            string      `json:"error,omitempty"`
	Success          bool        `json:"success"`
	ReturnValues     []ReturnValue `json:"returnValues,omitempty"`
	Calls            []CallTrace `json:"calls,omitempty"`
	StartIndex       int         `json:"startIndex"`
	EndIndex         int         `json:"endIndex"`
	DecodedFnName    string      `json:"decodedFnName,omitempty"`
	DecodedFnArgs    []FnArg     `json:"decodedFnArgs,omitempty"`
	DecodedFnSig     string      `json:"decodedFnSig,omitempty"`
	ContractCreation bool        `json:"contractCreation,omitempty"`
}

// ReturnValue represents a returned value in a call trace
type ReturnValue struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}

// FnArg represents a function argument in a decoded call
type FnArg struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}

// StateDiff represents a state diff from Sentio API
type StateDiff struct {
	BlockNumber   int64         `json:"blockNumber"`
	TxHash        string        `json:"txHash,omitempty"`
	SimulationId  string        `json:"simulationId,omitempty"`
	StateChanges  []StateChange `json:"stateChanges"`
}

// StateChange represents a state change for a specific address
type StateChange struct {
	Address        string           `json:"address"`
	BalanceChanges []BalanceChange  `json:"balanceChanges,omitempty"`
	StorageChanges []StorageChange  `json:"storageChanges,omitempty"`
	NonceChanges   []NonceChange    `json:"nonceChanges,omitempty"`
	CodeChanges    []CodeChange     `json:"codeChanges,omitempty"`
}

// BalanceChange represents a balance change in a state diff
type BalanceChange struct {
	PrevBalance string `json:"prevBalance"`
	Balance     string `json:"balance"`
}

// StorageChange represents a storage change in a state diff
type StorageChange struct {
	Slot      string `json:"slot"`
	PrevValue string `json:"prevValue"`
	Value     string `json:"value"`
}

// NonceChange represents a nonce change in a state diff
type NonceChange struct {
	PrevNonce int `json:"prevNonce"`
	Nonce     int `json:"nonce"`
}

// CodeChange represents a code change in a state diff
type CodeChange struct {
	PrevCode string `json:"prevCode"`
	Code     string `json:"code"`
}

// StateChangeSummary provides a human-readable summary of state changes
type StateChangeSummary struct {
	Address     string `json:"address"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ValueChange string `json:"valueChange,omitempty"`
}