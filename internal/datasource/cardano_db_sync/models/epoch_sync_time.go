package models

type EpochSyncTime struct {
	Id          int64  `gorm:"column:id"`
	EpochNumber int64  `gorm:"column:no"`
	Seconds     uint64 `gorm:"column:seconds"` // This is a "word63type" column
	State       string `gorm:"column:state"`   // This is a "syncstatetype" column
}

// Override default pluralized table name
func (EpochSyncTime) TableName() string {
	return "epoch_sync_time"
}
