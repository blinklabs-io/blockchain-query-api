package models

type AssetRegistryCache struct {
	AssetPolicy string `gorm:"column:asset_policy"`
	AssetName   string `gorm:"column:asset_name"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Ticker      string `gorm:"column:ticker"`
	Url         string `gorm:"column:url"`
	Logo        string `gorm:"column:logo"`
	Decimals    int    `gorm:"column:decimals"`
}

// Override default pluralized table name
func (AssetRegistryCache) TableName() string {
	return "asset_registry_cache"
}
