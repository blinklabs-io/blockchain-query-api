package models

type Delegation struct {
	Id                int64  `gorm:"column:id"`
	AddrId            int64  `gorm:"column:addr_id"` // stake_address(id)
	CertIndex         int32  `gorm:"column:cert_index"`
	PoolHashId        int64  `gorm:"column:pool_hash_id"` // pool_hash(id)
	ActiveEpochNumber int64  `gorm:"column:active_epoch_no"`
	TxId              int64  `gorm:"column:tx_id"` // tx(id)
	SlotNumber        uint32 `gorm:"column:slot_no"`
	RedeemerId        int64  `gorm:"column:redeemer_id"` // redeemer(id)
}

// Override default pluralized table name
func (Delegation) TableName() string {
	return "delegation"
}
