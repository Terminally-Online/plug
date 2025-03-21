package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"solver/internal/database/serializer"
	"solver/internal/database/types"
	"solver/internal/solver/signature"
	"solver/internal/utils"
	"time"

	"gorm.io/gorm"
)

type Run struct {
	Id              string        `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Status          string        `json:"status" gorm:"type:text"`
	Error           *string       `json:"error,omitempty" gorm:"type:text"`
	Errors          []string      `json:"errors,omitempty" gorm:"type:text[]"`
	GasUsed         uint64        `json:"gasEstimate,omitempty" gorm:"type:bigint"`
	GasPrice        *uint64       `json:"gasPrice,omitempty" gorm:"type:bigint"`
	From            string        `json:"from,omitempty" gorm:"type:text"`
	To              string        `json:"to,omitempty" gorm:"type:text"`
	Value           *types.BigInt `json:"value,omitempty" gorm:"type:bigint"`
	Data            RunOutputData `json:"data,omitempty" gorm:"type:jsonb"`
	TransactionHash *string       `json:"transactionHash,omitempty" gorm:"type:text"`

	IntentId    string               `json:"intentId,omitempty" gorm:"type:text"`
	Intent      Intent               `json:"-" gorm:"foreignKey:IntentId;references:Id"`
	LivePlugsId string               `json:"livePlugsId,omitempty" gorm:"type:text;column:live_plugs_id;default:null"`
	LivePlugs   *signature.LivePlugs `json:"livePlugs,omitempty" gorm:"foreignKey:LivePlugsId;references:Id"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type RunOutputData struct {
	Raw     []byte `json:"raw,omitempty" gorm:"type:bytea"`
	Decoded any    `json:"decoded,omitempty" gorm:"type:jsonb"`
}

// Add these methods to the RunOutputData type
func (r *RunOutputData) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("expected bytes but got %T", value)
	}
	return json.Unmarshal(bytes, r)
}

func (r RunOutputData) Value() (driver.Value, error) {
	return json.Marshal(r)
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

func (r Run) MarshalJSON() ([]byte, error) {
	type Alias Run

	// Create a temp struct without the LivePlugs
	tmp := &struct {
		*Alias
		LivePlugs *signature.LivePlugs `json:"livePlugs,omitempty"`
	}{
		Alias: (*Alias)(&r),
	}

	// Only include the LivePlugs if it has a non-zero Id
	if r.LivePlugs != nil && r.LivePlugs.Id != "" {
		tmp.LivePlugs = r.LivePlugs
	}

	return json.Marshal(tmp)
}
