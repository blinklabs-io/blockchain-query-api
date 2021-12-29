package models

type PoolMetadataRef struct {
	Id            int64  `gorm:"column:id"`
	HashId        int64  `gorm:"column:hash_id"` // pool_hash(id)
	Url           string `gorm:"column:url"`
	Hash          []byte `gorm:"column:hash"`             // This is a "hash32type" column
	RegisterdTxId int64  `gorm:"column:registered_tx_id"` // tx(id)
}

// Override default pluralized table name
func (PoolMetadataRef) TableName() string {
	return "pool_metadata_ref"
}
