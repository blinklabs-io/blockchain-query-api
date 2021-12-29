package models

type StakeAddress struct {
	Id            int64  `gorm:"column:id"`
	HashRaw       []byte `gorm:"column:hash_raw"` // This is a "addr29type" column
	View          string `gorm:"column:view"`
	ScriptHash    []byte `gorm:"column:script_hash"`
	RegisterdTxId int64  `gorm:"column:registered_tx_id"` // tx(id)
}

// Override default pluralized table name
func (StakeAddress) TableName() string {
	return "stake_address"
}
