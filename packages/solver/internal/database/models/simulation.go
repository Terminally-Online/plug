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
	GasLimit    *uint64          `json:"gasLimit,omitempty" gorm:"type:bigint"`
	valueStr    string           `gorm:"column:value"`
	Value       *big.Int         `json:"value,omitempty" gorm:"-"`
	AccessList  types.AccessList `json:"accessList,omitempty" gorm:"type:jsonb"`
	ABI         string           `json:"abi,omitempty" gorm:"type:text"`

	// These values are serialized back from strings when the struct is loaded from the database
	From common.Address `json:"from" gorm:"-"`
	To   common.Address `json:"to" gorm:"-"`
	Data hexutil.Bytes  `json:"data,omitempty" gorm:"-"`

	// Store these values as their string representations in the database.
	FromStr string `json:"-" gorm:"column:from;type:text"`
	ToStr   string `json:"-" gorm:"column:to;type:text"`
	DataStr string `json:"-" gorm:"column:data;type:text"`

	// Store the timestamps but do not expose them in the JSON response
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// SimulationResponse represents a simulation response in the database
type SimulationResponse struct {
	Id        string `json:"id" gorm:"type:text"`
	RequestId string `json:"requestId" gorm:"type:text"`
	// Adding this field to the struct with -:migration will prevent gorm from optimistically preloading the relationship
	SimulationRequest SimulationRequest    `json:"-" gorm:"foreignKey:RequestId;references:Id;-:migration"`
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

// Implement database serialization for big.Int
func (s *SimulationRequest) AfterFind(tx *gorm.DB) error {
	// Handle existing Value conversion
	if s.valueStr != "" {
		value := new(big.Int)
		_, ok := value.SetString(s.valueStr, 10)
		if !ok {
			return fmt.Errorf("failed to parse Value as big.Int: %s", s.valueStr)
		}
		s.Value = value
	}

	// Convert strings back to addresses and data
	s.From = common.HexToAddress(s.FromStr)
	s.To = common.HexToAddress(s.ToStr)
	if s.DataStr != "" {
		data, err := hexutil.Decode(s.DataStr)
		if err != nil {
			return fmt.Errorf("failed to decode data: %v", err)
		}
		s.Data = data
	}

	return nil
}

func (s *SimulationRequest) BeforeSave(tx *gorm.DB) error {
	// Handle existing Value conversion
	if s.Value != nil {
		s.valueStr = s.Value.String()
	} else {
		s.valueStr = "0"
	}

	// Convert addresses and data to strings
	s.FromStr = s.From.Hex()
	s.ToStr = s.To.Hex()
	if len(s.Data) > 0 {
		s.DataStr = hexutil.Encode(s.Data)
	}

	return nil
}
