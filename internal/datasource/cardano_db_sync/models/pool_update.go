package models

type PoolUpdate struct {
	Id                int64   `gorm:"column:id"`
	HashId            int64   `gorm:"column:hash_id"` // pool_hash(id)
	CertIndex         int32   `gorm:"column:cert_index"`
	VrfKeyHash        []byte  `gorm:"column:vrf_key_hash"` // This is a "hash32type" column
	Pledge            uint64  `gorm:"column:pledge"`       // This is a "lovelace" column
	RewardAddr        string  `gorm:"column:reward_addr"`  // This is a "addr29type" column
	ActiveEpochNumber int64   `gorm:"column:active_epoch_no"`
	MetaId            int64   `gorm:"column:meta_id"` // pool_metadata_ref(id)
	Margin            float32 `gorm:"column:margin"`
	FixedCost         uint64  `gorm:"column:fixed_cost"`       // This is a "lovelace" column
	RegisterdTxId     int64   `gorm:"column:registered_tx_id"` // tx(id)
}

// Override default pluralized table name
func (PoolUpdate) TableName() string {
	return "pool_update"
}
