package models

import (
	"math/big"
	"solver/internal/database/serializer"
	"solver/internal/database/types"
	"solver/internal/utils"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"gorm.io/gorm"
)

type Run struct {
	Id          string        `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status      string        `json:"status" gorm:"type:text"`
	Error       *string       `json:"error,omitempty" gorm:"type:text"`
	Errors      []string      `json:"errors,omitempty" gorm:"type:text[]"`
	GasEstimate uint64        `json:"gasEstimate,omitempty" gorm:"type:bigint"`
	From        string        `json:"from,omitempty" gorm:"type:text"`
	To          string        `json:"to,omitempty" gorm:"type:text"`
	Value       *big.Int      `json:"value,omitempty" db_field:"ValueStr" gorm:"-"`
	Inputs      types.Inputs  `json:"inputs,omitempty" gorm:"type:jsonb"`
	CallData    hexutil.Bytes `json:"-" db_field:"CallDataStr" gorm:"-"`
	ResultData  RunOutputData `json:"resultData" gorm:"type:jsonb"`

	// Relationships
	IntentId  string     `json:"intentId,omitempty" gorm:"type:text"`
	Intent    Intent     `json:"-" gorm:"foreignKey:IntentId;references:Id"`
	Execution *Execution `json:"-" gorm:"foreignKey:RunId"`

	// Store the timestamps but do not expose them in the JSON response
	ValueStr    string         `json:"-" gorm:"column:value;type:text"`
	CallDataStr string         `json:"-" gorm:"column:calldata;type:text"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}

type RunOutputData struct {
	Raw     []byte      `json:"raw" gorm:"type:bytea"`
	Decoded interface{} `json:"decoded,omitempty" gorm:"type:jsonb"`
}

func (r *Run) BeforeCreate(tx *gorm.DB) error {
	r.Id = utils.GenerateUUID()
	return nil
}

func (r *Run) BeforeSave(tx *gorm.DB) error {
	return serializer.HandleBeforeSave(r)
}

func (r *Run) AfterFind(tx *gorm.DB) error {
	return serializer.HandleAfterFind(r)
}
