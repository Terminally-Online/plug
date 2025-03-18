package types

import (
	"encoding/json"
	"math/big"
)

type Markets struct {
	Items    []Market `json:"items"`
	PageInfo struct {
		CountTotal int `json:"countTotal"`
		Count      int `json:"count"`
	} `json:"pageInfo"`
}

func (mi *Markets) UnmarshalJSON(data []byte) error {
	var raw struct {
		Items    []json.RawMessage `json:"items"`
		PageInfo struct {
			CountTotal int `json:"countTotal"`
			Count      int `json:"count"`
		} `json:"pageInfo"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	for _, itemData := range raw.Items {
		var market Market
		if err := json.Unmarshal(itemData, &market); err != nil {
			continue
		}

		if market.LLTV.Cmp(big.NewInt(0)) != 0 &&
			len(market.Metadata.Name) <= 30 &&
			market.State.Price != nil {
			mi.Items = append(mi.Items, market)
		}
	}

	return nil
}
