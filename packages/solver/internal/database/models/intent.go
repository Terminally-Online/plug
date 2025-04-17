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
	Id                         string              `json:"id" gorm:"primaryKey;type:text"`
	Status                     string              `json:"status" gorm:"type:text;default:'active'"`
	ChainId                    uint64              `json:"chainId" gorm:"type:int;not null"`
	From                       string              `json:"from" gorm:"type:text;not null"`
	Value                      *types.BigInt       `json:"value" gorm:"type:bigint"`
	GasLimit                   *uint64             `json:"gasLimit" gorm:"type:int"`
	Inputs                     types.Inputs        `json:"inputs" gorm:"type:jsonb"`
	Options                    types.Options       `json:"options" gorm:"type:jsonb"`
	Frequency                  int                 `json:"frequency" gorm:"type:int"`
	AccessList                 ethTypes.AccessList `json:"accessList,omitempty" gorm:"type:jsonb"`
	StartAt                    *time.Time          `json:"startAt" gorm:"type:timestamp"`
	EndAt                      *time.Time          `json:"endAt" gorm:"type:timestamp"`
	PeriodEndAt                *time.Time          `json:"periodEndAt" gorm:"type:timestamp"`
	NextSimulationAt           *time.Time          `json:"nextSimulationAt" gorm:"type:timestamp"`
	Saved                      bool                `json:"saved" gorm:"type:boolean"`
	Locked                     bool                `json:"locked" gorm:"type:boolean"`
	SignatureExpirationMinutes int                 `json:"signatureExpirationMinutes,omitempty" gorm:"type:int;default:5"`

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

func (i *Intent) ValidateFields() error {
	if i.ChainId == 0 {
		return fmt.Errorf("missing 'chainId'")
	}
	if i.From == "" {
		return fmt.Errorf("missing 'from' address")
	}
	if len(i.Inputs) == 0 {
		return fmt.Errorf("missing 'inputs'")
	}

	return nil
}
