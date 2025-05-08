package sentio

// Base simulation fields that are common between requests and responses
type SimulationSpec struct {
	NetworkID string `json:"networkId"`
	ChainID   string `json:"chainId,omitempty"`
	ChainSpec struct {
		ChainID string `json:"chainId"`
		ForkID  string `json:"forkId"`
	} `json:"chainSpec"`
	To                   string `json:"to"`
	From                 string `json:"from"`
	Input                string `json:"input"`
	BlockNumber          string `json:"blockNumber"`
	TransactionIndex     string `json:"transactionIndex"`
	Gas                  string `json:"gas"`
	GasPrice             string `json:"gasPrice"`
	MaxFeePerGas         string `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas string `json:"maxPriorityFeePerGas,omitempty"`
	Value                string `json:"value"`
	AccessList           []struct {
		Address     string   `json:"address"`
		StorageKeys []string `json:"storageKeys"`
	} `json:"accessList,omitempty"`
}

type TransactionReceipt struct {
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	EffectiveGasPrice string `json:"effectiveGasPrice"`
	Status            string `json:"status"`
	Error             string `json:"error,omitempty"`
	RevertReason      string `json:"revertReason,omitempty"`
	Logs              []any  `json:"logs"`
}

type SimulationResponse struct {
	Simulation struct {
		SimulationSpec
		Id        string `json:"id"`
		CreatedAt string `json:"createAt"`
		BundleID  string `json:"bundleId,omitempty"`
	} `json:"simulation"`
	Result struct {
		TransactionReceipt TransactionReceipt `json:"transactionReceipt"`
	} `json:"result"`
}

// // CallTrace represents a call trace from Sentio API
// type CallTrace struct {
// 	TxHash           string        `json:"txHash,omitempty"`
// 	SimulationId     string        `json:"simulationId,omitempty"`
// 	BlockNumber      int64         `json:"blockNumber"`
// 	From             string        `json:"from"`
// 	To               string        `json:"to"`
// 	Value            string        `json:"value,omitempty"`
// 	Gas              int64         `json:"gas"`
// 	GasUsed          int64         `json:"gasUsed"`
// 	Input            string        `json:"input"`
// 	Output           string        `json:"output"`
// 	Error            string        `json:"error,omitempty"`
// 	Success          bool          `json:"success"`
// 	ReturnValues     []ReturnValue `json:"returnValues,omitempty"`
// 	Calls            []CallTrace   `json:"calls,omitempty"`
// 	StartIndex       int           `json:"startIndex"`
// 	EndIndex         int           `json:"endIndex"`
// 	DecodedFnName    string        `json:"decodedFnName,omitempty"`
// 	DecodedFnArgs    []FnArg       `json:"decodedFnArgs,omitempty"`
// 	DecodedFnSig     string        `json:"decodedFnSig,omitempty"`
// 	ContractCreation bool          `json:"contractCreation,omitempty"`
// }

// // ReturnValue represents a returned value in a call trace
// type ReturnValue struct {
// 	Name  string `json:"name"`
// 	Type  string `json:"type"`
// 	Value any    `json:"value"`
// }

// // FnArg represents a function argument in a decoded call
// type FnArg struct {
// 	Name  string `json:"name"`
// 	Type  string `json:"type"`
// 	Value any    `json:"value"`
// }

// // StateDiff represents a state diff from Sentio API
// type StateDiff struct {
// 	BlockNumber  int64         `json:"blockNumber"`
// 	TxHash       string        `json:"txHash,omitempty"`
// 	SimulationId string        `json:"simulationId,omitempty"`
// 	StateChanges []StateChange `json:"stateChanges"`
// }

// // StateChange represents a state change for a specific address
// type StateChange struct {
// 	Address        string          `json:"address"`
// 	BalanceChanges []BalanceChange `json:"balanceChanges,omitempty"`
// 	StorageChanges []StorageChange `json:"storageChanges,omitempty"`
// 	NonceChanges   []NonceChange   `json:"nonceChanges,omitempty"`
// 	CodeChanges    []CodeChange    `json:"codeChanges,omitempty"`
// }

// // BalanceChange represents a balance change in a state diff
// type BalanceChange struct {
// 	PrevBalance string `json:"prevBalance"`
// 	Balance     string `json:"balance"`
// }

// // StorageChange represents a storage change in a state diff
// type StorageChange struct {
// 	Slot      string `json:"slot"`
// 	PrevValue string `json:"prevValue"`
// 	Value     string `json:"value"`
// }

// // NonceChange represents a nonce change in a state diff
// type NonceChange struct {
// 	PrevNonce int `json:"prevNonce"`
// 	Nonce     int `json:"nonce"`
// }

// // CodeChange represents a code change in a state diff
// type CodeChange struct {
// 	PrevCode string `json:"prevCode"`
// 	Code     string `json:"code"`
// }

// // StateChangeSummary provides a human-readable summary of state changes
// type StateChangeSummary struct {
// 	Address     string `json:"address"`
// 	Description string `json:"description"`
// 	Type        string `json:"type"`
// 	ValueChange string `json:"valueChange,omitempty"`
// }
