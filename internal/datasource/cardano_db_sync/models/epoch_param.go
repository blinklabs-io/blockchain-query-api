package models

type EpochParam struct {
	Id                  int64   `gorm:"column:id"`
	EpochNumber         uint32  `gorm:"column:epoch_num"`
	MinFeeA             uint32  `gorm:"column:min_fee_a"`
	MinFeeB             uint32  `gorm:"column:min_fee_b"`
	MaxBlockSize        uint32  `gorm:"column:max_block_size"`
	MaxTxSize           uint32  `gorm:"column:max_tx_size"`
	MaxBhSize           uint32  `gorm:"column:max_bh_size"`
	KeyDeposit          uint64  `gorm:"column:key_deposit"`
	PoolDeposit         uint64  `gorm:"column:pool_deposit"`
	MaxEpoch            uint32  `gorm:"column:max_epoch"`
	OptimalPoolCount    uint32  `gorm:"column:optimal_pool_count"`
	Influence           float32 `gorm:"column:influence"`
	MonetaryExpandRate  float32 `gorm:"column:monetary_expand_rate"`
	TreasuryGrowthRate  float32 `gorm:"column:treasury_growth_rate"`
	Decentralisation    float32 `gorm:"column:decentralisation"`
	Entropy             string  `gorm:"column:entropy"`
	ProtocolMajor       uint32  `gorm:"column:protocol_major"`
	ProtocolMinor       uint32  `gorm:"column:protocol_minor"`
	MinUtxoValue        uint32  `gorm:"column:min_utxo_value"`
	MinPoolCost         uint32  `gorm:"column:min_pool_cost"`
	Nonce               string  `gorm:"column:nonce"`
	CoinsPerUtxoWord    uint32  `gorm:"column:coins_per_utxo_word"`
	CostModelId         int64   `gorm:"column:cost_model_id"` // cost_model(id)
	PriceMem            float32 `gorm:"column:price_mem"`
	PriceStep           float32 `gorm:"column:price_step"`
	MaxTxExMem          string  `gorm:"column:max_tx_ex_mem"`
	MaxTxExSteps        string  `gorm:"column:max_tx_ex_steps"`
	MaxBlockExMem       string  `gorm:"column:max_block_ex_mem"`
	MaxBlockExSteps     string  `gorm:"column:max_block_ex_steps"`
	MaxValSize          string  `gorm:"column:max_val_size"`
	CollateralPercent   uint32  `gorm:"column:collateral_percent"`
	MaxCollateralInputs uint32  `gorm:"column:max_collateral_inputs"`
	BlockId             int64   `gorm:column:block_id"` // block(id)
}

// Override default pluralized table name
func (EpochParam) TableName() string {
	return "epoch_param"
}
