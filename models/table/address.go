package table

import (
	"gorm.io/gorm"
)

// 地址
type Address struct {
	gorm.Model
	Name      string // 收货人姓名
	Tel       string // 手机号
	Province  string // 省份
	City      string // 城市
	District  string // 区（县）
	Detail    string // 详细地址
	IsDefault int    // 是否默认
}
