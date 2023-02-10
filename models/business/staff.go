package model_business

import "gorm.io/gorm"

// 职工基础
type Staff struct {
	gorm.Model
	Name         string // 名字
	Sex          int    // 性别
	IDCardNum    string // 身份证号 加密
	PermissionLv int    // 特权等级
}
