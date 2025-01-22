package signature

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	eip712Types = map[string]interface{}{
		"EIP712Domain": []interface{}{
			map[string]interface{}{"name": "name", "type": "string"},
			map[string]interface{}{"name": "version", "type": "string"},
			map[string]interface{}{"name": "chainId", "type": "uint256"},
			map[string]interface{}{"name": "verifyingContract", "type": "address"},
		},
		"Plug": []interface{}{
			map[string]interface{}{"name": "target", "type": "address"},
			map[string]interface{}{"name": "value", "type": "uint256"},
			map[string]interface{}{"name": "data", "type": "bytes"},
		},
		"Plugs": []interface{}{
			map[string]interface{}{"name": "socket", "type": "address"},
			map[string]interface{}{"name": "plugs", "type": "Plug[]"},
			map[string]interface{}{"name": "solver", "type": "bytes"},
			map[string]interface{}{"name": "salt", "type": "bytes"},
		},
		"LivePlugs": []interface{}{
			map[string]interface{}{"name": "plugs", "type": "Plugs"},
			map[string]interface{}{"name": "signature", "type": "bytes"},
		},
	}
)

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
