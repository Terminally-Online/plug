package models

import (
	"solver/internal/utils"
	"time"

	"gorm.io/gorm"
)

type Run struct {
	Id          string        `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status      string        `json:"status" gorm:"type:text"`
	Result      string        `json:"result,omitempty" gorm:"type:jsonb"`
	Error       string        `json:"error,omitempty" gorm:"type:text"`
	Errors      []string      `json:"errors,omitempty" gorm:"type:text[]"`
	GasEstimate uint64        `json:"gasEstimate,omitempty" gorm:"type:bigint"`
	Data        RunOutputData `json:"data" gorm:"type:jsonb"`

	// Relationships
	IntentId  string     `json:"intentId,omitempty" gorm:"type:text"`
	Intent    Intent     `json:"intent,omitempty" gorm:"foreignKey:IntentId;references:Id"`
	Execution *Execution `json:"execution,omitempty" gorm:"foreignKey:RunId"`

	// Store the timestamps but do not expose them in the JSON response
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type RunOutputData struct {
	Raw     []byte      `json:"raw" gorm:"type:bytea"`
	Decoded interface{} `json:"decoded,omitempty" gorm:"type:jsonb"`
}

func (r *Run) BeforeCreate(tx *gorm.DB) error {
	r.Id = utils.GenerateUUID()

	return nil
}
