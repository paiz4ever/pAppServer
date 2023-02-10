package model_common

import "gorm.io/gorm"

// 订单
type Order struct {
	gorm.Model
	UserId        int     // 用户ID
	GoodsIdCounts string  // 商品ID和数目
	TotalPrice    float64 // 总价格
	AddressId     int     // 地址ID
	Status        int     // 状态
	ShopId        int     // 商铺ID
}
