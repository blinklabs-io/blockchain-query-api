package models

import (
	"time"
)

type EpochInfoCache struct {
	EpochNumber          uint64     `gorm:"column:epoch_no"`
	IOutSum              string     `gorm:"column:i_out_sum"` // This is a "word128type" column
	IFees                uint64     `gorm:"column:i_fees"`    // This is a "lovelace" column
	ITxCount             uint64     `gorm:"column:i_tx_count"`
	IBlockCount          uint64     `gorm:"column:i_blk_count"`
	IFirstBlockTime      *time.Time `gorm:"column:i_first_block_time"`
	ILastBlockTime       *time.Time `gorm:"column:i_last_block_time"`
	PMinFeeA             uint64     `gorm:"column:p_min_fee_a"`
	PMinFeeB             uint64     `gorm:"column:p_min_fee_b"`
	PMaxBlockSize        uint64     `gorm:"column:p_max_block_size"`
	PMaxTxSize           uint64     `gorm:"column:p_max_tx_size"`
	PMaxBhSize           uint64     `gorm:"column:p_max_bh_size"`
	PKeyDeposit          uint64     `gorm:"column:p_key_deposit"`  // This is a "lovelace" column
	PPoolDeposit         uint64     `gorm:"column:p_pool_deposit"` // This is a "lovelace" column
	PMaxEpoch            uint64     `gorm:"column:p_max_epoch"`
	POptimalPoolCount    uint64     `gorm:"column:p_optimal_pool_count"`
	PInfluence           float32    `gorm:"column:p_influence"`
	PMonetaryExpandRate  float32    `gorm:"column:p_monetary_expand_rate"`
	PTreasuryGrowthRate  float32    `gorm:"column:p_treasure_growth_rate"`
	PDecentralisation    float32    `gorm:"column:p_decentralisation"`
	PEntropy             string     `gorm:"column:p_entropy"`
	PProtocolMajor       uint64     `gorm:"column:p_protocol_major"`
	PProtocolMinor       uint64     `gorm:"column:p_protocol_minor"`
	PMinUtxoValue        uint64     `gorm:"column:p_min_utxo_value"` // This is a "lovelace" column
	PMinPoolCost         uint64     `gorm:"column:p_min_pool_cost"`  // This is a "lovelace" column
	PNonce               string     `gorm:"column:p_nonce"`
	PBlockHash           string     `gorm:"column:p_block_hash"`
	PCostModels          string     `gorm:"column:p_cost_models"`
	PPriceMem            float32    `gorm:"column:p_price_mem"`
	PPriceStep           float32    `gorm:"column:p_price_step"`
	PMaxTxExMem          string     `gorm:"column:p_max_tx_ex_mem"`      // This is a "word64type" column
	PMaxTxExStep         string     `gorm:"column:p_max_tx_ex_step"`     // This is a "word64type" column
	PMaxBlockExMem       string     `gorm:"column:p_max_block_ex_mem"`   // This is a "word64type" column
	PMaxBlockExSteps     string     `gorm:"column:p_max_block_ex_steps"` // This is a "word64type" column
	PMaxValSize          string     `gorm:"column:p_max_val_size"`       // This is a "word64type" column
	PCollateralPercent   uint64     `gorm:"column:p_collateral_percent"`
	PMaxCollateralInputs uint64     `gorm:"column:p_max_collateral_inputs"`
	PCoinsPerUtxoWord    uint64     `gorm:"column:p_coins_per_utxo_word"`
}

// Override default pluralized table name
func (EpochInfoCache) TableName() string {
	return "epoch_info_cache"
}
