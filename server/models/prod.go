package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 15:54
//

type TzProd struct {
	ProdId             uint64     `xorm:"not null pk autoincr comment('产品ID') UNSIGNED BIGINT" json:"prodId"`
	ProdName           string     `xorm:"not null default '' comment('商品名称') VARCHAR(300)" json:"prodName"`
	ShopId             int64      `xorm:"comment('店铺id') index BIGINT" json:"shopId"`
	OriPrice           string     `xorm:"default 0.00 comment('原价') DECIMAL(15,2)" json:"oriPrice"`
	Price              string     `xorm:"comment('现价') DECIMAL(15,2)" json:"price"`
	Brief              string     `xorm:"default '' comment('简要描述,卖点等') VARCHAR(500)" json:"brief"`
	Content            string     `xorm:"comment('详细描述') TEXT" json:"content"`
	Pic                string     `xorm:"comment('商品主图') VARCHAR(255)" json:"pic"`
	Imgs               string     `xorm:"comment('商品图片，以,分割') VARCHAR(1000)" json:"imgs"`
	Status             int        `xorm:"default 0 comment('默认是1，表示正常状态, -1表示删除, 0下架') INT" json:"status"`
	CategoryId         uint64     `xorm:"comment('商品分类') UNSIGNED BIGINT" json:"categoryId"`
	SoldNum            int        `xorm:"comment('销量') INT" json:"soldNum"`
	TotalStocks        int        `xorm:"default 0 comment('总库存') INT" json:"totalStocks"`
	DeliveryMode       string     `xorm:"comment('配送方式json见TransportModeVO') JSON" json:"deliveryMode"`
	DeliveryTemplateId int64      `xorm:"comment('运费模板id') BIGINT" json:"deliveryTemplateId"`
	CreateTime         *time.Time `xorm:"comment('录入时间') DATETIME" json:"createTime"`
	UpdateTime         *time.Time `xorm:"comment('修改时间') DATETIME" json:"updateTime"`
	PutawayTime        *time.Time `xorm:"comment('上架时间') DATETIME" json:"putawayTime"`
	Version            int        `xorm:"comment('版本 乐观锁') INT" json:"version"`
}
