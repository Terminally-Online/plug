package models

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

// SimulationRequest represents a simulation request in the database
type SimulationRequest struct {
	Id          string           `json:"id,omitempty" gorm:"type:text"`
	ReferenceId string           `json:"referenceId,omitempty" gorm:"type:text;index"`
	ChainId     uint64           `json:"chainId" gorm:"type:bigint"`
	From        common.Address   `json:"from" gorm:"type:bytea"`
	To          common.Address   `json:"to" gorm:"type:bytea"`
	Data        hexutil.Bytes    `json:"data,omitempty" gorm:"type:bytea"`
	GasLimit    *uint64          `json:"gasLimit,omitempty" gorm:"type:bigint"`
	valueStr    string           `gorm:"column:value"`
	Value       *big.Int         `json:"value,omitempty" gorm:"-"`
	AccessList  types.AccessList `json:"accessList,omitempty" gorm:"type:jsonb"`
	ABI         string           `json:"abi,omitempty" gorm:"type:text"`
	CreatedAt   time.Time        `json:"-"`
	UpdatedAt   time.Time        `json:"-"`
	DeletedAt   gorm.DeletedAt   `json:"-" gorm:"index"`
}

// SimulationResponse represents a simulation response in the database
type SimulationResponse struct {
	Id                string               `json:"id" gorm:"type:text"`
	RequestId         string               `json:"requestId" gorm:"type:text"`
	SimulationRequest SimulationRequest    `json:"request" gorm:"foreignKey:RequestId;references:Id"`
	GasUsed           uint64               `json:"gasUsed" gorm:"type:bigint"`
	Success           bool                 `json:"success"`
	Data              SimulationOutputData `json:"data" gorm:"type:jsonb"`
	ErrorMessage      string               `json:"errorMessage,omitempty" gorm:"type:text"`
	CreatedAt         time.Time            `json:"-"`
	UpdatedAt         time.Time            `json:"-"`
	DeletedAt         gorm.DeletedAt       `json:"-" gorm:"index"`
}

// OutputData represents the output data of a simulation
type SimulationOutputData struct {
	Raw     []byte      `json:"raw" gorm:"type:bytea"`
	Decoded interface{} `json:"decoded,omitempty" gorm:"type:jsonb"`
}

// Implement database serialization for big.Int
func (s *SimulationRequest) AfterFind(tx *gorm.DB) error {
	if s.valueStr == "" {
		s.Value = new(big.Int)
		return nil
	}

	value := new(big.Int)
	_, ok := value.SetString(s.valueStr, 10)
	if !ok {
		return fmt.Errorf("failed to parse Value as big.Int: %s", s.valueStr)
	}
	s.Value = value
	return nil
}

func (s *SimulationRequest) BeforeSave(tx *gorm.DB) error {
	if s.Value != nil {
		s.valueStr = s.Value.String()
	} else {
		s.valueStr = "0"
	}
	return nil
}
