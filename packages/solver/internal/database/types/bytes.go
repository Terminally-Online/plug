package types

import (
	"database/sql/driver"
	"fmt"
)

// ByteArray is a wrapper type for []byte that implements sql.Scanner and driver.Valuer
type ByteArray []byte

func (b *ByteArray) Scan(value interface{}) error {
	if value == nil {
		*b = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		*b = make([]byte, len(v))
		copy(*b, v)
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into ByteArray", value)
	}
}

func (b ByteArray) Value() (driver.Value, error) {
	if b == nil {
		return nil, nil
	}
	return []byte(b), nil
}

// ByteArrayPtr is a wrapper type for *[]byte that implements sql.Scanner and driver.Valuer
type ByteArrayPtr struct {
	Bytes *[]byte
}

func (b *ByteArrayPtr) Scan(value interface{}) error {
	if value == nil {
		b.Bytes = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		bytes := make([]byte, len(v))
		copy(bytes, v)
		b.Bytes = &bytes
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into ByteArrayPtr", value)
	}
}

func (b ByteArrayPtr) Value() (driver.Value, error) {
	if b.Bytes == nil {
		return nil, nil
	}
	return []byte(*b.Bytes), nil
}
