package utils

import (
	"encoding/json"
)

// StructToMap converts a struct to a map using JSON as an intermediate format
func StructToMap(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	
	return result, nil
}

// MapToStruct converts a map to a struct using JSON as an intermediate format
func MapToStruct(input map[string]interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}
	
	return json.Unmarshal(data, output)
} 