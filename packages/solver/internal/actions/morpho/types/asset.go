package types

type Asset struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}
