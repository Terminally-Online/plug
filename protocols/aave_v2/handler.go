package aave_v2

import (
	"solver/types"
)

type Handler struct {
	depositSchema types.ActionSchema
	borrowSchema  types.ActionSchema
}

func New() *Handler {
	h := &Handler{}
	h.initSchemas()
	return h
}

func (h *Handler) initSchemas() {
	h.depositSchema = types.ActionSchema{
		Protocol: types.ProtocolAave,
		Schema: types.Schema{
			Fields: []types.SchemaField{
				{
					Name:        "tokenIn",
					Type:        "address",
					Description: "Token to deposit",
					Options: []types.Option{
						{
							Value:       "0x6b175474e89094c44da98b954eedeac495271d0f",
							Label:       "DAI",
							Description: "Dai Stablecoin",
						},
						{
							Value:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
							Label:       "WETH",
							Description: "Wrapped Ether",
						},
					},
				},
				{
					Name:        "amountIn",
					Type:        "uint256",
					Description: "Amount to deposit in wei",
				},
			},
			Required: []string{"tokenIn", "amountIn"},
		},
	}

	h.borrowSchema = types.ActionSchema{
		Protocol: types.ProtocolAave,
		Schema: types.Schema{
			Fields: []types.SchemaField{
				{
					Name:        "tokenToBorrow",
					Type:        "address",
					Description: "Token to borrow",
					Options: []types.Option{
						{
							Value:       "0x6b175474e89094c44da98b954eedeac495271d0f",
							Label:       "DAI",
							Description: "Dai Stablecoin",
						},
					},
				},
			},
			Required: []string{"tokenToBorrow", "amount"},
		},
	}
}

func (h *Handler) SupportedActions() []types.Action {
	return []types.Action{types.ActionDeposit, types.ActionBorrow}
}

func (h *Handler) HandleGetDeposit() types.ActionSchema {
	return types.ActionSchema{
		Protocol: types.ProtocolAave,
		Schema: types.Schema{
			Fields: []types.SchemaField{
				{
					Name:        "tokenIn",
					Type:        "address",
					Description: "Token to deposit",
					Options: []types.Option{
						{
							Value:       "0x6b175474e89094c44da98b954eedeac495271d0f",
							Label:       "DAI",
							Description: "Dai Stablecoin",
						},
						{
							Value:       "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
							Label:       "WETH",
							Description: "Wrapped Ether",
						},
					},
				},
				{
					Name:        "amountIn",
					Type:        "uint256",
					Description: "Amount to deposit in wei",
				},
			},
			Required: []string{"tokenIn", "amountIn"},
		},
	}
}

func (h *Handler) HandleGetBorrow() types.ActionSchema {
	return h.borrowSchema
}

func (h *Handler) HandlePostDeposit(inputs *types.DepositInputs) (*types.Transaction, error) {
	// Implement Aave deposit logic
	return &types.Transaction{
		To: inputs.Target,
		// Add other transaction details
	}, nil
}

func (h *Handler) HandlePostBorrow(inputs *types.BorrowInputs) (*types.Transaction, error) {
	// Implement Aave borrow logic
	return &types.Transaction{
		To: inputs.Target,
		// Add other transaction details
	}, nil
}
