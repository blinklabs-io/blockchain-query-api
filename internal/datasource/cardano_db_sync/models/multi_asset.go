package models

type MultiAsset struct {
	Id          int64  `gorm:"column:id"`
	Policy      []byte `gorm:"column:policy"` // This is a "hash28type" column
	Name        []byte `gorm:"column:name"`   // This is a "asset32type" column
	Fingerprint string `gorm:"column:fingerprint"`
}

// Override default table name
func (MultiAsset) TableName() string {
	return "multi_asset"
}
