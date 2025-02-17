package ens

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"solver/bindings/ens_registrar_controller"
	"solver/internal/actions"
	"solver/internal/bindings/references"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"github.com/ethereum/go-ethereum/common"
)

var (
	resolver = "0x231b0Ee14048e9dCcD1d247744d114a4EB5E8E63"
)

func HandleActionBuy(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Name     string `json:"name"`
		MaxPrice string `json:"maxPrice"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal ens buy inputs: %v", err)
	}

	maxPrice, err := utils.StringToUint(inputs.MaxPrice, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert ens max price to uint: %v", err)
	}

	name, err := GetName(inputs.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err)
	}

	registrar, err := ens_registrar_controller.NewEnsRegistrarController(
		common.HexToAddress(references.Mainnet.References["ens"]["registrar_controller"]),
		params.Client,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get registrar: %v", err)
	}
	var secret [32]byte
	nameAndAddr := append([]byte(*name), common.HexToAddress(params.From).Bytes()...)
	copy(secret[:], nameAndAddr)
	commitment, err := registrar.MakeCommitment(
		params.Client.ReadOptions(params.From),
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
		params.Client.ReadOptions(params.From),
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
		return nil, utils.ErrABI("ENSRegistrarController")
	}

	if !hasCommitment {
		commitCalldata, err := registrarAbi.Pack("commit", commitment)
		if err != nil {
			return nil, utils.ErrTransaction(err.Error())
		}
		return []signature.Plug{{
			To:        common.HexToAddress(references.Mainnet.References["ens"]["registrar_controller"]),
			Data:      commitCalldata,
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
		return nil, utils.ErrTransaction(err.Error())
	}

	price, err := GetRentPrice(params.ChainId, *name, big.NewInt(secondsPerYear))
	if err != nil {
		return nil, fmt.Errorf("failed to get rent price: %v", err)
	}

	if price.Base.Cmp(maxPrice) > 0 {
		return nil, fmt.Errorf("rent price (%v wei) is higher than maximum allowed (%v wei)", price.Base, maxPrice)
	}

	return []signature.Plug{{
		To:    common.HexToAddress(references.Mainnet.References["ens"]["registrar_controller"]),
		Data:  registerCalldata,
		Value: price.Base,
	}}, nil
}

func HandleActionRenew(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
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

	price, err := GetRentPrice(params.ChainId, *name, inputs.Duration)
	if err != nil {
		return nil, fmt.Errorf("failed to get rent price: %v", err)
	}

	registrarAbi, err := ens_registrar_controller.EnsRegistrarControllerMetaData.GetAbi()
	if err != nil {
		return nil, utils.ErrABI("ENSRegistrarController")
	}

	renewCalldata, err := registrarAbi.Pack("renew", name, inputs.Duration)
	if err != nil {
		return nil, utils.ErrTransaction(err.Error())
	}

	return []signature.Plug{{
		To:    common.HexToAddress(references.Mainnet.References["ens"]["registrar_controller"]),
		Data:  renewCalldata,
		Value: price.Base,
	}}, nil
}

func HandleConstraintGracePeriod(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
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

	expiry, err := GetNameExpiry(params.ChainId, *name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name expiry: %v", err)
	}

	gracePeriodEnd := new(big.Int).Add(expiry, big.NewInt(60*60*24*90))
	now := big.NewInt(time.Now().Unix())
	if now.Cmp(expiry) > 0 && now.Cmp(gracePeriodEnd) < 0 {
		return nil, nil
	}

	return nil, fmt.Errorf("ENS %s is not in renewal grace period", inputs.Name)
}

func HandleConstraintTimeLeft(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
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

	expiry, err := GetNameExpiry(params.ChainId, *name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name expiry: %v", err)
	}

	now := big.NewInt(time.Now().Unix())
	remaining := new(big.Int).Sub(expiry, now)

	if remaining.Cmp(inputs.Duration) >= 0 {
		return nil, fmt.Errorf("time remaining (%v seconds) is not less than threshold (%v seconds)", remaining, inputs.Duration)
	}

	return nil, nil
}

func HandleConstraintRenewalPrice(rawInputs json.RawMessage, params actions.HandlerParams) ([]signature.Plug, error) {
	var inputs struct {
		Name     string   `json:"name"`
		Duration *big.Int `json:"duration"`
		MaxPrice string   `json:"price"`
	}
	if err := json.Unmarshal(rawInputs, &inputs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal renewal price inputs: %v", err)
	}

	maxPrice, err := utils.StringToUint(inputs.MaxPrice, 18)
	if err != nil {
		return nil, fmt.Errorf("failed to convert ens max price to uint: %v", err)
	}

	name, err := GetName(inputs.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to get name: %v", err)
	}

	price, err := GetRentPrice(params.ChainId, *name, inputs.Duration)
	if err != nil {
		return nil, fmt.Errorf("failed to get renewal price: %v", err)
	}

	if price.Base.Cmp(maxPrice) > 0 {
		return nil, fmt.Errorf("renewal price (%v wei) is higher than maximum allowed (%v wei)", price.Base, maxPrice)
	}

	return nil, nil
}
