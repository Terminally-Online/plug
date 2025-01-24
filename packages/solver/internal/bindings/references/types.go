package references

type Network struct {
	ChainIds   []int
	Explorer   string
	References map[string]map[string]string
}

type AdditionalSource struct {
	Filename   string
	SourceCode string
}
