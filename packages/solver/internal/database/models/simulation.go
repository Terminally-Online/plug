package models

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm"
)

// SimulationRequest represents a simulation request in the database
type SimulationRequest struct {
	gorm.Model
	Id         string           `json:"id,omitempty" gorm:"type:text"`
	ChainId    uint64           `json:"chainId" gorm:"type:bigint"`
	From       common.Address   `json:"from" gorm:"type:bytea"`
	To         common.Address   `json:"to" gorm:"type:bytea"`
	Data       hexutil.Bytes    `json:"data,omitempty" gorm:"type:bytea"`
	GasLimit   *uint64          `json:"gasLimit,omitempty" gorm:"type:bigint"`
	Value      *big.Int         `json:"value,omitempty" gorm:"type:text"`
	AccessList types.AccessList `json:"accessList,omitempty" gorm:"type:jsonb"`
	ABI        string           `json:"abi,omitempty" gorm:"type:text"`
}

// Implement database serialization for big.Int
func (s *SimulationRequest) AfterFind(tx *gorm.DB) error {
	if s.Value != nil {
		// Convert stored string back to big.Int
		value := new(big.Int)
		_, ok := value.SetString(s.Value.String(), 10)
		if !ok {
			return fmt.Errorf("failed to parse Value as big.Int: %s", s.Value.String())
		}
		s.Value = value
	}
	return nil
}

// SimulationResponse represents a simulation response in the database
type SimulationResponse struct {
	gorm.Model
	Id                string               `json:"id" gorm:"type:text"`
	RequestId         string               `json:"requestId" gorm:"type:text"`
	SimulationRequest SimulationRequest    `json:"request" gorm:"foreignKey:RequestId;references:Id"`
	GasUsed           uint64               `json:"gasUsed" gorm:"type:bigint"`
	Success           bool                 `json:"success"`
	Data              SimulationOutputData `json:"data" gorm:"type:jsonb"`
	ErrorMessage      string               `json:"errorMessage,omitempty" gorm:"type:text"`
}

// OutputData represents the output data of a simulation
type SimulationOutputData struct {
	Raw     []byte      `json:"raw" gorm:"type:bytea"`
	Decoded interface{} `json:"decoded,omitempty" gorm:"type:jsonb"`
}

func (s *SimulationRequest) BeforeSave(tx *gorm.DB) error {
	// Value is already a *big.Int, and its String() method will be used
	// automatically by GORM when saving to the database
	return nil
}
