package models

import (
	"fmt"
	"solver/internal/database/serializer"
	"solver/internal/database/types"
	"solver/internal/solver/signature"
	"solver/internal/utils"

	"time"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type Intent struct {
	Id               string              `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status           string              `json:"status,omitempty" gorm:"type:text;default:'active'"`
	ChainId          uint64              `json:"chainId" gorm:"type:int"`
	From             string              `json:"from,omitempty" gorm:"type:text"`
	Value            *types.BigInt       `json:"value,omitempty" db_field:"ValueStr" gorm:"-"`
	GasLimit         *uint64             `json:"gasLimit,omitempty" gorm:"type:int"`
	Inputs           types.Inputs        `json:"inputs,omitempty" gorm:"type:jsonb"`
	Options          types.Options       `json:"options,omitempty" gorm:"type:jsonb"`
	Frequency        int                 `json:"frequency,omitempty" gorm:"type:int"`
	AccessList       ethTypes.AccessList `json:"accessList,omitempty" gorm:"type:jsonb"`
	StartAt          *time.Time          `json:"startAt,omitempty" gorm:"type:timestamp"`
	EndAt            *time.Time          `json:"endAt,omitempty" gorm:"type:timestamp"`
	PeriodEndAt      *time.Time          `json:"periodEndAt,omitempty" gorm:"type:timestamp"`
	NextSimulationAt *time.Time          `json:"nextSimulationAt,omitempty" gorm:"type:timestamp"`
	Saved            bool                `json:"saved,omitempty" gorm:"type:boolean"`
	Locked           bool                `json:"locked,omitempty" gorm:"type:boolean"`

	Runs      []Run                 `json:"runs" gorm:"foreignKey:IntentId;references:Id"`
	LivePlugs []signature.LivePlugs `json:"-" gorm:"foreignKey:IntentId;references:Id"`
	ApiKeyId  string                `json:"-" gorm:"column:api_key_id;type:text"`
	ApiKey    ApiKey                `json:"-" gorm:"foreignKey:ApiKeyId;references:Id"`

	// Database storage fields
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
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

func (i *Intent) BeforeSave(tx *gorm.DB) error {
	return serializer.HandleBeforeSave(i)
}

func (i *Intent) AfterCreate(tx *gorm.DB) error {
	i.Runs = []Run{}

	return nil
}

func (i *Intent) AfterFind(tx *gorm.DB) error {
	return serializer.HandleAfterFind(i)
}

func (i *Intent) GetOrCreate(db *gorm.DB) (*Intent, error) {
	var intent Intent
	if err := db.Where("id = ?", i.Id).First(&intent).Error; err == gorm.ErrRecordNotFound {
		if err := db.Create(i).Error; err != nil {
			return nil, fmt.Errorf("failed to create intent: %w", err)
		}
		return i, nil
	}

	return &intent, nil
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
