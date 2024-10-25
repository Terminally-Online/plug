package protocols

import (
	"solver/types"
)

// ProtocolHandler interface with both GET and POST operations
type ProtocolHandler interface {
	SupportedActions() []types.Action

	// POST handlers - execute actions
	HandlePostDeposit(inputs *types.DepositInputs) (*types.Transaction, error)
	HandlePostBorrow(inputs *types.BorrowInputs) (*types.Transaction, error)

	// GET handlers - return schemas
	HandleGetDeposit() types.ActionSchema
	HandleGetBorrow() types.ActionSchema
}
