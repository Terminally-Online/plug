package aave_v3

import (
	"math/big"
	"solver/internal/actions"
	aave_actions "solver/internal/actions/aave_v3/actions"
	aave_options "solver/internal/actions/aave_v3/options"
	"solver/internal/bindings/references"
)

var (
	InterestRateMode = new(big.Int).SetUint64(2)

	ActionBorrowSentence     = "Borrow {0<amount:float>} {1<token:address:uint8>}"
	ActionDepositSentence    = "Deposit {0<amount:float>} {1<token:address:uint8>}"
	ActionRepaySentence      = "Repay {0<amount:float>} {1<token:address:uint8>}"
	ActionWithdrawSentence   = "Withdraw {0<amount:float>} {1<token:address:uint8>}"
	ReadApySentence          = "{0<action:int8>} APY of {1<token:address:uint8>} is {2<operator:int8>} than {3<threshold:float>}%"
	ReadHealthFactorSentence = "Get health factor"

	ActionBorrow = actions.NewActionDefinition(
		ActionBorrowSentence,
		aave_actions.Borrow,
		aave_options.BorrowOptions,
		nil,
		&aave_actions.BorrowFunc,
	)
	ActionDeposit = actions.NewActionDefinition(
		ActionDepositSentence,
		aave_actions.Deposit,
		aave_options.CollateralOptions,
		nil,
		&aave_actions.DepositFunc,
	)
	ActionRepay = actions.NewActionDefinition(
		ActionRepaySentence,
		aave_actions.Repay,
		aave_options.BorrowOptions,
		nil,
		&aave_actions.RepayFunc,
	)
	ActionWithdraw = actions.NewActionDefinition(
		ActionWithdrawSentence,
		aave_actions.Withdraw,
		aave_options.CollateralOptions,
		nil,
		&aave_actions.WithdrawFunc,
	)
	ReadApy = actions.NewActionDefinition(
		ReadApySentence,
		aave_actions.APY,
		aave_options.APYOptions,
		nil,
		actions.IsEmptyOnchainFunc,
	)
	ReadHealthFactor = actions.NewActionDefinition(
		ReadHealthFactorSentence,
		aave_actions.HealthFactor,
		aave_options.HealthFactorOptions,
		nil,
		actions.IsEmptyOnchainFunc,
	)
)

func New() actions.Protocol {
	return actions.NewProtocol(actions.Protocol{
		Name:   "Aave V3",
		Icon:   "https://cdn.onplug.io/protocols/aave.png",
		Tags:   []string{"lending", "defi"},
		Chains: []*references.Network{references.Mainnet, references.Base},
		Actions: map[string]actions.ActionDefinitionInterface{
			actions.ActionBorrow:     ActionBorrow,
			actions.ActionDeposit:    ActionDeposit,
			actions.ActionRepay:      ActionRepay,
			actions.ActionWithdraw:   ActionWithdraw,
			actions.ReadAPY:          ReadApy,
			actions.ReadHealthFactor: ReadHealthFactor,
		},
	})
}
