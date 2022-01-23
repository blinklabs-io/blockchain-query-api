package models

type MaTxOut struct {
	Id       int64 `gorm:"column:id"`
	Quantity int64 `gorm:"column:quantity"`  // This is a "word64type" column
	TxOutId  int64 `gorm:"column:tx_out_id"` // tx_out(id)
	Ident    int64 `gorm:"column:ident"`     // multi_asset(id)
}

// Override default table name
func (MaTxOut) TableName() string {
	return "ma_tx_out"
}
