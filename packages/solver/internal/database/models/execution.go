package models

import (
	"solver/internal/utils"
	"time"

	"gorm.io/gorm"
)

type Execution struct {
	Id              string `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status          string `json:"status,omitempty" gorm:"type:text"`
	GasUsed         uint64 `json:"gasUsed,omitempty" gorm:"type:bigint"`
	GasPrice        uint64 `json:"gasPrice,omitempty" gorm:"type:bigint"`
	TransactionHash string `json:"transactionHash,omitempty" gorm:"type:text"`

	// Relationships
	TransactionId string      `json:"transactionId,omitempty" gorm:"type:text"`
	Transaction   Transaction `json:"-" gorm:"foreignKey:TransactionId;references:Id"`

	// Store the timestamps but do not expose them in the JSON response
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (e *Execution) BeforeCreate(tx *gorm.DB) error {
	e.Id = utils.GenerateUUID()

	return nil
}
