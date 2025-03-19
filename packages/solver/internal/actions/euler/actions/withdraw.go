package actions

import (
	"fmt"
	"math/big"
	"solver/bindings/euler_evault_implementation"
	"solver/internal/actions"
	"solver/internal/actions/euler/reads"
	euler_utils "solver/internal/actions/euler/utils"
	"solver/internal/solver/signature"
	"solver/internal/utils"
)

type WithdrawCollateralRequest struct {
	Amount          string `json:"amount"`
	Token           string `json:"token"`
	Vault           string `json:"vault"`
	SubAccountIndex uint8  `json:"sub-account"`
}

type WithdrawRequest struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Vault  string `json:"vault"`
}

func HandleWithdrawCollateral(lookup *actions.SchemaLookup[WithdrawCollateralRequest]) ([]signature.Plug, error) {
	_, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert withdraw amount to uint: %w", err)
	}

	vault, err := reads.GetVault(lookup.Inputs.Vault, lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vault: %w", err)
	}

	vaultAbi, err := euler_evault_implementation.EulerEvaultImplementationMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("EulerEvaultImplementation")
	}

	subAccountAddress := euler_utils.GetSubAccountAddress(lookup.From, lookup.Inputs.SubAccountIndex)

	calldata, err := vaultAbi.Pack(
		"withdraw",
		amount,
		lookup.From,
		subAccountAddress,
	)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	call, err := euler_utils.WrapEVCCall(
		lookup.ChainId,
		vault.Vault,
		subAccountAddress,
		big.NewInt(0),
		calldata,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to wrap withdraw collateral call: %w", err)
	}

	return []signature.Plug{call}, nil
}

// NOTE: I believe this is exactly the same as WithdrawCollateral, but I don't
// want to rope it in with the other handler with a default sub account
// index of 0 in case there is some difference I'm not aware of yet. On
// the options side, this one does not have a sub account index option to
// remove complexity around earn deposits not being indexed in a contract like
// borrow and lending positions.
// - Mason
func HandleWithdraw(lookup *actions.SchemaLookup[WithdrawRequest]) ([]signature.Plug, error) {
	collateralLookup := &actions.SchemaLookup[WithdrawCollateralRequest]{
		ChainId: lookup.ChainId,
		Client:  lookup.Client,
		From:    lookup.From,
		Search:  lookup.Search,
		Inputs: &WithdrawCollateralRequest{
			Amount:          lookup.Inputs.Amount,
			Token:           lookup.Inputs.Token,
			Vault:           lookup.Inputs.Vault,
			SubAccountIndex: 0,
		},
	}

	return HandleWithdrawCollateral(collateralLookup)
}
