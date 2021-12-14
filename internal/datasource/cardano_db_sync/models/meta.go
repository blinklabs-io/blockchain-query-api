package models

import (
	"time"
)

type Meta struct {
	Id        int64      `gorm:"column:id"`
	StartTime *time.Time `gorm:"column:start_time"`
	Network   string     `gorm:"column:network_name"`
}

// Override default pluralized table name
func (Meta) TableName() string {
	return "meta"
}
