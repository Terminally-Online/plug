package models

import (
	"solver/internal/utils"
	"time"

	"gorm.io/gorm"
)

type ApiKey struct {
	Id        string `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Key       string `json:"key" gorm:"uniqueIndex;type:text;not null"`
	SocketId  string `json:"socketId,omitempty" gorm:"type:text"`
	RateLimit int    `json:"rateLimit" gorm:"type:integer;not null;default:100"`
	Role      string `json:"role" gorm:"type:text;not null;default:'user'"`

	// Store the timestamps but do not expose them in the JSON response
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (a *ApiKey) BeforeCreate(tx *gorm.DB) error {
	a.Id = utils.GenerateUUID()
	a.Key = utils.GenerateApiKey()

	return nil
}
