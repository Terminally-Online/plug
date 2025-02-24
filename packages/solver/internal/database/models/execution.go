package models

import (
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
	Run Run `json:"run,omitempty" gorm:"foreignKey:ExecutionId;references:Id"`

	// Database storage fields
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
