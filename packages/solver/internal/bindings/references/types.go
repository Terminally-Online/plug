package references

import "github.com/ethereum/go-ethereum/accounts/abi"

type Network struct {
	Name       string                       `json:"name"`
	ChainIds   []uint64                     `json:"chainIds"`
	References map[string]map[string]string `json:"-"`
	Icon       struct {
		Default string `json:"default"`
	} `json:"icon"`
}

type AdditionalSource struct {
	Filename   string
	SourceCode string
}

type EventReference struct {
	Name      string
	Signature string
	Inputs    []abi.Argument
	Contract  string
}

type EventInput struct {
	Name       string       `json:"name"`
	Type       string       `json:"type"`
	Indexed    bool         `json:"indexed"`
	Components []EventInput `json:"components"`
}
