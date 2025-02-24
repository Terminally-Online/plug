package serializer

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gorm.io/gorm"
)

type DBSerializer interface {
	BeforeSave(tx *gorm.DB) error
	AfterFind(tx *gorm.DB) error
}

// Define supported types and their conversion rules
var typeConverters = map[string]struct {
	toStr   func(interface{}) string
	fromStr func(string) (interface{}, error)
}{
	"common.Address": {
		toStr: func(v interface{}) string {
			return v.(common.Address).Hex()
		},
		fromStr: func(s string) (interface{}, error) {
			return common.HexToAddress(s), nil
		},
	},
	"hexutil.Bytes": {
		toStr: func(v interface{}) string {
			return hexutil.Encode(v.(hexutil.Bytes))
		},
		fromStr: func(s string) (interface{}, error) {
			return hexutil.Decode(s)
		},
	},
	"*big.Int": {
		toStr: func(v interface{}) string {
			if v.(*big.Int) == nil {
				return "0"
			}
			return v.(*big.Int).String()
		},
		fromStr: func(s string) (interface{}, error) {
			n := new(big.Int)
			if s == "" || s == "0" {
				return n, nil
			}
			if _, ok := n.SetString(s, 10); !ok {
				return nil, fmt.Errorf("failed to parse BigInt: %s", s)
			}
			return n, nil
		},
	},
}

// HandleBeforeSave handles the conversion of fields to their string representation
func HandleBeforeSave(model interface{}) error {
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		dbField := field.Tag.Get("db_field")
		if dbField == "" {
			continue
		}

		fieldValue := modelValue.Field(i)
		strField := modelValue.FieldByName(dbField)

		if !strField.IsValid() || !strField.CanSet() {
			continue
		}

		converter, ok := typeConverters[fieldValue.Type().String()]
		if !ok {
			continue
		}

		if !fieldValue.IsZero() {
			strValue := converter.toStr(fieldValue.Interface())
			strField.SetString(strValue)
		}
	}
	return nil
}

// HandleAfterFind handles the conversion of string fields back to their original types
func HandleAfterFind(model interface{}) error {
	modelValue := reflect.ValueOf(model).Elem()
	modelType := modelValue.Type()

	for i := 0; i < modelType.NumField(); i++ {
		field := modelType.Field(i)
		dbField := field.Tag.Get("db_field")
		if dbField == "" {
			continue
		}

		fieldValue := modelValue.Field(i)
		strField := modelValue.FieldByName(dbField)

		if !strField.IsValid() {
			continue
		}

		converter, ok := typeConverters[fieldValue.Type().String()]
		if !ok {
			continue
		}

		strValue := strField.String()
		if strValue == "" {
			continue
		}

		value, err := converter.fromStr(strValue)
		if err != nil {
			return fmt.Errorf("failed to convert field %s: %v", field.Name, err)
		}

		fieldValue.Set(reflect.ValueOf(value))
	}
	return nil
}
