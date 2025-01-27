package ens

import (
	"fmt"
	"solver/internal/actions"
)

const (
	secondsPerMonth = 2629800
	secondsPerYear  = 31557600
)

type EnsOptionsProvider struct{}

func (p *EnsOptionsProvider) GetOptions(chainId uint64, action string) (map[int]actions.Options, error) {
	durationOptions, err := GetDurationOptions()
	if err != nil {
		return nil, err
	}

	switch action {
	case actions.ActionRenew:
		return map[int]actions.Options{
			1: {Simple: durationOptions},
		}, nil
	case actions.ConstraintPrice:
		return map[int]actions.Options{
			1: {Simple: durationOptions},
		}, nil
	case TimeLeft:
		return map[int]actions.Options{
			1: {Simple: durationOptions},
		}, nil
	default:
		return nil, nil
	}
}

func GetDurationOptions() ([]actions.Option, error) {
	return []actions.Option{
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
