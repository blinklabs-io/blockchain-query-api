package models

type Redeemer struct {
	Id         int64  `gorm:"column:id"`
	TxId       int64  `gorm:"column:tx_id"`      // tx(id)
	UnitMem    uint32 `gorm:"column:unit_mem"`   // This is a "word63type" column
	UnitSteps  uint32 `gorm:"column:unit_steps"` // This is a "word63type" column
	fee        uint64 `gorm:"column:fee"`        // This is a "lovelace" column
	purpose    string `gorm:"column:purpose"`    // This is a "scriptpurposetype" column
	index      uint32 `gorm:"column:index"`
	ScriptHash []byte `gorm:"column:script_hash"` // This is a "hash28type" column
	DatumId    int64  `gorm:"column:datum_id"`    // datum(id)
}

// Override default pluralized table name
func (Redeemer) TableName() string {
	return "redeemer"
}
