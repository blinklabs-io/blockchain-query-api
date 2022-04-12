package models

import (
	"github.com/cloudstruct/blockchain-query-api/internal/datasource/postgres/types"
)

type TxMetadata struct {
	Id    int64       `gorm:"column:id"`
	Key   string      `gorm:"column:key"`  // This is a "word64type" column
	Json  types.Jsonb `gorm:"column:json"` // This is a "jsonb" type
	Bytes []byte      `gorm:"column:bytes"`
	TxId  int64       `gorm:"column:tx_id"` // tx(id)
}

// Override default pluralized table name
func (TxMetadata) TableName() string {
	return "tx_metadata"
}
