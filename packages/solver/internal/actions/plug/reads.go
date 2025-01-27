package plug

import (
	"context"
	"encoding/hex"
	"fmt"
	"solver/bindings/erc_20"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
TODO: (#12) This does not properly handle all proxies because the code is actually
deployed at a different address than the one that is typically
interfaced with and there is not a universal way to define where the code
actually lives at. We do however already handle contracts that have an implementation().
*/

/*
NOTE: This is a bytecode lookup for function selectors. As the ethereum
standards are poorly written and out of date it is impossible to take
a "standard" mechanism that determines what token type a contract is.

For those reading in the future that think ERC-165:

	(https://github.com/ethereum/ERCs/blob/master/ERCS/erc-165.md)

solves this it unfortunately does not as ERC20 does not implement it,
nor do we really have a way to make sure that all 721 & 1155s do. Really,
it is a pointless standard that has no reason to exist.

Because of this, the simplest way to determine if a token is within standard
is to get the code and then determine if that code contains one of the
function selectors that we *will* use to build the transaction.
The logic of this function is the definition of:

	DUP PUSH (selector) EQ

Which will translate to the bytecode of:

	8063XXXXXXXX14
	8063selector14
*/
func getTokenType(chainId uint64, address string) (*int, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	code, err := provider.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get code at address %s: %v", address, err)
	}
	hexCode := hex.EncodeToString(code)

	if strings.Contains(hexCode, fmt.Sprintf(utils.SelectorLookup, utils.UpgradableImplementationSelector)) {
		return getProxyTokenType(provider, address)
	}

	for _, selector := range utils.TokenSelectors {
		lookup := fmt.Sprintf(utils.SelectorLookup, selector.Selector)
		if strings.Contains(hexCode, lookup) {
			return &selector.Type, nil
		}
	}

	return nil, fmt.Errorf("unsupported token type")
}

func getERC20Decimals(chainId uint64, address string) (*uint8, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	contract, err := erc_20.NewErc20(common.HexToAddress(address), provider)
	if err != nil {
		return nil, fmt.Errorf("failed to create ERC20 contract instance: %v", err)
	}

	decimals, err := contract.Decimals(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get decimals: %v", err)
	}

	return &decimals, nil
}

func getProxyTokenType(_ *ethclient.Client, _ string) (*int, error) {
	return nil, fmt.Errorf("proxies not supported at this time")
}
