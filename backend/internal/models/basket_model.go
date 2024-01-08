package models

import (
	"time"
)

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/8 13:49
//

// Basket 模型
type BasketModel struct {
	BaseModel
	BasketID           uint       `gorm:"primaryKey;autoIncrement" json:"basket_id"`
	ShopID             uint       `gorm:"not null" json:"shop_id"`
	ProdID             uint       `gorm:"not null;default:0" json:"prod_id"`
	SKUID              uint       `gorm:"not null;default:0" json:"sku_id"`
	UserID             string     `gorm:"not null" json:"user_id"`
	BasketCount        int        `gorm:"not null;default:0" json:"basket_count"`
	BasketDate         *time.Time `gorm:"not null" json:"basket_date"`
	DiscountID         uint       `gorm:"default:null" json:"discount_id"`
	DistributionCardNo string     `gorm:"default:null" json:"distribution_card_no"`
}

// TableName 指定表名
func (BasketModel) TableName() string {
	return "tz_basket"
}
