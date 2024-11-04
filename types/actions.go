package types

type Action string

const (
	ActionDeposit      Action = "deposit"
	ActionBorrow       Action = "borrow"
	ActionRedeem       Action = "redeem"
	ActionRepay        Action = "repay"
	ActionHarvest      Action = "harvest"
	ActionTransfer     Action = "transfer"
	ActionTransferFrom Action = "transfer_from"
	ActionApprove      Action = "approve"
	ActionSwap         Action = "swap"
	ActionRoute        Action = "route"
)

type Protocol string

const (
	ProtocolAaveV2  Protocol = "aave_v2"
	ProtocolAaveV3  Protocol = "aave_v3"
	ProtocolYearnV3 Protocol = "yearn_v3"
)
