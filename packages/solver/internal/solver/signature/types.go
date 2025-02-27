package signature

import (
	"math/big"
	"solver/bindings/plug_router"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EIP712_DOMAIN_TYPEHASH = "0x8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f"
	PLUG_TYPEHASH          = "0x0d73e94823fdaacb148d9146f00bc268b7834e768ced483d796db05a52e1e395"
	PLUGS_TYPEHASH         = "0x4ddfe68cf187b28815da9c19a2cb9477b3f3293c6170c7e3e56842b550ac141d"
	LIVE_PLUGS_TYPEHASH    = "0x049e34029d287aa78a0f5a45ebdf78081b3357f85f5ab4cfcffb786eff0b3375"
)

type Transaction struct {
	From  common.Address `json:"from"`
	To    common.Address `json:"to"`
	Data  []byte         `json:"data"`
	Value *big.Int       `json:"value"`
	Gas   *big.Int       `json:"gas"`
}

type EIP712Domain struct {
	Name              string         `json:"name"`
	Version           string         `json:"version"`
	ChainId           *big.Int       `json:"chainId"`
	VerifyingContract common.Address `json:"verifyingContract"`
}

type Plug struct {
	To        common.Address `json:"to"`
	Data      []byte         `json:"data"`
	Value     *big.Int       `json:"value"`
	Gas       *big.Int       `json:"gas"`
	Exclusive bool           `json:"exclusive,omitempty"`
	Meta      interface{}    `json:"meta,omitempty"`
}

func (p Plug) Wrap() plug_router.PlugTypesLibPlug {
	return plug_router.PlugTypesLibPlug{
		To:    p.To,
		Data:  p.Data,
		Value: p.Value,
		Gas:   p.Gas,
	}
}

type Plugs struct {
	Socket common.Address `json:"socket"`
	Plugs  []Plug         `json:"plugs"`
	Solver []byte         `json:"solver"`
	Salt   []byte         `json:"salt"`
}

func (p Plugs) Wrap() plug_router.PlugTypesLibPlugs {
	var plugs []plug_router.PlugTypesLibPlug
	for _, plug := range p.Plugs {
		plugs = append(plugs, plug.Wrap())
	}

	return plug_router.PlugTypesLibPlugs{
		Socket: p.Socket,
		Plugs:  plugs,
		Solver: p.Solver,
		Salt:   p.Salt,
	}
}

type LivePlugs struct {
	Plugs     Plugs  `json:"plugs"`
	Signature []byte `json:"signature"`
}

func (l LivePlugs) Wrap() plug_router.PlugTypesLibLivePlugs {
	return plug_router.PlugTypesLibLivePlugs{
		Plugs:     l.Plugs.Wrap(),
		Signature: l.Signature,
	}
}

type Result struct {
	Success bool   `json:"success"`
	Result  []byte `json:"result"`
}
