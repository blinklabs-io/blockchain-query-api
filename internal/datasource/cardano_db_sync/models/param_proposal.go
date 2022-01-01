package models

type ParamProposal struct {
	Id                  int64   `gorm:"column:id"`
	EpochNumber         uint32  `gorm:"column:epoch_num"`
	Key                 []byte  `gorm:"column:key"`                // This is a "hash28type" column
	MinFeeA             string  `gorm:"column:min_fee_a"`          // This is a "word64type" column
	MinFeeB             string  `gorm:"column:min_fee_b"`          // This is a "word64type" column
	MaxBlockSize        string  `gorm:"column:max_block_size"`     // This is a "word64type" column
	MaxTxSize           string  `gorm:"column:max_tx_size"`        // This is a "word64type" column
	MaxBhSize           string  `gorm:"column:max_bh_size"`        // This is a "word64type" column
	KeyDeposit          uint64  `gorm:"column:key_deposit"`        // This is a "lovelace" column
	PoolDeposit         uint64  `gorm:"column:pool_deposit"`       // This is a "lovelace" column
	MaxEpoch            uint32  `gorm:"column:max_epoch"`          // This is a "word64type" column
	OptimalPoolCount    uint32  `gorm:"column:optimal_pool_count"` // This is a "word64type" column
	Influence           float32 `gorm:"column:influence"`
	MonetaryExpandRate  float32 `gorm:"column:monetary_expand_rate"`
	TreasuryGrowthRate  float32 `gorm:"column:treasury_growth_rate"`
	Decentralisation    float32 `gorm:"column:decentralisation"`
	Entropy             []byte  `gorm:"column:entropy"` // This is a "hash32type" column
	ProtocolMajor       uint32  `gorm:"column:protocol_major"`
	ProtocolMinor       uint32  `gorm:"column:protocol_minor"`
	MinUtxoValue        uint64  `gorm:"column:min_utxo_value"`      // This is a "lovelace" column
	MinPoolCost         uint64  `gorm:"column:min_pool_cost"`       // This is a "lovelace" column
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
	RegisteredTxId      int64   `gorm:"column:registered_tx_id"` // tx(id)
}

// Override default pluralized table name
func (ParamProposal) TableName() string {
	return "param_proposal"
}
