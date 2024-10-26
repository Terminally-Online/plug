package types

type Action string

const (
	ActionDeposit Action = "deposit"
	ActionBorrow  Action = "borrow"
	ActionRedeem  Action = "redeem"
	ActionRepay   Action = "repay"
)

type Protocol string

const (
	ProtocolAaveV2  Protocol = "aave_v2"
	ProtocolAaveV3  Protocol = "aave_v3"
	ProtocolYearnV3 Protocol = "yearn_v3"
)
