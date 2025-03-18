package actions

var (
	BaseLendActionTypeFields = []Option{
		{Label: "Borrow", Name: "Borrow", Value: "-1"},
		{Label: "Deposit", Name: "Deposit", Value: "1"},
	}

	BaseThresholdFields = []Option{
		{Label: "less than", Name: "Less Than", Value: "-1"},
		{Label: "greater than", Name: "Greater Than", Value: "1"},
	}
)

type OptionIcon struct {
	Default   string `json:"default,omitempty"`
	Secondary string `json:"secondary,omitempty"`
}

type OptionInfo struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type Option struct {
	Label string      `json:"label"`
	Value string      `json:"value"`
	Name  string      `json:"name,omitempty"`
	Icon  *OptionIcon `json:"icon,omitempty"`
	Info  *OptionInfo `json:"info,omitempty"`
}
