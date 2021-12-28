package models

type EpochStake struct {
	Id          int64  `gorm:"column:id"`
	AddrId      int64  `gorm:"column:addr_id"` // stake_address(id)
	PoolId      int64  `gorm:"column:pool_id"` // pool_hash(id)
	Amount      uint32 `gorm:"column:amount"`
	EpochNumber uint32 `gorm:"column:epoch_no"`
}

// Override default table name
func (EpochStake) TableName() string {
	return "epoch_stake"
}
