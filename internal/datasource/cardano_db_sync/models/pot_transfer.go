package models

type PotTransfer struct {
	Id        int64  `gorm:"column:id"`
	CertIndex int32  `gorm:"column:cert_index"`
	Treasury  uint64 `gorm:"column:treasury"` // This is a "int65type" column
	Reserves  uint64 `gorm:"column:reserves"` // This is a "int65type" column
	TxId      int64  `gorm:"column:tx_id"`    // tx(id)
}

// Override default table name
func (PotTransfer) TableName() string {
	return "pot_transfer"
}
