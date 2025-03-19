package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"solver/internal/actions/morpho/utils"
)

type MarketState struct {
	Price          *big.Int `json:"price,omitempty"`
	BorrowShares   *big.Int `json:"borrowShares"`
	BorrowAssets   *big.Int `json:"borrowAssets"`
	BorrowApy      float64  `json:"borrowApy"`
	DailyBorrowApy float64  `json:"dailyBorrowApy"`
	SupplyShares   *big.Int `json:"supplyShares"`
	SupplyAssets   *big.Int `json:"supplyAssets"`
	SupplyApy      float64  `json:"supplyApy"`
	DailySupplyApy float64  `json:"dailySupplyApy"`
}

func (ms *MarketState) UnmarshalJSON(data []byte) error {
	var raw struct {
		Price          interface{} `json:"price,omitempty"`
		BorrowShares   interface{} `json:"borrowShares"`
		BorrowAssets   interface{} `json:"borrowAssets"`
		BorrowApy      float64     `json:"borrowApy"`
		DailyBorrowApy float64     `json:"dailyBorrowApy"`
		SupplyShares   interface{} `json:"supplyShares"`
		SupplyAssets   interface{} `json:"supplyAssets"`
		SupplyApy      float64     `json:"supplyApy"`
		DailySupplyApy float64     `json:"dailySupplyApy"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("unmarshaling raw market state: %w", err)
	}

	ms.BorrowApy = raw.BorrowApy
	ms.DailyBorrowApy = raw.DailyBorrowApy
	ms.SupplyApy = raw.SupplyApy
	ms.DailySupplyApy = raw.DailySupplyApy

	var err error

	ms.Price, err = utils.ParseBigInt(raw.Price)
	if err != nil {
		return fmt.Errorf("parsing price: %w", err)
	}

	ms.BorrowShares, err = utils.ParseBigInt(raw.BorrowShares)
	if err != nil {
		return fmt.Errorf("parsing borrowShares: %w", err)
	}

	ms.BorrowAssets, err = utils.ParseBigInt(raw.BorrowAssets)
	if err != nil {
		return fmt.Errorf("parsing borrowAssets: %w", err)
	}

	ms.SupplyShares, err = utils.ParseBigInt(raw.SupplyShares)
	if err != nil {
		return fmt.Errorf("parsing supplyShares: %w", err)
	}

	ms.SupplyAssets, err = utils.ParseBigInt(raw.SupplyAssets)
	if err != nil {
		return fmt.Errorf("parsing supplyAssets: %w", err)
	}

	return nil
}
