package types

type ProtocolMetadata struct {
	Icon string `json:"icon"`
}

type ProtocolSchema struct {
	Metadata ProtocolMetadata  `json:"metadata"`
	Schema   map[Action]Schema `json:"schema"`
}

type Schema []SchemaField

type SchemaField struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Options []Option `json:"options,omitempty"`
}

type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Icon  string `json:"icon,omitempty"`
}

var (
	BaseDepositSchema = Schema{
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
	}

	BaseBorrowSchema = Schema{
		{
			Name: "tokenOut",
			Type: "address",
		},
		{
			Name: "amountOut",
			Type: "uint256",
		},
	}
)
