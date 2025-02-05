package signature

import (
	"math/big"

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

type Plugs struct {
	Socket common.Address `json:"socket"`
	Plugs  []Plug         `json:"plugs"`
	Solver []byte         `json:"solver"`
	Salt   []byte         `json:"salt"`
}

type LivePlugs struct {
	Plugs     Plugs  `json:"plugs"`
	Signature []byte `json:"signature"`
}

type Result struct {
	Success bool   `json:"success"`
	Result  []byte `json:"result"`
}
