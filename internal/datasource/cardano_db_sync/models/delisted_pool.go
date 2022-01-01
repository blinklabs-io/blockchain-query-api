package models

type DelistedPool struct {
	Id      int64  `gorm:"column:id"`
	HashRaw []byte `gorm:"column:hash_raw"` // This is a "hash28type" column
}

// Override default pluralized table name
func (DelistedPool) TableName() string {
	return "delisted_pool"
}
