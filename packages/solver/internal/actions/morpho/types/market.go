package types

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/url"
	"solver/bindings/morpho_router"
	"solver/internal/actions/morpho/utils"

	"github.com/ethereum/go-ethereum/common"
)

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
	DailyApys       struct {
		BorrowApy    float64 `json:"borrowApy"`
		SupplyApy    float64 `json:"supplyApy"`
		NetBorrowApy float64 `json:"netBorrowApy"`
		NetSupplyApy float64 `json:"netSupplyApy"`
	} `json:"dailyApys"`
	State MarketState `json:"state"`
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
		DailyApys       struct {
			BorrowApy    float64 `json:"borrowApy"`
			SupplyApy    float64 `json:"supplyApy"`
			NetBorrowApy float64 `json:"netBorrowApy"`
			NetSupplyApy float64 `json:"netSupplyApy"`
		} `json:"dailyApys"`
		State MarketState `json:"state"`
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
	m.DailyApys = raw.DailyApys
	m.State = raw.State

	lltv, err := utils.ParseBigInt(raw.LLTV)
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
