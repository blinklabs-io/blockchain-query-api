package models

type Reward struct {
	Id             int64  `gorm:"column:id"`
	AddrId         int64  `gorm:"column:addr_id"` // stake_address(id)
	Type           string `gorm:"column:type"`    // This is a "rewardtype" column
	Amount         uint64 `gorm:"column:amount"`  // This is a "lovelace" column
	EarnedEpoch    uint32 `gorm:"column:earned_epoch"`
	SpendableEpoch uint32 `gorm:"column:spendable_epoch"`
	PoolId         int64  `gorm:"column:pool_id"` // pool_hash(id)
}

// Override default pluralized table name
func (Reward) TableName() string {
	return "reward"
}
