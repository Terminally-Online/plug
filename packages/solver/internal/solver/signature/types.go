package signature

import (
	"math/big"
	"solver/bindings/plug_router"

	"github.com/ethereum/go-ethereum/common"
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

type Slice struct {
	Index  uint8    `json:"index"`  // Index of the plug in the sequence that produces this data
	Start  *big.Int `json:"start"`  // Starting byte position within the result data
	Length *big.Int `json:"length"` // Length of bytes to extract
}

func (s Slice) Wrap() plug_router.PlugTypesLibSlice {
	return plug_router.PlugTypesLibSlice{
		Index:  s.Index,
		Start:  s.Start,
		Length: s.Length,
	}
}

type Update struct {
	Start *big.Int `json:"start"` // Starting position where the slice should be inserted
	Slice Slice    `json:"slice"` // The slice specification
}

func (p Update) Wrap() plug_router.PlugTypesLibUpdate {
	return plug_router.PlugTypesLibUpdate{
		Start: p.Start,
		Slice: p.Slice.Wrap(),
	}
}

type Plug struct {
	Selector uint8          `json:"selector"` // 0x00 for call, 0x01 for delegatecall
	To       common.Address `json:"to"`       // Target contract address
	Data     []byte         `json:"data"`     // Calldata for the interaction
	Value    *big.Int       `json:"value"`    // ETH value to send with the call
	Updates  []Update       `json:"updates"`  // List of updates to apply to the data

	Exclusive bool        `json:"exclusive,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
}

func (p Plug) Wrap() plug_router.PlugTypesLibPlug {
	updates := make([]plug_router.PlugTypesLibUpdate, len(p.Updates))
	for index, update := range p.Updates {
		updates[index] = update.Wrap()
	}

	return plug_router.PlugTypesLibPlug{
		Selector: p.Selector,
		To:       p.To,
		Data:     p.Data,
		Value:    p.Value,
		Updates:  updates,
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
