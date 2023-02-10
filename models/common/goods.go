package model_common

import "gorm.io/gorm"

// 商品类目
type GoodsCategory struct {
	gorm.Model
	Name     string // 商品类名
	ParentId int    // 父级类目ID
	ShopId   int    // 商铺ID
}

// 商品详情
type Goods struct {
	gorm.Model
	Name       string  // 商品类名
	Title      string  // 商品标题
	Introduce  string  // 商品介绍
	Remark     string  // 商品备注
	Quantity   int     // 商品储量
	Price      float64 // 商品价格
	Sales      int     // 商品销量
	ImageUrls  string  // 商品图片
	Status     int     // 商品状态
	CategoryId int     // 商品类目ID
	ShopId     int     // 商铺ID
}
