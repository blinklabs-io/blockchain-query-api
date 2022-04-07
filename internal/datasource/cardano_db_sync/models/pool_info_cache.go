package models

import (
	"time"
)

type PoolInfoCache struct {
	Id                int64      `gorm:"column:id"`
	TxId              int64      `gorm:"column:tx_id"`
	TxHash            string     `gorm:"column:tx_hash"`
	BlockTime         *time.Time `gorm:"column:block_time"`
	PoolHashId        int64      `gorm:"column:pool_hash_id"`
	PoolIdBech32      string     `gorm:"column:pool_id_bech32"`
	PoolIdHex         string     `gorm:"column:pool_id_hex"`
	ActiveEpochNumber int64      `gorm:"column:active_epoch_no"`
	VrfHashKey        string     `gorm:"column:vrf_hash_key"`
	Margin            float32    `gorm:"column:margin"`
	FixedCost         uint64     `gorm:"column:fixed_cost"` // This is a "lovelace" column
	Pledge            uint64     `gorm:"column:pledge"`     // This is a "lovelace" column
	RewardAddress     string     `gorm:"column:reward_address"`
	Owners            string     `gorm:"column:owners"`
	Relays            jsonb      `gorm:"column:relays"`
	MetaId            int64      `gorm:"column:meta_id"`
	MetaUrl           string     `gorm:"column:meta_url"`
	MetaHash          string     `gorm:"column:meta_hash"`
	PoolStatus        string     `gorm:"column:pool_status"`
	RetiringEpoch     uint64     `gorm:"column:retiring_epoch"`
}

// Override default pluralized table name
func (PoolInfoCache) TableName() string {
	return "pool_info_cache"
}
