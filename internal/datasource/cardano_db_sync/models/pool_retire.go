package models

type PoolRetire struct {
	Id            int64  `gorm:"column:id"`
	HashId        int64  `gorm:"column:hash_id"` // pool_hash(id)
	CertIndex     int32  `gorm:"column:cert_index"`
	AnnouncedTxId int64  `gorm:"column:announced_tx_id"` // tx(id)
	RetiringEpoch uint64 `gorm:"column:retiring_epoch"`
}

// Override default pluralized table name
func (PoolRetire) TableName() string {
	return "pool_retire"
}
