package models

import (
	"math/big"
	"time"

	"gorm.io/gorm"
)

type Intent struct {
	Id               string     `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status           string     `json:"status,omitempty" gorm:"type:text"`
	ChainId          int        `json:"chainId" gorm:"type:int"`
	From             string     `json:"from,omitempty" gorm:"type:text"`
	To               string     `json:"to,omitempty" gorm:"type:text"`
	Actions          []string   `json:"actions,omitempty" gorm:"type:text[]"`
	Value            *big.Int   `json:"value,omitempty" db_field:"ValueStr" gorm:"-"`
	Frequency        int        `json:"frequency,omitempty" gorm:"type:int"`
	StartAt          time.Time  `json:"startAt,omitempty" gorm:"type:timestamp"`
	EndAt            time.Time  `json:"endAt,omitempty" gorm:"type:timestamp"`
	PeriodEndAt      *time.Time `json:"periodEndAt,omitempty" gorm:"type:timestamp"`
	NextSimulationAt *time.Time `json:"nextSimulationAt,omitempty" gorm:"type:timestamp"`

	// Relationships
	Runs   []Run  `json:"runs,omitempty" gorm:"foreignKey:IntentId;references:Id"`
	ApiKey ApiKey `json:"-" gorm:"foreignKey:ApiKey;references:Id"`

	// Database storage fields
	ValueStr  string         `json:"-" gorm:"column:value;type:text"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
