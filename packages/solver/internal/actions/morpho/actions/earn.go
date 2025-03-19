package actions

import (
	"fmt"
	"solver/bindings/erc_20"
	"solver/bindings/morpho_vault"
	"solver/internal/actions"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

type EarnRequest struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Vault  string `json:"vault"`
}

func Earn(lookup *actions.SchemaLookup[EarnRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert earn amount to uint: %w", err)
	}

	erc20Abi, err := erc_20.Erc20MetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ERC20")
	}
	approveCalldata, err := erc20Abi.Pack(
		"approve",
		common.HexToAddress(lookup.Inputs.Vault),
		amount,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack approve calldata: %w", err)
	}

	vaultAbi, err := morpho_vault.MorphoVaultMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get morpho vault abi: %w", err)
	}
	depositCalldata, err := vaultAbi.Pack(
		"deposit",
		amount,
		lookup.From,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to pack deposit calldata: %w", err)
	}

	return []signature.Plug{{
		To:   *token,
		Data: approveCalldata,
	}, {
		To:   common.HexToAddress(lookup.Inputs.Vault),
		Data: depositCalldata,
	}}, nil
}
