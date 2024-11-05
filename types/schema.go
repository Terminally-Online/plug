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

var (
	BaseDepositSchema = Schema{
		Fields: []SchemaField{
			{
				Name:        "tokenIn",
				Type:        "address",
				Description: "Address of the token to deposit",
			},
			{
				Name:        "tokenOut",
				Type:        "address",
				Description: "Address of the token to receive",
			},
			{
				Name:        "amountIn",
				Type:        "uint256",
				Description: "Amount of tokens to deposit",
			},
		},
		Required: []string{"tokenIn", "tokenOut", "amountIn"},
	}

	BaseBorrowSchema = Schema{
		Fields: []SchemaField{
			{
				Name:        "tokenOut",
				Type:        "address",
				Description: "Address of the token to borrow",
			},
			{
				Name:        "amountOut",
				Type:        "uint256",
				Description: "Amount of tokens to borrow",
			},
		},
		Required: []string{"tokenOut", "amountOut"},
	}
)
