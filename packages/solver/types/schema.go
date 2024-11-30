package types

type ProtocolMetadata struct {
	Icon string   `json:"icon"`
	Tags []string `json:"tags"`
}

type ProtocolSchema struct {
	Metadata ProtocolMetadata  `json:"metadata"`
	Schema   map[Action]Schema `json:"schema"`
}

type Option struct {
	Value string `json:"value"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Info  string `json:"info,omitempty"`
	Icon  string `json:"icon,omitempty"`
}

type Schema struct {
	Sentence string           `json:"sentence"`
	Options  map[int][]Option `json:"options,omitempty"`
}

var (
	BaseThresholdFields = []Option{
		{Label: "less than", Name: "Less Than", Value: "-1"},
		{Label: "greater than", Name: "Greater Than", Value: "1"},
	}
)
