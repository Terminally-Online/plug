package actions

import (
	"fmt"
	"solver/bindings/yearn_v3_pool"
	"solver/internal/actions"
	"solver/internal/actions/yearn_v3/reads"
	"solver/internal/actions/yearn_v3/types"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawRequest struct {
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Vault  string `json:"vault"`
}

var WithdrawFunc = actions.ActionOnchainFunctionResponse{
	Metadata:     yearn_v3_pool.YearnV3PoolMetaData,
	FunctionName: "withdraw",
}

func Withdraw(lookup *actions.SchemaLookup[WithdrawRequest]) ([]signature.Plug, error) {
	token, decimals, err := utils.ParseAddressAndDecimals(lookup.Inputs.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token with decimals: %w", err)
	}

	amount, err := utils.StringToUint(lookup.Inputs.Amount, decimals)
	if err != nil {
		return nil, fmt.Errorf("failed to convert deposit amount to uint: %w", err)
	}

	vaults, err := reads.GetVaults(lookup.ChainId)
	if err != nil {
		return nil, fmt.Errorf("failed to get vaults: %v", err)
	}
	var targetVault *types.YearnVault
	for _, vault := range vaults {
		if strings.EqualFold(vault.Address, lookup.Inputs.Vault) {
			targetVault = &vault
			break
		}
	}
	if targetVault == nil {
		return nil, fmt.Errorf("deposit not available for vault: %s", lookup.Inputs.Vault)
	}
	if !strings.EqualFold(token.Hex(), targetVault.Token.Address) {
		return nil, fmt.Errorf("asset %s cannot be used in vault: %s", token, lookup.Inputs.Vault)
	}

	withdrawCalldata, err := WithdrawFunc.GetCalldata(amount, lookup.From)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:   common.HexToAddress(lookup.Inputs.Vault),
		Data: withdrawCalldata,
	}}, nil
}
