package model_business

import "gorm.io/gorm"

// 商铺基础
type Shop struct {
	gorm.Model
	Address   string // 地址
	Tels      string // 电话
	ImageUrls string // 图片
	Introduce string // 介绍
}

// 商铺职工
type ShopStaff struct {
	gorm.Model
	StaffIds string // 员工
}
