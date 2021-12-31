package models

type TxIn struct {
	Id         int64  `gorm:"column:id"`
	TxInId     int64  `gorm:"column:tx_in_id"`     // tx(id)
	TxOutId    int64  `gorm:"column:tx_out_id"`    // tx(id)
	TxOutIndex string `gorm:"column:tx_out_index"` // This is a "txindex" column
	RedeemerId int64  `gorm:"column:redeemer_id"`  // redeemer(id)
}

// Override default pluralized table name
func (TxIn) TableName() string {
	return "tx_in"
}
