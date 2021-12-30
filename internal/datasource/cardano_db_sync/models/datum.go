package models

type Datum struct {
	Id    int64  `gorm:"column:id"`
	Hash  []byte `gorm:"column:hash"`  // This is a "hash32type" column
	TxId  int64  `gorm:"column:tx_id"` // tx(id)
	Value jsonb  `gorm:"column:value"` // This is a "jsonb" column
}

// Override default pluralized table name
func (Datum) TableName() string {
	return "datum"
}
