package models

type StakeRegistration struct {
	Id          int64 `gorm:"column:id"`
	AddrId      int64 `gorm:"column:addr_id"` // stake_address(id)
	CertIndex   int32 `gorm:"column:cert_index"`
	EpochNumber int64 `gorm:"column:epoch_no"`
	TxId        int64 `gorm:"column:tx_id"` // tx(id)
}

// Override default pluralized table name
func (StakeRegistration) TableName() string {
	return "stake_registration"
}
