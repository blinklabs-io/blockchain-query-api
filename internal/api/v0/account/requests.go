package account

import (
	"gorm.io/gorm"
	"strings"
)

// Return a query to get the ID of a given address
func AddressIdQuery(db *gorm.DB, address string) *gorm.DB {
	var addrIdQuery *gorm.DB
	if strings.HasPrefix(address, "stake") {
		// Shelley stake address
		addrIdQuery = db.Select("id").Limit(1).Where("view = ?", address).Table("stake_address")
	} else {
		// Payment address
		addrIdQuery = db.Select("stake_address_id").Limit(1).Where("address = ?", address).Table("tx_out")
	}
	return addrIdQuery
}

// Returns a query to get the latest withdrawal epoch given a transaction ID
func LatestWithdrawalEpochQuery(db *gorm.DB, i uint64) *gorm.DB {
	return db.Table("block").
		Select("epoch_no").
		Joins("INNER JOIN tx ON tx.block_id = block.id").
		Where("tx.id = ?", i).
		Limit(1)
}

// Returns a query to get the latest withdrawal transaction given an address ID
func LatestWithdrawalTxQuery(db *gorm.DB, i uint64) *gorm.DB {
	return db.Table("withdrawal").
		Select("tx_id").
		Where("addr_id = ?", i).
		Order("tx_id desc").
		Limit(1)
}

// Used as part of HandleGetAccountAssets to filter out assets which have been transferred away
func TxOutQuery(db *gorm.DB) *gorm.DB {
	return db.Table("tx_out").
		Select("tx_out.id").
		Joins("INNER JOIN tx_in ON tx_out.id = tx_in.tx_out_id AND tx_out.index = tx_in.tx_out_index").
		Where("txo.id = tx_out.id")
}
