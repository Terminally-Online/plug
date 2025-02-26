package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"math/big"
	"solver/internal/database/serializer"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

// InputsSlice is a custom type that implements sql.Scanner for proper JSONB handling
type InputsSlice []map[string]interface{}

// Scan implements the sql.Scanner interface for InputsSlice
func (is *InputsSlice) Scan(value interface{}) error {
	if value == nil {
		*is = make(InputsSlice, 0)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value: invalid data type")
	}

	if err := json.Unmarshal(bytes, is); err != nil {
		return err
	}
	return nil
}

// Value implements the driver.Valuer interface for InputsSlice
func (is InputsSlice) Value() (driver.Value, error) {
	if is == nil {
		return nil, nil
	}
	return json.Marshal(is)
}

type Intent struct {
	Id         string                 `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status     string                 `json:"status,omitempty" gorm:"type:text;default:'active'"`
	ChainId    uint64                 `json:"chainId" gorm:"type:int"`
	From       string                 `json:"from,omitempty" gorm:"type:text"`
	Value      *big.Int               `json:"value,omitempty" db_field:"ValueStr" gorm:"-"`
	GasLimit   *uint64                `json:"gasLimit,omitempty" gorm:"type:int"`
	Inputs     InputsSlice            `json:"inputs,omitempty" gorm:"type:jsonb"`
	Options    map[string]interface{} `json:"options,omitempty" gorm:"type:jsonb"`
	Frequency  int                    `json:"frequency,omitempty" gorm:"type:int"`
	AccessList types.AccessList       `json:"accessList,omitempty" gorm:"type:jsonb"`

	StartAt          *time.Time `json:"startAt,omitempty" gorm:"type:timestamp"`
	EndAt            *time.Time `json:"endAt,omitempty" gorm:"type:timestamp"`
	PeriodEndAt      *time.Time `json:"periodEndAt,omitempty" gorm:"type:timestamp"`
	NextSimulationAt *time.Time `json:"nextSimulationAt,omitempty" gorm:"type:timestamp"`

	// Relationships
	Runs     []Run  `json:"runs" gorm:"foreignKey:IntentId;references:Id"`
	ApiKeyId string `json:"-" gorm:"column:api_key_id;type:text"`
	ApiKey   ApiKey `json:"-" gorm:"foreignKey:ApiKeyId;references:Id"`

	// Database storage fields
	ValueStr  string         `json:"-" gorm:"column:value;type:text"`
	CreatedAt time.Time      `json:"createdAt"`
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

func (i *Intent) BeforeSave(tx *gorm.DB) error {
	return serializer.HandleBeforeSave(i)
}

func (i *Intent) BeforeUpdate(tx *gorm.DB) error {
	// TODO: @masonchain - We can't do it on every update, but after every simulation e need to call GetNextSimulationAt and update the PeriodEndAt and NextSimulationAt
	return nil
}

func (i *Intent) AfterFind(tx *gorm.DB) error {
	return serializer.HandleAfterFind(i)
}
