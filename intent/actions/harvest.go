package intent

import (
	"solver/utils"
)

type HarvestInputs struct {
	Token          string `json:"token"`
	PrimaryAddress string `json:"primaryAddress"`
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
