package models

type AdaPots struct {
	Id          int64  `gorm:"column:id"`
	SlotNumber  uint32 `gorm:"column:slot_no"`
	EpochNumber uint32 `gorm:"column:epoch_no"`
	Treasury    uint64 `gorm:"column:treasury"` // This is a "lovelace" column
	Reserves    uint64 `gorm:"column:reserves"` // This is a "lovelace" column
	Rewards     uint64 `gorm:"column:rewards"`  // This is a "lovelace" column
	Utxo        uint64 `gorm:"column:utxo"`     // This is a "lovelace" column
	Deposits    uint64 `gorm:"column:deposits"` // This is a "lovelace" column
	Fees        uint64 `gorm:"column:fees"`     // This is a "lovelace" column
	BlockId     int64  `gorm:"column:block_id"` // block(id)
}

// Override default table name
func (AdaPots) TableName() string {
	return "ada_pots"
}
