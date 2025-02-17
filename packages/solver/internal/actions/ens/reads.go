package ens

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"regexp"
	"solver/bindings/ens_base_registrar"
	"solver/bindings/ens_registrar_controller"
	"solver/internal/bindings/references"
	"solver/internal/client"
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

func GetRentPrice(chainId uint64, name string, duration *big.Int) (*ens_registrar_controller.IPriceOraclePrice, error) {
	client, err := client.New(chainId)
	if err != nil {
		return nil, err
	}

	// TODO: What is this address doing hardcoded here? ENS integration has so many issues
	//       in it right now that need to be handled before pushing it into production.
	registrar, err := ens_registrar_controller.NewEnsRegistrarController(common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b"), client)
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
func GetNameExpiry(chainId uint64, name string) (*big.Int, error) {
	client, err := client.New(chainId)
	if err != nil {
		return nil, err
	}

	// Calculate the label hash (uint256) for the name
	nameHash := sha256.Sum256([]byte(name))
	labelHash := new(big.Int).SetBytes(nameHash[:])

	// Create contract instance for BaseRegistrar
	baseRegistrar, err := ens_base_registrar.NewEnsBaseRegistrar( // Update constructor name
		common.HexToAddress(references.Mainnet.References["ens"]["base_registrar"]),
		client,
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
