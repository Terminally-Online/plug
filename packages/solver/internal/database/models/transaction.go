package models

import (
	"solver/internal/database/serializer"
	"solver/internal/utils"
	"time"

	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

type Transaction struct {
	Id         string              `json:"id,omitempty" gorm:"primaryKey;type:text"`
	From       string              `json:"from,omitempty" gorm:"type:text"`
	To         string              `json:"to,omitempty" gorm:"type:text"`
	ChainId    uint64              `json:"chainId" gorm:"type:int"`
	Value      string              `json:"value,omitempty" gorm:"type:text"`
	Data       string              `json:"data,omitempty" gorm:"type:text"`
	GasLimit   *string             `json:"gasLimit,omitempty" gorm:"type:text"`
	AccessList ethTypes.AccessList `json:"accessList,omitempty" gorm:"type:jsonb"`

	IntentId string `json:"intentId,omitempty" gorm:"type:text"`
	Intent   Intent `json:"-" gorm:"foreignKey:IntentId;references:Id"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	t.Id = utils.GenerateUUID()
	return nil
}

func (t *Transaction) BeforeSave(tx *gorm.DB) error {
	return serializer.HandleBeforeSave(t)
}

func (t *Transaction) AfterFind(tx *gorm.DB) error {
	return serializer.HandleAfterFind(t)
}
