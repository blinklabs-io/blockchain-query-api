package models

import (
	"database/sql/driver"
	"encoding/json"
)

// support jsonb column type
type jsonb map[string]interface{}

func (j jsonb) Value() (driver.Value, error) {
	retVal, err := json.Marshal(j)
	return string(retVal), err
}

func (j *jsonb) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
