package models

type CostModel struct {
	Id      int64 `gorm:"column:id"`
	Costs   jsonb `gorm:"column:costs"`    // This is a "jsonb" column
	BlockId int64 `gorm:"column:block_id"` // block(id)
}

// Override default pluralized table name
func (CostModel) TableName() string {
	return "cost_model"
}
