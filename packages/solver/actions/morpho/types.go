package morpho

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"solver/bindings/morpho_router"

	"github.com/ethereum/go-ethereum/common"
)

type Asset struct {
	Address  string `json:"address"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

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

	ms.Price, err = parseBigInt(raw.Price)
	if err != nil {
		return fmt.Errorf("parsing price: %w", err)
	}

	ms.BorrowShares, err = parseBigInt(raw.BorrowShares)
	if err != nil {
		return fmt.Errorf("parsing borrowShares: %w", err)
	}

	ms.BorrowAssets, err = parseBigInt(raw.BorrowAssets)
	if err != nil {
		return fmt.Errorf("parsing borrowAssets: %w", err)
	}

	ms.SupplyShares, err = parseBigInt(raw.SupplyShares)
	if err != nil {
		return fmt.Errorf("parsing supplyShares: %w", err)
	}

	ms.SupplyAssets, err = parseBigInt(raw.SupplyAssets)
	if err != nil {
		return fmt.Errorf("parsing supplyAssets: %w", err)
	}

	return nil
}

type Market struct {
	Whitelisted bool   `json:"whitelisted"`
	UniqueKey   string `json:"uniqueKey"`
	Metadata    struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	} `json:"metadata"`
	Params          morpho_router.MarketParams `json:"params"`
	LLTV            *big.Int                   `json:"lltv"`
	OracleAddress   string                     `json:"oracleAddress"`
	IRMAddress      string                     `json:"irmAddress"`
	LoanAsset       Asset                      `json:"loanAsset,omitempty"`
	CollateralAsset Asset                      `json:"collateralAsset,omitempty"`
	State           MarketState                `json:"state"`
}

func (m *Market) UnmarshalJSON(data []byte) error {
	var raw struct {
		Whitelisted     bool        `json:"whitelisted"`
		UniqueKey       string      `json:"uniqueKey"`
		LLTV            interface{} `json:"lltv"`
		OracleAddress   string      `json:"oracleAddress"`
		IRMAddress      string      `json:"irmAddress"`
		LoanAsset       Asset       `json:"loanAsset"`
		CollateralAsset Asset       `json:"collateralAsset"`
		State           MarketState `json:"state"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("unmarshaling raw market: %w", err)
	}

	m.Whitelisted = raw.Whitelisted
	m.UniqueKey = raw.UniqueKey
	m.OracleAddress = raw.OracleAddress
	m.IRMAddress = raw.IRMAddress
	m.LoanAsset = raw.LoanAsset
	m.CollateralAsset = raw.CollateralAsset
	m.State = raw.State

	lltv, err := parseBigInt(raw.LLTV)
	if err != nil {
		return fmt.Errorf("parsing LLTV: %w", err)
	}
	m.LLTV = lltv

	m.Metadata.Name = fmt.Sprintf(
		"%s/%s (LLTV: %.2f%%)",
		m.CollateralAsset.Symbol,
		m.LoanAsset.Symbol,
		new(big.Float).Quo(
			new(big.Float).SetInt(lltv),
			new(big.Float).SetInt(big.NewInt(10).Exp(big.NewInt(10), big.NewInt(16), nil)),
		),
	)
	m.Metadata.Icon = url.QueryEscape(fmt.Sprintf(
		"%s|%s",
		m.CollateralAsset.LogoURI,
		m.LoanAsset.LogoURI,
	))
	m.Params = morpho_router.MarketParams{
		LoanToken:       common.HexToAddress(m.LoanAsset.Address),
		CollateralToken: common.HexToAddress(m.CollateralAsset.Address),
		Oracle:          common.HexToAddress(m.OracleAddress),
		Irm:             common.HexToAddress(m.IRMAddress),
		Lltv:            lltv,
	}

	return nil
}

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

type Vault struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Metadata struct {
		Image       string `json:"image"`
		Description string `json:"description"`
		Curators    []struct {
			Name     string `json:"name"`
			Image    string `json:"image"`
			URL      string `json:"url"`
			Verified bool   `json:"verified"`
		} `json:"curators"`
	} `json:"metadata"`
	State struct {
		Rewards []struct {
			Asset Asset `json:"asset"`
		} `json:"rewards"`
		NetApy               float64 `json:"netApy"`
		NetApyWithoutRewards float64 `json:"netApyWithoutRewards"`
		Allocation           struct {
			Enabled bool     `json:"enabled"`
			Market  []Market `json:"market"`
		} `json:"allocation"`
	} `json:"state"`
}

type Distribution struct {
	Claimable   string   `json:"claimable"`
	Proof       []string `json:"proof"`
	TxData      string   `json:"txData"`
	Distributor struct {
		Id      string `json:"id"`
		Address string `json:"address"`
		ChainId int    `json:"chain_id"`
	} `json:"distributor"`
	Asset struct {
		Id      string `json:"id"`
		Address string `json:"address"`
		ChainId int    `json:"chain_id"`
	} `json:"asset"`
}

type DistributionResponse struct {
	Timestamp  int64 `json:"timestamp"`
	Pagination struct {
		PerPage    int    `json:"per_page"`
		Page       int    `json:"page"`
		TotalPages int    `json:"total_pages"`
		Next       string `json:"next,omitempty"`
		Prev       string `json:"prev,omitempty"`
	} `json:"pagination"`
	Data []Distribution `json:"data"`
}
