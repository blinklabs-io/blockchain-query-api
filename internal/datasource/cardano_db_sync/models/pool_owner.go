package models

type PoolOwner struct {
	Id            int64 `gorm:"column:id"`
	AddrId        int64 `gorm:"column:addr_id"`          // stake_address(id)
	PoolHashId    int64 `gorm:"column:pool_hash_id"`     // pool_hash(id)
	RegisterdTxId int64 `gorm:"column:registered_tx_id"` // tx(id)
}

// Override default pluralized table name
func (PoolOwner) TableName() string {
	return "pool_owner"
}
