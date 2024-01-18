package models

import "time"

//
// @Description
// @Author 代码小学生王木木
// @Date 2024/1/18 14:47
//

type ProductModel struct {
	BaseModel
	ProdName           string     `gorm:"type:varchar(255);column:prod_name;comment:商品名称" json:"prodName"`
	ShopID             int64      `gorm:"bigint;column:shop_id;comment:店铺id" json:"shopID"`
	OriPrice           float64    `gorm:"type:decimal(15,2);column:ori_price;comment:原价" json:"oriPrice"`
	Price              float64    `gorm:"type:decimal(15,2);column:price;comment:现价" json:"price"`
	Brief              string     `gorm:"type:varchar(500);column:brief;comment:简要描述,卖点等" json:"brief"`
	Content            string     `gorm:"type:text;column:content;comment:详细描述" json:"content"`
	Pic                string     `gorm:"type:varchar(255);column:pic;comment:商品主图" json:"pic"`
	Imgs               string     `gorm:"type:varchar(1000);column:imgs;comment:商品图片，以,分割" json:"imgs"`
	Status             int64      `gorm:"type:int;column:status;comment:默认是1，表示正常状态, -1表示删除, 0下架" json:"status"`
	CategoryID         int64      `gorm:"type:int;column:category_id;comment:商品分类" json:"categoryID"`
	SoldNum            int64      `gorm:"type:int;column:sold_num;comment:销量" json:"soldNum"`
	TotalStocks        int64      `gorm:"type:int;column:total_stocks;comment:总库存" json:"totalStocks"`
	DeliveryMode       string     `gorm:"type:int;column:delivery_mode;comment:配送方式json见TransportModeVO" json:"deliveryMode"`
	DeliveryTemplateID int64      `gorm:"type:json;column:delivery_template_id;comment:运费模板id" json:"deliveryTemplateID"`
	PutAwayTime        *time.Time `gorm:"type:datetime;column:putaway_time;comment:上架时间" json:"putAwayTime"`
	Version            int64      `gorm:"type:int;column:version;comment:版本 乐观锁" json:"version"`
}

func (ProductModel) TableName() string {
	return ""
}
