package intent

import (
	"solver/utils"
)

type HarvestInputs struct {
	Token          string `json:"token"`          // Address of the token to harvest.
	PrimaryAddress string `json:"primaryAddress"` // Address of the smart contract to interact with.
}

func (i HarvestInputs) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	if !utils.IsAddress(i.PrimaryAddress) {
		return utils.ErrInvalidHex("primaryAddress", i.PrimaryAddress)
	}
	return nil
}

func (i HarvestInputs) Build(from string) (*utils.Transaction, error) {
	if err := i.Validate(); err != nil {
		return nil, err
	}

	return nil, utils.ErrNotImplemented("HarvestInputs.Build")
}
