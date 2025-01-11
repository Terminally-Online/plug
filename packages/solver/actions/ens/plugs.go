package ens

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"solver/actions"
	"solver/bindings/ens_registrar_controller"
	"solver/types"
	"solver/utils"

	"github.com/ethereum/go-ethereum/common"
)

var (
	resolver = "0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63"
)

func HandleActionBuy(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Name     string `json:"name"`
		MaxPrice string `json:"maxPrice"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ens buy inputs: %v", err)
	}

	name, err := GetName(inputs.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err)
	}

	provider, err := utils.GetProvider(1)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider: %v", err)
	}
	registrar, err := ens_registrar_controller.NewEnsRegistrarController(common.HexToAddress(utils.Mainnet.References["ens"]["registrar_controller"]), provider)
	if err != nil {
		return nil, fmt.Errorf("failed to get registrar: %v", err)
	}
	var secret [32]byte
	nameAndAddr := append([]byte(*name), common.HexToAddress(params.From).Bytes()...)
	copy(secret[:], nameAndAddr)
	commitment, err := registrar.MakeCommitment(
		utils.BuildCallOpts(params.From, big.NewInt(1)),
		*name,
		common.HexToAddress(params.From),
		big.NewInt(secondsPerYear),
		secret,
		common.HexToAddress(resolver),
		[][]byte{},
		true,
		0,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to make commitment: %v", err)
	}

	commitmentTimestamp, err := registrar.Commitments(
		utils.BuildCallOpts(params.From, big.NewInt(1)),
		commitment,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to read commitment: %v", err)
	}

	currentTimestamp := big.NewInt(time.Now().Unix())
	hasCommitment := currentTimestamp.Sub(currentTimestamp, commitmentTimestamp).Cmp(big.NewInt(60)) >= 0 &&
		currentTimestamp.Sub(currentTimestamp, commitmentTimestamp).Cmp(big.NewInt(86400)) <= 0

	registrarAbi, err := ens_registrar_controller.EnsRegistrarControllerMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ENSRegistrarController")
	}

	if !hasCommitment {
		commitCalldata, err := registrarAbi.Pack("commit", commitment)
		if err != nil {
			return nil, utils.ErrTransactionFailed(err.Error())
		}
		return []*types.Transaction{{
			To:        utils.Mainnet.References["ens"]["registrar_controller"],
			Data:      "0x" + common.Bytes2Hex(commitCalldata),
			Exclusive: true,
		}}, nil
	}

	registerCalldata, err := registrarAbi.Pack(
		"register",
		name,
		common.HexToAddress(params.From),
		big.NewInt(secondsPerYear),
		secret,
		common.HexToAddress(resolver),
		[][]byte{},
		true,
		0,
	)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	price, err := GetRentPrice(*name, big.NewInt(secondsPerYear))
	if err != nil {
		return nil, fmt.Errorf("failed to get rent price: %v", err)
	}

	maxPrice, ok := new(big.Int).SetString(inputs.MaxPrice, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse max price: %v", inputs.MaxPrice)
	}
	if price.Base.Cmp(maxPrice) > 0 {
		return nil, fmt.Errorf("rent price (%v wei) is higher than maximum allowed (%v wei)", price.Base, maxPrice)
	}

	return []*types.Transaction{{
		To:    utils.Mainnet.References["ens"]["registrar_controller"],
		Data:  "0x" + common.Bytes2Hex(registerCalldata),
		Value: *price.Base,
	}}, nil
}

func HandleActionRenew(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Name     string   `json:"name"`
		Duration *big.Int `json:"duration"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal renew inputs: %v", err)
	}

	name, err := GetName(inputs.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err)
	}

	price, err := GetRentPrice(*name, inputs.Duration)
	if err != nil {
		return nil, fmt.Errorf("failed to get rent price: %v", err)
	}

	registrarAbi, err := ens_registrar_controller.EnsRegistrarControllerMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABIFailed("ENSRegistrarController")
	}

	renewCalldata, err := registrarAbi.Pack("renew", name, inputs.Duration)
	if err != nil {
		return nil, utils.ErrTransactionFailed(err.Error())
	}

	return []*types.Transaction{{
		To:    utils.Mainnet.References["ens"]["registrar_controller"],
		Data:  "0x" + common.Bytes2Hex(renewCalldata),
		Value: *price.Base,
	}}, nil
}

func HandleConstraintGracePeriod(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ENS time left inputs: %v", err)
	}

	name, err := GetName(inputs.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err)
	}

	expiry, err := GetNameExpiry(*name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name expiry: %v", err)
	}

	gracePeriodEnd := new(big.Int).Add(expiry, big.NewInt(60*60*24*90))
	now := big.NewInt(time.Now().Unix())
	if now.Cmp(expiry) > 0 && now.Cmp(gracePeriodEnd) < 0 {
		return []*types.Transaction{}, nil
	}

	return nil, fmt.Errorf("ENS %s is not in renewal grace period", inputs.Name)
}

func HandleConstraintTimeLeft(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Name     string   `json:"name"`
		Duration *big.Int `json:"duration"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ENS time left inputs: %v", err)
	}

	name, err := GetName(inputs.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err)
	}

	expiry, err := GetNameExpiry(*name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name expiry: %v", err)
	}

	now := big.NewInt(time.Now().Unix())
	remaining := new(big.Int).Sub(expiry, now)

	if remaining.Cmp(inputs.Duration) >= 0 {
		return nil, fmt.Errorf("time remaining (%v seconds) is not less than threshold (%v seconds)", remaining, inputs.Duration)
	}

	return []*types.Transaction{}, nil
}

func HandleConstraintRenewalPrice(rawInputs json.RawMessage, params actions.HandlerParams) ([]*types.Transaction, error) {
	var inputs struct {
		Name     string   `json:"name"`
		Duration *big.Int `json:"duration"`
		MaxPrice *big.Int `json:"price"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal renewal price inputs: %v", err)
	}

	name, err := GetName(inputs.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err)
	}

	price, err := GetRentPrice(*name, inputs.Duration)
	if err != nil {
		return nil, fmt.Errorf("failed to get renewal price: %v", err)
	}

	maxPriceWei := new(big.Int).Mul(inputs.MaxPrice, big.NewInt(1e18))

	if price.Base.Cmp(maxPriceWei) > 0 {
		return nil, fmt.Errorf("renewal price (%v wei) is higher than maximum allowed (%v wei)", price.Base, maxPriceWei)
	}

	return []*types.Transaction{}, nil
}
