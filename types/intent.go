package types

type Intent struct {
	Socket       string        `json:"socket"`
	Transactions []Transaction `json:"plugs"`
	Solver       string        `json:"solver"`
	Salt         string        `json:"salt"`
}
