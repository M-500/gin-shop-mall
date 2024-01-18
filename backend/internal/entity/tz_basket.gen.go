// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameBasket = "tz_basket"

// Basket 购物车
type Basket struct {
	BasketID           int64     `gorm:"column:basket_id;primaryKey;autoIncrement:true;comment:主键" json:"basket_id"` // 主键
	ShopID             int64     `gorm:"column:shop_id;not null;comment:店铺ID" json:"shop_id"`                        // 店铺ID
	ProdID             int64     `gorm:"column:prod_id;not null;comment:产品ID" json:"prod_id"`                        // 产品ID
	SkuID              int64     `gorm:"column:sku_id;not null;comment:SkuID" json:"sku_id"`                         // SkuID
	UserID             string    `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                        // 用户ID
	BasketCount        int32     `gorm:"column:basket_count;not null;comment:购物车产品个数" json:"basket_count"`           // 购物车产品个数
	BasketDate         time.Time `gorm:"column:basket_date;not null;comment:购物时间" json:"basket_date"`                // 购物时间
	DiscountID         int64     `gorm:"column:discount_id;comment:满减活动ID" json:"discount_id"`                       // 满减活动ID
	DistributionCardNo string    `gorm:"column:distribution_card_no;comment:分销推广人卡号" json:"distribution_card_no"`    // 分销推广人卡号
}

// TableName Basket's table name
func (*Basket) TableName() string {
	return TableNameBasket
}