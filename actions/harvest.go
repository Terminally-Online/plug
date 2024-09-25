package actions

import (
	"solver/types"
	"solver/utils"
)

type HarvestInputsImpl struct {
	Protocol string `json:"protocol"` // Slug of the protocol to use.
	Token    string `json:"token"`    // Address of the token to harvest.
}

func (i *HarvestInputsImpl) Validate() error {
	if !utils.IsAddress(i.Token) {
		return utils.ErrInvalidAddress("token", i.Token)
	}
	return nil
}

func (i *HarvestInputsImpl) Build(chainId int, from string) (*types.Transaction, error) {
	switch i.Protocol {
	default:
		return nil, utils.ErrInvalidProtocol("protocol", i.Protocol)
	}
}

func (i *HarvestInputsImpl) GetProtocol() string { return i.Protocol }
func (i *HarvestInputsImpl) GetToken() string    { return i.Token }
