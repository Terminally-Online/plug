package models

import (
	"time"

	"gorm.io/gorm"
)

type Run struct {
	Id string `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status string `json:"status" gorm:"type:text"`
	Result string `json:"result,omitempty" gorm:"type:jsonb"`
	Error string `json:"error,omitempty" gorm:"type:text"`
	Errors []string `json:"errors,omitempty" gorm:"type:text[]"`
	GasEstimate uint64 `json:"gasEstimate,omitempty" gorm:"type:bigint"`

	// Relationships
	IntentId string `json:"intentId,omitempty" gorm:"type:text"`
	Intent Intent `json:"intent,omitempty" gorm:"foreignKey:IntentId;references:Id"`

	// Database storage fields
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
