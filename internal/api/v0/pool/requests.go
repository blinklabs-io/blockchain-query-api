package pool

import (
	"gorm.io/gorm"
)

// Return a query to get the ID of a given pool address
func PoolIdQuery(db *gorm.DB, address string) *gorm.DB {
	return db.Select("id").Where("view = ?", address).Table("pool_hash")
}
