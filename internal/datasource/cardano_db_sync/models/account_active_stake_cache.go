package models

type AccountActiveStakeCache struct {
	StakeAddress string `gorm:"column:stake_address"`
	PoolId       string `gorm:"column:pool_id"`
	EpochNumber  int64  `gorm:"column:epoch_no"`
	Amount       uint64 `gorm:"column:amount"` // This is a "lovelace" column
}

// Override default pluralized table name
func (AccountActiveStakeCache) TableName() string {
	return "account_active_stake_cache"
}
