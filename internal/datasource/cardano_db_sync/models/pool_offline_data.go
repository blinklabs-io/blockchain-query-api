package models

type PoolOfflineData struct {
	Id         int64  `gorm:"column:id"`
	PoolId     int64  `gorm:"column:pool_id"` // pool_hash(id)
	TickerName string `gorm:"column:ticker_name"`
	Hash       []byte `gorm:"column:hash"` // This is a "hash32type" column
	Json       jsonb  `gorm:"column:json"`
	Bytes      []byte `gorm:"column:bytes"`
	PmrId      int64  `gorm:"column:pmr_id"` // pool_metadata_ref(id)
}

// Override default pluralized table name
func (PoolOfflineData) TableName() string {
	return "pool_offline_data"
}
