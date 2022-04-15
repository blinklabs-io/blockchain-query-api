package pool

import (
	"fmt"
	"gorm.io/gorm"
)

// Return a query to get the ID of a given pool address
func PoolIdQuery(db *gorm.DB, address string) *gorm.DB {
	return db.Select("id").Where("view = ?", address).Table("pool_hash")
}

func CurrentEpochQuery(db *gorm.DB) *gorm.DB {
	return db.Select("no").Order("no DESC").Limit(1).Table("epoch")
}

func SaturationLimitQuery(db *gorm.DB, epoch uint64) *gorm.DB {
	return db.Select("FLOOR(?)",
		db.Raw("supply::bigint / (?)",
			db.Select("p_optimal_pool_count").
				Where(fmt.Sprint("epoch_no = %s", epoch)).
				Table("grest.epoch_info_cache")))
}
