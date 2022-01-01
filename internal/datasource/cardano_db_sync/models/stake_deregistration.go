package models

type StakeDeregistration struct {
	Id          int64 `gorm:"column:id"`
	AddrId      int64 `gorm:"column:addr_id"` // stake_address(id)
	CertIndex   int32 `gorm:"column:cert_index"`
	EpochNumber int64 `gorm:"column:epoch_no"`
	TxId        int64 `gorm:"column:tx_id"`       // tx(id)
	RedeemerId  int64 `gorm:"column:redeemer_id"` // redeemer(id)
}

// Override default pluralized table name
func (StakeDeregistration) TableName() string {
	return "stake_deregistration"
}
