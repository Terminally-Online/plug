package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"solver/internal/database/serializer"
	"solver/internal/database/types"
	"solver/internal/utils"
	"time"

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
	Value       *types.BigInt `json:"value,omitempty" db_field:"ValueStr" gorm:"-"`
	ResultData  RunOutputData `json:"resultData" gorm:"type:jsonb"`

	IntentId   string   `json:"intentId,omitempty" gorm:"type:text"`
	Intent     Intent   `json:"-" gorm:"foreignKey:IntentId;references:Id"`
	LivePlugId string   `json:"livePlugId,omitempty" gorm:"type:text"`
	LivePlug   LivePlug `json:"livePlug,omitempty" gorm:"foreignKey:LivePlugId;references:Id"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type RunOutputData struct {
	Raw     []byte      `json:"raw" gorm:"type:bytea"`
	Decoded interface{} `json:"decoded,omitempty" gorm:"type:jsonb"`
}

// Add these methods to the RunOutputData type
func (r *RunOutputData) Scan(value interface{}) error {
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

	// Create a temp struct without the bundle
	tmp := &struct {
		*Alias
		LivePlug *LivePlug `json:"transactionBundle,omitempty"`
	}{
		Alias: (*Alias)(&r),
	}

	// Only include the bundle if it has a non-zero ID
	if r.LivePlug.Id != "" {
		tmp.LivePlug = &r.LivePlug
	}

	return json.Marshal(tmp)
}
