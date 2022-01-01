package models

type Script struct {
	Id             int64  `gorm:"column:id"`
	TxId           int64  `gorm:"column:tx_id"` // tx(id)
	Hash           []byte `gorm:"column:hash"`  // This is a "hash28type" column
	Type           string `gorm:"column:type"`  // This is a "scripttype" column
	Json           jsonb  `gorm:"column:json"`  // This is a "jsonb" type
	Bytes          []byte `gorm:"column:bytes"`
	SerialisedSize uint32 `gorm:"column:serialised_size"`
}

// Override default pluralized table name
func (Script) TableName() string {
	return "script"
}
