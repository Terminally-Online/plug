package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Inputs []map[string]any

func (is *Inputs) Scan(value any) error {
	if value == nil {
		*is = make(Inputs, 0)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value: invalid data type")
	}

	if err := json.Unmarshal(bytes, is); err != nil {
		return err
	}
	return nil
}

func (is Inputs) Value() (driver.Value, error) {
	if is == nil {
		return nil, nil
	}
	return json.Marshal(is)
}
