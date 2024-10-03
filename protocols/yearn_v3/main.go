package yearn_v3

import (
	"math/big"
	"os"
	"solver/bindings/erc_20"
	"solver/bindings/yearn_v3_pool"
	"solver/bindings/yearn_v3_registry"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**
Documentation for Yearn V3:
https://docs.yearn.finance/developers/addresses/v3-contracts

Documentation for the Registry:
https://docs.yearn.fi/developers/smart-contracts/V3/periphery/Registry#getendorsedvaults

Specification: ERC4626 Vaults
*/

var (
	Key = "yearn_v3"

	registryAddress    = "0xff31A1B020c868F6eA3f61Eb953344920EeCA3af"
	registryHexAddress = common.HexToAddress(registryAddress)
)

func ValidateVault(provider *ethclient.Client, asset string, target string) (*common.Address, error) {
	registry, err := yearn_v3_registry.NewYearnV3Registry(registryHexAddress, provider)
	if err != nil {
		return nil, utils.ErrContractFailed(registryAddress)
	}
	vaults, err := registry.GetEndorsedVaults(
		utils.BuildCallOpts(os.Getenv("SOLVER_ADDRESS"), big.NewInt(0)),
		common.HexToAddress(asset),
	)
	if err != nil {
		return nil, utils.ErrCallFailed(err.Error())
	}

	// Confirm the returned list of addresses (`vaults`) contains the target address.
	for _, vault := range vaults {
		if vault.String() == target {
			return &vault, nil
		}
	}

	return nil, utils.ErrInvalidField("target", target)
}

/*
*
BuildDeposit is the primary function for building a deposit transaction for Yearn V3.

It includes top-level validation for the token is supported by a yearn v3 vault.
Deposit, requires an approval first because the vaults pull from the sender directly.
By default this means that we must first approve the vault, and then deposit. We cannot
run these actions in separate transactions because the tokens could be stolen. This
must be done in a multicall.

If type(uint256).max is used as the amount, the vault will use the entire allowed balance
between the token balance of the sender and the max allowed deposit amount of the vault.
*/
func BuildDeposit(i types.DepositInputs, provider *ethclient.Client, chainId int, from string) ([]*ethtypes.Transaction, error) {
	target, err := ValidateVault(provider, i.GetTokenIn(), i.GetTokenOut())
	if err != nil {
		return nil, err
	}

	token, err := erc_20.NewErc20(common.HexToAddress(i.GetTokenIn()), provider)
	if err != nil {
		return nil, utils.ErrContractFailed(i.GetTokenIn())
	}

	amountIn := i.GetAmountIn()

	// When the amount is set to the uint256 max we will let the system determine the amount
	// amount to deposit based on the users desire to deposit the "maximum". The maximum
	// amount is calculated by the balance of the token held by the user and the maximum
	// deposit that can be made into the vault right now.
	if amountIn.Cmp(utils.Uint256Max) == 0 {
		balance, err := token.BalanceOf(
			utils.BuildCallOpts(from, big.NewInt(0)),
			common.HexToAddress(from),
		)
		if err != nil {
			return nil, utils.ErrCallFailed(err.Error())
		}

		pool, err := yearn_v3_pool.NewYearnV3Pool(*target, provider)
		if err != nil {
			return nil, utils.ErrContractFailed(i.GetTokenIn())
		}
		maxDeposit, err := pool.MaxDeposit(
			utils.BuildCallOpts(from, big.NewInt(0)),
			common.HexToAddress(from),
		)
		if err != nil {
			return nil, utils.ErrCallFailed(err.Error())
		}

		if balance.Cmp(maxDeposit) < 0 {
			amountIn = balance
		} else {
			amountIn = maxDeposit
		}

		if amountIn.Cmp(big.NewInt(0)) == 0 {
			return nil, utils.ErrInvalidField("amount: max", amountIn.String())
		}
	}

	// Approve function is built to allow token transfers to the vault by the vault.
	approve, err := token.Approve(
		utils.BuildTransactionOpts(from, big.NewInt(0)),
		*target,
		amountIn,
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	// Build the deposit calldata for the vault that has been provided.
	poolAbi, err := yearn_v3_pool.YearnV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("YearnV3Pool")
	}
	calldata, err := poolAbi.Pack("deposit", amountIn, common.HexToAddress(from))
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}
	deposit := ethtypes.NewTx(&ethtypes.LegacyTx{
		To:   target,
		Data: calldata,
	})

	return []*ethtypes.Transaction{approve, deposit}, nil
}

/*
*
BuildRedeem is the primary function for building a redemption for a Yearn V3 vault.

It includes top-level validation for the token is supported by a yearn v3 vault.
We cannot run these actions in separate transactions because the tokens could be stolen.
This must be done in a multicall.

If type(uint256).max is used as the amount, the vault will use determine the amount of shares
to redeem based on the users desire to redeem the "maximum". The maximum amount is calculated
by the contract and utilized through a read call to the vault.
*/
func BuildRedeem(i types.RedeemInputs, provider *ethclient.Client, chainId int, from string) ([]*ethtypes.Transaction, error) {
	// Validate that providing the pool token as the target is valid to receive the intended token out.
	target, err := ValidateVault(provider, i.GetTokenOut(), i.GetTokenIn())
	if err != nil {
		return nil, err
	}

	amountIn := i.GetAmountIn()
	if amountIn.Cmp(utils.Uint256Max) == 0 {
		pool, err := yearn_v3_pool.NewYearnV3Pool(*target, provider)
		if err != nil {
			return nil, utils.ErrContractFailed(target.String())
		}
		amountIn, err = pool.MaxRedeem(
			utils.BuildCallOpts(from, big.NewInt(0)),
			common.HexToAddress(from),
		)
		if err != nil {
			return nil, utils.ErrCallFailed(err.Error())
		}

		if amountIn.Cmp(big.NewInt(0)) == 0 {
			return nil, utils.ErrInvalidField("amountIn: max", amountIn.String())
		}
	}

	poolAbi, err := yearn_v3_pool.YearnV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("YearnV3Pool")
	}
	calldata, err := poolAbi.Pack("redeem", amountIn, common.HexToAddress(from), common.HexToAddress(from))
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}
	redeem := ethtypes.NewTx(&ethtypes.LegacyTx{
		To:   target,
		Data: calldata,
	})

	return []*ethtypes.Transaction{redeem}, nil
}
