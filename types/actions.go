package types

type Action string

const (
	ActionDeposit Action = "DEPOSIT"
	ActionBorrow  Action = "BORROW"
)

type Protocol string

const (
	ProtocolAave     Protocol = "AAVE"
	ProtocolCompound Protocol = "COMPOUND"
)
