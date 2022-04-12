package models

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/postgres/types"
)

type Datum struct {
	Id    int64       `gorm:"column:id"`
	Hash  []byte      `gorm:"column:hash"`  // This is a "hash32type" column
	TxId  int64       `gorm:"column:tx_id"` // tx(id)
	Value types.Jsonb `gorm:"column:value"` // This is a "jsonb" column
}

// Override default pluralized table name
func (Datum) TableName() string {
	return "datum"
}
