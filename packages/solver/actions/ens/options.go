package ens

import (
	"fmt"
	"solver/types"
)

const (
	secondsPerMonth = 2629800
	secondsPerYear  = 31557600
)

func GetDurationOptions() ([]types.Option, error) {
	return []types.Option{
		{
			Value: fmt.Sprint(secondsPerMonth),
			Label: "1 Month",
		},
		{
			Value: fmt.Sprint(secondsPerMonth * 3),
			Label: "3 Months",
		},
		{
			Value: fmt.Sprint(secondsPerMonth * 6),
			Label: "6 Months",
		},
		{
			Value: fmt.Sprint(secondsPerYear),
			Label: "1 Year",
		},
		{
			Value: fmt.Sprint(secondsPerYear * 3),
			Label: "3 Years",
		},
		{
			Value: fmt.Sprint(secondsPerYear * 5),
			Label: "5 Years",
		},
	}, nil
}
