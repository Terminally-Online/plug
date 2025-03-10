package models
//
// import (
// 	"solver/internal/database/serializer"
// 	"solver/internal/database/types"
// 	"solver/internal/utils"
// 	"time"
//
// 	"gorm.io/gorm"
// )
//
// type LivePlug struct {
// 	Id        string        `json:"id,omitempty" gorm:"primaryKey;type:text"`
// 	ChainId   uint64        `json:"chainId" gorm:"type:int"`
// 	From      string        `json:"from,omitempty" gorm:"type:text"`
// 	To        string        `json:"to,omitempty" gorm:"type:text"`
// 	Value     *types.BigInt `json:"value,omitempty" gorm:"type:bigint"`
// 	Gas       *types.BigInt `json:"gas,omitempty" gorm:"type:bigint"`
// 	Data      string        `json:"data,omitempty" gorm:"type:bytea"`
// 	Signature *string       `json:"signature,omitempty" gorm:"type:bytea"`
//
// 	IntentId string `json:"intentId,omitempty" gorm:"type:text"`
// 	Intent   Intent `json:"-" gorm:"foreignKey:IntentId;references:Id"`
// 	Plugs    []Plug `json:"plugs" gorm:"foreignKey:LivePlugId;references:Id"`
//
// 	CreatedAt time.Time      `json:"-"`
// 	UpdatedAt time.Time      `json:"-"`
// 	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
// }
//
// type Plug struct {
// 	Id        string        `json:"id,omitempty" gorm:"primaryKey;type:text"`
// 	From      string        `json:"from,omitempty" gorm:"type:text"`
// 	To        string        `json:"to,omitempty" gorm:"type:text"`
// 	Data      string        `json:"data,omitempty" gorm:"type:bytea"`
// 	Value     *types.BigInt `json:"value,omitempty" gorm:"type:bigint"`
// 	Gas       *types.BigInt `json:"gas,omitempty" gorm:"type:bigint"`
// 	Exclusive bool          `json:"exclusive,omitempty" gorm:"type:boolean"`
//
// 	LivePlugId string   `json:"livePlugId,omitempty" gorm:"type:text"`
// 	LivePlug   LivePlug `json:"-" gorm:"foreignKey:LivePlugId;references:Id"`
//
// 	CreatedAt time.Time      `json:"-"`
// 	UpdatedAt time.Time      `json:"-"`
// 	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
// }
//
// func (t *LivePlug) BeforeCreate(tx *gorm.DB) error {
// 	t.Id = utils.GenerateUUID()
// 	return nil
// }
//
// func (t *LivePlug) BeforeSave(tx *gorm.DB) error {
// 	return serializer.HandleBeforeSave(t)
// }
//
// func (t *LivePlug) AfterFind(tx *gorm.DB) error {
// 	return serializer.HandleAfterFind(t)
// }
//
// func (t *Plug) BeforeCreate(tx *gorm.DB) error {
// 	t.Id = utils.GenerateUUID()
// 	return nil
// }
//
// func (t *Plug) BeforeSave(tx *gorm.DB) error {
// 	return serializer.HandleBeforeSave(t)
// }
//
// func (t *Plug) AfterFind(tx *gorm.DB) error {
// 	return serializer.HandleAfterFind(t)
// }
