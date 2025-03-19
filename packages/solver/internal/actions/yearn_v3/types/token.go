package types

type Token struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	LogoURI  string `json:"logoURI"`
	ChainID  int    `json:"chainId"`
	Decimals int    `json:"decimals"`
}

type TokenList struct {
	Tokens []Token `json:"tokens"`
}
