package models

import (
	"time"

	"gorm.io/gorm"
)

type ApiKey struct {
	Id        string `json:"id,omitempty" gorm:"primaryKey;type:text"`
	Key       string `json:"key" gorm:"uniqueIndex;type:text;not null"`
	SocketId  string `json:"socketId,omitempty" gorm:"type:text"`
	RateLimit int    `json:"rateLimit" gorm:"type:integer;not null;default:0"`
	Role 	string `json:"role" gorm:"type:text;not null;default:'user'"`

	// Database storage fields
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
