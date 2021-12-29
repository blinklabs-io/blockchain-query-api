package models

type SlotLeader struct {
	Id          int64  `gorm:"column:id"`
	Hash        string `gorm:"column:hash"`
	PoolId      int64  `gorm:"column:pool_id"` // pool_hash(id)
	Description string `gorm:"column:description"`
}

// Override default table name
func (SlotLeader) TableName() string {
	return "slot_leader"
}
