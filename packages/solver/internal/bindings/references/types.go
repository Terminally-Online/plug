package references

type Network struct {
	Name       string                       `json:"name"`
	ChainIds   []uint64                     `json:"chainIds"`
	Explorer   string                       `json:"explorer"`
	References map[string]map[string]string `json:"references"`
	Icon       struct {
		Default string `json:"default"`
	} `json:"icon"`
}

type AdditionalSource struct {
	Filename   string
	SourceCode string
}
