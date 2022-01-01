package models

type EpochRewardTotalReceived struct {
	Id          int64  `gorm:"column:id"`
	EarnedEpoch uint32 `gorm:"column:earned_epoch"`
	Amount      uint64 `gorm:"column:amount"` // This is a "lovelace" column
}

// Override default pluralized table name
func (EpochRewardTotalReceived) TableName() string {
	return "epoch_reward_total_received"
}
