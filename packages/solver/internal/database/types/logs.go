package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Log struct {
	Address common.Address `json:"address"`
	Topics  []common.Hash  `json:"topics"`
	Data    hexutil.Bytes  `json:"data"`
}

type EventParameter struct {
	Name    string      `json:"name"`
	Type    string      `json:"type"`
	Indexed bool        `json:"-"`
	Value   interface{} `json:"value"`
}

type DecodedLog struct {
	Address    common.Address   `json:"address"`
	Name       string           `json:"name"`
	Parameters []EventParameter `json:"parameters"`
	Raw        Log              `json:"raw"`
}

type DecodedLogs []DecodedLog

func (l DecodedLog) MarshalJSON() ([]byte, error) {
	type Alias DecodedLog
	return json.Marshal(struct {
		Address string `json:"address"`
		*Alias
	}{
		Address: l.Address.Hex(),
		Alias:   (*Alias)(&l),
	})
}

func (l Log) MarshalJSON() ([]byte, error) {
	topics := make([]string, len(l.Topics))
	for i, t := range l.Topics {
		topics[i] = t.Hex()
	}
	return json.Marshal(struct {
		Address string   `json:"address"`
		Topics  []string `json:"topics"`
		Data    string   `json:"data"`
	}{
		Address: l.Address.Hex(),
		Topics:  topics,
		Data:    l.Data.String(),
	})
}

func (p EventParameter) MarshalJSON() ([]byte, error) {
	value := p.Value

	switch v := p.Value.(type) {
	case []byte:
		value = hexutil.Encode(v)
	case [32]byte:
		value = hexutil.Encode(v[:])
	case common.Address:
		value = v.String()
	case common.Hash:
		if p.Type == "address" {
			value = common.BytesToAddress(v.Bytes()).String()
		} else {
			value = v.Hex()
		}
	}

	return json.Marshal(struct {
		Name    string      `json:"name"`
		Type    string      `json:"type"`
		Indexed bool        `json:"indexed"`
		Value   interface{} `json:"value"`
	}{
		Name:    p.Name,
		Type:    p.Type,
		Indexed: p.Indexed,
		Value:   value,
	})
}
