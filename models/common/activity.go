package model_common

import "gorm.io/gorm"

// 商品活动
type Activity struct {
	gorm.Model
	Name      string // 活动名称
	ImageUrls string // 活动图片
	Introduce string // 活动介绍
	GoodsIds  string // 活动相关商品ID
	Type      int    // 活动类型
	Status    int    //  活动状态
}
