package models

type PoolActiveStakeCache struct {
	PoolId      string `gorm:"column:pool_id"`
	EpochNumber int64  `gorm:"column:epoch_no"`
	Amount      uint64 `gorm:"coumn:amount"` // This is a "lovelace" column
}

// Override default pluralized table name
func (PoolActiveStakeCache) TableName() string {
	return "pool_active_stake_cache"
}
