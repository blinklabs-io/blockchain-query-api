package models

type ReservedPoolTicker struct {
	Id       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	PoolHash string `gorm:"column:pool_hash"`
}

// Override default pluralized table name
func (ReservedPoolTicker) TableName() string {
	return "reserved_pool_ticker"
}
