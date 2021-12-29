package models

import (
	"time"
)

type Epoch struct {
	Id          int64      `gorm:"column:id"`
	OutSum      uint64     `gorm:"column:out_sum"` // This type may not be large enough. The DB column uses NUMERIC(38, 0)
	Fees        uint64     `gorm:"column:fees"`    // This is a "lovelace" column
	TxCount     uint32     `gorm:"column:tx_count"`
	BlockCount  uint32     `gorm:"column:blk_count"`
	EpochNumber uint32     `gorm:"column:no"`
	StartTime   *time.Time `gorm:"column:start_time"`
	EndTime     *time.Time `gorm:"column:end_time"`
}

// Override default pluralized table name
func (Epoch) TableName() string {
	return "epoch"
}
