package actions

import (
	"math/big"
	"solver/bindings/plug_factory"
	"solver/internal/actions"
	"solver/internal/solver/signature"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type DeployRequest struct {
	Factory        common.Address `json:"factory"`
	Nonce          *big.Int       `json:"nonce"`
	Admin          common.Address `json:"admin"`
	Delegate       common.Address `json:"delegate"`
	Implementation common.Address `json:"implementation"`
}

func Deploy(lookup *actions.SchemaLookup[DeployRequest]) ([]signature.Plug, error) {
	// NOTE: Instead of using the inline deployment that is enabled by the Plug router we
	//    	 call the factory directly because going through the router requires a signature
	//    	 and this action should be able to be executed without a user signature.
	//    	 -- The protocol was designed with a path explicitly for this, so it is chill.
	plugFactoryAbi, err := plug_factory.PlugFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	salt, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 96}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.AddressTy}},
	}.Pack(
		lookup.Inputs.Nonce,
		lookup.Inputs.Admin,
		lookup.Inputs.Delegate,
		lookup.Inputs.Implementation,
	)
	if err != nil {
		return nil, err
	}

	deployCalldata, err := plugFactoryAbi.Pack("deploy", salt)
	if err != nil {
		return nil, err
	}

	return []signature.Plug{{
		To:   lookup.Inputs.Factory,
		Data: deployCalldata,
	}}, nil
}
