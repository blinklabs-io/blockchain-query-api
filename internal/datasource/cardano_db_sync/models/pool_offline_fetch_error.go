package models

import (
	"time"
)

type PoolOfflineFetchError struct {
	Id         int64      `gorm:"column:id"`
	PoolId     int64      `gorm:"column:pool_id"` // pool_hash(id)
	FetchTime  *time.Time `gorm:"column:fetch_time"`
	PmrId      int64      `gorm:"column:pmr_id"` // pool_metadata_ref(id)
	FetchError string     `gorm:"column:fetch_error"`
	RetryCount uint32     `gorm:"column:retry_count"`
}

// Override default pluralized table name
func (PoolOfflineFetchError) TableName() string {
	return "pool_offline_fetch_error"
}
