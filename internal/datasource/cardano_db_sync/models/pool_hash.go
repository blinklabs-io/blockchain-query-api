package models

type PoolHash struct {
	Id      int64  `gorm:"column:id"`
	HashRaw []byte `gorm:"column:hash_raw"` // This is a "hash28type" column
	View    string `gorm:"column:view"`
}

// Override default pluralized table name
func (PoolHash) TableName() string {
	return "pool_hash"
}
