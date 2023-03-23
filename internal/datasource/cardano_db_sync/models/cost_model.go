package models

import (
	"github.com/blinklabs-io/blockchain-query-api/internal/datasource/postgres/types"
)

type CostModel struct {
	Id      int64       `gorm:"column:id"`
	Costs   types.Jsonb `gorm:"column:costs"`    // This is a "jsonb" column
	BlockId int64       `gorm:"column:block_id"` // block(id)
}

// Override default pluralized table name
func (CostModel) TableName() string {
	return "cost_model"
}
