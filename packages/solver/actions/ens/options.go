package ens

import (
	"fmt"
	"solver/types"
)

const (
	secondsPerMonth = 2629800
	secondsPerYear  = 31557600
)

type EnsOptionsProvider struct{}

func (p *EnsOptionsProvider) GetOptions(chainId int, action types.Action) (map[int]types.SchemaOptions, error) {
	durationOptions, err := GetDurationOptions()
	if err != nil {
		return nil, err
	}

	switch action {
	case types.ActionRenew:
		return map[int]types.SchemaOptions{
			1: {Simple: durationOptions},
		}, nil
	case types.Action(RenewalPrice):
		return map[int]types.SchemaOptions{
			1: {Simple: durationOptions},
		}, nil
	case types.Action(TimeLeft):
		return map[int]types.SchemaOptions{
			1: {Simple: durationOptions},
		}, nil
	default:
		return nil, nil
	}
}

func GetDurationOptions() ([]types.Option, error) {
	return []types.Option{
		{
			Value: fmt.Sprint(secondsPerMonth),
			Label: "1 Month",
			Name:  "1 Month",
		},
		{
			Value: fmt.Sprint(secondsPerMonth * 3),
			Label: "3 Months",
			Name:  "3 Months",
		},
		{
			Value: fmt.Sprint(secondsPerMonth * 6),
			Label: "6 Months",
			Name:  "6 Months",
		},
		{
			Value: fmt.Sprint(secondsPerYear),
			Label: "1 Year",
			Name:  "1 Year",
		},
		{
			Value: fmt.Sprint(secondsPerYear * 3),
			Label: "3 Years",
			Name:  "3 Years",
		},
		{
			Value: fmt.Sprint(secondsPerYear * 5),
			Label: "5 Years",
			Name:  "5 Years",
		},
	}, nil
}
