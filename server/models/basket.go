package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:46
//

type TzBasket struct {
	BasketId           uint64     `xorm:"not null pk autoincr comment('主键') UNSIGNED BIGINT" json:"basketId"`
	ShopId             int64      `xorm:"not null comment('店铺ID') index unique(uk_user_shop_sku) BIGINT" json:"shopId"`
	ProdId             uint64     `xorm:"not null default 0 comment('产品ID') UNSIGNED BIGINT" json:"prodId"`
	SkuId              uint64     `xorm:"not null default 0 comment('SkuID') unique(uk_user_shop_sku) UNSIGNED BIGINT" json:"skuId"`
	UserId             string     `xorm:"not null comment('用户ID') unique(uk_user_shop_sku) index VARCHAR(50)" json:"userId"`
	BasketCount        int        `xorm:"not null default 0 comment('购物车产品个数') INT" json:"basketCount"`
	BasketDate         *time.Time `xorm:"not null comment('购物时间') DATETIME" json:"basketDate"`
	DiscountId         int64      `xorm:"comment('满减活动ID') BIGINT" json:"discountId"`
	DistributionCardNo string     `xorm:"comment('分销推广人卡号') VARCHAR(36)" json:"distributionCardNo"`
}
