package morpho

import (
	"fmt"
	"net/url"
	"solver/types"
	"strconv"
)

func GetMarketOptions() ([]types.Option, error) {
	markets, err := GetMarkets()
	if err != nil {
		return nil, err
	}

	options := make([]types.Option, 0)
	for _, market := range markets {
		lltv := func() float64 {
			switch v := market.LLTV.(type) {
			case string:
				f, _ := strconv.ParseFloat(v, 64)
				return f / 1e16
			case float64:
				return v / 1e16
			default:
				return 0
			}
		}()

		reference := fmt.Sprintf(
			"%s/%s (LLTV: %.0f%%)",
			market.CollateralAsset.Symbol,
			market.LoanAsset.Symbol,
			lltv,
		)
		if lltv == 0 || len(reference) > 30 {
			continue
		}

		urlEncoded := url.QueryEscape(fmt.Sprintf(
			"%s|%s",
			market.CollateralAsset.LogoURI,
			market.LoanAsset.LogoURI,
		))

		options = append(options, types.Option{
			Label: reference,
			Name:  reference,
			Value: market.UniqueKey,
			Icon:  urlEncoded,
		})
	}

	return options, nil
}
