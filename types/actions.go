package types

type Action string

const (
	ActionDeposit      Action = "deposit"
	ActionBorrow       Action = "borrow"
	ActionRedeem       Action = "redeem"
	ActionWithdraw     Action = "withdraw"
	ActionRepay        Action = "repay"
	ActionHarvest      Action = "harvest"
	ActionTransfer     Action = "transfer"
	ActionTransferFrom Action = "transfer_from"
	ActionApprove      Action = "approve"
	ActionSwap         Action = "swap"
	ActionRoute        Action = "route"

	ConstraintHealthFactor       Action = "health_factor"
	ConstraintAPY                Action = "apy"
	ConstraintAPYDifferential    Action = "apy_differential"
	ConstraintAvailableLiquidity Action = "available_liquidity"
)

type Protocol string

const (
	ProtocolPlug    Protocol = "plug"
	ProtocolAaveV3  Protocol = "aave_v3"
	ProtocolYearnV3 Protocol = "yearn_v3"
	ProtocolNouns   Protocol = "nouns"
)
