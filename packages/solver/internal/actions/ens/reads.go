package ens

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"regexp"
	"solver/bindings/ens_base_registrar"
	"solver/bindings/ens_registrar_controller"
	"solver/internal/bindings/references"
	"solver/internal/utils"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func GetName(name string) (*string, error) {
	name = strings.TrimSuffix(name, ".eth")
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(name) {
		return nil, fmt.Errorf("name must only contain letters and numbers: %s", name)
	}
	return &name, nil
}

func GetRentPrice(chainId int, name string, duration *big.Int) (*ens_registrar_controller.IPriceOraclePrice, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	registrar, err := ens_registrar_controller.NewEnsRegistrarController(common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b"), provider)
	if err != nil {
		return nil, err
	}

	price, err := registrar.RentPrice(nil, name, duration)
	if err != nil {
		return nil, err
	}

	return &price, nil
}

// GetNameExpiry returns the unix timestamp at which an ENS name registration expires
func GetNameExpiry(chainId int, name string) (*big.Int, error) {
	provider, err := utils.GetProvider(chainId)
	if err != nil {
		return nil, err
	}

	// Calculate the label hash (uint256) for the name
	nameHash := sha256.Sum256([]byte(name))
	labelHash := new(big.Int).SetBytes(nameHash[:])

	// Create contract instance for BaseRegistrar
	baseRegistrar, err := ens_base_registrar.NewEnsBaseRegistrar( // Update constructor name
		common.HexToAddress(references.Mainnet.References["ens"]["base_registrar"]),
		provider,
	)
	if err != nil {
		return nil, err
	}

	// Call nameExpires function
	expiry, err := baseRegistrar.NameExpires(nil, labelHash)
	if err != nil {
		return nil, err
	}

	return expiry, nil
}
