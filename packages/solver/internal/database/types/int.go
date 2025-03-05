package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math/big"
)

// BigInt wraps big.Int to implement SQL and JSON interfaces
type BigInt struct {
	*big.Int
}

// NewBigInt creates a new BigInt
func NewBigInt(x *big.Int) *BigInt {
	if x == nil {
		return nil
	}
	return &BigInt{x}
}

// Scan implements sql.Scanner interface
func (b *BigInt) Scan(value interface{}) error {
	if value == nil {
		b.Int = nil
		return nil
	}

	if b.Int == nil {
		b.Int = new(big.Int)
	}

	switch v := value.(type) {
	case int64:
		b.Int.SetInt64(v)
	case string:
		if v == "" {
			b.Int = nil
			return nil
		}
		_, ok := b.Int.SetString(v, 10)
		if !ok {
			return fmt.Errorf("failed to parse BigInt from string: %s", v)
		}
	case []byte:
		str := string(v)
		if str == "" {
			b.Int = nil
			return nil
		}
		_, ok := b.Int.SetString(str, 10)
		if !ok {
			return fmt.Errorf("failed to parse BigInt from bytes: %s", str)
		}
	default:
		return fmt.Errorf("unsupported Scan type for BigInt: %T", value)
	}

	return nil
}

// Value implements sql.Valuer interface
func (b *BigInt) Value() (driver.Value, error) {
	if b == nil || b.Int == nil {
		return nil, nil
	}
	return b.String(), nil
}

// MarshalJSON implements json.Marshaler interface
func (b *BigInt) MarshalJSON() ([]byte, error) {
	if b == nil || b.Int == nil {
		return []byte("null"), nil
	}
	return json.Marshal(b.String())
}

// UnmarshalJSON implements json.Unmarshaler interface
func (b *BigInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}

	if b.Int == nil {
		b.Int = new(big.Int)
	}

	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return fmt.Errorf("failed to unmarshal BigInt: %w", err)
	}

	if str == "" {
		return nil
	}

	_, ok := b.Int.SetString(str, 10)
	if !ok {
		return fmt.Errorf("failed to parse BigInt: %s", str)
	}

	return nil
}
