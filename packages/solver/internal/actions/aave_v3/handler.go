package aave_v3

import (
	"math/big"
	"solver/internal/actions"
	"solver/internal/bindings/references"
)

var (
	interestRateMode = new(big.Int).SetUint64(2)
)

func New() actions.Protocol {
	return actions.New(actions.Protocol{
		Name:   "Aave V3",
		Icon:   "https://cdn.onplug.io/protocols/aave.png",
		Tags:   []string{"lending", "defi"},
		Chains: []*references.Network{references.Mainnet, references.Base},
		Actions: map[string]actions.ActionDefinition{
			actions.ActionDeposit: {
				Sentence: "Deposit {0<amount:float>} {1<token:address:uint8>}",
				Handler:  Deposit,
				Options:  CollateralOptions,
			},
			actions.ActionBorrow: {
				Sentence: "Borrow {0<amount:float>} {1<token:address:uint8>}",
				Handler:  Borrow,
				Options:  BorrowOptions,
			},
			actions.ActionRepay: {
				Sentence: "Repay {0<amount:float>} {1<token:address:uint8>}",
				Handler:  Repay,
				Options:  BorrowOptions,
			},
			actions.ActionWithdraw: {
				Sentence: "Withdraw {0<amount:float>} {1<token:address:uint8>}",
				Handler:  Withdraw,
				Options:  CollateralOptions,
			},
			actions.ConstraintHealthFactor: {
				Type:     actions.TypeConstraint,
				Sentence: "Health factor is {0<operator:int8>} than {1<threshold:float>}",
				Handler:  HealthFactor,
				Options:  HealthFactorOptions,
			},
			actions.ConstraintAPY: {
				Type:     actions.TypeConstraint,
				Sentence: "{0<action:int8>} APY of {1<token:address:uint8>} is {2<operator:int8>} than {3<threshold:float>}%",
				Handler:  APY,
				Options:  APYOptions,
			},
		},
	})
}
