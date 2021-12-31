package models

type TxOut struct {
	Id               int64  `gorm:"column:id"`
	TxId             int64  `gorm:"column:tx_id"` // tx(id)
	Index            string `gorm:"column:index"` // This is a "txindex" column
	Address          string `gorm:"column:address"`
	AddressRaw       []byte `gorm:"column:address_raw"`
	AddressHasScript bool   `gorm:"column:address_has_script"`
	PaymentCred      []byte `gorm:"column:payment_cred"`     // This is a "hash28type" column
	StakeAddressId   int64  `gorm:"column:stake_address_id"` // stake_address(id)
	Value            uint64 `gorm:"column:value"`            // This is a "lovelace" column
	DataHash         []byte `gorm:"column:data_hash"`        // This is a "hash32type" column
}

// Override default pluralized table name
func (TxOut) TableName() string {
	return "tx_out"
}
