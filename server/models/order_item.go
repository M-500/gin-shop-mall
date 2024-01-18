package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:51
//

type TzOrderItem struct {
	OrderItemId        uint64     `xorm:"not null pk autoincr comment('订单项ID') UNSIGNED BIGINT" json:"orderItemId"`
	ShopId             int64      `xorm:"not null comment('店铺id') BIGINT" json:"shopId"`
	OrderNumber        string     `xorm:"not null comment('订单order_number') index VARCHAR(50)" json:"orderNumber"`
	ProdId             uint64     `xorm:"not null comment('产品ID') UNSIGNED BIGINT" json:"prodId"`
	SkuId              uint64     `xorm:"not null comment('产品SkuID') UNSIGNED BIGINT" json:"skuId"`
	ProdCount          int        `xorm:"not null default 0 comment('购物车产品个数') INT" json:"prodCount"`
	ProdName           string     `xorm:"not null default '' comment('产品名称') VARCHAR(120)" json:"prodName"`
	SkuName            string     `xorm:"comment('sku名称') VARCHAR(120)" json:"skuName"`
	Pic                string     `xorm:"not null default '' comment('产品主图片路径') VARCHAR(255)" json:"pic"`
	Price              string     `xorm:"not null comment('产品价格') DECIMAL(15,2)" json:"price"`
	UserId             string     `xorm:"not null default '' comment('用户Id') VARCHAR(36)" json:"userId"`
	ProductTotalAmount string     `xorm:"not null comment('商品总金额') DECIMAL(15,2)" json:"productTotalAmount"`
	RecTime            *time.Time `xorm:"not null comment('购物时间') DATETIME" json:"recTime"`
	CommSts            int        `xorm:"not null default 0 comment('评论状态： 0 未评价  1 已评价') INT" json:"commSts"`
	DistributionCardNo string     `xorm:"comment('推广员使用的推销卡号') VARCHAR(36)" json:"distributionCardNo"`
	BasketDate         *time.Time `xorm:"comment('加入购物车时间') DATETIME" json:"basketDate"`
}
