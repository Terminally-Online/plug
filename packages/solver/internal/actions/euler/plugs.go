package euler

import (
	"encoding/json"
	"solver/internal/actions"
	"solver/internal/solver/signature"
)

func HandleSupply(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}

func HandleRepay(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}

func HandleRepayWithShares(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}

func HandleWithdraw(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}

func HandleBorrow(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}

func HandleReturn(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}

func HandleConstraintAPY(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}

func HandleConstraintHealthFactor(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}

func HandleConstraintTimeToLiquidation(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {}