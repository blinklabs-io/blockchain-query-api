package models

type MaTxMint struct {
	Id       int64 `gorm:"column:id"`
	Quantity int64 `gorm:"column:quantity"` // This is a "int65type" column
	TxId     int64 `gorm:"column:tx_id"`    // tx(id)
	ident    int64 `gorm:"column:ident"`    // multi_asset(id)
}

// Override default table name
func (MaTxMint) TableName() string {
	return "ma_tx_mint"
}
