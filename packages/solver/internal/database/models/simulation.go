package models

import (
	"math/big"
	"solver/internal/database/serializer"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

// SimulationRequest represents a simulation request in the database
type SimulationRequest struct {
	Id          string           `json:"id,omitempty" gorm:"primaryKey;type:text"`
	ReferenceId string           `json:"referenceId,omitempty" gorm:"type:text;index"`
	ChainId     uint64           `json:"chainId" gorm:"type:bigint"`
	From        common.Address   `json:"from" db_field:"FromStr" gorm:"-"`
	To          common.Address   `json:"to" db_field:"ToStr" gorm:"-"`
	Data        hexutil.Bytes    `json:"data,omitempty" db_field:"DataStr" gorm:"-"`
	Value       *big.Int         `json:"value,omitempty" db_field:"ValueStr" gorm:"-"`
	GasLimit    *uint64          `json:"gasLimit,omitempty" gorm:"type:bigint"`
	AccessList  types.AccessList `json:"accessList,omitempty" gorm:"type:jsonb"`
	ABI         string           `json:"abi,omitempty" gorm:"type:text"`

	// Database storage fields
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	FromStr   string         `json:"-" gorm:"column:from;type:text"`
	ToStr     string         `json:"-" gorm:"column:to;type:text"`
	DataStr   string         `json:"-" gorm:"column:data;type:text"`
	ValueStr  string         `json:"-" gorm:"column:value;type:text"`
}

// SimulationResponse represents a simulation response in the database
type SimulationResponse struct {
	Id                string               `json:"id" gorm:"primaryKey;type:text"`
	RequestId         string               `json:"requestId" gorm:"type:text"`
	SimulationRequest SimulationRequest    `json:"-" gorm:"foreignKey:RequestId;references:Id;-:migration"` // Adding this field to the struct with -:migration will prevent gorm from optimistically preloading the relationship
	GasUsed           uint64               `json:"gasUsed" gorm:"type:bigint"`
	Success           bool                 `json:"success"`
	Data              SimulationOutputData `json:"data" gorm:"type:jsonb"`
	ErrorMessage      string               `json:"errorMessage,omitempty" gorm:"type:text"`

	// Store the timestamps but do not expose them in the JSON response
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// OutputData represents the output data of a simulation
type SimulationOutputData struct {
	Raw     []byte      `json:"raw" gorm:"type:bytea"`
	Decoded interface{} `json:"decoded,omitempty" gorm:"type:jsonb"`
}

func (s *SimulationRequest) BeforeSave(tx *gorm.DB) error {
	return serializer.HandleBeforeSave(s)
}

func (s *SimulationRequest) AfterFind(tx *gorm.DB) error {
	return serializer.HandleAfterFind(s)
}
