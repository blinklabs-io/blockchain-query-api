package models

import (
	"database/sql/driver"
	"encoding/json"
)

type CostModel struct {
	Id      int64 `gorm:"column:id"`
	Costs   jsonb `gorm:"column:costs"`    // This is a "jsonb" column
	BlockId int64 `gorm:"column:block_id"` // block(id)
}

// Override default pluralized table name
func (CostModel) TableName() string {
	return "cost_model"
}

// TODO: this goes elsewhere
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
