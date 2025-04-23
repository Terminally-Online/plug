package morpho

import (
	"solver/internal/actions"
	morpho_actions "solver/internal/actions/morpho/actions"
	morpho_options "solver/internal/actions/morpho/options"
	"solver/internal/bindings/references"
)

var (
	ActionBorrowSentence            = "Borrow {0<amount:float>} {1<token:address:uint8>} from {1=>2<market:string>}"
	ActionClaimRewardsSentence      = "Claim all reward distributions"
	ActionEarnSentence              = "Earn by depositing {0<amount:float>} {1<token:address:uint8>} to {1=>2<vault:string>}"
	ActionRepaySentence             = "Repay {0<amount:float>} {1<token:address:uint8>} to {1=>2<market:string>}"
	ActionDepositCollateralSentence = "Supply {0<amount:float>} {1<token:address:uint8>} as collateral to {1=>2<market:string>}"
	ActionWithdrawSentence          = "Withdraw {0<amount:float>} {1<token:address:uint8>} from {1=>2<target:string>}"
	ReadAPYSentence                 = "Get {0<action:int8>} APY in {0=>1<target:string>}"
	ReadHealthFactorSentence        = "Get health factor in {0<market:string>}"

	ActionBorrow = actions.NewActionDefinition(
		ActionBorrowSentence,
		morpho_actions.Borrow,
		morpho_options.BorrowTokenToMarketOptions,
		nil,
		&morpho_actions.BorrowFunc,
	)
	ActionClaimRewards = actions.NewActionDefinition(
		ActionClaimRewardsSentence,
		morpho_actions.ClaimRewards,
		nil,
		nil,
		&morpho_actions.ClaimRewardsFunc,
	)
	ActionEarn = actions.NewActionDefinition(
		ActionEarnSentence,
		morpho_actions.Earn,
		morpho_options.SupplyTokenToVaultOptions,
		nil,
		&morpho_actions.EarnFunc,
	)
	ActionRepay = actions.NewActionDefinition(
		ActionRepaySentence,
		morpho_actions.Repay,
		morpho_options.BorrowTokenToMarketOptions,
		nil,
		&morpho_actions.RepayFunc,
	)
	ActionDepositCollateral = actions.NewActionDefinition(
		ActionDepositCollateralSentence,
		morpho_actions.DepositCollateral,
		morpho_options.CollateralTokenToMarketOptions,
		nil,
		&morpho_actions.DepositCollateralFunc,
	)
	ActionWithdraw = actions.NewActionDefinition(
		ActionWithdrawSentence,
		morpho_actions.Withdraw,
		morpho_options.SupplyAndCollateralTokenToMarketOptions,
		nil,
		// TODO MASON: we have two different return function signatures here if we do actually support vaults through this.
		&morpho_actions.WithdrawMarketFunc,
	)
	ReadAPY = actions.NewActionDefinition(
		ReadAPYSentence,
		morpho_actions.APY,
		morpho_options.APYOptions,
		nil,
		&morpho_actions.ApyFunc,
	)
	ReadHealthFactor = actions.NewActionDefinition(
		ReadHealthFactorSentence,
		morpho_actions.HealthFactor,
		morpho_options.HealthFactorOptions,
		nil,
		actions.IsEmptyOnchainFunc,
	)
)

func New() actions.Protocol {
	return actions.NewProtocol(
		actions.Protocol{
			Name:   "Morpho",
			Icon:   "https://cdn.onplug.io/protocols/morpho.png",
			Tags:   []string{"lending", "defi"},
			Chains: []*references.Network{references.Mainnet, references.Base},
			Actions: map[string]actions.ActionDefinitionInterface{
				actions.ActionBorrow:            ActionBorrow,
				actions.ActionClaimRewards:      ActionClaimRewards,
				actions.ActionEarn:              ActionEarn,
				actions.ActionRepay:             ActionRepay,
				actions.ActionDepositCollateral: ActionDepositCollateral,
				actions.ActionWithdraw:          ActionWithdraw,
				actions.ReadAPY:                 ReadAPY,
				actions.ReadHealthFactor:        ReadHealthFactor,
			},
		},
	)
}
