package table

import "gorm.io/gorm"

// 商铺基础
type Shop struct {
	gorm.Model
	Address   string `gorm:"column:addr"` // 地址
	Tels      string // 电话
	ImageUrls string // 图片
	Introduce string // 介绍
}

// 商铺职工
type CShopStaff struct {
	// gorm.Model
	StaffIds string `gorm:"index:idx_ids,unique"` // 员工
}
