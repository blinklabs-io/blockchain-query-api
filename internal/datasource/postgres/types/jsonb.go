package types

import (
	"database/sql/driver"
	"encoding/json"
)

// support jsonb column type
type Jsonb map[string]interface{}

func (j Jsonb) Value() (driver.Value, error) {
	retVal, err := json.Marshal(j)
	return string(retVal), err
}

func (j *Jsonb) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}
