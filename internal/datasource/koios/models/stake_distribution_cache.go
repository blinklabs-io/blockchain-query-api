package models

type StakeDistributionCache struct {
	StakeAddress     string  `gorm:"column:stake_address"`
	PoolId           string  `gorm:"column:pool_id"`
	TotalBalance     float32 `gorm:"column:total_balance"`
	Utxo             float32 `gorm:"column:utxo"`
	Rewards          float32 `gorm:"column:rewards"`
	Withdrawals      float32 `gorm:"column:withdrawals"`
	RewardsAvailable float32 `gorm:"column:rewards_available"`
}

// Override default pluralized table name
func (StakeDistributionCache) TableName() string {
	return "stake_distribution_cache"
}
