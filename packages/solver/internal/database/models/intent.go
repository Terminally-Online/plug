package models

import (
	"solver/internal/utils"
	"time"

	"gorm.io/gorm"
)

type Intent struct {
	Id               string                   `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status           string                   `json:"status,omitempty" gorm:"type:text;default:'active'"`
	ChainId          int                      `json:"chainId" gorm:"type:int"`
	From             string                   `json:"from,omitempty" gorm:"type:text"`
	Actions          []map[string]interface{} `json:"actions,omitempty" gorm:"type:jsonb"`
	Options          map[string]interface{}   `json:"options,omitempty" gorm:"type:jsonb"`
	Frequency        int                      `json:"frequency,omitempty" gorm:"type:int"`
	StartAt          *time.Time               `json:"startAt,omitempty" gorm:"type:timestamp"`
	EndAt            *time.Time               `json:"endAt,omitempty" gorm:"type:timestamp"`
	PeriodEndAt      *time.Time               `json:"periodEndAt,omitempty" gorm:"type:timestamp"`
	NextSimulationAt *time.Time               `json:"nextSimulationAt,omitempty" gorm:"type:timestamp"`

	// Relationships
	Runs     []Run  `json:"runs,omitempty" gorm:"foreignKey:IntentId;references:Id"`
	ApiKeyId string `json:"-" gorm:"column:api_key_id;type:text"`
	ApiKey   ApiKey `json:"-" gorm:"foreignKey:ApiKeyId;references:Id"`

	// Store the timestamps but do not expose them in the JSON response
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (i *Intent) GetNextSimulationAt() (periodEndAt *time.Time, nextSimulationAt *time.Time) {
	now := time.Now()

	if i.EndAt == nil || i.EndAt.Before(now) || i.Frequency == 0 {
		return nil, nil
	}

	if i.PeriodEndAt != nil && i.Frequency > 0 {
		var latest *Run
		for _, run := range i.Runs {
			if latest == nil || run.CreatedAt.After(latest.CreatedAt) {
				latest = &run
			}
		}
		if latest.Status == "success" {
			executionFrequency := time.Duration(i.Frequency) * 24 * time.Hour
			nextPeriodEnd := i.PeriodEndAt.Add(executionFrequency)

			if nextPeriodEnd.After(*i.EndAt) {
				return i.EndAt, nil
			}

			return &nextPeriodEnd, i.PeriodEndAt
		}

		nextSimAt := time.Now().Add(5 * time.Minute)
		return i.PeriodEndAt, &nextSimAt
	}

	return nil, nil
}

func (i *Intent) BeforeCreate(tx *gorm.DB) error {
	i.Id = utils.GenerateUUID()
	if i.Frequency > 0 {
		periodEndAt := i.StartAt.Add(time.Duration(i.Frequency) * 24 * time.Hour)
		i.PeriodEndAt = &periodEndAt
	}
	nextSim := i.StartAt
	i.NextSimulationAt = nextSim

	return nil
}

func (i *Intent) BeforeUpdate(tx *gorm.DB) error {
	// TODO: @masonchain - We can't do it on every update, but after every simulation e need to call GetNextSimulationAt and update the PeriodEndAt and NextSimulationAt
	return nil
}

func (i *Intent) AfterFind(tx *gorm.DB) error {
	return nil
}
