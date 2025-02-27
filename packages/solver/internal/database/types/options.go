package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Options map[string]interface{}

func (o *Options) Scan(value any) error {
	if value == nil {
		*o = make(Options, 0)
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal JSONB value: invalid data type")
	}

	if err := json.Unmarshal(bytes, o); err != nil {
		return err
	}
	return nil
}

func (o Options) Value() (driver.Value, error) {
	if o == nil {
		return nil, nil
	}
	return json.Marshal(o)
}
