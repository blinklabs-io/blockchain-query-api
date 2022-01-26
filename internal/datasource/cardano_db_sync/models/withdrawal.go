package models

type Withdrawal struct {
	Id         int64  `gorm:"column:id"`
	AddrId     int64  `gorm:"column:addr_id"`     // stake_address(id)
	Amount     uint64 `gorm:"column:amount"`      // This is a "lovelace" column
	RedeemerId int64  `gorm:"column:redeemer_id"` // redeemer(id)
	TxId       int64  `gorm:"column:tx_id"`       // tx(id)
}

// Override default pluralized table name
func (Withdrawal) TableName() string {
	return "withdrawal"
}
