package types

type Action string

const (
	ActionDeposit      Action = "deposit"
	ActionBorrow       Action = "borrow"
	ActionRedeem       Action = "redeem"
	ActionRedeemMax    Action = "redeem_max"
	ActionWithdraw     Action = "withdraw"
	ActionWithdrawMax  Action = "withdraw_max"
	ActionRepay        Action = "repay"
	ActionHarvest      Action = "harvest"
	ActionTransfer     Action = "transfer"
	ActionTransferFrom Action = "transfer_from"
	ActionApprove      Action = "approve"
	ActionSwap         Action = "swap"
	ActionRoute        Action = "route"
	ActionStake        Action = "stake"
	ActionStakeMax     Action = "stake_max"
	ActionBuy          Action = "buy"
	ActionBid          Action = "bid"
	ActionRenew        Action = "renew"

	ConstraintHealthFactor       Action = "health_factor"
	ConstraintAPY                Action = "apy"
	ConstraintAPYDifferential    Action = "apy_differential"
	ConstraintAvailableLiquidity Action = "available_liquidity"
	HandleConstraintTimeLeft     Action = "time_left"
	HandleConstraintRenewalPrice Action = "renewal_price"
)

type Protocol string

const (
	ProtocolPlug    Protocol = "plug"
	ProtocolAaveV3  Protocol = "aave_v3"
	ProtocolYearnV3 Protocol = "yearn_v3"
	ProtocolNouns   Protocol = "nouns"
	ProtocolENS     Protocol = "ens"
)
