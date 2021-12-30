package models

type Tx struct {
	Id            int64  `gorm:"column:id"`
	Hash          []byte `gorm:"column:hash"`     // This is a "hash32type" column
	BlockId       int64  `gorm:"column:block_id"` // block(id)
	BlockIndex    uint32 `gorm:"column:block_index"`
	OutSum        uint64 `gorm:"column:out_sum"` // This is a "lovelace" column
	Fee           uint64 `gorm:"column:fee"`     // This is a "lovelace" column
	Deposit       int64  `gorm:"column:deposit"`
	Size          uint32 `gorm:"column:size"`
	InvalidBefore string `gorm:"column:invalid_before"` // This is a "word64type" column
	InvalidAfter  string `gorm:"column:invalid_after"`  // This is a "word64type" column
	ValidContract bool   `gorm:"column:valid_contract"`
	ScriptSize    uint32 `gorm:"column:script_size"`
}

// Override default pluralized table name
func (Tx) TableName() string {
	return "tx"
}
