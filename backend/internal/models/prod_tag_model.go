package models

type ProdTagModel struct {
	BaseModel
	Title     string `gorm:"type:;comment:分组标题;column:title" json:"title"`
	ShopId    int64  `gorm:"type:;comment:店铺Id;column:shopId" json:"shopId"`
	Status    int    `gorm:"type:;comment:状态(1为正常,0为删除);column:status" json:"status"`
	IsDefault int    `gorm:"type:;comment:默认类型(0:商家自定义,1:系统默认);column:isDefault" json:"isDefault"`
	ProdCount int64  `gorm:"type:;comment:商品数量;column:prodCount" json:"prodCount"`
	Style     int    `gorm:"type:;comment:列表样式(0:一列一个,1:一列两个,2:一列三个);column:style" json:"style"`
	Seq       int    `gorm:"type:;comment:排序;column:seq" json:"seq"`
}

func (ProdTagModel) TableName() string {
	return "m4_prod_tag"
}
