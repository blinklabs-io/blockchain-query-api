package models

type CollateralTxIn struct {
	Id         int64  `gorm:"column:id"`
	TxInId     int64  `gorm:"column:tx_in_id"`     // tx(id)
	TxOutId    int64  `gorm:"column:tx_out_id"`    // tx(id)
	TxOutIndex string `gorm:"column:tx_out_index"` // This is a "txindex" column
}

// Override default pluralized table name
func (CollateralTxIn) TableName() string {
	return "collateral_tx_in"
}
