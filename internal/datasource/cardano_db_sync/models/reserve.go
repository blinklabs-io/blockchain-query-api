package models

type Reserve struct {
	Id        int64  `gorm:"column:id"`
	AddrId    int64  `gorm:"column:addr_id"` // stake_address(id)
	CertIndex int32  `gorm:"column:cert_index"`
	Amount    uint64 `gorm:"column:amount"` // This is a "int65type" column
	TxId      int64  `gorm:"column:tx_id"`  // tx(id)
}

// Override default table name
func (Reserve) TableName() string {
	return "reserve"
}
