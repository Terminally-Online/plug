package types

type ProtocolMetadata struct {
	Icon string `json:"icon"`
}

type ProtocolSchema struct {
	Metadata ProtocolMetadata  `json:"metadata"`
	Schema   map[Action]Schema `json:"schema"`
}

type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Icon  string `json:"icon,omitempty"`
}

type SchemaField struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Options []Option `json:"options,omitempty"`
}

type Schema struct {
	Sentence string        `json:"sentence"`
	Fields   []SchemaField `json:"fields,omitempty"`
}

var (
	BaseDepositSchema = Schema{
		Sentence: "Deposit {0} {1}.",
		Fields: []SchemaField{
			{
				Name: "tokenIn",
				Type: "address",
			},
			{
				Name: "tokenOut",
				Type: "address",
			},
			{
				Name: "amountIn",
				Type: "uint256",
			},
		},
	}

	BaseBorrowSchema = Schema{
		Sentence: "Borrow {0} {1}.",
		Fields: []SchemaField{
			{
				Name: "tokenOut",
				Type: "address",
			},
			{
				Name: "amountOut",
				Type: "uint256",
			},
		},
	}
)
