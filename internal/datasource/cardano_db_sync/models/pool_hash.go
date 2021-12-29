package models

type PoolHash struct {
	Id      int64  `gorm:"column:id"`
	HashRaw string `gorm:"column:hash_raw"` // This is a "hash28type"
	View    string `gorm:"column:view"`
}

// Override default pluralized table name
func (PoolHash) TableName() string {
	return "pool_hash"
}
