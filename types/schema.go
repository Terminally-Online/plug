package types

type Schema struct {
	Fields   []SchemaField `json:"fields"`
	Required []string      `json:"required"`
}

type SchemaField struct {
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Options     []Option `json:"options,omitempty"`
}

type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Icon  string `json:"icon,omitempty"`
}

type ActionSchema struct {
	Protocol Protocol `json:"protocol"`
	Schema   Schema   `json:"schema"`
}
