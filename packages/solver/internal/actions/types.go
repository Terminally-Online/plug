package actions

import "solver/bindings/erc_20"

const (
	TypeAction     string = "action"
	TypeConstraint string = "constraint"

	AaveV3    string = "aave_v3"
	Assert    string = "assert"
	BasePaint string = "basepaint"
	Bebop     string = "bebop"
	Boolean   string = "boolean"
	Database  string = "database"
	ENS       string = "ens"
	Euler     string = "euler"
	Math      string = "math"
	Morpho    string = "morpho"
	Nouns     string = "nouns"
	Plug      string = "plug"
	Scripting string = "scripting"
	YearnV3   string = "yearn_v3"

	ActionEarn               string = "earn"
	ActionDeposit            string = "deposit"
	ActionDepositCollateral  string = "deposit_collateral"
	ActionBorrow             string = "borrow"
	ActionRedeem             string = "redeem"
	ActionWithdraw           string = "withdraw"
	ActionWithdrawCollateral string = "withdraw_collateral"
	ActionRepay              string = "repay"
	ActionHarvest            string = "harvest"
	ActionTransfer           string = "transfer"
	ActionTransferFrom       string = "transfer_from"
	ActionApprove            string = "approve"
	ActionSwap               string = "swap"
	ActionRoute              string = "route"
	ActionStake              string = "stake"
	ActionUnstake            string = "unstake"
	ActionBuy                string = "buy"
	ActionBid                string = "bid"
	ActionRenew              string = "renew"
	ActionMint               string = "mint"
	ActionDeploy             string = "deploy"
	ActionClaimRewards       string = "claim_rewards"

	ReadBalance            string = "balance"
	ReadPrice              string = "price"
	ReadAPY                string = "apy"
	ReadHealthFactor       string = "health_factor"
	ReadTimeToLiquidiation string = "time_to_liquidation"
	ReadCurrentAuction     string = "current_auction"
	ReadHasTrait           string = "has_trait"

	IsGlobal  = false
	IsUser    = true
	IsStatic  = false
	IsDynamic = true
)

var (
	IsEmptyOnchainFunc *ActionOnchainFunctionResponse = nil

	Erc20ApprovalFunc = ActionOnchainFunctionResponse{
		Metadata:     erc_20.Erc20MetaData,
		FunctionName: "approve",
	}
)
