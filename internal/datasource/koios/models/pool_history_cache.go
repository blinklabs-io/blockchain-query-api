package models

type PoolHistoryCache struct {
	PoolId             string  `gorm:"column:pool_id"`
	EpochNumber        uint64  `gorm:"column:epoch_no"`
	ActiveStake        uint64  `gorm:"column:active_stake"` // This is a "lovelace" column
	ActiveStakePercent float32 `gorm:"column:active_stake_pct"`
	SaturationPct      float32 `gorm:"column:saturation_pct"`
	BlockCount         uint64  `gorm:"column:block_cnt"`
	DelegatorCount     uint64  `gorm:"column:delegator_cnt"`
	PoolFeeVariable    float32 `gorm:"column:pool_fee_variable"`
	PoolFeeFixed       uint64  `gorm:"column:pool_fee_fixed"` // This is a "lovelace" column
	PoolFees           float32 `gorm:"column:pool_fees"`
	DelegRewards       float32 `gorm:"column:deleg_rewards"`
	EpochRos           float32 `gorm:"column:epoch_ros"`
}

// Override default pluralized table name
func (PoolHistoryCache) TableName() string {
	return "pool_history_cache"
}
