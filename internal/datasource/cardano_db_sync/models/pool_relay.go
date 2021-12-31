package models

type PoolRelay struct {
	Id         int64  `gorm:"column:id"`
	UpdateId   int64  `gorm:"column:update_id"` // pool_update(id)
	Ipv4       string `gorm:"column:ipv4"`
	Ipv6       string `gorm:"column:ipv6"`
	DnsName    string `gorm:"column:dns_name"`
	DnsSrvName string `gorm:"column:dns_srv_name"`
	Port       int32  `gorm:"column:port"`
}

// Override default pluralized table name
func (PoolRelay) TableName() string {
	return "pool_relay"
}
