package models

type EpochParam struct {
	Id                  int64   `gorm:"column:id"`
	EpochNumber         uint32  `gorm:"column:epoch_num"`
	MinFeeA             uint32  `gorm:"column:min_fee_a"`
	MinFeeB             uint32  `gorm:"column:min_fee_b"`
	MaxBlockSize        uint32  `gorm:"column:max_block_size"`
	MaxTxSize           uint32  `gorm:"column:max_tx_size"`
	MaxBhSize           uint32  `gorm:"column:max_bh_size"`
	KeyDeposit          uint64  `gorm:"column:key_deposit"`  // This is a "lovelace" column
	PoolDeposit         uint64  `gorm:"column:pool_deposit"` // This is a "lovelace" column
	MaxEpoch            uint32  `gorm:"column:max_epoch"`
	OptimalPoolCount    uint32  `gorm:"column:optimal_pool_count"`
	Influence           float32 `gorm:"column:influence"`
	MonetaryExpandRate  float32 `gorm:"column:monetary_expand_rate"`
	TreasuryGrowthRate  float32 `gorm:"column:treasury_growth_rate"`
	Decentralisation    float32 `gorm:"column:decentralisation"`
	Entropy             []byte  `gorm:"column:entropy"` // This is a "hash32type" column
	ProtocolMajor       uint32  `gorm:"column:protocol_major"`
	ProtocolMinor       uint32  `gorm:"column:protocol_minor"`
	MinUtxoValue        uint64  `gorm:"column:min_utxo_value"`      // This is a "lovelace" column
	MinPoolCost         uint64  `gorm:"column:min_pool_cost"`       // This is a "lovelace" column
	Nonce               []byte  `gorm:"column:nonce"`               // This is a "hash32type" column
	CoinsPerUtxoWord    uint64  `gorm:"column:coins_per_utxo_word"` // This is a "lovelace" column
	CostModelId         int64   `gorm:"column:cost_model_id"`       // cost_model(id)
	PriceMem            float32 `gorm:"column:price_mem"`
	PriceStep           float32 `gorm:"column:price_step"`
	MaxTxExMem          string  `gorm:"column:max_tx_ex_mem"`      // This is a "word64type" column
	MaxTxExSteps        string  `gorm:"column:max_tx_ex_steps"`    // This is a "word64type" column
	MaxBlockExMem       string  `gorm:"column:max_block_ex_mem"`   // This is a "word64type" column
	MaxBlockExSteps     string  `gorm:"column:max_block_ex_steps"` // This is a "word64type" column
	MaxValSize          string  `gorm:"column:max_val_size"`       // This is a "word64type" column
	CollateralPercent   uint32  `gorm:"column:collateral_percent"`
	MaxCollateralInputs uint32  `gorm:"column:max_collateral_inputs"`
	BlockId             int64   `gorm:"column:block_id"` // block(id)
}

// Override default pluralized table name
func (EpochParam) TableName() string {
	return "epoch_param"
}
