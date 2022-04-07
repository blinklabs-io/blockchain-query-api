package models

type EpochActiveStakeCache struct {
	EpochNumber int64  `gorm:"column:epoch_no"`
	Amount      uint64 `gorm:"coumn:amount"` // This is a "lovelace" column
}

// Override default pluralized table name
func (EpochActiveStakeCache) TableName() string {
	return "epoch_active_stake_cache"
}
